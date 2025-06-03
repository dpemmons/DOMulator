package events

import (
	"github.com/dop251/goja"
	"github.com/dpemmons/DOMulator/internal/dom"
)

// CustomEvent represents a custom event that can carry additional data
// This implements the Web API CustomEvent interface
type CustomEvent struct {
	*dom.BaseEvent
	detail interface{}
}

// NewCustomEvent creates a new CustomEvent with the specified type, options, and detail data
func NewCustomEvent(eventType string, options *CustomEventInit) *CustomEvent {
	bubbles := false
	cancelable := false
	var detail interface{}

	if options != nil {
		bubbles = options.Bubbles
		cancelable = options.Cancelable
		detail = options.Detail
	}

	baseEvent := dom.NewEvent(eventType, bubbles, cancelable)

	return &CustomEvent{
		BaseEvent: baseEvent,
		detail:    detail,
	}
}

// CustomEventInit represents the initialization options for a CustomEvent
type CustomEventInit struct {
	Bubbles    bool
	Cancelable bool
	Detail     interface{}
}

// Detail returns the custom data associated with the event
func (ce *CustomEvent) Detail() interface{} {
	return ce.detail
}

// SetDetail sets the custom data for the event
func (ce *CustomEvent) SetDetail(detail interface{}) {
	ce.detail = detail
}

// CreateCustomEventConstructor creates the CustomEvent constructor for JavaScript binding
func CreateCustomEventConstructor(vm *goja.Runtime) func(goja.ConstructorCall) *goja.Object {
	return func(call goja.ConstructorCall) *goja.Object {
		if len(call.Arguments) < 1 {
			panic(vm.NewTypeError("CustomEvent constructor requires at least 1 argument"))
		}

		eventType := call.Arguments[0].String()
		var options *CustomEventInit

		// Parse options if provided
		if len(call.Arguments) > 1 && !goja.IsUndefined(call.Arguments[1]) {
			options = parseCustomEventInit(vm, call.Arguments[1])
		}

		customEvent := NewCustomEvent(eventType, options)
		obj := vm.NewObject()

		// Set up the CustomEvent properties and methods
		setupCustomEventObject(vm, obj, customEvent)

		return obj
	}
}

// parseCustomEventInit parses JavaScript CustomEventInit options into Go struct
func parseCustomEventInit(vm *goja.Runtime, value goja.Value) *CustomEventInit {
	options := &CustomEventInit{
		Bubbles:    false,
		Cancelable: false,
		Detail:     nil,
	}

	if value == nil || goja.IsNull(value) || goja.IsUndefined(value) {
		return options
	}

	// Safely convert to object
	var obj *goja.Object
	if value != nil {
		obj = value.ToObject(vm)
	}

	if obj == nil {
		return options
	}

	// Parse bubbles
	if bubblesVal := obj.Get("bubbles"); bubblesVal != nil && !goja.IsUndefined(bubblesVal) {
		options.Bubbles = bubblesVal.ToBoolean()
	}

	// Parse cancelable
	if cancelableVal := obj.Get("cancelable"); cancelableVal != nil && !goja.IsUndefined(cancelableVal) {
		options.Cancelable = cancelableVal.ToBoolean()
	}

	// Parse detail
	if detailVal := obj.Get("detail"); detailVal != nil && !goja.IsUndefined(detailVal) {
		options.Detail = detailVal.Export()
	}

	return options
}

// setupCustomEventObject sets up the JavaScript CustomEvent object with all properties and methods
func setupCustomEventObject(vm *goja.Runtime, obj *goja.Object, customEvent *CustomEvent) {
	// Set basic event properties
	obj.Set("type", customEvent.Type())
	obj.Set("bubbles", customEvent.Bubbles())
	obj.Set("cancelable", customEvent.Cancelable())
	obj.Set("detail", vm.ToValue(customEvent.Detail()))

	// Set event state properties
	obj.Set("defaultPrevented", customEvent.DefaultPrevented())
	obj.Set("isTrusted", customEvent.IsTrusted())
	obj.Set("timeStamp", customEvent.TimeStamp())

	// Set target and currentTarget (these will be set during dispatch)
	obj.Set("target", goja.Null())
	obj.Set("currentTarget", goja.Null())

	// Set event phase constants
	obj.Set("NONE", dom.NONE)
	obj.Set("CAPTURING_PHASE", dom.CAPTURING_PHASE)
	obj.Set("AT_TARGET", dom.AT_TARGET)
	obj.Set("BUBBLING_PHASE", dom.BUBBLING_PHASE)

	// Set current event phase
	obj.Set("eventPhase", customEvent.EventPhase())

	// Set up event methods
	obj.Set("preventDefault", func(call goja.FunctionCall) goja.Value {
		customEvent.PreventDefault()
		obj.Set("defaultPrevented", customEvent.DefaultPrevented())
		return goja.Undefined()
	})

	obj.Set("stopPropagation", func(call goja.FunctionCall) goja.Value {
		customEvent.StopPropagation()
		return goja.Undefined()
	})

	obj.Set("stopImmediatePropagation", func(call goja.FunctionCall) goja.Value {
		customEvent.StopImmediatePropagation()
		return goja.Undefined()
	})

	// initCustomEvent method (legacy, but part of the spec)
	obj.Set("initCustomEvent", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 4 {
			panic(vm.NewTypeError("initCustomEvent requires 4 arguments"))
		}

		eventType := call.Arguments[0].String()
		bubbles := call.Arguments[1].ToBoolean()
		cancelable := call.Arguments[2].ToBoolean()
		detail := call.Arguments[3].Export()

		// Create new custom event with the provided parameters
		newEvent := NewCustomEvent(eventType, &CustomEventInit{
			Bubbles:    bubbles,
			Cancelable: cancelable,
			Detail:     detail,
		})

		// Update the object with new values
		setupCustomEventObject(vm, obj, newEvent)

		return goja.Undefined()
	})

	// Add internal reference to Go CustomEvent for other APIs to use
	obj.Set("__customevent__", vm.ToValue(customEvent))
}

// CreateCustomEventFromJS creates a CustomEvent from a JavaScript event object
func CreateCustomEventFromJS(vm *goja.Runtime, jsEvent *goja.Object) *CustomEvent {
	// Extract the internal CustomEvent reference if it exists
	if internalRef := jsEvent.Get("__customevent__"); !goja.IsUndefined(internalRef) {
		if customEvent, ok := internalRef.Export().(*CustomEvent); ok {
			return customEvent
		}
	}

	// If no internal reference, create a new CustomEvent from the JS object properties
	eventType := jsEvent.Get("type").String()
	bubbles := jsEvent.Get("bubbles").ToBoolean()
	cancelable := jsEvent.Get("cancelable").ToBoolean()
	detail := jsEvent.Get("detail").Export()

	return NewCustomEvent(eventType, &CustomEventInit{
		Bubbles:    bubbles,
		Cancelable: cancelable,
		Detail:     detail,
	})
}

// UpdateJSEventFromCustomEvent updates a JavaScript event object with current CustomEvent state
func UpdateJSEventFromCustomEvent(vm *goja.Runtime, jsEvent *goja.Object, customEvent *CustomEvent) {
	// Update dynamic properties that can change during event lifecycle
	jsEvent.Set("target", vm.ToValue(customEvent.Target()))
	jsEvent.Set("currentTarget", vm.ToValue(customEvent.CurrentTarget()))
	jsEvent.Set("eventPhase", customEvent.EventPhase())
	jsEvent.Set("defaultPrevented", customEvent.DefaultPrevented())
}

// GetDetailAsObject safely extracts the detail property as a JavaScript object
func GetDetailAsObject(vm *goja.Runtime, detail interface{}) goja.Value {
	if detail == nil {
		return goja.Null()
	}
	return vm.ToValue(detail)
}
