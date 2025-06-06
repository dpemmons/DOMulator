package integration

import (
	"testing"
	"time"

	domulator "github.com/dpemmons/DOMulator"
)

func TestResizeObserverBasicFunctionality(t *testing.T) {
	test := domulator.NewTest(t)
	test.SetDebugMode(true)

	test.LoadHTML(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>ResizeObserver Basic Test</title>
		</head>
		<body>
			<div id="observed-element">Resizable Element</div>
			<div id="results"></div>
		</body>
		</html>
		<script>
		console.log("Script starting...");
		
		// Track state
		var initialCallbackFired = false;
		var resizeCallbackFired = false;
		
		// Create ResizeObserver
		var observer = new ResizeObserver(function(entries) {
			console.log("ResizeObserver callback called with", entries.length, "entries");
			
			if (!initialCallbackFired) {
				initialCallbackFired = true;
				document.getElementById('results').setAttribute('data-initial', 'true');
			} else {
				resizeCallbackFired = true;
				document.getElementById('results').setAttribute('data-resized', 'true');
			}
			
			entries.forEach(function(entry) {
				console.log("Entry for:", entry.target.id);
				console.log("Width:", entry.contentRect.width);
				console.log("Height:", entry.contentRect.height);
			});
		});
		
		// Start observing
		var element = document.getElementById('observed-element');
		observer.observe(element);
		console.log("Started observing element");
		
		// Set up test completion
		setTimeout(function() {
			document.body.setAttribute('data-test-complete', 'true');
		}, 100);
		</script>
	`)

	// Process initial observation
	test.FlushMicrotasks()

	// Verify initial callback was fired
	test.AssertElement("#results").HasAttribute("data-initial", "true")
	t.Log("✓ Initial ResizeObserver callback fired")

	// Now trigger a resize using the test harness
	test.TriggerElementResize("#observed-element", domulator.ResizeOptions{
		Width:  400,
		Height: 300,
	})

	// Process the resize callback
	test.FlushMicrotasks()

	// Verify resize callback was fired
	test.AssertElement("#results").HasAttribute("data-resized", "true")
	t.Log("✓ ResizeObserver detected element resize")

	// Advance time for test completion
	test.AdvanceTime(150 * time.Millisecond)
	test.ProcessPendingTimers()
	test.FlushMicrotasks()

	// Verify test completed
	test.AssertElement("body").HasAttribute("data-test-complete", "true")
	t.Log("✓ Test completed successfully")
}

func TestResizeObserverMultipleElements(t *testing.T) {
	test := domulator.NewTest(t)

	test.LoadHTML(`
		<!DOCTYPE html>
		<html>
		<body>
			<div id="element1">Element 1</div>
			<div id="element2">Element 2</div>
			<div id="results" data-callbacks="0"></div>
		</body>
		</html>
		<script>
		var callbackCount = 0;
		var observedElements = [];
		
		var observer = new ResizeObserver(function(entries) {
			callbackCount++;
			entries.forEach(function(entry) {
				observedElements.push(entry.target.id);
			});
			document.getElementById('results').setAttribute('data-callbacks', callbackCount);
			document.getElementById('results').setAttribute('data-elements', observedElements.join(','));
		});
		
		// Observe multiple elements
		observer.observe(document.getElementById('element1'));
		observer.observe(document.getElementById('element2'));
		</script>
	`)

	// Process initial observations
	test.FlushMicrotasks()

	// Should have received 2 callbacks (one for each element)
	test.AssertElement("#results").HasAttribute("data-callbacks", "2")
	test.AssertElement("#results").Exists()

	// Trigger resize on element1
	test.TriggerElementResize("#element1", domulator.ResizeOptions{
		Width: 500, Height: 400,
	})
	test.FlushMicrotasks()

	// Should have received another callback
	test.AssertElement("#results").HasAttribute("data-callbacks", "3")
}

func TestResizeObserverUnobserve(t *testing.T) {
	test := domulator.NewTest(t)

	test.LoadHTML(`
		<!DOCTYPE html>
		<html>
		<body>
			<div id="observed">Observable</div>
			<div id="results" data-callbacks="0"></div>
		</body>
		</html>
		<script>
		var callbackCount = 0;
		
		var observer = new ResizeObserver(function(entries) {
			callbackCount++;
			document.getElementById('results').setAttribute('data-callbacks', callbackCount);
		});
		
		var element = document.getElementById('observed');
		observer.observe(element);
		
		// Set up unobserve after initial callback
		setTimeout(function() {
			observer.unobserve(element);
			document.getElementById('results').setAttribute('data-unobserved', 'true');
		}, 50);
		</script>
	`)

	// Process initial observation
	test.FlushMicrotasks()
	test.AssertElement("#results").HasAttribute("data-callbacks", "1")

	// Advance time to trigger unobserve
	test.AdvanceTime(100 * time.Millisecond)
	test.ProcessPendingTimers()
	test.FlushMicrotasks()

	test.AssertElement("#results").HasAttribute("data-unobserved", "true")

	// Try to trigger resize after unobserve - should not trigger callback
	test.TriggerElementResize("#observed", domulator.ResizeOptions{
		Width: 300, Height: 200,
	})
	test.FlushMicrotasks()

	// Callback count should still be 1
	test.AssertElement("#results").HasAttribute("data-callbacks", "1")
}
