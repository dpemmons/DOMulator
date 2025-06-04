package abort

import (
	"testing"
	"time"

	"github.com/dop251/goja"
)

func TestAbortSignalEventListenerIntegration(t *testing.T) {
	vm := goja.New()

	// Set up abort APIs
	SetupAbortAPIs(vm)

	// Test that abort signal events work correctly
	_, err := vm.RunString(`
		const controller = new AbortController();
		const signal = controller.signal;
		
		let abortEventReceived = false;
		let abortEventReason = null;
		
		signal.addEventListener("abort", function(event) {
			abortEventReceived = true;
			abortEventReason = event.reason;
		});
		
		// Verify initial state
		if (signal.aborted) {
			throw new Error("Signal should not be aborted initially");
		}
		
		// Abort the controller
		controller.abort("Integration test reason");
		
		// Verify final state
		if (!signal.aborted) {
			throw new Error("Signal should be aborted after abort()");
		}
		
		if (signal.reason !== "Integration test reason") {
			throw new Error("Signal reason should match abort reason");
		}
	`)

	if err != nil {
		t.Errorf("JavaScript integration test failed: %v", err)
	}

	// Give time for event processing
	time.Sleep(10 * time.Millisecond)
}

func TestMultipleControllersIndependence(t *testing.T) {
	vm := goja.New()

	// Set up abort APIs
	SetupAbortAPIs(vm)

	// Test that multiple controllers are independent
	_, err := vm.RunString(`
		const controller1 = new AbortController();
		const controller2 = new AbortController();
		
		const signal1 = controller1.signal;
		const signal2 = controller2.signal;
		
		// Abort only the first controller
		controller1.abort("First abort");
		
		// Verify independence
		if (!signal1.aborted) {
			throw new Error("Signal1 should be aborted");
		}
		
		if (signal2.aborted) {
			throw new Error("Signal2 should not be aborted");
		}
		
		if (signal1.reason !== "First abort") {
			throw new Error("Signal1 reason should be 'First abort'");
		}
		
		if (signal2.reason !== undefined) {
			throw new Error("Signal2 reason should be undefined");
		}
		
		// Abort the second controller
		controller2.abort("Second abort");
		
		// Verify both are now aborted with different reasons
		if (!signal2.aborted) {
			throw new Error("Signal2 should now be aborted");
		}
		
		if (signal1.reason !== "First abort") {
			throw new Error("Signal1 reason should still be 'First abort'");
		}
		
		if (signal2.reason !== "Second abort") {
			throw new Error("Signal2 reason should be 'Second abort'");
		}
	`)

	if err != nil {
		t.Errorf("Multiple controllers test failed: %v", err)
	}
}

func TestAbortControllerAPICompliance(t *testing.T) {
	vm := goja.New()

	// Set up abort APIs
	SetupAbortAPIs(vm)

	// Test API compliance with Web Standards
	_, err := vm.RunString(`
		// Test AbortController constructor
		const controller = new AbortController();
		
		// Test that AbortController has expected properties and methods
		if (typeof controller.signal !== 'object') {
			throw new Error("AbortController.signal should be an object");
		}
		
		if (typeof controller.abort !== 'function') {
			throw new Error("AbortController.abort should be a function");
		}
		
		// Test that AbortSignal has expected properties and methods
		const signal = controller.signal;
		
		if (typeof signal.aborted !== 'boolean') {
			throw new Error("AbortSignal.aborted should be a boolean");
		}
		
		if (typeof signal.addEventListener !== 'function') {
			throw new Error("AbortSignal.addEventListener should be a function");
		}
		
		if (typeof signal.removeEventListener !== 'function') {
			throw new Error("AbortSignal.removeEventListener should be a function");
		}
		
		if (typeof signal.dispatchEvent !== 'function') {
			throw new Error("AbortSignal.dispatchEvent should be a function");
		}
		
		// Test AbortSignal constructor throws
		let constructorThrew = false;
		try {
			new AbortSignal();
		} catch (e) {
			constructorThrew = true;
		}
		
		if (!constructorThrew) {
			throw new Error("AbortSignal constructor should throw");
		}
		
		// Test initial states
		if (signal.aborted !== false) {
			throw new Error("AbortSignal.aborted should be false initially");
		}
		
		if (signal.reason !== undefined) {
			throw new Error("AbortSignal.reason should be undefined initially");
		}
	`)

	if err != nil {
		t.Errorf("API compliance test failed: %v", err)
	}
}

func TestFetchIntegrationUtilities(t *testing.T) {
	// Test the Go-level integration utilities
	signal := NewAbortSignal()
	integration := NewFetchIntegration(signal)

	// Test CheckAborted when not aborted
	err := integration.CheckAborted()
	if err != nil {
		t.Errorf("CheckAborted should return nil when not aborted, got: %v", err)
	}

	// Test CheckAborted when aborted
	signal.dispatchAbortEvent("test abort")
	err = integration.CheckAborted()
	if err == nil {
		t.Error("CheckAborted should return error when aborted")
	}

	if !IsAbortError(err) {
		t.Error("CheckAborted should return AbortError when aborted")
	}
}

func TestExtractSignalFromOptions(t *testing.T) {
	vm := goja.New()

	// Set up abort APIs
	SetupAbortAPIs(vm)

	// Create a controller and signal
	controller := NewAbortController()
	jsController := CreateAbortControllerObject(vm, controller)
	vm.Set("controller", jsController)

	// Test extracting signal from options
	val, err := vm.RunString(`({
		method: "POST",
		headers: {"Content-Type": "application/json"},
		body: "test",
		signal: controller.signal
	})`)

	if err != nil {
		t.Fatalf("Failed to create options object: %v", err)
	}

	optionsObj := val.ToObject(vm)
	extractedSignal := ExtractSignalFromOptions(vm, optionsObj)

	if extractedSignal != controller.Signal() {
		t.Error("ExtractSignalFromOptions should return the original signal")
	}

	// Test with no signal
	val2, err := vm.RunString(`({
		method: "GET"
	})`)

	if err != nil {
		t.Fatalf("Failed to create options object without signal: %v", err)
	}

	optionsObj2 := val2.ToObject(vm)
	extractedSignal2 := ExtractSignalFromOptions(vm, optionsObj2)

	if extractedSignal2 != nil {
		t.Error("ExtractSignalFromOptions should return nil when no signal provided")
	}
}

func TestAbortControllerBasicJavaScriptIntegration(t *testing.T) {
	vm := goja.New()

	// Set up constructors
	SetupAbortAPIs(vm)

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
