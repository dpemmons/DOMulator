package dom

import (
	"testing"
)

// TestNodeLength tests the DOM specification requirement for node length calculation
func TestNodeLength(t *testing.T) {
	doc := NewDocument()

	// Test DocumentType and Attr nodes - should return 0
	doctype := NewDocumentType("html", "", "", doc)
	if doctype.Length() != 0 {
		t.Errorf("Expected DocumentType length to be 0, got %d", doctype.Length())
	}

	attr := NewAttr("id", "test", doc)
	if attr.Length() != 0 {
		t.Errorf("Expected Attr length to be 0, got %d", attr.Length())
	}

	// Test CharacterData nodes - should return data length in runes (Unicode code points)
	text := NewText("Hello 世界", doc) // Contains Unicode characters
	expectedLength := len([]rune("Hello 世界"))
	if text.Length() != expectedLength {
		t.Errorf("Expected Text length to be %d, got %d", expectedLength, text.Length())
	}

	comment := NewComment("Test comment", doc)
	expectedCommentLength := len([]rune("Test comment"))
	if comment.Length() != expectedCommentLength {
		t.Errorf("Expected Comment length to be %d, got %d", expectedCommentLength, comment.Length())
	}

	pi := NewProcessingInstruction("xml", "version=\"1.0\"", doc)
	expectedPILength := len([]rune("version=\"1.0\""))
	if pi.Length() != expectedPILength {
		t.Errorf("Expected ProcessingInstruction length to be %d, got %d", expectedPILength, pi.Length())
	}

	// Test other nodes - should return number of children
	element := NewElement("div", doc)
	if element.Length() != 0 {
		t.Errorf("Expected empty Element length to be 0, got %d", element.Length())
	}

	// Add children and test length
	child1 := NewElement("span", doc)
	child2 := NewText("text", doc)
	element.AppendChild(child1)
	element.AppendChild(child2)
	if element.Length() != 2 {
		t.Errorf("Expected Element with 2 children to have length 2, got %d", element.Length())
	}

	// Test Document length
	if doc.Length() != 0 {
		t.Errorf("Expected empty Document length to be 0, got %d", doc.Length())
	}

	htmlElement := NewElement("html", doc)
	doc.AppendChild(htmlElement)
	if doc.Length() != 1 {
		t.Errorf("Expected Document with 1 child to have length 1, got %d", doc.Length())
	}

	// Test DocumentFragment length
	fragment := NewDocumentFragment(doc)
	if fragment.Length() != 0 {
		t.Errorf("Expected empty DocumentFragment length to be 0, got %d", fragment.Length())
	}

	fragment.AppendChild(NewElement("div", doc))
	fragment.AppendChild(NewText("text", doc))
	if fragment.Length() != 2 {
		t.Errorf("Expected DocumentFragment with 2 children to have length 2, got %d", fragment.Length())
	}
}

// TestNodeIsEmpty tests the IsEmpty() method
func TestNodeIsEmpty(t *testing.T) {
	doc := NewDocument()

	// Test empty nodes
	emptyText := NewText("", doc)
	if !emptyText.IsEmpty() {
		t.Errorf("Expected empty Text to be empty")
	}

	emptyElement := NewElement("div", doc)
	if !emptyElement.IsEmpty() {
		t.Errorf("Expected empty Element to be empty")
	}

	// DocumentType and Attr are always empty
	doctype := NewDocumentType("html", "", "", doc)
	if !doctype.IsEmpty() {
		t.Errorf("Expected DocumentType to be empty")
	}

	attr := NewAttr("id", "test", doc)
	if !attr.IsEmpty() {
		t.Errorf("Expected Attr to be empty")
	}

	// Test non-empty nodes
	nonEmptyText := NewText("content", doc)
	if nonEmptyText.IsEmpty() {
		t.Errorf("Expected non-empty Text to not be empty")
	}

	elementWithChild := NewElement("div", doc)
	elementWithChild.AppendChild(NewElement("span", doc))
	if elementWithChild.IsEmpty() {
		t.Errorf("Expected Element with child to not be empty")
	}
}

// TestDocumentChildOrderConstraints tests the DOM specification constraints for Document children
func TestDocumentChildOrderConstraints(t *testing.T) {
	doc := NewDocument()

	// Test that Document can only have one Element child
	html1 := NewElement("html", doc)
	html2 := NewElement("html", doc)

	doc.AppendChild(html1) // Should succeed

	// Attempting to add a second Element should panic
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic when adding second Element to Document")
		}
	}()
	doc.AppendChild(html2) // Should panic
}

// TestDocumentTypeConstraints tests DocumentType uniqueness constraint
func TestDocumentTypeConstraints(t *testing.T) {
	doc := NewDocument()

	// Test that Document can only have one DocumentType child
	doctype1 := NewDocumentType("html", "", "", doc)
	doctype2 := NewDocumentType("html", "", "", doc)

	doc.AppendChild(doctype1) // Should succeed

	// Attempting to add a second DocumentType should panic
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic when adding second DocumentType to Document")
		}
	}()
	doc.AppendChild(doctype2) // Should panic
}

// TestAttrNodeConstraints tests that Attr nodes cannot have children
func TestAttrNodeConstraints(t *testing.T) {
	doc := NewDocument()
	attr := NewAttr("id", "test", doc)
	text := NewText("value", doc)

	// Attempting to add a child to an Attr should panic
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic when adding child to Attr node")
		}
	}()
	attr.AppendChild(text) // Should panic
}

// TestCharacterDataNodeConstraints tests that CharacterData nodes cannot have children
func TestCharacterDataNodeConstraints(t *testing.T) {
	doc := NewDocument()
	text := NewText("content", doc)
	child := NewElement("span", doc)

	// Attempting to add a child to a Text node should panic
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic when adding child to Text node")
		}
	}()
	text.AppendChild(child) // Should panic
}

// TestDocumentTypeNodeConstraints tests that DocumentType nodes cannot have children
func TestDocumentTypeNodeConstraints(t *testing.T) {
	doc := NewDocument()
	doctype := NewDocumentType("html", "", "", doc)
	child := NewElement("span", doc)

	// Attempting to add a child to a DocumentType node should panic
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic when adding child to DocumentType node")
		}
	}()
	doctype.AppendChild(child) // Should panic
}

// TestElementChildConstraints tests Element child type constraints
func TestElementChildConstraints(t *testing.T) {
	doc := NewDocument()
	element := NewElement("div", doc)
	doctype := NewDocumentType("html", "", "", doc)

	// Attempting to add a DocumentType to an Element should panic
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic when adding DocumentType to Element")
		}
	}()
	element.AppendChild(doctype) // Should panic
}

// TestCircularReferenceValidation tests that circular references are prevented
func TestCircularReferenceValidation(t *testing.T) {
	doc := NewDocument()
	parent := NewElement("div", doc)
	child := NewElement("span", doc)
	grandchild := NewElement("p", doc)

	parent.AppendChild(child)
	child.AppendChild(grandchild)

	// Attempting to make grandchild a parent of parent should panic (circular reference)
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic when creating circular reference")
		}
	}()
	grandchild.AppendChild(parent) // Should panic
}

// TestValidDocumentStructure tests a valid document structure according to DOM spec
func TestValidDocumentStructure(t *testing.T) {
	doc := NewDocument()

	// Valid document structure: comments, doctype, comments, element, comments
	comment1 := NewComment("Start comment", doc)
	doctype := NewDocumentType("html", "", "", doc)
	comment2 := NewComment("Middle comment", doc)
	html := NewElement("html", doc)
	comment3 := NewComment("End comment", doc)

	// All of these should succeed
	doc.AppendChild(comment1)
	doc.AppendChild(doctype)
	doc.AppendChild(comment2)
	doc.AppendChild(html)
	doc.AppendChild(comment3)

	// Verify structure
	children := doc.ChildNodes()
	if len(children) != 5 {
		t.Errorf("Expected Document to have 5 children, got %d", len(children))
	}

	if children[0].NodeType() != CommentNode {
		t.Errorf("Expected first child to be Comment, got %v", children[0].NodeType())
	}
	if children[1].NodeType() != DocumentTypeNode {
		t.Errorf("Expected second child to be DocumentType, got %v", children[1].NodeType())
	}
	if children[2].NodeType() != CommentNode {
		t.Errorf("Expected third child to be Comment, got %v", children[2].NodeType())
	}
	if children[3].NodeType() != ElementNode {
		t.Errorf("Expected fourth child to be Element, got %v", children[3].NodeType())
	}
	if children[4].NodeType() != CommentNode {
		t.Errorf("Expected fifth child to be Comment, got %v", children[4].NodeType())
	}
}

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

	// Try removing a non-existent child - should panic with NotFoundError
	nonExistentChild := NewElement("a", doc)
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic when removing non-existent child")
		}
	}()
	parent.RemoveChild(nonExistentChild)
	t.Errorf("Should have panicked before reaching this point")
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

	// Insert an already parented child (child1 is already immediately before child2, so this should be a no-op)
	parent.InsertBefore(child1, child2) // child1 is already a child of parent and already before child2

	// Since child1 is already immediately before child2, no change should occur
	if parent.ChildNodes()[1] != child1 {
		t.Errorf("Expected child1 to remain at index 1, got %v at index 1", parent.ChildNodes()[1])
	}
	if parent.ChildNodes()[2] != child2 {
		t.Errorf("Expected child2 to remain at index 2, got %v at index 2", parent.ChildNodes()[2])
	}
	if len(parent.ChildNodes()) != 4 { // Should not change count
		t.Errorf("Expected parent to still have 4 children, got %d", len(parent.ChildNodes()))
	}

	// Test inserting a node that's not immediately before the reference
	// Insert child2 before anotherNewChild (should move child2 from index 2 to index 3)
	parent.InsertBefore(child2, anotherNewChild)
	if parent.ChildNodes()[2] != child2 {
		t.Errorf("Expected child2 to be at index 2 after moving, got %v", parent.ChildNodes()[2])
	}
	if parent.ChildNodes()[3] != anotherNewChild {
		t.Errorf("Expected anotherNewChild to be at index 3, got %v", parent.ChildNodes()[3])
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
