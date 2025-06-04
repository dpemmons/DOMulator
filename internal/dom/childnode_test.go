package dom

import (
	"testing"
)

func TestElement_ChildNodeMethods(t *testing.T) {
	doc := NewDocument()
	parent := NewElement("div", doc)
	elem1 := NewElement("span", doc)
	elem2 := NewElement("p", doc)
	text1 := NewText("text1", doc)

	// Set up parent with children
	parent.AppendChild(elem1)
	parent.AppendChild(text1)
	parent.AppendChild(elem2)

	t.Run("Before with single node", func(t *testing.T) {
		newElem := NewElement("strong", doc)
		err := elem2.Before(newElem)
		if err != nil {
			t.Errorf("Before() error = %v", err)
		}

		children := parent.ChildNodes()
		if children.Length() != 4 {
			t.Errorf("Expected 4 children, got %d", children.Length())
		}
		if children.Item(2) != newElem {
			t.Errorf("Expected newElem at index 2")
		}
		if children.Item(3) != elem2 {
			t.Errorf("Expected elem2 at index 3")
		}
	})

	t.Run("After with multiple nodes", func(t *testing.T) {
		newText := NewText("new text", doc)
		newElem := NewElement("em", doc)
		err := text1.After(newText, newElem)
		if err != nil {
			t.Errorf("After() error = %v", err)
		}

		children := parent.ChildNodes()
		if children.Length() != 6 {
			t.Errorf("Expected 6 children, got %d", children.Length())
		}
		// text1 should be at index 1, new nodes at 2 and 3
		if children.Item(1) != text1 {
			t.Errorf("Expected text1 at index 1")
		}
		if children.Item(2) != newText {
			t.Errorf("Expected newText at index 2")
		}
		if children.Item(3) != newElem {
			t.Errorf("Expected newElem at index 3")
		}
	})

	t.Run("ReplaceWith with string", func(t *testing.T) {
		children := parent.ChildNodes()
		oldChildCount := children.Length()
		err := elem1.ReplaceWith("replacement text")
		if err != nil {
			t.Errorf("ReplaceWith() error = %v", err)
		}

		children = parent.ChildNodes()
		if children.Length() != oldChildCount {
			t.Errorf("Expected same number of children after replace")
		}
		// First child should now be a text node
		if children.Item(0).NodeType() != TextNode {
			t.Errorf("Expected first child to be text node")
		}
		if children.Item(0).NodeValue() != "replacement text" {
			t.Errorf("Expected replacement text, got %s", children.Item(0).NodeValue())
		}
	})

	t.Run("Remove", func(t *testing.T) {
		children := parent.ChildNodes()
		childrenBefore := children.Length()
		targetNode := children.Item(2) // Get a node to remove

		if elem, ok := targetNode.(*Element); ok {
			err := elem.Remove()
			if err != nil {
				t.Errorf("Remove() error = %v", err)
			}
		} else if text, ok := targetNode.(*Text); ok {
			err := text.Remove()
			if err != nil {
				t.Errorf("Remove() error = %v", err)
			}
		}

		children = parent.ChildNodes()
		childrenAfter := children.Length()
		if childrenAfter != childrenBefore-1 {
			t.Errorf("Expected one less child after remove, got %d -> %d", childrenBefore, childrenAfter)
		}
	})
}

func TestText_ChildNodeMethods(t *testing.T) {
	doc := NewDocument()
	parent := NewElement("div", doc)
	text1 := NewText("first", doc)
	text2 := NewText("second", doc)
	text3 := NewText("third", doc)

	parent.AppendChild(text1)
	parent.AppendChild(text2)
	parent.AppendChild(text3)

	t.Run("Before with mixed nodes", func(t *testing.T) {
		newElem := NewElement("span", doc)
		err := text2.Before("prefix", newElem)
		if err != nil {
			t.Errorf("Before() error = %v", err)
		}

		children := parent.ChildNodes()
		if children.Length() != 5 {
			t.Errorf("Expected 5 children, got %d", children.Length())
		}
		// Should have: text1, "prefix", newElem, text2, text3
		if children.Item(1).NodeValue() != "prefix" {
			t.Errorf("Expected 'prefix' at index 1, got %s", children.Item(1).NodeValue())
		}
		if children.Item(2) != newElem {
			t.Errorf("Expected newElem at index 2")
		}
		if children.Item(3) != text2 {
			t.Errorf("Expected text2 at index 3")
		}
	})

	t.Run("After with no parent", func(t *testing.T) {
		orphanText := NewText("orphan", doc)
		err := orphanText.After("after")
		if err != nil {
			t.Errorf("After() should not error for orphan node, got %v", err)
		}
		// Should be a no-op
	})

	t.Run("ReplaceWith empty nodes", func(t *testing.T) {
		children := parent.ChildNodes()
		childrenBefore := children.Length()
		err := text3.ReplaceWith()
		if err != nil {
			t.Errorf("ReplaceWith() error = %v", err)
		}

		children = parent.ChildNodes()
		childrenAfter := children.Length()
		if childrenAfter != childrenBefore-1 {
			t.Errorf("Expected one less child after empty replace, got %d -> %d", childrenBefore, childrenAfter)
		}
	})
}

func TestComment_ChildNodeMethods(t *testing.T) {
	doc := NewDocument()
	parent := NewElement("div", doc)
	comment1 := NewComment("comment 1", doc)
	comment2 := NewComment("comment 2", doc)
	elem := NewElement("span", doc)

	parent.AppendChild(comment1)
	parent.AppendChild(elem)
	parent.AppendChild(comment2)

	t.Run("Before", func(t *testing.T) {
		newComment := NewComment("new comment", doc)
		err := comment2.Before(newComment)
		if err != nil {
			t.Errorf("Before() error = %v", err)
		}

		children := parent.ChildNodes()
		if children.Length() != 4 {
			t.Errorf("Expected 4 children, got %d", children.Length())
		}
		if children.Item(2) != newComment {
			t.Errorf("Expected newComment at index 2")
		}
	})

	t.Run("After", func(t *testing.T) {
		textNode := NewText("after comment", doc)
		err := comment1.After(textNode)
		if err != nil {
			t.Errorf("After() error = %v", err)
		}

		children := parent.ChildNodes()
		if children.Length() != 5 {
			t.Errorf("Expected 5 children, got %d", children.Length())
		}
		if children.Item(1) != textNode {
			t.Errorf("Expected textNode at index 1")
		}
	})

	t.Run("Remove", func(t *testing.T) {
		children := parent.ChildNodes()
		childrenBefore := children.Length()
		err := comment1.Remove()
		if err != nil {
			t.Errorf("Remove() error = %v", err)
		}

		children = parent.ChildNodes()
		childrenAfter := children.Length()
		if childrenAfter != childrenBefore-1 {
			t.Errorf("Expected one less child after remove")
		}
	})
}

func TestDocumentType_ChildNodeMethods(t *testing.T) {
	doc := NewDocument()
	doctype := NewDocumentType("html", "", "", doc)
	html := NewElement("html", doc)

	// Add doctype and html to document
	doc.AppendChild(doctype)
	doc.AppendChild(html)

	t.Run("Before", func(t *testing.T) {
		comment := NewComment("before html", doc)
		err := doctype.Before(comment)
		if err != nil {
			t.Errorf("Before() error = %v", err)
		}

		children := doc.ChildNodes()
		if children.Length() != 3 {
			t.Errorf("Expected 3 children, got %d", children.Length())
		}
		if children.Item(0) != comment {
			t.Errorf("Expected comment at index 0")
		}
		if children.Item(1) != doctype {
			t.Errorf("Expected doctype at index 1")
		}
	})

	t.Run("After", func(t *testing.T) {
		pi := NewProcessingInstruction("xml-stylesheet", "type='text/css' href='style.css'", doc)
		err := doctype.After(pi)
		if err != nil {
			t.Errorf("After() error = %v", err)
		}

		children := doc.ChildNodes()
		if children.Length() != 4 {
			t.Errorf("Expected 4 children, got %d", children.Length())
		}
		if children.Item(2) != pi {
			t.Errorf("Expected PI at index 2")
		}
	})

	t.Run("ReplaceWith", func(t *testing.T) {
		newDoctype := NewDocumentType("xhtml", "-//W3C//DTD XHTML 1.0//EN", "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd", doc)
		err := doctype.ReplaceWith(newDoctype)
		if err != nil {
			t.Errorf("ReplaceWith() error = %v", err)
		}

		children := doc.ChildNodes()
		// Should still have same number of children
		found := false
		for i := 0; i < children.Length(); i++ {
			child := children.Item(i)
			if child == newDoctype {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected to find new doctype in document")
		}
		if newDoctype.Name() != "xhtml" {
			t.Errorf("Expected new doctype name 'xhtml', got '%s'", newDoctype.Name())
		}
	})
}

func TestProcessingInstruction_ChildNodeMethods(t *testing.T) {
	doc := NewDocument()
	parent := NewElement("root", doc)
	pi1 := NewProcessingInstruction("xml-stylesheet", "type='text/css'", doc)
	pi2 := NewProcessingInstruction("xml-model", "href='schema.rng'", doc)
	elem := NewElement("content", doc)

	parent.AppendChild(pi1)
	parent.AppendChild(elem)
	parent.AppendChild(pi2)

	t.Run("Before", func(t *testing.T) {
		newPI := NewProcessingInstruction("target", "data", doc)
		err := pi2.Before(newPI)
		if err != nil {
			t.Errorf("Before() error = %v", err)
		}

		children := parent.ChildNodes()
		if children.Length() != 4 {
			t.Errorf("Expected 4 children, got %d", children.Length())
		}
		if children.Item(2) != newPI {
			t.Errorf("Expected newPI at index 2")
		}
	})

	t.Run("After with string", func(t *testing.T) {
		err := pi1.After("text after PI")
		if err != nil {
			t.Errorf("After() error = %v", err)
		}

		children := parent.ChildNodes()
		if children.Length() != 5 {
			t.Errorf("Expected 5 children, got %d", children.Length())
		}
		if children.Item(1).NodeType() != TextNode {
			t.Errorf("Expected text node at index 1")
		}
		if children.Item(1).NodeValue() != "text after PI" {
			t.Errorf("Expected 'text after PI', got '%s'", children.Item(1).NodeValue())
		}
	})

	t.Run("Remove", func(t *testing.T) {
		children := parent.ChildNodes()
		childrenBefore := children.Length()
		err := pi1.Remove()
		if err != nil {
			t.Errorf("Remove() error = %v", err)
		}

		children = parent.ChildNodes()
		childrenAfter := children.Length()
		if childrenAfter != childrenBefore-1 {
			t.Errorf("Expected one less child after remove")
		}

		// Verify pi1 is no longer in the tree
		children = parent.ChildNodes()
		for i := 0; i < children.Length(); i++ {
			child := children.Item(i)
			if child == pi1 {
				t.Errorf("pi1 should have been removed from parent")
			}
		}
	})
}

func TestChildNode_ErrorHandling(t *testing.T) {
	doc := NewDocument()

	t.Run("Invalid node type", func(t *testing.T) {
		elem := NewElement("div", doc)
		parent := NewElement("parent", doc)
		parent.AppendChild(elem)

		err := elem.Before(123) // Invalid type
		if err == nil {
			t.Errorf("Expected error for invalid node type")
		}
	})

	t.Run("Methods on node without parent", func(t *testing.T) {
		orphan := NewElement("orphan", doc)

		// These should not error, just be no-ops
		err1 := orphan.Before("text")
		err2 := orphan.After("text")
		err3 := orphan.ReplaceWith("replacement")
		err4 := orphan.Remove()

		if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
			t.Errorf("Methods on orphan nodes should not error")
		}
	})
}

func TestChildNode_DocumentFragmentCreation(t *testing.T) {
	doc := NewDocument()
	parent := NewElement("div", doc)
	elem := NewElement("span", doc)
	parent.AppendChild(elem)

	t.Run("Multiple nodes create DocumentFragment", func(t *testing.T) {
		text1 := NewText("text1", doc)
		text2 := NewText("text2", doc)
		elem2 := NewElement("strong", doc)

		err := elem.Before(text1, text2, elem2)
		if err != nil {
			t.Errorf("Before() error = %v", err)
		}

		children := parent.ChildNodes()
		if children.Length() != 4 {
			t.Errorf("Expected 4 children, got %d", children.Length())
		}

		// Verify all nodes were inserted in correct order
		if children.Item(0) != text1 {
			t.Errorf("Expected text1 at index 0")
		}
		if children.Item(1) != text2 {
			t.Errorf("Expected text2 at index 1")
		}
		if children.Item(2) != elem2 {
			t.Errorf("Expected elem2 at index 2")
		}
		if children.Item(3) != elem {
			t.Errorf("Expected original elem at index 3")
		}
	})

	t.Run("Mixed strings and nodes", func(t *testing.T) {
		newElem := NewElement("em", doc)
		err := elem.After("prefix", newElem, "suffix")
		if err != nil {
			t.Errorf("After() error = %v", err)
		}

		children := parent.ChildNodes()
		// Should have text1, text2, elem2, elem, "prefix", newElem, "suffix"
		if children.Length() != 7 {
			t.Errorf("Expected 7 children, got %d", children.Length())
		}

		// Verify the last three are our new nodes
		if children.Item(4).NodeValue() != "prefix" {
			t.Errorf("Expected 'prefix' at index 4")
		}
		if children.Item(5) != newElem {
			t.Errorf("Expected newElem at index 5")
		}
		if children.Item(6).NodeValue() != "suffix" {
			t.Errorf("Expected 'suffix' at index 6")
		}
	})
}
