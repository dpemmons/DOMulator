package ranges

import (
	"strings"
	"sync"
)

// Range represents a live range that updates when the node tree mutates
// per WHATWG DOM Section 5.5
type Range struct {
	AbstractRangeImpl
	mutex sync.RWMutex // For thread safety
	root  Node         // The root of the live range
}

// Range constants for compareBoundaryPoints
const (
	START_TO_START = 0
	START_TO_END   = 1
	END_TO_END     = 2
	END_TO_START   = 3
)

// NewRange creates a new live range
// The new Range() constructor steps are to set this's start and end to
// (current global object's associated Document, 0).
func NewRange(document Document) *Range {
	// Get the actual document node from the document adapter
	var docNode Node = document
	if adapter, ok := document.(interface{ GetDocument() Document }); ok {
		docNode = adapter.GetDocument()
	}

	r := &Range{
		AbstractRangeImpl: AbstractRangeImpl{
			start: BoundaryPoint{Node: docNode, Offset: 0},
			end:   BoundaryPoint{Node: docNode, Offset: 0},
		},
		root: docNode,
	}
	return r
}

// CommonAncestorContainer returns the node, furthest away from the document,
// that is an ancestor of both range's start node and end node.
// The commonAncestorContainer getter steps are:
// 1. Let container be start node.
// 2. While container is not an inclusive ancestor of end node, let container be container's parent.
// 3. Return container.
func (r *Range) CommonAncestorContainer() Node {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	container := r.start.Node
	for !isInclusiveAncestor(container, r.end.Node) {
		container = container.ParentNode()
		if container == nil {
			// This shouldn't happen in a valid range
			return r.root
		}
	}
	return container
}

// SetStart sets the start of this range to boundary point (node, offset)
func (r *Range) SetStart(node Node, offset int) error {
	return r.setStart(node, offset)
}

// SetEnd sets the end of this range to boundary point (node, offset)
func (r *Range) SetEnd(node Node, offset int) error {
	return r.setEnd(node, offset)
}

// setStart implements the "set the start" algorithm from the specification
func (r *Range) setStart(node Node, offset int) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	// If node is a doctype, then throw an "InvalidNodeTypeError" DOMException
	if node.NodeType() == DocumentTypeNode {
		return &DOMException{
			Code:    InvalidNodeTypeError,
			Message: "InvalidNodeTypeError",
		}
	}

	// If offset is greater than node's length, then throw an "IndexSizeError" DOMException
	if offset > getNodeLength(node) {
		return &DOMException{
			Code:    IndexSizeError,
			Message: "IndexSizeError",
		}
	}

	bp := BoundaryPoint{Node: node, Offset: offset}

	// If range's root is not equal to node's root, or if bp is after the range's end,
	// set range's end to bp
	nodeRoot := getRoot(node)
	if r.root != nodeRoot || CompareBoundaryPoints(&bp, &r.end) == 1 {
		r.end = bp
		r.root = nodeRoot
	}

	// Set range's start to bp
	r.start = bp
	return nil
}

// setEnd implements the "set the end" algorithm from the specification
func (r *Range) setEnd(node Node, offset int) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	// If node is a doctype, then throw an "InvalidNodeTypeError" DOMException
	if node.NodeType() == DocumentTypeNode {
		return &DOMException{
			Code:    InvalidNodeTypeError,
			Message: "InvalidNodeTypeError",
		}
	}

	// If offset is greater than node's length, then throw an "IndexSizeError" DOMException
	if offset > getNodeLength(node) {
		return &DOMException{
			Code:    IndexSizeError,
			Message: "IndexSizeError",
		}
	}

	bp := BoundaryPoint{Node: node, Offset: offset}

	// If range's root is not equal to node's root, or if bp is before the range's start,
	// set range's start to bp
	nodeRoot := getRoot(node)
	if r.root != nodeRoot || CompareBoundaryPoints(&bp, &r.start) == -1 {
		r.start = bp
		r.root = nodeRoot
	}

	// Set range's end to bp
	r.end = bp
	return nil
}

// SetStartBefore sets the start of this range to immediately before the given node
func (r *Range) SetStartBefore(node Node) error {
	parent := node.ParentNode()
	if parent == nil {
		return &DOMException{
			Code:    InvalidNodeTypeError,
			Message: "InvalidNodeTypeError",
		}
	}
	return r.setStart(parent, getNodeIndex(node))
}

// SetStartAfter sets the start of this range to immediately after the given node
func (r *Range) SetStartAfter(node Node) error {
	parent := node.ParentNode()
	if parent == nil {
		return &DOMException{
			Code:    InvalidNodeTypeError,
			Message: "node has no parent",
		}
	}
	return r.setStart(parent, getNodeIndex(node)+1)
}

// SetEndBefore sets the end of this range to immediately before the given node
func (r *Range) SetEndBefore(node Node) error {
	parent := node.ParentNode()
	if parent == nil {
		return &DOMException{
			Code:    InvalidNodeTypeError,
			Message: "node has no parent",
		}
	}
	return r.setEnd(parent, getNodeIndex(node))
}

// SetEndAfter sets the end of this range to immediately after the given node
func (r *Range) SetEndAfter(node Node) error {
	parent := node.ParentNode()
	if parent == nil {
		return &DOMException{
			Code:    InvalidNodeTypeError,
			Message: "node has no parent",
		}
	}
	return r.setEnd(parent, getNodeIndex(node)+1)
}

// Collapse collapses the range to one of its endpoints
// If toStart is true, set end to start; otherwise set start to end
func (r *Range) Collapse(toStart bool) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if toStart {
		r.end = r.start
	} else {
		r.start = r.end
	}
}

// SelectNode selects the given node within this range
func (r *Range) SelectNode(node Node) error {
	parent := node.ParentNode()
	if parent == nil {
		return &DOMException{
			Code:    InvalidNodeTypeError,
			Message: "InvalidNodeTypeError",
		}
	}

	index := getNodeIndex(node)
	if err := r.setStart(parent, index); err != nil {
		return err
	}
	return r.setEnd(parent, index+1)
}

// SelectNodeContents selects the contents of the given node
func (r *Range) SelectNodeContents(node Node) error {
	if node.NodeType() == DocumentTypeNode {
		return &DOMException{
			Code:    InvalidNodeTypeError,
			Message: "InvalidNodeTypeError",
		}
	}

	length := getNodeLength(node)
	if err := r.setStart(node, 0); err != nil {
		return err
	}
	return r.setEnd(node, length)
}

// CompareBoundaryPoints compares boundary points between this range and another range
func (r *Range) CompareBoundaryPoints(how int, sourceRange *Range) (int, error) {
	// Validate the 'how' parameter
	if how != START_TO_START && how != START_TO_END && how != END_TO_END && how != END_TO_START {
		return 0, &DOMException{
			Code:    NotSupportedError,
			Message: "NotSupportedError",
		}
	}

	// Check if ranges are in the same tree
	if r.root != sourceRange.root {
		return 0, &DOMException{
			Code:    WrongDocumentError,
			Message: "ranges are not in the same document",
		}
	}

	r.mutex.RLock()
	sourceRange.mutex.RLock()
	defer r.mutex.RUnlock()
	defer sourceRange.mutex.RUnlock()

	var thisPoint, otherPoint *BoundaryPoint

	switch how {
	case START_TO_START:
		thisPoint = &r.start
		otherPoint = &sourceRange.start
	case START_TO_END:
		thisPoint = &r.start
		otherPoint = &sourceRange.end
	case END_TO_END:
		thisPoint = &r.end
		otherPoint = &sourceRange.end
	case END_TO_START:
		thisPoint = &r.end
		otherPoint = &sourceRange.start
	}

	return CompareBoundaryPoints(thisPoint, otherPoint), nil
}

// IsPointInRange checks if the given point is within this range
func (r *Range) IsPointInRange(node Node, offset int) (bool, error) {
	// If node's root is different from this's root, return false
	if getRoot(node) != r.root {
		return false, nil
	}

	// If node is a doctype, then throw an "InvalidNodeTypeError" DOMException
	if node.NodeType() == DocumentTypeNode {
		return false, &DOMException{
			Code:    InvalidNodeTypeError,
			Message: "InvalidNodeTypeError",
		}
	}

	// If offset is greater than node's length, then throw an "IndexSizeError" DOMException
	if offset > getNodeLength(node) {
		return false, &DOMException{
			Code:    IndexSizeError,
			Message: "IndexSizeError",
		}
	}

	r.mutex.RLock()
	defer r.mutex.RUnlock()

	point := BoundaryPoint{Node: node, Offset: offset}

	// If (node, offset) is before start or after end, return false
	// Note: end is exclusive, so equal to end also returns false
	if CompareBoundaryPoints(&point, &r.start) == -1 || CompareBoundaryPoints(&point, &r.end) >= 0 {
		return false, nil
	}

	return true, nil
}

// ComparePoint compares a point to this range
// Returns -1 if the point is before the range, 0 if in the range, 1 if after
func (r *Range) ComparePoint(node Node, offset int) (int, error) {
	// If node's root is different from this's root, throw a "WrongDocumentError"
	if getRoot(node) != r.root {
		return 0, &DOMException{
			Code:    WrongDocumentError,
			Message: "node is not in the same document",
		}
	}

	// If node is a doctype, then throw an "InvalidNodeTypeError" DOMException
	if node.NodeType() == DocumentTypeNode {
		return 0, &DOMException{
			Code:    InvalidNodeTypeError,
			Message: "node cannot be a DocumentType",
		}
	}

	// If offset is greater than node's length, then throw an "IndexSizeError" DOMException
	if offset > getNodeLength(node) {
		return 0, &DOMException{
			Code:    IndexSizeError,
			Message: "offset is greater than node's length",
		}
	}

	r.mutex.RLock()
	defer r.mutex.RUnlock()

	point := BoundaryPoint{Node: node, Offset: offset}

	// If (node, offset) is before start, return -1
	if CompareBoundaryPoints(&point, &r.start) == -1 {
		return -1, nil
	}

	// If (node, offset) is after end, return 1
	if CompareBoundaryPoints(&point, &r.end) == 1 {
		return 1, nil
	}

	// Return 0 (point is in range)
	return 0, nil
}

// IntersectsNode checks if this range intersects the given node
func (r *Range) IntersectsNode(node Node) bool {
	// If node's root is different from this's root, return false
	if getRoot(node) != r.root {
		return false
	}

	parent := node.ParentNode()
	if parent == nil {
		return true
	}

	r.mutex.RLock()
	defer r.mutex.RUnlock()

	offset := getNodeIndex(node)
	beforePoint := BoundaryPoint{Node: parent, Offset: offset}
	afterPoint := BoundaryPoint{Node: parent, Offset: offset + 1}

	// If (parent, offset) is before end and (parent, offset + 1) is after start, return true
	if CompareBoundaryPoints(&beforePoint, &r.end) == -1 && CompareBoundaryPoints(&afterPoint, &r.start) == 1 {
		return true
	}

	return false
}

// CloneRange creates a copy of this range
func (r *Range) CloneRange() *Range {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	clone := &Range{
		AbstractRangeImpl: AbstractRangeImpl{
			start: r.start,
			end:   r.end,
		},
		root: r.root,
	}
	return clone
}

// Detach does nothing (legacy method)
// Its functionality (disabling a Range object) was removed, but the method itself
// is preserved for compatibility.
func (r *Range) Detach() {
	// Do nothing per specification
}

// String implements the stringification behavior
// Returns the concatenated text content of all contained text nodes
func (r *Range) String() string {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	// If this's start node is this's end node and it is a Text node,
	// then return the substring of that Text node's data beginning at
	// this's start offset and ending at this's end offset
	if r.start.Node == r.end.Node {
		if textNode, ok := r.start.Node.(CharacterData); ok && r.start.Node.NodeType() == TextNode {
			data := textNode.Data()
			if r.start.Offset >= len(data) {
				return ""
			}
			endOffset := r.end.Offset
			if endOffset > len(data) {
				endOffset = len(data)
			}
			return data[r.start.Offset:endOffset]
		}
	}

	var result strings.Builder

	// If this's start node is a Text node, then append the substring of that node's data
	// from this's start offset until the end to s
	if textNode, ok := r.start.Node.(CharacterData); ok && r.start.Node.NodeType() == TextNode {
		data := textNode.Data()
		if r.start.Offset < len(data) {
			result.WriteString(data[r.start.Offset:])
		}
	}

	// Append the concatenation of the data of all Text nodes that are contained in this,
	// in tree order, to s
	r.appendContainedTextNodes(&result)

	// If this's end node is a Text node, then append the substring of that node's data
	// from its start until this's end offset to s
	if textNode, ok := r.end.Node.(CharacterData); ok && r.end.Node.NodeType() == TextNode && r.end.Node != r.start.Node {
		data := textNode.Data()
		if r.end.Offset > 0 && r.end.Offset <= len(data) {
			result.WriteString(data[:r.end.Offset])
		}
	}

	return result.String()
}

// Helper functions

// isInclusiveAncestor checks if ancestor is an inclusive ancestor of descendant
func isInclusiveAncestor(ancestor, descendant Node) bool {
	if ancestor == descendant {
		return true
	}
	return isAncestor(ancestor, descendant)
}

// appendContainedTextNodes appends text content from all contained text nodes
func (r *Range) appendContainedTextNodes(builder *strings.Builder) {
	// This is a simplified implementation
	// TODO: Implement proper "contained" node detection and tree order traversal
	// For now, this is a placeholder
}

// TODO: Implement the complex mutation methods:
// - DeleteContents()
// - ExtractContents()
// - CloneContents()
// - InsertNode()
// - SurroundContents()

// These will be implemented in the next phase as they require extensive
// DOM manipulation and fragment handling
