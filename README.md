# DOMulator

Fast, lightweight DOM emulation and JavaScript runtime for testing web applications in Go. Run JavaScript frameworks like HTMX, Alpine.js, or vanilla JavaScript in-process without a headless browser.

This project is intended to support vide-coded Go/HTMX projects.

> [!IMPORTANT]
> This project is itself an experiment in ***vibe-coding***. I don't know yet how well it works and it will almost certainly change. It's primarily vibed from WHATWG standards but I do not recommend anyone use it until I've incorporated browser conformance tests.

## ✨ Features

- 🚀 **High Performance** - Keep client/server tests in-process; should be significantly faster than testing with a headless browser.
- 🎯 **Real JavaScript** - Powered by Goja (ECMAScript 5.1+)
- 🔧 **DOM API Implementation** - Aims to provide a comprehensive DOM API.
  - Supports core DOM manipulation, event handling, and an internal CSS selector engine.
  - The selector engine ([`internal/css/selectors.go`](internal/css/selectors.go:1)) supports tag, ID, class, descendant combinators (`div p`), and various attribute selectors (`[attr=value]`, `[attr^=val]`, etc.).
- 🧪 **Built for testing** - Assertions, mocking, and debugging tools.
- 📦 **Pure Go** - No CGO, no external dependencies.
- 🎭 **Framework agnostic** - Test HTMX, Alpine.js, React, Vue, or vanilla JS.
  - ⚠️ Only HTMX has been verified so far, and only the basics.

## 📦 Installation

```bash
go get github.com/dpemmons/DOMulator
```

## 🚀 Quick Start

```go
package main

import (
    "net/http"
    "net/http/httptest"
    "testing"

    domulator "github.com/dpemmons/DOMulator"
)

func TestVanillaJSApp(t *testing.T) {
    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        switch r.URL.Path {
        case "/":
            w.Write([]byte(`
                <input id="name" type="text">
                <button id="greet">Greet</button>
                <div id="message"></div>
                <script src="/app.js"></script>
            `))
        case "/app.js":
            w.Header().Set("Content-Type", "application/javascript")
            w.Write([]byte(`
                document.getElementById('greet').addEventListener('click', () => {
                    const name = document.getElementById('name').value;
                    document.getElementById('message').textContent = 'Hello, ' + name + '!';
                });
            `))
        }
    }))
    defer server.Close()

    test := domulator.NewTest(t)
    defer test.Close()
    test.WithServer(server).Navigate("/")

    test.Type("#name", "World")
    test.Click("#greet")
    test.AssertElement("#message").HasText("Hello, World!")
}
```

> See the complete working example: [`examples/vanilla_js_app_test.go`](examples/vanilla_js_app_test.go)

## 📖 Documentation

### Core Concepts

DOMulator provides three main components:

- **DOM** - A DOM implementation that mirrors browser behavior, with an internal CSS selector engine and event system.
- **JavaScript Runtime** - ECMAScript 5.1+ execution via Goja.
- **Test Utilities** - High-level helpers for interaction and assertions.

### Basic Usage

#### Creating a Test Environment

```go
// Create a browser-like test instance
test := domulator.NewTest(t)

// Navigate to a URL (automatic resource loading)
test.Navigate("http://localhost:3000")

// Or use with httptest server for mocking
server := httptest.NewServer(myHandler)
defer server.Close()

test.WithServer(server).Navigate("/")
```

#### Working with httptest

```go
func TestMyApp(t *testing.T) {
    // Create your test server with whatever behavior you want
    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        switch r.URL.Path {
        case "/":
            w.Write([]byte(`<h1>Home</h1><script src="/app.js"></script>`))
        case "/app.js":
            w.Header().Set("Content-Type", "application/javascript")
            w.Write([]byte(`console.log("App loaded");`))
        case "/api/data":
            w.Header().Set("Content-Type", "application/json")
            w.Write([]byte(`{"status": "ok"}`))
        }
    }))
    defer server.Close()
    
    // Use your real application handler instead
    // server := httptest.NewServer(myapp.Handler())
    
    test := domulator.NewTest(t)
    test.WithServer(server)
    
    // Navigate automatically loads HTML, scripts, CSS, etc.
    test.Navigate("/")
    
    // Relative URLs work automatically
    // test.Navigate("/some/other/path") // Example
}
```

#### Executing JavaScript

```go
// Execute custom JavaScript
test.ExecuteScript(`console.log("Hello from JS!")`)

// JavaScript loaded via <script> tags executes automatically
test.Navigate("/") // Loads page with <script> tags
```


#### Simulating User Interactions

DOMulator provides a rich set of methods to simulate user interactions. Many events like mousedown, mouseup, touch events, drag-and-drop, clipboard events, and media events are also available but not listed exhaustively here. Refer to [`testing/interactions.go`](testing/interactions.go:1) for a complete list.

```go
// Mouse events
test.Click("#submit-button")
test.DoubleClick(".item")
test.RightClick("#context-menu-trigger")
test.Hover(".tooltip-trigger") // Simulates mouseenter and mouseover
// test.Unhover(".tooltip-trigger") // Simulates mouseleave and mouseout

// Keyboard events (provide key names as strings)
test.Type("#search", "golang testing") // Sets value and triggers input/change events
test.KeyDown("#editor", "Control")     // Pass key name as string, e.g., "Enter", "Escape", "A", "a"
test.KeyUp("#editor", "Control")
// test.KeyPress("#editor", "a") // Deprecated, but available

// Form interactions
test.Check("#agree-terms")
test.Uncheck("#newsletter")
test.Select("#country", "USA") // Sets value of select element and triggers change
// test.UploadFile(...) // This feature is not yet implemented

// Focus events
test.Focus("#email")
test.Blur("#email")
```

#### Making Assertions

```go
// Text content
test.AssertElement("h1").HasText("Welcome")
test.AssertElement(".description").ContainsText("powerful")
// test.AssertTextMatches(...) // Regex-based text matching is not yet implemented

// Attributes
test.AssertElement("a").HasAttribute("href", "/home")
// To assert attribute presence (e.g., 'disabled'):
// For boolean attributes, check for its name or an empty string depending on convention
// e.g., test.AssertElement("button[disabled]").Exists() or test.AssertElement("button").HasAttribute("disabled", "")
// test.AssertAttributeMatches(...) // Regex-based attribute matching is not yet implemented

// CSS classes
test.AssertElement("#modal").HasClass("visible")
// To assert no class, you would typically check that HasClass fails or count elements without the class.
// A direct AssertNoClass is not available.

// Visibility (Simplified: checks inline style for display:none or visibility:hidden)
test.AssertElement("#success-message").IsVisible()
test.AssertElement("#error-message").IsHidden()

// Element existence
test.AssertElement("#dynamic-content").Exists()
test.AssertElement(".deleted-item").NotExists()
test.AssertElement(".list-item").Count(5)
```


#### Waiting for Conditions (Planned Feature)

> [!NOTE]
> The following WaitFor... utilities are planned for future releases and are not yet implemented.

```go
// // Wait for element (Example of planned API)
// test.WaitForElement("#lazy-loaded", 5*time.Second)

// // Wait for text (Example of planned API)
// test.WaitForText("#status", "Complete", 10*time.Second)

// // Wait for custom condition (Example of planned API)
// test.WaitFor(func() bool {
//     counter := test.QuerySelector("#counter") // QuerySelector is available on Document/Element
//     return counter.TextContent() == "100"
// }, 5*time.Second)

// // Wait for network idle (Example of planned API)
// test.WaitForNetworkIdle()
```

#### Time Manipulation & Advanced Testing

DOMulator provides precise control over JavaScript timing for testing async behavior:

```go
// Control JavaScript timers
test.AdvanceTime(100 * time.Millisecond)  // Move time forward
test.ProcessPendingTimers()               // Execute scheduled timers
test.FlushMicrotasks()                   // Process microtask queue

// ResizeObserver testing
test.TriggerElementResize("#sidebar", domulator.ResizeOptions{
    Width: 300, Height: 500,
    ContentWidth: 280, ContentHeight: 480,  // Optional: inner dimensions
})

// Window resize
test.ResizeWindow(1024, 768)  // Changes window dimensions and fires resize event
test.Resize()                 // Fires window resize event only (no dimension change)

// Custom initial window size
test := domulator.NewTestWithConfig(t, &domulator.TestConfig{
    WindowWidth:  1280,
    WindowHeight: 720,
    ElementWidth: 200,   // Default element dimensions for ResizeObserver
    ElementHeight: 150,
})
```


**Key Principles for Resize Testing:**

- [`TriggerElementResize()`](testing/harness.go:1) only affects the specific element you target for ResizeObserver callbacks.
- No automatic cascading layout - explicitly resize each element you want to test if its dimensions affect others in a layout-dependent way (though DOMulator currently has no layout engine).
- [`Resize()`](testing/harness.go:1) only fires window resize events, not ResizeObserver callbacks directly on elements.
- [`ResizeWindow()`](testing/harness.go:1) changes window dimensions AND fires the window resize event.

## 🔌 Framework Examples

> [!NOTE]
> Complete working examples can be found in the [`examples/`](examples/) directory.

### Form Handling

```go
func TestHTMXForm(t *testing.T) {
    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        switch r.URL.Path {
        case "/":
            w.Header().Set("Content-Type", "text/html")
            w.Write([]byte(`
                <form id="contact-form">
                    <input name="email" type="email" required>
                    <button type="submit">Send</button>
                </form>
                <div id="result"></div>
                <script>
                    // Simple form handler without HTMX
                    document.getElementById('contact-form').addEventListener('submit', async function(e) {
                        e.preventDefault();
                        const email = this.querySelector('input[name="email"]').value;
                        
                        // Simulate API call
                        if (email) {
                            document.getElementById('result').innerHTML = '<div class="success">Message sent!</div>';
                        }
                    });
                </script>
            `))
        }
    }))
    defer server.Close()

    test := domulator.NewTest(t)
    defer test.Close()
    test.WithServer(server).Navigate("/")

    test.Type("input[name=email]", "test@example.com")
    test.Click("button")
    test.AssertElement("#result .success").HasText("Message sent!")
}
```

> See the complete working example: [`examples/htmx_form_test.go`](examples/htmx_form_test.go)

### Infinite Scroll / Dynamic Content

```go
func TestHTMXInfiniteScroll(t *testing.T) {
    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        switch r.URL.Path {
        case "/":
            w.Write([]byte(`
                <div id="items">
                    <div class="item">Item 1</div>
                    <div class="item">Item 2</div>
                </div>
                <button id="load-more">Load More</button>
                <script>
                    let itemCount = 2;
                    document.getElementById('load-more').addEventListener('click', function() {
                        const container = document.getElementById('items');
                        
                        // Add two new items
                        for (let i = 0; i < 2; i++) {
                            itemCount++;
                            const item = document.createElement('div');
                            item.className = 'item';
                            item.textContent = 'Item ' + itemCount;
                            container.appendChild(item);
                        }
                    });
                </script>
            `))
        }
    }))
    defer server.Close()

    test := domulator.NewTest(t)
    defer test.Close()
    test.WithServer(server).Navigate("/")
    test.AssertElement(".item").Count(2) // Initial load

    // Trigger loading more items
    test.Click("#load-more")
    test.FlushMicrotasks()

    test.AssertElement(".item").Count(4)
}
```

> See the complete working example: [`examples/htmx_infinite_scroll_test.go`](examples/htmx_infinite_scroll_test.go)

### Component-like Behavior

```go
func TestAlpineJSComponent(t *testing.T) {
    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte(`
            <div id="app">
                <button id="btn">Test</button>
                <span id="result"></span>
            </div>
            <script>
                // Simple test - manually create Alpine functionality
                document.getElementById('btn').addEventListener('click', () => {
                    const span = document.getElementById('result');
                    const current = parseInt(span.textContent) || 0;
                    span.textContent = current + 1;
                });
                
                // Initialize the counter
                document.getElementById('result').textContent = '0';
            </script>
        `))
    }))
    defer server.Close()

    test := domulator.NewTest(t)
    defer test.Close()
    test.WithServer(server).Navigate("/")

    test.AssertElement("#result").HasText("0")
    test.Click("#btn")
    test.AssertElement("#result").HasText("1")
    test.Click("#btn")
    test.AssertElement("#result").HasText("2")
}
```

> See the complete working example: [`examples/alpinejs_component_test.go`](examples/alpinejs_component_test.go)

### Vanilla JavaScript

```go
func TestVanillaJSApp(t *testing.T) {
    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        switch r.URL.Path {
        case "/":
            w.Write([]byte(`
                <input id="name" type="text">
                <button id="greet">Greet</button>
                <div id="message"></div>
                <script src="/app.js"></script>
            `))
        case "/app.js":
            w.Header().Set("Content-Type", "application/javascript")
            w.Write([]byte(`
                document.getElementById('greet').addEventListener('click', () => {
                    const name = document.getElementById('name').value;
                    document.getElementById('message').textContent = 'Hello, ' + name + '!';
                });
            `))
        }
    }))
    defer server.Close()

    test := domulator.NewTest(t)
    defer test.Close()
    test.WithServer(server).Navigate("/")

    test.Type("#name", "World")
    test.Click("#greet")
    test.AssertElement("#message").HasText("Hello, World!")
}
```

> See the complete working example: [`examples/vanilla_js_app_test.go`](examples/vanilla_js_app_test.go)

### Additional Examples

More examples are available in the [`examples/`](examples/) directory:

- [`examples/async_js_test.go`](examples/async_js_test.go) - Testing asynchronous JavaScript operations
- [`examples/debugging_test.go`](examples/debugging_test.go) - Console output and debugging techniques  
- [`examples/httptest_server_test.go`](examples/httptest_server_test.go) - HTTP API testing patterns


## 🔧 Advanced Usage

### Custom Plugins

> [!NOTE]
> The InstallPlugin method and domulator.Environment interface are not directly visible in the current public API of domulator.Test or testing.TestHarness. This example may reflect a planned or internal mechanism.

```go
// // Example of a potential plugin structure (API may vary)
// type GeolocationPlugin struct{}

// func (g *GeolocationPlugin) Install(runtime *js.Runtime, window *goja.Object) { // Assuming direct access
//     runtime.Set("navigator.geolocation", map[string]interface{}{
//         "getCurrentPosition": func(call goja.FunctionCall) goja.Value {
//             success := call.Argument(0).Export().(goja.Callable)
//             // error := call.Argument(1).Export().(goja.Callable) // Handle error callback
            
//             position := map[string]interface{}{
//                 "coords": map[string]float64{
//                     "latitude":  37.7749,
//                     "longitude": -122.4194,
//                 },
//                 "timestamp": time.Now().Unix(),
//             }
//             // Call success callback from JS runtime
//             // This requires careful handling of Goja values and execution context
//             // success(nil, runtime.ToValue(position))
//             return goja.Undefined()
//         },
//     })
// }

// // Usage (conceptual)
// // test := domulator.NewTest(t)
// // runtime := test.Runtime() // Get the underlying JS runtime
// // windowObj := runtime.GlobalObject().Get("window").(*goja.Object) // Get window object
// // (&GeolocationPlugin{}).Install(runtime, windowObj)
```

### Debugging

```go
// Enable console output from JavaScript and capture logs
test := domulator.NewTest(t)
test.SetDebugMode(true) // Enables verbose JS console output to its default (e.g., terminal)
consoleCapture := test.CaptureConsole(t) // Captures console messages for programmatic access

// ... perform actions that might log to console ...
test.ExecuteScript(`console.log("Hello from test!", {detail: 123}); console.error("An error occurred");`)

// Retrieve and print captured messages
messages := consoleCapture.GetAllMessages() // Returns []testing.ConsoleMessage
for _, msg := range messages {
    // msg.Args is []interface{}, format them as needed
    argStrings := []string{}
    for _, arg := range msg.Args {
        argStrings = append(argStrings, fmt.Sprintf("%v", arg))
    }
    fmt.Printf("[%s] %s\n", msg.Level, strings.Join(argStrings, " "))
}
// You can also use consoleCapture.GetLogs(), .GetErrors(), etc.
// And assertion helpers like consoleCapture.AssertLogContains("expected text")

// Take DOM snapshots (This feature is not yet implemented)
// snapshot := test.TakeSnapshot()
// fmt.Println(snapshot.HTML)

// Debug specific elements
elemNode := test.QuerySelector("#mysterious-div") // Returns dom.Node, cast to *dom.Element
if elem, ok := elemNode.(*dom.Element); ok {
    fmt.Printf("Element: %s\n", elem.TagName())
    fmt.Printf("ID: %s\n", elem.GetAttribute("id"))
    fmt.Printf("Classes: %s\n", elem.GetAttribute("class")) // Or iterate elem.ClassList().Values()
    
    // To get all attributes:
    // attrs := elem.AttributesNamedNodeMap().ToSlice() // Access internal NamedNodeMap
    // for _, attr := range attrs {
    //     fmt.Printf("Attribute: %s = %s\n", attr.Name(), attr.Value())
    // }
}
// fmt.Printf("Computed Style: %v\n", elem.ComputedStyle()) // ComputedStyle is not yet implemented
```

## 📄 License

MIT License - see LICENSE for details.

## 🙏 Acknowledgments

DOMulator is built on these excellent libraries:

- [Goja](https://github.com/dop251/goja) - JavaScript runtime
- [golang.org/x/net/html](https://golang.org/x/net/html) - HTML parsing

## Known Limitations / Future Work

- **ES6 modules**: are not supported.
- **Advanced CSS Selectors**: Support for child (`>`), sibling (`+`, `~`) combinators, and most pseudo-classes/elements is planned.
- **Layout Engine**: Currently, there is no layout engine. Element dimensions and positions are not computed.
- **Testing Utilities**:
  - [`WaitFor...`](testing/harness.go:1) conditions (e.g., WaitForElement, WaitForText) are planned.
  - [`TakeSnapshot()`](testing/harness.go:1) for DOM snapshots is planned.
  - [`UploadFile()`](testing/harness.go:1) for form interactions is planned.
- **DOM API Completeness**:
  - [`ComputedStyle()`](internal/dom/element.go:1) for elements is not yet implemented.
  - [`innerHTML`](internal/dom/element.go:1) parsing and TextContent extraction have simplified implementations and are marked for improvement.
  - No layout engine (resize events can be triggered manually).
  - Advanced CSS selectors like child (`>`), sibling (`+`, `~`) combinators, and most pseudo-classes/elements are not yet supported.
  - Some DOM features like `innerHTML`/`TextContent` parsing and `ComputedStyle` are placeholders or not yet implemented.
- **Framework Verification**: While designed to be framework-agnostic, thorough verification beyond basic HTMX is ongoing.
