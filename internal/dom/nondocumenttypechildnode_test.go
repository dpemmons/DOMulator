package dom

import (
	"testing"
)

func TestElementPreviousElementSibling(t *testing.T) {
	doc := NewDocument()
	parent := NewElement("div", doc)

	// Create various node types
	text1 := NewText("text1", doc)
	elem1 := NewElement("span", doc)
	comment1 := NewComment("comment1", doc)
	elem2 := NewElement("p", doc)
	text2 := NewText("text2", doc)
	elem3 := NewElement("div", doc)

	// Add nodes to parent in order: text1, elem1, comment1, elem2, text2, elem3
	parent.AppendChild(text1)
	parent.AppendChild(elem1)
	parent.AppendChild(comment1)
	parent.AppendChild(elem2)
	parent.AppendChild(text2)
	parent.AppendChild(elem3)

	// Test elem1.PreviousElementSibling() - should be nil (only text before it)
	if prev := elem1.PreviousElementSibling(); prev != nil {
		t.Errorf("Expected elem1.PreviousElementSibling() to be nil, got %v", prev)
	}

	// Test elem2.PreviousElementSibling() - should be elem1
	if prev := elem2.PreviousElementSibling(); prev != elem1 {
		t.Errorf("Expected elem2.PreviousElementSibling() to be elem1, got %v", prev)
	}

	// Test elem3.PreviousElementSibling() - should be elem2
	if prev := elem3.PreviousElementSibling(); prev != elem2 {
		t.Errorf("Expected elem3.PreviousElementSibling() to be elem2, got %v", prev)
	}
}

func TestElementNextElementSibling(t *testing.T) {
	doc := NewDocument()
	parent := NewElement("div", doc)

	// Create various node types
	elem1 := NewElement("span", doc)
	text1 := NewText("text1", doc)
	comment1 := NewComment("comment1", doc)
	elem2 := NewElement("p", doc)
	text2 := NewText("text2", doc)
	elem3 := NewElement("div", doc)

	// Add nodes to parent in order: elem1, text1, comment1, elem2, text2, elem3
	parent.AppendChild(elem1)
	parent.AppendChild(text1)
	parent.AppendChild(comment1)
	parent.AppendChild(elem2)
	parent.AppendChild(text2)
	parent.AppendChild(elem3)

	// Test elem1.NextElementSibling() - should be elem2
	if next := elem1.NextElementSibling(); next != elem2 {
		t.Errorf("Expected elem1.NextElementSibling() to be elem2, got %v", next)
	}

	// Test elem2.NextElementSibling() - should be elem3
	if next := elem2.NextElementSibling(); next != elem3 {
		t.Errorf("Expected elem2.NextElementSibling() to be elem3, got %v", next)
	}

	// Test elem3.NextElementSibling() - should be nil (no element after it)
	if next := elem3.NextElementSibling(); next != nil {
		t.Errorf("Expected elem3.NextElementSibling() to be nil, got %v", next)
	}
}

func TestElementSiblingsWithNoParent(t *testing.T) {
	doc := NewDocument()
	elem := NewElement("div", doc)

	// Element with no parent should return nil for both methods
	if prev := elem.PreviousElementSibling(); prev != nil {
		t.Errorf("Expected PreviousElementSibling() to be nil for element with no parent, got %v", prev)
	}

	if next := elem.NextElementSibling(); next != nil {
		t.Errorf("Expected NextElementSibling() to be nil for element with no parent, got %v", next)
	}
}

func TestTextPreviousElementSibling(t *testing.T) {
	doc := NewDocument()
	parent := NewElement("div", doc)

	// Create various node types
	elem1 := NewElement("span", doc)
	text1 := NewText("text1", doc)
	comment1 := NewComment("comment1", doc)
	text2 := NewText("text2", doc)
	elem2 := NewElement("p", doc)

	// Add nodes to parent in order: elem1, text1, comment1, text2, elem2
	parent.AppendChild(elem1)
	parent.AppendChild(text1)
	parent.AppendChild(comment1)
	parent.AppendChild(text2)
	parent.AppendChild(elem2)

	// Test text1.PreviousElementSibling() - should be elem1
	if prev := text1.PreviousElementSibling(); prev != elem1 {
		t.Errorf("Expected text1.PreviousElementSibling() to be elem1, got %v", prev)
	}

	// Test text2.PreviousElementSibling() - should be elem1 (skipping text1 and comment1)
	if prev := text2.PreviousElementSibling(); prev != elem1 {
		t.Errorf("Expected text2.PreviousElementSibling() to be elem1, got %v", prev)
	}
}

func TestTextNextElementSibling(t *testing.T) {
	doc := NewDocument()
	parent := NewElement("div", doc)

	// Create various node types
	elem1 := NewElement("span", doc)
	text1 := NewText("text1", doc)
	comment1 := NewComment("comment1", doc)
	text2 := NewText("text2", doc)
	elem2 := NewElement("p", doc)

	// Add nodes to parent in order: elem1, text1, comment1, text2, elem2
	parent.AppendChild(elem1)
	parent.AppendChild(text1)
	parent.AppendChild(comment1)
	parent.AppendChild(text2)
	parent.AppendChild(elem2)

	// Test text1.NextElementSibling() - should be elem2
	if next := text1.NextElementSibling(); next != elem2 {
		t.Errorf("Expected text1.NextElementSibling() to be elem2, got %v", next)
	}

	// Test text2.NextElementSibling() - should be elem2
	if next := text2.NextElementSibling(); next != elem2 {
		t.Errorf("Expected text2.NextElementSibling() to be elem2, got %v", next)
	}
}

func TestCommentPreviousElementSibling(t *testing.T) {
	doc := NewDocument()
	parent := NewElement("div", doc)

	// Create various node types
	elem1 := NewElement("span", doc)
	text1 := NewText("text1", doc)
	comment1 := NewComment("comment1", doc)
	elem2 := NewElement("p", doc)

	// Add nodes to parent in order: elem1, text1, comment1, elem2
	parent.AppendChild(elem1)
	parent.AppendChild(text1)
	parent.AppendChild(comment1)
	parent.AppendChild(elem2)

	// Test comment1.PreviousElementSibling() - should be elem1
	if prev := comment1.PreviousElementSibling(); prev != elem1 {
		t.Errorf("Expected comment1.PreviousElementSibling() to be elem1, got %v", prev)
	}
}

func TestCommentNextElementSibling(t *testing.T) {
	doc := NewDocument()
	parent := NewElement("div", doc)

	// Create various node types
	elem1 := NewElement("span", doc)
	comment1 := NewComment("comment1", doc)
	text1 := NewText("text1", doc)
	elem2 := NewElement("p", doc)

	// Add nodes to parent in order: elem1, comment1, text1, elem2
	parent.AppendChild(elem1)
	parent.AppendChild(comment1)
	parent.AppendChild(text1)
	parent.AppendChild(elem2)

	// Test comment1.NextElementSibling() - should be elem2
	if next := comment1.NextElementSibling(); next != elem2 {
		t.Errorf("Expected comment1.NextElementSibling() to be elem2, got %v", next)
	}
}

func TestCDATASectionPreviousElementSibling(t *testing.T) {
	doc := NewDocument()
	parent := NewElement("div", doc)

	// Create various node types
	elem1 := NewElement("span", doc)
	text1 := NewText("text1", doc)
	cdata1 := NewCDATASection("cdata content", doc)
	elem2 := NewElement("p", doc)

	// Add nodes to parent in order: elem1, text1, cdata1, elem2
	parent.AppendChild(elem1)
	parent.AppendChild(text1)
	parent.AppendChild(cdata1)
	parent.AppendChild(elem2)

	// Test cdata1.PreviousElementSibling() - should be elem1
	if prev := cdata1.PreviousElementSibling(); prev != elem1 {
		t.Errorf("Expected cdata1.PreviousElementSibling() to be elem1, got %v", prev)
	}
}

func TestCDATASectionNextElementSibling(t *testing.T) {
	doc := NewDocument()
	parent := NewElement("div", doc)

	// Create various node types
	elem1 := NewElement("span", doc)
	cdata1 := NewCDATASection("cdata content", doc)
	text1 := NewText("text1", doc)
	elem2 := NewElement("p", doc)

	// Add nodes to parent in order: elem1, cdata1, text1, elem2
	parent.AppendChild(elem1)
	parent.AppendChild(cdata1)
	parent.AppendChild(text1)
	parent.AppendChild(elem2)

	// Test cdata1.NextElementSibling() - should be elem2
	if next := cdata1.NextElementSibling(); next != elem2 {
		t.Errorf("Expected cdata1.NextElementSibling() to be elem2, got %v", next)
	}
}

func TestCharacterDataNoParent(t *testing.T) {
	doc := NewDocument()
	text := NewText("text", doc)
	comment := NewComment("comment", doc)
	cdata := NewCDATASection("cdata", doc)

	// All character data nodes with no parent should return nil
	tests := []struct {
		name string
		node Node
	}{
		{"text", text},
		{"comment", comment},
		{"cdata", cdata},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Test PreviousElementSibling
			var prev *Element
			switch n := test.node.(type) {
			case *Text:
				prev = n.PreviousElementSibling()
			case *Comment:
				prev = n.PreviousElementSibling()
			case *CDATASection:
				prev = n.PreviousElementSibling()
			}

			if prev != nil {
				t.Errorf("Expected PreviousElementSibling() to be nil for %s with no parent, got %v", test.name, prev)
			}

			// Test NextElementSibling
			var next *Element
			switch n := test.node.(type) {
			case *Text:
				next = n.NextElementSibling()
			case *Comment:
				next = n.NextElementSibling()
			case *CDATASection:
				next = n.NextElementSibling()
			}

			if next != nil {
				t.Errorf("Expected NextElementSibling() to be nil for %s with no parent, got %v", test.name, next)
			}
		})
	}
}

func TestOnlyNonElementSiblings(t *testing.T) {
	doc := NewDocument()
	parent := NewElement("div", doc)

	// Create only non-element nodes
	text1 := NewText("text1", doc)
	comment1 := NewComment("comment1", doc)
	text2 := NewText("text2", doc)
	cdata1 := NewCDATASection("cdata", doc)
	text3 := NewText("text3", doc)

	// Add nodes to parent
	parent.AppendChild(text1)
	parent.AppendChild(comment1)
	parent.AppendChild(text2)
	parent.AppendChild(cdata1)
	parent.AppendChild(text3)

	// All nodes should return nil for element siblings since there are no elements
	tests := []struct {
		name string
		node Node
	}{
		{"text1", text1},
		{"comment1", comment1},
		{"text2", text2},
		{"cdata1", cdata1},
		{"text3", text3},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Test PreviousElementSibling
			var prev *Element
			switch n := test.node.(type) {
			case *Text:
				prev = n.PreviousElementSibling()
			case *Comment:
				prev = n.PreviousElementSibling()
			case *CDATASection:
				prev = n.PreviousElementSibling()
			}

			if prev != nil {
				t.Errorf("Expected PreviousElementSibling() to be nil for %s with no element siblings, got %v", test.name, prev)
			}

			// Test NextElementSibling
			var next *Element
			switch n := test.node.(type) {
			case *Text:
				next = n.NextElementSibling()
			case *Comment:
				next = n.NextElementSibling()
			case *CDATASection:
				next = n.NextElementSibling()
			}

			if next != nil {
				t.Errorf("Expected NextElementSibling() to be nil for %s with no element siblings, got %v", test.name, next)
			}
		})
	}
}

func TestDocumentTypeExclusion(t *testing.T) {
	doc := NewDocument()
	doctype := NewDocumentType("html", "", "", doc)

	// DocumentType should not have these methods - this is verified by compilation
	// The spec says these methods are specifically excluded from DocumentType
	// We can't test the absence of methods directly, but we can verify that
	// DocumentType doesn't implement NonDocumentTypeChildNode interface

	// This test verifies that DocumentType nodes exist and work
	if doctype.NodeName() != "html" {
		t.Errorf("Expected DocumentType node name to be 'html', got %s", doctype.NodeName())
	}

	if doctype.NodeType() != DocumentTypeNode {
		t.Errorf("Expected DocumentType node type to be DocumentTypeNode, got %d", doctype.NodeType())
	}
}

func TestComplexSiblingStructure(t *testing.T) {
	doc := NewDocument()
	parent := NewElement("div", doc)

	// Create a complex structure with multiple elements and non-elements
	elem1 := NewElement("h1", doc)
	text1 := NewText("text1", doc)
	comment1 := NewComment("comment1", doc)
	elem2 := NewElement("p", doc)
	cdata1 := NewCDATASection("cdata1", doc)
	text2 := NewText("text2", doc)
	elem3 := NewElement("span", doc)
	comment2 := NewComment("comment2", doc)
	elem4 := NewElement("div", doc)

	// Add nodes: elem1, text1, comment1, elem2, cdata1, text2, elem3, comment2, elem4
	nodes := []Node{elem1, text1, comment1, elem2, cdata1, text2, elem3, comment2, elem4}
	for _, node := range nodes {
		parent.AppendChild(node)
	}

	// Test element siblings
	if prev := elem1.PreviousElementSibling(); prev != nil {
		t.Errorf("Expected elem1.PreviousElementSibling() to be nil, got %v", prev)
	}
	if next := elem1.NextElementSibling(); next != elem2 {
		t.Errorf("Expected elem1.NextElementSibling() to be elem2, got %v", next)
	}

	if prev := elem2.PreviousElementSibling(); prev != elem1 {
		t.Errorf("Expected elem2.PreviousElementSibling() to be elem1, got %v", prev)
	}
	if next := elem2.NextElementSibling(); next != elem3 {
		t.Errorf("Expected elem2.NextElementSibling() to be elem3, got %v", next)
	}

	if prev := elem3.PreviousElementSibling(); prev != elem2 {
		t.Errorf("Expected elem3.PreviousElementSibling() to be elem2, got %v", prev)
	}
	if next := elem3.NextElementSibling(); next != elem4 {
		t.Errorf("Expected elem3.NextElementSibling() to be elem4, got %v", next)
	}

	if prev := elem4.PreviousElementSibling(); prev != elem3 {
		t.Errorf("Expected elem4.PreviousElementSibling() to be elem3, got %v", prev)
	}
	if next := elem4.NextElementSibling(); next != nil {
		t.Errorf("Expected elem4.NextElementSibling() to be nil, got %v", next)
	}

	// Test non-element siblings
	if prev := text1.PreviousElementSibling(); prev != elem1 {
		t.Errorf("Expected text1.PreviousElementSibling() to be elem1, got %v", prev)
	}
	if next := text1.NextElementSibling(); next != elem2 {
		t.Errorf("Expected text1.NextElementSibling() to be elem2, got %v", next)
	}

	if prev := comment1.PreviousElementSibling(); prev != elem1 {
		t.Errorf("Expected comment1.PreviousElementSibling() to be elem1, got %v", prev)
	}
	if next := comment1.NextElementSibling(); next != elem2 {
		t.Errorf("Expected comment1.NextElementSibling() to be elem2, got %v", next)
	}

	if prev := cdata1.PreviousElementSibling(); prev != elem2 {
		t.Errorf("Expected cdata1.PreviousElementSibling() to be elem2, got %v", prev)
	}
	if next := cdata1.NextElementSibling(); next != elem3 {
		t.Errorf("Expected cdata1.NextElementSibling() to be elem3, got %v", next)
	}
}
