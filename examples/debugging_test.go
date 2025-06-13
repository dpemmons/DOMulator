package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	domulator "github.com/dpemmons/DOMulator"
)

func TestDebugging(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
			<div id="mysterious-div" class="container active" data-value="123">
				<p>Some content</p>
			</div>
			<script>
				console.log("Hello from test!", {detail: 123});
				console.error("An error occurred");
			</script>
		`))
	}))
	defer server.Close()

	// Enable console output from JavaScript
	test := domulator.NewTest(t)
	defer test.Close()
	test.SetDebugMode(true) // Enables verbose JS console output to its default (e.g., terminal)

	test.WithServer(server).Navigate("/")

	// ... perform actions that might log to console ...
	test.ExecuteScript(`console.log("Hello from test!", {detail: 123}); console.error("An error occurred");`)

	// Disable debug mode to return to clean output
	test.SetDebugMode(false)
	test.ExecuteScript(`
		console.log('This will be suppressed again');
	`)

	// Console capture and other debugging features are planned but not yet implemented
	// consoleCapture := test.CaptureConsole(t) // Planned feature
	// messages := consoleCapture.GetAllMessages()

	// Take DOM snapshots (This feature is not yet implemented)
	// snapshot := test.TakeSnapshot()
	// fmt.Println(snapshot.HTML)

	// Debug specific elements using the correct API
	elemNode := test.Document().QuerySelector("#mysterious-div") // Returns dom.Node
	if elemNode != nil {
		fmt.Printf("Found element: %s\n", elemNode.NodeName())
		fmt.Printf("Text content: %s\n", elemNode.TextContent())
		// Additional debugging would require casting to *dom.Element for attribute access
		// if elem, ok := elemNode.(*dom.Element); ok {
		//     fmt.Printf("Element: %s\n", elem.TagName())
		//     fmt.Printf("ID: %s\n", elem.GetAttribute("id"))
		//     fmt.Printf("Classes: %s\n", elem.GetAttribute("class"))
		// }
	}
}
