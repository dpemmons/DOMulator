package dom

import (
	"github.com/dop251/goja"
)

// Document represents the entire HTML document.
type Document struct {
	nodeImpl
	// Add document-specific properties here
	// For example, a reference to the window object, if we implement one
	// defaultView *Window
}

// NewDocument creates a new Document node.
func NewDocument() *Document {
	doc := &Document{
		nodeImpl: nodeImpl{
			nodeType: DocumentNode,
			nodeName: "#document",
		},
	}
	doc.ownerDocument = doc // A document is its own owner document
	doc.nodeImpl.self = doc // Set the self reference
	return doc
}

// CreateElement creates a new element with the given tag name
func (d *Document) CreateElement(tagName string) *Element {
	return NewElement(tagName, d)
}

// CreateElementNS creates a new element with the given namespace URI and qualified name
func (d *Document) CreateElementNS(namespaceURI, qualifiedName string) (*Element, error) {
	return NewElementNS(namespaceURI, qualifiedName, d)
}

// CreateTextNode creates a new text node with the given data
func (d *Document) CreateTextNode(data string) *Text {
	return NewText(data, d)
}

// CreateComment creates a new comment node with the given data
func (d *Document) CreateComment(data string) *Comment {
	return NewComment(data, d)
}

// CreateDocumentFragment creates a new document fragment
func (d *Document) CreateDocumentFragment() *DocumentFragment {
	return NewDocumentFragment(d)
}

// GetElementById returns the element with the specified ID
func (d *Document) GetElementById(id string) *Element {
	if id == "" {
		return nil // Don't search for empty IDs
	}
	var result *Element
	Traverse(d, func(node Node) bool {
		if elem, ok := node.(*Element); ok {
			if elem.GetAttribute("id") == id {
				result = elem
				return false // Stop traversal
			}
		}
		return true // Continue traversal
	})
	return result
}

// GetElementsByTagName returns all elements with the specified tag name
func (d *Document) GetElementsByTagName(tagName string) []*Element {
	var elements []*Element
	Traverse(d, func(node Node) bool {
		if elem, ok := node.(*Element); ok {
			if elem.TagName() == tagName || tagName == "*" {
				elements = append(elements, elem)
			}
		}
		return true // Continue traversal
	})
	return elements
}

// GetElementsByTagNameNS returns all elements with the specified namespace URI and local name
func (d *Document) GetElementsByTagNameNS(namespaceURI, localName string) []*Element {
	var elements []*Element
	Traverse(d, func(node Node) bool {
		if elem, ok := node.(*Element); ok {
			if (namespaceURI == "*" || elem.NamespaceURI() == namespaceURI) &&
				(localName == "*" || elem.LocalName() == localName) {
				elements = append(elements, elem)
			}
		}
		return true // Continue traversal
	})
	return elements
}

// GetElementsByClassName returns all elements with the specified class name
func (d *Document) GetElementsByClassName(className string) []*Element {
	var elements []*Element
	Traverse(d, func(node Node) bool {
		if elem, ok := node.(*Element); ok {
			if elem.HasClass(className) {
				elements = append(elements, elem)
			}
		}
		return true // Continue traversal
	})
	return elements
}

// DocumentElement returns the document element (usually <html>)
func (d *Document) DocumentElement() *Element {
	for _, child := range d.childNodes {
		if elem, ok := child.(*Element); ok && elem.TagName() == "html" {
			return elem
		}
	}
	return nil
}

// Body returns the body element
func (d *Document) Body() *Element {
	if docElem := d.DocumentElement(); docElem != nil {
		for _, child := range docElem.ChildNodes() {
			if elem, ok := child.(*Element); ok && elem.TagName() == "body" {
				return elem
			}
		}
	}
	return nil
}

// Head returns the head element
func (d *Document) Head() *Element {
	if docElem := d.DocumentElement(); docElem != nil {
		for _, child := range docElem.ChildNodes() {
			if elem, ok := child.(*Element); ok && elem.TagName() == "head" {
				return elem
			}
		}
	}
	return nil
}

// toJS is a placeholder for JavaScript binding.
func (d *Document) toJS(vm *goja.Runtime) goja.Value {
	// This will be implemented later when integrating with Goja
	return vm.ToValue(nil)
}
