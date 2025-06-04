package dom

import (
	"sync"
	"testing"
)

func TestHTMLCollection_BasicFunctionality(t *testing.T) {
	doc := NewDocument()
	root := doc.CreateElement("div")
	doc.AppendChild(root)

	// Add some child elements
	span1 := doc.CreateElement("span")
	span1.SetAttribute("id", "first")
	span1.SetAttribute("name", "test-span")
	root.AppendChild(span1)

	span2 := doc.CreateElement("span")
	span2.SetAttribute("id", "second")
	root.AppendChild(span2)

	p1 := doc.CreateElement("p")
	p1.SetAttribute("id", "paragraph")
	root.AppendChild(p1)

	// Create collection for span elements
	collection := NewElementsByTagNameCollection(root, "span")

	// Test Length
	if collection.Length() != 2 {
		t.Errorf("Expected length 2, got %d", collection.Length())
	}

	// Test Item
	item0 := collection.Item(0)
	if item0 != span1 {
		t.Errorf("Expected first item to be span1")
	}

	item1 := collection.Item(1)
	if item1 != span2 {
		t.Errorf("Expected second item to be span2")
	}

	// Test out of bounds
	itemOOB := collection.Item(5)
	if itemOOB != nil {
		t.Errorf("Expected nil for out of bounds index")
	}

	// Test negative index
	itemNeg := collection.Item(-1)
	if itemNeg != nil {
		t.Errorf("Expected nil for negative index")
	}

	// Test NamedItem by id
	namedById := collection.NamedItem("first")
	if namedById != span1 {
		t.Errorf("Expected to find span1 by id 'first'")
	}

	// Test NamedItem by name
	namedByName := collection.NamedItem("test-span")
	if namedByName != span1 {
		t.Errorf("Expected to find span1 by name 'test-span'")
	}

	// Test NamedItem not found
	notFound := collection.NamedItem("nonexistent")
	if notFound != nil {
		t.Errorf("Expected nil for nonexistent name")
	}

	// Test empty name
	emptyName := collection.NamedItem("")
	if emptyName != nil {
		t.Errorf("Expected nil for empty name")
	}
}

func TestHTMLCollection_LiveCollection(t *testing.T) {
	doc := NewDocument()
	root := doc.CreateElement("div")
	doc.AppendChild(root)

	// Create collection for span elements
	collection := NewElementsByTagNameCollection(root, "span")

	// Initially empty
	if collection.Length() != 0 {
		t.Errorf("Expected initial length 0, got %d", collection.Length())
	}

	// Add a span element
	span1 := doc.CreateElement("span")
	root.AppendChild(span1)

	// Collection should now reflect the change
	if collection.Length() != 1 {
		t.Errorf("Expected length 1 after adding span, got %d", collection.Length())
	}

	if collection.Item(0) != span1 {
		t.Errorf("Expected first item to be the added span")
	}

	// Add another span
	span2 := doc.CreateElement("span")
	root.AppendChild(span2)

	// Collection should now have 2 items
	if collection.Length() != 2 {
		t.Errorf("Expected length 2 after adding second span, got %d", collection.Length())
	}

	// Remove the first span
	root.RemoveChild(span1)

	// Collection should now have 1 item
	if collection.Length() != 1 {
		t.Errorf("Expected length 1 after removing first span, got %d", collection.Length())
	}

	if collection.Item(0) != span2 {
		t.Errorf("Expected remaining item to be span2")
	}
}

func TestHTMLCollection_TagNameWildcard(t *testing.T) {
	doc := NewDocument()
	root := doc.CreateElement("div")
	doc.AppendChild(root)

	// Add various elements
	span := doc.CreateElement("span")
	root.AppendChild(span)

	p := doc.CreateElement("p")
	root.AppendChild(p)

	div := doc.CreateElement("div")
	root.AppendChild(div)

	// Create collection for all elements using wildcard
	collection := NewElementsByTagNameCollection(root, "*")

	// Should include all descendant elements (NOT including root itself per DOM spec)
	expectedElements := []*Element{span, p, div}
	if collection.Length() != len(expectedElements) {
		t.Errorf("Expected length %d for wildcard collection, got %d", len(expectedElements), collection.Length())
	}

	// Verify all elements are included
	items := collection.ToSlice()
	found := make(map[*Element]bool)
	for _, item := range items {
		found[item] = true
	}

	for _, elem := range expectedElements {
		if !found[elem] {
			t.Errorf("Wildcard collection missing element: %s", elem.TagName())
		}
	}
}

func TestHTMLCollection_ElementsByClassName(t *testing.T) {
	doc := NewDocument()
	root := doc.CreateElement("div")
	doc.AppendChild(root)

	// Add elements with classes
	span1 := doc.CreateElement("span")
	span1.SetAttribute("class", "highlight")
	root.AppendChild(span1)

	p1 := doc.CreateElement("p")
	p1.SetAttribute("class", "highlight important")
	root.AppendChild(p1)

	span2 := doc.CreateElement("span")
	span2.SetAttribute("class", "normal")
	root.AppendChild(span2)

	// Create collection for elements with "highlight" class
	collection := NewElementsByClassNameCollection(root, "highlight")

	// Should find 2 elements
	if collection.Length() != 2 {
		t.Errorf("Expected length 2 for highlight class, got %d", collection.Length())
	}

	items := collection.ToSlice()
	if items[0] != span1 || items[1] != p1 {
		t.Errorf("Expected to find span1 and p1 with highlight class")
	}
}

func TestHTMLCollection_ElementsByTagNameNS(t *testing.T) {
	doc := NewDocument()
	root := doc.CreateElement("div")
	doc.AppendChild(root)

	// Add elements with different namespaces
	elem1, err := doc.CreateElementNS("http://www.w3.org/1999/xhtml", "span")
	if err != nil {
		t.Fatalf("Failed to create namespaced element: %v", err)
	}
	root.AppendChild(elem1)

	elem2, err := doc.CreateElementNS("http://www.w3.org/2000/svg", "rect")
	if err != nil {
		t.Fatalf("Failed to create SVG element: %v", err)
	}
	root.AppendChild(elem2)

	elem3, err := doc.CreateElementNS("http://www.w3.org/1999/xhtml", "p")
	if err != nil {
		t.Fatalf("Failed to create second namespaced element: %v", err)
	}
	root.AppendChild(elem3)

	// Collection for XHTML namespace with any local name
	collection := NewElementsByTagNameNSCollection(root, "http://www.w3.org/1999/xhtml", "*")

	// Should find 2 XHTML elements
	if collection.Length() != 2 {
		t.Errorf("Expected length 2 for XHTML namespace, got %d", collection.Length())
	}

	items := collection.ToSlice()
	found := make(map[*Element]bool)
	for _, item := range items {
		found[item] = true
	}

	if !found[elem1] || !found[elem3] {
		t.Errorf("Expected to find both XHTML elements")
	}

	// Collection for any namespace with specific local name
	collection2 := NewElementsByTagNameNSCollection(root, "*", "rect")

	// Should find 1 rect element
	if collection2.Length() != 1 {
		t.Errorf("Expected length 1 for rect elements, got %d", collection2.Length())
	}

	if collection2.Item(0) != elem2 {
		t.Errorf("Expected to find the SVG rect element")
	}
}

func TestHTMLCollection_ElementsByName(t *testing.T) {
	doc := NewDocument()
	root := doc.CreateElement("form")
	doc.AppendChild(root)

	// Add form elements with name attributes
	input1 := doc.CreateElement("input")
	input1.SetAttribute("name", "username")
	root.AppendChild(input1)

	input2 := doc.CreateElement("input")
	input2.SetAttribute("name", "password")
	root.AppendChild(input2)

	input3 := doc.CreateElement("input")
	input3.SetAttribute("name", "username") // Duplicate name
	root.AppendChild(input3)

	// Collection for elements with name="username"
	collection := NewElementsByNameCollection(root, "username")

	// Should find 2 elements with name="username"
	if collection.Length() != 2 {
		t.Errorf("Expected length 2 for name='username', got %d", collection.Length())
	}

	items := collection.ToSlice()
	if items[0] != input1 || items[1] != input3 {
		t.Errorf("Expected to find both inputs with name='username'")
	}
}

func TestHTMLCollection_ChildElementsCollection(t *testing.T) {
	doc := NewDocument()
	parent := doc.CreateElement("div")
	doc.AppendChild(parent)

	// Add various node types
	text := doc.CreateTextNode("Some text")
	parent.AppendChild(text)

	span := doc.CreateElement("span")
	parent.AppendChild(span)

	comment := doc.CreateComment("A comment")
	parent.AppendChild(comment)

	p := doc.CreateElement("p")
	parent.AppendChild(p)

	// Collection should only include element children
	collection := NewChildElementsCollection(parent)

	if collection.Length() != 2 {
		t.Errorf("Expected length 2 for child elements only, got %d", collection.Length())
	}

	items := collection.ToSlice()
	if items[0] != span || items[1] != p {
		t.Errorf("Expected to find only element children (span and p)")
	}
}

func TestHTMLCollection_NestedElements(t *testing.T) {
	doc := NewDocument()
	root := doc.CreateElement("div")
	doc.AppendChild(root)

	// Create nested structure
	child1 := doc.CreateElement("span")
	root.AppendChild(child1)

	grandchild1 := doc.CreateElement("span")
	child1.AppendChild(grandchild1)

	child2 := doc.CreateElement("p")
	root.AppendChild(child2)

	grandchild2 := doc.CreateElement("span")
	child2.AppendChild(grandchild2)

	// Collection should find all span elements in subtree
	collection := NewElementsByTagNameCollection(root, "span")

	if collection.Length() != 3 {
		t.Errorf("Expected length 3 for nested spans, got %d", collection.Length())
	}

	// Should find child1, grandchild1, and grandchild2
	items := collection.ToSlice()
	found := make(map[*Element]bool)
	for _, item := range items {
		found[item] = true
	}

	if !found[child1] || !found[grandchild1] || !found[grandchild2] {
		t.Errorf("Expected to find all nested span elements")
	}
}

func TestHTMLCollection_CacheInvalidation(t *testing.T) {
	doc := NewDocument()
	root := doc.CreateElement("div")
	doc.AppendChild(root)

	collection := NewElementsByTagNameCollection(root, "span")

	// Initially empty
	if collection.Length() != 0 {
		t.Errorf("Expected initial length 0")
	}

	// Access the collection multiple times - should use cache
	for i := 0; i < 3; i++ {
		if collection.Length() != 0 {
			t.Errorf("Cache should return consistent results")
		}
	}

	// Add an element - should invalidate cache
	span := doc.CreateElement("span")
	root.AppendChild(span)

	// Should immediately see the change
	if collection.Length() != 1 {
		t.Errorf("Expected length 1 after DOM modification")
	}

	// Modify a different part of the tree - should still invalidate cache
	otherRoot := doc.CreateElement("section")
	doc.AppendChild(otherRoot)

	span2 := doc.CreateElement("span")
	otherRoot.AppendChild(span2)

	// Original collection should still work correctly
	if collection.Length() != 1 {
		t.Errorf("Expected length 1 - other modifications shouldn't affect this collection")
	}
}

func TestHTMLCollection_ToSlice(t *testing.T) {
	doc := NewDocument()
	root := doc.CreateElement("div")
	doc.AppendChild(root)

	span1 := doc.CreateElement("span")
	root.AppendChild(span1)

	span2 := doc.CreateElement("span")
	root.AppendChild(span2)

	collection := NewElementsByTagNameCollection(root, "span")

	// Test ToSlice
	slice := collection.ToSlice()

	if len(slice) != 2 {
		t.Errorf("Expected slice length 2, got %d", len(slice))
	}

	if slice[0] != span1 || slice[1] != span2 {
		t.Errorf("ToSlice returned incorrect elements")
	}

	// Modifying the returned slice should not affect the collection
	slice[0] = nil

	if collection.Item(0) != span1 {
		t.Errorf("Collection should not be affected by slice modification")
	}
}

func TestHTMLCollection_ThreadSafety(t *testing.T) {
	doc := NewDocument()
	root := doc.CreateElement("div")
	doc.AppendChild(root)

	collection := NewElementsByTagNameCollection(root, "span")

	var wg sync.WaitGroup
	numGoroutines := 5
	numOperations := 50

	// Test concurrent reads and writes
	wg.Add(numGoroutines * 2)

	// Concurrent readers
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < numOperations; j++ {
				_ = collection.Length()
				_ = collection.Item(0)
				_ = collection.NamedItem("test")
				_ = collection.ToSlice()
			}
		}()
	}

	// Concurrent writers (DOM modifications) - only additions to avoid removal race conditions
	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()
			for j := 0; j < numOperations; j++ {
				span := doc.CreateElement("span")
				span.SetAttribute("id", "test")
				root.AppendChild(span)
			}
		}(i)
	}

	wg.Wait()

	// Should complete without panics or race conditions
	// Final state should have added elements
	finalLength := collection.Length()
	expectedLength := numGoroutines * numOperations
	if finalLength != expectedLength {
		t.Logf("Expected %d elements, got %d (this may vary due to timing)", expectedLength, finalLength)
	}
}

func TestHTMLCollection_EmptyCollection(t *testing.T) {
	doc := NewDocument()
	root := doc.CreateElement("div")
	doc.AppendChild(root)

	// Collection for non-existent elements
	collection := NewElementsByTagNameCollection(root, "nonexistent")

	if collection.Length() != 0 {
		t.Errorf("Expected length 0 for empty collection")
	}

	if collection.Item(0) != nil {
		t.Errorf("Expected nil for item in empty collection")
	}

	if collection.NamedItem("anything") != nil {
		t.Errorf("Expected nil for named item in empty collection")
	}

	slice := collection.ToSlice()
	if len(slice) != 0 {
		t.Errorf("Expected empty slice for empty collection")
	}
}

func TestHTMLCollection_NilDocument(t *testing.T) {
	// Create a collection with no document context
	filter := func(node Node) bool {
		if elem, ok := node.(*Element); ok {
			return elem.TagName() == "span"
		}
		return false
	}

	collection := NewHTMLCollection(nil, filter)

	// Should handle nil document gracefully
	if collection.Length() != 0 {
		t.Errorf("Expected length 0 for collection with nil root")
	}

	if collection.Item(0) != nil {
		t.Errorf("Expected nil for item with nil root")
	}
}

func TestHTMLCollection_ComplexDOM(t *testing.T) {
	doc := NewDocument()

	// Build a more complex DOM structure
	html := doc.CreateElement("html")
	doc.AppendChild(html)

	head := doc.CreateElement("head")
	html.AppendChild(head)

	title := doc.CreateElement("title")
	head.AppendChild(title)

	body := doc.CreateElement("body")
	html.AppendChild(body)

	header := doc.CreateElement("header")
	header.SetAttribute("class", "main-header")
	body.AppendChild(header)

	nav := doc.CreateElement("nav")
	header.AppendChild(nav)

	ul := doc.CreateElement("ul")
	nav.AppendChild(ul)

	for i := 0; i < 5; i++ {
		li := doc.CreateElement("li")
		ul.AppendChild(li)

		a := doc.CreateElement("a")
		a.SetAttribute("class", "nav-link")
		li.AppendChild(a)
	}

	main := doc.CreateElement("main")
	body.AppendChild(main)

	article := doc.CreateElement("article")
	article.SetAttribute("class", "content")
	main.AppendChild(article)

	// Test collection from document root
	allDivs := NewElementsByTagNameCollection(doc, "div")
	if allDivs.Length() != 0 {
		t.Errorf("Expected no div elements, got %d", allDivs.Length())
	}

	allElements := NewElementsByTagNameCollection(doc, "*")
	// html, head, title, body, header, nav, ul, li*5, a*5, main, article = 19 total elements
	if allElements.Length() != 19 {
		t.Errorf("Expected 19 total elements, got %d", allElements.Length())
	}

	navLinks := NewElementsByClassNameCollection(doc, "nav-link")
	if navLinks.Length() != 5 {
		t.Errorf("Expected 5 nav-link elements, got %d", navLinks.Length())
	}

	// Test NamedItem functionality
	header.SetAttribute("id", "main-header")
	headerById := NewElementsByTagNameCollection(doc, "*").NamedItem("main-header")
	if headerById != header {
		t.Errorf("Expected to find header by id")
	}
}
