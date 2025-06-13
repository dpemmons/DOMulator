package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	domulator "github.com/dpemmons/DOMulator"
)

func TestVanillaJSApp(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/":
			w.Write([]byte(`
				<input id="name" type="text">
				<button id="greet">Greet</button>
				<div id="message"></div>
				<script src="/app.js"></script>
			`))
		case "/app.js":
			w.Header().Set("Content-Type", "application/javascript")
			w.Write([]byte(`
				document.getElementById('greet').addEventListener('click', () => {
					const name = document.getElementById('name').value;
					document.getElementById('message').textContent = 'Hello, ' + name + '!';
				});
			`))
		}
	}))
	defer server.Close()

	test := domulator.NewTest(t)
	defer test.Close()
	test.WithServer(server).Navigate("/")

	test.Type("#name", "World")
	test.Click("#greet")
	test.AssertElement("#message").HasText("Hello, World!")
}
