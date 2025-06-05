package dom

import (
	"testing"
)

// TestDOMTokenListFullSpecCompliance provides comprehensive testing against WHATWG DOM Section 7.1
// This ensures our implementation matches the specification exactly.
func TestDOMTokenListFullSpecCompliance(t *testing.T) {
	doc := NewDocument()
	element := NewElement("div", doc)

	t.Run("Interface DOMTokenList - Basic Properties", func(t *testing.T) {
		classList := NewDOMTokenList(element, "class")
		element.SetAttribute("class", "foo bar baz")

		// Test readonly attribute unsigned long length
		if classList.Length() != 3 {
			t.Errorf("expected length 3, got %d", classList.Length())
		}

		// Test getter DOMString? item(unsigned long index)
		if classList.Item(0) != "foo" {
			t.Errorf("expected item(0) to be 'foo', got '%s'", classList.Item(0))
		}
		if classList.Item(999) != "" {
			t.Errorf("expected item(999) to return empty string for out of bounds, got '%s'", classList.Item(999))
		}
	})

	t.Run("Token Set Management", func(t *testing.T) {
		// Use a fresh element to ensure clean state
		freshElement := NewElement("div", doc)
		classList := NewDOMTokenList(freshElement, "class")

		// Test that DOMTokenList has an associated token set (initially empty)
		if classList.Length() != 0 {
			t.Errorf("expected initial token set to be empty, got length %d", classList.Length())
		}

		// Test that token set is ordered and unique
		freshElement.SetAttribute("class", "c b a b c")
		if classList.Length() != 3 {
			t.Errorf("expected 3 unique tokens, got %d", classList.Length())
		}
		if classList.Item(0) != "c" || classList.Item(1) != "b" || classList.Item(2) != "a" {
			t.Error("expected tokens to maintain first-occurrence order: c, b, a")
		}
	})

	t.Run("Add Method - Full Specification", func(t *testing.T) {
		classList := NewDOMTokenList(element, "class")
		element.SetAttribute("class", "")

		// Test adding multiple tokens in one call
		classList.Add("first", "second", "third")
		if element.GetAttribute("class") != "first second third" {
			t.Errorf("expected 'first second third', got '%s'", element.GetAttribute("class"))
		}

		// Test that duplicate tokens are ignored
		classList.Add("first", "fourth", "second")
		if element.GetAttribute("class") != "first second third fourth" {
			t.Errorf("expected 'first second third fourth', got '%s'", element.GetAttribute("class"))
		}

		// Test all ASCII whitespace characters
		testWhitespaceChars := []string{" ", "\t", "\n", "\r", "\f"}
		for _, ws := range testWhitespaceChars {
			t.Run("Add with "+ws+" character", func(t *testing.T) {
				defer func() {
					r := recover()
					if r == nil {
						t.Errorf("expected panic for token containing ASCII whitespace: %q", ws)
						return
					}
					if exception, ok := r.(*DOMException); ok {
						if exception.Name() != "InvalidCharacterError" {
							t.Errorf("expected InvalidCharacterError, got %s", exception.Name())
						}
					} else {
						t.Errorf("expected DOMException, got %T", r)
					}
				}()
				classList.Add("invalid" + ws + "token")
			})
		}
	})

	t.Run("Remove Method - Full Specification", func(t *testing.T) {
		classList := NewDOMTokenList(element, "class")
		element.SetAttribute("class", "one two three four five")

		// Test removing multiple tokens in one call
		classList.Remove("two", "four")
		if element.GetAttribute("class") != "one three five" {
			t.Errorf("expected 'one three five', got '%s'", element.GetAttribute("class"))
		}

		// Test removing non-existent tokens (should be no-op)
		classList.Remove("nonexistent", "alsomissing")
		if element.GetAttribute("class") != "one three five" {
			t.Errorf("expected no change when removing non-existent tokens, got '%s'", element.GetAttribute("class"))
		}

		// Test removing all tokens
		classList.Remove("one", "three", "five")
		if element.GetAttribute("class") != "" {
			t.Errorf("expected empty class after removing all tokens, got '%s'", element.GetAttribute("class"))
		}
	})

	t.Run("Toggle Method - Complete Specification Compliance", func(t *testing.T) {
		classList := NewDOMTokenList(element, "class")

		// Test toggle without force parameter
		element.SetAttribute("class", "existing")

		// Toggle existing token (should remove and return false)
		result := classList.Toggle("existing", nil)
		if result != false {
			t.Error("expected toggle of existing token to return false")
		}
		if classList.Contains("existing") {
			t.Error("expected 'existing' to be removed")
		}

		// Toggle non-existing token (should add and return true)
		result = classList.Toggle("new", nil)
		if result != true {
			t.Error("expected toggle of non-existing token to return true")
		}
		if !classList.Contains("new") {
			t.Error("expected 'new' to be added")
		}

		// Test force=true on existing token (should not remove, return true)
		forceTrue := true
		result = classList.Toggle("new", &forceTrue)
		if result != true {
			t.Error("expected toggle with force=true on existing token to return true")
		}
		if !classList.Contains("new") {
			t.Error("expected 'new' to remain when force=true")
		}

		// Test force=false on existing token (should remove, return false)
		forceFalse := false
		result = classList.Toggle("new", &forceFalse)
		if result != false {
			t.Error("expected toggle with force=false on existing token to return false")
		}
		if classList.Contains("new") {
			t.Error("expected 'new' to be removed when force=false")
		}

		// Test force=true on non-existing token (should add, return true)
		result = classList.Toggle("another", &forceTrue)
		if result != true {
			t.Error("expected toggle with force=true on non-existing token to return true")
		}
		if !classList.Contains("another") {
			t.Error("expected 'another' to be added when force=true")
		}

		// Test force=false on non-existing token (should not add, return false)
		result = classList.Toggle("missing", &forceFalse)
		if result != false {
			t.Error("expected toggle with force=false on non-existing token to return false")
		}
		if classList.Contains("missing") {
			t.Error("expected 'missing' to not be added when force=false")
		}
	})

	t.Run("Replace Method - Complete Specification Compliance", func(t *testing.T) {
		classList := NewDOMTokenList(element, "class")

		// Test replace existing token with new token
		element.SetAttribute("class", "alpha beta gamma")
		result := classList.Replace("beta", "newbeta")
		if result != true {
			t.Error("expected replace of existing token to return true")
		}
		if element.GetAttribute("class") != "alpha newbeta gamma" {
			t.Errorf("expected 'alpha newbeta gamma', got '%s'", element.GetAttribute("class"))
		}

		// Test replace non-existing token (should return false, no change)
		result = classList.Replace("missing", "replacement")
		if result != false {
			t.Error("expected replace of non-existing token to return false")
		}
		if element.GetAttribute("class") != "alpha newbeta gamma" {
			t.Errorf("expected no change when replacing non-existing token, got '%s'", element.GetAttribute("class"))
		}

		// Test replace with existing token (deduplication behavior)
		element.SetAttribute("class", "foo bar baz")
		result = classList.Replace("bar", "foo") // Replace bar with existing foo
		if result != true {
			t.Error("expected replace with existing token to return true")
		}
		// Should remove the old foo and replace bar with foo, resulting in deduplication
		if element.GetAttribute("class") != "foo baz" {
			t.Errorf("expected 'foo baz' after deduplication, got '%s'", element.GetAttribute("class"))
		}

		// Test replace token with itself (should be no-op but return true)
		element.SetAttribute("class", "same other")
		result = classList.Replace("same", "same")
		if result != true {
			t.Error("expected replace token with itself to return true")
		}
		if element.GetAttribute("class") != "same other" {
			t.Errorf("expected no change when replacing token with itself, got '%s'", element.GetAttribute("class"))
		}
	})

	t.Run("Supports Method - Validation Steps", func(t *testing.T) {
		// Test with no supported tokens defined (should return true for compatibility)
		classList := NewDOMTokenList(element, "class")
		if !classList.Supports("anything") {
			t.Error("expected Supports to return true when no supported tokens defined")
		}

		// Test with supported tokens defined
		supportedTokens := []string{"valid", "allowed", "UPPERCASE"}
		classListWithSupport := NewDOMTokenListWithSupportedTokens(element, "rel", supportedTokens)

		// Test case-insensitive matching (spec step 2: ASCII lowercase)
		if !classListWithSupport.Supports("VALID") {
			t.Error("expected case-insensitive match for 'VALID'")
		}
		if !classListWithSupport.Supports("uppercase") {
			t.Error("expected case-insensitive match for 'uppercase'")
		}
		if classListWithSupport.Supports("invalid") {
			t.Error("expected 'invalid' to not be supported")
		}
	})

	t.Run("Value Getter and Setter", func(t *testing.T) {
		classList := NewDOMTokenList(element, "class")

		// Test value getter (serialize steps)
		element.SetAttribute("class", "serialized value")
		if classList.Value() != "serialized value" {
			t.Errorf("expected Value() to return attribute value directly, got '%s'", classList.Value())
		}

		// Test value setter
		classList.SetValue("new token set")
		if element.GetAttribute("class") != "new token set" {
			t.Errorf("expected SetValue to update attribute, got '%s'", element.GetAttribute("class"))
		}
		if classList.Value() != "new token set" {
			t.Errorf("expected Value() to reflect new value, got '%s'", classList.Value())
		}

		// Test that setter properly invalidates cache
		if classList.Length() != 3 { // "new", "token", "set"
			t.Errorf("expected length 3 after setValue, got %d", classList.Length())
		}
	})

	t.Run("Update Steps - Specification Compliance", func(t *testing.T) {
		// Test case 1: Element has no attribute and token set is empty
		freshElement := NewElement("span", doc)
		classList := NewDOMTokenList(freshElement, "class")

		// This should not create the attribute per spec step 1
		classList.runUpdateSteps()
		if freshElement.HasAttribute("class") {
			t.Error("expected update steps to not create attribute when element has no attribute and token set is empty")
		}

		// Test case 2: Token set has content, should create/update attribute
		classList.Add("test")
		if !freshElement.HasAttribute("class") {
			t.Error("expected update steps to create attribute when token set is not empty")
		}
		if freshElement.GetAttribute("class") != "test" {
			t.Errorf("expected attribute value 'test', got '%s'", freshElement.GetAttribute("class"))
		}
	})

	t.Run("Serialize Steps", func(t *testing.T) {
		// Use a fresh element to ensure clean state
		freshElement := NewElement("div", doc)
		classList := NewDOMTokenList(freshElement, "class")
		freshElement.SetAttribute("class", "  multiple   spaces   between  ")

		// Serialize steps should return the raw attribute value
		serialized := classList.runSerializeSteps()
		if serialized != "  multiple   spaces   between  " {
			t.Errorf("expected serialize steps to return raw attribute value, got '%s'", serialized)
		}

		// But parsed token set should be normalized to 3 tokens: "multiple", "spaces", "between"
		if classList.Length() != 3 {
			t.Errorf("expected 3 parsed tokens, got %d", classList.Length())
		}
	})

	t.Run("Attribute Change Steps", func(t *testing.T) {
		// Use a fresh element to ensure clean state
		freshElement := NewElement("div", doc)
		classList := NewDOMTokenList(freshElement, "class")

		// Test that DOMTokenList responds to attribute changes
		freshElement.SetAttribute("class", "new value")
		if classList.Length() != 2 { // "new", "value"
			t.Errorf("expected 2 tokens after attribute change, got %d", classList.Length())
		}
		if !classList.Contains("new") || !classList.Contains("value") {
			t.Error("expected token set to be updated with new attribute value")
		}

		// Test with empty value (should empty token set)
		freshElement.SetAttribute("class", "")
		if classList.Length() != 0 {
			t.Errorf("expected empty token set when value is empty, got %d tokens", classList.Length())
		}

		// Test removing the attribute entirely
		freshElement.RemoveAttribute("class")
		if classList.Length() != 0 {
			t.Errorf("expected empty token set when attribute is removed, got %d tokens", classList.Length())
		}

		// Test that changes to other attributes don't affect the token list
		originalLength := classList.Length()
		freshElement.SetAttribute("id", "test-id")
		if classList.Length() != originalLength {
			t.Error("expected no change when modifying different attribute")
		}
	})

	t.Run("Ordered Set Parser - All ASCII Whitespace", func(t *testing.T) {
		// Test all ASCII whitespace characters as defined in the spec
		input := "  \t\n\r\f  token1  \t\n\r\f  token2  \t\n\r\f  token1  \t\n\r\f  "
		tokens := parseOrderedSet(input)

		expected := []string{"token1", "token2"}
		if len(tokens) != len(expected) {
			t.Errorf("expected %d tokens, got %d", len(expected), len(tokens))
		}
		for i, token := range tokens {
			if token != expected[i] {
				t.Errorf("expected token[%d] to be '%s', got '%s'", i, expected[i], token)
			}
		}

		// Test empty string
		tokens = parseOrderedSet("")
		if len(tokens) != 0 {
			t.Errorf("expected empty token set for empty string, got %d tokens", len(tokens))
		}

		// Test only whitespace
		tokens = parseOrderedSet("  \t\n\r\f  ")
		if len(tokens) != 0 {
			t.Errorf("expected empty token set for whitespace-only string, got %d tokens", len(tokens))
		}
	})

	t.Run("Ordered Set Serializer", func(t *testing.T) {
		// Test empty set
		result := serializeOrderedSet([]string{})
		if result != "" {
			t.Errorf("expected empty string for empty set, got '%s'", result)
		}

		// Test single token
		result = serializeOrderedSet([]string{"single"})
		if result != "single" {
			t.Errorf("expected 'single', got '%s'", result)
		}

		// Test multiple tokens
		result = serializeOrderedSet([]string{"first", "second", "third"})
		if result != "first second third" {
			t.Errorf("expected 'first second third', got '%s'", result)
		}
	})

	t.Run("Live Collection Behavior", func(t *testing.T) {
		classList := NewDOMTokenList(element, "class")
		element.SetAttribute("class", "initial")

		// Verify initial state
		if !classList.Contains("initial") {
			t.Error("expected classList to contain 'initial'")
		}

		// Modify attribute externally and verify live update
		element.SetAttribute("class", "modified external")
		if classList.Contains("initial") {
			t.Error("expected classList to no longer contain 'initial' after external modification")
		}
		if !classList.Contains("modified") || !classList.Contains("external") {
			t.Error("expected classList to reflect external modifications")
		}

		// Remove attribute and verify live update
		element.RemoveAttribute("class")
		if classList.Length() != 0 {
			t.Errorf("expected empty classList after attribute removal, got %d tokens", classList.Length())
		}
	})

	t.Run("Error Handling - All Validation Cases", func(t *testing.T) {
		classList := NewDOMTokenList(element, "class")

		// Test all methods with empty string
		emptyStringMethods := []struct {
			name string
			test func()
		}{
			{"Add", func() { classList.Add("") }},
			{"Remove", func() { classList.Remove("") }},
			{"Toggle", func() { classList.Toggle("", nil) }},
			{"Replace-token", func() { classList.Replace("", "valid") }},
			{"Replace-newToken", func() { classList.Replace("valid", "") }},
		}

		for _, method := range emptyStringMethods {
			t.Run("SyntaxError for empty string in "+method.name, func(t *testing.T) {
				defer func() {
					r := recover()
					if r == nil {
						t.Errorf("expected panic for empty string in %s", method.name)
						return
					}
					if exception, ok := r.(*DOMException); ok {
						if exception.Name() != "SyntaxError" {
							t.Errorf("expected SyntaxError, got %s", exception.Name())
						}
					} else {
						t.Errorf("expected DOMException, got %T", r)
					}
				}()
				method.test()
			})
		}

		// Test all methods with whitespace
		whitespaceStringMethods := []struct {
			name string
			test func()
		}{
			{"Add", func() { classList.Add("has space") }},
			{"Remove", func() { classList.Remove("has\ttab") }},
			{"Toggle", func() { classList.Toggle("has\nnewline", nil) }},
			{"Replace-token", func() { classList.Replace("has\rcarriage", "valid") }},
			{"Replace-newToken", func() { classList.Replace("valid", "has\fformfeed") }},
		}

		for _, method := range whitespaceStringMethods {
			t.Run("InvalidCharacterError for whitespace in "+method.name, func(t *testing.T) {
				defer func() {
					r := recover()
					if r == nil {
						t.Errorf("expected panic for whitespace in %s", method.name)
						return
					}
					if exception, ok := r.(*DOMException); ok {
						if exception.Name() != "InvalidCharacterError" {
							t.Errorf("expected InvalidCharacterError, got %s", exception.Name())
						}
					} else {
						t.Errorf("expected DOMException, got %T", r)
					}
				}()
				method.test()
			})
		}
	})

	t.Run("Edge Cases and Corner Cases", func(t *testing.T) {
		classList := NewDOMTokenList(element, "class")

		// Test with very long token
		longToken := string(make([]rune, 1000))
		for i := range longToken {
			longToken = longToken[:i] + "a" + longToken[i+1:]
		}
		classList.Add(longToken)
		if !classList.Contains(longToken) {
			t.Error("expected very long token to be handled correctly")
		}

		// Test with Unicode characters
		unicodeToken := "token-with-ñ-and-emoji-🎉"
		classList.Add(unicodeToken)
		if !classList.Contains(unicodeToken) {
			t.Error("expected Unicode token to be handled correctly")
		}

		// Test rapid successive modifications
		classList.SetValue("")
		for i := 0; i < 100; i++ {
			classList.Add("token" + string(rune('0'+i%10)))
		}
		if classList.Length() != 10 { // Should only have token0-token9
			t.Errorf("expected 10 unique tokens after rapid modifications, got %d", classList.Length())
		}
	})
}
