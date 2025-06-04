package js

import (
	"testing"

	"github.com/dop251/goja"
	"github.com/dpemmons/DOMulator/internal/dom"
)

func TestDocumentCookieIntegration(t *testing.T) {
	// Setup runtime and bindings
	vm := goja.New()
	document := dom.NewDocument()
	bindings := NewDOMBindings(vm, document)

	// Setup document global
	vm.Set("document", bindings.WrapDocument())

	// Test 1: Initial empty cookie string
	t.Run("InitialEmptyCookies", func(t *testing.T) {
		result, err := vm.RunString("document.cookie")
		if err != nil {
			t.Fatalf("Error getting initial cookies: %v", err)
		}

		if result.String() != "" {
			t.Errorf("Expected empty cookie string initially, got '%s'", result.String())
		}
	})

	// Test 2: Setting a simple cookie
	t.Run("SetSimpleCookie", func(t *testing.T) {
		_, err := vm.RunString("document.cookie = 'test=value'")
		if err != nil {
			t.Fatalf("Error setting cookie: %v", err)
		}

		result, err := vm.RunString("document.cookie")
		if err != nil {
			t.Fatalf("Error getting cookies: %v", err)
		}

		expected := "test=value"
		if result.String() != expected {
			t.Errorf("Expected cookie string '%s', got '%s'", expected, result.String())
		}
	})

	// Test 3: Setting multiple cookies
	t.Run("SetMultipleCookies", func(t *testing.T) {
		_, err := vm.RunString(`
			document.cookie = 'cookie1=value1';
			document.cookie = 'cookie2=value2';
		`)
		if err != nil {
			t.Fatalf("Error setting multiple cookies: %v", err)
		}

		result, err := vm.RunString("document.cookie")
		if err != nil {
			t.Fatalf("Error getting cookies: %v", err)
		}

		cookieStr := result.String()
		// Should contain both cookies (order may vary)
		if !contains(cookieStr, "cookie1=value1") {
			t.Error("Expected cookie string to contain 'cookie1=value1'")
		}
		if !contains(cookieStr, "cookie2=value2") {
			t.Error("Expected cookie string to contain 'cookie2=value2'")
		}
		if !contains(cookieStr, "test=value") {
			t.Error("Expected cookie string to contain previous 'test=value'")
		}
	})

	// Test 4: Setting cookie with attributes
	t.Run("SetCookieWithAttributes", func(t *testing.T) {
		_, err := vm.RunString("document.cookie = 'secure_cookie=secret; Path=/admin; Secure; HttpOnly; SameSite=Strict'")
		if err != nil {
			t.Fatalf("Error setting cookie with attributes: %v", err)
		}

		// Check that cookie manager has the cookie with correct attributes
		cookie := bindings.cookieManager.GetCookie("secure_cookie")
		if cookie == nil {
			t.Fatal("Expected to find secure_cookie")
		}

		if cookie.Value != "secret" {
			t.Errorf("Expected value 'secret', got '%s'", cookie.Value)
		}
		if cookie.Path != "/admin" {
			t.Errorf("Expected path '/admin', got '%s'", cookie.Path)
		}
		if !cookie.Secure {
			t.Error("Expected Secure flag to be true")
		}
		if !cookie.HttpOnly {
			t.Error("Expected HttpOnly flag to be true")
		}
	})

	// Test 5: Cookie deletion with Max-Age
	t.Run("DeleteCookieWithMaxAge", func(t *testing.T) {
		// First set a cookie
		_, err := vm.RunString("document.cookie = 'temp_cookie=temp_value'")
		if err != nil {
			t.Fatalf("Error setting temp cookie: %v", err)
		}

		// Verify it exists
		if bindings.cookieManager.GetCookie("temp_cookie") == nil {
			t.Fatal("Expected temp_cookie to exist")
		}

		// Delete it with Max-Age=0
		_, err = vm.RunString("document.cookie = 'temp_cookie=; Max-Age=-1'")
		if err != nil {
			t.Fatalf("Error deleting cookie: %v", err)
		}

		// Verify it's gone
		if bindings.cookieManager.GetCookie("temp_cookie") != nil {
			t.Error("Expected temp_cookie to be deleted")
		}
	})

	// Test 6: JavaScript cookie manipulation
	t.Run("JavaScriptCookieHelpers", func(t *testing.T) {
		// Test JavaScript cookie helper functions
		script := `
			// Helper function to get a specific cookie
			function getCookie(name) {
				let cookies = document.cookie.split(';');
				for (let cookie of cookies) {
					let [key, value] = cookie.trim().split('=');
					if (key === name) {
						return value;
					}
				}
				return null;
			}
			
			// Set a test cookie
			document.cookie = 'js_test=js_value';
			
			// Get the cookie value
			getCookie('js_test');
		`

		result, err := vm.RunString(script)
		if err != nil {
			t.Fatalf("Error running JavaScript cookie helpers: %v", err)
		}

		if result.String() != "js_value" {
			t.Errorf("Expected 'js_value', got '%s'", result.String())
		}
	})

	// Test 7: Error handling for invalid cookies
	t.Run("InvalidCookieHandling", func(t *testing.T) {
		// Invalid cookies should be silently ignored (browser behavior)
		_, err := vm.RunString("document.cookie = '=invalid'") // Empty name
		if err != nil {
			t.Fatalf("Setting invalid cookie should not throw error: %v", err)
		}

		_, err = vm.RunString("document.cookie = 'invalid_no_equals'") // No equals sign
		if err != nil {
			t.Fatalf("Setting invalid cookie should not throw error: %v", err)
		}

		// Cookie string should not contain invalid cookies
		result, err := vm.RunString("document.cookie")
		if err != nil {
			t.Fatalf("Error getting cookies: %v", err)
		}

		cookieStr := result.String()
		if contains(cookieStr, "invalid") {
			t.Error("Cookie string should not contain invalid cookies")
		}
	})
}

func TestCookieManagerDomainAndPath(t *testing.T) {
	// Test that cookie manager can be configured with different domains/paths
	vm := goja.New()
	document := dom.NewDocument()
	bindings := NewDOMBindings(vm, document)
	vm.Set("document", bindings.WrapDocument())

	// Set a cookie and verify defaults
	_, err := vm.RunString("document.cookie = 'domain_test=value'")
	if err != nil {
		t.Fatalf("Error setting cookie: %v", err)
	}

	cookie := bindings.cookieManager.GetCookie("domain_test")
	if cookie == nil {
		t.Fatal("Expected to find domain_test cookie")
	}

	// Should use default domain and path
	if cookie.Domain != "localhost" {
		t.Errorf("Expected default domain 'localhost', got '%s'", cookie.Domain)
	}
	if cookie.Path != "/" {
		t.Errorf("Expected default path '/', got '%s'", cookie.Path)
	}
}

// Helper function to check if a string contains a substring
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > len(substr) &&
		(s[:len(substr)] == substr || s[len(s)-len(substr):] == substr ||
			indexOfSubstring(s, substr) >= 0))
}

func indexOfSubstring(s, substr string) int {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}
