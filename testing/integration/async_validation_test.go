package integration

import (
	"testing"
	"time"

	domulator "github.com/dpemmons/DOMulator"
)

// TestAsyncEventValidation demonstrates how the event loop control
// solves the original problem of testing asynchronous JavaScript
func TestAsyncEventValidation(t *testing.T) {
	test := domulator.NewTest(t)
	defer test.Close()

	t.Log("🔄 Testing asynchronous JavaScript event handling...")

	test.LoadHTML(`
		<html>
		<body>
			<button id="async-btn">Click me</button>
			<div id="result">Initial</div>
			<div id="timeout-result">Waiting...</div>
			<div id="microtask-result">Not processed</div>
			
			<script>
				// Test 1: Event listener added asynchronously
				setTimeout(() => {
					document.getElementById('async-btn').addEventListener('click', () => {
						document.getElementById('result').textContent = 'Async event handled!';
						console.log('✅ Async event listener worked');
					});
					console.log('Event listener added after 100ms');
				}, 100);
				
				// Test 2: Timeout that updates DOM
				setTimeout(() => {
					document.getElementById('timeout-result').textContent = 'Timeout executed!';
					console.log('✅ Timeout executed');
				}, 200);
				
				// Test 3: Microtask that updates DOM
				queueMicrotask(() => {
					document.getElementById('microtask-result').textContent = 'Microtask processed!';
					console.log('✅ Microtask processed');
				});
				
				console.log('All async operations scheduled');
			</script>
		</body>
		</html>
	`)

	// Initially, click should not work because event listener isn't added yet
	test.Click("#async-btn")
	test.AssertElement("#result").HasText("Initial")
	t.Log("✅ Click before listener addition correctly has no effect")

	// Process microtasks immediately
	test.FlushMicrotasks()
	test.AssertElement("#microtask-result").HasText("Microtask processed!")
	t.Log("✅ Microtask processed immediately")

	// Advance time to add event listener
	test.AdvanceTime(100 * time.Millisecond)

	// Now click should work
	test.Click("#async-btn")
	test.AssertElement("#result").HasText("Async event handled!")
	t.Log("✅ Click after listener addition works correctly")

	// Advance time to trigger timeout
	test.AdvanceTime(100 * time.Millisecond) // Total 200ms
	test.AssertElement("#timeout-result").HasText("Timeout executed!")
	t.Log("✅ Timeout executed at correct time")

	t.Log("🎉 All asynchronous JavaScript events validated successfully!")
}

// TestComplexAsyncValidation tests nested timers and microtasks
func TestComplexAsyncValidation(t *testing.T) {
	test := domulator.NewTest(t)
	defer test.Close()

	t.Log("🔄 Testing complex async flow with nested operations...")

	test.LoadHTML(`
		<html>
		<body>
			<div id="sequence">Step 0</div>
			
			<script>
				let step = 0;
				
				function updateStep(description) {
					step++;
					document.getElementById('sequence').textContent = 'Step ' + step + ': ' + description;
					console.log('Step', step, ':', description);
				}
				
				// Immediate microtask
				queueMicrotask(() => {
					updateStep('Immediate microtask');
				});
				
				// Timer that creates nested operations
				setTimeout(() => {
					updateStep('First timer');
					
					// Nested microtask
					queueMicrotask(() => {
						updateStep('Nested microtask');
					});
					
					// Nested timer
					setTimeout(() => {
						updateStep('Nested timer');
						
						// Double-nested microtask
						queueMicrotask(() => {
							updateStep('Double-nested microtask');
						});
					}, 50);
				}, 100);
				
				console.log('Complex async flow scheduled');
			</script>
		</body>
		</html>
	`)

	// Check initial state
	test.AssertElement("#sequence").HasText("Step 0")

	// Process immediate microtask
	test.FlushMicrotasks()
	test.AssertElement("#sequence").HasText("Step 1: Immediate microtask")
	t.Log("✅ Immediate microtask processed")

	// Advance to first timer (this also processes microtasks queued by the timer)
	test.AdvanceTime(100 * time.Millisecond)
	test.AssertElement("#sequence").HasText("Step 3: Nested microtask")
	t.Log("✅ First timer executed and nested microtask processed")

	// Advance to nested timer (this also processes its microtask)
	test.AdvanceTime(50 * time.Millisecond)
	test.AssertElement("#sequence").HasText("Step 5: Double-nested microtask")
	t.Log("✅ Nested timer executed and double-nested microtask processed")

	t.Log("🎉 Complex async flow completed successfully!")
}

// TestEventHandlingWithAsyncSetup tests the original problem scenario
func TestEventHandlingWithAsyncSetup(t *testing.T) {
	test := domulator.NewTest(t)
	defer test.Close()

	t.Log("🔄 Testing event handling with async setup (original problem scenario)...")

	test.LoadHTML(`
		<html>
		<body>
			<button id="btn1">Button 1</button>
			<button id="btn2">Button 2</button>
			<div id="events-log">No events</div>
			
			<script>
				let eventCount = 0;
				
				function logEvent(eventType, buttonId) {
					eventCount++;
					document.getElementById('events-log').textContent = 
						'Event ' + eventCount + ': ' + eventType + ' on ' + buttonId;
					console.log('Event logged:', eventType, 'on', buttonId);
				}
				
				// Setup event listeners asynchronously (simulating dynamic loading)
				setTimeout(() => {
					document.getElementById('btn1').addEventListener('click', () => {
						logEvent('click', 'btn1');
					});
					
					document.getElementById('btn1').addEventListener('mousedown', () => {
						logEvent('mousedown', 'btn1');
					});
					
					console.log('Listeners added to btn1');
				}, 50);
				
				setTimeout(() => {
					document.getElementById('btn2').addEventListener('click', () => {
						logEvent('click', 'btn2');
					});
					
					document.getElementById('btn2').addEventListener('mouseup', () => {
						logEvent('mouseup', 'btn2');
					});
					
					console.log('Listeners added to btn2');
				}, 100);
				
				console.log('Async event listener setup scheduled');
			</script>
		</body>
		</html>
	`)

	// Initially no events should be logged when clicking
	test.Click("#btn1")
	test.Click("#btn2")
	test.AssertElement("#events-log").HasText("No events")
	t.Log("✅ Events ignored before listeners are set up")

	// Advance time to add btn1 listeners
	test.AdvanceTime(50 * time.Millisecond)

	// Now btn1 should work but btn2 shouldn't
	test.Click("#btn1")
	test.AssertElement("#events-log").HasText("Event 1: click on btn1")
	t.Log("✅ btn1 click works after its listener is added")

	test.Click("#btn2")
	test.AssertElement("#events-log").HasText("Event 1: click on btn1") // Should be unchanged
	t.Log("✅ btn2 click still ignored")

	// Advance time to add btn2 listeners
	test.AdvanceTime(50 * time.Millisecond) // Total 100ms

	// Now both buttons should work
	test.MouseDown("#btn1")
	test.AssertElement("#events-log").HasText("Event 2: mousedown on btn1")

	test.MouseUp("#btn2")
	test.AssertElement("#events-log").HasText("Event 3: mouseup on btn2")

	test.Click("#btn2")
	test.AssertElement("#events-log").HasText("Event 4: click on btn2")

	t.Log("✅ Both buttons work correctly after all listeners are added")

	t.Log("🎉 Event handling with async setup validated successfully!")
	t.Log("    This demonstrates that DOMulator now properly tests the full event flow,")
	t.Log("    including JavaScript event listeners receiving and processing events!")
}
