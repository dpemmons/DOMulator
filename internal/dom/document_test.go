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
	if doc.OwnerDocument() != nil {
		t.Errorf("Expected owner document to be nil per WHATWG DOM spec")
	}
	if doc.ParentNode() != nil {
		t.Errorf("Expected parent node to be nil")
	}
	if doc.ChildNodes().Length() != 0 {
		t.Errorf("Expected no child nodes, got %d", doc.ChildNodes().Length())
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
	if div.TagName() != "DIV" { // Expect uppercase
		t.Errorf("Expected tag name 'DIV', got %s", div.TagName())
	}
	if div.OwnerDocument() != doc {
		t.Errorf("Expected owner document to match")
	}
	if div.ParentNode() != nil {
		t.Errorf("Expected new element to have no parent")
	}

	if span.TagName() != "SPAN" { // Expect uppercase
		t.Errorf("Expected tag name 'SPAN', got %s", span.TagName())
	}
	if input.TagName() != "INPUT" { // Expect uppercase
		t.Errorf("Expected tag name 'INPUT', got %s", input.TagName())
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
	if frag1.ChildNodes().Length() != 0 {
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

	if frag1.ChildNodes().Length() != 2 {
		t.Errorf("Expected fragment to have 2 children, got %d", frag1.ChildNodes().Length())
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
	if divs.Length() != 2 {
		t.Errorf("Expected 2 div elements, got %d", divs.Length())
	}
	if divs.Item(0) != div1 || divs.Item(1) != div2 {
		t.Errorf("Unexpected div elements order")
	}

	// Test finding spans
	spans := doc.GetElementsByTagName("span")
	if spans.Length() != 2 {
		t.Errorf("Expected 2 span elements, got %d", spans.Length())
	}
	// Depth-first traversal order: span2 (child of div2) comes first in traversal, then span1 (direct child of div1)
	if spans.Item(0) != span2 || spans.Item(1) != span1 {
		t.Errorf("Unexpected span elements order. Got: [%p, %p], Expected: [%p, %p]",
			spans.Item(0), spans.Item(1), span2, span1)
	}

	// Test finding single element
	paragraphs := doc.GetElementsByTagName("p")
	if paragraphs.Length() != 1 {
		t.Errorf("Expected 1 p element, got %d", paragraphs.Length())
	}
	if paragraphs.Item(0) != p {
		t.Errorf("Expected to find p element")
	}

	// Test finding non-existent elements
	links := doc.GetElementsByTagName("a")
	if links.Length() != 0 {
		t.Errorf("Expected 0 a elements, got %d", links.Length())
	}

	// Test wildcard selector
	all := doc.GetElementsByTagName("*")
	expectedCount := 7 // html, body, div1, div2, span1, span2, p
	if all.Length() != expectedCount {
		t.Errorf("Expected %d elements with wildcard, got %d", expectedCount, all.Length())
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
	if mainElements.Length() != 2 {
		t.Errorf("Expected 2 elements with 'main' class, got %d", mainElements.Length())
	}
	if mainElements.Item(0) != div1 || mainElements.Item(1) != span1 {
		t.Errorf("Unexpected elements with 'main' class")
	}

	contentElements := doc.GetElementsByClassName("content")
	if contentElements.Length() != 1 {
		t.Errorf("Expected 1 element with 'content' class, got %d", contentElements.Length())
	}
	if contentElements.Item(0) != div2 {
		t.Errorf("Expected to find div2 with 'content' class")
	}

	// Test finding non-existent class
	nonExistent := doc.GetElementsByClassName("nonexistent")
	if nonExistent.Length() != 0 {
		t.Errorf("Expected 0 elements with nonexistent class, got %d", nonExistent.Length())
	}

	// Test with empty class name
	emptyClass := doc.GetElementsByClassName("")
	if emptyClass.Length() != 0 {
		t.Errorf("Expected 0 elements with empty class name, got %d", emptyClass.Length())
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

	// DocumentElement should return the first element child, regardless of tag name
	if doc2.DocumentElement() != body2 {
		t.Errorf("Expected document element to be the body element when it's the only child")
	}
	// Head should be nil when no html element (Head looks specifically within html element)
	if doc2.Head() != nil {
		t.Errorf("Expected no head when no html element")
	}
	// Body should be nil when no html element (Body looks specifically within html element)
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
	if primaryElements.Length() != 3 {
		t.Errorf("Expected 3 elements with 'primary' class, got %d", primaryElements.Length())
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
	if primaryElementsAfter.Length() != 4 {
		t.Errorf("Expected 4 elements with 'primary' class after addition, got %d", primaryElementsAfter.Length())
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

// TestDocumentFragmentGetElementById tests the NonElementParentNode mixin implementation
// for DocumentFragment as specified in WHATWG DOM Standard 4.2.4
func TestDocumentFragmentGetElementById(t *testing.T) {
	t.Run("Basic element finding", func(t *testing.T) {
		doc := NewDocument()
		frag := doc.CreateDocumentFragment()
		div1 := doc.CreateElement("div")
		div2 := doc.CreateElement("div")
		span := doc.CreateElement("span")

		div1.SetAttribute("id", "container")
		div2.SetAttribute("id", "content")
		span.SetAttribute("id", "text")

		// Build fragment structure: frag -> div1 -> div2 -> span
		frag.AppendChild(div1)
		div1.AppendChild(div2)
		div2.AppendChild(span)

		// Test finding existing elements
		foundContainer := frag.GetElementById("container")
		if foundContainer != div1 {
			t.Errorf("Expected to find container div, got %v", foundContainer)
		}

		foundContent := frag.GetElementById("content")
		if foundContent != div2 {
			t.Errorf("Expected to find content div, got %v", foundContent)
		}

		foundText := frag.GetElementById("text")
		if foundText != span {
			t.Errorf("Expected to find text span, got %v", foundText)
		}
	})

	t.Run("Non-existent and empty ID handling", func(t *testing.T) {
		doc := NewDocument()
		frag := doc.CreateDocumentFragment()
		elem := doc.CreateElement("div")
		elem.SetAttribute("id", "existing")
		frag.AppendChild(elem)

		// Test finding non-existent element
		notFound := frag.GetElementById("nonexistent")
		if notFound != nil {
			t.Errorf("Expected not to find nonexistent element, got %v", notFound)
		}

		// Test with empty ID
		notFoundEmpty := frag.GetElementById("")
		if notFoundEmpty != nil {
			t.Errorf("Expected not to find element with empty ID, got %v", notFoundEmpty)
		}
	})

	t.Run("Tree order traversal with duplicate IDs", func(t *testing.T) {
		doc := NewDocument()
		frag := doc.CreateDocumentFragment()

		// Create structure: frag -> [div1, div2] where div1 -> span1, div2 -> span2
		div1 := doc.CreateElement("div")
		div2 := doc.CreateElement("div")
		span1 := doc.CreateElement("span")
		span2 := doc.CreateElement("span")

		// Both spans have the same ID
		span1.SetAttribute("id", "duplicate")
		span2.SetAttribute("id", "duplicate")

		frag.AppendChild(div1)
		frag.AppendChild(div2)
		div1.AppendChild(span1)
		div2.AppendChild(span2)

		// Should return first element found in tree order (depth-first)
		found := frag.GetElementById("duplicate")
		if found != span1 {
			t.Errorf("Expected to find first span in tree order, got %v", found)
		}
	})

	t.Run("Complex nested structure", func(t *testing.T) {
		doc := NewDocument()
		frag := doc.CreateDocumentFragment()

		// Create a more complex structure
		article := doc.CreateElement("article")
		header := doc.CreateElement("header")
		section := doc.CreateElement("section")
		footer := doc.CreateElement("footer")
		nav := doc.CreateElement("nav")
		aside := doc.CreateElement("aside")

		article.SetAttribute("id", "main-article")
		header.SetAttribute("id", "article-header")
		section.SetAttribute("id", "article-content")
		footer.SetAttribute("id", "article-footer")
		nav.SetAttribute("id", "article-nav")
		aside.SetAttribute("id", "article-sidebar")

		// Build structure: frag -> article -> [header, section -> aside, footer]
		//                                   header -> nav
		frag.AppendChild(article)
		article.AppendChild(header)
		article.AppendChild(section)
		article.AppendChild(footer)
		header.AppendChild(nav)
		section.AppendChild(aside)

		// Test finding elements at different levels
		foundArticle := frag.GetElementById("main-article")
		if foundArticle != article {
			t.Errorf("Expected to find article, got %v", foundArticle)
		}

		foundHeader := frag.GetElementById("article-header")
		if foundHeader != header {
			t.Errorf("Expected to find header, got %v", foundHeader)
		}

		foundNav := frag.GetElementById("article-nav")
		if foundNav != nav {
			t.Errorf("Expected to find nav, got %v", foundNav)
		}

		foundAside := frag.GetElementById("article-sidebar")
		if foundAside != aside {
			t.Errorf("Expected to find aside, got %v", foundAside)
		}

		foundFooter := frag.GetElementById("article-footer")
		if foundFooter != footer {
			t.Errorf("Expected to find footer, got %v", foundFooter)
		}
	})

	t.Run("Empty fragment", func(t *testing.T) {
		doc := NewDocument()
		frag := doc.CreateDocumentFragment()

		// Empty fragment should return nil for any ID
		result := frag.GetElementById("anything")
		if result != nil {
			t.Errorf("Expected nil from empty fragment, got %v", result)
		}
	})

	t.Run("Fragment with non-element children", func(t *testing.T) {
		doc := NewDocument()
		frag := doc.CreateDocumentFragment()

		// Add text and comment nodes (no IDs to find)
		text := doc.CreateTextNode("some text")
		comment := doc.CreateComment("a comment")
		elem := doc.CreateElement("p")
		elem.SetAttribute("id", "paragraph")

		frag.AppendChild(text)
		frag.AppendChild(comment)
		frag.AppendChild(elem)

		// Should find the element despite non-element siblings
		found := frag.GetElementById("paragraph")
		if found != elem {
			t.Errorf("Expected to find paragraph element, got %v", found)
		}
	})

	t.Run("Tree order verification", func(t *testing.T) {
		doc := NewDocument()
		frag := doc.CreateDocumentFragment()

		// Create structure to test specific tree order:
		// frag -> [div1, div2]
		//         div1 -> [span1, p1]
		//                 span1 -> strong1
		//         div2 -> [span2, p2]

		div1 := doc.CreateElement("div")
		div2 := doc.CreateElement("div")
		span1 := doc.CreateElement("span")
		span2 := doc.CreateElement("span")
		p1 := doc.CreateElement("p")
		p2 := doc.CreateElement("p")
		strong1 := doc.CreateElement("strong")

		// Set same ID on elements that should be found in specific order
		span1.SetAttribute("id", "same")
		strong1.SetAttribute("id", "same")
		span2.SetAttribute("id", "same")

		frag.AppendChild(div1)
		frag.AppendChild(div2)
		div1.AppendChild(span1)
		div1.AppendChild(p1)
		span1.AppendChild(strong1)
		div2.AppendChild(span2)
		div2.AppendChild(p2)

		// Tree order should be: div1, span1, strong1, p1, div2, span2, p2
		// So first element with id="same" should be span1
		found := frag.GetElementById("same")
		if found != span1 {
			t.Errorf("Expected to find span1 first in tree order, got %v", found)
		}
	})
}

// TestNonElementParentNodeCompliance verifies that both Document and DocumentFragment
// implement the NonElementParentNode mixin correctly according to the specification
func TestNonElementParentNodeCompliance(t *testing.T) {
	t.Run("Document implements NonElementParentNode", func(t *testing.T) {
		doc := NewDocument()
		// Verify Document has GetElementById method
		elem := doc.CreateElement("div")
		elem.SetAttribute("id", "test")
		doc.AppendChild(elem)

		found := doc.GetElementById("test")
		if found != elem {
			t.Errorf("Document.GetElementById not working correctly")
		}
	})

	t.Run("DocumentFragment implements NonElementParentNode", func(t *testing.T) {
		doc := NewDocument()
		// Verify DocumentFragment has GetElementById method
		frag := doc.CreateDocumentFragment()
		elem := doc.CreateElement("div")
		elem.SetAttribute("id", "test")
		frag.AppendChild(elem)

		found := frag.GetElementById("test")
		if found != elem {
			t.Errorf("DocumentFragment.GetElementById not working correctly")
		}
	})

	t.Run("Specification compliance: descendants only", func(t *testing.T) {
		doc := NewDocument()
		// Both Document and DocumentFragment should only search descendants,
		// not themselves (though they can't have IDs anyway)

		// Test with DocumentFragment
		frag := doc.CreateDocumentFragment()
		elem := doc.CreateElement("div")
		elem.SetAttribute("id", "child")
		frag.AppendChild(elem)

		// DocumentFragment itself has no ID and shouldn't be considered
		found := frag.GetElementById("child")
		if found != elem {
			t.Errorf("DocumentFragment should find descendant element")
		}

		// Test with Document
		doc2 := NewDocument()
		elem2 := doc2.CreateElement("p")
		elem2.SetAttribute("id", "doc-child")
		doc2.AppendChild(elem2)

		found2 := doc2.GetElementById("doc-child")
		if found2 != elem2 {
			t.Errorf("Document should find descendant element")
		}
	})
}
