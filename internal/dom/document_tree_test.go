package dom

import (
	"testing"
)

// TestDocumentTreeSpecCompliance tests WHATWG DOM Section 4.2.1 "Document tree" concepts
func TestDocumentTreeSpecCompliance(t *testing.T) {
	t.Run("IsConnected uses shadow-including root", func(t *testing.T) {
		// Create document tree
		doc := NewDocument()
		body := doc.CreateElement("body")
		doc.AppendChild(body)

		// Create element with shadow root
		host := doc.CreateElement("div")
		body.AppendChild(host)

		// Test: element in document tree should be connected
		if !host.IsConnected() {
			t.Error("Element in document tree should be connected")
		}

		// Attach shadow root
		shadowRoot, err := host.AttachShadow(ShadowRootInit{Mode: "open"})
		if err != nil {
			t.Fatalf("Failed to attach shadow root: %v", err)
		}

		// Create element in shadow tree
		shadowElement := doc.CreateElement("span")
		shadowRoot.AppendChild(shadowElement)

		// Test: element in shadow tree should be connected (shadow-including root is document)
		if !shadowElement.IsConnected() {
			t.Error("Element in shadow tree should be connected - shadow-including root is document")
		}

		// Create orphaned element
		orphan := doc.CreateElement("orphan")

		// Test: orphaned element should not be connected
		if orphan.IsConnected() {
			t.Error("Orphaned element should not be connected")
		}

		// Test: shadow root itself should be connected (its host is in document)
		if !shadowRoot.IsConnected() {
			t.Error("Shadow root should be connected when host is in document")
		}

		// Remove host from document
		body.RemoveChild(host)

		// Test: host should no longer be connected
		if host.IsConnected() {
			t.Error("Host removed from document should not be connected")
		}

		// Test: shadow element should no longer be connected (shadow-including root is now orphaned host)
		if shadowElement.IsConnected() {
			t.Error("Element in orphaned shadow tree should not be connected")
		}

		// Test: shadow root should no longer be connected
		if shadowRoot.IsConnected() {
			t.Error("Shadow root of orphaned host should not be connected")
		}
	})

	t.Run("Document element concept", func(t *testing.T) {
		doc := NewDocument()

		// Test: new document has no document element
		if doc.DocumentElement() != nil {
			t.Error("New document should have no document element")
		}

		// Add a doctype first
		doctype := NewDocumentType("html", "", "", doc)
		doc.AppendChild(doctype)

		// Test: document with only doctype has no document element
		if doc.DocumentElement() != nil {
			t.Error("Document with only doctype should have no document element")
		}

		// Add the document element
		html := doc.CreateElement("html")
		doc.AppendChild(html)

		// Test: document element should be the html element
		if doc.DocumentElement() != html {
			t.Error("Document element should be the html element")
		}

		// Add some children to html
		head := doc.CreateElement("head")
		body := doc.CreateElement("body")
		html.AppendChild(head)
		html.AppendChild(body)

		// Test: document element should still be html (not its children)
		if doc.DocumentElement() != html {
			t.Error("Document element should remain the html element")
		}

		// Test: there can be only one document element
		html2 := doc.CreateElement("html2")
		defer func() {
			if r := recover(); r == nil {
				t.Error("Should not be able to add second element to document")
			}
		}()
		doc.AppendChild(html2) // Should panic due to hierarchy validation
	})

	t.Run("In a document tree concept", func(t *testing.T) {
		doc := NewDocument()
		html := doc.CreateElement("html")
		body := doc.CreateElement("body")
		div := doc.CreateElement("div")

		doc.AppendChild(html)
		html.AppendChild(body)
		body.AppendChild(div)

		// Test: all nodes should be in document tree
		if !html.IsConnected() {
			t.Error("Document element should be in document tree")
		}
		if !body.IsConnected() {
			t.Error("Body element should be in document tree")
		}
		if !div.IsConnected() {
			t.Error("Div element should be in document tree")
		}

		// Test: document itself has special behavior (no parent, but is connected)
		if doc.ParentNode() != nil {
			t.Error("Document should have no parent")
		}
		if !doc.IsConnected() {
			t.Error("Document should be connected to itself")
		}

		// Create shadow tree
		shadowRoot, err := div.AttachShadow(ShadowRootInit{Mode: "open"})
		if err != nil {
			t.Fatalf("Failed to attach shadow root: %v", err)
		}
		shadowSpan := doc.CreateElement("span")
		shadowRoot.AppendChild(shadowSpan)

		// Test: shadow tree nodes should also be in document tree (via shadow-including root)
		if !shadowSpan.IsConnected() {
			t.Error("Element in shadow tree should be in document tree via shadow-including root")
		}

		// Remove div from document tree
		body.RemoveChild(div)

		// Test: removed nodes should no longer be in document tree
		if div.IsConnected() {
			t.Error("Removed div should not be in document tree")
		}
		if shadowSpan.IsConnected() {
			t.Error("Element in orphaned shadow tree should not be in document tree")
		}
	})

	t.Run("Shadow-including root behavior", func(t *testing.T) {
		doc := NewDocument()
		host := doc.CreateElement("div")
		doc.AppendChild(host)

		// Test: regular root vs shadow-including root
		normalRoot := host.GetRootNode(nil)
		if normalRoot != doc {
			t.Error("Normal root should be document")
		}

		shadowIncludingRoot := host.GetRootNode(&GetRootNodeOptions{Composed: true})
		if shadowIncludingRoot != doc {
			t.Error("Shadow-including root should be document")
		}

		// Create shadow tree
		shadowRoot, err := host.AttachShadow(ShadowRootInit{Mode: "open"})
		if err != nil {
			t.Fatalf("Failed to attach shadow root: %v", err)
		}
		shadowElement := doc.CreateElement("div")
		shadowRoot.AppendChild(shadowElement)

		// Test: shadow element roots
		normalRoot = shadowElement.GetRootNode(nil)
		if normalRoot != shadowRoot {
			t.Error("Normal root of shadow element should be shadow root")
		}

		shadowIncludingRoot = shadowElement.GetRootNode(&GetRootNodeOptions{Composed: true})
		if shadowIncludingRoot != doc {
			t.Error("Shadow-including root of shadow element should be document")
		}

		// Test: IsConnected uses shadow-including root
		if !shadowElement.IsConnected() {
			t.Error("Shadow element should be connected via shadow-including root")
		}
	})
}
