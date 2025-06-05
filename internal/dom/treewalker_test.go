package dom

import (
	"testing"
)

func TestTreeWalkerBasicProperties(t *testing.T) {
	doc := NewDocument()
	root := doc.CreateElement("div")

	walker := NewTreeWalker(root, NodeFilterShowAll, nil)

	if walker.Root() != root {
		t.Error("Expected walker root to be the specified root")
	}

	if walker.WhatToShow() != NodeFilterShowAll {
		t.Errorf("Expected whatToShow to be %d, got %d", NodeFilterShowAll, walker.WhatToShow())
	}

	if walker.CurrentNode() != root {
		t.Error("Expected initial current node to be root")
	}
}

func TestTreeWalkerSetCurrentNode(t *testing.T) {
	doc := NewDocument()
	root := doc.CreateElement("div")
	child := doc.CreateElement("span")
	root.AppendChild(child)

	walker := NewTreeWalker(root, NodeFilterShowAll, nil)

	// Set current node to child
	walker.SetCurrentNode(child)
	if walker.CurrentNode() != child {
		t.Error("Expected current node to be set to child")
	}

	// Setting to nil should be ignored
	walker.SetCurrentNode(nil)
	if walker.CurrentNode() != child {
		t.Error("Expected current node to remain child when setting to nil")
	}
}

func TestTreeWalkerParentNode(t *testing.T) {
	doc := NewDocument()
	root := doc.CreateElement("div")
	child1 := doc.CreateElement("span")
	grandchild := doc.CreateElement("p")

	root.AppendChild(child1)
	child1.AppendChild(grandchild)

	walker := NewTreeWalker(root, NodeFilterShowElement, nil)

	// Start at grandchild
	walker.SetCurrentNode(grandchild)

	// Move to parent (child1)
	parent := walker.ParentNode()
	if parent != child1 {
		t.Errorf("Expected ParentNode to return child1, got %v", parent)
	}
	if walker.CurrentNode() != child1 {
		t.Error("Expected current node to be updated to child1")
	}

	// Move to parent (root)
	parent = walker.ParentNode()
	if parent != root {
		t.Errorf("Expected ParentNode to return root, got %v", parent)
	}
	if walker.CurrentNode() != root {
		t.Error("Expected current node to be updated to root")
	}

	// At root, should return nil
	parent = walker.ParentNode()
	if parent != nil {
		t.Errorf("Expected ParentNode at root to return nil, got %v", parent)
	}
	if walker.CurrentNode() != root {
		t.Error("Expected current node to remain root")
	}
}

func TestTreeWalkerFirstChild(t *testing.T) {
	doc := NewDocument()
	root := doc.CreateElement("div")
	child1 := doc.CreateElement("span")
	child2 := doc.CreateElement("p")
	text1 := doc.CreateTextNode("Hello")

	root.AppendChild(child1)
	root.AppendChild(child2)
	child1.AppendChild(text1)

	walker := NewTreeWalker(root, NodeFilterShowElement, nil)

	// First child of root should be child1
	firstChild := walker.FirstChild()
	if firstChild != child1 {
		t.Errorf("Expected FirstChild to return child1, got %v", firstChild)
	}
	if walker.CurrentNode() != child1 {
		t.Error("Expected current node to be updated to child1")
	}

	// First child of child1 should be nil (text node filtered out)
	firstChild = walker.FirstChild()
	if firstChild != nil {
		t.Errorf("Expected FirstChild of child1 to return nil (text filtered), got %v", firstChild)
	}
	if walker.CurrentNode() != child1 {
		t.Error("Expected current node to remain child1")
	}
}

func TestTreeWalkerLastChild(t *testing.T) {
	doc := NewDocument()
	root := doc.CreateElement("div")
	child1 := doc.CreateElement("span")
	child2 := doc.CreateElement("p")
	child3 := doc.CreateElement("a")

	root.AppendChild(child1)
	root.AppendChild(child2)
	root.AppendChild(child3)

	walker := NewTreeWalker(root, NodeFilterShowElement, nil)

	// Last child of root should be child3
	lastChild := walker.LastChild()
	if lastChild != child3 {
		t.Errorf("Expected LastChild to return child3, got %v", lastChild)
	}
	if walker.CurrentNode() != child3 {
		t.Error("Expected current node to be updated to child3")
	}
}

func TestTreeWalkerNextSibling(t *testing.T) {
	doc := NewDocument()
	root := doc.CreateElement("div")
	child1 := doc.CreateElement("span")
	child2 := doc.CreateElement("p")
	child3 := doc.CreateElement("a")

	root.AppendChild(child1)
	root.AppendChild(child2)
	root.AppendChild(child3)

	walker := NewTreeWalker(root, NodeFilterShowElement, nil)

	// Start at child1
	walker.SetCurrentNode(child1)

	// Next sibling should be child2
	nextSibling := walker.NextSibling()
	if nextSibling != child2 {
		t.Errorf("Expected NextSibling to return child2, got %v", nextSibling)
	}
	if walker.CurrentNode() != child2 {
		t.Error("Expected current node to be updated to child2")
	}

	// Next sibling should be child3
	nextSibling = walker.NextSibling()
	if nextSibling != child3 {
		t.Errorf("Expected NextSibling to return child3, got %v", nextSibling)
	}

	// No more siblings
	nextSibling = walker.NextSibling()
	if nextSibling != nil {
		t.Errorf("Expected NextSibling to return nil, got %v", nextSibling)
	}
}

func TestTreeWalkerPreviousSibling(t *testing.T) {
	doc := NewDocument()
	root := doc.CreateElement("div")
	child1 := doc.CreateElement("span")
	child2 := doc.CreateElement("p")
	child3 := doc.CreateElement("a")

	root.AppendChild(child1)
	root.AppendChild(child2)
	root.AppendChild(child3)

	walker := NewTreeWalker(root, NodeFilterShowElement, nil)

	// Start at child3
	walker.SetCurrentNode(child3)

	// Previous sibling should be child2
	prevSibling := walker.PreviousSibling()
	if prevSibling != child2 {
		t.Errorf("Expected PreviousSibling to return child2, got %v", prevSibling)
	}
	if walker.CurrentNode() != child2 {
		t.Error("Expected current node to be updated to child2")
	}

	// Previous sibling should be child1
	prevSibling = walker.PreviousSibling()
	if prevSibling != child1 {
		t.Errorf("Expected PreviousSibling to return child1, got %v", prevSibling)
	}

	// No more siblings
	prevSibling = walker.PreviousSibling()
	if prevSibling != nil {
		t.Errorf("Expected PreviousSibling to return nil, got %v", prevSibling)
	}
}

func TestTreeWalkerNextNode(t *testing.T) {
	doc := NewDocument()
	root := doc.CreateElement("div")
	child1 := doc.CreateElement("span")
	grandchild := doc.CreateElement("em")
	child2 := doc.CreateElement("p")

	root.AppendChild(child1)
	root.AppendChild(child2)
	child1.AppendChild(grandchild)

	walker := NewTreeWalker(root, NodeFilterShowElement, nil)

	// Start at root, next should be child1
	next := walker.NextNode()
	if next != child1 {
		t.Errorf("Expected NextNode to return child1, got %v", next)
	}

	// From child1, next should be grandchild
	next = walker.NextNode()
	if next != grandchild {
		t.Errorf("Expected NextNode to return grandchild, got %v", next)
	}

	// From grandchild, next should be child2
	next = walker.NextNode()
	if next != child2 {
		t.Errorf("Expected NextNode to return child2, got %v", next)
	}

	// No more nodes
	next = walker.NextNode()
	if next != nil {
		t.Errorf("Expected NextNode to return nil, got %v", next)
	}
}

func TestTreeWalkerPreviousNode(t *testing.T) {
	doc := NewDocument()
	root := doc.CreateElement("div")
	child1 := doc.CreateElement("span")
	grandchild := doc.CreateElement("em")
	child2 := doc.CreateElement("p")

	root.AppendChild(child1)
	root.AppendChild(child2)
	child1.AppendChild(grandchild)

	walker := NewTreeWalker(root, NodeFilterShowElement, nil)

	// Start at child2
	walker.SetCurrentNode(child2)

	// Previous should be grandchild
	prev := walker.PreviousNode()
	if prev != grandchild {
		t.Errorf("Expected PreviousNode to return grandchild, got %v", prev)
	}

	// From grandchild, previous should be child1
	prev = walker.PreviousNode()
	if prev != child1 {
		t.Errorf("Expected PreviousNode to return child1, got %v", prev)
	}

	// From child1, previous should be root
	prev = walker.PreviousNode()
	if prev != root {
		t.Errorf("Expected PreviousNode to return root, got %v", prev)
	}

	// At root, should return nil
	prev = walker.PreviousNode()
	if prev != nil {
		t.Errorf("Expected PreviousNode at root to return nil, got %v", prev)
	}
}

func TestTreeWalkerWithTextFilter(t *testing.T) {
	doc := NewDocument()
	root := doc.CreateElement("div")
	child1 := doc.CreateElement("span")
	text1 := doc.CreateTextNode("Hello")
	text2 := doc.CreateTextNode("World")

	root.AppendChild(child1)
	root.AppendChild(text2)
	child1.AppendChild(text1)

	walker := NewTreeWalker(root, NodeFilterShowText, nil)

	// Start at root, first child should be text1
	firstChild := walker.FirstChild()
	if firstChild != text1 {
		t.Errorf("Expected FirstChild to return text1, got %v", firstChild)
	}

	// Next sibling should be text2
	nextSibling := walker.NextSibling()
	if nextSibling != text2 {
		t.Errorf("Expected NextSibling to return text2, got %v", nextSibling)
	}
}

func TestTreeWalkerWithCustomFilter(t *testing.T) {
	doc := NewDocument()
	root := doc.CreateElement("div")
	div1 := doc.CreateElement("div")
	span1 := doc.CreateElement("span")
	div2 := doc.CreateElement("div")

	div1.SetAttribute("class", "target")
	div2.SetAttribute("class", "target")

	root.AppendChild(div1)
	root.AppendChild(span1)
	root.AppendChild(div2)

	// Custom filter: only accept divs with class="target"
	customFilter := NodeFilterFunc(func(node Node) int {
		if elem, ok := node.(*Element); ok {
			if elem.TagName() == "div" && elem.GetAttribute("class") == "target" {
				return NodeFilterAccept
			}
		}
		return NodeFilterSkip
	})

	walker := NewTreeWalker(root, NodeFilterShowElement, customFilter)

	// First child should be div1
	firstChild := walker.FirstChild()
	if firstChild != div1 {
		t.Errorf("Expected FirstChild to return div1, got %v", firstChild)
	}

	// Next sibling should be div2 (span1 is skipped)
	nextSibling := walker.NextSibling()
	if nextSibling != div2 {
		t.Errorf("Expected NextSibling to return div2, got %v", nextSibling)
	}
}

func TestTreeWalkerWithRejectingFilter(t *testing.T) {
	doc := NewDocument()
	root := doc.CreateElement("div")
	span1 := doc.CreateElement("span")
	span2 := doc.CreateElement("span")
	div1 := doc.CreateElement("div")

	root.AppendChild(span1)
	root.AppendChild(span2)
	root.AppendChild(div1)

	// Filter that rejects span elements
	rejectingFilter := NodeFilterFunc(func(node Node) int {
		if elem, ok := node.(*Element); ok {
			if elem.TagName() == "span" {
				return NodeFilterReject
			}
		}
		return NodeFilterAccept
	})

	walker := NewTreeWalker(root, NodeFilterShowElement, rejectingFilter)

	// First child should be div1 (spans are rejected)
	firstChild := walker.FirstChild()
	if firstChild != div1 {
		t.Errorf("Expected FirstChild to return div1 (spans rejected), got %v", firstChild)
	}
}

func TestTreeWalkerComplexNavigation(t *testing.T) {
	doc := NewDocument()
	root := doc.CreateElement("div")
	child1 := doc.CreateElement("span")
	child2 := doc.CreateElement("p")
	grandchild1 := doc.CreateElement("em")
	grandchild2 := doc.CreateElement("strong")

	root.AppendChild(child1)
	root.AppendChild(child2)
	child1.AppendChild(grandchild1)
	child2.AppendChild(grandchild2)

	walker := NewTreeWalker(root, NodeFilterShowElement, nil)

	// Navigate to grandchild1
	walker.FirstChild() // child1
	walker.FirstChild() // grandchild1

	if walker.CurrentNode() != grandchild1 {
		t.Error("Expected to be at grandchild1")
	}

	// Go to parent, then next sibling, then first child
	walker.ParentNode()  // child1
	walker.NextSibling() // child2
	walker.FirstChild()  // grandchild2

	if walker.CurrentNode() != grandchild2 {
		t.Error("Expected to be at grandchild2")
	}

	// Go back to root via parent chain
	walker.ParentNode() // child2
	walker.ParentNode() // root

	if walker.CurrentNode() != root {
		t.Error("Expected to be back at root")
	}
}

func TestTreeWalkerFromDocument(t *testing.T) {
	doc := NewDocument()
	html := doc.CreateElement("html")
	body := doc.CreateElement("body")
	div := doc.CreateElement("div")

	doc.AppendChild(html)
	html.AppendChild(body)
	body.AppendChild(div)

	// Create walker with document as root
	walker := doc.CreateTreeWalker(doc, NodeFilterShowElement, nil)

	// Navigate through structure
	firstChild := walker.FirstChild() // html
	if firstChild != html {
		t.Errorf("Expected first child to be html, got %v", firstChild)
	}

	nextChild := walker.FirstChild() // body
	if nextChild != body {
		t.Errorf("Expected next child to be body, got %v", nextChild)
	}

	lastChild := walker.FirstChild() // div
	if lastChild != div {
		t.Errorf("Expected last child to be div, got %v", lastChild)
	}
}

func TestTreeWalkerEmptyStructure(t *testing.T) {
	doc := NewDocument()
	root := doc.CreateElement("div")
	// Root has no children

	walker := NewTreeWalker(root, NodeFilterShowAll, nil)

	// All navigation methods should return nil
	if walker.FirstChild() != nil {
		t.Error("Expected FirstChild to return nil for empty structure")
	}

	if walker.LastChild() != nil {
		t.Error("Expected LastChild to return nil for empty structure")
	}

	if walker.NextSibling() != nil {
		t.Error("Expected NextSibling to return nil for root")
	}

	if walker.PreviousSibling() != nil {
		t.Error("Expected PreviousSibling to return nil for root")
	}

	if walker.NextNode() != nil {
		t.Error("Expected NextNode to return nil for empty structure")
	}

	if walker.PreviousNode() != nil {
		t.Error("Expected PreviousNode to return nil at root")
	}
}
