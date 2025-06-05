package dom

import (
	"testing"
)

// TestTextInterface_SpecCompliance tests WHATWG DOM Section 4.11 Text interface compliance
func TestTextInterface_SpecCompliance(t *testing.T) {
	t.Run("Text constructor", func(t *testing.T) {
		// Test new Text([data = ""])
		doc := NewDocument()

		// Test with data
		text1 := NewText("Hello World", doc)
		if text1.Data() != "Hello World" {
			t.Errorf("Expected data 'Hello World', got '%s'", text1.Data())
		}
		if text1.OwnerDocument() != doc {
			t.Errorf("Expected owner document to be set correctly")
		}

		// Test with empty data (default)
		text2 := NewText("", doc)
		if text2.Data() != "" {
			t.Errorf("Expected empty data, got '%s'", text2.Data())
		}

		// Test CharacterData inheritance
		if _, ok := interface{}(text1).(CharacterData); !ok {
			t.Errorf("Text should implement CharacterData interface")
		}
	})

	t.Run("Text.splitText method compliance", func(t *testing.T) {
		doc := NewDocument()

		// Test basic splitting
		text := NewText("Hello World", doc)
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
		text = NewText("Hello World", doc)
		newText = text.SplitText(0)
		if text.Data() != "" {
			t.Errorf("Expected original text empty, got '%s'", text.Data())
		}
		if newText.Data() != "Hello World" {
			t.Errorf("Expected new text 'Hello World', got '%s'", newText.Data())
		}

		// Test splitting at end
		text = NewText("Hello World", doc)
		newText = text.SplitText(11)
		if text.Data() != "Hello World" {
			t.Errorf("Expected original text unchanged, got '%s'", text.Data())
		}
		if newText.Data() != "" {
			t.Errorf("Expected new text empty, got '%s'", newText.Data())
		}
	})

	t.Run("Text.splitText with parent insertion", func(t *testing.T) {
		doc := NewDocument()
		parent := NewElement("div", doc)
		text := NewText("Hello World", doc)
		parent.AppendChild(text)

		// Split text and verify DOM tree structure
		newText := text.SplitText(6)

		// Check parent-child relationships
		if text.ParentNode() != parent {
			t.Errorf("Expected original text to remain in parent")
		}
		if newText.ParentNode() != parent {
			t.Errorf("Expected new text to be inserted in parent")
		}

		// Check sibling relationships
		if text.NextSibling() != newText {
			t.Errorf("Expected new text to be next sibling of original")
		}
		if newText.PreviousSibling() != text {
			t.Errorf("Expected original text to be previous sibling of new")
		}

		// Check parent's children
		children := parent.ChildNodes()
		if children.Length() != 2 {
			t.Errorf("Expected parent to have 2 children, got %d", children.Length())
		}
		if children.Item(0) != text {
			t.Errorf("Expected first child to be original text")
		}
		if children.Item(1) != newText {
			t.Errorf("Expected second child to be new text")
		}
	})

	t.Run("Text.splitText with existing siblings", func(t *testing.T) {
		doc := NewDocument()
		parent := NewElement("div", doc)

		// Create DOM structure: <div>before[text]after</div>
		beforeElement := NewElement("span", doc)
		beforeElement.SetTextContent("before")
		text := NewText("Hello World", doc)
		afterElement := NewElement("span", doc)
		afterElement.SetTextContent("after")

		parent.AppendChild(beforeElement)
		parent.AppendChild(text)
		parent.AppendChild(afterElement)

		// Split text
		newText := text.SplitText(6)

		// Verify structure: <div>before[text][newText]after</div>
		children := parent.ChildNodes()
		if children.Length() != 4 {
			t.Errorf("Expected parent to have 4 children, got %d", children.Length())
		}
		if children.Item(0) != beforeElement {
			t.Errorf("Expected first child to be before element")
		}
		if children.Item(1) != text {
			t.Errorf("Expected second child to be original text")
		}
		if children.Item(2) != newText {
			t.Errorf("Expected third child to be new text")
		}
		if children.Item(3) != afterElement {
			t.Errorf("Expected fourth child to be after element")
		}
	})

	t.Run("Text.wholeText property compliance", func(t *testing.T) {
		doc := NewDocument()
		parent := NewElement("div", doc)

		// Test single text node
		text1 := NewText("Hello", doc)
		parent.AppendChild(text1)

		if text1.WholeText() != "Hello" {
			t.Errorf("Expected wholeText 'Hello', got '%s'", text1.WholeText())
		}

		// Test multiple contiguous text nodes
		text2 := NewText(" ", doc)
		text3 := NewText("World", doc)
		parent.AppendChild(text2)
		parent.AppendChild(text3)

		// All text nodes should return the same wholeText
		expectedWholeText := "Hello World"
		if text1.WholeText() != expectedWholeText {
			t.Errorf("Expected text1 wholeText '%s', got '%s'", expectedWholeText, text1.WholeText())
		}
		if text2.WholeText() != expectedWholeText {
			t.Errorf("Expected text2 wholeText '%s', got '%s'", expectedWholeText, text2.WholeText())
		}
		if text3.WholeText() != expectedWholeText {
			t.Errorf("Expected text3 wholeText '%s', got '%s'", expectedWholeText, text3.WholeText())
		}
	})

	t.Run("Text.wholeText with mixed node types", func(t *testing.T) {
		doc := NewDocument()
		parent := NewElement("div", doc)

		// Create structure: text1 + element + text2 + text3
		text1 := NewText("Hello", doc)
		element := NewElement("span", doc)
		text2 := NewText("Beautiful", doc)
		text3 := NewText("World", doc)

		parent.AppendChild(text1)
		parent.AppendChild(element)
		parent.AppendChild(text2)
		parent.AppendChild(text3)

		// text1 should only include itself (no contiguous text nodes)
		if text1.WholeText() != "Hello" {
			t.Errorf("Expected text1 wholeText 'Hello', got '%s'", text1.WholeText())
		}

		// text2 and text3 should be contiguous
		expectedWholeText := "BeautifulWorld"
		if text2.WholeText() != expectedWholeText {
			t.Errorf("Expected text2 wholeText '%s', got '%s'", expectedWholeText, text2.WholeText())
		}
		if text3.WholeText() != expectedWholeText {
			t.Errorf("Expected text3 wholeText '%s', got '%s'", expectedWholeText, text3.WholeText())
		}
	})

	t.Run("Text.wholeText with complex DOM structure", func(t *testing.T) {
		doc := NewDocument()
		parent := NewElement("div", doc)

		// Create structure: text1 + text2 + comment + text3 + element + text4 + text5
		text1 := NewText("Start", doc)
		text2 := NewText("Middle", doc)
		comment := NewComment("comment", doc)
		text3 := NewText("End", doc)
		element := NewElement("br", doc)
		text4 := NewText("After", doc)
		text5 := NewText("Element", doc)

		parent.AppendChild(text1)
		parent.AppendChild(text2)
		parent.AppendChild(comment)
		parent.AppendChild(text3)
		parent.AppendChild(element)
		parent.AppendChild(text4)
		parent.AppendChild(text5)

		// text1 and text2 should be contiguous (comment interrupts)
		expectedWholeText1 := "StartMiddle"
		if text1.WholeText() != expectedWholeText1 {
			t.Errorf("Expected text1 wholeText '%s', got '%s'", expectedWholeText1, text1.WholeText())
		}
		if text2.WholeText() != expectedWholeText1 {
			t.Errorf("Expected text2 wholeText '%s', got '%s'", expectedWholeText1, text2.WholeText())
		}

		// text3 should be alone (comment before, element after)
		if text3.WholeText() != "End" {
			t.Errorf("Expected text3 wholeText 'End', got '%s'", text3.WholeText())
		}

		// text4 and text5 should be contiguous
		expectedWholeText2 := "AfterElement"
		if text4.WholeText() != expectedWholeText2 {
			t.Errorf("Expected text4 wholeText '%s', got '%s'", expectedWholeText2, text4.WholeText())
		}
		if text5.WholeText() != expectedWholeText2 {
			t.Errorf("Expected text5 wholeText '%s', got '%s'", expectedWholeText2, text5.WholeText())
		}
	})
}

// TestTextInterface_CharacterDataInheritance tests that Text properly inherits from CharacterData
func TestTextInterface_CharacterDataInheritance(t *testing.T) {
	doc := NewDocument()
	text := NewText("Hello World", doc)

	t.Run("CharacterData methods work correctly", func(t *testing.T) {
		// Test Data() and SetData()
		if text.Data() != "Hello World" {
			t.Errorf("Expected data 'Hello World', got '%s'", text.Data())
		}

		text.SetData("Updated")
		if text.Data() != "Updated" {
			t.Errorf("Expected data 'Updated', got '%s'", text.Data())
		}

		// Test Length()
		if text.Length() != 7 {
			t.Errorf("Expected length 7, got %d", text.Length())
		}

		// Test SubstringData()
		substr := text.SubstringData(0, 3)
		if substr != "Upd" {
			t.Errorf("Expected substring 'Upd', got '%s'", substr)
		}

		// Test AppendData()
		text.AppendData(" Data")
		if text.Data() != "Updated Data" {
			t.Errorf("Expected data 'Updated Data', got '%s'", text.Data())
		}

		// Test InsertData()
		text.InsertData(7, " More")
		if text.Data() != "Updated More Data" {
			t.Errorf("Expected data 'Updated More Data', got '%s'", text.Data())
		}

		// Test ReplaceData()
		text.ReplaceData(8, 4, "Extra")
		if text.Data() != "Updated Extra Data" {
			t.Errorf("Expected data 'Updated Extra Data', got '%s'", text.Data())
		}

		// Test DeleteData()
		text.DeleteData(7, 6)
		if text.Data() != "Updated Data" {
			t.Errorf("Expected data 'Updated Data', got '%s'", text.Data())
		}
	})

	t.Run("CharacterData spec-compliant algorithms", func(t *testing.T) {
		text.SetData("Test Data")

		// Test error handling with negative offsets (should be handled gracefully)
		text.InsertData(-1, "Before")
		if text.Data() != "BeforeTest Data" {
			t.Errorf("Expected data 'BeforeTest Data', got '%s'", text.Data())
		}

		text.SetData("Test Data")
		text.DeleteData(-1, 4)
		if text.Data() != " Data" {
			t.Errorf("Expected data ' Data', got '%s'", text.Data())
		}

		text.SetData("Test Data")
		text.ReplaceData(-1, 4, "New")
		if text.Data() != "New Data" {
			t.Errorf("Expected data 'New Data', got '%s'", text.Data())
		}

		// Test beyond-length offsets
		text.SetData("Short")
		text.InsertData(100, " Extended")
		if text.Data() != "Short Extended" {
			t.Errorf("Expected data 'Short Extended', got '%s'", text.Data())
		}
	})
}

// TestTextInterface_SpecificationExamples tests examples from WHATWG DOM specification
func TestTextInterface_SpecificationExamples(t *testing.T) {
	t.Run("splitText example", func(t *testing.T) {
		// Example: splitting "foobar" at offset 3
		doc := NewDocument()
		text := NewText("foobar", doc)

		newText := text.SplitText(3)

		if text.Data() != "foo" {
			t.Errorf("Expected original text 'foo', got '%s'", text.Data())
		}
		if newText.Data() != "bar" {
			t.Errorf("Expected new text 'bar', got '%s'", newText.Data())
		}
	})

	t.Run("wholeText example", func(t *testing.T) {
		// Example: multiple contiguous text nodes
		doc := NewDocument()
		parent := NewElement("p", doc)

		text1 := NewText("Hello ", doc)
		text2 := NewText("beautiful ", doc)
		text3 := NewText("world!", doc)

		parent.AppendChild(text1)
		parent.AppendChild(text2)
		parent.AppendChild(text3)

		expectedWholeText := "Hello beautiful world!"
		if text1.WholeText() != expectedWholeText {
			t.Errorf("Expected wholeText '%s', got '%s'", expectedWholeText, text1.WholeText())
		}
		if text2.WholeText() != expectedWholeText {
			t.Errorf("Expected wholeText '%s', got '%s'", expectedWholeText, text2.WholeText())
		}
		if text3.WholeText() != expectedWholeText {
			t.Errorf("Expected wholeText '%s', got '%s'", expectedWholeText, text3.WholeText())
		}
	})

	t.Run("exclusive Text node definition", func(t *testing.T) {
		// "An exclusive Text node is a Text node that is not a CDATASection node"
		doc := NewDocument()
		text := NewText("text", doc)
		cdata := NewCDATASection("cdata", doc)

		// Text should be exclusive (not a CDATASection)
		if text.NodeType() != TextNode {
			t.Errorf("Expected Text node type, got %d", text.NodeType())
		}

		// CDATASection should not be exclusive Text (it's a CDATASection)
		if cdata.NodeType() != CDataSectionNode {
			t.Errorf("Expected CDATASection node type, got %d", cdata.NodeType())
		}

		// Both should be CharacterData
		if _, ok := interface{}(text).(CharacterData); !ok {
			t.Errorf("Text should implement CharacterData")
		}
		if _, ok := interface{}(cdata).(CharacterData); !ok {
			t.Errorf("CDATASection should implement CharacterData")
		}
	})
}
