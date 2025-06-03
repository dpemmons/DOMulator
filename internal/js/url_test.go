package js

import (
	"testing"

	"github.com/dop251/goja"
	"github.com/dpemmons/DOMulator/internal/dom"
)

func TestURLJavaScriptBindings(t *testing.T) {
	// Create runtime and bindings
	vm := goja.New()
	document := dom.NewDocument()
	bindings := NewDOMBindings(vm, document)
	bindings.SetupBrowserAPIs()

	tests := []struct {
		name   string
		script string
		want   interface{}
	}{
		{
			name:   "URL constructor basic",
			script: `new URL("https://example.com/path").href`,
			want:   "https://example.com/path",
		},
		{
			name:   "URL constructor with base",
			script: `new URL("path", "https://example.com").href`,
			want:   "https://example.com/path",
		},
		{
			name:   "URL protocol property",
			script: `new URL("https://example.com").protocol`,
			want:   "https:",
		},
		{
			name:   "URL hostname property",
			script: `new URL("https://example.com:8080").hostname`,
			want:   "example.com",
		},
		{
			name:   "URL port property",
			script: `new URL("https://example.com:8080").port`,
			want:   "8080",
		},
		{
			name:   "URL pathname property",
			script: `new URL("https://example.com/path/to/resource").pathname`,
			want:   "/path/to/resource",
		},
		{
			name:   "URL search property",
			script: `new URL("https://example.com?foo=bar").search`,
			want:   "?foo=bar",
		},
		{
			name:   "URL hash property",
			script: `new URL("https://example.com#section").hash`,
			want:   "#section",
		},
		{
			name:   "URL origin property",
			script: `new URL("https://example.com:8080").origin`,
			want:   "https://example.com:8080",
		},
		{
			name:   "URL toString method",
			script: `new URL("https://example.com/path").toString()`,
			want:   "https://example.com/path",
		},
		{
			name:   "URL toJSON method",
			script: `new URL("https://example.com/path").toJSON()`,
			want:   "https://example.com/path",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := vm.RunString(test.script)
			if err != nil {
				t.Fatalf("Script execution failed: %v", err)
			}

			actual := result.Export()
			if actual != test.want {
				t.Errorf("Expected %v, got %v", test.want, actual)
			}
		})
	}
}

func TestURLSearchParamsJavaScriptBindings(t *testing.T) {
	// Create runtime and bindings
	vm := goja.New()
	document := dom.NewDocument()
	bindings := NewDOMBindings(vm, document)
	bindings.SetupBrowserAPIs()

	tests := []struct {
		name   string
		script string
		want   interface{}
	}{
		{
			name:   "URLSearchParams constructor empty",
			script: `new URLSearchParams().toString()`,
			want:   "",
		},
		{
			name:   "URLSearchParams constructor with string",
			script: `new URLSearchParams("foo=bar&baz=qux").get("foo")`,
			want:   "bar",
		},
		{
			name:   "URLSearchParams append method",
			script: `var p = new URLSearchParams(); p.append("foo", "bar"); p.get("foo")`,
			want:   "bar",
		},
		{
			name:   "URLSearchParams set method",
			script: `var p = new URLSearchParams("foo=bar&foo=baz"); p.set("foo", "newval"); p.get("foo")`,
			want:   "newval",
		},
		{
			name:   "URLSearchParams has method",
			script: `new URLSearchParams("foo=bar").has("foo")`,
			want:   true,
		},
		{
			name:   "URLSearchParams has method false",
			script: `new URLSearchParams("foo=bar").has("missing")`,
			want:   false,
		},
		{
			name:   "URLSearchParams delete method",
			script: `var p = new URLSearchParams("foo=bar&baz=qux"); p.delete("foo"); p.has("foo")`,
			want:   false,
		},
		{
			name:   "URLSearchParams getAll method",
			script: `var p = new URLSearchParams("foo=bar&foo=baz"); p.getAll("foo").length`,
			want:   int64(2),
		},
		{
			name:   "URLSearchParams keys method",
			script: `new URLSearchParams("foo=bar&baz=qux").keys().length`,
			want:   int64(2),
		},
		{
			name:   "URLSearchParams values method",
			script: `new URLSearchParams("foo=bar&baz=qux").values().length`,
			want:   int64(2),
		},
		{
			name:   "URLSearchParams entries method",
			script: `new URLSearchParams("foo=bar").entries().length`,
			want:   int64(1),
		},
		{
			name:   "URLSearchParams size property",
			script: `new URLSearchParams("foo=bar&baz=qux").size`,
			want:   int64(2),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := vm.RunString(test.script)
			if err != nil {
				t.Fatalf("Script execution failed: %v", err)
			}

			actual := result.Export()
			if actual != test.want {
				t.Errorf("Expected %v, got %v", test.want, actual)
			}
		})
	}
}

func TestURLSearchParamsForEach(t *testing.T) {
	// Create runtime and bindings
	vm := goja.New()
	document := dom.NewDocument()
	bindings := NewDOMBindings(vm, document)
	bindings.SetupBrowserAPIs()

	script := `
		var params = new URLSearchParams("foo=bar&baz=qux");
		var results = [];
		params.forEach(function(value, name) {
			results.push(name + "=" + value);
		});
		results.join(",");
	`

	result, err := vm.RunString(script)
	if err != nil {
		t.Fatalf("Script execution failed: %v", err)
	}

	actual := result.Export()
	expected := "foo=bar,baz=qux"
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestURLSearchParamsSort(t *testing.T) {
	// Create runtime and bindings
	vm := goja.New()
	document := dom.NewDocument()
	bindings := NewDOMBindings(vm, document)
	bindings.SetupBrowserAPIs()

	script := `
		var params = new URLSearchParams("z=1&a=2&m=3");
		params.sort();
		params.keys().join(",");
	`

	result, err := vm.RunString(script)
	if err != nil {
		t.Fatalf("Script execution failed: %v", err)
	}

	actual := result.Export()
	expected := "a,m,z"
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestURLInvalidConstructor(t *testing.T) {
	// Create runtime and bindings
	vm := goja.New()
	document := dom.NewDocument()
	bindings := NewDOMBindings(vm, document)
	bindings.SetupBrowserAPIs()

	// Test invalid URL - relative URL without base
	script := `
		try {
			new URL("relative/path");
			"should have thrown";
		} catch (e) {
			e.name;
		}
	`

	result, err := vm.RunString(script)
	if err != nil {
		t.Fatalf("Script execution failed: %v", err)
	}

	actual := result.Export()
	if actual != "TypeError" {
		t.Errorf("Expected TypeError, got %v", actual)
	}
}

func TestURLSearchParamsInvalidMethods(t *testing.T) {
	// Create runtime and bindings
	vm := goja.New()
	document := dom.NewDocument()
	bindings := NewDOMBindings(vm, document)
	bindings.SetupBrowserAPIs()

	tests := []struct {
		name   string
		script string
	}{
		{
			name: "append without arguments",
			script: `
				try {
					new URLSearchParams().append();
					"should have thrown";
				} catch (e) {
					e.name;
				}
			`,
		},
		{
			name: "set without arguments",
			script: `
				try {
					new URLSearchParams().set("name");
					"should have thrown";
				} catch (e) {
					e.name;
				}
			`,
		},
		{
			name: "forEach without callback",
			script: `
				try {
					new URLSearchParams().forEach();
					"should have thrown";
				} catch (e) {
					e.name;
				}
			`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := vm.RunString(test.script)
			if err != nil {
				t.Fatalf("Script execution failed: %v", err)
			}

			actual := result.Export()
			if actual != "TypeError" {
				t.Errorf("Expected TypeError, got %v", actual)
			}
		})
	}
}

func TestURLSearchParamsIntegration(t *testing.T) {
	// Create runtime and bindings
	vm := goja.New()
	document := dom.NewDocument()
	bindings := NewDOMBindings(vm, document)
	bindings.SetupBrowserAPIs()

	// Test URL and URLSearchParams integration
	script := `
		var url = new URL("https://example.com?foo=bar&baz=qux");
		var params = url.searchParams;
		params.append("new", "value");
		params.get("new");
	`

	result, err := vm.RunString(script)
	if err != nil {
		t.Fatalf("Script execution failed: %v", err)
	}

	actual := result.Export()
	if actual != "value" {
		t.Errorf("Expected 'value', got %v", actual)
	}
}

func TestURLConstructorRequiresArguments(t *testing.T) {
	// Create runtime and bindings
	vm := goja.New()
	document := dom.NewDocument()
	bindings := NewDOMBindings(vm, document)
	bindings.SetupBrowserAPIs()

	script := `
		try {
			new URL();
			"should have thrown";
		} catch (e) {
			e.name;
		}
	`

	result, err := vm.RunString(script)
	if err != nil {
		t.Fatalf("Script execution failed: %v", err)
	}

	actual := result.Export()
	if actual != "TypeError" {
		t.Errorf("Expected TypeError, got %v", actual)
	}
}

func TestURLSearchParamsSpecialCharacters(t *testing.T) {
	// Create runtime and bindings
	vm := goja.New()
	document := dom.NewDocument()
	bindings := NewDOMBindings(vm, document)
	bindings.SetupBrowserAPIs()

	script := `
		var params = new URLSearchParams();
		params.append("special", "!@#$%^&*()");
		params.append("unicode", "café");
		params.append("spaces", "hello world");
		params.get("special") + "|" + params.get("unicode") + "|" + params.get("spaces");
	`

	result, err := vm.RunString(script)
	if err != nil {
		t.Fatalf("Script execution failed: %v", err)
	}

	actual := result.Export()
	expected := "!@#$%^&*()|café|hello world"
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
