package dom

import (
	"github.com/dop251/goja"
)

// CDATASection represents a CDATA section node in the DOM tree.
type CDATASection struct {
	characterDataImpl
}

// NewCDATASection creates a new CDATASection node.
func NewCDATASection(data string, doc *Document) *CDATASection {
	cdata := &CDATASection{
		characterDataImpl: characterDataImpl{
			nodeImpl: nodeImpl{
				nodeType:      CDataSectionNode, // Corrected to CDataSectionNode
				nodeName:      "#cdata-section",
				nodeValue:     data,
				ownerDocument: doc,
			},
		},
	}
	return cdata
}

// toJS is a placeholder for JavaScript binding.
func (c *CDATASection) toJS(vm *goja.Runtime) goja.Value {
	// This will be implemented later when integrating with Goja
	return vm.ToValue(nil)
}
