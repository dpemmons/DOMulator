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

// ChildNode mixin methods

// Before inserts nodes just before this CDATA section node in its parent's children list.
func (c *CDATASection) Before(nodes ...interface{}) error {
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

	// Insert before this CDATA section node
	parent.InsertBefore(convertedNode, c)
	return nil
}

// After inserts nodes just after this CDATA section node in its parent's children list.
func (c *CDATASection) After(nodes ...interface{}) error {
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

// ReplaceWith replaces this CDATA section node with the given nodes.
func (c *CDATASection) ReplaceWith(nodes ...interface{}) error {
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
		// Just remove this CDATA section node
		parent.RemoveChild(c)
		return nil
	}

	// Replace this CDATA section node with the converted node
	parent.ReplaceChild(convertedNode, c)
	return nil
}

// Remove removes this CDATA section node from its parent.
func (c *CDATASection) Remove() error {
	parent := c.ParentNode()
	if parent == nil {
		return nil // No parent, nothing to do
	}

	parent.RemoveChild(c)
	return nil
}

// NonDocumentTypeChildNode mixin methods

// PreviousElementSibling returns the first preceding sibling that is an element; otherwise null.
func (c *CDATASection) PreviousElementSibling() *Element {
	parent := c.ParentNode()
	if parent == nil {
		return nil
	}

	siblings := parent.ChildNodes()
	for i := siblings.Length() - 1; i >= 0; i-- {
		if siblings.Item(i) == c {
			// Found this CDATA section node, now look backwards for an element sibling
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
func (c *CDATASection) NextElementSibling() *Element {
	parent := c.ParentNode()
	if parent == nil {
		return nil
	}

	siblings := parent.ChildNodes()
	for i := 0; i < siblings.Length(); i++ {
		sibling := siblings.Item(i)
		if sibling == c {
			// Found this CDATA section node, now look forwards for an element sibling
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
