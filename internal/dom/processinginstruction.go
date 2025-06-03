package dom

import (
	"github.com/dop251/goja"
)

// ProcessingInstruction represents a processing instruction node in the DOM tree.
type ProcessingInstruction struct {
	characterDataImpl
	target string
}

// NewProcessingInstruction creates a new ProcessingInstruction node.
func NewProcessingInstruction(target, data string, doc *Document) *ProcessingInstruction {
	pi := &ProcessingInstruction{
		characterDataImpl: characterDataImpl{
			nodeImpl: nodeImpl{
				nodeType:      ProcessingInstructionNode,
				nodeName:      target, // NodeName for PI is its target
				nodeValue:     data,
				ownerDocument: doc,
			},
		},
		target: target,
	}
	return pi
}

// Target returns the target of the processing instruction.
func (pi *ProcessingInstruction) Target() string {
	return pi.target
}

// toJS is a placeholder for JavaScript binding.
func (pi *ProcessingInstruction) toJS(vm *goja.Runtime) goja.Value {
	// This will be implemented later when integrating with Goja
	return vm.ToValue(nil)
}
