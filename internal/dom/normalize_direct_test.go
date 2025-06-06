package dom

import (
	"testing"
)

func TestNormalizeDirectDOMTest(t *testing.T) {
	// Create a document and elements directly
	doc := NewDocument()

	// Create a proper document structure
	html := NewElement("html", doc)
	body := NewElement("body", doc)
	doc.AppendChild(html)
	html.AppendChild(body)

	// Create a paragraph element
	para := NewElement("p", doc)
	body.AppendChild(para)

	// Create the initial text node
	text1 := NewText("First", doc)
	para.AppendChild(text1)

	t.Logf("=== DIRECT DOM NORMALIZE TEST ===")
	t.Logf("Initial children count: %d", len(para.nodeImpl.childNodes))
	t.Logf("Initial text content: %s", para.TextContent())

	// Add additional text nodes
	text2 := NewText(" Second", doc)
	text3 := NewText(" Third", doc)
	para.AppendChild(text2)
	para.AppendChild(text3)

	t.Logf("After adding nodes - children count: %d", len(para.nodeImpl.childNodes))
	for i, child := range para.nodeImpl.childNodes {
		if textNode, ok := child.(*Text); ok {
			t.Logf("Child %d: nodeType=%d data='%s'", i, child.NodeType(), textNode.Data())
		} else {
			t.Logf("Child %d: nodeType=%d", i, child.NodeType())
		}
	}

	// Now call normalize directly
	t.Logf("Calling normalize...")
	para.Normalize()

	t.Logf("After normalize - children count: %d", len(para.nodeImpl.childNodes))
	for i, child := range para.nodeImpl.childNodes {
		if textNode, ok := child.(*Text); ok {
			t.Logf("Child %d: nodeType=%d data='%s'", i, child.NodeType(), textNode.Data())
		} else {
			t.Logf("Child %d: nodeType=%d", i, child.NodeType())
		}
	}

	t.Logf("Final text content: %s", para.TextContent())
	t.Logf("=== END DIRECT DOM NORMALIZE TEST ===")

	// Test that normalization worked
	if len(para.nodeImpl.childNodes) != 1 {
		t.Errorf("Expected 1 child node after normalization, got %d", len(para.nodeImpl.childNodes))
	}

	if para.TextContent() != "First Second Third" {
		t.Errorf("Expected text content 'First Second Third', got '%s'", para.TextContent())
	}
}

func TestNormalizeEmptyTextNodes(t *testing.T) {
	// Test normalization with empty text nodes
	doc := NewDocument()
	html := NewElement("html", doc)
	body := NewElement("body", doc)
	doc.AppendChild(html)
	html.AppendChild(body)

	para := NewElement("p", doc)
	body.AppendChild(para)

	// Create text nodes including empty ones
	text1 := NewText("Hello", doc)
	text2 := NewText("", doc) // Empty
	text3 := NewText(" World", doc)
	text4 := NewText("", doc) // Empty

	para.AppendChild(text1)
	para.AppendChild(text2)
	para.AppendChild(text3)
	para.AppendChild(text4)

	t.Logf("Before normalize: %d children", len(para.nodeImpl.childNodes))
	para.Normalize()
	t.Logf("After normalize: %d children", len(para.nodeImpl.childNodes))

	// Should have 1 child (empty nodes removed, others merged)
	if len(para.nodeImpl.childNodes) != 1 {
		t.Errorf("Expected 1 child after normalization, got %d", len(para.nodeImpl.childNodes))
	}

	if para.TextContent() != "Hello World" {
		t.Errorf("Expected 'Hello World', got '%s'", para.TextContent())
	}
}

func TestNormalizeWithNonTextNodes(t *testing.T) {
	// Test normalization with mixed content
	doc := NewDocument()
	html := NewElement("html", doc)
	body := NewElement("body", doc)
	doc.AppendChild(html)
	html.AppendChild(body)

	para := NewElement("p", doc)
	body.AppendChild(para)

	// Create mixed content: text, element, text, text
	text1 := NewText("Start ", doc)
	elem := NewElement("em", doc)
	elem.AppendChild(NewText("emphasized", doc))
	text2 := NewText(" middle", doc)
	text3 := NewText(" end", doc)

	para.AppendChild(text1)
	para.AppendChild(elem)
	para.AppendChild(text2)
	para.AppendChild(text3)

	t.Logf("Before normalize: %d children", len(para.nodeImpl.childNodes))
	para.Normalize()
	t.Logf("After normalize: %d children", len(para.nodeImpl.childNodes))

	// Should have 3 children: text1 (unchanged), elem, merged text2+text3
	if len(para.nodeImpl.childNodes) != 3 {
		t.Errorf("Expected 3 children after normalization, got %d", len(para.nodeImpl.childNodes))
	}

	// Check that the last two text nodes were merged
	lastChild := para.nodeImpl.childNodes[2]
	if textNode, ok := lastChild.(*Text); ok {
		if textNode.Data() != " middle end" {
			t.Errorf("Expected last text node to be ' middle end', got '%s'", textNode.Data())
		}
	} else {
		t.Errorf("Expected last child to be a text node")
	}
}
