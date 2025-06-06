package integration

import (
	"testing"
	"time"

	domulator "github.com/dpemmons/DOMulator"
)

func TestDebugMutationObserver(t *testing.T) {
	test := domulator.NewTest(t).SetDebugMode(true)

	test.LoadHTML(`
		<!DOCTYPE html>
		<html>
		<body>
			<div id="observed">
				<p id="child1">Original content</p>
			</div>
		</body>
		</html>
		<script>
		console.log("Script starting...");
		
		// Test if MutationObserver exists
		console.log("MutationObserver type:", typeof MutationObserver);
		console.log("MutationObserver constructor:", MutationObserver);
		
		if (typeof MutationObserver === 'undefined') {
			document.body.setAttribute('data-debug', 'no-mutationobserver');
			throw new Error("MutationObserver not defined");
		}
		
		let mutationCount = 0;
		
		const observer = new MutationObserver(function(mutations) {
			console.log("MutationObserver callback triggered with", mutations.length, "mutations");
			mutationCount += mutations.length;
			
			mutations.forEach(function(mutation, index) {
				console.log("Mutation", index + ":", mutation.type, "on", mutation.target.tagName);
			});
			
			document.body.setAttribute('data-mutation-count', mutationCount.toString());
		});
		
		console.log("MutationObserver created successfully");
		
		const observed = document.getElementById('observed');
		console.log("Found observing element:", observed.id);
		
		observer.observe(observed, {
			childList: true,
			attributes: true,
			characterData: true,
			subtree: true
		});
		
		console.log("Observer started");
		
		// Test immediate mutation (no setTimeout)
		console.log("Triggering immediate mutation...");
		const newChild = document.createElement('p');
		newChild.textContent = 'New paragraph';
		observed.appendChild(newChild);
		console.log("Child appended");
		
		// Mark that we got this far
		document.body.setAttribute('data-debug', 'immediate-mutation-done');
		
		// Test setTimeout mutation
		setTimeout(() => {
			console.log("setTimeout callback executing...");
			observed.setAttribute('data-test', 'value');
			console.log("Attribute set");
			
			document.body.setAttribute('data-debug', 'timeout-mutation-done');
			
			setTimeout(() => {
				console.log("Second setTimeout executing...");
				observer.disconnect();
				document.body.setAttribute('data-debug', 'complete');
			}, 10);
		}, 10);
		
		console.log("Script complete");
		</script>
	`)

	// Check immediate state
	test.AssertElement("body").HasAttribute("data-debug", "immediate-mutation-done")

	// Check if immediate mutation was observed
	mutationCount := test.Document().Body().GetAttribute("data-mutation-count")
	t.Logf("Mutation count after immediate mutation: %s", mutationCount)

	// Advance time to process setTimeout
	test.AdvanceTime(50 * time.Millisecond)
	test.FlushMicrotasks()

	// Check final state
	finalDebug := test.Document().Body().GetAttribute("data-debug")
	finalMutationCount := test.Document().Body().GetAttribute("data-mutation-count")

	t.Logf("Final debug state: %s", finalDebug)
	t.Logf("Final mutation count: %s", finalMutationCount)

	// This test just logs info, doesn't assert success
}
