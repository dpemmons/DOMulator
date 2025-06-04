package dom

import (
	"strings"
	"testing"
)

// TestNodeSpecificationCompliance tests exact WHATWG DOM specification compliance
func TestNodeSpecificationCompliance(t *testing.T) {
	t.Run("NodeName Specification Compliance", func(t *testing.T) {
		doc := NewDocument()

		// Test Element - should return HTML-uppercased qualified name
		elem := doc.CreateElement("div")
		if elem.NodeName() != "DIV" {
			t.Errorf("Element NodeName: expected 'DIV', got '%s'", elem.NodeName())
		}

		// Test lowercase element name - should be uppercased
		elem2 := doc.CreateElement("span")
		if elem2.NodeName() != "SPAN" {
			t.Errorf("Element NodeName: expected 'SPAN', got '%s'", elem2.NodeName())
		}

		// Test Text node - should return "#text"
		text := doc.CreateTextNode("hello")
		if text.NodeName() != "#text" {
			t.Errorf("Text NodeName: expected '#text', got '%s'", text.NodeName())
		}

		// Test Comment node - should return "#comment"
		comment := doc.CreateComment("test")
		if comment.NodeName() != "#comment" {
			t.Errorf("Comment NodeName: expected '#comment', got '%s'", comment.NodeName())
		}

		// Test Document - should return "#document"
		if doc.NodeName() != "#document" {
			t.Errorf("Document NodeName: expected '#document', got '%s'", doc.NodeName())
		}

		// Test DocumentFragment - should return "#document-fragment"
		fragment := doc.CreateDocumentFragment()
		if fragment.NodeName() != "#document-fragment" {
			t.Errorf("DocumentFragment NodeName: expected '#document-fragment', got '%s'", fragment.NodeName())
		}

		// Test ProcessingInstruction - should return target
		pi := NewProcessingInstruction("xml-stylesheet", "href='style.css' type='text/css'", doc)
		if pi.NodeName() != "xml-stylesheet" {
			t.Errorf("ProcessingInstruction NodeName: expected 'xml-stylesheet', got '%s'", pi.NodeName())
		}

		// Test CDATASection - should return "#cdata-section"
		cdata := NewCDATASection("some data", doc)
		if cdata.NodeName() != "#cdata-section" {
			t.Errorf("CDATASection NodeName: expected '#cdata-section', got '%s'", cdata.NodeName())
		}

		// Test DocumentType - should return name
		doctype := NewDocumentType("html", "", "", doc)
		if doctype.NodeName() != "html" {
			t.Errorf("DocumentType NodeName: expected 'html', got '%s'", doctype.NodeName())
		}

		// Test Attribute - should return qualified name
		attr := NewAttr("id", "test", doc)
		if attr.NodeName() != "id" {
			t.Errorf("Attribute NodeName: expected 'id', got '%s'", attr.NodeName())
		}
	})

	t.Run("NodeValue Specification Compliance", func(t *testing.T) {
		doc := NewDocument()

		// Test CharacterData nodes - should return data
		text := doc.CreateTextNode("hello world")
		if text.NodeValue() != "hello world" {
			t.Errorf("Text NodeValue: expected 'hello world', got '%s'", text.NodeValue())
		}

		comment := doc.CreateComment("test comment")
		if comment.NodeValue() != "test comment" {
			t.Errorf("Comment NodeValue: expected 'test comment', got '%s'", comment.NodeValue())
		}

		cdata := NewCDATASection("cdata content", doc)
		if cdata.NodeValue() != "cdata content" {
			t.Errorf("CDATASection NodeValue: expected 'cdata content', got '%s'", cdata.NodeValue())
		}

		pi := NewProcessingInstruction("target", "data content", doc)
		if pi.NodeValue() != "data content" {
			t.Errorf("ProcessingInstruction NodeValue: expected 'data content', got '%s'", pi.NodeValue())
		}

		// Test Attribute nodes - should return value
		attr := NewAttr("id", "test-value", doc)
		if attr.NodeValue() != "test-value" {
			t.Errorf("Attribute NodeValue: expected 'test-value', got '%s'", attr.NodeValue())
		}

		// Test other nodes - should return empty string (null)
		elem := doc.CreateElement("div")
		if elem.NodeValue() != "" {
			t.Errorf("Element NodeValue: expected empty string, got '%s'", elem.NodeValue())
		}

		if doc.NodeValue() != "" {
			t.Errorf("Document NodeValue: expected empty string, got '%s'", doc.NodeValue())
		}

		fragment := doc.CreateDocumentFragment()
		if fragment.NodeValue() != "" {
			t.Errorf("DocumentFragment NodeValue: expected empty string, got '%s'", fragment.NodeValue())
		}

		doctype := NewDocumentType("html", "", "", doc)
		if doctype.NodeValue() != "" {
			t.Errorf("DocumentType NodeValue: expected empty string, got '%s'", doctype.NodeValue())
		}
	})

	t.Run("NodeValue Setter Specification Compliance", func(t *testing.T) {
		doc := NewDocument()

		// Test CharacterData nodes - should set data
		text := doc.CreateTextNode("original")
		text.SetNodeValue("modified")
		if text.NodeValue() != "modified" {
			t.Errorf("Text SetNodeValue: expected 'modified', got '%s'", text.NodeValue())
		}

		// Test null handling - should treat as empty string
		text.SetNodeValue("")
		if text.NodeValue() != "" {
			t.Errorf("Text SetNodeValue with empty: expected empty string, got '%s'", text.NodeValue())
		}

		// Test Attribute nodes - should set value
		attr := NewAttr("id", "original", doc)
		attr.SetNodeValue("new-value")
		if attr.NodeValue() != "new-value" {
			t.Errorf("Attribute SetNodeValue: expected 'new-value', got '%s'", attr.NodeValue())
		}

		// Test other nodes - should do nothing
		elem := doc.CreateElement("div")
		elem.SetNodeValue("ignored")
		if elem.NodeValue() != "" {
			t.Errorf("Element SetNodeValue: should remain empty, got '%s'", elem.NodeValue())
		}
	})

	t.Run("TextContent Specification Compliance", func(t *testing.T) {
		doc := NewDocument()

		// Test DocumentFragment and Element - should return descendant text content
		elem := doc.CreateElement("div")
		text1 := doc.CreateTextNode("Hello ")
		text2 := doc.CreateTextNode("World")
		span := doc.CreateElement("span")
		text3 := doc.CreateTextNode("!")

		elem.AppendChild(text1)
		elem.AppendChild(span)
		span.AppendChild(text2)
		elem.AppendChild(text3)

		expected := "Hello World!"
		if elem.TextContent() != expected {
			t.Errorf("Element TextContent: expected '%s', got '%s'", expected, elem.TextContent())
		}

		// Test with nested structure
		fragment := doc.CreateDocumentFragment()
		div1 := doc.CreateElement("div")
		div2 := doc.CreateElement("div")
		text4 := doc.CreateTextNode("Fragment ")
		text5 := doc.CreateTextNode("Text")

		fragment.AppendChild(div1)
		div1.AppendChild(text4)
		fragment.AppendChild(div2)
		div2.AppendChild(text5)

		expectedFragment := "Fragment Text"
		if fragment.TextContent() != expectedFragment {
			t.Errorf("DocumentFragment TextContent: expected '%s', got '%s'", expectedFragment, fragment.TextContent())
		}

		// Test Attribute - should return value
		attr := NewAttr("title", "tooltip text", doc)
		if attr.TextContent() != "tooltip text" {
			t.Errorf("Attribute TextContent: expected 'tooltip text', got '%s'", attr.TextContent())
		}

		// Test CharacterData - should return data
		text := doc.CreateTextNode("text data")
		if text.TextContent() != "text data" {
			t.Errorf("Text TextContent: expected 'text data', got '%s'", text.TextContent())
		}

		comment := doc.CreateComment("comment data")
		if comment.TextContent() != "comment data" {
			t.Errorf("Comment TextContent: expected 'comment data', got '%s'", comment.TextContent())
		}

		// Test other nodes - should return empty string (null)
		if doc.TextContent() != "" {
			t.Errorf("Document TextContent: expected empty string, got '%s'", doc.TextContent())
		}

		doctype := NewDocumentType("html", "", "", doc)
		if doctype.TextContent() != "" {
			t.Errorf("DocumentType TextContent: expected empty string, got '%s'", doctype.TextContent())
		}
	})

	t.Run("TextContent Setter Specification Compliance", func(t *testing.T) {
		doc := NewDocument()

		// Test Element and DocumentFragment - should string replace all
		elem := doc.CreateElement("div")
		elem.AppendChild(doc.CreateTextNode("old"))
		elem.AppendChild(doc.CreateElement("span"))
		elem.AppendChild(doc.CreateTextNode("content"))

		elem.SetTextContent("new content")
		if elem.TextContent() != "new content" {
			t.Errorf("Element SetTextContent: expected 'new content', got '%s'", elem.TextContent())
		}

		// Should have only one Text child
		if elem.ChildNodes().Length() != 1 {
			t.Errorf("Element SetTextContent: expected 1 child, got %d", elem.ChildNodes().Length())
		}

		if elem.FirstChild().NodeType() != TextNode {
			t.Error("Element SetTextContent: child should be Text node")
		}

		// Test setting empty string - should remove all children
		elem.SetTextContent("")
		if elem.ChildNodes().Length() != 0 {
			t.Errorf("Element SetTextContent empty: expected 0 children, got %d", elem.ChildNodes().Length())
		}

		// Test Attribute - should set value
		attr := NewAttr("class", "old-class", doc)
		attr.SetTextContent("new-class")
		if attr.TextContent() != "new-class" {
			t.Errorf("Attribute SetTextContent: expected 'new-class', got '%s'", attr.TextContent())
		}

		// Test CharacterData - should set data
		text := doc.CreateTextNode("old data")
		text.SetTextContent("new data")
		if text.TextContent() != "new data" {
			t.Errorf("Text SetTextContent: expected 'new data', got '%s'", text.TextContent())
		}

		// Test other nodes - should do nothing
		doc.SetTextContent("ignored")
		if doc.TextContent() != "" {
			t.Errorf("Document SetTextContent: should remain empty, got '%s'", doc.TextContent())
		}
	})

	t.Run("Normalize Method Specification Compliance", func(t *testing.T) {
		doc := NewDocument()
		container := doc.CreateElement("div")

		// Test removing empty Text nodes
		text1 := doc.CreateTextNode("hello")
		emptyText := doc.CreateTextNode("")
		text2 := doc.CreateTextNode("world")

		container.AppendChild(text1)
		container.AppendChild(emptyText)
		container.AppendChild(text2)

		// Before normalize: 3 children
		if container.ChildNodes().Length() != 3 {
			t.Errorf("Before normalize: expected 3 children, got %d", container.ChildNodes().Length())
		}

		container.Normalize()

		// After normalize: should be 1 child (empty text removed, others merged)
		if container.ChildNodes().Length() != 1 {
			t.Errorf("After normalize: expected 1 child, got %d", container.ChildNodes().Length())
		}

		if container.FirstChild().NodeValue() != "helloworld" {
			t.Errorf("After normalize: expected 'helloworld', got '%s'", container.FirstChild().NodeValue())
		}

		// Test with nested structure
		nested := doc.CreateElement("span")
		nestedText1 := doc.CreateTextNode("nested ")
		nestedText2 := doc.CreateTextNode("text")
		nested.AppendChild(nestedText1)
		nested.AppendChild(nestedText2)

		container.AppendChild(nested)
		container.Normalize()

		// Normalize should affect descendant Text nodes
		if nested.ChildNodes().Length() != 1 {
			t.Errorf("Nested after normalize: expected 1 child, got %d", nested.ChildNodes().Length())
		}

		if nested.FirstChild().NodeValue() != "nested text" {
			t.Errorf("Nested after normalize: expected 'nested text', got '%s'", nested.FirstChild().NodeValue())
		}
	})

	t.Run("IsEqualNode Deep Comparison Specification", func(t *testing.T) {
		doc := NewDocument()

		// Test Element equality with attributes
		elem1 := doc.CreateElement("div")
		elem1.SetAttribute("id", "test")
		elem1.SetAttribute("class", "example")

		elem2 := doc.CreateElement("div")
		elem2.SetAttribute("id", "test")
		elem2.SetAttribute("class", "example")

		if !elem1.IsEqualNode(elem2) {
			t.Error("Elements with same attributes should be equal")
		}

		// Test with different attribute values
		elem2.SetAttribute("id", "different")
		if elem1.IsEqualNode(elem2) {
			t.Error("Elements with different attribute values should not be equal")
		}

		// Test with different number of attributes
		elem2.SetAttribute("id", "test")
		elem2.SetAttribute("data-extra", "value")
		if elem1.IsEqualNode(elem2) {
			t.Error("Elements with different number of attributes should not be equal")
		}

		// Test ProcessingInstruction equality
		pi1 := NewProcessingInstruction("xml-stylesheet", "href='a.css'", doc)
		pi2 := NewProcessingInstruction("xml-stylesheet", "href='a.css'", doc)
		pi3 := NewProcessingInstruction("xml-stylesheet", "href='b.css'", doc)
		pi4 := NewProcessingInstruction("other-target", "href='a.css'", doc)

		if !pi1.IsEqualNode(pi2) {
			t.Error("ProcessingInstructions with same target and data should be equal")
		}

		if pi1.IsEqualNode(pi3) {
			t.Error("ProcessingInstructions with different data should not be equal")
		}

		if pi1.IsEqualNode(pi4) {
			t.Error("ProcessingInstructions with different target should not be equal")
		}

		// Test DocumentType equality
		dt1 := NewDocumentType("html", "public1", "system1", doc)
		dt2 := NewDocumentType("html", "public1", "system1", doc)
		dt3 := NewDocumentType("html", "public2", "system1", doc)

		if !dt1.IsEqualNode(dt2) {
			t.Error("DocumentTypes with same name, publicId, systemId should be equal")
		}

		if dt1.IsEqualNode(dt3) {
			t.Error("DocumentTypes with different publicId should not be equal")
		}

		// Test with children
		parent1 := doc.CreateElement("parent")
		parent1.AppendChild(doc.CreateTextNode("child text"))

		parent2 := doc.CreateElement("parent")
		parent2.AppendChild(doc.CreateTextNode("child text"))

		if !parent1.IsEqualNode(parent2) {
			t.Error("Elements with equal children should be equal")
		}

		parent2.AppendChild(doc.CreateTextNode(" extra"))
		if parent1.IsEqualNode(parent2) {
			t.Error("Elements with different children should not be equal")
		}
	})

	t.Run("CompareDocumentPosition Bitmask Specification", func(t *testing.T) {
		doc := NewDocument()
		root := doc.CreateElement("root")
		child1 := doc.CreateElement("child1")
		child2 := doc.CreateElement("child2")
		grandchild := doc.CreateElement("grandchild")

		doc.AppendChild(root)
		root.AppendChild(child1)
		root.AppendChild(child2)
		child1.AppendChild(grandchild)

		// Test same node
		if root.CompareDocumentPosition(root) != 0 {
			t.Error("CompareDocumentPosition with same node should return 0")
		}

		// Test contains relationship
		position := root.CompareDocumentPosition(grandchild)
		if (position & DOCUMENT_POSITION_CONTAINED_BY) == 0 {
			t.Error("Ancestor should have CONTAINED_BY bit set for descendant")
		}
		if (position & DOCUMENT_POSITION_FOLLOWING) == 0 {
			t.Error("Ancestor should have FOLLOWING bit set for descendant")
		}

		// Test contained by relationship
		position = grandchild.CompareDocumentPosition(root)
		if (position & DOCUMENT_POSITION_CONTAINS) == 0 {
			t.Error("Descendant should have CONTAINS bit set for ancestor")
		}
		if (position & DOCUMENT_POSITION_PRECEDING) == 0 {
			t.Error("Descendant should have PRECEDING bit set for ancestor")
		}

		// Test sibling relationships
		position = child1.CompareDocumentPosition(child2)
		if (position & DOCUMENT_POSITION_FOLLOWING) == 0 {
			t.Error("First sibling should have FOLLOWING bit set for second sibling")
		}

		position = child2.CompareDocumentPosition(child1)
		if (position & DOCUMENT_POSITION_PRECEDING) == 0 {
			t.Error("Second sibling should have PRECEDING bit set for first sibling")
		}

		// Test disconnected nodes
		orphan := doc.CreateElement("orphan")
		position = root.CompareDocumentPosition(orphan)
		if (position & DOCUMENT_POSITION_DISCONNECTED) == 0 {
			t.Error("Disconnected nodes should have DISCONNECTED bit set")
		}
		if (position & DOCUMENT_POSITION_IMPLEMENTATION_SPECIFIC) == 0 {
			t.Error("Disconnected nodes should have IMPLEMENTATION_SPECIFIC bit set")
		}
	})

	t.Run("BaseURI Specification Compliance", func(t *testing.T) {
		doc := NewDocument()
		elem := doc.CreateElement("div")

		// BaseURI should return document's base URL serialized
		baseURI := elem.BaseURI()
		if baseURI == "" {
			t.Error("BaseURI should return non-empty string")
		}

		// For now, should return "about:blank" as default
		if baseURI != "about:blank" {
			t.Errorf("BaseURI: expected 'about:blank', got '%s'", baseURI)
		}

		// Document itself should also return base URI
		docBaseURI := doc.BaseURI()
		if docBaseURI != baseURI {
			t.Error("Document and element should have same base URI")
		}
	})

	t.Run("OwnerDocument Specification Compliance", func(t *testing.T) {
		doc := NewDocument()

		// Document itself should return null
		if doc.OwnerDocument() != nil {
			t.Error("Document.OwnerDocument() should return nil")
		}

		// All other nodes should return their document
		testNodes := []Node{
			doc.CreateElement("div"),
			doc.CreateTextNode("text"),
			doc.CreateComment("comment"),
			doc.CreateDocumentFragment(),
			NewProcessingInstruction("target", "data", doc),
			NewCDATASection("data", doc),
			NewDocumentType("html", "", "", doc),
			NewAttr("id", "value", doc),
		}

		for _, node := range testNodes {
			if node.OwnerDocument() != doc {
				t.Errorf("Node %s should have document as owner", node.NodeName())
			}
		}
	})

	t.Run("Contains Method Specification Compliance", func(t *testing.T) {
		doc := NewDocument()
		elem := doc.CreateElement("div")
		text := doc.CreateTextNode("content")
		elem.AppendChild(text)

		// Self-containment (inclusive descendant)
		if !elem.Contains(elem) {
			t.Error("Node should contain itself")
		}

		// Ancestor contains descendant
		if !elem.Contains(text) {
			t.Error("Element should contain its text node")
		}

		// Descendant does not contain ancestor
		if text.Contains(elem) {
			t.Error("Text node should not contain its parent element")
		}

		// Null parameter
		if elem.Contains(nil) {
			t.Error("Contains(nil) should return false")
		}

		// Unrelated nodes
		other := doc.CreateElement("other")
		if elem.Contains(other) {
			t.Error("Element should not contain unrelated element")
		}
	})

	t.Run("GetRootNode Method Specification Compliance", func(t *testing.T) {
		doc := NewDocument()
		elem := doc.CreateElement("div")
		text := doc.CreateTextNode("content")

		doc.AppendChild(elem)
		elem.AppendChild(text)

		// Default options (nil should work)
		root := text.GetRootNode(nil)
		if root != doc {
			t.Error("GetRootNode(nil) should return document")
		}

		// Explicit composed: false
		options := &GetRootNodeOptions{Composed: false}
		root = text.GetRootNode(options)
		if root != doc {
			t.Error("GetRootNode({composed: false}) should return document")
		}

		// Composed: true (same result for now, no shadow DOM)
		options.Composed = true
		root = text.GetRootNode(options)
		if root != doc {
			t.Error("GetRootNode({composed: true}) should return document")
		}

		// Orphaned node should return itself as root
		orphan := doc.CreateElement("orphan")
		root = orphan.GetRootNode(nil)
		if root != orphan {
			t.Error("Orphaned node should return itself as root")
		}
	})
}

// TestNodeSpecificationEdgeCases tests edge cases from the WHATWG DOM specification
func TestNodeSpecificationEdgeCases(t *testing.T) {
	t.Run("NodeName Case Sensitivity", func(t *testing.T) {
		doc := NewDocument()

		// Element names should be HTML-uppercased
		testCases := []string{"div", "DIV", "Div", "dIv"}
		for _, name := range testCases {
			elem := doc.CreateElement(name)
			if elem.NodeName() != strings.ToUpper(name) {
				t.Errorf("Element created with '%s' should have NodeName '%s', got '%s'",
					name, strings.ToUpper(name), elem.NodeName())
			}
		}
	})

	t.Run("Empty Text Node Normalization", func(t *testing.T) {
		doc := NewDocument()
		container := doc.CreateElement("div")

		// Add only empty text nodes
		container.AppendChild(doc.CreateTextNode(""))
		container.AppendChild(doc.CreateTextNode(""))

		if container.ChildNodes().Length() != 2 {
			t.Error("Should have 2 empty text nodes before normalize")
		}

		container.Normalize()

		// All empty text nodes should be removed
		if container.ChildNodes().Length() != 0 {
			t.Errorf("After normalize: expected 0 children, got %d", container.ChildNodes().Length())
		}
	})

	t.Run("Mixed Empty and Non-Empty Text Nodes", func(t *testing.T) {
		doc := NewDocument()
		container := doc.CreateElement("div")

		// Mix of empty and non-empty text nodes
		container.AppendChild(doc.CreateTextNode(""))
		container.AppendChild(doc.CreateTextNode("hello"))
		container.AppendChild(doc.CreateTextNode(""))
		container.AppendChild(doc.CreateTextNode("world"))
		container.AppendChild(doc.CreateTextNode(""))

		container.Normalize()

		// Should result in one text node with concatenated content
		if container.ChildNodes().Length() != 1 {
			t.Errorf("After normalize: expected 1 child, got %d", container.ChildNodes().Length())
		}

		if container.FirstChild().NodeValue() != "helloworld" {
			t.Errorf("After normalize: expected 'helloworld', got '%s'", container.FirstChild().NodeValue())
		}
	})

	t.Run("CloneNode Type Preservation", func(t *testing.T) {
		doc := NewDocument()

		// Test that cloned nodes preserve their specific types
		originalElement := doc.CreateElement("div")
		originalElement.SetAttribute("id", "test")

		clonedElement := originalElement.CloneNode(false)

		// Should be an Element, not just a Node
		if clonedElement.NodeType() != ElementNode {
			t.Error("Cloned element should have ElementNode type")
		}

		// Should have attributes (shallow clone of element includes attributes)
		if clonedElem, ok := clonedElement.(*Element); ok {
			if clonedElem.GetAttribute("id") != "test" {
				t.Error("Cloned element should have copied attributes")
			}
		} else {
			t.Error("Cloned node should be castable to *Element")
		}

		// Deep clone should include children
		text := doc.CreateTextNode("content")
		originalElement.AppendChild(text)

		deepClone := originalElement.CloneNode(true)
		if deepClone.ChildNodes().Length() != 1 {
			t.Error("Deep clone should include children")
		}

		if deepClone.FirstChild().NodeValue() != "content" {
			t.Error("Deep clone should clone child content")
		}
	})

	t.Run("IsEqualNode with Complex Hierarchy", func(t *testing.T) {
		doc := NewDocument()

		// Create complex hierarchy
		createComplexElement := func() *Element {
			elem := doc.CreateElement("article")
			elem.SetAttribute("class", "post")

			header := doc.CreateElement("header")
			header.AppendChild(doc.CreateTextNode("Title"))
			elem.AppendChild(header)

			content := doc.CreateElement("div")
			content.SetAttribute("class", "content")
			content.AppendChild(doc.CreateTextNode("Content text"))
			elem.AppendChild(content)

			return elem
		}

		elem1 := createComplexElement()
		elem2 := createComplexElement()

		if !elem1.IsEqualNode(elem2) {
			t.Error("Identical complex hierarchies should be equal")
		}

		// Modify one slightly
		elem2.FirstChild().AppendChild(doc.CreateTextNode(" Extra"))

		if elem1.IsEqualNode(elem2) {
			t.Error("Modified hierarchy should not be equal")
		}
	})
}
