package dom

import (
	"github.com/dop251/goja"
)

// DocumentType represents a DocumentType node in the DOM tree.
type DocumentType struct {
	nodeImpl
	name     string
	publicID string
	systemID string
}

// NewDocumentType creates a new DocumentType node.
func NewDocumentType(name, publicID, systemID string, doc *Document) *DocumentType {
	dt := &DocumentType{
		nodeImpl: nodeImpl{
			nodeType:      DocumentTypeNode,
			nodeName:      name,
			ownerDocument: doc,
		},
		name:     name,
		publicID: publicID,
		systemID: systemID,
	}
	return dt
}

// Name returns the name of the document type.
func (dt *DocumentType) Name() string {
	return dt.name
}

// PublicID returns the public ID of the document type.
func (dt *DocumentType) PublicID() string {
	return dt.publicID
}

// SystemID returns the system ID of the document type.
func (dt *DocumentType) SystemID() string {
	return dt.systemID
}

// toJS is a placeholder for JavaScript binding.
func (dt *DocumentType) toJS(vm *goja.Runtime) goja.Value {
	// This will be implemented later when integrating with Goja
	return vm.ToValue(nil)
}

// ChildNode mixin methods

// Before inserts nodes just before this DocumentType node in its parent's children list.
func (dt *DocumentType) Before(nodes ...interface{}) error {
	parent := dt.ParentNode()
	if parent == nil {
		return nil // No parent, nothing to do
	}

	// Convert nodes to a single node or document fragment
	convertedNode, err := convertNodesToNode(nodes, dt.ownerDocument)
	if err != nil {
		return err
	}

	if convertedNode == nil {
		return nil // Nothing to insert
	}

	// Insert before this DocumentType node
	parent.InsertBefore(convertedNode, dt)
	return nil
}

// After inserts nodes just after this DocumentType node in its parent's children list.
func (dt *DocumentType) After(nodes ...interface{}) error {
	parent := dt.ParentNode()
	if parent == nil {
		return nil // No parent, nothing to do
	}

	// Convert nodes to a single node or document fragment
	convertedNode, err := convertNodesToNode(nodes, dt.ownerDocument)
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

	nextSibling := findViableNextSibling(dt, excludeNodes)
	if nextSibling != nil {
		parent.InsertBefore(convertedNode, nextSibling)
	} else {
		parent.AppendChild(convertedNode)
	}
	return nil
}

// ReplaceWith replaces this DocumentType node with the given nodes.
func (dt *DocumentType) ReplaceWith(nodes ...interface{}) error {
	parent := dt.ParentNode()
	if parent == nil {
		return nil // No parent, nothing to do
	}

	// Convert nodes to a single node or document fragment
	convertedNode, err := convertNodesToNode(nodes, dt.ownerDocument)
	if err != nil {
		return err
	}

	if convertedNode == nil {
		// Just remove this DocumentType node
		parent.RemoveChild(dt)
		return nil
	}

	// Replace this DocumentType node with the converted node
	parent.ReplaceChild(convertedNode, dt)
	return nil
}

// Remove removes this DocumentType node from its parent.
func (dt *DocumentType) Remove() error {
	parent := dt.ParentNode()
	if parent == nil {
		return nil // No parent, nothing to do
	}

	parent.RemoveChild(dt)
	return nil
}
