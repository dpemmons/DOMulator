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
	return doc
}

// toJS is a placeholder for JavaScript binding.
func (d *Document) toJS(vm *goja.Runtime) goja.Value {
	// This will be implemented later when integrating with Goja
	return vm.ToValue(nil)
}
