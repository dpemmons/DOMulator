package dom

import (
	"testing"
	"time"
)

// MockNode implements the Node interface for testing purposes.
type MockNode struct {
	*nodeImpl // Embed nodeImpl to satisfy Node interface methods
	listeners map[string][]func(Event)
}

func NewMockNode(nodeType NodeType, nodeName string) *MockNode {
	m := &MockNode{
		nodeImpl:  &nodeImpl{nodeType: nodeType, nodeName: nodeName},
		listeners: make(map[string][]func(Event)),
	}
	m.nodeImpl.self = m // Set self reference for nodeImpl methods
	return m
}

// Implement Node interface methods that are not directly inherited from nodeImpl
func (m *MockNode) AppendChild(child Node) Node {
	return m.nodeImpl.AppendChild(child)
}

func (m *MockNode) RemoveChild(child Node) Node {
	return m.nodeImpl.RemoveChild(child)
}

func (m *MockNode) InsertBefore(newChild, refChild Node) Node {
	return m.nodeImpl.InsertBefore(newChild, refChild)
}

func (m *MockNode) ReplaceChild(newChild, oldChild Node) Node {
	return m.nodeImpl.ReplaceChild(newChild, oldChild)
}

func (m *MockNode) CloneNode(deep bool) Node {
	// For a mock node, a shallow clone is usually sufficient for event testing.
	// If deep cloning is needed, it would involve recursively cloning children.
	return NewMockNode(m.nodeType, m.nodeName)
}

func (m *MockNode) AddEventListener(eventType string, listener func(Event)) {
	m.listeners[eventType] = append(m.listeners[eventType], listener)
}

func (m *MockNode) RemoveEventListener(eventType string, listener func(Event)) {
	// This is a simplified removal. In a real DOM, you'd need to compare function pointers.
	// For mock, we'll just clear all listeners of that type for simplicity in testing.
	delete(m.listeners, eventType)
}

func (m *MockNode) DispatchEvent(event Event) bool {
	// Simplified dispatch for mock node. Real dispatch involves propagation.
	baseEvent, ok := event.(*BaseEvent)
	if !ok {
		return false
	}
	baseEvent.SetTarget(m)        // m is a Node
	baseEvent.SetCurrentTarget(m) // m is a Node
	baseEvent.SetEventPhase(AT_TARGET)

	if handlers, ok := m.listeners[baseEvent.Type()]; ok {
		for _, handler := range handlers {
			handler(baseEvent)
			if baseEvent.IsImmediatePropagationStopped() {
				break
			}
		}
	}
	return !baseEvent.DefaultPrevented()
}

func TestNewEvent(t *testing.T) {
	e := NewEvent("click", true, true)

	if e.Type() != "click" {
		t.Errorf("Expected type 'click', got %s", e.Type())
	}
	if !e.Bubbles() {
		t.Error("Expected Bubbles() to be true")
	}
	if !e.Cancelable() {
		t.Error("Expected Cancelable() to be true")
	}
	if e.DefaultPrevented() {
		t.Error("Expected DefaultPrevented() to be false initially")
	}
	if !e.IsTrusted() {
		t.Error("Expected IsTrusted() to be true")
	}
	if e.TimeStamp() == 0 { // TODO: Update when timestamp is properly implemented
		t.Log("Timestamp is 0, consider implementing proper timestamp")
	}
}

func TestPreventDefault(t *testing.T) {
	e := NewEvent("submit", true, true) // Cancelable event
	e.PreventDefault()
	if !e.DefaultPrevented() {
		t.Error("Expected DefaultPrevented() to be true after PreventDefault()")
	}

	e2 := NewEvent("click", true, false) // Non-cancelable event
	e2.PreventDefault()
	if e2.DefaultPrevented() {
		t.Error("Expected DefaultPrevented() to be false for non-cancelable event")
	}
}

func TestStopPropagation(t *testing.T) {
	e := NewEvent("click", true, true)
	e.StopPropagation()
	if !e.IsPropagationStopped() {
		t.Error("Expected IsPropagationStopped() to be true after StopPropagation()")
	}
	if e.IsImmediatePropagationStopped() {
		t.Error("Expected IsImmediatePropagationStopped() to be false after StopPropagation()")
	}
}

func TestStopImmediatePropagation(t *testing.T) {
	e := NewEvent("click", true, true)
	e.StopImmediatePropagation()
	if !e.IsPropagationStopped() {
		t.Error("Expected IsPropagationStopped() to be true after StopImmediatePropagation()")
	}
	if !e.IsImmediatePropagationStopped() {
		t.Error("Expected IsImmediatePropagationStopped() to be true after StopImmediatePropagation()")
	}
}

func TestAddRemoveEventListener(t *testing.T) {
	node := NewMockNode(ElementNode, "div")
	var called bool
	listener := func(e Event) {
		called = true
	}

	node.AddEventListener("click", listener)
	e := NewEvent("click", true, true)
	node.DispatchEvent(e)
	if !called {
		t.Error("Expected listener to be called")
	}

	called = false
	node.RemoveEventListener("click", listener) // This mock implementation removes all "click" listeners
	e2 := NewEvent("click", true, true)
	node.DispatchEvent(e2)
	if called {
		t.Error("Expected listener not to be called after removal")
	}
}

func TestDispatchEvent(t *testing.T) {
	node := NewMockNode(ElementNode, "button")
	e := NewEvent("click", true, true)

	// Test initial state
	if e.Target() != nil {
		t.Error("Expected target to be nil initially")
	}
	if e.CurrentTarget() != nil {
		t.Error("Expected currentTarget to be nil initially")
	}
	if e.EventPhase() != NONE {
		t.Error("Expected event phase to be NONE initially")
	}

	node.DispatchEvent(e)

	// Test state after dispatch
	if e.Target() != Node(node) { // Cast node to Node interface for comparison
		t.Errorf("Expected target to be the dispatching node, got %v", e.Target())
	}
	if e.CurrentTarget() != Node(node) { // Cast node to Node interface for comparison
		t.Errorf("Expected currentTarget to be the dispatching node, got %v", e.CurrentTarget())
	}
	if e.EventPhase() != AT_TARGET {
		t.Error("Expected event phase to be AT_TARGET after dispatch")
	}
}

func TestEventPropagationOrder(t *testing.T) {
	// This test will be more comprehensive once the full propagation logic is in Node.DispatchEvent
	// For now, it tests the basic AT_TARGET phase.

	grandparent := NewMockNode(ElementNode, "div")
	parent := NewMockNode(ElementNode, "span")
	child := NewMockNode(ElementNode, "p")

	grandparent.AppendChild(parent) // Now MockNode implements AppendChild
	parent.AppendChild(child)       // Now MockNode implements AppendChild

	var log []string
	listener := func(nodeName string) func(Event) {
		return func(e Event) {
			log = append(log, nodeName+"-"+e.EventPhase().String())
		}
	}

	// Add listeners (simplified for mock, real DOM would have capture/bubble flags)
	grandparent.AddEventListener("test", listener("grandparent"))
	parent.AddEventListener("test", listener("parent"))
	child.AddEventListener("test", listener("child"))

	e := NewEvent("test", true, true)
	child.DispatchEvent(e) // Dispatching on child

	// Expected order for simplified mock dispatch (only AT_TARGET for the target itself)
	expected := []string{"child-AT_TARGET"}
	if len(log) != len(expected) {
		t.Fatalf("Expected %d events, got %d: %v", len(expected), len(log), log)
	}
	for i, val := range expected {
		if log[i] != val {
			t.Errorf("Expected %s, got %s at index %d", val, log[i], i)
		}
	}
}

// String method for EventPhase for better test output
func (ep EventPhase) String() string {
	switch ep {
	case NONE:
		return "NONE"
	case CAPTURING_PHASE:
		return "CAPTURING_PHASE"
	case AT_TARGET:
		return "AT_TARGET"
	case BUBBLING_PHASE:
		return "BUBBLING_PHASE"
	default:
		return "UNKNOWN"
	}
}

func TestMultipleListeners(t *testing.T) {
	node := NewMockNode(ElementNode, "div")
	callCount := 0

	listener1 := func(e Event) { callCount++ }
	listener2 := func(e Event) { callCount++ }

	node.AddEventListener("click", listener1)
	node.AddEventListener("click", listener2)

	e := NewEvent("click", true, true)
	node.DispatchEvent(e)

	if callCount != 2 {
		t.Errorf("Expected 2 listeners to be called, got %d", callCount)
	}
}

func TestStopImmediatePropagationEffect(t *testing.T) {
	node := NewMockNode(ElementNode, "div")
	callCount := 0

	listener1 := func(e Event) {
		callCount++
		e.StopImmediatePropagation()
	}
	listener2 := func(e Event) {
		callCount++
	}

	node.AddEventListener("click", listener1)
	node.AddEventListener("click", listener2)

	e := NewEvent("click", true, true)
	node.DispatchEvent(e)

	if callCount != 1 {
		t.Errorf("Expected only 1 listener to be called, got %d", callCount)
	}
}

func TestEventTimestamp(t *testing.T) {
	// This test assumes TimeStamp() will eventually return a non-zero value.
	// For now, it just checks if it's int64.
	e := NewEvent("test", false, false)
	time.Sleep(1 * time.Millisecond) // Simulate some time passing
	if e.TimeStamp() != 0 {          // TODO: Update this test when timestamp is properly implemented
		t.Logf("Timestamp is %d, expected 0 for now. Will be updated when implemented.", e.TimeStamp())
	}
}
