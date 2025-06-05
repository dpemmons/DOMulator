package dom

import (
	"testing"
)

func TestDOMTokenListSpecCompliance(t *testing.T) {
	doc := NewDocument()
	element := NewElement("div", doc)

	t.Run("Length", func(t *testing.T) {
		classList := NewDOMTokenList(element, "class")

		// Empty list should have length 0
		if classList.Length() != 0 {
			t.Errorf("expected length 0, got %d", classList.Length())
		}

		element.SetAttribute("class", "foo bar baz")
		if classList.Length() != 3 {
			t.Errorf("expected length 3, got %d", classList.Length())
		}
	})

	t.Run("Item", func(t *testing.T) {
		classList := NewDOMTokenList(element, "class")
		element.SetAttribute("class", "foo bar baz")

		// Test valid indices
		if classList.Item(0) != "foo" {
			t.Errorf("expected 'foo', got '%s'", classList.Item(0))
		}
		if classList.Item(1) != "bar" {
			t.Errorf("expected 'bar', got '%s'", classList.Item(1))
		}
		if classList.Item(2) != "baz" {
			t.Errorf("expected 'baz', got '%s'", classList.Item(2))
		}

		// Test out of bounds - should return empty string per Go conventions
		if classList.Item(3) != "" {
			t.Errorf("expected empty string for out of bounds, got '%s'", classList.Item(3))
		}
		if classList.Item(-1) != "" {
			t.Errorf("expected empty string for negative index, got '%s'", classList.Item(-1))
		}
	})

	t.Run("Contains", func(t *testing.T) {
		classList := NewDOMTokenList(element, "class")
		element.SetAttribute("class", "foo bar baz")

		if !classList.Contains("foo") {
			t.Error("expected Contains('foo') to be true")
		}
		if !classList.Contains("bar") {
			t.Error("expected Contains('bar') to be true")
		}
		if !classList.Contains("baz") {
			t.Error("expected Contains('baz') to be true")
		}
		if classList.Contains("notfound") {
			t.Error("expected Contains('notfound') to be false")
		}
	})

	t.Run("Add - Valid Tokens", func(t *testing.T) {
		classList := NewDOMTokenList(element, "class")
		element.SetAttribute("class", "")

		// Add single token
		classList.Add("foo")
		if !classList.Contains("foo") {
			t.Error("expected 'foo' to be added")
		}
		if element.GetAttribute("class") != "foo" {
			t.Errorf("expected class attribute to be 'foo', got '%s'", element.GetAttribute("class"))
		}

		// Add multiple tokens
		classList.Add("bar", "baz")
		if !classList.Contains("bar") || !classList.Contains("baz") {
			t.Error("expected 'bar' and 'baz' to be added")
		}
		if element.GetAttribute("class") != "foo bar baz" {
			t.Errorf("expected class attribute to be 'foo bar baz', got '%s'", element.GetAttribute("class"))
		}

		// Add duplicate token (should be ignored)
		classList.Add("foo")
		if element.GetAttribute("class") != "foo bar baz" {
			t.Errorf("expected class attribute unchanged after adding duplicate, got '%s'", element.GetAttribute("class"))
		}
	})

	t.Run("Add - SyntaxError for Empty Token", func(t *testing.T) {
		classList := NewDOMTokenList(element, "class")

		defer func() {
			r := recover()
			if r == nil {
				t.Error("expected panic for empty token")
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

		classList.Add("")
	})

	t.Run("Add - InvalidCharacterError for Whitespace", func(t *testing.T) {
		classList := NewDOMTokenList(element, "class")

		defer func() {
			r := recover()
			if r == nil {
				t.Error("expected panic for token with whitespace")
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

		classList.Add("invalid token")
	})

	t.Run("Remove - Valid Tokens", func(t *testing.T) {
		classList := NewDOMTokenList(element, "class")
		element.SetAttribute("class", "foo bar baz")

		// Remove single token
		classList.Remove("bar")
		if classList.Contains("bar") {
			t.Error("expected 'bar' to be removed")
		}
		if element.GetAttribute("class") != "foo baz" {
			t.Errorf("expected class attribute to be 'foo baz', got '%s'", element.GetAttribute("class"))
		}

		// Remove multiple tokens
		classList.Remove("foo", "baz")
		if classList.Contains("foo") || classList.Contains("baz") {
			t.Error("expected 'foo' and 'baz' to be removed")
		}
		if element.GetAttribute("class") != "" {
			t.Errorf("expected class attribute to be empty, got '%s'", element.GetAttribute("class"))
		}

		// Remove non-existent token (should be ignored)
		classList.Remove("notfound")
		if element.GetAttribute("class") != "" {
			t.Errorf("expected class attribute unchanged after removing non-existent token, got '%s'", element.GetAttribute("class"))
		}
	})

	t.Run("Remove - SyntaxError for Empty Token", func(t *testing.T) {
		classList := NewDOMTokenList(element, "class")

		defer func() {
			r := recover()
			if r == nil {
				t.Error("expected panic for empty token")
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

		classList.Remove("")
	})

	t.Run("Remove - InvalidCharacterError for Whitespace", func(t *testing.T) {
		classList := NewDOMTokenList(element, "class")

		defer func() {
			r := recover()
			if r == nil {
				t.Error("expected panic for token with whitespace")
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

		classList.Remove("invalid token")
	})

	t.Run("Toggle - Without Force", func(t *testing.T) {
		classList := NewDOMTokenList(element, "class")
		element.SetAttribute("class", "foo")

		// Toggle existing token (should remove)
		result := classList.Toggle("foo", nil)
		if result != false {
			t.Error("expected Toggle to return false when removing token")
		}
		if classList.Contains("foo") {
			t.Error("expected 'foo' to be removed")
		}

		// Toggle non-existing token (should add)
		result = classList.Toggle("bar", nil)
		if result != true {
			t.Error("expected Toggle to return true when adding token")
		}
		if !classList.Contains("bar") {
			t.Error("expected 'bar' to be added")
		}
	})

	t.Run("Toggle - With Force", func(t *testing.T) {
		classList := NewDOMTokenList(element, "class")
		element.SetAttribute("class", "foo")

		// Force add existing token (should not change but return true)
		forceTrue := true
		result := classList.Toggle("foo", &forceTrue)
		if result != true {
			t.Error("expected Toggle with force=true to return true")
		}
		if !classList.Contains("foo") {
			t.Error("expected 'foo' to remain")
		}

		// Force remove existing token
		forceFalse := false
		result = classList.Toggle("foo", &forceFalse)
		if result != false {
			t.Error("expected Toggle with force=false to return false")
		}
		if classList.Contains("foo") {
			t.Error("expected 'foo' to be removed")
		}

		// Force add non-existing token
		result = classList.Toggle("bar", &forceTrue)
		if result != true {
			t.Error("expected Toggle with force=true to return true")
		}
		if !classList.Contains("bar") {
			t.Error("expected 'bar' to be added")
		}

		// Force remove non-existing token
		result = classList.Toggle("baz", &forceFalse)
		if result != false {
			t.Error("expected Toggle with force=false to return false")
		}
		if classList.Contains("baz") {
			t.Error("expected 'baz' to not be added")
		}
	})

	t.Run("Toggle - SyntaxError for Empty Token", func(t *testing.T) {
		classList := NewDOMTokenList(element, "class")

		defer func() {
			r := recover()
			if r == nil {
				t.Error("expected panic for empty token")
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

		classList.Toggle("", nil)
	})

	t.Run("Toggle - InvalidCharacterError for Whitespace", func(t *testing.T) {
		classList := NewDOMTokenList(element, "class")

		defer func() {
			r := recover()
			if r == nil {
				t.Error("expected panic for token with whitespace")
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

		classList.Toggle("invalid token", nil)
	})

	t.Run("Replace", func(t *testing.T) {
		classList := NewDOMTokenList(element, "class")
		element.SetAttribute("class", "foo bar baz")

		// Replace existing token
		result := classList.Replace("bar", "newbar")
		if result != true {
			t.Error("expected Replace to return true for existing token")
		}
		if classList.Contains("bar") {
			t.Error("expected 'bar' to be removed")
		}
		if !classList.Contains("newbar") {
			t.Error("expected 'newbar' to be added")
		}
		if element.GetAttribute("class") != "foo newbar baz" {
			t.Errorf("expected class attribute to be 'foo newbar baz', got '%s'", element.GetAttribute("class"))
		}

		// Replace non-existing token
		result = classList.Replace("notfound", "replacement")
		if result != false {
			t.Error("expected Replace to return false for non-existing token")
		}
		if classList.Contains("replacement") {
			t.Error("expected 'replacement' to not be added")
		}

		// Replace with existing token (should deduplicate)
		element.SetAttribute("class", "foo bar baz")
		result = classList.Replace("bar", "foo")
		if result != true {
			t.Error("expected Replace to return true")
		}
		if element.GetAttribute("class") != "foo baz" {
			t.Errorf("expected class attribute to be 'foo baz' after deduplication, got '%s'", element.GetAttribute("class"))
		}
	})

	t.Run("Replace - SyntaxError for Empty Token", func(t *testing.T) {
		classList := NewDOMTokenList(element, "class")

		defer func() {
			r := recover()
			if r == nil {
				t.Error("expected panic for empty token")
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

		classList.Replace("", "valid")
	})

	t.Run("Replace - SyntaxError for Empty NewToken", func(t *testing.T) {
		classList := NewDOMTokenList(element, "class")

		defer func() {
			r := recover()
			if r == nil {
				t.Error("expected panic for empty newToken")
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

		classList.Replace("valid", "")
	})

	t.Run("Replace - InvalidCharacterError for Whitespace", func(t *testing.T) {
		classList := NewDOMTokenList(element, "class")

		defer func() {
			r := recover()
			if r == nil {
				t.Error("expected panic for token with whitespace")
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

		classList.Replace("valid", "invalid token")
	})

	t.Run("Supports", func(t *testing.T) {
		classList := NewDOMTokenList(element, "class")

		// Without supported tokens defined, should return true
		result := classList.Supports("anything")
		if result != true {
			t.Error("expected Supports to return true when no supported tokens defined")
		}

		// With supported tokens defined
		supportedTokens := []string{"foo", "bar", "BAZ"}
		classListWithSupport := NewDOMTokenListWithSupportedTokens(element, "class", supportedTokens)

		if !classListWithSupport.Supports("foo") {
			t.Error("expected Supports('foo') to be true")
		}
		if !classListWithSupport.Supports("FOO") { // Should be case-insensitive
			t.Error("expected Supports('FOO') to be true (case-insensitive)")
		}
		if !classListWithSupport.Supports("baz") { // Lowercase of BAZ
			t.Error("expected Supports('baz') to be true")
		}
		if classListWithSupport.Supports("notfound") {
			t.Error("expected Supports('notfound') to be false")
		}
	})

	t.Run("Value and SetValue", func(t *testing.T) {
		classList := NewDOMTokenList(element, "class")
		element.SetAttribute("class", "foo bar baz")

		// Test Value getter
		if classList.Value() != "foo bar baz" {
			t.Errorf("expected Value() to be 'foo bar baz', got '%s'", classList.Value())
		}

		// Test SetValue
		classList.SetValue("new token list")
		if element.GetAttribute("class") != "new token list" {
			t.Errorf("expected class attribute to be 'new token list', got '%s'", element.GetAttribute("class"))
		}
		if classList.Value() != "new token list" {
			t.Errorf("expected Value() to be 'new token list', got '%s'", classList.Value())
		}
	})

	t.Run("String", func(t *testing.T) {
		classList := NewDOMTokenList(element, "class")
		element.SetAttribute("class", "foo bar baz")

		if classList.String() != "foo bar baz" {
			t.Errorf("expected String() to be 'foo bar baz', got '%s'", classList.String())
		}
	})

	t.Run("Update Steps - Empty Attribute and Empty Token Set", func(t *testing.T) {
		// Use a fresh element to ensure no class attribute is present
		freshElement := NewElement("div", doc)
		classList := NewDOMTokenList(freshElement, "class")

		// This should not set the attribute per spec
		classList.runUpdateSteps()
		if freshElement.HasAttribute("class") {
			t.Error("expected class attribute to not be set when both attribute and token set are empty")
		}
	})

	t.Run("Ordered Set Parser", func(t *testing.T) {
		// Test with various whitespace characters
		tokens := parseOrderedSet("  foo\t\tbar\n\nbaz\r\rfiz\f\fbuz  ")
		expected := []string{"foo", "bar", "baz", "fiz", "buz"}
		if len(tokens) != len(expected) {
			t.Errorf("expected %d tokens, got %d", len(expected), len(tokens))
		}
		for i, token := range tokens {
			if token != expected[i] {
				t.Errorf("expected token[%d] to be '%s', got '%s'", i, expected[i], token)
			}
		}

		// Test deduplication while preserving order
		tokens = parseOrderedSet("foo bar foo baz bar")
		expected = []string{"foo", "bar", "baz"}
		if len(tokens) != len(expected) {
			t.Errorf("expected %d unique tokens, got %d", len(expected), len(tokens))
		}
		for i, token := range tokens {
			if token != expected[i] {
				t.Errorf("expected token[%d] to be '%s', got '%s'", i, expected[i], token)
			}
		}
	})

	t.Run("Serialize Steps", func(t *testing.T) {
		classList := NewDOMTokenList(element, "class")
		element.SetAttribute("class", "foo bar baz")

		serialized := classList.runSerializeSteps()
		if serialized != "foo bar baz" {
			t.Errorf("expected serialize steps to return 'foo bar baz', got '%s'", serialized)
		}
	})

	t.Run("Live Updates", func(t *testing.T) {
		classList := NewDOMTokenList(element, "class")
		element.SetAttribute("class", "foo bar")

		// Verify initial state
		if classList.Length() != 2 {
			t.Errorf("expected length 2, got %d", classList.Length())
		}

		// Modify attribute directly
		element.SetAttribute("class", "foo bar baz qux")

		// DOMTokenList should reflect the change
		if classList.Length() != 4 {
			t.Errorf("expected length 4 after direct attribute modification, got %d", classList.Length())
		}
		if !classList.Contains("baz") || !classList.Contains("qux") {
			t.Error("expected DOMTokenList to contain newly added tokens")
		}
	})
}
