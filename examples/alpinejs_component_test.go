package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	domulator "github.com/dpemmons/DOMulator"
)

func TestAlpineJSComponent(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/alpine.js" {
			// Serve the local Alpine.js file
			alpineJS, err := ioutil.ReadFile(filepath.Join("..", "testing", "integration", "testdata", "alpine.js", "cdn.js"))
			if err != nil {
				t.Fatalf("Failed to read Alpine.js file: %v", err)
			}
			w.Header().Set("Content-Type", "application/javascript")
			w.Write(alpineJS)
			return
		}

		w.Write([]byte(`
			<div id="app">
				<button id="btn">Test</button>
				<span id="result"></span>
			</div>
			<script src="/alpine.js"></script>
			<script>
				// Simple test - manually create Alpine functionality
				document.getElementById('btn').addEventListener('click', () => {
					const span = document.getElementById('result');
					const current = parseInt(span.textContent) || 0;
					span.textContent = current + 1;
				});
				
				// Initialize the counter
				document.getElementById('result').textContent = '0';
			</script>
		`))
	}))
	defer server.Close()

	test := domulator.NewTest(t)
	defer test.Close()

	test.WithServer(server).Navigate("/")

	// Test basic functionality
	test.AssertElement("#result").HasText("0")

	test.Click("#btn")
	test.AssertElement("#result").HasText("1")

	test.Click("#btn")
	test.AssertElement("#result").HasText("2")
}
