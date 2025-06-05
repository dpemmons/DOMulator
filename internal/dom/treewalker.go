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
func (tw *TreeWalker) ParentNode() Node {
	tw.mutex.Lock()
	defer tw.mutex.Unlock()

	current := tw.currentNode.ParentNode()
	for current != nil && current != tw.root {
		result := tw.filter.AcceptNode(current)
		if result == NodeFilterAccept {
			tw.currentNode = current
			return current
		}
		current = current.ParentNode()
	}

	// Check the root itself
	if current == tw.root {
		result := tw.filter.AcceptNode(current)
		if result == NodeFilterAccept {
			tw.currentNode = current
			return current
		}
	}

	return nil
}

// FirstChild moves to and returns the first child that passes the filter
func (tw *TreeWalker) FirstChild() Node {
	tw.mutex.Lock()
	defer tw.mutex.Unlock()

	return tw.traverseChildren(tw.currentNode, true)
}

// LastChild moves to and returns the last child that passes the filter
func (tw *TreeWalker) LastChild() Node {
	tw.mutex.Lock()
	defer tw.mutex.Unlock()

	return tw.traverseChildren(tw.currentNode, false)
}

// PreviousSibling moves to and returns the previous sibling that passes the filter
func (tw *TreeWalker) PreviousSibling() Node {
	tw.mutex.Lock()
	defer tw.mutex.Unlock()

	return tw.traverseSiblings(tw.currentNode, false)
}

// NextSibling moves to and returns the next sibling that passes the filter
func (tw *TreeWalker) NextSibling() Node {
	tw.mutex.Lock()
	defer tw.mutex.Unlock()

	return tw.traverseSiblings(tw.currentNode, true)
}

// PreviousNode moves to and returns the previous node in document order that passes the filter
func (tw *TreeWalker) PreviousNode() Node {
	tw.mutex.Lock()
	defer tw.mutex.Unlock()

	current := tw.currentNode
	if current == tw.root {
		return nil
	}

	// Get the previous node in document order
	var prev Node

	// Try previous sibling
	if sibling := current.PreviousSibling(); sibling != nil {
		// Go to the deepest last descendant of the sibling
		prev = tw.getLastDescendantOrSelf(sibling)
	} else {
		// Go to parent
		prev = current.ParentNode()
	}

	// Find the previous acceptable node
	for prev != nil && tw.isNodeWithinRoot(prev) {
		result := tw.filter.AcceptNode(prev)
		if result == NodeFilterAccept {
			tw.currentNode = prev
			return prev
		}

		// Continue searching backwards
		if prev == tw.root {
			break
		}

		// Get next previous node
		if prevSibling := prev.PreviousSibling(); prevSibling != nil {
			prev = tw.getLastDescendantOrSelf(prevSibling)
		} else {
			prev = prev.ParentNode()
		}
	}

	return nil
}

// NextNode moves to and returns the next node in document order that passes the filter
func (tw *TreeWalker) NextNode() Node {
	tw.mutex.Lock()
	defer tw.mutex.Unlock()

	current := tw.currentNode

	// Try descendants first
	if current.HasChildNodes() {
		child := current.FirstChild()
		for child != nil {
			if tw.isNodeWithinRoot(child) {
				result := tw.filter.AcceptNode(child)
				if result == NodeFilterAccept {
					tw.currentNode = child
					return child
				}
				if result == NodeFilterSkip && child.HasChildNodes() {
					// Skip this node but check its children
					subChild := child.FirstChild()
					for subChild != nil {
						if tw.isNodeWithinRoot(subChild) {
							result := tw.filter.AcceptNode(subChild)
							if result == NodeFilterAccept {
								tw.currentNode = subChild
								return subChild
							}
						}
						subChild = subChild.NextSibling()
					}
				}
			}
			child = child.NextSibling()
		}
	}

	// Try siblings and ancestors' siblings
	node := current
	for node != nil && node != tw.root {
		if sibling := node.NextSibling(); sibling != nil && tw.isNodeWithinRoot(sibling) {
			result := tw.filter.AcceptNode(sibling)
			if result == NodeFilterAccept {
				tw.currentNode = sibling
				return sibling
			}
			// Check sibling's descendants
			if result == NodeFilterSkip {
				descendant := tw.getFirstDescendant(sibling)
				for descendant != nil {
					if tw.isNodeWithinRoot(descendant) {
						result := tw.filter.AcceptNode(descendant)
						if result == NodeFilterAccept {
							tw.currentNode = descendant
							return descendant
						}
					}
					descendant = tw.getNextInSubtree(descendant, sibling)
				}
			}
			// Continue with next sibling
			siblingNext := sibling.NextSibling()
			for siblingNext != nil && tw.isNodeWithinRoot(siblingNext) {
				result := tw.filter.AcceptNode(siblingNext)
				if result == NodeFilterAccept {
					tw.currentNode = siblingNext
					return siblingNext
				}
				siblingNext = siblingNext.NextSibling()
			}
		}
		node = node.ParentNode()
	}

	return nil
}

// traverseChildren traverses children of a node
func (tw *TreeWalker) traverseChildren(node Node, forward bool) Node {
	if !node.HasChildNodes() {
		return nil
	}

	var child Node
	if forward {
		child = node.FirstChild()
	} else {
		child = node.LastChild()
	}

	for child != nil {
		if tw.isNodeWithinRoot(child) {
			result := tw.filter.AcceptNode(child)
			if result == NodeFilterAccept {
				tw.currentNode = child
				return child
			}
			// For SKIP, check children of this node
			if result == NodeFilterSkip {
				grandchild := tw.traverseChildren(child, forward)
				if grandchild != nil {
					return grandchild
				}
			}
		}

		if forward {
			child = child.NextSibling()
		} else {
			child = child.PreviousSibling()
		}
	}

	return nil
}

// traverseSiblings traverses siblings of a node
func (tw *TreeWalker) traverseSiblings(node Node, forward bool) Node {
	var sibling Node
	if forward {
		sibling = node.NextSibling()
	} else {
		sibling = node.PreviousSibling()
	}

	for sibling != nil {
		if tw.isNodeWithinRoot(sibling) {
			result := tw.filter.AcceptNode(sibling)
			if result == NodeFilterAccept {
				tw.currentNode = sibling
				return sibling
			}
			// For SKIP, check children of this sibling
			if result == NodeFilterSkip {
				child := tw.traverseChildren(sibling, forward)
				if child != nil {
					return child
				}
			}
		}

		if forward {
			sibling = sibling.NextSibling()
		} else {
			sibling = sibling.PreviousSibling()
		}
	}

	return nil
}

// isNodeWithinRoot checks if a node is within the walker's root
func (tw *TreeWalker) isNodeWithinRoot(node Node) bool {
	if node == nil {
		return false
	}

	// Check if node is the root
	if node == tw.root {
		return true
	}

	// Check if node is a descendant of root
	current := node.ParentNode()
	for current != nil {
		if current == tw.root {
			return true
		}
		current = current.ParentNode()
	}

	return false
}

// getFirstDescendant gets the first descendant of a node
func (tw *TreeWalker) getFirstDescendant(node Node) Node {
	if !node.HasChildNodes() {
		return nil
	}
	return node.FirstChild()
}

// getLastDescendantOrSelf gets the deepest last descendant or the node itself
func (tw *TreeWalker) getLastDescendantOrSelf(node Node) Node {
	current := node
	for current.HasChildNodes() {
		current = current.LastChild()
	}
	return current
}

// getNextInSubtree gets the next node within a subtree
func (tw *TreeWalker) getNextInSubtree(node, root Node) Node {
	if node == root {
		return nil
	}

	// Try first child
	if node.HasChildNodes() {
		return node.FirstChild()
	}

	// Try next sibling
	if sibling := node.NextSibling(); sibling != nil {
		return sibling
	}

	// Go up to find ancestor with next sibling
	current := node.ParentNode()
	for current != nil && current != root {
		if sibling := current.NextSibling(); sibling != nil {
			return sibling
		}
		current = current.ParentNode()
	}

	return nil
}

// getPreviousNodeInSubtree gets the previous node within a subtree
func (tw *TreeWalker) getPreviousNodeInSubtree(node, root Node) Node {
	if node == root {
		return nil
	}

	// Try previous sibling and its descendants
	if sibling := node.PreviousSibling(); sibling != nil {
		return tw.getLastDescendantOrSelf(sibling)
	}

	// Go to parent
	parent := node.ParentNode()
	if parent != nil && parent != root {
		return parent
	}

	return nil
}
