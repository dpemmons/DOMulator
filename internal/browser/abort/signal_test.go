package abort

import (
	"testing"
	"time"

	"github.com/dop251/goja"
)

func TestAbortSignal(t *testing.T) {
	signal := NewAbortSignal()

	// Test initial state
	if signal.Aborted() {
		t.Error("Signal should not be aborted initially")
	}

	if signal.Reason() != nil {
		t.Error("Signal reason should be nil initially")
	}

	// Test abort
	testReason := "test abort"
	signal.dispatchAbortEvent(testReason)

	if !signal.Aborted() {
		t.Error("Signal should be aborted after dispatchAbortEvent")
	}

	if signal.Reason() != testReason {
		t.Errorf("Expected reason '%s', got '%v'", testReason, signal.Reason())
	}

	// Test multiple aborts (should be idempotent)
	signal.dispatchAbortEvent("second abort")
	if signal.Reason() != testReason {
		t.Error("Reason should not change after second abort")
	}
}

func TestAbortSignalEventListeners(t *testing.T) {
	signal := NewAbortSignal()
	eventReceived := false
	var receivedEvent *AbortEvent

	// Add event listener
	signal.AddEventListener("abort", func(event *AbortEvent) {
		eventReceived = true
		receivedEvent = event
	})

	// Dispatch abort event
	testReason := "test event"
	signal.dispatchAbortEvent(testReason)

	if !eventReceived {
		t.Error("Event listener should have been called")
	}

	if receivedEvent == nil {
		t.Error("Received event should not be nil")
	}

	if receivedEvent.AbortReason() != testReason {
		t.Errorf("Expected event reason '%s', got '%v'", testReason, receivedEvent.AbortReason())
	}

	if receivedEvent.Signal() != signal {
		t.Error("Event signal should match the original signal")
	}
}

func TestAbortSignalJavaScriptBinding(t *testing.T) {
	vm := goja.New()
	signal := NewAbortSignal()

	// Test JavaScript object creation
	jsSignal := CreateAbortSignalObject(vm, signal)
	if jsSignal == nil {
		t.Fatal("JavaScript signal object should not be nil")
	}

	// Test initial properties
	abortedValue := jsSignal.Get("aborted")
	if abortedValue.ToBoolean() {
		t.Error("JavaScript signal should not be aborted initially")
	}

	reasonValue := jsSignal.Get("reason")
	if !goja.IsUndefined(reasonValue) && !goja.IsNull(reasonValue) {
		t.Error("JavaScript signal reason should be undefined/null initially")
	}

	// Test abort
	signal.dispatchAbortEvent("js test")

	// Create new JS object to reflect updated state
	jsSignal2 := CreateAbortSignalObject(vm, signal)
	abortedValue2 := jsSignal2.Get("aborted")
	if !abortedValue2.ToBoolean() {
		t.Error("JavaScript signal should be aborted after abort")
	}

	reasonValue2 := jsSignal2.Get("reason")
	if reasonValue2.String() != "js test" {
		t.Errorf("Expected reason 'js test', got '%s'", reasonValue2.String())
	}
}

func TestAbortSignalConstructor(t *testing.T) {
	vm := goja.New()

	// Set up AbortSignal constructor
	vm.Set("AbortSignal", CreateAbortSignalConstructor(vm))

	// Test that constructor throws
	_, err := vm.RunString("new AbortSignal()")
	if err == nil {
		t.Error("AbortSignal constructor should throw an error")
	}
}

func TestAbortEvent(t *testing.T) {
	vm := goja.New()
	signal := NewAbortSignal()

	testReason := "test event creation"
	event := &AbortEvent{
		BaseEvent: signal.BaseEvent,
		signal:    signal,
		reason:    testReason,
	}

	// Test event properties
	if event.Signal() != signal {
		t.Error("Event signal should match the original signal")
	}

	if event.AbortReason() != testReason {
		t.Errorf("Expected event reason '%s', got '%v'", testReason, event.AbortReason())
	}

	// Test JavaScript object creation
	jsEvent := CreateAbortEventObject(vm, event)
	if jsEvent == nil {
		t.Fatal("JavaScript event object should not be nil")
	}

	// Test JavaScript event properties
	typeValue := jsEvent.Get("type")
	if typeValue.String() != "abort" {
		t.Errorf("Expected event type 'abort', got '%s'", typeValue.String())
	}

	reasonValue := jsEvent.Get("reason")
	if reasonValue.String() != testReason {
		t.Errorf("Expected event reason '%s', got '%s'", testReason, reasonValue.String())
	}

	bubblesValue := jsEvent.Get("bubbles")
	if bubblesValue.ToBoolean() {
		t.Error("Abort events should not bubble")
	}

	cancelableValue := jsEvent.Get("cancelable")
	if cancelableValue.ToBoolean() {
		t.Error("Abort events should not be cancelable")
	}
}

func TestAbortError(t *testing.T) {
	// Test with reason
	err := &AbortError{reason: "test abort reason"}
	if err.Error() != "test abort reason" {
		t.Errorf("Expected error message 'test abort reason', got '%s'", err.Error())
	}

	// Test without reason
	err2 := &AbortError{reason: nil}
	if err2.Error() != "The operation was aborted" {
		t.Errorf("Expected default error message, got '%s'", err2.Error())
	}

	// Test IsAbortError
	if !IsAbortError(err) {
		t.Error("IsAbortError should return true for AbortError")
	}

	// Test with non-abort error
	otherErr := &testError{}
	if IsAbortError(otherErr) {
		t.Error("IsAbortError should return false for non-AbortError")
	}
}

func TestGetAbortSignalFromJS(t *testing.T) {
	vm := goja.New()
	signal := NewAbortSignal()

	// Test with valid JS object
	jsSignal := CreateAbortSignalObject(vm, signal)
	extractedSignal := GetAbortSignalFromJS(jsSignal)
	if extractedSignal != signal {
		t.Error("GetAbortSignalFromJS should return the original signal")
	}

	// Test with invalid JS object
	invalidObj := vm.NewObject()
	extractedSignal2 := GetAbortSignalFromJS(invalidObj)
	if extractedSignal2 != nil {
		t.Error("GetAbortSignalFromJS should return nil for invalid objects")
	}
}

func TestAbortSignalJavaScriptEventListeners(t *testing.T) {
	vm := goja.New()
	signal := NewAbortSignal()
	jsSignal := CreateAbortSignalObject(vm, signal)

	// Test addEventListener
	eventReceived := false
	vm.Set("testCallback", func(call goja.FunctionCall) goja.Value {
		eventReceived = true
		return goja.Undefined()
	})

	addEventListenerFunc := jsSignal.Get("addEventListener")
	if addEventListenerFunc == nil || goja.IsUndefined(addEventListenerFunc) {
		t.Fatal("addEventListener should be available")
	}

	// Add listener via JavaScript
	_, err := vm.RunString(`
		signal = arguments[0];
		signal.addEventListener("abort", testCallback);
	`)
	if err != nil {
		vm.Set("signal", jsSignal)
		_, err = vm.RunString(`signal.addEventListener("abort", testCallback)`)
		if err != nil {
			t.Fatalf("Failed to add event listener: %v", err)
		}
	}

	// Trigger abort
	signal.dispatchAbortEvent("js listener test")

	// Give some time for async event handling
	time.Sleep(10 * time.Millisecond)

	if !eventReceived {
		t.Error("JavaScript event listener should have been called")
	}
}

// Helper type for testing
type testError struct{}

func (e *testError) Error() string {
	return "test error"
}
