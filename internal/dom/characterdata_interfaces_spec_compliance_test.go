package dom

import (
	"testing"
)

// TestCDATASectionInterface_SpecCompliance tests WHATWG DOM Section 4.12 CDATASection interface compliance
func TestCDATASectionInterface_SpecCompliance(t *testing.T) {
	t.Run("CDATASection inherits from Text", func(t *testing.T) {
		doc := NewDocument()
		cdata := NewCDATASection("test data", doc)

		// CDATASection should be a Text node in terms of interface
		if _, ok := interface{}(cdata).(*Text); !ok {
			// Check if it embeds Text (which it should)
			if _, ok := interface{}(cdata).(interface{ SplitText(int) *Text }); !ok {
				t.Errorf("CDATASection should inherit from Text interface")
			}
		}

		// But NodeType should be CDATASection, not Text
		if cdata.NodeType() != CDataSectionNode {
			t.Errorf("Expected CDATASection NodeType %d, got %d", CDataSectionNode, cdata.NodeType())
		}

		// Should have proper node name
		if cdata.NodeName() != "#cdata-section" {
			t.Errorf("Expected CDATASection NodeName '#cdata-section', got '%s'", cdata.NodeName())
		}

		// Should have CharacterData interface methods
		if cdata.Data() != "test data" {
			t.Errorf("Expected data 'test data', got '%s'", cdata.Data())
		}
	})

	t.Run("CDATASection has Text interface methods", func(t *testing.T) {
		doc := NewDocument()
		cdata := NewCDATASection("Hello World", doc)

		// Should have SplitText method from Text interface
		newCDATA := cdata.SplitText(6)
		if cdata.Data() != "Hello " {
			t.Errorf("Expected original CDATA 'Hello ', got '%s'", cdata.Data())
		}
		if newCDATA.Data() != "World" {
			t.Errorf("Expected new CDATA 'World', got '%s'", newCDATA.Data())
		}

		// New node should also be CDATASection
		if newCDATA.NodeType() != CDataSectionNode {
			t.Errorf("Expected split result to be CDATASection, got NodeType %d", newCDATA.NodeType())
		}
	})

	t.Run("CDATASection WholeText behavior", func(t *testing.T) {
		doc := NewDocument()
		parent := NewElement("div", doc)

		// Create contiguous text nodes including CDATA
		text1 := NewText("Hello ", doc)
		cdata := NewCDATASection("CDATA", doc)
		text2 := NewText(" World", doc)

		parent.AppendChild(text1)
		parent.AppendChild(cdata)
		parent.AppendChild(text2)

		// WholeText should concatenate all contiguous text nodes
		expectedWholeText := "Hello CDATA World"
		if cdata.WholeText() != expectedWholeText {
			t.Errorf("Expected WholeText '%s', got '%s'", expectedWholeText, cdata.WholeText())
		}
	})

	t.Run("CDATASection exclusive text node behavior", func(t *testing.T) {
		// Per Text interface spec: "An exclusive Text node is a Text node that is not a CDATASection node"
		doc := NewDocument()
		text := NewText("text", doc)
		cdata := NewCDATASection("cdata", doc)

		// Text should be exclusive (TextNode type)
		if text.NodeType() != TextNode {
			t.Errorf("Text node should have TextNode type, got %d", text.NodeType())
		}

		// CDATASection should NOT be exclusive Text (CDataSectionNode type)
		if cdata.NodeType() != CDataSectionNode {
			t.Errorf("CDATASection should have CDataSectionNode type, got %d", cdata.NodeType())
		}

		// Both should implement CharacterData
		if _, ok := interface{}(text).(CharacterData); !ok {
			t.Errorf("Text should implement CharacterData")
		}
		if _, ok := interface{}(cdata).(CharacterData); !ok {
			t.Errorf("CDATASection should implement CharacterData")
		}
	})
}

// TestProcessingInstructionInterface_SpecCompliance tests WHATWG DOM Section 4.13 ProcessingInstruction interface compliance
func TestProcessingInstructionInterface_SpecCompliance(t *testing.T) {
	t.Run("ProcessingInstruction inherits from CharacterData", func(t *testing.T) {
		doc := NewDocument()
		pi := NewProcessingInstruction("xml-stylesheet", "type=\"text/css\" href=\"style.css\"", doc)

		// Should implement CharacterData interface
		if _, ok := interface{}(pi).(CharacterData); !ok {
			t.Errorf("ProcessingInstruction should implement CharacterData interface")
		}

		// Should have proper NodeType
		if pi.NodeType() != ProcessingInstructionNode {
			t.Errorf("Expected ProcessingInstruction NodeType %d, got %d", ProcessingInstructionNode, pi.NodeType())
		}

		// NodeName should be the target
		if pi.NodeName() != "xml-stylesheet" {
			t.Errorf("Expected NodeName 'xml-stylesheet', got '%s'", pi.NodeName())
		}

		// Data should be accessible via CharacterData interface
		if pi.Data() != "type=\"text/css\" href=\"style.css\"" {
			t.Errorf("Expected data 'type=\"text/css\" href=\"style.css\"', got '%s'", pi.Data())
		}
	})

	t.Run("ProcessingInstruction target getter compliance", func(t *testing.T) {
		doc := NewDocument()
		target := "xml-stylesheet"
		data := "type=\"text/css\" href=\"style.css\""
		pi := NewProcessingInstruction(target, data, doc)

		// Test target getter per specification: "The target getter steps are to return this's target"
		if pi.Target() != target {
			t.Errorf("Expected target '%s', got '%s'", target, pi.Target())
		}

		// Target should remain constant
		pi.SetData("new data")
		if pi.Target() != target {
			t.Errorf("Target should remain constant after data change, got '%s'", pi.Target())
		}
	})

	t.Run("ProcessingInstruction target storage", func(t *testing.T) {
		doc := NewDocument()

		// Test various target values
		testCases := []struct {
			target string
			data   string
		}{
			{"xml", "version=\"1.0\""},
			{"xml-stylesheet", "type=\"text/css\" href=\"style.css\""},
			{"target", "data"},
			{"", ""}, // Empty target should work
		}

		for _, tc := range testCases {
			pi := NewProcessingInstruction(tc.target, tc.data, doc)

			if pi.Target() != tc.target {
				t.Errorf("Expected target '%s', got '%s'", tc.target, pi.Target())
			}

			if pi.Data() != tc.data {
				t.Errorf("Expected data '%s', got '%s'", tc.data, pi.Data())
			}

			// NodeName should match target
			if pi.NodeName() != tc.target {
				t.Errorf("Expected NodeName to match target '%s', got '%s'", tc.target, pi.NodeName())
			}
		}
	})

	t.Run("ProcessingInstruction CharacterData methods", func(t *testing.T) {
		doc := NewDocument()
		pi := NewProcessingInstruction("test", "initial data", doc)

		// Test CharacterData interface methods
		if pi.Length() != 12 {
			t.Errorf("Expected length 12, got %d", pi.Length())
		}

		// Test SubstringData
		substr := pi.SubstringData(0, 7)
		if substr != "initial" {
			t.Errorf("Expected substring 'initial', got '%s'", substr)
		}

		// Test AppendData
		pi.AppendData(" appended")
		if pi.Data() != "initial data appended" {
			t.Errorf("Expected 'initial data appended', got '%s'", pi.Data())
		}

		// Target should remain unchanged
		if pi.Target() != "test" {
			t.Errorf("Target should remain 'test', got '%s'", pi.Target())
		}
	})
}

// TestCommentInterface_SpecCompliance tests WHATWG DOM Section 4.14 Comment interface compliance
func TestCommentInterface_SpecCompliance(t *testing.T) {
	t.Run("Comment constructor compliance", func(t *testing.T) {
		doc := NewDocument()

		// Test constructor per specification steps:
		// "The new Comment(data) constructor steps are to set this's data to data
		// and this's node document to current global object's associated Document"

		data := "This is a comment"
		comment := NewComment(data, doc)

		// Step 1: data should be set
		if comment.Data() != data {
			t.Errorf("Expected data '%s', got '%s'", data, comment.Data())
		}

		// Step 2: node document should be set
		if comment.OwnerDocument() != doc {
			t.Errorf("Expected owner document to be set correctly")
		}

		// Should have proper NodeType
		if comment.NodeType() != CommentNode {
			t.Errorf("Expected Comment NodeType %d, got %d", CommentNode, comment.NodeType())
		}

		// Should have proper NodeName
		if comment.NodeName() != "#comment" {
			t.Errorf("Expected NodeName '#comment', got '%s'", comment.NodeName())
		}
	})

	t.Run("Comment inherits from CharacterData", func(t *testing.T) {
		doc := NewDocument()
		comment := NewComment("test comment", doc)

		// Should implement CharacterData interface
		if _, ok := interface{}(comment).(CharacterData); !ok {
			t.Errorf("Comment should implement CharacterData interface")
		}

		// Should have CharacterData methods working
		if comment.Length() != 12 {
			t.Errorf("Expected length 12, got %d", comment.Length())
		}

		// Test data manipulation
		comment.SetData("new comment data")
		if comment.Data() != "new comment data" {
			t.Errorf("Expected 'new comment data', got '%s'", comment.Data())
		}
	})

	t.Run("Comment constructor with empty data", func(t *testing.T) {
		doc := NewDocument()

		// Test with empty data (default parameter behavior)
		comment := NewComment("", doc)

		if comment.Data() != "" {
			t.Errorf("Expected empty data, got '%s'", comment.Data())
		}

		if comment.OwnerDocument() != doc {
			t.Errorf("Expected owner document to be set correctly")
		}
	})

	t.Run("Comment CharacterData operations", func(t *testing.T) {
		doc := NewDocument()
		comment := NewComment("Original comment", doc)

		// Test various CharacterData operations
		comment.AppendData(" - modified")
		if comment.Data() != "Original comment - modified" {
			t.Errorf("Expected 'Original comment - modified', got '%s'", comment.Data())
		}

		comment.InsertData(8, " new")
		if comment.Data() != "Original new comment - modified" {
			t.Errorf("Expected 'Original new comment - modified', got '%s'", comment.Data())
		}

		comment.ReplaceData(0, 8, "Updated")
		if comment.Data() != "Updated new comment - modified" {
			t.Errorf("Expected 'Updated new comment - modified', got '%s'", comment.Data())
		}

		comment.DeleteData(7, 4)
		if comment.Data() != "Updated comment - modified" {
			t.Errorf("Expected 'Updated comment - modified', got '%s'", comment.Data())
		}
	})

	t.Run("Comment specification examples", func(t *testing.T) {
		// Test examples that might appear in the specification
		doc := NewDocument()

		// HTML-style comment
		htmlComment := NewComment(" This is an HTML comment ", doc)
		if htmlComment.Data() != " This is an HTML comment " {
			t.Errorf("HTML comment data mismatch")
		}

		// Empty comment
		emptyComment := NewComment("", doc)
		if emptyComment.Data() != "" {
			t.Errorf("Empty comment should have empty data")
		}

		// Comment with special characters
		specialComment := NewComment("Comment with <>&\"' characters", doc)
		if specialComment.Data() != "Comment with <>&\"' characters" {
			t.Errorf("Special characters should be preserved in comment data")
		}
	})
}

// TestCharacterDataInterfacesIntegration tests how all three interfaces work together
func TestCharacterDataInterfacesIntegration(t *testing.T) {
	t.Run("All interfaces implement CharacterData", func(t *testing.T) {
		doc := NewDocument()

		// Create instances of all three
		text := NewText("text data", doc)
		cdata := NewCDATASection("cdata data", doc)
		comment := NewComment("comment data", doc)
		pi := NewProcessingInstruction("target", "pi data", doc)

		// All should implement CharacterData
		interfaces := []struct {
			name string
			node CharacterData
		}{
			{"Text", text},
			{"CDATASection", cdata},
			{"Comment", comment},
			{"ProcessingInstruction", pi},
		}

		for _, iface := range interfaces {
			if iface.node.Data() == "" {
				t.Errorf("%s should have non-empty data", iface.name)
			}
			if iface.node.Length() == 0 {
				t.Errorf("%s should have non-zero length", iface.name)
			}
		}
	})

	t.Run("Node type distinctions", func(t *testing.T) {
		doc := NewDocument()

		text := NewText("data", doc)
		cdata := NewCDATASection("data", doc)
		comment := NewComment("data", doc)
		pi := NewProcessingInstruction("target", "data", doc)

		// Each should have distinct node types
		if text.NodeType() != TextNode {
			t.Errorf("Text should have TextNode type")
		}
		if cdata.NodeType() != CDataSectionNode {
			t.Errorf("CDATASection should have CDataSectionNode type")
		}
		if comment.NodeType() != CommentNode {
			t.Errorf("Comment should have CommentNode type")
		}
		if pi.NodeType() != ProcessingInstructionNode {
			t.Errorf("ProcessingInstruction should have ProcessingInstructionNode type")
		}

		// All types should be different
		types := []NodeType{text.NodeType(), cdata.NodeType(), comment.NodeType(), pi.NodeType()}
		for i := 0; i < len(types); i++ {
			for j := i + 1; j < len(types); j++ {
				if types[i] == types[j] {
					t.Errorf("Node types should be distinct: %d == %d", types[i], types[j])
				}
			}
		}
	})

	t.Run("Mixed DOM tree with all character data types", func(t *testing.T) {
		doc := NewDocument()
		root := NewElement("root", doc)

		// Create a mixed tree
		text := NewText("text content", doc)
		cdata := NewCDATASection("cdata content", doc)
		comment := NewComment("comment content", doc)
		pi := NewProcessingInstruction("xml", "version=\"1.0\"", doc)

		root.AppendChild(text)
		root.AppendChild(cdata)
		root.AppendChild(comment)
		root.AppendChild(pi)

		// Verify tree structure
		children := root.ChildNodes()
		if children.Length() != 4 {
			t.Errorf("Expected 4 children, got %d", children.Length())
		}

		// Verify each child type and content
		if children.Item(0).NodeType() != TextNode {
			t.Errorf("First child should be Text")
		}
		if children.Item(1).NodeType() != CDataSectionNode {
			t.Errorf("Second child should be CDATASection")
		}
		if children.Item(2).NodeType() != CommentNode {
			t.Errorf("Third child should be Comment")
		}
		if children.Item(3).NodeType() != ProcessingInstructionNode {
			t.Errorf("Fourth child should be ProcessingInstruction")
		}
	})
}
