package dom

import (
	"testing"
)

// TestDocumentOrShadowRootMixin tests the DocumentOrShadowRoot mixin implementation
// per WHATWG DOM Section 4.2.5
func TestDocumentOrShadowRootMixin(t *testing.T) {
	t.Run("Document CustomElementRegistry getter", func(t *testing.T) {
		doc := NewDocument()

		// By default, customElementRegistry should be nil
		registry := doc.CustomElementRegistry()
		if registry != nil {
			t.Errorf("Expected CustomElementRegistry to be nil by default, got %v", registry)
		}
	})

	t.Run("ShadowRoot CustomElementRegistry getter", func(t *testing.T) {
		doc := NewDocument()
		host := doc.CreateElement("div")
		shadowRoot := NewShadowRoot(host, ShadowRootModeOpen)

		// By default, customElementRegistry should be nil
		registry := shadowRoot.CustomElementRegistry()
		if registry != nil {
			t.Errorf("Expected CustomElementRegistry to be nil by default, got %v", registry)
		}
	})

	t.Run("CustomElementRegistry creation and association", func(t *testing.T) {
		doc := NewDocument()

		// Create a custom element registry
		registry := NewCustomElementRegistry(doc)
		if registry == nil {
			t.Fatal("Failed to create CustomElementRegistry")
		}

		// Check registry document association
		if registry.Document() != doc {
			t.Errorf("Expected registry document to be %v, got %v", doc, registry.Document())
		}
	})

	t.Run("CustomElementRegistry placeholder methods", func(t *testing.T) {
		doc := NewDocument()
		registry := NewCustomElementRegistry(doc)

		// Test Define method (should return NotSupportedError)
		err := registry.Define("my-element", nil)
		if err == nil {
			t.Error("Expected Define to return an error")
		}
		if _, ok := err.(*DOMException); !ok {
			t.Errorf("Expected DOMException, got %T", err)
		}

		// Test Get method (should return nil)
		constructor := registry.Get("my-element")
		if constructor != nil {
			t.Errorf("Expected Get to return nil, got %v", constructor)
		}

		// Test WhenDefined method (should return nil for now)
		promise := registry.WhenDefined("my-element")
		if promise != nil {
			t.Errorf("Expected WhenDefined to return nil, got %v", promise)
		}

		// Test Upgrade method (should not panic)
		registry.Upgrade(doc)
	})

	t.Run("DocumentOrShadowRoot mixin specification compliance", func(t *testing.T) {
		// Test that both Document and ShadowRoot implement the mixin
		doc := NewDocument()
		host := doc.CreateElement("div")
		shadowRoot := NewShadowRoot(host, ShadowRootModeOpen)

		// Both should have CustomElementRegistry method and return nil by default
		docRegistry := doc.CustomElementRegistry()
		shadowRegistry := shadowRoot.CustomElementRegistry()

		// Both should return nil by default
		if docRegistry != nil {
			t.Errorf("Document CustomElementRegistry should be nil by default, got %v", docRegistry)
		}
		if shadowRegistry != nil {
			t.Errorf("ShadowRoot CustomElementRegistry should be nil by default, got %v", shadowRegistry)
		}
	})
}

// TestCustomElementRegistryPlaceholderImplementation tests the placeholder methods
func TestCustomElementRegistryPlaceholderImplementation(t *testing.T) {
	doc := NewDocument()
	registry := NewCustomElementRegistry(doc)

	t.Run("Define method placeholder", func(t *testing.T) {
		// Test with various parameter combinations
		testCases := []struct {
			name        string
			constructor interface{}
			options     []interface{}
		}{
			{"simple-element", nil, nil},
			{"complex-element", func() {}, []interface{}{"option1"}},
		}

		for _, tc := range testCases {
			err := registry.Define(tc.name, tc.constructor, tc.options...)
			if err == nil {
				t.Errorf("Expected error for Define(%s), got nil", tc.name)
			}

			// Should be a NotSupportedError
			if domErr, ok := err.(*DOMException); ok {
				if domErr.Name() != "NotSupportedError" {
					t.Errorf("Expected NotSupportedError, got %s", domErr.Name())
				}
			} else {
				t.Errorf("Expected DOMException, got %T", err)
			}
		}
	})

	t.Run("Get method placeholder", func(t *testing.T) {
		testCases := []string{
			"my-element",
			"custom-button",
			"special-div",
			"",
		}

		for _, name := range testCases {
			result := registry.Get(name)
			if result != nil {
				t.Errorf("Expected Get(%s) to return nil, got %v", name, result)
			}
		}
	})

	t.Run("WhenDefined method placeholder", func(t *testing.T) {
		testCases := []string{
			"my-element",
			"custom-button",
			"special-div",
			"",
		}

		for _, name := range testCases {
			result := registry.WhenDefined(name)
			if result != nil {
				t.Errorf("Expected WhenDefined(%s) to return nil, got %v", name, result)
			}
		}
	})

	t.Run("Upgrade method placeholder", func(t *testing.T) {
		// Test with various node types
		testNodes := []Node{
			doc,
			doc.CreateElement("div"),
			doc.CreateTextNode("text"),
			doc.CreateDocumentFragment(),
		}

		for i, node := range testNodes {
			// Should not panic
			func() {
				defer func() {
					if r := recover(); r != nil {
						t.Errorf("Upgrade panicked on test case %d: %v", i, r)
					}
				}()
				registry.Upgrade(node)
			}()
		}
	})
}

// TestCustomElementRegistryAssociation tests the association between registry and document
func TestCustomElementRegistryAssociation(t *testing.T) {
	t.Run("Document association", func(t *testing.T) {
		doc := NewDocument()
		registry := NewCustomElementRegistry(doc)

		if registry.Document() != doc {
			t.Errorf("Expected registry document to be %v, got %v", doc, registry.Document())
		}
	})

	t.Run("Multiple registries for different documents", func(t *testing.T) {
		doc1 := NewDocument()
		doc2 := NewDocument()

		registry1 := NewCustomElementRegistry(doc1)
		registry2 := NewCustomElementRegistry(doc2)

		if registry1.Document() != doc1 {
			t.Errorf("Expected registry1 document to be doc1, got %v", registry1.Document())
		}
		if registry2.Document() != doc2 {
			t.Errorf("Expected registry2 document to be doc2, got %v", registry2.Document())
		}

		// Registries should be different instances
		if registry1 == registry2 {
			t.Error("Expected different registry instances for different documents")
		}
	})

	t.Run("Nil document handling", func(t *testing.T) {
		// This should create a registry with nil document
		registry := NewCustomElementRegistry(nil)
		if registry.Document() != nil {
			t.Errorf("Expected registry document to be nil, got %v", registry.Document())
		}
	})
}

// TestDocumentOrShadowRootSpecificationCompliance tests specification compliance
func TestDocumentOrShadowRootSpecificationCompliance(t *testing.T) {
	t.Run("getter steps implementation", func(t *testing.T) {
		// Test step 1: If this is a document, return this's custom element registry
		doc := NewDocument()
		docRegistry := doc.CustomElementRegistry()
		if docRegistry != nil {
			t.Error("Document should return nil custom element registry by default")
		}

		// Test step 2: Assert: this is a ShadowRoot node
		// Test step 3: Return this's custom element registry
		host := doc.CreateElement("div")
		shadowRoot := NewShadowRoot(host, ShadowRootModeOpen)
		shadowRegistry := shadowRoot.CustomElementRegistry()
		if shadowRegistry != nil {
			t.Error("ShadowRoot should return nil custom element registry by default")
		}
	})

	t.Run("mixin interface coverage", func(t *testing.T) {
		// Verify that both Document and ShadowRoot implement the mixin
		doc := NewDocument()
		host := doc.CreateElement("div")
		shadowRoot := NewShadowRoot(host, ShadowRootModeOpen)

		// Both should have the CustomElementRegistry method
		// This test ensures the method exists and is callable
		_ = doc.CustomElementRegistry()
		_ = shadowRoot.CustomElementRegistry()
	})

	t.Run("readonly attribute behavior", func(t *testing.T) {
		// The customElementRegistry attribute should be readonly
		// In our implementation, this means we only provide getters
		doc := NewDocument()
		host := doc.CreateElement("div")
		shadowRoot := NewShadowRoot(host, ShadowRootModeOpen)

		// Test that getting the registry multiple times returns the same value
		docRegistry1 := doc.CustomElementRegistry()
		docRegistry2 := doc.CustomElementRegistry()
		if docRegistry1 != docRegistry2 {
			t.Error("Document CustomElementRegistry should return consistent values")
		}

		shadowRegistry1 := shadowRoot.CustomElementRegistry()
		shadowRegistry2 := shadowRoot.CustomElementRegistry()
		if shadowRegistry1 != shadowRegistry2 {
			t.Error("ShadowRoot CustomElementRegistry should return consistent values")
		}
	})
}
