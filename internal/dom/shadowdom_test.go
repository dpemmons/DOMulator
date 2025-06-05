package dom

import (
	"testing"
)

func TestShadowDOMBasics(t *testing.T) {
	// Create a document
	doc := NewDocument()

	// Create a div element (can have shadow root)
	host := doc.CreateElement("div")

	// Initially, element should not have a shadow root
	if host.ShadowRoot() != nil {
		t.Error("Element should not have a shadow root initially")
	}

	// Create shadow root init
	init := ShadowRootInit{
		Mode:           ShadowRootModeOpen,
		SlotAssignment: SlotAssignmentNamed,
	}

	// Attach shadow root
	shadowRoot, err := host.AttachShadow(init)
	if err != nil {
		t.Fatalf("Failed to attach shadow root: %v", err)
	}

	// Verify shadow root is attached
	if host.ShadowRoot() != shadowRoot {
		t.Error("Shadow root not properly attached to host")
	}

	// Verify shadow root properties
	if shadowRoot.Host() != host {
		t.Error("Shadow root host not set correctly")
	}

	if shadowRoot.Mode() != ShadowRootModeOpen {
		t.Error("Shadow root mode not set correctly")
	}

	if shadowRoot.SlotAssignment() != SlotAssignmentNamed {
		t.Error("Shadow root slot assignment not set correctly")
	}

	// Verify shadow root is a DocumentFragment
	if shadowRoot.NodeType() != DocumentFragmentNode {
		t.Error("Shadow root should be a DocumentFragment")
	}

	// Try to attach another shadow root (should fail)
	_, err = host.AttachShadow(init)
	if err == nil {
		t.Error("Should not be able to attach second shadow root")
	}
}

func TestShadowRootCannotAttachToInvalidElements(t *testing.T) {
	doc := NewDocument()

	// Create elements that cannot have shadow roots
	invalidElements := []string{"br", "img", "input", "meta", "link"}

	for _, tagName := range invalidElements {
		elem := doc.CreateElement(tagName)
		init := ShadowRootInit{
			Mode:           ShadowRootModeOpen,
			SlotAssignment: SlotAssignmentNamed,
		}

		_, err := elem.AttachShadow(init)
		if err == nil {
			t.Errorf("Should not be able to attach shadow root to %s element", tagName)
		}
	}
}

func TestSlotBasics(t *testing.T) {
	// Create a document
	doc := NewDocument()

	// Create a slot element
	slot := NewSlot(doc)

	// Verify slot properties
	if slot.NodeType() != ElementNode {
		t.Error("Slot should be an Element")
	}

	if slot.TagName() != "SLOT" {
		t.Error("Slot should have tag name SLOT")
	}

	// Initially no slot name
	if slot.GetSlotName() != "" {
		t.Error("Slot should have empty name initially")
	}

	// Set slot name
	slot.SetSlotName("content")
	if slot.GetSlotName() != "content" {
		t.Error("Slot name not set correctly")
	}

	// Verify name attribute is set
	if slot.GetAttribute("name") != "content" {
		t.Error("Slot name attribute not set correctly")
	}

	// Initially no assigned nodes
	if len(slot.GetAssignedNodes()) != 0 {
		t.Error("Slot should have no assigned nodes initially")
	}
}

func TestSlottableInterface(t *testing.T) {
	doc := NewDocument()

	// Create an element (implements Slottable)
	elem := doc.CreateElement("div")

	// Initially no slot name
	if elem.GetSlotName() != "" {
		t.Error("Element should have empty slot name initially")
	}

	// Set slot name
	elem.SetSlotName("header")
	if elem.GetSlotName() != "header" {
		t.Error("Element slot name not set correctly")
	}

	// Verify slot attribute is set
	if elem.GetAttribute("slot") != "header" {
		t.Error("Element slot attribute not set correctly")
	}

	// Initially no assigned slot
	if elem.GetAssignedSlot() != nil {
		t.Error("Element should have no assigned slot initially")
	}

	// Create a slot and assign it
	slot := NewSlot(doc)
	elem.SetAssignedSlot(slot)
	if elem.GetAssignedSlot() != slot {
		t.Error("Element assigned slot not set correctly")
	}
}

func TestShadowRootSlotAssignment(t *testing.T) {
	// Create a document
	doc := NewDocument()

	// Create host element
	host := doc.CreateElement("div")

	// Create shadow root
	init := ShadowRootInit{
		Mode:           ShadowRootModeOpen,
		SlotAssignment: SlotAssignmentNamed,
	}
	shadowRoot, err := host.AttachShadow(init)
	if err != nil {
		t.Fatalf("Failed to attach shadow root: %v", err)
	}

	// Create a slot in the shadow root
	slot := NewSlot(doc)
	slot.SetSlotName("content")
	shadowRoot.AppendChild(slot)

	// Create a slottable element and add it to the host
	content := doc.CreateElement("p")
	content.SetSlotName("content")
	host.AppendChild(content)

	// Test slot finding
	foundSlot := shadowRoot.FindSlot(content, false)
	if foundSlot != slot {
		t.Error("FindSlot should find the correct slot")
	}

	// Test slottables finding
	slottables := shadowRoot.FindSlottables(slot)
	if len(slottables) != 1 || slottables[0] != content {
		t.Error("FindSlottables should find the correct slottables")
	}
}

func TestShadowRootManualSlotAssignment(t *testing.T) {
	// Create a document
	doc := NewDocument()

	// Create host element
	host := doc.CreateElement("div")

	// Create shadow root with manual slot assignment
	init := ShadowRootInit{
		Mode:           ShadowRootModeOpen,
		SlotAssignment: SlotAssignmentManual,
	}
	shadowRoot, err := host.AttachShadow(init)
	if err != nil {
		t.Fatalf("Failed to attach shadow root: %v", err)
	}

	// Create a slot in the shadow root
	slot := NewSlot(doc)
	shadowRoot.AppendChild(slot)

	// Create slottable elements
	elem1 := doc.CreateElement("p")
	elem2 := doc.CreateElement("span")
	host.AppendChild(elem1)
	host.AppendChild(elem2)

	// Manually assign elements to the slot
	shadowRoot.SetManualSlotAssignment(slot, []*Element{elem1, elem2})

	// Test manual assignment retrieval
	assignments := shadowRoot.GetManualSlotAssignments()
	if len(assignments) != 1 {
		t.Error("Should have one manual slot assignment")
	}

	if assigned, exists := assignments[slot]; !exists || len(assigned) != 2 {
		t.Error("Manual slot assignment not set correctly")
	}

	// Test slottables finding with manual assignment
	slottables := shadowRoot.FindSlottables(slot)
	if len(slottables) != 2 {
		t.Error("FindSlottables should find manually assigned elements")
	}
}

func TestShadowRootClosedMode(t *testing.T) {
	// Create a document
	doc := NewDocument()

	// Create host element
	host := doc.CreateElement("div")

	// Create closed shadow root
	init := ShadowRootInit{
		Mode:           ShadowRootModeClosed,
		SlotAssignment: SlotAssignmentNamed,
	}
	shadowRoot, err := host.AttachShadow(init)
	if err != nil {
		t.Fatalf("Failed to attach shadow root: %v", err)
	}

	// Verify closed mode
	if shadowRoot.Mode() != ShadowRootModeClosed {
		t.Error("Shadow root should be in closed mode")
	}

	// Create a slot and slottable
	slot := NewSlot(doc)
	slot.SetSlotName("content")
	shadowRoot.AppendChild(slot)

	content := doc.CreateElement("p")
	content.SetSlotName("content")
	host.AppendChild(content)

	// Test slot finding with open=true (should return null for closed shadow root)
	foundSlot := shadowRoot.FindSlot(content, true)
	if foundSlot != nil {
		t.Error("FindSlot with open=true should return null for closed shadow root")
	}

	// Test slot finding with open=false (should work for closed shadow root)
	foundSlot = shadowRoot.FindSlot(content, false)
	if foundSlot != slot {
		t.Error("FindSlot with open=false should work for closed shadow root")
	}
}

func TestShadowIncludingRoot(t *testing.T) {
	// Create a document
	doc := NewDocument()

	// Create host element and add to document
	host := doc.CreateElement("div")
	doc.AppendChild(host)

	// Create shadow root
	init := ShadowRootInit{
		Mode:           ShadowRootModeOpen,
		SlotAssignment: SlotAssignmentNamed,
	}
	shadowRoot, err := host.AttachShadow(init)
	if err != nil {
		t.Fatalf("Failed to attach shadow root: %v", err)
	}

	// Create element in shadow root
	shadowElement := doc.CreateElement("span")
	shadowRoot.AppendChild(shadowElement)

	// Test normal root (should not cross shadow boundary)
	normalRoot := shadowElement.GetRootNode(nil)
	if normalRoot != shadowRoot {
		t.Error("Normal root should be the shadow root")
	}

	// Test shadow-including root (should cross shadow boundary)
	options := &GetRootNodeOptions{Composed: true}
	shadowIncludingRoot := shadowElement.GetRootNode(options)
	if shadowIncludingRoot != doc {
		t.Error("Shadow-including root should be the document")
	}
}
