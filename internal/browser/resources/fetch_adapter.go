package resources

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/dpemmons/DOMulator/internal/browser/fetch"
)

// FetchAdapter adapts the existing fetch.FetchAPI to our FetchAPI interface
type FetchAdapter struct {
	fetchAPI *fetch.FetchAPI
}

// NewFetchAdapter creates a new fetch adapter
func NewFetchAdapter(fetchAPI *fetch.FetchAPI) *FetchAdapter {
	return &FetchAdapter{
		fetchAPI: fetchAPI,
	}
}

// Fetch performs an HTTP request and returns the response
func (fa *FetchAdapter) Fetch(url string, options map[string]interface{}) (*fetch.Response, error) {
	// Use the real fetch API if available
	if fa.fetchAPI != nil {
		// Convert options to RequestOptions and call the private fetch method
		requestOptions := &fetch.RequestOptions{
			Method:  "GET",
			Headers: make(map[string]string),
		}

		if options != nil {
			if method, ok := options["method"].(string); ok {
				requestOptions.Method = method
			}
			if headers, ok := options["headers"].(map[string]string); ok {
				requestOptions.Headers = headers
			}
		}

		// Use the real HTTP client to fetch the URL
		return fa.realFetch(url, requestOptions)
	}

	// Fall back to basic fetch for testing when no fetch API is provided
	return fa.basicFetch(url, options)
}

// realFetch performs a real HTTP request using Go's http package
func (fa *FetchAdapter) realFetch(url string, options *fetch.RequestOptions) (*fetch.Response, error) {
	// Create HTTP client
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	// Create request body
	var body io.Reader
	if options.Body != "" {
		body = strings.NewReader(options.Body)
	}

	// Create HTTP request
	req, err := http.NewRequest(options.Method, url, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	for key, value := range options.Headers {
		req.Header.Set(key, value)
	}

	// Perform request
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Convert headers
	headers := make(map[string]string)
	for key, values := range resp.Header {
		if len(values) > 0 {
			headers[key] = values[0]
		}
	}

	return &fetch.Response{
		Status:     resp.StatusCode,
		StatusText: resp.Status,
		Headers:    headers,
		Body:       string(bodyBytes),
		URL:        url,
		Ok:         resp.StatusCode >= 200 && resp.StatusCode < 300,
	}, nil
}

// basicFetch is a simplified fetch implementation for resource loading
func (fa *FetchAdapter) basicFetch(url string, options map[string]interface{}) (*fetch.Response, error) {
	// This is a simplified version - we'll enhance it later
	// For now, we just return a mock response to enable testing
	method := "GET"
	if options != nil {
		if m, ok := options["method"].(string); ok {
			method = m
		}
	}

	// Create a basic response based on URL patterns
	body := `console.log("Script loaded via resource loader from: ` + url + `");`

	// Handle different file types
	if method == "GET" {
		// You could add logic here to return different content based on URL
		// For now, return JavaScript content for all requests
	}

	return &fetch.Response{
		Status:     200,
		StatusText: "OK",
		Headers:    make(map[string]string),
		Body:       body,
		URL:        url,
		Ok:         true,
	}, nil
}
