package fetch

import (
	"testing"

	"github.com/dop251/goja"
	domtesting "github.com/dpemmons/DOMulator/testing"
)

func TestFetchAPI_CreateFetchFunction(t *testing.T) {
	vm := goja.New()
	mocks := domtesting.NewNetworkMocks()
	fetchAPI := New(mocks)

	// Bind fetch to JavaScript
	vm.Set("fetch", fetchAPI.CreateFetchFunction(vm))

	// Test that fetch function is properly bound
	result, err := vm.RunString(`typeof fetch`)
	if err != nil {
		t.Fatalf("Failed to check fetch type: %v", err)
	}

	if result.String() != "function" {
		t.Errorf("Expected fetch to be a function, got %s", result.String())
	}

	// Test that fetch returns a promise
	mocks.MockGET("/api/test", domtesting.MockResponse{
		Status: 200,
		Body:   "Hello World",
		Headers: map[string]string{
			"Content-Type": "text/plain",
		},
	})

	result, err = vm.RunString(`
		let promise = fetch('/api/test');
		typeof promise;
	`)

	if err != nil {
		t.Fatalf("Failed to execute fetch: %v", err)
	}

	if result.String() != "object" {
		t.Errorf("Expected fetch to return an object (Promise), got %s", result.String())
	}
}

func TestFetchAPI_MockIntegration(t *testing.T) {
	mocks := domtesting.NewNetworkMocks()
	fetchAPI := New(mocks)

	// Test GET request mock
	mocks.MockGET("/api/data", domtesting.MockResponse{
		Status: 200,
		Body:   `{"message": "success"}`,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	})

	response, err := fetchAPI.fetch("/api/data", &RequestOptions{Method: "GET"})
	if err != nil {
		t.Fatalf("Fetch failed: %v", err)
	}

	if response.Status != 200 {
		t.Errorf("Expected status 200, got %d", response.Status)
	}

	if response.Body != `{"message": "success"}` {
		t.Errorf("Expected JSON body, got %s", response.Body)
	}

	if !response.Ok {
		t.Error("Expected response.Ok to be true")
	}
}

func TestFetchAPI_POSTRequest(t *testing.T) {
	mocks := domtesting.NewNetworkMocks()
	fetchAPI := New(mocks)

	// Test POST request mock
	mocks.MockPOST("/api/submit", domtesting.MockResponse{
		Status: 201,
		Body:   "Created",
	})

	response, err := fetchAPI.fetch("/api/submit", &RequestOptions{
		Method: "POST",
		Body:   "test data",
		Headers: map[string]string{
			"Content-Type": "text/plain",
		},
	})

	if err != nil {
		t.Fatalf("POST request failed: %v", err)
	}

	if response.Status != 201 {
		t.Errorf("Expected status 201, got %d", response.Status)
	}
}

func TestFetchAPI_RequestOptions(t *testing.T) {
	vm := goja.New()
	fetchAPI := New(nil)

	// Test parsing of request options
	optionsJS, _ := vm.RunString(`({
		method: 'POST',
		headers: {
			'Content-Type': 'application/json',
			'Authorization': 'Bearer token123'
		},
		body: '{"test": true}'
	})`)

	options := fetchAPI.parseRequestOptions(vm, optionsJS)

	if options.Method != "POST" {
		t.Errorf("Expected method POST, got %s", options.Method)
	}

	if options.Body != `{"test": true}` {
		t.Errorf("Expected JSON body, got %s", options.Body)
	}

	if options.Headers["Content-Type"] != "application/json" {
		t.Errorf("Expected Content-Type header, got %s", options.Headers["Content-Type"])
	}

	if options.Headers["Authorization"] != "Bearer token123" {
		t.Errorf("Expected Authorization header, got %s", options.Headers["Authorization"])
	}
}

func TestResponse_JSON(t *testing.T) {
	response := &Response{
		Body: `{"name": "test", "value": 42}`,
	}

	var result map[string]interface{}
	err := response.JSON(&result)
	if err != nil {
		t.Fatalf("JSON parsing failed: %v", err)
	}

	if result["name"] != "test" {
		t.Errorf("Expected name 'test', got %v", result["name"])
	}

	if result["value"] != float64(42) {
		t.Errorf("Expected value 42, got %v", result["value"])
	}
}

func TestResponse_JSONError(t *testing.T) {
	response := &Response{
		Body: `{"invalid": json}`,
	}

	var result map[string]interface{}
	err := response.JSON(&result)
	if err == nil {
		t.Fatal("Expected JSON parsing to fail")
	}
}

func TestFetchAPI_ErrorHandling(t *testing.T) {
	mocks := domtesting.NewNetworkMocks()
	fetchAPI := New(mocks)

	// Test 404 response
	mocks.MockGET("/api/notfound", domtesting.MockResponse{
		Status: 404,
		Body:   "Not Found",
	})

	response, err := fetchAPI.fetch("/api/notfound", &RequestOptions{Method: "GET"})
	if err != nil {
		t.Fatalf("Fetch failed: %v", err)
	}

	if response.Status != 404 {
		t.Errorf("Expected status 404, got %d", response.Status)
	}

	if response.Ok {
		t.Error("Expected response.Ok to be false for 404")
	}
}

func TestFetchAPI_Headers(t *testing.T) {
	mocks := domtesting.NewNetworkMocks()
	fetchAPI := New(mocks)

	// Test response headers
	mocks.MockGET("/api/headers", domtesting.MockResponse{
		Status: 200,
		Body:   "test",
		Headers: map[string]string{
			"Content-Type":    "text/plain",
			"Cache-Control":   "no-cache",
			"X-Custom-Header": "custom-value",
		},
	})

	response, err := fetchAPI.fetch("/api/headers", &RequestOptions{Method: "GET"})
	if err != nil {
		t.Fatalf("Fetch failed: %v", err)
	}

	if response.Headers["Content-Type"] != "text/plain" {
		t.Errorf("Expected Content-Type header, got %s", response.Headers["Content-Type"])
	}

	if response.Headers["X-Custom-Header"] != "custom-value" {
		t.Errorf("Expected custom header, got %s", response.Headers["X-Custom-Header"])
	}
}

func TestFetchAPI_JavaScriptIntegration(t *testing.T) {
	vm := goja.New()
	mocks := domtesting.NewNetworkMocks()
	fetchAPI := New(mocks)

	// Bind fetch to JavaScript
	vm.Set("fetch", fetchAPI.CreateFetchFunction(vm))

	// Mock response
	mocks.MockGET("/api/js-test", domtesting.MockResponse{
		Status: 200,
		Body:   `{"success": true}`,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	})

	// Test simple JavaScript fetch call without complex options
	_, err := vm.RunString(`
		let promise = fetch('/api/js-test');
		// Just test that the call doesn't throw an error
		typeof promise;
	`)

	if err != nil {
		t.Fatalf("JavaScript execution failed: %v", err)
	}

	// This tests that the JavaScript binding works without complex object parsing
}
