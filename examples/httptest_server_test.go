package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	domulator "github.com/dpemmons/DOMulator"
)

func TestMyApp(t *testing.T) {
	// Create your test server with whatever behavior you want
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/":
			w.Write([]byte(`<h1>Home</h1><script src="/app.js"></script>`))
		case "/app.js":
			w.Header().Set("Content-Type", "application/javascript")
			w.Write([]byte(`console.log("App loaded");`))
		case "/api/data":
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"status": "ok"}`))
		}
	}))
	defer server.Close()

	// Use your real application handler instead
	// server := httptest.NewServer(myapp.Handler())

	test := domulator.NewTest(t)
	defer test.Close()
	test.WithServer(server)

	// Navigate automatically loads HTML, scripts, CSS, etc.
	test.Navigate("/")

	// Verify the page loaded correctly
	test.AssertElement("h1").HasText("Home")

	// Relative URLs work automatically
	// test.Navigate("/some/other/path") // Example
}
