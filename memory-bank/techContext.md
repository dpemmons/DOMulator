# Technical Context: DOMulator

## Technologies Used
- **Go (Golang)**: The primary programming language for the entire DOMulator library.
- **Goja**: A pure Go implementation of an ECMAScript 5.1+ JavaScript runtime. This is the core engine for executing JavaScript within DOMulator.
- **`golang.org/x/net/html`**: Go's standard library extension for parsing HTML5 documents. Used for the initial parsing of HTML content into a token stream.
- **PuerkitoBio/goquery**: A Go library that brings a jQuery-like syntax to Go for parsing and manipulating HTML documents. Used for CSS selector parsing and matching.
- **Standard Go Libraries**: Extensive use of `net/http`, `context`, `time`, `sync`, `fmt`, `encoding/json`, etc., for various functionalities like network handling, concurrency, and data manipulation.

## Development Setup
- **Go Modules**: Project dependencies are managed using Go Modules.
- **Testing Framework**: Standard Go `testing` package is used for unit and integration tests. Benchmarks are also written using Go's built-in benchmarking tools.
- **No External Build Tools**: The project is designed to be built and tested purely with standard Go commands (`go build`, `go test`).

## Technical Constraints
- **Pure Go**: A strict constraint to avoid CGO, ensuring maximum portability and ease of integration into any Go project without requiring C compilers or specific system libraries.
- **ECMAScript 5.1+ Compatibility**: The JavaScript runtime (Goja) supports ECMAScript 5.1+, meaning newer JavaScript features (ES6+) might require polyfills or transpilation if not natively supported by Goja.
- **Performance Targets**: The core design must prioritize performance, aiming for orders of magnitude faster execution than headless browsers. This influences decisions on data structures, algorithms, and lazy evaluation.
- **Browser API Emulation**: The goal is to emulate *sufficient* browser API compatibility to run real-world frameworks, not necessarily a pixel-perfect, exhaustive implementation of every single browser API. Focus on APIs critical for DOM manipulation, event handling, and network requests.

## Dependencies
- `github.com/dop251/goja`: The JavaScript runtime.
- `github.com/PuerkitoBio/goquery`: For CSS selector parsing and DOM querying.
- `golang.org/x/net/html`: For HTML parsing.
- Other standard Go libraries.

## Tool Usage Patterns
- **`go test ./...`**: For running all tests.
- **`go test -bench=. ./...`**: For running performance benchmarks.
- **`go mod download`**: For managing Go module dependencies.
