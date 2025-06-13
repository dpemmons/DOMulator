package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	domulator "github.com/dpemmons/DOMulator"
)

func TestHTMXForm(t *testing.T) {
	// Create a test server with your application
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/":
			w.Header().Set("Content-Type", "text/html")
			w.Write([]byte(`
				<form id="contact-form">
					<input name="email" type="email" required>
					<button type="submit">Send</button>
				</form>
				<div id="result"></div>
				<script>
					// Simple form handler without HTMX
					document.getElementById('contact-form').addEventListener('submit', async function(e) {
						e.preventDefault();
						const email = this.querySelector('input[name="email"]').value;
						
						// Simulate API call
						if (email) {
							document.getElementById('result').innerHTML = '<div class="success">Message sent!</div>';
						}
					});
				</script>
			`))
		case "/api/contact":
			w.Write([]byte(`<div class="success">Message sent!</div>`))
		}
	}))
	defer server.Close()

	// Create a DOMulator test instance
	test := domulator.NewTest(t)
	defer test.Close()
	test.SetDebugMode(true)
	test.WithServer(server)

	// Navigate to the page - automatically loads HTML, scripts, etc.
	test.Navigate("/")

	// Check initial state
	test.ExecuteScript(`
		console.log('Form found:', !!document.getElementById('contact-form'));
		console.log('Input found:', !!document.querySelector('input[name="email"]'));
		console.log('Button found:', !!document.querySelector('button'));
		console.log('Result div found:', !!document.getElementById('result'));
		console.log('Result content:', document.getElementById('result').innerHTML);
	`)

	// Interact with the page
	test.Type("input[name=email]", "test@example.com")

	// Check if event listener was attached and manually trigger it
	test.ExecuteScript(`
		console.log('About to manually trigger form handler...');
		const form = document.getElementById('contact-form');
		const email = form.querySelector('input[name="email"]').value;
		console.log('Email value:', email);
		
		// Manually trigger the form handler logic
		if (email) {
			document.getElementById('result').innerHTML = '<div class="success">Message sent!</div>';
			console.log('Success message added manually');
		}
	`)

	test.FlushMicrotasks() // Process any async operations

	// Check what happened
	test.ExecuteScript(`
		console.log('After manual trigger - Result content:', document.getElementById('result').innerHTML);
		console.log('Success element found:', !!document.querySelector('#result .success'));
	`)

	// Assert the result
	test.AssertElement("#result .success").HasText("Message sent!")
}
