package dom

import (
	"testing"
)

// TestNodeIteratorSpecCompliance tests exact compliance with WHATWG DOM spec section 6.1
func TestNodeIteratorSpecCompliance(t *testing.T) {
	t.Run("Initial State", func(t *testing.T) {
		// Per spec: Each NodeIterator has reference (a node) and pointer before reference (a boolean)
		// Initial state should have reference = root and pointer before reference = true
		doc := NewDocument()
		root := doc.CreateElement("div")
		iterator := NewNodeIterator(root, NodeFilterShowAll, nil)

		if iterator.ReferenceNode() != root {
			t.Errorf("Initial reference node should be root, got %v", iterator.ReferenceNode())
		}
		if !iterator.PointerBeforeReferenceNode() {
			t.Error("Initial pointer before reference should be true")
		}
	})

	t.Run("Traverse Algorithm - Next Direction", func(t *testing.T) {
		doc := NewDocument()
		root := doc.CreateElement("div")
		child1 := doc.CreateElement("span")
		child2 := doc.CreateElement("p")
		root.AppendChild(child1)
		root.AppendChild(child2)

		iterator := NewNodeIterator(root, NodeFilterShowAll, nil)

		// First call: pointer is before reference (true), reference is root
		// Step: beforeNode is true, so set it to false
		// Filter root -> ACCEPT
		// Update: reference = root, pointer before = false
		node := iterator.NextNode()
		if node != root {
			t.Errorf("First NextNode should return root, got %v", node)
		}
		if iterator.ReferenceNode() != root {
			t.Errorf("After first NextNode, reference should be root, got %v", iterator.ReferenceNode())
		}
		if iterator.PointerBeforeReferenceNode() {
			t.Error("After first NextNode, pointer should be after reference (false)")
		}

		// Second call: pointer is after reference (false), reference is root
		// Step: beforeNode is false, so set node to first following node (child1)
		// Filter child1 -> ACCEPT
		// Update: reference = child1, pointer before = false
		node = iterator.NextNode()
		if node != child1 {
			t.Errorf("Second NextNode should return child1, got %v", node)
		}
		if iterator.ReferenceNode() != child1 {
			t.Errorf("After second NextNode, reference should be child1, got %v", iterator.ReferenceNode())
		}
		if iterator.PointerBeforeReferenceNode() {
			t.Error("After second NextNode, pointer should still be after reference (false)")
		}
	})

	t.Run("Traverse Algorithm - Previous Direction", func(t *testing.T) {
		doc := NewDocument()
		root := doc.CreateElement("div")
		child := doc.CreateElement("span")
		root.AppendChild(child)

		iterator := NewNodeIterator(root, NodeFilterShowAll, nil)

		// Move to child first
		iterator.NextNode() // root
		iterator.NextNode() // child

		// Now: reference = child, pointer after reference (false)
		// Call PreviousNode:
		// Step: beforeNode is false, so set it to true
		// Filter child -> ACCEPT
		// Update: reference = child, pointer before = true
		// Return child
		node := iterator.PreviousNode()
		if node != child {
			t.Errorf("First PreviousNode should return current reference (child), got %v", node)
		}
		if iterator.ReferenceNode() != child {
			t.Errorf("After first PreviousNode, reference should still be child, got %v", iterator.ReferenceNode())
		}
		if !iterator.PointerBeforeReferenceNode() {
			t.Error("After first PreviousNode, pointer should be before reference (true)")
		}

		// Second call: pointer is before reference (true), reference is child
		// Step: beforeNode is true, so set node to first preceding node (root)
		// Filter root -> ACCEPT
		// Update: reference = root, pointer before = true
		node = iterator.PreviousNode()
		if node != root {
			t.Errorf("Second PreviousNode should return root, got %v", node)
		}
		if iterator.ReferenceNode() != root {
			t.Errorf("After second PreviousNode, reference should be root, got %v", iterator.ReferenceNode())
		}
		if !iterator.PointerBeforeReferenceNode() {
			t.Error("After second PreviousNode, pointer should still be before reference (true)")
		}
	})

	t.Run("Boundary Behavior - End of Iteration", func(t *testing.T) {
		doc := NewDocument()
		root := doc.CreateElement("div")
		child := doc.CreateElement("span")
		root.AppendChild(child)

		iterator := NewNodeIterator(root, NodeFilterShowAll, nil)

		// Traverse to end
		iterator.NextNode() // root
		iterator.NextNode() // child

		// Try to go past end
		node := iterator.NextNode()
		if node != nil {
			t.Errorf("NextNode at end should return nil, got %v", node)
		}
		// State should be unchanged when returning nil
		if iterator.ReferenceNode() != child {
			t.Errorf("Reference should remain at child after hitting end, got %v", iterator.ReferenceNode())
		}
		if iterator.PointerBeforeReferenceNode() {
			t.Error("Pointer should remain after reference when NextNode returns nil")
		}

		// Now call PreviousNode - should return current reference
		node = iterator.PreviousNode()
		if node != child {
			t.Errorf("PreviousNode after end should return last node (child), got %v", node)
		}
	})

	t.Run("Boundary Behavior - Beginning of Iteration", func(t *testing.T) {
		doc := NewDocument()
		root := doc.CreateElement("div")
		child := doc.CreateElement("span")
		root.AppendChild(child)

		iterator := NewNodeIterator(root, NodeFilterShowAll, nil)

		// Try to go before beginning (initial state has pointer before root)
		node := iterator.PreviousNode()
		if node != nil {
			t.Errorf("PreviousNode at beginning should return nil, got %v", node)
		}
		// State should be unchanged when returning nil
		if iterator.ReferenceNode() != root {
			t.Errorf("Reference should remain at root after hitting beginning, got %v", iterator.ReferenceNode())
		}
		if !iterator.PointerBeforeReferenceNode() {
			t.Error("Pointer should remain before reference when PreviousNode returns nil")
		}
	})

	t.Run("Filter Algorithm - whatToShow Bit Checking", func(t *testing.T) {
		doc := NewDocument()
		root := doc.CreateElement("div")
		elem := doc.CreateElement("span")
		text := doc.CreateTextNode("text")
		comment := doc.CreateComment("comment")

		root.AppendChild(elem)
		root.AppendChild(text)
		root.AppendChild(comment)

		// Test with NodeFilterShowElement (0x00000001)
		// Per spec: Let n be node's nodeType attribute value − 1
		// Element has nodeType = 1, so n = 0, check bit 0
		iterator := NewNodeIterator(root, NodeFilterShowElement, nil)

		var nodes []Node
		for {
			node := iterator.NextNode()
			if node == nil {
				break
			}
			nodes = append(nodes, node)
		}

		if len(nodes) != 2 {
			t.Errorf("Expected 2 elements, got %d", len(nodes))
		}
		if len(nodes) >= 1 && nodes[0] != root {
			t.Errorf("First node should be root element, got %v", nodes[0])
		}
		if len(nodes) >= 2 && nodes[1] != elem {
			t.Errorf("Second node should be span element, got %v", nodes[1])
		}

		// Test with NodeFilterShowText (0x00000004)
		// Text has nodeType = 3, so n = 2, check bit 2
		iterator2 := NewNodeIterator(root, NodeFilterShowText, nil)

		nodes = nil
		for {
			node := iterator2.NextNode()
			if node == nil {
				break
			}
			nodes = append(nodes, node)
		}

		if len(nodes) != 1 {
			t.Errorf("Expected 1 text node, got %d", len(nodes))
		}
		if len(nodes) >= 1 && nodes[0] != text {
			t.Errorf("Should return text node, got %v", nodes[0])
		}

		// Test with NodeFilterShowComment (0x00000080)
		// Comment has nodeType = 8, so n = 7, check bit 7
		iterator3 := NewNodeIterator(root, NodeFilterShowComment, nil)

		nodes = nil
		for {
			node := iterator3.NextNode()
			if node == nil {
				break
			}
			nodes = append(nodes, node)
		}

		if len(nodes) != 1 {
			t.Errorf("Expected 1 comment node, got %d", len(nodes))
		}
		if len(nodes) >= 1 && nodes[0] != comment {
			t.Errorf("Should return comment node, got %v", nodes[0])
		}
	})

	t.Run("Filter Algorithm - Active Flag", func(t *testing.T) {
		doc := NewDocument()
		root := doc.CreateElement("div")
		child := doc.CreateElement("span")
		root.AppendChild(child)

		var iterator *NodeIterator
		filterCalls := 0

		// Create filter that attempts recursive traversal
		filter := NodeFilterFunc(func(node Node) int {
			filterCalls++
			t.Logf("Filter called for node %v, call #%d", node, filterCalls)

			// On the second call (which should be for child), try recursive call
			if filterCalls == 2 {
				t.Log("Attempting recursive NextNode call")
				// This should panic with InvalidStateError
				iterator.NextNode()
			}
			return NodeFilterAccept
		})

		iterator = NewNodeIterator(root, NodeFilterShowAll, filter)

		// First call should succeed (will filter root)
		node := iterator.NextNode()
		if node != root {
			t.Errorf("Expected root, got %v", node)
		}

		// Second call should trigger recursive attempt and panic
		defer func() {
			if r := recover(); r != nil {
				t.Logf("Recovered panic: %v (type: %T)", r, r)
				if err, ok := r.(*DOMException); ok {
					if err.Name() != "InvalidStateError" {
						t.Errorf("Expected InvalidStateError, got %s", err.Name())
					}
				} else {
					t.Errorf("Expected DOMException, got %v", r)
				}
			} else {
				t.Errorf("Expected panic with InvalidStateError for recursive filter call, filterCalls=%d", filterCalls)
			}
		}()

		// This should cause filter to be called for child, triggering recursive call
		iterator.NextNode()
	})

	t.Run("Filter Algorithm - Result Values", func(t *testing.T) {
		doc := NewDocument()
		root := doc.CreateElement("div")
		child1 := doc.CreateElement("span")
		child2 := doc.CreateElement("p")
		child3 := doc.CreateElement("a")

		root.AppendChild(child1)
		root.AppendChild(child2)
		root.AppendChild(child3)

		// Test FILTER_SKIP behavior
		skipFilter := NodeFilterFunc(func(node Node) int {
			if node == child1 {
				return NodeFilterSkip
			}
			return NodeFilterAccept
		})

		iterator := NewNodeIterator(root, NodeFilterShowAll, skipFilter)

		nodes := []Node{
			iterator.NextNode(), // root (accept)
			iterator.NextNode(), // child1 (skip) -> child2 (accept)
			iterator.NextNode(), // child3 (accept)
		}

		if nodes[0] != root {
			t.Errorf("First node should be root, got %v", nodes[0])
		}
		if nodes[1] != child2 {
			t.Errorf("Second node should be child2 (child1 skipped), got %v", nodes[1])
		}
		if nodes[2] != child3 {
			t.Errorf("Third node should be child3, got %v", nodes[2])
		}

		// Test FILTER_REJECT behavior (should skip node and its descendants)
		rejectFilter := NodeFilterFunc(func(node Node) int {
			if node == child1 {
				return NodeFilterReject
			}
			return NodeFilterAccept
		})

		// Add a descendant to child1 to verify it's also skipped
		grandchild := doc.CreateElement("em")
		child1.AppendChild(grandchild)

		iterator2 := NewNodeIterator(root, NodeFilterShowAll, rejectFilter)

		nodes = []Node{
			iterator2.NextNode(), // root (accept)
			iterator2.NextNode(), // child1 (reject, skip it and descendants) -> child2 (accept)
			iterator2.NextNode(), // child3 (accept)
		}

		if nodes[0] != root {
			t.Errorf("First node should be root, got %v", nodes[0])
		}
		if nodes[1] != child2 {
			t.Errorf("Second node should be child2 (child1 and descendants rejected), got %v", nodes[1])
		}
		if nodes[2] != child3 {
			t.Errorf("Third node should be child3, got %v", nodes[2])
		}
	})

	t.Run("PreRemove Algorithm - Reference is Descendant", func(t *testing.T) {
		doc := NewDocument()
		root := doc.CreateElement("div")
		child1 := doc.CreateElement("span")
		child2 := doc.CreateElement("p")
		child3 := doc.CreateElement("a")

		root.AppendChild(child1)
		root.AppendChild(child2)
		root.AppendChild(child3)

		iterator := NewNodeIterator(root, NodeFilterShowAll, nil)

		// Position at child2
		iterator.NextNode() // root
		iterator.NextNode() // child1
		iterator.NextNode() // child2

		// Verify state
		if iterator.ReferenceNode() != child2 {
			t.Errorf("Reference should be child2, got %v", iterator.ReferenceNode())
		}
		if iterator.PointerBeforeReferenceNode() {
			t.Error("Pointer should be after reference")
		}

		// Simulate removing child2 (pointer after reference)
		// Per spec step 3: reference moves to previous sibling's last descendant
		iterator.preRemoveNode(child2)

		if iterator.ReferenceNode() != child1 {
			t.Errorf("After removal, reference should be child1, got %v", iterator.ReferenceNode())
		}
		if iterator.PointerBeforeReferenceNode() {
			t.Error("Pointer should remain after reference")
		}
	})

	t.Run("PreRemove Algorithm - Pointer Before Reference", func(t *testing.T) {
		doc := NewDocument()
		root := doc.CreateElement("div")
		child1 := doc.CreateElement("span")
		child2 := doc.CreateElement("p")
		child3 := doc.CreateElement("a")

		root.AppendChild(child1)
		root.AppendChild(child2)
		root.AppendChild(child3)

		iterator := NewNodeIterator(root, NodeFilterShowAll, nil)

		// Position at child2 with pointer before
		iterator.NextNode()     // root
		iterator.NextNode()     // child1
		iterator.NextNode()     // child2
		iterator.PreviousNode() // child2 (pointer now before)

		// Verify state
		if iterator.ReferenceNode() != child2 {
			t.Errorf("Reference should be child2, got %v", iterator.ReferenceNode())
		}
		if !iterator.PointerBeforeReferenceNode() {
			t.Error("Pointer should be before reference")
		}

		// Simulate removing child2 (pointer before reference)
		// Per spec step 2: reference moves to next node (child3)
		iterator.preRemoveNode(child2)

		if iterator.ReferenceNode() != child3 {
			t.Errorf("After removal, reference should be child3, got %v", iterator.ReferenceNode())
		}
		if !iterator.PointerBeforeReferenceNode() {
			t.Error("Pointer should remain before reference")
		}
	})

	t.Run("Mixed Direction Traversal", func(t *testing.T) {
		doc := NewDocument()
		root := doc.CreateElement("div")
		child1 := doc.CreateElement("span")
		child2 := doc.CreateElement("p")

		root.AppendChild(child1)
		root.AppendChild(child2)

		iterator := NewNodeIterator(root, NodeFilterShowAll, nil)

		// Forward: root
		node := iterator.NextNode()
		if node != root {
			t.Errorf("Step 1: Expected root, got %v", node)
		}

		// Forward: child1
		node = iterator.NextNode()
		if node != child1 {
			t.Errorf("Step 2: Expected child1, got %v", node)
		}

		// Backward: child1 (toggle pointer)
		node = iterator.PreviousNode()
		if node != child1 {
			t.Errorf("Step 3: Expected child1 (pointer toggle), got %v", node)
		}

		// Backward: root
		node = iterator.PreviousNode()
		if node != root {
			t.Errorf("Step 4: Expected root, got %v", node)
		}

		// Forward: root (toggle pointer)
		node = iterator.NextNode()
		if node != root {
			t.Errorf("Step 5: Expected root (pointer toggle), got %v", node)
		}

		// Forward: child1
		node = iterator.NextNode()
		if node != child1 {
			t.Errorf("Step 6: Expected child1, got %v", node)
		}

		// Forward: child2
		node = iterator.NextNode()
		if node != child2 {
			t.Errorf("Step 7: Expected child2, got %v", node)
		}
	})
}

// TestNodeIteratorDocumentOrder verifies correct document order traversal
func TestNodeIteratorDocumentOrder(t *testing.T) {
	doc := NewDocument()

	// Create tree structure:
	//   root
	//   ├── A
	//   │   ├── A1
	//   │   └── A2
	//   │       └── A2a
	//   └── B
	//       └── B1
	root := doc.CreateElement("root")
	a := doc.CreateElement("A")
	a1 := doc.CreateElement("A1")
	a2 := doc.CreateElement("A2")
	a2a := doc.CreateElement("A2a")
	b := doc.CreateElement("B")
	b1 := doc.CreateElement("B1")

	root.AppendChild(a)
	root.AppendChild(b)
	a.AppendChild(a1)
	a.AppendChild(a2)
	a2.AppendChild(a2a)
	b.AppendChild(b1)

	iterator := NewNodeIterator(root, NodeFilterShowAll, nil)

	// Expected document order
	expectedOrder := []Node{root, a, a1, a2, a2a, b, b1}

	// Forward traversal
	var forwardNodes []Node
	for {
		node := iterator.NextNode()
		if node == nil {
			break
		}
		forwardNodes = append(forwardNodes, node)
	}

	if len(forwardNodes) != len(expectedOrder) {
		t.Errorf("Expected %d nodes in forward traversal, got %d", len(expectedOrder), len(forwardNodes))
	}

	for i, expected := range expectedOrder {
		if i < len(forwardNodes) && forwardNodes[i] != expected {
			t.Errorf("Forward traversal position %d: expected %v, got %v",
				i, expected.(*Element).TagName(), forwardNodes[i].(*Element).TagName())
		}
	}

	// Backward traversal from end
	var backwardNodes []Node
	for {
		node := iterator.PreviousNode()
		if node == nil {
			break
		}
		backwardNodes = append(backwardNodes, node)
	}

	// Should traverse in reverse order
	expectedBackward := []Node{b1, b, a2a, a2, a1, a, root}

	if len(backwardNodes) != len(expectedBackward) {
		t.Errorf("Expected %d nodes in backward traversal, got %d", len(expectedBackward), len(backwardNodes))
	}

	for i, expected := range expectedBackward {
		if i < len(backwardNodes) && backwardNodes[i] != expected {
			t.Errorf("Backward traversal position %d: expected %v, got %v",
				i, expected.(*Element).TagName(), backwardNodes[i].(*Element).TagName())
		}
	}
}
