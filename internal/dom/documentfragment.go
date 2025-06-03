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
	return df
}

// toJS is a placeholder for JavaScript binding.
func (df *DocumentFragment) toJS(vm *goja.Runtime) goja.Value {
	// This will be implemented later when integrating with Goja
	return vm.ToValue(nil)
}
