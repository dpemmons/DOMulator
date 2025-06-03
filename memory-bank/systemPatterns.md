# System Patterns: DOMulator Architecture

## System Architecture
DOMulator is designed with a layered architecture to separate concerns and facilitate modular development.

```mermaid
graph TD
    A[DOMulator API] --> B(Test Harness Layer)
    B --> C(JavaScript Runtime - Goja)
    C --> D(DOM Core)
    D --> E(Browser APIs)
    E --> F(Foundation Layer)

    subgraph Test Harness Layer
        B1[Assertions]
        B2[Interactions]
        B3[Network Mocks]
        B4[Snapshots]
    end

    subgraph JavaScript Runtime (Goja)
        C1[VM]
        C2[Bindings]
        C3[Polyfills]
        C4[Module Loader]
    end

    subgraph DOM Core
        D1[Document]
        D2[Elements]
        D3[Events]
        D4[Mutations]
    end

    subgraph Browser APIs
        E1[Fetch]
        E2[XHR]
        E3[Timers]
        E4[Storage]
        E5[Console]
    end

    subgraph Foundation Layer
        F1[HTML Parser]
        F2[CSS Selectors]
        F3[Event Loop]
    end
```

## Key Technical Decisions
- **Pure Go Implementation**: Avoids CGO for cross-platform compatibility and ease of deployment.
- **Goja for JavaScript**: Chosen for its pure Go implementation, performance, and excellent Go/JS interop.
- **Custom DOM**: A custom-built DOM implementation for precise control, lazy evaluation, and built-in mutation tracking.
- **Event Loop**: Implements the HTML specification's event loop model for accurate asynchronous behavior.
- **Network Interception**: Provides comprehensive mocking capabilities for Fetch and XMLHttpRequest.

## Design Patterns in Use
- **Layered Architecture**: Clear separation of concerns, promoting modularity and maintainability.
- **Interface-based Design**: Used extensively in the DOM Core (e.g., `Node` interface) to allow for specialized element types and extensibility.
- **Lazy Evaluation**: Properties like `innerHTML` and `outerHTML` are computed on demand to optimize memory usage and performance.
- **Interceptor Pattern**: For network mocking, allowing dynamic handling of requests.
- **Plugin Architecture**: Enables extending DOMulator's capabilities with custom browser APIs.

## Component Relationships
- **DOM Core & JavaScript Runtime**: Tightly coupled through `DOMBindings` which expose Go DOM objects to the Goja VM.
- **Test Harness & Core**: The `TestHarness` orchestrates interactions with the `DOMulator` instance, leveraging its core functionalities for assertions and simulations.
- **Foundation & Upper Layers**: The HTML Parser, CSS Selectors, and Event Loop provide fundamental services consumed by the DOM Core and JavaScript Runtime.

## Critical Implementation Paths
1. **HTML Parsing to DOM Construction**: Efficiently converting raw HTML into a functional DOM tree.
2. **JavaScript Binding**: Seamlessly exposing Go DOM objects and browser APIs to the Goja JavaScript runtime.
3. **Event Dispatching**: Accurate propagation of events through the DOM tree (capture, target, bubble phases).
4. **Network Mocking**: Intercepting and responding to HTTP requests made from JavaScript within the emulated environment.
5. **Event Loop Management**: Correctly scheduling and executing tasks, microtasks, timers, and animation frames.

## 🎯 **NEW ARCHITECTURE DIRECTION**: Browser APIs for Modern Framework Support

### **Phase 1: HTMX Critical APIs Architecture** 🚀

With the DOM foundation complete, the next architectural focus is implementing browser APIs to support modern web frameworks, starting with HTMX.

#### **HTTP/Fetch API Implementation Pattern**
```go
// New package: internal/browser/fetch
type FetchAPI struct {
    client *http.Client
    interceptors []RequestInterceptor
}

type Response struct {
    status int
    headers map[string]string
    body io.Reader
}

// JavaScript binding integration
func (b *DOMBindings) AddFetchAPI() {
    b.vm.Set("fetch", b.createFetchFunction())
}
```

#### **FormData API Implementation Pattern**
```go
// New package: internal/browser/forms
type FormData struct {
    fields map[string][]FormField
}

type FormField struct {
    value string
    filename string  // for file uploads
    contentType string
}

// JavaScript binding integration
func (b *DOMBindings) AddFormDataAPI() {
    b.vm.Set("FormData", b.createFormDataConstructor())
}
```

#### **CustomEvent API Implementation Pattern**
```go
// Extension to internal/dom/events
type CustomEvent struct {
    *eventImpl
    detail interface{}
}

// JavaScript binding integration
func (b *DOMBindings) AddCustomEventAPI() {
    b.vm.Set("CustomEvent", b.createCustomEventConstructor())
}
```

### **Architectural Principles for Browser APIs**

#### **1. Modular Design**
- Each browser API in separate package under `internal/browser/`
- Clear separation between Go implementation and JavaScript bindings
- Pluggable architecture for optional API inclusion

#### **2. JavaScript Compatibility**
- Faithful reproduction of browser API surfaces
- Proper error handling and edge case behavior
- Compatible Promise/async patterns using Goja

#### **3. Testing Integration**
- Browser APIs integrate with existing testing framework
- Network mocking extends to new fetch API
- Event simulation works with CustomEvent

#### **4. Performance Considerations**
- Lazy initialization of browser APIs
- Efficient request/response handling
- Memory-conscious FormData implementation

### **Implementation Phases**

#### **Phase 1 Architecture: HTMX Support**
```
internal/browser/
├── fetch/           # HTTP/Fetch API implementation
│   ├── fetch.go     # Core fetch functionality
│   ├── response.go  # Response object implementation
│   └── request.go   # Request object implementation
├── forms/           # FormData API implementation
│   ├── formdata.go  # FormData constructor and methods
│   └── multipart.go # Multipart encoding support
└── events/          # CustomEvent extensions
    └── custom.go    # CustomEvent implementation
```

#### **Phase 2 Architecture: Extended Browser APIs**
```
internal/browser/
├── storage/         # localStorage, sessionStorage
├── url/            # URL, URLSearchParams APIs
├── history/        # History API (pushState, replaceState)
└── observers/      # MutationObserver, IntersectionObserver
```

### **Integration with Existing Architecture**

#### **JavaScript Bindings Extension**
```go
// Enhanced DOMBindings with browser API support
type DOMBindings struct {
    vm *goja.Runtime
    document *dom.Document
    nodeCache map[dom.Node]*goja.Object
    
    // New: Browser API managers
    fetchAPI *fetch.FetchAPI
    formAPI *forms.FormAPI
    eventAPI *events.CustomEventAPI
}
```

#### **Testing Framework Integration**
- **TestHarness**: Extended with browser API configuration
- **Network Mocking**: Integrated with fetch API
- **Assertions**: New assertions for HTTP responses and form data

This architectural approach ensures that DOMulator can support modern web frameworks while maintaining its core principles of speed, reliability, and ease of use.

## Future Extensibility Patterns

### **Plugin Architecture for Browser APIs**
- **Interface-based Design**: Common interfaces for browser API implementations
- **Registration System**: Dynamic registration of custom browser APIs
- **Configuration**: Runtime configuration of which APIs to enable
- **Testing Support**: Framework for testing custom browser API implementations

### **Performance Optimization Strategies**
- **Object Pooling**: Reuse DOM node objects and browser API objects to reduce allocation
- **Lazy Evaluation**: Defer expensive operations until needed
- **Caching Strategies**: Cache parsed selectors, compiled scripts, and HTTP responses
- **Memory Management**: Proper cleanup of JavaScript object references and browser API resources
