package integration

import (
	"testing"

	domulator "github.com/dpemmons/DOMulator"
)

func TestBasicJavaScriptExecution(t *testing.T) {
	test := domulator.NewTest(t)
	test.LoadHTML(`
		<button id="btn">Click me</button>
		
		<script>
			console.log('JavaScript is executing...');
			
			// Set an attribute to show JS ran
			document.body.setAttribute('data-js-executed', 'yes');
			
			// Simple event listener
			document.getElementById('btn').addEventListener('click', function() {
				console.log('Click event received!');
				document.getElementById('btn').setAttribute('data-clicked', 'true');
			});
			
			console.log('Event listener setup complete!');
		</script>
	`)

	// First verify JavaScript executed
	test.AssertElement("body").HasAttribute("data-js-executed", "yes")

	// Now try a simple click and see if the event is received
	test.Click("#btn")
	test.AssertElement("#btn").HasAttribute("data-clicked", "true")

	t.Log("✅ Basic JavaScript execution and event dispatch working!")
}

func TestDirectEventDispatch(t *testing.T) {
	test := domulator.NewTest(t)
	test.LoadHTML(`
		<input id="input" type="text">
		
		<script>
			console.log('Setting up keydown listener...');
			
			var inputElement = document.getElementById('input');
			inputElement.addEventListener('keydown', function(event) {
				console.log('Keydown event received:', event);
				console.log('this:', this);
				console.log('inputElement:', inputElement);
				
				// Try both this and the explicit element reference
				inputElement.setAttribute('data-keydown-received', 'yes');
				console.log('Attribute set on inputElement');
				
				// Also try with this
				if (this && this.setAttribute) {
					this.setAttribute('data-keydown-received-this', 'yes');
					console.log('Attribute set on this');
				} else {
					console.log('this does not have setAttribute method');
				}
			});
			
			console.log('Keydown listener setup complete!');
		</script>
	`)

	// Try the keydown event that was failing
	test.KeyDown("#input", "a")
	test.AssertElement("#input").HasAttribute("data-keydown-received", "yes")

	t.Log("✅ Direct keydown event dispatch working!")
}
