package dom

import (
	"testing"
)

func TestValidateQualifiedName(t *testing.T) {
	tests := []struct {
		name          string
		qualifiedName string
		expectError   bool
		errorType     string
	}{
		// Valid names
		{
			name:          "simple element name",
			qualifiedName: "div",
			expectError:   false,
		},
		{
			name:          "element with prefix",
			qualifiedName: "xml:lang",
			expectError:   false,
		},
		{
			name:          "element with underscores",
			qualifiedName: "my_element",
			expectError:   false,
		},
		{
			name:          "element with numbers",
			qualifiedName: "element123",
			expectError:   false,
		},
		{
			name:          "element with hyphens",
			qualifiedName: "my-element",
			expectError:   false,
		},
		{
			name:          "complex namespace prefix",
			qualifiedName: "ns:complex-element_123",
			expectError:   false,
		},

		// Invalid names - empty
		{
			name:          "empty name",
			qualifiedName: "",
			expectError:   true,
			errorType:     "InvalidCharacterError",
		},

		// Invalid names - multiple colons
		{
			name:          "multiple colons",
			qualifiedName: "ns:sub:element",
			expectError:   true,
			errorType:     "NamespaceError",
		},

		// Invalid names - empty prefix
		{
			name:          "empty prefix",
			qualifiedName: ":element",
			expectError:   true,
			errorType:     "NamespaceError",
		},

		// Invalid names - empty local name
		{
			name:          "empty local name",
			qualifiedName: "prefix:",
			expectError:   true,
			errorType:     "NamespaceError",
		},

		// Invalid names - invalid characters
		{
			name:          "starts with number",
			qualifiedName: "123element",
			expectError:   true,
			errorType:     "InvalidCharacterError",
		},
		{
			name:          "contains space",
			qualifiedName: "my element",
			expectError:   true,
			errorType:     "InvalidCharacterError",
		},
		{
			name:          "contains invalid symbols",
			qualifiedName: "element!",
			expectError:   true,
			errorType:     "InvalidCharacterError",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateQualifiedName(tt.qualifiedName)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error for '%s', but got none", tt.qualifiedName)
					return
				}

				// Check error type
				switch tt.errorType {
				case "InvalidCharacterError":
					if _, ok := err.(*InvalidCharacterError); !ok {
						t.Errorf("Expected InvalidCharacterError, got %T: %v", err, err)
					}
				case "NamespaceError":
					if _, ok := err.(*NamespaceError); !ok {
						t.Errorf("Expected NamespaceError, got %T: %v", err, err)
					}
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error for '%s': %v", tt.qualifiedName, err)
				}
			}
		})
	}
}

func TestValidateAndExtract(t *testing.T) {
	tests := []struct {
		name           string
		namespace      string
		qualifiedName  string
		expectError    bool
		errorType      string
		expectedNS     string
		expectedPrefix string
		expectedLocal  string
	}{
		// Valid combinations
		{
			name:           "simple name without namespace",
			namespace:      "",
			qualifiedName:  "div",
			expectError:    false,
			expectedNS:     "",
			expectedPrefix: "",
			expectedLocal:  "div",
		},
		{
			name:           "name with prefix and namespace",
			namespace:      "http://example.com/ns",
			qualifiedName:  "ex:element",
			expectError:    false,
			expectedNS:     "http://example.com/ns",
			expectedPrefix: "ex",
			expectedLocal:  "element",
		},
		{
			name:           "xml prefix with correct namespace",
			namespace:      "http://www.w3.org/XML/1998/namespace",
			qualifiedName:  "xml:lang",
			expectError:    false,
			expectedNS:     "http://www.w3.org/XML/1998/namespace",
			expectedPrefix: "xml",
			expectedLocal:  "lang",
		},
		{
			name:           "xmlns attribute with correct namespace",
			namespace:      "http://www.w3.org/2000/xmlns/",
			qualifiedName:  "xmlns",
			expectError:    false,
			expectedNS:     "http://www.w3.org/2000/xmlns/",
			expectedPrefix: "",
			expectedLocal:  "xmlns",
		},
		{
			name:           "xmlns prefix with correct namespace",
			namespace:      "http://www.w3.org/2000/xmlns/",
			qualifiedName:  "xmlns:prefix",
			expectError:    false,
			expectedNS:     "http://www.w3.org/2000/xmlns/",
			expectedPrefix: "xmlns",
			expectedLocal:  "prefix",
		},

		// Invalid combinations - prefix without namespace
		{
			name:          "prefix without namespace",
			namespace:     "",
			qualifiedName: "prefix:element",
			expectError:   true,
			errorType:     "NamespaceError",
		},

		// Invalid combinations - xml prefix with wrong namespace
		{
			name:          "xml prefix with wrong namespace",
			namespace:     "http://example.com/ns",
			qualifiedName: "xml:lang",
			expectError:   true,
			errorType:     "NamespaceError",
		},

		// Invalid combinations - xmlns with wrong namespace
		{
			name:          "xmlns attribute with wrong namespace",
			namespace:     "http://example.com/ns",
			qualifiedName: "xmlns",
			expectError:   true,
			errorType:     "NamespaceError",
		},
		{
			name:          "xmlns prefix with wrong namespace",
			namespace:     "http://example.com/ns",
			qualifiedName: "xmlns:prefix",
			expectError:   true,
			errorType:     "NamespaceError",
		},

		// Invalid combinations - xmlns namespace with wrong name
		{
			name:          "xmlns namespace with regular element",
			namespace:     "http://www.w3.org/2000/xmlns/",
			qualifiedName: "element",
			expectError:   true,
			errorType:     "NamespaceError",
		},

		// Invalid qualified names should still be caught
		{
			name:          "invalid qualified name",
			namespace:     "http://example.com/ns",
			qualifiedName: "prefix:element:invalid",
			expectError:   true,
			errorType:     "NamespaceError",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ns, prefix, localName, err := ValidateAndExtract(tt.namespace, tt.qualifiedName)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error for namespace='%s', qualifiedName='%s', but got none", tt.namespace, tt.qualifiedName)
					return
				}

				// Check error type
				switch tt.errorType {
				case "InvalidCharacterError":
					if _, ok := err.(*InvalidCharacterError); !ok {
						t.Errorf("Expected InvalidCharacterError, got %T: %v", err, err)
					}
				case "NamespaceError":
					if _, ok := err.(*NamespaceError); !ok {
						t.Errorf("Expected NamespaceError, got %T: %v", err, err)
					}
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error for namespace='%s', qualifiedName='%s': %v", tt.namespace, tt.qualifiedName, err)
					return
				}

				// Check extracted values
				if ns != tt.expectedNS {
					t.Errorf("Expected namespace '%s', got '%s'", tt.expectedNS, ns)
				}
				if prefix != tt.expectedPrefix {
					t.Errorf("Expected prefix '%s', got '%s'", tt.expectedPrefix, prefix)
				}
				if localName != tt.expectedLocal {
					t.Errorf("Expected local name '%s', got '%s'", tt.expectedLocal, localName)
				}
			}
		})
	}
}

func TestIsValidXMLName(t *testing.T) {
	tests := []struct {
		name     string
		xmlName  string
		expected bool
	}{
		// Valid names
		{"simple name", "element", true},
		{"with underscore", "my_element", true},
		{"with hyphen", "my-element", true},
		{"with numbers", "element123", true},
		{"with colon", "prefix:element", true},
		{"unicode characters", "élément", true},
		{"single character", "a", true},
		{"single underscore", "_", true},

		// Invalid names
		{"empty string", "", false},
		{"starts with number", "123element", false},
		{"starts with hyphen", "-element", false},
		{"starts with dot", ".element", false},
		{"contains space", "my element", false},
		{"contains invalid char", "element!", false},
		{"contains @", "element@", false},
		{"contains #", "element#", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isValidXMLName(tt.xmlName)
			if result != tt.expected {
				t.Errorf("isValidXMLName('%s') = %v, expected %v", tt.xmlName, result, tt.expected)
			}
		})
	}
}

func TestGetWellKnownNamespace(t *testing.T) {
	tests := []struct {
		prefix   string
		expected string
	}{
		{"xml", "http://www.w3.org/XML/1998/namespace"},
		{"xmlns", "http://www.w3.org/2000/xmlns/"},
		{"html", "http://www.w3.org/1999/xhtml"},
		{"svg", "http://www.w3.org/2000/svg"},
		{"mathml", "http://www.w3.org/1998/Math/MathML"},
		{"math", "http://www.w3.org/1998/Math/MathML"},
		{"unknown", ""},
		{"", ""},
	}

	for _, tt := range tests {
		t.Run(tt.prefix, func(t *testing.T) {
			result := GetWellKnownNamespace(tt.prefix)
			if result != tt.expected {
				t.Errorf("GetWellKnownNamespace('%s') = '%s', expected '%s'", tt.prefix, result, tt.expected)
			}
		})
	}
}

func TestParseQualifiedName(t *testing.T) {
	tests := []struct {
		name           string
		qualifiedName  string
		expectedNS     string
		expectedPrefix string
		expectedLocal  string
	}{
		{
			name:           "simple name",
			qualifiedName:  "div",
			expectedNS:     "",
			expectedPrefix: "",
			expectedLocal:  "div",
		},
		{
			name:           "xml prefixed name",
			qualifiedName:  "xml:lang",
			expectedNS:     "http://www.w3.org/XML/1998/namespace",
			expectedPrefix: "xml",
			expectedLocal:  "lang",
		},
		{
			name:           "svg prefixed name",
			qualifiedName:  "svg:rect",
			expectedNS:     "http://www.w3.org/2000/svg",
			expectedPrefix: "svg",
			expectedLocal:  "rect",
		},
		{
			name:           "unknown prefix",
			qualifiedName:  "custom:element",
			expectedNS:     "",
			expectedPrefix: "custom",
			expectedLocal:  "element",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			info := ParseQualifiedName(tt.qualifiedName)

			if info.NamespaceURI != tt.expectedNS {
				t.Errorf("Expected namespace '%s', got '%s'", tt.expectedNS, info.NamespaceURI)
			}
			if info.Prefix != tt.expectedPrefix {
				t.Errorf("Expected prefix '%s', got '%s'", tt.expectedPrefix, info.Prefix)
			}
			if info.LocalName != tt.expectedLocal {
				t.Errorf("Expected local name '%s', got '%s'", tt.expectedLocal, info.LocalName)
			}
		})
	}
}

func TestNamespaceErrorTypes(t *testing.T) {
	// Test NamespaceError
	nsErr := &NamespaceError{Message: "test namespace error"}
	if nsErr.Error() != "test namespace error" {
		t.Errorf("NamespaceError.Error() = '%s', expected 'test namespace error'", nsErr.Error())
	}

	// Test InvalidCharacterError
	charErr := &InvalidCharacterError{Message: "test character error"}
	if charErr.Error() != "test character error" {
		t.Errorf("InvalidCharacterError.Error() = '%s', expected 'test character error'", charErr.Error())
	}
}

// Unicode tests for XML Name validation
func TestUnicodeNameValidation(t *testing.T) {
	tests := []struct {
		name     string
		xmlName  string
		expected bool
	}{
		// Valid Unicode characters
		{"latin extended", "café", true},
		{"greek letters", "αβγ", true},
		{"cyrillic", "тест", true},
		{"chinese characters", "测试", true},
		{"combining marks", "e\u0301", true}, // e with acute accent combining mark

		// Invalid Unicode characters
		{"control characters", "test\u0001", false},
		{"private use area", "test\uE000", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isValidXMLName(tt.xmlName)
			if result != tt.expected {
				t.Errorf("isValidXMLName('%s') = %v, expected %v", tt.xmlName, result, tt.expected)
			}
		})
	}
}
