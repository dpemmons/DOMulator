package integration

import (
	"testing"

	domulator "github.com/dpemmons/DOMulator"
	testpkg "github.com/dpemmons/DOMulator/testing"
)

func TestMutationObserver(t *testing.T) {
	test := domulator.NewTest(t)
	test.LoadHTML(`
		<div id="container">
			<p id="target">Original content</p>
		</div>
		
		<script>
			// Track mutations received by the observer
			let mutations = [];
			let callbackCount = 0;
			
			// Create a MutationObserver
			const observer = new MutationObserver((mutationsList, observer) => {
				callbackCount++;
				console.log('MutationObserver callback fired! Count:', callbackCount);
				
				for (let mutation of mutationsList) {
					mutations.push({
						type: mutation.type,
						target: mutation.target.id || mutation.target.tagName,
						attributeName: mutation.attributeName,
						addedNodesCount: mutation.addedNodes.length,
						removedNodesCount: mutation.removedNodes.length
					});
				}
				
				// Store results in DOM for testing
				document.getElementById('container').setAttribute('data-callback-count', callbackCount.toString());
				document.getElementById('container').setAttribute('data-mutations', JSON.stringify(mutations));
			});
			
			// Start observing
			observer.observe(document.getElementById('target'), {
				attributes: true,
				childList: true,
				characterData: true,
				subtree: true
			});
			
			console.log('MutationObserver setup complete');
			
			// Initialize the callback count attribute
			document.getElementById('container').setAttribute('data-callback-count', '0');
		</script>
	`)

	// Verify observer is set up but no mutations yet
	test.AssertElement("#container").HasAttribute("data-callback-count", "0")

	// Test 1: Attribute mutation
	test.ExecuteScript(`
		document.getElementById('target').setAttribute('data-test', 'value1');
	`)

	// Process microtasks to trigger observer callback
	test.FlushMicrotasks()

	// Verify the mutation was observed
	test.AssertElement("#container").HasAttribute("data-callback-count", "1")

	// Test 2: Child list mutation (adding node)
	test.ExecuteScript(`
		const newElement = document.createElement('span');
		newElement.textContent = 'New element';
		document.getElementById('target').appendChild(newElement);
	`)

	test.FlushMicrotasks()
	test.AssertElement("#container").HasAttribute("data-callback-count", "2")

	// Test 3: Character data mutation
	test.ExecuteScript(`
		document.getElementById('target').firstChild.textContent = 'Modified content';
	`)

	test.FlushMicrotasks()
	test.AssertElement("#container").HasAttribute("data-callback-count", "3")

	// Test 4: Multiple mutations in one synchronous block
	test.ExecuteScript(`
		const target = document.getElementById('target');
		target.setAttribute('data-test', 'value2');
		target.setAttribute('data-another', 'value3');
		target.className = 'modified';
	`)

	test.FlushMicrotasks()
	// Should be at least 4 now (could be more if each attribute change triggers separately)
	test.AssertElement("#container").HasAttributeContaining("data-callback-count", "4")

	t.Log("✅ MutationObserver works correctly with microtask processing!")
}

func TestMutationObserverDisconnect(t *testing.T) {
	test := domulator.NewTest(t)
	test.LoadHTML(`
		<div id="container">
			<p id="target">Content</p>
		</div>
		
		<script>
			let callbackCount = 0;
			
			const observer = new MutationObserver(() => {
				callbackCount++;
				document.getElementById('container').setAttribute('data-callback-count', callbackCount.toString());
			});
			
			observer.observe(document.getElementById('target'), {
				attributes: true,
				childList: true
			});
			
			// Store observer reference globally so we can disconnect it
			window.testObserver = observer;
		</script>
	`)

	// Make a change - should trigger observer
	test.ExecuteScript(`document.getElementById('target').setAttribute('data-test', 'value1');`)
	test.FlushMicrotasks()
	test.AssertElement("#container").HasAttribute("data-callback-count", "1")

	// Disconnect the observer
	test.ExecuteScript(`window.testObserver.disconnect();`)

	// Make another change - should NOT trigger observer
	test.ExecuteScript(`document.getElementById('target').setAttribute('data-test', 'value2');`)
	test.FlushMicrotasks()

	// Count should still be 1
	test.AssertElement("#container").HasAttribute("data-callback-count", "1")

	t.Log("✅ MutationObserver disconnect works correctly!")
}

func TestMutationObserverTakeRecords(t *testing.T) {
	test := domulator.NewTest(t)
	test.LoadHTML(`
		<div id="container">
			<p id="target">Content</p>
		</div>
		
		<script>
			let callbackCount = 0;
			let takenRecords = [];
			
			const observer = new MutationObserver(() => {
				callbackCount++;
				document.getElementById('container').setAttribute('data-callback-count', callbackCount.toString());
			});
			
			observer.observe(document.getElementById('target'), {
				attributes: true,
				childList: true
			});
			
			// Store observer reference globally
			window.testObserver = observer;
		</script>
	`)

	// Make changes without processing microtasks
	test.ExecuteScript(`
		document.getElementById('target').setAttribute('data-test', 'value1');
		document.getElementById('target').setAttribute('data-test2', 'value2');
	`)

	// Take records before microtasks are processed
	test.ExecuteScript(`
		const records = window.testObserver.takeRecords();
		document.getElementById('container').setAttribute('data-taken-count', records.length.toString());
	`)

	// Now process microtasks - callback should not fire since we took the records
	test.FlushMicrotasks()

	// Verify records were taken and callback didn't fire
	test.AssertElement("#container").HasAttribute("data-taken-count", "2")
	test.AssertElement("#container").HasAttribute("data-callback-count", "0")

	t.Log("✅ MutationObserver takeRecords works correctly!")
}

func TestMutationObserverWithEventLoop(t *testing.T) {
	test := domulator.NewTest(t)
	test.LoadHTML(`
		<div id="container">
			<p id="target">Content</p>
		</div>
		
		<script>
			let sequence = [];
			
			// Set up observer
			const observer = new MutationObserver(() => {
				sequence.push('mutation-callback');
				document.getElementById('container').setAttribute('data-sequence', sequence.join(','));
			});
			
			observer.observe(document.getElementById('target'), {
				attributes: true
			});
			
			// Function to test async timing
			function testAsyncSequence() {
				sequence.push('sync-start');
				
				// Mutation (will queue microtask)
				document.getElementById('target').setAttribute('data-test', 'value');
				
				// Promise (will queue microtask)
				Promise.resolve().then(() => {
					sequence.push('promise-then');
					document.getElementById('container').setAttribute('data-sequence', sequence.join(','));
				});
				
				// Timeout (will queue task)
				setTimeout(() => {
					sequence.push('timeout');
					document.getElementById('container').setAttribute('data-sequence', sequence.join(','));
				}, 0);
				
				sequence.push('sync-end');
				document.getElementById('container').setAttribute('data-sequence', sequence.join(','));
			}
			
			window.testAsyncSequence = testAsyncSequence;
		</script>
	`)

	// Run the async sequence test
	test.ExecuteScript(`window.testAsyncSequence();`)

	// Should have sync operations first
	test.AssertElement("#container").HasAttributeContaining("data-sequence", "sync-start")
	test.AssertElement("#container").HasAttributeContaining("data-sequence", "sync-end")

	// Process microtasks - should execute Promise and MutationObserver callbacks
	test.FlushMicrotasks()
	test.AssertElement("#container").HasAttributeContaining("data-sequence", "mutation-callback")
	test.AssertElement("#container").HasAttributeContaining("data-sequence", "promise-then")

	// Process task queue - should execute setTimeout
	test.ProcessPendingTimers()
	test.AssertElement("#container").HasAttributeContaining("data-sequence", "timeout")

	// Verify correct order: sync operations, then microtasks, then tasks
	test.ExecuteScript(`
		const seq = document.getElementById('container').getAttribute('data-sequence');
		const parts = seq.split(',');
		const syncEndIndex = parts.indexOf('sync-end');
		const mutationIndex = parts.indexOf('mutation-callback');
		const promiseIndex = parts.indexOf('promise-then');
		const timeoutIndex = parts.indexOf('timeout');
		
		// Verify order is correct
		const orderCorrect = syncEndIndex < mutationIndex && 
		                    syncEndIndex < promiseIndex && 
		                    mutationIndex < timeoutIndex && 
		                    promiseIndex < timeoutIndex;
		
		document.getElementById('container').setAttribute('data-order-correct', orderCorrect.toString());
	`)

	test.AssertElement("#container").HasAttribute("data-order-correct", "true")

	t.Log("✅ MutationObserver respects proper event loop ordering!")
}

func TestMutationObserverWithConsoleCapture(t *testing.T) {
	harness := testpkg.NewTestHarness()
	defer harness.Close()

	consoleCapture := harness.CaptureConsole(t)

	harness.LoadHTML(`
		<div id="container">
			<p id="target">Original content</p>
		</div>
		
		<script>
			console.log('Setting up MutationObserver...');
			
			// Track mutations received by the observer
			let mutations = [];
			let callbackCount = 0;
			
			// Create a MutationObserver
			const observer = new MutationObserver((mutationsList, observer) => {
				callbackCount++;
				console.log('MutationObserver callback fired! Count:', callbackCount);
				console.log('Mutations received:', mutationsList.length);
				
				for (let mutation of mutationsList) {
					console.log('Mutation type:', mutation.type);
					if (mutation.type === 'childList') {
						console.log('Child nodes added:', mutation.addedNodes.length);
						console.log('Child nodes removed:', mutation.removedNodes.length);
					}
					
					mutations.push({
						type: mutation.type,
						target: mutation.target.id || mutation.target.tagName,
						addedNodesCount: mutation.addedNodes.length,
						removedNodesCount: mutation.removedNodes.length
					});
				}
				
				// Store results in DOM for testing
				document.getElementById('container').setAttribute('data-callback-count', callbackCount.toString());
				document.getElementById('container').setAttribute('data-mutations', JSON.stringify(mutations));
			});
			
			// Start observing (focus on childList mutations which seem to work)
			observer.observe(document.getElementById('target'), {
				childList: true,
				subtree: true
			});
			
			console.log('MutationObserver setup complete');
			console.log('Observing element:', document.getElementById('target').id);
			
			// Initialize the callback count attribute
			document.getElementById('container').setAttribute('data-callback-count', '0');
		</script>
	`)

	// Assert initial console output
	consoleCapture.AssertLogContains("Setting up MutationObserver...")
	consoleCapture.AssertLogContains("MutationObserver setup complete")
	consoleCapture.AssertLogContains("Observing element:")
	consoleCapture.AssertLogContains("target")

	// Verify observer is set up but no mutations yet
	container := harness.Document().QuerySelector("#container")
	if container == nil {
		t.Fatal("Could not find container element")
	}
	if container.GetAttribute("data-callback-count") != "0" {
		t.Errorf("Expected callback count 0, got %s", container.GetAttribute("data-callback-count"))
	}

	// Test 1: Child list mutation (adding node)
	harness.ExecuteScript(`
		console.log('Testing child list mutation...');
		const newElement = document.createElement('span');
		newElement.textContent = 'New element';
		document.getElementById('target').appendChild(newElement);
	`)

	harness.FlushMicrotasks()

	// Assert console output for child mutation
	consoleCapture.AssertLogContains("Testing child list mutation...")
	consoleCapture.AssertLogContains("MutationObserver callback fired!")

	if container.GetAttribute("data-callback-count") != "1" {
		t.Errorf("Expected callback count 1, got %s", container.GetAttribute("data-callback-count"))
	}

	// Test 2: Adding another child node
	harness.ExecuteScript(`
		console.log('Testing second child list mutation...');
		const anotherElement = document.createElement('div');
		anotherElement.textContent = 'Another element';
		document.getElementById('target').appendChild(anotherElement);
	`)

	harness.FlushMicrotasks()

	// Assert console output for second mutation
	consoleCapture.AssertLogContains("Testing second child list mutation...")

	if container.GetAttribute("data-callback-count") != "2" {
		t.Errorf("Expected callback count 2, got %s", container.GetAttribute("data-callback-count"))
	}

	// Test 3: Removing a child node
	harness.ExecuteScript(`
		console.log('Testing child removal...');
		const target = document.getElementById('target');
		const spans = target.getElementsByTagName('span');
		if (spans.length > 0) {
			target.removeChild(spans[0]);
		}
	`)

	harness.FlushMicrotasks()

	// Assert console output for removal
	consoleCapture.AssertLogContains("Testing child removal...")

	if container.GetAttribute("data-callback-count") != "3" {
		t.Errorf("Expected callback count 3, got %s", container.GetAttribute("data-callback-count"))
	}

	// Verify that MutationObserver logs are captured
	consoleCapture.AssertLogContains("Mutation type:")
	consoleCapture.AssertLogContains("Child nodes")

	t.Log("✅ MutationObserver with console capture completed successfully!")
}
