package dom

import (
	"sync"
)

// TreeWalker represents a DOM TreeWalker per WHATWG DOM specification
type TreeWalker struct {
	root        Node
	currentNode Node
	whatToShow  uint32
	filter      NodeFilter

	// Document reference
	document *Document

	// Mutex for thread safety
	mutex sync.RWMutex
}

// NewTreeWalker creates a new TreeWalker
func NewTreeWalker(root Node, whatToShow uint32, filter NodeFilter) *TreeWalker {
	// Get the document for this walker
	var doc *Document
	if root.NodeType() == DocumentNode {
		doc = root.(*Document)
	} else {
		doc = root.OwnerDocument()
	}

	return &TreeWalker{
		root:        root,
		currentNode: root,
		whatToShow:  whatToShow,
		filter:      combinedFilter(whatToShow, filter),
		document:    doc,
	}
}

// Root returns the root node of this walker
func (tw *TreeWalker) Root() Node {
	tw.mutex.RLock()
	defer tw.mutex.RUnlock()
	return tw.root
}

// WhatToShow returns the whatToShow bitmask
func (tw *TreeWalker) WhatToShow() uint32 {
	tw.mutex.RLock()
	defer tw.mutex.RUnlock()
	return tw.whatToShow
}

// Filter returns the NodeFilter
func (tw *TreeWalker) Filter() NodeFilter {
	tw.mutex.RLock()
	defer tw.mutex.RUnlock()
	return tw.filter
}

// CurrentNode returns the current node
func (tw *TreeWalker) CurrentNode() Node {
	tw.mutex.RLock()
	defer tw.mutex.RUnlock()
	return tw.currentNode
}

// SetCurrentNode sets the current node
func (tw *TreeWalker) SetCurrentNode(node Node) {
	tw.mutex.Lock()
	defer tw.mutex.Unlock()
	if node != nil {
		tw.currentNode = node
	}
}

// ParentNode moves to and returns the parent node that passes the filter
// Per WHATWG DOM spec Section 6.2
func (tw *TreeWalker) ParentNode() Node {
	tw.mutex.Lock()
	defer tw.mutex.Unlock()

	node := tw.currentNode

	// While node is non-null and is not this's root:
	for node != nil && node != tw.root {
		// Set node to node's parent.
		node = node.ParentNode()

		// If node is non-null and filtering node within this returns FILTER_ACCEPT,
		// then set this's current to node and return node.
		if node != nil && tw.filteringNodeWithin(node) == NodeFilterAccept {
			tw.currentNode = node
			return node
		}
	}

	// Return null.
	return nil
}

// FirstChild moves to and returns the first child that passes the filter
// Per WHATWG DOM spec Section 6.2
func (tw *TreeWalker) FirstChild() Node {
	tw.mutex.Lock()
	defer tw.mutex.Unlock()

	return tw.traverseChildren(tw.currentNode, "first")
}

// LastChild moves to and returns the last child that passes the filter
// Per WHATWG DOM spec Section 6.2
func (tw *TreeWalker) LastChild() Node {
	tw.mutex.Lock()
	defer tw.mutex.Unlock()

	return tw.traverseChildren(tw.currentNode, "last")
}

// NextSibling moves to and returns the next sibling that passes the filter
// Per WHATWG DOM spec Section 6.2
func (tw *TreeWalker) NextSibling() Node {
	tw.mutex.Lock()
	defer tw.mutex.Unlock()

	return tw.traverseSiblings(tw.currentNode, "next")
}

// PreviousSibling moves to and returns the previous sibling that passes the filter
// Per WHATWG DOM spec Section 6.2
func (tw *TreeWalker) PreviousSibling() Node {
	tw.mutex.Lock()
	defer tw.mutex.Unlock()

	return tw.traverseSiblings(tw.currentNode, "previous")
}

// PreviousNode moves to and returns the previous node in document order that passes the filter
// Per WHATWG DOM spec Section 6.2
func (tw *TreeWalker) PreviousNode() Node {
	tw.mutex.Lock()
	defer tw.mutex.Unlock()

	node := tw.currentNode

	// While node is not this's root:
	for node != tw.root {
		// Let sibling be node's previous sibling.
		sibling := node.PreviousSibling()

		// While sibling is non-null:
		for sibling != nil {
			// Set node to sibling.
			node = sibling

			// Let result be the result of filtering node within this.
			result := tw.filteringNodeWithin(node)

			// While result is not FILTER_REJECT and node has a child:
			for result != NodeFilterReject && node.HasChildNodes() {
				// Set node to node's last child.
				node = node.LastChild()

				// Set result to the result of filtering node within this.
				result = tw.filteringNodeWithin(node)
			}

			// If result is FILTER_ACCEPT, then set this's current to node and return node.
			if result == NodeFilterAccept {
				tw.currentNode = node
				return node
			}

			// Set sibling to node's previous sibling.
			sibling = node.PreviousSibling()
		}

		// If node is this's root or node's parent is null, then return null.
		if node == tw.root || node.ParentNode() == nil {
			return nil
		}

		// Set node to node's parent.
		node = node.ParentNode()

		// If the return value of filtering node within this is FILTER_ACCEPT,
		// then set this's current to node and return node.
		if tw.filteringNodeWithin(node) == NodeFilterAccept {
			tw.currentNode = node
			return node
		}
	}

	// Return null.
	return nil
}

// NextNode moves to and returns the next node in document order that passes the filter
// Per WHATWG DOM spec Section 6.2
func (tw *TreeWalker) NextNode() Node {
	tw.mutex.Lock()
	defer tw.mutex.Unlock()

	node := tw.currentNode
	result := NodeFilterAccept

	// While true:
	for {
		// While result is not FILTER_REJECT and node has a child:
		for result != NodeFilterReject && node.HasChildNodes() {
			// Set node to its first child.
			node = node.FirstChild()

			// Set result to the result of filtering node within this.
			result = tw.filteringNodeWithin(node)

			// If result is FILTER_ACCEPT, then set this's current to node and return node.
			if result == NodeFilterAccept {
				tw.currentNode = node
				return node
			}
		}

		// Let sibling be null.
		var sibling Node

		// Let temporary be node.
		temporary := node

		// While temporary is non-null:
		for temporary != nil {
			// If temporary is this's root, then return null.
			if temporary == tw.root {
				return nil
			}

			// Set sibling to temporary's next sibling.
			sibling = temporary.NextSibling()

			// If sibling is non-null, then set node to sibling and break.
			if sibling != nil {
				node = sibling
				break
			}

			// Set temporary to temporary's parent.
			temporary = temporary.ParentNode()
		}

		// Set result to the result of filtering node within this.
		result = tw.filteringNodeWithin(node)

		// If result is FILTER_ACCEPT, then set this's current to node and return node.
		if result == NodeFilterAccept {
			tw.currentNode = node
			return node
		}
	}
}

// traverseChildren implements the "traverse children" algorithm from WHATWG DOM spec
func (tw *TreeWalker) traverseChildren(walker Node, childType string) Node {
	// Let node be walker's current.
	node := walker

	// Set node to node's first child if type is first, and node's last child if type is last.
	if childType == "first" {
		node = node.FirstChild()
	} else {
		node = node.LastChild()
	}

	// While node is non-null:
	for node != nil {
		// Let result be the result of filtering node within walker.
		result := tw.filteringNodeWithin(node)

		// If result is FILTER_ACCEPT, then set walker's current to node and return node.
		if result == NodeFilterAccept {
			tw.currentNode = node
			return node
		}

		// If result is FILTER_SKIP:
		if result == NodeFilterSkip {
			// Let child be node's first child if type is first, and node's last child if type is last.
			var child Node
			if childType == "first" {
				child = node.FirstChild()
			} else {
				child = node.LastChild()
			}

			// If child is non-null, then set node to child and continue.
			if child != nil {
				node = child
				continue
			}
		}

		// While node is non-null:
		for node != nil {
			// Let sibling be node's next sibling if type is first, and node's previous sibling if type is last.
			var sibling Node
			if childType == "first" {
				sibling = node.NextSibling()
			} else {
				sibling = node.PreviousSibling()
			}

			// If sibling is non-null, then set node to sibling and break.
			if sibling != nil {
				node = sibling
				break
			}

			// Let parent be node's parent.
			parent := node.ParentNode()

			// If parent is null, walker's root, or walker's current, then return null.
			if parent == nil || parent == tw.root || parent == tw.currentNode {
				return nil
			}

			// Set node to parent.
			node = parent
		}
	}

	// Return null.
	return nil
}

// traverseSiblings implements the "traverse siblings" algorithm from WHATWG DOM spec
func (tw *TreeWalker) traverseSiblings(walker Node, siblingType string) Node {
	// Let node be walker's current.
	node := walker

	// If node is root, then return null.
	if node == tw.root {
		return nil
	}

	// While true:
	for {
		// Let sibling be node's next sibling if type is next, and node's previous sibling if type is previous.
		var sibling Node
		if siblingType == "next" {
			sibling = node.NextSibling()
		} else {
			sibling = node.PreviousSibling()
		}

		// While sibling is non-null:
		for sibling != nil {
			// Set node to sibling.
			node = sibling

			// Let result be the result of filtering node within walker.
			result := tw.filteringNodeWithin(node)

			// If result is FILTER_ACCEPT, then set walker's current to node and return node.
			if result == NodeFilterAccept {
				tw.currentNode = node
				return node
			}

			// Set sibling to node's first child if type is next, and node's last child if type is previous.
			if siblingType == "next" {
				sibling = node.FirstChild()
			} else {
				sibling = node.LastChild()
			}

			// If result is FILTER_REJECT or sibling is null, then set sibling to node's next sibling if type is next, and node's previous sibling if type is previous.
			if result == NodeFilterReject || sibling == nil {
				if siblingType == "next" {
					sibling = node.NextSibling()
				} else {
					sibling = node.PreviousSibling()
				}
			}
		}

		// Set node to node's parent.
		node = node.ParentNode()

		// If node is null or walker's root, then return null.
		if node == nil || node == tw.root {
			return nil
		}

		// If the return value of filtering node within walker is FILTER_ACCEPT, then return null.
		if tw.filteringNodeWithin(node) == NodeFilterAccept {
			return nil
		}
	}
}

// filteringNodeWithin implements the "filtering node within" algorithm
func (tw *TreeWalker) filteringNodeWithin(node Node) int {
	return tw.filter.AcceptNode(node)
}
