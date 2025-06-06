package integration

import (
	"testing"

	domulator "github.com/dpemmons/DOMulator"
)

func TestDebugFocus(t *testing.T) {
	test := domulator.NewTest(t)
	test.LoadHTML(`
		<input id="input" type="text">
		
		<script>
			console.log('Setting up event listener...');
			let callCount = 0;
			
			document.getElementById('input').addEventListener('focus', () => {
				callCount++;
				console.log('Focus event received! Call count:', callCount);
				document.getElementById('input').setAttribute('data-focus-count', callCount.toString());
			});
			
			console.log('Event listener setup complete');
		</script>
	`)

	// Test focus once
	t.Log("Calling Focus once...")
	test.Focus("#input")

	// Check how many times the event fired
	test.AssertElement("#input").HasAttribute("data-focus-count", "1")
}
