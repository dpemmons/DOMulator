package dom

import (
	"github.com/dop251/goja"
)

// Text represents a text node in the DOM tree.
type Text struct {
	nodeImpl
}

// NewText creates a new Text node.
func NewText(data string, doc *Document) *Text {
	text := &Text{
		nodeImpl: nodeImpl{
			nodeType:      TextNode,
			nodeName:      "#text",
			nodeValue:     data,
			ownerDocument: doc,
		},
	}
	text.nodeImpl.self = text // Set the self reference
	return text
}

// toJS is a placeholder for JavaScript binding.
func (t *Text) toJS(vm *goja.Runtime) goja.Value {
	// This will be implemented later when integrating with Goja
	return vm.ToValue(nil)
}
