package dom

import (
	"github.com/dop251/goja"
)

// Text represents a text node in the DOM tree.
type Text struct {
	characterDataImpl
}

// NewText creates a new Text node following WHATWG DOM specification.
func NewText(data string, doc *Document) *Text {
	text := &Text{
		characterDataImpl: characterDataImpl{
			nodeImpl: nodeImpl{
				nodeType:      TextNode,
				nodeName:      "#text",
				nodeValue:     data,
				ownerDocument: doc,
			},
		},
	}
	text.characterDataImpl.nodeImpl.self = text // Set the self reference
	return text
}

// WholeText returns the concatenation of the data of all the contiguous Text nodes
// of this node, in tree order, as specified in WHATWG DOM Section 4.11.
func (t *Text) WholeText() string {
	// Step 1: Get the contiguous text nodes of this node
	contiguousNodes := t.getContiguousTextNodes()

	// Step 2: Concatenate their data in tree order
	var result string
	for _, node := range contiguousNodes {
		if textNode, ok := node.(*Text); ok {
			result += textNode.Data()
		}
	}

	return result
}

// getContiguousTextNodes returns the contiguous Text nodes of this node.
// The contiguous Text nodes of a node are: node, node's previous sibling Text node,
// if any, and its contiguous Text nodes, and node's next sibling Text node, if any,
// and its contiguous Text nodes, avoiding any duplicates.
func (t *Text) getContiguousTextNodes() []Node {
	var nodes []Node

	// Collect previous contiguous text nodes (in reverse order)
	var prevNodes []Node
	current := t.PreviousSibling()
	for current != nil && current.NodeType() == TextNode {
		prevNodes = append(prevNodes, current)
		current = current.PreviousSibling()
	}

	// Add previous nodes in correct order (reverse the slice)
	for i := len(prevNodes) - 1; i >= 0; i-- {
		nodes = append(nodes, prevNodes[i])
	}

	// Add this node
	nodes = append(nodes, t)

	// Collect next contiguous text nodes
	current = t.NextSibling()
	for current != nil && current.NodeType() == TextNode {
		nodes = append(nodes, current)
		current = current.NextSibling()
	}

	return nodes
}

// SplitText splits this text node into two text nodes at the specified offset
// following the WHATWG DOM specification algorithm with backward compatibility.
func (t *Text) SplitText(offset int) *Text {
	// Handle negative offset (backward compatibility)
	if offset < 0 {
		offset = 0
	}

	// Step 1: Let length be node's length
	length := t.Length()

	// Backward compatibility: If offset is greater than length, treat as splitting at end
	if offset > length {
		offset = length
	}

	// Step 3: Let count be length minus offset
	count := length - offset

	// Step 4: Let new data be the result of substringing data with node, offset, and count
	var newData string
	if count > 0 {
		newData = t.substringDataInternal(offset, count)
	} else {
		newData = ""
	}

	// Step 5: Let new node be a new Text node, with the same node document as node
	newNode := NewText(newData, t.ownerDocument)

	// Step 6: Let parent be node's parent
	parent := t.ParentNode()

	// Step 7: If parent is not null:
	if parent != nil {
		// Step 7.1: Insert new node into parent before node's next sibling
		nextSibling := t.NextSibling()
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
		t.replaceDataInternal(offset, count, "")
	}

	// Step 9: Return new node
	return newNode
}

// CloneNode creates a copy of the text node using the spec-compliant cloning implementation.
func (t *Text) CloneNode(deep bool) Node {
	// Use the spec-compliant cloning implementation
	return CloneNodeSpec(t, deep)
}

// toJS is a placeholder for JavaScript binding.
func (t *Text) toJS(vm *goja.Runtime) goja.Value {
	// This will be implemented later when integrating with Goja
	return vm.ToValue(nil)
}

// ChildNode mixin methods

// Before inserts nodes just before this text node in its parent's children list.
func (t *Text) Before(nodes ...interface{}) error {
	parent := t.ParentNode()
	if parent == nil {
		return nil // No parent, nothing to do
	}

	// Convert nodes to a single node or document fragment
	convertedNode, err := convertNodesToNode(nodes, t.ownerDocument)
	if err != nil {
		return err
	}

	if convertedNode == nil {
		return nil // Nothing to insert
	}

	// Insert before this text node
	parent.InsertBefore(convertedNode, t)
	return nil
}

// After inserts nodes just after this text node in its parent's children list.
func (t *Text) After(nodes ...interface{}) error {
	parent := t.ParentNode()
	if parent == nil {
		return nil // No parent, nothing to do
	}

	// Convert nodes to a single node or document fragment
	convertedNode, err := convertNodesToNode(nodes, t.ownerDocument)
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

	nextSibling := findViableNextSibling(t, excludeNodes)
	if nextSibling != nil {
		parent.InsertBefore(convertedNode, nextSibling)
	} else {
		parent.AppendChild(convertedNode)
	}
	return nil
}

// ReplaceWith replaces this text node with the given nodes.
func (t *Text) ReplaceWith(nodes ...interface{}) error {
	parent := t.ParentNode()
	if parent == nil {
		return nil // No parent, nothing to do
	}

	// Convert nodes to a single node or document fragment
	convertedNode, err := convertNodesToNode(nodes, t.ownerDocument)
	if err != nil {
		return err
	}

	if convertedNode == nil {
		// Just remove this text node
		parent.RemoveChild(t)
		return nil
	}

	// Replace this text node with the converted node
	parent.ReplaceChild(convertedNode, t)
	return nil
}

// Remove removes this text node from its parent.
func (t *Text) Remove() error {
	parent := t.ParentNode()
	if parent == nil {
		return nil // No parent, nothing to do
	}

	parent.RemoveChild(t)
	return nil
}

// NonDocumentTypeChildNode mixin methods

// PreviousElementSibling returns the first preceding sibling that is an element; otherwise null.
func (t *Text) PreviousElementSibling() *Element {
	parent := t.ParentNode()
	if parent == nil {
		return nil
	}

	siblings := parent.ChildNodes()
	for i := siblings.Length() - 1; i >= 0; i-- {
		if siblings.Item(i) == t {
			// Found this text node, now look backwards for an element sibling
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
func (t *Text) NextElementSibling() *Element {
	parent := t.ParentNode()
	if parent == nil {
		return nil
	}

	siblings := parent.ChildNodes()
	for i := 0; i < siblings.Length(); i++ {
		sibling := siblings.Item(i)
		if sibling == t {
			// Found this text node, now look forwards for an element sibling
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
