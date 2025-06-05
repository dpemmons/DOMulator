package dom

import (
	"testing"
)

// TestTreeWalkerSpecCompliance tests TreeWalker implementation against WHATWG DOM specification
func TestTreeWalkerSpecCompliance(t *testing.T) {
	t.Run("Interface Compliance", testTreeWalkerInterfaceCompliance)
	t.Run("ParentNode Algorithm", testTreeWalkerParentNodeAlgorithm)
	t.Run("TraverseChildren Algorithm", testTreeWalkerTraverseChildrenAlgorithm)
	t.Run("TraverseSiblings Algorithm", testTreeWalkerTraverseSiblingsAlgorithm)
	t.Run("PreviousNode Algorithm", testTreeWalkerPreviousNodeAlgorithm)
	t.Run("NextNode Algorithm", testTreeWalkerNextNodeAlgorithm)
	t.Run("Filter Reject vs Skip", testTreeWalkerFilterRejectVsSkip)
	t.Run("Root Node Filtering", testTreeWalkerRootNodeFiltering)
	t.Run("Complex Traversal Patterns", testTreeWalkerComplexTraversalPatterns)
}

func testTreeWalkerInterfaceCompliance(t *testing.T) {
	doc := NewDocument()
	root := doc.CreateElement("div")
	walker := NewTreeWalker(root, NodeFilterShowAll, nil)

	// Test all required properties exist and have correct values
	if walker.Root() != root {
		t.Error("Root() should return the specified root node")
	}

	if walker.WhatToShow() != NodeFilterShowAll {
		t.Errorf("WhatToShow() should return %d, got %d", NodeFilterShowAll, walker.WhatToShow())
	}

	if walker.Filter() == nil {
		t.Error("Filter() should not return nil")
	}

	if walker.CurrentNode() != root {
		t.Error("CurrentNode() should initially be the root")
	}

	// Test currentNode setter
	child := doc.CreateElement("span")
	root.AppendChild(child)
	walker.SetCurrentNode(child)
	if walker.CurrentNode() != child {
		t.Error("SetCurrentNode should update current node")
	}

	// Test setting to nil should be ignored
	walker.SetCurrentNode(nil)
	if walker.CurrentNode() != child {
		t.Error("SetCurrentNode(nil) should be ignored")
	}
}

func testTreeWalkerParentNodeAlgorithm(t *testing.T) {
	doc := NewDocument()
	root := doc.CreateElement("div")
	child1 := doc.CreateElement("span")
	grandchild := doc.CreateElement("p")

	root.AppendChild(child1)
	child1.AppendChild(grandchild)

	// Test with filter that accepts all elements
	walker := NewTreeWalker(root, NodeFilterShowElement, nil)
	walker.SetCurrentNode(grandchild)

	// Move to parent (child1)
	parent := walker.ParentNode()
	if parent != child1 {
		t.Errorf("Expected parent to be child1, got %v", parent)
	}
	if walker.CurrentNode() != child1 {
		t.Error("Current node should be updated to child1")
	}

	// Move to parent (root)
	parent = walker.ParentNode()
	if parent != root {
		t.Errorf("Expected parent to be root, got %v", parent)
	}

	// At root, should return nil (cannot go beyond root)
	parent = walker.ParentNode()
	if parent != nil {
		t.Errorf("ParentNode() at root should return nil, got %v", parent)
	}

	// Test with rejecting filter
	rejectFilter := NodeFilterFunc(func(node Node) int {
		if elem, ok := node.(*Element); ok && elem.LocalName() == "span" { // Use LocalName
			return NodeFilterReject
		}
		return NodeFilterAccept
	})
	walker2 := NewTreeWalker(root, NodeFilterShowElement, rejectFilter)
	walker2.SetCurrentNode(grandchild)

	// Should skip rejected span and go directly to root
	parent = walker2.ParentNode()
	if parent != root {
		t.Errorf("Expected parent to be root (skipping rejected span), got %v", parent)
	}
}

func testTreeWalkerTraverseChildrenAlgorithm(t *testing.T) {
	doc := NewDocument()
	root := doc.CreateElement("div")
	child1 := doc.CreateElement("span")
	child2 := doc.CreateElement("p")
	grandchild1 := doc.CreateElement("a")
	grandchild2 := doc.CreateElement("b")
	text1 := doc.CreateTextNode("text1")
	text2 := doc.CreateTextNode("text2")

	root.AppendChild(child1)
	root.AppendChild(text1)
	root.AppendChild(child2)
	child1.AppendChild(grandchild1)
	child1.AppendChild(text2)
	child2.AppendChild(grandchild2)

	// Test FirstChild with element filter
	walker := NewTreeWalker(root, NodeFilterShowElement, nil)
	firstChild := walker.FirstChild()
	if firstChild != child1 {
		t.Errorf("FirstChild should return child1, got %v", firstChild)
	}

	// Test LastChild with element filter
	walker.SetCurrentNode(root)
	lastChild := walker.LastChild()
	if lastChild != child2 {
		t.Errorf("LastChild should return child2, got %v", lastChild)
	}

	// Test with SKIP filter - should traverse into skipped nodes
	skipFilter := NodeFilterFunc(func(node Node) int {
		if elem, ok := node.(*Element); ok && elem.LocalName() == "span" { // Use LocalName
			return NodeFilterSkip
		}
		return NodeFilterAccept
	})
	walker2 := NewTreeWalker(root, NodeFilterShowElement, skipFilter)
	firstChild = walker2.FirstChild()
	// Should skip span but find its child (grandchild1)
	if firstChild != grandchild1 {
		t.Errorf("FirstChild with SKIP filter should return grandchild1, got %v", firstChild)
	}

	// Test with REJECT filter - should not traverse into rejected nodes
	rejectFilter := NodeFilterFunc(func(node Node) int {
		if elem, ok := node.(*Element); ok && elem.LocalName() == "span" { // Use LocalName
			return NodeFilterReject
		}
		return NodeFilterAccept
	})
	walker3 := NewTreeWalker(root, NodeFilterShowElement, rejectFilter)
	firstChild = walker3.FirstChild()
	// Should reject span and its subtree, go to next child (child2)
	if firstChild != child2 {
		t.Errorf("FirstChild with REJECT filter should return child2, got %v", firstChild)
	}
}

func testTreeWalkerTraverseSiblingsAlgorithm(t *testing.T) {
	doc := NewDocument()
	root := doc.CreateElement("div")
	child1 := doc.CreateElement("span")
	child2 := doc.CreateElement("p")
	child3 := doc.CreateElement("a")
	grandchild1 := doc.CreateElement("em")
	grandchild2 := doc.CreateElement("strong")

	root.AppendChild(child1)
	root.AppendChild(child2)
	root.AppendChild(child3)
	child1.AppendChild(grandchild1)
	child2.AppendChild(grandchild2)

	// Test NextSibling
	walker := NewTreeWalker(root, NodeFilterShowElement, nil)
	walker.SetCurrentNode(child1)

	nextSibling := walker.NextSibling()
	if nextSibling != child2 {
		t.Errorf("NextSibling should return child2, got %v", nextSibling)
	}

	nextSibling = walker.NextSibling()
	if nextSibling != child3 {
		t.Errorf("NextSibling should return child3, got %v", nextSibling)
	}

	nextSibling = walker.NextSibling()
	if nextSibling != nil {
		t.Errorf("NextSibling should return nil, got %v", nextSibling)
	}

	// Test PreviousSibling
	walker.SetCurrentNode(child3)
	prevSibling := walker.PreviousSibling()
	if prevSibling != child2 {
		t.Errorf("PreviousSibling should return child2, got %v", prevSibling)
	}

	// Test from root should return null
	walker.SetCurrentNode(root)
	nextSibling = walker.NextSibling()
	if nextSibling != nil {
		t.Errorf("NextSibling from root should return nil, got %v", nextSibling)
	}

	prevSibling = walker.PreviousSibling()
	if prevSibling != nil {
		t.Errorf("PreviousSibling from root should return nil, got %v", prevSibling)
	}

	// Test with SKIP filter that descends into children
	skipFilter := NodeFilterFunc(func(node Node) int {
		if elem, ok := node.(*Element); ok && elem.LocalName() == "p" { // Use LocalName
			return NodeFilterSkip
		}
		return NodeFilterAccept
	})
	walker2 := NewTreeWalker(root, NodeFilterShowElement, skipFilter)
	walker2.SetCurrentNode(child1)

	nextSibling = walker2.NextSibling()
	// Should skip p but find its child (grandchild2)
	if nextSibling != grandchild2 {
		t.Errorf("NextSibling with SKIP filter should return grandchild2, got %v", nextSibling)
	}
}

func testTreeWalkerPreviousNodeAlgorithm(t *testing.T) {
	doc := NewDocument()
	root := doc.CreateElement("div")
	child1 := doc.CreateElement("span")
	child2 := doc.CreateElement("p")
	grandchild1 := doc.CreateElement("a")
	grandchild2 := doc.CreateElement("b")

	root.AppendChild(child1)
	root.AppendChild(child2)
	child1.AppendChild(grandchild1)
	child2.AppendChild(grandchild2)

	walker := NewTreeWalker(root, NodeFilterShowElement, nil)

	// Start at child2, should go to grandchild1 (deepest last descendant of previous sibling)
	walker.SetCurrentNode(child2)
	prev := walker.PreviousNode()
	if prev != grandchild1 {
		t.Errorf("PreviousNode from child2 should return grandchild1, got %v", prev)
	}

	// From grandchild1, should go to child1 (parent)
	prev = walker.PreviousNode()
	if prev != child1 {
		t.Errorf("PreviousNode from grandchild1 should return child1, got %v", prev)
	}

	// From child1, should go to root
	prev = walker.PreviousNode()
	if prev != root {
		t.Errorf("PreviousNode from child1 should return root, got %v", prev)
	}

	// From root, should return nil
	prev = walker.PreviousNode()
	if prev != nil {
		t.Errorf("PreviousNode from root should return nil, got %v", prev)
	}

	// Test with REJECT filter
	rejectFilter := NodeFilterFunc(func(node Node) int {
		if elem, ok := node.(*Element); ok && elem.LocalName() == "span" { // Use LocalName
			return NodeFilterReject
		}
		return NodeFilterAccept
	})
	walker2 := NewTreeWalker(root, NodeFilterShowElement, rejectFilter)
	walker2.SetCurrentNode(child2)

	// Should skip rejected span subtree entirely
	prev = walker2.PreviousNode()
	if prev != root {
		t.Errorf("PreviousNode with REJECT filter should skip to root, got %v", prev)
	}
}

func testTreeWalkerNextNodeAlgorithm(t *testing.T) {
	doc := NewDocument()
	root := doc.CreateElement("div")
	child1 := doc.CreateElement("span")
	child2 := doc.CreateElement("p")
	grandchild1 := doc.CreateElement("a")
	grandchild2 := doc.CreateElement("b")

	root.AppendChild(child1)
	root.AppendChild(child2)
	child1.AppendChild(grandchild1)
	child2.AppendChild(grandchild2)

	walker := NewTreeWalker(root, NodeFilterShowElement, nil)

	// Start at root, should go to child1
	next := walker.NextNode()
	if next != child1 {
		t.Errorf("NextNode from root should return child1, got %v", next)
	}

	// From child1, should go to grandchild1 (first child)
	next = walker.NextNode()
	if next != grandchild1 {
		t.Errorf("NextNode from child1 should return grandchild1, got %v", next)
	}

	// From grandchild1, should go to child2 (next in document order)
	next = walker.NextNode()
	if next != child2 {
		t.Errorf("NextNode from grandchild1 should return child2, got %v", next)
	}

	// From child2, should go to grandchild2
	next = walker.NextNode()
	if next != grandchild2 {
		t.Errorf("NextNode from child2 should return grandchild2, got %v", next)
	}

	// From grandchild2, should return nil (end of tree)
	next = walker.NextNode()
	if next != nil {
		t.Errorf("NextNode from grandchild2 should return nil, got %v", next)
	}

	// Test with REJECT filter
	rejectFilter := NodeFilterFunc(func(node Node) int {
		if elem, ok := node.(*Element); ok && elem.LocalName() == "span" { // Use LocalName
			return NodeFilterReject
		}
		return NodeFilterAccept
	})
	walker2 := NewTreeWalker(root, NodeFilterShowElement, rejectFilter)

	// Should skip entire span subtree
	next = walker2.NextNode()
	if next != child2 {
		t.Errorf("NextNode with REJECT filter should skip to child2, got %v", next)
	}
}

func testTreeWalkerFilterRejectVsSkip(t *testing.T) {
	doc := NewDocument()
	root := doc.CreateElement("div")
	target := doc.CreateElement("target")
	child1 := doc.CreateElement("span")
	child2 := doc.CreateElement("p")
	grandchild := doc.CreateElement("a")

	root.AppendChild(target)
	root.AppendChild(child1)
	root.AppendChild(child2)
	target.AppendChild(grandchild)

	// Test SKIP behavior - should traverse into children
	skipFilter := NodeFilterFunc(func(node Node) int {
		if elem, ok := node.(*Element); ok && elem.LocalName() == "target" { // Use LocalName
			return NodeFilterSkip
		}
		return NodeFilterAccept
	})
	walker1 := NewTreeWalker(root, NodeFilterShowElement, skipFilter)

	// Should find grandchild inside skipped target
	next := walker1.NextNode()
	if next != grandchild {
		t.Errorf("NextNode with SKIP should find grandchild, got %v", next)
	}

	// Test REJECT behavior - should not traverse into children
	rejectFilter := NodeFilterFunc(func(node Node) int {
		if elem, ok := node.(*Element); ok && elem.LocalName() == "target" { // Use LocalName
			return NodeFilterReject
		}
		return NodeFilterAccept
	})
	walker2 := NewTreeWalker(root, NodeFilterShowElement, rejectFilter)

	// Should skip entire target subtree and go to child1
	next = walker2.NextNode()
	if next != child1 {
		t.Errorf("NextNode with REJECT should skip to child1, got %v", next)
	}
}

func testTreeWalkerRootNodeFiltering(t *testing.T) {
	doc := NewDocument()
	root := doc.CreateElement("div")
	child := doc.CreateElement("span")
	root.AppendChild(child)

	// Create filter that would reject the root
	rejectRootFilter := NodeFilterFunc(func(node Node) int {
		if elem, ok := node.(*Element); ok && elem.LocalName() == "div" { // Use LocalName
			return NodeFilterReject
		}
		return NodeFilterAccept
	})

	walker := NewTreeWalker(root, NodeFilterShowElement, rejectRootFilter)

	// Root should not be affected by filtering in ParentNode
	walker.SetCurrentNode(child)
	parent := walker.ParentNode()
	if parent != nil {
		t.Errorf("ParentNode should not return root when it would be rejected, got %v", parent)
	}

	// But NextNode from root should work (root is current, not being filtered)
	walker.SetCurrentNode(root)
	next := walker.NextNode()
	if next != child {
		t.Errorf("NextNode from root should find child, got %v", next)
	}
}

func testTreeWalkerComplexTraversalPatterns(t *testing.T) {
	doc := NewDocument()
	root := doc.CreateElement("div")

	// Create complex tree structure
	level1a := doc.CreateElement("level1a")
	level1b := doc.CreateElement("level1b")
	level2a := doc.CreateElement("level2a")
	level2b := doc.CreateElement("level2b")
	level2c := doc.CreateElement("level2c")
	level3a := doc.CreateElement("level3a")

	root.AppendChild(level1a)
	root.AppendChild(level1b)
	level1a.AppendChild(level2a)
	level1a.AppendChild(level2b)
	level1b.AppendChild(level2c)
	level2a.AppendChild(level3a)

	// Test complex filter that skips some and rejects others
	complexFilter := NodeFilterFunc(func(node Node) int {
		if elem, ok := node.(*Element); ok {
			switch elem.LocalName() { // Use LocalName
			case "level1a":
				return NodeFilterSkip // Skip but traverse children
			case "level2b":
				return NodeFilterReject // Reject entire subtree
			default:
				return NodeFilterAccept
			}
		}
		return NodeFilterAccept
	})

	walker := NewTreeWalker(root, NodeFilterShowElement, complexFilter)

	// Expected traversal order: root -> level2a -> level3a -> level1b -> level2c
	expected := []Node{level2a, level3a, level1b, level2c}

	for i, expectedNode := range expected {
		next := walker.NextNode()
		if next != expectedNode {
			t.Errorf("NextNode step %d should return %v, got %v", i+1, expectedNode, next)
		}
	}

	// Should be at end
	next := walker.NextNode()
	if next != nil {
		t.Errorf("NextNode at end should return nil, got %v", next)
	}

	// Test reverse traversal
	walker.SetCurrentNode(level2c)
	expected = []Node{level1b, level3a, level2a, root}

	for i, expectedNode := range expected {
		prev := walker.PreviousNode()
		if prev != expectedNode {
			t.Errorf("PreviousNode step %d should return %v, got %v", i+1, expectedNode, prev)
		}
	}

	// Should be at beginning
	prev := walker.PreviousNode()
	if prev != nil {
		t.Errorf("PreviousNode at beginning should return nil, got %v", prev)
	}
}

// TestTreeWalkerSpecExamples tests specific examples from the WHATWG DOM spec
func TestTreeWalkerSpecExamples(t *testing.T) {
	// Create a document structure similar to spec examples
	doc := NewDocument()
	html := doc.CreateElement("html")
	head := doc.CreateElement("head")
	title := doc.CreateElement("title")
	titleText := doc.CreateTextNode("Example")
	body := doc.CreateElement("body")
	h1 := doc.CreateElement("h1")
	h1Text := doc.CreateTextNode("Heading")
	p := doc.CreateElement("p")
	pText := doc.CreateTextNode("Paragraph")

	doc.AppendChild(html)
	html.AppendChild(head)
	html.AppendChild(body)
	head.AppendChild(title)
	title.AppendChild(titleText)
	body.AppendChild(h1)
	body.AppendChild(p)
	h1.AppendChild(h1Text)
	p.AppendChild(pText)

	// Test TreeWalker with element filter starting from body
	walker := NewTreeWalker(body, NodeFilterShowElement, nil)

	// First child should be h1
	firstChild := walker.FirstChild()
	if firstChild != h1 {
		t.Errorf("FirstChild should return h1, got %v", firstChild)
	}

	// Next sibling should be p
	nextSibling := walker.NextSibling()
	if nextSibling != p {
		t.Errorf("NextSibling should return p, got %v", nextSibling)
	}

	// Parent should be body
	parent := walker.ParentNode()
	if parent != body {
		t.Errorf("ParentNode should return body, got %v", parent)
	}

	// Test with text filter
	textWalker := NewTreeWalker(body, NodeFilterShowText, nil)

	// First text node should be h1Text
	firstText := textWalker.FirstChild()
	if firstText != h1Text {
		t.Errorf("FirstChild with text filter should return h1Text, got %v", firstText)
	}

	// Next text node should be pText
	nextText := textWalker.NextNode()
	if nextText != pText {
		t.Errorf("NextNode with text filter should return pText, got %v", nextText)
	}
}
