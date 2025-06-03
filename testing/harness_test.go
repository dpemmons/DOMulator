package testing

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestNewTestHarness(t *testing.T) {
	harness := NewTestHarness()

	if harness == nil {
		t.Fatal("NewTestHarness() returned nil")
	}

	if harness.client == nil {
		t.Error("TestHarness.client should not be nil")
	}

	if harness.mocks == nil {
		t.Error("TestHarness.mocks should not be nil")
	}

	if harness.config == nil {
		t.Error("TestHarness.config should not be nil")
	}

	if harness.config.DefaultTimeout != 5*time.Second {
		t.Errorf("Expected default timeout to be 5s, got %v", harness.config.DefaultTimeout)
	}

	if harness.config.Headers == nil {
		t.Error("Config.Headers should not be nil")
	}
}

func TestTestHarness_LoadHTML(t *testing.T) {
	harness := NewTestHarness()

	html := `<html><head><title>Test</title></head><body><h1>Hello World</h1></body></html>`

	result := harness.LoadHTML(html)

	// Should return self for chaining
	if result != harness {
		t.Error("LoadHTML should return self for chaining")
	}

	// Should have created domulator
	if harness.domulator == nil {
		t.Fatal("LoadHTML should create domulator")
	}

	if harness.domulator.document == nil {
		t.Error("LoadHTML should create document")
	}

	if harness.domulator.runtime == nil {
		t.Error("LoadHTML should create JavaScript runtime")
	}

	// Test Document() method
	doc := harness.Document()
	if doc == nil {
		t.Error("Document() should return document after LoadHTML")
	}

	// Test Runtime() method
	runtime := harness.Runtime()
	if runtime == nil {
		t.Error("Runtime() should return runtime after LoadHTML")
	}
}

func TestTestHarness_LoadHTML_InvalidHTML(t *testing.T) {
	harness := NewTestHarness()

	// Even invalid HTML should parse (HTML parser is forgiving)
	html := `<invalid><unclosed><malformed`

	// Should not panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("LoadHTML should not panic on invalid HTML, got panic: %v", r)
		}
	}()

	harness.LoadHTML(html)

	if harness.domulator == nil {
		t.Error("LoadHTML should still create domulator even with invalid HTML")
	}
}

func TestTestHarness_ExecuteScript(t *testing.T) {
	harness := NewTestHarness()

	// Test error when no document loaded
	err := harness.ExecuteScript("console.log('test')")
	if err == nil {
		t.Error("ExecuteScript should return error when no document loaded")
	}

	// Load HTML first
	html := `<html><body><div id="test">Initial</div></body></html>`
	harness.LoadHTML(html)

	// Test simple script execution
	err = harness.ExecuteScript("console.log('Hello from JavaScript')")
	if err != nil {
		t.Errorf("ExecuteScript failed: %v", err)
	}

	// Test DOM manipulation
	err = harness.ExecuteScript(`
		var div = document.getElementById('test');
		div.textContent = 'Modified by JS';
	`)
	if err != nil {
		t.Errorf("ExecuteScript DOM manipulation failed: %v", err)
	}

	// Verify the DOM was actually modified
	// Note: We'd need to check this through DOM inspection
	// This tests that the script executed without error
}

func TestTestHarness_ExecuteScript_JSError(t *testing.T) {
	harness := NewTestHarness()
	html := `<html><body></body></html>`
	harness.LoadHTML(html)

	// Test JavaScript syntax error
	err := harness.ExecuteScript("invalid javascript syntax !!!")
	if err == nil {
		t.Error("ExecuteScript should return error for invalid JavaScript")
	}
}

func TestTestHarness_SetHandler(t *testing.T) {
	harness := NewTestHarness()

	// Create a simple test handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("<html><body>Handler Response</body></html>"))
	})

	result := harness.SetHandler(handler)

	// Should return self for chaining
	if result != harness {
		t.Error("SetHandler should return self for chaining")
	}

	// Test that handler is actually used
	response := harness.GET("/test")
	if response.Error != nil {
		t.Errorf("GET request failed: %v", response.Error)
	}

	if response.Status != http.StatusOK {
		t.Errorf("Expected status 200, got %d", response.Status)
	}

	if !strings.Contains(response.Body, "Handler Response") {
		t.Error("Response should contain handler content")
	}
}

func TestTestHarness_SetBaseURL(t *testing.T) {
	harness := NewTestHarness()

	result := harness.SetBaseURL("http://example.com")

	// Should return self for chaining
	if result != harness {
		t.Error("SetBaseURL should return self for chaining")
	}

	if harness.config.BaseURL != "http://example.com" {
		t.Errorf("Expected base URL to be set to 'http://example.com', got '%s'", harness.config.BaseURL)
	}
}

func TestTestHarness_Navigate_WithHandler(t *testing.T) {
	harness := NewTestHarness()

	// Create test handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/page1" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`<html><head><title>Page 1</title></head><body><h1>Page 1 Content</h1></body></html>`))
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Not Found"))
		}
	})

	harness.SetHandler(handler)

	result := harness.Navigate("/page1")

	// Should return self for chaining
	if result != harness {
		t.Error("Navigate should return self for chaining")
	}

	// Should have loaded the page
	if harness.domulator == nil {
		t.Fatal("Navigate should create domulator")
	}

	doc := harness.Document()
	if doc == nil {
		t.Error("Navigate should create document")
	}
}

func TestTestHarness_Navigate_HandlerError(t *testing.T) {
	harness := NewTestHarness()

	// Create handler that returns error status
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Server Error"))
	})

	harness.SetHandler(handler)

	// HTTP error status codes are valid responses, not transport errors
	// Navigate should succeed but load the error content
	result := harness.Navigate("/error")

	// Should return self for chaining
	if result != harness {
		t.Error("Navigate should return self for chaining even with error status")
	}

	// Should have loaded the error content
	if harness.domulator == nil {
		t.Error("Navigate should create domulator even with error status")
	}
}

func TestTestHarness_GET_POST(t *testing.T) {
	harness := NewTestHarness()

	// Create test handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" && r.URL.Path == "/get" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("GET Response"))
		} else if r.Method == "POST" && r.URL.Path == "/post" {
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("POST Response"))
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	})

	harness.SetHandler(handler)

	// Test GET
	getResp := harness.GET("/get")
	if getResp.Error != nil {
		t.Errorf("GET failed: %v", getResp.Error)
	}
	if getResp.Status != http.StatusOK {
		t.Errorf("Expected GET status 200, got %d", getResp.Status)
	}
	if getResp.Body != "GET Response" {
		t.Errorf("Expected GET body 'GET Response', got '%s'", getResp.Body)
	}

	// Test POST
	postResp := harness.POST("/post", "test data")
	if postResp.Error != nil {
		t.Errorf("POST failed: %v", postResp.Error)
	}
	if postResp.Status != http.StatusCreated {
		t.Errorf("Expected POST status 201, got %d", postResp.Status)
	}
	if postResp.Body != "POST Response" {
		t.Errorf("Expected POST body 'POST Response', got '%s'", postResp.Body)
	}
}

func TestTestHarness_SetTimeout(t *testing.T) {
	harness := NewTestHarness()

	timeout := 10 * time.Second
	result := harness.SetTimeout(timeout)

	// Should return self for chaining
	if result != harness {
		t.Error("SetTimeout should return self for chaining")
	}

	if harness.config.DefaultTimeout != timeout {
		t.Errorf("Expected timeout to be %v, got %v", timeout, harness.config.DefaultTimeout)
	}
}

func TestTestHarness_SetHeader(t *testing.T) {
	harness := NewTestHarness()

	result := harness.SetHeader("X-Test-Header", "test-value")

	// Should return self for chaining
	if result != harness {
		t.Error("SetHeader should return self for chaining")
	}

	if harness.config.Headers["X-Test-Header"] != "test-value" {
		t.Error("Header should be set in config")
	}

	// Test that header is included in requests
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		headerValue := r.Header.Get("X-Test-Header")
		if headerValue != "test-value" {
			t.Errorf("Expected header value 'test-value', got '%s'", headerValue)
		}
		w.WriteHeader(http.StatusOK)
	})

	harness.SetHandler(handler)
	harness.GET("/test")
}

func TestTestHarness_Document_Runtime_NilBeforeLoad(t *testing.T) {
	harness := NewTestHarness()

	// Should return nil before loading HTML
	if harness.Document() != nil {
		t.Error("Document() should return nil before LoadHTML")
	}

	if harness.Runtime() != nil {
		t.Error("Runtime() should return nil before LoadHTML")
	}
}

func TestTestHarness_Close(t *testing.T) {
	harness := NewTestHarness()

	// Test close without loading anything
	err := harness.Close()
	if err != nil {
		t.Errorf("Close() should not error when nothing loaded: %v", err)
	}

	// Test close after loading HTML
	html := `<html><body></body></html>`
	harness.LoadHTML(html)

	err = harness.Close()
	if err != nil {
		t.Errorf("Close() should not error after LoadHTML: %v", err)
	}
}

func TestTestHarness_ChainableAPI(t *testing.T) {
	harness := NewTestHarness()

	// Create test handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`<html><body><div id="test">Hello</div></body></html>`))
	})

	// Test that all methods are chainable
	result := harness.
		SetTimeout(1*time.Second).
		SetHeader("User-Agent", "TestHarness").
		SetHandler(handler).
		Navigate("/test")

	if result != harness {
		t.Error("API should be fully chainable")
	}

	if harness.Document() == nil {
		t.Error("Chained operations should result in loaded document")
	}
}

// Integration test with real HTTP server
func TestTestHarness_RealServerIntegration(t *testing.T) {
	// Create test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`<html><head><title>Test Server</title></head><body><h1>Hello from server</h1></body></html>`))
	}))
	defer server.Close()

	harness := NewTestHarness()
	harness.SetBaseURL(server.URL)

	// Test navigation to real server
	harness.Navigate("/")

	if harness.Document() == nil {
		t.Error("Should have loaded document from real server")
	}
}
