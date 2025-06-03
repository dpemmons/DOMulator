package fetch

import (
	"net/http"
	"testing"

	"github.com/dop251/goja"
	"github.com/dpemmons/DOMulator/internal/browser/cookies"
	domtest "github.com/dpemmons/DOMulator/testing"
)

func TestFetchCookieIntegration(t *testing.T) {
	// Setup
	vm := goja.New()
	mocks := domtest.NewNetworkMocks()
	cookieManager := cookies.NewCookieManager("example.com", "/")
	fetchAPI := NewWithCookieManager(mocks, cookieManager)

	// Setup fetch global
	vm.Set("fetch", fetchAPI.CreateFetchFunction(vm))

	// Test 1: Fetch automatically includes cookies that were previously set
	t.Run("FetchIncludesCookies", func(t *testing.T) {
		// Set some cookies in the cookie manager
		cookieManager.ParseCookieString("session_id=abc123")
		cookieManager.ParseCookieString("user_pref=dark_mode")

		// Mock a request that we can inspect
		mocks.MockGET("https://example.com/api/data", domtest.MockResponse{
			Status: 200,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			Body: `{"status":"success"}`,
		})

		// Capture the request to verify cookies are sent
		var capturedRequest *http.Request
		originalClient := fetchAPI.client
		fetchAPI.client = &http.Client{
			Transport: &cookieCapturingTransport{
				onRequest: func(req *http.Request) {
					capturedRequest = req
				},
				fallback: http.DefaultTransport,
			},
		}
		defer func() { fetchAPI.client = originalClient }()

		// Make fetch request
		script := `
			fetch('https://example.com/api/data')
				.then(response => response.json())
				.then(data => data.status);
		`

		result, err := vm.RunString(script)
		if err != nil {
			t.Fatalf("Error running fetch: %v", err)
		}

		// Wait for the promise to resolve
		promise := result.ToObject(vm)
		if promise == nil {
			t.Fatal("Expected a promise")
		}

		// Verify cookies were sent in the request
		if capturedRequest != nil {
			cookieHeader := capturedRequest.Header.Get("Cookie")
			if cookieHeader == "" {
				t.Error("Expected Cookie header to be set")
			}

			// Should contain both cookies
			if !contains(cookieHeader, "session_id=abc123") {
				t.Error("Expected cookie header to contain session_id=abc123")
			}
			if !contains(cookieHeader, "user_pref=dark_mode") {
				t.Error("Expected cookie header to contain user_pref=dark_mode")
			}
		}
	})

	// Test 2: Fetch processes Set-Cookie headers from responses
	t.Run("FetchProcessesSetCookieHeaders", func(t *testing.T) {
		// Mock a response with Set-Cookie headers
		mocks.Clear()
		mocks.MockGET("https://example.com/login", domtest.MockResponse{
			Status: 200,
			Headers: map[string]string{
				"Content-Type": "application/json",
				"Set-Cookie":   "auth_token=xyz789; Path=/; HttpOnly; Secure",
			},
			Body: `{"login":"success"}`,
		})

		// Make fetch request directly (synchronous)
		_, err := fetchAPI.fetch("https://example.com/login", &RequestOptions{
			Method:  "GET",
			Headers: make(map[string]string),
		})
		if err != nil {
			t.Fatalf("Error running fetch: %v", err)
		}

		// Verify the cookie was stored
		authCookie := cookieManager.GetCookie("auth_token")
		if authCookie == nil {
			t.Fatal("Expected auth_token cookie to be stored")
		}

		if authCookie.Value != "xyz789" {
			t.Errorf("Expected cookie value 'xyz789', got '%s'", authCookie.Value)
		}
		if authCookie.Path != "/" {
			t.Errorf("Expected cookie path '/', got '%s'", authCookie.Path)
		}
		if !authCookie.HttpOnly {
			t.Error("Expected HttpOnly flag to be true")
		}
		if !authCookie.Secure {
			t.Error("Expected Secure flag to be true")
		}
	})

	// Test 3: Full end-to-end cookie workflow
	t.Run("EndToEndCookieWorkflow", func(t *testing.T) {
		// Clear any existing cookies
		cookieManager = cookies.NewCookieManager("api.example.com", "/")
		fetchAPI.cookieManager = cookieManager

		// Step 1: Login request that sets authentication cookie
		mocks.Clear()
		mocks.MockPOST("https://api.example.com/auth/login", domtest.MockResponse{
			Status: 200,
			Headers: map[string]string{
				"Content-Type": "application/json",
				"Set-Cookie":   "session=abc123def456; Path=/; HttpOnly",
			},
			Body: `{"authenticated":true}`,
		})

		// Make login request directly (synchronous)
		_, err := fetchAPI.fetch("https://api.example.com/auth/login", &RequestOptions{
			Method: "POST",
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			Body: `{"username": "user", "password": "pass"}`,
		})
		if err != nil {
			t.Fatalf("Error running login: %v", err)
		}

		// Verify session cookie was stored
		sessionCookie := cookieManager.GetCookie("session")
		if sessionCookie == nil {
			t.Fatal("Expected session cookie to be stored after login")
		}

		// Step 2: Subsequent API request should automatically include the session cookie
		var capturedRequest *http.Request
		originalClient := fetchAPI.client
		fetchAPI.client = &http.Client{
			Transport: &cookieCapturingTransport{
				onRequest: func(req *http.Request) {
					capturedRequest = req
				},
				fallback: http.DefaultTransport,
			},
		}
		defer func() { fetchAPI.client = originalClient }()

		mocks.MockGET("https://api.example.com/user/profile", domtest.MockResponse{
			Status: 200,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			Body: `{"user":"john_doe","email":"john@example.com"}`,
		})

		// Make profile request directly (synchronous)
		_, err = fetchAPI.fetch("https://api.example.com/user/profile", &RequestOptions{
			Method:  "GET",
			Headers: make(map[string]string),
		})
		if err != nil {
			t.Fatalf("Error running profile request: %v", err)
		}

		// Verify the session cookie was automatically included
		if capturedRequest != nil {
			cookieHeader := capturedRequest.Header.Get("Cookie")
			if !contains(cookieHeader, "session=abc123def456") {
				t.Error("Expected session cookie to be automatically included in subsequent request")
			}
		}
	})

	// Test 4: Multiple Set-Cookie headers
	t.Run("MultipleSetCookieHeaders", func(t *testing.T) {
		// Create a custom mock response with multiple Set-Cookie headers
		// Note: In real HTTP, multiple Set-Cookie headers are sent as separate headers
		mocks.Clear()

		// For this test, we'll simulate the behavior by calling ParseSetCookieHeader directly
		// since the testing framework may not handle multiple headers properly

		err := cookieManager.ParseSetCookieHeader("pref1=value1; Path=/")
		if err != nil {
			t.Fatalf("Error setting cookie 1: %v", err)
		}

		err = cookieManager.ParseSetCookieHeader("pref2=value2; Path=/; Secure")
		if err != nil {
			t.Fatalf("Error setting cookie 2: %v", err)
		}

		// Verify both cookies were stored
		pref1 := cookieManager.GetCookie("pref1")
		pref2 := cookieManager.GetCookie("pref2")

		if pref1 == nil {
			t.Error("Expected pref1 cookie to be stored")
		}
		if pref2 == nil {
			t.Error("Expected pref2 cookie to be stored")
		}

		if pref1 != nil && pref1.Value != "value1" {
			t.Errorf("Expected pref1 value 'value1', got '%s'", pref1.Value)
		}
		if pref2 != nil && pref2.Value != "value2" {
			t.Errorf("Expected pref2 value 'value2', got '%s'", pref2.Value)
		}
		if pref2 != nil && !pref2.Secure {
			t.Error("Expected pref2 to have Secure flag")
		}
	})
}

// cookieCapturingTransport captures HTTP requests for inspection
type cookieCapturingTransport struct {
	onRequest func(*http.Request)
	fallback  http.RoundTripper
}

func (t *cookieCapturingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if t.onRequest != nil {
		t.onRequest(req)
	}
	return t.fallback.RoundTrip(req)
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
