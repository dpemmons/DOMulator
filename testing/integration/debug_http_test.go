package integration

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	domulator "github.com/dpemmons/DOMulator"
)

func TestDebugHTTPRequests(t *testing.T) {
	// Create a simple test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Server received request: %s %s\n", r.Method, r.URL.Path)
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`<html><head><title>Debug Test</title></head><body><h1>Hello World</h1></body></html>`))
	}))
	defer server.Close()

	fmt.Printf("Test server URL: %s\n", server.URL)

	// Test direct HTTP request first
	test := domulator.NewTest(t)
	response := test.GET(server.URL + "/")
	fmt.Printf("Direct GET response: %s\n", response.Body)

	// Test with server configuration
	test2 := domulator.NewTest(t)
	test2.WithServer(server)
	response2 := test2.GET("/")
	fmt.Printf("Server-configured GET response: %s\n", response2.Body)

	// Test Navigate
	test3 := domulator.NewTest(t)
	test3.WithServer(server).Navigate("/")

	// Check what document we got
	doc := test3.Document()
	if doc != nil {
		fmt.Printf("Navigate document HTML: %s\n", doc.DocumentElement().OuterHTML())
		title := doc.QuerySelector("title")
		if title != nil {
			fmt.Printf("Title element found: %s\n", title.TextContent())
		} else {
			fmt.Printf("No title element found\n")
		}
	} else {
		fmt.Printf("No document found after Navigate\n")
	}
}
