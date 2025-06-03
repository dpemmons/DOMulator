package js

import (
	"strings"
	"testing"

	"github.com/dop251/goja"
	"github.com/dpemmons/DOMulator/internal/dom"
	"github.com/dpemmons/DOMulator/internal/parser"
)

func TestDOMBindingsBasics(t *testing.T) {
	doc := dom.NewDocument()
	vm := goja.New()

	bindings := NewDOMBindings(vm, doc)

	if bindings.vm != vm {
		t.Error("DOMBindings should store the provided VM")
	}

	if bindings.document != doc {
		t.Error("DOMBindings should store the provided document")
	}

	if bindings.nodeCache == nil {
		t.Error("DOMBindings should initialize nodeCache")
	}
}

func TestDocumentWrapping(t *testing.T) {
	doc := dom.NewDocument()
	vm := goja.New()
	bindings := NewDOMBindings(vm, doc)

	docJS := bindings.WrapDocument()

	// Test basic document properties
	nodeType := docJS.Get("nodeType").ToInteger()
	if nodeType != int64(dom.DocumentNode) {
		t.Errorf("Expected nodeType %d, got %d", dom.DocumentNode, nodeType)
	}

	nodeName := docJS.Get("nodeName").String()
	if nodeName != "#document" {
		t.Errorf("Expected nodeName '#document', got '%s'", nodeName)
	}

	// Test document methods exist
	if docJS.Get("createElement") == nil {
		t.Error("document.createElement should be defined")
	}

	if docJS.Get("createTextNode") == nil {
		t.Error("document.createTextNode should be defined")
	}

	if docJS.Get("getElementById") == nil {
		t.Error("document.getElementById should be defined")
	}
}

func TestElementWrapping(t *testing.T) {
	doc := dom.NewDocument()
	vm := goja.New()
	bindings := NewDOMBindings(vm, doc)

	element := doc.CreateElement("div")
	element.SetAttribute("id", "test")
	element.SetAttribute("class", "sample")

	elemJS := bindings.WrapElement(element)

	// Test basic element properties
	nodeType := elemJS.Get("nodeType").ToInteger()
	if nodeType != int64(dom.ElementNode) {
		t.Errorf("Expected nodeType %d, got %d", dom.ElementNode, nodeType)
	}

	tagName := elemJS.Get("tagName").String()
	if tagName != "div" {
		t.Errorf("Expected tagName 'div', got '%s'", tagName)
	}

	nodeName := elemJS.Get("nodeName").String()
	if nodeName != "div" {
		t.Errorf("Expected nodeName 'div', got '%s'", nodeName)
	}

	// Test innerHTML/outerHTML properties exist
	innerHTML := elemJS.Get("innerHTML")
	if innerHTML == nil {
		t.Error("element.innerHTML should be defined")
	}

	outerHTML := elemJS.Get("outerHTML")
	if outerHTML == nil {
		t.Error("element.outerHTML should be defined")
	}

	textContent := elemJS.Get("textContent")
	if textContent == nil {
		t.Error("element.textContent should be defined")
	}
}

func TestObjectIdentityCache(t *testing.T) {
	doc := dom.NewDocument()
	vm := goja.New()
	bindings := NewDOMBindings(vm, doc)

	element := doc.CreateElement("div")

	// Wrap the same element twice
	elemJS1 := bindings.WrapElement(element)
	elemJS2 := bindings.WrapElement(element)

	// Should return the exact same JavaScript object
	if elemJS1 != elemJS2 {
		t.Error("Same DOM element should return identical JavaScript object (object identity)")
	}

	// Verify cache is working
	if len(bindings.nodeCache) != 1 {
		t.Errorf("Expected 1 cached object, got %d", len(bindings.nodeCache))
	}
}

func TestNodeExtractionFromJS(t *testing.T) {
	doc := dom.NewDocument()
	vm := goja.New()
	bindings := NewDOMBindings(vm, doc)

	element := doc.CreateElement("div")
	elemJS := bindings.WrapElement(element)

	// Test extraction
	extracted := bindings.extractNodeFromJS(elemJS)
	if extracted != element {
		t.Error("extractNodeFromJS should return the original DOM node")
	}

	// Test with invalid value
	invalidValue := vm.ToValue("not a node")
	extractedInvalid := bindings.extractNodeFromJS(invalidValue)
	if extractedInvalid != nil {
		t.Error("extractNodeFromJS should return nil for invalid values")
	}
}

func TestAttributeMethodsInJS(t *testing.T) {
	doc := dom.NewDocument()
	vm := goja.New()
	bindings := NewDOMBindings(vm, doc)

	element := doc.CreateElement("div")
	elemJS := bindings.WrapElement(element)

	// Set up JavaScript context
	vm.Set("element", elemJS)

	// Test setAttribute
	_, err := vm.RunString(`element.setAttribute('id', 'test-id')`)
	if err != nil {
		t.Fatalf("setAttribute failed: %v", err)
	}

	// Test getAttribute
	result, err := vm.RunString(`element.getAttribute('id')`)
	if err != nil {
		t.Fatalf("getAttribute failed: %v", err)
	}

	if result.String() != "test-id" {
		t.Errorf("Expected 'test-id', got '%s'", result.String())
	}

	// Test hasAttribute
	result, err = vm.RunString(`element.hasAttribute('id')`)
	if err != nil {
		t.Fatalf("hasAttribute failed: %v", err)
	}

	if !result.ToBoolean() {
		t.Error("hasAttribute should return true for existing attribute")
	}

	// Test removeAttribute
	_, err = vm.RunString(`element.removeAttribute('id')`)
	if err != nil {
		t.Fatalf("removeAttribute failed: %v", err)
	}

	result, err = vm.RunString(`element.hasAttribute('id')`)
	if err != nil {
		t.Fatalf("hasAttribute after remove failed: %v", err)
	}

	if result.ToBoolean() {
		t.Error("hasAttribute should return false after removeAttribute")
	}
}

func TestNodeMethodsInJS(t *testing.T) {
	doc := dom.NewDocument()
	vm := goja.New()
	bindings := NewDOMBindings(vm, doc)

	parent := doc.CreateElement("div")
	child := doc.CreateTextNode("Hello")

	parentJS := bindings.WrapElement(parent)
	childJS := bindings.WrapNode(child)

	vm.Set("parent", parentJS)
	vm.Set("child", childJS)

	// Test appendChild
	_, err := vm.RunString(`parent.appendChild(child)`)
	if err != nil {
		t.Fatalf("appendChild failed: %v", err)
	}

	// Verify DOM was modified
	if parent.FirstChild() != child {
		t.Error("appendChild should have added child to parent in DOM")
	}

	// Test childNodes property is updated
	result, err := vm.RunString(`parent.childNodes.length`)
	if err != nil {
		t.Fatalf("childNodes.length access failed: %v", err)
	}

	if result.ToInteger() != 1 {
		t.Errorf("Expected childNodes.length 1, got %d", result.ToInteger())
	}
}

func TestDOMManipulationWithComplexStructure(t *testing.T) {
	doc := dom.NewDocument()
	vm := goja.New()
	bindings := NewDOMBindings(vm, doc)

	docJS := bindings.WrapDocument()
	vm.Set("document", docJS)

	// Create complex DOM structure in JavaScript
	_, err := vm.RunString(`
		var div = document.createElement('div');
		div.setAttribute('id', 'container');
		
		var span = document.createElement('span');
		span.setAttribute('class', 'text');
		
		var text = document.createTextNode('Hello World');
		
		span.appendChild(text);
		div.appendChild(span);
	`)
	if err != nil {
		t.Fatalf("Complex DOM creation failed: %v", err)
	}

	// Verify structure was created correctly
	result, err := vm.RunString(`div.childNodes.length`)
	if err != nil {
		t.Fatalf("Failed to access childNodes: %v", err)
	}

	if result.ToInteger() != 1 {
		t.Errorf("Expected 1 child, got %d", result.ToInteger())
	}

	// Test text content
	result, err = vm.RunString(`div.textContent`)
	if err != nil {
		t.Fatalf("Failed to access textContent: %v", err)
	}

	if !strings.Contains(result.String(), "Hello World") {
		t.Errorf("Expected text content to contain 'Hello World', got '%s'", result.String())
	}
}

func TestErrorHandlingInBindings(t *testing.T) {
	doc := dom.NewDocument()
	vm := goja.New()
	bindings := NewDOMBindings(vm, doc)

	docJS := bindings.WrapDocument()
	vm.Set("document", docJS)

	// Test createElement with no arguments (should throw error)
	_, err := vm.RunString(`
		try {
			document.createElement();
		} catch(e) {
			'error caught';
		}
	`)
	if err != nil {
		t.Fatalf("Error handling test failed: %v", err)
	}

	// Test setAttribute with insufficient arguments
	_, err = vm.RunString(`
		var div = document.createElement('div');
		try {
			div.setAttribute('id');
		} catch(e) {
			'error caught';
		}
	`)
	if err != nil {
		t.Fatalf("setAttribute error handling test failed: %v", err)
	}
}

func TestEventMethodsInBindings(t *testing.T) {
	doc := dom.NewDocument()
	vm := goja.New()
	bindings := NewDOMBindings(vm, doc)

	element := doc.CreateElement("div")
	elemJS := bindings.WrapElement(element)

	vm.Set("element", elemJS)

	// Test addEventListener exists
	_, err := vm.RunString(`
		if (typeof element.addEventListener !== 'function') {
			throw new Error('addEventListener should be a function');
		}
	`)
	if err != nil {
		t.Fatalf("addEventListener availability test failed: %v", err)
	}

	// Test removeEventListener exists
	_, err = vm.RunString(`
		if (typeof element.removeEventListener !== 'function') {
			throw new Error('removeEventListener should be a function');
		}
	`)
	if err != nil {
		t.Fatalf("removeEventListener availability test failed: %v", err)
	}

	// Test dispatchEvent exists
	_, err = vm.RunString(`
		if (typeof element.dispatchEvent !== 'function') {
			throw new Error('dispatchEvent should be a function');
		}
	`)
	if err != nil {
		t.Fatalf("dispatchEvent availability test failed: %v", err)
	}
}

func TestNodeListWrapping(t *testing.T) {
	doc := dom.NewDocument()
	vm := goja.New()
	bindings := NewDOMBindings(vm, doc)

	// Create test elements
	elements := []*dom.Element{
		doc.CreateElement("div"),
		doc.CreateElement("span"),
		doc.CreateElement("p"),
	}

	nodeList := bindings.WrapNodeList(elements)

	// Test length property
	length := nodeList.Get("length").ToInteger()
	if length != int64(len(elements)) {
		t.Errorf("Expected length %d, got %d", len(elements), length)
	}

	// Test array-like access
	firstElement := nodeList.Get("0")
	if firstElement == nil {
		t.Error("NodeList[0] should be accessible")
	}

	// Verify it's the wrapped version of our first element
	tagName := firstElement.ToObject(vm).Get("tagName").String()
	if tagName != "div" {
		t.Errorf("Expected first element tagName 'div', got '%s'", tagName)
	}
}

func TestTextNodeWrapping(t *testing.T) {
	doc := dom.NewDocument()
	vm := goja.New()
	bindings := NewDOMBindings(vm, doc)

	textNode := doc.CreateTextNode("Hello, World!")
	textJS := bindings.WrapNode(textNode)

	// Test basic text node properties
	nodeType := textJS.Get("nodeType").ToInteger()
	if nodeType != int64(dom.TextNode) {
		t.Errorf("Expected nodeType %d, got %d", dom.TextNode, nodeType)
	}

	nodeName := textJS.Get("nodeName").String()
	if nodeName != "#text" {
		t.Errorf("Expected nodeName '#text', got '%s'", nodeName)
	}

	nodeValue := textJS.Get("nodeValue").String()
	if nodeValue != "Hello, World!" {
		t.Errorf("Expected nodeValue 'Hello, World!', got '%s'", nodeValue)
	}

	// Test data property
	data := textJS.Get("data").String()
	if data != "Hello, World!" {
		t.Errorf("Expected data 'Hello, World!', got '%s'", data)
	}
}

func TestCommentNodeWrapping(t *testing.T) {
	doc := dom.NewDocument()
	vm := goja.New()
	bindings := NewDOMBindings(vm, doc)

	comment := doc.CreateComment("This is a comment")
	commentJS := bindings.WrapNode(comment)

	// Test basic comment node properties
	nodeType := commentJS.Get("nodeType").ToInteger()
	if nodeType != int64(dom.CommentNode) {
		t.Errorf("Expected nodeType %d, got %d", dom.CommentNode, nodeType)
	}

	nodeName := commentJS.Get("nodeName").String()
	if nodeName != "#comment" {
		t.Errorf("Expected nodeName '#comment', got '%s'", nodeName)
	}

	nodeValue := commentJS.Get("nodeValue").String()
	if nodeValue != "This is a comment" {
		t.Errorf("Expected nodeValue 'This is a comment', got '%s'", nodeValue)
	}

	// Test data property
	data := commentJS.Get("data").String()
	if data != "This is a comment" {
		t.Errorf("Expected data 'This is a comment', got '%s'", data)
	}
}

func TestCloneNodeInJS(t *testing.T) {
	doc := dom.NewDocument()
	vm := goja.New()
	bindings := NewDOMBindings(vm, doc)

	element := doc.CreateElement("div")
	element.SetAttribute("id", "original")

	elemJS := bindings.WrapElement(element)
	vm.Set("element", elemJS)

	// Test shallow clone
	result, err := vm.RunString(`
		var cloned = element.cloneNode(false);
		cloned.getAttribute('id');
	`)
	if err != nil {
		t.Fatalf("cloneNode test failed: %v", err)
	}

	if result.String() != "original" {
		t.Errorf("Cloned element should preserve attributes, got '%s'", result.String())
	}

	// Test that cloned node is different object
	_, err = vm.RunString(`
		var cloned = element.cloneNode(false);
		if (cloned === element) {
			throw new Error('Cloned node should be different object');
		}
	`)
	if err != nil {
		t.Fatalf("Clone identity test failed: %v", err)
	}
}

func TestDocumentMethodsWithParser(t *testing.T) {
	// Create document with parsed HTML
	html := `<html><body><div id="test" class="sample">Content</div><p class="sample">Paragraph</p></body></html>`

	p := parser.NewParser()
	doc, err := p.Parse(strings.NewReader(html))
	if err != nil {
		t.Fatalf("Failed to parse HTML: %v", err)
	}

	vm := goja.New()
	bindings := NewDOMBindings(vm, doc)
	docJS := bindings.WrapDocument()
	vm.Set("document", docJS)

	// Test getElementById
	result, err := vm.RunString(`
		var element = document.getElementById('test');
		element ? element.tagName : null;
	`)
	if err != nil {
		t.Fatalf("getElementById test failed: %v", err)
	}

	if result.String() != "div" {
		t.Errorf("getElementById should find div element, got '%s'", result.String())
	}

	// Test getElementsByClassName
	result, err = vm.RunString(`
		var elements = document.getElementsByClassName('sample');
		elements.length;
	`)
	if err != nil {
		t.Fatalf("getElementsByClassName test failed: %v", err)
	}

	if result.ToInteger() != 2 {
		t.Errorf("getElementsByClassName should find 2 elements, got %d", result.ToInteger())
	}

	// Test getElementsByTagName
	result, err = vm.RunString(`
		var divs = document.getElementsByTagName('div');
		divs.length;
	`)
	if err != nil {
		t.Fatalf("getElementsByTagName test failed: %v", err)
	}

	if result.ToInteger() < 1 {
		t.Errorf("getElementsByTagName should find at least 1 div, got %d", result.ToInteger())
	}
}

func TestGetTextContentHelper(t *testing.T) {
	doc := dom.NewDocument()
	vm := goja.New()
	bindings := NewDOMBindings(vm, doc)

	// Test with Element
	element := doc.CreateElement("div")
	text := doc.CreateTextNode("Hello")
	element.AppendChild(text)

	textContent := bindings.getTextContent(element)
	if textContent != "Hello" {
		t.Errorf("Expected 'Hello', got '%s'", textContent)
	}

	// Test with Text node
	textNode := doc.CreateTextNode("World")
	textContent = bindings.getTextContent(textNode)
	if textContent != "World" {
		t.Errorf("Expected 'World', got '%s'", textContent)
	}

	// Test with Comment node
	comment := doc.CreateComment("Comment")
	textContent = bindings.getTextContent(comment)
	if textContent != "" {
		t.Errorf("Comment textContent should be empty, got '%s'", textContent)
	}
}

func TestNullElementHandling(t *testing.T) {
	doc := dom.NewDocument()
	vm := goja.New()
	bindings := NewDOMBindings(vm, doc)

	// Test wrapping null element
	result := bindings.WrapElement(nil)
	if result != nil {
		t.Error("WrapElement(nil) should return nil")
	}

	// Test wrapping null node
	result = bindings.WrapNode(nil)
	if result != nil {
		t.Error("WrapNode(nil) should return nil")
	}
}

func TestCacheMemoryManagement(t *testing.T) {
	doc := dom.NewDocument()
	vm := goja.New()
	bindings := NewDOMBindings(vm, doc)

	// Create multiple elements and wrap them
	for i := 0; i < 10; i++ {
		element := doc.CreateElement("div")
		bindings.WrapElement(element)
	}

	// Verify cache contains expected number of elements
	if len(bindings.nodeCache) != 10 {
		t.Errorf("Expected 10 cached elements, got %d", len(bindings.nodeCache))
	}

	// Test that same elements return same objects
	element := doc.CreateElement("span")
	obj1 := bindings.WrapElement(element)
	obj2 := bindings.WrapElement(element)

	if obj1 != obj2 {
		t.Error("Same element should return identical cached object")
	}

	// Cache should still be 11 (10 divs + 1 span)
	if len(bindings.nodeCache) != 11 {
		t.Errorf("Expected 11 cached elements, got %d", len(bindings.nodeCache))
	}
}
