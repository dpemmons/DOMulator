package dom

import (
	"testing"
)

func TestNodeIteratorBasicProperties(t *testing.T) {
	doc := NewDocument()
	root := doc.CreateElement("div")

	iterator := NewNodeIterator(root, NodeFilterShowAll, nil)

	if iterator.Root() != root {
		t.Error("Expected iterator root to be the specified root")
	}

	if iterator.WhatToShow() != NodeFilterShowAll {
		t.Errorf("Expected whatToShow to be %d, got %d", NodeFilterShowAll, iterator.WhatToShow())
	}

	if iterator.ReferenceNode() != root {
		t.Error("Expected initial reference node to be root")
	}

	if !iterator.PointerBeforeReferenceNode() {
		t.Error("Expected pointer to be before reference node initially")
	}
}

func TestNodeIteratorNextNode(t *testing.T) {
	doc := NewDocument()
	root := doc.CreateElement("div")
	child1 := doc.CreateElement("span")
	child2 := doc.CreateElement("p")
	text1 := doc.CreateTextNode("Hello")
	text2 := doc.CreateTextNode("World")

	// Build structure: root > child1 > text1, child2 > text2
	root.AppendChild(child1)
	root.AppendChild(child2)
	child1.AppendChild(text1)
	child2.AppendChild(text2)

	iterator := NewNodeIterator(root, NodeFilterShowAll, nil)

	// First call should return root itself
	node := iterator.NextNode()
	if node != root {
		t.Errorf("Expected first NextNode to return root, got %v", node)
	}

	// Second call should return child1
	node = iterator.NextNode()
	if node != child1 {
		t.Errorf("Expected second NextNode to return child1, got %v", node)
	}

	// Third call should return text1
	node = iterator.NextNode()
	if node != text1 {
		t.Errorf("Expected third NextNode to return text1, got %v", node)
	}

	// Fourth call should return child2
	node = iterator.NextNode()
	if node != child2 {
		t.Errorf("Expected fourth NextNode to return child2, got %v", node)
	}

	// Fifth call should return text2
	node = iterator.NextNode()
	if node != text2 {
		t.Errorf("Expected fifth NextNode to return text2, got %v", node)
	}

	// Sixth call should return nil (end of iteration)
	node = iterator.NextNode()
	if node != nil {
		t.Errorf("Expected sixth NextNode to return nil, got %v", node)
	}
}

func TestNodeIteratorPreviousNode(t *testing.T) {
	doc := NewDocument()
	root := doc.CreateElement("div")
	child1 := doc.CreateElement("span")
	child2 := doc.CreateElement("p")
	text1 := doc.CreateTextNode("Hello")
	text2 := doc.CreateTextNode("World")

	// Build structure: root > child1 > text1, child2 > text2
	root.AppendChild(child1)
	root.AppendChild(child2)
	child1.AppendChild(text1)
	child2.AppendChild(text2)

	iterator := NewNodeIterator(root, NodeFilterShowAll, nil)

	// Move to end first
	var lastNode Node
	for {
		node := iterator.NextNode()
		if node == nil {
			break
		}
		lastNode = node
	}

	// Verify we're at text2
	if lastNode != text2 {
		t.Errorf("Expected to end at text2, got %v", lastNode)
	}

	// Per spec: pointer is after reference (text2)
	// PreviousNode should toggle pointer to before reference and return text2
	node := iterator.PreviousNode()
	if node != text2 {
		t.Errorf("Expected first PreviousNode to return text2, got %v", node)
	}

	// Now pointer is before reference (text2)
	// PreviousNode should move to child2
	node = iterator.PreviousNode()
	if node != child2 {
		t.Errorf("Expected second PreviousNode to return child2, got %v", node)
	}

	node = iterator.PreviousNode()
	if node != text1 {
		t.Errorf("Expected third PreviousNode to return text1, got %v", node)
	}

	node = iterator.PreviousNode()
	if node != child1 {
		t.Errorf("Expected fourth PreviousNode to return child1, got %v", node)
	}

	node = iterator.PreviousNode()
	if node != root {
		t.Errorf("Expected fifth PreviousNode to return root, got %v", node)
	}

	// Should return nil when at beginning
	node = iterator.PreviousNode()
	if node != nil {
		t.Errorf("Expected sixth PreviousNode to return nil, got %v", node)
	}
}

func TestNodeIteratorWithElementFilter(t *testing.T) {
	doc := NewDocument()
	root := doc.CreateElement("div")
	child1 := doc.CreateElement("span")
	child2 := doc.CreateElement("p")
	text1 := doc.CreateTextNode("Hello")
	text2 := doc.CreateTextNode("World")

	// Build structure: root > child1 > text1, child2 > text2
	root.AppendChild(child1)
	root.AppendChild(child2)
	child1.AppendChild(text1)
	child2.AppendChild(text2)

	// Only show elements
	iterator := NewNodeIterator(root, NodeFilterShowElement, nil)

	// Should return only elements
	var nodes []Node
	for {
		node := iterator.NextNode()
		if node == nil {
			break
		}
		nodes = append(nodes, node)
	}

	expectedNodes := []Node{root, child1, child2}
	if len(nodes) != len(expectedNodes) {
		t.Errorf("Expected %d nodes, got %d", len(expectedNodes), len(nodes))
	}

	for i, expected := range expectedNodes {
		if i >= len(nodes) || nodes[i] != expected {
			t.Errorf("Expected node %d to be %v, got %v", i, expected, nodes[i])
		}
	}
}

func TestNodeIteratorWithTextFilter(t *testing.T) {
	doc := NewDocument()
	root := doc.CreateElement("div")
	child1 := doc.CreateElement("span")
	text1 := doc.CreateTextNode("Hello")
	text2 := doc.CreateTextNode("World")

	// Build structure: root > child1 > text1, text2
	root.AppendChild(child1)
	root.AppendChild(text2)
	child1.AppendChild(text1)

	// Only show text nodes
	iterator := NewNodeIterator(root, NodeFilterShowText, nil)

	// Should return only text nodes
	var nodes []Node
	for {
		node := iterator.NextNode()
		if node == nil {
			break
		}
		nodes = append(nodes, node)
	}

	expectedNodes := []Node{text1, text2}
	if len(nodes) != len(expectedNodes) {
		t.Errorf("Expected %d nodes, got %d", len(expectedNodes), len(nodes))
	}

	for i, expected := range expectedNodes {
		if i >= len(nodes) || nodes[i] != expected {
			t.Errorf("Expected node %d to be %v, got %v", i, expected, nodes[i])
		}
	}
}

func TestNodeIteratorWithCustomFilter(t *testing.T) {
	doc := NewDocument()
	root := doc.CreateElement("div")
	div1 := doc.CreateElement("div")
	span1 := doc.CreateElement("span")
	div2 := doc.CreateElement("div")

	div1.SetAttribute("class", "target")
	div2.SetAttribute("class", "target")

	// Build structure
	root.AppendChild(div1)
	root.AppendChild(span1)
	root.AppendChild(div2)

	// Custom filter: only accept divs with class="target"
	customFilter := NodeFilterFunc(func(node Node) int {
		if elem, ok := node.(*Element); ok {
			// Use LocalName for case-insensitive matching against "div"
			if elem.LocalName() == "div" && elem.GetAttribute("class") == "target" {
				return NodeFilterAccept
			}
		}
		return NodeFilterSkip
	})

	iterator := NewNodeIterator(root, NodeFilterShowElement, customFilter)

	// Should return only the divs with class="target"
	var nodes []Node
	for {
		node := iterator.NextNode()
		if node == nil {
			break
		}
		nodes = append(nodes, node)
	}

	expectedNodes := []Node{div1, div2}
	if len(nodes) != len(expectedNodes) {
		t.Errorf("Expected %d nodes, got %d", len(expectedNodes), len(nodes))
	}

	for i, expected := range expectedNodes {
		if i >= len(nodes) || nodes[i] != expected {
			t.Errorf("Expected node %d to be %v, got %v", i, expected, nodes[i])
		}
	}
}

func TestNodeIteratorWithRejectingFilter(t *testing.T) {
	doc := NewDocument()
	root := doc.CreateElement("div")
	child1 := doc.CreateElement("span")
	child2 := doc.CreateElement("p")

	root.AppendChild(child1)
	root.AppendChild(child2)

	// Filter that rejects span elements
	rejectingFilter := NodeFilterFunc(func(node Node) int {
		if elem, ok := node.(*Element); ok {
			// Use LocalName for case-insensitive matching against "span"
			if elem.LocalName() == "span" {
				return NodeFilterReject
			}
		}
		return NodeFilterAccept
	})

	iterator := NewNodeIterator(root, NodeFilterShowElement, rejectingFilter)

	// Should return root and p, but not span
	var nodes []Node
	for {
		node := iterator.NextNode()
		if node == nil {
			break
		}
		nodes = append(nodes, node)
	}

	expectedNodes := []Node{root, child2}
	if len(nodes) != len(expectedNodes) {
		t.Errorf("Expected %d nodes, got %d", len(expectedNodes), len(nodes))
	}

	for i, expected := range expectedNodes {
		if i >= len(nodes) || nodes[i] != expected {
			t.Errorf("Expected node %d to be %v, got %v", i, expected, nodes[i])
		}
	}
}

func TestNodeIteratorDetach(t *testing.T) {
	doc := NewDocument()
	root := doc.CreateElement("div")

	iterator := NewNodeIterator(root, NodeFilterShowAll, nil)

	// Detach should do nothing per spec
	iterator.Detach()

	// Iterator should still work
	node := iterator.NextNode()
	if node != root {
		t.Errorf("Expected iterator to still work after detach, got %v", node)
	}
}

func TestNodeIteratorMixedTraversal(t *testing.T) {
	doc := NewDocument()
	root := doc.CreateElement("div")
	child1 := doc.CreateElement("span")
	child2 := doc.CreateElement("p")

	root.AppendChild(child1)
	root.AppendChild(child2)

	iterator := NewNodeIterator(root, NodeFilterShowElement, nil)

	// Move forward
	node := iterator.NextNode() // root (pointer was before, now after)
	if node != root {
		t.Errorf("Expected first node to be root, got %v", node)
	}

	node = iterator.NextNode() // child1
	if node != child1 {
		t.Errorf("Expected second node to be child1, got %v", node)
	}

	// Move backward - pointer is after child1
	// First PreviousNode toggles pointer to before child1 and returns child1
	node = iterator.PreviousNode()
	if node != child1 {
		t.Errorf("Expected PreviousNode to return child1 (toggle pointer), got %v", node)
	}

	// Now pointer is before child1, so PreviousNode moves to root
	node = iterator.PreviousNode()
	if node != root {
		t.Errorf("Expected PreviousNode to return root, got %v", node)
	}

	// Move forward again - pointer is before root
	// First NextNode toggles pointer to after root and returns root
	node = iterator.NextNode()
	if node != root {
		t.Errorf("Expected NextNode to return root (toggle pointer), got %v", node)
	}

	// Now pointer is after root, so NextNode moves to child1
	node = iterator.NextNode()
	if node != child1 {
		t.Errorf("Expected NextNode to return child1, got %v", node)
	}

	node = iterator.NextNode() // child2
	if node != child2 {
		t.Errorf("Expected NextNode to return child2, got %v", node)
	}
}

func TestNodeIteratorEmptyRoot(t *testing.T) {
	doc := NewDocument()
	root := doc.CreateElement("div")
	// Root has no children

	iterator := NewNodeIterator(root, NodeFilterShowAll, nil)

	// Should return root first
	node := iterator.NextNode()
	if node != root {
		t.Errorf("Expected first node to be root, got %v", node)
	}

	// Should return nil after root
	node = iterator.NextNode()
	if node != nil {
		t.Errorf("Expected second node to be nil, got %v", node)
	}

	// PreviousNode should return root again
	node = iterator.PreviousNode()
	if node != root {
		t.Errorf("Expected PreviousNode to return root, got %v", node)
	}
}

func TestNodeIteratorFromDocument(t *testing.T) {
	doc := NewDocument()
	html := doc.CreateElement("html")
	body := doc.CreateElement("body")
	div := doc.CreateElement("div")

	doc.AppendChild(html)
	html.AppendChild(body)
	body.AppendChild(div)

	// Create iterator with document as root
	iterator := doc.CreateNodeIterator(doc, NodeFilterShowElement, nil)

	// Should traverse document structure
	var nodes []Node
	for {
		node := iterator.NextNode()
		if node == nil {
			break
		}
		if node.NodeType() == ElementNode {
			nodes = append(nodes, node)
		}
	}

	expectedNodes := []Node{html, body, div}
	if len(nodes) != len(expectedNodes) {
		t.Errorf("Expected %d element nodes, got %d", len(expectedNodes), len(nodes))
	}
}

// Test spec-compliant whatToShow bit checking
func TestNodeIteratorWhatToShowBitChecking(t *testing.T) {
	doc := NewDocument()
	root := doc.CreateElement("div")
	elem := doc.CreateElement("span")
	text := doc.CreateTextNode("Hello")
	comment := doc.CreateComment("comment")

	root.AppendChild(elem)
	root.AppendChild(text)
	root.AppendChild(comment)

	// Test with only ELEMENT bit set (0x1)
	iterator := NewNodeIterator(root, 0x1, nil)

	// Should return root (element) first
	node := iterator.NextNode()
	if node != root {
		t.Errorf("Expected root element, got %v", node)
	}

	// Should return span element
	node = iterator.NextNode()
	if node != elem {
		t.Errorf("Expected span element, got %v", node)
	}

	// Should skip text and comment nodes
	node = iterator.NextNode()
	if node != nil {
		t.Errorf("Expected nil (no more elements), got %v", node)
	}

	// Test with TEXT bit set (0x4)
	iterator2 := NewNodeIterator(root, 0x4, nil)

	// Should skip root element
	node = iterator2.NextNode()
	if node != text {
		t.Errorf("Expected text node, got %v", node)
	}

	// Should be no more text nodes
	node = iterator2.NextNode()
	if node != nil {
		t.Errorf("Expected nil (no more text nodes), got %v", node)
	}
}

// Test recursive filter invocation (should throw InvalidStateError)
func TestNodeIteratorRecursiveFilterInvocation(t *testing.T) {
	doc := NewDocument()
	root := doc.CreateElement("div")
	child := doc.CreateElement("span")
	root.AppendChild(child)

	var iterator *NodeIterator

	// Create a filter that tries to call NextNode recursively
	recursiveFilter := NodeFilterFunc(func(node Node) int {
		if node == child {
			// This should panic with InvalidStateError
			iterator.NextNode()
		}
		return NodeFilterAccept
	})

	iterator = NewNodeIterator(root, NodeFilterShowAll, recursiveFilter)

	// NextNode should eventually call the filter with child node
	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(*DOMException); ok {
				if err.Name() != "InvalidStateError" {
					t.Errorf("Expected InvalidStateError, got %s", err.Name())
				}
			} else {
				t.Errorf("Expected DOMException, got %v", r)
			}
		} else {
			t.Error("Expected panic with InvalidStateError")
		}
	}()

	// This should work for root
	node := iterator.NextNode()
	if node != root {
		t.Errorf("Expected root, got %v", node)
	}

	// This should trigger the recursive call and panic
	iterator.NextNode()
}

// Test filter exception handling
func TestNodeIteratorFilterException(t *testing.T) {
	doc := NewDocument()
	root := doc.CreateElement("div")
	child := doc.CreateElement("span")
	root.AppendChild(child)

	// Create a filter that panics
	panicFilter := NodeFilterFunc(func(node Node) int {
		if node == child {
			panic("test error")
		}
		return NodeFilterAccept
	})

	iterator := NewNodeIterator(root, NodeFilterShowAll, panicFilter)

	// First node should work
	node := iterator.NextNode()
	if node != root {
		t.Errorf("Expected root, got %v", node)
	}

	// Second node should trigger panic
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic from filter")
		} else if r != "test error" {
			t.Errorf("Expected 'test error', got %v", r)
		}
	}()

	iterator.NextNode()
}

// Test pre-remove algorithm
func TestNodeIteratorPreRemove(t *testing.T) {
	doc := NewDocument()
	root := doc.CreateElement("div")
	child1 := doc.CreateElement("span")
	child2 := doc.CreateElement("p")
	child3 := doc.CreateElement("a")

	root.AppendChild(child1)
	root.AppendChild(child2)
	root.AppendChild(child3)

	iterator := NewNodeIterator(root, NodeFilterShowAll, nil)

	// Move to child2
	iterator.NextNode() // root
	iterator.NextNode() // child1
	iterator.NextNode() // child2

	// Reference is now at child2
	if iterator.ReferenceNode() != child2 {
		t.Errorf("Expected reference to be child2, got %v", iterator.ReferenceNode())
	}

	// Simulate removing child2 (pointer is after reference)
	iterator.preRemoveNode(child2)

	// Reference should move to child1 (previous sibling's last descendant)
	if iterator.ReferenceNode() != child1 {
		t.Errorf("Expected reference to be child1 after removal, got %v", iterator.ReferenceNode())
	}
}

// Test pre-remove with pointer before reference
func TestNodeIteratorPreRemovePointerBefore(t *testing.T) {
	doc := NewDocument()
	root := doc.CreateElement("div")
	child1 := doc.CreateElement("span")
	child2 := doc.CreateElement("p")
	child3 := doc.CreateElement("a")

	root.AppendChild(child1)
	root.AppendChild(child2)
	root.AppendChild(child3)

	iterator := NewNodeIterator(root, NodeFilterShowAll, nil)

	// Move to child2 with pointer before
	iterator.NextNode()     // root (pointer before->after)
	iterator.NextNode()     // child1 (pointer after)
	iterator.NextNode()     // child2 (pointer after)
	iterator.PreviousNode() // child2 (pointer after->before)

	// Verify state
	if iterator.ReferenceNode() != child2 {
		t.Errorf("Reference should be child2, got %v", iterator.ReferenceNode())
	}
	if !iterator.PointerBeforeReferenceNode() {
		t.Error("Expected pointer to be before reference")
	}

	// Simulate removing child2 (pointer before reference)
	// Per spec step 2: reference moves to next node (child3)
	iterator.preRemoveNode(child2)

	if iterator.ReferenceNode() != child3 {
		t.Errorf("After removal, reference should be child3, got %v", iterator.ReferenceNode())
	}
	// Pointer position should be preserved
	if !iterator.PointerBeforeReferenceNode() {
		t.Error("Pointer should remain before reference after removal")
	}
}

// Test NodeIterator with all node types
func TestNodeIteratorAllNodeTypes(t *testing.T) {
	doc := NewDocument()
	root := doc.CreateElement("div")

	// Add various node types
	elem := doc.CreateElement("span")
	text := doc.CreateTextNode("text")
	comment := doc.CreateComment("comment")
	cdata, _ := doc.CreateCDATASection("cdata")
	pi, _ := doc.CreateProcessingInstruction("target", "data")

	root.AppendChild(elem)
	root.AppendChild(text)
	root.AppendChild(comment)
	if cdata != nil {
		root.AppendChild(cdata)
	}
	if pi != nil {
		root.AppendChild(pi)
	}

	// Test with NodeFilterShowAll
	iterator := NewNodeIterator(root, NodeFilterShowAll, nil)

	var nodeTypes []int
	for {
		node := iterator.NextNode()
		if node == nil {
			break
		}
		nodeTypes = append(nodeTypes, int(node.NodeType()))
	}

	expectedTypes := []int{
		int(ElementNode), // root
		int(ElementNode), // span
		int(TextNode),    // text
		int(CommentNode), // comment
	}

	if cdata != nil {
		expectedTypes = append(expectedTypes, int(CDataSectionNode))
	}
	if pi != nil {
		expectedTypes = append(expectedTypes, int(ProcessingInstructionNode))
	}

	if len(nodeTypes) != len(expectedTypes) {
		t.Errorf("Expected %d nodes, got %d", len(expectedTypes), len(nodeTypes))
		t.Logf("Node types found: %v", nodeTypes)
		t.Logf("Expected types: %v", expectedTypes)
	}

	for i := 0; i < len(expectedTypes) && i < len(nodeTypes); i++ {
		if nodeTypes[i] != expectedTypes[i] {
			t.Errorf("At index %d: expected node type %d, got %d", i, expectedTypes[i], nodeTypes[i])
		}
	}
}

// Test specification example scenario
func TestNodeIteratorSpecExample(t *testing.T) {
	// Create structure from spec example
	doc := NewDocument()
	root := doc.CreateElement("root")

	// Create A, B, C nodes
	nodeA := doc.CreateElement("A")
	nodeB := doc.CreateElement("B")
	nodeC := doc.CreateElement("C")

	root.AppendChild(nodeA)
	root.AppendChild(nodeB)
	root.AppendChild(nodeC)

	// Create iterator that only shows elements
	iterator := NewNodeIterator(root, NodeFilterShowElement, nil)

	// Traverse forward
	nodes := []Node{
		iterator.NextNode(), // root
		iterator.NextNode(), // A
		iterator.NextNode(), // B
		iterator.NextNode(), // C
	}

	expectedNodes := []Node{root, nodeA, nodeB, nodeC}
	for i, expected := range expectedNodes {
		if nodes[i] != expected {
			t.Errorf("Forward traversal: expected %v at position %d, got %v", expected, i, nodes[i])
		}
	}

	// Traverse backward
	backNodes := []Node{
		iterator.PreviousNode(), // C
		iterator.PreviousNode(), // B
		iterator.PreviousNode(), // A
		iterator.PreviousNode(), // root
	}

	expectedBack := []Node{nodeC, nodeB, nodeA, root}
	for i, expected := range expectedBack {
		if backNodes[i] != expected {
			t.Errorf("Backward traversal: expected %v at position %d, got %v", expected, i, backNodes[i])
		}
	}
}
