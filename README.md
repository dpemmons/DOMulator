# DOMulator

Fast, lightweight DOM emulation and JavaScript runtime for testing web applications in Go. Run JavaScript frameworks like HTMX, Alpine.js, or vanilla JavaScript without a headless browser.

## ✨ Features

- 🚀 **High Performance** - Significantly faster than headless browser testing
- 🎯 **Real JavaScript** - Powered by Goja (ECMAScript 5.1+)
- 🔧 **Full DOM API** - Run real frameworks without modifications
- 🧪 **Built for testing** - Assertions, mocking, and debugging tools
- 📦 **Pure Go** - No CGO, no external dependencies
- 🎭 **Framework agnostic** - Test HTMX, Alpine.js, React, Vue, or vanilla JS

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
    "github.com/dpemmons/DOMulator"
)

func TestHTMXForm(t *testing.T) {
    // Create a test server with your application
    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        switch r.URL.Path {
        case "/":
            w.Header().Set("Content-Type", "text/html")
            w.Write([]byte(`
                <form hx-post="/api/contact" hx-target="#result">
                    <input name="email" type="email" required>
                    <button type="submit">Send</button>
                </form>
                <div id="result"></div>
                <script src="https://unpkg.com/htmx.org"></script>
            `))
        case "/api/contact":
            w.Write([]byte(`<div class="success">Message sent!</div>`))
        }
    }))
    defer server.Close()

    // Create a DOMulator test instance
    test := domulator.NewTest(t)
    test.WithServer(server)
    
    // Navigate to the page - automatically loads HTML, scripts, etc.
    test.Navigate("/")
    
    // Interact with the page
    test.Type("input[name=email]", "test@example.com")
    test.Click("button")
    
    // Assert the result
    test.AssertElement("#result .success").HasText("Message sent!")
}
```

## 📖 Documentation

### Core Concepts

DOMulator provides three main components:

- **DOM** - A complete DOM implementation that mirrors browser behavior
- **JavaScript Runtime** - ECMAScript 5.1+ execution via Goja
- **Test Utilities** - High-level helpers for interaction and assertions

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
    test.Navigate("/api/data")
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

```go
// Mouse events
test.Click("#submit-button")
test.DoubleClick(".item")
test.RightClick("#context-menu-trigger")
test.Hover(".tooltip-trigger")

// Keyboard events
test.Type("#search", "golang testing")
test.PressKey("#search", domulator.KeyEnter)
test.KeyDown("#editor", domulator.KeyCtrl)
test.KeyUp("#editor", domulator.KeyCtrl)

// Form interactions
test.Check("#agree-terms")
test.Uncheck("#newsletter")
test.SelectOption("#country", "USA")
test.UploadFile("#avatar", "avatar.jpg", imageBytes)

// Focus events
test.Focus("#email")
test.Blur("#email")
```

#### Making Assertions

```go
// Text content
test.AssertText("h1", "Welcome")
test.AssertTextContains(".description", "powerful")
test.AssertTextMatches(".code", `[A-Z]{3}-\d{3}`)

// Attributes
test.AssertAttribute("a", "href", "/home")
test.AssertHasAttribute("button", "disabled")
test.AssertAttributeMatches("input", "pattern", `\d+`)

// CSS classes
test.AssertHasClass("#modal", "visible")
test.AssertNoClass("#modal", "hidden")

// Visibility
test.AssertVisible("#success-message")
test.AssertHidden("#error-message")

// Element existence
test.AssertExists("#dynamic-content")
test.AssertNotExists(".deleted-item")
test.AssertCount(".list-item", 5)
```


#### Waiting for Conditions

```go
// Wait for element
test.WaitForElement("#lazy-loaded", 5*time.Second)

// Wait for text
test.WaitForText("#status", "Complete", 10*time.Second)

// Wait for custom condition
test.WaitFor(func() bool {
    counter := test.QuerySelector("#counter")
    return counter.TextContent() == "100"
}, 5*time.Second)

// Wait for network idle
test.WaitForNetworkIdle()
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

**Key Principles:**
- `TriggerElementResize()` only affects the specific element you target
- No automatic cascading - explicitly resize each element you want to test
- `Resize()` only fires window events, not ResizeObserver callbacks
- `ResizeWindow()` changes window dimensions AND fires resize event

## 🔌 Framework Examples

### HTMX

```go
func TestHTMXInfiniteScroll(t *testing.T) {
    page := 1
    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        switch r.URL.Path {
        case "/":
            w.Write([]byte(`
                <div id="items" hx-get="/api/items" hx-trigger="load"></div>
                <script src="https://unpkg.com/htmx.org"></script>
            `))
        case "/api/items":
            html := fmt.Sprintf(`
                <div class="item">Item %d</div>
                <div class="item">Item %d</div>
                <div hx-get="/api/items?page=%d" 
                     hx-trigger="revealed" 
                     hx-swap="outerHTML">
                    Load More
                </div>
            `, page*2-1, page*2, page+1)
            page++
            w.Write([]byte(html))
        }
    }))
    defer server.Close()

    test := domulator.NewTest(t)
    test.WithServer(server).Navigate("/")

    // Trigger infinite scroll
    test.ScrollIntoView("[hx-trigger='revealed']")
    test.WaitForCount(".item", 4)

    test.AssertElement(".item").Count(4)
}
```

### Alpine.js

```go
func TestAlpineJSComponent(t *testing.T) {
    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte(`
            <div x-data="{ count: 0 }">
                <button @click="count++">Increment</button>
                <span x-text="count"></span>
            </div>
            <script src="https://unpkg.com/alpinejs"></script>
        `))
    }))
    defer server.Close()

    test := domulator.NewTest(t)
    test.WithServer(server).Navigate("/")

    test.AssertElement("span").HasText("0")

    test.Click("button")
    test.AssertElement("span").HasText("1")

    test.Click("button")
    test.AssertElement("span").HasText("2")
}
```

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
    test.WithServer(server).Navigate("/")

    test.Type("#name", "World")
    test.Click("#greet")
    test.AssertElement("#message").HasText("Hello, World!")
}
```

## 🔧 Advanced Usage

### Custom Plugins

```go
// Create a custom browser API
type GeolocationPlugin struct{}

func (g *GeolocationPlugin) Install(env domulator.Environment) {
    env.Window.Set("navigator.geolocation", map[string]interface{}{
        "getCurrentPosition": func(success, error goja.Callable) {
            position := map[string]interface{}{
                "coords": map[string]float64{
                    "latitude":  37.7749,
                    "longitude": -122.4194,
                },
                "timestamp": time.Now().Unix(),
            }
            success(nil, env.Runtime.ToValue(position))
        },
    })
}

// Use the plugin
test := domulator.NewTest(t)
test.InstallPlugin(&GeolocationPlugin{})
```

### Debugging

```go
// Enable console output
test := domulator.NewTest(t, domulator.WithConsoleLog(true))

// Capture console logs
logs := test.GetConsoleLogs()
for _, log := range logs {
    fmt.Printf("[%s] %s\n", log.Level, log.Message)
}

// Take DOM snapshots
snapshot := test.TakeSnapshot()
fmt.Println(snapshot.HTML)

// Debug specific elements
elem := test.QuerySelector("#mysterious-div")
fmt.Printf("Classes: %v\n", elem.ClassList())
fmt.Printf("Attributes: %v\n", elem.Attributes())
fmt.Printf("Computed Style: %v\n", elem.ComputedStyle())
```

## 🤝 Contributing

Contributions are welcome! Please read our Contributing Guide for details.

```bash
# Clone the repo
git clone https://github.com/dpemmons/DOMulator.git

# Install dependencies
cd domulator
go mod download

# Run tests
go test ./...

# Run benchmarks
go test -bench=. ./...
```

## 📄 License

MIT License - see LICENSE for details.

## 🙏 Acknowledgments

DOMulator is built on these excellent libraries:

- [Goja](https://github.com/dop251/goja) - JavaScript runtime
- [goquery](https://github.com/PuerkitoBio/goquery) - CSS selectors
- [golang.org/x/net/html](https://golang.org/x/net/html) - HTML parsing
