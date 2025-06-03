package testing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
)

// HTTPClient handles HTTP requests for both handler and server integration modes
type HTTPClient struct {
	handler  http.Handler
	baseURL  string
	headers  map[string]string
	mocks    *NetworkMocks
	recorder *httptest.ResponseRecorder
}

// Response represents an HTTP response with additional metadata
type Response struct {
	Status     int
	StatusText string
	Headers    map[string]string
	Body       string
	Error      error
}

// NewHTTPClient creates a new HTTP client with network mocking support
func NewHTTPClient(mocks *NetworkMocks) *HTTPClient {
	return &HTTPClient{
		headers: make(map[string]string),
		mocks:   mocks,
	}
}

// SetHandler configures the client to use a direct HTTP handler
func (c *HTTPClient) SetHandler(handler http.Handler) {
	c.handler = handler
	c.baseURL = "" // Clear base URL when using handler
}

// SetBaseURL configures the client to make requests to a real server
func (c *HTTPClient) SetBaseURL(baseURL string) {
	c.baseURL = strings.TrimSuffix(baseURL, "/")
	c.handler = nil // Clear handler when using base URL
}

// SetHeader sets a default header for all requests
func (c *HTTPClient) SetHeader(key, value string) {
	c.headers[key] = value
}

// GET performs a GET request
func (c *HTTPClient) GET(path string) *Response {
	return c.request("GET", path, nil)
}

// POST performs a POST request with the given body
func (c *HTTPClient) POST(path string, body interface{}) *Response {
	var reader io.Reader

	if body != nil {
		switch v := body.(type) {
		case string:
			reader = strings.NewReader(v)
		case []byte:
			reader = bytes.NewReader(v)
		case io.Reader:
			reader = v
		default:
			// Try to JSON marshal
			jsonData, err := json.Marshal(body)
			if err != nil {
				return &Response{
					Error: fmt.Errorf("failed to marshal request body: %w", err),
				}
			}
			reader = bytes.NewReader(jsonData)
		}
	}

	return c.request("POST", path, reader)
}

// request performs the actual HTTP request using either handler or HTTP client
func (c *HTTPClient) request(method, path string, body io.Reader) *Response {
	// First check if we have a mock for this request
	if mockResponse := c.mocks.Match(method, path); mockResponse != nil {
		return &Response{
			Status:     mockResponse.Status,
			StatusText: http.StatusText(mockResponse.Status),
			Headers:    mockResponse.Headers,
			Body:       mockResponse.Body,
		}
	}

	if c.handler != nil {
		return c.requestViaHandler(method, path, body)
	}

	if c.baseURL != "" {
		return c.requestViaHTTP(method, path, body)
	}

	return &Response{
		Error: fmt.Errorf("no handler or base URL configured"),
	}
}

// requestViaHandler handles requests using a direct HTTP handler
func (c *HTTPClient) requestViaHandler(method, path string, body io.Reader) *Response {
	// Create request
	req := httptest.NewRequest(method, path, body)

	// Add headers
	for key, value := range c.headers {
		req.Header.Set(key, value)
	}

	// Create response recorder
	recorder := httptest.NewRecorder()

	// Execute handler
	c.handler.ServeHTTP(recorder, req)

	// Convert headers
	headers := make(map[string]string)
	for key, values := range recorder.Header() {
		if len(values) > 0 {
			headers[key] = values[0]
		}
	}

	return &Response{
		Status:     recorder.Code,
		StatusText: http.StatusText(recorder.Code),
		Headers:    headers,
		Body:       recorder.Body.String(),
	}
}

// requestViaHTTP handles requests using the standard HTTP client
func (c *HTTPClient) requestViaHTTP(method, path string, body io.Reader) *Response {
	// Build full URL
	fullURL := c.baseURL + path
	if !strings.HasPrefix(path, "/") {
		fullURL = c.baseURL + "/" + path
	}

	// Parse URL to validate
	parsedURL, err := url.Parse(fullURL)
	if err != nil {
		return &Response{
			Error: fmt.Errorf("invalid URL: %w", err),
		}
	}

	// Create request
	req, err := http.NewRequest(method, parsedURL.String(), body)
	if err != nil {
		return &Response{
			Error: fmt.Errorf("failed to create request: %w", err),
		}
	}

	// Add headers
	for key, value := range c.headers {
		req.Header.Set(key, value)
	}

	// Execute request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return &Response{
			Error: fmt.Errorf("request failed: %w", err),
		}
	}
	defer resp.Body.Close()

	// Read response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return &Response{
			Error: fmt.Errorf("failed to read response body: %w", err),
		}
	}

	// Convert headers
	headers := make(map[string]string)
	for key, values := range resp.Header {
		if len(values) > 0 {
			headers[key] = values[0]
		}
	}

	return &Response{
		Status:     resp.StatusCode,
		StatusText: resp.Status,
		Headers:    headers,
		Body:       string(bodyBytes),
	}
}

// LoadHTML is a convenience method that loads the response body as HTML
// This method is designed to be chainable with harness operations
func (r *Response) LoadHTML() string {
	if r.Error != nil {
		panic(fmt.Sprintf("Cannot load HTML from failed response: %v", r.Error))
	}
	return r.Body
}
