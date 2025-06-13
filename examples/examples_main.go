package main

import (
	"fmt"
	"testing"
	"time"

	domulator "github.com/dpemmons/DOMulator"
)

func main() {
	DebugModeExample()
	CleanOutputExample()
	ExampleAsyncJavaScriptTesting()
}

// Example showing how to use debug mode for verbose console output
func DebugModeExample() {
	println("Running DebugModeExample...")
	t := &testing.T{} // In real tests, this would be provided by the test framework
	test := domulator.NewTest(t)

	// By default, console.log output is suppressed for clean test output
	test.LoadHTML(`
		<button id="btn">Click me</button>
		<script>
			console.log('This will not appear in output by default');
			document.getElementById('btn').addEventListener('click', function() {
				console.log('Button clicked - this is also suppressed by default');
			});
		</script>
	`)

	// Enable debug mode to see verbose console output
	test.SetDebugMode(true)

	// Now console.log statements will be visible
	test.ExecuteScript(`
		console.log('This will appear because debug mode is enabled');
		console.log('Object details:', document.getElementById('btn'));
	`)

	test.Click("#btn")

	// Disable debug mode to return to clean output
	test.SetDebugMode(false)

	test.ExecuteScript(`
		console.log('This will be suppressed again');
	`)
	println("DebugModeExample completed.")
}

// Example showing the clean test output by default
func CleanOutputExample() {
	println("Running CleanOutputExample...")

	t := &testing.T{}
	test := domulator.NewTest(t)

	test.LoadHTML(`
		<form id="form">
			<input id="email" type="email" name="email">
			<button type="submit">Submit</button>
		</form>
		<script>
			// All of these console.log statements are suppressed by default
			console.log('Form loaded');
			
			document.getElementById('form').addEventListener('submit', function(e) {
				console.log('Form submitted with event:', e);
				console.log('Email value:', document.getElementById('email').value);
			});
		</script>
	`)

	// This produces clean output without verbose console dumps
	test.Type("#email", "test@example.com")
	test.Submit("#form")

	// Assertions work normally
	test.AssertElement("#email").HasValue("test@example.com")
	println("CleanOutputExample completed.")
}

// ExampleAsyncJavaScriptTesting demonstrates how to test asynchronous JavaScript
// code using DOMulator's event loop control methods.
func ExampleAsyncJavaScriptTesting() {
	println("Running ExampleAsyncJavaScriptTesting...")

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

	println("ExampleAsyncJavaScriptTesting completed.")
}

// Helper function to get element text content
func getElementText(test *domulator.Test, selector string) string {
	element := test.Document().QuerySelector(selector)
	if element == nil {
		return "<element not found>"
	}
	return element.TextContent()
}
