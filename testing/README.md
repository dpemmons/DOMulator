# DOMulator Testing Framework

A comprehensive testing framework for web applications that provides DOM manipulation, JavaScript execution, and HTTP integration testing capabilities. Built on top of DOMulator's DOM and JavaScript runtime.

## Features

- **DOM Testing**: Parse HTML, query elements with CSS selectors, and verify DOM state
- **JavaScript Integration**: Execute JavaScript code and test DOM interactions  
- **HTTP Integration**: Test HTTP handlers directly or make requests to running servers
- **Network Mocking**: Mock HTTP responses for isolated testing
- **User Interactions**: Simulate clicks, typing, form submissions, and other user actions
- **Fluent API**: Chainable assertions and interactions for readable tests

## Quick Start

```go
package main

import (
    "testing"
    "github.com/dpemmons/DOMulator/testing"
)

func TestSimpleExample(t *testing.T) {
    harness := testing.NewTestHarness()
    defer harness.Close()

    html := `
    <!DOCTYPE html>
    <html>
    <body>
        <h1 id="title">Hello World</h1>
        <button id="btn">Click me</button>
    </body>
    </html>`

    harness.LoadHTML(html)
    harness.AssertElement("#title").Exists().HasText("Hello World")
    harness.AssertElement("#btn").Exists()
    harness.Click("#btn")
}
```

## Core Components

### TestHarness

The main entry point for all testing operations. Provides methods for:

- Loading HTML content
- Making HTTP requests  
- Simulating user interactions
- Creating assertions
- Managing test lifecycle

```go
harness := testing.NewTestHarness()
defer harness.Close()

// Load HTML directly
harness.LoadHTML("<html><body>...</body></html>")

// Or navigate to a URL (requires handler/server setup)
harness.Navigate("/path")
```

### HTTP Integration

#### Direct Handler Testing

Test HTTP handlers directly without starting a server:

```go
mux := http.NewServeMux()
mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("<html><body>Hello</body></html>"))
})

harness := testing.NewTestHarness().SetHandler(mux)
harness.Navigate("/")
```

#### Server Testing  

Test against a running server:

```go
harness := testing.NewTestHarness().SetBaseURL("http://localhost:8080")
harness.Navigate("/api/endpoint")
```

#### Network Mocking

Mock HTTP responses for isolated testing:

```go
harness := testing.NewTestHarness()

// Mock a GET request
harness.mocks.MockGET("/api/users", testing.MockResponse{
    Status: 200,
    Headers: map[string]string{"Content-Type": "application/json"},
    Body: `[{"name": "John"}]`,
})

response := harness.GET("/api/users")
// response.Status == 200
```

### DOM Assertions

#### Element Assertions

Query and assert on DOM elements using CSS selectors:

```go
// Basic existence and counting
harness.AssertElement("div").Exists()
harness.AssertElement("div").NotExists()  
harness.AssertElement("li").Count(3)

// Element selection
harness.AssertElement("li").First().HasText("First item")
harness.AssertElement("li").Last().HasText("Last item")
harness.AssertElement("li").At(1).HasText("Second item")

// Text content
harness.AssertElement("#title").HasText("Exact text")
harness.AssertElement("#content").ContainsText("partial text")

// Attributes
harness.AssertElement("#link").HasAttribute("href", "/path")
harness.AssertElement("#input").HasAttributeContaining("class", "form")
harness.AssertElement("#div").HasClass("active")

// Form elements
harness.AssertElement("#input").HasValue("input value")
harness.AssertElement("#checkbox").IsChecked()
harness.AssertElement("#checkbox").IsNotChecked()

// Visibility (basic style attribute checking)
harness.AssertElement("#element").IsVisible()
harness.AssertElement("#element").IsHidden()
```

#### Document Assertions

Assert on document-level properties:

```go
harness.AssertDocument().HasTitle("Page Title")
harness.AssertDocument().ContainsText("text in body")
```

### User Interactions

Simulate user interactions with the page:

```go
// Click events
harness.Click("#button")
harness.DoubleClick("#element")

// Form interactions  
harness.Type("#input", "text to type")
harness.Check("#checkbox")
harness.Uncheck("#checkbox")
harness.Select("#select", "option-value")
harness.Submit("#form")

// Focus events
harness.Focus("#input")
harness.Blur("#input")

// Mouse events
harness.Hover("#element")
harness.Unhover("#element")

// Custom events
harness.TriggerEvent("#element", "custom-event")
```

### JavaScript Execution

Execute JavaScript code in the context of the loaded page:

```go
// Execute JavaScript
err := harness.ExecuteScript(`
    document.getElementById('title').textContent = 'New Title';
`)

// Verify the result
harness.AssertElement("#title").HasText("New Title")
```

### Direct DOM/Runtime Access

For advanced use cases, access the underlying DOM and JavaScript runtime:

```go
document := harness.Document()
runtime := harness.Runtime()

// Use DOM API directly
elements := css.QuerySelectorAll(document, "div")

// Use JavaScript runtime directly  
result, err := runtime.RunString("document.title")
```

## API Reference

### TestHarness Methods

#### Lifecycle
- `NewTestHarness() *TestHarness` - Create new test harness
- `Close() error` - Clean up resources

#### Configuration  
- `SetHandler(handler http.Handler) *TestHarness` - Use HTTP handler
- `SetBaseURL(url string) *TestHarness` - Set server base URL
- `SetTimeout(timeout time.Duration) *TestHarness` - Set default timeout
- `SetHeader(key, value string) *TestHarness` - Set default HTTP header

#### Content Loading
- `LoadHTML(html string) *TestHarness` - Load HTML content
- `Navigate(url string) *TestHarness` - Navigate to URL

#### HTTP Requests
- `GET(path string) *Response` - Make GET request
- `POST(path string, body interface{}) *Response` - Make POST request

#### DOM Assertions
- `AssertElement(selector string) *ElementAssertion` - Assert on element(s)
- `AssertDocument() *DocumentAssertion` - Assert on document

#### User Interactions
- `Click(selector string) *TestHarness` - Click element
- `DoubleClick(selector string) *TestHarness` - Double-click element  
- `Type(selector, text string) *TestHarness` - Type into element
- `Focus(selector string) *TestHarness` - Focus element
- `Blur(selector string) *TestHarness` - Blur element
- `Submit(selector string) *TestHarness` - Submit form
- `Check(selector string) *TestHarness` - Check checkbox/radio
- `Uncheck(selector string) *TestHarness` - Uncheck checkbox
- `Select(selector, value string) *TestHarness` - Select option
- `Hover(selector string) *TestHarness` - Hover over element
- `Unhover(selector string) *TestHarness` - Move mouse away
- `TriggerEvent(selector, eventType string) *TestHarness` - Trigger custom event

#### JavaScript
- `ExecuteScript(js string) error` - Execute JavaScript code
- `Document() *dom.Document` - Get DOM document
- `Runtime() *js.Runtime` - Get JavaScript runtime

### ElementAssertion Methods

#### Existence and Counting
- `Exists() *ElementAssertion` - Assert element exists
- `NotExists() *ElementAssertion` - Assert element doesn't exist
- `Count(expected int) *ElementAssertion` - Assert element count

#### Element Selection
- `First() *ElementAssertion` - Get first matching element
- `Last() *ElementAssertion` - Get last matching element  
- `At(index int) *ElementAssertion` - Get element at index

#### Text Content
- `HasText(expected string) *ElementAssertion` - Assert exact text
- `ContainsText(expected string) *ElementAssertion` - Assert contains text

#### Attributes
- `HasAttribute(name, expected string) *ElementAssertion` - Assert attribute value
- `HasAttributeContaining(name, expected string) *ElementAssertion` - Assert attribute contains
- `HasClass(className string) *ElementAssertion` - Assert has CSS class

#### Form Elements
- `HasValue(expected string) *ElementAssertion` - Assert input value
- `IsChecked() *ElementAssertion` - Assert checkbox is checked
- `IsNotChecked() *ElementAssertion` - Assert checkbox is not checked

#### Visibility
- `IsVisible() *ElementAssertion` - Assert element is visible
- `IsHidden() *ElementAssertion` - Assert element is hidden

### DocumentAssertion Methods

- `HasTitle(expected string) *DocumentAssertion` - Assert document title
- `ContainsText(expected string) *DocumentAssertion` - Assert body contains text

### NetworkMocks Methods

- `MockGET(pattern string, response MockResponse)` - Mock GET request
- `MockPOST(pattern string, response MockResponse)` - Mock POST request
- `MockPUT(pattern string, response MockResponse)` - Mock PUT request
- `MockDELETE(pattern string, response MockResponse)` - Mock DELETE request
- `Mock(method, pattern string, response MockResponse)` - Mock any method
- `Clear()` - Clear all mocks
- `Reset()` - Reset all mocks (alias for Clear)

### Response Type

```go
type Response struct {
    Status     int                 // HTTP status code
    StatusText string             // HTTP status text
    Headers    map[string]string  // Response headers
    Body       string             // Response body
    Error      error              // Request error if any
}
```

### MockResponse Type

```go
type MockResponse struct {
    Status  int                 // HTTP status code to return
    Headers map[string]string  // Response headers to return
    Body    string             // Response body to return
}
```

## Best Practices

### Test Organization

1. **Use descriptive test names** that explain what is being tested
2. **Group related assertions** together for clarity
3. **Clean up resources** using `defer harness.Close()`
4. **Use meaningful selectors** that reflect the semantic purpose

### Error Handling

The framework uses panics for assertion failures to integrate naturally with Go's testing framework. Failed assertions will cause the test to fail with a descriptive message.

### Performance

- **Reuse harnesses** when possible for multiple related tests
- **Use direct HTML loading** for unit tests (faster than HTTP)
- **Use HTTP integration** for integration tests
- **Mock external dependencies** to keep tests fast and reliable

### Debugging

- **Use `harness.ExecuteScript()`** to inspect JavaScript state
- **Access `harness.Document()`** to examine DOM directly
- **Print response bodies** when debugging HTTP interactions
- **Use simple CSS selectors** for better error messages

## Examples

See `examples_test.go` for comprehensive examples including:

- Direct HTML testing
- HTTP handler integration  
- Network mocking
- Complex JavaScript interactions
- Form validation testing

## Architecture

The testing framework is built on several core DOMulator components:

- **DOM**: Full W3C-compatible DOM implementation
- **CSS Selectors**: CSS selector engine for element querying
- **JavaScript Runtime**: Goja-based JavaScript execution with DOM bindings
- **HTML Parser**: Standards-compliant HTML parsing
- **Event System**: DOM event propagation and handling

This architecture provides a realistic testing environment that closely matches browser behavior while being fast enough for unit testing.
