package dom

// EventPhase constants define the current phase of the event flow.
type EventPhase int

const (
	NONE            EventPhase = 0
	CAPTURING_PHASE EventPhase = 1 // The event is currently in the capturing phase.
	AT_TARGET       EventPhase = 2 // The event is currently at the event's target.
	BUBBLING_PHASE  EventPhase = 3 // The event is currently in the bubbling phase.
)

// Event represents an event which takes place in the DOM.
// It is the base interface for any event.
type Event interface {
	Type() string
	Target() Node
	CurrentTarget() Node
	EventPhase() EventPhase
	Bubbles() bool
	Cancelable() bool
	DefaultPrevented() bool
	IsTrusted() bool
	TimeStamp() int64
	PreventDefault()
	StopPropagation()
	StopImmediatePropagation()
}

// BaseEvent is a concrete implementation of the Event interface,
// providing common fields and methods for all events.
type BaseEvent struct {
	eventType          string
	eventTarget        Node
	eventCurrentTarget Node
	eventPhase         EventPhase
	bubblesFlag        bool
	cancelableFlag     bool
	defaultPrevented   bool
	isTrustedFlag      bool
	timeStamp          int64
	propagationStopped bool
	immediateStopped   bool
}

// NewEvent creates a new BaseEvent with the given type and options.
func NewEvent(eventType string, bubbles, cancelable bool) *BaseEvent {
	return &BaseEvent{
		eventType:      eventType,
		bubblesFlag:    bubbles,
		cancelableFlag: cancelable,
		isTrustedFlag:  true, // Events created by the system are trusted
		timeStamp:      0,    // TODO: Implement proper timestamp
	}
}

// Type returns the name of the event.
func (e *BaseEvent) Type() string {
	return e.eventType
}

// Target returns the object to which the event was dispatched.
func (e *BaseEvent) Target() Node {
	return e.eventTarget
}

// CurrentTarget returns the object whose EventListener is currently being processed.
func (e *BaseEvent) CurrentTarget() Node {
	return e.eventCurrentTarget
}

// EventPhase returns the current phase of the event flow.
func (e *BaseEvent) EventPhase() EventPhase {
	return e.eventPhase
}

// Bubbles returns a boolean indicating whether the event bubbles up through the DOM or not.
func (e *BaseEvent) Bubbles() bool {
	return e.bubblesFlag
}

// Cancelable returns a boolean indicating whether the event is cancelable.
func (e *BaseEvent) Cancelable() bool {
	return e.cancelableFlag
}

// DefaultPrevented returns a boolean indicating whether the default action of the event was prevented.
func (e *BaseEvent) DefaultPrevented() bool {
	return e.defaultPrevented
}

// IsTrusted returns a boolean indicating whether the event was created by a user agent (true)
// or by a script (false).
func (e *BaseEvent) IsTrusted() bool {
	return e.isTrustedFlag
}

// TimeStamp returns the time at which the event was created.
func (e *BaseEvent) TimeStamp() int64 {
	return e.timeStamp
}

// PreventDefault prevents the default action of the event.
func (e *BaseEvent) PreventDefault() {
	if e.cancelableFlag {
		e.defaultPrevented = true
	}
}

// StopPropagation prevents further propagation of the current event.
func (e *BaseEvent) StopPropagation() {
	e.propagationStopped = true
}

// StopImmediatePropagation prevents other listeners of the same event from being called.
func (e *BaseEvent) StopImmediatePropagation() {
	e.propagationStopped = true
	e.immediateStopped = true
}

// SetTarget sets the event's target. This is for internal use during event dispatch.
func (e *BaseEvent) SetTarget(n Node) {
	e.eventTarget = n
}

// SetCurrentTarget sets the event's current target. This is for internal use during event dispatch.
func (e *BaseEvent) SetCurrentTarget(n Node) {
	e.eventCurrentTarget = n
}

// SetEventPhase sets the event's phase. This is for internal use during event dispatch.
func (e *BaseEvent) SetEventPhase(phase EventPhase) {
	e.eventPhase = phase
}

// IsPropagationStopped returns true if StopPropagation or StopImmediatePropagation was called.
func (e *BaseEvent) IsPropagationStopped() bool {
	return e.propagationStopped
}

// IsImmediatePropagationStopped returns true if StopImmediatePropagation was called.
func (e *BaseEvent) IsImmediatePropagationStopped() bool {
	return e.immediateStopped
}
