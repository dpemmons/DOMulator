package dom

import (
	"testing"
)

func TestTextCreation(t *testing.T) {
	doc := NewDocument()
	text := NewText("Hello World", doc)

	if text.NodeType() != TextNode {
		t.Errorf("Expected node type %v, got %v", TextNode, text.NodeType())
	}
	if text.NodeName() != "#text" {
		t.Errorf("Expected node name '#text', got %s", text.NodeName())
	}
	if text.NodeValue() != "Hello World" {
		t.Errorf("Expected text 'Hello World', got %s", text.NodeValue())
	}
	if text.OwnerDocument() != doc {
		t.Errorf("Expected owner document to match")
	}
	if text.ParentNode() != nil {
		t.Errorf("Expected parent node to be nil")
	}
	if len(text.ChildNodes()) != 0 {
		t.Errorf("Expected no child nodes, got %d", len(text.ChildNodes()))
	}
}

func TestTextVariousContent(t *testing.T) {
	doc := NewDocument()

	// Test with empty string
	emptyText := NewText("", doc)
	if emptyText.NodeValue() != "" {
		t.Errorf("Expected empty string, got '%s'", emptyText.NodeValue())
	}

	// Test with special characters
	specialText := NewText("Hello <world> & \"friends\"!", doc)
	if specialText.NodeValue() != "Hello <world> & \"friends\"!" {
		t.Errorf("Expected special chars text, got '%s'", specialText.NodeValue())
	}

	// Test with whitespace
	whitespaceText := NewText("  \t\n  Hello  \t\n  ", doc)
	if whitespaceText.NodeValue() != "  \t\n  Hello  \t\n  " {
		t.Errorf("Expected whitespace preserved, got '%s'", whitespaceText.NodeValue())
	}

	// Test with unicode
	unicodeText := NewText("Hello 世界 🌍", doc)
	if unicodeText.NodeValue() != "Hello 世界 🌍" {
		t.Errorf("Expected unicode text, got '%s'", unicodeText.NodeValue())
	}
}

func TestTextData(t *testing.T) {
	doc := NewDocument()
	text := NewText("Initial text", doc)

	// Test Data() method
	if text.Data() != "Initial text" {
		t.Errorf("Expected 'Initial text', got '%s'", text.Data())
	}

	// Test SetData() method
	text.SetData("Updated text")
	if text.Data() != "Updated text" {
		t.Errorf("Expected 'Updated text', got '%s'", text.Data())
	}
	if text.NodeValue() != "Updated text" {
		t.Errorf("Expected NodeValue to match Data, got '%s'", text.NodeValue())
	}

	// Test setting empty data
	text.SetData("")
	if text.Data() != "" {
		t.Errorf("Expected empty data, got '%s'", text.Data())
	}

	// Test setting data with special characters
	text.SetData("<script>alert('xss')</script>")
	if text.Data() != "<script>alert('xss')</script>" {
		t.Errorf("Expected script text preserved, got '%s'", text.Data())
	}
}

func TestTextLength(t *testing.T) {
	doc := NewDocument()

	// Test empty text
	emptyText := NewText("", doc)
	if emptyText.Length() != 0 {
		t.Errorf("Expected length 0 for empty text, got %d", emptyText.Length())
	}

	// Test regular text
	text := NewText("Hello", doc)
	if text.Length() != 5 {
		t.Errorf("Expected length 5, got %d", text.Length())
	}

	// Test unicode text
	unicodeText := NewText("Hello 世界", doc)
	if unicodeText.Length() != 8 { // 6 ASCII + 2 unicode chars
		t.Errorf("Expected length 8 for unicode text, got %d", unicodeText.Length())
	}

	// Test after modification
	text.SetData("Updated text content")
	if text.Length() != 20 {
		t.Errorf("Expected length 20 after update, got %d", text.Length())
	}
}

func TestTextSubstringData(t *testing.T) {
	doc := NewDocument()
	text := NewText("Hello World", doc)

	// Test valid substring
	substr := text.SubstringData(0, 5)
	if substr != "Hello" {
		t.Errorf("Expected 'Hello', got '%s'", substr)
	}

	// Test substring from middle
	substr = text.SubstringData(6, 5)
	if substr != "World" {
		t.Errorf("Expected 'World', got '%s'", substr)
	}

	// Test substring with count exceeding length
	substr = text.SubstringData(6, 10)
	if substr != "World" {
		t.Errorf("Expected 'World' (truncated), got '%s'", substr)
	}

	// Test substring starting from end
	substr = text.SubstringData(11, 5)
	if substr != "" {
		t.Errorf("Expected empty string when starting beyond text, got '%s'", substr)
	}

	// Test negative offset (should be treated as 0)
	substr = text.SubstringData(-1, 5)
	if substr != "Hello" {
		t.Errorf("Expected 'Hello' for negative offset, got '%s'", substr)
	}

	// Test zero count
	substr = text.SubstringData(0, 0)
	if substr != "" {
		t.Errorf("Expected empty string for zero count, got '%s'", substr)
	}
}

func TestTextAppendData(t *testing.T) {
	doc := NewDocument()
	text := NewText("Hello", doc)

	// Test appending text
	text.AppendData(" World")
	if text.Data() != "Hello World" {
		t.Errorf("Expected 'Hello World', got '%s'", text.Data())
	}

	// Test appending empty string
	text.AppendData("")
	if text.Data() != "Hello World" {
		t.Errorf("Expected 'Hello World' (unchanged), got '%s'", text.Data())
	}

	// Test appending special characters
	text.AppendData(" & <friends>!")
	if text.Data() != "Hello World & <friends>!" {
		t.Errorf("Expected special chars appended, got '%s'", text.Data())
	}

	// Test length after appending
	if text.Length() != 24 {
		t.Errorf("Expected length 24 after appending, got %d", text.Length())
	}
}

func TestTextInsertData(t *testing.T) {
	doc := NewDocument()
	text := NewText("Hello World", doc)

	// Test inserting at beginning
	text.InsertData(0, "Hi ")
	if text.Data() != "Hi Hello World" {
		t.Errorf("Expected 'Hi Hello World', got '%s'", text.Data())
	}

	// Test inserting in middle
	text.SetData("Hello World")
	text.InsertData(5, " Beautiful")
	if text.Data() != "Hello Beautiful World" {
		t.Errorf("Expected 'Hello Beautiful World', got '%s'", text.Data())
	}

	// Test inserting at end
	text.SetData("Hello World")
	text.InsertData(11, "!")
	if text.Data() != "Hello World!" {
		t.Errorf("Expected 'Hello World!', got '%s'", text.Data())
	}

	// Test inserting beyond end (should append)
	text.SetData("Hello")
	text.InsertData(10, " World")
	if text.Data() != "Hello World" {
		t.Errorf("Expected 'Hello World' (appended), got '%s'", text.Data())
	}

	// Test inserting empty string
	text.SetData("Hello World")
	text.InsertData(5, "")
	if text.Data() != "Hello World" {
		t.Errorf("Expected 'Hello World' (unchanged), got '%s'", text.Data())
	}
}

func TestTextDeleteData(t *testing.T) {
	doc := NewDocument()
	text := NewText("Hello World", doc)

	// Test deleting from beginning
	text.DeleteData(0, 6)
	if text.Data() != "World" {
		t.Errorf("Expected 'World', got '%s'", text.Data())
	}

	// Test deleting from middle
	text.SetData("Hello Beautiful World")
	text.DeleteData(5, 10) // Remove " Beautiful"
	if text.Data() != "Hello World" {
		t.Errorf("Expected 'Hello World', got '%s'", text.Data())
	}

	// Test deleting from end
	text.SetData("Hello World")
	text.DeleteData(5, 6) // Remove " World"
	if text.Data() != "Hello" {
		t.Errorf("Expected 'Hello', got '%s'", text.Data())
	}

	// Test deleting with count exceeding remaining length
	text.SetData("Hello World")
	text.DeleteData(6, 20) // Should delete "World" only
	if text.Data() != "Hello " {
		t.Errorf("Expected 'Hello ', got '%s'", text.Data())
	}

	// Test deleting beyond text length (should do nothing)
	text.SetData("Hello")
	text.DeleteData(10, 5)
	if text.Data() != "Hello" {
		t.Errorf("Expected 'Hello' (unchanged), got '%s'", text.Data())
	}

	// Test deleting with zero count
	text.SetData("Hello World")
	text.DeleteData(5, 0)
	if text.Data() != "Hello World" {
		t.Errorf("Expected 'Hello World' (unchanged), got '%s'", text.Data())
	}
}

func TestTextReplaceData(t *testing.T) {
	doc := NewDocument()
	text := NewText("Hello World", doc)

	// Test replacing from beginning
	text.ReplaceData(0, 5, "Hi")
	if text.Data() != "Hi World" {
		t.Errorf("Expected 'Hi World', got '%s'", text.Data())
	}

	// Test replacing in middle
	text.SetData("Hello World")
	text.ReplaceData(6, 5, "Earth")
	if text.Data() != "Hello Earth" {
		t.Errorf("Expected 'Hello Earth', got '%s'", text.Data())
	}

	// Test replacing with longer text
	text.SetData("Hello World")
	text.ReplaceData(6, 5, "Beautiful Planet")
	if text.Data() != "Hello Beautiful Planet" {
		t.Errorf("Expected 'Hello Beautiful Planet', got '%s'", text.Data())
	}

	// Test replacing with shorter text
	text.SetData("Hello Beautiful World")
	text.ReplaceData(6, 10, "Nice")
	if text.Data() != "Hello NiceWorld" {
		t.Errorf("Expected 'Hello NiceWorld', got '%s'", text.Data())
	}

	// Test replacing with empty string (equivalent to delete)
	text.SetData("Hello World")
	text.ReplaceData(5, 6, "")
	if text.Data() != "Hello" {
		t.Errorf("Expected 'Hello', got '%s'", text.Data())
	}

	// Test replacing beyond text length
	text.SetData("Hello")
	text.ReplaceData(10, 5, " World")
	if text.Data() != "Hello World" {
		t.Errorf("Expected 'Hello World' (appended), got '%s'", text.Data())
	}
}

func TestTextSplitText(t *testing.T) {
	doc := NewDocument()
	text := NewText("Hello World", doc)

	// Test splitting in middle
	newText := text.SplitText(6)
	if text.Data() != "Hello " {
		t.Errorf("Expected original text 'Hello ', got '%s'", text.Data())
	}
	if newText.Data() != "World" {
		t.Errorf("Expected new text 'World', got '%s'", newText.Data())
	}
	if newText.OwnerDocument() != doc {
		t.Errorf("Expected new text to have same owner document")
	}

	// Test splitting at beginning
	text.SetData("Hello World")
	newText = text.SplitText(0)
	if text.Data() != "" {
		t.Errorf("Expected original text empty, got '%s'", text.Data())
	}
	if newText.Data() != "Hello World" {
		t.Errorf("Expected new text 'Hello World', got '%s'", newText.Data())
	}

	// Test splitting at end
	text.SetData("Hello World")
	newText = text.SplitText(11)
	if text.Data() != "Hello World" {
		t.Errorf("Expected original text unchanged, got '%s'", text.Data())
	}
	if newText.Data() != "" {
		t.Errorf("Expected new text empty, got '%s'", newText.Data())
	}

	// Test splitting beyond text length
	text.SetData("Hello")
	newText = text.SplitText(10)
	if text.Data() != "Hello" {
		t.Errorf("Expected original text unchanged, got '%s'", text.Data())
	}
	if newText.Data() != "" {
		t.Errorf("Expected new text empty when splitting beyond length, got '%s'", newText.Data())
	}
}

func TestTextNodeHierarchyIntegration(t *testing.T) {
	doc := NewDocument()
	parent := NewElement("div", doc)
	text1 := NewText("Hello ", doc)
	text2 := NewText("World", doc)

	// Test that Text properly integrates with Node hierarchy
	parent.AppendChild(text1)
	parent.AppendChild(text2)

	if parent.FirstChild() != text1 {
		t.Errorf("Expected first child to be text1")
	}
	if parent.LastChild() != text2 {
		t.Errorf("Expected last child to be text2")
	}
	if text1.NextSibling() != text2 {
		t.Errorf("Expected text1's next sibling to be text2")
	}
	if text2.PreviousSibling() != text1 {
		t.Errorf("Expected text2's previous sibling to be text1")
	}
	if text1.ParentNode() != parent {
		t.Errorf("Expected text1's parent to be div")
	}

	// Test text content aggregation
	if parent.TextContent() != "Hello World" {
		t.Errorf("Expected combined text content 'Hello World', got '%s'", parent.TextContent())
	}
}

func TestTextCloning(t *testing.T) {
	doc := NewDocument()
	text := NewText("Original text", doc)

	// Test shallow clone (deep parameter should not matter for text nodes)
	clonedText := text.CloneNode(false)
	if clonedTextNode, ok := clonedText.(*Text); ok {
		if clonedTextNode.Data() != "Original text" {
			t.Errorf("Expected cloned text 'Original text', got '%s'", clonedTextNode.Data())
		}
		if clonedTextNode.OwnerDocument() != doc {
			t.Errorf("Expected cloned text to have same owner document")
		}
		if clonedTextNode.ParentNode() != nil {
			t.Errorf("Expected cloned text to have no parent")
		}
		if clonedTextNode == text {
			t.Errorf("Expected different text node instances")
		}
	} else {
		t.Errorf("Expected cloned node to be a Text node")
	}

	// Test deep clone
	deepClonedText := text.CloneNode(true)
	if deepClonedTextNode, ok := deepClonedText.(*Text); ok {
		if deepClonedTextNode.Data() != "Original text" {
			t.Errorf("Expected deep cloned text 'Original text', got '%s'", deepClonedTextNode.Data())
		}
	} else {
		t.Errorf("Expected deep cloned node to be a Text node")
	}
}

func TestTextWhitespaceHandling(t *testing.T) {
	doc := NewDocument()

	// Test preservation of whitespace
	whitespaceText := NewText("   \t\n\r   ", doc)
	if whitespaceText.Data() != "   \t\n\r   " {
		t.Errorf("Expected whitespace preserved, got '%s'", whitespaceText.Data())
	}

	// Test mixed content with whitespace
	mixedText := NewText("Hello\t\nWorld  ", doc)
	if mixedText.Data() != "Hello\t\nWorld  " {
		t.Errorf("Expected mixed content preserved, got '%s'", mixedText.Data())
	}

	// Test length calculation with whitespace
	if whitespaceText.Length() != 9 {
		t.Errorf("Expected length 9 for whitespace text, got %d", whitespaceText.Length())
	}
}

func TestTextEdgeCases(t *testing.T) {
	doc := NewDocument()
	text := NewText("Test", doc)

	// Test negative offset handling in various methods
	text.SetData("Hello World")

	// SubstringData with negative offset
	substr := text.SubstringData(-5, 5)
	if substr != "Hello" {
		t.Errorf("Expected 'Hello' for negative offset in SubstringData, got '%s'", substr)
	}

	// InsertData with negative offset
	text.SetData("Hello")
	text.InsertData(-1, "Hi ")
	if text.Data() != "Hi Hello" {
		t.Errorf("Expected 'Hi Hello' for negative offset in InsertData, got '%s'", text.Data())
	}

	// DeleteData with negative offset
	text.SetData("Hello World")
	text.DeleteData(-1, 5)
	if text.Data() != " World" {
		t.Errorf("Expected ' World' for negative offset in DeleteData, got '%s'", text.Data())
	}

	// ReplaceData with negative offset
	text.SetData("Hello World")
	text.ReplaceData(-1, 5, "Hi")
	if text.Data() != "Hi World" {
		t.Errorf("Expected 'Hi World' for negative offset in ReplaceData, got '%s'", text.Data())
	}

	// SplitText with negative offset
	text.SetData("Hello World")
	newText := text.SplitText(-1)
	if text.Data() != "" {
		t.Errorf("Expected empty original text for negative offset in SplitText, got '%s'", text.Data())
	}
	if newText.Data() != "Hello World" {
		t.Errorf("Expected 'Hello World' in new text for negative offset in SplitText, got '%s'", newText.Data())
	}
}
