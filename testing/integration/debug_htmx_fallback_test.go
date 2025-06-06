package integration

import (
	"net/http"
	"net/http/httptest"
	"testing"

	domulator "github.com/dpemmons/DOMulator"
)

func TestDebugHTMXFallbackLogic(t *testing.T) {
	// Create test that replicates the exact HTML structure and JS logic from the failing HTMX test
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/":
			w.Header().Set("Content-Type", "text/html")
			w.Write([]byte(`<!DOCTYPE html>
<html>
<head>
    <title>HTMX Fallback Debug Test</title>
</head>
<body>
    <h1>HTMX Integration Test</h1>
    
    <!-- Same button structure as HTMX test -->
    <button hx-get="/api/hello" hx-target="#result" id="htmx-button">
        Get Hello Message
    </button>
    
    <div id="result"></div>
    
    <!-- Same script logic as HTMX test but simplified for debugging -->
    <script>
        console.log('Starting HTMX check script');
        
        function setupStatusAndFallbacks() {
            console.log('Setting up status and fallbacks');
            
            var statusDiv = document.createElement('div');
            statusDiv.id = 'htmx-status';
            
            // HTMX won't be available in our test environment
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
                
                if (button) {
                    console.log('Adding button click handler');
                    button.addEventListener('click', function(e) {
                        e.preventDefault();
                        console.log('Button clicked - fallback mode');
                        var resultDiv = document.getElementById('result');
                        if (resultDiv) {
                            resultDiv.innerHTML = '<div class="fallback">Fallback: Button clicked (HTMX not available)</div>';
                            console.log('Result div updated with fallback content');
                        } else {
                            console.log('Could not find result div!');
                        }
                    });
                    console.log('Button click handler attached');
                } else {
                    console.log('Could not find htmx-button!');
                }
            }
            
            document.body.appendChild(statusDiv);
            console.log('Status div added to page');
        }
        
        // Try the exact same setup methods as the HTMX test
        if (document.readyState === 'loading') {
            console.log('Document state is loading, adding DOMContentLoaded listener');
            document.addEventListener('DOMContentLoaded', setupStatusAndFallbacks);
        } else {
            console.log('Document already loaded, calling setup directly');
            setupStatusAndFallbacks();
        }
        
        // Also try with a small delay like the original
        setTimeout(function() {
            console.log('Timeout executing, calling setup again');
            setupStatusAndFallbacks();
        }, 100);
        
        console.log('Script setup completed');
    </script>
</body>
</html>`))
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	test := domulator.NewTest(t)
	test.SetDebugMode(true)
	test.WithServer(server).Navigate("/")

	// Check if the status div was created
	statusDiv := test.Document().QuerySelector("#htmx-status")
	if statusDiv != nil {
		statusText := statusDiv.TextContent()
		t.Logf("Status div text: '%s'", statusText)

		if statusText == "HTMX failed to load or is not available" {
			t.Log("✅ Fallback status detection working")
		} else {
			t.Log("❌ Unexpected status text")
		}
	} else {
		t.Log("❌ Status div not found")
	}

	// Check initial result div state
	resultDiv := test.Document().QuerySelector("#result")
	if resultDiv != nil {
		initialText := resultDiv.TextContent()
		t.Logf("Initial result div text: '%s'", initialText)
	}

	// Click the button
	test.Click("#htmx-button")

	// Check if fallback content was added
	if resultDiv != nil {
		resultHTML := resultDiv.InnerHTML()
		resultText := resultDiv.TextContent()
		t.Logf("After click - Result div HTML: '%s'", resultHTML)
		t.Logf("After click - Result div text: '%s'", resultText)

		if resultText == "Fallback: Button clicked (HTMX not available)" {
			t.Log("✅ Fallback button functionality working")
		} else {
			t.Log("❌ Fallback button functionality not working")
		}
	}
}
