package dom

import (
	"testing"
)

// TestNodeInterfaceSpecCompliance tests our Node implementation against the full WHATWG DOM specification
func TestNodeInterfaceSpecCompliance(t *testing.T) {
	t.Run("Node Constants", func(t *testing.T) {
		// Test all node type constants as defined in the spec
		nodeTypes := map[string]NodeType{
			"ELEMENT_NODE":                ElementNode,
			"ATTRIBUTE_NODE":              AttributeNode,
			"TEXT_NODE":                   TextNode,
			"CDATA_SECTION_NODE":          CDataSectionNode,
			"ENTITY_REFERENCE_NODE":       EntityReferenceNode,
			"ENTITY_NODE":                 EntityNode,
			"PROCESSING_INSTRUCTION_NODE": ProcessingInstructionNode,
			"COMMENT_NODE":                CommentNode,
			"DOCUMENT_NODE":               DocumentNode,
			"DOCUMENT_TYPE_NODE":          DocumentTypeNode,
			"DOCUMENT_FRAGMENT_NODE":      DocumentFragmentNode,
			"NOTATION_NODE":               NotationNode,
		}

		expectedValues := map[string]int{
			"ELEMENT_NODE":                1,
			"ATTRIBUTE_NODE":              2,
			"TEXT_NODE":                   3,
			"CDATA_SECTION_NODE":          4,
			"ENTITY_REFERENCE_NODE":       5,
			"ENTITY_NODE":                 6,
			"PROCESSING_INSTRUCTION_NODE": 7,
			"COMMENT_NODE":                8,
			"DOCUMENT_NODE":               9,
			"DOCUMENT_TYPE_NODE":          10,
			"DOCUMENT_FRAGMENT_NODE":      11,
			"NOTATION_NODE":               12,
		}

		for name, nodeType := range nodeTypes {
			if int(nodeType) != expectedValues[name] {
				t.Errorf("Node type %s: expected %d, got %d", name, expectedValues[name], int(nodeType))
			}
		}
	})

	t.Run("Document Position Constants", func(t *testing.T) {
		// Test document position constants
		constants := map[string]int{
			"DOCUMENT_POSITION_DISCONNECTED":            DOCUMENT_POSITION_DISCONNECTED,
			"DOCUMENT_POSITION_PRECEDING":               DOCUMENT_POSITION_PRECEDING,
			"DOCUMENT_POSITION_FOLLOWING":               DOCUMENT_POSITION_FOLLOWING,
			"DOCUMENT_POSITION_CONTAINS":                DOCUMENT_POSITION_CONTAINS,
			"DOCUMENT_POSITION_CONTAINED_BY":            DOCUMENT_POSITION_CONTAINED_BY,
			"DOCUMENT_POSITION_IMPLEMENTATION_SPECIFIC": DOCUMENT_POSITION_IMPLEMENTATION_SPECIFIC,
		}

		expectedValues := map[string]int{
			"DOCUMENT_POSITION_DISCONNECTED":            0x01,
			"DOCUMENT_POSITION_PRECEDING":               0x02,
			"DOCUMENT_POSITION_FOLLOWING":               0x04,
			"DOCUMENT_POSITION_CONTAINS":                0x08,
			"DOCUMENT_POSITION_CONTAINED_BY":            0x10,
			"DOCUMENT_POSITION_IMPLEMENTATION_SPECIFIC": 0x20,
		}

		for name, value := range constants {
			if value != expectedValues[name] {
				t.Errorf("Constant %s: expected %d, got %d", name, expectedValues[name], value)
			}
		}
	})

	t.Run("NodeType and NodeName Getters", func(t *testing.T) {
		doc := NewDocument()

		// Test Element
		elem := doc.CreateElement("div")
		if elem.NodeType() != ElementNode {
			t.Errorf("Element.NodeType(): expected %d, got %d", ElementNode, elem.NodeType())
		}
		if elem.NodeName() != "DIV" {
			t.Errorf("Element.NodeName(): expected 'DIV', got '%s'", elem.NodeName())
		}

		// Test Text
		text := doc.CreateTextNode("hello")
		if text.NodeType() != TextNode {
			t.Errorf("Text.NodeType(): expected %d, got %d", TextNode, text.NodeType())
		}
		if text.NodeName() != "#text" {
			t.Errorf("Text.NodeName(): expected '#text', got '%s'", text.NodeName())
		}

		// Test Comment
		comment := doc.CreateComment("test comment")
		if comment.NodeType() != CommentNode {
			t.Errorf("Comment.NodeType(): expected %d, got %d", CommentNode, comment.NodeType())
		}
		if comment.NodeName() != "#comment" {
			t.Errorf("Comment.NodeName(): expected '#comment', got '%s'", comment.NodeName())
		}

		// Test Document
		if doc.NodeType() != DocumentNode {
			t.Errorf("Document.NodeType(): expected %d, got %d", DocumentNode, doc.NodeType())
		}
		if doc.NodeName() != "#document" {
			t.Errorf("Document.NodeName(): expected '#document', got '%s'", doc.NodeName())
		}

		// Test DocumentFragment
		frag := doc.CreateDocumentFragment()
		if frag.NodeType() != DocumentFragmentNode {
			t.Errorf("DocumentFragment.NodeType(): expected %d, got %d", DocumentFragmentNode, frag.NodeType())
		}
		if frag.NodeName() != "#document-fragment" {
			t.Errorf("DocumentFragment.NodeName(): expected '#document-fragment', got '%s'", frag.NodeName())
		}
	})

	t.Run("BaseURI Getter", func(t *testing.T) {
		doc := NewDocument()
		elem := doc.CreateElement("div")

		// BaseURI should return the document's base URL
		baseURI := elem.BaseURI()
		if baseURI == "" {
			t.Error("BaseURI should return a non-empty string")
		}
	})

	t.Run("IsConnected Getter", func(t *testing.T) {
		doc := NewDocument()
		elem := doc.CreateElement("div")
		text := doc.CreateTextNode("hello")

		// Not connected initially
		if elem.IsConnected() {
			t.Error("Element should not be connected when not in document")
		}
		if text.IsConnected() {
			t.Error("Text should not be connected when not in document")
		}

		// Connected when appended to document
		doc.AppendChild(elem)
		if !elem.IsConnected() {
			t.Error("Element should be connected when in document")
		}

		elem.AppendChild(text)
		if !text.IsConnected() {
			t.Error("Text should be connected when child of connected element")
		}
	})

	t.Run("OwnerDocument Getter", func(t *testing.T) {
		doc := NewDocument()
		elem := doc.CreateElement("div")
		text := doc.CreateTextNode("hello")

		// Document itself returns null
		if doc.OwnerDocument() != nil {
			t.Error("Document.OwnerDocument() should return nil")
		}

		// Other nodes return their document
		if elem.OwnerDocument() != doc {
			t.Error("Element.OwnerDocument() should return the document")
		}
		if text.OwnerDocument() != doc {
			t.Error("Text.OwnerDocument() should return the document")
		}
	})

	t.Run("GetRootNode Method", func(t *testing.T) {
		doc := NewDocument()
		elem := doc.CreateElement("div")
		text := doc.CreateTextNode("hello")

		doc.AppendChild(elem)
		elem.AppendChild(text)

		// Default options (composed: false)
		if text.GetRootNode(nil) != doc {
			t.Error("GetRootNode() should return document root")
		}

		// Explicit options
		options := &GetRootNodeOptions{Composed: false}
		if text.GetRootNode(options) != doc {
			t.Error("GetRootNode() with composed:false should return document root")
		}

		// Composed true (we don't support shadow DOM yet, so same result)
		options.Composed = true
		if text.GetRootNode(options) != doc {
			t.Error("GetRootNode() with composed:true should return document root")
		}
	})

	t.Run("ParentElement Getter", func(t *testing.T) {
		doc := NewDocument()
		elem := doc.CreateElement("div")
		text := doc.CreateTextNode("hello")

		// No parent initially
		if elem.ParentElement() != nil {
			t.Error("Element.ParentElement() should return nil when no parent")
		}

		// Parent is document (not Element)
		doc.AppendChild(elem)
		if elem.ParentElement() != nil {
			t.Error("Element.ParentElement() should return nil when parent is Document")
		}

		// Parent is element
		elem.AppendChild(text)
		if text.ParentElement() != elem {
			t.Error("Text.ParentElement() should return parent element")
		}
	})

	t.Run("HasChildNodes Method", func(t *testing.T) {
		doc := NewDocument()
		elem := doc.CreateElement("div")
		text := doc.CreateTextNode("hello")

		// No children initially
		if elem.HasChildNodes() {
			t.Error("Element.HasChildNodes() should return false when no children")
		}

		// Has children after appending
		elem.AppendChild(text)
		if !elem.HasChildNodes() {
			t.Error("Element.HasChildNodes() should return true when has children")
		}
	})

	t.Run("NodeValue Getter and Setter", func(t *testing.T) {
		doc := NewDocument()
		text := doc.CreateTextNode("hello")
		comment := doc.CreateComment("test")
		elem := doc.CreateElement("div")

		// CharacterData nodes return their data
		if text.NodeValue() != "hello" {
			t.Errorf("Text.NodeValue(): expected 'hello', got '%s'", text.NodeValue())
		}
		if comment.NodeValue() != "test" {
			t.Errorf("Comment.NodeValue(): expected 'test', got '%s'", comment.NodeValue())
		}

		// Element returns null (empty string)
		if elem.NodeValue() != "" {
			t.Errorf("Element.NodeValue(): expected empty string, got '%s'", elem.NodeValue())
		}

		// Setting nodeValue on CharacterData
		text.SetNodeValue("world")
		if text.NodeValue() != "world" {
			t.Errorf("After SetNodeValue('world'), expected 'world', got '%s'", text.NodeValue())
		}

		// Setting nodeValue on Element does nothing
		elem.SetNodeValue("ignored")
		if elem.NodeValue() != "" {
			t.Errorf("Element.NodeValue() should remain empty after SetNodeValue")
		}
	})

	t.Run("TextContent Getter and Setter", func(t *testing.T) {
		doc := NewDocument()
		elem := doc.CreateElement("div")
		text1 := doc.CreateTextNode("hello")
		text2 := doc.CreateTextNode(" world")

		elem.AppendChild(text1)
		elem.AppendChild(text2)

		// TextContent concatenates descendant text
		if elem.TextContent() != "hello world" {
			t.Errorf("Element.TextContent(): expected 'hello world', got '%s'", elem.TextContent())
		}

		// Setting textContent replaces all children
		elem.SetTextContent("new content")
		if elem.TextContent() != "new content" {
			t.Errorf("After SetTextContent(), expected 'new content', got '%s'", elem.TextContent())
		}

		// Should have only one Text child now
		if elem.ChildNodes().Length() != 1 {
			t.Errorf("After SetTextContent(), expected 1 child, got %d", elem.ChildNodes().Length())
		}
	})

	t.Run("Normalize Method", func(t *testing.T) {
		doc := NewDocument()
		elem := doc.CreateElement("div")
		text1 := doc.CreateTextNode("hello")
		text2 := doc.CreateTextNode(" ")
		text3 := doc.CreateTextNode("world")

		elem.AppendChild(text1)
		elem.AppendChild(text2)
		elem.AppendChild(text3)

		// Before normalize: 3 text nodes
		if elem.ChildNodes().Length() != 3 {
			t.Errorf("Before normalize: expected 3 children, got %d", elem.ChildNodes().Length())
		}

		elem.Normalize()

		// After normalize: should be 1 text node
		if elem.ChildNodes().Length() != 1 {
			t.Errorf("After normalize: expected 1 child, got %d", elem.ChildNodes().Length())
		}

		// Content should be concatenated
		if elem.FirstChild().NodeValue() != "hello world" {
			t.Errorf("After normalize: expected 'hello world', got '%s'", elem.FirstChild().NodeValue())
		}
	})

	t.Run("CloneNode Method", func(t *testing.T) {
		doc := NewDocument()
		elem := doc.CreateElement("div")
		text := doc.CreateTextNode("hello")
		elem.AppendChild(text)

		// Shallow clone
		shallowClone := elem.CloneNode(false)
		if shallowClone.NodeType() != ElementNode {
			t.Error("Cloned node should have same type")
		}
		if shallowClone.NodeName() != "DIV" {
			t.Error("Cloned node should have same name")
		}
		if shallowClone.ChildNodes().Length() != 0 {
			t.Error("Shallow clone should have no children")
		}

		// Deep clone
		deepClone := elem.CloneNode(true)
		if deepClone.ChildNodes().Length() != 1 {
			t.Error("Deep clone should have children")
		}
		if deepClone.FirstChild().NodeValue() != "hello" {
			t.Error("Deep clone children should be cloned")
		}

		// Clone should not be same node
		if elem.IsSameNode(shallowClone) {
			t.Error("Clone should not be same node as original")
		}
	})

	t.Run("IsEqualNode Method", func(t *testing.T) {
		doc := NewDocument()
		elem1 := doc.CreateElement("div")
		elem2 := doc.CreateElement("div")
		elem3 := doc.CreateElement("span")

		// Same type and name
		if !elem1.IsEqualNode(elem2) {
			t.Error("Elements with same type and name should be equal")
		}

		// Different names
		if elem1.IsEqualNode(elem3) {
			t.Error("Elements with different names should not be equal")
		}

		// With children
		text1 := doc.CreateTextNode("hello")
		text2 := doc.CreateTextNode("hello")
		elem1.AppendChild(text1)
		elem2.AppendChild(text2)

		if !elem1.IsEqualNode(elem2) {
			t.Error("Elements with equal children should be equal")
		}

		// Null comparison
		if elem1.IsEqualNode(nil) {
			t.Error("Node should not be equal to nil")
		}
	})

	t.Run("IsSameNode Method", func(t *testing.T) {
		doc := NewDocument()
		elem1 := doc.CreateElement("div")
		elem2 := doc.CreateElement("div")

		// Same reference
		if !elem1.IsSameNode(elem1) {
			t.Error("Node should be same as itself")
		}

		// Different references
		if elem1.IsSameNode(elem2) {
			t.Error("Different nodes should not be same")
		}

		// Null comparison
		if elem1.IsSameNode(nil) {
			t.Error("Node should not be same as nil")
		}
	})

	t.Run("CompareDocumentPosition Method", func(t *testing.T) {
		doc := NewDocument()
		elem1 := doc.CreateElement("div")
		elem2 := doc.CreateElement("span")
		text := doc.CreateTextNode("hello")

		doc.AppendChild(elem1)
		elem1.AppendChild(elem2)
		elem2.AppendChild(text)

		// Same node
		if elem1.CompareDocumentPosition(elem1) != 0 {
			t.Error("CompareDocumentPosition with same node should return 0")
		}

		// Contains relationship
		position := elem1.CompareDocumentPosition(text)
		if (position & DOCUMENT_POSITION_CONTAINED_BY) == 0 {
			t.Error("Ancestor should contain descendant")
		}
		if (position & DOCUMENT_POSITION_FOLLOWING) == 0 {
			t.Error("Ancestor should precede descendant")
		}

		// Contained by relationship
		position = text.CompareDocumentPosition(elem1)
		if (position & DOCUMENT_POSITION_CONTAINS) == 0 {
			t.Error("Descendant should be contained by ancestor")
		}
		if (position & DOCUMENT_POSITION_PRECEDING) == 0 {
			t.Error("Descendant should follow ancestor")
		}
	})

	t.Run("Contains Method", func(t *testing.T) {
		doc := NewDocument()
		elem1 := doc.CreateElement("div")
		elem2 := doc.CreateElement("span")
		text := doc.CreateTextNode("hello")

		doc.AppendChild(elem1)
		elem1.AppendChild(elem2)
		elem2.AppendChild(text)

		// Self-containment
		if !elem1.Contains(elem1) {
			t.Error("Node should contain itself")
		}

		// Ancestor contains descendant
		if !elem1.Contains(text) {
			t.Error("Ancestor should contain descendant")
		}

		// Descendant does not contain ancestor
		if text.Contains(elem1) {
			t.Error("Descendant should not contain ancestor")
		}

		// Null containment
		if elem1.Contains(nil) {
			t.Error("Node should not contain nil")
		}
	})

	t.Run("Tree Navigation Methods", func(t *testing.T) {
		doc := NewDocument()
		elem := doc.CreateElement("div")
		text1 := doc.CreateTextNode("first")
		text2 := doc.CreateTextNode("second")

		elem.AppendChild(text1)
		elem.AppendChild(text2)

		// ParentNode
		if text1.ParentNode() != elem {
			t.Error("ParentNode should return parent")
		}

		// FirstChild and LastChild
		if elem.FirstChild() != text1 {
			t.Error("FirstChild should return first child")
		}
		if elem.LastChild() != text2 {
			t.Error("LastChild should return last child")
		}

		// PreviousSibling and NextSibling
		if text1.PreviousSibling() != nil {
			t.Error("First child should have no previous sibling")
		}
		if text1.NextSibling() != text2 {
			t.Error("First child's next sibling should be second child")
		}
		if text2.PreviousSibling() != text1 {
			t.Error("Second child's previous sibling should be first child")
		}
		if text2.NextSibling() != nil {
			t.Error("Last child should have no next sibling")
		}
	})

	t.Run("Child Manipulation Methods", func(t *testing.T) {
		doc := NewDocument()
		elem := doc.CreateElement("div")
		text1 := doc.CreateTextNode("first")
		text2 := doc.CreateTextNode("second")
		text3 := doc.CreateTextNode("third")

		// AppendChild
		elem.AppendChild(text1)
		if elem.FirstChild() != text1 {
			t.Error("AppendChild should add child")
		}

		// InsertBefore
		elem.InsertBefore(text2, text1)
		if elem.FirstChild() != text2 {
			t.Error("InsertBefore should insert at correct position")
		}

		// ReplaceChild
		elem.ReplaceChild(text3, text1)
		if elem.LastChild() != text3 {
			t.Error("ReplaceChild should replace child")
		}

		// RemoveChild
		elem.RemoveChild(text3)
		if elem.ChildNodes().Length() != 1 {
			t.Error("RemoveChild should remove child")
		}
	})
}

// TestNodeInterfaceNamespaceSupport tests namespace-related functionality
func TestNodeInterfaceNamespaceSupport(t *testing.T) {
	// Note: Full namespace support is not yet implemented
	// This test documents the expected behavior per the spec

	t.Run("LookupPrefix Method", func(t *testing.T) {
		doc := NewDocument()
		elem := doc.CreateElement("div")

		// LookupPrefix should be implemented (currently returns empty)
		prefix := elem.LookupPrefix("http://www.w3.org/1999/xhtml")
		_ = prefix // Currently not fully implemented
	})

	t.Run("LookupNamespaceURI Method", func(t *testing.T) {
		doc := NewDocument()
		elem := doc.CreateElement("div")

		// LookupNamespaceURI should be implemented (currently returns empty)
		namespace := elem.LookupNamespaceURI("xmlns")
		_ = namespace // Currently not fully implemented
	})

	t.Run("IsDefaultNamespace Method", func(t *testing.T) {
		doc := NewDocument()
		elem := doc.CreateElement("div")

		// IsDefaultNamespace should be implemented (currently returns false)
		isDefault := elem.IsDefaultNamespace("http://www.w3.org/1999/xhtml")
		_ = isDefault // Currently not fully implemented
	})
}
