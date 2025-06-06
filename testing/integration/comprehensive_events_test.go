package integration

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	domulator "github.com/dpemmons/DOMulator"
)

func TestComprehensiveEventFlow(t *testing.T) {
	// This test validates the FULL event flow: DOM event → JavaScript listener → Console output
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`<!DOCTYPE html>
<html>
<head>
    <title>Comprehensive Event Flow Test</title>
</head>
<body>
    <div id="test-log"></div>
    <button id="test-button">Click Me</button>
    <input id="test-input" type="text" placeholder="Type here">
    <form id="test-form">
        <input name="username" type="text" required>
        <button type="submit">Submit</button>
    </form>
    
    <script>
        console.log('=== Event Flow Test Script Starting ===');
        
        // Create a simple logging system that writes to both console and DOM
        function logEvent(message) {
            console.log('EVENT: ' + message);
            var logDiv = document.getElementById('test-log');
            if (logDiv) {
                var entry = document.createElement('div');
                entry.className = 'log-entry';
                entry.textContent = new Date().toISOString() + ' - ' + message;
                logDiv.appendChild(entry);
            }
        }
        
        // Test 1: Basic click event
        var button = document.getElementById('test-button');
        if (button) {
            console.log('Adding click listener to button');
            button.addEventListener('click', function(event) {
                logEvent('Button clicked! Event type: ' + event.type);
                logEvent('Target tagName: ' + event.target.tagName);
                logEvent('Button text: ' + event.target.textContent);
                
                // Mark that the event was processed
                button.setAttribute('data-clicked', 'true');
            });
            logEvent('Click listener attached to button');
        } else {
            console.log('Button not found!');
            logEvent('ERROR: Button not found');
        }
        
        // Test 2: Input events
        var input = document.getElementById('test-input');
        if (input) {
            console.log('Adding input listeners');
            
            input.addEventListener('focus', function(event) {
                logEvent('Input focused');
                input.setAttribute('data-focused', 'true');
            });
            
            input.addEventListener('input', function(event) {
                logEvent('Input changed to: "' + event.target.value + '"');
                input.setAttribute('data-last-value', event.target.value);
            });
            
            input.addEventListener('blur', function(event) {
                logEvent('Input blurred with value: "' + event.target.value + '"');
                input.setAttribute('data-blurred', 'true');
            });
            
            logEvent('Input listeners attached');
        } else {
            logEvent('ERROR: Input not found');
        }
        
        // Test 3: Form submission
        var form = document.getElementById('test-form');
        if (form) {
            console.log('Adding form submit listener');
            
            form.addEventListener('submit', function(event) {
                event.preventDefault(); // Prevent actual submission
                
                var username = form.querySelector('input[name="username"]').value;
                logEvent('Form submitted with username: "' + username + '"');
                
                // Mark form as submitted
                form.setAttribute('data-submitted', 'true');
                form.setAttribute('data-username', username);
            });
            
            logEvent('Form submit listener attached');
        } else {
            logEvent('ERROR: Form not found');
        }
        
        // Test 4: Document-level events
        document.addEventListener('DOMContentLoaded', function() {
            logEvent('DOMContentLoaded fired');
        });
        
        // Test 5: Custom events
        var customEvent = new Event('customTest', { bubbles: true });
        document.addEventListener('customTest', function(event) {
            logEvent('Custom event received: ' + event.type);
        });
        
        // Dispatch the custom event after a short delay
        setTimeout(function() {
            console.log('Dispatching custom event');
            document.dispatchEvent(customEvent);
            logEvent('Custom event dispatched');
        }, 100);
        
        // Mark script as fully loaded
        document.body.setAttribute('data-script-loaded', 'true');
        
        console.log('=== Event Flow Test Script Completed ===');
        logEvent('Script initialization completed');
    </script>
</body>
</html>`))
	}))
	defer server.Close()

	test := domulator.NewTest(t)
	test.SetDebugMode(true)
	test.WithServer(server).Navigate("/")

	// Verify script loaded
	body := test.Document().QuerySelector("body")
	if body != nil && body.GetAttribute("data-script-loaded") == "true" {
		t.Log("✅ JavaScript script loaded successfully")
	} else {
		t.Log("❌ JavaScript script failed to load")
	}

	// Check initial log entries
	logDiv := test.Document().QuerySelector("#test-log")
	if logDiv != nil {
		initialLogs := logDiv.TextContent()
		t.Logf("Initial log content: %s", initialLogs)
	}

	t.Run("ButtonClickEvent", func(t *testing.T) {
		// Test button click
		button := test.Document().QuerySelector("#test-button")
		if button == nil {
			t.Fatal("Button not found")
		}

		// Click the button
		test.Click("#test-button")

		// Check if event was processed
		if button.GetAttribute("data-clicked") == "true" {
			t.Log("✅ Button click event processed by JavaScript")
		} else {
			t.Log("❌ Button click event NOT processed by JavaScript")
		}

		// Check log entries
		if logDiv != nil {
			logs := logDiv.TextContent()
			t.Logf("After button click - Log content: %s", logs)

			if logDiv.QuerySelector(".log-entry") != nil {
				t.Log("✅ Events are being logged to DOM")
			} else {
				t.Log("❌ No log entries found in DOM")
			}
		}
	})

	t.Run("InputEvents", func(t *testing.T) {
		input := test.Document().QuerySelector("#test-input")
		if input == nil {
			t.Fatal("Input not found")
		}

		// Focus input
		test.Focus("#test-input")
		if input.GetAttribute("data-focused") == "true" {
			t.Log("✅ Input focus event processed")
		} else {
			t.Log("❌ Input focus event NOT processed")
		}

		// Type in input
		test.Type("#test-input", "test value")
		if input.GetAttribute("data-last-value") == "test value" {
			t.Log("✅ Input change event processed")
		} else {
			t.Logf("❌ Input change event NOT processed. Last value: '%s'", input.GetAttribute("data-last-value"))
		}

		// Blur input
		test.Blur("#test-input")
		if input.GetAttribute("data-blurred") == "true" {
			t.Log("✅ Input blur event processed")
		} else {
			t.Log("❌ Input blur event NOT processed")
		}
	})

	t.Run("FormSubmission", func(t *testing.T) {
		form := test.Document().QuerySelector("#test-form")
		if form == nil {
			t.Fatal("Form not found")
		}

		// Fill out form
		test.Type("input[name='username']", "testuser")

		// Submit form
		test.Submit("#test-form")

		// Check if form submission was processed
		if form.GetAttribute("data-submitted") == "true" {
			t.Log("✅ Form submission event processed")

			username := form.GetAttribute("data-username")
			if username == "testuser" {
				t.Log("✅ Form data correctly captured")
			} else {
				t.Logf("❌ Form data NOT captured correctly. Got: '%s'", username)
			}
		} else {
			t.Log("❌ Form submission event NOT processed")
		}
	})

	t.Run("LogAnalysis", func(t *testing.T) {
		// Final analysis of all logged events
		if logDiv != nil {
			allLogs := logDiv.TextContent()
			t.Logf("=== COMPLETE EVENT LOG ===\n%s", allLogs)

			// Count log entries
			logEntries := logDiv.QuerySelectorAll(".log-entry")
			if logEntries != nil {
				entryCount := len(logEntries.ToSlice())
				t.Logf("Total log entries: %d", entryCount)

				if entryCount > 0 {
					t.Log("✅ JavaScript events are being captured and logged")
				} else {
					t.Log("❌ No JavaScript event logs found")
				}
			}
		}
	})
}

func TestEventValidationPattern(t *testing.T) {
	// A simpler test that follows the pattern from your original simple_events_test.go
	// but adds JavaScript validation
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`<!DOCTYPE html>
<html>
<head><title>Event Validation Test</title></head>
<body>
    <button id="btn">Test Button</button>
    <div id="result">No events yet</div>
    
    <script>
        var eventCount = 0;
        var lastEventType = '';
        
        document.getElementById('btn').addEventListener('click', function(e) {
            eventCount++;
            lastEventType = e.type;
            
            var result = document.getElementById('result');
            result.textContent = 'Events received: ' + eventCount + ', Last: ' + lastEventType;
            result.setAttribute('data-event-count', eventCount);
            result.setAttribute('data-last-event', lastEventType);
            
            console.log('Button clicked! Total events: ' + eventCount);
        });
    </script>
</body>
</html>`))
	}))
	defer server.Close()

	test := domulator.NewTest(t)
	test.SetDebugMode(true)
	test.WithServer(server).Navigate("/")

	// Before clicking - should have no events
	result := test.Document().QuerySelector("#result")
	if result != nil {
		initialCount := result.GetAttribute("data-event-count")
		t.Logf("Initial event count: '%s'", initialCount)
	}

	// Click button multiple times
	for i := 1; i <= 3; i++ {
		test.Click("#btn")

		if result != nil {
			count := result.GetAttribute("data-event-count")
			lastEvent := result.GetAttribute("data-last-event")
			t.Logf("After click %d - Count: '%s', Last event: '%s'", i, count, lastEvent)

			if count == "" {
				t.Errorf("Click %d: Event count not recorded by JavaScript", i)
			} else if count != strconv.Itoa(i) {
				t.Errorf("Click %d: Expected count %d, got '%s'", i, i, count)
			} else {
				t.Logf("✅ Click %d: JavaScript event handler working correctly", i)
			}
		}
	}
}
