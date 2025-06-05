package dom

import (
	"strings"
	"testing"
	"unicode/utf8"
)

// TestDocumentTypeSpecCompliance tests compliance with WHATWG DOM Section 4.6
// Interface DocumentType specification
func TestDocumentTypeSpecCompliance(t *testing.T) {
	doc := NewDocument()

	t.Run("WHATWG DOM Section 4.6 - Interface DocumentType", func(t *testing.T) {
		t.Run("DocumentType nodes are simply known as doctypes", func(t *testing.T) {
			// Verify that DocumentType nodes are recognized as doctypes
			doctype := NewDocumentType("html", "", "", doc)

			// Should be a DocumentType node
			if doctype.NodeType() != DocumentTypeNode {
				t.Errorf("Expected DocumentType node type %d, got %d", DocumentTypeNode, doctype.NodeType())
			}

			// Node name should be the same as the doctype name
			if doctype.NodeName() != "html" {
				t.Errorf("Expected node name to match doctype name 'html', got %q", doctype.NodeName())
			}
		})

		t.Run("Doctypes have an associated name, public ID, and system ID", func(t *testing.T) {
			name := "html"
			publicID := "-//W3C//DTD HTML 4.01//EN"
			systemID := "http://www.w3.org/TR/html4/strict.dtd"

			doctype := NewDocumentType(name, publicID, systemID, doc)

			// Verify all three properties are associated with the doctype
			if doctype.Name() != name {
				t.Errorf("Expected name %q, got %q", name, doctype.Name())
			}

			if doctype.PublicID() != publicID {
				t.Errorf("Expected publicID %q, got %q", publicID, doctype.PublicID())
			}

			if doctype.SystemID() != systemID {
				t.Errorf("Expected systemID %q, got %q", systemID, doctype.SystemID())
			}
		})

		t.Run("When a doctype is created, its name is always given", func(t *testing.T) {
			// Test that name is required and preserved
			testNames := []string{
				"html",
				"xml",
				"svg",
				"custom-doctype",
				"_underscore",
				"name-with-dashes",
				"NameWithMixedCase",
			}

			for _, name := range testNames {
				t.Run("name_"+name, func(t *testing.T) {
					doctype := NewDocumentType(name, "", "", doc)

					if doctype.Name() != name {
						t.Errorf("Expected name %q to be preserved, got %q", name, doctype.Name())
					}

					// Verify name is accessible via the getter
					if doctype.Name() == "" {
						t.Error("Name should never be empty when given")
					}
				})
			}
		})

		t.Run("Unless explicitly given when a doctype is created, its public ID and system ID are the empty string", func(t *testing.T) {
			// Test default empty strings
			doctype1 := NewDocumentType("html", "", "", doc)

			if doctype1.PublicID() != "" {
				t.Errorf("Expected default publicID to be empty string, got %q", doctype1.PublicID())
			}

			if doctype1.SystemID() != "" {
				t.Errorf("Expected default systemID to be empty string, got %q", doctype1.SystemID())
			}

			// Test explicitly set empty strings
			doctype2 := NewDocumentType("xml", "", "", doc)

			if doctype2.PublicID() != "" {
				t.Errorf("Expected explicitly empty publicID to be empty string, got %q", doctype2.PublicID())
			}

			if doctype2.SystemID() != "" {
				t.Errorf("Expected explicitly empty systemID to be empty string, got %q", doctype2.SystemID())
			}

			// Test with non-empty values to ensure they're preserved
			doctype3 := NewDocumentType("svg", "public", "system", doc)

			if doctype3.PublicID() != "public" {
				t.Errorf("Expected publicID 'public', got %q", doctype3.PublicID())
			}

			if doctype3.SystemID() != "system" {
				t.Errorf("Expected systemID 'system', got %q", doctype3.SystemID())
			}
		})
	})

	t.Run("WHATWG DOM Section 4.6 - Getter methods specification", func(t *testing.T) {
		t.Run("The name getter steps are to return this's name", func(t *testing.T) {
			testCases := []struct {
				name        string
				description string
			}{
				{"html", "HTML5 doctype"},
				{"xml", "XML doctype"},
				{"svg:svg", "namespaced doctype"},
				{"custom", "custom doctype"},
				{"123", "numeric name"},
				{"", "empty name"},
				{"a", "single character"},
				{"very-long-doctype-name-with-many-characters", "long name"},
			}

			for _, tc := range testCases {
				t.Run(tc.description, func(t *testing.T) {
					doctype := NewDocumentType(tc.name, "public", "system", doc)

					// The name getter should return exactly the name that was set
					if doctype.Name() != tc.name {
						t.Errorf("Name getter should return %q, got %q", tc.name, doctype.Name())
					}

					// Multiple calls should return the same value
					name1 := doctype.Name()
					name2 := doctype.Name()
					if name1 != name2 {
						t.Errorf("Name getter should be consistent: %q != %q", name1, name2)
					}
				})
			}
		})

		t.Run("The publicId getter steps are to return this's public ID", func(t *testing.T) {
			testCases := []struct {
				publicID    string
				description string
			}{
				{"", "empty public ID"},
				{"-//W3C//DTD HTML 4.01//EN", "HTML 4.01 public ID"},
				{"-//W3C//DTD XHTML 1.0 Strict//EN", "XHTML 1.0 public ID"},
				{"public", "simple public ID"},
				{"//EXAMPLE//DTD Example//EN", "custom public ID"},
				{"a", "single character"},
				{strings.Repeat("x", 1000), "very long public ID"},
			}

			for _, tc := range testCases {
				t.Run(tc.description, func(t *testing.T) {
					doctype := NewDocumentType("test", tc.publicID, "system", doc)

					// The publicId getter should return exactly the public ID that was set
					if doctype.PublicID() != tc.publicID {
						t.Errorf("PublicID getter should return %q, got %q", tc.publicID, doctype.PublicID())
					}

					// Multiple calls should return the same value
					publicID1 := doctype.PublicID()
					publicID2 := doctype.PublicID()
					if publicID1 != publicID2 {
						t.Errorf("PublicID getter should be consistent: %q != %q", publicID1, publicID2)
					}
				})
			}
		})

		t.Run("The systemId getter steps are to return this's system ID", func(t *testing.T) {
			testCases := []struct {
				systemID    string
				description string
			}{
				{"", "empty system ID"},
				{"http://www.w3.org/TR/html4/strict.dtd", "HTML 4.01 system ID"},
				{"http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd", "XHTML 1.0 system ID"},
				{"system", "simple system ID"},
				{"file:///path/to/dtd", "file URL system ID"},
				{"https://example.com/custom.dtd", "HTTPS URL system ID"},
				{"a", "single character"},
				{strings.Repeat("y", 1000), "very long system ID"},
			}

			for _, tc := range testCases {
				t.Run(tc.description, func(t *testing.T) {
					doctype := NewDocumentType("test", "public", tc.systemID, doc)

					// The systemId getter should return exactly the system ID that was set
					if doctype.SystemID() != tc.systemID {
						t.Errorf("SystemID getter should return %q, got %q", tc.systemID, doctype.SystemID())
					}

					// Multiple calls should return the same value
					systemID1 := doctype.SystemID()
					systemID2 := doctype.SystemID()
					if systemID1 != systemID2 {
						t.Errorf("SystemID getter should be consistent: %q != %q", systemID1, systemID2)
					}
				})
			}
		})
	})

	t.Run("Properties are read-only", func(t *testing.T) {
		// DocumentType properties should be immutable after creation
		doctype := NewDocumentType("original", "orig-public", "orig-system", doc)

		// Verify initial values
		if doctype.Name() != "original" {
			t.Errorf("Expected initial name 'original', got %q", doctype.Name())
		}
		if doctype.PublicID() != "orig-public" {
			t.Errorf("Expected initial publicID 'orig-public', got %q", doctype.PublicID())
		}
		if doctype.SystemID() != "orig-system" {
			t.Errorf("Expected initial systemID 'orig-system', got %q", doctype.SystemID())
		}

		// The spec doesn't define setters, so properties should remain unchanged
		// We can't test setter absence directly, but we can verify they don't change
		// through any other operations

		// Verify values remain the same after various operations
		_ = doctype.CloneNode(false)
		_ = doctype.NodeName()
		_ = doctype.NodeType()

		if doctype.Name() != "original" {
			t.Errorf("Name should remain unchanged: expected 'original', got %q", doctype.Name())
		}
		if doctype.PublicID() != "orig-public" {
			t.Errorf("PublicID should remain unchanged: expected 'orig-public', got %q", doctype.PublicID())
		}
		if doctype.SystemID() != "orig-system" {
			t.Errorf("SystemID should remain unchanged: expected 'orig-system', got %q", doctype.SystemID())
		}
	})
}

func TestDocumentTypeEdgeCases(t *testing.T) {
	doc := NewDocument()

	t.Run("Unicode characters in properties", func(t *testing.T) {
		// Test various Unicode characters in name, publicID, and systemID
		testCases := []struct {
			name     string
			publicID string
			systemID string
			desc     string
		}{
			{"html", "Ñoñó", "sýstém", "Latin characters with diacritics"},
			{"тест", "публик", "систем", "Cyrillic characters"},
			{"测试", "公共", "系统", "Chinese characters"},
			{"🌟", "⭐", "✨", "Emoji characters"},
			{"", "👨‍💻", "🚀", "Empty name with emoji IDs"},
		}

		for _, tc := range testCases {
			t.Run(tc.desc, func(t *testing.T) {
				doctype := NewDocumentType(tc.name, tc.publicID, tc.systemID, doc)

				// Verify Unicode characters are preserved correctly
				if doctype.Name() != tc.name {
					t.Errorf("Unicode name not preserved: expected %q, got %q", tc.name, doctype.Name())
				}
				if doctype.PublicID() != tc.publicID {
					t.Errorf("Unicode publicID not preserved: expected %q, got %q", tc.publicID, doctype.PublicID())
				}
				if doctype.SystemID() != tc.systemID {
					t.Errorf("Unicode systemID not preserved: expected %q, got %q", tc.systemID, doctype.SystemID())
				}

				// Verify strings are valid UTF-8
				if !utf8.ValidString(doctype.Name()) {
					t.Error("Name is not valid UTF-8")
				}
				if !utf8.ValidString(doctype.PublicID()) {
					t.Error("PublicID is not valid UTF-8")
				}
				if !utf8.ValidString(doctype.SystemID()) {
					t.Error("SystemID is not valid UTF-8")
				}
			})
		}
	})

	t.Run("Very long strings", func(t *testing.T) {
		// Test with very long strings to ensure no truncation
		longName := strings.Repeat("name", 1000)     // 4000 chars
		longPublic := strings.Repeat("public", 1000) // 6000 chars
		longSystem := strings.Repeat("system", 1000) // 6000 chars

		doctype := NewDocumentType(longName, longPublic, longSystem, doc)

		if doctype.Name() != longName {
			t.Errorf("Long name truncated: expected length %d, got %d", len(longName), len(doctype.Name()))
		}
		if doctype.PublicID() != longPublic {
			t.Errorf("Long publicID truncated: expected length %d, got %d", len(longPublic), len(doctype.PublicID()))
		}
		if doctype.SystemID() != longSystem {
			t.Errorf("Long systemID truncated: expected length %d, got %d", len(longSystem), len(doctype.SystemID()))
		}
	})

	t.Run("Special characters and whitespace", func(t *testing.T) {
		// Test various special characters and whitespace handling
		testCases := []struct {
			name     string
			publicID string
			systemID string
			desc     string
		}{
			{" ", " ", " ", "single space"},
			{"\t", "\t", "\t", "tab character"},
			{"\n", "\n", "\n", "newline character"},
			{"\r\n", "\r\n", "\r\n", "CRLF sequence"},
			{"  leading", "  leading", "  leading", "leading spaces"},
			{"trailing  ", "trailing  ", "trailing  ", "trailing spaces"},
			{" both ", " both ", " both ", "both leading and trailing"},
			{"inner spaces", "inner spaces", "inner spaces", "internal spaces"},
			{"!@#$%^&*()", "!@#$%^&*()", "!@#$%^&*()", "special symbols"},
			{"\"quotes\"", "\"quotes\"", "\"quotes\"", "double quotes"},
			{"'apostrophes'", "'apostrophes'", "'apostrophes'", "single quotes"},
		}

		for _, tc := range testCases {
			t.Run(tc.desc, func(t *testing.T) {
				doctype := NewDocumentType(tc.name, tc.publicID, tc.systemID, doc)

				// Verify all characters are preserved exactly as given
				if doctype.Name() != tc.name {
					t.Errorf("Name with %s not preserved: expected %q, got %q", tc.desc, tc.name, doctype.Name())
				}
				if doctype.PublicID() != tc.publicID {
					t.Errorf("PublicID with %s not preserved: expected %q, got %q", tc.desc, tc.publicID, doctype.PublicID())
				}
				if doctype.SystemID() != tc.systemID {
					t.Errorf("SystemID with %s not preserved: expected %q, got %q", tc.desc, tc.systemID, doctype.SystemID())
				}
			})
		}
	})

	t.Run("Null bytes and control characters", func(t *testing.T) {
		// Test handling of null bytes and control characters
		testCases := []struct {
			name     string
			publicID string
			systemID string
			desc     string
		}{
			{"\x00", "\x00", "\x00", "null bytes"},
			{"\x01\x02\x03", "\x01\x02\x03", "\x01\x02\x03", "control characters"},
			{"before\x00after", "before\x00after", "before\x00after", "null byte in middle"},
		}

		for _, tc := range testCases {
			t.Run(tc.desc, func(t *testing.T) {
				doctype := NewDocumentType(tc.name, tc.publicID, tc.systemID, doc)

				// Verify all bytes are preserved exactly
				if doctype.Name() != tc.name {
					t.Errorf("Name with %s not preserved correctly", tc.desc)
				}
				if doctype.PublicID() != tc.publicID {
					t.Errorf("PublicID with %s not preserved correctly", tc.desc)
				}
				if doctype.SystemID() != tc.systemID {
					t.Errorf("SystemID with %s not preserved correctly", tc.desc)
				}
			})
		}
	})
}

func TestDocumentTypeIntegrationWithDOM(t *testing.T) {
	doc := NewDocument()

	t.Run("DocumentType as child of Document", func(t *testing.T) {
		// Create a DocumentType and add it to a document
		doctype := NewDocumentType("html", "-//W3C//DTD HTML 4.01//EN", "http://www.w3.org/TR/html4/strict.dtd", doc)

		// Add to document (AppendChild returns the added node, not an error)
		result := doc.AppendChild(doctype)
		if result == nil {
			t.Error("AppendChild returned nil")
		}

		// Verify it's accessible via Document.Doctype()
		retrievedDoctype := doc.Doctype()
		if retrievedDoctype == nil {
			t.Error("Document.Doctype() returned nil after adding DocumentType")
		} else {
			if retrievedDoctype.Name() != "html" {
				t.Errorf("Retrieved doctype name: expected 'html', got %q", retrievedDoctype.Name())
			}
		}

		// Verify owner document relationship
		if doctype.OwnerDocument() != doc {
			t.Error("DocumentType owner document should be the document it was created with")
		}

		// Verify parent-child relationship
		if doctype.ParentNode() != doc {
			t.Error("DocumentType parent should be the document after being appended")
		}
	})

	t.Run("DocumentType created via DOMImplementation", func(t *testing.T) {
		// Test DocumentType created through DOMImplementation interface
		impl := doc.Implementation()
		doctype, err := impl.CreateDocumentType("svg", "-//W3C//DTD SVG 1.1//EN", "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd")
		if err != nil {
			t.Errorf("Failed to create DocumentType via DOMImplementation: %v", err)
		}

		// Verify all properties are correctly set
		if doctype.Name() != "svg" {
			t.Errorf("Expected name 'svg', got %q", doctype.Name())
		}
		if doctype.PublicID() != "-//W3C//DTD SVG 1.1//EN" {
			t.Errorf("Expected publicID '-//W3C//DTD SVG 1.1//EN', got %q", doctype.PublicID())
		}
		if doctype.SystemID() != "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd" {
			t.Errorf("Expected systemID 'http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd', got %q", doctype.SystemID())
		}

		// Verify owner document is set correctly
		if doctype.OwnerDocument() != doc {
			t.Error("DocumentType created via DOMImplementation should have correct owner document")
		}
	})

	t.Run("DocumentType cloning preserves all properties", func(t *testing.T) {
		// Test that cloning preserves all DocumentType properties correctly
		original := NewDocumentType("mathml", "-//W3C//DTD MathML 2.0//EN", "http://www.w3.org/Math/DTD/mathml2/mathml2.dtd", doc)

		// Clone the DocumentType
		cloned := original.CloneNode(false)
		clonedDT, ok := cloned.(*DocumentType)
		if !ok {
			t.Fatal("Cloned node is not a DocumentType")
		}

		// Verify all properties are preserved
		if clonedDT.Name() != original.Name() {
			t.Errorf("Cloned name mismatch: expected %q, got %q", original.Name(), clonedDT.Name())
		}
		if clonedDT.PublicID() != original.PublicID() {
			t.Errorf("Cloned publicID mismatch: expected %q, got %q", original.PublicID(), clonedDT.PublicID())
		}
		if clonedDT.SystemID() != original.SystemID() {
			t.Errorf("Cloned systemID mismatch: expected %q, got %q", original.SystemID(), clonedDT.SystemID())
		}

		// Verify it's a different instance
		if cloned == original {
			t.Error("Cloned DocumentType should be a different instance")
		}

		// Verify owner document is preserved
		if clonedDT.OwnerDocument() != original.OwnerDocument() {
			t.Error("Cloned DocumentType should have same owner document")
		}
	})
}

func BenchmarkDocumentTypeOperations(b *testing.B) {
	doc := NewDocument()

	b.Run("NewDocumentType", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = NewDocumentType("html", "-//W3C//DTD HTML 4.01//EN", "http://www.w3.org/TR/html4/strict.dtd", doc)
		}
	})

	b.Run("Name getter", func(b *testing.B) {
		doctype := NewDocumentType("html", "-//W3C//DTD HTML 4.01//EN", "http://www.w3.org/TR/html4/strict.dtd", doc)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = doctype.Name()
		}
	})

	b.Run("PublicID getter", func(b *testing.B) {
		doctype := NewDocumentType("html", "-//W3C//DTD HTML 4.01//EN", "http://www.w3.org/TR/html4/strict.dtd", doc)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = doctype.PublicID()
		}
	})

	b.Run("SystemID getter", func(b *testing.B) {
		doctype := NewDocumentType("html", "-//W3C//DTD HTML 4.01//EN", "http://www.w3.org/TR/html4/strict.dtd", doc)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = doctype.SystemID()
		}
	})

	b.Run("CloneNode", func(b *testing.B) {
		doctype := NewDocumentType("html", "-//W3C//DTD HTML 4.01//EN", "http://www.w3.org/TR/html4/strict.dtd", doc)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = doctype.CloneNode(false)
		}
	})
}
