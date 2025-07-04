package dom

import (
	"testing"
)

func TestCloneNodeSpec(t *testing.T) {
	doc := NewDocument()

	t.Run("Text node cloning", func(t *testing.T) {
		// Create a text node
		original := NewText("Hello, World!", doc)

		// Test shallow clone
		cloned := original.CloneNode(false)
		if cloned == nil {
			t.Fatal("CloneNode returned nil")
		}

		// Verify it's a different instance
		if cloned == original {
			t.Error("CloneNode returned the same instance")
		}

		// Verify properties
		if cloned.NodeType() != TextNode {
			t.Errorf("Expected node type %d, got %d", TextNode, cloned.NodeType())
		}

		if cloned.NodeName() != "#text" {
			t.Errorf("Expected node name '#text', got %q", cloned.NodeName())
		}

		if cloned.NodeValue() != "Hello, World!" {
			t.Errorf("Expected node value 'Hello, World!', got %q", cloned.NodeValue())
		}

		// Verify owner document
		if cloned.OwnerDocument() != doc {
			t.Error("Cloned node has incorrect owner document")
		}

		// Verify no children (Text nodes can't have children)
		children := cloned.ChildNodes()
		if children.Length() != 0 {
			t.Errorf("Expected 0 children, got %d", children.Length())
		}
	})

	t.Run("Comment node cloning", func(t *testing.T) {
		original := NewComment("This is a comment", doc)
		cloned := original.CloneNode(false)

		if cloned.NodeType() != CommentNode {
			t.Errorf("Expected node type %d, got %d", CommentNode, cloned.NodeType())
		}

		if cloned.NodeName() != "#comment" {
			t.Errorf("Expected node name '#comment', got %q", cloned.NodeName())
		}

		if cloned.NodeValue() != "This is a comment" {
			t.Errorf("Expected node value 'This is a comment', got %q", cloned.NodeValue())
		}
	})

	t.Run("ProcessingInstruction node cloning", func(t *testing.T) {
		original := NewProcessingInstruction("xml-stylesheet", "type=\"text/css\" href=\"style.css\"", doc)
		cloned := original.CloneNode(false)

		if cloned.NodeType() != ProcessingInstructionNode {
			t.Errorf("Expected node type %d, got %d", ProcessingInstructionNode, cloned.NodeType())
		}

		if cloned.NodeName() != "xml-stylesheet" {
			t.Errorf("Expected node name 'xml-stylesheet', got %q", cloned.NodeName())
		}

		clonedPI := cloned.(*ProcessingInstruction)
		if clonedPI.Target() != "xml-stylesheet" {
			t.Errorf("Expected target 'xml-stylesheet', got %q", clonedPI.Target())
		}

		if clonedPI.Data() != "type=\"text/css\" href=\"style.css\"" {
			t.Errorf("Expected data 'type=\"text/css\" href=\"style.css\"', got %q", clonedPI.Data())
		}
	})

	t.Run("DocumentType node cloning", func(t *testing.T) {
		original := NewDocumentType("html", "-//W3C//DTD HTML 4.01//EN", "http://www.w3.org/TR/html4/strict.dtd", doc)
		cloned := original.CloneNode(false)

		if cloned.NodeType() != DocumentTypeNode {
			t.Errorf("Expected node type %d, got %d", DocumentTypeNode, cloned.NodeType())
		}

		clonedDT := cloned.(*DocumentType)
		if clonedDT.Name() != "html" {
			t.Errorf("Expected name 'html', got %q", clonedDT.Name())
		}

		if clonedDT.PublicID() != "-//W3C//DTD HTML 4.01//EN" {
			t.Errorf("Expected publicID '-//W3C//DTD HTML 4.01//EN', got %q", clonedDT.PublicID())
		}

		if clonedDT.SystemID() != "http://www.w3.org/TR/html4/strict.dtd" {
			t.Errorf("Expected systemID 'http://www.w3.org/TR/html4/strict.dtd', got %q", clonedDT.SystemID())
		}
	})

	t.Run("Attr node cloning", func(t *testing.T) {
		original := NewAttr("class", "button primary", doc)
		cloned := original.CloneNode(false)

		if cloned.NodeType() != AttributeNode {
			t.Errorf("Expected node type %d, got %d", AttributeNode, cloned.NodeType())
		}

		clonedAttr := cloned.(*Attr)
		if clonedAttr.Name() != "class" {
			t.Errorf("Expected name 'class', got %q", clonedAttr.Name())
		}

		if clonedAttr.Value() != "button primary" {
			t.Errorf("Expected value 'button primary', got %q", clonedAttr.Value())
		}
	})
}

func TestCloneNode_ShadowRootProhibition(t *testing.T) {
	doc := NewDocument()
	host := doc.CreateElement("div")
	shadowRoot, err := host.AttachShadow(ShadowRootInit{Mode: ShadowRootModeOpen})
	if err != nil {
		t.Fatalf("Failed to attach shadow root: %v", err)
	}

	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("Cloning a ShadowRoot should panic")
			return
		}
		domErr, ok := r.(*DOMException)
		if !ok {
			t.Errorf("Panic was not a DOMException, got %T", r)
			return
		}
		if domErr.Name() != "NotSupportedError" {
			t.Errorf("Expected NotSupportedError, got %s", domErr.Name())
		}
	}()

	// This should panic
	shadowRoot.CloneNode(false)
}

func TestCloneNode_ElementWithShadowRoot(t *testing.T) {
	doc := NewDocument()
	originalHost := doc.CreateElement("div")
	originalHost.SetAttribute("id", "host")

	// Attach a clonable shadow root
	originalShadowInit := ShadowRootInit{
		Mode:           ShadowRootModeOpen,
		DelegatesFocus: true,
		Clonable:       true,
		Serializable:   true,
		SlotAssignment: SlotAssignmentManual,
	}
	originalSR, err := originalHost.AttachShadow(originalShadowInit)
	if err != nil {
		t.Fatalf("Failed to attach original shadow root: %v", err)
	}
	// Manually set declarative and customElementRegistry for testing their transfer
	originalSR.declarative = true
	originalSR.customElementRegistry = NewCustomElementRegistry(doc)

	shadowText := NewText("Shadow Content", doc)
	shadowSpan := NewElement("span", doc)
	shadowSpan.AppendChild(NewText("Inside Span", doc))
	originalSR.AppendChild(shadowText)
	originalSR.AppendChild(shadowSpan)

	t.Run("Clone host with clonable ShadowRoot", func(t *testing.T) {
		clonedHostNode := originalHost.CloneNode(true) // Deep clone the host

		clonedHost, ok := clonedHostNode.(*Element)
		if !ok {
			t.Fatalf("Cloned host is not an Element")
		}
		if clonedHost == originalHost {
			t.Fatal("Cloned host is the same instance as original")
		}

		clonedSR := clonedHost.ShadowRoot()
		if clonedSR == nil {
			t.Fatal("Cloned host does not have a shadow root")
		}
		if clonedSR == originalSR {
			t.Fatal("Cloned shadow root is the same instance as original")
		}

		// Verify properties of the cloned shadow root
		if clonedSR.Mode() != originalSR.Mode() {
			t.Errorf("Expected cloned SR mode %s, got %s", originalSR.Mode(), clonedSR.Mode())
		}
		if clonedSR.DelegatesFocus() != originalSR.DelegatesFocus() {
			t.Errorf("Expected cloned SR delegatesFocus %t, got %t", originalSR.DelegatesFocus(), clonedSR.DelegatesFocus())
		}
		if !clonedSR.Clonable() { // Spec: "Attach a shadow root with ... true (for clonable)..."
			t.Error("Expected cloned SR to be clonable (true by spec for this operation)")
		}
		if clonedSR.Serializable() != originalSR.Serializable() {
			t.Errorf("Expected cloned SR serializable %t, got %t", originalSR.Serializable(), clonedSR.Serializable())
		}
		if clonedSR.SlotAssignment() != originalSR.SlotAssignment() {
			t.Errorf("Expected cloned SR slotAssignment %s, got %s", originalSR.SlotAssignment(), clonedSR.SlotAssignment())
		}
		if clonedSR.customElementRegistry != originalSR.customElementRegistry {
			t.Error("Expected cloned SR customElementRegistry to be copied")
		}
		if clonedSR.declarative != originalSR.declarative {
			t.Errorf("Expected cloned SR declarative %t, got %t", originalSR.declarative, clonedSR.declarative)
		}

		// Verify children of the cloned shadow root
		clonedSRChildren := clonedSR.ChildNodes()
		if clonedSRChildren.Length() != 2 {
			t.Fatalf("Expected 2 children in cloned shadow root, got %d", clonedSRChildren.Length())
		}
		if clonedSRChildren.Item(0).NodeValue() != "Shadow Content" {
			t.Errorf("Incorrect text node in cloned shadow root: %s", clonedSRChildren.Item(0).NodeValue())
		}
		clonedSRSpan, ok := clonedSRChildren.Item(1).(*Element)
		if !ok || clonedSRSpan.TagName() != "SPAN" {
			t.Errorf("Incorrect element node in cloned shadow root")
		}
		if clonedSRSpan.FirstChild().NodeValue() != "Inside Span" {
			t.Errorf("Incorrect text in span within cloned shadow root")
		}
		if clonedSRChildren.Item(0).OwnerDocument() != clonedHost.OwnerDocument() {
			t.Error("Cloned shadow child has incorrect owner document")
		}
	})

	t.Run("Clone host with non-clonable ShadowRoot", func(t *testing.T) {
		nonClonableHost := NewElement("div", doc)
		nonClonableSRInit := ShadowRootInit{Mode: ShadowRootModeOpen, Clonable: false}
		_, err := nonClonableHost.AttachShadow(nonClonableSRInit)
		if err != nil {
			t.Fatalf("Failed to attach non-clonable shadow root: %v", err)
		}

		clonedHostNode := nonClonableHost.CloneNode(true)
		clonedHost := clonedHostNode.(*Element)

		if clonedHost.ShadowRoot() != nil {
			t.Error("Non-clonable shadow root should not be cloned")
		}
	})
}

func TestCloneNode_ElementIsValue(t *testing.T) {
	doc := NewDocument()
	original := doc.CreateElement("button", ElementCreationOptions{Is: "custom-button"})
	original.SetAttribute("id", "mybtn")

	clonedNode := original.CloneNode(false)
	clonedElem, ok := clonedNode.(*Element)
	if !ok {
		t.Fatalf("Cloned node is not an element")
	}

	if clonedElem.IsValue() != "custom-button" {
		t.Errorf("Expected 'is' value 'custom-button', got '%s'", clonedElem.IsValue())
	}
	if clonedElem.GetAttribute("id") != "mybtn" { // Ensure other parts are still cloned
		t.Errorf("Attribute 'id' not cloned correctly")
	}
}

func TestCloneNode_DocumentProperties(t *testing.T) {
	originalDoc := NewDocument()
	originalDoc.url = "http://example.com/original"
	originalDoc.documentURI = "http://example.com/original"
	originalDoc.contentType = "text/html"
	originalDoc.characterSet = "ISO-8859-1"
	originalDoc.charset = "ISO-8859-1"
	originalDoc.inputEncoding = "ISO-8859-1"
	originalDoc.mode = "quirks" // Should result in CompatMode "BackCompat"
	originalDoc.documentType = "html"
	originalDoc.customElementRegistry = NewCustomElementRegistry(originalDoc)

	el := originalDoc.CreateElement("div")
	originalDoc.AppendChild(el)

	clonedDocNode := originalDoc.CloneNode(true) // Deep clone
	clonedDoc, ok := clonedDocNode.(*Document)
	if !ok {
		t.Fatalf("Cloned node is not a Document")
	}

	if clonedDoc.URL() != originalDoc.URL() {
		t.Errorf("Expected URL '%s', got '%s'", originalDoc.URL(), clonedDoc.URL())
	}
	if clonedDoc.ContentType() != originalDoc.ContentType() {
		t.Errorf("Expected ContentType '%s', got '%s'", originalDoc.ContentType(), clonedDoc.ContentType())
	}
	if clonedDoc.CharacterSet() != originalDoc.CharacterSet() {
		t.Errorf("Expected CharacterSet '%s', got '%s'", originalDoc.CharacterSet(), clonedDoc.CharacterSet())
	}
	if clonedDoc.CompatMode() != "BackCompat" {
		t.Errorf("Expected CompatMode 'BackCompat', got '%s'", clonedDoc.CompatMode())
	}
	if clonedDoc.documentType != originalDoc.documentType {
		t.Errorf("Expected documentType (type) '%s', got '%s'", originalDoc.documentType, clonedDoc.documentType)
	}
	if clonedDoc.CustomElementRegistry() == nil {
		t.Error("Expected CustomElementRegistry to be copied, but it was nil")
	} else if clonedDoc.CustomElementRegistry() == originalDoc.CustomElementRegistry() {
		if clonedDoc.CustomElementRegistry().Document() != clonedDoc {
			t.Error("Cloned document's CustomElementRegistry is not associated with the cloned document")
		}
	}

	if clonedDoc.ChildNodes().Length() != 1 {
		t.Errorf("Expected 1 child for deep-cloned document, got %d", clonedDoc.ChildNodes().Length())
	}
}

func TestElementCloning(t *testing.T) {
	doc := NewDocument()

	t.Run("Element shallow clone", func(t *testing.T) {
		// Create element with attributes
		original := NewElement("div", doc)
		original.SetAttribute("id", "test-div")
		original.SetAttribute("class", "container")
		original.SetAttribute("data-value", "123")

		// Add some children
		text1 := NewText("Hello ", doc)
		span := NewElement("span", doc)
		text2 := NewText("World", doc)
		span.AppendChild(text2)
		original.AppendChild(text1)
		original.AppendChild(span)

		// Shallow clone
		cloned := original.CloneNode(false)

		// Verify basic properties
		if cloned.NodeType() != ElementNode {
			t.Errorf("Expected node type %d, got %d", ElementNode, cloned.NodeType())
		}

		clonedElem := cloned.(*Element)
		if clonedElem.TagName() != "DIV" { // Expect uppercase
			t.Errorf("Expected tag name 'DIV', got %q", clonedElem.TagName())
		}

		// Verify attributes are copied
		if clonedElem.GetAttribute("id") != "test-div" {
			t.Errorf("Expected id 'test-div', got %q", clonedElem.GetAttribute("id"))
		}

		if clonedElem.GetAttribute("class") != "container" {
			t.Errorf("Expected class 'container', got %q", clonedElem.GetAttribute("class"))
		}

		if clonedElem.GetAttribute("data-value") != "123" {
			t.Errorf("Expected data-value '123', got %q", clonedElem.GetAttribute("data-value"))
		}

		// Verify no children in shallow clone
		children := cloned.ChildNodes()
		if children.Length() != 0 {
			t.Errorf("Expected 0 children in shallow clone, got %d", children.Length())
		}
	})

	t.Run("Element deep clone", func(t *testing.T) {
		// Create element with attributes and nested children
		original := NewElement("div", doc)
		original.SetAttribute("id", "test-div")
		original.SetAttribute("class", "container")

		text1 := NewText("Hello ", doc)
		span := NewElement("span", doc)
		span.SetAttribute("style", "font-weight: bold")
		text2 := NewText("World", doc)
		text3 := NewText("!", doc)

		span.AppendChild(text2)
		original.AppendChild(text1)
		original.AppendChild(span)
		original.AppendChild(text3)

		// Deep clone
		cloned := original.CloneNode(true)

		// Verify basic properties
		clonedElem := cloned.(*Element)
		if clonedElem.TagName() != "DIV" { // Expect uppercase
			t.Errorf("Expected tag name 'DIV', got %q", clonedElem.TagName())
		}

		// Verify attributes are copied
		if clonedElem.GetAttribute("id") != "test-div" {
			t.Errorf("Expected id 'test-div', got %q", clonedElem.GetAttribute("id"))
		}

		// Verify children are cloned
		children := cloned.ChildNodes()
		if children.Length() != 3 {
			t.Fatalf("Expected 3 children in deep clone, got %d", children.Length())
		}

		// Check first child (text node)
		child1 := children.Item(0)
		if child1.NodeType() != TextNode {
			t.Errorf("Expected first child to be text node, got %d", child1.NodeType())
		}
		if child1.NodeValue() != "Hello " {
			t.Errorf("Expected first child value 'Hello ', got %q", child1.NodeValue())
		}

		// Check second child (span element)
		child2 := children.Item(1)
		if child2.NodeType() != ElementNode {
			t.Errorf("Expected second child to be element node, got %d", child2.NodeType())
		}

		spanElem := child2.(*Element)
		if spanElem.TagName() != "SPAN" { // Expect uppercase
			t.Errorf("Expected span tag name 'SPAN', got %q", spanElem.TagName())
		}

		if spanElem.GetAttribute("style") != "font-weight: bold" {
			t.Errorf("Expected style attribute 'font-weight: bold', got %q", spanElem.GetAttribute("style"))
		}

		// Check span's child
		spanChildren := spanElem.ChildNodes()
		if spanChildren.Length() != 1 {
			t.Errorf("Expected span to have 1 child, got %d", spanChildren.Length())
		}

		//TODO(dpemmons) this ends up empty!
		spanChild := spanChildren.Item(0)
		if spanChild.NodeValue() != "World" {
			t.Errorf("Expected span child value 'World', got %q", spanChild.NodeValue())
		}

		// Check third child (text node)
		child3 := children.Item(2)
		if child3.NodeValue() != "!" {
			t.Errorf("Expected third child value '!', got %q", child3.NodeValue())
		}

		// Verify cloned nodes are different instances
		if child1 == text1 {
			t.Error("Cloned text node is the same instance as original")
		}

		if child2 == span {
			t.Error("Cloned span element is the same instance as original")
		}
	})
}

func TestDocumentFragmentCloning(t *testing.T) {
	doc := NewDocument()

	t.Run("DocumentFragment shallow clone", func(t *testing.T) {
		original := NewDocumentFragment(doc)

		// Add some children
		text := NewText("Fragment text", doc)
		div := NewElement("div", doc)
		original.AppendChild(text)
		original.AppendChild(div)

		// Shallow clone
		cloned := original.CloneNode(false)

		if cloned.NodeType() != DocumentFragmentNode {
			t.Errorf("Expected node type %d, got %d", DocumentFragmentNode, cloned.NodeType())
		}

		// Should have no children in shallow clone
		children := cloned.ChildNodes()
		if children.Length() != 0 {
			t.Errorf("Expected 0 children in shallow clone, got %d", children.Length())
		}
	})

	t.Run("DocumentFragment deep clone", func(t *testing.T) {
		original := NewDocumentFragment(doc)

		// Add some children
		text := NewText("Fragment text", doc)
		div := NewElement("div", doc)
		div.SetAttribute("class", "test")
		original.AppendChild(text)
		original.AppendChild(div)

		// Deep clone
		cloned := original.CloneNode(true)

		// Should have all children cloned
		children := cloned.ChildNodes()
		if children.Length() != 2 {
			t.Fatalf("Expected 2 children in deep clone, got %d", children.Length())
		}

		// Check first child
		child1 := children.Item(0)
		if child1.NodeType() != TextNode {
			t.Errorf("Expected first child to be text node, got %d", child1.NodeType())
		}
		if child1.NodeValue() != "Fragment text" {
			t.Errorf("Expected first child value 'Fragment text', got %q", child1.NodeValue())
		}

		// Check second child
		child2 := children.Item(1)
		if child2.NodeType() != ElementNode {
			t.Errorf("Expected second child to be element node, got %d", child2.NodeType())
		}

		clonedDiv := child2.(*Element)
		if clonedDiv.GetAttribute("class") != "test" {
			t.Errorf("Expected class attribute 'test', got %q", clonedDiv.GetAttribute("class"))
		}

		// Verify different instances
		if child1 == text {
			t.Error("Cloned text node is the same instance as original")
		}
		if child2 == div {
			t.Error("Cloned div element is the same instance as original")
		}
	})
}

func TestDocumentCloning(t *testing.T) {
	t.Run("Document shallow clone", func(t *testing.T) {
		original := NewDocument()

		// Add some children
		doctype := NewDocumentType("html", "", "", original)
		html := NewElement("html", original)
		original.AppendChild(doctype)
		original.AppendChild(html)

		// Shallow clone
		cloned := original.CloneNode(false)

		if cloned.NodeType() != DocumentNode {
			t.Errorf("Expected node type %d, got %d", DocumentNode, cloned.NodeType())
		}

		// Should have no children in shallow clone
		children := cloned.ChildNodes()
		if children.Length() != 0 {
			t.Errorf("Expected 0 children in shallow clone, got %d", children.Length())
		}

		// Verify owner document is nil for documents
		if cloned.OwnerDocument() != nil {
			t.Error("Cloned document should have nil owner document")
		}
	})

	t.Run("Document deep clone", func(t *testing.T) {
		original := NewDocument()

		// Add DOCTYPE and HTML element
		doctype := NewDocumentType("html", "", "", original)
		html := NewElement("html", original)
		head := NewElement("head", original)
		body := NewElement("body", original)
		title := NewElement("title", original)
		titleText := NewText("Test Document", original)

		title.AppendChild(titleText)
		head.AppendChild(title)
		html.AppendChild(head)
		html.AppendChild(body)
		original.AppendChild(doctype)
		original.AppendChild(html)

		// Deep clone
		cloned := original.CloneNode(true)

		// Verify structure is preserved
		children := cloned.ChildNodes()
		if children.Length() != 2 {
			t.Fatalf("Expected 2 children in deep clone, got %d", children.Length())
		}

		// Check DOCTYPE
		child1 := children.Item(0)
		if child1.NodeType() != DocumentTypeNode {
			t.Errorf("Expected first child to be DocumentType, got %d", child1.NodeType())
		}

		// Check HTML element
		child2 := children.Item(1)
		if child2.NodeType() != ElementNode {
			t.Errorf("Expected second child to be Element, got %d", child2.NodeType())
		}

		htmlElem := child2.(*Element)
		if htmlElem.TagName() != "HTML" { // Expect uppercase
			t.Errorf("Expected html tag 'HTML', got %q", htmlElem.TagName())
		}

		// Check nested structure
		htmlChildren := htmlElem.ChildNodes()
		if htmlChildren.Length() != 2 {
			t.Errorf("Expected html to have 2 children, got %d", htmlChildren.Length())
		}

		headElem := htmlChildren.Item(0).(*Element)
		if headElem.TagName() != "HEAD" { // Expect uppercase
			t.Errorf("Expected head tag 'HEAD', got %q", headElem.TagName())
		}

		// Verify all cloned nodes are different instances
		if child1 == doctype {
			t.Error("Cloned DOCTYPE is the same instance as original")
		}
		if child2 == html {
			t.Error("Cloned HTML element is the same instance as original")
		}
	})
}

func TestCrossDocumentCloning(t *testing.T) {
	// Create two documents
	doc1 := NewDocument()
	doc2 := NewDocument()

	t.Run("Clone node to different document", func(t *testing.T) {
		// Create element in doc1
		original := NewElement("div", doc1)
		original.SetAttribute("id", "test")
		text := NewText("Hello", doc1)
		original.AppendChild(text)

		// Clone to doc2
		cloned := CloneNodeWithDocument(original, doc2, true)

		// Verify owner document is changed
		if cloned.OwnerDocument() != doc2 {
			t.Error("Cloned node should have new owner document")
		}

		// Verify children also have new owner document
		children := cloned.ChildNodes()
		if children.Length() != 1 {
			t.Fatalf("Expected 1 child, got %d", children.Length())
		}

		child := children.Item(0)
		if child.OwnerDocument() != doc2 {
			t.Error("Cloned child should have new owner document")
		}

		// Verify properties are preserved
		clonedElem := cloned.(*Element)
		if clonedElem.GetAttribute("id") != "test" {
			t.Errorf("Expected id 'test', got %q", clonedElem.GetAttribute("id"))
		}

		if child.NodeValue() != "Hello" {
			t.Errorf("Expected child value 'Hello', got %q", child.NodeValue())
		}
	})
}

func TestCloningValidation(t *testing.T) {
	doc := NewDocument()

	t.Run("Validation of cloned nodes", func(t *testing.T) {
		// Create a complex structure
		div := NewElement("div", doc)
		div.SetAttribute("class", "test")
		span := NewElement("span", doc)
		text := NewText("Test text", doc)

		span.AppendChild(text)
		div.AppendChild(span)

		// Clone it
		cloned := div.CloneNode(true)

		// Validate using the validation function
		err := validateClonedNode(div, cloned, true)
		if err != nil {
			t.Errorf("Validation failed: %v", err)
		}

		// Test shallow clone validation
		shallowCloned := div.CloneNode(false)
		err = validateClonedNode(div, shallowCloned, false)
		if err != nil {
			t.Errorf("Shallow clone validation failed: %v", err)
		}
	})

	t.Run("Validation should catch mismatched node types", func(t *testing.T) {
		text := NewText("Test", doc)
		div := NewElement("div", doc)

		// This should fail validation
		err := validateClonedNode(text, div, false)
		if err == nil {
			t.Error("Expected validation to fail for mismatched node types")
		}
	})

	t.Run("Validation should catch incorrect child count", func(t *testing.T) {
		// Create original with children
		original := NewElement("div", doc)
		text := NewText("Test", doc)
		original.AppendChild(text)

		// Create "cloned" without children
		fakeClone := NewElement("div", doc)

		// This should fail validation for deep clone
		err := validateClonedNode(original, fakeClone, true)
		if err == nil {
			t.Error("Expected validation to fail for incorrect child count")
		}
	})
}

func TestCDATASectionCloning(t *testing.T) {
	doc := NewDocument()

	t.Run("CDATASection cloning", func(t *testing.T) {
		original := NewCDATASection("Some CDATA content", doc)
		cloned := original.CloneNode(false)

		if cloned.NodeType() != CDataSectionNode {
			t.Errorf("Expected node type %d, got %d", CDataSectionNode, cloned.NodeType())
		}

		if cloned.NodeName() != "#cdata-section" {
			t.Errorf("Expected node name '#cdata-section', got %q", cloned.NodeName())
		}

		if cloned.NodeValue() != "Some CDATA content" {
			t.Errorf("Expected node value 'Some CDATA content', got %q", cloned.NodeValue())
		}

		// Verify it's a different instance
		if cloned == original {
			t.Error("CloneNode returned the same instance")
		}
	})
}

func TestCloneNode_ElementWithNamespace(t *testing.T) { // Renamed for clarity and consistency
	doc := NewDocument()

	t.Run("Element with namespace cloning", func(t *testing.T) {
		// Create element with namespace
		original, err := NewElementNS("http://www.w3.org/2000/svg", "svg:rect", doc)
		if err != nil {
			t.Fatalf("Failed to create namespaced element: %v", err)
		}

		original.SetAttribute("width", "100")
		original.SetAttribute("height", "50")

		// Clone it
		cloned := original.CloneNode(false)

		// Verify namespace is preserved
		clonedElem := cloned.(*Element)
		if clonedElem.NamespaceURI() != "http://www.w3.org/2000/svg" {
			t.Errorf("Expected namespace URI 'http://www.w3.org/2000/svg', got %q", clonedElem.NamespaceURI())
		}

		if clonedElem.TagName() != "svg:rect" {
			t.Errorf("Expected tag name 'svg:rect', got %q", clonedElem.TagName())
		}

		// Verify attributes are preserved
		if clonedElem.GetAttribute("width") != "100" {
			t.Errorf("Expected width '100', got %q", clonedElem.GetAttribute("width"))
		}
	})
}

func TestCloneNode_AttrNamespace(t *testing.T) {
	doc := NewDocument()
	// elem := doc.CreateElement("div") // Owner element not strictly needed for Attr.CloneNode

	originalAttr, err := NewAttrNS("http://www.example.com/ns", "ex:myattr", "value123", doc)
	if err != nil {
		t.Fatalf("Failed to create namespaced Attr: %v", err)
	}
	// Setting ownerElement is usually done when attribute is added to an element's NamedNodeMap
	// For direct Attr.CloneNode testing, it's not essential but good for completeness if it were part of Attr struct.
	// originalAttr.ownerElement = elem // This field is not directly settable on Attr from outside package

	clonedNode := originalAttr.CloneNode(false)
	clonedAttr, ok := clonedNode.(*Attr)
	if !ok {
		t.Fatalf("Cloned node is not an Attr, got %T", clonedNode)
	}

	if clonedAttr.NamespaceURI() != "http://www.example.com/ns" {
		t.Errorf("Expected NamespaceURI '%s', got '%s'", "http://www.example.com/ns", clonedAttr.NamespaceURI())
	}
	if clonedAttr.Prefix() != "ex" {
		t.Errorf("Expected Prefix '%s', got '%s'", "ex", clonedAttr.Prefix())
	}
	if clonedAttr.LocalName() != "myattr" {
		t.Errorf("Expected LocalName '%s', got '%s'", "myattr", clonedAttr.LocalName())
	}
	if clonedAttr.Value() != "value123" {
		t.Errorf("Expected Value '%s', got '%s'", "value123", clonedAttr.Value())
	}
	if clonedAttr.Name() != "ex:myattr" {
		t.Errorf("Expected Name (NodeName) '%s', got '%s'", "ex:myattr", clonedAttr.Name())
	}
	if clonedAttr.OwnerDocument() != doc {
		t.Error("Cloned Attr has incorrect owner document")
	}
	if clonedAttr == originalAttr {
		t.Error("Cloned Attr is the same instance as original")
	}
}
