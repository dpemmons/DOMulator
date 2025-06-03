package testing

import (
	"net/http"
	"testing"
)

// Example using direct HTML loading
func TestDirectHTMLExample(t *testing.T) {
	harness := NewTestHarness()
	defer harness.Close()

	html := `
	<!DOCTYPE html>
	<html>
	<head><title>Test Page</title></head>
	<body>
		<h1 id="main-title">Welcome</h1>
		<form id="contact-form">
			<input type="text" id="name" name="name" placeholder="Enter your name" />
			<input type="email" id="email" name="email" placeholder="Enter your email" />
			<button type="submit" id="submit-btn">Submit</button>
		</form>
		<div id="output" style="display:block;">Initial content</div>
	</body>
	</html>`

	// Load the HTML and test interactions
	harness.LoadHTML(html)

	harness.AssertDocument().HasTitle("Test Page")
	harness.AssertElement("#main-title").Exists().HasText("Welcome")
	harness.AssertElement("#contact-form").Exists()
	harness.AssertElement("#output").IsVisible().HasText("Initial content")

	// Simulate user interaction
	harness.Type("#name", "John Doe")
	harness.Type("#email", "john@example.com")
	harness.Click("#submit-btn")

	// Verify the form values are set
	harness.AssertElement("#name").HasValue("John Doe")
	harness.AssertElement("#email").HasValue("john@example.com")

	// Event was dispatched successfully (no crashes)
	harness.AssertElement("#submit-btn").Exists()
}

// Example using HTTP handler integration
func TestHTTPHandlerExample(t *testing.T) {
	// Create a simple HTTP handler
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`
		<!DOCTYPE html>
		<html>
		<head><title>Home Page</title></head>
		<body>
			<h1>Welcome to our site!</h1>
			<nav>
				<a href="/about" id="about-link">About</a>
				<a href="/contact" id="contact-link">Contact</a>
			</nav>
		</body>
		</html>`))
	})

	mux.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`
		<!DOCTYPE html>
		<html>
		<head><title>About Us</title></head>
		<body>
			<h1>About Our Company</h1>
			<p>We are a leading technology company.</p>
		</body>
		</html>`))
	})

	harness := NewTestHarness().SetHandler(mux)
	defer harness.Close()

	// Test the home page
	harness.Navigate("/")

	harness.AssertDocument().HasTitle("Home Page")
	harness.AssertElement("h1").HasText("Welcome to our site!")
	harness.AssertElement("#about-link").Exists().HasAttribute("href", "/about")
	harness.AssertElement("#contact-link").Exists().HasAttribute("href", "/contact")

	// Test navigation
	response := harness.GET("/about")

	if response.Status != 200 {
		t.Fatalf("Expected status 200, got %d", response.Status)
	}

	harness.LoadHTML(response.Body)
	harness.AssertDocument().HasTitle("About Us")
	harness.AssertElement("h1").HasText("About Our Company")
	harness.AssertElement("p").ContainsText("leading technology company")
}

// Example with network mocking
func TestNetworkMockingExample(t *testing.T) {
	harness := NewTestHarness()
	defer harness.Close()

	// Set up network mocks
	harness.mocks.MockGET("/api/users", MockResponse{
		Status: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: `[{"id": 1, "name": "John Doe"}, {"id": 2, "name": "Jane Smith"}]`,
	})

	// Test the mocked response
	response := harness.GET("/api/users")

	if response.Status != 200 {
		t.Fatalf("Expected status 200, got %d", response.Status)
	}

	if response.Headers["Content-Type"] != "application/json" {
		t.Fatalf("Expected JSON content type, got %s", response.Headers["Content-Type"])
	}

	if !containsText(response.Body, "John Doe") {
		t.Fatalf("Expected response to contain 'John Doe', got: %s", response.Body)
	}
}

// Example with basic interaction testing
func TestInteractionExample(t *testing.T) {
	harness := NewTestHarness()
	defer harness.Close()

	html := `
	<!DOCTYPE html>
	<html>
	<head><title>Interactive Page</title></head>
	<body>
		<button id="counter-btn">Count: 0</button>
		<ul id="todo-list">
			<li>Existing todo</li>
		</ul>
		<input type="text" id="todo-input" placeholder="Add a todo" />
		<button id="add-todo">Add Todo</button>
		<input type="checkbox" id="checkbox" />
		<select id="select">
			<option value="option1">Option 1</option>
			<option value="option2">Option 2</option>
		</select>
	</body>
	</html>`

	harness.LoadHTML(html)

	// Test basic element interactions
	harness.AssertElement("#counter-btn").HasText("Count: 0")
	harness.Click("#counter-btn") // Events are dispatched successfully
	harness.AssertElement("#counter-btn").Exists()

	// Test form interactions
	harness.Type("#todo-input", "Buy groceries")
	harness.AssertElement("#todo-input").HasValue("Buy groceries")

	harness.Check("#checkbox")
	harness.AssertElement("#checkbox").IsChecked()

	harness.Uncheck("#checkbox")
	harness.AssertElement("#checkbox").IsNotChecked()

	harness.Select("#select", "option2")
	harness.AssertElement("#select").HasValue("option2")

	// Test existing todo
	harness.AssertElement("#todo-list li").Count(1).HasText("Existing todo")

	// Test hover and focus interactions
	harness.Hover("#add-todo")
	harness.Focus("#todo-input")
	harness.Blur("#todo-input")
	harness.Unhover("#add-todo")
}

// Example demonstrating form testing
func TestFormExample(t *testing.T) {
	harness := NewTestHarness()
	defer harness.Close()

	html := `
	<!DOCTYPE html>
	<html>
	<body>
		<form id="signup-form">
			<input type="email" id="email" required />
			<input type="password" id="password" minlength="8" required />
			<input type="password" id="confirm-password" required />
			<button type="submit" id="submit">Sign Up</button>
			<div id="error-message" style="display:none;color:red;">Hidden message</div>
		</form>
	</body>
	</html>`

	harness.LoadHTML(html)

	// Test form element interactions
	harness.Type("#email", "user@example.com")
	harness.AssertElement("#email").HasValue("user@example.com")

	harness.Type("#password", "password123")
	harness.AssertElement("#password").HasValue("password123")

	harness.Type("#confirm-password", "password123")
	harness.AssertElement("#confirm-password").HasValue("password123")

	// Test form submission (event dispatching)
	harness.Submit("#signup-form")

	// Test element visibility
	harness.AssertElement("#error-message").IsHidden()
	harness.AssertElement("#error-message").HasText("Hidden message")
}

// Helper function for the example
func containsText(text, substring string) bool {
	return len(text) >= len(substring) &&
		text[len(text)-len(substring):len(text)] == substring ||
		text[:len(substring)] == substring ||
		findSubstring(text, substring)
}

func findSubstring(text, substring string) bool {
	for i := 0; i <= len(text)-len(substring); i++ {
		if text[i:i+len(substring)] == substring {
			return true
		}
	}
	return false
}
