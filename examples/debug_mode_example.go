package examples

import (
	"testing"

	domulator "github.com/dpemmons/DOMulator"
)

// Example showing how to use debug mode for verbose console output
func main() {
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
}

// Example showing the clean test output by default
func ExampleCleanOutput() {
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
}
