package dom

import (
	"github.com/dop251/goja"
)

// Element represents an element node in the DOM tree.
type Element struct {
	nodeImpl

	tagName    string
	attributes map[string]string // Map to store attributes
	// classList      *ClassList    // To be implemented
	// dataset        *Dataset      // To be implemented
	// style          *CSSStyleDeclaration // To be implemented

	// Lazy-loaded properties
	innerHTML   *string
	outerHTML   *string
	textContent *string

	// Event handling
	// eventListeners map[EventPhase]map[string][]*EventListener // To be implemented

	// Shadow DOM support
	// shadowRoot     *ShadowRoot // To be implemented

	// Form elements
	value    string
	checked  bool
	selected bool

	// Mutation tracking
	// mutationRecord *MutationRecord // To be implemented
}

// NewElement creates a new Element node.
func NewElement(tagName string, doc *Document) *Element {
	elem := &Element{
		nodeImpl: nodeImpl{
			nodeType:      ElementNode,
			nodeName:      tagName, // NodeName for Element is its tagName
			ownerDocument: doc,
		},
		tagName:    tagName,
		attributes: make(map[string]string), // Initialize the attributes map
	}
	elem.nodeImpl.self = elem // Set the self reference
	return elem
}

// TagName returns the tag name of the element.
func (e *Element) TagName() string {
	return e.tagName
}

// SetAttribute sets the value of an attribute on the specified element.
// If the attribute already exists, its value is updated; otherwise, a new attribute is added.
func (e *Element) SetAttribute(name, value string) {
	e.attributes[name] = value
}

// GetAttribute returns the value of the named attribute on the specified element.
// If the named attribute does not exist, the value will be an empty string.
func (e *Element) GetAttribute(name string) string {
	return e.attributes[name]
}

// InnerHTML returns the HTML content of the element.
func (e *Element) InnerHTML() string {
	if e.innerHTML == nil {
		// TODO: Implement actual HTML serialization
		// For now, a placeholder
		content := ""
		for _, child := range e.childNodes {
			// This is a very basic placeholder.
			// A proper implementation would recursively serialize children.
			if child.NodeType() == TextNode {
				content += child.NodeValue()
			} else if child.NodeType() == ElementNode {
				content += child.(*Element).OuterHTML()
			}
		}
		e.innerHTML = &content
	}
	return *e.innerHTML
}

// SetInnerHTML sets the HTML content of the element.
func (e *Element) SetInnerHTML(html string) {
	// TODO: Implement HTML parsing and DOM manipulation
	// For now, clear existing children and set a text node
	e.childNodes = nil
	textNode := NewText(html, e.ownerDocument)
	e.AppendChild(textNode)
	e.innerHTML = &html // Invalidate cached innerHTML
	e.outerHTML = nil   // Invalidate cached outerHTML
	e.textContent = nil // Invalidate cached textContent
}

// OuterHTML returns the HTML content of the element including itself.
func (e *Element) OuterHTML() string {
	if e.outerHTML == nil {
		// TODO: Implement actual HTML serialization
		// For now, a placeholder
		content := "<" + e.tagName + ">" + e.InnerHTML() + "</" + e.tagName + ">"
		e.outerHTML = &content
	}
	return *e.outerHTML
}

// TextContent returns the text content of the element.
func (e *Element) TextContent() string {
	if e.textContent == nil {
		// TODO: Implement actual text content extraction
		// For now, a placeholder
		content := ""
		for _, child := range e.childNodes {
			if child.NodeType() == TextNode {
				content += child.NodeValue()
			} else if child.NodeType() == ElementNode {
				content += child.(*Element).TextContent()
			}
		}
		e.textContent = &content
	}
	return *e.textContent
}

// SetTextContent sets the text content of the element.
func (e *Element) SetTextContent(text string) {
	// TODO: Implement proper text content setting
	e.childNodes = nil
	textNode := NewText(text, e.ownerDocument)
	e.AppendChild(textNode)
	e.innerHTML = nil     // Invalidate cached innerHTML
	e.outerHTML = nil     // Invalidate cached outerHTML
	e.textContent = &text // Invalidate cached textContent
}

// toJS is a placeholder for JavaScript binding.
func (e *Element) toJS(vm *goja.Runtime) goja.Value {
	// This will be implemented later when integrating with Goja
	return vm.ToValue(nil)
}
