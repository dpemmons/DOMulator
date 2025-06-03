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
    "testing"
    "github.com/example/domulator"
)

func TestHTMXForm(t *testing.T) {
    // Create a test environment
    test := domulator.NewTest(t)

    // Mock your server endpoint
    test.MockPOST("/api/contact", func(req domulator.Request) domulator.Response {
        return domulator.Response{
            Status: 200,
            Body: `<div class="success">Message sent!</div>`,
        }
    })

    // Load your HTML
    test.LoadHTML(`
        <form hx-post="/api/contact" hx-target="#result">
            <input name="email" type="email" required>
            <button type="submit">Send</button>
        </form>
        <div id="result"></div>
    `)

    // Load HTMX
    test.LoadScript("https://unpkg.com/htmx.org")

    // Interact with the page
    test.Type("input[name=email]", "test@example.com")
    test.Click("button")

    // Assert the result
    test.AssertText("#result .success", "Message sent!")
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
// Basic setup
dom := domulator.New()

// With test harness (recommended)
test := domulator.NewTest(t)

// With custom options
dom := domulator.New(
    domulator.WithConsoleLog(true),
    domulator.WithTimeout(30 * time.Second),
    domulator.WithBaseURL("https://example.com"),
)
```

#### Loading Content

```go
// HTML string
test.LoadHTML(`<div>Hello World</div>`)

// HTML file
test.LoadHTMLFile("templates/index.html")

// JavaScript
test.ExecuteScript(`console.log("Hello from JS!")`)
test.LoadScriptFile("app.js")
test.LoadScript("https://cdn.jsdelivr.net/npm/alpinejs")
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

#### Network Mocking

```go
// Simple response
test.MockGET("/api/user", domulator.Response{
    Status: 200,
    Body: `{"name": "John Doe"}`,
})

// Dynamic handler
test.MockPOST("/api/items", func(req domulator.Request) domulator.Response {
    var item Item
    json.Unmarshal(req.Body, &item)

    return domulator.Response{
        Status: 201,
        Headers: map[string]string{
            "Location": fmt.Sprintf("/api/items/%d", item.ID),
        },
        Body: req.Body, // Echo back
    }
})

// Pattern matching
test.MockPattern("GET /api/users/*", func(req domulator.Request) domulator.Response {
    userID := path.Base(req.URL)
    return domulator.Response{
        Status: 200,
        Body: fmt.Sprintf(`{"id": "%s"}`, userID),
    }
})

// Verify requests
requests := test.GetRequests("/api/submit")
assert.Len(t, requests, 1)
assert.Equal(t, "application/json", requests[0].Header("Content-Type"))
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

## 🔌 Framework Examples

### HTMX

```go
func TestHTMXInfiniteScroll(t *testing.T) {
    test := domulator.NewTest(t)

    page := 1
    test.MockGET("/api/items", func(req domulator.Request) domulator.Response {
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

        return domulator.Response{Status: 200, Body: html}
    })

    test.LoadHTML(`
        <div id="items" hx-get="/api/items" hx-trigger="load"></div>
    `)
    test.LoadScript("https://unpkg.com/htmx.org")

    // Trigger infinite scroll
    test.ScrollIntoView("[hx-trigger='revealed']")
    test.WaitForCount(".item", 4)

    test.AssertCount(".item", 4)
}
```

### Alpine.js

```go
func TestAlpineJSComponent(t *testing.T) {
    test := domulator.NewTest(t)

    test.LoadHTML(`
        <div x-data="{ count: 0 }">
            <button @click="count++">Increment</button>
            <span x-text="count"></span>
        </div>
    `)
    test.LoadScript("https://unpkg.com/alpinejs")

    test.AssertText("span", "0")

    test.Click("button")
    test.AssertText("span", "1")

    test.Click("button")
    test.AssertText("span", "2")
}
```

### Vanilla JavaScript

```go
func TestVanillaJSApp(t *testing.T) {
    test := domulator.NewTest(t)

    test.LoadHTML(`
        <input id="name" type="text">
        <button id="greet">Greet</button>
        <div id="message"></div>
    `)

    test.ExecuteScript(`
        document.getElementById('greet').addEventListener('click', () => {
            const name = document.getElementById('name').value;
            document.getElementById('message').textContent = 'Hello, ' + name + '!';
        });
    `)

    test.Type("#name", "World")
    test.Click("#greet")
    test.AssertText("#message", "Hello, World!")
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
