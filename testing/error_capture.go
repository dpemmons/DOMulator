package testing

import (
	"strings"
	"testing"

	"github.com/dpemmons/DOMulator/internal/js"
)

// ErrorCapture provides a testing helper to capture and assert on JavaScript errors
type ErrorCapture struct {
	errors []js.JavaScriptError
	t      testing.TB // Store test context for assertions
}

// NewErrorCapture creates a new error capture instance
func NewErrorCapture(t testing.TB) *ErrorCapture {
	return &ErrorCapture{
		errors: make([]js.JavaScriptError, 0),
		t:      t,
	}
}

// Callback is the callback function for the runtime error handler
func (e *ErrorCapture) Callback(err *js.JavaScriptError) {
	e.errors = append(e.errors, *err)
}

// AssertErrorType checks that at least one error of the expected type was captured
func (e *ErrorCapture) AssertErrorType(expectedType string) *ErrorCapture {
	found := false
	for _, err := range e.errors {
		if err.Type == expectedType {
			found = true
			break
		}
	}

	if !found {
		e.t.Errorf("Expected JavaScript error of type '%s', but none found", expectedType)
		if len(e.errors) > 0 {
			e.t.Errorf("Captured errors:")
			for i, err := range e.errors {
				e.t.Errorf("  %d: %s: %s", i+1, err.Type, err.Message)
			}
		}
	}
	return e
}

// AssertErrorMessage checks that at least one error contains the expected message text
func (e *ErrorCapture) AssertErrorMessage(expectedMessage string) *ErrorCapture {
	found := false
	for _, err := range e.errors {
		if strings.Contains(strings.ToLower(err.Message), strings.ToLower(expectedMessage)) {
			found = true
			break
		}
	}

	if !found {
		e.t.Errorf("Expected JavaScript error containing message '%s', but none found", expectedMessage)
		if len(e.errors) > 0 {
			e.t.Errorf("Captured error messages:")
			for i, err := range e.errors {
				e.t.Errorf("  %d: %s", i+1, err.Message)
			}
		}
	}
	return e
}

// AssertErrorCount checks the number of captured errors
func (e *ErrorCapture) AssertErrorCount(expected int) *ErrorCapture {
	actual := len(e.errors)
	if actual != expected {
		e.t.Errorf("Expected %d JavaScript errors, got %d", expected, actual)
		if actual > 0 {
			e.t.Errorf("Captured errors:")
			for i, err := range e.errors {
				e.t.Errorf("  %d: %s: %s", i+1, err.Type, err.Message)
			}
		}
	}
	return e
}

// AssertNoErrors checks that no JavaScript errors were captured
func (e *ErrorCapture) AssertNoErrors() *ErrorCapture {
	if count := len(e.errors); count > 0 {
		e.t.Errorf("Expected no JavaScript errors, but found %d", count)
		for i, err := range e.errors {
			e.t.Errorf("  %d: %s: %s", i+1, err.Type, err.Message)
		}
	}
	return e
}

// AssertTypeError checks for a TypeError
func (e *ErrorCapture) AssertTypeError() *ErrorCapture {
	return e.AssertErrorType("TypeError")
}

// AssertReferenceError checks for a ReferenceError
func (e *ErrorCapture) AssertReferenceError() *ErrorCapture {
	return e.AssertErrorType("ReferenceError")
}

// AssertSyntaxError checks for a SyntaxError
func (e *ErrorCapture) AssertSyntaxError() *ErrorCapture {
	return e.AssertErrorType("SyntaxError")
}

// GetErrors returns all captured errors
func (e *ErrorCapture) GetErrors() []js.JavaScriptError {
	return e.errors
}

// GetLastError returns the most recently captured error, or nil if none
func (e *ErrorCapture) GetLastError() *js.JavaScriptError {
	if len(e.errors) == 0 {
		return nil
	}
	return &e.errors[len(e.errors)-1]
}

// Clear removes all captured errors
func (e *ErrorCapture) Clear() *ErrorCapture {
	e.errors = e.errors[:0]
	return e
}

// Count returns the total number of captured errors
func (e *ErrorCapture) Count() int {
	return len(e.errors)
}

// HasErrorType checks if any error of the specified type was captured
func (e *ErrorCapture) HasErrorType(errorType string) bool {
	for _, err := range e.errors {
		if err.Type == errorType {
			return true
		}
	}
	return false
}

// HasErrorMessage checks if any error contains the specified message
func (e *ErrorCapture) HasErrorMessage(message string) bool {
	for _, err := range e.errors {
		if strings.Contains(strings.ToLower(err.Message), strings.ToLower(message)) {
			return true
		}
	}
	return false
}
