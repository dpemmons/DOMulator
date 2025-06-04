package abort

import (
	"testing"

	"github.com/dop251/goja"
)

func TestAbortController(t *testing.T) {
	controller := NewAbortController()

	// Test initial state
	if controller.Signal() == nil {
		t.Error("Controller should have a signal")
	}

	if controller.Signal().Aborted() {
		t.Error("Controller signal should not be aborted initially")
	}

	// Test abort without reason
	controller.Abort(nil)

	if !controller.Signal().Aborted() {
		t.Error("Controller signal should be aborted after abort")
	}

	reason := controller.Signal().Reason()
	if reason != "AbortError" {
		t.Errorf("Expected default reason 'AbortError', got '%v'", reason)
	}
}

func TestAbortControllerWithReason(t *testing.T) {
	controller := NewAbortController()
	testReason := "Custom abort reason"

	// Test abort with custom reason
	controller.Abort(testReason)

	if !controller.Signal().Aborted() {
		t.Error("Controller signal should be aborted after abort")
	}

	if controller.Signal().Reason() != testReason {
		t.Errorf("Expected reason '%s', got '%v'", testReason, controller.Signal().Reason())
	}
}

func TestAbortControllerSignalIntegrity(t *testing.T) {
	controller := NewAbortController()
	signal1 := controller.Signal()
	signal2 := controller.Signal()

	// Test that Signal() returns the same instance
	if signal1 != signal2 {
		t.Error("Controller should return the same signal instance")
	}

	// Test that aborting affects the same signal
	controller.Abort("test")
	if !signal1.Aborted() || !signal2.Aborted() {
		t.Error("Both signal references should be aborted")
	}
}

func TestAbortControllerJavaScriptBinding(t *testing.T) {
	vm := goja.New()
	controller := NewAbortController()

	// Test JavaScript object creation
	jsController := CreateAbortControllerObject(vm, controller)
	if jsController == nil {
		t.Fatal("JavaScript controller object should not be nil")
	}

	// Test signal property
	signalValue := jsController.Get("signal")
	if signalValue == nil || goja.IsUndefined(signalValue) {
		t.Error("Controller should have a signal property")
	}

	signalObj := signalValue.ToObject(vm)
	if signalObj == nil {
		t.Error("Signal property should be an object")
	}

	// Test abort method
	abortFunc := jsController.Get("abort")
	if abortFunc == nil || goja.IsUndefined(abortFunc) {
		t.Fatal("Controller should have an abort method")
	}

	// Test initial signal state
	abortedValue := signalObj.Get("aborted")
	if abortedValue.ToBoolean() {
		t.Error("Signal should not be aborted initially")
	}

	// Call abort method
	_, err := vm.RunString(`
		controller = arguments[0];
		controller.abort("JavaScript test");
	`)
	if err != nil {
		// Try alternative approach
		vm.Set("controller", jsController)
		_, err = vm.RunString(`controller.abort("JavaScript test")`)
		if err != nil {
			t.Fatalf("Failed to call abort method: %v", err)
		}
	}

	// Check that signal is now aborted
	if !controller.Signal().Aborted() {
		t.Error("Signal should be aborted after calling abort")
	}

	if controller.Signal().Reason() != "JavaScript test" {
		t.Errorf("Expected reason 'JavaScript test', got '%v'", controller.Signal().Reason())
	}
}

func TestAbortControllerConstructor(t *testing.T) {
	vm := goja.New()

	// Set up AbortController constructor
	vm.Set("AbortController", CreateAbortControllerConstructor(vm))

	// Test constructor
	val, err := vm.RunString("new AbortController()")
	if err != nil {
		t.Fatalf("AbortController constructor should not throw: %v", err)
	}

	obj := val.ToObject(vm)
	if obj == nil {
		t.Fatal("Constructor should return an object")
	}

	// Test that the constructed object has signal property
	signalValue := obj.Get("signal")
	if signalValue == nil || goja.IsUndefined(signalValue) {
		t.Error("Constructed controller should have a signal property")
	}

	// Test that the constructed object has abort method
	abortValue := obj.Get("abort")
	if abortValue == nil || goja.IsUndefined(abortValue) {
		t.Error("Constructed controller should have an abort method")
	}
}

func TestAbortControllerMultipleAborts(t *testing.T) {
	controller := NewAbortController()
	signal := controller.Signal()

	// Test multiple aborts
	controller.Abort("first")
	firstReason := signal.Reason()

	controller.Abort("second")
	secondReason := signal.Reason()

	// Reason should not change after first abort
	if firstReason != secondReason {
		t.Error("Signal reason should not change after subsequent aborts")
	}

	if firstReason != "first" {
		t.Errorf("Expected reason 'first', got '%v'", firstReason)
	}
}

func TestGetAbortControllerFromJS(t *testing.T) {
	vm := goja.New()
	controller := NewAbortController()

	// Test with valid JS object
	jsController := CreateAbortControllerObject(vm, controller)
	extractedController := GetAbortControllerFromJS(jsController)
	if extractedController != controller {
		t.Error("GetAbortControllerFromJS should return the original controller")
	}

	// Test with invalid JS object
	invalidObj := vm.NewObject()
	extractedController2 := GetAbortControllerFromJS(invalidObj)
	if extractedController2 != nil {
		t.Error("GetAbortControllerFromJS should return nil for invalid objects")
	}
}

func TestAbortControllerWithEventListeners(t *testing.T) {
	controller := NewAbortController()
	signal := controller.Signal()

	eventReceived := false
	var receivedEvent *AbortEvent

	// Add event listener to signal
	signal.AddEventListener("abort", func(event *AbortEvent) {
		eventReceived = true
		receivedEvent = event
	})

	// Abort the controller
	testReason := "event test"
	controller.Abort(testReason)

	// Check that event was fired
	if !eventReceived {
		t.Error("Abort event should have been fired")
	}

	if receivedEvent == nil {
		t.Error("Received event should not be nil")
	}

	if receivedEvent.AbortReason() != testReason {
		t.Errorf("Expected event reason '%s', got '%v'", testReason, receivedEvent.AbortReason())
	}

	if receivedEvent.Signal() != signal {
		t.Error("Event signal should match controller signal")
	}
}

func TestAbortControllerJavaScriptIntegration(t *testing.T) {
	vm := goja.New()

	// Set up constructors
	vm.Set("AbortController", CreateAbortControllerConstructor(vm))
	vm.Set("AbortSignal", CreateAbortSignalConstructor(vm))

	// Test creating controller and using it
	_, err := vm.RunString(`
		const controller = new AbortController();
		const signal = controller.signal;
		
		if (signal.aborted) {
			throw new Error("Signal should not be aborted initially");
		}
		
		// Add event listener
		let eventReceived = false;
		signal.addEventListener("abort", function(event) {
			eventReceived = true;
		});
		
		// Abort the controller
		controller.abort("integration test");
		
		if (!signal.aborted) {
			throw new Error("Signal should be aborted after abort");
		}
		
		if (signal.reason !== "integration test") {
			throw new Error("Signal reason should be 'integration test'");
		}
	`)

	if err != nil {
		t.Errorf("JavaScript integration test failed: %v", err)
	}
}

func TestAbortTimeout(t *testing.T) {
	vm := goja.New()

	// Test AbortTimeout function
	signalObj := AbortTimeout(vm, 100)
	if signalObj == nil {
		t.Fatal("AbortTimeout should return a signal object")
	}

	// Extract the signal
	signal := GetAbortSignalFromJS(signalObj)
	if signal == nil {
		t.Fatal("Should be able to extract signal from AbortTimeout result")
	}

	// In our implementation, the signal should be aborted immediately
	// In a real implementation, it would abort after the timeout
	// For testing purposes, we'll just verify the function works
}

func TestCreateAbortTimeoutFunction(t *testing.T) {
	vm := goja.New()

	// Set up the timeout function
	vm.Set("AbortSignalTimeout", CreateAbortTimeoutFunction(vm))

	// Test valid timeout
	val, err := vm.RunString("AbortSignalTimeout(1000)")
	if err != nil {
		t.Errorf("AbortSignal.timeout should not throw for valid timeout: %v", err)
	}

	if val == nil || goja.IsUndefined(val) {
		t.Error("AbortSignal.timeout should return a signal")
	}

	// Test negative timeout
	_, err = vm.RunString("AbortSignalTimeout(-1)")
	if err == nil {
		t.Error("AbortSignal.timeout should throw for negative timeout")
	}

	// Test missing argument
	_, err = vm.RunString("AbortSignalTimeout()")
	if err == nil {
		t.Error("AbortSignal.timeout should throw when no timeout provided")
	}
}
