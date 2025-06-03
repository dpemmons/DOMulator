package dom

import (
	"testing"
)

func TestNodeAppendChild(t *testing.T) {
	doc := NewDocument()
	parent := NewElement("div", doc)
	child1 := NewElement("span", doc)
	child2 := NewElement("p", doc)

	parent.AppendChild(child1)
	if parent.FirstChild() != child1 {
		t.Errorf("Expected first child to be child1, got %v", parent.FirstChild())
	}
	if child1.ParentNode() != parent {
		t.Errorf("Expected child1's parent to be parent, got %v", child1.ParentNode())
	}
	if len(parent.ChildNodes()) != 1 {
		t.Errorf("Expected parent to have 1 child, got %d", len(parent.ChildNodes()))
	}

	parent.AppendChild(child2)
	if parent.LastChild() != child2 {
		t.Errorf("Expected last child to be child2, got %v", parent.LastChild())
	}
	if child2.ParentNode() != parent {
		t.Errorf("Expected child2's parent to be parent, got %v", child2.ParentNode())
	}
	if len(parent.ChildNodes()) != 2 {
		t.Errorf("Expected parent to have 2 children, got %d", len(parent.ChildNodes()))
	}

	// Test appending an already parented child
	grandparent := NewElement("body", doc)
	grandparent.AppendChild(parent)
	if parent.ParentNode() != grandparent {
		t.Errorf("Expected parent's parent to be grandparent, got %v", parent.ParentNode())
	}

	// Appending child1 to a new parent should remove it from the old parent
	newParent := NewElement("section", doc)
	newParent.AppendChild(child1)
	if child1.ParentNode() != newParent {
		t.Errorf("Expected child1's parent to be newParent, got %v", child1.ParentNode())
	}
	if len(parent.ChildNodes()) != 1 {
		t.Errorf("Expected parent to have 1 child after child1 moved, got %d", len(parent.ChildNodes()))
	}
	if parent.FirstChild() != child2 {
		t.Errorf("Expected parent's first child to be child2, got %v", parent.FirstChild())
	}
}

func TestNodeRemoveChild(t *testing.T) {
	doc := NewDocument()
	parent := NewElement("div", doc)
	child1 := NewElement("span", doc)
	child2 := NewElement("p", doc)

	parent.AppendChild(child1)
	parent.AppendChild(child2)

	removedChild := parent.RemoveChild(child1)
	if removedChild != child1 {
		t.Errorf("Expected removed child to be child1, got %v", removedChild)
	}
	if child1.ParentNode() != nil {
		t.Errorf("Expected child1's parent to be nil, got %v", child1.ParentNode())
	}
	if len(parent.ChildNodes()) != 1 {
		t.Errorf("Expected parent to have 1 child, got %d", len(parent.ChildNodes()))
	}
	if parent.FirstChild() != child2 {
		t.Errorf("Expected parent's first child to be child2, got %v", parent.FirstChild())
	}

	// Try removing a non-existent child
	nonExistentChild := NewElement("a", doc)
	removedNonExistent := parent.RemoveChild(nonExistentChild)
	if removedNonExistent != nil {
		t.Errorf("Expected removing non-existent child to return nil, got %v", removedNonExistent)
	}
	if len(parent.ChildNodes()) != 1 {
		t.Errorf("Expected parent to still have 1 child, got %d", len(parent.ChildNodes()))
	}
}

func TestNodeInsertBefore(t *testing.T) {
	doc := NewDocument()
	parent := NewElement("div", doc)
	child1 := NewElement("span", doc)
	child2 := NewElement("p", doc)
	newChild := NewElement("a", doc)

	parent.AppendChild(child1)
	parent.AppendChild(child2)

	// Insert before child1
	insertedChild := parent.InsertBefore(newChild, child1)
	if insertedChild != newChild {
		t.Errorf("Expected inserted child to be newChild, got %v", insertedChild)
	}
	if newChild.ParentNode() != parent {
		t.Errorf("Expected newChild's parent to be parent, got %v", newChild.ParentNode())
	}
	if parent.ChildNodes()[0] != newChild {
		t.Errorf("Expected newChild to be at index 0, got %v", parent.ChildNodes()[0])
	}
	if len(parent.ChildNodes()) != 3 {
		t.Errorf("Expected parent to have 3 children, got %d", len(parent.ChildNodes()))
	}

	// Insert at end (refChild is nil)
	anotherNewChild := NewElement("b", doc)
	parent.InsertBefore(anotherNewChild, nil)
	if parent.LastChild() != anotherNewChild {
		t.Errorf("Expected anotherNewChild to be last child, got %v", parent.LastChild())
	}
	if len(parent.ChildNodes()) != 4 {
		t.Errorf("Expected parent to have 4 children, got %d", len(parent.ChildNodes()))
	}

	// Insert an already parented child
	parent.InsertBefore(child1, child2) // child1 is already a child of parent
	if parent.ChildNodes()[2] != child1 {
		t.Errorf("Expected child1 to be at index 2, got %v", parent.ChildNodes()[2])
	}
	if len(parent.ChildNodes()) != 4 { // Should not change count
		t.Errorf("Expected parent to still have 4 children, got %d", len(parent.ChildNodes()))
	}
}

func TestNodeReplaceChild(t *testing.T) {
	doc := NewDocument()
	parent := NewElement("div", doc)
	child1 := NewElement("span", doc)
	child2 := NewElement("p", doc)
	newChild := NewElement("a", doc)

	parent.AppendChild(child1)
	parent.AppendChild(child2)

	replacedChild := parent.ReplaceChild(newChild, child1)
	if replacedChild != child1 {
		t.Errorf("Expected replaced child to be child1, got %v", replacedChild)
	}
	if child1.ParentNode() != nil {
		t.Errorf("Expected child1's parent to be nil, got %v", child1.ParentNode())
	}
	if newChild.ParentNode() != parent {
		t.Errorf("Expected newChild's parent to be parent, got %v", newChild.ParentNode())
	}
	if parent.ChildNodes()[0] != newChild {
		t.Errorf("Expected newChild to be at index 0, got %v", parent.ChildNodes()[0])
	}
	if len(parent.ChildNodes()) != 2 {
		t.Errorf("Expected parent to have 2 children, got %d", len(parent.ChildNodes()))
	}

	// Replace with an already parented child
	parent.ReplaceChild(child2, newChild) // child2 is already a child of parent
	if parent.ChildNodes()[0] != child2 {
		t.Errorf("Expected child2 to be at index 0, got %v", parent.ChildNodes()[0])
	}
	if len(parent.ChildNodes()) != 2 { // Should not change count
		t.Errorf("Expected parent to still have 2 children, got %d", len(parent.ChildNodes()))
	}
}

func TestNodeCloneNode(t *testing.T) {
	doc := NewDocument()
	parent := NewElement("div", doc)
	child1 := NewElement("span", doc)
	textNode := NewText("hello", doc)

	parent.AppendChild(child1)
	child1.AppendChild(textNode)

	// Shallow clone
	shallowClone := parent.CloneNode(false)
	if shallowClone.NodeType() != parent.NodeType() {
		t.Errorf("Expected shallow clone node type %v, got %v", parent.NodeType(), shallowClone.NodeType())
	}
	if shallowClone.NodeName() != parent.NodeName() {
		t.Errorf("Expected shallow clone node name %s, got %s", parent.NodeName(), shallowClone.NodeName())
	}
	if len(shallowClone.ChildNodes()) != 0 {
		t.Errorf("Expected shallow clone to have 0 children, got %d", len(shallowClone.ChildNodes()))
	}

	// Deep clone
	deepClone := parent.CloneNode(true)
	if deepClone.NodeType() != parent.NodeType() {
		t.Errorf("Expected deep clone node type %v, got %v", parent.NodeType(), deepClone.NodeType())
	}
	if deepClone.NodeName() != parent.NodeName() {
		t.Errorf("Expected deep clone node name %s, got %s", parent.NodeName(), deepClone.NodeName())
	}
	if len(deepClone.ChildNodes()) != 1 {
		t.Errorf("Expected deep clone to have 1 child, got %d", len(deepClone.ChildNodes()))
	}
	clonedChild := deepClone.FirstChild()
	if clonedChild.NodeName() != child1.NodeName() {
		t.Errorf("Expected cloned child name %s, got %s", child1.NodeName(), clonedChild.NodeName())
	}
	if len(clonedChild.ChildNodes()) != 1 {
		t.Errorf("Expected cloned child to have 1 child, got %d", len(clonedChild.ChildNodes()))
	}
	clonedText := clonedChild.FirstChild()
	if clonedText.NodeValue() != textNode.NodeValue() {
		t.Errorf("Expected cloned text value %s, got %s", textNode.NodeValue(), clonedText.NodeValue())
	}
}

func TestNodeTraversal(t *testing.T) {
	doc := NewDocument()
	root := NewElement("root", doc)
	child1 := NewElement("c1", doc)
	child2 := NewElement("c2", doc)
	child3 := NewElement("c3", doc)
	grandchild1 := NewElement("gc1", doc)

	root.AppendChild(child1)
	root.AppendChild(child2)
	root.AppendChild(child3)
	child2.AppendChild(grandchild1)

	// Test ParentNode
	if child1.ParentNode() != root {
		t.Errorf("Expected child1's parent to be root")
	}
	if root.ParentNode() != nil {
		t.Errorf("Expected root's parent to be nil")
	}

	// Test FirstChild
	if root.FirstChild() != child1 {
		t.Errorf("Expected root's first child to be child1")
	}
	if child3.FirstChild() != nil {
		t.Errorf("Expected child3's first child to be nil")
	}

	// Test LastChild
	if root.LastChild() != child3 {
		t.Errorf("Expected root's last child to be child3")
	}
	if child3.LastChild() != nil {
		t.Errorf("Expected child3's last child to be nil")
	}

	// Test PreviousSibling
	if child2.PreviousSibling() != child1 {
		t.Errorf("Expected child2's previous sibling to be child1")
	}
	if child1.PreviousSibling() != nil {
		t.Errorf("Expected child1's previous sibling to be nil")
	}

	// Test NextSibling
	if child2.NextSibling() != child3 {
		t.Errorf("Expected child2's next sibling to be child3")
	}
	if child3.NextSibling() != nil {
		t.Errorf("Expected child3's next sibling to be nil")
	}

	// Test ChildNodes
	childNodes := root.ChildNodes()
	if len(childNodes) != 3 {
		t.Errorf("Expected root to have 3 children, got %d", len(childNodes))
	}
	if childNodes[0] != child1 || childNodes[1] != child2 || childNodes[2] != child3 {
		t.Errorf("ChildNodes mismatch")
	}
}
