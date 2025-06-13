package main

import (
	"testing"
	"time"

	domulator "github.com/dpemmons/DOMulator"
)

// TestAsyncJavaScriptTesting demonstrates how to test asynchronous JavaScript
// code using DOMulator's event loop control methods.
func TestAsyncJavaScriptTesting(t *testing.T) {
	test := domulator.NewTest(t)
	defer test.Close()

	// Example 1: Testing setTimeout with DOM manipulation
	t.Log("=== Example 1: Testing setTimeout ===")

	test.LoadHTML(`
		<html>
		<body>
			<div id="status">Loading...</div>
			<script>
				// This would be hard to test without event loop control
				setTimeout(() => {
					document.getElementById('status').textContent = 'Ready!';
					console.log('Status updated to Ready!');
				}, 1000);
			</script>
		</body>
		</html>
	`)

	// Initially shows "Loading..."
	t.Logf("Initial status: %s", getElementText(test, "#status"))

	// Advance time to trigger the timeout
	test.AdvanceTime(1 * time.Second)

	// Now shows "Ready!"
	t.Logf("After 1 second: %s", getElementText(test, "#status"))

	// Example 2: Testing microtasks (Promises)
	t.Log("=== Example 2: Testing Microtasks ===")

	test.LoadHTML(`
		<html>
		<body>
			<div id="result">Initial</div>
			<script>
				queueMicrotask(() => {
					document.getElementById('result').textContent = 'Microtask 1';
				});
				
				queueMicrotask(() => {
					document.getElementById('result').textContent = 'Microtask 2';
				});
				
				console.log('Microtasks queued');
			</script>
		</body>
		</html>
	`)

	t.Logf("Before processing microtasks: %s", getElementText(test, "#result"))

	// Process all pending microtasks
	test.FlushMicrotasks()

	t.Logf("After processing microtasks: %s", getElementText(test, "#result"))

	// Example 3: Testing complex async flows
	t.Log("=== Example 3: Complex Async Flow ===")

	test.LoadHTML(`
		<html>
		<body>
			<div id="progress">Step 0</div>
			<script>
				let step = 0;
				
				function updateProgress() {
					step++;
					document.getElementById('progress').textContent = 'Step ' + step;
				}
				
				// Immediate microtask
				queueMicrotask(() => {
					updateProgress();
					console.log('Completed step 1 (microtask)');
				});
				
				// Timer that schedules another timer
				setTimeout(() => {
					updateProgress();
					console.log('Completed step 2 (timer)');
					
					// Nested timer
					setTimeout(() => {
						updateProgress();
						console.log('Completed step 3 (nested timer)');
					}, 100);
				}, 500);
				
				console.log('Async operations scheduled');
			</script>
		</body>
		</html>
	`)

	t.Logf("Initial: %s", getElementText(test, "#progress"))

	// Process immediate microtask
	test.FlushMicrotasks()
	t.Logf("After microtasks: %s", getElementText(test, "#progress"))

	// Advance time for first timer
	test.AdvanceTime(500 * time.Millisecond)
	t.Logf("After 500ms: %s", getElementText(test, "#progress"))

	// Advance time for nested timer
	test.AdvanceTime(100 * time.Millisecond)
	t.Logf("After 600ms total: %s", getElementText(test, "#progress"))

	// Example 4: Testing event listeners added asynchronously
	t.Log("=== Example 4: Async Event Listeners ===")

	test.LoadHTML(`
		<html>
		<body>
			<button id="btn">Click me</button>
			<div id="result">Not clicked</div>
			<script>
				// Event listener is added after a delay
				setTimeout(() => {
					document.getElementById('btn').addEventListener('click', () => {
						document.getElementById('result').textContent = 'Button clicked!';
					});
					console.log('Event listener added');
				}, 200);
			</script>
		</body>
		</html>
	`)

	// Button click should not work initially
	test.Click("#btn")
	t.Logf("Clicked before listener added: %s", getElementText(test, "#result"))

	// Advance time to add event listener
	test.AdvanceTime(200 * time.Millisecond)

	// Now button click should work
	test.Click("#btn")
	t.Logf("Clicked after listener added: %s", getElementText(test, "#result"))

	t.Log("✅ All async testing examples completed successfully!")
}

// Helper function to get element text content
func getElementText(test *domulator.Test, selector string) string {
	element := test.Document().QuerySelector(selector)
	if element == nil {
		return "<element not found>"
	}
	return element.TextContent()
}
