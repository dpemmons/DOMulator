package dom

import (
	"testing"
)

func TestIsEqualNode_Null(t *testing.T) {
	doc := NewDocument()
	elem := doc.CreateElement("div")

	// Per spec: Returns false if otherNode is null
	if elem.IsEqualNode(nil) {
		t.Error("IsEqualNode should return false for nil input")
	}
}

func TestIsEqualNode_SameNode(t *testing.T) {
	doc := NewDocument()
	elem := doc.CreateElement("div")

	// Same node should always equal itself
	if !elem.IsEqualNode(elem) {
		t.Error("IsEqualNode should return true for same node")
	}
}

func TestIsEqualNode_DifferentNodeTypes(t *testing.T) {
	doc := NewDocument()
	elem := doc.CreateElement("div")
	text := doc.CreateTextNode("text")

	// Different node types should not be equal
	if elem.IsEqualNode(text) {
		t.Error("IsEqualNode should return false for different node types")
	}
}

func TestIsEqualNode_Text(t *testing.T) {
	doc := NewDocument()

	// Equal text nodes
	text1 := doc.CreateTextNode("hello")
	text2 := doc.CreateTextNode("hello")
	if !text1.IsEqualNode(text2) {
		t.Error("Text nodes with same data should be equal")
	}

	// Different text nodes
	text3 := doc.CreateTextNode("world")
	if text1.IsEqualNode(text3) {
		t.Error("Text nodes with different data should not be equal")
	}

	// Empty text nodes
	text4 := doc.CreateTextNode("")
	text5 := doc.CreateTextNode("")
	if !text4.IsEqualNode(text5) {
		t.Error("Empty text nodes should be equal")
	}
}

func TestIsEqualNode_Comment(t *testing.T) {
	doc := NewDocument()

	// Equal comment nodes
	comment1 := doc.CreateComment("comment")
	comment2 := doc.CreateComment("comment")
	if !comment1.IsEqualNode(comment2) {
		t.Error("Comment nodes with same data should be equal")
	}

	// Different comment nodes
	comment3 := doc.CreateComment("different")
	if comment1.IsEqualNode(comment3) {
		t.Error("Comment nodes with different data should not be equal")
	}

	// Empty comments
	comment4 := doc.CreateComment("")
	comment5 := doc.CreateComment("")
	if !comment4.IsEqualNode(comment5) {
		t.Error("Empty comment nodes should be equal")
	}
}

func TestIsEqualNode_ProcessingInstruction(t *testing.T) {
	doc := NewDocument()

	// Equal processing instructions
	pi1, _ := doc.CreateProcessingInstruction("target", "data")
	pi2, _ := doc.CreateProcessingInstruction("target", "data")
	if !pi1.IsEqualNode(pi2) {
		t.Error("ProcessingInstruction nodes with same target and data should be equal")
	}

	// Different targets
	pi3, _ := doc.CreateProcessingInstruction("different", "data")
	if pi1.IsEqualNode(pi3) {
		t.Error("ProcessingInstruction nodes with different targets should not be equal")
	}

	// Different data
	pi4, _ := doc.CreateProcessingInstruction("target", "different")
	if pi1.IsEqualNode(pi4) {
		t.Error("ProcessingInstruction nodes with different data should not be equal")
	}
}

func TestIsEqualNode_DocumentType(t *testing.T) {
	doc := NewDocument()

	// Equal document types
	dt1 := NewDocumentType("html", "-//W3C//DTD HTML 4.01//EN", "http://www.w3.org/TR/html4/strict.dtd", doc)
	dt2 := NewDocumentType("html", "-//W3C//DTD HTML 4.01//EN", "http://www.w3.org/TR/html4/strict.dtd", doc)
	if !dt1.IsEqualNode(dt2) {
		t.Error("DocumentType nodes with same name, publicId, and systemId should be equal")
	}

	// Different names
	dt3 := NewDocumentType("xml", "-//W3C//DTD HTML 4.01//EN", "http://www.w3.org/TR/html4/strict.dtd", doc)
	if dt1.IsEqualNode(dt3) {
		t.Error("DocumentType nodes with different names should not be equal")
	}

	// Different public IDs
	dt4 := NewDocumentType("html", "different", "http://www.w3.org/TR/html4/strict.dtd", doc)
	if dt1.IsEqualNode(dt4) {
		t.Error("DocumentType nodes with different public IDs should not be equal")
	}

	// Different system IDs
	dt5 := NewDocumentType("html", "-//W3C//DTD HTML 4.01//EN", "different", doc)
	if dt1.IsEqualNode(dt5) {
		t.Error("DocumentType nodes with different system IDs should not be equal")
	}
}

func TestIsEqualNode_Attr(t *testing.T) {
	doc := NewDocument()

	// Equal attributes
	attr1, _ := doc.CreateAttribute("class")
	attr1.SetValue("test")
	attr2, _ := doc.CreateAttribute("class")
	attr2.SetValue("test")
	if !attr1.IsEqualNode(attr2) {
		t.Error("Attr nodes with same name and value should be equal")
	}

	// Different names
	attr3, _ := doc.CreateAttribute("id")
	attr3.SetValue("test")
	if attr1.IsEqualNode(attr3) {
		t.Error("Attr nodes with different names should not be equal")
	}

	// Different values
	attr4, _ := doc.CreateAttribute("class")
	attr4.SetValue("different")
	if attr1.IsEqualNode(attr4) {
		t.Error("Attr nodes with different values should not be equal")
	}

	// Empty values
	attr5, _ := doc.CreateAttribute("empty")
	attr5.SetValue("")
	attr6, _ := doc.CreateAttribute("empty")
	attr6.SetValue("")
	if !attr5.IsEqualNode(attr6) {
		t.Error("Attr nodes with empty values should be equal")
	}
}

func TestIsEqualNode_Element_Basic(t *testing.T) {
	doc := NewDocument()

	// Equal elements with same tag name
	elem1 := doc.CreateElement("div")
	elem2 := doc.CreateElement("div")
	if !elem1.IsEqualNode(elem2) {
		t.Error("Elements with same tag name should be equal")
	}

	// Different tag names
	elem3 := doc.CreateElement("span")
	if elem1.IsEqualNode(elem3) {
		t.Error("Elements with different tag names should not be equal")
	}
}

func TestIsEqualNode_Element_WithAttributes(t *testing.T) {
	doc := NewDocument()

	// Equal elements with same attributes
	elem1 := doc.CreateElement("div")
	elem1.SetAttribute("class", "test")
	elem1.SetAttribute("id", "myid")

	elem2 := doc.CreateElement("div")
	elem2.SetAttribute("class", "test")
	elem2.SetAttribute("id", "myid")

	if !elem1.IsEqualNode(elem2) {
		t.Error("Elements with same attributes should be equal")
	}

	// Different attribute values
	elem3 := doc.CreateElement("div")
	elem3.SetAttribute("class", "different")
	elem3.SetAttribute("id", "myid")

	if elem1.IsEqualNode(elem3) {
		t.Error("Elements with different attribute values should not be equal")
	}

	// Missing attributes
	elem4 := doc.CreateElement("div")
	elem4.SetAttribute("class", "test")
	// Missing id attribute

	if elem1.IsEqualNode(elem4) {
		t.Error("Elements with missing attributes should not be equal")
	}

	// Extra attributes
	elem5 := doc.CreateElement("div")
	elem5.SetAttribute("class", "test")
	elem5.SetAttribute("id", "myid")
	elem5.SetAttribute("data-extra", "value")

	if elem1.IsEqualNode(elem5) {
		t.Error("Elements with extra attributes should not be equal")
	}
}

func TestIsEqualNode_Element_AttributeOrder(t *testing.T) {
	doc := NewDocument()

	// Elements with same attributes in different order should be equal
	elem1 := doc.CreateElement("div")
	elem1.SetAttribute("class", "test")
	elem1.SetAttribute("id", "myid")

	elem2 := doc.CreateElement("div")
	elem2.SetAttribute("id", "myid")
	elem2.SetAttribute("class", "test")

	if !elem1.IsEqualNode(elem2) {
		t.Error("Elements with same attributes in different order should be equal")
	}
}

func TestIsEqualNode_Element_WithChildren(t *testing.T) {
	doc := NewDocument()

	// Equal elements with same children
	elem1 := doc.CreateElement("div")
	child1 := doc.CreateElement("span")
	child1.SetAttribute("class", "child")
	elem1.AppendChild(child1)

	elem2 := doc.CreateElement("div")
	child2 := doc.CreateElement("span")
	child2.SetAttribute("class", "child")
	elem2.AppendChild(child2)

	if !elem1.IsEqualNode(elem2) {
		t.Error("Elements with equal children should be equal")
	}

	// Different number of children
	elem3 := doc.CreateElement("div")
	child3 := doc.CreateElement("span")
	child4 := doc.CreateElement("p")
	elem3.AppendChild(child3)
	elem3.AppendChild(child4)

	if elem1.IsEqualNode(elem3) {
		t.Error("Elements with different number of children should not be equal")
	}

	// Children in different order
	elem4 := doc.CreateElement("div")
	child5 := doc.CreateElement("p")
	child6 := doc.CreateElement("span")
	elem4.AppendChild(child5)
	elem4.AppendChild(child6)

	if elem3.IsEqualNode(elem4) {
		t.Error("Elements with children in different order should not be equal")
	}
}

func TestIsEqualNode_Element_NestedChildren(t *testing.T) {
	doc := NewDocument()

	// Create deeply nested equal structures
	elem1 := doc.CreateElement("div")
	child1 := doc.CreateElement("span")
	grandchild1 := doc.CreateTextNode("text")
	child1.AppendChild(grandchild1)
	elem1.AppendChild(child1)

	elem2 := doc.CreateElement("div")
	child2 := doc.CreateElement("span")
	grandchild2 := doc.CreateTextNode("text")
	child2.AppendChild(grandchild2)
	elem2.AppendChild(child2)

	if !elem1.IsEqualNode(elem2) {
		t.Error("Elements with nested equal children should be equal")
	}

	// Different nested content
	elem3 := doc.CreateElement("div")
	child3 := doc.CreateElement("span")
	grandchild3 := doc.CreateTextNode("different")
	child3.AppendChild(grandchild3)
	elem3.AppendChild(child3)

	if elem1.IsEqualNode(elem3) {
		t.Error("Elements with different nested content should not be equal")
	}
}

func TestIsEqualNode_Document(t *testing.T) {
	// Create two documents with same structure
	doc1 := NewDocument()
	doctype1 := NewDocumentType("html", "", "", doc1)
	doc1.AppendChild(doctype1)
	root1 := doc1.CreateElement("html")
	doc1.AppendChild(root1)

	doc2 := NewDocument()
	doctype2 := NewDocumentType("html", "", "", doc2)
	doc2.AppendChild(doctype2)
	root2 := doc2.CreateElement("html")
	doc2.AppendChild(root2)

	if !doc1.IsEqualNode(doc2) {
		t.Error("Documents with same structure should be equal")
	}

	// Different document structure
	doc3 := NewDocument()
	root3 := doc3.CreateElement("xml")
	doc3.AppendChild(root3)

	if doc1.IsEqualNode(doc3) {
		t.Error("Documents with different structure should not be equal")
	}
}

func TestIsEqualNode_DocumentFragment(t *testing.T) {
	doc := NewDocument()

	// Equal document fragments
	frag1 := doc.CreateDocumentFragment()
	elem1 := doc.CreateElement("div")
	frag1.AppendChild(elem1)

	frag2 := doc.CreateDocumentFragment()
	elem2 := doc.CreateElement("div")
	frag2.AppendChild(elem2)

	if !frag1.IsEqualNode(frag2) {
		t.Error("DocumentFragments with equal children should be equal")
	}

	// Empty fragments
	frag3 := doc.CreateDocumentFragment()
	frag4 := doc.CreateDocumentFragment()

	if !frag3.IsEqualNode(frag4) {
		t.Error("Empty DocumentFragments should be equal")
	}

	// Different children
	frag5 := doc.CreateDocumentFragment()
	elem3 := doc.CreateElement("span")
	frag5.AppendChild(elem3)

	if frag1.IsEqualNode(frag5) {
		t.Error("DocumentFragments with different children should not be equal")
	}
}

func TestIsEqualNode_MixedContent(t *testing.T) {
	doc := NewDocument()

	// Create elements with mixed content (text and elements)
	elem1 := doc.CreateElement("p")
	text1 := doc.CreateTextNode("Hello ")
	bold1 := doc.CreateElement("b")
	boldText1 := doc.CreateTextNode("world")
	bold1.AppendChild(boldText1)
	text2 := doc.CreateTextNode("!")
	elem1.AppendChild(text1)
	elem1.AppendChild(bold1)
	elem1.AppendChild(text2)

	elem2 := doc.CreateElement("p")
	text3 := doc.CreateTextNode("Hello ")
	bold2 := doc.CreateElement("b")
	boldText2 := doc.CreateTextNode("world")
	bold2.AppendChild(boldText2)
	text4 := doc.CreateTextNode("!")
	elem2.AppendChild(text3)
	elem2.AppendChild(bold2)
	elem2.AppendChild(text4)

	if !elem1.IsEqualNode(elem2) {
		t.Error("Elements with same mixed content should be equal")
	}

	// Different mixed content
	elem3 := doc.CreateElement("p")
	text5 := doc.CreateTextNode("Hello ")
	italic := doc.CreateElement("i") // Different tag
	italicText := doc.CreateTextNode("world")
	italic.AppendChild(italicText)
	text6 := doc.CreateTextNode("!")
	elem3.AppendChild(text5)
	elem3.AppendChild(italic)
	elem3.AppendChild(text6)

	if elem1.IsEqualNode(elem3) {
		t.Error("Elements with different mixed content should not be equal")
	}
}

func TestIsEqualNode_Namespace(t *testing.T) {
	doc := NewDocument()

	// Elements with same namespace
	elem1, _ := doc.CreateElementNS("http://www.w3.org/1999/xhtml", "div")
	elem2, _ := doc.CreateElementNS("http://www.w3.org/1999/xhtml", "div")
	if !elem1.IsEqualNode(elem2) {
		t.Error("Elements with same namespace should be equal")
	}

	// Elements with different namespaces
	elem3, _ := doc.CreateElementNS("http://www.w3.org/2000/svg", "div")
	if elem1.IsEqualNode(elem3) {
		t.Error("Elements with different namespaces should not be equal")
	}

	// Element with namespace vs without
	elem4 := doc.CreateElement("div")
	if elem1.IsEqualNode(elem4) {
		t.Error("Element with namespace should not equal element without namespace")
	}
}

func TestIsEqualNode_WithPrefix(t *testing.T) {
	doc := NewDocument()

	// Elements with same prefix
	elem1, _ := doc.CreateElementNS("http://example.com", "pre:element")
	elem2, _ := doc.CreateElementNS("http://example.com", "pre:element")
	if !elem1.IsEqualNode(elem2) {
		t.Error("Elements with same prefix should be equal")
	}

	// Elements with different prefixes
	elem3, _ := doc.CreateElementNS("http://example.com", "other:element")
	if elem1.IsEqualNode(elem3) {
		t.Error("Elements with different prefixes should not be equal")
	}
}

func TestIsEqualNode_ComplexExample(t *testing.T) {
	doc := NewDocument()

	// Create complex equal structures
	createComplexElement := func() *Element {
		div := doc.CreateElement("div")
		div.SetAttribute("class", "container")
		div.SetAttribute("id", "main")

		header := doc.CreateElement("header")
		title := doc.CreateElement("h1")
		titleText := doc.CreateTextNode("Title")
		title.AppendChild(titleText)
		header.AppendChild(title)
		div.AppendChild(header)

		content := doc.CreateElement("main")
		p := doc.CreateElement("p")
		pText := doc.CreateTextNode("Content goes here")
		p.AppendChild(pText)
		content.AppendChild(p)
		div.AppendChild(content)

		return div
	}

	elem1 := createComplexElement()
	elem2 := createComplexElement()

	if !elem1.IsEqualNode(elem2) {
		t.Error("Complex equal structures should be equal")
	}

	// Modify one slightly
	elem3 := createComplexElement()
	elem3.SetAttribute("data-modified", "true")

	if elem1.IsEqualNode(elem3) {
		t.Error("Modified complex structure should not be equal")
	}
}
