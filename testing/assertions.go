package testing

import (
	"fmt"
	"strings"

	"github.com/dpemmons/DOMulator/internal/css"
	"github.com/dpemmons/DOMulator/internal/dom"
)

// AssertElement creates an element assertion for the given CSS selector
func (h *TestHarness) AssertElement(selector string) *ElementAssertion {
	if h.domulator == nil {
		panic("no document loaded - call LoadHTML() or Navigate() first")
	}

	elements := css.QuerySelectorAll(h.domulator.document, selector)

	return &ElementAssertion{
		harness:  h,
		selector: selector,
		elements: elements.ToSlice(),
	}
}

// AssertDocument creates a document-level assertion
func (h *TestHarness) AssertDocument() *DocumentAssertion {
	if h.domulator == nil {
		panic("no document loaded - call LoadHTML() or Navigate() first")
	}

	return &DocumentAssertion{
		harness:  h,
		document: h.domulator.document,
	}
}

// ElementAssertion provides fluent assertions for DOM elements
type ElementAssertion struct {
	harness  *TestHarness
	selector string
	elements []dom.Node
}

// DocumentAssertion provides fluent assertions for the document
type DocumentAssertion struct {
	harness  *TestHarness
	document *dom.Document
}

// Exists asserts that at least one element matching the selector exists
func (e *ElementAssertion) Exists() *ElementAssertion {
	if len(e.elements) == 0 {
		panic(fmt.Sprintf("Expected element '%s' to exist, but found none", e.selector))
	}
	return e
}

// NotExists asserts that no elements matching the selector exist
func (e *ElementAssertion) NotExists() *ElementAssertion {
	if len(e.elements) > 0 {
		panic(fmt.Sprintf("Expected element '%s' to not exist, but found %d", e.selector, len(e.elements)))
	}
	return e
}

// Count asserts that exactly the specified number of elements exist
func (e *ElementAssertion) Count(expected int) *ElementAssertion {
	actual := len(e.elements)
	if actual != expected {
		panic(fmt.Sprintf("Expected %d elements matching '%s', but found %d", expected, e.selector, actual))
	}
	return e
}

// First returns an assertion for the first matching element
func (e *ElementAssertion) First() *ElementAssertion {
	if len(e.elements) == 0 {
		panic(fmt.Sprintf("Cannot get first element: no elements matching '%s' found", e.selector))
	}
	return &ElementAssertion{
		harness:  e.harness,
		selector: e.selector + ":first",
		elements: []dom.Node{e.elements[0]},
	}
}

// Last returns an assertion for the last matching element
func (e *ElementAssertion) Last() *ElementAssertion {
	if len(e.elements) == 0 {
		panic(fmt.Sprintf("Cannot get last element: no elements matching '%s' found", e.selector))
	}
	return &ElementAssertion{
		harness:  e.harness,
		selector: e.selector + ":last",
		elements: []dom.Node{e.elements[len(e.elements)-1]},
	}
}

// At returns an assertion for the element at the specified index
func (e *ElementAssertion) At(index int) *ElementAssertion {
	if index < 0 || index >= len(e.elements) {
		panic(fmt.Sprintf("Index %d out of range: found %d elements matching '%s'", index, len(e.elements), e.selector))
	}
	return &ElementAssertion{
		harness:  e.harness,
		selector: e.selector + fmt.Sprintf("[%d]", index),
		elements: []dom.Node{e.elements[index]},
	}
}

// HasText asserts that the first element contains the specified text
func (e *ElementAssertion) HasText(expected string) *ElementAssertion {
	e.Exists() // Ensure element exists first

	element := e.elements[0]
	actual := getTextContent(element)

	if actual != expected {
		panic(fmt.Sprintf("Expected element '%s' to have text '%s', but got '%s'", e.selector, expected, actual))
	}
	return e
}

// ContainsText asserts that the first element contains the specified text as a substring
func (e *ElementAssertion) ContainsText(expected string) *ElementAssertion {
	e.Exists() // Ensure element exists first

	element := e.elements[0]
	actual := getTextContent(element)

	if !strings.Contains(actual, expected) {
		panic(fmt.Sprintf("Expected element '%s' to contain text '%s', but got '%s'", e.selector, expected, actual))
	}
	return e
}

// HasAttribute asserts that the first element has the specified attribute with the given value
func (e *ElementAssertion) HasAttribute(name, expected string) *ElementAssertion {
	e.Exists() // Ensure element exists first

	element := e.elements[0]
	if elem, ok := element.(*dom.Element); ok {
		actual := elem.GetAttribute(name)
		if actual != expected {
			panic(fmt.Sprintf("Expected element '%s' to have attribute '%s'='%s', but got '%s'", e.selector, name, expected, actual))
		}
	} else {
		panic(fmt.Sprintf("Element '%s' is not an Element node, cannot check attributes", e.selector))
	}
	return e
}

// HasAttributeContaining asserts that the first element has an attribute containing the specified value
func (e *ElementAssertion) HasAttributeContaining(name, expected string) *ElementAssertion {
	e.Exists() // Ensure element exists first

	element := e.elements[0]
	if elem, ok := element.(*dom.Element); ok {
		actual := elem.GetAttribute(name)
		if !strings.Contains(actual, expected) {
			panic(fmt.Sprintf("Expected element '%s' attribute '%s' to contain '%s', but got '%s'", e.selector, name, expected, actual))
		}
	} else {
		panic(fmt.Sprintf("Element '%s' is not an Element node, cannot check attributes", e.selector))
	}
	return e
}

// HasClass asserts that the first element has the specified CSS class
func (e *ElementAssertion) HasClass(className string) *ElementAssertion {
	e.Exists() // Ensure element exists first

	element := e.elements[0]
	if elem, ok := element.(*dom.Element); ok {
		classAttr := elem.GetAttribute("class")
		classes := strings.Fields(classAttr)

		for _, class := range classes {
			if class == className {
				return e
			}
		}

		panic(fmt.Sprintf("Expected element '%s' to have class '%s', but classes are: %v", e.selector, className, classes))
	} else {
		panic(fmt.Sprintf("Element '%s' is not an Element node, cannot check classes", e.selector))
	}
}

// IsVisible asserts that the first element is visible (not hidden via display:none or visibility:hidden)
// Note: This is a simplified implementation that only checks style attributes
func (e *ElementAssertion) IsVisible() *ElementAssertion {
	e.Exists() // Ensure element exists first

	element := e.elements[0]
	if elem, ok := element.(*dom.Element); ok {
		style := elem.GetAttribute("style")

		// Simple checks for common hiding patterns
		if strings.Contains(style, "display:none") || strings.Contains(style, "display: none") ||
			strings.Contains(style, "visibility:hidden") || strings.Contains(style, "visibility: hidden") {
			panic(fmt.Sprintf("Expected element '%s' to be visible, but it appears to be hidden", e.selector))
		}
	}
	return e
}

// IsHidden asserts that the first element is hidden
func (e *ElementAssertion) IsHidden() *ElementAssertion {
	e.Exists() // Ensure element exists first

	element := e.elements[0]
	if elem, ok := element.(*dom.Element); ok {
		style := elem.GetAttribute("style")

		// Simple checks for common hiding patterns
		if strings.Contains(style, "display:none") || strings.Contains(style, "display: none") ||
			strings.Contains(style, "visibility:hidden") || strings.Contains(style, "visibility: hidden") {
			return e
		}

		panic(fmt.Sprintf("Expected element '%s' to be hidden, but it appears to be visible", e.selector))
	}
	return e
}

// HasValue asserts that the first element (input/textarea/select) has the specified value
func (e *ElementAssertion) HasValue(expected string) *ElementAssertion {
	e.Exists() // Ensure element exists first

	element := e.elements[0]
	if elem, ok := element.(*dom.Element); ok {
		actual := elem.GetAttribute("value")
		if actual != expected {
			panic(fmt.Sprintf("Expected element '%s' to have value '%s', but got '%s'", e.selector, expected, actual))
		}
	} else {
		panic(fmt.Sprintf("Element '%s' is not an Element node, cannot check value", e.selector))
	}
	return e
}

// IsChecked asserts that the first element (checkbox/radio) is checked
func (e *ElementAssertion) IsChecked() *ElementAssertion {
	e.Exists() // Ensure element exists first

	element := e.elements[0]
	if elem, ok := element.(*dom.Element); ok {
		checked := elem.GetAttribute("checked")
		if checked == "" {
			panic(fmt.Sprintf("Expected element '%s' to be checked, but it is not", e.selector))
		}
	} else {
		panic(fmt.Sprintf("Element '%s' is not an Element node, cannot check checked state", e.selector))
	}
	return e
}

// IsNotChecked asserts that the first element (checkbox/radio) is not checked
func (e *ElementAssertion) IsNotChecked() *ElementAssertion {
	e.Exists() // Ensure element exists first

	element := e.elements[0]
	if elem, ok := element.(*dom.Element); ok {
		checked := elem.GetAttribute("checked")
		if checked != "" {
			panic(fmt.Sprintf("Expected element '%s' to not be checked, but it is", e.selector))
		}
	} else {
		panic(fmt.Sprintf("Element '%s' is not an Element node, cannot check checked state", e.selector))
	}
	return e
}

// HasTitle asserts that the document has the specified title
func (d *DocumentAssertion) HasTitle(expected string) *DocumentAssertion {
	// Find title element
	titleElements := css.QuerySelectorAll(d.document, "title")
	if titleElements.Length() == 0 {
		panic(fmt.Sprintf("Expected document title to be '%s', but no title element found", expected))
	}

	actual := getTextContent(titleElements.Item(0))
	if actual != expected {
		panic(fmt.Sprintf("Expected document title to be '%s', but got '%s'", expected, actual))
	}
	return d
}

// ContainsText asserts that the document body contains the specified text
func (d *DocumentAssertion) ContainsText(expected string) *DocumentAssertion {
	bodyElements := css.QuerySelectorAll(d.document, "body")
	if bodyElements.Length() == 0 {
		panic(fmt.Sprintf("Expected document to contain text '%s', but no body element found", expected))
	}

	actual := getTextContent(bodyElements.Item(0))
	if !strings.Contains(actual, expected) {
		panic(fmt.Sprintf("Expected document to contain text '%s', but body text is '%s'", expected, actual))
	}
	return d
}

// getTextContent recursively extracts text content from a DOM node
func getTextContent(node dom.Node) string {
	switch n := node.(type) {
	case *dom.Text:
		return n.NodeValue()
	case *dom.Element:
		var text strings.Builder
		for child := n.FirstChild(); child != nil; child = child.NextSibling() {
			text.WriteString(getTextContent(child))
		}
		return text.String()
	default:
		// For other node types, recursively get text from children
		var text strings.Builder
		for child := node.FirstChild(); child != nil; child = child.NextSibling() {
			text.WriteString(getTextContent(child))
		}
		return text.String()
	}
}
