package dom

// NodeIterator represents a DOM NodeIterator per WHATWG DOM specification
type NodeIterator struct {
	root                       Node
	referenceNode              Node
	pointerBeforeReferenceNode bool
	whatToShow                 uint32
	filter                     NodeFilter

	// Document reference to track when nodes are removed
	document *Document

	// Active flag to avoid recursive invocations per spec
	activeFlag bool
}

// NewNodeIterator creates a new NodeIterator
func NewNodeIterator(root Node, whatToShow uint32, filter NodeFilter) *NodeIterator {
	// Get the document for this iterator
	var doc *Document
	if root.NodeType() == DocumentNode {
		doc = root.(*Document)
	} else {
		doc = root.OwnerDocument()
	}

	iterator := &NodeIterator{
		root:                       root,
		referenceNode:              root,
		pointerBeforeReferenceNode: true,
		whatToShow:                 whatToShow,
		filter:                     filter, // Store raw filter, not combined
		document:                   doc,
		activeFlag:                 false, // Initially unset per spec
	}

	// Register with document for node removal notifications
	if doc != nil {
		// TODO: Register for node removal notifications when mutation system supports it
	}

	return iterator
}

// Root returns the root node of this iterator
func (ni *NodeIterator) Root() Node {
	return ni.root
}

// WhatToShow returns the whatToShow bitmask
func (ni *NodeIterator) WhatToShow() uint32 {
	return ni.whatToShow
}

// Filter returns the NodeFilter
func (ni *NodeIterator) Filter() NodeFilter {
	return ni.filter
}

// ReferenceNode returns the current reference node
func (ni *NodeIterator) ReferenceNode() Node {
	return ni.referenceNode
}

// PointerBeforeReferenceNode returns true if the pointer is before the reference node
func (ni *NodeIterator) PointerBeforeReferenceNode() bool {
	return ni.pointerBeforeReferenceNode
}

// filterNode implements the spec's "filter node" algorithm
func (ni *NodeIterator) filterNode(node Node) int {
	// Step 1: If active flag is set, throw InvalidStateError
	if ni.activeFlag {
		panic(&DOMException{
			name:    "InvalidStateError",
			message: "filter invoked while NodeIterator is active",
		})
	}

	// Step 2: Let n be node's nodeType attribute value − 1
	n := node.NodeType() - 1

	// Step 3: If the nth bit (where 0 is the least significant bit) of whatToShow is not set, return SKIP
	if ni.whatToShow&(1<<uint(n)) == 0 {
		return NodeFilterSkip
	}

	// Step 4: If filter is null, return ACCEPT
	if ni.filter == nil {
		return NodeFilterAccept
	}

	// Step 5: Set traverser's active flag
	ni.activeFlag = true

	// Step 6: Let result be the return value of calling filter's acceptNode with node
	// If this throws an exception, unset active flag and rethrow
	var result int
	func() {
		defer func() {
			// Step 7: Unset traverser's active flag
			ni.activeFlag = false
			// Rethrow any panic
			if r := recover(); r != nil {
				panic(r)
			}
		}()
		result = ni.filter.AcceptNode(node)
	}()

	// Step 8: Return result
	return result
}

// traverse implements the spec's traverse algorithm
func (ni *NodeIterator) traverse(direction string) Node {
	node := ni.referenceNode
	beforeNode := ni.pointerBeforeReferenceNode

	for {
		// Branch on direction per spec
		moved := false
		switch direction {
		case "next":
			if !beforeNode {
				// Set node to the first node following node in iterator's iterator collection
				next := ni.getNextNodeInDocumentOrder(node)
				if next == nil || !ni.isNodeWithinRoot(next) {
					return nil
				}
				node = next
				moved = true
			} else {
				// If beforeNode is true, then set it to false
				beforeNode = false
			}
		case "previous":
			if beforeNode {
				// Set node to the first node preceding node in iterator's iterator collection
				prev := ni.getPreviousNodeInDocumentOrder(node)
				if prev == nil || !ni.isNodeWithinRoot(prev) {
					return nil
				}
				node = prev
				moved = true
			} else {
				// If beforeNode is false, then set it to true
				beforeNode = true
			}
		}

		// Let result be the result of filtering node within iterator
		result := ni.filterNode(node)
		if result == NodeFilterAccept {
			break
		}

		// If we only toggled the pointer position but didn't move to a new node,
		// and the current node was rejected/skipped, we need to continue traversing
		if !moved && result != NodeFilterAccept {
			continue
		}

		// Handle FILTER_REJECT by skipping the node and its descendants
		if moved && result == NodeFilterReject && direction == "next" {
			// Skip descendants of rejected node by finding next node that's not a descendant
			next := ni.getNextNodeSkippingDescendants(node)
			if next == nil || !ni.isNodeWithinRoot(next) {
				return nil
			}
			node = next
			// Re-filter the new node
			result = ni.filterNode(node)
			if result == NodeFilterAccept {
				break
			}
		}
	}

	// Set iterator's reference to node
	ni.referenceNode = node
	// Set iterator's pointer before reference to beforeNode
	ni.pointerBeforeReferenceNode = beforeNode
	// Return node
	return node
}

// NextNode returns the next node in document order that passes the filter
func (ni *NodeIterator) NextNode() Node {
	return ni.traverse("next")
}

// PreviousNode returns the previous node in document order that passes the filter
func (ni *NodeIterator) PreviousNode() Node {
	return ni.traverse("previous")
}

// Detach makes the iterator inactive (legacy method, does nothing per spec)
func (ni *NodeIterator) Detach() {
	// Per WHATWG DOM spec: "The detach() method must do nothing"
	// This method is kept for compatibility but has no effect
}

// isNodeWithinRoot checks if a node is within the iterator's root
func (ni *NodeIterator) isNodeWithinRoot(node Node) bool {
	if node == nil {
		return false
	}

	// Check if node is the root
	if node == ni.root {
		return true
	}

	// Check if node is a descendant of root
	current := node.ParentNode()
	for current != nil {
		if current == ni.root {
			return true
		}
		current = current.ParentNode()
	}

	return false
}

// getNextNodeInDocumentOrder returns the next node in document order
func (ni *NodeIterator) getNextNodeInDocumentOrder(node Node) Node {
	if node == nil {
		return nil
	}

	// First try to go to first child
	if node.HasChildNodes() {
		return node.FirstChild()
	}

	// Then try next sibling
	if sibling := node.NextSibling(); sibling != nil {
		return sibling
	}

	// Go up the tree looking for an ancestor with a next sibling
	current := node.ParentNode()
	for current != nil && current != ni.root {
		if sibling := current.NextSibling(); sibling != nil {
			return sibling
		}
		current = current.ParentNode()
	}

	return nil
}

// getPreviousNodeInDocumentOrder returns the previous node in document order
func (ni *NodeIterator) getPreviousNodeInDocumentOrder(node Node) Node {
	if node == nil {
		return nil
	}

	// Try previous sibling first
	if sibling := node.PreviousSibling(); sibling != nil {
		// Go to the deepest last child of the sibling
		current := sibling
		for current.HasChildNodes() {
			current = current.LastChild()
		}
		return current
	}

	// Go to parent (including the root)
	parent := node.ParentNode()
	if parent != nil && ni.isNodeWithinRoot(parent) {
		return parent
	}

	return nil
}

// preRemoveNode implements the NodeIterator pre-remove steps per WHATWG DOM spec
func (ni *NodeIterator) preRemoveNode(toBeRemovedNode Node) {
	// Step 1: If toBeRemovedNode is not an inclusive ancestor of nodeIterator's reference,
	// or toBeRemovedNode is nodeIterator's root, then return
	if !ni.isInclusiveAncestor(toBeRemovedNode, ni.referenceNode) || toBeRemovedNode == ni.root {
		return
	}

	// Step 2: If nodeIterator's pointer before reference is true
	if ni.pointerBeforeReferenceNode {
		// Step 2.1: Let next be toBeRemovedNode's first following node that is an
		// inclusive descendant of nodeIterator's root and is not an inclusive descendant
		// of toBeRemovedNode, and null if there is no such node
		next := ni.getFirstFollowingNode(toBeRemovedNode)

		// Step 2.2: If next is non-null, then set nodeIterator's reference to next and return
		if next != nil {
			ni.referenceNode = next
			return
		}

		// Step 2.3: Otherwise, set nodeIterator's pointer before reference to false
		ni.pointerBeforeReferenceNode = false
		// Steps are not terminated here - continue to step 3
	}

	// Step 3: Set nodeIterator's reference to toBeRemovedNode's parent, if
	// toBeRemovedNode's previous sibling is null, and to the inclusive descendant
	// of toBeRemovedNode's previous sibling that appears last in tree order otherwise
	if toBeRemovedNode.PreviousSibling() == nil {
		ni.referenceNode = toBeRemovedNode.ParentNode()
	} else {
		// Find the last inclusive descendant of the previous sibling
		ni.referenceNode = ni.getLastInclusiveDescendant(toBeRemovedNode.PreviousSibling())
	}
}

// isInclusiveAncestor checks if ancestor is an inclusive ancestor of node
func (ni *NodeIterator) isInclusiveAncestor(ancestor, node Node) bool {
	if ancestor == node {
		return true
	}
	current := node.ParentNode()
	for current != nil {
		if current == ancestor {
			return true
		}
		current = current.ParentNode()
	}
	return false
}

// getFirstFollowingNode returns the first following node that is an inclusive
// descendant of root and not an inclusive descendant of toBeRemoved
func (ni *NodeIterator) getFirstFollowingNode(toBeRemoved Node) Node {
	current := toBeRemoved

	// Try next sibling first
	if sibling := current.NextSibling(); sibling != nil {
		return sibling
	}

	// Go up the tree looking for an ancestor with a next sibling
	current = current.ParentNode()
	for current != nil && current != ni.root {
		if sibling := current.NextSibling(); sibling != nil {
			return sibling
		}
		current = current.ParentNode()
	}

	return nil
}

// getLastInclusiveDescendant returns the last inclusive descendant in tree order
func (ni *NodeIterator) getLastInclusiveDescendant(node Node) Node {
	current := node
	for current.HasChildNodes() {
		current = current.LastChild()
	}
	return current
}

// getNextNodeSkippingDescendants returns the next node in document order, skipping descendants
func (ni *NodeIterator) getNextNodeSkippingDescendants(node Node) Node {
	if node == nil {
		return nil
	}

	// Try next sibling
	if sibling := node.NextSibling(); sibling != nil {
		return sibling
	}

	// Go up the tree looking for an ancestor with a next sibling
	current := node.ParentNode()
	for current != nil && current != ni.root {
		if sibling := current.NextSibling(); sibling != nil {
			return sibling
		}
		current = current.ParentNode()
	}

	return nil
}
