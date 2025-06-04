package dom

import (
	"strings"
	"testing"
)

// TestDocumentSpecCompliance tests Document interface compliance with WHATWG DOM specification
func TestDocumentSpecCompliance(t *testing.T) {
	tests := []struct {
		name string
		test func(t *testing.T)
	}{
		{"Constructor", testDocumentConstructor},
		{"Implementation", testDocumentImplementation},
		{"URL and DocumentURI", testDocumentURL},
		{"CompatMode", testDocumentCompatMode},
		{"CharacterSet", testDocumentCharacterSet},
		{"ContentType", testDocumentContentType},
		{"Doctype", testDocumentDoctype},
		{"DocumentElement", testDocumentElement},
		{"GetElementsByTagName", testDocumentGetElementsByTagName},
		{"GetElementsByTagNameNS", testDocumentGetElementsByTagNameNS},
		{"GetElementsByClassName", testDocumentGetElementsByClassName},
		{"CreateElement", testDocumentCreateElement},
		{"CreateElementNS", testDocumentCreateElementNS},
		{"CreateDocumentFragment", testDocumentCreateDocumentFragment},
		{"CreateTextNode", testDocumentCreateTextNode},
		{"CreateCDATASection", testDocumentCreateCDATASection},
		{"CreateComment", testDocumentCreateComment},
		{"CreateProcessingInstruction", testDocumentCreateProcessingInstruction},
		{"ImportNode", testDocumentImportNode},
		{"AdoptNode", testDocumentAdoptNode},
		{"CreateAttribute", testDocumentCreateAttribute},
		{"CreateAttributeNS", testDocumentCreateAttributeNS},
		{"CreateEvent", testDocumentCreateEvent},
		{"CreateRange", testDocumentCreateRange},
		{"CreateNodeIterator", testDocumentCreateNodeIterator},
		{"CreateTreeWalker", testDocumentCreateTreeWalker},
	}

	for _, tt := range tests {
		t.Run(tt.name, tt.test)
	}
}

func testDocumentConstructor(t *testing.T) {
	doc := NewDocument()

	// Test default properties per specification
	if doc.URL() != "about:blank" {
		t.Errorf("Expected URL to be 'about:blank', got %s", doc.URL())
	}

	if doc.DocumentURI() != "about:blank" {
		t.Errorf("Expected DocumentURI to be 'about:blank', got %s", doc.DocumentURI())
	}

	if doc.CompatMode() != "CSS1Compat" {
		t.Errorf("Expected CompatMode to be 'CSS1Compat', got %s", doc.CompatMode())
	}

	if doc.CharacterSet() != "UTF-8" {
		t.Errorf("Expected CharacterSet to be 'UTF-8', got %s", doc.CharacterSet())
	}

	if doc.ContentType() != "application/xml" {
		t.Errorf("Expected ContentType to be 'application/xml', got %s", doc.ContentType())
	}
}

func testDocumentImplementation(t *testing.T) {
	doc := NewDocument()
	impl := doc.Implementation()

	if impl == nil {
		t.Error("Expected Implementation() to return a DOMImplementation object")
	}

	// Test DOMImplementation methods
	if !impl.HasFeature("", "") {
		t.Error("Expected HasFeature to always return true")
	}

	doctype, err := impl.CreateDocumentType("html", "", "")
	if err != nil {
		t.Errorf("Expected CreateDocumentType to succeed, got error: %v", err)
	}
	if doctype == nil {
		t.Error("Expected CreateDocumentType to return a DocumentType")
	}
}

func testDocumentURL(t *testing.T) {
	doc := NewDocument()

	// URL and DocumentURI should be aliases
	if doc.URL() != doc.DocumentURI() {
		t.Error("Expected URL and DocumentURI to be aliases")
	}
}

func testDocumentCompatMode(t *testing.T) {
	doc := NewDocument()

	// Test no-quirks mode
	doc.mode = "no-quirks"
	if doc.CompatMode() != "CSS1Compat" {
		t.Errorf("Expected CompatMode to be 'CSS1Compat' for no-quirks mode, got %s", doc.CompatMode())
	}

	// Test quirks mode
	doc.mode = "quirks"
	if doc.CompatMode() != "BackCompat" {
		t.Errorf("Expected CompatMode to be 'BackCompat' for quirks mode, got %s", doc.CompatMode())
	}

	// Test limited-quirks mode
	doc.mode = "limited-quirks"
	if doc.CompatMode() != "CSS1Compat" {
		t.Errorf("Expected CompatMode to be 'CSS1Compat' for limited-quirks mode, got %s", doc.CompatMode())
	}
}

func testDocumentCharacterSet(t *testing.T) {
	doc := NewDocument()

	// CharacterSet, Charset, and InputEncoding should be aliases
	if doc.CharacterSet() != doc.Charset() || doc.CharacterSet() != doc.InputEncoding() {
		t.Error("Expected CharacterSet, Charset, and InputEncoding to be aliases")
	}
}

func testDocumentContentType(t *testing.T) {
	doc := NewDocument()

	// Default should be application/xml
	if doc.ContentType() != "application/xml" {
		t.Errorf("Expected default ContentType to be 'application/xml', got %s", doc.ContentType())
	}
}

func testDocumentDoctype(t *testing.T) {
	doc := NewDocument()

	// Initially no doctype
	if doc.Doctype() != nil {
		t.Error("Expected Doctype() to return nil for new document")
	}

	// Add a doctype
	doctype := NewDocumentType("html", "", "", doc)
	doc.AppendChild(doctype)

	if doc.Doctype() != doctype {
		t.Error("Expected Doctype() to return the added DocumentType")
	}
}

func testDocumentElement(t *testing.T) {
	doc := NewDocument()

	// Initially no document element
	if doc.DocumentElement() != nil {
		t.Error("Expected DocumentElement() to return nil for new document")
	}

	// Add an html element
	html := doc.CreateElement("html")
	doc.AppendChild(html)

	if doc.DocumentElement() != html {
		t.Error("Expected DocumentElement() to return the html element")
	}
}

func testDocumentGetElementsByTagName(t *testing.T) {
	doc := NewDocument()
	html := doc.CreateElement("html")
	body := doc.CreateElement("body")
	div1 := doc.CreateElement("div")
	div2 := doc.CreateElement("div")

	doc.AppendChild(html)
	html.AppendChild(body)
	body.AppendChild(div1)
	body.AppendChild(div2)

	// Test getElementsByTagName
	divs := doc.GetElementsByTagName("div")
	if divs.Length() != 2 {
		t.Errorf("Expected 2 div elements, got %d", divs.Length())
	}

	// Test wildcard
	all := doc.GetElementsByTagName("*")
	if all.Length() != 4 { // html, body, div, div
		t.Errorf("Expected 4 elements with '*', got %d", all.Length())
	}
}

func testDocumentGetElementsByTagNameNS(t *testing.T) {
	doc := NewDocument()

	collection := doc.GetElementsByTagNameNS("http://www.w3.org/1999/xhtml", "div")
	if collection == nil {
		t.Error("Expected GetElementsByTagNameNS to return a collection")
	}
}

func testDocumentGetElementsByClassName(t *testing.T) {
	doc := NewDocument()
	html := doc.CreateElement("html")
	div1 := doc.CreateElement("div")
	div2 := doc.CreateElement("div")

	div1.SetAttribute("class", "aaa bbb")
	div2.SetAttribute("class", "aaa ccc")

	doc.AppendChild(html)
	html.AppendChild(div1)
	html.AppendChild(div2)

	// Test getElementsByClassName
	elements := doc.GetElementsByClassName("aaa")
	if elements.Length() != 2 {
		t.Errorf("Expected 2 elements with class 'aaa', got %d", elements.Length())
	}

	elements = doc.GetElementsByClassName("bbb")
	if elements.Length() != 1 {
		t.Errorf("Expected 1 element with class 'bbb', got %d", elements.Length())
	}
}

func testDocumentCreateElement(t *testing.T) {
	doc := NewDocument()

	// Test normal case
	elem := doc.CreateElement("div")
	if elem == nil {
		t.Error("Expected CreateElement to return an element")
	}
	if elem.TagName() != "div" {
		t.Errorf("Expected tagName to be 'div', got %s", elem.TagName())
	}

	// Test HTML document lowercasing
	doc.documentType = "html"
	elem = doc.CreateElement("DIV")
	if elem.TagName() != "DIV" { // Note: TagName() returns the original case, but internally it's lowercased
		t.Errorf("Expected tagName to be 'DIV', got %s", elem.TagName())
	}
}

func testDocumentCreateElementNS(t *testing.T) {
	doc := NewDocument()

	elem, err := doc.CreateElementNS("http://www.w3.org/1999/xhtml", "div")
	if err != nil {
		t.Errorf("Expected CreateElementNS to succeed, got error: %v", err)
	}
	if elem == nil {
		t.Error("Expected CreateElementNS to return an element")
	}
}

func testDocumentCreateDocumentFragment(t *testing.T) {
	doc := NewDocument()

	fragment := doc.CreateDocumentFragment()
	if fragment == nil {
		t.Error("Expected CreateDocumentFragment to return a DocumentFragment")
	}
	if fragment.NodeType() != DocumentFragmentNode {
		t.Error("Expected DocumentFragment to have correct node type")
	}
}

func testDocumentCreateTextNode(t *testing.T) {
	doc := NewDocument()

	text := doc.CreateTextNode("Hello, World!")
	if text == nil {
		t.Error("Expected CreateTextNode to return a Text node")
	}
	if text.NodeType() != TextNode {
		t.Error("Expected Text node to have correct node type")
	}
	if text.Data() != "Hello, World!" {
		t.Errorf("Expected text data to be 'Hello, World!', got %s", text.Data())
	}
}

func testDocumentCreateCDATASection(t *testing.T) {
	doc := NewDocument()

	// Test XML document (should succeed)
	doc.documentType = "xml"
	cdata, err := doc.CreateCDATASection("some data")
	if err != nil {
		t.Errorf("Expected CreateCDATASection to succeed in XML document, got error: %v", err)
	}
	if cdata == nil {
		t.Error("Expected CreateCDATASection to return a CDATASection")
	}

	// Test HTML document (should fail)
	doc.documentType = "html"
	_, err = doc.CreateCDATASection("some data")
	if err == nil {
		t.Error("Expected CreateCDATASection to fail in HTML document")
	}
	if !strings.Contains(err.Error(), "NotSupportedError") {
		t.Errorf("Expected NotSupportedError, got %v", err)
	}

	// Test invalid data
	doc.documentType = "xml"
	_, err = doc.CreateCDATASection("data with ]]> end marker")
	if err == nil {
		t.Error("Expected CreateCDATASection to fail with invalid data")
	}
	if !strings.Contains(err.Error(), "InvalidCharacterError") {
		t.Errorf("Expected InvalidCharacterError, got %v", err)
	}
}

func testDocumentCreateComment(t *testing.T) {
	doc := NewDocument()

	comment := doc.CreateComment("This is a comment")
	if comment == nil {
		t.Error("Expected CreateComment to return a Comment node")
	}
	if comment.NodeType() != CommentNode {
		t.Error("Expected Comment node to have correct node type")
	}
	if comment.Data() != "This is a comment" {
		t.Errorf("Expected comment data to be 'This is a comment', got %s", comment.Data())
	}
}

func testDocumentCreateProcessingInstruction(t *testing.T) {
	doc := NewDocument()

	// Test valid case
	pi, err := doc.CreateProcessingInstruction("xml-stylesheet", "href='style.css'")
	if err != nil {
		t.Errorf("Expected CreateProcessingInstruction to succeed, got error: %v", err)
	}
	if pi == nil {
		t.Error("Expected CreateProcessingInstruction to return a ProcessingInstruction")
	}

	// Test invalid target
	_, err = doc.CreateProcessingInstruction("", "data")
	if err == nil {
		t.Error("Expected CreateProcessingInstruction to fail with empty target")
	}

	// Test invalid data
	_, err = doc.CreateProcessingInstruction("target", "data with ?> end marker")
	if err == nil {
		t.Error("Expected CreateProcessingInstruction to fail with invalid data")
	}
}

func testDocumentImportNode(t *testing.T) {
	doc1 := NewDocument()
	doc2 := NewDocument()

	// Create a node in doc1
	elem := doc1.CreateElement("div")
	text := doc1.CreateTextNode("Hello")
	elem.AppendChild(text)

	// Import into doc2
	imported, err := doc2.ImportNode(elem, true)
	if err != nil {
		t.Errorf("Expected ImportNode to succeed, got error: %v", err)
	}
	if imported == nil {
		t.Error("Expected ImportNode to return a node")
	}

	// Test importing document (should fail)
	_, err = doc2.ImportNode(doc1, false)
	if err == nil {
		t.Error("Expected ImportNode to fail when importing a document")
	}
}

func testDocumentAdoptNode(t *testing.T) {
	doc1 := NewDocument()
	doc2 := NewDocument()

	// Create a node in doc1
	elem := doc1.CreateElement("div")

	// Adopt into doc2
	adopted, err := doc2.AdoptNode(elem)
	if err != nil {
		t.Errorf("Expected AdoptNode to succeed, got error: %v", err)
	}
	if adopted != elem {
		t.Error("Expected AdoptNode to return the same node")
	}

	// Test adopting document (should fail)
	_, err = doc2.AdoptNode(doc1)
	if err == nil {
		t.Error("Expected AdoptNode to fail when adopting a document")
	}
}

func testDocumentCreateAttribute(t *testing.T) {
	doc := NewDocument()

	// Test normal case
	attr, err := doc.CreateAttribute("id")
	if err != nil {
		t.Errorf("Expected CreateAttribute to succeed, got error: %v", err)
	}
	if attr == nil {
		t.Error("Expected CreateAttribute to return an attribute")
	}
	if attr.Name() != "id" {
		t.Errorf("Expected attribute name to be 'id', got %s", attr.Name())
	}

	// Test empty name
	_, err = doc.CreateAttribute("")
	if err == nil {
		t.Error("Expected CreateAttribute to fail with empty name")
	}
}

func testDocumentCreateAttributeNS(t *testing.T) {
	doc := NewDocument()

	// Test normal case
	attr, err := doc.CreateAttributeNS("http://www.w3.org/1999/xhtml", "id")
	if err != nil {
		t.Errorf("Expected CreateAttributeNS to succeed, got error: %v", err)
	}
	if attr == nil {
		t.Error("Expected CreateAttributeNS to return an attribute")
	}

	// Test empty qualified name
	_, err = doc.CreateAttributeNS("http://www.w3.org/1999/xhtml", "")
	if err == nil {
		t.Error("Expected CreateAttributeNS to fail with empty qualified name")
	}
}

func testDocumentCreateEvent(t *testing.T) {
	doc := NewDocument()

	// Test supported interfaces
	supportedInterfaces := []string{
		"Event", "events", "HTMLEvents", "SVGEvents",
		"CustomEvent",
		"MouseEvent", "mouseevents",
		"KeyboardEvent",
		"UIEvent", "uievents",
	}

	for _, interfaceName := range supportedInterfaces {
		event, err := doc.CreateEvent(interfaceName)
		if err != nil {
			t.Errorf("Expected CreateEvent to succeed for %s, got error: %v", interfaceName, err)
		}
		if event == nil {
			t.Errorf("Expected CreateEvent to return an event for %s", interfaceName)
		}
	}

	// Test unsupported interface
	_, err := doc.CreateEvent("UnsupportedEvent")
	if err == nil {
		t.Error("Expected CreateEvent to fail for unsupported interface")
	}
}

func testDocumentCreateRange(t *testing.T) {
	doc := NewDocument()

	range_ := doc.CreateRange()
	if range_ == nil {
		t.Error("Expected CreateRange to return a range object")
	}

	// Should be a map with expected properties
	if rangeMap, ok := range_.(map[string]interface{}); ok {
		if rangeMap["startContainer"] != doc {
			t.Error("Expected range startContainer to be the document")
		}
		if rangeMap["collapsed"] != true {
			t.Error("Expected range to be collapsed")
		}
	} else {
		t.Error("Expected CreateRange to return a map")
	}
}

func testDocumentCreateNodeIterator(t *testing.T) {
	doc := NewDocument()
	root := doc.CreateElement("div")

	iterator := doc.CreateNodeIterator(root, 0xFFFFFFFF, nil)
	if iterator == nil {
		t.Error("Expected CreateNodeIterator to return an iterator")
	}

	// Should be a map with expected properties
	if iteratorMap, ok := iterator.(map[string]interface{}); ok {
		if iteratorMap["root"] != root {
			t.Error("Expected iterator root to be the specified root")
		}
		if iteratorMap["whatToShow"] != uint32(0xFFFFFFFF) {
			t.Error("Expected iterator whatToShow to match input")
		}
	} else {
		t.Error("Expected CreateNodeIterator to return a map")
	}
}

func testDocumentCreateTreeWalker(t *testing.T) {
	doc := NewDocument()
	root := doc.CreateElement("div")

	walker := doc.CreateTreeWalker(root, 0xFFFFFFFF, nil)
	if walker == nil {
		t.Error("Expected CreateTreeWalker to return a walker")
	}

	// Should be a map with expected properties
	if walkerMap, ok := walker.(map[string]interface{}); ok {
		if walkerMap["root"] != root {
			t.Error("Expected walker root to be the specified root")
		}
		if walkerMap["currentNode"] != root {
			t.Error("Expected walker currentNode to be the root")
		}
	} else {
		t.Error("Expected CreateTreeWalker to return a map")
	}
}

// TestDocumentElementCreationOptions tests element creation with options
func TestDocumentElementCreationOptions(t *testing.T) {
	doc := NewDocument()

	// Test createElement with default options
	elem := doc.CreateElement("div")
	if elem == nil {
		t.Error("Expected createElement to return an element")
	}

	// Test createElementNS with default options
	elemNS, err := doc.CreateElementNS("http://www.w3.org/1999/xhtml", "div")
	if err != nil {
		t.Errorf("Expected createElementNS to succeed, got error: %v", err)
	}
	if elemNS == nil {
		t.Error("Expected createElementNS to return an element")
	}
}

// TestDocumentSpec4_5Examples tests examples from the specification section 4.5
func TestDocumentSpec4_5Examples(t *testing.T) {
	doc := NewDocument()

	// Create structure for getElementsByClassName example
	html := doc.CreateElement("html")
	body := doc.CreateElement("body")
	div := doc.CreateElement("div")
	div.SetAttribute("id", "example")

	p1 := doc.CreateElement("p")
	p1.SetAttribute("id", "p1")
	p1.SetAttribute("class", "aaa bbb")

	p2 := doc.CreateElement("p")
	p2.SetAttribute("id", "p2")
	p2.SetAttribute("class", "aaa ccc")

	p3 := doc.CreateElement("p")
	p3.SetAttribute("id", "p3")
	p3.SetAttribute("class", "bbb ccc")

	doc.AppendChild(html)
	html.AppendChild(body)
	body.AppendChild(div)
	div.AppendChild(p1)
	div.AppendChild(p2)
	div.AppendChild(p3)

	// Test the examples from the specification

	// document.getElementById("example").getElementsByClassName("aaa") should return p1 and p2
	example := doc.GetElementById("example")
	if example == nil {
		t.Error("Expected to find element with id 'example'")
	}

	aaa := example.GetElementsByClassName("aaa")
	if aaa.Length() != 2 {
		t.Errorf("Expected 2 elements with class 'aaa', got %d", aaa.Length())
	}

	// getElementsByClassName("ccc bbb") should return only p3
	cccBbb := doc.GetElementsByClassName("ccc bbb")
	if cccBbb.Length() != 1 {
		t.Errorf("Expected 1 element with classes 'ccc bbb', got %d", cccBbb.Length())
	}

	if cccBbb.Length() > 0 {
		elem := cccBbb.Item(0)
		if elem == nil {
			t.Error("Expected HTMLCollection to contain Element nodes")
		} else if elem.GetAttribute("id") != "p3" {
			t.Errorf("Expected element with id 'p3', got %s", elem.GetAttribute("id"))
		}
	}

	// getElementsByClassName("aaa,bbb") should return no nodes
	aaaBbb := doc.GetElementsByClassName("aaa,bbb")
	if aaaBbb.Length() != 0 {
		t.Errorf("Expected 0 elements with class 'aaa,bbb', got %d", aaaBbb.Length())
	}
}
