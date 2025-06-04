package js

import (
	"testing"

	"github.com/dop251/goja"
	"github.com/dpemmons/DOMulator/internal/dom"
)

func TestDOMExceptionJavaScriptIntegration(t *testing.T) {
	vm := goja.New()
	document := dom.NewDocument()
	bindings := NewDOMBindings(vm, document)
	bindings.SetupBrowserAPIs()

	tests := []struct {
		name     string
		script   string
		expected interface{}
	}{
		{
			name:     "DOMException constructor exists",
			script:   "typeof DOMException",
			expected: "function",
		},
		{
			name:     "DOMException with default parameters",
			script:   "var e = new DOMException(); e.name",
			expected: "Error",
		},
		{
			name:     "DOMException with message only",
			script:   "var e = new DOMException('Test message'); e.message",
			expected: "Test message",
		},
		{
			name:     "DOMException with message and name",
			script:   "var e = new DOMException('Test message', 'TestError'); e.name",
			expected: "TestError",
		},
		{
			name:     "DOMException code for HierarchyRequestError",
			script:   "var e = new DOMException('Test', 'HierarchyRequestError'); e.code",
			expected: int64(3),
		},
		{
			name:     "DOMException code for unknown error",
			script:   "var e = new DOMException('Test', 'UnknownError'); e.code",
			expected: int64(0),
		},
		{
			name:     "DOMException toString method",
			script:   "var e = new DOMException('Test message', 'TestError'); e.toString()",
			expected: "TestError: Test message",
		},
		{
			name:     "DOMException toString without message",
			script:   "var e = new DOMException('', 'TestError'); e.toString()",
			expected: "TestError",
		},
		{
			name:     "DOMException INDEX_SIZE_ERR constant",
			script:   "DOMException.INDEX_SIZE_ERR",
			expected: int64(1),
		},
		{
			name:     "DOMException HIERARCHY_REQUEST_ERR constant",
			script:   "DOMException.HIERARCHY_REQUEST_ERR",
			expected: int64(3),
		},
		{
			name:     "DOMException NOT_FOUND_ERR constant",
			script:   "DOMException.NOT_FOUND_ERR",
			expected: int64(8),
		},
		{
			name:     "DOMException SYNTAX_ERR constant",
			script:   "DOMException.SYNTAX_ERR",
			expected: int64(12),
		},
		{
			name:     "DOMException NAMESPACE_ERR constant",
			script:   "DOMException.NAMESPACE_ERR",
			expected: int64(14),
		},
		{
			name:     "DOMException ABORT_ERR constant",
			script:   "DOMException.ABORT_ERR",
			expected: int64(20),
		},
		{
			name:     "DOMException DATA_CLONE_ERR constant",
			script:   "DOMException.DATA_CLONE_ERR",
			expected: int64(25),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := vm.RunString(tt.script)
			if err != nil {
				t.Fatalf("Script execution failed: %v", err)
			}

			var actual interface{}
			if result.ExportType().Kind().String() == "string" {
				actual = result.String()
			} else {
				actual = result.Export()
			}

			if actual != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, actual)
			}
		})
	}
}

func TestDOMExceptionErrorHandling(t *testing.T) {
	vm := goja.New()
	document := dom.NewDocument()
	bindings := NewDOMBindings(vm, document)
	bindings.SetupBrowserAPIs()

	// Test try/catch with DOMException
	script := `
		try {
			var e = new DOMException('Test error', 'TestError');
			throw e;
		} catch (err) {
			if (err.name === 'TestError' && err.message === 'Test error') {
				'caught';
			} else {
				'failed';
			}
		}
	`

	result, err := vm.RunString(script)
	if err != nil {
		t.Fatalf("Script execution failed: %v", err)
	}

	if result.String() != "caught" {
		t.Errorf("Expected 'caught', got %v", result.String())
	}
}

func TestDOMExceptionJSPropertyAccess(t *testing.T) {
	vm := goja.New()
	document := dom.NewDocument()
	bindings := NewDOMBindings(vm, document)
	bindings.SetupBrowserAPIs()

	// Test property enumeration and access
	script := `
		var e = new DOMException('Test message', 'HierarchyRequestError');
		var props = {
			name: e.name,
			message: e.message,
			code: e.code,
			toString: typeof e.toString
		};
		JSON.stringify(props);
	`

	result, err := vm.RunString(script)
	if err != nil {
		t.Fatalf("Script execution failed: %v", err)
	}

	expected := `{"name":"HierarchyRequestError","message":"Test message","code":3,"toString":"function"}`
	if result.String() != expected {
		t.Errorf("Expected %s, got %s", expected, result.String())
	}
}

func TestDOMExceptionValidNames(t *testing.T) {
	vm := goja.New()
	document := dom.NewDocument()
	bindings := NewDOMBindings(vm, document)
	bindings.SetupBrowserAPIs()

	// Test all standard DOMException names have correct codes
	exceptions := map[string]int{
		"IndexSizeError":             1,
		"HierarchyRequestError":      3,
		"WrongDocumentError":         4,
		"InvalidCharacterError":      5,
		"NoModificationAllowedError": 7,
		"NotFoundError":              8,
		"NotSupportedError":          9,
		"InUseAttributeError":        10,
		"InvalidStateError":          11,
		"SyntaxError":                12,
		"InvalidModificationError":   13,
		"NamespaceError":             14,
		"InvalidAccessError":         15,
		"SecurityError":              18,
		"NetworkError":               19,
		"AbortError":                 20,
		"URLMismatchError":           21,
		"QuotaExceededError":         22,
		"TimeoutError":               23,
		"InvalidNodeTypeError":       24,
		"DataCloneError":             25,
	}

	for name, expectedCode := range exceptions {
		t.Run(name, func(t *testing.T) {
			script := `var e = new DOMException('test', '` + name + `'); e.code`
			result, err := vm.RunString(script)
			if err != nil {
				t.Fatalf("Script execution failed: %v", err)
			}

			code := int(result.ToInteger())
			if code != expectedCode {
				t.Errorf("Expected code %d for %s, got %d", expectedCode, name, code)
			}
		})
	}
}
