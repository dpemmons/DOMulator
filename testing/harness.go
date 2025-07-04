package testing

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/dpemmons/DOMulator/internal/dom"
	"github.com/dpemmons/DOMulator/internal/js"
	htmlparser "github.com/dpemmons/DOMulator/internal/parser"
)

// TestHarness provides a high-level API for testing web applications
// with DOMulator. It supports both direct handler integration and
// server-based testing for maximum flexibility.
type TestHarness struct {
	client          *HTTPClient
	domulator       *DOMulator
	mocks           *NetworkMocks
	config          *Config
	consoleCallback func(js.ConsoleLevel, []interface{}) // Stored callback
	errorCallback   func(*js.JavaScriptError)            // Stored callback
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
	DebugMode      bool // Controls JavaScript console verbosity
	WindowWidth    int  // Initial window width (default: 800)
	WindowHeight   int  // Initial window height (default: 600)
	ElementWidth   int  // Default element width for ResizeObserver (default: 100)
	ElementHeight  int  // Default element height for ResizeObserver (default: 100)
}

// NewTestHarness creates a new test harness instance
func NewTestHarness() *TestHarness {
	config := &Config{
		DefaultTimeout: 5 * time.Second,
		Headers:        make(map[string]string),
		WindowWidth:    800,
		WindowHeight:   600,
		ElementWidth:   100,
		ElementHeight:  100,
	}

	mocks := NewNetworkMocks()
	client := NewHTTPClient(mocks)

	return &TestHarness{
		client: client,
		mocks:  mocks,
		config: config,
	}
}

// NewTestHarnessWithConfig creates a new test harness instance with custom configuration
func NewTestHarnessWithConfig(config *Config) *TestHarness {
	// Set defaults for any missing values
	if config.DefaultTimeout == 0 {
		config.DefaultTimeout = 5 * time.Second
	}
	if config.Headers == nil {
		config.Headers = make(map[string]string)
	}
	if config.WindowWidth == 0 {
		config.WindowWidth = 800
	}
	if config.WindowHeight == 0 {
		config.WindowHeight = 600
	}
	if config.ElementWidth == 0 {
		config.ElementWidth = 100
	}
	if config.ElementHeight == 0 {
		config.ElementHeight = 100
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

	// Set window dimensions from config
	if h.config.WindowWidth > 0 && h.config.WindowHeight > 0 {
		runtime.SetWindowDimensions(h.config.WindowWidth, h.config.WindowHeight)
	}

	// Set up event loop for proper asynchronous behavior
	runtime.SetupEventLoop(true) // Use virtual time for deterministic tests

	// Store in domulator wrapper
	h.domulator = &DOMulator{
		document: document,
		runtime:  runtime,
	}

	// Apply debug mode if it was set
	runtime.SetDebugMode(h.config.DebugMode)

	// Apply stored callbacks
	h.applyCallbacks()

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

	// Set window dimensions from config
	if h.config.WindowWidth > 0 && h.config.WindowHeight > 0 {
		runtime.SetWindowDimensions(h.config.WindowWidth, h.config.WindowHeight)
	}

	// Set up event loop for proper asynchronous behavior
	runtime.SetupEventLoop(true) // Use virtual time for deterministic tests

	// Store in domulator wrapper
	h.domulator = &DOMulator{
		document: document,
		runtime:  runtime,
	}

	// Apply debug mode if it was set
	runtime.SetDebugMode(h.config.DebugMode)

	// Apply stored callbacks
	h.applyCallbacks()

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

// SetDebugMode enables or disables verbose console output from JavaScript
func (h *TestHarness) SetDebugMode(debug bool) *TestHarness {
	h.config.DebugMode = debug
	if h.domulator != nil && h.domulator.runtime != nil {
		h.domulator.runtime.SetDebugMode(debug)
	}
	return h
}

// AdvanceTime advances virtual time by the specified duration
func (h *TestHarness) AdvanceTime(duration time.Duration) {
	if h.domulator == nil || h.domulator.runtime == nil {
		panic("no document loaded - call LoadHTML() or Navigate() first")
	}

	eventLoop := h.domulator.runtime.EventLoop()
	if eventLoop == nil {
		panic("event loop not initialized")
	}

	eventLoop.AdvanceTime(duration)
}

// FlushMicrotasks processes all pending microtasks immediately
func (h *TestHarness) FlushMicrotasks() {
	if h.domulator == nil || h.domulator.runtime == nil {
		panic("no document loaded - call LoadHTML() or Navigate() first")
	}

	eventLoop := h.domulator.runtime.EventLoop()
	if eventLoop == nil {
		panic("event loop not initialized")
	}

	eventLoop.ProcessMicrotasks()
}

// ProcessPendingTimers executes all timers that are ready to fire
func (h *TestHarness) ProcessPendingTimers() {
	if h.domulator == nil || h.domulator.runtime == nil {
		panic("no document loaded - call LoadHTML() or Navigate() first")
	}

	eventLoop := h.domulator.runtime.EventLoop()
	if eventLoop == nil {
		panic("event loop not initialized")
	}

	// Process one task if available (timers, animation frames, etc.)
	eventLoop.ProcessNextTask()
}

// ResizeWindow sets the window dimensions and fires a resize event
func (h *TestHarness) ResizeWindow(width, height int) *TestHarness {
	if h.domulator == nil || h.domulator.runtime == nil {
		panic("no document loaded - call LoadHTML() or Navigate() first")
	}

	// Update window dimensions in JavaScript runtime
	h.domulator.runtime.SetWindowDimensions(width, height)

	// Fire resize event on window
	h.ExecuteScript(`
		var event = new Event('resize', { bubbles: false, cancelable: false });
		window.dispatchEvent(event);
	`)

	return h
}

// SetConsoleCallback sets a callback for console output
func (h *TestHarness) SetConsoleCallback(cb func(js.ConsoleLevel, []interface{})) *TestHarness {
	h.consoleCallback = cb
	if h.domulator != nil && h.domulator.runtime != nil {
		h.domulator.runtime.SetConsoleCallback(cb)
	}
	return h
}

// SetErrorCallback sets a callback for JavaScript errors
func (h *TestHarness) SetErrorCallback(cb func(*js.JavaScriptError)) *TestHarness {
	h.errorCallback = cb
	if h.domulator != nil && h.domulator.runtime != nil {
		h.domulator.runtime.SetErrorCallback(cb)
	}
	return h
}

// applyCallbacks applies stored callbacks to the runtime
func (h *TestHarness) applyCallbacks() {
	if h.domulator != nil && h.domulator.runtime != nil {
		if h.consoleCallback != nil {
			h.domulator.runtime.SetConsoleCallback(h.consoleCallback)
		}
		if h.errorCallback != nil {
			h.domulator.runtime.SetErrorCallback(h.errorCallback)
		}
	}
}

// CaptureConsole creates and sets up a ConsoleCapture helper with a testing.TB
func (h *TestHarness) CaptureConsole(t testing.TB) *ConsoleCapture {
	capture := NewConsoleCapture(t)
	h.SetConsoleCallback(capture.Callback)
	return capture
}

// CaptureErrors creates and sets up an ErrorCapture helper with a testing.TB
func (h *TestHarness) CaptureErrors(t testing.TB) *ErrorCapture {
	capture := NewErrorCapture(t)
	h.SetErrorCallback(capture.Callback)
	return capture
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

	// Process any queued microtasks (e.g., Alpine.js queues itself to start)
	if h.domulator.runtime.EventLoop() != nil {
		h.domulator.runtime.EventLoop().ProcessMicrotasks()
	}

	// Fire DOMContentLoaded event after all scripts have been loaded and executed
	// This is when scripts can attach event listeners to DOM elements
	h.domulator.document.FireDOMContentLoaded()

	// Process microtasks again after DOMContentLoaded
	if h.domulator.runtime.EventLoop() != nil {
		h.domulator.runtime.EventLoop().ProcessMicrotasks()
	}

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
