package dom

import (
	"testing"
)

func TestNodeFilterConstants(t *testing.T) {
	// Test show constants
	if NodeFilterShowAll != 0xFFFFFFFF {
		t.Errorf("Expected NodeFilterShowAll to be 0xFFFFFFFF, got 0x%x", NodeFilterShowAll)
	}

	if NodeFilterShowElement != 0x1 {
		t.Errorf("Expected NodeFilterShowElement to be 0x1, got 0x%x", NodeFilterShowElement)
	}

	if NodeFilterShowText != 0x4 {
		t.Errorf("Expected NodeFilterShowText to be 0x4, got 0x%x", NodeFilterShowText)
	}

	// Test accept constants
	if NodeFilterAccept != 1 {
		t.Errorf("Expected NodeFilterAccept to be 1, got %d", NodeFilterAccept)
	}

	if NodeFilterReject != 2 {
		t.Errorf("Expected NodeFilterReject to be 2, got %d", NodeFilterReject)
	}

	if NodeFilterSkip != 3 {
		t.Errorf("Expected NodeFilterSkip to be 3, got %d", NodeFilterSkip)
	}
}

func TestNodeFilterFunc(t *testing.T) {
	doc := NewDocument()
	elem := doc.CreateElement("div")

	// Test function that accepts all nodes
	acceptAll := NodeFilterFunc(func(node Node) int {
		return NodeFilterAccept
	})

	result := acceptAll.AcceptNode(elem)
	if result != NodeFilterAccept {
		t.Errorf("Expected NodeFilterAccept, got %d", result)
	}

	// Test function that rejects all nodes
	rejectAll := NodeFilterFunc(func(node Node) int {
		return NodeFilterReject
	})

	result = rejectAll.AcceptNode(elem)
	if result != NodeFilterReject {
		t.Errorf("Expected NodeFilterReject, got %d", result)
	}

	// Test function that skips all nodes
	skipAll := NodeFilterFunc(func(node Node) int {
		return NodeFilterSkip
	})

	result = skipAll.AcceptNode(elem)
	if result != NodeFilterSkip {
		t.Errorf("Expected NodeFilterSkip, got %d", result)
	}
}

func TestDefaultNodeFilter(t *testing.T) {
	doc := NewDocument()
	elem := doc.CreateElement("div")

	result := DefaultNodeFilter.AcceptNode(elem)
	if result != NodeFilterAccept {
		t.Errorf("Expected DefaultNodeFilter to accept all nodes, got %d", result)
	}

	text := doc.CreateTextNode("hello")
	result = DefaultNodeFilter.AcceptNode(text)
	if result != NodeFilterAccept {
		t.Errorf("Expected DefaultNodeFilter to accept text nodes, got %d", result)
	}
}

func TestWhatToShowFilter(t *testing.T) {
	doc := NewDocument()

	// Test element filter
	elementFilter := whatToShowFilter(NodeFilterShowElement)

	elem := doc.CreateElement("div")
	result := elementFilter.AcceptNode(elem)
	if result != NodeFilterAccept {
		t.Errorf("Expected element filter to accept element, got %d", result)
	}

	text := doc.CreateTextNode("hello")
	result = elementFilter.AcceptNode(text)
	if result != NodeFilterSkip {
		t.Errorf("Expected element filter to skip text node, got %d", result)
	}

	// Test text filter
	textFilter := whatToShowFilter(NodeFilterShowText)

	result = textFilter.AcceptNode(text)
	if result != NodeFilterAccept {
		t.Errorf("Expected text filter to accept text node, got %d", result)
	}

	result = textFilter.AcceptNode(elem)
	if result != NodeFilterSkip {
		t.Errorf("Expected text filter to skip element, got %d", result)
	}

	// Test combined filter (elements and text)
	combinedFilter := whatToShowFilter(NodeFilterShowElement | NodeFilterShowText)

	result = combinedFilter.AcceptNode(elem)
	if result != NodeFilterAccept {
		t.Errorf("Expected combined filter to accept element, got %d", result)
	}

	result = combinedFilter.AcceptNode(text)
	if result != NodeFilterAccept {
		t.Errorf("Expected combined filter to accept text node, got %d", result)
	}

	comment := doc.CreateComment("test")
	result = combinedFilter.AcceptNode(comment)
	if result != NodeFilterSkip {
		t.Errorf("Expected combined filter to skip comment, got %d", result)
	}
}

func TestCombinedFilter(t *testing.T) {
	doc := NewDocument()
	elem := doc.CreateElement("div")
	elem.SetAttribute("class", "test")

	// Test whatToShow only
	filter := combinedFilter(NodeFilterShowElement, nil)
	result := filter.AcceptNode(elem)
	if result != NodeFilterAccept {
		t.Errorf("Expected combined filter with no custom filter to accept element, got %d", result)
	}

	// Test whatToShow with custom filter that accepts
	customFilter := NodeFilterFunc(func(node Node) int {
		if elem, ok := node.(*Element); ok {
			if elem.GetAttribute("class") == "test" {
				return NodeFilterAccept
			}
		}
		return NodeFilterReject
	})

	filter = combinedFilter(NodeFilterShowElement, customFilter)
	result = filter.AcceptNode(elem)
	if result != NodeFilterAccept {
		t.Errorf("Expected combined filter with accepting custom filter to accept element, got %d", result)
	}

	// Test whatToShow with custom filter that rejects
	rejectingFilter := NodeFilterFunc(func(node Node) int {
		return NodeFilterReject
	})

	filter = combinedFilter(NodeFilterShowElement, rejectingFilter)
	result = filter.AcceptNode(elem)
	if result != NodeFilterReject {
		t.Errorf("Expected combined filter with rejecting custom filter to reject element, got %d", result)
	}

	// Test whatToShow that skips
	text := doc.CreateTextNode("hello")
	filter = combinedFilter(NodeFilterShowElement, customFilter)
	result = filter.AcceptNode(text)
	if result != NodeFilterSkip {
		t.Errorf("Expected combined filter to skip text node when only showing elements, got %d", result)
	}
}

func TestNodeFilterWithAllNodeTypes(t *testing.T) {
	doc := NewDocument()

	// Create nodes of different types
	elem := doc.CreateElement("div")
	text := doc.CreateTextNode("hello")
	comment := doc.CreateComment("comment")
	fragment := doc.CreateDocumentFragment()

	// CDATA and ProcessingInstruction require special handling
	cdata, _ := doc.CreateCDATASection("cdata")
	pi, _ := doc.CreateProcessingInstruction("target", "data")

	nodes := []struct {
		node     Node
		nodeType NodeType
		showBit  uint32
	}{
		{elem, ElementNode, NodeFilterShowElement},
		{text, TextNode, NodeFilterShowText},
		{comment, CommentNode, NodeFilterShowComment},
		{fragment, DocumentFragmentNode, NodeFilterShowDocumentFragment},
		{doc, DocumentNode, NodeFilterShowDocument},
	}

	if cdata != nil {
		nodes = append(nodes, struct {
			node     Node
			nodeType NodeType
			showBit  uint32
		}{cdata, CDataSectionNode, NodeFilterShowCDataSection})
	}

	if pi != nil {
		nodes = append(nodes, struct {
			node     Node
			nodeType NodeType
			showBit  uint32
		}{pi, ProcessingInstructionNode, NodeFilterShowProcessingInstruction})
	}

	for _, testCase := range nodes {
		// Test filter that shows this specific node type
		filter := whatToShowFilter(testCase.showBit)
		result := filter.AcceptNode(testCase.node)
		if result != NodeFilterAccept {
			t.Errorf("Expected filter to accept node type %d, got %d", testCase.nodeType, result)
		}

		// Test filter that shows all except this node type
		excludeFilter := whatToShowFilter(NodeFilterShowAll &^ testCase.showBit)
		result = excludeFilter.AcceptNode(testCase.node)
		if result != NodeFilterSkip {
			t.Errorf("Expected exclude filter to skip node type %d, got %d", testCase.nodeType, result)
		}
	}
}

func TestNodeFilterWithComplexCustomFilter(t *testing.T) {
	doc := NewDocument()

	// Create a complex DOM structure
	html := doc.CreateElement("html")
	body := doc.CreateElement("body")
	div1 := doc.CreateElement("div")
	div1.SetAttribute("id", "container")
	div2 := doc.CreateElement("div")
	div2.SetAttribute("class", "content")
	text1 := doc.CreateTextNode("Hello")
	text2 := doc.CreateTextNode("World")

	html.AppendChild(body)
	body.AppendChild(div1)
	div1.AppendChild(div2)
	div1.AppendChild(text1)
	div2.AppendChild(text2)

	// Custom filter that only accepts divs with attributes
	customFilter := NodeFilterFunc(func(node Node) int {
		if elem, ok := node.(*Element); ok {
			// Use LocalName for case-insensitive matching against "div"
			if elem.LocalName() == "div" && (elem.GetAttribute("id") != "" || elem.GetAttribute("class") != "") {
				return NodeFilterAccept
			}
		}
		return NodeFilterSkip
	})

	// Test the custom filter
	result := customFilter.AcceptNode(div1) // Has id
	if result != NodeFilterAccept {
		t.Errorf("Expected custom filter to accept div with id, got %d", result)
	}

	result = customFilter.AcceptNode(div2) // Has class
	if result != NodeFilterAccept {
		t.Errorf("Expected custom filter to accept div with class, got %d", result)
	}

	result = customFilter.AcceptNode(html) // Not a div
	if result != NodeFilterSkip {
		t.Errorf("Expected custom filter to skip non-div element, got %d", result)
	}

	result = customFilter.AcceptNode(text1) // Not an element
	if result != NodeFilterSkip {
		t.Errorf("Expected custom filter to skip text node, got %d", result)
	}

	// Test combined filter
	combinedFilt := combinedFilter(NodeFilterShowElement, customFilter)

	result = combinedFilt.AcceptNode(div1)
	if result != NodeFilterAccept {
		t.Errorf("Expected combined filter to accept qualifying div, got %d", result)
	}

	result = combinedFilt.AcceptNode(text1)
	if result != NodeFilterSkip {
		t.Errorf("Expected combined filter to skip text node, got %d", result)
	}

	result = combinedFilt.AcceptNode(html)
	if result != NodeFilterSkip {
		t.Errorf("Expected combined filter to skip non-qualifying element, got %d", result)
	}
}
