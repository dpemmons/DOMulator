package dom

import (
	"strings"
	"testing"
)

func TestElementCreation(t *testing.T) {
	doc := NewDocument()
	elem := NewElement("div", doc)

	if elem.NodeType() != ElementNode {
		t.Errorf("Expected node type %v, got %v", ElementNode, elem.NodeType())
	}
	if elem.NodeName() != "div" {
		t.Errorf("Expected node name 'div', got %s", elem.NodeName())
	}
	if elem.TagName() != "div" {
		t.Errorf("Expected tag name 'div', got %s", elem.TagName())
	}
	if elem.OwnerDocument() != doc {
		t.Errorf("Expected owner document to match")
	}
	if elem.ParentNode() != nil {
		t.Errorf("Expected parent node to be nil")
	}
	if len(elem.ChildNodes()) != 0 {
		t.Errorf("Expected no child nodes, got %d", len(elem.ChildNodes()))
	}
}

func TestElementAttributes(t *testing.T) {
	doc := NewDocument()
	elem := NewElement("div", doc)

	// Test setting and getting attributes
	elem.SetAttribute("id", "test-id")
	elem.SetAttribute("class", "test-class")
	elem.SetAttribute("data-value", "123")

	if elem.GetAttribute("id") != "test-id" {
		t.Errorf("Expected id attribute 'test-id', got %s", elem.GetAttribute("id"))
	}
	if elem.GetAttribute("class") != "test-class" {
		t.Errorf("Expected class attribute 'test-class', got %s", elem.GetAttribute("class"))
	}
	if elem.GetAttribute("data-value") != "123" {
		t.Errorf("Expected data-value attribute '123', got %s", elem.GetAttribute("data-value"))
	}

	// Test non-existent attribute
	if elem.GetAttribute("nonexistent") != "" {
		t.Errorf("Expected empty string for non-existent attribute, got %s", elem.GetAttribute("nonexistent"))
	}

	// Test HasAttribute
	if !elem.HasAttribute("id") {
		t.Errorf("Expected element to have 'id' attribute")
	}
	if !elem.HasAttribute("class") {
		t.Errorf("Expected element to have 'class' attribute")
	}
	if elem.HasAttribute("nonexistent") {
		t.Errorf("Expected element to not have 'nonexistent' attribute")
	}

	// Test updating existing attribute
	elem.SetAttribute("id", "updated-id")
	if elem.GetAttribute("id") != "updated-id" {
		t.Errorf("Expected updated id attribute 'updated-id', got %s", elem.GetAttribute("id"))
	}

	// Test RemoveAttribute
	elem.RemoveAttribute("data-value")
	if elem.HasAttribute("data-value") {
		t.Errorf("Expected 'data-value' attribute to be removed")
	}
	if elem.GetAttribute("data-value") != "" {
		t.Errorf("Expected empty string for removed attribute, got %s", elem.GetAttribute("data-value"))
	}

	// Test removing non-existent attribute (should not crash)
	elem.RemoveAttribute("nonexistent")
}

func TestElementClassHandling(t *testing.T) {
	doc := NewDocument()
	elem := NewElement("div", doc)

	// Test HasClass with no class attribute
	if elem.HasClass("test") {
		t.Errorf("Expected element to not have class 'test' when no class attribute")
	}

	// Test HasClass with single class
	elem.SetAttribute("class", "test")
	if !elem.HasClass("test") {
		t.Errorf("Expected element to have class 'test'")
	}
	if elem.HasClass("other") {
		t.Errorf("Expected element to not have class 'other'")
	}

	// Test HasClass with multiple classes
	elem.SetAttribute("class", "test class1 class2")
	if !elem.HasClass("test") {
		t.Errorf("Expected element to have class 'test'")
	}
	if !elem.HasClass("class1") {
		t.Errorf("Expected element to have class 'class1'")
	}
	if !elem.HasClass("class2") {
		t.Errorf("Expected element to have class 'class2'")
	}
	if elem.HasClass("class3") {
		t.Errorf("Expected element to not have class 'class3'")
	}

	// Test HasClass with whitespace
	elem.SetAttribute("class", "  test   class1  class2  ")
	if !elem.HasClass("test") {
		t.Errorf("Expected element to have class 'test' with extra whitespace")
	}
	if !elem.HasClass("class1") {
		t.Errorf("Expected element to have class 'class1' with extra whitespace")
	}

	// Test HasClass with empty class attribute
	elem.SetAttribute("class", "")
	if elem.HasClass("test") {
		t.Errorf("Expected element to not have class 'test' when class attribute is empty")
	}
}

func TestElementGetElementsByTagName(t *testing.T) {
	doc := NewDocument()
	root := NewElement("div", doc)
	child1 := NewElement("span", doc)
	child2 := NewElement("p", doc)
	child3 := NewElement("span", doc)
	grandchild := NewElement("span", doc)

	root.AppendChild(child1)
	root.AppendChild(child2)
	root.AppendChild(child3)
	child2.AppendChild(grandchild)

	// Test finding specific tag
	spans := root.GetElementsByTagName("span")
	if len(spans) != 3 {
		t.Errorf("Expected 3 span elements, got %d", len(spans))
	}
	// Depth-first traversal order: child1, grandchild (under child2), child3
	if spans[0] != child1 || spans[1] != grandchild || spans[2] != child3 {
		t.Errorf("Unexpected span elements returned. Got: [%p, %p, %p], Expected: [%p, %p, %p]",
			spans[0], spans[1], spans[2], child1, grandchild, child3)
	}

	// Test finding non-existent tag
	divs := root.GetElementsByTagName("div")
	if len(divs) != 0 {
		t.Errorf("Expected 0 div elements (excluding self), got %d", len(divs))
	}

	// Test wildcard selector
	all := root.GetElementsByTagName("*")
	if len(all) != 4 {
		t.Errorf("Expected 4 elements with wildcard, got %d", len(all))
	}

	// Test on element with no children
	child2spans := child1.GetElementsByTagName("span")
	if len(child2spans) != 0 {
		t.Errorf("Expected 0 span elements in child1, got %d", len(child2spans))
	}
}

func TestElementGetElementsByClassName(t *testing.T) {
	doc := NewDocument()
	root := NewElement("div", doc)
	child1 := NewElement("span", doc)
	child2 := NewElement("p", doc)
	child3 := NewElement("span", doc)
	grandchild := NewElement("div", doc)

	child1.SetAttribute("class", "test")
	child2.SetAttribute("class", "test other")
	child3.SetAttribute("class", "other")
	grandchild.SetAttribute("class", "test nested")

	root.AppendChild(child1)
	root.AppendChild(child2)
	root.AppendChild(child3)
	child2.AppendChild(grandchild)

	// Test finding elements by class
	testElements := root.GetElementsByClassName("test")
	if len(testElements) != 3 {
		t.Errorf("Expected 3 elements with class 'test', got %d", len(testElements))
	}
	if testElements[0] != child1 || testElements[1] != child2 || testElements[2] != grandchild {
		t.Errorf("Unexpected elements with class 'test' returned")
	}

	otherElements := root.GetElementsByClassName("other")
	if len(otherElements) != 2 {
		t.Errorf("Expected 2 elements with class 'other', got %d", len(otherElements))
	}
	if otherElements[0] != child2 || otherElements[1] != child3 {
		t.Errorf("Unexpected elements with class 'other' returned")
	}

	// Test finding non-existent class
	nonExistent := root.GetElementsByClassName("nonexistent")
	if len(nonExistent) != 0 {
		t.Errorf("Expected 0 elements with non-existent class, got %d", len(nonExistent))
	}

	// Test on element with no children
	emptyResult := child1.GetElementsByClassName("test")
	if len(emptyResult) != 0 {
		t.Errorf("Expected 0 elements in child1, got %d", len(emptyResult))
	}
}

func TestElementTextContent(t *testing.T) {
	doc := NewDocument()
	elem := NewElement("div", doc)

	// Test empty element
	if elem.TextContent() != "" {
		t.Errorf("Expected empty text content, got '%s'", elem.TextContent())
	}

	// Test setting text content
	elem.SetTextContent("Hello World")
	if elem.TextContent() != "Hello World" {
		t.Errorf("Expected text content 'Hello World', got '%s'", elem.TextContent())
	}

	// Verify it creates a text node child
	if len(elem.ChildNodes()) != 1 {
		t.Errorf("Expected 1 child node after setting text content, got %d", len(elem.ChildNodes()))
	}
	if elem.FirstChild().NodeType() != TextNode {
		t.Errorf("Expected first child to be text node")
	}
	if elem.FirstChild().NodeValue() != "Hello World" {
		t.Errorf("Expected text node value 'Hello World', got '%s'", elem.FirstChild().NodeValue())
	}

	// Test overwriting text content
	elem.SetTextContent("New Text")
	if elem.TextContent() != "New Text" {
		t.Errorf("Expected updated text content 'New Text', got '%s'", elem.TextContent())
	}
	if len(elem.ChildNodes()) != 1 {
		t.Errorf("Expected 1 child node after updating text content, got %d", len(elem.ChildNodes()))
	}

	// Test with nested elements
	elem2 := NewElement("div", doc)
	span := NewElement("span", doc)
	text1 := NewText("Hello ", doc)
	text2 := NewText("World", doc)

	elem2.AppendChild(text1)
	elem2.AppendChild(span)
	span.AppendChild(text2)

	expectedText := "Hello World"
	if elem2.TextContent() != expectedText {
		t.Errorf("Expected nested text content '%s', got '%s'", expectedText, elem2.TextContent())
	}
}

func TestElementInnerHTML(t *testing.T) {
	doc := NewDocument()
	elem := NewElement("div", doc)

	// Test empty element
	if elem.InnerHTML() != "" {
		t.Errorf("Expected empty innerHTML, got '%s'", elem.InnerHTML())
	}

	// Test with text node
	text := NewText("Hello", doc)
	elem.AppendChild(text)
	if elem.InnerHTML() != "Hello" {
		t.Errorf("Expected innerHTML 'Hello', got '%s'", elem.InnerHTML())
	}

	// Test SetInnerHTML (basic implementation)
	elem.SetInnerHTML("New Content")
	if elem.InnerHTML() != "New Content" {
		t.Errorf("Expected innerHTML 'New Content', got '%s'", elem.InnerHTML())
	}

	// Verify it clears existing children and creates text node
	if len(elem.ChildNodes()) != 1 {
		t.Errorf("Expected 1 child node after setting innerHTML, got %d", len(elem.ChildNodes()))
	}
	if elem.FirstChild().NodeType() != TextNode {
		t.Errorf("Expected first child to be text node")
	}
}

func TestElementOuterHTML(t *testing.T) {
	doc := NewDocument()
	elem := NewElement("div", doc)

	// Test empty element
	expected := "<div></div>"
	if elem.OuterHTML() != expected {
		t.Errorf("Expected outerHTML '%s', got '%s'", expected, elem.OuterHTML())
	}

	// Test with content
	elem.SetTextContent("Hello")
	expected = "<div>Hello</div>"
	if elem.OuterHTML() != expected {
		t.Errorf("Expected outerHTML '%s', got '%s'", expected, elem.OuterHTML())
	}

	// Test different tag
	span := NewElement("span", doc)
	span.SetTextContent("World")
	expected = "<span>World</span>"
	if span.OuterHTML() != expected {
		t.Errorf("Expected span outerHTML '%s', got '%s'", expected, span.OuterHTML())
	}
}

func TestElementFormProperties(t *testing.T) {
	doc := NewDocument()
	input := NewElement("input", doc)

	// Test default values
	if input.value != "" {
		t.Errorf("Expected default value to be empty, got '%s'", input.value)
	}
	if input.checked != false {
		t.Errorf("Expected default checked to be false, got %v", input.checked)
	}
	if input.selected != false {
		t.Errorf("Expected default selected to be false, got %v", input.selected)
	}

	// Test direct property access (this tests the struct fields exist)
	input.value = "test-value"
	input.checked = true
	input.selected = true

	if input.value != "test-value" {
		t.Errorf("Expected value 'test-value', got '%s'", input.value)
	}
	if input.checked != true {
		t.Errorf("Expected checked to be true, got %v", input.checked)
	}
	if input.selected != true {
		t.Errorf("Expected selected to be true, got %v", input.selected)
	}
}

func TestElementCacheInvalidation(t *testing.T) {
	doc := NewDocument()
	elem := NewElement("div", doc)

	// Set initial content and access cached values
	elem.SetTextContent("Initial")
	initialText := elem.TextContent()
	initialInner := elem.InnerHTML()
	initialOuter := elem.OuterHTML()

	// Modify content through SetTextContent
	elem.SetTextContent("Modified")

	// Verify caches are properly invalidated
	if elem.TextContent() == initialText {
		t.Errorf("TextContent cache was not invalidated")
	}
	if elem.InnerHTML() == initialInner {
		t.Errorf("InnerHTML cache was not invalidated")
	}
	if elem.OuterHTML() == initialOuter {
		t.Errorf("OuterHTML cache was not invalidated")
	}

	// Verify new values are correct
	if elem.TextContent() != "Modified" {
		t.Errorf("Expected text content 'Modified', got '%s'", elem.TextContent())
	}

	// Test cache invalidation with SetInnerHTML
	elem.SetInnerHTML("HTML Modified")
	if elem.InnerHTML() != "HTML Modified" {
		t.Errorf("Expected innerHTML 'HTML Modified', got '%s'", elem.InnerHTML())
	}
}

func TestElementNodeHierarchyIntegration(t *testing.T) {
	doc := NewDocument()
	parent := NewElement("div", doc)
	child1 := NewElement("span", doc)
	child2 := NewElement("p", doc)

	// Test that Element properly integrates with Node hierarchy
	parent.AppendChild(child1)
	parent.AppendChild(child2)

	if parent.FirstChild() != child1 {
		t.Errorf("Expected first child to be child1")
	}
	if parent.LastChild() != child2 {
		t.Errorf("Expected last child to be child2")
	}
	if child1.NextSibling() != child2 {
		t.Errorf("Expected child1's next sibling to be child2")
	}
	if child2.PreviousSibling() != child1 {
		t.Errorf("Expected child2's previous sibling to be child1")
	}

	// Test removal
	parent.RemoveChild(child1)
	if parent.FirstChild() != child2 {
		t.Errorf("Expected first child to be child2 after removal")
	}
	if child1.ParentNode() != nil {
		t.Errorf("Expected child1's parent to be nil after removal")
	}
}

func TestElementSplitBySpace(t *testing.T) {
	// Test the internal splitBySpace function through class operations
	doc := NewDocument()
	elem := NewElement("div", doc)

	// Test various whitespace scenarios
	testCases := []struct {
		classAttr string
		className string
		expected  bool
	}{
		{"test", "test", true},
		{"test", "other", false},
		{"test class1 class2", "class1", true},
		{"test class1 class2", "class3", false},
		{"  test  class1  class2  ", "test", true},
		{"  test  class1  class2  ", "class1", true},
		{"\ttest\nclass1\rclass2", "test", true},
		{"\ttest\nclass1\rclass2", "class1", true},
		{"", "test", false},
		{"   ", "test", false},
	}

	for i, tc := range testCases {
		elem.SetAttribute("class", tc.classAttr)
		result := elem.HasClass(tc.className)
		if result != tc.expected {
			t.Errorf("Test case %d: class='%s', searching for '%s', expected %v, got %v",
				i, tc.classAttr, tc.className, tc.expected, result)
		}
	}
}

func TestElementInsertAdjacentHTML(t *testing.T) {
	doc := NewDocument()

	// Create test structure: container > [before, target, after]
	container := NewElement("div", doc)
	before := NewElement("span", doc)
	target := NewElement("p", doc)
	after := NewElement("div", doc)

	before.SetTextContent("before")
	target.SetTextContent("target")
	after.SetTextContent("after")

	container.AppendChild(before)
	container.AppendChild(target)
	container.AppendChild(after)

	// Test beforebegin - insert before target element
	err := target.InsertAdjacentHTML("beforebegin", "Hello")
	if err != nil {
		t.Errorf("Unexpected error for beforebegin: %v", err)
	}

	children := container.ChildNodes()
	if len(children) != 4 {
		t.Errorf("Expected 4 children after beforebegin, got %d", len(children))
	}

	// Check order: before, inserted text, target, after
	if children[0] != before {
		t.Errorf("Expected first child to still be 'before' element")
	}
	if children[1].NodeType() != TextNode || children[1].NodeValue() != "Hello" {
		t.Errorf("Expected second child to be text node 'Hello', got type %d value '%s'",
			children[1].NodeType(), children[1].NodeValue())
	}
	if children[2] != target {
		t.Errorf("Expected third child to be target element")
	}
	if children[3] != after {
		t.Errorf("Expected fourth child to be 'after' element")
	}

	// Test afterbegin - insert as first child of target
	err = target.InsertAdjacentHTML("afterbegin", "<em>emphasis</em>")
	if err != nil {
		t.Errorf("Unexpected error for afterbegin: %v", err)
	}

	targetChildren := target.ChildNodes()
	if len(targetChildren) != 2 {
		t.Errorf("Expected 2 children in target after afterbegin, got %d", len(targetChildren))
	}

	// First child should be the inserted element
	if targetChildren[0].NodeType() != ElementNode {
		t.Errorf("Expected first child to be element node")
	}
	emElement := targetChildren[0].(*Element)
	if emElement.TagName() != "em" {
		t.Errorf("Expected first child to be 'em' element, got '%s'", emElement.TagName())
	}

	// Test beforeend - insert as last child of target
	err = target.InsertAdjacentHTML("beforeend", "<strong>strong</strong>")
	if err != nil {
		t.Errorf("Unexpected error for beforeend: %v", err)
	}

	targetChildren = target.ChildNodes()
	if len(targetChildren) != 3 {
		t.Errorf("Expected 3 children in target after beforeend, got %d", len(targetChildren))
	}

	// Last child should be the inserted element
	lastChild := targetChildren[len(targetChildren)-1]
	if lastChild.NodeType() != ElementNode {
		t.Errorf("Expected last child to be element node")
	}
	strongElement := lastChild.(*Element)
	if strongElement.TagName() != "strong" {
		t.Errorf("Expected last child to be 'strong' element, got '%s'", strongElement.TagName())
	}

	// Test afterend - insert after target element
	err = target.InsertAdjacentHTML("afterend", "World")
	if err != nil {
		t.Errorf("Unexpected error for afterend: %v", err)
	}

	children = container.ChildNodes()
	if len(children) != 5 {
		t.Errorf("Expected 5 children after afterend, got %d", len(children))
	}

	// Check that "World" was inserted after target
	if children[3].NodeType() != TextNode || children[3].NodeValue() != "World" {
		t.Errorf("Expected fourth child to be text node 'World', got type %d value '%s'",
			children[3].NodeType(), children[3].NodeValue())
	}
}

func TestElementInsertAdjacentHTMLPositions(t *testing.T) {
	doc := NewDocument()
	container := NewElement("div", doc)
	target := NewElement("p", doc)
	container.AppendChild(target)

	// Test all position values with case insensitivity
	positions := []string{"beforebegin", "afterbegin", "beforeend", "afterend"}
	casedPositions := []string{"BeforeBegin", "AFTERBEGIN", "beforeEnd", "AfterEnd"}

	for i, pos := range positions {
		// Test normal case
		err := target.InsertAdjacentHTML(pos, "test")
		if err != nil {
			t.Errorf("Unexpected error for position '%s': %v", pos, err)
		}

		// Test case insensitive
		err = target.InsertAdjacentHTML(casedPositions[i], "test")
		if err != nil {
			t.Errorf("Unexpected error for case-insensitive position '%s': %v", casedPositions[i], err)
		}
	}
}

func TestElementInsertAdjacentHTMLErrors(t *testing.T) {
	doc := NewDocument()
	elem := NewElement("div", doc)

	// Test invalid position
	err := elem.InsertAdjacentHTML("invalid", "content")
	if err == nil {
		t.Errorf("Expected error for invalid position")
	}
	if err.Error() != "invalid position: must be 'beforebegin', 'afterbegin', 'beforeend', or 'afterend'" {
		t.Errorf("Unexpected error message: %v", err)
	}

	// Test beforebegin with no parent
	err = elem.InsertAdjacentHTML("beforebegin", "content")
	if err == nil {
		t.Errorf("Expected error for beforebegin with no parent")
	}
	if err.Error() != "cannot insert beforebegin: element has no parent" {
		t.Errorf("Unexpected error message: %v", err)
	}

	// Test afterend with no parent
	err = elem.InsertAdjacentHTML("afterend", "content")
	if err == nil {
		t.Errorf("Expected error for afterend with no parent")
	}
	if err.Error() != "cannot insert afterend: element has no parent" {
		t.Errorf("Unexpected error message: %v", err)
	}
}

func TestElementInsertAdjacentHTMLBasicParsing(t *testing.T) {
	doc := NewDocument()
	container := NewElement("div", doc)
	target := NewElement("p", doc)
	container.AppendChild(target)

	// Test text content
	err := target.InsertAdjacentHTML("beforeend", "Plain text")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	children := target.ChildNodes()
	if len(children) != 1 || children[0].NodeType() != TextNode {
		t.Errorf("Expected one text node child")
	}
	if children[0].NodeValue() != "Plain text" {
		t.Errorf("Expected text content 'Plain text', got '%s'", children[0].NodeValue())
	}

	// Test simple element
	err = target.InsertAdjacentHTML("beforeend", "<span>element</span>")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	children = target.ChildNodes()
	if len(children) != 2 {
		t.Errorf("Expected 2 children after adding element, got %d", len(children))
	}

	if children[1].NodeType() != ElementNode {
		t.Errorf("Expected second child to be element")
	}

	spanElement := children[1].(*Element)
	if spanElement.TagName() != "span" {
		t.Errorf("Expected span element, got '%s'", spanElement.TagName())
	}

	// Test element with attributes
	err = target.InsertAdjacentHTML("beforeend", `<div class="test" id="myid">content</div>`)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	children = target.ChildNodes()
	if len(children) != 3 {
		t.Errorf("Expected 3 children after adding element with attributes, got %d", len(children))
	}

	divElement := children[2].(*Element)
	if divElement.TagName() != "div" {
		t.Errorf("Expected div element, got '%s'", divElement.TagName())
	}
	if divElement.GetAttribute("class") != "test" {
		t.Errorf("Expected class attribute 'test', got '%s'", divElement.GetAttribute("class"))
	}
	if divElement.GetAttribute("id") != "myid" {
		t.Errorf("Expected id attribute 'myid', got '%s'", divElement.GetAttribute("id"))
	}

	// Test self-closing element
	err = target.InsertAdjacentHTML("beforeend", `<input type="text" />`)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	children = target.ChildNodes()
	if len(children) != 4 {
		t.Errorf("Expected 4 children after adding self-closing element, got %d", len(children))
	}

	inputElement := children[3].(*Element)
	if inputElement.TagName() != "input" {
		t.Errorf("Expected input element, got '%s'", inputElement.TagName())
	}
	if inputElement.GetAttribute("type") != "text" {
		t.Errorf("Expected type attribute 'text', got '%s'", inputElement.GetAttribute("type"))
	}
}

func TestElementInsertAdjacentHTMLEmpty(t *testing.T) {
	doc := NewDocument()
	container := NewElement("div", doc)
	target := NewElement("p", doc)
	container.AppendChild(target)

	// Test empty string
	err := target.InsertAdjacentHTML("beforeend", "")
	if err != nil {
		t.Errorf("Unexpected error for empty string: %v", err)
	}

	// Should not add any children
	if len(target.ChildNodes()) != 0 {
		t.Errorf("Expected no children for empty string, got %d", len(target.ChildNodes()))
	}

	// Test whitespace only
	err = target.InsertAdjacentHTML("beforeend", "   ")
	if err != nil {
		t.Errorf("Unexpected error for whitespace: %v", err)
	}

	// Should not add any children (whitespace gets trimmed)
	if len(target.ChildNodes()) != 0 {
		t.Errorf("Expected no children for whitespace only, got %d", len(target.ChildNodes()))
	}
}

func TestElementInsertAdjacentHTMLCacheInvalidation(t *testing.T) {
	doc := NewDocument()
	container := NewElement("div", doc)
	target := NewElement("p", doc)
	target.SetTextContent("original")
	container.AppendChild(target)

	// Access cached values
	originalText := target.TextContent()
	originalInner := target.InnerHTML()
	originalOuter := target.OuterHTML()

	// Insert content
	err := target.InsertAdjacentHTML("beforeend", "<span>added</span>")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Verify caches are invalidated
	if target.TextContent() == originalText {
		t.Errorf("TextContent cache was not invalidated")
	}
	if target.InnerHTML() == originalInner {
		t.Errorf("InnerHTML cache was not invalidated")
	}
	if target.OuterHTML() == originalOuter {
		t.Errorf("OuterHTML cache was not invalidated")
	}

	// Verify new content is reflected
	newText := target.TextContent()
	if !strings.Contains(newText, "original") {
		t.Errorf("Expected new text content to contain 'original'")
	}
}
