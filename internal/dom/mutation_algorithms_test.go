package dom

import (
	"testing"
)

// TestMutationAlgorithmsSpecCompliance tests the WHATWG DOM Section 4.2.3 mutation algorithms
func TestMutationAlgorithmsSpecCompliance(t *testing.T) {
	t.Run("ensurePreInsertValidity", testEnsurePreInsertValidity)
	t.Run("preInsert", testPreInsert)
	t.Run("insert", testInsert)
	t.Run("preRemove", testPreRemove)
	t.Run("remove", testRemove)
	t.Run("replaceAll", testReplaceAll)
	t.Run("DocumentFragmentInsertion", testDocumentFragmentInsertion)
	t.Run("MutationObserverIntegration", testMutationObserverIntegration)
}

// testEnsurePreInsertValidity tests the "ensure pre-insert validity" algorithm
func testEnsurePreInsertValidity(t *testing.T) {
	doc := NewDocument()
	element := NewElement("div", doc)
	text := NewText("test", doc)
	doctype := NewDocumentType("html", "", "", doc)
	fragment := NewDocumentFragment(doc)

	t.Run("Valid parent types", func(t *testing.T) {
		// DocumentFragment and Element are valid parents for text
		if err := ensurePreInsertValidity(text, fragment, nil); err != nil {
			t.Errorf("DocumentFragment should be valid parent: %v", err)
		}
		if err := ensurePreInsertValidity(text, element, nil); err != nil {
			t.Errorf("Element should be valid parent: %v", err)
		}

		// Document is valid parent for elements and doctypes
		elementChild := NewElement("html", doc)
		if err := ensurePreInsertValidity(elementChild, doc, nil); err != nil {
			t.Errorf("Document should be valid parent for elements: %v", err)
		}
		if err := ensurePreInsertValidity(doctype, doc, nil); err != nil {
			t.Errorf("Document should be valid parent for doctype: %v", err)
		}
	})

	t.Run("Invalid parent types", func(t *testing.T) {
		// Text, Comment, etc. are not valid parents
		comment := NewComment("comment", doc)
		if err := ensurePreInsertValidity(text, comment, nil); err == nil {
			t.Error("Comment should not be valid parent")
		}
		if err := ensurePreInsertValidity(text, text, nil); err == nil {
			t.Error("Text should not be valid parent")
		}
	})

	t.Run("Host-including ancestor check", func(t *testing.T) {
		// Node cannot be ancestor of parent
		element.AppendChild(text)
		if err := ensurePreInsertValidity(element, text, nil); err == nil {
			t.Error("Node cannot be ancestor of parent")
		}
	})

	t.Run("Child parent validation", func(t *testing.T) {
		parent := NewElement("parent", doc)
		child := NewElement("child", doc)
		parent.AppendChild(child)

		wrongParent := NewElement("wrong", doc)
		if err := ensurePreInsertValidity(text, wrongParent, child); err == nil {
			t.Error("Child's parent must match specified parent")
		}
	})

	t.Run("Valid node types", func(t *testing.T) {
		// Test valid nodes into appropriate parents
		testCases := []struct {
			node   Node
			parent Node
			name   string
		}{
			{text, fragment, "Text into DocumentFragment"},
			{text, element, "Text into Element"},
			{NewComment("comment", doc), fragment, "Comment into DocumentFragment"},
			{NewComment("comment", doc), element, "Comment into Element"},
			{NewProcessingInstruction("pi", "data", doc), fragment, "ProcessingInstruction into DocumentFragment"},
			{NewProcessingInstruction("pi", "data", doc), element, "ProcessingInstruction into Element"},
			{NewElement("span", doc), fragment, "Element into DocumentFragment"},
			{NewElement("span", doc), element, "Element into Element"},
			{NewElement("html", doc), doc, "Element into Document"},
			{doctype, doc, "DocumentType into Document"},
		}

		for _, tc := range testCases {
			if err := ensurePreInsertValidity(tc.node, tc.parent, nil); err != nil {
				t.Errorf("%s should be valid: %v", tc.name, err)
			}
		}
	})

	t.Run("Document constraints", func(t *testing.T) {
		// Text nodes cannot be inserted into documents
		if err := ensurePreInsertValidity(text, doc, nil); err == nil {
			t.Error("Text nodes cannot be inserted into documents")
		}

		// DocumentType can only be inserted into documents
		if err := ensurePreInsertValidity(doctype, element, nil); err == nil {
			t.Error("DocumentType can only be inserted into documents")
		}
	})

	t.Run("Document element constraints", func(t *testing.T) {
		emptyDoc := NewDocument()
		htmlElement := NewElement("html", emptyDoc)

		// Document can only have one element
		emptyDoc.AppendChild(htmlElement)
		anotherElement := NewElement("body", emptyDoc)
		if err := ensurePreInsertValidity(anotherElement, emptyDoc, nil); err == nil {
			t.Error("Document can only have one element child")
		}
	})

	t.Run("DocumentFragment with multiple elements", func(t *testing.T) {
		emptyDoc := NewDocument()
		frag := NewDocumentFragment(emptyDoc)
		frag.AppendChild(NewElement("div", emptyDoc))
		frag.AppendChild(NewElement("span", emptyDoc))

		if err := ensurePreInsertValidity(frag, emptyDoc, nil); err == nil {
			t.Error("DocumentFragment with multiple elements cannot be inserted into document")
		}
	})

	t.Run("DocumentFragment with text child", func(t *testing.T) {
		emptyDoc := NewDocument()
		frag := NewDocumentFragment(emptyDoc)
		frag.AppendChild(NewText("text", emptyDoc))

		if err := ensurePreInsertValidity(frag, emptyDoc, nil); err == nil {
			t.Error("DocumentFragment with text child cannot be inserted into document")
		}
	})
}

// testPreInsert tests the "pre-insert" algorithm
func testPreInsert(t *testing.T) {
	doc := NewDocument()
	parent := NewElement("div", doc)
	child1 := NewElement("span", doc)
	child2 := NewElement("p", doc)

	t.Run("Basic insertion", func(t *testing.T) {
		result, err := preInsert(child1, parent, nil)
		if err != nil {
			t.Errorf("preInsert failed: %v", err)
		}
		if result != child1 {
			t.Error("preInsert should return the inserted node")
		}
		if child1.ParentNode() != parent {
			t.Error("Child should be parented to the specified parent")
		}
	})

	t.Run("Insert before existing child", func(t *testing.T) {
		result, err := preInsert(child2, parent, child1)
		if err != nil {
			t.Errorf("preInsert failed: %v", err)
		}
		if result != child2 {
			t.Error("preInsert should return the inserted node")
		}

		children := parent.ChildNodes()
		if children.Item(0) != child2 || children.Item(1) != child1 {
			t.Error("Child2 should be inserted before child1")
		}
	})

	t.Run("Reference child is node", func(t *testing.T) {
		// When reference child is the node being inserted, use next sibling
		// This test case demonstrates what happens when you try to insert a node before itself
		// The node must already be in the parent for this to make sense
		child3 := NewElement("em", doc)
		parent.AppendChild(child3) // First add the child to the parent

		// Now try to insert it before itself - should use next sibling as reference
		result, err := preInsert(child3, parent, child3)
		if err != nil {
			t.Errorf("preInsert failed: %v", err)
		}
		if result != child3 {
			t.Error("preInsert should return the inserted node")
		}
		// The child should still be in the parent
		if child3.ParentNode() != parent {
			t.Error("Child should still be parented to parent")
		}
	})

	t.Run("Validation errors", func(t *testing.T) {
		invalidChild := NewText("text", doc)
		_, err := preInsert(invalidChild, doc, nil)
		if err == nil {
			t.Error("preInsert should fail when validation fails")
		}
	})
}

// testInsert tests the "insert" algorithm
func testInsert(t *testing.T) {
	doc := NewDocument()
	parent := NewElement("div", doc)

	t.Run("Single node insertion", func(t *testing.T) {
		child := NewElement("span", doc)
		insertNode(child, parent, nil, false)

		if child.ParentNode() != parent {
			t.Error("Child should be adopted into parent")
		}
		if parent.FirstChild() != child {
			t.Error("Child should be appended to parent")
		}
	})

	t.Run("DocumentFragment insertion", func(t *testing.T) {
		fragment := NewDocumentFragment(doc)
		child1 := NewElement("p", doc)
		child2 := NewElement("em", doc)
		fragment.AppendChild(child1)
		fragment.AppendChild(child2)

		insertNode(fragment, parent, nil, false)

		// Fragment children should be moved to parent
		if child1.ParentNode() != parent || child2.ParentNode() != parent {
			t.Error("Fragment children should be adopted into parent")
		}
		if fragment.HasChildNodes() {
			t.Error("Fragment should be empty after insertion")
		}

		children := parent.ChildNodes()
		if children.Length() != 3 { // original span + p + em
			t.Errorf("Parent should have 3 children, got %d", children.Length())
		}
	})

	t.Run("Insert before existing child", func(t *testing.T) {
		newChild := NewElement("strong", doc)
		firstChild := parent.FirstChild()
		insertNode(newChild, parent, firstChild, false)

		if parent.FirstChild() != newChild {
			t.Error("New child should be first child")
		}
	})

	t.Run("Suppress observers", func(t *testing.T) {
		// Test that observers are suppressed when flag is set
		observer := NewMutationObserver(func(records []*MutationRecord, observer *MutationObserver) {
			t.Error("Observer should be suppressed")
		})
		observer.Observe(parent, MutationObserverInit{ChildList: true})

		suppressedChild := NewElement("suppressed", doc)
		insertNode(suppressedChild, parent, nil, true) // suppress observers

		observer.Disconnect()
	})
}

// testPreRemove tests the "pre-remove" algorithm
func testPreRemove(t *testing.T) {
	doc := NewDocument()
	parent := NewElement("div", doc)
	child := NewElement("span", doc)
	parent.AppendChild(child)

	t.Run("Valid removal", func(t *testing.T) {
		result, err := preRemove(child, parent)
		if err != nil {
			t.Errorf("preRemove failed: %v", err)
		}
		if result != child {
			t.Error("preRemove should return the removed child")
		}
		if child.ParentNode() != nil {
			t.Error("Child should be detached from parent")
		}
		if parent.HasChildNodes() {
			t.Error("Parent should have no children")
		}
	})

	t.Run("Invalid parent", func(t *testing.T) {
		child2 := NewElement("p", doc)
		wrongParent := NewElement("wrong", doc)

		_, err := preRemove(child2, wrongParent)
		if err == nil {
			t.Error("preRemove should fail when child's parent doesn't match")
		}
	})
}

// testRemove tests the "remove" algorithm
func testRemove(t *testing.T) {
	doc := NewDocument()
	parent := NewElement("div", doc)
	child := NewElement("span", doc)
	grandchild := NewText("text", doc)

	parent.AppendChild(child)
	child.AppendChild(grandchild)

	t.Run("Remove with descendants", func(t *testing.T) {
		removeNode(child, false)

		if child.ParentNode() != nil {
			t.Error("Child should be detached")
		}
		if parent.HasChildNodes() {
			t.Error("Parent should have no children")
		}
		// Grandchild should still be parented to child
		if grandchild.ParentNode() != child {
			t.Error("Grandchild should still be parented to child")
		}
	})

	t.Run("Remove already detached node", func(t *testing.T) {
		// Should be safe to call on already detached node
		removeNode(child, false)
		// Should not panic or error
	})
}

// testReplaceAll tests the "replace all" algorithm
func testReplaceAll(t *testing.T) {
	doc := NewDocument()
	parent := NewElement("div", doc)
	child1 := NewElement("span", doc)
	child2 := NewElement("p", doc)

	parent.AppendChild(child1)
	parent.AppendChild(child2)

	t.Run("Replace with single node", func(t *testing.T) {
		newChild := NewElement("em", doc)
		replaceAllWithNode(newChild, parent)

		if parent.ChildNodes().Length() != 1 {
			t.Error("Parent should have exactly one child")
		}
		if parent.FirstChild() != newChild {
			t.Error("Parent should contain the new child")
		}
		if child1.ParentNode() != nil || child2.ParentNode() != nil {
			t.Error("Old children should be detached")
		}
	})

	t.Run("Replace with null", func(t *testing.T) {
		replaceAllWithNode(nil, parent)

		if parent.HasChildNodes() {
			t.Error("Parent should have no children")
		}
	})

	t.Run("Replace with DocumentFragment", func(t *testing.T) {
		fragment := NewDocumentFragment(doc)
		newChild1 := NewElement("strong", doc)
		newChild2 := NewElement("em", doc)
		fragment.AppendChild(newChild1)
		fragment.AppendChild(newChild2)

		replaceAllWithNode(fragment, parent)

		if parent.ChildNodes().Length() != 2 {
			t.Error("Parent should have two children")
		}
		if parent.FirstChild() != newChild1 || parent.LastChild() != newChild2 {
			t.Error("Fragment children should be moved to parent")
		}
		if fragment.HasChildNodes() {
			t.Error("Fragment should be empty")
		}
	})
}

// testDocumentFragmentInsertion tests DocumentFragment-specific insertion behavior
func testDocumentFragmentInsertion(t *testing.T) {
	doc := NewDocument()
	parent := NewElement("div", doc)

	t.Run("Fragment children are moved", func(t *testing.T) {
		fragment := NewDocumentFragment(doc)
		child1 := NewElement("span", doc)
		child2 := NewElement("p", doc)
		child3 := NewText("text", doc)

		fragment.AppendChild(child1)
		fragment.AppendChild(child2)
		fragment.AppendChild(child3)

		parent.AppendChild(fragment)

		// Fragment should be empty
		if fragment.HasChildNodes() {
			t.Error("Fragment should be empty after insertion")
		}

		// Children should be moved to parent
		if parent.ChildNodes().Length() != 3 {
			t.Errorf("Parent should have 3 children, got %d", parent.ChildNodes().Length())
		}

		children := parent.ChildNodes()
		if children.Item(0) != child1 || children.Item(1) != child2 || children.Item(2) != child3 {
			t.Error("Fragment children should be in correct order")
		}

		// All children should be parented to parent
		if child1.ParentNode() != parent || child2.ParentNode() != parent || child3.ParentNode() != parent {
			t.Error("All fragment children should be parented to parent")
		}
	})

	t.Run("Empty fragment insertion", func(t *testing.T) {
		emptyFragment := NewDocumentFragment(doc)
		initialChildCount := parent.ChildNodes().Length()

		parent.AppendChild(emptyFragment)

		if parent.ChildNodes().Length() != initialChildCount {
			t.Error("Empty fragment insertion should not change parent")
		}
	})
}

// testMutationObserverIntegration tests integration with MutationObserver
func testMutationObserverIntegration(t *testing.T) {
	doc := NewDocument()
	parent := NewElement("div", doc)

	t.Run("Observer receives insertion notifications", func(t *testing.T) {
		var records []*MutationRecord
		observer := NewMutationObserver(func(r []*MutationRecord, o *MutationObserver) {
			records = append(records, r...)
		})

		observer.Observe(parent, MutationObserverInit{ChildList: true})

		child := NewElement("span", doc)
		parent.AppendChild(child)

		// Process pending mutation records
		registry := doc.getObserverRegistry()
		registry.ProcessMutationObservers()

		if len(records) == 0 {
			t.Error("Observer should receive mutation records")
		}

		record := records[0]
		if record.Type != "childList" {
			t.Error("Record type should be childList")
		}
		if record.Target != parent {
			t.Error("Record target should be parent")
		}
		if len(record.AddedNodes) != 1 || record.AddedNodes[0] != child {
			t.Error("Record should show added child")
		}

		observer.Disconnect()
	})

	t.Run("Observer receives removal notifications", func(t *testing.T) {
		child := NewElement("span", doc)
		parent.AppendChild(child)

		var records []*MutationRecord
		observer := NewMutationObserver(func(r []*MutationRecord, o *MutationObserver) {
			records = append(records, r...)
		})

		observer.Observe(parent, MutationObserverInit{ChildList: true})

		parent.RemoveChild(child)

		// Process pending mutation records
		registry := doc.getObserverRegistry()
		registry.ProcessMutationObservers()

		if len(records) == 0 {
			t.Error("Observer should receive mutation records")
		}

		record := records[0]
		if record.Type != "childList" {
			t.Error("Record type should be childList")
		}
		if len(record.RemovedNodes) != 1 || record.RemovedNodes[0] != child {
			t.Error("Record should show removed child")
		}

		observer.Disconnect()
	})

	t.Run("Suppressed observers don't receive notifications", func(t *testing.T) {
		var records []*MutationRecord
		observer := NewMutationObserver(func(r []*MutationRecord, o *MutationObserver) {
			records = append(records, r...)
		})

		observer.Observe(parent, MutationObserverInit{ChildList: true})

		child := NewElement("span", doc)
		insertNode(child, parent, nil, true) // suppress observers

		// Process pending mutation records
		registry := doc.getObserverRegistry()
		registry.ProcessMutationObservers()

		if len(records) > 0 {
			t.Error("Suppressed observers should not receive notifications")
		}

		observer.Disconnect()
	})
}

// TestSpecComplianceWithExistingTests verifies new algorithms work with existing test cases
func TestSpecComplianceWithExistingTests(t *testing.T) {
	t.Run("Existing AppendChild behavior", func(t *testing.T) {
		doc := NewDocument()
		parent := NewElement("div", doc)
		child := NewElement("span", doc)

		result := parent.AppendChild(child)

		if result != child {
			t.Error("AppendChild should return the appended child")
		}
		if child.ParentNode() != parent {
			t.Error("Child should be parented to parent")
		}
		if parent.FirstChild() != child {
			t.Error("Child should be first child")
		}
	})

	t.Run("Existing RemoveChild behavior", func(t *testing.T) {
		doc := NewDocument()
		parent := NewElement("div", doc)
		child := NewElement("span", doc)
		parent.AppendChild(child)

		result := parent.RemoveChild(child)

		if result != child {
			t.Error("RemoveChild should return the removed child")
		}
		if child.ParentNode() != nil {
			t.Error("Child should be detached")
		}
		if parent.HasChildNodes() {
			t.Error("Parent should have no children")
		}
	})

	t.Run("Exception handling", func(t *testing.T) {
		doc := NewDocument()
		text := NewText("text", doc)

		// Should panic when trying to insert text into document
		defer func() {
			if r := recover(); r == nil {
				t.Error("Should panic when inserting text into document")
			}
		}()

		doc.AppendChild(text)
	})
}
