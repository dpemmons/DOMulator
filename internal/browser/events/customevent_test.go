package events

import (
	"testing"

	"github.com/dop251/goja"
	"github.com/dpemmons/DOMulator/internal/dom"
)

func TestCustomEvent_BasicCreation(t *testing.T) {
	// Test with minimal options
	event := NewCustomEvent("test-event", nil)

	if event.Type() != "test-event" {
		t.Errorf("Expected type 'test-event', got '%s'", event.Type())
	}

	if event.Bubbles() {
		t.Error("Expected bubbles to be false by default")
	}

	if event.Cancelable() {
		t.Error("Expected cancelable to be false by default")
	}

	if event.Detail() != nil {
		t.Errorf("Expected detail to be nil, got %v", event.Detail())
	}

	if event.DefaultPrevented() {
		t.Error("Expected defaultPrevented to be false initially")
	}

	if !event.IsTrusted() {
		t.Error("Expected isTrusted to be true for system-created events")
	}
}

func TestCustomEvent_WithStringDetail(t *testing.T) {
	detail := "test string detail"

	options := &CustomEventInit{
		Bubbles:    true,
		Cancelable: true,
		Detail:     detail,
	}

	event := NewCustomEvent("custom-event", options)

	if event.Type() != "custom-event" {
		t.Errorf("Expected type 'custom-event', got '%s'", event.Type())
	}

	if !event.Bubbles() {
		t.Error("Expected bubbles to be true")
	}

	if !event.Cancelable() {
		t.Error("Expected cancelable to be true")
	}

	if event.Detail() != detail {
		t.Errorf("Expected detail '%s', got %v", detail, event.Detail())
	}
}

func TestCustomEvent_EventMethods(t *testing.T) {
	options := &CustomEventInit{
		Bubbles:    true,
		Cancelable: true,
		Detail:     "test data",
	}

	event := NewCustomEvent("test", options)

	// Test preventDefault
	if event.DefaultPrevented() {
		t.Error("Expected defaultPrevented to be false initially")
	}

	event.PreventDefault()
	if !event.DefaultPrevented() {
		t.Error("Expected defaultPrevented to be true after calling preventDefault")
	}

	// Test stopPropagation
	if event.IsPropagationStopped() {
		t.Error("Expected propagation not to be stopped initially")
	}

	event.StopPropagation()
	if !event.IsPropagationStopped() {
		t.Error("Expected propagation to be stopped after calling stopPropagation")
	}

	// Test stopImmediatePropagation
	event2 := NewCustomEvent("test2", options)

	if event2.IsImmediatePropagationStopped() {
		t.Error("Expected immediate propagation not to be stopped initially")
	}

	event2.StopImmediatePropagation()
	if !event2.IsImmediatePropagationStopped() {
		t.Error("Expected immediate propagation to be stopped")
	}

	if !event2.IsPropagationStopped() {
		t.Error("Expected propagation to also be stopped when immediate propagation is stopped")
	}
}

func TestCustomEvent_EventPhases(t *testing.T) {
	event := NewCustomEvent("test", nil)

	// Initial phase should be NONE
	if event.EventPhase() != dom.NONE {
		t.Errorf("Expected initial phase to be NONE (%d), got %d", dom.NONE, event.EventPhase())
	}

	// Test setting different phases
	event.SetEventPhase(dom.CAPTURING_PHASE)
	if event.EventPhase() != dom.CAPTURING_PHASE {
		t.Errorf("Expected phase to be CAPTURING_PHASE (%d), got %d", dom.CAPTURING_PHASE, event.EventPhase())
	}

	event.SetEventPhase(dom.AT_TARGET)
	if event.EventPhase() != dom.AT_TARGET {
		t.Errorf("Expected phase to be AT_TARGET (%d), got %d", dom.AT_TARGET, event.EventPhase())
	}

	event.SetEventPhase(dom.BUBBLING_PHASE)
	if event.EventPhase() != dom.BUBBLING_PHASE {
		t.Errorf("Expected phase to be BUBBLING_PHASE (%d), got %d", dom.BUBBLING_PHASE, event.EventPhase())
	}
}

func TestCustomEvent_DetailProperty(t *testing.T) {
	// Test with string detail
	event1 := NewCustomEvent("test", &CustomEventInit{Detail: "string data"})
	if event1.Detail() != "string data" {
		t.Errorf("Expected string detail, got %v", event1.Detail())
	}

	// Test with number detail
	event2 := NewCustomEvent("test", &CustomEventInit{Detail: 42})
	if event2.Detail() != 42 {
		t.Errorf("Expected number detail, got %v", event2.Detail())
	}

	// Test SetDetail method
	event3 := NewCustomEvent("test", nil)
	event3.SetDetail("new detail")
	if event3.Detail() != "new detail" {
		t.Errorf("Expected updated detail, got %v", event3.Detail())
	}
}

func TestCustomEvent_ObjectDetail(t *testing.T) {
	// Test with object detail - just verify it stores and retrieves
	obj := map[string]interface{}{"key": "value", "num": 123}
	event := NewCustomEvent("test", &CustomEventInit{Detail: obj})

	// Just verify it's not nil and is the same type
	detail := event.Detail()
	if detail == nil {
		t.Error("Expected detail to not be nil")
	}

	if detailMap, ok := detail.(map[string]interface{}); ok {
		if detailMap["key"] != "value" {
			t.Errorf("Expected key 'value', got %v", detailMap["key"])
		}
		if detailMap["num"] != 123 {
			t.Errorf("Expected num 123, got %v", detailMap["num"])
		}
	} else {
		t.Error("Expected detail to be a map[string]interface{}")
	}
}

func TestCustomEvent_JavaScriptConstructor(t *testing.T) {
	vm := goja.New()
	constructor := CreateCustomEventConstructor(vm)

	// Test basic constructor call
	args := []goja.Value{vm.ToValue("test-event")}
	call := goja.ConstructorCall{Arguments: args}
	obj := constructor(call)

	if obj == nil {
		t.Fatal("CustomEvent constructor returned nil")
	}

	// Test event properties
	if obj.Get("type").String() != "test-event" {
		t.Errorf("Expected type 'test-event', got '%s'", obj.Get("type").String())
	}

	if obj.Get("bubbles").ToBoolean() {
		t.Error("Expected bubbles to be false by default")
	}

	if obj.Get("cancelable").ToBoolean() {
		t.Error("Expected cancelable to be false by default")
	}

	if !goja.IsNull(obj.Get("detail")) {
		t.Error("Expected detail to be null by default")
	}
}

func TestCustomEvent_JavaScriptConstructorWithOptions(t *testing.T) {
	vm := goja.New()
	constructor := CreateCustomEventConstructor(vm)

	// Create options object
	options := vm.NewObject()
	options.Set("bubbles", true)
	options.Set("cancelable", true)
	options.Set("detail", vm.ToValue(map[string]interface{}{"msg": "hello"}))

	args := []goja.Value{vm.ToValue("custom-event"), options}
	call := goja.ConstructorCall{Arguments: args}
	obj := constructor(call)

	// Test properties
	if obj.Get("type").String() != "custom-event" {
		t.Errorf("Expected type 'custom-event', got '%s'", obj.Get("type").String())
	}

	if !obj.Get("bubbles").ToBoolean() {
		t.Error("Expected bubbles to be true")
	}

	if !obj.Get("cancelable").ToBoolean() {
		t.Error("Expected cancelable to be true")
	}

	// Test detail object
	detail := obj.Get("detail").ToObject(vm)
	if detail.Get("msg").String() != "hello" {
		t.Error("Expected detail to contain correct data")
	}
}

func TestCustomEvent_JavaScriptMethods(t *testing.T) {
	vm := goja.New()
	constructor := CreateCustomEventConstructor(vm)

	options := vm.NewObject()
	options.Set("cancelable", true)

	args := []goja.Value{vm.ToValue("test"), options}
	call := goja.ConstructorCall{Arguments: args}
	obj := constructor(call)

	// Test preventDefault method
	preventDefaultFn, exists := goja.AssertFunction(obj.Get("preventDefault"))
	if !exists {
		t.Fatal("CustomEvent object missing preventDefault method")
	}

	// Initially not prevented
	if obj.Get("defaultPrevented").ToBoolean() {
		t.Error("Expected defaultPrevented to be false initially")
	}

	// Call preventDefault
	_, err := preventDefaultFn(goja.Undefined())
	if err != nil {
		t.Errorf("Failed to call preventDefault: %v", err)
	}

	// Should be prevented now
	if !obj.Get("defaultPrevented").ToBoolean() {
		t.Error("Expected defaultPrevented to be true after calling preventDefault")
	}

	// Test stopPropagation method
	stopPropFn, exists := goja.AssertFunction(obj.Get("stopPropagation"))
	if !exists {
		t.Fatal("CustomEvent object missing stopPropagation method")
	}

	_, err = stopPropFn(goja.Undefined())
	if err != nil {
		t.Errorf("Failed to call stopPropagation: %v", err)
	}

	// Test stopImmediatePropagation method
	stopImmFn, exists := goja.AssertFunction(obj.Get("stopImmediatePropagation"))
	if !exists {
		t.Fatal("CustomEvent object missing stopImmediatePropagation method")
	}

	_, err = stopImmFn(goja.Undefined())
	if err != nil {
		t.Errorf("Failed to call stopImmediatePropagation: %v", err)
	}
}

func TestCustomEvent_InitCustomEvent(t *testing.T) {
	vm := goja.New()
	constructor := CreateCustomEventConstructor(vm)

	args := []goja.Value{vm.ToValue("initial")}
	call := goja.ConstructorCall{Arguments: args}
	obj := constructor(call)

	// Test initCustomEvent method
	initFn, exists := goja.AssertFunction(obj.Get("initCustomEvent"))
	if !exists {
		t.Fatal("CustomEvent object missing initCustomEvent method")
	}

	// Call initCustomEvent
	_, err := initFn(goja.Undefined(),
		vm.ToValue("new-type"),
		vm.ToValue(true), // bubbles
		vm.ToValue(true), // cancelable
		vm.ToValue("new detail"))
	if err != nil {
		t.Errorf("Failed to call initCustomEvent: %v", err)
	}

	// Check updated properties
	if obj.Get("type").String() != "new-type" {
		t.Errorf("Expected type 'new-type', got '%s'", obj.Get("type").String())
	}

	if !obj.Get("bubbles").ToBoolean() {
		t.Error("Expected bubbles to be true after initCustomEvent")
	}

	if !obj.Get("cancelable").ToBoolean() {
		t.Error("Expected cancelable to be true after initCustomEvent")
	}

	if obj.Get("detail").String() != "new detail" {
		t.Errorf("Expected detail 'new detail', got '%s'", obj.Get("detail").String())
	}
}

func TestCustomEvent_EventPhaseConstants(t *testing.T) {
	vm := goja.New()
	constructor := CreateCustomEventConstructor(vm)

	args := []goja.Value{vm.ToValue("test")}
	call := goja.ConstructorCall{Arguments: args}
	obj := constructor(call)

	// Test event phase constants
	if obj.Get("NONE").ToInteger() != int64(dom.NONE) {
		t.Errorf("Expected NONE constant to be %d, got %d", dom.NONE, obj.Get("NONE").ToInteger())
	}

	if obj.Get("CAPTURING_PHASE").ToInteger() != int64(dom.CAPTURING_PHASE) {
		t.Errorf("Expected CAPTURING_PHASE constant to be %d, got %d", dom.CAPTURING_PHASE, obj.Get("CAPTURING_PHASE").ToInteger())
	}

	if obj.Get("AT_TARGET").ToInteger() != int64(dom.AT_TARGET) {
		t.Errorf("Expected AT_TARGET constant to be %d, got %d", dom.AT_TARGET, obj.Get("AT_TARGET").ToInteger())
	}

	if obj.Get("BUBBLING_PHASE").ToInteger() != int64(dom.BUBBLING_PHASE) {
		t.Errorf("Expected BUBBLING_PHASE constant to be %d, got %d", dom.BUBBLING_PHASE, obj.Get("BUBBLING_PHASE").ToInteger())
	}
}

func TestCustomEvent_JavaScriptIntegration(t *testing.T) {
	vm := goja.New()

	// Register CustomEvent constructor
	vm.Set("CustomEvent", CreateCustomEventConstructor(vm))

	// Test JavaScript code
	script := `
		// Test basic creation
		var event1 = new CustomEvent('test1');
		
		// Test with options
		var event2 = new CustomEvent('test2', {
			bubbles: true,
			cancelable: true,
			detail: {message: 'Hello', count: 42}
		});
		
		// Test methods
		event2.preventDefault();
		event2.stopPropagation();
		
		// Test initCustomEvent
		var event3 = new CustomEvent('old');
		event3.initCustomEvent('new', false, false, 'updated');
		
		// Return results
		({
			event1Type: event1.type,
			event1Bubbles: event1.bubbles,
			event1Detail: event1.detail,
			
			event2Type: event2.type,
			event2Bubbles: event2.bubbles,
			event2Cancelable: event2.cancelable,
			event2DetailMessage: event2.detail.message,
			event2DetailCount: event2.detail.count,
			event2DefaultPrevented: event2.defaultPrevented,
			
			event3Type: event3.type,
			event3Bubbles: event3.bubbles,
			event3Detail: event3.detail,
			
			// Test constants
			NONE: CustomEvent.prototype.NONE || event1.NONE,
			CAPTURING_PHASE: event1.CAPTURING_PHASE,
			AT_TARGET: event1.AT_TARGET,
			BUBBLING_PHASE: event1.BUBBLING_PHASE
		})
	`

	result, err := vm.RunString(script)
	if err != nil {
		t.Fatalf("JavaScript execution failed: %v", err)
	}

	obj := result.ToObject(vm)

	// Test event1 results
	if obj.Get("event1Type").String() != "test1" {
		t.Errorf("Expected event1 type 'test1', got '%s'", obj.Get("event1Type").String())
	}

	if obj.Get("event1Bubbles").ToBoolean() {
		t.Error("Expected event1 bubbles to be false")
	}

	if !goja.IsNull(obj.Get("event1Detail")) {
		t.Error("Expected event1 detail to be null")
	}

	// Test event2 results
	if obj.Get("event2Type").String() != "test2" {
		t.Errorf("Expected event2 type 'test2', got '%s'", obj.Get("event2Type").String())
	}

	if !obj.Get("event2Bubbles").ToBoolean() {
		t.Error("Expected event2 bubbles to be true")
	}

	if !obj.Get("event2Cancelable").ToBoolean() {
		t.Error("Expected event2 cancelable to be true")
	}

	if obj.Get("event2DetailMessage").String() != "Hello" {
		t.Errorf("Expected detail message 'Hello', got '%s'", obj.Get("event2DetailMessage").String())
	}

	if obj.Get("event2DetailCount").ToInteger() != 42 {
		t.Errorf("Expected detail count 42, got %d", obj.Get("event2DetailCount").ToInteger())
	}

	if !obj.Get("event2DefaultPrevented").ToBoolean() {
		t.Error("Expected event2 to have default prevented")
	}

	// Test event3 (initCustomEvent) results
	if obj.Get("event3Type").String() != "new" {
		t.Errorf("Expected event3 type 'new', got '%s'", obj.Get("event3Type").String())
	}

	if obj.Get("event3Detail").String() != "updated" {
		t.Errorf("Expected event3 detail 'updated', got '%s'", obj.Get("event3Detail").String())
	}
}

func TestCustomEvent_ErrorHandling(t *testing.T) {
	vm := goja.New()
	constructor := CreateCustomEventConstructor(vm)

	// Test missing required argument
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic when creating CustomEvent without type")
		}
	}()

	args := []goja.Value{}
	call := goja.ConstructorCall{Arguments: args}
	constructor(call)
}

func TestCustomEvent_EdgeCases(t *testing.T) {
	// Test with nil options
	event1 := NewCustomEvent("test", nil)
	if event1.Type() != "test" {
		t.Errorf("Expected type 'test', got '%s'", event1.Type())
	}

	// Test with empty detail
	event2 := NewCustomEvent("test", &CustomEventInit{Detail: ""})
	if event2.Detail() != "" {
		t.Errorf("Expected empty string detail, got %v", event2.Detail())
	}

	// Test with complex detail object - just verify it's preserved
	complexDetail := map[string]interface{}{
		"nested": map[string]interface{}{
			"value": 123,
		},
		"array": []interface{}{1, 2, 3},
	}
	event3 := NewCustomEvent("test", &CustomEventInit{Detail: complexDetail})

	// Verify it's not nil and has the right type
	detail := event3.Detail()
	if detail == nil {
		t.Error("Expected complex detail to be preserved")
	}

	if detailMap, ok := detail.(map[string]interface{}); ok {
		if nested, ok := detailMap["nested"].(map[string]interface{}); ok {
			if nested["value"] != 123 {
				t.Errorf("Expected nested value 123, got %v", nested["value"])
			}
		} else {
			t.Error("Expected nested map in detail")
		}
	} else {
		t.Error("Expected detail to be a map")
	}
}

func TestParseCustomEventInit_EdgeCases(t *testing.T) {
	vm := goja.New()

	// Test with null value
	options1 := parseCustomEventInit(vm, goja.Null())
	if options1.Bubbles || options1.Cancelable || options1.Detail != nil {
		t.Error("Expected default values for null input")
	}

	// Test with undefined value
	options2 := parseCustomEventInit(vm, goja.Undefined())
	if options2.Bubbles || options2.Cancelable || options2.Detail != nil {
		t.Error("Expected default values for undefined input")
	}

	// Test with empty object
	emptyObj := vm.NewObject()
	options3 := parseCustomEventInit(vm, emptyObj)
	if options3.Bubbles || options3.Cancelable || options3.Detail != nil {
		t.Error("Expected default values for empty object")
	}
}
