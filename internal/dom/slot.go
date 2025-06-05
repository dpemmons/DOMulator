package dom

// Slottable interface represents elements and text nodes that can be assigned to slots
// Element and Text nodes are slottables as per WHATWG DOM Section 4.2.2.2
type Slottable interface {
	Node
	GetSlotName() string
	SetSlotName(string)
	GetAssignedSlot() *Slot
	SetAssignedSlot(*Slot)
	GetManualSlot() *Slot
	SetManualSlot(*Slot)
}

// Slot represents a slot element as defined in WHATWG DOM Section 4.2.2.1
type Slot struct {
	*Element
	slotName      string      // The slot's name
	assignedNodes []Slottable // Currently assigned slottables
}

// NewSlot creates a new slot element
func NewSlot(doc *Document) *Slot {
	// Create the underlying element with tag name "slot"
	elem := NewElement("slot", doc)

	slot := &Slot{
		Element:       elem,
		slotName:      "", // Default to empty string
		assignedNodes: make([]Slottable, 0),
	}

	// Override the self reference to point to the slot
	elem.nodeImpl.self = slot

	return slot
}

// GetSlotName returns the slot's name
func (s *Slot) GetSlotName() string {
	return s.slotName
}

// SetSlotName sets the slot's name and updates the name attribute
func (s *Slot) SetSlotName(name string) {
	s.slotName = name
	if name == "" {
		s.RemoveAttribute("name")
	} else {
		s.SetAttribute("name", name)
	}

	// Run assign slottables for a tree with element's root
	if root := s.GetRootNode(nil); root != nil {
		if shadowRoot, ok := root.(*ShadowRoot); ok {
			shadowRoot.AssignSlottablesForTree(root)
		}
	}
}

// GetAssignedNodes returns the slot's currently assigned slottables
func (s *Slot) GetAssignedNodes() []Slottable {
	// Return a copy to prevent external modification
	result := make([]Slottable, len(s.assignedNodes))
	copy(result, s.assignedNodes)
	return result
}

// SetAssignedNodes sets the slot's assigned slottables
func (s *Slot) SetAssignedNodes(slottables []Slottable) {
	// Make a copy to prevent external modification
	s.assignedNodes = make([]Slottable, len(slottables))
	copy(s.assignedNodes, slottables)
}

// Implement Slottable interface for Slot (a slot can be a slottable)
func (s *Slot) GetAssignedSlot() *Slot {
	// Element already implements Slottable, so we can directly call its methods
	return s.Element.GetAssignedSlot()
}

func (s *Slot) SetAssignedSlot(slot *Slot) {
	// Element already implements Slottable, so we can directly call its methods
	s.Element.SetAssignedSlot(slot)
}

func (s *Slot) GetManualSlot() *Slot {
	// Element already implements Slottable, so we can directly call its methods
	return s.Element.GetManualSlot()
}

func (s *Slot) SetManualSlot(slot *Slot) {
	// Element already implements Slottable, so we can directly call its methods
	s.Element.SetManualSlot(slot)
}

// handleSlotNameChange implements slot attribute change steps from WHATWG DOM Section 4.2.2.1
func (s *Slot) handleSlotNameChange(oldValue, newValue string) {
	// If value is oldValue, then return
	if newValue == oldValue {
		return
	}

	// If value is null and oldValue is the empty string, then return
	if newValue == "" && oldValue == "" {
		return
	}

	// If value is the empty string and oldValue is null, then return
	if newValue == "" && oldValue == "" {
		return
	}

	// If value is null or the empty string, then set element's name to the empty string
	if newValue == "" {
		s.slotName = ""
	} else {
		// Otherwise, set element's name to value
		s.slotName = newValue
	}

	// Run assign slottables for a tree with element's root
	if root := s.GetRootNode(nil); root != nil {
		if shadowRoot, ok := root.(*ShadowRoot); ok {
			shadowRoot.AssignSlottablesForTree(root)
		}
	}
}

// Override SetAttribute to handle name attribute changes
func (s *Slot) SetAttribute(name, value string) {
	if name == "name" {
		oldValue := s.GetAttribute("name")
		s.Element.SetAttribute(name, value)
		s.handleSlotNameChange(oldValue, value)
	} else {
		s.Element.SetAttribute(name, value)
	}
}

// Override RemoveAttribute to handle name attribute removal
func (s *Slot) RemoveAttribute(name string) {
	if name == "name" {
		oldValue := s.GetAttribute("name")
		s.Element.RemoveAttribute(name)
		s.handleSlotNameChange(oldValue, "")
	} else {
		s.Element.RemoveAttribute(name)
	}
}
