package loop

import (
	"sync"
	"testing"
	"time"

	"github.com/dop251/goja"
	"github.com/dpemmons/DOMulator/internal/dom"
)

func TestEventLoopBasics(t *testing.T) {
	// Create VM and document for testing
	vm := goja.New()
	doc := &dom.Document{} // Simplified for test

	// Create event loop with virtual time for deterministic testing
	opts := &EventLoopOptions{
		FrameRate:        16667 * time.Microsecond, // 60fps
		RenderingEnabled: true,
		VirtualTime:      true,
	}
	el := New(vm, doc, opts)

	// Test basic creation
	if el == nil {
		t.Fatal("Failed to create event loop")
	}

	// Test initial state
	if el.IsRunning() {
		t.Error("Event loop should not be running initially")
	}

	// Test queue counts
	counts := el.GetQueueCounts()
	if counts.Tasks != 0 || counts.Microtasks != 0 || counts.AnimationFrames != 0 || counts.IdleCallbacks != 0 {
		t.Error("All queues should be empty initially")
	}
}

func TestTaskScheduling(t *testing.T) {
	vm := goja.New()
	doc := &dom.Document{}
	opts := &EventLoopOptions{VirtualTime: true}
	el := New(vm, doc, opts)

	executed := false
	callback := goja.Callable(func(this goja.Value, args ...goja.Value) (goja.Value, error) {
		executed = true
		return goja.Undefined(), nil
	})

	// Schedule a task
	taskID := el.ScheduleTask(TimerTask, TimerTaskSource, callback, nil, 0)
	if taskID == 0 {
		t.Error("Task ID should not be zero")
	}

	// Verify task was queued
	counts := el.GetQueueCounts()
	if counts.Tasks != 1 {
		t.Errorf("Expected 1 task, got %d", counts.Tasks)
	}

	// Process the task
	processedTask := el.ProcessNextTask()
	if processedTask == nil {
		t.Error("Expected to process a task")
	}

	if !executed {
		t.Error("Task callback should have been executed")
	}

	// Verify task queue is empty
	counts = el.GetQueueCounts()
	if counts.Tasks != 0 {
		t.Errorf("Expected 0 tasks after processing, got %d", counts.Tasks)
	}
}

func TestMicrotaskProcessing(t *testing.T) {
	vm := goja.New()
	doc := &dom.Document{}
	opts := &EventLoopOptions{VirtualTime: true}
	el := New(vm, doc, opts)

	executionOrder := []string{}
	var mu sync.Mutex

	// Helper to track execution order
	addToOrder := func(name string) {
		mu.Lock()
		executionOrder = append(executionOrder, name)
		mu.Unlock()
	}

	// Schedule a task that queues microtasks
	taskCallback := goja.Callable(func(this goja.Value, args ...goja.Value) (goja.Value, error) {
		addToOrder("task")

		// Queue microtasks from the task
		el.QueueMicrotask(func() {
			addToOrder("microtask1")
		}, QueueMicrotaskSource)

		el.QueueMicrotask(func() {
			addToOrder("microtask2")
		}, QueueMicrotaskSource)

		return goja.Undefined(), nil
	})

	el.ScheduleTask(TimerTask, TimerTaskSource, taskCallback, nil, 0)

	// Process one task (this should also process all microtasks)
	el.ProcessNextTask()

	// Verify execution order: task, then all microtasks
	expected := []string{"task", "microtask1", "microtask2"}
	mu.Lock()
	if len(executionOrder) != len(expected) {
		t.Errorf("Expected %d executions, got %d", len(expected), len(executionOrder))
	}

	for i, exp := range expected {
		if i >= len(executionOrder) || executionOrder[i] != exp {
			t.Errorf("Expected execution order %v, got %v", expected, executionOrder)
			break
		}
	}
	mu.Unlock()
}

func TestAnimationFrameCallbacks(t *testing.T) {
	vm := goja.New()
	doc := &dom.Document{}
	opts := &EventLoopOptions{VirtualTime: true}
	el := New(vm, doc, opts)

	executed := false
	var receivedTimestamp float64

	callback := goja.Callable(func(this goja.Value, args ...goja.Value) (goja.Value, error) {
		executed = true
		if len(args) > 0 {
			if ts, ok := args[0].Export().(float64); ok {
				receivedTimestamp = ts
			}
		}
		return goja.Undefined(), nil
	})

	// Schedule animation frame callback
	id := el.RequestAnimationFrame(callback)
	if id == 0 {
		t.Error("Animation frame ID should not be zero")
	}

	// Verify it was queued
	counts := el.GetQueueCounts()
	if counts.AnimationFrames != 1 {
		t.Errorf("Expected 1 animation frame callback, got %d", counts.AnimationFrames)
	}

	// Process animation frame
	el.ProcessAnimationFrame()

	if !executed {
		t.Error("Animation frame callback should have been executed")
	}

	if receivedTimestamp == 0 {
		t.Error("Animation frame callback should receive a timestamp")
	}

	// Verify queue is empty
	counts = el.GetQueueCounts()
	if counts.AnimationFrames != 0 {
		t.Errorf("Expected 0 animation frame callbacks after processing, got %d", counts.AnimationFrames)
	}
}

func TestIdleCallbacks(t *testing.T) {
	vm := goja.New()
	doc := &dom.Document{}
	opts := &EventLoopOptions{VirtualTime: true}
	el := New(vm, doc, opts)

	executed := false
	callback := goja.Callable(func(this goja.Value, args ...goja.Value) (goja.Value, error) {
		executed = true
		return goja.Undefined(), nil
	})

	// Schedule idle callback
	id := el.RequestIdleCallback(callback, time.Second)
	if id == 0 {
		t.Error("Idle callback ID should not be zero")
	}

	// Verify it was queued
	counts := el.GetQueueCounts()
	if counts.IdleCallbacks != 1 {
		t.Errorf("Expected 1 idle callback, got %d", counts.IdleCallbacks)
	}

	// Process idle callbacks (should execute when no other work is pending)
	el.processIdleCallbacks()

	if !executed {
		t.Error("Idle callback should have been executed")
	}
}

func TestVirtualTimeAdvancement(t *testing.T) {
	vm := goja.New()
	doc := &dom.Document{}
	opts := &EventLoopOptions{VirtualTime: true}
	el := New(vm, doc, opts)

	executed := false
	callback := goja.Callable(func(this goja.Value, args ...goja.Value) (goja.Value, error) {
		executed = true
		return goja.Undefined(), nil
	})

	// Schedule task with 100ms delay
	el.ScheduleTask(TimerTask, TimerTaskSource, callback, nil, 100*time.Millisecond)

	// Task should not execute immediately
	task := el.ProcessNextTask()
	if task != nil {
		t.Error("Task should not execute before its delay")
	}

	if executed {
		t.Error("Task should not have executed yet")
	}

	// Advance virtual time by 100ms
	el.AdvanceTime(100 * time.Millisecond)

	if !executed {
		t.Error("Task should have executed after advancing time")
	}
}

func TestMetrics(t *testing.T) {
	vm := goja.New()
	doc := &dom.Document{}
	opts := &EventLoopOptions{VirtualTime: true}
	el := New(vm, doc, opts)

	// Initial metrics
	metrics := el.GetMetrics()
	if metrics.TasksProcessed != 0 || metrics.MicrotasksProcessed != 0 {
		t.Error("Initial metrics should be zero")
	}

	if metrics.IsRunning {
		t.Error("Event loop should not be running initially")
	}

	if !metrics.UsingVirtualTime {
		t.Error("Should be using virtual time")
	}

	// Execute some work
	taskCallback := goja.Callable(func(this goja.Value, args ...goja.Value) (goja.Value, error) {
		el.QueueMicrotask(func() {
			// Microtask work
		}, QueueMicrotaskSource)
		return goja.Undefined(), nil
	})

	el.ScheduleTask(TimerTask, TimerTaskSource, taskCallback, nil, 0)
	el.ProcessNextTask()

	// Check updated metrics
	metrics = el.GetMetrics()
	if metrics.TasksProcessed != 1 {
		t.Errorf("Expected 1 task processed, got %d", metrics.TasksProcessed)
	}

	if metrics.MicrotasksProcessed != 1 {
		t.Errorf("Expected 1 microtask processed, got %d", metrics.MicrotasksProcessed)
	}
}

func TestTaskCancellation(t *testing.T) {
	vm := goja.New()
	doc := &dom.Document{}
	opts := &EventLoopOptions{VirtualTime: true}
	el := New(vm, doc, opts)

	callback := goja.Callable(func(this goja.Value, args ...goja.Value) (goja.Value, error) {
		t.Error("Cancelled task should not execute")
		return goja.Undefined(), nil
	})

	// Schedule and cancel animation frame
	id := el.RequestAnimationFrame(callback)
	el.CancelAnimationFrame(id)

	// Process animation frames - should not execute the cancelled callback
	el.ProcessAnimationFrame()

	// Schedule and cancel idle callback
	id = el.RequestIdleCallback(callback, time.Second)
	el.CancelIdleCallback(id)

	// Process idle callbacks - should not execute the cancelled callback
	el.processIdleCallbacks()
}

func TestQueueLimits(t *testing.T) {
	vm := goja.New()
	doc := &dom.Document{}
	opts := &EventLoopOptions{VirtualTime: true}
	el := New(vm, doc, opts)

	// Start a microtask that would create infinite recursion
	el.QueueMicrotask(func() {
		// Each microtask queues another microtask
		el.QueueMicrotask(func() {
			// This would create infinite recursion without limits
		}, QueueMicrotaskSource)
	}, QueueMicrotaskSource)

	// This should not hang due to the microtask limit
	el.ProcessMicrotasks()

	// If we reach here, the limit worked
	metrics := el.GetMetrics()
	if metrics.MicrotasksProcessed == 0 {
		t.Error("Should have processed at least some microtasks")
	}
}
