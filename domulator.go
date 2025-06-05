// Package domulator provides a fast, lightweight DOM emulation and JavaScript runtime
// for testing web applications in Go. Run JavaScript frameworks like HTMX, Alpine.js,
// or vanilla JavaScript without a headless browser.
package domulator

import (
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/dpemmons/DOMulator/internal/dom"
	"github.com/dpemmons/DOMulator/internal/js"
	testingpkg "github.com/dpemmons/DOMulator/testing"
)

// Test represents a browser-like test environment for testing web applications.
// It provides a simple API for navigating to pages, interacting with elements,
// and making assertions about the resulting DOM state.
type Test struct {
	t       *testing.T
	harness *testingpkg.TestHarness
	server  *httptest.Server
}

// NewTest creates a new DOMulator test instance.
// This is the main entry point for browser-like testing.
func NewTest(t *testing.T) *Test {
	return &Test{
		t:       t,
		harness: testingpkg.NewTestHarness(),
	}
}

// WithServer configures the test to use the provided httptest server.
// This allows you to use relative URLs in Navigate() calls.
// The server URL will be automatically prepended to relative paths.
func (test *Test) WithServer(server *httptest.Server) *Test {
	test.server = server
	// Configure the underlying harness to use this server's base URL
	test.harness.SetBaseURL(server.URL)
	return test
}

// Navigate loads a page from the specified URL and initializes the DOM.
// If a server has been configured with WithServer(), relative URLs will
// be automatically converted to absolute URLs using the server's base URL.
//
// This method automatically:
// - Fetches the HTML content
// - Parses it into a DOM tree
// - Loads and executes any <script> tags found
// - Loads CSS stylesheets and other resources
//
// Example:
//
//	test.Navigate("/login")  // Uses test server if configured
//	test.Navigate("http://localhost:3000/login")  // Direct URL
func (test *Test) Navigate(url string) *Test {
	test.harness.Navigate(url)
	return test
}

// ExecuteScript executes JavaScript code in the current browser context.
// The script has access to the full DOM API and any previously loaded scripts.
func (test *Test) ExecuteScript(script string) *Test {
	if err := test.harness.ExecuteScript(script); err != nil {
		test.t.Fatalf("Failed to execute script: %v", err)
	}
	return test
}

// Document returns the current DOM document for advanced use cases.
// Most users should prefer the higher-level interaction and assertion methods.
func (test *Test) Document() *dom.Document {
	return test.harness.Document()
}

// Runtime returns the JavaScript runtime for advanced use cases.
// Most users should prefer ExecuteScript() for running JavaScript.
func (test *Test) Runtime() *js.Runtime {
	return test.harness.Runtime()
}

// Close cleans up resources used by the test instance.
// This is automatically called when the test completes, but can be
// called manually if needed.
func (test *Test) Close() error {
	return test.harness.Close()
}

// SetTimeout configures the default timeout for wait operations.
func (test *Test) SetTimeout(timeout time.Duration) *Test {
	test.harness.SetTimeout(timeout)
	return test
}

// SetHeader sets a default header for all HTTP requests made by the browser.
func (test *Test) SetHeader(key, value string) *Test {
	test.harness.SetHeader(key, value)
	return test
}

// Interaction Methods - These methods simulate user interactions with the page

// Click simulates a click event on the first element matching the selector.
func (test *Test) Click(selector string) *Test {
	test.harness.Click(selector)
	return test
}

// DoubleClick simulates a double-click event on the first element matching the selector.
func (test *Test) DoubleClick(selector string) *Test {
	test.harness.DoubleClick(selector)
	return test
}

// Type simulates typing text into an input element.
// This will clear any existing content and replace it with the provided text.
func (test *Test) Type(selector, text string) *Test {
	test.harness.Type(selector, text)
	return test
}

// Focus simulates focusing on an element (clicking into an input field, etc.).
func (test *Test) Focus(selector string) *Test {
	test.harness.Focus(selector)
	return test
}

// Blur simulates removing focus from an element.
func (test *Test) Blur(selector string) *Test {
	test.harness.Blur(selector)
	return test
}

// Submit simulates submitting a form.
func (test *Test) Submit(selector string) *Test {
	test.harness.Submit(selector)
	return test
}

// Check simulates checking a checkbox or radio button.
func (test *Test) Check(selector string) *Test {
	test.harness.Check(selector)
	return test
}

// Uncheck simulates unchecking a checkbox.
func (test *Test) Uncheck(selector string) *Test {
	test.harness.Uncheck(selector)
	return test
}

// Select simulates selecting an option in a select element.
func (test *Test) Select(selector, value string) *Test {
	test.harness.Select(selector, value)
	return test
}

// Hover simulates hovering over an element with the mouse.
func (test *Test) Hover(selector string) *Test {
	test.harness.Hover(selector)
	return test
}

// Unhover simulates moving the mouse away from an element.
func (test *Test) Unhover(selector string) *Test {
	test.harness.Unhover(selector)
	return test
}

// TriggerEvent dispatches a custom event on the first element matching the selector.
func (test *Test) TriggerEvent(selector, eventType string) *Test {
	test.harness.TriggerEvent(selector, eventType)
	return test
}

// Assertion Methods - These methods provide fluent assertions about the DOM state

// AssertElement creates an element assertion for the given CSS selector.
// Returns a fluent assertion object that can be chained with various checks.
//
// Example:
//
//	test.AssertElement("#submit-btn").Exists().HasText("Submit").HasClass("primary")
func (test *Test) AssertElement(selector string) *testingpkg.ElementAssertion {
	return test.harness.AssertElement(selector)
}

// AssertDocument creates a document-level assertion for checking document properties.
// Returns a fluent assertion object for document-wide checks.
//
// Example:
//
//	test.AssertDocument().HasTitle("My App").ContainsText("Welcome")
func (test *Test) AssertDocument() *testingpkg.DocumentAssertion {
	return test.harness.AssertDocument()
}

// HTTP Methods - These methods allow making direct HTTP requests for API testing

// GET performs a GET request to the specified path and returns the response.
// If a server is configured, relative paths will be resolved against the server URL.
func (test *Test) GET(path string) *testingpkg.Response {
	if test.server != nil && !strings.HasPrefix(path, "http") {
		path = test.server.URL + path
	}
	return test.harness.GET(path)
}

// POST performs a POST request with the given body and returns the response.
// If a server is configured, relative paths will be resolved against the server URL.
func (test *Test) POST(path string, body interface{}) *testingpkg.Response {
	if test.server != nil && !strings.HasPrefix(path, "http") {
		path = test.server.URL + path
	}
	return test.harness.POST(path, body)
}
