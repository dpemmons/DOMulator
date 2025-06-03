package dom

import (
	"testing"
)

func TestDocumentCreation(t *testing.T) {
	doc := NewDocument()

	if doc.NodeType() != DocumentNode {
		t.Errorf("Expected node type %v, got %v", DocumentNode, doc.NodeType())
	}
	if doc.NodeName() != "#document" {
		t.Errorf("Expected node name '#document', got %s", doc.NodeName())
	}
	if doc.OwnerDocument() != doc {
		t.Errorf("Expected owner document to be itself")
	}
	if doc.ParentNode() != nil {
		t.Errorf("Expected parent node to be nil")
	}
	if len(doc.ChildNodes()) != 0 {
		t.Errorf("Expected no child nodes, got %d", len(doc.ChildNodes()))
	}
}

func TestDocumentCreateElement(t *testing.T) {
	doc := NewDocument()

	// Test creating various elements
	div := doc.CreateElement("div")
	span := doc.CreateElement("span")
	input := doc.CreateElement("input")

	if div.NodeType() != ElementNode {
		t.Errorf("Expected element node type, got %v", div.NodeType())
	}
	if div.TagName() != "div" {
		t.Errorf("Expected tag name 'div', got %s", div.TagName())
	}
	if div.OwnerDocument() != doc {
		t.Errorf("Expected owner document to match")
	}
	if div.ParentNode() != nil {
		t.Errorf("Expected new element to have no parent")
	}

	if span.TagName() != "span" {
		t.Errorf("Expected tag name 'span', got %s", span.TagName())
	}
	if input.TagName() != "input" {
		t.Errorf("Expected tag name 'input', got %s", input.TagName())
	}

	// Test that elements are independent
	if div == span {
		t.Errorf("Expected different element instances")
	}
}

func TestDocumentCreateTextNode(t *testing.T) {
	doc := NewDocument()

	text1 := doc.CreateTextNode("Hello World")
	text2 := doc.CreateTextNode("")
	text3 := doc.CreateTextNode("Special chars: <>&\"'")

	if text1.NodeType() != TextNode {
		t.Errorf("Expected text node type, got %v", text1.NodeType())
	}
	if text1.NodeValue() != "Hello World" {
		t.Errorf("Expected text 'Hello World', got %s", text1.NodeValue())
	}
	if text1.OwnerDocument() != doc {
		t.Errorf("Expected owner document to match")
	}
	if text1.ParentNode() != nil {
		t.Errorf("Expected new text node to have no parent")
	}

	if text2.NodeValue() != "" {
		t.Errorf("Expected empty text, got %s", text2.NodeValue())
	}
	if text3.NodeValue() != "Special chars: <>&\"'" {
		t.Errorf("Expected special chars text, got %s", text3.NodeValue())
	}

	// Test that text nodes are independent
	if text1 == text2 {
		t.Errorf("Expected different text node instances")
	}
}

func TestDocumentCreateComment(t *testing.T) {
	doc := NewDocument()

	comment1 := doc.CreateComment("This is a comment")
	comment2 := doc.CreateComment("")
	comment3 := doc.CreateComment("Multi-line\ncomment\nwith breaks")

	if comment1.NodeType() != CommentNode {
		t.Errorf("Expected comment node type, got %v", comment1.NodeType())
	}
	if comment1.NodeValue() != "This is a comment" {
		t.Errorf("Expected comment 'This is a comment', got %s", comment1.NodeValue())
	}
	if comment1.OwnerDocument() != doc {
		t.Errorf("Expected owner document to match")
	}
	if comment1.ParentNode() != nil {
		t.Errorf("Expected new comment node to have no parent")
	}

	if comment2.NodeValue() != "" {
		t.Errorf("Expected empty comment, got %s", comment2.NodeValue())
	}
	if comment3.NodeValue() != "Multi-line\ncomment\nwith breaks" {
		t.Errorf("Expected multi-line comment, got %s", comment3.NodeValue())
	}

	// Test that comment nodes are independent
	if comment1 == comment2 {
		t.Errorf("Expected different comment node instances")
	}
}

func TestDocumentCreateDocumentFragment(t *testing.T) {
	doc := NewDocument()

	frag1 := doc.CreateDocumentFragment()
	frag2 := doc.CreateDocumentFragment()

	if frag1.NodeType() != DocumentFragmentNode {
		t.Errorf("Expected document fragment node type, got %v", frag1.NodeType())
	}
	if frag1.OwnerDocument() != doc {
		t.Errorf("Expected owner document to match")
	}
	if frag1.ParentNode() != nil {
		t.Errorf("Expected new fragment to have no parent")
	}
	if len(frag1.ChildNodes()) != 0 {
		t.Errorf("Expected new fragment to have no children")
	}

	// Test that fragments are independent
	if frag1 == frag2 {
		t.Errorf("Expected different fragment instances")
	}

	// Test adding children to fragment
	elem := doc.CreateElement("div")
	text := doc.CreateTextNode("content")
	frag1.AppendChild(elem)
	frag1.AppendChild(text)

	if len(frag1.ChildNodes()) != 2 {
		t.Errorf("Expected fragment to have 2 children, got %d", len(frag1.ChildNodes()))
	}
	if frag1.FirstChild() != elem {
		t.Errorf("Expected first child to be element")
	}
	if frag1.LastChild() != text {
		t.Errorf("Expected last child to be text")
	}
}

func TestDocumentGetElementById(t *testing.T) {
	doc := NewDocument()

	// Create a simple DOM structure
	html := doc.CreateElement("html")
	body := doc.CreateElement("body")
	div1 := doc.CreateElement("div")
	div2 := doc.CreateElement("div")
	span := doc.CreateElement("span")

	div1.SetAttribute("id", "container")
	div2.SetAttribute("id", "content")
	span.SetAttribute("id", "text")

	doc.AppendChild(html)
	html.AppendChild(body)
	body.AppendChild(div1)
	div1.AppendChild(div2)
	div2.AppendChild(span)

	// Test finding existing elements
	foundContainer := doc.GetElementById("container")
	if foundContainer != div1 {
		t.Errorf("Expected to find container div, got %v", foundContainer)
	}

	foundContent := doc.GetElementById("content")
	if foundContent != div2 {
		t.Errorf("Expected to find content div, got %v", foundContent)
	}

	foundText := doc.GetElementById("text")
	if foundText != span {
		t.Errorf("Expected to find text span, got %v", foundText)
	}

	// Test finding non-existent element
	notFound := doc.GetElementById("nonexistent")
	if notFound != nil {
		t.Errorf("Expected not to find nonexistent element, got %v", notFound)
	}

	// Test with empty ID
	notFoundEmpty := doc.GetElementById("")
	if notFoundEmpty != nil {
		t.Errorf("Expected not to find element with empty ID, got %v", notFoundEmpty)
	}

	// Test duplicate IDs (should return first found in traversal order)
	div3 := doc.CreateElement("div")
	div3.SetAttribute("id", "container") // Duplicate ID
	body.AppendChild(div3)

	foundDuplicate := doc.GetElementById("container")
	// The traversal finds div3 first (the one at the end of body's children)
	if foundDuplicate != div3 { // Should find the first one encountered in traversal
		t.Errorf("Expected to find div3 with duplicate ID, got div1")
	}
}

func TestDocumentGetElementsByTagName(t *testing.T) {
	doc := NewDocument()

	// Create a simple DOM structure
	html := doc.CreateElement("html")
	body := doc.CreateElement("body")
	div1 := doc.CreateElement("div")
	div2 := doc.CreateElement("div")
	span1 := doc.CreateElement("span")
	span2 := doc.CreateElement("span")
	p := doc.CreateElement("p")

	doc.AppendChild(html)
	html.AppendChild(body)
	body.AppendChild(div1)
	body.AppendChild(p)
	div1.AppendChild(div2)
	div1.AppendChild(span1)
	div2.AppendChild(span2)

	// Test finding divs
	divs := doc.GetElementsByTagName("div")
	if len(divs) != 2 {
		t.Errorf("Expected 2 div elements, got %d", len(divs))
	}
	if divs[0] != div1 || divs[1] != div2 {
		t.Errorf("Unexpected div elements order")
	}

	// Test finding spans
	spans := doc.GetElementsByTagName("span")
	if len(spans) != 2 {
		t.Errorf("Expected 2 span elements, got %d", len(spans))
	}
	// Depth-first traversal order: span2 (child of div2) comes first in traversal, then span1 (direct child of div1)
	if spans[0] != span2 || spans[1] != span1 {
		t.Errorf("Unexpected span elements order. Got: [%p, %p], Expected: [%p, %p]",
			spans[0], spans[1], span2, span1)
	}

	// Test finding single element
	paragraphs := doc.GetElementsByTagName("p")
	if len(paragraphs) != 1 {
		t.Errorf("Expected 1 p element, got %d", len(paragraphs))
	}
	if paragraphs[0] != p {
		t.Errorf("Expected to find p element")
	}

	// Test finding non-existent elements
	links := doc.GetElementsByTagName("a")
	if len(links) != 0 {
		t.Errorf("Expected 0 a elements, got %d", len(links))
	}

	// Test wildcard selector
	all := doc.GetElementsByTagName("*")
	expectedCount := 7 // html, body, div1, div2, span1, span2, p
	if len(all) != expectedCount {
		t.Errorf("Expected %d elements with wildcard, got %d", expectedCount, len(all))
	}
}

func TestDocumentGetElementsByClassName(t *testing.T) {
	doc := NewDocument()

	// Create elements with various classes
	html := doc.CreateElement("html")
	body := doc.CreateElement("body")
	div1 := doc.CreateElement("div")
	div2 := doc.CreateElement("div")
	span1 := doc.CreateElement("span")
	span2 := doc.CreateElement("span")

	div1.SetAttribute("class", "container main")
	div2.SetAttribute("class", "content")
	span1.SetAttribute("class", "text main")
	span2.SetAttribute("class", "highlight")

	doc.AppendChild(html)
	html.AppendChild(body)
	body.AppendChild(div1)
	body.AppendChild(div2)
	div1.AppendChild(span1)
	div2.AppendChild(span2)

	// Test finding by single class
	mainElements := doc.GetElementsByClassName("main")
	if len(mainElements) != 2 {
		t.Errorf("Expected 2 elements with 'main' class, got %d", len(mainElements))
	}
	if mainElements[0] != div1 || mainElements[1] != span1 {
		t.Errorf("Unexpected elements with 'main' class")
	}

	contentElements := doc.GetElementsByClassName("content")
	if len(contentElements) != 1 {
		t.Errorf("Expected 1 element with 'content' class, got %d", len(contentElements))
	}
	if contentElements[0] != div2 {
		t.Errorf("Expected to find div2 with 'content' class")
	}

	// Test finding non-existent class
	nonExistent := doc.GetElementsByClassName("nonexistent")
	if len(nonExistent) != 0 {
		t.Errorf("Expected 0 elements with nonexistent class, got %d", len(nonExistent))
	}

	// Test with empty class name
	emptyClass := doc.GetElementsByClassName("")
	if len(emptyClass) != 0 {
		t.Errorf("Expected 0 elements with empty class name, got %d", len(emptyClass))
	}
}

func TestDocumentStructureHelpers(t *testing.T) {
	doc := NewDocument()

	// Test with empty document
	if doc.DocumentElement() != nil {
		t.Errorf("Expected no document element in empty document")
	}
	if doc.Head() != nil {
		t.Errorf("Expected no head in empty document")
	}
	if doc.Body() != nil {
		t.Errorf("Expected no body in empty document")
	}

	// Create proper HTML structure
	html := doc.CreateElement("html")
	head := doc.CreateElement("head")
	body := doc.CreateElement("body")
	title := doc.CreateElement("title")
	div := doc.CreateElement("div")

	doc.AppendChild(html)
	html.AppendChild(head)
	html.AppendChild(body)
	head.AppendChild(title)
	body.AppendChild(div)

	// Test DocumentElement
	docElem := doc.DocumentElement()
	if docElem != html {
		t.Errorf("Expected document element to be html, got %v", docElem)
	}

	// Test Head
	headElem := doc.Head()
	if headElem != head {
		t.Errorf("Expected head element, got %v", headElem)
	}

	// Test Body
	bodyElem := doc.Body()
	if bodyElem != body {
		t.Errorf("Expected body element, got %v", bodyElem)
	}

	// Test with non-standard structure (no html element)
	doc2 := NewDocument()
	body2 := doc2.CreateElement("body")
	doc2.AppendChild(body2)

	if doc2.DocumentElement() != nil {
		t.Errorf("Expected no document element when no html element")
	}
	if doc2.Head() != nil {
		t.Errorf("Expected no head when no html element")
	}
	if doc2.Body() != nil {
		t.Errorf("Expected no body when no html element")
	}
}

func TestDocumentNodeHierarchyIntegration(t *testing.T) {
	doc := NewDocument()

	// Test that Document properly integrates with Node hierarchy
	html := doc.CreateElement("html")
	comment := doc.CreateComment("HTML5 document")

	doc.AppendChild(comment)
	doc.AppendChild(html)

	if doc.FirstChild() != comment {
		t.Errorf("Expected first child to be comment")
	}
	if doc.LastChild() != html {
		t.Errorf("Expected last child to be html")
	}
	if comment.NextSibling() != html {
		t.Errorf("Expected comment's next sibling to be html")
	}
	if html.PreviousSibling() != comment {
		t.Errorf("Expected html's previous sibling to be comment")
	}

	// Test removal
	doc.RemoveChild(comment)
	if doc.FirstChild() != html {
		t.Errorf("Expected first child to be html after removal")
	}
	if comment.ParentNode() != nil {
		t.Errorf("Expected comment's parent to be nil after removal")
	}
}

func TestDocumentComplexTraversal(t *testing.T) {
	doc := NewDocument()

	// Create a complex nested structure
	html := doc.CreateElement("html")
	head := doc.CreateElement("head")
	body := doc.CreateElement("body")
	header := doc.CreateElement("header")
	nav := doc.CreateElement("nav")
	main := doc.CreateElement("main")
	footer := doc.CreateElement("footer")

	// Add some IDs and classes
	header.SetAttribute("id", "main-header")
	header.SetAttribute("class", "header primary")
	nav.SetAttribute("class", "navigation primary")
	main.SetAttribute("id", "content")
	main.SetAttribute("class", "main-content")
	footer.SetAttribute("class", "footer primary")

	// Build the tree
	doc.AppendChild(html)
	html.AppendChild(head)
	html.AppendChild(body)
	body.AppendChild(header)
	body.AppendChild(nav)
	body.AppendChild(main)
	body.AppendChild(footer)

	// Test complex queries
	primaryElements := doc.GetElementsByClassName("primary")
	if len(primaryElements) != 3 {
		t.Errorf("Expected 3 elements with 'primary' class, got %d", len(primaryElements))
	}

	headerById := doc.GetElementById("main-header")
	if headerById != header {
		t.Errorf("Expected to find header by ID")
	}

	contentById := doc.GetElementById("content")
	if contentById != main {
		t.Errorf("Expected to find main by ID")
	}

	// Test that modifications work correctly
	section := doc.CreateElement("section")
	section.SetAttribute("class", "primary")
	main.AppendChild(section)

	primaryElementsAfter := doc.GetElementsByClassName("primary")
	if len(primaryElementsAfter) != 4 {
		t.Errorf("Expected 4 elements with 'primary' class after addition, got %d", len(primaryElementsAfter))
	}
}

func TestDocumentOwnershipConsistency(t *testing.T) {
	doc := NewDocument()

	// Create elements and verify ownership
	elem := doc.CreateElement("div")
	text := doc.CreateTextNode("content")
	comment := doc.CreateComment("note")
	frag := doc.CreateDocumentFragment()

	// All created nodes should have correct owner document
	if elem.OwnerDocument() != doc {
		t.Errorf("Element should have correct owner document")
	}
	if text.OwnerDocument() != doc {
		t.Errorf("Text node should have correct owner document")
	}
	if comment.OwnerDocument() != doc {
		t.Errorf("Comment node should have correct owner document")
	}
	if frag.OwnerDocument() != doc {
		t.Errorf("Fragment should have correct owner document")
	}

	// When we add to the document, ownership should be maintained
	elem.AppendChild(text)
	elem.AppendChild(comment)
	doc.AppendChild(elem)

	if text.OwnerDocument() != doc {
		t.Errorf("Text node should maintain owner document when added to tree")
	}
	if comment.OwnerDocument() != doc {
		t.Errorf("Comment node should maintain owner document when added to tree")
	}
}
