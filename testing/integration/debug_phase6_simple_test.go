package integration

import (
	"testing"
	"time"

	domulator "github.com/dpemmons/DOMulator"
)

func TestDebugDOMAPIPhase6Simple(t *testing.T) {
	test := domulator.NewTest(t)
	test.SetDebugMode(true)

	test.LoadHTML(`
		<!DOCTYPE html>
		<html>
		<body>
			<div id="observed-element">Element</div>
		</body>
		</html>
		<script>
		console.log("Starting simple ResizeObserver test...");
		
		// Test that ResizeObserver exists and can be created
		let observer = new ResizeObserver(function(entries) {
			console.log("ResizeObserver callback triggered");
		});
		
		let element = document.getElementById('observed-element');
		observer.observe(element);
		
		console.log("ResizeObserver setup complete");
		
		setTimeout(() => {
			console.log("First timeout executing...");
			document.body.setAttribute('data-step', '1');
			
			setTimeout(() => {
				console.log("Second timeout executing...");
				document.body.setAttribute('data-step', '2');
				document.body.setAttribute('data-test', 'complete');
				console.log("Test complete!");
			}, 50);
		}, 10);
		
		console.log("Timeouts scheduled");
		</script>
	`)

	test.FlushMicrotasks()

	// Advance time for first timeout
	test.AdvanceTime(20 * time.Millisecond)
	test.ProcessPendingTimers()
	test.FlushMicrotasks()

	// Check that first timeout executed
	body := test.Document().Body()
	if body.HasAttribute("data-step") {
		t.Logf("Found data-step: %s", body.GetAttribute("data-step"))
	} else {
		t.Logf("No data-step attribute found after first timeout")
	}

	// Advance time for second timeout
	test.AdvanceTime(60 * time.Millisecond)
	test.ProcessPendingTimers()
	test.FlushMicrotasks()

	// Check final state
	if body.HasAttribute("data-test") {
		t.Logf("Found data-test: %s", body.GetAttribute("data-test"))
	} else {
		t.Logf("No data-test attribute found after second timeout")
	}

	if body.HasAttribute("data-step") {
		t.Logf("Final data-step: %s", body.GetAttribute("data-step"))
	}

	test.AssertElement("body").HasAttribute("data-test", "complete")
}
