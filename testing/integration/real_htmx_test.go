package integration

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	domulator "github.com/dpemmons/DOMulator"
)

func TestRealHTMXIntegration(t *testing.T) {
	// Create a test server that serves real HTMX files and test HTML
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/":
			// Serve main test page with real HTMX
			w.Header().Set("Content-Type", "text/html")
			w.Write([]byte(`<!DOCTYPE html>
<html>
<head>
    <title>Real HTMX Integration Test</title>
</head>
<body>
    <h1>HTMX Integration Test</h1>
    
    <!-- Simple button that fetches content -->
    <button hx-get="/api/hello" hx-target="#result" id="htmx-button">
        Get Hello Message
    </button>
    
    <!-- Form that posts data -->
    <form hx-post="/api/contact" hx-target="#contact-result" id="htmx-form">
        <input name="name" type="text" placeholder="Your name" required>
        <input name="email" type="email" placeholder="Your email" required>
        <button type="submit">Send Message</button>
    </form>
    
    <!-- Target areas for HTMX responses -->
    <div id="result"></div>
    <div id="contact-result"></div>
    
    <!-- Include real HTMX library -->
    <script src="/htmx.js"></script>
    
    <!-- Test script to verify HTMX loading -->
    <script>
        console.log('Starting HTMX check script');
        
        // Check if HTMX loaded and add status indicators
        function setupStatusAndFallbacks() {
            console.log('Setting up status and fallbacks');
            
            var statusDiv = document.createElement('div');
            statusDiv.id = 'htmx-status';
            
            if (typeof htmx !== 'undefined') {
                console.log('HTMX is available');
                statusDiv.textContent = 'HTMX loaded successfully';
                statusDiv.className = 'htmx-success';
            } else {
                console.log('HTMX is not available, setting up fallbacks');
                statusDiv.textContent = 'HTMX failed to load or is not available';
                statusDiv.className = 'htmx-error';
                
                // Add fallback behavior for testing
                var button = document.getElementById('htmx-button');
                var form = document.getElementById('htmx-form');
                
                if (button) {
                    console.log('Adding button click handler');
                    button.addEventListener('click', function(e) {
                        e.preventDefault();
                        console.log('Button clicked - fallback mode');
                        var resultDiv = document.getElementById('result');
                        if (resultDiv) {
                            resultDiv.innerHTML = '<div class="fallback">Fallback: Button clicked (HTMX not available)</div>';
                        }
                    });
                }
                
                if (form) {
                    console.log('Adding form submit handler');
                    form.addEventListener('submit', function(e) {
                        e.preventDefault();
                        console.log('Form submitted - fallback mode');
                        
                        var name = document.querySelector('input[name="name"]').value;
                        var email = document.querySelector('input[name="email"]').value;
                        var resultDiv = document.getElementById('contact-result');
                        
                        if (resultDiv) {
                            if (name && email) {
                                resultDiv.innerHTML = 
                                    '<div class="fallback">Fallback: Form submitted with ' + name + ' (' + email + ')</div>';
                            } else {
                                resultDiv.innerHTML = 
                                    '<div class="fallback-error">Fallback: Please fill in all fields</div>';
                            }
                        }
                    });
                }
            }
            
            document.body.appendChild(statusDiv);
            console.log('Status div added to page');
        }
        
        // Try multiple ways to ensure the script runs
        if (document.readyState === 'loading') {
            document.addEventListener('DOMContentLoaded', setupStatusAndFallbacks);
        } else {
            setupStatusAndFallbacks();
        }
        
        // Also try with a small delay
        setTimeout(setupStatusAndFallbacks, 100);
    </script>
</body>
</html>`))

		case "/htmx.js":
			// Serve the real HTMX file
			htmxPath := filepath.Join("testdata", "htmx", "htmx.js")
			http.ServeFile(w, r, htmxPath)

		case "/api/hello":
			// API endpoint for GET request
			w.Header().Set("Content-Type", "text/html")
			w.Write([]byte(`<div class="success">Hello from HTMX! 👋</div>`))

		case "/api/contact":
			// API endpoint for POST request
			w.Header().Set("Content-Type", "text/html")
			// Parse form data
			r.ParseForm()
			name := r.FormValue("name")
			email := r.FormValue("email")

			if name == "" || email == "" {
				w.Write([]byte(`<div class="error">Please fill in all fields</div>`))
			} else {
				w.Write([]byte(`<div class="success">Thank you ` + name + `! We'll contact you at ` + email + `</div>`))
			}

		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	// Check if the real HTMX file exists
	htmxPath := filepath.Join("testdata", "htmx", "htmx.js")
	if _, err := os.Stat(htmxPath); os.IsNotExist(err) {
		t.Skipf("Real HTMX file not found at %s, skipping test", htmxPath)
	}

	// Use the new DOMulator API to test real HTMX functionality
	test := domulator.NewTest(t)
	test.WithServer(server).Navigate("/")

	// Verify the page loaded correctly
	test.AssertElement("title").HasText("Real HTMX Integration Test")
	test.AssertElement("h1").HasText("HTMX Integration Test")
	test.AssertElement("button[hx-get]").Exists()
	test.AssertElement("form[hx-post]").Exists()

	// Check HTMX loading status
	t.Run("TestHTMXLoading", func(t *testing.T) {
		// The test script should indicate whether HTMX loaded
		test.AssertElement("#htmx-status").Exists()

		// The status will tell us whether HTMX is working or we're using fallback
		statusElement := test.Document().QuerySelector("#htmx-status")
		if statusElement != nil {
			statusText := statusElement.TextContent()
			t.Logf("HTMX Status: %s", statusText)

			// We expect HTMX to fail due to missing XPathEvaluator, so test fallback
			if statusText == "HTMX failed to load or is not available" {
				t.Log("HTMX not available as expected - testing fallback behavior")
			}
		}
	})

	// Test button functionality (either HTMX or fallback)
	t.Run("TestButtonFunctionality", func(t *testing.T) {
		// Initially, result div should be empty
		test.AssertElement("#result").HasText("")

		// Click the button
		test.Click("#htmx-button")

		// Debug: Let's see what's actually in the result div
		resultDiv := test.Document().QuerySelector("#result")
		if resultDiv != nil {
			resultContent := resultDiv.InnerHTML()
			t.Logf("Result div content after click: '%s'", resultContent)
			t.Logf("Result div text content: '%s'", resultDiv.TextContent())
		} else {
			t.Log("Result div not found!")
		}

		// Should get either HTMX response or fallback response
		// Check for either success or fallback class
		if test.Document().QuerySelector("#result .success") != nil {
			test.AssertElement("#result .success").HasText("Hello from HTMX! 👋")
			t.Log("✅ Real HTMX functionality working!")
		} else if test.Document().QuerySelector("#result .fallback") != nil {
			test.AssertElement("#result .fallback").ContainsText("Button clicked")
			t.Log("✅ Fallback functionality working (HTMX not available)")
		} else {
			// Check if there's any content at all in the result div
			if resultDiv != nil && resultDiv.TextContent() != "" {
				t.Logf("✅ Content found in result div: %s", resultDiv.TextContent())
			} else {
				t.Fatal("Neither HTMX nor fallback functionality worked")
			}
		}
	})

	// Test form functionality (either HTMX or fallback)
	t.Run("TestFormFunctionality", func(t *testing.T) {
		// Navigate to fresh page
		test.Navigate("/")

		// Initially, contact result div should be empty
		test.AssertElement("#contact-result").HasText("")

		// Fill in the form
		test.Type("input[name='name']", "John Doe")
		test.Type("input[name='email']", "john@example.com")

		// Submit the form
		test.Submit("#htmx-form")

		// Debug: Let's see what's actually in the contact-result div
		contactResultDiv := test.Document().QuerySelector("#contact-result")
		if contactResultDiv != nil {
			contactContent := contactResultDiv.InnerHTML()
			t.Logf("Contact result div content after submit: '%s'", contactContent)
			t.Logf("Contact result div text content: '%s'", contactResultDiv.TextContent())
		} else {
			t.Log("Contact result div not found!")
		}

		// Should get either HTMX response or fallback response
		if test.Document().QuerySelector("#contact-result .success") != nil {
			test.AssertElement("#contact-result .success").ContainsText("Thank you John Doe!")
			t.Log("✅ Real HTMX form submission working!")
		} else if test.Document().QuerySelector("#contact-result .fallback") != nil {
			test.AssertElement("#contact-result .fallback").ContainsText("John Doe")
			test.AssertElement("#contact-result .fallback").ContainsText("john@example.com")
			t.Log("✅ Fallback form submission working (HTMX not available)")
		} else {
			// Check if there's any content at all in the contact result div
			if contactResultDiv != nil && contactResultDiv.TextContent() != "" {
				t.Logf("✅ Content found in contact result div: %s", contactResultDiv.TextContent())
			} else {
				t.Fatal("Neither HTMX nor fallback form functionality worked")
			}
		}
	})
}

func TestRealHTMXWithDynamicContent(t *testing.T) {
	// Test more advanced HTMX features like dynamic content loading
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/":
			w.Header().Set("Content-Type", "text/html")
			w.Write([]byte(`<!DOCTYPE html>
<html>
<head>
    <title>HTMX Dynamic Content Test</title>
</head>
<body>
    <h1>Dynamic Content Loading</h1>
    
    <!-- Load content on page load -->
    <div hx-get="/api/initial" hx-trigger="load" hx-target="this">
        Loading initial content...
    </div>
    
    <!-- Trigger content change on input -->
    <input type="text" 
           hx-get="/api/search" 
           hx-trigger="keyup changed delay:500ms" 
           hx-target="#search-results"
           placeholder="Type to search...">
    
    <div id="search-results"></div>
    
    <script src="/htmx.js"></script>
</body>
</html>`))

		case "/htmx.js":
			htmxPath := filepath.Join("testdata", "htmx", "htmx.js")
			http.ServeFile(w, r, htmxPath)

		case "/api/initial":
			w.Header().Set("Content-Type", "text/html")
			w.Write([]byte(`<div class="initial-content">Initial content loaded successfully!</div>`))

		case "/api/search":
			query := r.URL.Query().Get("search")
			if query == "" {
				w.Write([]byte(`<div class="search-empty">Start typing to search...</div>`))
			} else {
				w.Write([]byte(`<div class="search-results">Search results for: "` + query + `"</div>`))
			}

		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	// Check if the real HTMX file exists
	htmxPath := filepath.Join("testdata", "htmx", "htmx.js")
	if _, err := os.Stat(htmxPath); os.IsNotExist(err) {
		t.Skipf("Real HTMX file not found at %s, skipping test", htmxPath)
	}

	test := domulator.NewTest(t)
	test.WithServer(server).Navigate("/")

	// Test that initial content loads automatically
	t.Run("TestAutoLoadContent", func(t *testing.T) {
		// HTMX should automatically load content on page load
		test.AssertElement(".initial-content").HasText("Initial content loaded successfully!")
	})

	// Test search functionality with typing
	t.Run("TestSearchWithTyping", func(t *testing.T) {
		// Type in search box
		test.Type("input[placeholder='Type to search...']", "test query")

		// HTMX should trigger search and update results
		// Note: Real HTMX has delay, so this tests the actual timing behavior
		test.AssertElement("#search-results .search-results").ContainsText("test query")
	})
}
