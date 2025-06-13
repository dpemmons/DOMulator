package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	domulator "github.com/dpemmons/DOMulator"
)

func TestHTMXInfiniteScroll(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/":
			w.Write([]byte(`
				<div id="items">
					<div class="item">Item 1</div>
					<div class="item">Item 2</div>
				</div>
				<button id="load-more">Load More</button>
				<script>
					let itemCount = 2;
					document.getElementById('load-more').addEventListener('click', function() {
						const container = document.getElementById('items');
						
						// Add two new items
						for (let i = 0; i < 2; i++) {
							itemCount++;
							const item = document.createElement('div');
							item.className = 'item';
							item.textContent = 'Item ' + itemCount;
							container.appendChild(item);
						}
					});
				</script>
			`))
		}
	}))
	defer server.Close()

	test := domulator.NewTest(t)
	defer test.Close()
	test.WithServer(server).Navigate("/")
	test.AssertElement(".item").Count(2) // Initial load

	// Trigger loading more items
	test.Click("#load-more")
	test.FlushMicrotasks()

	test.AssertElement(".item").Count(4)
}
