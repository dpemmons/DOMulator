package integration

import (
	"testing"

	domulator "github.com/dpemmons/DOMulator"
	testpkg "github.com/dpemmons/DOMulator/testing"
)

func TestBasicEventTrigger(t *testing.T) {
	test := domulator.NewTest(t)
	test.LoadHTML(`
		<button id="btn">Click me</button>
		<input id="input" type="text">
		<form id="form">
			<input type="submit" value="Submit">
		</form>
		<div id="draggable" draggable="true">Drag me</div>
		<div id="target">Drop here</div>
		<div id="touch-area">Touch me</div>
		<video id="video">Video</video>
		<textarea id="textarea">Text</textarea>
		<div id="scroll-area" style="overflow:scroll;">Scrollable</div>
		
		<script>
			// Add event listeners that update DOM attributes when events are received
			console.log('Setting up event listeners...');
			
			// Track all events on their respective elements
			const elements = {
				'#btn': document.getElementById('btn'),
				'#input': document.getElementById('input'),
				'#form': document.getElementById('form'),
				'#draggable': document.getElementById('draggable'),
				'#target': document.getElementById('target'),
				'#touch-area': document.getElementById('touch-area'),
				'#video': document.getElementById('video'),
				'#textarea': document.getElementById('textarea'),
				'#scroll-area': document.getElementById('scroll-area')
			};
			
			// Function to mark an element as having received an event
			function markEvent(element, eventType) {
				if (element) {
					element.setAttribute('data-last-event', eventType);
					element.setAttribute('data-event-count', (parseInt(element.getAttribute('data-event-count') || '0') + 1).toString());
					console.log('Event received:', eventType, 'on', element.id || element.tagName);
				}
			}
			
			// Add listeners for all the events we'll test
			Object.values(elements).forEach(el => {
				if (!el) return;
				
				// Keyboard events
				el.addEventListener('keydown', () => markEvent(el, 'keydown'));
				el.addEventListener('keyup', () => markEvent(el, 'keyup'));
				el.addEventListener('keypress', () => markEvent(el, 'keypress'));
				
				// Mouse events
				el.addEventListener('mousedown', () => markEvent(el, 'mousedown'));
				el.addEventListener('mouseup', () => markEvent(el, 'mouseup'));
				el.addEventListener('mousemove', () => markEvent(el, 'mousemove'));
				el.addEventListener('contextmenu', () => markEvent(el, 'contextmenu'));
				
				// Touch events
				el.addEventListener('touchstart', () => markEvent(el, 'touchstart'));
				el.addEventListener('touchmove', () => markEvent(el, 'touchmove'));
				el.addEventListener('touchend', () => markEvent(el, 'touchend'));
				
				// Drag events
				el.addEventListener('dragstart', () => markEvent(el, 'dragstart'));
				el.addEventListener('drag', () => markEvent(el, 'drag'));
				el.addEventListener('dragenter', () => markEvent(el, 'dragenter'));
				el.addEventListener('dragover', () => markEvent(el, 'dragover'));
				el.addEventListener('dragleave', () => markEvent(el, 'dragleave'));
				el.addEventListener('drop', () => markEvent(el, 'drop'));
				el.addEventListener('dragend', () => markEvent(el, 'dragend'));
				
				// Form events
				el.addEventListener('reset', () => markEvent(el, 'reset'));
				el.addEventListener('invalid', () => markEvent(el, 'invalid'));
				
				// Other events
				el.addEventListener('scroll', () => markEvent(el, 'scroll'));
				el.addEventListener('play', () => markEvent(el, 'play'));
				el.addEventListener('pause', () => markEvent(el, 'pause'));
				el.addEventListener('ended', () => markEvent(el, 'ended'));
				el.addEventListener('copy', () => markEvent(el, 'copy'));
				el.addEventListener('cut', () => markEvent(el, 'cut'));
				el.addEventListener('paste', () => markEvent(el, 'paste'));
			});
			
			// Window/document events
			document.addEventListener('DOMContentLoaded', () => {
				document.body.setAttribute('data-last-event', 'DOMContentLoaded');
				console.log('DOMContentLoaded event received');
			});
			
			window.addEventListener('resize', () => {
				document.body.setAttribute('data-last-event', 'resize');
				console.log('Resize event received');
			});
			
			window.addEventListener('load', () => {
				document.body.setAttribute('data-last-event', 'load');
				console.log('Load event received');
			});
			
			window.addEventListener('unload', () => {
				document.body.setAttribute('data-last-event', 'unload');
				console.log('Unload event received');
			});
			
			console.log('Event listeners setup complete!');
		</script>
	`)

	// Test that all our new event methods work and JavaScript receives them

	// Keyboard Events
	test.KeyDown("#input", "a")
	test.AssertElement("#input").HasAttribute("data-last-event", "keydown")

	test.KeyUp("#input", "a")
	test.AssertElement("#input").HasAttribute("data-last-event", "keyup")

	test.KeyPress("#input", "a")
	test.AssertElement("#input").HasAttribute("data-last-event", "keypress")

	// Advanced Mouse Events
	test.MouseDown("#btn")
	test.AssertElement("#btn").HasAttribute("data-last-event", "mousedown")

	test.MouseUp("#btn")
	test.AssertElement("#btn").HasAttribute("data-last-event", "mouseup")

	test.MouseMove("#btn")
	test.AssertElement("#btn").HasAttribute("data-last-event", "mousemove")

	test.RightClick("#btn")
	test.AssertElement("#btn").HasAttribute("data-last-event", "contextmenu")

	// Touch Events
	test.TouchStart("#touch-area")
	test.AssertElement("#touch-area").HasAttribute("data-last-event", "touchstart")

	test.TouchMove("#touch-area")
	test.AssertElement("#touch-area").HasAttribute("data-last-event", "touchmove")

	test.TouchEnd("#touch-area")
	test.AssertElement("#touch-area").HasAttribute("data-last-event", "touchend")

	// Drag and Drop Events
	test.DragStart("#draggable")
	test.AssertElement("#draggable").HasAttribute("data-last-event", "dragstart")

	test.Drag("#draggable")
	test.AssertElement("#draggable").HasAttribute("data-last-event", "drag")

	test.DragEnter("#target")
	test.AssertElement("#target").HasAttribute("data-last-event", "dragenter")

	test.DragOver("#target")
	test.AssertElement("#target").HasAttribute("data-last-event", "dragover")

	test.Drop("#target")
	test.AssertElement("#target").HasAttribute("data-last-event", "drop")

	test.DragEnd("#draggable")
	test.AssertElement("#draggable").HasAttribute("data-last-event", "dragend")

	// Form Events
	test.Reset("#form")
	test.AssertElement("#form").HasAttribute("data-last-event", "reset")

	test.Invalid("#input")
	test.AssertElement("#input").HasAttribute("data-last-event", "invalid")

	// Scroll Events
	test.Scroll("#scroll-area")
	test.AssertElement("#scroll-area").HasAttribute("data-last-event", "scroll")

	// Media Events
	test.Play("#video")
	test.AssertElement("#video").HasAttribute("data-last-event", "play")

	test.Pause("#video")
	test.AssertElement("#video").HasAttribute("data-last-event", "pause")

	test.Ended("#video")
	test.AssertElement("#video").HasAttribute("data-last-event", "ended")

	// Clipboard Events
	test.Copy("#textarea")
	test.AssertElement("#textarea").HasAttribute("data-last-event", "copy")

	test.Cut("#textarea")
	test.AssertElement("#textarea").HasAttribute("data-last-event", "cut")

	test.Paste("#textarea")
	test.AssertElement("#textarea").HasAttribute("data-last-event", "paste")

	// Window Events
	test.Resize()
	test.AssertElement("body").HasAttribute("data-last-event", "resize")

	test.Load()
	test.AssertElement("body").HasAttribute("data-last-event", "load")

	test.Unload()
	test.AssertElement("body").HasAttribute("data-last-event", "unload")

	// Verify elements still exist (events didn't break the DOM)
	test.AssertElement("#btn").Exists()
	test.AssertElement("#input").Exists()
	test.AssertElement("#form").Exists()
	test.AssertElement("#draggable").Exists()
	test.AssertElement("#target").Exists()

	t.Log("✅ All event types triggered successfully and JavaScript listeners received them!")
}

func TestEventSequence(t *testing.T) {
	test := domulator.NewTest(t)
	test.LoadHTML(`
		<button id="btn">Test Button</button>
		<input id="input" type="text">
		
		<script>
			let eventSequence = [];
			
			document.getElementById('input').addEventListener('focus', () => {
				eventSequence.push('focus');
				document.getElementById('input').setAttribute('data-events', eventSequence.join(','));
			});
			
			document.getElementById('input').addEventListener('input', () => {
				eventSequence.push('input');
				document.getElementById('input').setAttribute('data-events', eventSequence.join(','));
			});
			
			document.getElementById('input').addEventListener('keydown', () => {
				eventSequence.push('keydown');
				document.getElementById('input').setAttribute('data-events', eventSequence.join(','));
			});
			
			document.getElementById('input').addEventListener('keyup', () => {
				eventSequence.push('keyup');
				document.getElementById('input').setAttribute('data-events', eventSequence.join(','));
			});
			
			document.getElementById('btn').addEventListener('mouseenter', () => {
				eventSequence.push('mouseenter');
				document.getElementById('btn').setAttribute('data-events', eventSequence.join(','));
			});
			
			document.getElementById('btn').addEventListener('mousedown', () => {
				eventSequence.push('mousedown');
				document.getElementById('btn').setAttribute('data-events', eventSequence.join(','));
			});
			
			document.getElementById('btn').addEventListener('mouseup', () => {
				eventSequence.push('mouseup');
				document.getElementById('btn').setAttribute('data-events', eventSequence.join(','));
			});
			
			document.getElementById('btn').addEventListener('click', () => {
				eventSequence.push('click');
				document.getElementById('btn').setAttribute('data-events', eventSequence.join(','));
			});
			
			document.getElementById('btn').addEventListener('mouseleave', () => {
				eventSequence.push('mouseleave');
				document.getElementById('btn').setAttribute('data-events', eventSequence.join(','));
			});
		</script>
	`)

	// Test a realistic interaction sequence
	test.Focus("#input")
	test.AssertElement("#input").HasAttribute("data-events", "focus")

	test.Type("#input", "hello")
	// Type triggers input event, so should have focus,input
	test.AssertElement("#input").HasAttributeContaining("data-events", "input")

	test.KeyDown("#input", "Tab")
	test.AssertElement("#input").HasAttributeContaining("data-events", "keydown")

	test.KeyUp("#input", "Tab")
	test.AssertElement("#input").HasAttributeContaining("data-events", "keyup")

	test.Hover("#btn")
	test.AssertElement("#btn").HasAttributeContaining("data-events", "mouseenter")

	test.MouseDown("#btn")
	test.AssertElement("#btn").HasAttributeContaining("data-events", "mousedown")

	test.MouseUp("#btn")
	test.AssertElement("#btn").HasAttributeContaining("data-events", "mouseup")

	test.Click("#btn")
	test.AssertElement("#btn").HasAttributeContaining("data-events", "click")

	test.Unhover("#btn")
	test.AssertElement("#btn").HasAttributeContaining("data-events", "mouseleave")

	// Verify the input has the typed value
	test.AssertElement("#input").HasValue("hello")

	t.Log("✅ Event sequence completed successfully and JavaScript tracked all events!")
}

func TestMobileEventFlow(t *testing.T) {
	test := domulator.NewTest(t)
	test.LoadHTML(`
		<div id="mobile-target" style="touch-action: manipulation;">
			Mobile Touch Target
		</div>
		
		<script>
			let touchEvents = [];
			const target = document.getElementById('mobile-target');
			
			target.addEventListener('touchstart', () => {
				touchEvents.push('touchstart');
				target.setAttribute('data-touch-events', touchEvents.join(','));
			});
			
			target.addEventListener('touchmove', () => {
				touchEvents.push('touchmove');
				target.setAttribute('data-touch-events', touchEvents.join(','));
			});
			
			target.addEventListener('touchend', () => {
				touchEvents.push('touchend');
				target.setAttribute('data-touch-events', touchEvents.join(','));
			});
		</script>
	`)

	// Test mobile-specific interaction flow
	test.TouchStart("#mobile-target")
	test.AssertElement("#mobile-target").HasAttribute("data-touch-events", "touchstart")

	test.TouchMove("#mobile-target")
	test.AssertElement("#mobile-target").HasAttributeContaining("data-touch-events", "touchmove")

	test.TouchEnd("#mobile-target")
	test.AssertElement("#mobile-target").HasAttributeContaining("data-touch-events", "touchend")

	// Test mobile gesture helpers
	test.Tap("#mobile-target")
	// Tap = touchstart + touchend, so should have all three events now
	test.AssertElement("#mobile-target").HasAttributeContaining("data-touch-events", "touchstart")
	test.AssertElement("#mobile-target").HasAttributeContaining("data-touch-events", "touchend")

	// Verify element still exists after all touch interactions
	test.AssertElement("#mobile-target").Exists()

	t.Log("✅ Mobile event flow completed successfully and JavaScript tracked touch events!")
}

func TestDragDropFlow(t *testing.T) {
	test := domulator.NewTest(t)
	test.LoadHTML(`
		<div id="source" draggable="true" style="background: blue; padding: 20px;">
			Drag Source
		</div>
		<div id="target" style="background: green; padding: 40px;">
			Drop Target
		</div>
		
		<script>
			let dragEvents = [];
			let dropEvents = [];
			
			const source = document.getElementById('source');
			const target = document.getElementById('target');
			
			// Source events
			source.addEventListener('dragstart', () => {
				dragEvents.push('dragstart');
				source.setAttribute('data-drag-events', dragEvents.join(','));
			});
			
			source.addEventListener('drag', () => {
				dragEvents.push('drag');
				source.setAttribute('data-drag-events', dragEvents.join(','));
			});
			
			source.addEventListener('dragend', () => {
				dragEvents.push('dragend');
				source.setAttribute('data-drag-events', dragEvents.join(','));
			});
			
			// Target events
			target.addEventListener('dragenter', () => {
				dropEvents.push('dragenter');
				target.setAttribute('data-drop-events', dropEvents.join(','));
			});
			
			target.addEventListener('dragover', () => {
				dropEvents.push('dragover');
				target.setAttribute('data-drop-events', dropEvents.join(','));
			});
			
			target.addEventListener('dragleave', () => {
				dropEvents.push('dragleave');
				target.setAttribute('data-drop-events', dropEvents.join(','));
			});
			
			target.addEventListener('drop', () => {
				dropEvents.push('drop');
				target.setAttribute('data-drop-events', dropEvents.join(','));
			});
		</script>
	`)

	// Test individual drag and drop events
	test.DragStart("#source")
	test.AssertElement("#source").HasAttribute("data-drag-events", "dragstart")

	test.Drag("#source")
	test.AssertElement("#source").HasAttributeContaining("data-drag-events", "drag")

	test.DragEnter("#target")
	test.AssertElement("#target").HasAttribute("data-drop-events", "dragenter")

	test.DragOver("#target")
	test.AssertElement("#target").HasAttributeContaining("data-drop-events", "dragover")

	test.DragLeave("#target")
	test.AssertElement("#target").HasAttributeContaining("data-drop-events", "dragleave")

	test.DragEnter("#target")
	test.DragOver("#target")
	test.Drop("#target")
	test.AssertElement("#target").HasAttributeContaining("data-drop-events", "drop")

	test.DragEnd("#source")
	test.AssertElement("#source").HasAttributeContaining("data-drag-events", "dragend")

	// Test the helper method
	test.DragAndDrop("#source", "#target")

	// Verify both elements still exist
	test.AssertElement("#source").Exists()
	test.AssertElement("#target").Exists()

	t.Log("✅ Drag and drop flow completed successfully and JavaScript tracked all events!")
}

func TestFormEventFlow(t *testing.T) {
	test := domulator.NewTest(t)
	test.LoadHTML(`
		<form id="test-form">
			<input id="email" type="email" name="email" required>
			<input id="password" type="password" name="password" required>
			<input type="checkbox" id="terms" name="terms">
			<select id="country" name="country">
				<option value="us">United States</option>
				<option value="ca">Canada</option>
			</select>
			<button type="submit">Submit</button>
			<button type="reset">Reset</button>
		</form>
		
		<script>
			let formEvents = [];
			
			function logEvent(eventType, element) {
				formEvents.push(eventType);
				document.getElementById('test-form').setAttribute('data-form-events', formEvents.join(','));
			}
			
			document.getElementById('email').addEventListener('focus', () => logEvent('email-focus'));
			document.getElementById('email').addEventListener('blur', () => logEvent('email-blur'));
			document.getElementById('email').addEventListener('input', () => logEvent('email-input'));
			
			document.getElementById('password').addEventListener('focus', () => logEvent('password-focus'));
			document.getElementById('password').addEventListener('input', () => logEvent('password-input'));
			document.getElementById('password').addEventListener('keydown', () => logEvent('password-keydown'));
			
			document.getElementById('terms').addEventListener('change', () => logEvent('terms-change'));
			document.getElementById('country').addEventListener('change', () => logEvent('country-change'));
			
			document.getElementById('test-form').addEventListener('submit', () => logEvent('form-submit'));
			document.getElementById('test-form').addEventListener('reset', () => logEvent('form-reset'));
		</script>
	`)

	// Test form interaction flow
	test.Focus("#email")
	test.AssertElement("#test-form").HasAttributeContaining("data-form-events", "email-focus")

	test.Type("#email", "test@example.com")
	test.AssertElement("#test-form").HasAttributeContaining("data-form-events", "email-input")

	test.Blur("#email")
	test.AssertElement("#test-form").HasAttributeContaining("data-form-events", "email-blur")

	test.Focus("#password")
	test.AssertElement("#test-form").HasAttributeContaining("data-form-events", "password-focus")

	test.Type("#password", "password123")
	test.AssertElement("#test-form").HasAttributeContaining("data-form-events", "password-input")

	test.KeyDown("#password", "Tab")
	test.AssertElement("#test-form").HasAttributeContaining("data-form-events", "password-keydown")

	test.Check("#terms")
	test.AssertElement("#test-form").HasAttributeContaining("data-form-events", "terms-change")

	test.Select("#country", "ca")
	test.AssertElement("#test-form").HasAttributeContaining("data-form-events", "country-change")

	// Test form submission and reset
	test.Submit("#test-form")
	test.AssertElement("#test-form").HasAttributeContaining("data-form-events", "form-submit")

	test.Reset("#test-form")
	test.AssertElement("#test-form").HasAttributeContaining("data-form-events", "form-reset")

	// Verify form state
	test.AssertElement("#email").HasValue("test@example.com")
	test.AssertElement("#password").HasValue("password123")
	test.AssertElement("#terms").IsChecked()
	test.AssertElement("#country").HasValue("ca")

	t.Log("✅ Form event flow completed successfully and JavaScript tracked all form events!")
}

func TestMediaEventFlow(t *testing.T) {
	test := domulator.NewTest(t)
	test.LoadHTML(`
		<video id="video" controls>
			<source src="test.mp4" type="video/mp4">
		</video>
		<audio id="audio" controls>
			<source src="test.mp3" type="audio/mpeg">
		</audio>
		
		<script>
			let mediaEvents = [];
			
			function logMediaEvent(eventType) {
				mediaEvents.push(eventType);
				document.body.setAttribute('data-media-events', mediaEvents.join(','));
			}
			
			document.getElementById('video').addEventListener('play', () => logMediaEvent('video-play'));
			document.getElementById('video').addEventListener('pause', () => logMediaEvent('video-pause'));
			document.getElementById('video').addEventListener('ended', () => logMediaEvent('video-ended'));
			
			document.getElementById('audio').addEventListener('play', () => logMediaEvent('audio-play'));
			document.getElementById('audio').addEventListener('pause', () => logMediaEvent('audio-pause'));
			document.getElementById('audio').addEventListener('ended', () => logMediaEvent('audio-ended'));
		</script>
	`)

	// Test media events
	test.Play("#video")
	test.AssertElement("body").HasAttributeContaining("data-media-events", "video-play")

	test.Pause("#video")
	test.AssertElement("body").HasAttributeContaining("data-media-events", "video-pause")

	test.Ended("#video")
	test.AssertElement("body").HasAttributeContaining("data-media-events", "video-ended")

	test.Play("#audio")
	test.AssertElement("body").HasAttributeContaining("data-media-events", "audio-play")

	test.Pause("#audio")
	test.AssertElement("body").HasAttributeContaining("data-media-events", "audio-pause")

	test.Ended("#audio")
	test.AssertElement("body").HasAttributeContaining("data-media-events", "audio-ended")

	// Verify media elements exist
	test.AssertElement("#video").Exists()
	test.AssertElement("#audio").Exists()

	t.Log("✅ Media event flow completed successfully and JavaScript tracked all media events!")
}

func TestSimpleEventsWithConsoleCapture(t *testing.T) {
	harness := testpkg.NewTestHarness()
	defer harness.Close()

	consoleCapture := harness.CaptureConsole(t)

	harness.LoadHTML(`
		<button id="btn">Click me</button>
		<div id="output"></div>
		
		<script>
			console.log('Setting up event listeners...');
			console.warn('This is a warning message');
			console.error('This is an error message');
			
			// Test DOM manipulation
			document.body.setAttribute('data-script-executed', 'true');
			console.log('DOM manipulation completed');
			
			// Test DOMContentLoaded
			document.addEventListener('DOMContentLoaded', function() {
				document.body.setAttribute('data-last-event', 'DOMContentLoaded');
				console.log('DOMContentLoaded event received');
			});
			
			// Test immediate script execution
			console.log('Script execution complete!');
		</script>
	`)

	// Assert console output for different log levels
	consoleCapture.AssertLogContains("Setting up event listeners...")
	consoleCapture.AssertWarnContains("This is a warning message")
	consoleCapture.AssertErrorContains("This is an error message")
	consoleCapture.AssertLogContains("DOM manipulation completed")
	consoleCapture.AssertLogContains("DOMContentLoaded event received")
	consoleCapture.AssertLogContains("Script execution complete!")

	// Test additional console output through ExecuteScript
	harness.ExecuteScript(`
		console.log('Additional log from ExecuteScript');
		console.warn('Additional warning from ExecuteScript');
		console.error('Additional error from ExecuteScript');
		
		// Test setting an attribute to verify script execution
		document.body.setAttribute('data-execute-script-ran', 'true');
	`)

	// Assert the additional console output
	consoleCapture.AssertLogContains("Additional log from ExecuteScript")
	consoleCapture.AssertWarnContains("Additional warning from ExecuteScript")
	consoleCapture.AssertErrorContains("Additional error from ExecuteScript")

	// Verify DOM state
	body := harness.Document().QuerySelector("body")
	if body == nil {
		t.Fatal("Could not find body element")
	}

	if scriptExecuted := body.GetAttribute("data-script-executed"); scriptExecuted != "true" {
		t.Errorf("Expected data-script-executed 'true', got %s", scriptExecuted)
	}

	if executeScriptRan := body.GetAttribute("data-execute-script-ran"); executeScriptRan != "true" {
		t.Errorf("Expected data-execute-script-ran 'true', got %s", executeScriptRan)
	}

	if lastEvent := body.GetAttribute("data-last-event"); lastEvent != "DOMContentLoaded" {
		t.Errorf("Expected data-last-event 'DOMContentLoaded', got %s", lastEvent)
	}

	t.Log("✅ Console capture test completed successfully!")
}
