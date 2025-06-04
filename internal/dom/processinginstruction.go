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

// ChildNode mixin methods

// Before inserts nodes just before this processing instruction node in its parent's children list.
func (pi *ProcessingInstruction) Before(nodes ...interface{}) error {
	parent := pi.ParentNode()
	if parent == nil {
		return nil // No parent, nothing to do
	}

	// Convert nodes to a single node or document fragment
	convertedNode, err := convertNodesToNode(nodes, pi.ownerDocument)
	if err != nil {
		return err
	}

	if convertedNode == nil {
		return nil // Nothing to insert
	}

	// Insert before this processing instruction node
	parent.InsertBefore(convertedNode, pi)
	return nil
}

// After inserts nodes just after this processing instruction node in its parent's children list.
func (pi *ProcessingInstruction) After(nodes ...interface{}) error {
	parent := pi.ParentNode()
	if parent == nil {
		return nil // No parent, nothing to do
	}

	// Convert nodes to a single node or document fragment
	convertedNode, err := convertNodesToNode(nodes, pi.ownerDocument)
	if err != nil {
		return err
	}

	if convertedNode == nil {
		return nil // Nothing to insert
	}

	// Find the next viable sibling (not in the nodes being inserted)
	var excludeNodes []Node
	if fragment, ok := convertedNode.(*DocumentFragment); ok {
		fragmentChildren := fragment.ChildNodes()
		for i := 0; i < fragmentChildren.Length(); i++ {
			excludeNodes = append(excludeNodes, fragmentChildren.Item(i))
		}
	} else {
		excludeNodes = []Node{convertedNode}
	}

	nextSibling := findViableNextSibling(pi, excludeNodes)
	if nextSibling != nil {
		parent.InsertBefore(convertedNode, nextSibling)
	} else {
		parent.AppendChild(convertedNode)
	}
	return nil
}

// ReplaceWith replaces this processing instruction node with the given nodes.
func (pi *ProcessingInstruction) ReplaceWith(nodes ...interface{}) error {
	parent := pi.ParentNode()
	if parent == nil {
		return nil // No parent, nothing to do
	}

	// Convert nodes to a single node or document fragment
	convertedNode, err := convertNodesToNode(nodes, pi.ownerDocument)
	if err != nil {
		return err
	}

	if convertedNode == nil {
		// Just remove this processing instruction node
		parent.RemoveChild(pi)
		return nil
	}

	// Replace this processing instruction node with the converted node
	parent.ReplaceChild(convertedNode, pi)
	return nil
}

// Remove removes this processing instruction node from its parent.
func (pi *ProcessingInstruction) Remove() error {
	parent := pi.ParentNode()
	if parent == nil {
		return nil // No parent, nothing to do
	}

	parent.RemoveChild(pi)
	return nil
}
