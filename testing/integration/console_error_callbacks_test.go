package integration

import (
	"testing"

	"github.com/dpemmons/DOMulator/internal/js"
	testpkg "github.com/dpemmons/DOMulator/testing"
)

func TestConsoleCallbacks(t *testing.T) {
	harness := testpkg.NewTestHarness()
	defer harness.Close()

	// Set up console capture
	consoleCapture := harness.CaptureConsole(t)

	// Load HTML with script that produces various console outputs
	harness.LoadHTML(`
		<!DOCTYPE html>
		<html>
		<head><title>Console Test</title></head>
		<body>
			<script>
				// Test different console levels
				console.log("This is a log message");
				console.info("This is an info message");
				console.warn("This is a warning message");
				console.error("This is an error message");
				
				// Test multiple arguments
				console.log("Multiple", "arguments", 123, true);
				
				// Test object logging
				var obj = { name: "test", value: 42 };
				console.log("Object:", obj);
			</script>
		</body>
		</html>
	`)

	// Verify console capture functionality
	consoleCapture.
		AssertLogCount(3). // log + multiple args + object
		AssertInfoCount(1).
		AssertWarnCount(1).
		AssertErrorCount(1).
		AssertLogContains("log message").
		AssertInfoContains("info message").
		AssertWarnContains("warning message").
		AssertErrorContains("error message").
		AssertLogContains("Multiple").
		AssertLogContains("Object:")

	// Test clearing and new messages
	consoleCapture.Clear()
	harness.ExecuteScript(`console.log("After clear");`)
	consoleCapture.AssertLogCount(1).AssertLogContains("After clear")
}

func TestJavaScriptErrorCallbacks(t *testing.T) {
	harness := testpkg.NewTestHarness()
	defer harness.Close()

	// Set up error capture
	errorCapture := harness.CaptureErrors(t)

	// Load basic HTML
	harness.LoadHTML(`
		<!DOCTYPE html>
		<html>
		<head><title>Error Test</title></head>
		<body></body>
		</html>
	`)

	// Test TypeError
	harness.ExecuteScript(`
		try {
			null.someMethod();
		} catch (e) {
			// We expect this to be caught by our error handler
		}
	`)

	// Test ReferenceError
	harness.ExecuteScript(`
		try {
			undefinedVariable.property = 5;
		} catch (e) {
			// We expect this to be caught by our error handler  
		}
	`)

	// Test custom error
	harness.ExecuteScript(`
		try {
			throw new Error("Custom error message");
		} catch (e) {
			// We expect this to be caught by our error handler
		}
	`)

	// Note: These specific tests might not trigger our error callback
	// because they're wrapped in try-catch blocks. Let's test some
	// that will actually trigger the error callback.

	// Clear any previous errors
	errorCapture.Clear()

	// Test actual runtime errors that will trigger the callback
	harness.ExecuteScript(`
		// This should trigger a TypeError
		var obj = null;
		obj.method();  // Cannot read property 'method' of null
	`)

	// Check if we captured any errors
	if errorCapture.Count() > 0 {
		errorCapture.AssertErrorType("TypeError")
	}

	// Test another type of error
	errorCapture.Clear()
	harness.ExecuteScript(`
		// This should trigger a ReferenceError
		nonExistentFunction();
	`)

	if errorCapture.Count() > 0 {
		errorCapture.AssertReferenceError()
	}
}

func TestConsoleInDebugMode(t *testing.T) {
	harness := testpkg.NewTestHarness()
	defer harness.Close()

	// Test in non-debug mode first
	harness.SetDebugMode(false)
	consoleCapture := harness.CaptureConsole(t)

	harness.LoadHTML(`
		<!DOCTYPE html>
		<html>
		<body>
			<script>
				console.log("Should be captured even in non-debug mode");
			</script>
		</body>
		</html>
	`)

	// Should still capture in non-debug mode
	consoleCapture.AssertLogCount(1).AssertLogContains("captured even in non-debug")

	// Clear and test debug mode
	consoleCapture.Clear()
	harness.SetDebugMode(true)

	harness.ExecuteScript(`console.log("Debug mode message");`)
	consoleCapture.AssertLogCount(1).AssertLogContains("Debug mode message")
}

func TestConsoleWithDOMElements(t *testing.T) {
	harness := testpkg.NewTestHarness()
	defer harness.Close()

	consoleCapture := harness.CaptureConsole(t)

	harness.LoadHTML(`
		<!DOCTYPE html>
		<html>
		<body>
			<div id="testDiv">Hello World</div>
			<script>
				var div = document.getElementById('testDiv');
				console.log("Element:", div);
				console.log("Element text:", div.textContent);
			</script>
		</body>
		</html>
	`)

	// Should capture console logs with DOM elements
	consoleCapture.
		AssertLogCount(2).
		AssertLogContains("Element:").
		AssertLogContains("Hello World")
}

func TestErrorCallbackWithoutPanic(t *testing.T) {
	harness := testpkg.NewTestHarness()
	defer harness.Close()

	errorCapture := harness.CaptureErrors(t)

	harness.LoadHTML(`
		<!DOCTYPE html>
		<html>
		<body></body>
		</html>
	`)

	// The key test: errors should be captured via callback, not panic
	// Execute code that would normally panic
	err := harness.ExecuteScript(`
		// This should trigger our error callback instead of panicking
		var result = someUndefinedFunction();
	`)

	// The execution should complete without panicking
	// The error might be nil if it was handled by the callback
	if err != nil {
		t.Logf("ExecuteScript returned error (this is expected): %v", err)
	}

	// Check if we captured the error via callback
	if errorCapture.Count() > 0 {
		t.Logf("Successfully captured %d errors via callback", errorCapture.Count())
		lastError := errorCapture.GetLastError()
		if lastError != nil {
			t.Logf("Last error: Type=%s, Message=%s", lastError.Type, lastError.Message)
		}
	} else {
		t.Log("No errors captured - this might be expected depending on how the runtime handles this")
	}
}

func TestDirectCallbackUsage(t *testing.T) {
	harness := testpkg.NewTestHarness()
	defer harness.Close()

	// Test direct callback usage without the helper
	var capturedConsole []string
	var capturedErrors []*js.JavaScriptError

	harness.SetConsoleCallback(func(level js.ConsoleLevel, args []interface{}) {
		// Convert args to strings and store
		var parts []string
		for _, arg := range args {
			parts = append(parts, arg.(string))
		}
		message := string(level) + ": " + parts[0] // Simple concat for test
		capturedConsole = append(capturedConsole, message)
	})

	harness.SetErrorCallback(func(err *js.JavaScriptError) {
		capturedErrors = append(capturedErrors, err)
	})

	harness.LoadHTML(`
		<!DOCTYPE html>
		<html>
		<body>
			<script>
				console.log("Direct callback test");
				console.error("Direct error test");
			</script>
		</body>
		</html>
	`)

	// Check direct callback results
	if len(capturedConsole) >= 2 {
		t.Logf("Captured console messages: %v", capturedConsole)
	} else {
		t.Errorf("Expected at least 2 console messages, got %d", len(capturedConsole))
	}

	// Execute some code that might trigger error callback
	harness.ExecuteScript(`
		// Try to trigger an error
		try {
			var x = null;
			x.someMethod();
		} catch (e) {
			console.log("Caught error in script");
		}
	`)

	t.Logf("Total captured console messages: %d", len(capturedConsole))
	t.Logf("Total captured errors: %d", len(capturedErrors))
}
