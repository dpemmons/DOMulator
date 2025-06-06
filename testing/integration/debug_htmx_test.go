package integration

import (
	"net/http"
	"net/http/httptest"
	"testing"

	domulator "github.com/dpemmons/DOMulator"
)

func TestDebugHTMXJavaScript(t *testing.T) {
	// Create a simple test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/":
			w.Header().Set("Content-Type", "text/html")
			w.Write([]byte(`<!DOCTYPE html>
<html>
<head>
    <title>Debug JS Test</title>
</head>
<body>
    <button id="test-button">Click Me</button>
    <div id="result"></div>
    
    <script>
        console.log('Script starting...');
        
        // Test 1: Basic JavaScript execution
        document.body.setAttribute('data-js-loaded', 'true');
        console.log('Basic JS executed');
        
        // Test 2: DOM manipulation
        var resultDiv = document.getElementById('result');
        if (resultDiv) {
            resultDiv.textContent = 'JavaScript is working!';
            console.log('DOM manipulation worked');
        } else {
            console.log('Could not find result div');
        }
        
        // Test 3: Event handler setup
        function setupEventHandler() {
            console.log('Setting up event handler');
            var button = document.getElementById('test-button');
            if (button) {
                button.addEventListener('click', function() {
                    console.log('Button clicked!');
                    var resultDiv = document.getElementById('result');
                    if (resultDiv) {
                        resultDiv.innerHTML = '<span class="clicked">Button was clicked!</span>';
                        console.log('Result updated after click');
                    }
                });
                console.log('Event handler attached');
            } else {
                console.log('Could not find button');
            }
        }
        
        // Try multiple ways to ensure the handler is set up
        if (document.readyState === 'loading') {
            console.log('Document still loading, adding DOMContentLoaded listener');
            document.addEventListener('DOMContentLoaded', setupEventHandler);
        } else {
            console.log('Document already loaded, setting up handler now');
            setupEventHandler();
        }
        
        // Also try with a delay
        setTimeout(function() {
            console.log('Timeout handler executing');
            setupEventHandler();
        }, 100);
        
        console.log('Script completed');
    </script>
</body>
</html>`))
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	test := domulator.NewTest(t)

	// Enable debug mode to see console output
	test.SetDebugMode(true)

	// Navigate to the test page
	test.WithServer(server).Navigate("/")

	// Check if JavaScript executed at all
	bodyElement := test.Document().QuerySelector("body")
	if bodyElement != nil {
		jsLoaded := bodyElement.GetAttribute("data-js-loaded")
		t.Logf("JavaScript loaded attribute: '%s'", jsLoaded)
	}

	// Check if DOM manipulation worked
	resultDiv := test.Document().QuerySelector("#result")
	if resultDiv != nil {
		resultText := resultDiv.TextContent()
		t.Logf("Result div text: '%s'", resultText)

		if resultText == "JavaScript is working!" {
			t.Log("✅ Basic JavaScript execution working")
		} else {
			t.Log("❌ Basic JavaScript execution failed")
		}
	} else {
		t.Log("❌ Result div not found")
	}

	// Test button click
	test.Click("#test-button")

	// Check if event handler worked
	if resultDiv != nil {
		resultHTML := resultDiv.InnerHTML()
		resultText := resultDiv.TextContent()
		t.Logf("After click - Result div HTML: '%s'", resultHTML)
		t.Logf("After click - Result div text: '%s'", resultText)

		if resultText == "Button was clicked!" {
			t.Log("✅ Event handler working")
		} else {
			t.Log("❌ Event handler not working")
		}
	}
}
