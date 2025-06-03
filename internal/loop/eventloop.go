// Package loop implements the HTML5 event loop specification for DOMulator.
// This provides the foundation for asynchronous JavaScript execution with
// proper task and microtask scheduling.
package loop

import (
	"sync"
	"time"

	"github.com/dop251/goja"
	"github.com/dpemmons/DOMulator/internal/dom"
)

// EventLoop implements the HTML5 event loop specification with main thread architecture.
// This ensures compatibility with Goja's single-threaded design while providing
// deterministic behavior for testing scenarios.
type EventLoop struct {
	// Core queues (priority order: microtasks > tasks > animation > idle)
	microtaskQueue *MicrotaskQueue // Highest priority
	taskQueue      *TaskQueue      // Normal priority
	animationQueue *AnimationQueue // Frame-locked
	idleQueue      *IdleQueue      // Lowest priority

	// Runtime integration
	vm       *goja.Runtime
	document *dom.Document

	// State management (single-threaded, no atomic operations needed)
	running          bool
	blocked          bool
	renderingEnabled bool

	// Timing control
	frameRate        time.Duration // 16.67ms for 60fps
	lastFrameTime    time.Time
	performanceStart time.Time

	// Testing support for deterministic execution
	virtualTime  time.Time
	timeOffset   time.Duration
	usingVirtual bool

	// Synchronization (protects state, not VM access)
	mu       sync.RWMutex
	shutdown chan struct{}
	done     chan struct{}

	// Metrics for monitoring and debugging
	taskCount         int64
	microtaskCount    int64
	frameCount        int64
	idleCallbackCount int64
}

// EventLoopOptions configures the event loop behavior
type EventLoopOptions struct {
	FrameRate        time.Duration // Default: 16.67ms (60fps)
	RenderingEnabled bool          // Default: true
	VirtualTime      bool          // Default: false (for testing)
}

// New creates a new event loop with the specified JavaScript runtime and DOM document.
// The event loop uses main thread architecture for Goja VM compatibility.
func New(vm *goja.Runtime, document *dom.Document, opts *EventLoopOptions) *EventLoop {
	if opts == nil {
		opts = &EventLoopOptions{
			FrameRate:        16667 * time.Microsecond, // 60fps
			RenderingEnabled: true,
			VirtualTime:      false,
		}
	}

	el := &EventLoop{
		microtaskQueue:   NewMicrotaskQueue(),
		taskQueue:        NewTaskQueue(),
		animationQueue:   NewAnimationQueue(),
		idleQueue:        NewIdleQueue(),
		vm:               vm,
		document:         document,
		frameRate:        opts.FrameRate,
		renderingEnabled: opts.RenderingEnabled,
		performanceStart: time.Now(),
		usingVirtual:     opts.VirtualTime,
		shutdown:         make(chan struct{}),
		done:             make(chan struct{}),
	}

	if el.usingVirtual {
		el.virtualTime = time.Now()
	}

	return el
}

// Start begins the event loop processing. This method blocks until the event loop
// is stopped. For main thread architecture, this should be called from the main goroutine.
func (el *EventLoop) Start() error {
	el.mu.Lock()
	if el.running {
		el.mu.Unlock()
		return ErrAlreadyRunning
	}
	el.running = true
	el.mu.Unlock()

	defer func() {
		close(el.done)
		el.mu.Lock()
		el.running = false
		el.mu.Unlock()
	}()

	// Main event loop - runs on calling goroutine (main thread)
	for {
		select {
		case <-el.shutdown:
			return nil
		default:
			// Process one iteration of the event loop
			el.processEventLoopIteration()
		}
	}
}

// Stop gracefully shuts down the event loop
func (el *EventLoop) Stop() {
	el.mu.Lock()
	defer el.mu.Unlock()

	if !el.running {
		return
	}

	close(el.shutdown)
}

// Wait waits for the event loop to complete shutdown
func (el *EventLoop) Wait() {
	<-el.done
}

// IsRunning returns true if the event loop is currently running
func (el *EventLoop) IsRunning() bool {
	el.mu.RLock()
	defer el.mu.RUnlock()
	return el.running
}

// processEventLoopIteration implements one iteration of the HTML5 event loop algorithm
func (el *EventLoop) processEventLoopIteration() {
	currentTime := el.getCurrentTime()

	// 1. Select and execute one task from the task queue
	task := el.taskQueue.SelectTask(currentTime)
	if task != nil {
		el.executeTask(task)
		el.incrementTaskCount()
	}

	// 2. Process ALL microtasks (microtask checkpoint)
	el.processMicrotasks()

	// 3. Update rendering if needed
	if el.shouldRender() {
		el.performRenderSteps()
		el.incrementFrameCount()
	}

	// 4. Process idle callbacks if time permits
	if el.hasIdleTime() {
		el.processIdleCallbacks()
	}
}

// executeTask executes a single task with proper error handling
func (el *EventLoop) executeTask(task *Task) {
	defer func() {
		if r := recover(); r != nil {
			// Log error but continue event loop processing
			// In a real browser, this would be reported to the console
			// TODO: Integrate with console API for error reporting
		}
	}()

	// Execute the task callback - goja.Callable is invoked directly
	_, err := task.Callback(goja.Undefined(), task.Args...)
	if err != nil {
		// Handle JavaScript execution errors
		// TODO: Integrate with error reporting system
	}
}

// processMicrotasks processes all pending microtasks until the queue is empty
func (el *EventLoop) processMicrotasks() {
	// Prevent infinite loops with a reasonable limit
	const maxMicrotasks = 1000
	processed := 0

	for processed < maxMicrotasks {
		microtask := el.microtaskQueue.Dequeue()
		if microtask == nil {
			break // Queue is empty
		}

		el.executeMicrotask(microtask)
		el.incrementMicrotaskCount()
		processed++

		// New microtasks may have been queued during execution
		// so we continue the loop
	}

	if processed >= maxMicrotasks {
		// TODO: Report microtask starvation warning
	}
}

// executeMicrotask executes a single microtask with proper error handling
func (el *EventLoop) executeMicrotask(microtask *Microtask) {
	defer func() {
		if r := recover(); r != nil {
			// Log error but continue microtask processing
			// TODO: Integrate with console API for error reporting
		}
	}()

	// Execute the microtask callback - goja.Callable is invoked directly
	_, err := microtask.Callback(goja.Undefined(), microtask.Args...)
	if err != nil {
		// Handle JavaScript execution errors
		// TODO: Integrate with error reporting system
	}
}

// shouldRender determines if rendering should occur this iteration
func (el *EventLoop) shouldRender() bool {
	if !el.renderingEnabled {
		return false
	}

	now := el.getCurrentTime()
	return now.Sub(el.lastFrameTime) >= el.frameRate
}

// performRenderSteps executes the rendering process simulation
func (el *EventLoop) performRenderSteps() {
	now := el.getCurrentTime()

	// 1. Process animation frame callbacks
	el.processAnimationFrameCallbacks(now)

	// 2. Update computed styles (simulated)
	el.updateComputedStyles()

	// 3. Layout calculations (simulated)
	el.performLayout()

	// 4. Paint (simulated - for testing framework)
	el.performPaint()

	// Update frame timing
	el.lastFrameTime = now
}

// processAnimationFrameCallbacks executes all scheduled animation frame callbacks
func (el *EventLoop) processAnimationFrameCallbacks(now time.Time) {
	callbacks := el.animationQueue.GetCallbacks()

	// Calculate high-resolution timestamp (DOMHighResTimeStamp)
	timestamp := float64(now.Sub(el.performanceStart).Nanoseconds()) / 1e6

	for _, callback := range callbacks {
		// Execute callback with timestamp directly
		_, err := callback.Callback(goja.Undefined(), el.vm.ToValue(timestamp))
		if err != nil {
			// Handle JavaScript execution errors
			// TODO: Integrate with error reporting system
		}
	}

	// Clear processed callbacks
	el.animationQueue.Clear()
}

// updateComputedStyles simulates CSS style recalculation
func (el *EventLoop) updateComputedStyles() {
	// Placeholder for style recalculation simulation
	// In a real browser, this would trigger CSS style computation
}

// performLayout simulates layout calculations
func (el *EventLoop) performLayout() {
	// Placeholder for layout calculation simulation
	// In a real browser, this would trigger layout/reflow
}

// performPaint simulates painting operations
func (el *EventLoop) performPaint() {
	// Placeholder for paint simulation
	// In a real browser, this would trigger painting operations
}

// hasIdleTime determines if there's time available for idle callbacks
func (el *EventLoop) hasIdleTime() bool {
	// Simple heuristic: if no tasks or microtasks are pending
	return el.taskQueue.IsEmpty() && el.microtaskQueue.IsEmpty()
}

// processIdleCallbacks executes idle callbacks if time permits
func (el *EventLoop) processIdleCallbacks() {
	// Calculate available idle time
	deadline := el.getCurrentTime().Add(time.Millisecond) // 1ms deadline

	for el.getCurrentTime().Before(deadline) {
		callback := el.idleQueue.GetNextCallback()
		if callback == nil {
			break // No more idle callbacks
		}

		el.executeIdleCallback(callback, deadline)
		el.incrementIdleCallbackCount()
	}
}

// executeIdleCallback executes an idle callback with deadline information
func (el *EventLoop) executeIdleCallback(callback *IdleCallback, deadline time.Time) {
	defer func() {
		if r := recover(); r != nil {
			// Log error but continue processing
			// TODO: Integrate with console API for error reporting
		}
	}()

	// Create IdleDeadline object for the callback
	timeRemaining := deadline.Sub(el.getCurrentTime())
	idleDeadline := el.createIdleDeadline(timeRemaining)

	// Execute the callback - goja.Callable is invoked directly
	_, err := callback.Callback(goja.Undefined(), idleDeadline)
	if err != nil {
		// Handle JavaScript execution errors
		// TODO: Integrate with error reporting system
	}
}

// createIdleDeadline creates an IdleDeadline object for idle callbacks
func (el *EventLoop) createIdleDeadline(timeRemaining time.Duration) goja.Value {
	idleDeadline := el.vm.NewObject()
	idleDeadline.Set("timeRemaining", func() float64 {
		return float64(timeRemaining.Nanoseconds()) / 1e6 // Convert to milliseconds
	})
	idleDeadline.Set("didTimeout", false) // TODO: Implement timeout tracking
	return idleDeadline
}

// getCurrentTime returns the current time, using virtual time if enabled
func (el *EventLoop) getCurrentTime() time.Time {
	if el.usingVirtual {
		return el.virtualTime.Add(el.timeOffset)
	}
	return time.Now()
}

// QueueMicrotask queues a microtask for execution at the next microtask checkpoint
func (el *EventLoop) QueueMicrotask(callback func(), source MicrotaskSource) int64 {
	// Create a wrapper that converts the Go function to a goja.Callable
	gojaCallback := goja.Callable(func(this goja.Value, args ...goja.Value) (goja.Value, error) {
		callback()
		return goja.Undefined(), nil
	})

	return el.microtaskQueue.Queue(source, gojaCallback, nil)
}

// ScheduleTask schedules a task with the specified delay
func (el *EventLoop) ScheduleTask(taskType TaskType, source TaskSource, callback goja.Callable, args []goja.Value, delay time.Duration) int64 {
	return el.taskQueue.ScheduleTask(taskType, source, callback, args, delay, el.getCurrentTime())
}

// RequestAnimationFrame schedules an animation frame callback
func (el *EventLoop) RequestAnimationFrame(callback goja.Callable) int64 {
	return el.animationQueue.Schedule(callback)
}

// CancelAnimationFrame cancels an animation frame callback
func (el *EventLoop) CancelAnimationFrame(id int64) {
	el.animationQueue.Cancel(id)
}

// RequestIdleCallback schedules an idle callback
func (el *EventLoop) RequestIdleCallback(callback goja.Callable, timeout time.Duration) int64 {
	return el.idleQueue.Schedule(callback, timeout)
}

// CancelIdleCallback cancels an idle callback
func (el *EventLoop) CancelIdleCallback(id int64) {
	el.idleQueue.Cancel(id)
}

// Testing utilities for deterministic async testing

// AdvanceTime advances virtual time by the specified duration (testing only)
func (el *EventLoop) AdvanceTime(duration time.Duration) {
	if !el.usingVirtual {
		return
	}

	el.mu.Lock()
	el.timeOffset += duration
	el.mu.Unlock()

	// Process any tasks that should execute in this time window
	el.processTasksUntilTime(el.getCurrentTime())
}

// processTasksUntilTime processes all tasks scheduled up to the specified time
func (el *EventLoop) processTasksUntilTime(targetTime time.Time) {
	for {
		task := el.taskQueue.SelectTaskByTime(targetTime)
		if task == nil {
			break
		}
		el.executeTask(task)
		el.incrementTaskCount()
		el.processMicrotasks() // Process microtasks after each task
	}
}

// ProcessMicrotasks forces processing of all pending microtasks (testing only)
func (el *EventLoop) ProcessMicrotasks() {
	el.processMicrotasks()
}

// ProcessNextTask processes exactly one task from the queue (testing only)
func (el *EventLoop) ProcessNextTask() *Task {
	task := el.taskQueue.SelectTask(el.getCurrentTime())
	if task != nil {
		el.executeTask(task)
		el.incrementTaskCount()
		el.processMicrotasks() // Always process microtasks after a task
	}
	return task
}

// ProcessAnimationFrame triggers one animation frame (testing only)
func (el *EventLoop) ProcessAnimationFrame() {
	if el.animationQueue.HasCallbacks() {
		el.processAnimationFrameCallbacks(el.getCurrentTime())
		el.incrementFrameCount()
	}
}

// GetQueueCounts returns the current count of items in each queue (testing/debugging)
func (el *EventLoop) GetQueueCounts() QueueStats {
	return QueueStats{
		Tasks:           el.taskQueue.Count(),
		Microtasks:      el.microtaskQueue.Count(),
		AnimationFrames: el.animationQueue.Count(),
		IdleCallbacks:   el.idleQueue.Count(),
	}
}

// QueueStats provides queue count information
type QueueStats struct {
	Tasks           int
	Microtasks      int
	AnimationFrames int
	IdleCallbacks   int
}

// GetMetrics returns performance metrics for the event loop
func (el *EventLoop) GetMetrics() EventLoopMetrics {
	el.mu.RLock()
	defer el.mu.RUnlock()

	return EventLoopMetrics{
		TasksProcessed:        el.taskCount,
		MicrotasksProcessed:   el.microtaskCount,
		FramesRendered:        el.frameCount,
		IdleCallbacksExecuted: el.idleCallbackCount,
		IsRunning:             el.running,
		UsingVirtualTime:      el.usingVirtual,
	}
}

// EventLoopMetrics provides performance and status information
type EventLoopMetrics struct {
	TasksProcessed        int64
	MicrotasksProcessed   int64
	FramesRendered        int64
	IdleCallbacksExecuted int64
	IsRunning             bool
	UsingVirtualTime      bool
}

// Metric increment helpers (thread-safe)
func (el *EventLoop) incrementTaskCount() {
	el.mu.Lock()
	el.taskCount++
	el.mu.Unlock()
}

func (el *EventLoop) incrementMicrotaskCount() {
	el.mu.Lock()
	el.microtaskCount++
	el.mu.Unlock()
}

func (el *EventLoop) incrementFrameCount() {
	el.mu.Lock()
	el.frameCount++
	el.mu.Unlock()
}

func (el *EventLoop) incrementIdleCallbackCount() {
	el.mu.Lock()
	el.idleCallbackCount++
	el.mu.Unlock()
}
