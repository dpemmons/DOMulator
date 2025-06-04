package dom

import (
	"testing"
)

func TestNodeNamespaceMethods(t *testing.T) {
	// Create a document with namespace-aware elements
	doc := NewDocument()

	// Create an HTML element with XHTML namespace
	htmlElem, err := NewElementNS("http://www.w3.org/1999/xhtml", "html", doc)
	if err != nil {
		t.Fatalf("Failed to create HTML element: %v", err)
	}
	doc.AppendChild(htmlElem)

	// Add xmlns declaration
	htmlElem.SetAttribute("xmlns", "http://www.w3.org/1999/xhtml")
	htmlElem.SetAttribute("xmlns:svg", "http://www.w3.org/2000/svg")

	// Create SVG element with SVG namespace
	svgElem, err := NewElementNS("http://www.w3.org/2000/svg", "svg:rect", doc)
	if err != nil {
		t.Fatalf("Failed to create SVG element: %v", err)
	}
	htmlElem.AppendChild(svgElem)

	// Create text node
	textNode := NewText("Hello World", doc)
	svgElem.AppendChild(textNode)

	t.Run("LookupPrefix", func(t *testing.T) {
		// Test cases for LookupPrefix method
		tests := []struct {
			name      string
			node      Node
			namespace string
			expected  string
		}{
			{
				name:      "HTML element looking up XHTML namespace",
				node:      htmlElem,
				namespace: "http://www.w3.org/1999/xhtml",
				expected:  "", // Default namespace has no prefix
			},
			{
				name:      "HTML element looking up SVG namespace",
				node:      htmlElem,
				namespace: "http://www.w3.org/2000/svg",
				expected:  "svg",
			},
			{
				name:      "SVG element looking up SVG namespace",
				node:      svgElem,
				namespace: "http://www.w3.org/2000/svg",
				expected:  "svg", // Should inherit from parent
			},
			{
				name:      "Text node looking up SVG namespace",
				node:      textNode,
				namespace: "http://www.w3.org/2000/svg",
				expected:  "svg", // Should inherit from parent
			},
			{
				name:      "Unknown namespace",
				node:      htmlElem,
				namespace: "http://example.com/unknown",
				expected:  "",
			},
			{
				name:      "Empty namespace",
				node:      htmlElem,
				namespace: "",
				expected:  "",
			},
			{
				name:      "Document node looking up namespace",
				node:      doc,
				namespace: "http://www.w3.org/1999/xhtml",
				expected:  "", // Document should delegate to document element
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := tt.node.LookupPrefix(tt.namespace)
				if result != tt.expected {
					t.Errorf("LookupPrefix(%q) = %q, expected %q", tt.namespace, result, tt.expected)
				}
			})
		}
	})

	t.Run("LookupNamespaceURI", func(t *testing.T) {
		// Test cases for LookupNamespaceURI method
		tests := []struct {
			name     string
			node     Node
			prefix   string
			expected string
		}{
			{
				name:     "HTML element looking up default namespace",
				node:     htmlElem,
				prefix:   "",
				expected: "http://www.w3.org/1999/xhtml",
			},
			{
				name:     "HTML element looking up svg prefix",
				node:     htmlElem,
				prefix:   "svg",
				expected: "http://www.w3.org/2000/svg",
			},
			{
				name:     "SVG element looking up svg prefix",
				node:     svgElem,
				prefix:   "svg",
				expected: "http://www.w3.org/2000/svg",
			},
			{
				name:     "Text node looking up svg prefix",
				node:     textNode,
				prefix:   "svg",
				expected: "http://www.w3.org/2000/svg", // Should inherit from parent
			},
			{
				name:     "Built-in xml prefix",
				node:     htmlElem,
				prefix:   "xml",
				expected: "http://www.w3.org/XML/1998/namespace",
			},
			{
				name:     "Built-in xmlns prefix",
				node:     htmlElem,
				prefix:   "xmlns",
				expected: "http://www.w3.org/2000/xmlns/",
			},
			{
				name:     "Unknown prefix",
				node:     htmlElem,
				prefix:   "unknown",
				expected: "",
			},
			{
				name:     "Document node looking up prefix",
				node:     doc,
				prefix:   "svg",
				expected: "http://www.w3.org/2000/svg", // Should delegate to document element
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := tt.node.LookupNamespaceURI(tt.prefix)
				if result != tt.expected {
					t.Errorf("LookupNamespaceURI(%q) = %q, expected %q", tt.prefix, result, tt.expected)
				}
			})
		}
	})

	t.Run("IsDefaultNamespace", func(t *testing.T) {
		// Test cases for IsDefaultNamespace method
		tests := []struct {
			name      string
			node      Node
			namespace string
			expected  bool
		}{
			{
				name:      "HTML element with XHTML default namespace",
				node:      htmlElem,
				namespace: "http://www.w3.org/1999/xhtml",
				expected:  true,
			},
			{
				name:      "HTML element with SVG namespace (not default)",
				node:      htmlElem,
				namespace: "http://www.w3.org/2000/svg",
				expected:  false,
			},
			{
				name:      "SVG element inheriting XHTML default namespace",
				node:      svgElem,
				namespace: "http://www.w3.org/1999/xhtml",
				expected:  true, // Should inherit default namespace from parent
			},
			{
				name:      "Text node inheriting default namespace",
				node:      textNode,
				namespace: "http://www.w3.org/1999/xhtml",
				expected:  true, // Should inherit from parent
			},
			{
				name:      "Empty namespace",
				node:      htmlElem,
				namespace: "",
				expected:  false,
			},
			{
				name:      "Unknown namespace",
				node:      htmlElem,
				namespace: "http://example.com/unknown",
				expected:  false,
			},
			{
				name:      "Document node checking default namespace",
				node:      doc,
				namespace: "http://www.w3.org/1999/xhtml",
				expected:  true, // Should delegate to document element
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := tt.node.IsDefaultNamespace(tt.namespace)
				if result != tt.expected {
					t.Errorf("IsDefaultNamespace(%q) = %t, expected %t", tt.namespace, result, tt.expected)
				}
			})
		}
	})
}

func TestNodeNamespaceEdgeCases(t *testing.T) {
	doc := NewDocument()

	t.Run("Document without document element", func(t *testing.T) {
		// Test document methods when there's no document element
		prefix := doc.LookupPrefix("http://www.w3.org/1999/xhtml")
		if prefix != "" {
			t.Errorf("Document without document element should return empty prefix, got %q", prefix)
		}

		namespace := doc.LookupNamespaceURI("html")
		if namespace != "" {
			t.Errorf("Document without document element should return empty namespace, got %q", namespace)
		}

		isDefault := doc.IsDefaultNamespace("http://www.w3.org/1999/xhtml")
		if isDefault {
			t.Errorf("Document without document element should return false for IsDefaultNamespace")
		}
	})

	t.Run("DocumentType and DocumentFragment nodes", func(t *testing.T) {
		// Create DocumentType and DocumentFragment nodes
		docType := NewDocumentType("html", "", "", doc)
		docFrag := NewDocumentFragment(doc)

		// All namespace methods should return empty/false for these node types
		if prefix := docType.LookupPrefix("http://www.w3.org/1999/xhtml"); prefix != "" {
			t.Errorf("DocumentType.LookupPrefix should return empty, got %q", prefix)
		}

		if namespace := docType.LookupNamespaceURI("html"); namespace != "" {
			t.Errorf("DocumentType.LookupNamespaceURI should return empty, got %q", namespace)
		}

		if isDefault := docType.IsDefaultNamespace("http://www.w3.org/1999/xhtml"); isDefault {
			t.Errorf("DocumentType.IsDefaultNamespace should return false")
		}

		if prefix := docFrag.LookupPrefix("http://www.w3.org/1999/xhtml"); prefix != "" {
			t.Errorf("DocumentFragment.LookupPrefix should return empty, got %q", prefix)
		}

		if namespace := docFrag.LookupNamespaceURI("html"); namespace != "" {
			t.Errorf("DocumentFragment.LookupNamespaceURI should return empty, got %q", namespace)
		}

		if isDefault := docFrag.IsDefaultNamespace("http://www.w3.org/1999/xhtml"); isDefault {
			t.Errorf("DocumentFragment.IsDefaultNamespace should return false")
		}
	})

	t.Run("Attribute nodes", func(t *testing.T) {
		// Create an element and attribute
		elem := NewElement("div", doc)
		attr := NewAttr("class", "test", doc)

		// Set the owner element (simulating attribute attachment)
		attr.ownerElement = elem

		// Attribute should delegate to its owner element
		// Since elem has no namespace declarations, results should be empty/false
		if prefix := attr.LookupPrefix("http://www.w3.org/1999/xhtml"); prefix != "" {
			t.Errorf("Attr.LookupPrefix should delegate to owner element, got %q", prefix)
		}

		if namespace := attr.LookupNamespaceURI("html"); namespace != "" {
			t.Errorf("Attr.LookupNamespaceURI should delegate to owner element, got %q", namespace)
		}

		if isDefault := attr.IsDefaultNamespace("http://www.w3.org/1999/xhtml"); isDefault {
			t.Errorf("Attr.IsDefaultNamespace should delegate to owner element")
		}
	})

	t.Run("Orphaned attribute", func(t *testing.T) {
		// Create an attribute without an owner element
		attr := NewAttr("id", "test", doc)

		// Should return empty/false since there's no owner element
		if prefix := attr.LookupPrefix("http://www.w3.org/1999/xhtml"); prefix != "" {
			t.Errorf("Orphaned Attr.LookupPrefix should return empty, got %q", prefix)
		}

		if namespace := attr.LookupNamespaceURI("html"); namespace != "" {
			t.Errorf("Orphaned Attr.LookupNamespaceURI should return empty, got %q", namespace)
		}

		if isDefault := attr.IsDefaultNamespace("http://www.w3.org/1999/xhtml"); isDefault {
			t.Errorf("Orphaned Attr.IsDefaultNamespace should return false")
		}
	})
}

func TestNodeNamespaceInheritance(t *testing.T) {
	doc := NewDocument()

	// Create a hierarchy: html > body > div > span > text
	htmlElem, _ := NewElementNS("http://www.w3.org/1999/xhtml", "html", doc)
	htmlElem.SetAttribute("xmlns", "http://www.w3.org/1999/xhtml")
	htmlElem.SetAttribute("xmlns:custom", "http://example.com/custom")
	doc.AppendChild(htmlElem)

	bodyElem := NewElement("body", doc)
	htmlElem.AppendChild(bodyElem)

	divElem := NewElement("div", doc)
	divElem.SetAttribute("xmlns:local", "http://example.com/local")
	bodyElem.AppendChild(divElem)

	spanElem := NewElement("span", doc)
	divElem.AppendChild(spanElem)

	textNode := NewText("Hello", doc)
	spanElem.AppendChild(textNode)

	t.Run("Namespace inheritance through hierarchy", func(t *testing.T) {
		// Test that namespace declarations are inherited down the tree
		tests := []struct {
			name     string
			node     Node
			prefix   string
			expected string
		}{
			{
				name:     "html element - default namespace",
				node:     htmlElem,
				prefix:   "",
				expected: "http://www.w3.org/1999/xhtml",
			},
			{
				name:     "body element - inherited default namespace",
				node:     bodyElem,
				prefix:   "",
				expected: "http://www.w3.org/1999/xhtml",
			},
			{
				name:     "div element - inherited default namespace",
				node:     divElem,
				prefix:   "",
				expected: "http://www.w3.org/1999/xhtml",
			},
			{
				name:     "span element - inherited default namespace",
				node:     spanElem,
				prefix:   "",
				expected: "http://www.w3.org/1999/xhtml",
			},
			{
				name:     "text node - inherited default namespace",
				node:     textNode,
				prefix:   "",
				expected: "http://www.w3.org/1999/xhtml",
			},
			{
				name:     "body element - inherited custom prefix",
				node:     bodyElem,
				prefix:   "custom",
				expected: "http://example.com/custom",
			},
			{
				name:     "span element - local prefix from div",
				node:     spanElem,
				prefix:   "local",
				expected: "http://example.com/local",
			},
			{
				name:     "text node - local prefix inherited from div",
				node:     textNode,
				prefix:   "local",
				expected: "http://example.com/local",
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := tt.node.LookupNamespaceURI(tt.prefix)
				if result != tt.expected {
					t.Errorf("LookupNamespaceURI(%q) = %q, expected %q", tt.prefix, result, tt.expected)
				}
			})
		}
	})

	t.Run("Prefix lookup through hierarchy", func(t *testing.T) {
		// Test LookupPrefix inheritance
		tests := []struct {
			name      string
			node      Node
			namespace string
			expected  string
		}{
			{
				name:      "text node looking up custom namespace",
				node:      textNode,
				namespace: "http://example.com/custom",
				expected:  "custom",
			},
			{
				name:      "span element looking up local namespace",
				node:      spanElem,
				namespace: "http://example.com/local",
				expected:  "local",
			},
			{
				name:      "body element looking up local namespace (not found)",
				node:      bodyElem,
				namespace: "http://example.com/local",
				expected:  "", // Should not find local prefix defined on div
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := tt.node.LookupPrefix(tt.namespace)
				if result != tt.expected {
					t.Errorf("LookupPrefix(%q) = %q, expected %q", tt.namespace, result, tt.expected)
				}
			})
		}
	})
}

func TestNamespaceBuiltInPrefixes(t *testing.T) {
	doc := NewDocument()
	elem := NewElement("test", doc)

	t.Run("Built-in xml prefix", func(t *testing.T) {
		namespace := elem.LookupNamespaceURI("xml")
		expected := "http://www.w3.org/XML/1998/namespace"
		if namespace != expected {
			t.Errorf("LookupNamespaceURI('xml') = %q, expected %q", namespace, expected)
		}
	})

	t.Run("Built-in xmlns prefix", func(t *testing.T) {
		namespace := elem.LookupNamespaceURI("xmlns")
		expected := "http://www.w3.org/2000/xmlns/"
		if namespace != expected {
			t.Errorf("LookupNamespaceURI('xmlns') = %q, expected %q", namespace, expected)
		}
	})

	t.Run("Built-in prefixes work on any element", func(t *testing.T) {
		textNode := NewText("test", doc)
		elem.AppendChild(textNode)

		// Even text nodes should resolve built-in prefixes
		xmlNS := textNode.LookupNamespaceURI("xml")
		expectedXML := "http://www.w3.org/XML/1998/namespace"
		if xmlNS != expectedXML {
			t.Errorf("Text node LookupNamespaceURI('xml') = %q, expected %q", xmlNS, expectedXML)
		}

		xmlnsNS := textNode.LookupNamespaceURI("xmlns")
		expectedXMLNS := "http://www.w3.org/2000/xmlns/"
		if xmlnsNS != expectedXMLNS {
			t.Errorf("Text node LookupNamespaceURI('xmlns') = %q, expected %q", xmlnsNS, expectedXMLNS)
		}
	})
}

func TestNamespaceEmptyValues(t *testing.T) {
	doc := NewDocument()
	elem := NewElement("test", doc)
	elem.SetAttribute("xmlns:empty", "") // Empty namespace URI

	t.Run("Empty namespace URI in xmlns attribute", func(t *testing.T) {
		// Per spec, empty string namespace URI should be returned as empty string
		namespace := elem.LookupNamespaceURI("empty")
		if namespace != "" {
			t.Errorf("LookupNamespaceURI('empty') = %q, expected empty string", namespace)
		}
	})

	t.Run("Empty default namespace", func(t *testing.T) {
		elem.SetAttribute("xmlns", "") // Remove default namespace

		namespace := elem.LookupNamespaceURI("")
		if namespace != "" {
			t.Errorf("LookupNamespaceURI('') with empty xmlns = %q, expected empty string", namespace)
		}

		isDefault := elem.IsDefaultNamespace("")
		if !isDefault {
			t.Errorf("IsDefaultNamespace('') with empty xmlns should return true")
		}
	})
}
