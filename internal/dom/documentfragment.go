package dom

import (
	"github.com/dop251/goja"
)

// DocumentFragment represents a lightweight Document object whose contents are not part of the main document.
type DocumentFragment struct {
	nodeImpl
}

// NewDocumentFragment creates a new DocumentFragment node.
func NewDocumentFragment(doc *Document) *DocumentFragment {
	df := &DocumentFragment{
		nodeImpl: nodeImpl{
			nodeType:      DocumentFragmentNode,
			nodeName:      "#document-fragment",
			ownerDocument: doc,
		},
	}
	df.nodeImpl.self = df // Set the self reference
	return df
}

// GetElementById returns the first element within the fragment's descendants whose ID is elementId.
// This implements the NonElementParentNode mixin as specified in WHATWG DOM Standard 4.2.4.
// The method steps are to return the first element, in tree order, within this's descendants,
// whose ID is elementId; otherwise, if there is no such element, null.
func (df *DocumentFragment) GetElementById(id string) *Element {
	if id == "" {
		return nil // Don't search for empty IDs
	}
	return df.findElementByIdRecursive(id)
}

// findElementByIdRecursive performs a depth-first search for an element with the given ID.
// It properly stops traversal when the first match is found.
func (df *DocumentFragment) findElementByIdRecursive(id string) *Element {
	// Search through all direct children
	for _, child := range df.ChildNodes() {
		if elem, ok := child.(*Element); ok {
			// Check if this element has the target ID
			if elem.GetAttribute("id") == id {
				return elem
			}
			// Recursively search in this element's descendants
			if found := elem.findElementByIdRecursive(id); found != nil {
				return found
			}
		}
	}
	return nil
}

// toJS is a placeholder for JavaScript binding.
func (df *DocumentFragment) toJS(vm *goja.Runtime) goja.Value {
	// This will be implemented later when integrating with Goja
	return vm.ToValue(nil)
}
