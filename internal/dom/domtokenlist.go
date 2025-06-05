package dom

import (
	"strings"
)

// DOMTokenList represents a set of space-separated tokens per WHATWG DOM Section 7.1.
// It provides a live, ordered set of unique tokens from a string attribute.
// https://dom.spec.whatwg.org/#interface-domtokenlist
type DOMTokenList struct {
	element       *Element // Element that owns this token list
	attributeName string   // Attribute's local name (e.g., "class", "rel")
	tokenSet      []string // Token set (cached ordered set of unique tokens)
	version       uint64   // Version for cache invalidation
	lastValue     string   // Last seen attribute value for change detection

	// supportedTokens defines the supported tokens for validation steps
	// nil means no supported tokens are defined
	supportedTokens map[string]bool
}

// NewDOMTokenList creates a new DOMTokenList for the specified element and attribute.
// This implements the DOMTokenList object creation steps from the specification.
func NewDOMTokenList(element *Element, attributeName string) *DOMTokenList {
	dtl := &DOMTokenList{
		element:       element,
		attributeName: attributeName,
		tokenSet:      nil, // Will be populated on first access
		version:       0,
	}

	// Run the attribute change steps for initial creation
	value := element.GetAttribute(attributeName)
	dtl.runAttributeChangeSteps(attributeName, "", value, value, "")

	return dtl
}

// NewDOMTokenListWithSupportedTokens creates a DOMTokenList with defined supported tokens.
func NewDOMTokenListWithSupportedTokens(element *Element, attributeName string, supportedTokens []string) *DOMTokenList {
	dtl := NewDOMTokenList(element, attributeName)
	if len(supportedTokens) > 0 {
		dtl.supportedTokens = make(map[string]bool)
		for _, token := range supportedTokens {
			dtl.supportedTokens[strings.ToLower(token)] = true
		}
	}
	return dtl
}

// Length getter - returns the number of tokens.
// The length attribute' getter must return this's token set's size.
func (dtl *DOMTokenList) Length() int {
	dtl.updateIfNeeded()
	return len(dtl.tokenSet)
}

// Item method - returns the token at the specified index.
// The item(index) method steps are:
// 1. If index is equal to or greater than this's token set's size, then return null.
// 2. Return this's token set[index].
func (dtl *DOMTokenList) Item(index int) string {
	dtl.updateIfNeeded()
	if index < 0 || index >= len(dtl.tokenSet) {
		return "" // Return empty string instead of null for Go compatibility
	}
	return dtl.tokenSet[index]
}

// Contains method - returns true if the token list contains the specified token.
// The contains(token) method steps are to return true if this's token set[token] exists; otherwise false.
func (dtl *DOMTokenList) Contains(token string) bool {
	// Note: The spec doesn't validate tokens for contains(), it just checks existence
	dtl.updateIfNeeded()
	for _, t := range dtl.tokenSet {
		if t == token {
			return true
		}
	}
	return false
}

// Add method - adds one or more tokens to the list per WHATWG DOM specification.
// The add(tokens…) method steps are:
//  1. For each token of tokens:
//     a. If token is the empty string, then throw a "SyntaxError" DOMException.
//     b. If token contains any ASCII whitespace, then throw an "InvalidCharacterError" DOMException.
//  2. For each token of tokens, append token to this's token set.
//  3. Run the update steps.
func (dtl *DOMTokenList) Add(tokens ...string) {
	// Step 1: Validate all tokens first per specification
	for _, token := range tokens {
		if token == "" {
			panic(NewSyntaxError("Token must not be empty"))
		}
		if containsASCIIWhitespace(token) {
			panic(NewInvalidCharacterError("Token must not contain ASCII whitespace"))
		}
	}

	dtl.updateIfNeeded()

	// Step 2: For each token of tokens, append token to this's token set
	modified := false
	for _, token := range tokens {
		if !dtl.containsToken(token) {
			dtl.tokenSet = append(dtl.tokenSet, token)
			modified = true
		}
	}

	// Step 3: Run the update steps
	if modified {
		dtl.runUpdateSteps()
	}
}

// Remove method - removes one or more tokens from the list per WHATWG DOM specification.
// The remove(tokens…) method steps are:
//  1. For each token of tokens:
//     a. If token is the empty string, then throw a "SyntaxError" DOMException.
//     b. If token contains any ASCII whitespace, then throw an "InvalidCharacterError" DOMException.
//  2. For each token in tokens, remove token from this's token set.
//  3. Run the update steps.
func (dtl *DOMTokenList) Remove(tokens ...string) {
	// Step 1: Validate all tokens first per specification
	for _, token := range tokens {
		if token == "" {
			panic(NewSyntaxError("Token must not be empty"))
		}
		if containsASCIIWhitespace(token) {
			panic(NewInvalidCharacterError("Token must not contain ASCII whitespace"))
		}
	}

	dtl.updateIfNeeded()

	// Step 2: For each token in tokens, remove token from this's token set
	modified := false
	newTokens := make([]string, 0, len(dtl.tokenSet))

	for _, existing := range dtl.tokenSet {
		shouldRemove := false
		for _, toRemove := range tokens {
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

	// Step 3: Run the update steps
	if modified {
		dtl.tokenSet = newTokens
		dtl.runUpdateSteps()
	}
}

// Toggle method - toggles a token per WHATWG DOM specification.
// The toggle(token, force) method steps are:
//  1. If token is the empty string, then throw a "SyntaxError" DOMException.
//  2. If token contains any ASCII whitespace, then throw an "InvalidCharacterError" DOMException.
//  3. If this's token set[token] exists:
//     a. If force is either not given or is false, then remove token from this's token set, run the update steps and return false.
//     b. Return true.
//  4. Otherwise, if force not given or is true, append token to this's token set, run the update steps, and return true.
//  5. Return false.
func (dtl *DOMTokenList) Toggle(token string, force *bool) bool {
	// Step 1: If token is the empty string, then throw a "SyntaxError" DOMException
	if token == "" {
		panic(NewSyntaxError("Token must not be empty"))
	}

	// Step 2: If token contains any ASCII whitespace, then throw an "InvalidCharacterError" DOMException
	if containsASCIIWhitespace(token) {
		panic(NewInvalidCharacterError("Token must not contain ASCII whitespace"))
	}

	dtl.updateIfNeeded()

	exists := dtl.containsToken(token)

	// Step 3: If this's token set[token] exists
	if exists {
		// Step 3a: If force is either not given or is false, then remove token from this's token set, run the update steps and return false
		if force == nil || !*force {
			dtl.Remove(token)
			return false
		}
		// Step 3b: Return true
		return true
	}

	// Step 4: Otherwise, if force not given or is true, append token to this's token set, run the update steps, and return true
	if force == nil || *force {
		dtl.Add(token)
		return true
	}

	// Step 5: Return false
	return false
}

// Replace method - replaces an existing token with a new token per WHATWG DOM specification.
// The replace(token, newToken) method steps are:
// 1. If either token or newToken is the empty string, then throw a "SyntaxError" DOMException.
// 2. If either token or newToken contains any ASCII whitespace, then throw an "InvalidCharacterError" DOMException.
// 3. If this's token set does not contain token, then return false.
// 4. Replace token in this's token set with newToken.
// 5. Run the update steps.
// 6. Return true.
func (dtl *DOMTokenList) Replace(token, newToken string) bool {
	// Step 1: If either token or newToken is the empty string, then throw a "SyntaxError" DOMException
	if token == "" || newToken == "" {
		panic(NewSyntaxError("Token must not be empty"))
	}

	// Step 2: If either token or newToken contains any ASCII whitespace, then throw an "InvalidCharacterError" DOMException
	if containsASCIIWhitespace(token) || containsASCIIWhitespace(newToken) {
		panic(NewInvalidCharacterError("Token must not contain ASCII whitespace"))
	}

	dtl.updateIfNeeded()

	// Step 3: If this's token set does not contain token, then return false
	tokenIndex := -1
	for i, t := range dtl.tokenSet {
		if t == token {
			tokenIndex = i
			break
		}
	}
	if tokenIndex == -1 {
		return false
	}

	// Step 4: Replace token in this's token set with newToken
	// Check if newToken already exists elsewhere in the set
	for i, existing := range dtl.tokenSet {
		if existing == newToken && i != tokenIndex {
			// Remove the old token since new token already exists
			dtl.tokenSet = append(dtl.tokenSet[:tokenIndex], dtl.tokenSet[tokenIndex+1:]...)
			dtl.runUpdateSteps()
			return true
		}
	}

	// Replace the old token with the new token
	dtl.tokenSet[tokenIndex] = newToken

	// Step 5: Run the update steps
	dtl.runUpdateSteps()

	// Step 6: Return true
	return true
}

// Supports method - returns true if token is in the associated attribute's supported tokens per WHATWG DOM specification.
// The supports(token) method steps are:
// 1. Let result be the return value of validation steps called with token.
// 2. Return result.
func (dtl *DOMTokenList) Supports(token string) bool {
	// Run validation steps
	return dtl.runValidationSteps(token)
}

// Value getter/setter - returns/sets the associated set as string per WHATWG DOM specification.
// The value attribute must return the result of running this's serialize steps.
// Setting the value attribute must set an attribute value for the associated element using associated attribute's local name and the given value.
func (dtl *DOMTokenList) Value() string {
	return dtl.runSerializeSteps()
}

// SetValue sets the entire token list from a string value.
func (dtl *DOMTokenList) SetValue(value string) {
	dtl.element.SetAttribute(dtl.attributeName, value)
	dtl.invalidate()
}

// String returns the string representation of the token list (same as Value()).
func (dtl *DOMTokenList) String() string {
	return dtl.Value()
}

// Internal implementation methods

// updateIfNeeded checks if the underlying attribute has changed and updates the cache.
func (dtl *DOMTokenList) updateIfNeeded() {
	currentValue := dtl.element.GetAttribute(dtl.attributeName)
	if currentValue != dtl.lastValue || dtl.tokenSet == nil {
		dtl.lastValue = currentValue
		dtl.tokenSet = parseOrderedSet(currentValue)
		dtl.version++
	}
}

// runUpdateSteps implements the DOMTokenList update steps from the specification:
// 1. If the associated element does not have an associated attribute and token set is empty, then return.
// 2. Set an attribute value for the associated element using associated attribute's local name and the result of running the ordered set serializer for token set.
func (dtl *DOMTokenList) runUpdateSteps() {
	// Step 1: If the associated element does not have an associated attribute and token set is empty, then return
	if !dtl.element.HasAttribute(dtl.attributeName) && len(dtl.tokenSet) == 0 {
		return
	}

	// Step 2: Set an attribute value using the ordered set serializer
	value := serializeOrderedSet(dtl.tokenSet)
	dtl.element.SetAttribute(dtl.attributeName, value)
	dtl.lastValue = value
	dtl.version++
}

// runSerializeSteps implements the DOMTokenList serialize steps from the specification:
// Return the result of running get an attribute value given the associated element and the associated attribute's local name.
func (dtl *DOMTokenList) runSerializeSteps() string {
	return dtl.element.GetAttribute(dtl.attributeName)
}

// runValidationSteps implements the DOMTokenList validation steps from the specification:
// 1. If the associated attribute's local name does not define supported tokens, throw a TypeError.
// 2. Let lowercase token be a copy of token, in ASCII lowercase.
// 3. If lowercase token is present in supported tokens, return true.
// 4. Return false.
func (dtl *DOMTokenList) runValidationSteps(token string) bool {
	// Step 1: If the associated attribute's local name does not define supported tokens, throw a TypeError
	if dtl.supportedTokens == nil {
		// In Go, we can't throw a TypeError directly, but we can panic with a message
		// For now, return true to maintain compatibility until we have a proper JavaScript integration
		return true
	}

	// Step 2: Let lowercase token be a copy of token, in ASCII lowercase
	lowercaseToken := strings.ToLower(token)

	// Step 3: If lowercase token is present in supported tokens, return true
	if dtl.supportedTokens[lowercaseToken] {
		return true
	}

	// Step 4: Return false
	return false
}

// runAttributeChangeSteps implements the DOMTokenList attribute change steps from the specification:
// 1. If localName is associated attribute's local name, namespace is null, and value is null, then empty token set.
// 2. Otherwise, if localName is associated attribute's local name, namespace is null, then set token set to value, parsed.
func (dtl *DOMTokenList) runAttributeChangeSteps(localName, namespace, oldValue, value, newValue string) {
	if localName != dtl.attributeName || namespace != "" {
		return
	}

	// Step 1: If value is null (empty in Go), then empty token set
	if value == "" {
		dtl.tokenSet = []string{}
	} else {
		// Step 2: Set token set to value, parsed
		dtl.tokenSet = parseOrderedSet(value)
	}

	dtl.lastValue = value
	dtl.version++
}

// containsToken is an internal helper to check if a token exists.
func (dtl *DOMTokenList) containsToken(token string) bool {
	for _, t := range dtl.tokenSet {
		if t == token {
			return true
		}
	}
	return false
}

// invalidate forces the cache to be refreshed on next access.
func (dtl *DOMTokenList) invalidate() {
	dtl.tokenSet = nil
	dtl.lastValue = ""
}

// containsASCIIWhitespace checks if a string contains any ASCII whitespace characters.
func containsASCIIWhitespace(s string) bool {
	for _, r := range s {
		if r == ' ' || r == '\t' || r == '\n' || r == '\r' || r == '\f' {
			return true
		}
	}
	return false
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
