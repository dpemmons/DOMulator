package dom

import (
	"github.com/dpemmons/DOMulator/internal/dom/ranges"
)

// DOMRange is the live range implementation
type DOMRange = ranges.Range

// DOMStaticRange is the static range implementation
type DOMStaticRange = ranges.StaticRange

// DOMStaticRangeInit is the static range initialization parameters
type DOMStaticRangeInit struct {
	StartContainer Node
	StartOffset    int
	EndContainer   Node
	EndOffset      int
}

// DOMAbstractRange is the abstract range interface
type DOMAbstractRange = ranges.AbstractRange

// Range constants for compareBoundaryPoints
const (
	START_TO_START = ranges.START_TO_START
	START_TO_END   = ranges.START_TO_END
	END_TO_END     = ranges.END_TO_END
	END_TO_START   = ranges.END_TO_START
)

// CreateStaticRange creates a new static range with the given parameters
func (d *Document) CreateStaticRange(init DOMStaticRangeInit) (*DOMStaticRange, error) {
	// Convert the init parameters to use proper DOM types
	rangeInit := ranges.StaticRangeInit{
		StartContainer: nodeToRangeNode(init.StartContainer),
		StartOffset:    init.StartOffset,
		EndContainer:   nodeToRangeNode(init.EndContainer),
		EndOffset:      init.EndOffset,
	}
	return ranges.NewStaticRange(rangeInit)
}

// RangeAdapter adapts the ranges.Node interface to our DOM Node interface
type RangeAdapter struct {
	node Node
}

// Implement the ranges.Node interface methods by delegating to the DOM Node
func (ra *RangeAdapter) NodeType() int {
	return int(ra.node.NodeType())
}

func (ra *RangeAdapter) ParentNode() ranges.Node {
	parent := ra.node.ParentNode()
	if parent == nil {
		return nil
	}
	return &RangeAdapter{node: parent}
}

func (ra *RangeAdapter) ChildNodes() ranges.NodeList {
	children := ra.node.ChildNodes()
	return &RangeNodeListAdapter{nodeList: children}
}

func (ra *RangeAdapter) OwnerDocument() ranges.Document {
	doc := ra.node.OwnerDocument()
	if doc == nil {
		return nil
	}
	return &RangeDocumentAdapter{document: doc}
}

func (ra *RangeAdapter) NextSibling() ranges.Node {
	sibling := ra.node.NextSibling()
	if sibling == nil {
		return nil
	}
	return &RangeAdapter{node: sibling}
}

func (ra *RangeAdapter) PreviousSibling() ranges.Node {
	sibling := ra.node.PreviousSibling()
	if sibling == nil {
		return nil
	}
	return &RangeAdapter{node: sibling}
}

func (ra *RangeAdapter) Length() int {
	return ra.node.Length()
}

// RangeNodeListAdapter adapts NodeList to ranges.NodeList
type RangeNodeListAdapter struct {
	nodeList *NodeList
}

func (rna *RangeNodeListAdapter) Length() int {
	return rna.nodeList.Length()
}

func (rna *RangeNodeListAdapter) Item(index int) ranges.Node {
	node := rna.nodeList.Item(index)
	if node == nil {
		return nil
	}
	return &RangeAdapter{node: node}
}

// RangeDocumentAdapter adapts Document to ranges.Document
type RangeDocumentAdapter struct {
	document *Document
}

func (rda *RangeDocumentAdapter) NodeType() int {
	return int(rda.document.NodeType())
}

func (rda *RangeDocumentAdapter) ParentNode() ranges.Node {
	parent := rda.document.ParentNode()
	if parent == nil {
		return nil
	}
	return &RangeAdapter{node: parent}
}

func (rda *RangeDocumentAdapter) ChildNodes() ranges.NodeList {
	children := rda.document.ChildNodes()
	return &RangeNodeListAdapter{nodeList: children}
}

func (rda *RangeDocumentAdapter) OwnerDocument() ranges.Document {
	// A document is its own owner document
	return rda
}

func (rda *RangeDocumentAdapter) NextSibling() ranges.Node {
	sibling := rda.document.NextSibling()
	if sibling == nil {
		return nil
	}
	return &RangeAdapter{node: sibling}
}

func (rda *RangeDocumentAdapter) PreviousSibling() ranges.Node {
	sibling := rda.document.PreviousSibling()
	if sibling == nil {
		return nil
	}
	return &RangeAdapter{node: sibling}
}

func (rda *RangeDocumentAdapter) Length() int {
	return rda.document.Length()
}

// Helper function to convert DOM nodes to range nodes
func nodeToRangeNode(node Node) ranges.Node {
	if node == nil {
		return nil
	}
	if doc, ok := node.(*Document); ok {
		return &RangeDocumentAdapter{document: doc}
	}
	return &RangeAdapter{node: node}
}

// Helper function to convert range nodes back to DOM nodes
func rangeNodeToNode(rangeNode ranges.Node) Node {
	if rangeNode == nil {
		return nil
	}
	if adapter, ok := rangeNode.(*RangeAdapter); ok {
		return adapter.node
	}
	if docAdapter, ok := rangeNode.(*RangeDocumentAdapter); ok {
		return docAdapter.document
	}
	return nil
}

// TODO: JavaScript bindings for DOMRange and DOMStaticRange will be implemented
// when we add full JavaScript integration
