package dom

import (
	"github.com/dop251/goja"
)

// CDATASection represents a CDATA section node in the DOM tree.
// Per WHATWG DOM Section 4.12: interface CDATASection : Text {};
type CDATASection struct {
	Text
}

// NewCDATASection creates a new CDATASection node following WHATWG DOM specification.
func NewCDATASection(data string, doc *Document) *CDATASection {
	// Create a Text node first
	text := NewText(data, doc)

	// Create CDATASection that inherits from Text
	cdata := &CDATASection{
		Text: *text,
	}

	// Override the node type to be CDATASection while keeping all Text functionality
	cdata.nodeImpl.nodeType = CDataSectionNode
	cdata.nodeImpl.nodeName = "#cdata-section"
	cdata.nodeImpl.self = cdata // Set the self reference to the CDATASection

	return cdata
}

// SplitText splits this CDATA section into two CDATA sections at the specified offset
// following the WHATWG DOM specification algorithm but maintaining CDATASection type.
func (c *CDATASection) SplitText(offset int) *Text {
	// Handle negative offset (backward compatibility)
	if offset < 0 {
		offset = 0
	}

	// Step 1: Let length be node's length
	length := c.Length()

	// Backward compatibility: If offset is greater than length, treat as splitting at end
	if offset > length {
		offset = length
	}

	// Step 3: Let count be length minus offset
	count := length - offset

	// Step 4: Let new data be the result of substringing data with node, offset, and count
	var newData string
	if count > 0 {
		newData = c.substringDataInternal(offset, count)
	} else {
		newData = ""
	}

	// Step 5: Let new node be a new CDATASection node (not Text), with the same node document as node
	newNode := NewCDATASection(newData, c.ownerDocument)

	// Step 6: Let parent be node's parent
	parent := c.ParentNode()

	// Step 7: If parent is not null:
	if parent != nil {
		// Step 7.1: Insert new node into parent before node's next sibling
		nextSibling := c.NextSibling()
		if nextSibling != nil {
			parent.InsertBefore(newNode, nextSibling)
		} else {
			parent.AppendChild(newNode)
		}

		// Steps 7.2-7.5: Range updates (not implemented in this DOM - placeholder for future)
		// For each live range whose start node is node and start offset is greater than offset,
		// set its start node to new node and decrease its start offset by offset.
		// For each live range whose end node is node and end offset is greater than offset,
		// set its end node to new node and decrease its end offset by offset.
		// For each live range whose start node is parent and start offset is equal to
		// the index of node plus 1, increase its start offset by 1.
		// For each live range whose end node is parent and end offset is equal to
		// the index of node plus 1, increase its end offset by 1.
	}

	// Step 8: Replace data with node, offset, count, and data the empty string
	if count > 0 {
		c.replaceDataInternal(offset, count, "")
	}

	// Step 9: Return new node (access embedded Text while preserving CDATASection type)
	return &newNode.Text
}

// CloneNode creates a copy of the CDATA section using the spec-compliant cloning implementation.
func (c *CDATASection) CloneNode(deep bool) Node {
	// Use the spec-compliant cloning implementation
	return CloneNodeSpec(c, deep)
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
