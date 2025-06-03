# DOMulator Architecture

## Overview

DOMulator is a pure Go library that provides a complete DOM emulation environment with JavaScript execution capabilities. It enables testing of JavaScript-driven web applications without the overhead of headless browsers.

## Design Goals

- **Performance** - Minimal overhead, lazy evaluation, efficient memory usage
- **Compatibility** - Sufficient browser API compliance to run real JavaScript frameworks
- **Testability** - Built-in testing utilities and introspection capabilities
- **Extensibility** - Plugin system for custom browser APIs
- **Simplicity** - Clean API that feels natural to Go developers

## Core Architecture

```
┌─────────────────────────────────────────────────────────────┐
│                        DOMulator API                         │
├─────────────────────────────────────────────────────────────┤
│                      Test Harness Layer                      │
│  Assertions │ Interactions │ Network Mocks │ Snapshots      │
├─────────────────────────────────────────────────────────────┤
│                    JavaScript Runtime (Goja)                 │
│  VM │ Bindings │ Polyfills │ Module Loader                  │
├─────────────────────────────────────────────────────────────┤
│                        DOM Core                              │
│  Document │ Elements │ Events │ Mutations                   │
├─────────────────────────────────────────────────────────────┤
│                     Browser APIs                             │
│  Fetch │ XHR │ Timers │ Storage │ Console                  │
├─────────────────────────────────────────────────────────────┤
│                    Foundation Layer                          │
│  HTML Parser │ CSS Selectors │ Event Loop                   │
└─────────────────────────────────────────────────────────────┘
```

## Component Design

### 1. DOM Core (internal/dom)

The DOM implementation is custom-built for precise control over behavior and optimal integration with the JavaScript runtime.

```go
// internal/dom/node.go
type Node interface {
    NodeType() NodeType
    NodeName() string
    NodeValue() string

    ParentNode() Node
    ChildNodes() NodeList
    FirstChild() Node
    LastChild() Node
    PreviousSibling() Node
    NextSibling() Node

    AppendChild(child Node) Node
    RemoveChild(child Node) Node
    InsertBefore(newChild, refChild Node) Node
    ReplaceChild(newChild, oldChild Node) Node

    CloneNode(deep bool) Node

    // Internal methods
    setParent(parent Node)
    toJS(vm *goja.Runtime) goja.Value
}

// internal/dom/element.go
type Element struct {
    nodeImpl

    tagName        string
    attributes     *AttributeMap
    classList      *ClassList
    dataset        *Dataset
    style          *CSSStyleDeclaration

    // Lazy-loaded properties
    innerHTML      *string
    outerHTML      *string
    textContent    *string

    // Event handling
    eventListeners map[EventPhase]map[string][]*EventListener

    // Shadow DOM support
    shadowRoot     *ShadowRoot

    // Form elements
    value          string
    checked        bool
    selected       bool

    // Mutation tracking
    mutationRecord *MutationRecord
}
```

**Key Design Decisions:**

- **Interface-based** - Allows for specialized element types
- **Lazy evaluation** - Properties computed on demand
- **Mutation tracking** - Built-in change detection
- **Memory efficiency** - String pointers for large text content

### 2. JavaScript Runtime (internal/js)

Built on Goja for ECMAScript 5.1+ support.

```go
// internal/js/runtime.go
type Runtime struct {
    vm          *goja.Runtime
    eventLoop   *EventLoop
    modules     *ModuleLoader
    console     *Console

    // Global objects
    window      *Window
    document    *Document
    navigator   *Navigator
    location    *Location

    // Polyfills
    polyfills   map[string]string
}

// internal/js/bindings.go
type DOMBindings struct {
    runtime *Runtime

    // Cached constructors
    elementConstructor     goja.Callable
    eventConstructor       goja.Callable
    nodeListConstructor    goja.Callable
}

func (b *DOMBindings) BindElement(elem *dom.Element) goja.Value {
    obj := b.runtime.vm.NewObject()

    // Define properties with proper descriptors
    obj.DefineAccessorProperty("innerHTML",
        b.runtime.vm.ToValue(func() goja.Value {
            return b.runtime.vm.ToValue(elem.InnerHTML())
        }),
        b.runtime.vm.ToValue(func(val goja.Value) {
            elem.SetInnerHTML(val.String())
        }),
        goja.FLAG_FALSE, goja.FLAG_TRUE,
    )

    // Bind methods
    obj.Set("addEventListener", elem.AddEventListener)
    obj.Set("removeEventListener", elem.RemoveEventListener)
    obj.Set("dispatchEvent", elem.DispatchEvent)

    return obj
}
```

**Why Goja:**

- Pure Go implementation (no CGO)
- 6-7x faster than Otto
- Excellent Go/JS interop
- Active development and community

### 3. Event System (internal/events)

Implements full DOM Level 3 Events specification.

```go
// internal/events/event.go
type Event struct {
    typ            string
    target         EventTarget
    currentTarget  EventTarget
    eventPhase     EventPhase
    timeStamp      int64

    bubbles        bool
    cancelable     bool
    defaultPrevented bool
    propagationStopped bool
    immediatePropagationStopped bool

    // Additional properties for specific event types
    detail         interface{} // CustomEvent
    data           string      // InputEvent
    key            string      // KeyboardEvent
    button         int         // MouseEvent
    touches        []Touch     // TouchEvent
}

// internal/events/dispatcher.go
type EventDispatcher struct {
    captureListeners map[string][]*EventListener
    bubbleListeners  map[string][]*EventListener
}

func (d *EventDispatcher) Dispatch(event *Event) bool {
    // Capture phase
    path := buildEventPath(event.target)
    for i := len(path) - 1; i >= 0; i-- {
        if event.propagationStopped {
            break
        }
        executeListeners(path[i], event, CAPTURING_PHASE)
    }

    // Target phase
    if !event.propagationStopped {
        executeListeners(event.target, event, AT_TARGET)
    }

    // Bubble phase
    if event.bubbles && !event.propagationStopped {
        for i := 1; i < len(path); i++ {
            if event.propagationStopped {
                break
            }
            executeListeners(path[i], event, BUBBLING_PHASE)
        }
    }

    return !event.defaultPrevented
}
```

### 4. Network Layer (internal/net)

Provides Fetch and XMLHttpRequest implementations with full mocking capabilities.

```go
// internal/net/fetch.go
type FetchAPI struct {
    interceptor Interceptor
    cache       *ResponseCache
}

type Request struct {
    method      string
    url         *url.URL
    headers     http.Header
    body        io.ReadCloser

    // Fetch-specific
    mode        RequestMode
    credentials RequestCredentials
    cache       RequestCache
    redirect    RequestRedirect
    referrer    string
    integrity   string

    // Internal
    context     context.Context
    startTime   time.Time
}

// internal/net/interceptor.go
type Interceptor interface {
    ShouldIntercept(req *Request) bool
    Intercept(req *Request) (*Response, error)
}

type NetworkMock struct {
    mu       sync.RWMutex
    handlers []MockHandler
    fallback Interceptor
}

type MockHandler struct {
    matcher  RequestMatcher
    handler  ResponseGenerator
    priority int
}

type RequestMatcher func(*Request) bool
type ResponseGenerator func(*Request) *Response
```

### 5. Event Loop (internal/loop)

Implements HTML specification's event loop model.

```go
// internal/loop/eventloop.go
type EventLoop struct {
    // Task queues
    tasks       *TaskQueue
    microtasks  *MicrotaskQueue

    // Rendering
    animationCallbacks []AnimationFrameCallback
    nextFrameTime      time.Time

    // Timers
    timers      *TimerHeap
    intervals   map[int]*Interval

    // Control
    running     bool
    stopCh      chan struct{}
}

func (el *EventLoop) Run() {
    el.running = true
    defer func() { el.running = false }()

    for {
        select {
        case <-el.stopCh:
            return

        default:
            // Process microtasks
            el.processMicrotasks()

            // Process next task
            if task := el.tasks.Dequeue(); task != nil {
                task.Execute()
                el.processMicrotasks()
            }

            // Check timers
            el.processTimers()

            // Animation frames (60 FPS)
            if time.Now().After(el.nextFrameTime) {
                el.processAnimationFrame()
                el.nextFrameTime = time.Now().Add(16666667 * time.Nanosecond)
            }
        }
    }
}
```

### 6. HTML Parsing (internal/parser)

Built on golang.org/x/net/html with custom DOM construction.

```go
// internal/parser/parser.go
type Parser struct {
    tokenizer *html.Tokenizer
    doc       *dom.Document
    errors    []error
}

func (p *Parser) Parse(r io.Reader) (*dom.Document, error) {
    p.tokenizer = html.NewTokenizer(r)
    p.doc = dom.NewDocument()

    // Build DOM tree
    if err := p.parseDocument(); err != nil {
        return nil, err
    }

    // Post-process
    p.assignIDs()
    p.computeStyles()

    return p.doc, nil
}
```

### 7. CSS Selectors (internal/css)

Leverages PuerkitoBio/goquery for selector parsing with custom matching.

```go
// internal/css/selector.go
type Selector struct {
    raw      string
    compiled cascadia.Selector
}

type SelectorEngine struct {
    cache map[string]*Selector
    mu    sync.RWMutex
}

func (se *SelectorEngine) QuerySelector(root dom.Node, selector string) dom.Element {
    sel := se.getOrCompile(selector)
    return se.findFirst(root, sel)
}

func (se *SelectorEngine) QuerySelectorAll(root dom.Node, selector string) dom.NodeList {
    sel := se.getOrCompile(selector)
    return se.findAll(root, sel)
}
```

## Testing Framework

### Test Harness Design

```go
// testing/harness.go
type TestHarness struct {
    dom         *DOMulator
    network     *NetworkMock
    console     *ConsoleCapture

    // Assertion tracking
    assertions  []Assertion
    failures    []AssertionFailure

    // Interaction recording
    interactions []Interaction

    // Timing
    clock       *Clock
}
```

### Assertion System

```go
// testing/assertions.go
type Assertion interface {
    Assert() error
    Description() string
}

type ElementAssertion struct {
    selector string
    element  dom.Element
    harness  *TestHarness
}

func (ea *ElementAssertion) HasText(expected string) *ElementAssertion {
    ea.harness.recordAssertion(&TextAssertion{
        element:  ea.element,
        expected: expected,
    })
    return ea
}
```

## Performance Considerations

### 1. Memory Management

- Object pooling for frequently created objects (Events, Nodes)
- Lazy loading of expensive properties (innerHTML, computed styles)
- Weak references for event listeners to prevent leaks

### 2. Execution Optimization

- Selector caching with LRU eviction
- Batch DOM updates within microtasks
- Throttled mutation observers

### 3. Benchmarking

Performance benchmarks will be implemented to ensure DOMulator meets its speed targets compared to headless browser solutions.

## Security Considerations

### JavaScript Isolation

- Each DOMulator instance has its own JavaScript runtime
- No access to host filesystem or network unless explicitly provided
- Configurable resource limits (memory, execution time)

### Content Security

- Optional CSP enforcement
- Script sanitization options
- Configurable feature policies

## Extension Points

### Plugin Architecture

```go
type Plugin interface {
    Name() string
    Install(env *Environment) error
}

type Environment struct {
    Window   *Window
    Document *Document
    Runtime  *Runtime
}
```

### Custom Elements

```go
type CustomElementRegistry struct {
    definitions map[string]CustomElementDefinition
}

type CustomElementDefinition struct {
    Constructor goja.Callable
    Callbacks   CustomElementCallbacks
}
```

## Future Considerations

### Planned Features

- **WebComponents** - Full custom elements and shadow DOM
- **Service Workers** - For PWA testing
- **WebAssembly** - WASM module loading
- **CSS Animations** - CSSOM and animation APIs
- **WebRTC** - Mock peer connections

### Potential Optimizations

- Parallel selector matching
- JIT compilation for hot paths
- Shared memory for worker threads
- Persistent caching of parsed JavaScript
