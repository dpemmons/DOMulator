package integration

import (
	"testing"
	"time"

	domulator "github.com/dpemmons/DOMulator"
)

func TestDebugDOMAPIPhase6ResizeObserver(t *testing.T) {
	test := domulator.NewTest(t)
	test.SetDebugMode(true) // Enable debug output

	test.LoadHTML(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>Debug Phase 6 ResizeObserver Test</title>
		</head>
		<body>
			<div id="observed-element" style="width: 100px; height: 100px; background: blue;">
				Resizable Element
			</div>
			<div id="results"></div>
		</body>
		</html>
		<script>
		console.log("Script starting...");
		let observerCallbackCount = 0;
		let observedElements = [];
		
		// Test ResizeObserver creation
		console.log("ResizeObserver constructor:", typeof ResizeObserver);
		
		const observer = new ResizeObserver(function(entries) {
			observerCallbackCount++;
			console.log("ResizeObserver callback called, entries:", entries.length);
			
			entries.forEach(function(entry) {
				observedElements.push(entry.target.id);
				console.log("Observed element:", entry.target.id);
				console.log("Content rect:", entry.contentRect);
				
				// Store results in DOM for verification
				const results = document.getElementById('results');
				results.setAttribute('data-callback-count', observerCallbackCount.toString());
				results.setAttribute('data-observed-elements', observedElements.join(','));
			});
		});
		
		console.log("ResizeObserver created");
		
		// Test observer methods
		const element = document.getElementById('observed-element');
		
		try {
			observer.observe(element);
			console.log("observe() completed");
		} catch (e) {
			console.log("observe() error:", e.message);
		}
		
		console.log("Setting up timeouts...");
		
		// Simulate a resize by changing the element's style
		setTimeout(() => {
			console.log("First timeout executing...");
			
			try {
				element.style.width = '200px';
				element.style.height = '150px';
				console.log("Element resized");
				
				// Verify observer functionality immediately
				console.log("Building results array...");
				let results = [];
				results.push("constructor: " + (typeof ResizeObserver === 'function'));
				results.push("creation: " + (observer !== null));
				results.push("methods: " + (typeof observer.observe === 'function'));
				results.push("callbacks: " + (observerCallbackCount > 0));
				console.log("Results built:", results.join(','));
				
				console.log("Setting final attributes...");
				document.body.setAttribute('data-resize-results', results.join(','));
				document.body.setAttribute('data-resize-test', 'complete');
				console.log("Test complete!");
			} catch (e) {
				console.log("Error in timeout:", e.message);
				// Set error state
				document.body.setAttribute('data-resize-test', 'error');
				document.body.setAttribute('data-resize-error', e.message);
			}
		}, 10);
		
		console.log("Script finished setup");
		</script>
	`)

	// Process the initial script execution
	test.FlushMicrotasks()

	// Advance time to trigger first timeout (10ms)
	test.AdvanceTime(20 * time.Millisecond)
	test.ProcessPendingTimers()
	test.FlushMicrotasks()

	// Advance time to trigger second timeout (50ms more)
	test.AdvanceTime(60 * time.Millisecond)
	test.ProcessPendingTimers()
	test.FlushMicrotasks()

	// Process any additional timers that may have been scheduled
	test.AdvanceTime(100 * time.Millisecond)
	test.ProcessPendingTimers()
	test.FlushMicrotasks()

	// Check what we have
	body := test.Document().Body()
	if body != nil {
		if body.HasAttribute("data-resize-test") {
			t.Logf("Found data-resize-test: %s", body.GetAttribute("data-resize-test"))
		} else {
			t.Logf("No data-resize-test attribute found")
		}

		if body.HasAttribute("data-resize-results") {
			t.Logf("Found data-resize-results: %s", body.GetAttribute("data-resize-results"))
		} else {
			t.Logf("No data-resize-results attribute found")
		}
	} else {
		t.Logf("No body element found")
	}

	test.AssertElement("body").HasAttribute("data-resize-test", "complete")
}
