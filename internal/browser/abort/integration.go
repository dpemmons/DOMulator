package abort

import (
	"context"
	"errors"

	"github.com/dop251/goja"
)

// FetchIntegration provides methods to integrate AbortSignal with fetch requests
type FetchIntegration struct {
	signal *AbortSignal
}

// NewFetchIntegration creates a new fetch integration with the given signal
func NewFetchIntegration(signal *AbortSignal) *FetchIntegration {
	return &FetchIntegration{
		signal: signal,
	}
}

// CreateContext creates a context that will be cancelled when the signal is aborted
func (fi *FetchIntegration) CreateContext(parent context.Context) (context.Context, context.CancelFunc) {
	if fi.signal == nil {
		return parent, func() {}
	}

	ctx, cancel := context.WithCancel(parent)

	// If already aborted, cancel immediately
	if fi.signal.Aborted() {
		cancel()
		return ctx, cancel
	}

	// Listen for abort events
	fi.signal.AddEventListener("abort", func(event *AbortEvent) {
		cancel()
	})

	return ctx, cancel
}

// CheckAborted checks if the signal has been aborted and returns an appropriate error
func (fi *FetchIntegration) CheckAborted() error {
	if fi.signal != nil && fi.signal.Aborted() {
		reason := fi.signal.Reason()
		if reason == nil {
			reason = "The operation was aborted"
		}
		return &AbortError{reason: reason}
	}
	return nil
}

// ExtractSignalFromOptions extracts an AbortSignal from JavaScript fetch options
func ExtractSignalFromOptions(vm *goja.Runtime, options *goja.Object) *AbortSignal {
	if options == nil {
		return nil
	}

	signalValue := options.Get("signal")
	if signalValue == nil || goja.IsUndefined(signalValue) || goja.IsNull(signalValue) {
		return nil
	}

	signalObj := signalValue.ToObject(vm)
	if signalObj == nil {
		return nil
	}

	return GetAbortSignalFromJS(signalObj)
}

// AbortableError represents an error that can be caused by an abort operation
type AbortableError struct {
	err    error
	reason interface{}
}

func (e *AbortableError) Error() string {
	return e.err.Error()
}

func (e *AbortableError) Unwrap() error {
	return e.err
}

func (e *AbortableError) Reason() interface{} {
	return e.reason
}

// WrapAbortableError wraps an error with abort information
func WrapAbortableError(err error, reason interface{}) error {
	if errors.Is(err, context.Canceled) {
		return &AbortError{reason: reason}
	}
	return &AbortableError{err: err, reason: reason}
}

// SetupAbortAPIs sets up the AbortController and AbortSignal constructors in the JavaScript runtime
func SetupAbortAPIs(vm *goja.Runtime) {
	// AbortController constructor
	vm.Set("AbortController", CreateAbortControllerConstructor(vm))

	// AbortSignal constructor (should throw)
	vm.Set("AbortSignal", CreateAbortSignalConstructor(vm))

	// AbortSignal.timeout static method
	abortSignalConstructor := vm.Get("AbortSignal").ToObject(vm)
	if abortSignalConstructor != nil {
		abortSignalConstructor.Set("timeout", CreateAbortTimeoutFunction(vm))
	}
}
