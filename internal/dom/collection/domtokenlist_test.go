package collection

import (
	"testing"
)

// mockElement is a simple implementation of ElementInterface for testing.
type mockElement struct {
	attributes map[string]string
}

func newMockElement() *mockElement {
	return &mockElement{
		attributes: make(map[string]string),
	}
}

func (e *mockElement) GetAttribute(name string) string {
	return e.attributes[name]
}

func (e *mockElement) SetAttribute(name, value string) {
	e.attributes[name] = value
}

func (e *mockElement) HasAttribute(name string) bool {
	_, exists := e.attributes[name]
	return exists
}

func TestDOMTokenList_Length(t *testing.T) {
	tests := []struct {
		name     string
		initial  string
		expected int
	}{
		{"empty", "", 0},
		{"single token", "class1", 1},
		{"multiple tokens", "class1 class2 class3", 3},
		{"with extra spaces", "  class1   class2  class3  ", 3},
		{"with duplicates", "class1 class2 class1 class3", 3}, // Should deduplicate
		{"whitespace only", "   ", 0},
		{"mixed whitespace", "class1\t\nclass2\r\fclass3", 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			element := newMockElement()
			element.SetAttribute("class", tt.initial)

			dtl := NewDOMTokenList(element, "class")
			if got := dtl.Length(); got != tt.expected {
				t.Errorf("Length() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestDOMTokenList_Item(t *testing.T) {
	element := newMockElement()
	element.SetAttribute("class", "class1 class2 class3")

	dtl := NewDOMTokenList(element, "class")

	tests := []struct {
		index    int
		expected string
	}{
		{0, "class1"},
		{1, "class2"},
		{2, "class3"},
		{3, ""},   // Out of bounds
		{-1, ""},  // Negative index
		{100, ""}, // Way out of bounds
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := dtl.Item(tt.index); got != tt.expected {
				t.Errorf("Item(%d) = %v, want %v", tt.index, got, tt.expected)
			}
		})
	}
}

func TestDOMTokenList_Contains(t *testing.T) {
	element := newMockElement()
	element.SetAttribute("class", "class1 class2 class3")

	dtl := NewDOMTokenList(element, "class")

	tests := []struct {
		token    string
		expected bool
	}{
		{"class1", true},
		{"class2", true},
		{"class3", true},
		{"class4", false},
		{"", false},         // Empty token
		{"class 1", false},  // Token with space
		{"class\t1", false}, // Token with tab
	}

	for _, tt := range tests {
		t.Run(tt.token, func(t *testing.T) {
			if got := dtl.Contains(tt.token); got != tt.expected {
				t.Errorf("Contains(%q) = %v, want %v", tt.token, got, tt.expected)
			}
		})
	}
}

func TestDOMTokenList_Add(t *testing.T) {
	tests := []struct {
		name           string
		initial        string
		tokensToAdd    []string
		expectedResult string
		expectedError  bool
	}{
		{
			name:           "add to empty",
			initial:        "",
			tokensToAdd:    []string{"class1"},
			expectedResult: "class1",
			expectedError:  false,
		},
		{
			name:           "add to existing",
			initial:        "class1",
			tokensToAdd:    []string{"class2"},
			expectedResult: "class1 class2",
			expectedError:  false,
		},
		{
			name:           "add duplicate",
			initial:        "class1 class2",
			tokensToAdd:    []string{"class1"},
			expectedResult: "class1 class2",
			expectedError:  false,
		},
		{
			name:           "add multiple",
			initial:        "class1",
			tokensToAdd:    []string{"class2", "class3"},
			expectedResult: "class1 class2 class3",
			expectedError:  false,
		},
		{
			name:           "add multiple with duplicates",
			initial:        "class1 class2",
			tokensToAdd:    []string{"class2", "class3", "class1"},
			expectedResult: "class1 class2 class3",
			expectedError:  false,
		},
		{
			name:           "add empty token",
			initial:        "class1",
			tokensToAdd:    []string{""},
			expectedResult: "class1",
			expectedError:  true,
		},
		{
			name:           "add token with space",
			initial:        "class1",
			tokensToAdd:    []string{"class 2"},
			expectedResult: "class1",
			expectedError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			element := newMockElement()
			element.SetAttribute("class", tt.initial)

			dtl := NewDOMTokenList(element, "class")
			err := dtl.Add(tt.tokensToAdd...)

			if tt.expectedError && err == nil {
				t.Errorf("Add() expected error but got none")
			}
			if !tt.expectedError && err != nil {
				t.Errorf("Add() unexpected error: %v", err)
			}

			if !tt.expectedError {
				result := element.GetAttribute("class")
				if result != tt.expectedResult {
					t.Errorf("Add() result = %q, want %q", result, tt.expectedResult)
				}
			}
		})
	}
}

func TestDOMTokenList_Remove(t *testing.T) {
	tests := []struct {
		name           string
		initial        string
		tokensToRemove []string
		expectedResult string
	}{
		{
			name:           "remove from empty",
			initial:        "",
			tokensToRemove: []string{"class1"},
			expectedResult: "",
		},
		{
			name:           "remove existing",
			initial:        "class1 class2 class3",
			tokensToRemove: []string{"class2"},
			expectedResult: "class1 class3",
		},
		{
			name:           "remove non-existing",
			initial:        "class1 class2",
			tokensToRemove: []string{"class3"},
			expectedResult: "class1 class2",
		},
		{
			name:           "remove multiple",
			initial:        "class1 class2 class3 class4",
			tokensToRemove: []string{"class2", "class4"},
			expectedResult: "class1 class3",
		},
		{
			name:           "remove all",
			initial:        "class1 class2",
			tokensToRemove: []string{"class1", "class2"},
			expectedResult: "",
		},
		{
			name:           "remove invalid token silently ignored",
			initial:        "class1 class2",
			tokensToRemove: []string{"class 3", "class1"},
			expectedResult: "class2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			element := newMockElement()
			element.SetAttribute("class", tt.initial)

			dtl := NewDOMTokenList(element, "class")
			dtl.Remove(tt.tokensToRemove...)

			result := element.GetAttribute("class")
			if result != tt.expectedResult {
				t.Errorf("Remove() result = %q, want %q", result, tt.expectedResult)
			}
		})
	}
}

func TestDOMTokenList_Toggle(t *testing.T) {
	tests := []struct {
		name           string
		initial        string
		token          string
		force          *bool
		expectedResult string
		expectedReturn bool
	}{
		{
			name:           "toggle add",
			initial:        "class1",
			token:          "class2",
			force:          nil,
			expectedResult: "class1 class2",
			expectedReturn: true,
		},
		{
			name:           "toggle remove",
			initial:        "class1 class2",
			token:          "class2",
			force:          nil,
			expectedResult: "class1",
			expectedReturn: false,
		},
		{
			name:           "force add existing",
			initial:        "class1 class2",
			token:          "class2",
			force:          &[]bool{true}[0],
			expectedResult: "class1 class2",
			expectedReturn: true,
		},
		{
			name:           "force add non-existing",
			initial:        "class1",
			token:          "class2",
			force:          &[]bool{true}[0],
			expectedResult: "class1 class2",
			expectedReturn: true,
		},
		{
			name:           "force remove existing",
			initial:        "class1 class2",
			token:          "class2",
			force:          &[]bool{false}[0],
			expectedResult: "class1",
			expectedReturn: false,
		},
		{
			name:           "force remove non-existing",
			initial:        "class1",
			token:          "class2",
			force:          &[]bool{false}[0],
			expectedResult: "class1",
			expectedReturn: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			element := newMockElement()
			element.SetAttribute("class", tt.initial)

			dtl := NewDOMTokenList(element, "class")
			result := dtl.Toggle(tt.token, tt.force)

			if result != tt.expectedReturn {
				t.Errorf("Toggle() return = %v, want %v", result, tt.expectedReturn)
			}

			actualResult := element.GetAttribute("class")
			if actualResult != tt.expectedResult {
				t.Errorf("Toggle() result = %q, want %q", actualResult, tt.expectedResult)
			}
		})
	}
}

func TestDOMTokenList_Replace(t *testing.T) {
	tests := []struct {
		name           string
		initial        string
		oldToken       string
		newToken       string
		expectedResult string
		expectedReturn bool
	}{
		{
			name:           "replace existing",
			initial:        "class1 class2 class3",
			oldToken:       "class2",
			newToken:       "newclass",
			expectedResult: "class1 newclass class3",
			expectedReturn: true,
		},
		{
			name:           "replace non-existing",
			initial:        "class1 class2",
			oldToken:       "class3",
			newToken:       "newclass",
			expectedResult: "class1 class2",
			expectedReturn: false,
		},
		{
			name:           "replace with existing token",
			initial:        "class1 class2 class3",
			oldToken:       "class2",
			newToken:       "class3",
			expectedResult: "class1 class3",
			expectedReturn: true,
		},
		{
			name:           "replace first occurrence",
			initial:        "class1 class2 class1",
			oldToken:       "class1",
			newToken:       "newclass",
			expectedResult: "newclass class2", // DOMTokenList deduplicates, so "class1 class2 class1" becomes "class1 class2"
			expectedReturn: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			element := newMockElement()
			element.SetAttribute("class", tt.initial)

			dtl := NewDOMTokenList(element, "class")
			result := dtl.Replace(tt.oldToken, tt.newToken)

			if result != tt.expectedReturn {
				t.Errorf("Replace() return = %v, want %v", result, tt.expectedReturn)
			}

			actualResult := element.GetAttribute("class")
			if actualResult != tt.expectedResult {
				t.Errorf("Replace() result = %q, want %q", actualResult, tt.expectedResult)
			}
		})
	}
}

func TestDOMTokenList_Supports(t *testing.T) {
	element := newMockElement()
	dtl := NewDOMTokenList(element, "class")

	tests := []struct {
		token    string
		expected bool
	}{
		{"validtoken", true},
		{"valid-token", true},
		{"valid_token", true},
		{"", false},              // Empty token
		{"invalid token", false}, // Token with space
		{"invalid\ttok", false},  // Token with tab
	}

	for _, tt := range tests {
		t.Run(tt.token, func(t *testing.T) {
			if got := dtl.Supports(tt.token); got != tt.expected {
				t.Errorf("Supports(%q) = %v, want %v", tt.token, got, tt.expected)
			}
		})
	}
}

func TestDOMTokenList_String(t *testing.T) {
	tests := []struct {
		name     string
		initial  string
		expected string
	}{
		{"empty", "", ""},
		{"single", "class1", "class1"},
		{"multiple", "class1 class2 class3", "class1 class2 class3"},
		{"with duplicates", "class1 class2 class1", "class1 class2"}, // Should deduplicate
		{"with extra spaces", "  class1   class2  ", "class1 class2"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			element := newMockElement()
			element.SetAttribute("class", tt.initial)

			dtl := NewDOMTokenList(element, "class")
			if got := dtl.String(); got != tt.expected {
				t.Errorf("String() = %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestDOMTokenList_SetValue(t *testing.T) {
	element := newMockElement()
	element.SetAttribute("class", "class1 class2")

	dtl := NewDOMTokenList(element, "class")

	// Verify initial state
	if dtl.String() != "class1 class2" {
		t.Errorf("Initial String() = %q, want %q", dtl.String(), "class1 class2")
	}

	// Set new value
	dtl.SetValue("newclass1 newclass2 newclass3")

	// Verify the value was set
	if got := element.GetAttribute("class"); got != "newclass1 newclass2 newclass3" {
		t.Errorf("SetValue() attribute = %q, want %q", got, "newclass1 newclass2 newclass3")
	}

	// Verify the token list reflects the change
	if got := dtl.String(); got != "newclass1 newclass2 newclass3" {
		t.Errorf("SetValue() String() = %q, want %q", got, "newclass1 newclass2 newclass3")
	}
}

func TestDOMTokenList_LiveUpdates(t *testing.T) {
	element := newMockElement()
	element.SetAttribute("class", "class1 class2")

	dtl := NewDOMTokenList(element, "class")

	// Verify initial state
	if dtl.Length() != 2 {
		t.Errorf("Initial Length() = %d, want 2", dtl.Length())
	}

	// Modify the attribute directly
	element.SetAttribute("class", "newclass1 newclass2 newclass3")

	// Verify the token list reflects the change (live updates)
	if dtl.Length() != 3 {
		t.Errorf("After attribute change Length() = %d, want 3", dtl.Length())
	}

	if !dtl.Contains("newclass1") {
		t.Errorf("After attribute change Contains(newclass1) = false, want true")
	}

	if dtl.Contains("class1") {
		t.Errorf("After attribute change Contains(class1) = true, want false")
	}
}

func TestParseOrderedSet(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{"", []string{}},
		{"class1", []string{"class1"}},
		{"class1 class2", []string{"class1", "class2"}},
		{"  class1   class2  ", []string{"class1", "class2"}},
		{"class1 class2 class1", []string{"class1", "class2"}}, // Deduplicate
		{"class1\tclass2\nclass3", []string{"class1", "class2", "class3"}},
		{"   ", []string{}}, // Whitespace only
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := parseOrderedSet(tt.input)
			if len(got) != len(tt.expected) {
				t.Errorf("parseOrderedSet(%q) length = %d, want %d", tt.input, len(got), len(tt.expected))
				return
			}
			for i, token := range got {
				if token != tt.expected[i] {
					t.Errorf("parseOrderedSet(%q)[%d] = %q, want %q", tt.input, i, token, tt.expected[i])
				}
			}
		})
	}
}

func TestValidateToken(t *testing.T) {
	tests := []struct {
		token     string
		expectErr bool
	}{
		{"validtoken", false},
		{"valid-token", false},
		{"valid_token", false},
		{"", true},                      // Empty
		{"token with space", true},      // Space
		{"token\twith\ttab", true},      // Tab
		{"token\nwith\nnewline", true},  // Newline
		{"token\rwith\rcarriage", true}, // Carriage return
		{"token\fwith\fform", true},     // Form feed
	}

	for _, tt := range tests {
		t.Run(tt.token, func(t *testing.T) {
			err := validateToken(tt.token)
			if tt.expectErr && err == nil {
				t.Errorf("validateToken(%q) expected error but got none", tt.token)
			}
			if !tt.expectErr && err != nil {
				t.Errorf("validateToken(%q) unexpected error: %v", tt.token, err)
			}
		})
	}
}
