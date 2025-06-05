package dom

import (
	"testing"
)

// TestShadowRootSection4_8_Interface tests WHATWG DOM Section 4.8 ShadowRoot interface compliance
func TestShadowRootSection4_8_Interface(t *testing.T) {
	t.Run("ShadowRoot Properties", func(t *testing.T) {
		doc := NewDocument()
		host := doc.CreateElement("div")

		// Test basic creation
		sr := NewShadowRoot(host, ShadowRootModeOpen)

		// Test mode getter
		if sr.Mode() != ShadowRootModeOpen {
			t.Errorf("Expected mode to be open, got %v", sr.Mode())
		}

		// Test host getter
		if sr.Host() != host {
			t.Errorf("Expected host to be %v, got %v", host, sr.Host())
		}

		// Test slotAssignment getter (default should be "named")
		if sr.SlotAssignment() != SlotAssignmentNamed {
			t.Errorf("Expected slotAssignment to be named, got %v", sr.SlotAssignment())
		}

		// Test delegates focus getter (default should be false)
		if sr.DelegatesFocus() != false {
			t.Errorf("Expected delegatesFocus to be false, got %v", sr.DelegatesFocus())
		}

		// Test clonable getter (default should be false)
		if sr.Clonable() != false {
			t.Errorf("Expected clonable to be false, got %v", sr.Clonable())
		}

		// Test serializable getter (default should be false)
		if sr.Serializable() != false {
			t.Errorf("Expected serializable to be false, got %v", sr.Serializable())
		}

		// Test onslotchange event handler
		if sr.OnSlotChange() != nil {
			t.Error("Expected initial onslotchange to be nil")
		}

		handler := func(event *Event) {}
		sr.SetOnSlotChange(handler)
		if sr.OnSlotChange() == nil {
			t.Error("Expected onslotchange handler to be set")
		}
	})

	t.Run("ShadowRootInit Properties", func(t *testing.T) {
		doc := NewDocument()
		host := doc.CreateElement("div")

		init := &ShadowRootInit{
			Mode:           ShadowRootModeClosed,
			SlotAssignment: SlotAssignmentManual,
			DelegatesFocus: true,
			Clonable:       true,
			Serializable:   true,
		}

		sr := NewShadowRootWithInit(host, init)

		// Test that all properties are set correctly
		if sr.Mode() != ShadowRootModeClosed {
			t.Errorf("Expected mode to be closed, got %v", sr.Mode())
		}

		if sr.SlotAssignment() != SlotAssignmentManual {
			t.Errorf("Expected slotAssignment to be manual, got %v", sr.SlotAssignment())
		}

		if sr.DelegatesFocus() != true {
			t.Errorf("Expected delegatesFocus to be true, got %v", sr.DelegatesFocus())
		}

		if sr.Clonable() != true {
			t.Errorf("Expected clonable to be true, got %v", sr.Clonable())
		}

		if sr.Serializable() != true {
			t.Errorf("Expected serializable to be true, got %v", sr.Serializable())
		}
	})
}

// TestShadowIncludingTreeAlgorithms tests shadow-including tree algorithms from WHATWG DOM Section 4.8
func TestShadowIncludingTreeAlgorithms(t *testing.T) {
	t.Run("ShadowIncludingRoot", func(t *testing.T) {
		doc := NewDocument()
		body := doc.CreateElement("body")
		doc.AppendChild(body)

		host := doc.CreateElement("div")
		body.AppendChild(host)

		init := ShadowRootInit{Mode: ShadowRootModeOpen}
		sr, _ := host.AttachShadow(init)

		child := doc.CreateElement("span")
		sr.AppendChild(child)

		// Test shadow-including root of child element in shadow DOM
		shadowIncludingRoot := ShadowIncludingRoot(child)
		if shadowIncludingRoot != doc {
			t.Errorf("Expected shadow-including root to be document, got %v", shadowIncludingRoot)
		}

		// Test shadow-including root of shadow root itself
		shadowIncludingRoot = ShadowIncludingRoot(sr)
		if shadowIncludingRoot != doc {
			t.Errorf("Expected shadow-including root of shadow root to be document, got %v", shadowIncludingRoot)
		}

		// Test shadow-including root of host
		shadowIncludingRoot = ShadowIncludingRoot(host)
		if shadowIncludingRoot != doc {
			t.Errorf("Expected shadow-including root of host to be document, got %v", shadowIncludingRoot)
		}
	})

	t.Run("IsShadowIncludingDescendant", func(t *testing.T) {
		doc := NewDocument()
		body := doc.CreateElement("body")
		doc.AppendChild(body)

		host := doc.CreateElement("div")
		body.AppendChild(host)

		init := ShadowRootInit{Mode: ShadowRootModeOpen}
		sr, _ := host.AttachShadow(init)

		child := doc.CreateElement("span")
		sr.AppendChild(child)

		// Test normal descendant relationship
		if !IsShadowIncludingDescendant(child, sr) {
			t.Error("Child should be shadow-including descendant of shadow root")
		}

		// Test shadow-including descendant relationship
		if !IsShadowIncludingDescendant(child, body) {
			t.Error("Child in shadow should be shadow-including descendant of host's parent")
		}

		// Test that unrelated nodes are not descendants
		other := doc.CreateElement("p")
		if IsShadowIncludingDescendant(other, sr) {
			t.Error("Unrelated element should not be shadow-including descendant")
		}
	})

	t.Run("IsShadowIncludingInclusiveDescendant", func(t *testing.T) {
		doc := NewDocument()
		host := doc.CreateElement("div")

		init := ShadowRootInit{Mode: ShadowRootModeOpen}
		sr, _ := host.AttachShadow(init)

		child := doc.CreateElement("span")
		sr.AppendChild(child)

		// Test inclusive relationship (node is descendant of itself)
		if !IsShadowIncludingInclusiveDescendant(sr, sr) {
			t.Error("Shadow root should be shadow-including inclusive descendant of itself")
		}

		if !IsShadowIncludingInclusiveDescendant(child, child) {
			t.Error("Child should be shadow-including inclusive descendant of itself")
		}

		// Test normal descendant relationship
		if !IsShadowIncludingInclusiveDescendant(child, sr) {
			t.Error("Child should be shadow-including inclusive descendant of shadow root")
		}
	})

	t.Run("IsShadowIncludingAncestor", func(t *testing.T) {
		doc := NewDocument()
		host := doc.CreateElement("div")

		init := ShadowRootInit{Mode: ShadowRootModeOpen}
		sr, _ := host.AttachShadow(init)

		child := doc.CreateElement("span")
		sr.AppendChild(child)

		// Test ancestor relationship (inverse of descendant)
		if !IsShadowIncludingAncestor(sr, child) {
			t.Error("Shadow root should be shadow-including ancestor of child")
		}

		if IsShadowIncludingAncestor(child, sr) {
			t.Error("Child should not be shadow-including ancestor of shadow root")
		}
	})
}

// TestClosedShadowHidden tests the closed-shadow-hidden algorithm from WHATWG DOM Section 4.8
func TestClosedShadowHidden(t *testing.T) {
	doc := NewDocument()
	body := doc.CreateElement("body")
	doc.AppendChild(body)

	// Create closed shadow DOM
	hostClosed := doc.CreateElement("div")
	body.AppendChild(hostClosed)

	initClosed := ShadowRootInit{Mode: ShadowRootModeClosed}
	srClosed, _ := hostClosed.AttachShadow(initClosed)

	childClosed := doc.CreateElement("span")
	srClosed.AppendChild(childClosed)

	// Create open shadow DOM
	hostOpen := doc.CreateElement("div")
	body.AppendChild(hostOpen)

	initOpen := ShadowRootInit{Mode: ShadowRootModeOpen}
	srOpen, _ := hostOpen.AttachShadow(initOpen)

	childOpen := doc.CreateElement("span")
	srOpen.AppendChild(childOpen)

	// Test closed shadow hidden
	t.Run("Closed Shadow Hidden", func(t *testing.T) {
		// Child in closed shadow should be hidden from elements outside
		if !IsClosedShadowHidden(childClosed, body) {
			t.Error("Child in closed shadow should be hidden from body")
		}

		// Child in open shadow should NOT be hidden
		if IsClosedShadowHidden(childOpen, body) {
			t.Error("Child in open shadow should not be hidden from body")
		}

		// Elements in same shadow tree should not be hidden from each other
		siblingClosed := doc.CreateElement("em")
		srClosed.AppendChild(siblingClosed)

		if IsClosedShadowHidden(childClosed, siblingClosed) {
			t.Error("Children in same shadow tree should not be hidden from each other")
		}
	})
}

// TestRetarget tests the retargeting algorithm from WHATWG DOM Section 4.8
func TestRetarget(t *testing.T) {
	doc := NewDocument()
	body := doc.CreateElement("body")
	doc.AppendChild(body)

	host := doc.CreateElement("div")
	body.AppendChild(host)

	init := ShadowRootInit{Mode: ShadowRootModeOpen}
	sr, _ := host.AttachShadow(init)

	child := doc.CreateElement("span")
	sr.AppendChild(child)

	t.Run("Retargeting Algorithm", func(t *testing.T) {
		// Retargeting from child in shadow to element outside shadow
		retargeted := Retarget(child, body)
		if retargeted != host {
			t.Errorf("Expected retargeted node to be host %v, got %v", host, retargeted)
		}

		// Retargeting within same shadow tree should return original
		sibling := doc.CreateElement("em")
		sr.AppendChild(sibling)

		retargeted = Retarget(child, sibling)
		if retargeted != child {
			t.Errorf("Expected retargeted node to be original child %v, got %v", child, retargeted)
		}

		// Retargeting to unrelated node
		unrelated := doc.CreateElement("p")
		retargeted = Retarget(child, unrelated)
		if retargeted != host {
			t.Errorf("Expected retargeted node to be host %v, got %v", host, retargeted)
		}
	})
}

// TestGetEventParent tests the shadow root's get the parent algorithm
func TestGetEventParent(t *testing.T) {
	doc := NewDocument()
	host := doc.CreateElement("div")

	init := ShadowRootInit{Mode: ShadowRootModeOpen}
	sr, _ := host.AttachShadow(init)

	event := NewEvent("test", true, true)

	t.Run("GetEventParent", func(t *testing.T) {
		// For now, our simplified implementation always returns the host
		// Use the event directly without casting
		parent := sr.GetEventParent(event)
		if parent != host {
			t.Errorf("Expected event parent to be host %v, got %v", host, parent)
		}

		// Test with nil event
		parent = sr.GetEventParent(nil)
		if parent != host {
			t.Errorf("Expected event parent to be host %v with nil event, got %v", host, parent)
		}
	})
}

// TestTraverseShadowIncludingTreeOrder tests shadow-including tree order traversal
func TestTraverseShadowIncludingTreeOrder(t *testing.T) {
	doc := NewDocument()
	body := doc.CreateElement("body")
	doc.AppendChild(body)

	// Create a nested structure with shadow DOM
	host1 := doc.CreateElement("div")
	body.AppendChild(host1)

	init1 := ShadowRootInit{Mode: ShadowRootModeOpen}
	sr1, _ := host1.AttachShadow(init1)

	child1 := doc.CreateElement("span")
	sr1.AppendChild(child1)

	// Create nested shadow DOM
	host2 := doc.CreateElement("div")
	sr1.AppendChild(host2)

	init2 := ShadowRootInit{Mode: ShadowRootModeOpen}
	sr2, _ := host2.AttachShadow(init2)

	child2 := doc.CreateElement("em")
	sr2.AppendChild(child2)

	// Add a sibling to host1
	sibling := doc.CreateElement("p")
	body.AppendChild(sibling)

	t.Run("Shadow-Including Tree Order", func(t *testing.T) {
		var visited []Node
		TraverseShadowIncludingTreeOrder(body, func(node Node) bool {
			visited = append(visited, node)
			return true
		})

		// Check that we visited nodes in correct order
		// Should be: body, host1, sr1, child1, host2, sr2, child2, sibling
		expectedOrder := []Node{body, host1, sr1, child1, host2, sr2, child2, sibling}

		if len(visited) != len(expectedOrder) {
			t.Errorf("Expected %d visited nodes, got %d", len(expectedOrder), len(visited))
		}

		for i, expected := range expectedOrder {
			if i >= len(visited) || visited[i] != expected {
				t.Errorf("At position %d, expected %v, got %v", i, expected, visited[i])
			}
		}
	})
}

// TestShadowRootGetRootNode tests the overridden GetRootNode method
func TestShadowRootGetRootNode(t *testing.T) {
	doc := NewDocument()
	body := doc.CreateElement("body")
	doc.AppendChild(body)

	host := doc.CreateElement("div")
	body.AppendChild(host)

	init := ShadowRootInit{Mode: ShadowRootModeOpen}
	sr, _ := host.AttachShadow(init)

	child := doc.CreateElement("span")
	sr.AppendChild(child)

	t.Run("GetRootNode with Composed", func(t *testing.T) {
		// Test composed: false (normal root)
		root := sr.GetRootNode(nil)
		if root != sr {
			t.Errorf("Expected normal root to be shadow root %v, got %v", sr, root)
		}

		options := &GetRootNodeOptions{Composed: false}
		root = sr.GetRootNode(options)
		if root != sr {
			t.Errorf("Expected normal root to be shadow root %v, got %v", sr, root)
		}

		// Test composed: true (shadow-including root)
		options.Composed = true
		root = sr.GetRootNode(options)
		if root != doc {
			t.Errorf("Expected shadow-including root to be document %v, got %v", doc, root)
		}
	})
}

// TestShadowRootNodeName tests the NodeName override
func TestShadowRootNodeName(t *testing.T) {
	doc := NewDocument()
	host := doc.CreateElement("div")

	sr := NewShadowRoot(host, ShadowRootModeOpen)

	if sr.NodeName() != "#shadow-root" {
		t.Errorf("Expected NodeName to be '#shadow-root', got '%s'", sr.NodeName())
	}
}

// TestShadowRootString tests the String method for debugging
func TestShadowRootString(t *testing.T) {
	doc := NewDocument()
	host := doc.CreateElement("div")

	t.Run("Open Shadow Root String", func(t *testing.T) {
		sr := NewShadowRoot(host, ShadowRootModeOpen)
		str := sr.String()
		expected := "#shadow-root (open) host: DIV"
		if str != expected {
			t.Errorf("Expected string '%s', got '%s'", expected, str)
		}
	})

	t.Run("Closed Shadow Root String", func(t *testing.T) {
		sr := NewShadowRoot(host, ShadowRootModeClosed)
		str := sr.String()
		expected := "#shadow-root (closed) host: DIV"
		if str != expected {
			t.Errorf("Expected string '%s', got '%s'", expected, str)
		}
	})
}

// Benchmark tests for performance
func BenchmarkShadowIncludingRoot(b *testing.B) {
	doc := NewDocument()
	host := doc.CreateElement("div")
	init := ShadowRootInit{Mode: ShadowRootModeOpen}
	sr, _ := host.AttachShadow(init)
	child := doc.CreateElement("span")
	sr.AppendChild(child)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ShadowIncludingRoot(child)
	}
}

func BenchmarkIsShadowIncludingDescendant(b *testing.B) {
	doc := NewDocument()
	host := doc.CreateElement("div")
	init := ShadowRootInit{Mode: ShadowRootModeOpen}
	sr, _ := host.AttachShadow(init)
	child := doc.CreateElement("span")
	sr.AppendChild(child)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = IsShadowIncludingDescendant(child, sr)
	}
}

func BenchmarkRetarget(b *testing.B) {
	doc := NewDocument()
	host := doc.CreateElement("div")
	init := ShadowRootInit{Mode: ShadowRootModeOpen}
	sr, _ := host.AttachShadow(init)
	child := doc.CreateElement("span")
	sr.AppendChild(child)
	target := doc.CreateElement("div")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Retarget(child, target)
	}
}
