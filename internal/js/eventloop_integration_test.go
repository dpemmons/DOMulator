package js

import (
	"testing"
	"time"

	"github.com/dpemmons/DOMulator/internal/dom"
)

func TestEventLoopIntegration(t *testing.T) {
	document := dom.NewDocument()
	runtime := New(document)

	// Setup event loop with virtual time for deterministic testing
	runtime.SetupEventLoop(true)

	if runtime.EventLoop() == nil {
		t.Fatal("Event loop should be initialized")
	}

	// Test that event loop APIs are available
	global := runtime.VM().GlobalObject()

	funcs := []string{
		"queueMicrotask",
		"requestAnimationFrame",
		"cancelAnimationFrame",
		"requestIdleCallback",
		"cancelIdleCallback",
	}

	for _, funcName := range funcs {
		if global.Get(funcName) == nil {
			t.Errorf("%s should be available globally", funcName)
		}
	}

	// Test on window object too
	window := global.Get("window").ToObject(runtime.VM())
	for _, funcName := range funcs {
		if window.Get(funcName) == nil {
			t.Errorf("%s should be available on window", funcName)
		}
	}

	runtime.Shutdown()
}

func TestQueueMicrotask(t *testing.T) {
	document := dom.NewDocument()
	runtime := New(document)
	runtime.SetupEventLoop(true)

	// Test queueMicrotask execution
	code := `
		let executed = false;
		queueMicrotask(() => {
			executed = true;
		});
		executed; // Should be false before microtask checkpoint
	`

	result, err := runtime.RunString(code)
	if err != nil {
		t.Fatalf("Error executing code: %v", err)
	}

	// Before microtask processing
	if result.ToBoolean() {
		t.Error("Microtask should not execute synchronously")
	}

	// Process microtasks
	runtime.EventLoop().ProcessMicrotasks()

	// Check if microtask executed
	result, err = runtime.RunString("executed")
	if err != nil {
		t.Fatalf("Error checking result: %v", err)
	}

	if !result.ToBoolean() {
		t.Error("Microtask should have executed after checkpoint")
	}

	runtime.Shutdown()
}

func TestRequestAnimationFrame(t *testing.T) {
	document := dom.NewDocument()
	runtime := New(document)
	runtime.SetupEventLoop(true)

	// Test requestAnimationFrame
	code := `
		let frameExecuted = false;
		let frameTimestamp = 0;
		
		let frameId = requestAnimationFrame((timestamp) => {
			frameExecuted = true;
			frameTimestamp = timestamp;
		});
		
		// Should return a valid ID
		frameId > 0;
	`

	result, err := runtime.RunString(code)
	if err != nil {
		t.Fatalf("Error executing code: %v", err)
	}

	if !result.ToBoolean() {
		t.Error("requestAnimationFrame should return a valid ID")
	}

	// Frame should not execute immediately
	result, err = runtime.RunString("frameExecuted")
	if err != nil {
		t.Fatalf("Error checking result: %v", err)
	}

	if result.ToBoolean() {
		t.Error("Animation frame should not execute synchronously")
	}

	// Process animation frame
	runtime.EventLoop().ProcessAnimationFrame()

	// Check if frame executed
	result, err = runtime.RunString("frameExecuted")
	if err != nil {
		t.Fatalf("Error checking result: %v", err)
	}

	if !result.ToBoolean() {
		t.Error("Animation frame should have executed")
	}

	// Check timestamp is provided
	result, err = runtime.RunString("frameTimestamp > 0")
	if err != nil {
		t.Fatalf("Error checking timestamp: %v", err)
	}

	if !result.ToBoolean() {
		t.Error("Animation frame should receive a valid timestamp")
	}

	runtime.Shutdown()
}

func TestCancelAnimationFrame(t *testing.T) {
	document := dom.NewDocument()
	runtime := New(document)
	runtime.SetupEventLoop(true)

	// Test cancelAnimationFrame
	code := `
		let frameExecuted = false;
		
		let frameId = requestAnimationFrame(() => {
			frameExecuted = true;
		});
		
		cancelAnimationFrame(frameId);
		frameId; // Return the ID for verification
	`

	result, err := runtime.RunString(code)
	if err != nil {
		t.Fatalf("Error executing code: %v", err)
	}

	if result.ToInteger() <= 0 {
		t.Error("Should have received a valid frame ID")
	}

	// Process animation frame - should not execute the cancelled callback
	runtime.EventLoop().ProcessAnimationFrame()

	// Check that frame was not executed
	result, err = runtime.RunString("frameExecuted")
	if err != nil {
		t.Fatalf("Error checking result: %v", err)
	}

	if result.ToBoolean() {
		t.Error("Cancelled animation frame should not execute")
	}

	runtime.Shutdown()
}

func TestRequestIdleCallback(t *testing.T) {
	document := dom.NewDocument()
	runtime := New(document)
	runtime.SetupEventLoop(true)

	// Test requestIdleCallback
	code := `
		let idleExecuted = false;
		let idleDeadline = null;
		
		let idleId = requestIdleCallback((deadline) => {
			idleExecuted = true;
			idleDeadline = deadline;
		});
		
		idleId > 0; // Should return a valid ID
	`

	result, err := runtime.RunString(code)
	if err != nil {
		t.Fatalf("Error executing code: %v", err)
	}

	if !result.ToBoolean() {
		t.Error("requestIdleCallback should return a valid ID")
	}

	// Idle callback should not execute immediately
	result, err = runtime.RunString("idleExecuted")
	if err != nil {
		t.Fatalf("Error checking result: %v", err)
	}

	if result.ToBoolean() {
		t.Error("Idle callback should not execute synchronously")
	}

	// Process a few event loop iterations to trigger idle processing
	for i := 0; i < 3; i++ {
		runtime.EventLoop().ProcessNextTask()
	}

	runtime.Shutdown()
}

func TestRequestIdleCallbackWithTimeout(t *testing.T) {
	document := dom.NewDocument()
	runtime := New(document)
	runtime.SetupEventLoop(true)

	// Test requestIdleCallback with timeout option
	code := `
		let idleExecuted = false;
		
		let idleId = requestIdleCallback(() => {
			idleExecuted = true;
		}, { timeout: 1000 });
		
		idleId > 0; // Should return a valid ID
	`

	result, err := runtime.RunString(code)
	if err != nil {
		t.Fatalf("Error executing code: %v", err)
	}

	if !result.ToBoolean() {
		t.Error("requestIdleCallback with timeout should return a valid ID")
	}

	runtime.Shutdown()
}

func TestCancelIdleCallback(t *testing.T) {
	document := dom.NewDocument()
	runtime := New(document)
	runtime.SetupEventLoop(true)

	// Test cancelIdleCallback
	code := `
		let idleExecuted = false;
		
		let idleId = requestIdleCallback(() => {
			idleExecuted = true;
		});
		
		cancelIdleCallback(idleId);
		idleId; // Return the ID for verification
	`

	result, err := runtime.RunString(code)
	if err != nil {
		t.Fatalf("Error executing code: %v", err)
	}

	if result.ToInteger() <= 0 {
		t.Error("Should have received a valid idle callback ID")
	}

	runtime.Shutdown()
}

func TestEventLoopProcessingOrder(t *testing.T) {
	document := dom.NewDocument()
	runtime := New(document)
	runtime.SetupEventLoop(true)

	// Test the correct processing order: tasks -> microtasks -> animation frames -> idle
	code := `
		let order = [];
		
		// Queue items in different order than they should execute
		requestIdleCallback(() => order.push('idle'));
		
		setTimeout(() => {
			order.push('task');
			queueMicrotask(() => order.push('microtask'));
		}, 0);
		
		requestAnimationFrame(() => order.push('animation'));
		
		order; // Should be empty initially
	`

	result, err := runtime.RunString(code)
	if err != nil {
		t.Fatalf("Error executing code: %v", err)
	}

	// Should be empty initially
	if result.ToObject(runtime.VM()).Get("length").ToInteger() != 0 {
		t.Error("Order array should be empty initially")
	}

	// Process one task (setTimeout)
	runtime.EventLoop().ProcessNextTask()

	// Check microtask executed after task
	result, err = runtime.RunString("order")
	if err != nil {
		t.Fatalf("Error checking order: %v", err)
	}

	arr := result.ToObject(runtime.VM())
	if arr.Get("length").ToInteger() != 2 {
		t.Error("Should have task and microtask processed")
	}

	// Verify correct order: task, then microtask
	if arr.Get("0").String() != "task" {
		t.Error("First item should be 'task'")
	}
	if arr.Get("1").String() != "microtask" {
		t.Error("Second item should be 'microtask'")
	}

	// Process animation frame
	runtime.EventLoop().ProcessAnimationFrame()

	result, err = runtime.RunString("order")
	if err != nil {
		t.Fatalf("Error checking order: %v", err)
	}

	arr = result.ToObject(runtime.VM())
	if arr.Get("length").ToInteger() >= 3 {
		if arr.Get("2").String() != "animation" {
			t.Error("Third item should be 'animation'")
		}
	}

	runtime.Shutdown()
}

func TestVirtualTimeAdvancement(t *testing.T) {
	document := dom.NewDocument()
	runtime := New(document)
	runtime.SetupEventLoop(true) // Virtual time enabled

	// Test virtual time advancement
	code := `
		let executed = false;
		
		setTimeout(() => {
			executed = true;
		}, 100); // 100ms delay
		
		executed; // Should be false initially
	`

	result, err := runtime.RunString(code)
	if err != nil {
		t.Fatalf("Error executing code: %v", err)
	}

	if result.ToBoolean() {
		t.Error("Timer should not execute immediately")
	}

	// Advance virtual time by 50ms - should not trigger
	runtime.EventLoop().AdvanceTime(50 * time.Millisecond)

	result, err = runtime.RunString("executed")
	if err != nil {
		t.Fatalf("Error checking result: %v", err)
	}

	if result.ToBoolean() {
		t.Error("Timer should not execute before 100ms")
	}

	// Advance virtual time by another 60ms (total 110ms) - should trigger
	runtime.EventLoop().AdvanceTime(60 * time.Millisecond)

	result, err = runtime.RunString("executed")
	if err != nil {
		t.Fatalf("Error checking result: %v", err)
	}

	if !result.ToBoolean() {
		t.Error("Timer should execute after 100ms virtual time")
	}

	runtime.Shutdown()
}

func TestEventLoopMetrics(t *testing.T) {
	document := dom.NewDocument()
	runtime := New(document)
	runtime.SetupEventLoop(true)

	// Execute some operations to generate metrics
	_, err := runtime.RunString(`
		queueMicrotask(() => {});
		setTimeout(() => {}, 0);
		requestAnimationFrame(() => {});
	`)
	if err != nil {
		t.Fatalf("Error executing code: %v", err)
	}

	// Process to generate some metrics
	runtime.EventLoop().ProcessNextTask()
	runtime.EventLoop().ProcessAnimationFrame()

	metrics := runtime.EventLoop().GetMetrics()

	if !metrics.UsingVirtualTime {
		t.Error("Should be using virtual time")
	}

	if metrics.TasksProcessed == 0 {
		t.Error("Should have processed at least one task")
	}

	if metrics.MicrotasksProcessed == 0 {
		t.Error("Should have processed at least one microtask")
	}

	runtime.Shutdown()
}

func TestEventLoopQueueStats(t *testing.T) {
	document := dom.NewDocument()
	runtime := New(document)
	runtime.SetupEventLoop(true)

	// Queue some operations
	_, err := runtime.RunString(`
		queueMicrotask(() => {});
		queueMicrotask(() => {});
		setTimeout(() => {}, 0);
		requestAnimationFrame(() => {});
		requestIdleCallback(() => {});
	`)
	if err != nil {
		t.Fatalf("Error executing code: %v", err)
	}

	stats := runtime.EventLoop().GetQueueCounts()

	if stats.Microtasks != 2 {
		t.Errorf("Expected 2 microtasks, got %d", stats.Microtasks)
	}

	if stats.Tasks != 1 {
		t.Errorf("Expected 1 task, got %d", stats.Tasks)
	}

	if stats.AnimationFrames != 1 {
		t.Errorf("Expected 1 animation frame, got %d", stats.AnimationFrames)
	}

	if stats.IdleCallbacks != 1 {
		t.Errorf("Expected 1 idle callback, got %d", stats.IdleCallbacks)
	}

	runtime.Shutdown()
}
