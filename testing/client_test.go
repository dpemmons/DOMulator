package testing

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestNewHTTPClient(t *testing.T) {
	mocks := NewNetworkMocks()
	client := NewHTTPClient(mocks)

	if client == nil {
		t.Fatal("NewHTTPClient() returned nil")
	}

	if client.headers == nil {
		t.Error("HTTPClient.headers should not be nil")
	}

	if client.mocks != mocks {
		t.Error("HTTPClient.mocks should be set to provided mocks")
	}

	if client.handler != nil {
		t.Error("HTTPClient.handler should be nil initially")
	}

	if client.baseURL != "" {
		t.Error("HTTPClient.baseURL should be empty initially")
	}
}

func TestHTTPClient_SetHandler(t *testing.T) {
	client := NewHTTPClient(NewNetworkMocks())

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	client.SetHandler(handler)

	if client.handler == nil {
		t.Error("SetHandler should set the handler")
	}

	if client.baseURL != "" {
		t.Error("SetHandler should clear baseURL")
	}
}

func TestHTTPClient_SetBaseURL(t *testing.T) {
	client := NewHTTPClient(NewNetworkMocks())

	// Set a handler first
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	client.SetHandler(handler)

	client.SetBaseURL("http://example.com/")

	if client.baseURL != "http://example.com" {
		t.Errorf("Expected baseURL to be 'http://example.com', got '%s'", client.baseURL)
	}

	if client.handler != nil {
		t.Error("SetBaseURL should clear handler")
	}
}

func TestHTTPClient_SetBaseURL_TrimsSuffix(t *testing.T) {
	client := NewHTTPClient(NewNetworkMocks())

	client.SetBaseURL("http://example.com/")

	if client.baseURL != "http://example.com" {
		t.Errorf("Expected trailing slash to be trimmed, got '%s'", client.baseURL)
	}
}

func TestHTTPClient_SetHeader(t *testing.T) {
	client := NewHTTPClient(NewNetworkMocks())

	client.SetHeader("X-Test", "test-value")

	if client.headers["X-Test"] != "test-value" {
		t.Error("SetHeader should set header in headers map")
	}
}

func TestHTTPClient_GET_WithHandler(t *testing.T) {
	client := NewHTTPClient(NewNetworkMocks())

	// Create test handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected GET method, got %s", r.Method)
		}

		if r.URL.Path != "/test" {
			t.Errorf("Expected path '/test', got '%s'", r.URL.Path)
		}

		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Test Response"))
	})

	client.SetHandler(handler)

	response := client.GET("/test")

	if response.Error != nil {
		t.Errorf("GET request failed: %v", response.Error)
	}

	if response.Status != http.StatusOK {
		t.Errorf("Expected status 200, got %d", response.Status)
	}

	if response.StatusText != "OK" {
		t.Errorf("Expected status text 'OK', got '%s'", response.StatusText)
	}

	if response.Body != "Test Response" {
		t.Errorf("Expected body 'Test Response', got '%s'", response.Body)
	}

	if response.Headers["Content-Type"] != "text/plain" {
		t.Errorf("Expected Content-Type header 'text/plain', got '%s'", response.Headers["Content-Type"])
	}
}

func TestHTTPClient_POST_WithHandler(t *testing.T) {
	client := NewHTTPClient(NewNetworkMocks())

	// Create test handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected POST method, got %s", r.Method)
		}

		// Read body
		buf := make([]byte, 1024)
		n, _ := r.Body.Read(buf)
		body := string(buf[:n])

		if body != "test data" {
			t.Errorf("Expected body 'test data', got '%s'", body)
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Created"))
	})

	client.SetHandler(handler)

	response := client.POST("/create", "test data")

	if response.Error != nil {
		t.Errorf("POST request failed: %v", response.Error)
	}

	if response.Status != http.StatusCreated {
		t.Errorf("Expected status 201, got %d", response.Status)
	}

	if response.Body != "Created" {
		t.Errorf("Expected body 'Created', got '%s'", response.Body)
	}
}

func TestHTTPClient_POST_WithJSONBody(t *testing.T) {
	client := NewHTTPClient(NewNetworkMocks())

	// Create test handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var data map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			t.Errorf("Failed to decode JSON: %v", err)
		}

		if data["name"] != "test" {
			t.Errorf("Expected name 'test', got '%v'", data["name"])
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("JSON received"))
	})

	client.SetHandler(handler)

	// Test with struct that will be JSON marshaled
	body := map[string]interface{}{
		"name": "test",
		"id":   123,
	}

	response := client.POST("/json", body)

	if response.Error != nil {
		t.Errorf("POST JSON request failed: %v", response.Error)
	}

	if response.Status != http.StatusOK {
		t.Errorf("Expected status 200, got %d", response.Status)
	}

	if response.Body != "JSON received" {
		t.Errorf("Expected body 'JSON received', got '%s'", response.Body)
	}
}

func TestHTTPClient_POST_WithBytesBody(t *testing.T) {
	client := NewHTTPClient(NewNetworkMocks())

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		buf := make([]byte, 1024)
		n, _ := r.Body.Read(buf)
		body := string(buf[:n])

		if body != "byte data" {
			t.Errorf("Expected body 'byte data', got '%s'", body)
		}

		w.WriteHeader(http.StatusOK)
	})

	client.SetHandler(handler)

	response := client.POST("/bytes", []byte("byte data"))

	if response.Error != nil {
		t.Errorf("POST bytes request failed: %v", response.Error)
	}
}

func TestHTTPClient_POST_WithReaderBody(t *testing.T) {
	client := NewHTTPClient(NewNetworkMocks())

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		buf := make([]byte, 1024)
		n, _ := r.Body.Read(buf)
		body := string(buf[:n])

		if body != "reader data" {
			t.Errorf("Expected body 'reader data', got '%s'", body)
		}

		w.WriteHeader(http.StatusOK)
	})

	client.SetHandler(handler)

	reader := strings.NewReader("reader data")
	response := client.POST("/reader", reader)

	if response.Error != nil {
		t.Errorf("POST reader request failed: %v", response.Error)
	}
}

func TestHTTPClient_POST_JSONMarshalError(t *testing.T) {
	client := NewHTTPClient(NewNetworkMocks())

	// Create an unmarshalable type
	unmarshalable := make(chan int)

	response := client.POST("/test", unmarshalable)

	if response.Error == nil {
		t.Error("Expected error for unmarshalable body")
	}

	if !strings.Contains(response.Error.Error(), "marshal") {
		t.Errorf("Expected marshal error, got: %v", response.Error)
	}
}

func TestHTTPClient_WithHeaders(t *testing.T) {
	client := NewHTTPClient(NewNetworkMocks())

	// Set headers
	client.SetHeader("X-Custom", "custom-value")
	client.SetHeader("Authorization", "Bearer token123")

	// Create handler that checks headers
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Custom") != "custom-value" {
			t.Errorf("Expected X-Custom header 'custom-value', got '%s'", r.Header.Get("X-Custom"))
		}

		if r.Header.Get("Authorization") != "Bearer token123" {
			t.Errorf("Expected Authorization header 'Bearer token123', got '%s'", r.Header.Get("Authorization"))
		}

		w.WriteHeader(http.StatusOK)
	})

	client.SetHandler(handler)

	response := client.GET("/test")

	if response.Error != nil {
		t.Errorf("Request with headers failed: %v", response.Error)
	}
}

func TestHTTPClient_WithMocks(t *testing.T) {
	mocks := NewNetworkMocks()
	client := NewHTTPClient(mocks)

	// Add mock response
	mocks.MockGET("/api/users", MockResponse{
		Status: http.StatusOK,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: `{"users": ["alice", "bob"]}`,
	})

	response := client.GET("/api/users")

	if response.Error != nil {
		t.Errorf("Mock request failed: %v", response.Error)
	}

	if response.Status != http.StatusOK {
		t.Errorf("Expected status 200, got %d", response.Status)
	}

	if response.Headers["Content-Type"] != "application/json" {
		t.Errorf("Expected Content-Type 'application/json', got '%s'", response.Headers["Content-Type"])
	}

	if !strings.Contains(response.Body, "alice") {
		t.Error("Response body should contain mock data")
	}
}

func TestHTTPClient_MocksOverrideHandler(t *testing.T) {
	mocks := NewNetworkMocks()
	client := NewHTTPClient(mocks)

	// Set up both handler and mock
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Handler Response"))
	})
	client.SetHandler(handler)

	// Mock should override handler
	mocks.MockGET("/test", MockResponse{
		Status: http.StatusOK,
		Body:   "Mock Response",
	})

	response := client.GET("/test")

	if response.Body != "Mock Response" {
		t.Errorf("Expected mock to override handler, got '%s'", response.Body)
	}
}

func TestHTTPClient_NoHandlerOrBaseURL(t *testing.T) {
	client := NewHTTPClient(NewNetworkMocks())

	response := client.GET("/test")

	if response.Error == nil {
		t.Error("Expected error when no handler or base URL configured")
	}

	if !strings.Contains(response.Error.Error(), "no handler or base URL") {
		t.Errorf("Expected configuration error, got: %v", response.Error)
	}
}

// Integration test with real HTTP server
func TestHTTPClient_RealServerIntegration(t *testing.T) {
	// Create test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/ping" {
			w.Header().Set("X-Server", "test")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("pong"))
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()

	client := NewHTTPClient(NewNetworkMocks())
	client.SetBaseURL(server.URL)

	response := client.GET("/ping")

	if response.Error != nil {
		t.Errorf("Real server request failed: %v", response.Error)
	}

	if response.Status != http.StatusOK {
		t.Errorf("Expected status 200, got %d", response.Status)
	}

	if response.Body != "pong" {
		t.Errorf("Expected body 'pong', got '%s'", response.Body)
	}

	if response.Headers["X-Server"] != "test" {
		t.Errorf("Expected X-Server header 'test', got '%s'", response.Headers["X-Server"])
	}
}

func TestHTTPClient_RealServerPOST(t *testing.T) {
	// Create test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" && r.URL.Path == "/data" {
			buf := new(bytes.Buffer)
			buf.ReadFrom(r.Body)
			body := buf.String()

			if body == "test payload" {
				w.WriteHeader(http.StatusCreated)
				w.Write([]byte("received"))
			} else {
				w.WriteHeader(http.StatusBadRequest)
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()

	client := NewHTTPClient(NewNetworkMocks())
	client.SetBaseURL(server.URL)

	response := client.POST("/data", "test payload")

	if response.Error != nil {
		t.Errorf("Real server POST failed: %v", response.Error)
	}

	if response.Status != http.StatusCreated {
		t.Errorf("Expected status 201, got %d", response.Status)
	}

	if response.Body != "received" {
		t.Errorf("Expected body 'received', got '%s'", response.Body)
	}
}

func TestHTTPClient_InvalidURL(t *testing.T) {
	client := NewHTTPClient(NewNetworkMocks())
	client.SetBaseURL("invalid://[::1]:namedport")

	response := client.GET("/test")

	if response.Error == nil {
		t.Error("Expected error for invalid URL")
	}

	if !strings.Contains(response.Error.Error(), "invalid URL") {
		t.Errorf("Expected URL error, got: %v", response.Error)
	}
}

func TestResponse_LoadHTML(t *testing.T) {
	response := &Response{
		Status: http.StatusOK,
		Body:   "<html><body>Test</body></html>",
	}

	html := response.LoadHTML()

	if html != "<html><body>Test</body></html>" {
		t.Errorf("LoadHTML should return body, got '%s'", html)
	}
}

func TestResponse_LoadHTML_Error(t *testing.T) {
	response := &Response{
		Error: http.ErrServerClosed,
	}

	// Should panic on error
	defer func() {
		if r := recover(); r == nil {
			t.Error("LoadHTML should panic when response has error")
		}
	}()

	response.LoadHTML()
}
