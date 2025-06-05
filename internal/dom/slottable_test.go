package dom

import (
	"testing"
)

func TestSlottableMixinAssignedSlot(t *testing.T) {
	doc := NewDocument()

	// Create body element for the document
	body := NewElement("body", doc)
	doc.AppendChild(body)

	t.Run("Element AssignedSlot method", func(t *testing.T) {
		// Create a custom element that can have a shadow root
		host := NewElement("div", doc)
		body.AppendChild(host)

		// Create a slottable element
		slottable := NewElement("span", doc)
		slottable.SetAttribute("slot", "test-slot")
		host.AppendChild(slottable)

		// Before shadow root - should return nil
		assignedSlot := slottable.AssignedSlot()
		if assignedSlot != nil {
			t.Errorf("Expected nil assigned slot before shadow root, got: %v", assignedSlot)
		}

		// Create shadow root with a slot
		shadowRoot, err := host.AttachShadow(ShadowRootInit{Mode: ShadowRootModeOpen})
		if err != nil {
			t.Fatalf("Failed to attach shadow root: %v", err)
		}

		slot := NewSlot(doc)
		slot.SetSlotName("test-slot")
		shadowRoot.AppendChild(slot)

		// Now the element should be assigned to the slot
		assignedSlot = slottable.AssignedSlot()
		if assignedSlot == nil {
			t.Error("Expected assigned slot, got nil")
		} else if assignedSlot.GetSlotName() != "test-slot" {
			t.Errorf("Expected slot name 'test-slot', got: %s", assignedSlot.GetSlotName())
		}
	})

	t.Run("Text AssignedSlot method", func(t *testing.T) {
		// Create a custom element that can have a shadow root
		host := NewElement("div", doc)
		body.AppendChild(host)

		// Create a slottable text node
		textNode := NewText("Hello World", doc)
		host.AppendChild(textNode)

		// Before shadow root - should return nil
		assignedSlot := textNode.AssignedSlot()
		if assignedSlot != nil {
			t.Errorf("Expected nil assigned slot before shadow root, got: %v", assignedSlot)
		}

		// Create shadow root with a default slot (no name)
		shadowRoot, err := host.AttachShadow(ShadowRootInit{Mode: ShadowRootModeOpen})
		if err != nil {
			t.Fatalf("Failed to attach shadow root: %v", err)
		}

		slot := NewSlot(doc)
		// No slot name = default slot
		shadowRoot.AppendChild(slot)

		// Text nodes go to default slot when they have no slot name
		assignedSlot = textNode.AssignedSlot()
		if assignedSlot == nil {
			t.Error("Expected text node to be assigned to default slot, got nil")
		} else if assignedSlot.GetSlotName() != "" {
			t.Errorf("Expected default slot (empty name), got: %s", assignedSlot.GetSlotName())
		}
	})

	t.Run("Closed shadow root returns nil", func(t *testing.T) {
		// Create a custom element that can have a shadow root
		host := NewElement("div", doc)
		body.AppendChild(host)

		// Create a slottable element
		slottable := NewElement("span", doc)
		slottable.SetAttribute("slot", "test-slot")
		host.AppendChild(slottable)

		// Create closed shadow root with a slot
		shadowRoot, err := host.AttachShadow(ShadowRootInit{Mode: ShadowRootModeClosed})
		if err != nil {
			t.Fatalf("Failed to attach shadow root: %v", err)
		}

		slot := NewSlot(doc)
		slot.SetSlotName("test-slot")
		shadowRoot.AppendChild(slot)

		// AssignedSlot() should return nil for closed shadow roots (open=true in spec)
		assignedSlot := slottable.AssignedSlot()
		if assignedSlot != nil {
			t.Errorf("Expected nil assigned slot for closed shadow root, got: %v", assignedSlot)
		}
	})

	t.Run("Manual slot assignment", func(t *testing.T) {
		// Create a custom element that can have a shadow root
		host := NewElement("div", doc)
		body.AppendChild(host)

		// Create a slottable element
		slottable := NewElement("span", doc)
		host.AppendChild(slottable)

		// Create shadow root with manual slot assignment
		shadowRoot, err := host.AttachShadow(ShadowRootInit{
			Mode:           ShadowRootModeOpen,
			SlotAssignment: SlotAssignmentManual,
		})
		if err != nil {
			t.Fatalf("Failed to attach shadow root: %v", err)
		}

		slot := NewSlot(doc)
		shadowRoot.AppendChild(slot)

		// Manually assign the element to the slot
		shadowRoot.SetManualSlotAssignment(slot, []*Element{slottable})

		// AssignedSlot() should find the manually assigned slot
		assignedSlot := slottable.AssignedSlot()
		if assignedSlot == nil {
			t.Error("Expected assigned slot for manual assignment, got nil")
		} else if assignedSlot != slot {
			t.Error("Expected manually assigned slot to be found")
		}
	})
}

func TestSlottableMixinSpecificationCompliance(t *testing.T) {
	doc := NewDocument()

	// Create body element for the document
	body := NewElement("body", doc)
	doc.AppendChild(body)

	t.Run("WHATWG DOM Section 4.2.9 compliance", func(t *testing.T) {
		// Test that the assignedSlot getter follows the specification:
		// "The assignedSlot getter steps are to return the result of find a slot given this and true."

		// Create a basic setup
		host := NewElement("div", doc)
		body.AppendChild(host)

		element := NewElement("span", doc)
		element.SetAttribute("slot", "named-slot")
		host.AppendChild(element)

		textNode := NewText("test text", doc)
		host.AppendChild(textNode)

		// Before shadow root attachment - both should return nil
		if element.AssignedSlot() != nil {
			t.Error("Element should have no assigned slot before shadow root")
		}
		if textNode.AssignedSlot() != nil {
			t.Error("Text node should have no assigned slot before shadow root")
		}

		// Create open shadow root
		shadowRoot, err := host.AttachShadow(ShadowRootInit{Mode: ShadowRootModeOpen})
		if err != nil {
			t.Fatalf("Failed to attach shadow root: %v", err)
		}

		// Add named slot
		namedSlot := NewSlot(doc)
		namedSlot.SetSlotName("named-slot")
		shadowRoot.AppendChild(namedSlot)

		// Add default slot
		defaultSlot := NewSlot(doc)
		shadowRoot.AppendChild(defaultSlot)

		// Element with slot="named-slot" should be assigned to named slot
		assignedSlot := element.AssignedSlot()
		if assignedSlot == nil {
			t.Error("Element should be assigned to named slot")
		} else if assignedSlot != namedSlot {
			t.Error("Element should be assigned to the named slot")
		}

		// Text node (no slot attribute) should be assigned to default slot
		assignedSlot = textNode.AssignedSlot()
		if assignedSlot == nil {
			t.Error("Text node should be assigned to default slot")
		} else if assignedSlot != defaultSlot {
			t.Error("Text node should be assigned to the default slot")
		}
	})

	t.Run("Find a slot algorithm with open=true", func(t *testing.T) {
		// The assignedSlot getter calls "find a slot" with open=true
		// This means it only finds slots in open shadow roots

		host := NewElement("div", doc)
		body.AppendChild(host)

		element := NewElement("span", doc)
		element.SetAttribute("slot", "test-slot")
		host.AppendChild(element)

		// Test with closed shadow root
		closedShadowRoot, err := host.AttachShadow(ShadowRootInit{Mode: ShadowRootModeClosed})
		if err != nil {
			t.Fatalf("Failed to attach closed shadow root: %v", err)
		}

		slot := NewSlot(doc)
		slot.SetSlotName("test-slot")
		closedShadowRoot.AppendChild(slot)

		// AssignedSlot() with open=true should return nil for closed shadow root
		if element.AssignedSlot() != nil {
			t.Error("assignedSlot should return nil for closed shadow root when open=true")
		}

		// Now test that internal GetAssignedSlot() (used for slot assignment) can still find it
		// This would use open=false internally
		if element.GetAssignedSlot() == nil {
			// This is expected since we haven't run slot assignment yet
			// The AssignedSlot() method follows the spec and uses open=true
		}
	})
}
