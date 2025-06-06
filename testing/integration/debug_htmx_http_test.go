package integration

import (
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	domulator "github.com/dpemmons/DOMulator"
)

func TestDebugHTMXHTTPRequest(t *testing.T) {
	// Test if HTMX can make HTTP requests programmatically
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/":
			w.Header().Set("Content-Type", "text/html")
			w.Write([]byte(`<!DOCTYPE html>
<html>
<head>
    <title>HTMX HTTP Test</title>
</head>
<body>
    <div id="result">Initial content</div>
    <button id="test-button">Test Button</button>
    
    <script src="/htmx.js"></script>
    <script>
        console.log('Script starting - testing HTMX HTTP functionality');
        
        function testHTMXRequest() {
            console.log('Testing HTMX request');
            
            if (typeof htmx !== 'undefined') {
                console.log('HTMX is available, testing request');
                
                // Try to use HTMX programmatically
                var resultDiv = document.getElementById('result');
                
                try {
                    // Test if htmx.ajax exists and works
                    if (typeof htmx.ajax === 'function') {
                        console.log('htmx.ajax is available');
                        htmx.ajax('GET', '/api/test', {target: '#result'});
                        console.log('htmx.ajax called');
                    } else {
                        console.log('htmx.ajax is not available');
                        resultDiv.textContent = 'HTMX loaded but ajax not available';
                    }
                } catch (e) {
                    console.log('Error calling htmx.ajax:', e);
                    resultDiv.textContent = 'HTMX ajax error: ' + e.message;
                }
            } else {
                console.log('HTMX not available');
                document.getElementById('result').textContent = 'HTMX not available';
            }
        }
        
        // Set up button click handler
        document.getElementById('test-button').addEventListener('click', function() {
            console.log('Button clicked, testing HTMX request');
            testHTMXRequest();
        });
        
        // Test immediately on load
        setTimeout(function() {
            console.log('Auto-testing HTMX after load');
            testHTMXRequest();
        }, 100);
    </script>
</body>
</html>`))

		case "/htmx.js":
			htmxPath := filepath.Join("testdata", "htmx", "htmx.js")
			http.ServeFile(w, r, htmxPath)

		case "/api/test":
			t.Logf("API /api/test called - HTMX HTTP request working!")
			w.Header().Set("Content-Type", "text/html")
			w.Write([]byte(`<div class="success">HTMX HTTP request successful!</div>`))

		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	test := domulator.NewTest(t)
	test.SetDebugMode(true)
	test.WithServer(server).Navigate("/")

	// Check the result div after auto-test
	resultDiv := test.Document().QuerySelector("#result")
	if resultDiv != nil {
		resultText := resultDiv.TextContent()
		t.Logf("Result after auto-test: '%s'", resultText)
	}

	// Also test manual button click
	test.Click("#test-button")

	if resultDiv != nil {
		resultHTML := resultDiv.InnerHTML()
		resultText := resultDiv.TextContent()
		t.Logf("Result after button click - HTML: '%s'", resultHTML)
		t.Logf("Result after button click - Text: '%s'", resultText)

		if resultText == "HTMX HTTP request successful!" {
			t.Log("✅ HTMX HTTP requests working")
		} else if resultText == "HTMX loaded but ajax not available" {
			t.Log("❌ HTMX loaded but ajax functionality missing")
		} else if resultText == "HTMX not available" {
			t.Log("❌ HTMX not available")
		} else {
			t.Logf("❓ Unexpected result: %s", resultText)
		}
	}
}

func TestDebugHTMXFetch(t *testing.T) {
	// Test if the Fetch API is available (HTMX might depend on it)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/":
			w.Header().Set("Content-Type", "text/html")
			w.Write([]byte(`<!DOCTYPE html>
<html>
<head>
    <title>Fetch API Test</title>
</head>
<body>
    <div id="result">Testing...</div>
    
    <script>
        console.log('Testing Fetch API availability');
        
        var resultDiv = document.getElementById('result');
        
        if (typeof fetch !== 'undefined') {
            console.log('Fetch API is available');
            resultDiv.textContent = 'Fetch API available';
        } else {
            console.log('Fetch API is NOT available');
            resultDiv.textContent = 'Fetch API not available';
        }
        
        // Also test XMLHttpRequest
        if (typeof XMLHttpRequest !== 'undefined') {
            console.log('XMLHttpRequest is available');
            resultDiv.textContent += ' - XMLHttpRequest available';
        } else {
            console.log('XMLHttpRequest is NOT available');
            resultDiv.textContent += ' - XMLHttpRequest not available';
        }
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

	resultDiv := test.Document().QuerySelector("#result")
	if resultDiv != nil {
		resultText := resultDiv.TextContent()
		t.Logf("Fetch/XMLHttpRequest test result: '%s'", resultText)
	}
}
