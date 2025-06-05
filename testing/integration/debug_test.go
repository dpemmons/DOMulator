package integration

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	domulator "github.com/dpemmons/DOMulator"
)

func TestDebugSimpleNavigation(t *testing.T) {
	// Create a very simple test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Server received request for: %s\n", r.URL.Path)
		w.Write([]byte(`<html><body><h1>Hello World</h1></body></html>`))
	}))
	defer server.Close()

	fmt.Printf("Server URL: %s\n", server.URL)

	// Use the new DOMulator API
	test := domulator.NewTest(t)
	test.WithServer(server)

	// Navigate to the page
	test.Navigate("/")

	// Check if we can get the document
	doc := test.Document()
	if doc == nil {
		t.Fatal("Document is nil")
	}

	// Print the document HTML to see what we got
	fmt.Printf("Document HTML: %s\n", doc.DocumentElement().OuterHTML())

	// Test a simple assertion
	test.AssertElement("h1").Exists()
}
