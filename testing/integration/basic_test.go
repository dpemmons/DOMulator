package integration

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/dpemmons/DOMulator/internal/browser/fetch"
	"github.com/dpemmons/DOMulator/internal/browser/resources"
	"github.com/dpemmons/DOMulator/internal/dom"
	"github.com/dpemmons/DOMulator/internal/js"
	"github.com/dpemmons/DOMulator/internal/parser"
)

// runtimeAdapter adapts js.Runtime to resources.JSRuntime interface
type runtimeAdapter struct {
	runtime *js.Runtime
}

func (ra *runtimeAdapter) RunScript(name, code string) (interface{}, error) {
	result, err := ra.runtime.RunScript(name, code)
	if err != nil {
		return nil, err
	}
	return result.Export(), nil // Convert goja.Value to interface{}
}

func (ra *runtimeAdapter) VM() interface{} {
	return ra.runtime.VM()
}

func TestBasicDOMManipulation(t *testing.T) {
	// Create HTTP test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/index.html":
			w.Header().Set("Content-Type", "text/html")
			w.Write([]byte(`<!DOCTYPE html>
<html>
<head>
    <title>Basic Integration Test</title>
</head>
<body>
    <div id="container">
        <p>Original content</p>
    </div>
    <script src="/app.js"></script>
</body>
</html>`))
		case "/app.js":
			w.Header().Set("Content-Type", "application/javascript")
			w.Write([]byte(`
// Add a new paragraph to the container
var container = document.getElementById('container');
var newPara = document.createElement('p');
newPara.textContent = 'Added by JavaScript';
newPara.id = 'dynamic-content';
container.appendChild(newPara);

// Add a class to show the script executed
document.body.className = 'script-executed';
`))
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	// Fetch and parse the HTML document
	htmlURL := server.URL + "/index.html"
	resp, err := http.Get(htmlURL)
	if err != nil {
		t.Fatalf("Failed to fetch HTML: %v", err)
	}
	defer resp.Body.Close()

	// Parse the HTML
	parser := parser.NewParser()
	document, err := parser.Parse(resp.Body)
	if err != nil {
		t.Fatalf("Failed to parse HTML: %v", err)
	}

	// Create JavaScript runtime with document
	runtime := js.New(document)
	runtimeAdapter := &runtimeAdapter{runtime: runtime}

	// Create fetch API and resource manager
	fetchAPI := fetch.New(nil) // No network mocks, use real HTTP
	resourceManager := resources.New(document, fetchAPI, runtimeAdapter)

	// Find all script elements and load them
	scripts := document.GetElementsByTagName("script")
	scriptsToLoad := make([]*dom.Element, 0)

	for i := 0; i < scripts.Length(); i++ {
		script := scripts.Item(i)
		if script.GetAttribute("src") != "" {
			scriptsToLoad = append(scriptsToLoad, script)
		}
	}

	// Load and execute scripts
	for _, script := range scriptsToLoad {
		src := script.GetAttribute("src")
		if src != "" {
			// Convert relative URL to absolute and set it on the script element
			if !strings.HasPrefix(src, "http") {
				src = server.URL + src
				script.SetAttribute("src", src)
			}

			t.Logf("Loading script from: %s", src)

			// Load the script using resource manager
			err := resourceManager.LoadResource(script)
			if err != nil {
				t.Fatalf("Failed to load script: %v", err)
			}

			// Wait for script to load and execute
			// Check if there are any active tasks
			maxWait := time.Second * 5
			start := time.Now()
			for resourceManager.IsLoading() {
				if time.Since(start) > maxWait {
					t.Fatalf("Script loading timed out")
				}
				time.Sleep(10 * time.Millisecond)
			}

			// Check if any tasks failed
			tasks := resourceManager.GetActiveTasks()
			for _, task := range tasks {
				if task.HasError() {
					t.Fatalf("Script failed to load: %v", task.Error)
				}
				t.Logf("Task state: %s, Content length: %d", task.State, len(task.Content))
			}

			// Add a small delay to ensure script execution completes
			time.Sleep(100 * time.Millisecond)
		}
	}

	// Debug: Check the DOM structure before verification
	t.Logf("DOM structure after script execution:")
	container := document.GetElementById("container")
	if container != nil {
		children := container.ChildNodes()
		t.Logf("Container has %d children", children.Length())
		for i := 0; i < children.Length(); i++ {
			child := children.Item(i)
			t.Logf("  Child %d: type=%d, name=%s", i, int(child.NodeType()), child.NodeName())
			if child.NodeType() == dom.ElementNode {
				elem := child.(*dom.Element)
				t.Logf("    Element: tag=%s, id=%s, text=%s", elem.TagName(), elem.GetAttribute("id"), elem.TextContent())
			}
		}
	}

	// Verify the DOM was modified by JavaScript
	t.Run("VerifyDOMModifications", func(t *testing.T) {
		// Check that the new paragraph was added
		dynamicContent := document.GetElementById("dynamic-content")
		if dynamicContent == nil {
			t.Error("Expected element with id 'dynamic-content' was not found")
			return
		}

		// Check the text content
		if dynamicContent.TextContent() != "Added by JavaScript" {
			t.Errorf("Expected text content 'Added by JavaScript', got '%s'", dynamicContent.TextContent())
		}

		// Check that it's a child of the container
		container := document.GetElementById("container")
		if container == nil {
			t.Error("Container element not found")
			return
		}

		found := false
		children := container.ChildNodes()
		for i := 0; i < children.Length(); i++ {
			child := children.Item(i)
			if child.NodeType() == dom.ElementNode {
				element := child.(*dom.Element)
				if element.GetAttribute("id") == "dynamic-content" {
					found = true
					break
				}
			}
		}

		if !found {
			t.Error("Dynamic content element was not found as child of container")
		}

		// Check that the body class was added
		body := document.Body()
		if body == nil {
			t.Error("Body element not found")
			return
		}

		if body.GetAttribute("class") != "script-executed" {
			t.Errorf("Expected body class 'script-executed', got '%s'", body.GetAttribute("class"))
		}
	})

	// Verify the original content is still there
	t.Run("VerifyOriginalContent", func(t *testing.T) {
		container := document.GetElementById("container")
		if container == nil {
			t.Error("Container element not found")
			return
		}

		// Should have both original and new content
		paragraphs := container.GetElementsByTagName("p")
		if paragraphs.Length() != 2 {
			t.Errorf("Expected 2 paragraphs in container, got %d", paragraphs.Length())
		}

		// Check original paragraph is still there
		originalFound := false
		for i := 0; i < paragraphs.Length(); i++ {
			p := paragraphs.Item(i)
			if p.TextContent() == "Original content" {
				originalFound = true
				break
			}
		}

		if !originalFound {
			t.Error("Original content paragraph not found")
		}
	})
}
