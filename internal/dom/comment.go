package dom

import (
	"github.com/dop251/goja"
)

// Comment represents a comment node in the DOM tree.
type Comment struct {
	nodeImpl
}

// NewComment creates a new Comment node.
func NewComment(data string, doc *Document) *Comment {
	comment := &Comment{
		nodeImpl: nodeImpl{
			nodeType:      CommentNode,
			nodeName:      "#comment",
			nodeValue:     data,
			ownerDocument: doc,
		},
	}
	comment.nodeImpl.self = comment // Set the self reference
	return comment
}

// toJS is a placeholder for JavaScript binding.
func (c *Comment) toJS(vm *goja.Runtime) goja.Value {
	// This will be implemented later when integrating with Goja
	return vm.ToValue(nil)
}

// ChildNode mixin methods

// Before inserts nodes just before this comment node in its parent's children list.
func (c *Comment) Before(nodes ...interface{}) error {
	parent := c.ParentNode()
	if parent == nil {
		return nil // No parent, nothing to do
	}

	// Convert nodes to a single node or document fragment
	convertedNode, err := convertNodesToNode(nodes, c.ownerDocument)
	if err != nil {
		return err
	}

	if convertedNode == nil {
		return nil // Nothing to insert
	}

	// Insert before this comment node
	parent.InsertBefore(convertedNode, c)
	return nil
}

// After inserts nodes just after this comment node in its parent's children list.
func (c *Comment) After(nodes ...interface{}) error {
	parent := c.ParentNode()
	if parent == nil {
		return nil // No parent, nothing to do
	}

	// Convert nodes to a single node or document fragment
	convertedNode, err := convertNodesToNode(nodes, c.ownerDocument)
	if err != nil {
		return err
	}

	if convertedNode == nil {
		return nil // Nothing to insert
	}

	// Find the next viable sibling (not in the nodes being inserted)
	var excludeNodes []Node
	if fragment, ok := convertedNode.(*DocumentFragment); ok {
		excludeNodes = fragment.ChildNodes()
	} else {
		excludeNodes = []Node{convertedNode}
	}

	nextSibling := findViableNextSibling(c, excludeNodes)
	if nextSibling != nil {
		parent.InsertBefore(convertedNode, nextSibling)
	} else {
		parent.AppendChild(convertedNode)
	}
	return nil
}

// ReplaceWith replaces this comment node with the given nodes.
func (c *Comment) ReplaceWith(nodes ...interface{}) error {
	parent := c.ParentNode()
	if parent == nil {
		return nil // No parent, nothing to do
	}

	// Convert nodes to a single node or document fragment
	convertedNode, err := convertNodesToNode(nodes, c.ownerDocument)
	if err != nil {
		return err
	}

	if convertedNode == nil {
		// Just remove this comment node
		parent.RemoveChild(c)
		return nil
	}

	// Replace this comment node with the converted node
	parent.ReplaceChild(convertedNode, c)
	return nil
}

// Remove removes this comment node from its parent.
func (c *Comment) Remove() error {
	parent := c.ParentNode()
	if parent == nil {
		return nil // No parent, nothing to do
	}

	parent.RemoveChild(c)
	return nil
}
