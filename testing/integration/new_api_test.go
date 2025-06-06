package integration

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	domulator "github.com/dpemmons/DOMulator"
	"github.com/dpemmons/DOMulator/internal/dom"
)

func TestNewAPIWithHTTPTestServer(t *testing.T) {
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

func TestScriptAutoLoading(t *testing.T) {
	// Test that external scripts are automatically loaded and executed
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Script server received: %s %s\n", r.Method, r.URL.Path)

		// Simple path matching without full URL prefix
		switch r.URL.Path {
		case "/":
			w.Write([]byte(`
				<html>
					<body>
						<div id="output"></div>
						<script src="/app.js"></script>
					</body>
				</html>
			`))
		case "/app.js":
			w.Header().Set("Content-Type", "application/javascript")
			w.Write([]byte(`
				document.getElementById('output').textContent = 'Script loaded and executed!';
			`))
		default:
			// Default response for unmatched paths
			w.Write([]byte(`<html><body><div>No match for path: ` + r.URL.Path + `</div></body></html>`))
		}
	}))
	defer server.Close()

	test := domulator.NewTest(t)
	test.WithServer(server).Navigate("/")

	// Debug: Print what HTML was loaded
	doc := test.Document()
	if doc != nil {
		fmt.Printf("Script test HTML: %s\n", doc.DocumentElement().OuterHTML())
	}

	// The script should have automatically executed and updated the div
	test.AssertElement("#output").HasText("Script loaded and executed!")
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
