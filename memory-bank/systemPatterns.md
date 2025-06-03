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
