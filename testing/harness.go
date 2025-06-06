package testing

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dpemmons/DOMulator/internal/dom"
	"github.com/dpemmons/DOMulator/internal/js"
	htmlparser "github.com/dpemmons/DOMulator/internal/parser"
)

// TestHarness provides a high-level API for testing web applications
// with DOMulator. It supports both direct handler integration and
// server-based testing for maximum flexibility.
type TestHarness struct {
	client    *HTTPClient
	domulator *DOMulator
	mocks     *NetworkMocks
	config    *Config
}

// DOMulator wraps the core DOMulator functionality for testing
type DOMulator struct {
	document *dom.Document
	runtime  *js.Runtime
}

// Config holds configuration options for the test harness
type Config struct {
	DefaultTimeout time.Duration
	BaseURL        string
	Headers        map[string]string
}

// NewTestHarness creates a new test harness instance
func NewTestHarness() *TestHarness {
	config := &Config{
		DefaultTimeout: 5 * time.Second,
		Headers:        make(map[string]string),
	}

	mocks := NewNetworkMocks()
	client := NewHTTPClient(mocks)

	return &TestHarness{
		client: client,
		mocks:  mocks,
		config: config,
	}
}

// SetHandler configures the harness to use a direct HTTP handler
// for requests. This is the fastest integration method for unit tests.
func (h *TestHarness) SetHandler(handler http.Handler) *TestHarness {
	h.client.SetHandler(handler)
	return h
}

// SetBaseURL configures the harness to make HTTP requests to a real server.
// This is useful for integration tests or testing against running servers.
func (h *TestHarness) SetBaseURL(url string) *TestHarness {
	h.config.BaseURL = url
	h.client.SetBaseURL(url)
	return h
}

// Navigate loads a page from the specified URL and initializes the DOM.
// This method automatically loads and executes any <script> tags found in the HTML,
// similar to how a real browser would behave.
func (h *TestHarness) Navigate(url string) *TestHarness {
	// Handle absolute URLs by automatically configuring the client for HTTP mode
	fullURL := url
	if strings.HasPrefix(url, "http") {
		// For absolute URLs, we need to ensure the client is configured to use HTTP mode
		// Extract the base URL (scheme + host + port) for future relative requests
		if h.config.BaseURL == "" && h.client.handler == nil {
			// Parse the URL to extract the base
			if idx := strings.Index(url[8:], "/"); idx != -1 { // Skip "http://" or "https://"
				baseURL := url[:8+idx] // Include scheme + host + port
				h.SetBaseURL(baseURL)
			} else {
				// URL has no path, use it as base URL
				h.SetBaseURL(url)
			}
		}
		fullURL = url
	} else if h.config.BaseURL != "" {
		// Resolve relative URLs using base URL if configured
		baseURL := strings.TrimSuffix(h.config.BaseURL, "/")
		fullURL = baseURL + url
	}

	response := h.client.GET(fullURL) // Use the resolved full URL
	if response.Error != nil {
		panic(fmt.Sprintf("Failed to navigate to %s: %v", fullURL, response.Error))
	}

	// Load HTML directly without executing scripts yet (LoadHTML would execute them)
	// Parse HTML into DOM using the parser
	parser := htmlparser.NewParser()
	document, err := parser.Parse(strings.NewReader(response.Body))
	if err != nil {
		panic(fmt.Sprintf("Failed to parse HTML: %v", err))
	}

	// Create JavaScript runtime with DOM integration
	runtime := js.New(document)

	// Store in domulator wrapper
	h.domulator = &DOMulator{
		document: document,
		runtime:  runtime,
	}

	// Set initial document readyState to "loading"
	document.SetReadyState("loading")

	// Now load and execute scripts with the proper base URL for external scripts
	h.loadPageResources(fullURL)

	// Fire window.load event after all resources are loaded
	if h.domulator != nil && h.domulator.document != nil {
		h.domulator.document.FireWindowLoad()
	}

	return h
}

// GET performs a GET request and returns the response
func (h *TestHarness) GET(path string) *Response {
	return h.client.GET(path)
}

// POST performs a POST request with the given body and returns the response
func (h *TestHarness) POST(path string, body interface{}) *Response {
	return h.client.POST(path, body)
}

// LoadHTML loads HTML content into the DOMulator environment
func (h *TestHarness) LoadHTML(html string) *TestHarness {
	// Parse HTML into DOM using the parser
	parser := htmlparser.NewParser()
	document, err := parser.Parse(strings.NewReader(html))
	if err != nil {
		panic(fmt.Sprintf("Failed to parse HTML: %v", err))
	}

	// Create JavaScript runtime with DOM integration
	runtime := js.New(document)

	// Store in domulator wrapper
	h.domulator = &DOMulator{
		document: document,
		runtime:  runtime,
	}

	// Set initial document readyState to "loading"
	document.SetReadyState("loading")

	// Execute scripts just like Navigate() does
	h.loadPageResources("")

	// Fire window.load event after all resources are loaded
	h.domulator.document.FireWindowLoad()

	return h
}

// ExecuteScript executes JavaScript code in the current context
func (h *TestHarness) ExecuteScript(js string) error {
	if h.domulator == nil {
		return fmt.Errorf("no document loaded - call LoadHTML() or Navigate() first")
	}

	_, err := h.domulator.runtime.RunString(js)
	return err
}

// Document returns the current DOM document (for advanced use cases)
func (h *TestHarness) Document() *dom.Document {
	if h.domulator == nil {
		return nil
	}
	return h.domulator.document
}

// Runtime returns the JavaScript runtime (for advanced use cases)
func (h *TestHarness) Runtime() *js.Runtime {
	if h.domulator == nil {
		return nil
	}
	return h.domulator.runtime
}

// Close cleans up resources used by the harness
func (h *TestHarness) Close() error {
	if h.domulator != nil && h.domulator.runtime != nil {
		h.domulator.runtime.Shutdown()
	}
	return nil
}

// SetTimeout configures the default timeout for wait operations
func (h *TestHarness) SetTimeout(timeout time.Duration) *TestHarness {
	h.config.DefaultTimeout = timeout
	return h
}

// SetHeader sets a default header for all HTTP requests
func (h *TestHarness) SetHeader(key, value string) *TestHarness {
	h.config.Headers[key] = value
	h.client.SetHeader(key, value)
	return h
}

// loadPageResources automatically loads and executes scripts found in the DOM,
// similar to how a real browser would behave when loading a page.
func (h *TestHarness) loadPageResources(baseURL string) {
	if h.domulator == nil || h.domulator.document == nil {
		return
	}

	// Find all script tags
	scripts := h.domulator.document.QuerySelectorAll("script")

	for i := 0; i < scripts.Length(); i++ {
		script := scripts.Item(i).(*dom.Element)

		// Check if script has a src attribute (external script)
		if src := script.GetAttribute("src"); src != "" {
			h.loadExternalScript(src, baseURL)
		} else {
			// Inline script - execute the content directly
			scriptContent := script.TextContent()
			if scriptContent != "" {
				if err := h.ExecuteScript(scriptContent); err != nil {
					// Log error but don't panic - scripts can fail in real browsers too
					fmt.Printf("Warning: Failed to execute inline script: %v\n", err)
				}
			}
		}
	}

	// Fire DOMContentLoaded event after all scripts have been loaded and executed
	// This is when scripts can attach event listeners to DOM elements
	h.domulator.document.FireDOMContentLoaded()

	// Set document readyState to "complete"
	h.domulator.document.SetReadyState("complete")
}

// loadExternalScript fetches and executes an external JavaScript file
func (h *TestHarness) loadExternalScript(src, baseURL string) {
	// Handle relative URLs
	scriptURL := src
	if !strings.HasPrefix(src, "http") {
		if strings.HasSuffix(baseURL, "/") {
			scriptURL = baseURL + strings.TrimPrefix(src, "/")
		} else {
			// Parse base URL to get the base path
			if strings.Contains(baseURL, "/") {
				lastSlash := strings.LastIndex(baseURL, "/")
				baseURL = baseURL[:lastSlash+1]
			} else {
				baseURL = baseURL + "/"
			}
			scriptURL = baseURL + strings.TrimPrefix(src, "/")
		}
	}

	// Fetch the script content
	response := h.client.GET(scriptURL)
	if response.Error != nil {
		fmt.Printf("Warning: Failed to load script from %s: %v\n", scriptURL, response.Error)
		return
	}

	// Execute the script
	if err := h.ExecuteScript(response.Body); err != nil {
		fmt.Printf("Warning: Failed to execute script from %s: %v\n", scriptURL, err)
	}
}
