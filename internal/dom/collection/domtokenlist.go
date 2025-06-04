package collection

import (
	"errors"
	"strings"
)

// DOMTokenList represents a set of space-separated tokens.
// It provides a live, ordered set of unique tokens from a string attribute.
type DOMTokenList struct {
	element   ElementInterface // Element that owns this token list
	attribute string           // Attribute name (e.g., "class", "rel")
	tokens    []string         // Cached ordered set of unique tokens
	version   uint64           // Version for cache invalidation
	lastValue string           // Last seen attribute value for change detection
}

// ElementInterface defines the interface for elements that can have token lists.
// This avoids circular dependency issues.
type ElementInterface interface {
	GetAttribute(name string) string
	SetAttribute(name, value string)
	HasAttribute(name string) bool
}

// NewDOMTokenList creates a new DOMTokenList for the specified element and attribute.
func NewDOMTokenList(element ElementInterface, attribute string) *DOMTokenList {
	dtl := &DOMTokenList{
		element:   element,
		attribute: attribute,
		tokens:    nil, // Will be populated on first access
		version:   0,
	}
	return dtl
}

// Length returns the number of tokens in the list.
func (dtl *DOMTokenList) Length() int {
	dtl.updateIfNeeded()
	return len(dtl.tokens)
}

// Item returns the token at the specified index, or empty string if out of bounds.
func (dtl *DOMTokenList) Item(index int) string {
	dtl.updateIfNeeded()
	if index < 0 || index >= len(dtl.tokens) {
		return ""
	}
	return dtl.tokens[index]
}

// Contains returns true if the token list contains the specified token.
func (dtl *DOMTokenList) Contains(token string) bool {
	if err := validateToken(token); err != nil {
		return false
	}

	dtl.updateIfNeeded()
	for _, t := range dtl.tokens {
		if t == token {
			return true
		}
	}
	return false
}

// Add adds one or more tokens to the list.
// Tokens that are already present are ignored.
// Returns an error if any token is invalid.
func (dtl *DOMTokenList) Add(tokens ...string) error {
	// Validate all tokens first
	for _, token := range tokens {
		if err := validateToken(token); err != nil {
			return err
		}
	}

	dtl.updateIfNeeded()

	// Add tokens that don't already exist
	modified := false
	for _, token := range tokens {
		if !dtl.containsToken(token) {
			dtl.tokens = append(dtl.tokens, token)
			modified = true
		}
	}

	if modified {
		dtl.updateAttribute()
	}

	return nil
}

// Remove removes one or more tokens from the list.
// Tokens that are not present are ignored.
func (dtl *DOMTokenList) Remove(tokens ...string) {
	// Validate tokens (silently ignore invalid ones per spec)
	validTokens := make([]string, 0, len(tokens))
	for _, token := range tokens {
		if validateToken(token) == nil {
			validTokens = append(validTokens, token)
		}
	}

	if len(validTokens) == 0 {
		return
	}

	dtl.updateIfNeeded()

	// Remove tokens that exist
	modified := false
	newTokens := make([]string, 0, len(dtl.tokens))

	for _, existing := range dtl.tokens {
		shouldRemove := false
		for _, toRemove := range validTokens {
			if existing == toRemove {
				shouldRemove = true
				modified = true
				break
			}
		}
		if !shouldRemove {
			newTokens = append(newTokens, existing)
		}
	}

	if modified {
		dtl.tokens = newTokens
		dtl.updateAttribute()
	}
}

// Toggle adds the token if it's not present, or removes it if it is present.
// If force is provided, it forces the token to be added (true) or removed (false).
// Returns true if the token is present after the operation.
func (dtl *DOMTokenList) Toggle(token string, force *bool) bool {
	if err := validateToken(token); err != nil {
		return false
	}

	dtl.updateIfNeeded()

	exists := dtl.containsToken(token)

	var shouldAdd bool
	if force != nil {
		shouldAdd = *force
	} else {
		shouldAdd = !exists
	}

	if shouldAdd && !exists {
		dtl.tokens = append(dtl.tokens, token)
		dtl.updateAttribute()
		return true
	} else if !shouldAdd && exists {
		dtl.Remove(token)
		return false
	}

	return exists
}

// Replace replaces an existing token with a new token.
// Returns true if the replacement was successful.
func (dtl *DOMTokenList) Replace(oldToken, newToken string) bool {
	if err := validateToken(oldToken); err != nil {
		return false
	}
	if err := validateToken(newToken); err != nil {
		return false
	}

	dtl.updateIfNeeded()

	// Find the old token
	for i, token := range dtl.tokens {
		if token == oldToken {
			// Check if new token already exists elsewhere
			for j, existing := range dtl.tokens {
				if existing == newToken && j != i {
					// Remove the old token since new token already exists
					dtl.tokens = append(dtl.tokens[:i], dtl.tokens[i+1:]...)
					dtl.updateAttribute()
					return true
				}
			}
			// Replace the old token
			dtl.tokens[i] = newToken
			dtl.updateAttribute()
			return true
		}
	}

	return false
}

// Supports checks if the token is supported for this token list.
// For now, this always returns true as we don't have specific validation rules.
func (dtl *DOMTokenList) Supports(token string) bool {
	// The supports method is primarily used for specific token lists
	// like rel attribute where certain values have special meaning.
	// For now, we'll return true for any valid token format.
	return validateToken(token) == nil
}

// String returns the string representation of the token list.
func (dtl *DOMTokenList) String() string {
	dtl.updateIfNeeded()
	return strings.Join(dtl.tokens, " ")
}

// Value returns the current string value (same as String()).
func (dtl *DOMTokenList) Value() string {
	return dtl.String()
}

// SetValue sets the entire token list from a string value.
func (dtl *DOMTokenList) SetValue(value string) {
	dtl.element.SetAttribute(dtl.attribute, value)
	dtl.invalidate()
}

// updateIfNeeded checks if the underlying attribute has changed and updates the cache.
func (dtl *DOMTokenList) updateIfNeeded() {
	currentValue := dtl.element.GetAttribute(dtl.attribute)
	if currentValue != dtl.lastValue || dtl.tokens == nil {
		dtl.lastValue = currentValue
		dtl.tokens = parseOrderedSet(currentValue)
		dtl.version++
	}
}

// updateAttribute updates the underlying attribute with the current token list.
func (dtl *DOMTokenList) updateAttribute() {
	value := strings.Join(dtl.tokens, " ")
	dtl.element.SetAttribute(dtl.attribute, value)
	dtl.lastValue = value
	dtl.version++
}

// containsToken is an internal helper to check if a token exists.
func (dtl *DOMTokenList) containsToken(token string) bool {
	for _, t := range dtl.tokens {
		if t == token {
			return true
		}
	}
	return false
}

// invalidate forces the cache to be refreshed on next access.
func (dtl *DOMTokenList) invalidate() {
	dtl.tokens = nil
	dtl.lastValue = ""
}

// validateToken validates that a token conforms to the DOM specification.
// Tokens must not be empty and must not contain ASCII whitespace.
func validateToken(token string) error {
	if token == "" {
		return errors.New("token must not be empty")
	}

	// Check for ASCII whitespace characters
	for _, r := range token {
		if r == ' ' || r == '\t' || r == '\n' || r == '\r' || r == '\f' {
			return errors.New("token must not contain ASCII whitespace")
		}
	}

	return nil
}

// parseOrderedSet parses a space-separated string into an ordered set of unique tokens.
// This implements the "ordered set parser" algorithm from the DOM specification.
func parseOrderedSet(input string) []string {
	if input == "" {
		return []string{}
	}

	// Split by ASCII whitespace and deduplicate while preserving order
	var tokens []string
	seen := make(map[string]bool)

	// Split by any ASCII whitespace
	parts := strings.FieldsFunc(input, func(r rune) bool {
		return r == ' ' || r == '\t' || r == '\n' || r == '\r' || r == '\f'
	})

	for _, part := range parts {
		if part != "" && !seen[part] {
			tokens = append(tokens, part)
			seen[part] = true
		}
	}

	return tokens
}

// serializeOrderedSet serializes a list of tokens to a space-separated string.
// This implements the "ordered set serializer" algorithm from the DOM specification.
func serializeOrderedSet(tokens []string) string {
	return strings.Join(tokens, " ")
}
