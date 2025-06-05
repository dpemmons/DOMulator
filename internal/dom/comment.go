package dom

import (
	"github.com/dop251/goja"
)

// Comment represents a comment node in the DOM tree.
type Comment struct {
	characterDataImpl
}

// NewComment creates a new Comment node following WHATWG DOM specification.
// Per WHATWG DOM Section 4.14: "The new Comment(data) constructor steps are to
// set this's data to data and this's node document to current global object's associated Document."
func NewComment(data string, doc *Document) *Comment {
	comment := &Comment{
		characterDataImpl: characterDataImpl{
			nodeImpl: nodeImpl{
				nodeType:      CommentNode,
				nodeName:      "#comment",
				nodeValue:     data,
				ownerDocument: doc,
			},
		},
	}
	comment.characterDataImpl.nodeImpl.self = comment // Set the self reference
	return comment
}

// CloneNode creates a copy of the comment node using the spec-compliant cloning implementation.
func (c *Comment) CloneNode(deep bool) Node {
	// Use the spec-compliant cloning implementation
	return CloneNodeSpec(c, deep)
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
		fragmentChildren := fragment.ChildNodes()
		for i := 0; i < fragmentChildren.Length(); i++ {
			excludeNodes = append(excludeNodes, fragmentChildren.Item(i))
		}
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

// NonDocumentTypeChildNode mixin methods

// PreviousElementSibling returns the first preceding sibling that is an element; otherwise null.
func (c *Comment) PreviousElementSibling() *Element {
	parent := c.ParentNode()
	if parent == nil {
		return nil
	}

	siblings := parent.ChildNodes()
	for i := siblings.Length() - 1; i >= 0; i-- {
		if siblings.Item(i) == c {
			// Found this comment node, now look backwards for an element sibling
			for j := i - 1; j >= 0; j-- {
				sibling := siblings.Item(j)
				if sibling.NodeType() == ElementNode {
					return sibling.(*Element)
				}
			}
			break
		}
	}
	return nil
}

// NextElementSibling returns the first following sibling that is an element; otherwise null.
func (c *Comment) NextElementSibling() *Element {
	parent := c.ParentNode()
	if parent == nil {
		return nil
	}

	siblings := parent.ChildNodes()
	for i := 0; i < siblings.Length(); i++ {
		sibling := siblings.Item(i)
		if sibling == c {
			// Found this comment node, now look forwards for an element sibling
			for j := i + 1; j < siblings.Length(); j++ {
				nextSibling := siblings.Item(j)
				if nextSibling.NodeType() == ElementNode {
					return nextSibling.(*Element)
				}
			}
			break
		}
	}
	return nil
}
