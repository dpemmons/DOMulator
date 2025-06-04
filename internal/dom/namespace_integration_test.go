package dom

import (
	"testing"
)

func TestElementNS(t *testing.T) {
	doc := NewDocument()

	tests := []struct {
		name           string
		namespaceURI   string
		qualifiedName  string
		expectError    bool
		expectedNS     string
		expectedPrefix string
		expectedLocal  string
	}{
		{
			name:           "simple HTML element",
			namespaceURI:   "http://www.w3.org/1999/xhtml",
			qualifiedName:  "div",
			expectError:    false,
			expectedNS:     "http://www.w3.org/1999/xhtml",
			expectedPrefix: "",
			expectedLocal:  "div",
		},
		{
			name:           "SVG element with prefix",
			namespaceURI:   "http://www.w3.org/2000/svg",
			qualifiedName:  "svg:rect",
			expectError:    false,
			expectedNS:     "http://www.w3.org/2000/svg",
			expectedPrefix: "svg",
			expectedLocal:  "rect",
		},
		{
			name:           "XML element",
			namespaceURI:   "http://www.w3.org/XML/1998/namespace",
			qualifiedName:  "xml:lang",
			expectError:    false,
			expectedNS:     "http://www.w3.org/XML/1998/namespace",
			expectedPrefix: "xml",
			expectedLocal:  "lang",
		},
		{
			name:          "prefix without namespace",
			namespaceURI:  "",
			qualifiedName: "ns:element",
			expectError:   true,
		},
		{
			name:          "xml prefix with wrong namespace",
			namespaceURI:  "http://example.com",
			qualifiedName: "xml:element",
			expectError:   true,
		},
		{
			name:          "xmlns prefix with wrong namespace",
			namespaceURI:  "http://example.com",
			qualifiedName: "xmlns:attr",
			expectError:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			elem, err := NewElementNS(tt.namespaceURI, tt.qualifiedName, doc)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			if elem.NamespaceURI() != tt.expectedNS {
				t.Errorf("Expected namespace URI %q, got %q", tt.expectedNS, elem.NamespaceURI())
			}

			if elem.Prefix() != tt.expectedPrefix {
				t.Errorf("Expected prefix %q, got %q", tt.expectedPrefix, elem.Prefix())
			}

			if elem.LocalName() != tt.expectedLocal {
				t.Errorf("Expected local name %q, got %q", tt.expectedLocal, elem.LocalName())
			}

			if elem.TagName() != tt.qualifiedName {
				t.Errorf("Expected tag name %q, got %q", tt.qualifiedName, elem.TagName())
			}
		})
	}
}

func TestDocumentCreateElementNS(t *testing.T) {
	doc := NewDocument()

	// Test successful creation
	elem, err := doc.CreateElementNS("http://www.w3.org/1999/xhtml", "div")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if elem == nil {
		t.Error("Expected element, got nil")
	}
	if elem.NamespaceURI() != "http://www.w3.org/1999/xhtml" {
		t.Errorf("Expected HTML namespace, got %q", elem.NamespaceURI())
	}

	// Test error case
	_, err = doc.CreateElementNS("", "ns:element")
	if err == nil {
		t.Error("Expected error for prefix without namespace")
	}
}

func TestElementAttributesNS(t *testing.T) {
	doc := NewDocument()
	elem, _ := doc.CreateElementNS("http://www.w3.org/1999/xhtml", "div")

	// Test SetAttributeNS
	err := elem.SetAttributeNS("http://www.w3.org/XML/1998/namespace", "xml:lang", "en")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Test GetAttributeNS
	value := elem.GetAttributeNS("http://www.w3.org/XML/1998/namespace", "lang")
	if value != "en" {
		t.Errorf("Expected 'en', got %q", value)
	}

	// Test HasAttributeNS
	if !elem.HasAttributeNS("http://www.w3.org/XML/1998/namespace", "lang") {
		t.Error("Expected attribute to exist")
	}

	if elem.HasAttributeNS("http://www.w3.org/XML/1998/namespace", "nonexistent") {
		t.Error("Expected attribute to not exist")
	}

	// Test RemoveAttributeNS
	elem.RemoveAttributeNS("http://www.w3.org/XML/1998/namespace", "lang")
	if elem.HasAttributeNS("http://www.w3.org/XML/1998/namespace", "lang") {
		t.Error("Expected attribute to be removed")
	}

	// Test non-namespace attributes
	elem.SetAttribute("class", "test")
	if !elem.HasAttributeNS("", "class") {
		t.Error("Expected non-namespace attribute to be found")
	}

	value = elem.GetAttributeNS("", "class")
	if value != "test" {
		t.Errorf("Expected 'test', got %q", value)
	}
}

func TestElementGetElementsByTagNameNS(t *testing.T) {
	doc := NewDocument()
	root, _ := doc.CreateElementNS("http://www.w3.org/1999/xhtml", "html")
	doc.AppendChild(root)

	// Create some elements
	div1, _ := doc.CreateElementNS("http://www.w3.org/1999/xhtml", "div")
	div2, _ := doc.CreateElementNS("http://www.w3.org/1999/xhtml", "div")
	svgRect, _ := doc.CreateElementNS("http://www.w3.org/2000/svg", "rect")

	root.AppendChild(div1)
	root.AppendChild(div2)
	root.AppendChild(svgRect)

	// Test finding HTML div elements
	divs := root.GetElementsByTagNameNS("http://www.w3.org/1999/xhtml", "div")
	if len(divs) != 2 {
		t.Errorf("Expected 2 div elements, got %d", len(divs))
	}

	// Test finding SVG elements
	rects := root.GetElementsByTagNameNS("http://www.w3.org/2000/svg", "rect")
	if len(rects) != 1 {
		t.Errorf("Expected 1 rect element, got %d", len(rects))
	}

	// Test wildcard namespace
	allDivs := root.GetElementsByTagNameNS("*", "div")
	if len(allDivs) != 2 {
		t.Errorf("Expected 2 div elements with wildcard namespace, got %d", len(allDivs))
	}

	// Test wildcard local name
	allElements := root.GetElementsByTagNameNS("*", "*")
	if len(allElements) != 3 {
		t.Errorf("Expected 3 elements with wildcard namespace and local name, got %d", len(allElements))
	}

	// Test no matches
	spans := root.GetElementsByTagNameNS("http://www.w3.org/1999/xhtml", "span")
	if len(spans) != 0 {
		t.Errorf("Expected 0 span elements, got %d", len(spans))
	}
}

func TestDocumentGetElementsByTagNameNS(t *testing.T) {
	doc := NewDocument()
	root, _ := doc.CreateElementNS("http://www.w3.org/1999/xhtml", "html")
	doc.AppendChild(root)

	// Create some elements
	div1, _ := doc.CreateElementNS("http://www.w3.org/1999/xhtml", "div")
	div2, _ := doc.CreateElementNS("http://www.w3.org/1999/xhtml", "div")
	svgRect, _ := doc.CreateElementNS("http://www.w3.org/2000/svg", "rect")

	root.AppendChild(div1)
	root.AppendChild(div2)
	root.AppendChild(svgRect)

	// Test finding HTML div elements
	divs := doc.GetElementsByTagNameNS("http://www.w3.org/1999/xhtml", "div")
	if len(divs) != 2 {
		t.Errorf("Expected 2 div elements, got %d", len(divs))
	}

	// Test finding SVG elements
	rects := doc.GetElementsByTagNameNS("http://www.w3.org/2000/svg", "rect")
	if len(rects) != 1 {
		t.Errorf("Expected 1 rect element, got %d", len(rects))
	}

	// Test finding HTML root element
	htmls := doc.GetElementsByTagNameNS("http://www.w3.org/1999/xhtml", "html")
	if len(htmls) != 1 {
		t.Errorf("Expected 1 html element, got %d", len(htmls))
	}
}

func TestNamespacePrefixParsing(t *testing.T) {
	doc := NewDocument()

	// Test creating element with well-known prefix
	elem := NewElement("svg:rect", doc)
	if elem.NamespaceURI() != "http://www.w3.org/2000/svg" {
		t.Errorf("Expected SVG namespace for svg: prefix, got %q", elem.NamespaceURI())
	}
	if elem.Prefix() != "svg" {
		t.Errorf("Expected prefix 'svg', got %q", elem.Prefix())
	}
	if elem.LocalName() != "rect" {
		t.Errorf("Expected local name 'rect', got %q", elem.LocalName())
	}

	// Test creating element without prefix
	elem2 := NewElement("div", doc)
	if elem2.NamespaceURI() != "" {
		t.Errorf("Expected empty namespace for unprefixed element, got %q", elem2.NamespaceURI())
	}
	if elem2.Prefix() != "" {
		t.Errorf("Expected empty prefix, got %q", elem2.Prefix())
	}
	if elem2.LocalName() != "div" {
		t.Errorf("Expected local name 'div', got %q", elem2.LocalName())
	}
}

func TestNamespaceErrorPropagation(t *testing.T) {
	doc := NewDocument()

	// Test that namespace errors are properly propagated from SetAttributeNS
	elem, _ := doc.CreateElementNS("http://www.w3.org/1999/xhtml", "div")

	err := elem.SetAttributeNS("", "ns:attr", "value")
	if err == nil {
		t.Error("Expected error for prefix without namespace in SetAttributeNS")
	}

	err = elem.SetAttributeNS("http://example.com", "xml:attr", "value")
	if err == nil {
		t.Error("Expected error for xml prefix with wrong namespace in SetAttributeNS")
	}
}

func TestMixedNamespaceDocument(t *testing.T) {
	doc := NewDocument()

	// Create a mixed namespace document structure
	html, _ := doc.CreateElementNS("http://www.w3.org/1999/xhtml", "html")
	doc.AppendChild(html)

	body, _ := doc.CreateElementNS("http://www.w3.org/1999/xhtml", "body")
	html.AppendChild(body)

	// Add SVG content
	svg, _ := doc.CreateElementNS("http://www.w3.org/2000/svg", "svg")
	body.AppendChild(svg)

	rect, _ := doc.CreateElementNS("http://www.w3.org/2000/svg", "rect")
	svg.AppendChild(rect)

	circle, _ := doc.CreateElementNS("http://www.w3.org/2000/svg", "circle")
	svg.AppendChild(circle)

	// Add some HTML content
	div, _ := doc.CreateElementNS("http://www.w3.org/1999/xhtml", "div")
	body.AppendChild(div)

	// Test searching across namespaces
	svgElements := doc.GetElementsByTagNameNS("http://www.w3.org/2000/svg", "*")
	if len(svgElements) != 3 { // svg, rect, circle
		t.Errorf("Expected 3 SVG elements, got %d", len(svgElements))
	}

	htmlElements := doc.GetElementsByTagNameNS("http://www.w3.org/1999/xhtml", "*")
	if len(htmlElements) != 3 { // html, body, div
		t.Errorf("Expected 3 HTML elements, got %d", len(htmlElements))
	}

	// Test finding specific elements
	rects := doc.GetElementsByTagNameNS("http://www.w3.org/2000/svg", "rect")
	if len(rects) != 1 {
		t.Errorf("Expected 1 rect element, got %d", len(rects))
	}

	// Test from element level
	svgElementsFromBody := body.GetElementsByTagNameNS("http://www.w3.org/2000/svg", "*")
	if len(svgElementsFromBody) != 3 {
		t.Errorf("Expected 3 SVG elements from body search, got %d", len(svgElementsFromBody))
	}
}
