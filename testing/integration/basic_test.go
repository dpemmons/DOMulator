package integration

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	domulator "github.com/dpemmons/DOMulator"
	"github.com/dpemmons/DOMulator/internal/dom"
	testpkg "github.com/dpemmons/DOMulator/testing"
)

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

	// Use the new DOMulator API - automatic HTML parsing and script loading
	test := domulator.NewTest(t)
	test.WithServer(server).Navigate("/index.html")

	// Verify the DOM was modified by JavaScript
	t.Run("VerifyDOMModifications", func(t *testing.T) {
		// Check that the new paragraph was added with the expected content
		test.AssertElement("#dynamic-content").Exists()
		test.AssertElement("#dynamic-content").HasText("Added by JavaScript")

		// Check that the body class was added
		test.AssertElement("body").HasClass("script-executed")
	})

	// Verify the original content is still there
	t.Run("VerifyOriginalContent", func(t *testing.T) {
		// Check that both original and new content exist
		test.AssertElement("#container").Exists()
		test.AssertElement("#container p").Count(2) // Should have 2 paragraphs

		// Check original paragraph is still there using contains text
		test.AssertElement("#container").ContainsText("Original content")
	})

	// Verify the title was set correctly
	test.AssertElement("title").HasText("Basic Integration Test")
}

func TestNavigateToAbsoluteURL(t *testing.T) {
	// Test that absolute URLs work too
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<h1>Direct Navigation</h1>`))
	}))
	defer server.Close()

	test := domulator.NewTest(t)

	// Navigate directly to the absolute URL (no WithServer needed)
	test.Navigate(server.URL)

	test.AssertElement("h1").HasText("Direct Navigation")
}

func TestInlineScriptExecution(t *testing.T) {
	// Test that inline scripts are automatically executed
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
			<html>
				<body>
					<div id="inline-result"></div>
					<script>
						document.getElementById('inline-result').textContent = 'Inline script executed!';
					</script>
				</body>
			</html>
		`))
	}))
	defer server.Close()

	test := domulator.NewTest(t)
	test.WithServer(server).Navigate("/")

	// The inline script should have automatically executed
	test.AssertElement("#inline-result").HasText("Inline script executed!")
}

func TestWithHTTPTestServer(t *testing.T) {
	// Create a test server with HTMX-style behavior
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/":
			w.Header().Set("Content-Type", "text/html")
			w.Write([]byte(`
				<html>
					<head><title>Test App</title></head>
					<body>
						<form hx-post="/api/contact" hx-target="#result">
							<input name="email" type="email" required>
							<button type="submit">Send</button>
						</form>
						<div id="result"></div>
						<script src="/htmx.js"></script>
					</body>
				</html>
			`))
		case "/htmx.js":
			// Serve a mock HTMX implementation that handles basic form submissions
			w.Header().Set("Content-Type", "application/javascript")
			w.Write([]byte(`
				// Mock HTMX implementation for testing
				(function() {
					console.log('HTMX script starting...');
					
					function handleSubmit(event) {
						console.log('handleSubmit called, event:', event);
						event.preventDefault();
						var form = event.target;
						console.log('form:', form);
						var target = form.getAttribute('hx-target');
						var url = form.getAttribute('hx-post');
						console.log('target:', target, 'url:', url);
						
						if (target && url) {
							// Simple fetch simulation using basic DOM manipulation
							var targetElement = document.querySelector(target);
							console.log('targetElement:', targetElement);
							if (targetElement) {
								targetElement.innerHTML = '<div class="success">Message sent!</div>';
								console.log('Updated target element innerHTML');
							} else {
								console.log('Target element not found!');
							}
						} else {
							console.log('Missing target or url attributes');
						}
					}
					
					// Initialize HTMX-like behavior
					document.addEventListener('DOMContentLoaded', function() {
						console.log('DOMContentLoaded fired in HTMX script');
						var forms = document.querySelectorAll('[hx-post]');
						console.log('Found', forms.length, 'forms with hx-post');
						for (var i = 0; i < forms.length; i++) {
							console.log('Adding submit listener to form', i);
							forms[i].addEventListener('submit', handleSubmit);
						}
					});
					
					console.log('HTMX script loaded');
				})();
			`))
		case "/api/contact":
			// Simulate form submission response
			w.Write([]byte(`<div class="success">Message sent!</div>`))
		}
	}))
	defer server.Close()

	// Use the new DOMulator API
	test := domulator.NewTest(t)

	// Navigate to the page - this should automatically load HTML and scripts
	test.WithServer(server).Navigate("/")

	// Debug: Print the actual HTML that was loaded
	doc := test.Document()
	if doc != nil {
		fmt.Printf("Loaded HTML: %s\n", doc.DocumentElement().OuterHTML())

		// Check if attributes are actually there
		inputs := doc.QuerySelectorAll("input")
		fmt.Printf("Found %d input elements\n", inputs.Length())
		if inputs.Length() > 0 {
			input := inputs.Item(0).(*dom.Element)
			fmt.Printf("Input attributes: name='%s', type='%s'\n",
				input.GetAttribute("name"), input.GetAttribute("type"))
		}
	}

	// Verify the page loaded correctly
	test.AssertElement("title").HasText("Test App")
	test.AssertElement("form").Exists()
	test.AssertElement("input[name=email]").Exists()
	test.AssertElement("#result").Exists()

	// Debug: Test if HTMX script was loaded by checking if the event handler is attached
	test.ExecuteScript("console.log('Testing HTMX setup...');")

	// Test form interaction
	test.Type("input[name=email]", "test@example.com")

	// Debug: Let's try triggering the form submit event directly
	test.Submit("form[hx-post]")

	// Debug: Check what's in the result div after form submission
	resultDiv := test.Document().QuerySelector("#result")
	if resultDiv != nil {
		fmt.Printf("Result div content after submission: '%s'\n", resultDiv.InnerHTML())
	}

	// This should trigger HTMX behavior and update the result div
	// Note: This will test both the automatic script loading and HTMX functionality
	test.AssertElement("#result .success").HasText("Message sent!")
}

func TestBasicDOMCreationWithConsoleCapture(t *testing.T) {
	harness := testpkg.NewTestHarness()
	defer harness.Close()

	consoleCapture := harness.CaptureConsole(t)

	harness.LoadHTML(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>Test Page</title>
		</head>
		<body>
			<h1 id="title">Hello World</h1>
			<p class="content">This is a test paragraph.</p>
			<div id="container">Content here</div>
			<script>
				console.log('DOM loaded successfully');
				console.log('Page title:', document.title);
				
				// Log main elements
				var title = document.getElementById('title');
				if (title) {
					console.log('Title found:', title.textContent);
				} else {
					console.log('Title element not found');
				}
				
				var content = document.querySelector('.content');
				if (content) {
					console.log('Content found:', content.textContent);
				} else {
					console.log('Content element not found');
				}
				
				var container = document.getElementById('container');
				if (container) {
					console.log('Container found:', container.textContent);
				} else {
					console.log('Container element not found');
				}
				
				console.log('DOM inspection complete');
			</script>
		</body>
		</html>
	`)

	// Assert console output
	consoleCapture.AssertLogContains("DOM loaded successfully")
	consoleCapture.AssertLogContains("Page title:")
	consoleCapture.AssertLogContains("Title found:")
	consoleCapture.AssertLogContains("Hello World")
	consoleCapture.AssertLogContains("Content found:")
	consoleCapture.AssertLogContains("This is a test paragraph.")
	consoleCapture.AssertLogContains("Container found:")
	consoleCapture.AssertLogContains("Content here")
	consoleCapture.AssertLogContains("DOM inspection complete")

	// Test basic element existence using harness
	title := harness.Document().QuerySelector("#title")
	if title == nil {
		t.Fatal("Could not find title element")
	}
	if title.TextContent() != "Hello World" {
		t.Errorf("Expected title 'Hello World', got %s", title.TextContent())
	}

	container := harness.Document().QuerySelector("#container")
	if container == nil {
		t.Fatal("Could not find container element")
	}
	if container.TextContent() != "Content here" {
		t.Errorf("Expected container 'Content here', got %s", container.TextContent())
	}

	t.Log("✅ Basic DOM creation with console capture completed successfully!")
}

func TestBasicJavaScriptWithConsoleCapture(t *testing.T) {
	harness := testpkg.NewTestHarness()
	defer harness.Close()

	consoleCapture := harness.CaptureConsole(t)

	harness.LoadHTML(`
		<!DOCTYPE html>
		<html>
		<body>
			<button id="btn">Click me</button>
			<div id="output"></div>
			<script>
				console.log('Setting up DOM manipulation...');
				
				// Set up a simple function we can call
				function updateOutput() {
					console.log('Updating output via function call...');
					document.getElementById('output').textContent = 'Function executed!';
					console.log('Output updated successfully');
				}
				
				// Make function available globally
				window.updateOutput = updateOutput;
				
				console.log('DOM manipulation setup complete');
			</script>
		</body>
		</html>
	`)

	// Assert initial console output
	consoleCapture.AssertLogContains("Setting up DOM manipulation...")
	consoleCapture.AssertLogContains("DOM manipulation setup complete")

	// Test initial state
	output := harness.Document().QuerySelector("#output")
	if output == nil {
		t.Fatal("Could not find output element")
	}
	if output.TextContent() != "" {
		t.Errorf("Expected empty output initially, got %s", output.TextContent())
	}

	// Test function call by executing script
	harness.ExecuteScript(`
		console.log('Calling updateOutput function...');
		window.updateOutput();
	`)

	// Assert console output for function call
	consoleCapture.AssertLogContains("Calling updateOutput function...")
	consoleCapture.AssertLogContains("Updating output via function call...")
	consoleCapture.AssertLogContains("Output updated successfully")

	// Verify the output was updated
	if output.TextContent() != "Function executed!" {
		t.Errorf("Expected output 'Function executed!', got %s", output.TextContent())
	}

	t.Log("✅ Basic JavaScript with console capture completed successfully!")
}
