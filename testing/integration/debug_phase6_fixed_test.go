package integration

import (
	"testing"
	"time"

	domulator "github.com/dpemmons/DOMulator"
)

func TestDebugDOMAPIPhase6ResizeObserverFixed(t *testing.T) {
	test := domulator.NewTest(t)
	test.SetDebugMode(true)

	test.LoadHTML(`
		<!DOCTYPE html>
		<html>
		<body>
			<div id="observed-element">Resizable Element</div>
		</body>
		</html>
		<script>
		console.log("Starting ResizeObserver test...");
		
		// Test ResizeObserver constructor
		let constructorCheck = (typeof ResizeObserver === 'function');
		console.log("ResizeObserver constructor available:", constructorCheck);
		
		// Test ResizeObserver creation and basic functionality
		let observerCreated = false;
		let observeMethodAvailable = false;
		let callbackTriggered = false;
		
		try {
			let observer = new ResizeObserver(function(entries) {
				console.log("ResizeObserver callback triggered with", entries.length, "entries");
				callbackTriggered = true;
				
				// Set results immediately in callback
				document.body.setAttribute('data-callback-triggered', 'true');
			});
			
			observerCreated = true;
			observeMethodAvailable = (typeof observer.observe === 'function');
			
			let element = document.getElementById('observed-element');
			observer.observe(element);
			
			console.log("Observer setup complete");
		} catch (e) {
			console.log("Error setting up ResizeObserver:", e.message);
		}
		
		// Set test results immediately
		document.body.setAttribute('data-constructor-check', constructorCheck.toString());
		document.body.setAttribute('data-observer-created', observerCreated.toString());
		document.body.setAttribute('data-observe-method', observeMethodAvailable.toString());
		document.body.setAttribute('data-resize-test', 'complete');
		
		console.log("ResizeObserver test complete");
		</script>
	`)

	test.FlushMicrotasks()

	// Wait a moment for any ResizeObserver callbacks
	test.AdvanceTime(50 * time.Millisecond)
	test.ProcessPendingTimers()
	test.FlushMicrotasks()

	// Check results
	body := test.Document().Body()

	constructorCheck := body.GetAttribute("data-constructor-check")
	observerCreated := body.GetAttribute("data-observer-created")
	observeMethod := body.GetAttribute("data-observe-method")
	callbackTriggered := body.GetAttribute("data-callback-triggered")

	t.Logf("Constructor available: %s", constructorCheck)
	t.Logf("Observer created: %s", observerCreated)
	t.Logf("Observe method available: %s", observeMethod)
	t.Logf("Callback triggered: %s", callbackTriggered)

	// Verify basic functionality
	test.AssertElement("body").HasAttribute("data-constructor-check", "true")
	test.AssertElement("body").HasAttribute("data-observer-created", "true")
	test.AssertElement("body").HasAttribute("data-observe-method", "true")
	test.AssertElement("body").HasAttribute("data-resize-test", "complete")

	t.Log("✅ ResizeObserver basic functionality test passed!")
}
