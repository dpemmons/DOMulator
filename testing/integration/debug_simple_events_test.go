package integration

import (
	"testing"

	domulator "github.com/dpemmons/DOMulator"
)

func TestDebugBasicEventTrigger(t *testing.T) {
	test := domulator.NewTest(t)
	test.SetDebugMode(true) // Enable debug mode to see console output

	test.LoadHTML(`
		<button id="btn">Click me</button>
		<input id="input" type="text">
		
		<script>
			console.log('Setting up event listeners...');
			
			const input = document.getElementById('input');
			const btn = document.getElementById('btn');
			
			// Function to mark an element as having received an event
			function markEvent(element, eventType) {
				if (element) {
					element.setAttribute('data-last-event', eventType);
					element.setAttribute('data-event-count', (parseInt(element.getAttribute('data-event-count') || '0') + 1).toString());
					console.log('🎯 Event received:', eventType, 'on', element.id || element.tagName);
				}
			}
			
			// Keyboard events
			input.addEventListener('keydown', () => markEvent(input, 'keydown'));
			input.addEventListener('keyup', () => markEvent(input, 'keyup'));
			input.addEventListener('keypress', () => markEvent(input, 'keypress'));
			
			// Mouse events
			btn.addEventListener('mousedown', () => markEvent(btn, 'mousedown'));
			btn.addEventListener('mouseup', () => markEvent(btn, 'mouseup'));
			btn.addEventListener('click', () => markEvent(btn, 'click'));
			
			console.log('✅ Event listeners setup complete!');
		</script>
	`)

	// Test keyboard events
	test.KeyDown("#input", "a")
	test.AssertElement("#input").HasAttribute("data-last-event", "keydown")
	t.Logf("✅ KeyDown event received by JavaScript")

	test.KeyUp("#input", "a")
	test.AssertElement("#input").HasAttribute("data-last-event", "keyup")
	t.Logf("✅ KeyUp event received by JavaScript")

	test.KeyPress("#input", "a")
	test.AssertElement("#input").HasAttribute("data-last-event", "keypress")
	t.Logf("✅ KeyPress event received by JavaScript")

	// Test mouse events
	test.MouseDown("#btn")
	test.AssertElement("#btn").HasAttribute("data-last-event", "mousedown")
	t.Logf("✅ MouseDown event received by JavaScript")

	test.MouseUp("#btn")
	test.AssertElement("#btn").HasAttribute("data-last-event", "mouseup")
	t.Logf("✅ MouseUp event received by JavaScript")

	test.Click("#btn")
	test.AssertElement("#btn").HasAttribute("data-last-event", "click")
	t.Logf("✅ Click event received by JavaScript")

	// Check event counts to prove multiple events were processed
	inputEventCount := test.Document().QuerySelector("#input").GetAttribute("data-event-count")
	btnEventCount := test.Document().QuerySelector("#btn").GetAttribute("data-event-count")

	t.Logf("📊 Input received %s events, Button received %s events", inputEventCount, btnEventCount)

	t.Log("🎉 All events successfully round-tripped: Go → DOM Events → JavaScript Listeners → DOM Updates!")
}
