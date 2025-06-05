package dom

import (
	"testing"
)

func TestDOMImplementation_CreateDocumentType(t *testing.T) {
	// Create a document for the DOMImplementation
	doc := NewDocument()
	impl := doc.Implementation()

	t.Run("Valid qualified names", func(t *testing.T) {
		testCases := []struct {
			qualifiedName string
			publicID      string
			systemID      string
		}{
			{"html", "", ""},
			{"xml", "", ""},
			{"svg:svg", "", ""},
			{"my-element", "", ""},
			{"_test", "", ""},
			{"prefix:localName", "public", "system"},
		}

		for _, tc := range testCases {
			t.Run(tc.qualifiedName, func(t *testing.T) {
				doctype, err := impl.CreateDocumentType(tc.qualifiedName, tc.publicID, tc.systemID)
				if err != nil {
					t.Errorf("Expected no error for valid qualified name %q, got: %v", tc.qualifiedName, err)
					return
				}

				if doctype == nil {
					t.Errorf("Expected non-nil DocumentType")
					return
				}

				if doctype.Name() != tc.qualifiedName {
					t.Errorf("Expected name %q, got %q", tc.qualifiedName, doctype.Name())
				}

				if doctype.PublicID() != tc.publicID {
					t.Errorf("Expected publicId %q, got %q", tc.publicID, doctype.PublicID())
				}

				if doctype.SystemID() != tc.systemID {
					t.Errorf("Expected systemId %q, got %q", tc.systemID, doctype.SystemID())
				}

				// Verify node document is set to the associated document
				if doctype.OwnerDocument() != doc {
					t.Errorf("Expected owner document to be the associated document")
				}
			})
		}
	})

	t.Run("Invalid qualified names - InvalidCharacterError", func(t *testing.T) {
		testCases := []string{
			"",           // Empty string
			"123invalid", // Starts with number
			"inv@lid",    // Contains invalid character
			"inv lid",    // Contains space
		}

		for _, qualifiedName := range testCases {
			t.Run(qualifiedName, func(t *testing.T) {
				_, err := impl.CreateDocumentType(qualifiedName, "", "")
				if err == nil {
					t.Errorf("Expected InvalidCharacterError for qualified name %q", qualifiedName)
					return
				}

				if domErr, ok := err.(*DOMException); ok {
					if domErr.Code() != INVALID_CHARACTER_ERR {
						t.Errorf("Expected InvalidCharacterError, got error code %d", domErr.Code())
					}
				} else {
					t.Errorf("Expected DOMException with InvalidCharacterError, got %T", err)
				}
			})
		}
	})

	t.Run("Invalid qualified names - NamespaceError", func(t *testing.T) {
		testCases := []string{
			"pre:fix:local", // More than one colon
			":",             // Empty prefix and local name
			"prefix:",       // Empty local name
		}

		for _, qualifiedName := range testCases {
			t.Run(qualifiedName, func(t *testing.T) {
				_, err := impl.CreateDocumentType(qualifiedName, "", "")
				if err == nil {
					t.Errorf("Expected NamespaceError for qualified name %q", qualifiedName)
					return
				}

				if domErr, ok := err.(*DOMException); ok {
					if domErr.Code() != NAMESPACE_ERR && domErr.Code() != INVALID_CHARACTER_ERR {
						t.Errorf("Expected NamespaceError or InvalidCharacterError, got error code %d", domErr.Code())
					}
				} else {
					t.Errorf("Expected DOMException, got %T", err)
				}
			})
		}
	})
}

func TestDOMImplementation_CreateDocument(t *testing.T) {
	// Create a document for the DOMImplementation
	doc := NewDocument()
	impl := doc.Implementation()

	t.Run("Create XML document with empty qualified name", func(t *testing.T) {
		xmlDoc, err := impl.CreateDocument("", "", nil)
		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
			return
		}

		if xmlDoc == nil {
			t.Errorf("Expected non-nil Document")
			return
		}

		// Verify it's an XML document
		if xmlDoc.ContentType() != "application/xml" {
			t.Errorf("Expected content type 'application/xml', got %q", xmlDoc.ContentType())
		}

		// Should have no document element
		if xmlDoc.DocumentElement() != nil {
			t.Errorf("Expected no document element for empty qualified name")
		}
	})

	t.Run("Create XML document with qualified name", func(t *testing.T) {
		xmlDoc, err := impl.CreateDocument("", "root", nil)
		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
			return
		}

		// Should have document element
		if xmlDoc.DocumentElement() == nil {
			t.Errorf("Expected document element")
			return
		}

		if xmlDoc.DocumentElement().TagName() != "root" {
			t.Errorf("Expected document element tag name 'root', got %q", xmlDoc.DocumentElement().TagName())
		}
	})

	t.Run("Create document with namespace - HTML", func(t *testing.T) {
		xmlDoc, err := impl.CreateDocument(HTMLNamespace, "html", nil)
		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
			return
		}

		// HTML namespace should set content type to application/xhtml+xml
		if xmlDoc.ContentType() != "application/xhtml+xml" {
			t.Errorf("Expected content type 'application/xhtml+xml', got %q", xmlDoc.ContentType())
		}
	})

	t.Run("Create document with namespace - SVG", func(t *testing.T) {
		xmlDoc, err := impl.CreateDocument(SVGNamespace, "svg", nil)
		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
			return
		}

		// SVG namespace should set content type to image/svg+xml
		if xmlDoc.ContentType() != "image/svg+xml" {
			t.Errorf("Expected content type 'image/svg+xml', got %q", xmlDoc.ContentType())
		}
	})

	t.Run("Create document with namespace - Other", func(t *testing.T) {
		xmlDoc, err := impl.CreateDocument("http://example.com/ns", "custom", nil)
		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
			return
		}

		// Other namespaces should set content type to application/xml
		if xmlDoc.ContentType() != "application/xml" {
			t.Errorf("Expected content type 'application/xml', got %q", xmlDoc.ContentType())
		}
	})

	t.Run("Create document with doctype", func(t *testing.T) {
		// Create a doctype first
		doctype, err := impl.CreateDocumentType("root", "", "")
		if err != nil {
			t.Errorf("Failed to create doctype: %v", err)
			return
		}

		xmlDoc, err := impl.CreateDocument("", "root", doctype)
		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
			return
		}

		// Verify doctype is attached
		if xmlDoc.Doctype() == nil {
			t.Errorf("Expected doctype to be attached")
		} else if xmlDoc.Doctype().Name() != "root" {
			t.Errorf("Expected doctype name 'root', got %q", xmlDoc.Doctype().Name())
		}
	})

	t.Run("Invalid qualified name should throw same exception as createElementNS", func(t *testing.T) {
		// Test with invalid qualified name
		_, err := impl.CreateDocument("", "inv@lid", nil)
		if err == nil {
			t.Errorf("Expected error for invalid qualified name")
			return
		}

		// Should be same type of error as createElementNS would throw
		// Accept both DOMException and the legacy error types
		switch e := err.(type) {
		case *DOMException:
			if e.Code() != INVALID_CHARACTER_ERR && e.Code() != NAMESPACE_ERR {
				t.Errorf("Expected InvalidCharacterError or NamespaceError, got error code %d", e.Code())
			}
		case *InvalidCharacterError, *NamespaceError:
			// These are valid error types from createElementNS
		default:
			t.Errorf("Expected DOMException, InvalidCharacterError, or NamespaceError, got %T", err)
		}
	})
}

func TestDOMImplementation_CreateHTMLDocument(t *testing.T) {
	// Create a document for the DOMImplementation
	doc := NewDocument()
	impl := doc.Implementation()

	t.Run("Create HTML document without title", func(t *testing.T) {
		htmlDoc := impl.CreateHTMLDocument()

		if htmlDoc == nil {
			t.Errorf("Expected non-nil Document")
			return
		}

		// Verify content type
		if htmlDoc.ContentType() != "text/html" {
			t.Errorf("Expected content type 'text/html', got %q", htmlDoc.ContentType())
		}

		// Verify doctype
		doctype := htmlDoc.Doctype()
		if doctype == nil {
			t.Errorf("Expected DOCTYPE")
		} else if doctype.Name() != "html" {
			t.Errorf("Expected DOCTYPE name 'html', got %q", doctype.Name())
		}

		// Verify document element
		documentElement := htmlDoc.DocumentElement()
		if documentElement == nil {
			t.Errorf("Expected document element")
		} else if documentElement.TagName() != "HTML" { // Expect uppercase
			t.Errorf("Expected document element tag name 'HTML', got %q", documentElement.TagName())
		}

		// Verify head element
		head := htmlDoc.Head()
		if head == nil {
			t.Errorf("Expected head element")
		} else if head.TagName() != "HEAD" { // Expect uppercase
			t.Errorf("Expected head element tag name 'HEAD', got %q", head.TagName())
		}

		// Verify body element
		body := htmlDoc.Body()
		if body == nil {
			t.Errorf("Expected body element")
		} else if body.TagName() != "BODY" { // Expect uppercase
			t.Errorf("Expected body element tag name 'BODY', got %q", body.TagName())
		}

		// Should have no title element
		titleElements := head.GetElementsByTagName("title")
		if titleElements.Length() != 0 {
			t.Errorf("Expected no title element when title not provided")
		}
	})

	t.Run("Create HTML document with title", func(t *testing.T) {
		title := "Test Document"
		htmlDoc := impl.CreateHTMLDocument(title)

		if htmlDoc == nil {
			t.Errorf("Expected non-nil Document")
			return
		}

		// Verify head element exists
		head := htmlDoc.Head()
		if head == nil {
			t.Errorf("Expected head element")
			return
		}

		// Verify title element
		titleElements := head.GetElementsByTagName("title")
		if titleElements.Length() != 1 {
			t.Errorf("Expected exactly one title element, got %d", titleElements.Length())
			return
		}

		titleElement := titleElements.Item(0)
		if titleElement != nil {
			if titleElement.TextContent() != title {
				t.Errorf("Expected title text content %q, got %q", title, titleElement.TextContent())
			}
		} else {
			t.Errorf("Expected title element but got nil")
		}
	})

	t.Run("Create HTML document with empty title", func(t *testing.T) {
		title := ""
		htmlDoc := impl.CreateHTMLDocument(title)

		head := htmlDoc.Head()
		if head == nil {
			t.Errorf("Expected head element")
			return
		}

		// Should have title element even with empty string
		titleElements := head.GetElementsByTagName("title")
		if titleElements.Length() != 1 {
			t.Errorf("Expected exactly one title element, got %d", titleElements.Length())
			return
		}

		titleElement := titleElements.Item(0)
		if titleElement != nil {
			if titleElement.TextContent() != "" {
				t.Errorf("Expected empty title text content, got %q", titleElement.TextContent())
			}
		} else {
			t.Errorf("Expected title element but got nil")
		}
	})

	t.Run("Verify HTML namespace", func(t *testing.T) {
		htmlDoc := impl.CreateHTMLDocument("Test")

		// Verify elements have HTML namespace
		documentElement := htmlDoc.DocumentElement()
		if documentElement.NamespaceURI() != HTMLNamespace {
			t.Errorf("Expected HTML namespace %q, got %q", HTMLNamespace, documentElement.NamespaceURI())
		}

		head := htmlDoc.Head()
		if head.NamespaceURI() != HTMLNamespace {
			t.Errorf("Expected HTML namespace %q, got %q", HTMLNamespace, head.NamespaceURI())
		}

		body := htmlDoc.Body()
		if body.NamespaceURI() != HTMLNamespace {
			t.Errorf("Expected HTML namespace %q, got %q", HTMLNamespace, body.NamespaceURI())
		}
	})
}

func TestDOMImplementation_HasFeature(t *testing.T) {
	// Create a document for the DOMImplementation
	doc := NewDocument()
	impl := doc.Implementation()

	t.Run("Always returns true", func(t *testing.T) {
		testCases := []struct {
			feature string
			version string
		}{
			{"", ""},
			{"Core", "1.0"},
			{"XML", "2.0"},
			{"HTML", ""},
			{"invalid", "invalid"},
			{"anything", "anything"},
		}

		for _, tc := range testCases {
			t.Run(tc.feature+"_"+tc.version, func(t *testing.T) {
				result := impl.HasFeature(tc.feature, tc.version)
				if !result {
					t.Errorf("Expected HasFeature to always return true, got false for feature=%q version=%q", tc.feature, tc.version)
				}
			})
		}
	})
}

func TestDOMImplementation_DocumentAssociation(t *testing.T) {
	t.Run("DOMImplementation is associated with document", func(t *testing.T) {
		doc := NewDocument()
		impl := doc.Implementation()

		if impl == nil {
			t.Errorf("Expected non-nil DOMImplementation")
			return
		}

		// The implementation should be associated with the document
		// We can't directly test the internal association, but we can test
		// that createDocumentType sets the correct owner document
		doctype, err := impl.CreateDocumentType("test", "", "")
		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
			return
		}

		if doctype.OwnerDocument() != doc {
			t.Errorf("Expected doctype owner document to be the associated document")
		}
	})
}

// Benchmark tests
func BenchmarkCreateDocumentType(b *testing.B) {
	doc := NewDocument()
	impl := doc.Implementation()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := impl.CreateDocumentType("html", "", "")
		if err != nil {
			b.Errorf("Unexpected error: %v", err)
		}
	}
}

func BenchmarkCreateDocument(b *testing.B) {
	doc := NewDocument()
	impl := doc.Implementation()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := impl.CreateDocument("", "root", nil)
		if err != nil {
			b.Errorf("Unexpected error: %v", err)
		}
	}
}

func BenchmarkCreateHTMLDocument(b *testing.B) {
	doc := NewDocument()
	impl := doc.Implementation()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = impl.CreateHTMLDocument("Test Document")
	}
}
