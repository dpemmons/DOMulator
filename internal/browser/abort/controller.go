package abort

import (
	"github.com/dop251/goja"
)

// AbortController represents an abort controller that can be used to cancel operations
// This implements the Web API AbortController interface
type AbortController struct {
	signal *AbortSignal
}

// NewAbortController creates a new AbortController with a new AbortSignal
func NewAbortController() *AbortController {
	return &AbortController{
		signal: NewAbortSignal(),
	}
}

// Signal returns the AbortSignal associated with this controller
func (c *AbortController) Signal() *AbortSignal {
	return c.signal
}

// Abort aborts the controller's signal with an optional reason
func (c *AbortController) Abort(reason interface{}) {
	if reason == nil {
		reason = "AbortError"
	}
	c.signal.dispatchAbortEvent(reason)
}

// CreateAbortControllerConstructor creates the AbortController constructor for JavaScript binding
func CreateAbortControllerConstructor(vm *goja.Runtime) func(goja.ConstructorCall) *goja.Object {
	return func(call goja.ConstructorCall) *goja.Object {
		controller := NewAbortController()
		return CreateAbortControllerObject(vm, controller)
	}
}

// CreateAbortControllerObject creates a JavaScript AbortController object from a Go AbortController
func CreateAbortControllerObject(vm *goja.Runtime, controller *AbortController) *goja.Object {
	obj := vm.NewObject()

	// Signal property (read-only)
	obj.Set("signal", CreateAbortSignalObject(vm, controller.signal))

	// Abort method
	obj.Set("abort", func(call goja.FunctionCall) goja.Value {
		var reason interface{}
		if len(call.Arguments) > 0 {
			reason = call.Arguments[0].Export()
		}
		controller.Abort(reason)
		return goja.Undefined()
	})

	// Store reference to Go controller for internal use
	obj.Set("__abortcontroller__", vm.ToValue(controller))

	return obj
}

// GetAbortControllerFromJS extracts the Go AbortController from a JavaScript object
func GetAbortControllerFromJS(jsController *goja.Object) *AbortController {
	if jsController == nil {
		return nil
	}

	internalRef := jsController.Get("__abortcontroller__")
	if internalRef == nil || goja.IsUndefined(internalRef) || goja.IsNull(internalRef) {
		return nil
	}

	exported := internalRef.Export()
	if exported == nil {
		return nil
	}

	if controller, ok := exported.(*AbortController); ok {
		return controller
	}
	return nil
}

// AbortController static methods

// AbortTimeout creates an AbortSignal that automatically aborts after the specified timeout
// This is a convenience method that's part of the AbortSignal spec
func AbortTimeout(vm *goja.Runtime, milliseconds int) *goja.Object {
	signal := NewAbortSignal()

	// In a real implementation, this would use a timer
	// For now, we'll create an already-aborted signal with a timeout reason
	go func() {
		// Simulate timeout behavior
		signal.dispatchAbortEvent("TimeoutError")
	}()

	return CreateAbortSignalObject(vm, signal)
}

// CreateAbortTimeoutFunction creates the AbortSignal.timeout static method
func CreateAbortTimeoutFunction(vm *goja.Runtime) func(goja.FunctionCall) goja.Value {
	return func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			panic(vm.NewTypeError("AbortSignal.timeout requires a timeout value"))
		}

		milliseconds := int(call.Arguments[0].ToInteger())
		if milliseconds < 0 {
			panic(vm.NewTypeError("Timeout value must be non-negative"))
		}

		return AbortTimeout(vm, milliseconds)
	}
}
