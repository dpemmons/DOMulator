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

	elements := css.QuerySelectorAll(h.domulator.document, selector)
	if elements.Length() == 0 {
		panic(fmt.Sprintf("Cannot type: no elements matching selector '%s' found", selector))
	}

	element := elements.Item(0)

	// Set the value attribute
	if elem, ok := element.(*dom.Element); ok {
		elem.SetAttribute("value", text)

		// Dispatch input event
		inputEvent := dom.NewEvent("input", true, false)
		element.DispatchEvent(inputEvent)

		// Dispatch change event
		changeEvent := dom.NewEvent("change", true, false)
		element.DispatchEvent(changeEvent)
	} else {
		panic(fmt.Sprintf("Cannot type: element '%s' is not an Element node", selector))
	}

	return h
}

// Focus simulates focusing on an element
func (h *TestHarness) Focus(selector string) *TestHarness {
	if h.domulator == nil {
		panic("no document loaded - call LoadHTML() or Navigate() first")
	}

	elements := css.QuerySelectorAll(h.domulator.document, selector)
	if elements.Length() == 0 {
		panic(fmt.Sprintf("Cannot focus: no elements matching selector '%s' found", selector))
	}

	element := elements.Item(0)

	// Create and dispatch focus event
	focusEvent := dom.NewEvent("focus", false, false)
	element.DispatchEvent(focusEvent)
	return h
}

// Blur simulates blurring (losing focus) on an element
func (h *TestHarness) Blur(selector string) *TestHarness {
	if h.domulator == nil {
		panic("no document loaded - call LoadHTML() or Navigate() first")
	}

	elements := css.QuerySelectorAll(h.domulator.document, selector)
	if elements.Length() == 0 {
		panic(fmt.Sprintf("Cannot blur: no elements matching selector '%s' found", selector))
	}

	element := elements.Item(0)

	// Create and dispatch blur event
	blurEvent := dom.NewEvent("blur", false, false)
	element.DispatchEvent(blurEvent)
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
