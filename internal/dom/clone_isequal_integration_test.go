package dom

import (
	"testing"
)

// TestCloneAndIsEqualNode_Integration tests that cloned nodes are equal to their originals
func TestCloneAndIsEqualNode_Integration(t *testing.T) {
	doc := NewDocument()

	t.Run("Text node clone equality", func(t *testing.T) {
		original := doc.CreateTextNode("Hello World")
		cloned := original.CloneNode(false)

		if !original.IsEqualNode(cloned) {
			t.Error("Cloned text node should be equal to original")
		}
	})

	t.Run("Comment node clone equality", func(t *testing.T) {
		original := doc.CreateComment("This is a comment")
		cloned := original.CloneNode(false)

		if !original.IsEqualNode(cloned) {
			t.Error("Cloned comment node should be equal to original")
		}
	})

	t.Run("ProcessingInstruction clone equality", func(t *testing.T) {
		original, _ := doc.CreateProcessingInstruction("xml-stylesheet", "href='style.css'")
		cloned := original.CloneNode(false)

		if !original.IsEqualNode(cloned) {
			t.Error("Cloned ProcessingInstruction should be equal to original")
		}
	})

	t.Run("DocumentType clone equality", func(t *testing.T) {
		original := NewDocumentType("html", "-//W3C//DTD HTML 4.01//EN", "http://www.w3.org/TR/html4/strict.dtd", doc)
		cloned := original.CloneNode(false)

		if !original.IsEqualNode(cloned) {
			t.Error("Cloned DocumentType should be equal to original")
		}
	})

	t.Run("Attribute clone equality", func(t *testing.T) {
		original, _ := doc.CreateAttribute("class")
		original.SetValue("test-class")
		cloned := original.CloneNode(false)

		if !original.IsEqualNode(cloned) {
			t.Error("Cloned Attribute should be equal to original")
		}
	})

	t.Run("Simple element clone equality", func(t *testing.T) {
		original := doc.CreateElement("div")
		original.SetAttribute("id", "test")
		original.SetAttribute("class", "container")

		cloned := original.CloneNode(false)

		if !original.IsEqualNode(cloned) {
			t.Error("Cloned element should be equal to original")
		}
	})

	t.Run("Element with children shallow clone", func(t *testing.T) {
		parent := doc.CreateElement("div")
		child := doc.CreateElement("span")
		child.SetAttribute("class", "child")
		parent.AppendChild(child)

		// Shallow clone (subtree = false)
		cloned := parent.CloneNode(false)

		// Shallow clone should not include children
		if cloned.HasChildNodes() {
			t.Error("Shallow clone should not have children")
		}

		// But the structure without children should be equal
		emptyParent := doc.CreateElement("div")
		if !cloned.IsEqualNode(emptyParent) {
			t.Error("Shallow cloned element should be equal to empty element with same tag")
		}
	})

	t.Run("Element with children deep clone", func(t *testing.T) {
		parent := doc.CreateElement("div")
		parent.SetAttribute("class", "parent")

		child1 := doc.CreateElement("span")
		child1.SetAttribute("class", "child1")
		text1 := doc.CreateTextNode("Child 1 text")
		child1.AppendChild(text1)

		child2 := doc.CreateElement("p")
		child2.SetAttribute("id", "paragraph")
		text2 := doc.CreateTextNode("Paragraph text")
		child2.AppendChild(text2)

		parent.AppendChild(child1)
		parent.AppendChild(child2)

		// Deep clone (subtree = true)
		cloned := parent.CloneNode(true)

		if !parent.IsEqualNode(cloned) {
			t.Error("Deep cloned element should be equal to original")
		}

		// Verify children are also equal
		if parent.ChildNodes().Length() != cloned.ChildNodes().Length() {
			t.Error("Cloned element should have same number of children")
		}

		for i := 0; i < parent.ChildNodes().Length(); i++ {
			originalChild := parent.ChildNodes().Item(i)
			clonedChild := cloned.ChildNodes().Item(i)

			if !originalChild.IsEqualNode(clonedChild) {
				t.Errorf("Child %d should be equal between original and clone", i)
			}
		}
	})

	t.Run("Document clone equality", func(t *testing.T) {
		// Create a simple document structure
		doctype := NewDocumentType("html", "", "", doc)
		doc.AppendChild(doctype)

		root := doc.CreateElement("html")
		head := doc.CreateElement("head")
		title := doc.CreateElement("title")
		titleText := doc.CreateTextNode("Test Document")
		title.AppendChild(titleText)
		head.AppendChild(title)
		root.AppendChild(head)
		doc.AppendChild(root)

		// Clone the document
		cloned := doc.CloneNode(true)

		if !doc.IsEqualNode(cloned) {
			t.Error("Cloned document should be equal to original")
		}
	})

	t.Run("DocumentFragment clone equality", func(t *testing.T) {
		fragment := doc.CreateDocumentFragment()

		elem1 := doc.CreateElement("div")
		elem1.SetAttribute("class", "first")
		text1 := doc.CreateTextNode("First element")
		elem1.AppendChild(text1)

		elem2 := doc.CreateElement("span")
		elem2.SetAttribute("id", "second")
		text2 := doc.CreateTextNode("Second element")
		elem2.AppendChild(text2)

		fragment.AppendChild(elem1)
		fragment.AppendChild(elem2)

		cloned := fragment.CloneNode(true)

		if !fragment.IsEqualNode(cloned) {
			t.Error("Cloned DocumentFragment should be equal to original")
		}
	})

	t.Run("Complex nested structure clone equality", func(t *testing.T) {
		// Create a complex nested structure
		article := doc.CreateElement("article")
		article.SetAttribute("class", "post")
		article.SetAttribute("data-id", "123")

		header := doc.CreateElement("header")
		h1 := doc.CreateElement("h1")
		h1Text := doc.CreateTextNode("Article Title")
		h1.AppendChild(h1Text)
		header.AppendChild(h1)

		meta := doc.CreateElement("div")
		meta.SetAttribute("class", "meta")
		author := doc.CreateElement("span")
		author.SetAttribute("class", "author")
		authorText := doc.CreateTextNode("John Doe")
		author.AppendChild(authorText)
		meta.AppendChild(author)
		header.AppendChild(meta)

		content := doc.CreateElement("div")
		content.SetAttribute("class", "content")
		p1 := doc.CreateElement("p")
		p1Text := doc.CreateTextNode("First paragraph of the article.")
		p1.AppendChild(p1Text)

		p2 := doc.CreateElement("p")
		p2Text1 := doc.CreateTextNode("Second paragraph with ")
		strong := doc.CreateElement("strong")
		strongText := doc.CreateTextNode("bold text")
		strong.AppendChild(strongText)
		p2Text2 := doc.CreateTextNode(" in the middle.")
		p2.AppendChild(p2Text1)
		p2.AppendChild(strong)
		p2.AppendChild(p2Text2)

		content.AppendChild(p1)
		content.AppendChild(p2)

		article.AppendChild(header)
		article.AppendChild(content)

		// Clone the complex structure
		cloned := article.CloneNode(true)

		if !article.IsEqualNode(cloned) {
			t.Error("Cloned complex structure should be equal to original")
		}
	})

	t.Run("Namespace element clone equality", func(t *testing.T) {
		original, _ := doc.CreateElementNS("http://www.w3.org/2000/svg", "svg")
		original.SetAttribute("width", "100")
		original.SetAttribute("height", "100")

		circle, _ := doc.CreateElementNS("http://www.w3.org/2000/svg", "circle")
		circle.SetAttribute("cx", "50")
		circle.SetAttribute("cy", "50")
		circle.SetAttribute("r", "25")
		original.AppendChild(circle)

		cloned := original.CloneNode(true)

		if !original.IsEqualNode(cloned) {
			t.Error("Cloned namespaced element should be equal to original")
		}
	})
}

// TestCloneModificationDoesNotAffectOriginal ensures that modifying a clone doesn't affect the original
func TestCloneModificationDoesNotAffectOriginal(t *testing.T) {
	doc := NewDocument()

	t.Run("Modifying cloned element doesn't affect original", func(t *testing.T) {
		original := doc.CreateElement("div")
		original.SetAttribute("class", "original")

		child := doc.CreateElement("span")
		child.SetAttribute("id", "child")
		originalText := doc.CreateTextNode("Original text")
		child.AppendChild(originalText)
		original.AppendChild(child)

		// Clone the element
		cloned := original.CloneNode(true)

		// Verify they're initially equal
		if !original.IsEqualNode(cloned) {
			t.Error("Clone should initially be equal to original")
		}

		// Modify the clone
		clonedElement := cloned.(*Element)
		clonedElement.SetAttribute("class", "modified")
		clonedElement.SetAttribute("data-test", "added")

		// Modify the child of the clone
		clonedChild := cloned.FirstChild().(*Element)
		clonedChild.SetAttribute("id", "modified-child")

		// Modify text content in clone
		clonedText := clonedChild.FirstChild().(*Text)
		clonedText.SetData("Modified text")

		// Original should remain unchanged
		if original.GetAttribute("class") != "original" {
			t.Error("Original element class should not have changed")
		}

		if original.HasAttribute("data-test") {
			t.Error("Original should not have data-test attribute")
		}

		originalChild := original.FirstChild().(*Element)
		if originalChild.GetAttribute("id") != "child" {
			t.Error("Original child id should not have changed")
		}

		originalChildText := originalChild.FirstChild().(*Text)
		if originalChildText.Data() != "Original text" {
			t.Error("Original text content should not have changed")
		}

		// They should no longer be equal
		if original.IsEqualNode(cloned) {
			t.Error("Modified clone should not be equal to original")
		}
	})
}
