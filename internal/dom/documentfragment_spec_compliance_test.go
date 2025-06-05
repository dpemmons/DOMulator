package dom

import (
	"testing"
)

// TestDocumentFragmentSpecCompliance tests compliance with WHATWG DOM Section 4.7
// Interface DocumentFragment specification
func TestDocumentFragmentSpecCompliance(t *testing.T) {
	doc := NewDocument()

	t.Run("WHATWG DOM Section 4.7 - Interface DocumentFragment", func(t *testing.T) {
		t.Run("Constructor returns a new DocumentFragment node", func(t *testing.T) {
			// Test that NewDocumentFragment creates a proper DocumentFragment
			fragment := NewDocumentFragment(doc)

			// Should be a DocumentFragment node
			if fragment.NodeType() != DocumentFragmentNode {
				t.Errorf("Expected DocumentFragment node type %d, got %d", DocumentFragmentNode, fragment.NodeType())
			}

			// Node name should be "#document-fragment"
			if fragment.NodeName() != "#document-fragment" {
				t.Errorf("Expected node name '#document-fragment', got %q", fragment.NodeName())
			}

			// Should not be nil
			if fragment == nil {
				t.Error("NewDocumentFragment should not return nil")
			}
		})

		t.Run("Constructor sets node document to current global object's associated Document", func(t *testing.T) {
			// Test that the DocumentFragment's owner document is set correctly
			fragment := NewDocumentFragment(doc)

			// The owner document should be the provided document
			if fragment.OwnerDocument() != doc {
				t.Error("DocumentFragment owner document should be the document provided to constructor")
			}

			// Test with different documents
			doc2 := NewDocument()
			fragment2 := NewDocumentFragment(doc2)

			if fragment2.OwnerDocument() != doc2 {
				t.Error("DocumentFragment should use the document provided to constructor")
			}

			// Fragments should have different owner documents
			if fragment.OwnerDocument() == fragment2.OwnerDocument() {
				t.Error("Different DocumentFragments should have different owner documents when created with different documents")
			}
		})

		t.Run("DocumentFragment has an associated host (null or an element in a different node tree)", func(t *testing.T) {
			fragment := NewDocumentFragment(doc)

			// Host should initially be null
			if fragment.Host() != nil {
				t.Error("DocumentFragment host should initially be null")
			}

			// Should be able to set a host
			element := doc.CreateElement("div")
			fragment.SetHost(element)

			if fragment.Host() != element {
				t.Error("DocumentFragment host should be settable")
			}

			// Should be able to set host back to nil
			fragment.SetHost(nil)
			if fragment.Host() != nil {
				t.Error("DocumentFragment host should be settable to nil")
			}
		})

		t.Run("Host is null unless otherwise stated", func(t *testing.T) {
			// Test multiple fragment creations to ensure host is always initially null
			for i := 0; i < 10; i++ {
				fragment := NewDocumentFragment(doc)
				if fragment.Host() != nil {
					t.Errorf("DocumentFragment %d host should be null unless otherwise stated", i)
				}
			}
		})
	})

	t.Run("WHATWG DOM Section 4.7 - Host-including inclusive ancestor concept", func(t *testing.T) {
		t.Run("Host-including inclusive ancestor of object B", func(t *testing.T) {
			// Create a fragment with a host
			fragment := NewDocumentFragment(doc)
			hostElement := doc.CreateElement("div")
			fragment.SetHost(hostElement)

			// Create a child element in the fragment
			childElement := doc.CreateElement("span")
			fragment.AppendChild(childElement)

			// Test case 1: A is an inclusive ancestor of B
			if !fragment.IsHostIncludingInclusiveAncestorOf(childElement) {
				t.Error("DocumentFragment should be host-including inclusive ancestor of its child")
			}

			if !fragment.IsHostIncludingInclusiveAncestorOf(fragment) {
				t.Error("DocumentFragment should be host-including inclusive ancestor of itself")
			}

			// Test case 2: B's root has a non-null host and A is a host-including inclusive ancestor of B's root's host
			// Create a parent document tree where hostElement lives
			parentDoc := NewDocument()
			parentElement := parentDoc.CreateElement("body")
			parentDoc.AppendChild(parentElement)
			parentElement.AppendChild(hostElement)

			// Now we test the host-including inclusive ancestor relationship
			// In a real implementation, we'd need to implement IsHostIncludingInclusiveAncestorOf on Element too
			// For now, let's verify the fragment's host relationship
			if fragment.Host() != hostElement {
				t.Error("Fragment should maintain its host relationship")
			}

			// Verify that the fragment is still a host-including inclusive ancestor of its children
			if !fragment.IsHostIncludingInclusiveAncestorOf(childElement) {
				t.Error("Fragment should remain host-including inclusive ancestor of its children")
			}
		})

		t.Run("IsInclusiveAncestorOf basic functionality", func(t *testing.T) {
			fragment := NewDocumentFragment(doc)
			child1 := doc.CreateElement("div")
			child2 := doc.CreateElement("span")
			grandchild := doc.CreateElement("p")

			fragment.AppendChild(child1)
			child1.AppendChild(child2)
			child2.AppendChild(grandchild)

			// Fragment should be inclusive ancestor of all descendants
			if !fragment.IsInclusiveAncestorOf(fragment) {
				t.Error("Node should be inclusive ancestor of itself")
			}
			if !fragment.IsInclusiveAncestorOf(child1) {
				t.Error("Fragment should be inclusive ancestor of direct child")
			}
			if !fragment.IsInclusiveAncestorOf(child2) {
				t.Error("Fragment should be inclusive ancestor of grandchild")
			}
			if !fragment.IsInclusiveAncestorOf(grandchild) {
				t.Error("Fragment should be inclusive ancestor of great-grandchild")
			}

			// For now, we only test DocumentFragment's IsInclusiveAncestorOf method
			// In a complete implementation, Element would also have this method

			// Test that fragment is not an inclusive ancestor of unrelated nodes
			unrelatedElement := doc.CreateElement("unrelated")
			if fragment.IsInclusiveAncestorOf(unrelatedElement) {
				t.Error("Fragment should not be inclusive ancestor of unrelated elements")
			}
		})

		t.Run("Host-including inclusive ancestor with complex trees", func(t *testing.T) {
			// Create a complex scenario with nested fragments and hosts
			fragment1 := NewDocumentFragment(doc)
			fragment2 := NewDocumentFragment(doc)

			host1 := doc.CreateElement("template")
			host2 := doc.CreateElement("slot")

			// Set up host relationships
			fragment1.SetHost(host1)
			fragment2.SetHost(host2)

			// Create tree: fragment1 -> host1 -> host2 -> fragment2
			fragment1.AppendChild(host1)
			host1.AppendChild(host2)

			// Add content to fragment2
			content := doc.CreateElement("div")
			fragment2.AppendChild(content)

			// Now fragment1 should be host-including inclusive ancestor of content
			// because content's root (fragment2) has host2, and fragment1 contains host2
			if !fragment1.IsHostIncludingInclusiveAncestorOf(content) {
				t.Error("Fragment1 should be host-including inclusive ancestor through host chain")
			}
		})
	})

	t.Run("DocumentFragment as ParentNode implementation", func(t *testing.T) {
		t.Run("Basic parent-child relationships", func(t *testing.T) {
			fragment := NewDocumentFragment(doc)

			// Initially no children
			if fragment.ChildElementCount() != 0 {
				t.Error("New DocumentFragment should have no children")
			}
			if fragment.FirstElementChild() != nil {
				t.Error("New DocumentFragment should have no first element child")
			}
			if fragment.LastElementChild() != nil {
				t.Error("New DocumentFragment should have no last element child")
			}

			// Add some elements
			div := doc.CreateElement("div")
			span := doc.CreateElement("span")
			p := doc.CreateElement("p")

			fragment.AppendChild(div)
			fragment.AppendChild(span)
			fragment.AppendChild(p)

			// Test children count and access
			if fragment.ChildElementCount() != 3 {
				t.Errorf("Expected 3 element children, got %d", fragment.ChildElementCount())
			}
			if fragment.FirstElementChild() != div {
				t.Error("FirstElementChild should return first appended element")
			}
			if fragment.LastElementChild() != p {
				t.Error("LastElementChild should return last appended element")
			}

			// Test children collection
			children := fragment.Children()
			if children.Length() != 3 {
				t.Errorf("Children collection should have 3 elements, got %d", children.Length())
			}
		})

		t.Run("Mixed content (elements and text nodes)", func(t *testing.T) {
			fragment := NewDocumentFragment(doc)

			// Add mixed content
			text1 := doc.CreateTextNode("before")
			div := doc.CreateElement("div")
			text2 := doc.CreateTextNode("between")
			span := doc.CreateElement("span")
			text3 := doc.CreateTextNode("after")

			fragment.AppendChild(text1)
			fragment.AppendChild(div)
			fragment.AppendChild(text2)
			fragment.AppendChild(span)
			fragment.AppendChild(text3)

			// Should only count element children
			if fragment.ChildElementCount() != 2 {
				t.Errorf("Expected 2 element children, got %d", fragment.ChildElementCount())
			}
			if fragment.FirstElementChild() != div {
				t.Error("FirstElementChild should skip text nodes")
			}
			if fragment.LastElementChild() != span {
				t.Error("LastElementChild should skip text nodes")
			}

			// Children collection should only include elements
			children := fragment.Children()
			if children.Length() != 2 {
				t.Errorf("Children collection should have 2 elements, got %d", children.Length())
			}

			// But ChildNodes should include everything
			childNodes := fragment.ChildNodes()
			if childNodes.Length() != 5 {
				t.Errorf("ChildNodes should have 5 nodes, got %d", childNodes.Length())
			}
		})
	})

	t.Run("DocumentFragment cloning with host preservation", func(t *testing.T) {
		t.Run("CloneNode preserves structure but resets host", func(t *testing.T) {
			fragment := NewDocumentFragment(doc)
			host := doc.CreateElement("template")
			fragment.SetHost(host)

			// Add some content
			div := doc.CreateElement("div")
			div.SetAttribute("id", "test")
			text := doc.CreateTextNode("content")
			div.AppendChild(text)
			fragment.AppendChild(div)

			// Clone without children
			shallowClone := fragment.CloneNode(false)
			clonedFragment, ok := shallowClone.(*DocumentFragment)
			if !ok {
				t.Fatal("Cloned node should be a DocumentFragment")
			}

			// Host should be reset to nil
			if clonedFragment.Host() != nil {
				t.Error("Cloned DocumentFragment host should be reset to nil")
			}

			// Should have same owner document
			if clonedFragment.OwnerDocument() != fragment.OwnerDocument() {
				t.Error("Cloned DocumentFragment should have same owner document")
			}

			// Should have no children (shallow clone)
			if clonedFragment.ChildNodes().Length() != 0 {
				t.Error("Shallow cloned DocumentFragment should have no children")
			}

			// Clone with children
			deepClone := fragment.CloneNode(true)
			deepClonedFragment, ok := deepClone.(*DocumentFragment)
			if !ok {
				t.Fatal("Deep cloned node should be a DocumentFragment")
			}

			// Should have cloned children
			if deepClonedFragment.ChildNodes().Length() != 1 {
				t.Error("Deep cloned DocumentFragment should have cloned children")
			}

			// Host should still be reset
			if deepClonedFragment.Host() != nil {
				t.Error("Deep cloned DocumentFragment host should be reset to nil")
			}

			// Child should be properly cloned
			clonedDiv := deepClonedFragment.FirstElementChild()
			if clonedDiv == nil {
				t.Fatal("Deep cloned fragment should have cloned element child")
			}
			if clonedDiv.GetAttribute("id") != "test" {
				t.Error("Cloned element should preserve attributes")
			}
			if clonedDiv.TextContent() != "content" {
				t.Error("Cloned element should preserve text content")
			}
		})
	})
}

func TestDocumentFragmentConstructorBehavior(t *testing.T) {
	t.Run("Constructor behavior matches spec", func(t *testing.T) {
		doc := NewDocument()

		// Test basic constructor behavior
		fragment := NewDocumentFragment(doc)

		// Should implement Node interface properly
		if fragment.NodeType() != DocumentFragmentNode {
			t.Error("DocumentFragment should have correct node type")
		}

		if fragment.NodeName() != "#document-fragment" {
			t.Error("DocumentFragment should have correct node name")
		}

		// Should have null parent initially
		if fragment.ParentNode() != nil {
			t.Error("New DocumentFragment should have null parent")
		}

		// Should have empty child list initially
		if fragment.ChildNodes().Length() != 0 {
			t.Error("New DocumentFragment should have empty child list")
		}

		// Should be associated with the provided document
		if fragment.OwnerDocument() != doc {
			t.Error("DocumentFragment should be associated with the provided document")
		}
	})

	t.Run("Multiple fragments have independent state", func(t *testing.T) {
		doc := NewDocument()

		fragment1 := NewDocumentFragment(doc)
		fragment2 := NewDocumentFragment(doc)

		// Should be different instances
		if fragment1 == fragment2 {
			t.Error("Multiple DocumentFragment instances should be different objects")
		}

		// Should have independent host settings
		host1 := doc.CreateElement("template")
		fragment1.SetHost(host1)

		if fragment2.Host() != nil {
			t.Error("Setting host on one fragment should not affect another")
		}

		// Should have independent child lists
		div1 := doc.CreateElement("div")
		fragment1.AppendChild(div1)

		if fragment2.ChildNodes().Length() != 0 {
			t.Error("Adding children to one fragment should not affect another")
		}
	})
}

func TestDocumentFragmentHostConcept(t *testing.T) {
	doc := NewDocument()

	t.Run("Host property getter and setter", func(t *testing.T) {
		fragment := NewDocumentFragment(doc)

		// Initial state
		if fragment.Host() != nil {
			t.Error("New DocumentFragment should have null host")
		}

		// Set host to an element
		template := doc.CreateElement("template")
		fragment.SetHost(template)

		if fragment.Host() != template {
			t.Error("Host should be set correctly")
		}

		// Set host to different element
		slot := doc.CreateElement("slot")
		fragment.SetHost(slot)

		if fragment.Host() != slot {
			t.Error("Host should be updatable")
		}

		// Set host back to nil
		fragment.SetHost(nil)

		if fragment.Host() != nil {
			t.Error("Host should be resettable to nil")
		}
	})

	t.Run("Host from different node tree", func(t *testing.T) {
		doc1 := NewDocument()
		doc2 := NewDocument()

		fragment := NewDocumentFragment(doc1)
		hostFromOtherTree := doc2.CreateElement("div")

		// Should be able to set host from different node tree
		fragment.SetHost(hostFromOtherTree)

		if fragment.Host() != hostFromOtherTree {
			t.Error("Should be able to set host from different node tree")
		}

		// Host's owner document should be different
		if fragment.OwnerDocument() == hostFromOtherTree.OwnerDocument() {
			t.Error("Fragment and its host should be able to have different owner documents")
		}
	})

	t.Run("Host relationship in template and shadow root scenarios", func(t *testing.T) {
		doc := NewDocument()

		// Simulate template element scenario
		template := doc.CreateElement("template")
		templateContent := NewDocumentFragment(doc)
		templateContent.SetHost(template)

		// Add content to template
		div := doc.CreateElement("div")
		templateContent.AppendChild(div)

		// Template content should be hosted by template element
		if templateContent.Host() != template {
			t.Error("Template content should be hosted by template element")
		}

		// The div should have the template content as its root
		root := div.GetRootNode(nil)
		if root != templateContent {
			t.Error("Element in template content should have template content as root")
		}
	})
}

func TestDocumentFragmentIntegrationWithDOM(t *testing.T) {
	doc := NewDocument()

	t.Run("DocumentFragment as container for manipulation", func(t *testing.T) {
		fragment := NewDocumentFragment(doc)

		// Build a complex structure in the fragment
		div := doc.CreateElement("div")
		div.SetAttribute("class", "container")

		span1 := doc.CreateElement("span")
		span1.SetTextContent("Hello")

		span2 := doc.CreateElement("span")
		span2.SetTextContent("World")

		div.AppendChild(span1)
		div.AppendChild(span2)
		fragment.AppendChild(div)

		// Verify structure in fragment
		if fragment.ChildElementCount() != 1 {
			t.Error("Fragment should contain the built structure")
		}

		divInFragment := fragment.FirstElementChild()
		if divInFragment.GetAttribute("class") != "container" {
			t.Error("Element structure should be preserved in fragment")
		}

		if divInFragment.ChildElementCount() != 2 {
			t.Error("Nested structure should be preserved in fragment")
		}
	})

	t.Run("DocumentFragment insertion into document", func(t *testing.T) {
		fragment := NewDocumentFragment(doc)

		// Create content in fragment
		h1 := doc.CreateElement("h1")
		h1.SetTextContent("Title")

		p := doc.CreateElement("p")
		p.SetTextContent("Paragraph")

		fragment.AppendChild(h1)
		fragment.AppendChild(p)

		// Create document structure
		body := doc.CreateElement("body")
		doc.AppendChild(body)

		// Insert fragment into document
		body.AppendChild(fragment)

		// Fragment should be empty after insertion (children moved)
		if fragment.ChildNodes().Length() != 0 {
			t.Error("DocumentFragment should be empty after insertion into document")
		}

		// Content should be in the document now
		if body.ChildElementCount() != 2 {
			t.Error("Document should contain the fragment's former children")
		}

		// Elements should have correct parent
		if h1.ParentNode() != body {
			t.Error("Inserted elements should have document element as parent")
		}
		if p.ParentNode() != body {
			t.Error("Inserted elements should have document element as parent")
		}
	})
}

func BenchmarkDocumentFragmentOperations(b *testing.B) {
	doc := NewDocument()

	b.Run("NewDocumentFragment", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = NewDocumentFragment(doc)
		}
	})

	b.Run("Host getter", func(b *testing.B) {
		fragment := NewDocumentFragment(doc)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = fragment.Host()
		}
	})

	b.Run("Host setter", func(b *testing.B) {
		fragment := NewDocumentFragment(doc)
		element := doc.CreateElement("div")
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			fragment.SetHost(element)
		}
	})

	b.Run("IsHostIncludingInclusiveAncestorOf", func(b *testing.B) {
		fragment := NewDocumentFragment(doc)
		child := doc.CreateElement("div")
		fragment.AppendChild(child)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = fragment.IsHostIncludingInclusiveAncestorOf(child)
		}
	})

	b.Run("CloneNode deep", func(b *testing.B) {
		fragment := NewDocumentFragment(doc)
		for i := 0; i < 10; i++ {
			div := doc.CreateElement("div")
			div.SetTextContent("content")
			fragment.AppendChild(div)
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = fragment.CloneNode(true)
		}
	})
}
