package integration

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	domulator "github.com/dpemmons/DOMulator"
)

func TestSimpleNavigation(t *testing.T) {
	// Create a simple test server identical to the debug test
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Simple server received: %s %s\n", r.Method, r.URL.Path)
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`
			<html>
				<head><title>Test App</title></head>
				<body>
					<form>
						<input name="email" type="email" required>
						<button type="submit">Send</button>
					</form>
					<div id="result"></div>
				</body>
			</html>
		`))
	}))
	defer server.Close()

	fmt.Printf("Simple test server URL: %s\n", server.URL)

	// Use the new DOMulator API exactly like the working debug test
	test := domulator.NewTest(t)
	test.WithServer(server).Navigate("/")

	// Debug: Print the actual HTML that was loaded
	doc := test.Document()
	if doc != nil {
		fmt.Printf("Simple test HTML: %s\n", doc.DocumentElement().OuterHTML())
	}

	// Try just one simple assertion
	test.AssertElement("title").HasText("Test App")
}
