package examples

import (
	"fmt"
	"testing"
	"time"

	domulator "github.com/dpemmons/DOMulator"
)

// ExampleAsyncJavaScriptTesting demonstrates how to test asynchronous JavaScript
// code using DOMulator's event loop control methods.
func ExampleAsyncJavaScriptTesting() {
	// Create a mock test environment
	t := &testing.T{}

	test := domulator.NewTest(t)
	defer test.Close()

	// Example 1: Testing setTimeout with DOM manipulation
	fmt.Println("=== Example 1: Testing setTimeout ===")

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
	fmt.Println("Initial status:", getElementText(test, "#status"))

	// Advance time to trigger the timeout
	test.AdvanceTime(1 * time.Second)

	// Now shows "Ready!"
	fmt.Println("After 1 second:", getElementText(test, "#status"))

	// Example 2: Testing microtasks (Promises)
	fmt.Println("\n=== Example 2: Testing Microtasks ===")

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

	fmt.Println("Before processing microtasks:", getElementText(test, "#result"))

	// Process all pending microtasks
	test.FlushMicrotasks()

	fmt.Println("After processing microtasks:", getElementText(test, "#result"))

	// Example 3: Testing complex async flows
	fmt.Println("\n=== Example 3: Complex Async Flow ===")

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

	fmt.Println("Initial:", getElementText(test, "#progress"))

	// Process immediate microtask
	test.FlushMicrotasks()
	fmt.Println("After microtasks:", getElementText(test, "#progress"))

	// Advance time for first timer
	test.AdvanceTime(500 * time.Millisecond)
	fmt.Println("After 500ms:", getElementText(test, "#progress"))

	// Advance time for nested timer
	test.AdvanceTime(100 * time.Millisecond)
	fmt.Println("After 600ms total:", getElementText(test, "#progress"))

	// Example 4: Testing event listeners added asynchronously
	fmt.Println("\n=== Example 4: Async Event Listeners ===")

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
	fmt.Println("Clicked before listener added:", getElementText(test, "#result"))

	// Advance time to add event listener
	test.AdvanceTime(200 * time.Millisecond)

	// Now button click should work
	test.Click("#btn")
	fmt.Println("Clicked after listener added:", getElementText(test, "#result"))

	fmt.Println("\n✅ All async testing examples completed successfully!")
}

// Helper function to get element text content
func getElementText(test *domulator.Test, selector string) string {
	element := test.Document().QuerySelector(selector)
	if element == nil {
		return "<element not found>"
	}
	return element.TextContent()
}

// This example can be run as a test:
// go test -run ExampleAsyncJavaScriptTesting ./examples/
func TestExampleAsyncJavaScriptTesting(t *testing.T) {
	ExampleAsyncJavaScriptTesting()
}
