package dom

import (
	"testing"
)

func TestParentNodeProperties(t *testing.T) {
	doc := NewDocument()

	t.Run("Element ParentNode properties", func(t *testing.T) {
		// Create parent element
		parent := doc.CreateElement("div")

		// Initially no element children
		if parent.FirstElementChild() != nil {
			t.Error("Expected FirstElementChild to be nil for empty element")
		}
		if parent.LastElementChild() != nil {
			t.Error("Expected LastElementChild to be nil for empty element")
		}
		if parent.ChildElementCount() != 0 {
			t.Errorf("Expected ChildElementCount to be 0, got %d", parent.ChildElementCount())
		}

		// Add text node (should not affect element counts)
		textNode := doc.CreateTextNode("text")
		parent.AppendChild(textNode)

		if parent.FirstElementChild() != nil {
			t.Error("Expected FirstElementChild to be nil with only text nodes")
		}
		if parent.LastElementChild() != nil {
			t.Error("Expected LastElementChild to be nil with only text nodes")
		}
		if parent.ChildElementCount() != 0 {
			t.Errorf("Expected ChildElementCount to be 0 with only text nodes, got %d", parent.ChildElementCount())
		}

		// Add first element child
		elem1 := doc.CreateElement("span")
		parent.AppendChild(elem1)

		if parent.FirstElementChild() != elem1 {
			t.Error("Expected FirstElementChild to be elem1")
		}
		if parent.LastElementChild() != elem1 {
			t.Error("Expected LastElementChild to be elem1")
		}
		if parent.ChildElementCount() != 1 {
			t.Errorf("Expected ChildElementCount to be 1, got %d", parent.ChildElementCount())
		}

		// Add another text node
		textNode2 := doc.CreateTextNode("more text")
		parent.AppendChild(textNode2)

		// Add second element child
		elem2 := doc.CreateElement("p")
		parent.AppendChild(elem2)

		if parent.FirstElementChild() != elem1 {
			t.Error("Expected FirstElementChild to be elem1")
		}
		if parent.LastElementChild() != elem2 {
			t.Error("Expected LastElementChild to be elem2")
		}
		if parent.ChildElementCount() != 2 {
			t.Errorf("Expected ChildElementCount to be 2, got %d", parent.ChildElementCount())
		}

		// Verify Children() HTMLCollection
		children := parent.Children()
		if children.Length() != 2 {
			t.Errorf("Expected Children().Length() to be 2, got %d", children.Length())
		}
		if children.Item(0) != elem1 {
			t.Error("Expected first child in collection to be elem1")
		}
		if children.Item(1) != elem2 {
			t.Error("Expected second child in collection to be elem2")
		}
	})

	t.Run("Document ParentNode properties", func(t *testing.T) {
		// Document with no element children
		if doc.FirstElementChild() != nil {
			t.Error("Expected Document FirstElementChild to be nil initially")
		}
		if doc.LastElementChild() != nil {
			t.Error("Expected Document LastElementChild to be nil initially")
		}
		if doc.ChildElementCount() != 0 {
			t.Errorf("Expected Document ChildElementCount to be 0, got %d", doc.ChildElementCount())
		}

		// Add document element
		html := doc.CreateElement("html")
		doc.AppendChild(html)

		if doc.FirstElementChild() != html {
			t.Error("Expected Document FirstElementChild to be html element")
		}
		if doc.LastElementChild() != html {
			t.Error("Expected Document LastElementChild to be html element")
		}
		if doc.ChildElementCount() != 1 {
			t.Errorf("Expected Document ChildElementCount to be 1, got %d", doc.ChildElementCount())
		}
	})

	t.Run("DocumentFragment ParentNode properties", func(t *testing.T) {
		fragment := doc.CreateDocumentFragment()

		// Initially no element children
		if fragment.FirstElementChild() != nil {
			t.Error("Expected DocumentFragment FirstElementChild to be nil initially")
		}
		if fragment.LastElementChild() != nil {
			t.Error("Expected DocumentFragment LastElementChild to be nil initially")
		}
		if fragment.ChildElementCount() != 0 {
			t.Errorf("Expected DocumentFragment ChildElementCount to be 0, got %d", fragment.ChildElementCount())
		}

		// Add elements
		elem1 := doc.CreateElement("div")
		elem2 := doc.CreateElement("span")
		fragment.AppendChild(elem1)
		fragment.AppendChild(elem2)

		if fragment.FirstElementChild() != elem1 {
			t.Error("Expected DocumentFragment FirstElementChild to be elem1")
		}
		if fragment.LastElementChild() != elem2 {
			t.Error("Expected DocumentFragment LastElementChild to be elem2")
		}
		if fragment.ChildElementCount() != 2 {
			t.Errorf("Expected DocumentFragment ChildElementCount to be 2, got %d", fragment.ChildElementCount())
		}
	})
}

func TestParentNodePrepend(t *testing.T) {
	doc := NewDocument()
	parent := doc.CreateElement("div")

	t.Run("Prepend single string", func(t *testing.T) {
		err := parent.Prepend("Hello")
		if err != nil {
			t.Errorf("Prepend failed: %v", err)
		}

		if parent.FirstChild().NodeType() != TextNode {
			t.Error("Expected first child to be a text node")
		}
		if parent.FirstChild().NodeValue() != "Hello" {
			t.Errorf("Expected text content 'Hello', got '%s'", parent.FirstChild().NodeValue())
		}
	})

	t.Run("Prepend element node", func(t *testing.T) {
		elem := doc.CreateElement("span")
		err := parent.Prepend(elem)
		if err != nil {
			t.Errorf("Prepend failed: %v", err)
		}

		if parent.FirstChild() != elem {
			t.Error("Expected span element to be first child")
		}
		if parent.FirstChild().NodeName() != "span" {
			t.Errorf("Expected first child node name 'span', got '%s'", parent.FirstChild().NodeName())
		}
	})

	t.Run("Prepend multiple nodes", func(t *testing.T) {
		parent2 := doc.CreateElement("div")
		elem1 := doc.CreateElement("p")
		elem2 := doc.CreateElement("strong")

		err := parent2.Prepend(elem1, "text", elem2)
		if err != nil {
			t.Errorf("Prepend failed: %v", err)
		}

		// Should create a document fragment with all nodes
		if len(parent2.ChildNodes()) != 3 {
			t.Errorf("Expected 3 children, got %d", len(parent2.ChildNodes()))
		}

		if parent2.ChildNodes()[0] != elem1 {
			t.Error("Expected first child to be elem1")
		}
		if parent2.ChildNodes()[1].NodeValue() != "text" {
			t.Error("Expected second child to be text node with 'text'")
		}
		if parent2.ChildNodes()[2] != elem2 {
			t.Error("Expected third child to be elem2")
		}
	})
}

func TestParentNodeAppend(t *testing.T) {
	doc := NewDocument()
	parent := doc.CreateElement("div")

	t.Run("Append single string", func(t *testing.T) {
		err := parent.Append("World")
		if err != nil {
			t.Errorf("Append failed: %v", err)
		}

		if parent.LastChild().NodeType() != TextNode {
			t.Error("Expected last child to be a text node")
		}
		if parent.LastChild().NodeValue() != "World" {
			t.Errorf("Expected text content 'World', got '%s'", parent.LastChild().NodeValue())
		}
	})

	t.Run("Append element node", func(t *testing.T) {
		elem := doc.CreateElement("em")
		err := parent.Append(elem)
		if err != nil {
			t.Errorf("Append failed: %v", err)
		}

		if parent.LastChild() != elem {
			t.Error("Expected em element to be last child")
		}
		if parent.LastChild().NodeName() != "em" {
			t.Errorf("Expected last child node name 'em', got '%s'", parent.LastChild().NodeName())
		}
	})

	t.Run("Append multiple nodes", func(t *testing.T) {
		parent2 := doc.CreateElement("div")
		elem1 := doc.CreateElement("h1")
		elem2 := doc.CreateElement("h2")

		err := parent2.Append(elem1, "heading", elem2)
		if err != nil {
			t.Errorf("Append failed: %v", err)
		}

		if len(parent2.ChildNodes()) != 3 {
			t.Errorf("Expected 3 children, got %d", len(parent2.ChildNodes()))
		}

		if parent2.ChildNodes()[0] != elem1 {
			t.Error("Expected first child to be elem1")
		}
		if parent2.ChildNodes()[1].NodeValue() != "heading" {
			t.Error("Expected second child to be text node with 'heading'")
		}
		if parent2.ChildNodes()[2] != elem2 {
			t.Error("Expected third child to be elem2")
		}
	})
}

func TestParentNodeReplaceChildren(t *testing.T) {
	doc := NewDocument()
	parent := doc.CreateElement("div")

	// Add some initial children
	parent.AppendChild(doc.CreateElement("old1"))
	parent.AppendChild(doc.CreateTextNode("old text"))
	parent.AppendChild(doc.CreateElement("old2"))

	t.Run("ReplaceChildren with new nodes", func(t *testing.T) {
		elem1 := doc.CreateElement("new1")
		elem2 := doc.CreateElement("new2")

		err := parent.ReplaceChildren(elem1, "new text", elem2)
		if err != nil {
			t.Errorf("ReplaceChildren failed: %v", err)
		}

		if len(parent.ChildNodes()) != 3 {
			t.Errorf("Expected 3 children after replacement, got %d", len(parent.ChildNodes()))
		}

		if parent.ChildNodes()[0] != elem1 {
			t.Error("Expected first child to be new elem1")
		}
		if parent.ChildNodes()[1].NodeValue() != "new text" {
			t.Error("Expected second child to be new text node")
		}
		if parent.ChildNodes()[2] != elem2 {
			t.Error("Expected third child to be new elem2")
		}
	})

	t.Run("ReplaceChildren with empty", func(t *testing.T) {
		err := parent.ReplaceChildren()
		if err != nil {
			t.Errorf("ReplaceChildren with no args failed: %v", err)
		}

		if len(parent.ChildNodes()) != 0 {
			t.Errorf("Expected 0 children after empty replacement, got %d", len(parent.ChildNodes()))
		}
	})
}

func TestParentNodeMoveBefore(t *testing.T) {
	doc := NewDocument()
	parent := doc.CreateElement("div")

	// Create some child elements
	elem1 := doc.CreateElement("span")
	elem2 := doc.CreateElement("p")
	elem3 := doc.CreateElement("strong")

	parent.AppendChild(elem1)
	parent.AppendChild(elem2)
	parent.AppendChild(elem3)

	t.Run("MoveBefore with valid nodes", func(t *testing.T) {
		// Move elem3 before elem1
		err := parent.MoveBefore(elem3, elem1)
		if err != nil {
			t.Errorf("MoveBefore failed: %v", err)
		}

		// Order should now be: elem3, elem1, elem2
		if parent.ChildNodes()[0] != elem3 {
			t.Error("Expected elem3 to be first after move")
		}
		if parent.ChildNodes()[1] != elem1 {
			t.Error("Expected elem1 to be second after move")
		}
		if parent.ChildNodes()[2] != elem2 {
			t.Error("Expected elem2 to be third after move")
		}
	})

	t.Run("MoveBefore with null child (append)", func(t *testing.T) {
		// Move elem1 to end
		err := parent.MoveBefore(elem1, nil)
		if err != nil {
			t.Errorf("MoveBefore to end failed: %v", err)
		}

		// Order should now be: elem3, elem2, elem1
		if parent.ChildNodes()[0] != elem3 {
			t.Error("Expected elem3 to be first")
		}
		if parent.ChildNodes()[1] != elem2 {
			t.Error("Expected elem2 to be second")
		}
		if parent.ChildNodes()[2] != elem1 {
			t.Error("Expected elem1 to be last after move to end")
		}
	})

	t.Run("MoveBefore with node as its own reference", func(t *testing.T) {
		// Initial order: elem3, elem2, elem1
		t.Logf("Before MoveBefore(elem2, elem2): %v, %v, %v",
			parent.ChildNodes()[0].NodeName(),
			parent.ChildNodes()[1].NodeName(),
			parent.ChildNodes()[2].NodeName())

		// Move elem2 before itself (should move to next sibling's position)
		err := parent.MoveBefore(elem2, elem2)
		if err != nil {
			t.Errorf("MoveBefore with self reference failed: %v", err)
		}

		t.Logf("After MoveBefore(elem2, elem2): %v, %v, %v",
			parent.ChildNodes()[0].NodeName(),
			parent.ChildNodes()[1].NodeName(),
			parent.ChildNodes()[2].NodeName())

		// Should move elem2 before its next sibling (elem1)
		// Order should be: elem3, elem2, elem1
		if parent.ChildNodes()[0] != elem3 {
			t.Error("Expected elem3 to be first")
		}
		if parent.ChildNodes()[1] != elem2 {
			t.Error("Expected elem2 to be second")
		}
		if parent.ChildNodes()[2] != elem1 {
			t.Error("Expected elem1 to be third")
		}
	})

	t.Run("MoveBefore with external node", func(t *testing.T) {
		// Create an external element and move it in
		external := doc.CreateElement("external")

		err := parent.MoveBefore(external, elem2)
		if err != nil {
			t.Errorf("MoveBefore with external node failed: %v", err)
		}

		// external should now be parent of elem2
		if len(parent.ChildNodes()) != 4 {
			t.Errorf("Expected 4 children after external move, got %d", len(parent.ChildNodes()))
		}

		// Find external's position
		var externalIndex = -1
		for i, child := range parent.ChildNodes() {
			if child == external {
				externalIndex = i
				break
			}
		}

		if externalIndex == -1 {
			t.Error("External element not found in parent")
		}

		// elem2 should be after external
		if externalIndex+1 >= len(parent.ChildNodes()) || parent.ChildNodes()[externalIndex+1] != elem2 {
			t.Error("Expected elem2 to be after external element")
		}
	})
}

func TestParentNodeQuerySelector(t *testing.T) {
	doc := NewDocument()
	parent := doc.CreateElement("div")

	// Create test structure
	span1 := doc.CreateElement("span")
	span1.SetAttribute("id", "first")
	span1.SetAttribute("class", "highlight")

	span2 := doc.CreateElement("span")
	span2.SetAttribute("class", "highlight bold")

	p := doc.CreateElement("p")
	p.SetAttribute("id", "paragraph")

	strong := doc.CreateElement("strong")
	strong.SetAttribute("class", "bold")

	parent.AppendChild(span1)
	parent.AppendChild(span2)
	parent.AppendChild(p)
	p.AppendChild(strong) // nested element

	t.Run("QuerySelector by tag name", func(t *testing.T) {
		result := parent.QuerySelector("span")
		if result != span1 {
			t.Errorf("Expected first span element, got %v (span1=%v)", result, span1)
		}

		result = parent.QuerySelector("p")
		if result != p {
			t.Errorf("Expected p element, got %v", result)
		}

		result = parent.QuerySelector("strong")
		if result != strong {
			t.Errorf("Expected strong element (nested), got %v", result)
		}
	})

	t.Run("QuerySelector by ID", func(t *testing.T) {
		result := parent.QuerySelector("#first")
		if result != span1 {
			t.Error("Expected span with id 'first'")
		}

		result = parent.QuerySelector("#paragraph")
		if result != p {
			t.Error("Expected p with id 'paragraph'")
		}

		result = parent.QuerySelector("#nonexistent")
		if result != nil {
			t.Error("Expected nil for nonexistent ID")
		}
	})

	t.Run("QuerySelector by class", func(t *testing.T) {
		result := parent.QuerySelector(".highlight")
		if result != span1 {
			t.Error("Expected first element with 'highlight' class")
		}

		result = parent.QuerySelector(".bold")
		if result != span2 {
			t.Error("Expected first element with 'bold' class")
		}

		result = parent.QuerySelector(".nonexistent")
		if result != nil {
			t.Error("Expected nil for nonexistent class")
		}
	})

	t.Run("QuerySelector with compound selectors", func(t *testing.T) {
		result := parent.QuerySelector("span.highlight")
		if result != span1 {
			t.Error("Expected span with highlight class")
		}

		result = parent.QuerySelector("span#first")
		if result != span1 {
			t.Error("Expected span with ID first")
		}

		result = parent.QuerySelector("p#paragraph")
		if result != p {
			t.Error("Expected p with ID paragraph")
		}
	})

	t.Run("QuerySelector wildcard", func(t *testing.T) {
		result := parent.QuerySelector("*")
		if result != span1 {
			t.Error("Expected first element for wildcard selector")
		}
	})
}

func TestParentNodeQuerySelectorAll(t *testing.T) {
	doc := NewDocument()
	parent := doc.CreateElement("div")

	// Create test structure with multiple matching elements
	span1 := doc.CreateElement("span")
	span1.SetAttribute("class", "highlight")

	span2 := doc.CreateElement("span")
	span2.SetAttribute("class", "highlight")

	span3 := doc.CreateElement("span")
	span3.SetAttribute("class", "other")

	p := doc.CreateElement("p")
	strong := doc.CreateElement("strong")
	strong.SetAttribute("class", "highlight")

	parent.AppendChild(span1)
	parent.AppendChild(span2)
	parent.AppendChild(span3)
	parent.AppendChild(p)
	p.AppendChild(strong) // nested element

	t.Run("QuerySelectorAll by tag name", func(t *testing.T) {
		results := parent.QuerySelectorAll("span")
		if len(results) != 3 {
			t.Errorf("Expected 3 span elements, got %d", len(results))
		}

		if results[0] != span1 || results[1] != span2 || results[2] != span3 {
			t.Error("Expected spans in document order")
		}
	})

	t.Run("QuerySelectorAll by class", func(t *testing.T) {
		results := parent.QuerySelectorAll(".highlight")
		if len(results) != 3 {
			t.Errorf("Expected 3 elements with 'highlight' class, got %d", len(results))
		}

		if results[0] != span1 || results[1] != span2 || results[2] != strong {
			t.Error("Expected elements with highlight class in document order")
		}
	})

	t.Run("QuerySelectorAll with no matches", func(t *testing.T) {
		results := parent.QuerySelectorAll(".nonexistent")
		if len(results) != 0 {
			t.Errorf("Expected 0 elements for nonexistent class, got %d", len(results))
		}
	})

	t.Run("QuerySelectorAll wildcard", func(t *testing.T) {
		results := parent.QuerySelectorAll("*")
		if len(results) != 5 { // 4 direct children + 1 nested strong
			t.Errorf("Expected 5 elements for wildcard, got %d", len(results))
		}
	})
}

func TestConvertNodesToNode(t *testing.T) {
	doc := NewDocument()

	t.Run("Convert empty slice", func(t *testing.T) {
		result, err := convertNodesToNode([]interface{}{}, doc)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if result != nil {
			t.Error("Expected nil result for empty slice")
		}
	})

	t.Run("Convert single string", func(t *testing.T) {
		result, err := convertNodesToNode([]interface{}{"hello"}, doc)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if result.NodeType() != TextNode {
			t.Error("Expected text node for single string")
		}
		if result.NodeValue() != "hello" {
			t.Errorf("Expected 'hello', got '%s'", result.NodeValue())
		}
	})

	t.Run("Convert single node", func(t *testing.T) {
		elem := doc.CreateElement("div")
		result, err := convertNodesToNode([]interface{}{elem}, doc)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if result != elem {
			t.Error("Expected same element for single node")
		}
	})

	t.Run("Convert multiple mixed nodes", func(t *testing.T) {
		elem := doc.CreateElement("span")
		result, err := convertNodesToNode([]interface{}{elem, "text", doc.CreateElement("p")}, doc)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if result.NodeType() != DocumentFragmentNode {
			t.Error("Expected DocumentFragment for multiple nodes")
		}

		fragment := result.(*DocumentFragment)
		if len(fragment.ChildNodes()) != 3 {
			t.Errorf("Expected 3 children in fragment, got %d", len(fragment.ChildNodes()))
		}

		if fragment.ChildNodes()[0] != elem {
			t.Error("Expected first child to be span element")
		}
		if fragment.ChildNodes()[1].NodeValue() != "text" {
			t.Error("Expected second child to be text node")
		}
		if fragment.ChildNodes()[2].NodeName() != "p" {
			t.Error("Expected third child to be p element")
		}
	})

	t.Run("Convert invalid type", func(t *testing.T) {
		_, err := convertNodesToNode([]interface{}{123}, doc)
		if err == nil {
			t.Error("Expected error for invalid type")
		}
	})
}

// Test spec compliance for the "converting nodes into a node" algorithm
func TestConvertNodesToNodeSpecCompliance(t *testing.T) {
	doc := NewDocument()

	t.Run("Spec step 1: node is null initially", func(t *testing.T) {
		// Empty nodes should return null
		result, err := convertNodesToNode([]interface{}{}, doc)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if result != nil {
			t.Error("Expected null/nil for empty nodes per spec step 1")
		}
	})

	t.Run("Spec step 2: replace strings with Text nodes", func(t *testing.T) {
		result, err := convertNodesToNode([]interface{}{"test string"}, doc)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if result.NodeType() != TextNode {
			t.Error("String should be converted to Text node per spec step 2")
		}
		if result.NodeValue() != "test string" {
			t.Errorf("Text node should contain string data, got '%s'", result.NodeValue())
		}
		if result.OwnerDocument() != doc {
			t.Error("Text node should have correct owner document per spec step 2")
		}
	})

	t.Run("Spec step 3: single node case", func(t *testing.T) {
		elem := doc.CreateElement("div")
		result, err := convertNodesToNode([]interface{}{elem}, doc)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if result != elem {
			t.Error("Single node should be returned as-is per spec step 3")
		}
	})

	t.Run("Spec step 4: multiple nodes create DocumentFragment", func(t *testing.T) {
		elem1 := doc.CreateElement("div")
		elem2 := doc.CreateElement("span")

		result, err := convertNodesToNode([]interface{}{elem1, "text", elem2}, doc)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if result.NodeType() != DocumentFragmentNode {
			t.Error("Multiple nodes should create DocumentFragment per spec step 4")
		}

		fragment := result.(*DocumentFragment)
		if fragment.OwnerDocument() != doc {
			t.Error("DocumentFragment should have correct owner document per spec step 4")
		}

		// Check that all nodes are appended in order
		if len(fragment.ChildNodes()) != 3 {
			t.Errorf("DocumentFragment should contain all nodes, expected 3, got %d", len(fragment.ChildNodes()))
		}

		if fragment.ChildNodes()[0] != elem1 {
			t.Error("First child should be elem1")
		}
		if fragment.ChildNodes()[1].NodeType() != TextNode || fragment.ChildNodes()[1].NodeValue() != "text" {
			t.Error("Second child should be text node with 'text'")
		}
		if fragment.ChildNodes()[2] != elem2 {
			t.Error("Third child should be elem2")
		}
	})
}
