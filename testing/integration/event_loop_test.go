package integration

import (
	"testing"
	"time"

	domulator "github.com/dpemmons/DOMulator"
	testpkg "github.com/dpemmons/DOMulator/testing"
)

func TestBasicTimeoutAdvancement(t *testing.T) {
	test := domulator.NewTest(t)
	defer test.Close()

	test.LoadHTML(`
		<html>
		<body>
			<div id="status">Initial</div>
			<script>
				setTimeout(() => {
					document.getElementById('status').textContent = 'Updated';
				}, 1000);
			</script>
		</body>
		</html>
	`)

	// Initially should show "Initial"
	test.AssertElement("#status").HasText("Initial")

	// Advance time by 1 second to trigger the timeout
	test.AdvanceTime(1 * time.Second)

	// Should now show "Updated"
	test.AssertElement("#status").HasText("Updated")
}

func TestMultipleTimersWithDifferentDelays(t *testing.T) {
	test := domulator.NewTest(t)
	defer test.Close()

	test.LoadHTML(`
		<html>
		<body>
			<div id="status">0</div>
			<script>
				let count = 0;
				
				setTimeout(() => {
					count++;
					document.getElementById('status').textContent = count.toString();
				}, 100);
				
				setTimeout(() => {
					count++;
					document.getElementById('status').textContent = count.toString();
				}, 200);
				
				setTimeout(() => {
					count++;
					document.getElementById('status').textContent = count.toString();
				}, 300);
			</script>
		</body>
		</html>
	`)

	test.AssertElement("#status").HasText("0")

	// Advance 100ms - first timer should fire
	test.AdvanceTime(100 * time.Millisecond)
	test.AssertElement("#status").HasText("1")

	// Advance another 100ms (200ms total) - second timer should fire
	test.AdvanceTime(100 * time.Millisecond)
	test.AssertElement("#status").HasText("2")

	// Advance another 100ms (300ms total) - third timer should fire
	test.AdvanceTime(100 * time.Millisecond)
	test.AssertElement("#status").HasText("3")
}

func TestSetInterval(t *testing.T) {
	test := domulator.NewTest(t)
	defer test.Close()

	test.LoadHTML(`
		<html>
		<body>
			<div id="counter">0</div>
			<script>
				let count = 0;
				setInterval(() => {
					count++;
					document.getElementById('counter').textContent = count.toString();
				}, 500);
			</script>
		</body>
		</html>
	`)

	test.AssertElement("#counter").HasText("0")

	// First interval
	test.AdvanceTime(500 * time.Millisecond)
	test.AssertElement("#counter").HasText("1")

	// Second interval
	test.AdvanceTime(500 * time.Millisecond)
	test.AssertElement("#counter").HasText("2")

	// Third interval
	test.AdvanceTime(500 * time.Millisecond)
	test.AssertElement("#counter").HasText("3")
}

func TestMicrotaskProcessing(t *testing.T) {
	test := domulator.NewTest(t)
	defer test.Close()

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
			</script>
		</body>
		</html>
	`)

	// Initially should still show "Initial" since microtasks haven't processed
	test.AssertElement("#result").HasText("Initial")

	// Flush microtasks - should run both microtasks in order
	test.FlushMicrotasks()
	test.AssertElement("#result").HasText("Microtask 2")
}

func TestNestedTimers(t *testing.T) {
	test := domulator.NewTest(t)
	defer test.Close()

	test.LoadHTML(`
		<html>
		<body>
			<div id="steps">0</div>
			<script>
				let step = 0;
				
				setTimeout(() => {
					step = 1;
					document.getElementById('steps').textContent = step.toString();
					
					setTimeout(() => {
						step = 2;
						document.getElementById('steps').textContent = step.toString();
					}, 100);
				}, 100);
			</script>
		</body>
		</html>
	`)

	test.AssertElement("#steps").HasText("0")

	// Advance 100ms - first timer should fire and schedule second timer
	test.AdvanceTime(100 * time.Millisecond)
	test.AssertElement("#steps").HasText("1")

	// Advance another 100ms - second timer should fire
	test.AdvanceTime(100 * time.Millisecond)
	test.AssertElement("#steps").HasText("2")
}

func TestTimerWithMicrotasks(t *testing.T) {
	test := domulator.NewTest(t)
	defer test.Close()

	test.LoadHTML(`
		<html>
		<body>
			<div id="log"></div>
			<script>
				let log = [];
				
				setTimeout(() => {
					log.push('timer');
					queueMicrotask(() => {
						log.push('microtask-from-timer');
						document.getElementById('log').textContent = log.join(',');
					});
				}, 100);
				
				queueMicrotask(() => {
					log.push('immediate-microtask');
				});
			</script>
		</body>
		</html>
	`)

	// Process immediate microtask first
	test.FlushMicrotasks()

	// Then advance time to trigger timer, which should also process its microtask
	test.AdvanceTime(100 * time.Millisecond)
	test.AssertElement("#log").HasText("immediate-microtask,timer,microtask-from-timer")
}

func TestEventListenersInTimers(t *testing.T) {
	test := domulator.NewTest(t)
	defer test.Close()

	test.LoadHTML(`
		<html>
		<body>
			<button id="btn">Click me</button>
			<div id="result">Not clicked</div>
			<script>
				setTimeout(() => {
					document.getElementById('btn').addEventListener('click', () => {
						document.getElementById('result').textContent = 'Clicked!';
					});
				}, 100);
			</script>
		</body>
		</html>
	`)

	// Button click should not work initially
	test.Click("#btn")
	test.AssertElement("#result").HasText("Not clicked")

	// Advance time to add event listener
	test.AdvanceTime(100 * time.Millisecond)

	// Now button click should work
	test.Click("#btn")
	test.AssertElement("#result").HasText("Clicked!")
}

func TestProcessPendingTimers(t *testing.T) {
	test := domulator.NewTest(t)
	defer test.Close()

	test.LoadHTML(`
		<html>
		<body>
			<div id="status">Initial</div>
			<script>
				// Zero-delay timer should be ready immediately
				setTimeout(() => {
					document.getElementById('status').textContent = 'Ready';
				}, 0);
			</script>
		</body>
		</html>
	`)

	test.AssertElement("#status").HasText("Initial")

	// Process pending timers without advancing time
	test.ProcessPendingTimers()
	test.AssertElement("#status").HasText("Ready")
}

func TestZeroDelayTimers(t *testing.T) {
	test := domulator.NewTest(t)
	defer test.Close()

	test.LoadHTML(`
		<html>
		<body>
			<div id="order"></div>
			<script>
				let sequence = [];
				
				sequence.push('1');
				
				setTimeout(() => {
					sequence.push('3');
					document.getElementById('order').textContent = sequence.join('');
				}, 0);
				
				sequence.push('2');
			</script>
		</body>
		</html>
	`)

	// Initially should just have synchronous execution
	test.AssertElement("#order").HasText("")

	// Process zero-delay timer
	test.ProcessPendingTimers()
	test.AssertElement("#order").HasText("123")
}

func TestClearTimeout(t *testing.T) {
	test := domulator.NewTest(t)
	defer test.Close()

	test.LoadHTML(`
		<html>
		<body>
			<div id="status">Initial</div>
			<script>
				let timeoutId = setTimeout(() => {
					document.getElementById('status').textContent = 'Should not appear';
				}, 100);
				
				clearTimeout(timeoutId);
			</script>
		</body>
		</html>
	`)

	test.AssertElement("#status").HasText("Initial")

	// Advance time - timer should not fire since it was cleared
	test.AdvanceTime(200 * time.Millisecond)
	test.AssertElement("#status").HasText("Initial")
}

func TestComplexAsyncFlow(t *testing.T) {
	test := domulator.NewTest(t)
	defer test.Close()

	test.LoadHTML(`
		<html>
		<body>
			<div id="progress">0</div>
			<script>
				let progress = 0;
				
				function updateProgress() {
					progress++;
					document.getElementById('progress').textContent = progress.toString();
				}
				
				// Immediate microtask
				queueMicrotask(updateProgress);
				
				// Timer that queues another microtask
				setTimeout(() => {
					updateProgress();
					queueMicrotask(updateProgress);
					
					// Nested timer
					setTimeout(updateProgress, 50);
				}, 100);
				
				// Another timer
				setTimeout(updateProgress, 200);
			</script>
		</body>
		</html>
	`)

	test.AssertElement("#progress").HasText("0")

	// Process immediate microtask
	test.FlushMicrotasks()
	test.AssertElement("#progress").HasText("1")

	// Advance 100ms - first timer fires, then its microtask
	test.AdvanceTime(100 * time.Millisecond)
	test.AssertElement("#progress").HasText("3")

	// Advance 50ms more (150ms total) - nested timer fires
	test.AdvanceTime(50 * time.Millisecond)
	test.AssertElement("#progress").HasText("4")

	// Advance 50ms more (200ms total) - final timer fires
	test.AdvanceTime(50 * time.Millisecond)
	test.AssertElement("#progress").HasText("5")
}

func TestRequestAnimationFrame(t *testing.T) {
	test := domulator.NewTest(t)
	defer test.Close()

	test.LoadHTML(`
		<html>
		<body>
			<div id="frame">0</div>
			<script>
				let frameCount = 0;
				
				function animate(timestamp) {
					frameCount++;
					document.getElementById('frame').textContent = frameCount.toString();
					
					if (frameCount < 3) {
						requestAnimationFrame(animate);
					}
				}
				
				requestAnimationFrame(animate);
			</script>
		</body>
		</html>
	`)

	test.AssertElement("#frame").HasText("0")

	// Process first animation frame
	test.ProcessPendingTimers()
	test.AssertElement("#frame").HasText("1")

	// Process second animation frame
	test.ProcessPendingTimers()
	test.AssertElement("#frame").HasText("2")

	// Process third animation frame
	test.ProcessPendingTimers()
	test.AssertElement("#frame").HasText("3")

	// No more frames should be scheduled
	test.ProcessPendingTimers()
	test.AssertElement("#frame").HasText("3")
}

func TestEventLoopOrderingWithConsoleCapture(t *testing.T) {
	harness := testpkg.NewTestHarness()
	defer harness.Close()

	consoleCapture := harness.CaptureConsole(t)

	harness.LoadHTML(`
		<html>
		<body>
			<div id="sequence"></div>
			<script>
				console.log('Script starting...');
				
				let sequence = [];
				
				// Function to log and update DOM
				function addToSequence(item) {
					sequence.push(item);
					console.log('Added to sequence:', item);
					console.log('Current sequence:', sequence.join(' -> '));
					document.getElementById('sequence').textContent = sequence.join(' -> ');
				}
				
				// Synchronous code
				addToSequence('sync-1');
				
				// Microtask
				queueMicrotask(() => {
					addToSequence('microtask-1');
					
					// Nested microtask
					queueMicrotask(() => {
						addToSequence('nested-microtask');
					});
				});
				
				// Timer (task)
				setTimeout(() => {
					addToSequence('timer-1');
					
					// Microtask from timer
					queueMicrotask(() => {
						addToSequence('timer-microtask');
					});
				}, 10);
				
				// Another timer
				setTimeout(() => {
					addToSequence('timer-2');
				}, 20);
				
				// More synchronous code
				addToSequence('sync-2');
				
				console.log('Script setup complete');
			</script>
		</body>
		</html>
	`)

	// Assert initial console output
	consoleCapture.AssertLogContains("Script starting...")
	consoleCapture.AssertLogContains("Added to sequence:")
	consoleCapture.AssertLogContains("Script setup complete")

	// Initially should just have synchronous execution
	sequence := harness.Document().QuerySelector("#sequence")
	if sequence == nil {
		t.Fatal("Could not find sequence element")
	}
	if sequence.TextContent() != "sync-1 -> sync-2" {
		t.Errorf("Expected 'sync-1 -> sync-2', got %s", sequence.TextContent())
	}

	// Process microtasks - should run microtask-1 then nested-microtask
	harness.FlushMicrotasks()

	consoleCapture.AssertLogContains("microtask-1")
	consoleCapture.AssertLogContains("nested-microtask")

	if sequence.TextContent() != "sync-1 -> sync-2 -> microtask-1 -> nested-microtask" {
		t.Errorf("Expected microtasks to be processed, got %s", sequence.TextContent())
	}

	// Advance time to trigger first timer (10ms)
	harness.AdvanceTime(10 * time.Millisecond)

	consoleCapture.AssertLogContains("timer-1")
	consoleCapture.AssertLogContains("timer-microtask")

	expectedAfterTimer1 := "sync-1 -> sync-2 -> microtask-1 -> nested-microtask -> timer-1 -> timer-microtask"
	if sequence.TextContent() != expectedAfterTimer1 {
		t.Errorf("Expected timer sequence, got %s", sequence.TextContent())
	}

	// Advance time to trigger second timer (20ms total)
	harness.AdvanceTime(10 * time.Millisecond)

	consoleCapture.AssertLogContains("timer-2")

	finalExpected := "sync-1 -> sync-2 -> microtask-1 -> nested-microtask -> timer-1 -> timer-microtask -> timer-2"
	if sequence.TextContent() != finalExpected {
		t.Errorf("Expected final sequence, got %s", sequence.TextContent())
	}

	// Verify sequence demonstrates proper event loop ordering:
	// 1. Synchronous code executes first
	// 2. Microtasks process before any timers
	// 3. Nested microtasks process in order
	// 4. Timers execute in chronological order
	// 5. Microtasks from timers execute immediately after timer callback
	consoleCapture.AssertLogContains("Current sequence:")

	t.Log("✅ Event loop ordering with console capture completed successfully!")
}
