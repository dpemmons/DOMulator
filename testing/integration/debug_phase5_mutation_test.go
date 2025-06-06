package integration

import (
	"testing"
	"time"

	domulator "github.com/dpemmons/DOMulator"
)

func TestDebugPhase5MutationPattern(t *testing.T) {
	test := domulator.NewTest(t).SetDebugMode(true)

	test.LoadHTML(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>Phase 5 MutationObserver Debug</title>
		</head>
		<body>
			<div id="observed">
				<p id="child1">Original content</p>
			</div>
		</body>
		</html>
		<script>
		let mutationCount = 0;
		let mutationTypes = [];
		
		console.log("Script starting...");
		console.log("MutationObserver constructor:", typeof MutationObserver);
		
		const observer = new MutationObserver(function(mutations) {
			console.log("=== MutationObserver callback triggered ===");
			console.log("Mutations received:", mutations.length);
			
			mutations.forEach(function(mutation, index) {
				mutationCount++;
				mutationTypes.push(mutation.type);
				console.log("Mutation", index, ":", mutation.type, "on", mutation.target.tagName || mutation.target.nodeType);
				
				// Update DOM to signal mutations were observed
				document.body.setAttribute('data-mutation-count', mutationCount.toString());
				document.body.setAttribute('data-mutation-types', mutationTypes.join(','));
				console.log("Updated body attributes");
			});
		});
		
		console.log("MutationObserver created");
		
		const observed = document.getElementById('observed');
		
		try {
			observer.observe(observed, {
				childList: true,
				attributes: true,
				characterData: true,
				subtree: true
			});
			console.log("Observer started on element:", observed.id);
		} catch (e) {
			console.log("Observer start error:", e.message);
		}
		
		console.log("Setting up first setTimeout...");
		setTimeout(() => {
			console.log("=== First setTimeout executing ===");
			
			// Child list mutation
			console.log("Creating new child element...");
			const newChild = document.createElement('p');
			newChild.textContent = 'New paragraph';
			console.log("Appending new child...");
			observed.appendChild(newChild);
			console.log("Child appended");
			
			// Attribute mutation
			console.log("Setting attribute...");
			observed.setAttribute('data-test', 'value');
			console.log("Attribute set");
			
			// Text mutation
			console.log("Modifying text content...");
			const child1 = document.getElementById('child1');
			if (child1) {
				child1.textContent = 'Modified content';
				console.log("Text content modified");
			} else {
				console.log("child1 not found!");
			}
			
			console.log("All mutations triggered, current mutation count:", mutationCount);
			console.log("Setting up second setTimeout...");
			
			setTimeout(() => {
				console.log("=== Second setTimeout executing ===");
				console.log("Final mutation count:", mutationCount);
				console.log("Final mutation types:", mutationTypes.join(','));
				
				observer.disconnect();
				console.log("Observer disconnected");
				
				// Final verification
				let results = [];
				results.push("constructor: " + (typeof MutationObserver === 'function'));
				results.push("creation: " + (observer !== null));
				results.push("methods: " + (typeof observer.observe === 'function'));
				results.push("mutations: " + (mutationCount > 0));
				
				document.body.setAttribute('data-observer-results', results.join(','));
				document.body.setAttribute('data-observer-test', 'complete');
				console.log("Set final attributes");
				
			}, 50);
		}, 10);
		
		console.log("Script setup complete");
		</script>
	`)

	// Check immediate state
	debugState := test.Document().Body().GetAttribute("data-observer-test")
	t.Logf("Immediate observer test state: '%s'", debugState)

	// Advance virtual time to execute the setTimeout calls and process microtasks
	test.AdvanceTime(100 * time.Millisecond)
	test.FlushMicrotasks()
	test.ProcessPendingTimers()

	// Process any additional timers that may have been scheduled
	test.AdvanceTime(100 * time.Millisecond)
	test.ProcessPendingTimers()

	// Check final state
	finalState := test.Document().Body().GetAttribute("data-observer-test")
	finalResults := test.Document().Body().GetAttribute("data-observer-results")
	mutationCount := test.Document().Body().GetAttribute("data-mutation-count")
	mutationTypes := test.Document().Body().GetAttribute("data-mutation-types")

	t.Logf("Final observer test state: '%s'", finalState)
	t.Logf("Final observer results: '%s'", finalResults)
	t.Logf("Final mutation count: '%s'", mutationCount)
	t.Logf("Final mutation types: '%s'", mutationTypes)

	// This test just logs info for now
}
