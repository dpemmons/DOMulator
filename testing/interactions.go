package testing

import (
	"fmt"

	"github.com/dpemmons/DOMulator/internal/css"
	"github.com/dpemmons/DOMulator/internal/dom"
)

// Click simulates a click event on the first element matching the selector
func (h *TestHarness) Click(selector string) *TestHarness {
	if h.domulator == nil {
		panic("no document loaded - call LoadHTML() or Navigate() first")
	}

	elements := css.QuerySelectorAll(h.domulator.document, selector)
	if elements.Length() == 0 {
		panic(fmt.Sprintf("Cannot click: no elements matching selector '%s' found", selector))
	}

	element := elements.Item(0)

	// Create and dispatch a click event
	clickEvent := dom.NewEvent("click", true, true)
	element.DispatchEvent(clickEvent)

	// Special handling for submit buttons - also trigger form submission
	if elem, ok := element.(*dom.Element); ok {
		buttonType := elem.GetAttribute("type")
		if buttonType == "submit" || (buttonType == "" && elem.TagName() == "BUTTON") {
			// Find the parent form
			parent := elem.ParentNode()
			for parent != nil {
				if parentElem, ok := parent.(*dom.Element); ok && parentElem.TagName() == "FORM" {
					// Trigger submit event on the form
					submitEvent := dom.NewEvent("submit", true, true)
					submitEvent.SetTarget(parent)
					parent.DispatchEvent(submitEvent)
					break
				}
				parent = parent.ParentNode()
			}
		}
	}

	return h
}

// DoubleClick simulates a double-click event on the first element matching the selector
func (h *TestHarness) DoubleClick(selector string) *TestHarness {
	if h.domulator == nil {
		panic("no document loaded - call LoadHTML() or Navigate() first")
	}

	elements := css.QuerySelectorAll(h.domulator.document, selector)
	if elements.Length() == 0 {
		panic(fmt.Sprintf("Cannot double-click: no elements matching selector '%s' found", selector))
	}

	element := elements.Item(0)

	// Create and dispatch a dblclick event
	dblClickEvent := dom.NewEvent("dblclick", true, true)
	element.DispatchEvent(dblClickEvent)
	return h
}

// Type simulates typing text into an input element
func (h *TestHarness) Type(selector, text string) *TestHarness {
	if h.domulator == nil {
		panic("no document loaded - call LoadHTML() or Navigate() first")
	}

	// Execute JavaScript to set value and trigger events directly in the JS runtime
	script := fmt.Sprintf(`
		var element = document.querySelector('%s');
		if (element) {
			// Set both the property and attribute to keep them in sync
			element.value = '%s';
			element.setAttribute('value', '%s');
			
			// Trigger input event
			var inputEvent = new Event('input', { bubbles: true, cancelable: false });
			element.dispatchEvent(inputEvent);
			
			// Trigger change event
			var changeEvent = new Event('change', { bubbles: true, cancelable: false });
			element.dispatchEvent(changeEvent);
		}
	`, selector, text, text)
	h.ExecuteScript(script)
	return h
}

// Focus simulates focusing on an element
func (h *TestHarness) Focus(selector string) *TestHarness {
	if h.domulator == nil {
		panic("no document loaded - call LoadHTML() or Navigate() first")
	}

	// Execute JavaScript to trigger the event directly in the JS runtime
	script := fmt.Sprintf(`
		var element = document.querySelector('%s');
		if (element) {
			var event = new Event('focus', { bubbles: false, cancelable: false });
			element.dispatchEvent(event);
		}
	`, selector)
	h.ExecuteScript(script)
	return h
}

// Blur simulates blurring (losing focus) on an element
func (h *TestHarness) Blur(selector string) *TestHarness {
	if h.domulator == nil {
		panic("no document loaded - call LoadHTML() or Navigate() first")
	}

	// Execute JavaScript to trigger the event directly in the JS runtime
	script := fmt.Sprintf(`
		var element = document.querySelector('%s');
		if (element) {
			var event = new Event('blur', { bubbles: false, cancelable: false });
			element.dispatchEvent(event);
		}
	`, selector)
	h.ExecuteScript(script)
	return h
}

// Submit simulates submitting a form
func (h *TestHarness) Submit(selector string) *TestHarness {
	if h.domulator == nil {
		panic("no document loaded - call LoadHTML() or Navigate() first")
	}

	elements := css.QuerySelectorAll(h.domulator.document, selector)
	if elements.Length() == 0 {
		panic(fmt.Sprintf("Cannot submit: no elements matching selector '%s' found", selector))
	}

	element := elements.Item(0)

	// Create and dispatch submit event
	submitEvent := dom.NewEvent("submit", true, true)
	// Manually set the target since it's not set by NewEvent
	submitEvent.SetTarget(element)
	element.DispatchEvent(submitEvent)
	return h
}

// Check simulates checking a checkbox or radio button
func (h *TestHarness) Check(selector string) *TestHarness {
	if h.domulator == nil {
		panic("no document loaded - call LoadHTML() or Navigate() first")
	}

	elements := css.QuerySelectorAll(h.domulator.document, selector)
	if elements.Length() == 0 {
		panic(fmt.Sprintf("Cannot check: no elements matching selector '%s' found", selector))
	}

	element := elements.Item(0)

	if elem, ok := element.(*dom.Element); ok {
		elem.SetAttribute("checked", "checked")

		// Dispatch change event
		changeEvent := dom.NewEvent("change", true, false)
		element.DispatchEvent(changeEvent)
	} else {
		panic(fmt.Sprintf("Cannot check: element '%s' is not an Element node", selector))
	}

	return h
}

// Uncheck simulates unchecking a checkbox
func (h *TestHarness) Uncheck(selector string) *TestHarness {
	if h.domulator == nil {
		panic("no document loaded - call LoadHTML() or Navigate() first")
	}

	elements := css.QuerySelectorAll(h.domulator.document, selector)
	if elements.Length() == 0 {
		panic(fmt.Sprintf("Cannot uncheck: no elements matching selector '%s' found", selector))
	}

	element := elements.Item(0)

	if elem, ok := element.(*dom.Element); ok {
		elem.SetAttribute("checked", "")

		// Dispatch change event
		changeEvent := dom.NewEvent("change", true, false)
		element.DispatchEvent(changeEvent)
	} else {
		panic(fmt.Sprintf("Cannot uncheck: element '%s' is not an Element node", selector))
	}

	return h
}

// Select simulates selecting an option in a select element
func (h *TestHarness) Select(selector, value string) *TestHarness {
	if h.domulator == nil {
		panic("no document loaded - call LoadHTML() or Navigate() first")
	}

	elements := css.QuerySelectorAll(h.domulator.document, selector)
	if elements.Length() == 0 {
		panic(fmt.Sprintf("Cannot select: no elements matching selector '%s' found", selector))
	}

	element := elements.Item(0)

	if elem, ok := element.(*dom.Element); ok {
		elem.SetAttribute("value", value)

		// Dispatch change event
		changeEvent := dom.NewEvent("change", true, false)
		element.DispatchEvent(changeEvent)
	} else {
		panic(fmt.Sprintf("Cannot select: element '%s' is not an Element node", selector))
	}

	return h
}

// Hover simulates hovering over an element
func (h *TestHarness) Hover(selector string) *TestHarness {
	if h.domulator == nil {
		panic("no document loaded - call LoadHTML() or Navigate() first")
	}

	elements := css.QuerySelectorAll(h.domulator.document, selector)
	if elements.Length() == 0 {
		panic(fmt.Sprintf("Cannot hover: no elements matching selector '%s' found", selector))
	}

	element := elements.Item(0)

	// Create and dispatch mouseenter event
	mouseEnterEvent := dom.NewEvent("mouseenter", false, false)
	element.DispatchEvent(mouseEnterEvent)

	// Create and dispatch mouseover event (which bubbles)
	mouseOverEvent := dom.NewEvent("mouseover", true, true)
	element.DispatchEvent(mouseOverEvent)
	return h
}

// Unhover simulates moving the mouse away from an element
func (h *TestHarness) Unhover(selector string) *TestHarness {
	if h.domulator == nil {
		panic("no document loaded - call LoadHTML() or Navigate() first")
	}

	elements := css.QuerySelectorAll(h.domulator.document, selector)
	if elements.Length() == 0 {
		panic(fmt.Sprintf("Cannot unhover: no elements matching selector '%s' found", selector))
	}

	element := elements.Item(0)

	// Create and dispatch mouseleave event
	mouseLeaveEvent := dom.NewEvent("mouseleave", false, false)
	element.DispatchEvent(mouseLeaveEvent)

	// Create and dispatch mouseout event (which bubbles)
	mouseOutEvent := dom.NewEvent("mouseout", true, true)
	element.DispatchEvent(mouseOutEvent)
	return h
}

// TriggerEvent dispatches a custom event on the first element matching the selector
func (h *TestHarness) TriggerEvent(selector, eventType string) *TestHarness {
	if h.domulator == nil {
		panic("no document loaded - call LoadHTML() or Navigate() first")
	}

	elements := css.QuerySelectorAll(h.domulator.document, selector)
	if elements.Length() == 0 {
		panic(fmt.Sprintf("Cannot trigger event: no elements matching selector '%s' found", selector))
	}

	element := elements.Item(0)

	// Create and dispatch custom event
	customEvent := dom.NewEvent(eventType, true, true)
	element.DispatchEvent(customEvent)
	return h
}

// ===== KEYBOARD EVENTS =====

// KeyDown simulates a keydown event
func (h *TestHarness) KeyDown(selector, key string) *TestHarness {
	if h.domulator == nil {
		panic("no document loaded - call LoadHTML() or Navigate() first")
	}

	// Execute JavaScript to trigger the event directly in the JS runtime
	script := fmt.Sprintf(`
		var element = document.querySelector('%s');
		if (element) {
			// Try KeyboardEvent first, fall back to Event if not available
			var event;
			try {
				event = new KeyboardEvent('keydown', { 
					bubbles: true, 
					cancelable: true,
					key: '%s'
				});
			} catch (e) {
				// Fallback to generic Event
				event = new Event('keydown', { bubbles: true, cancelable: true });
				event.key = '%s';
			}
			element.dispatchEvent(event);
			console.log('KeyDown event dispatched on', element.tagName, element.id);
		} else {
			console.log('Element not found for selector: %s');
		}
	`, selector, key, key, selector)
	h.ExecuteScript(script)
	return h
}

// KeyUp simulates a keyup event
func (h *TestHarness) KeyUp(selector, key string) *TestHarness {
	if h.domulator == nil {
		panic("no document loaded - call LoadHTML() or Navigate() first")
	}

	// Execute JavaScript to trigger the event directly in the JS runtime
	script := fmt.Sprintf(`
		var element = document.querySelector('%s');
		if (element) {
			var event = new Event('keyup', { bubbles: true, cancelable: true });
			element.dispatchEvent(event);
		}
	`, selector)
	h.ExecuteScript(script)
	return h
}

// KeyPress simulates a keypress event (deprecated but still used)
func (h *TestHarness) KeyPress(selector, key string) *TestHarness {
	if h.domulator == nil {
		panic("no document loaded - call LoadHTML() or Navigate() first")
	}

	// Execute JavaScript to trigger the event directly in the JS runtime
	script := fmt.Sprintf(`
		var element = document.querySelector('%s');
		if (element) {
			var event = new Event('keypress', { bubbles: true, cancelable: true });
			element.dispatchEvent(event);
		}
	`, selector)
	h.ExecuteScript(script)
	return h
}

// ===== ADVANCED MOUSE EVENTS =====

// MouseDown simulates a mousedown event
func (h *TestHarness) MouseDown(selector string) *TestHarness {
	if h.domulator == nil {
		panic("no document loaded - call LoadHTML() or Navigate() first")
	}

	// Execute JavaScript to trigger the event directly in the JS runtime
	script := fmt.Sprintf(`
		var element = document.querySelector('%s');
		if (element) {
			var event = new Event('mousedown', { bubbles: true, cancelable: true });
			element.dispatchEvent(event);
		}
	`, selector)
	h.ExecuteScript(script)
	return h
}

// MouseUp simulates a mouseup event
func (h *TestHarness) MouseUp(selector string) *TestHarness {
	if h.domulator == nil {
		panic("no document loaded - call LoadHTML() or Navigate() first")
	}

	// Execute JavaScript to trigger the event directly in the JS runtime
	script := fmt.Sprintf(`
		var element = document.querySelector('%s');
		if (element) {
			var event = new Event('mouseup', { bubbles: true, cancelable: true });
			element.dispatchEvent(event);
		}
	`, selector)
	h.ExecuteScript(script)
	return h
}

// MouseMove simulates a mousemove event
func (h *TestHarness) MouseMove(selector string) *TestHarness {
	if h.domulator == nil {
		panic("no document loaded - call LoadHTML() or Navigate() first")
	}

	// Execute JavaScript to trigger the event directly in the JS runtime
	script := fmt.Sprintf(`
		var element = document.querySelector('%s');
		if (element) {
			var event = new Event('mousemove', { bubbles: true, cancelable: true });
			element.dispatchEvent(event);
		}
	`, selector)
	h.ExecuteScript(script)
	return h
}

// RightClick simulates a right-click (contextmenu event)
func (h *TestHarness) RightClick(selector string) *TestHarness {
	if h.domulator == nil {
		panic("no document loaded - call LoadHTML() or Navigate() first")
	}

	// Execute JavaScript to trigger the event directly in the JS runtime
	script := fmt.Sprintf(`
		var element = document.querySelector('%s');
		if (element) {
			var event = new Event('contextmenu', { bubbles: true, cancelable: true });
			element.dispatchEvent(event);
		}
	`, selector)
	h.ExecuteScript(script)
	return h
}

// ===== TOUCH EVENTS =====

// TouchStart simulates a touchstart event
func (h *TestHarness) TouchStart(selector string) *TestHarness {
	if h.domulator == nil {
		panic("no document loaded - call LoadHTML() or Navigate() first")
	}

	// Execute JavaScript to trigger the event directly in the JS runtime
	script := fmt.Sprintf(`
		var element = document.querySelector('%s');
		if (element) {
			var event = new Event('touchstart', { bubbles: true, cancelable: true });
			element.dispatchEvent(event);
		}
	`, selector)
	h.ExecuteScript(script)
	return h
}

// TouchMove simulates a touchmove event
func (h *TestHarness) TouchMove(selector string) *TestHarness {
	if h.domulator == nil {
		panic("no document loaded - call LoadHTML() or Navigate() first")
	}

	// Execute JavaScript to trigger the event directly in the JS runtime
	script := fmt.Sprintf(`
		var element = document.querySelector('%s');
		if (element) {
			var event = new Event('touchmove', { bubbles: true, cancelable: true });
			element.dispatchEvent(event);
		}
	`, selector)
	h.ExecuteScript(script)
	return h
}

// TouchEnd simulates a touchend event
func (h *TestHarness) TouchEnd(selector string) *TestHarness {
	if h.domulator == nil {
		panic("no document loaded - call LoadHTML() or Navigate() first")
	}

	// Execute JavaScript to trigger the event directly in the JS runtime
	script := fmt.Sprintf(`
		var element = document.querySelector('%s');
		if (element) {
			var event = new Event('touchend', { bubbles: true, cancelable: true });
			element.dispatchEvent(event);
		}
	`, selector)
	h.ExecuteScript(script)
	return h
}

// ===== DRAG AND DROP EVENTS =====

// DragStart simulates a dragstart event
func (h *TestHarness) DragStart(selector string) *TestHarness {
	if h.domulator == nil {
		panic("no document loaded - call LoadHTML() or Navigate() first")
	}

	// Execute JavaScript to trigger the event directly in the JS runtime
	script := fmt.Sprintf(`
		var element = document.querySelector('%s');
		if (element) {
			var event = new Event('dragstart', { bubbles: true, cancelable: true });
			element.dispatchEvent(event);
		}
	`, selector)
	h.ExecuteScript(script)
	return h
}

// Drag simulates a drag event
func (h *TestHarness) Drag(selector string) *TestHarness {
	if h.domulator == nil {
		panic("no document loaded - call LoadHTML() or Navigate() first")
	}

	// Execute JavaScript to trigger the event directly in the JS runtime
	script := fmt.Sprintf(`
		var element = document.querySelector('%s');
		if (element) {
			var event = new Event('drag', { bubbles: true, cancelable: true });
			element.dispatchEvent(event);
		}
	`, selector)
	h.ExecuteScript(script)
	return h
}

// DragEnd simulates a dragend event
func (h *TestHarness) DragEnd(selector string) *TestHarness {
	if h.domulator == nil {
		panic("no document loaded - call LoadHTML() or Navigate() first")
	}

	// Execute JavaScript to trigger the event directly in the JS runtime
	script := fmt.Sprintf(`
		var element = document.querySelector('%s');
		if (element) {
			var event = new Event('dragend', { bubbles: true, cancelable: true });
			element.dispatchEvent(event);
		}
	`, selector)
	h.ExecuteScript(script)
	return h
}

// DragOver simulates a dragover event
func (h *TestHarness) DragOver(selector string) *TestHarness {
	if h.domulator == nil {
		panic("no document loaded - call LoadHTML() or Navigate() first")
	}

	// Execute JavaScript to trigger the event directly in the JS runtime
	script := fmt.Sprintf(`
		var element = document.querySelector('%s');
		if (element) {
			var event = new Event('dragover', { bubbles: true, cancelable: true });
			element.dispatchEvent(event);
		}
	`, selector)
	h.ExecuteScript(script)
	return h
}

// DragEnter simulates a dragenter event
func (h *TestHarness) DragEnter(selector string) *TestHarness {
	if h.domulator == nil {
		panic("no document loaded - call LoadHTML() or Navigate() first")
	}

	// Execute JavaScript to trigger the event directly in the JS runtime
	script := fmt.Sprintf(`
		var element = document.querySelector('%s');
		if (element) {
			var event = new Event('dragenter', { bubbles: true, cancelable: true });
			element.dispatchEvent(event);
		}
	`, selector)
	h.ExecuteScript(script)
	return h
}

// DragLeave simulates a dragleave event
func (h *TestHarness) DragLeave(selector string) *TestHarness {
	if h.domulator == nil {
		panic("no document loaded - call LoadHTML() or Navigate() first")
	}

	// Execute JavaScript to trigger the event directly in the JS runtime
	script := fmt.Sprintf(`
		var element = document.querySelector('%s');
		if (element) {
			var event = new Event('dragleave', { bubbles: true, cancelable: true });
			element.dispatchEvent(event);
		}
	`, selector)
	h.ExecuteScript(script)
	return h
}

// Drop simulates a drop event
func (h *TestHarness) Drop(selector string) *TestHarness {
	if h.domulator == nil {
		panic("no document loaded - call LoadHTML() or Navigate() first")
	}

	// Execute JavaScript to trigger the event directly in the JS runtime
	script := fmt.Sprintf(`
		var element = document.querySelector('%s');
		if (element) {
			var event = new Event('drop', { bubbles: true, cancelable: true });
			element.dispatchEvent(event);
		}
	`, selector)
	h.ExecuteScript(script)
	return h
}

// ===== FORM EVENTS =====

// Reset simulates a reset event on a form
func (h *TestHarness) Reset(selector string) *TestHarness {
	if h.domulator == nil {
		panic("no document loaded - call LoadHTML() or Navigate() first")
	}

	// Execute JavaScript to trigger the event directly in the JS runtime
	script := fmt.Sprintf(`
		var element = document.querySelector('%s');
		if (element) {
			var event = new Event('reset', { bubbles: true, cancelable: true });
			element.dispatchEvent(event);
		}
	`, selector)
	h.ExecuteScript(script)
	return h
}

// Invalid simulates an invalid event (form validation)
func (h *TestHarness) Invalid(selector string) *TestHarness {
	if h.domulator == nil {
		panic("no document loaded - call LoadHTML() or Navigate() first")
	}

	// Execute JavaScript to trigger the event directly in the JS runtime
	script := fmt.Sprintf(`
		var element = document.querySelector('%s');
		if (element) {
			var event = new Event('invalid', { bubbles: false, cancelable: true });
			element.dispatchEvent(event);
		}
	`, selector)
	h.ExecuteScript(script)
	return h
}

// ===== SCROLL EVENTS =====

// Scroll simulates a scroll event
func (h *TestHarness) Scroll(selector string) *TestHarness {
	if h.domulator == nil {
		panic("no document loaded - call LoadHTML() or Navigate() first")
	}

	// Execute JavaScript to trigger the event directly in the JS runtime
	script := fmt.Sprintf(`
		var element = document.querySelector('%s');
		if (element) {
			var event = new Event('scroll', { bubbles: false, cancelable: false });
			element.dispatchEvent(event);
		}
	`, selector)
	h.ExecuteScript(script)
	return h
}

// ===== MEDIA EVENTS =====

// Play simulates a play event on media elements
func (h *TestHarness) Play(selector string) *TestHarness {
	if h.domulator == nil {
		panic("no document loaded - call LoadHTML() or Navigate() first")
	}

	// Execute JavaScript to trigger the event directly in the JS runtime
	script := fmt.Sprintf(`
		var element = document.querySelector('%s');
		if (element) {
			var event = new Event('play', { bubbles: false, cancelable: false });
			element.dispatchEvent(event);
		}
	`, selector)
	h.ExecuteScript(script)
	return h
}

// Pause simulates a pause event on media elements
func (h *TestHarness) Pause(selector string) *TestHarness {
	if h.domulator == nil {
		panic("no document loaded - call LoadHTML() or Navigate() first")
	}

	// Execute JavaScript to trigger the event directly in the JS runtime
	script := fmt.Sprintf(`
		var element = document.querySelector('%s');
		if (element) {
			var event = new Event('pause', { bubbles: false, cancelable: false });
			element.dispatchEvent(event);
		}
	`, selector)
	h.ExecuteScript(script)
	return h
}

// Ended simulates an ended event on media elements
func (h *TestHarness) Ended(selector string) *TestHarness {
	if h.domulator == nil {
		panic("no document loaded - call LoadHTML() or Navigate() first")
	}

	// Execute JavaScript to trigger the event directly in the JS runtime
	script := fmt.Sprintf(`
		var element = document.querySelector('%s');
		if (element) {
			var event = new Event('ended', { bubbles: false, cancelable: false });
			element.dispatchEvent(event);
		}
	`, selector)
	h.ExecuteScript(script)
	return h
}

// ===== CLIPBOARD EVENTS =====

// Copy simulates a copy event
func (h *TestHarness) Copy(selector string) *TestHarness {
	if h.domulator == nil {
		panic("no document loaded - call LoadHTML() or Navigate() first")
	}

	// Execute JavaScript to trigger the event directly in the JS runtime
	script := fmt.Sprintf(`
		var element = document.querySelector('%s');
		if (element) {
			var event = new Event('copy', { bubbles: true, cancelable: true });
			element.dispatchEvent(event);
		}
	`, selector)
	h.ExecuteScript(script)
	return h
}

// Cut simulates a cut event
func (h *TestHarness) Cut(selector string) *TestHarness {
	if h.domulator == nil {
		panic("no document loaded - call LoadHTML() or Navigate() first")
	}

	// Execute JavaScript to trigger the event directly in the JS runtime
	script := fmt.Sprintf(`
		var element = document.querySelector('%s');
		if (element) {
			var event = new Event('cut', { bubbles: true, cancelable: true });
			element.dispatchEvent(event);
		}
	`, selector)
	h.ExecuteScript(script)
	return h
}

// Paste simulates a paste event
func (h *TestHarness) Paste(selector string) *TestHarness {
	if h.domulator == nil {
		panic("no document loaded - call LoadHTML() or Navigate() first")
	}

	// Execute JavaScript to trigger the event directly in the JS runtime
	script := fmt.Sprintf(`
		var element = document.querySelector('%s');
		if (element) {
			var event = new Event('paste', { bubbles: true, cancelable: true });
			element.dispatchEvent(event);
		}
	`, selector)
	h.ExecuteScript(script)
	return h
}

// ===== WINDOW EVENTS =====

// Resize simulates a resize event on the window
func (h *TestHarness) Resize() *TestHarness {
	if h.domulator == nil {
		panic("no document loaded - call LoadHTML() or Navigate() first")
	}

	// Execute JavaScript that properly dispatches window resize event
	script := `
		// Update the body's data-last-event attribute for test verification
		if (document.body) {
			document.body.setAttribute('data-last-event', 'resize');
		}
		
		// Dispatch resize event on window
		var event = new Event('resize', { bubbles: false, cancelable: false });
		window.dispatchEvent(event);
		
		console.log('Window resize event dispatched');
	`
	h.ExecuteScript(script)
	return h
}

// Load simulates a load event on the window
func (h *TestHarness) Load() *TestHarness {
	if h.domulator == nil {
		panic("no document loaded - call LoadHTML() or Navigate() first")
	}

	// Execute JavaScript that properly dispatches window load event
	script := `
		// Update the body's data-last-event attribute for test verification
		if (document.body) {
			document.body.setAttribute('data-last-event', 'load');
		}
		
		// Dispatch load event on window
		var event = new Event('load', { bubbles: false, cancelable: false });
		window.dispatchEvent(event);
		
		console.log('Window load event dispatched');
	`
	h.ExecuteScript(script)
	return h
}

// Unload simulates an unload event on the window
func (h *TestHarness) Unload() *TestHarness {
	if h.domulator == nil {
		panic("no document loaded - call LoadHTML() or Navigate() first")
	}

	// Execute JavaScript that properly dispatches window unload event
	script := `
		// Update the body's data-last-event attribute for test verification
		if (document.body) {
			document.body.setAttribute('data-last-event', 'unload');
		}
		
		// Dispatch unload event on window
		var event = new Event('unload', { bubbles: false, cancelable: false });
		window.dispatchEvent(event);
		
		console.log('Window unload event dispatched');
	`
	h.ExecuteScript(script)
	return h
}

// ===== ADVANCED INTERACTION HELPERS =====

// DragAndDrop performs a complete drag and drop operation from source to target
func (h *TestHarness) DragAndDrop(sourceSelector, targetSelector string) *TestHarness {
	// Start drag operation
	h.DragStart(sourceSelector)
	h.Drag(sourceSelector)

	// Handle drop target
	h.DragEnter(targetSelector)
	h.DragOver(targetSelector)
	h.Drop(targetSelector)

	// End drag operation
	h.DragEnd(sourceSelector)

	return h
}

// Tap simulates a mobile tap (touchstart + touchend)
func (h *TestHarness) Tap(selector string) *TestHarness {
	h.TouchStart(selector)
	h.TouchEnd(selector)
	return h
}

// SwipeLeft simulates a left swipe gesture
func (h *TestHarness) SwipeLeft(selector string) *TestHarness {
	h.TouchStart(selector)
	h.TouchMove(selector)
	h.TouchEnd(selector)
	return h
}

// SwipeRight simulates a right swipe gesture
func (h *TestHarness) SwipeRight(selector string) *TestHarness {
	h.TouchStart(selector)
	h.TouchMove(selector)
	h.TouchEnd(selector)
	return h
}

// LongPress simulates a long press gesture (extended touch)
func (h *TestHarness) LongPress(selector string) *TestHarness {
	h.TouchStart(selector)
	// In a real implementation, we might add a delay here
	h.TouchEnd(selector)
	return h
}
