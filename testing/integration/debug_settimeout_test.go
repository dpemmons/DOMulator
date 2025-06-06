package integration

import (
	"testing"
	"time"

	domulator "github.com/dpemmons/DOMulator"
)

func TestDebugSetTimeoutNested(t *testing.T) {
	test := domulator.NewTest(t)
	test.SetDebugMode(true)

	test.LoadHTML(`
		<!DOCTYPE html>
		<html>
		<body>
			<div id="results"></div>
		</body>
		</html>
		<script>
		console.log("Starting nested setTimeout test...");
		
		let step = 1;
		document.getElementById('results').setAttribute('data-step', step.toString());
		
		setTimeout(() => {
			console.log("First timeout executing...");
			step = 2;
			document.getElementById('results').setAttribute('data-step', step.toString());
			
			setTimeout(() => {
				console.log("Second timeout executing...");
				step = 3;
				document.getElementById('results').setAttribute('data-step', step.toString());
				document.body.setAttribute('data-test', 'complete');
				console.log("Nested setTimeout test complete!");
			}, 50);
			
			console.log("First timeout finished, nested timeout scheduled");
		}, 10);
		
		console.log("Setup complete");
		</script>
	`)

	test.FlushMicrotasks()

	// Check initial step
	test.AssertElement("#results").HasAttribute("data-step", "1")

	// Advance time for first timeout
	test.AdvanceTime(20 * time.Millisecond)
	test.ProcessPendingTimers()
	test.FlushMicrotasks()

	// Check that first timeout executed
	test.AssertElement("#results").HasAttribute("data-step", "2")

	// Advance time for second timeout
	test.AdvanceTime(60 * time.Millisecond)
	test.ProcessPendingTimers()
	test.FlushMicrotasks()

	// Check that second timeout executed
	test.AssertElement("#results").HasAttribute("data-step", "3")
	test.AssertElement("body").HasAttribute("data-test", "complete")
}
