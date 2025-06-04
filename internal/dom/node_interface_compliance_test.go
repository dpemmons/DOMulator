package dom

import (
	"testing"
)

// TestNodeInterfaceCompliance tests the Node interface methods for compliance with the WHATWG DOM specification
func TestNodeInterfaceCompliance(t *testing.T) {
	t.Run("Phase 1: Core Properties & Simple Methods", func(t *testing.T) {
		t.Run("IsConnected", func(t *testing.T) {
			// Create a document and elements
			doc := NewDocument()
			element := NewElement("div", doc)
			disconnectedElement := NewElement("span", doc)

			// Test disconnected node
			if disconnectedElement.IsConnected() {
				t.Error("disconnected element should not be connected")
			}

			// Connect element to document
			doc.AppendChild(element)

			// Test connected node
			if !element.IsConnected() {
				t.Error("element connected to document should be connected")
			}

			// Test child of connected node
			child := NewElement("p", doc)
			element.AppendChild(child)

			if !child.IsConnected() {
				t.Error("child of connected element should be connected")
			}

			// Test after removing from document
			doc.RemoveChild(element)
			if element.IsConnected() {
				t.Error("element removed from document should not be connected")
			}

			// Child should also be disconnected
			if child.IsConnected() {
				t.Error("child of disconnected element should not be connected")
			}
		})

		t.Run("ParentElement", func(t *testing.T) {
			doc := NewDocument()
			parent := NewElement("div", doc)
			child := NewElement("span", doc)
			textNode := NewText("text content", doc)

			// Test node with no parent
			if parent.ParentElement() != nil {
				t.Error("node with no parent should return nil for ParentElement")
			}

			// Test child with Element parent
			parent.AppendChild(child)
			if child.ParentElement() != parent {
				t.Error("child should return parent element for ParentElement")
			}

			// Test node with Document parent (not an Element)
			doc.AppendChild(parent)
			if parent.ParentElement() != nil {
				t.Error("node with Document parent should return nil for ParentElement")
			}

			// Test text node with Element parent
			child.AppendChild(textNode)
			if textNode.ParentElement() != child {
				t.Error("text node should return parent element for ParentElement")
			}

			// Test text node with Document parent (should return nil)
			// Note: Text nodes cannot be direct children of Document, so we create a more realistic scenario
			docFragment := NewDocumentFragment(doc)
			doc.AppendChild(docFragment)
			docFragmentText := NewText("fragment text", doc)
			docFragment.AppendChild(docFragmentText)

			// DocumentFragment is not an Element, so text node should return nil for ParentElement
			if docFragmentText.ParentElement() != nil {
				t.Error("text node with DocumentFragment parent should return nil for ParentElement")
			}
		})

		t.Run("BaseURI", func(t *testing.T) {
			doc := NewDocument()
			element := NewElement("div", doc)

			// Test element with document
			baseURI := element.BaseURI()
			if baseURI == "" {
				t.Error("element should have a base URI")
			}

			// For now, we expect the default "about:blank"
			if baseURI != "about:blank" {
				t.Errorf("expected base URI 'about:blank', got '%s'", baseURI)
			}

			// Test orphaned node
			orphan := NewElement("span", nil)
			if orphan.BaseURI() != "" {
				t.Error("orphaned node should have empty base URI")
			}
		})

		t.Run("HasChildNodes", func(t *testing.T) {
			doc := NewDocument()
			element := NewElement("div", doc)

			// Test empty element
			if element.HasChildNodes() {
				t.Error("empty element should not have child nodes")
			}

			// Add a child
			child := NewElement("span", doc)
			element.AppendChild(child)

			// Test element with children
			if !element.HasChildNodes() {
				t.Error("element with children should have child nodes")
			}

			// Remove child
			element.RemoveChild(child)

			// Test after removal
			if element.HasChildNodes() {
				t.Error("element after removing all children should not have child nodes")
			}

			// Test with text node
			textNode := NewText("content", doc)
			element.AppendChild(textNode)

			if !element.HasChildNodes() {
				t.Error("element with text node should have child nodes")
			}

			// Test text node itself (should not have children)
			if textNode.HasChildNodes() {
				t.Error("text node should not have child nodes")
			}
		})
	})

	t.Run("NodeValue getter/setter", func(t *testing.T) {
		doc := NewDocument()

		// Test with Text node
		textNode := NewText("initial text", doc)
		if textNode.NodeValue() != "initial text" {
			t.Errorf("Text nodeValue should be 'initial text', got '%s'", textNode.NodeValue())
		}

		textNode.SetNodeValue("updated text")
		if textNode.NodeValue() != "updated text" {
			t.Errorf("Text nodeValue after SetNodeValue should be 'updated text', got '%s'", textNode.NodeValue())
		}

		// Test with Comment node
		comment := NewComment("initial comment", doc)
		if comment.NodeValue() != "initial comment" {
			t.Errorf("Comment nodeValue should be 'initial comment', got '%s'", comment.NodeValue())
		}

		comment.SetNodeValue("updated comment")
		if comment.NodeValue() != "updated comment" {
			t.Errorf("Comment nodeValue after SetNodeValue should be 'updated comment', got '%s'", comment.NodeValue())
		}

		// Test with Attr node
		attr := NewAttr("id", "test-id", doc)
		if attr.NodeValue() != "test-id" {
			t.Errorf("Attr nodeValue should be 'test-id', got '%s'", attr.NodeValue())
		}

		attr.SetNodeValue("new-id")
		if attr.NodeValue() != "new-id" {
			t.Errorf("Attr nodeValue after SetNodeValue should be 'new-id', got '%s'", attr.NodeValue())
		}

		// Test with Element node (should return empty string)
		element := NewElement("div", doc)
		if element.NodeValue() != "" {
			t.Errorf("Element nodeValue should be empty string, got '%s'", element.NodeValue())
		}

		// SetNodeValue on Element should do nothing
		element.SetNodeValue("should be ignored")
		if element.NodeValue() != "" {
			t.Errorf("Element nodeValue should still be empty string after SetNodeValue, got '%s'", element.NodeValue())
		}

		// Test with Document node (should return empty string)
		if doc.NodeValue() != "" {
			t.Errorf("Document nodeValue should be empty string, got '%s'", doc.NodeValue())
		}

		// SetNodeValue on Document should do nothing
		doc.SetNodeValue("should be ignored")
		if doc.NodeValue() != "" {
			t.Errorf("Document nodeValue should still be empty string after SetNodeValue, got '%s'", doc.NodeValue())
		}
	})

	t.Run("TextContent getter/setter", func(t *testing.T) {
		doc := NewDocument()

		// Test with Text node
		textNode := NewText("text content", doc)
		if textNode.TextContent() != "text content" {
			t.Errorf("Text textContent should be 'text content', got '%s'", textNode.TextContent())
		}

		textNode.SetTextContent("new text content")
		if textNode.TextContent() != "new text content" {
			t.Errorf("Text textContent after SetTextContent should be 'new text content', got '%s'", textNode.TextContent())
		}

		// Test with Comment node
		comment := NewComment("comment content", doc)
		if comment.TextContent() != "comment content" {
			t.Errorf("Comment textContent should be 'comment content', got '%s'", comment.TextContent())
		}

		comment.SetTextContent("new comment content")
		if comment.TextContent() != "new comment content" {
			t.Errorf("Comment textContent after SetTextContent should be 'new comment content', got '%s'", comment.TextContent())
		}

		// Test with Element (should collect descendant text)
		element := NewElement("div", doc)
		text1 := NewText("Hello ", doc)
		text2 := NewText("World", doc)
		element.AppendChild(text1)
		element.AppendChild(text2)

		if element.TextContent() != "Hello World" {
			t.Errorf("Element textContent should be 'Hello World', got '%s'", element.TextContent())
		}

		// SetTextContent on Element should replace all children with single text node
		element.SetTextContent("New content")
		if element.TextContent() != "New content" {
			t.Errorf("Element textContent after SetTextContent should be 'New content', got '%s'", element.TextContent())
		}

		// Should have only one child (the text node)
		children := element.ChildNodes()
		if children.Length() != 1 {
			t.Errorf("Element should have 1 child after SetTextContent, got %d", children.Length())
		}

		if children.Item(0).NodeType() != TextNode {
			t.Error("Element's only child should be a Text node after SetTextContent")
		}

		// Test with nested elements
		parent := NewElement("div", doc)
		child1 := NewElement("span", doc)
		child2 := NewElement("em", doc)
		text3 := NewText("Text1", doc)
		text4 := NewText("Text2", doc)
		text5 := NewText("Text3", doc)

		child1.AppendChild(text3)
		child2.AppendChild(text4)
		parent.AppendChild(child1)
		parent.AppendChild(text5)
		parent.AppendChild(child2)

		if parent.TextContent() != "Text1Text3Text2" {
			t.Errorf("Nested element textContent should be 'Text1Text3Text2', got '%s'", parent.TextContent())
		}

		// Test with Attr node
		attr := NewAttr("class", "my-class", doc)
		if attr.TextContent() != "my-class" {
			t.Errorf("Attr textContent should be 'my-class', got '%s'", attr.TextContent())
		}

		attr.SetTextContent("new-class")
		if attr.TextContent() != "new-class" {
			t.Errorf("Attr textContent after SetTextContent should be 'new-class', got '%s'", attr.TextContent())
		}

		// Test with DocumentFragment
		fragment := NewDocumentFragment(doc)
		fragText1 := NewText("Frag1", doc)
		fragText2 := NewText("Frag2", doc)
		fragment.AppendChild(fragText1)
		fragment.AppendChild(fragText2)

		if fragment.TextContent() != "Frag1Frag2" {
			t.Errorf("DocumentFragment textContent should be 'Frag1Frag2', got '%s'", fragment.TextContent())
		}

		fragment.SetTextContent("New fragment content")
		if fragment.TextContent() != "New fragment content" {
			t.Errorf("DocumentFragment textContent after SetTextContent should be 'New fragment content', got '%s'", fragment.TextContent())
		}

		// Test with Document (should return empty string)
		if doc.TextContent() != "" {
			t.Errorf("Document textContent should be empty string, got '%s'", doc.TextContent())
		}

		// SetTextContent on Document should do nothing
		doc.SetTextContent("should be ignored")
		if doc.TextContent() != "" {
			t.Errorf("Document textContent should still be empty string after SetTextContent, got '%s'", doc.TextContent())
		}
	})

	t.Run("Normalize method", func(t *testing.T) {
		doc := NewDocument()
		parent := NewElement("div", doc)

		// Test removing empty text nodes
		text1 := NewText("", doc) // Empty text node
		text2 := NewText("Hello", doc)
		text3 := NewText("", doc) // Empty text node
		text4 := NewText("World", doc)
		text5 := NewText("", doc) // Empty text node

		parent.AppendChild(text1)
		parent.AppendChild(text2)
		parent.AppendChild(text3)
		parent.AppendChild(text4)
		parent.AppendChild(text5)

		// Before normalization: 5 children
		if parent.ChildNodes().Length() != 5 {
			t.Errorf("Expected 5 children before normalization, got %d", parent.ChildNodes().Length())
		}

		parent.Normalize()

		// After normalization: empty text nodes should be removed, adjacent text nodes should be merged
		children := parent.ChildNodes()
		if children.Length() != 1 {
			t.Errorf("Expected 1 child after normalization, got %d", children.Length())
		}

		if children.Item(0).NodeType() != TextNode {
			t.Error("Remaining child should be a Text node")
		}

		if children.Item(0).NodeValue() != "HelloWorld" {
			t.Errorf("Normalized text should be 'HelloWorld', got '%s'", children.Item(0).NodeValue())
		}

		// Test with mixed content (elements and text nodes)
		parent2 := NewElement("div", doc)
		text6 := NewText("Start", doc)
		element1 := NewElement("span", doc)
		text7 := NewText("Middle1", doc)
		text8 := NewText("Middle2", doc)
		element2 := NewElement("em", doc)
		text9 := NewText("End1", doc)
		text10 := NewText("End2", doc)

		parent2.AppendChild(text6)
		parent2.AppendChild(element1)
		parent2.AppendChild(text7)
		parent2.AppendChild(text8)
		parent2.AppendChild(element2)
		parent2.AppendChild(text9)
		parent2.AppendChild(text10)

		// Before normalization: 7 children
		if parent2.ChildNodes().Length() != 7 {
			t.Errorf("Expected 7 children before normalization, got %d", parent2.ChildNodes().Length())
		}

		parent2.Normalize()

		// After normalization: should have 5 children (text, element, text, element, text)
		children2 := parent2.ChildNodes()
		if children2.Length() != 5 {
			t.Errorf("Expected 5 children after normalization, got %d", children2.Length())
		}

		// Check the structure
		if children2.Item(0).NodeValue() != "Start" {
			t.Errorf("First child should be 'Start', got '%s'", children2.Item(0).NodeValue())
		}

		if children2.Item(1).NodeType() != ElementNode {
			t.Error("Second child should be an Element")
		}

		if children2.Item(2).NodeValue() != "Middle1Middle2" {
			t.Errorf("Third child should be 'Middle1Middle2', got '%s'", children2.Item(2).NodeValue())
		}

		if children2.Item(3).NodeType() != ElementNode {
			t.Error("Fourth child should be an Element")
		}

		if children2.Item(4).NodeValue() != "End1End2" {
			t.Errorf("Fifth child should be 'End1End2', got '%s'", children2.Item(4).NodeValue())
		}

		// Test normalization with no text nodes
		parent3 := NewElement("div", doc)
		elem1 := NewElement("p", doc)
		elem2 := NewElement("span", doc)
		parent3.AppendChild(elem1)
		parent3.AppendChild(elem2)

		lengthBefore := parent3.ChildNodes().Length()
		parent3.Normalize()
		lengthAfter := parent3.ChildNodes().Length()

		if lengthBefore != lengthAfter {
			t.Error("Normalization should not affect non-text nodes")
		}

		// Test empty parent
		parent4 := NewElement("div", doc)
		parent4.Normalize() // Should not crash

		if parent4.ChildNodes().Length() != 0 {
			t.Error("Empty parent should remain empty after normalization")
		}
	})
}

// TestNodeTypeConstants tests that the node type constants match the DOM specification
func TestNodeTypeConstants(t *testing.T) {
	tests := []struct {
		name     string
		nodeType NodeType
		expected int
	}{
		{"ELEMENT_NODE", ElementNode, 1},
		{"ATTRIBUTE_NODE", AttributeNode, 2},
		{"TEXT_NODE", TextNode, 3},
		{"CDATA_SECTION_NODE", CDataSectionNode, 4},
		{"ENTITY_REFERENCE_NODE", EntityReferenceNode, 5},
		{"ENTITY_NODE", EntityNode, 6},
		{"PROCESSING_INSTRUCTION_NODE", ProcessingInstructionNode, 7},
		{"COMMENT_NODE", CommentNode, 8},
		{"DOCUMENT_NODE", DocumentNode, 9},
		{"DOCUMENT_TYPE_NODE", DocumentTypeNode, 10},
		{"DOCUMENT_FRAGMENT_NODE", DocumentFragmentNode, 11},
		{"NOTATION_NODE", NotationNode, 12},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if int(test.nodeType) != test.expected {
				t.Errorf("expected %s to be %d, got %d", test.name, test.expected, int(test.nodeType))
			}
		})
	}
}

// TestDocumentPositionConstants tests the document position constants for compareDocumentPosition
func TestDocumentPositionConstants(t *testing.T) {
	tests := []struct {
		name     string
		constant int
		expected int
	}{
		{"DOCUMENT_POSITION_DISCONNECTED", DOCUMENT_POSITION_DISCONNECTED, 0x01},
		{"DOCUMENT_POSITION_PRECEDING", DOCUMENT_POSITION_PRECEDING, 0x02},
		{"DOCUMENT_POSITION_FOLLOWING", DOCUMENT_POSITION_FOLLOWING, 0x04},
		{"DOCUMENT_POSITION_CONTAINS", DOCUMENT_POSITION_CONTAINS, 0x08},
		{"DOCUMENT_POSITION_CONTAINED_BY", DOCUMENT_POSITION_CONTAINED_BY, 0x10},
		{"DOCUMENT_POSITION_IMPLEMENTATION_SPECIFIC", DOCUMENT_POSITION_IMPLEMENTATION_SPECIFIC, 0x20},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.constant != test.expected {
				t.Errorf("expected %s to be 0x%02x, got 0x%02x", test.name, test.expected, test.constant)
			}
		})
	}
}

// TestNodeInterfaceMethods tests basic Node interface methods
func TestNodeInterfaceMethods(t *testing.T) {
	t.Run("NodeType", func(t *testing.T) {
		doc := NewDocument()
		element := NewElement("div", doc)
		text := NewText("content", doc)
		comment := NewComment("comment", doc)

		if doc.NodeType() != DocumentNode {
			t.Errorf("expected Document NodeType to be %d, got %d", DocumentNode, doc.NodeType())
		}

		if element.NodeType() != ElementNode {
			t.Errorf("expected Element NodeType to be %d, got %d", ElementNode, element.NodeType())
		}

		if text.NodeType() != TextNode {
			t.Errorf("expected Text NodeType to be %d, got %d", TextNode, text.NodeType())
		}

		if comment.NodeType() != CommentNode {
			t.Errorf("expected Comment NodeType to be %d, got %d", CommentNode, comment.NodeType())
		}
	})

	t.Run("NodeName", func(t *testing.T) {
		doc := NewDocument()
		element := NewElement("div", doc)
		text := NewText("content", doc)
		comment := NewComment("comment", doc)

		if doc.NodeName() != "#document" {
			t.Errorf("expected Document NodeName to be '#document', got '%s'", doc.NodeName())
		}

		if element.NodeName() != "DIV" {
			t.Errorf("expected Element NodeName to be 'DIV', got '%s'", element.NodeName())
		}

		if text.NodeName() != "#text" {
			t.Errorf("expected Text NodeName to be '#text', got '%s'", text.NodeName())
		}

		if comment.NodeName() != "#comment" {
			t.Errorf("expected Comment NodeName to be '#comment', got '%s'", comment.NodeName())
		}
	})

	t.Run("OwnerDocument", func(t *testing.T) {
		doc := NewDocument()
		element := NewElement("div", doc)
		text := NewText("content", doc)

		// Document should return nil for OwnerDocument per WHATWG DOM spec
		if doc.OwnerDocument() != nil {
			t.Error("Document should return nil for OwnerDocument per WHATWG DOM spec")
		}

		if element.OwnerDocument() != doc {
			t.Error("Element should have the correct owner document")
		}

		if text.OwnerDocument() != doc {
			t.Error("Text node should have the correct owner document")
		}
	})
}

// TestNodeTreeNavigation tests tree navigation methods
func TestNodeTreeNavigation(t *testing.T) {
	doc := NewDocument()
	root := NewElement("root", doc)
	child1 := NewElement("child1", doc)
	child2 := NewElement("child2", doc)
	child3 := NewElement("child3", doc)

	doc.AppendChild(root)
	root.AppendChild(child1)
	root.AppendChild(child2)
	root.AppendChild(child3)

	t.Run("ParentNode", func(t *testing.T) {
		if root.ParentNode() != doc {
			t.Error("root element should have document as parent")
		}

		if child1.ParentNode() != root {
			t.Error("child1 should have root as parent")
		}

		if doc.ParentNode() != nil {
			t.Error("document should have no parent")
		}
	})

	t.Run("FirstChild/LastChild", func(t *testing.T) {
		if root.FirstChild() != child1 {
			t.Error("root's first child should be child1")
		}

		if root.LastChild() != child3 {
			t.Error("root's last child should be child3")
		}

		if child1.FirstChild() != nil {
			t.Error("child1 should have no children")
		}
	})

	t.Run("PreviousSibling/NextSibling", func(t *testing.T) {
		if child1.PreviousSibling() != nil {
			t.Error("child1 should have no previous sibling")
		}

		if child2.PreviousSibling() != child1 {
			t.Error("child2's previous sibling should be child1")
		}

		if child2.NextSibling() != child3 {
			t.Error("child2's next sibling should be child3")
		}

		if child3.NextSibling() != nil {
			t.Error("child3 should have no next sibling")
		}
	})

	t.Run("ChildNodes", func(t *testing.T) {
		children := root.ChildNodes()

		if children.Length() != 3 {
			t.Errorf("expected 3 children, got %d", children.Length())
		}

		if children.Item(0) != child1 {
			t.Error("first child should be child1")
		}

		if children.Item(1) != child2 {
			t.Error("second child should be child2")
		}

		if children.Item(2) != child3 {
			t.Error("third child should be child3")
		}
	})
}
