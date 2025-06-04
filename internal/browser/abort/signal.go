package abort

import (
	"sync"
	"sync/atomic"

	"github.com/dop251/goja"
	"github.com/dpemmons/DOMulator/internal/dom"
)

// AbortSignal represents an abort signal that can be used to cancel operations
// This implements the Web API AbortSignal interface
type AbortSignal struct {
	*dom.BaseEvent
	aborted   atomic.Bool
	reason    atomic.Value
	onabort   func(*AbortEvent)
	listeners []func(*AbortEvent)
	jsObjects []func() // Functions to update associated JS objects
	mu        sync.RWMutex
}

// NewAbortSignal creates a new AbortSignal
func NewAbortSignal() *AbortSignal {
	signal := &AbortSignal{
		BaseEvent: dom.NewEvent("abort", false, false),
		listeners: make([]func(*AbortEvent), 0),
		jsObjects: make([]func(), 0),
	}

	// Initialize reason to a special "undefined" value since atomic.Value can't store nil
	signal.reason.Store("__undefined__")

	return signal
}

// Aborted returns true if the signal has been aborted
func (s *AbortSignal) Aborted() bool {
	return s.aborted.Load()
}

// Reason returns the reason the signal was aborted, or nil if not aborted
func (s *AbortSignal) Reason() interface{} {
	reason := s.reason.Load()
	if reason == "__undefined__" {
		return nil
	}
	return reason
}

// SetOnAbort sets the onabort event handler
func (s *AbortSignal) SetOnAbort(handler func(*AbortEvent)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.onabort = handler
}

// AddEventListener adds an event listener for the abort event
func (s *AbortSignal) AddEventListener(eventType string, listener func(*AbortEvent)) {
	if eventType != "abort" {
		return // AbortSignal only supports abort events
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	s.listeners = append(s.listeners, listener)
}

// RemoveEventListener removes an event listener for the abort event
func (s *AbortSignal) RemoveEventListener(eventType string, listener func(*AbortEvent)) {
	if eventType != "abort" {
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	// Remove the listener (simplified implementation - in real implementation would need function comparison)
	// For now, we'll just clear all listeners when any removal is requested
	s.listeners = make([]func(*AbortEvent), 0)
}

// dispatchAbortEvent dispatches an abort event to all listeners
func (s *AbortSignal) dispatchAbortEvent(reason interface{}) {
	if s.Aborted() {
		return // Already aborted
	}

	// Set aborted state and reason atomically
	s.aborted.Store(true)
	s.reason.Store(reason)

	// Create abort event
	event := &AbortEvent{
		BaseEvent: dom.NewEvent("abort", false, false),
		signal:    s,
		reason:    reason,
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	// Call onabort handler
	if s.onabort != nil {
		s.onabort(event)
	}

	// Call all event listeners
	for _, listener := range s.listeners {
		listener(event)
	}

	// Update all associated JavaScript objects
	for _, updateFn := range s.jsObjects {
		updateFn()
	}
}

// CreateAbortSignalConstructor creates the AbortSignal constructor for JavaScript binding
func CreateAbortSignalConstructor(vm *goja.Runtime) func(goja.ConstructorCall) *goja.Object {
	return func(call goja.ConstructorCall) *goja.Object {
		// AbortSignal constructor should throw - signals are created by AbortController
		panic(vm.NewTypeError("AbortSignal constructor is not available"))
	}
}

// CreateAbortSignalObject creates a JavaScript AbortSignal object from a Go AbortSignal
func CreateAbortSignalObject(vm *goja.Runtime, signal *AbortSignal) *goja.Object {
	obj := vm.NewObject()

	// Set up properties (we'll update them dynamically)
	obj.Set("aborted", signal.Aborted())

	// Handle reason property - should be undefined if not aborted yet
	reason := signal.Reason()
	if reason == nil {
		obj.Set("reason", goja.Undefined())
	} else {
		obj.Set("reason", vm.ToValue(reason))
	}

	// Add an internal method to update properties when state changes
	updateProperties := func() {
		obj.Set("aborted", signal.Aborted())
		reason := signal.Reason()
		if reason == nil {
			obj.Set("reason", goja.Undefined())
		} else {
			obj.Set("reason", vm.ToValue(reason))
		}
	}

	// Register the update function with the signal so it can be called when state changes
	signal.mu.Lock()
	signal.jsObjects = append(signal.jsObjects, updateProperties)
	signal.mu.Unlock()

	// Store the update function so we can call it when the signal changes
	obj.Set("__updateProperties__", vm.ToValue(updateProperties))

	// Set up onabort property
	obj.Set("onabort", goja.Null())

	// addEventListener method
	obj.Set("addEventListener", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 2 {
			panic(vm.NewTypeError("addEventListener requires 2 arguments"))
		}

		eventType := call.Arguments[0].String()
		if eventType != "abort" {
			return goja.Undefined()
		}

		listener, ok := goja.AssertFunction(call.Arguments[1])
		if !ok {
			panic(vm.NewTypeError("Listener must be a function"))
		}

		// Wrap the JavaScript function
		signal.AddEventListener(eventType, func(event *AbortEvent) {
			jsEvent := CreateAbortEventObject(vm, event)
			_, _ = listener(goja.Undefined(), vm.ToValue(jsEvent))
		})

		return goja.Undefined()
	})

	// removeEventListener method
	obj.Set("removeEventListener", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 2 {
			panic(vm.NewTypeError("removeEventListener requires 2 arguments"))
		}

		eventType := call.Arguments[0].String()
		listener, ok := goja.AssertFunction(call.Arguments[1])
		if !ok {
			panic(vm.NewTypeError("Listener must be a function"))
		}

		// For simplicity, remove all listeners when any removal is requested
		signal.RemoveEventListener(eventType, func(event *AbortEvent) {
			_, _ = listener(goja.Undefined(), vm.ToValue(event))
		})

		return goja.Undefined()
	})

	// dispatchEvent method (for completeness)
	obj.Set("dispatchEvent", func(call goja.FunctionCall) goja.Value {
		// AbortSignal events are typically only dispatched internally
		return vm.ToValue(false)
	})

	// throwIfAborted static method (should be on the constructor)
	obj.Set("throwIfAborted", func(call goja.FunctionCall) goja.Value {
		if signal.Aborted() {
			reason := signal.Reason()
			if reason == nil {
				reason = "AbortError"
			}
			panic(vm.NewGoError(&AbortError{reason: reason}))
		}
		return goja.Undefined()
	})

	// Store reference to Go signal for internal use
	obj.Set("__abortsignal__", vm.ToValue(signal))

	return obj
}

// GetAbortSignalFromJS extracts the Go AbortSignal from a JavaScript object
func GetAbortSignalFromJS(jsSignal *goja.Object) *AbortSignal {
	if jsSignal == nil {
		return nil
	}

	internalRef := jsSignal.Get("__abortsignal__")
	if internalRef == nil || goja.IsUndefined(internalRef) || goja.IsNull(internalRef) {
		return nil
	}

	exported := internalRef.Export()
	if exported == nil {
		return nil
	}

	if signal, ok := exported.(*AbortSignal); ok {
		return signal
	}
	return nil
}

// AbortEvent represents an abort event
type AbortEvent struct {
	*dom.BaseEvent
	signal *AbortSignal
	reason interface{}
}

// Signal returns the AbortSignal that was aborted
func (e *AbortEvent) Signal() *AbortSignal {
	return e.signal
}

// AbortReason returns the reason for the abort
func (e *AbortEvent) AbortReason() interface{} {
	return e.reason
}

// CreateAbortEventObject creates a JavaScript AbortEvent object
func CreateAbortEventObject(vm *goja.Runtime, event *AbortEvent) *goja.Object {
	obj := vm.NewObject()

	// Set basic event properties
	obj.Set("type", "abort")
	obj.Set("bubbles", false)
	obj.Set("cancelable", false)
	obj.Set("defaultPrevented", false)
	obj.Set("isTrusted", true)
	obj.Set("timeStamp", event.TimeStamp())

	// Set abort-specific properties
	obj.Set("reason", vm.ToValue(event.reason))

	// Event methods
	obj.Set("preventDefault", func(call goja.FunctionCall) goja.Value {
		event.PreventDefault()
		return goja.Undefined()
	})

	obj.Set("stopPropagation", func(call goja.FunctionCall) goja.Value {
		event.StopPropagation()
		return goja.Undefined()
	})

	obj.Set("stopImmediatePropagation", func(call goja.FunctionCall) goja.Value {
		event.StopImmediatePropagation()
		return goja.Undefined()
	})

	return obj
}

// AbortError represents an error thrown when an operation is aborted
type AbortError struct {
	reason interface{}
}

func (e *AbortError) Error() string {
	if e.reason != nil {
		return e.reason.(string)
	}
	return "The operation was aborted"
}

// IsAbortError checks if an error is an AbortError
func IsAbortError(err error) bool {
	_, ok := err.(*AbortError)
	return ok
}
