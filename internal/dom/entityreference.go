package dom

import (
	"github.com/dop251/goja"
)

// EntityReference represents an EntityReference node in the DOM tree.
type EntityReference struct {
	nodeImpl
}

// NewEntityReference creates a new EntityReference node.
func NewEntityReference(name string, doc *Document) *EntityReference {
	er := &EntityReference{
		nodeImpl: nodeImpl{
			nodeType:      EntityReferenceNode,
			nodeName:      name,
			ownerDocument: doc,
		},
	}
	return er
}

// toJS is a placeholder for JavaScript binding.
func (er *EntityReference) toJS(vm *goja.Runtime) goja.Value {
	// This will be implemented later when integrating with Goja
	return vm.ToValue(nil)
}
