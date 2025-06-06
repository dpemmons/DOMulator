package integration

import (
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	domulator "github.com/dpemmons/DOMulator"
)

func TestDebugScriptLoading(t *testing.T) {
	// Test whether external script loading is working at all
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/":
			w.Header().Set("Content-Type", "text/html")
			w.Write([]byte(`<!DOCTYPE html>
<html>
<head>
    <title>Script Loading Test</title>
</head>
<body>
    <div id="result">Before scripts</div>
    
    <!-- Try to load the HTMX script -->
    <script src="/htmx.js"></script>
    
    <!-- Inline script that should run after HTMX -->
    <script>
        console.log('Inline script starting');
        
        var resultDiv = document.getElementById('result');
        if (resultDiv) {
            resultDiv.textContent = 'Inline script executed';
            console.log('Updated result div');
        }
        
        if (typeof htmx !== 'undefined') {
            console.log('HTMX is available');
            resultDiv.textContent = 'HTMX loaded successfully';
        } else {
            console.log('HTMX is NOT available');
            resultDiv.textContent = 'HTMX failed to load';
        }
        
        console.log('Inline script completed');
    </script>
</body>
</html>`))

		case "/htmx.js":
			// Serve the real HTMX file
			htmxPath := filepath.Join("testdata", "htmx", "htmx.js")
			t.Logf("Serving HTMX from: %s", htmxPath)
			http.ServeFile(w, r, htmxPath)

		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	test := domulator.NewTest(t)
	test.SetDebugMode(true)
	test.WithServer(server).Navigate("/")

	// Check what the result div contains
	resultDiv := test.Document().QuerySelector("#result")
	if resultDiv != nil {
		resultText := resultDiv.TextContent()
		t.Logf("Result div text: '%s'", resultText)

		switch resultText {
		case "Before scripts":
			t.Log("❌ No scripts executed at all")
		case "Inline script executed":
			t.Log("✅ Inline script executed, but HTMX check didn't run")
		case "HTMX loaded successfully":
			t.Log("✅ HTMX loaded and inline script executed")
		case "HTMX failed to load":
			t.Log("✅ Inline script executed, HTMX not available (expected)")
		default:
			t.Logf("❓ Unexpected result: %s", resultText)
		}
	} else {
		t.Log("❌ Result div not found")
	}
}

func TestDebugScriptLoadingSimple(t *testing.T) {
	// Test simpler case with just inline script (no external script)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`<!DOCTYPE html>
<html>
<head>
    <title>Simple Script Test</title>
</head>
<body>
    <div id="result">Before script</div>
    
    <script>
        console.log('Simple script starting');
        var resultDiv = document.getElementById('result');
        if (resultDiv) {
            resultDiv.textContent = 'Simple script executed';
            console.log('Simple script completed');
        }
    </script>
</body>
</html>`))
	}))
	defer server.Close()

	test := domulator.NewTest(t)
	test.SetDebugMode(true)
	test.WithServer(server).Navigate("/")

	resultDiv := test.Document().QuerySelector("#result")
	if resultDiv != nil {
		resultText := resultDiv.TextContent()
		t.Logf("Simple test result: '%s'", resultText)

		if resultText == "Simple script executed" {
			t.Log("✅ Simple inline script works")
		} else {
			t.Log("❌ Simple inline script failed")
		}
	}
}
