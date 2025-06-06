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

// LoadHTML loads HTML content directly into the browser without making an HTTP request.
func (test *Test) LoadHTML(html string) *Test {
	test.harness.LoadHTML(html)
	return test
}

// ===== KEYBOARD EVENTS =====

// KeyDown simulates a keydown event.
func (test *Test) KeyDown(selector, key string) *Test {
	test.harness.KeyDown(selector, key)
	return test
}

// KeyUp simulates a keyup event.
func (test *Test) KeyUp(selector, key string) *Test {
	test.harness.KeyUp(selector, key)
	return test
}

// KeyPress simulates a keypress event (deprecated but still used).
func (test *Test) KeyPress(selector, key string) *Test {
	test.harness.KeyPress(selector, key)
	return test
}

// ===== ADVANCED MOUSE EVENTS =====

// MouseDown simulates a mousedown event.
func (test *Test) MouseDown(selector string) *Test {
	test.harness.MouseDown(selector)
	return test
}

// MouseUp simulates a mouseup event.
func (test *Test) MouseUp(selector string) *Test {
	test.harness.MouseUp(selector)
	return test
}

// MouseMove simulates a mousemove event.
func (test *Test) MouseMove(selector string) *Test {
	test.harness.MouseMove(selector)
	return test
}

// RightClick simulates a right-click (contextmenu event).
func (test *Test) RightClick(selector string) *Test {
	test.harness.RightClick(selector)
	return test
}

// ===== TOUCH EVENTS =====

// TouchStart simulates a touchstart event.
func (test *Test) TouchStart(selector string) *Test {
	test.harness.TouchStart(selector)
	return test
}

// TouchMove simulates a touchmove event.
func (test *Test) TouchMove(selector string) *Test {
	test.harness.TouchMove(selector)
	return test
}

// TouchEnd simulates a touchend event.
func (test *Test) TouchEnd(selector string) *Test {
	test.harness.TouchEnd(selector)
	return test
}

// ===== DRAG AND DROP EVENTS =====

// DragStart simulates a dragstart event.
func (test *Test) DragStart(selector string) *Test {
	test.harness.DragStart(selector)
	return test
}

// Drag simulates a drag event.
func (test *Test) Drag(selector string) *Test {
	test.harness.Drag(selector)
	return test
}

// DragEnd simulates a dragend event.
func (test *Test) DragEnd(selector string) *Test {
	test.harness.DragEnd(selector)
	return test
}

// DragOver simulates a dragover event.
func (test *Test) DragOver(selector string) *Test {
	test.harness.DragOver(selector)
	return test
}

// DragEnter simulates a dragenter event.
func (test *Test) DragEnter(selector string) *Test {
	test.harness.DragEnter(selector)
	return test
}

// DragLeave simulates a dragleave event.
func (test *Test) DragLeave(selector string) *Test {
	test.harness.DragLeave(selector)
	return test
}

// Drop simulates a drop event.
func (test *Test) Drop(selector string) *Test {
	test.harness.Drop(selector)
	return test
}

// ===== FORM EVENTS =====

// Reset simulates a reset event on a form.
func (test *Test) Reset(selector string) *Test {
	test.harness.Reset(selector)
	return test
}

// Invalid simulates an invalid event (form validation).
func (test *Test) Invalid(selector string) *Test {
	test.harness.Invalid(selector)
	return test
}

// ===== SCROLL EVENTS =====

// Scroll simulates a scroll event.
func (test *Test) Scroll(selector string) *Test {
	test.harness.Scroll(selector)
	return test
}

// ===== MEDIA EVENTS =====

// Play simulates a play event on media elements.
func (test *Test) Play(selector string) *Test {
	test.harness.Play(selector)
	return test
}

// Pause simulates a pause event on media elements.
func (test *Test) Pause(selector string) *Test {
	test.harness.Pause(selector)
	return test
}

// Ended simulates an ended event on media elements.
func (test *Test) Ended(selector string) *Test {
	test.harness.Ended(selector)
	return test
}

// ===== CLIPBOARD EVENTS =====

// Copy simulates a copy event.
func (test *Test) Copy(selector string) *Test {
	test.harness.Copy(selector)
	return test
}

// Cut simulates a cut event.
func (test *Test) Cut(selector string) *Test {
	test.harness.Cut(selector)
	return test
}

// Paste simulates a paste event.
func (test *Test) Paste(selector string) *Test {
	test.harness.Paste(selector)
	return test
}

// ===== WINDOW EVENTS =====

// Resize simulates a resize event on the window/document.
func (test *Test) Resize() *Test {
	test.harness.Resize()
	return test
}

// Load simulates a load event on the window/document.
func (test *Test) Load() *Test {
	test.harness.Load()
	return test
}

// Unload simulates an unload event on the window/document.
func (test *Test) Unload() *Test {
	test.harness.Unload()
	return test
}

// ===== ADVANCED INTERACTION HELPERS =====

// DragAndDrop performs a complete drag and drop operation from source to target.
func (test *Test) DragAndDrop(sourceSelector, targetSelector string) *Test {
	test.harness.DragAndDrop(sourceSelector, targetSelector)
	return test
}

// Tap simulates a mobile tap (touchstart + touchend).
func (test *Test) Tap(selector string) *Test {
	test.harness.Tap(selector)
	return test
}

// SwipeLeft simulates a left swipe gesture.
func (test *Test) SwipeLeft(selector string) *Test {
	test.harness.SwipeLeft(selector)
	return test
}

// SwipeRight simulates a right swipe gesture.
func (test *Test) SwipeRight(selector string) *Test {
	test.harness.SwipeRight(selector)
	return test
}

// LongPress simulates a long press gesture (extended touch).
func (test *Test) LongPress(selector string) *Test {
	test.harness.LongPress(selector)
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
