package dom

import (
	"github.com/dop251/goja"
)

// NodeType represents the type of a Node.
type NodeType int

const (
	ElementNode               NodeType = 1
	AttributeNode             NodeType = 2
	TextNode                  NodeType = 3
	CDataSectionNode          NodeType = 4
	EntityReferenceNode       NodeType = 5
	EntityNode                NodeType = 6
	ProcessingInstructionNode NodeType = 7
	CommentNode               NodeType = 8
	DocumentNode              NodeType = 9
	DocumentTypeNode          NodeType = 10
	DocumentFragmentNode      NodeType = 11
	NotationNode              NodeType = 12
)

// Node represents a single node in the DOM tree.
type Node interface {
	NodeType() NodeType
	NodeName() string
	NodeValue() string
	OwnerDocument() *Document

	ParentNode() Node
	ChildNodes() NodeList
	FirstChild() Node
	LastChild() Node
	PreviousSibling() Node
	NextSibling() Node

	AppendChild(child Node) Node
	RemoveChild(child Node) Node
	InsertBefore(newChild, refChild Node) Node
	ReplaceChild(newChild, oldChild Node) Node

	CloneNode(deep bool) Node

	// Length method as defined in DOM specification
	Length() int
	IsEmpty() bool

	// Event Listener methods
	AddEventListener(eventType string, listener func(Event))
	RemoveEventListener(eventType string, listener func(Event))
	DispatchEvent(event Event) bool

	// Mutation Observer methods
	registerMutationObserver(observer *MutationObserver, options MutationObserverInit)
	unregisterMutationObserver(observer *MutationObserver)

	// Internal methods for DOM manipulation and JS binding
	setParent(parent Node)
	setOwnerDocument(doc *Document)
	toJS(vm *goja.Runtime) goja.Value
}

// NodeList represents a collection of nodes.
type NodeList []Node

// nodeImpl provides a common base for all Node types.
type nodeImpl struct {
	nodeType       NodeType
	nodeName       string
	nodeValue      string
	parentNode     Node
	childNodes     NodeList
	ownerDocument  *Document
	self           Node // Reference to the containing Node
	eventListeners map[string][]func(Event)
}

func (n *nodeImpl) AddEventListener(eventType string, listener func(Event)) {
	if n.eventListeners == nil {
		n.eventListeners = make(map[string][]func(Event))
	}
	n.eventListeners[eventType] = append(n.eventListeners[eventType], listener)
}

func (n *nodeImpl) RemoveEventListener(eventType string, listener func(Event)) {
	if n.eventListeners == nil {
		return
	}
	listeners := n.eventListeners[eventType]
	for i, l := range listeners {
		// This is a simplified comparison. In a real DOM, comparing function pointers is complex.
		// For now, we'll remove the first matching listener.
		// A more robust solution might involve unique listener IDs or a different storage mechanism.
		if &l == &listener { // Compare addresses of the functions
			n.eventListeners[eventType] = append(listeners[:i], listeners[i+1:]...)
			return
		}
	}
}

func (n *nodeImpl) getEventListeners() map[string][]func(Event) {
	if n.eventListeners == nil {
		n.eventListeners = make(map[string][]func(Event))
	}
	return n.eventListeners
}

func (n *nodeImpl) DispatchEvent(event Event) bool {
	baseEvent, ok := event.(*BaseEvent)
	if !ok {
		return false
	}

	// 1. Determine propagation path
	var path []Node
	curr := n.self
	for curr != nil {
		path = append(path, curr)
		curr = curr.ParentNode()
	}

	// 2. Capturing Phase (from parent to target)
	baseEvent.SetEventPhase(CAPTURING_PHASE)
	for i := len(path) - 1; i >= 0; i-- {
		node := path[i]
		baseEvent.SetCurrentTarget(node)
		// Access event listeners through the interface to handle different node types
		if nodeWithListeners, ok := node.(interface {
			getEventListeners() map[string][]func(Event)
		}); ok {
			if handlers, exists := nodeWithListeners.getEventListeners()[baseEvent.Type()]; exists {
				for _, handler := range handlers {
					handler(baseEvent)
					if baseEvent.IsPropagationStopped() {
						return !baseEvent.DefaultPrevented()
					}
				}
			}
		}
		if node == n.self { // Stop capturing phase at target
			break
		}
	}

	// 3. At Target Phase
	baseEvent.SetEventPhase(AT_TARGET)
	baseEvent.SetCurrentTarget(n.self)
	if handlers, ok := n.eventListeners[baseEvent.Type()]; ok {
		for _, handler := range handlers {
			handler(baseEvent)
			if baseEvent.IsPropagationStopped() {
				return !baseEvent.DefaultPrevented()
			}
		}
	}

	// 4. Bubbling Phase (from target to parent, if event bubbles)
	if baseEvent.Bubbles() {
		baseEvent.SetEventPhase(BUBBLING_PHASE)
		for i := 1; i < len(path); i++ { // Start from parent of target
			node := path[i]
			baseEvent.SetCurrentTarget(node)
			// Access event listeners through the interface to handle different node types
			if nodeWithListeners, ok := node.(interface {
				getEventListeners() map[string][]func(Event)
			}); ok {
				if handlers, exists := nodeWithListeners.getEventListeners()[baseEvent.Type()]; exists {
					for _, handler := range handlers {
						handler(baseEvent)
						if baseEvent.IsPropagationStopped() {
							return !baseEvent.DefaultPrevented()
						}
					}
				}
			}
		}
	}

	return !baseEvent.DefaultPrevented()
}

func (n *nodeImpl) NodeType() NodeType {
	return n.nodeType
}

func (n *nodeImpl) NodeName() string {
	return n.nodeName
}

func (n *nodeImpl) NodeValue() string {
	return n.nodeValue
}

func (n *nodeImpl) OwnerDocument() *Document {
	return n.ownerDocument
}

func (n *nodeImpl) ParentNode() Node {
	return n.parentNode
}

func (n *nodeImpl) ChildNodes() NodeList {
	return n.childNodes
}

func (n *nodeImpl) FirstChild() Node {
	if len(n.childNodes) > 0 {
		return n.childNodes[0]
	}
	return nil
}

func (n *nodeImpl) LastChild() Node {
	if len(n.childNodes) > 0 {
		return n.childNodes[len(n.childNodes)-1]
	}
	return nil
}

func (n *nodeImpl) PreviousSibling() Node {
	if n.parentNode == nil {
		return nil
	}
	// Iterate through the parent's children to find n.self and its previous sibling
	for i, child := range n.parentNode.ChildNodes() {
		if child == n.self {
			if i > 0 {
				return n.parentNode.ChildNodes()[i-1]
			}
			return nil
		}
	}
	return nil
}

func (n *nodeImpl) NextSibling() Node {
	if n.parentNode == nil {
		return nil
	}
	// Iterate through the parent's children to find n.self and its next sibling
	for i, child := range n.parentNode.ChildNodes() {
		if child == n.self {
			if i < len(n.parentNode.ChildNodes())-1 {
				return n.parentNode.ChildNodes()[i+1]
			}
			return nil
		}
	}
	return nil
}

func (n *nodeImpl) AppendChild(child Node) Node {
	if child == nil {
		panic(NewNotFoundError("child cannot be null"))
	}

	// Validate hierarchy constraints
	if err := n.validateHierarchy(child); err != nil {
		panic(err)
	}

	// Handle DocumentFragment insertion
	if child.NodeType() == DocumentFragmentNode {
		fragment := child
		// Move all children from the fragment to this node
		for len(fragment.ChildNodes()) > 0 {
			fragmentChild := fragment.ChildNodes()[0]
			fragment.RemoveChild(fragmentChild)
			n.AppendChild(fragmentChild)
		}
		return fragment
	}

	// If the child already has a parent, remove it from that parent first
	if child.ParentNode() != nil {
		child.ParentNode().RemoveChild(child)
	}
	n.childNodes = append(n.childNodes, child)
	child.setParent(n.self)
	child.setOwnerDocument(n.ownerDocument)

	// Increment modification time
	if n.ownerDocument != nil {
		n.ownerDocument.incrementModificationTime()
	}

	return child
}

func (n *nodeImpl) RemoveChild(child Node) Node {
	if child == nil {
		panic(NewNotFoundError("child cannot be null"))
	}

	// Check if child is actually a child of this node
	if child.ParentNode() != n.self {
		panic(NewNotFoundError("Node is not a child of this parent"))
	}

	for i, c := range n.childNodes {
		if c == child {
			n.childNodes = append(n.childNodes[:i], n.childNodes[i+1:]...)
			child.setParent(nil)

			// Increment modification time
			if n.ownerDocument != nil {
				n.ownerDocument.incrementModificationTime()
			}

			return child
		}
	}

	// This should never happen given the parent check above, but just in case
	panic(NewNotFoundError("Child not found in parent's child list"))
}

func (n *nodeImpl) InsertBefore(newChild, refChild Node) Node {
	if newChild == nil {
		panic(NewNotFoundError("newChild cannot be null"))
	}

	if refChild == nil {
		return n.AppendChild(newChild)
	}

	// Validate that refChild is actually a child of this node
	if refChild.ParentNode() != n.self {
		panic(NewNotFoundError("refChild is not a child of this node"))
	}

	// Validate hierarchy constraints
	if err := n.validateHierarchy(newChild); err != nil {
		panic(err)
	}

	// Handle DocumentFragment insertion
	if newChild.NodeType() == DocumentFragmentNode {
		fragment := newChild
		// Find the position of refChild
		refIndex := -1
		for i, c := range n.childNodes {
			if c == refChild {
				refIndex = i
				break
			}
		}

		if refIndex == -1 {
			return nil // refChild not found
		}

		// Insert all children from the fragment before refChild
		for len(fragment.ChildNodes()) > 0 {
			fragmentChild := fragment.ChildNodes()[0]
			fragment.RemoveChild(fragmentChild)
			n.childNodes = append(n.childNodes[:refIndex], append(NodeList{fragmentChild}, n.childNodes[refIndex:]...)...)
			fragmentChild.setParent(n.self)
			fragmentChild.setOwnerDocument(n.ownerDocument)
			refIndex++ // Increment position for next insertion
		}

		// Increment modification time
		if n.ownerDocument != nil {
			n.ownerDocument.incrementModificationTime()
		}

		return fragment
	}

	// Special case: if newChild is already a child of this parent, we need to move it
	if newChild.ParentNode() == n.self {
		// Find current position of newChild
		var currentIndex = -1
		for i, c := range n.childNodes {
			if c == newChild {
				currentIndex = i
				break
			}
		}

		// Find target position (before refChild)
		var targetIndex = -1
		for i, c := range n.childNodes {
			if c == refChild {
				targetIndex = i
				break
			}
		}

		if currentIndex == -1 || targetIndex == -1 {
			return nil
		}

		// Remove from current position first
		n.childNodes = append(n.childNodes[:currentIndex], n.childNodes[currentIndex+1:]...)

		// Insert at the original target index without adjustment
		// This matches DOM behavior where position is calculated relative to original state
		n.childNodes = append(n.childNodes[:targetIndex], append(NodeList{newChild}, n.childNodes[targetIndex:]...)...)
		return newChild
	}

	// Remove from old parent if it has one
	if newChild.ParentNode() != nil {
		newChild.ParentNode().RemoveChild(newChild)
	}

	for i, c := range n.childNodes {
		if c == refChild {
			n.childNodes = append(n.childNodes[:i], append(NodeList{newChild}, n.childNodes[i:]...)...)
			newChild.setParent(n.self)
			newChild.setOwnerDocument(n.ownerDocument)

			// Increment modification time
			if n.ownerDocument != nil {
				n.ownerDocument.incrementModificationTime()
			}

			return newChild
		}
	}
	return nil // refChild not found
}

func (n *nodeImpl) ReplaceChild(newChild, oldChild Node) Node {
	if newChild == nil {
		panic(NewNotFoundError("newChild cannot be null"))
	}

	if oldChild == nil {
		panic(NewNotFoundError("oldChild cannot be null"))
	}

	// Validate that oldChild is actually a child of this node
	if oldChild.ParentNode() != n.self {
		panic(NewNotFoundError("oldChild is not a child of this node"))
	}

	// Validate hierarchy constraints for newChild, excluding oldChild from uniqueness checks
	if err := n.validateHierarchyWithExclusion(newChild, oldChild); err != nil {
		panic(err)
	}

	// Handle DocumentFragment replacement
	if newChild.NodeType() == DocumentFragmentNode {
		fragment := newChild
		// Find the position of oldChild
		oldIndex := -1
		for i, c := range n.childNodes {
			if c == oldChild {
				oldIndex = i
				break
			}
		}

		if oldIndex == -1 {
			return nil // oldChild not found
		}

		// Remove the old child first
		n.childNodes = append(n.childNodes[:oldIndex], n.childNodes[oldIndex+1:]...)
		oldChild.setParent(nil)

		// Insert all children from the fragment at the old position
		insertIndex := oldIndex
		for len(fragment.ChildNodes()) > 0 {
			fragmentChild := fragment.ChildNodes()[0]
			fragment.RemoveChild(fragmentChild)
			n.childNodes = append(n.childNodes[:insertIndex], append(NodeList{fragmentChild}, n.childNodes[insertIndex:]...)...)
			fragmentChild.setParent(n.self)
			fragmentChild.setOwnerDocument(n.ownerDocument)
			insertIndex++ // Increment position for next insertion
		}

		// Increment modification time
		if n.ownerDocument != nil {
			n.ownerDocument.incrementModificationTime()
		}

		return oldChild
	}

	// Special case: if newChild is already a child of this parent, we need to handle it carefully
	if newChild.ParentNode() == n.self {
		// Find current position of newChild
		var newChildIndex = -1
		for i, c := range n.childNodes {
			if c == newChild {
				newChildIndex = i
				break
			}
		}

		// Find position of oldChild
		var oldChildIndex = -1
		for i, c := range n.childNodes {
			if c == oldChild {
				oldChildIndex = i
				break
			}
		}

		if newChildIndex == -1 || oldChildIndex == -1 {
			return nil
		}

		// If they're the same child, do nothing
		if newChildIndex == oldChildIndex {
			return oldChild
		}

		// When replacing with an already-parented child within the same parent,
		// we swap their positions instead of removing one
		n.childNodes[oldChildIndex] = newChild
		n.childNodes[newChildIndex] = oldChild

		// The old child remains parented (it just moved positions)
		return oldChild
	}

	// Remove from old parent if it has one
	if newChild.ParentNode() != nil {
		newChild.ParentNode().RemoveChild(newChild)
	}

	for i, c := range n.childNodes {
		if c == oldChild {
			n.childNodes[i] = newChild
			newChild.setParent(n.self)
			newChild.setOwnerDocument(n.ownerDocument)
			oldChild.setParent(nil)

			// Increment modification time
			if n.ownerDocument != nil {
				n.ownerDocument.incrementModificationTime()
			}

			return oldChild
		}
	}
	return nil // oldChild not found
}

func (n *nodeImpl) CloneNode(deep bool) Node {
	clone := &nodeImpl{
		nodeType:      n.nodeType,
		nodeName:      n.nodeName,
		nodeValue:     n.nodeValue,
		ownerDocument: n.ownerDocument,
	}
	if deep {
		for _, child := range n.childNodes {
			clone.AppendChild(child.CloneNode(true))
		}
	}
	return clone
}

func (n *nodeImpl) setParent(parent Node) {
	n.parentNode = parent
}

func (n *nodeImpl) setOwnerDocument(doc *Document) {
	n.ownerDocument = doc
	for _, child := range n.childNodes {
		child.setOwnerDocument(doc)
	}
}

// Length implements the DOM specification requirement for determining the length of a node
func (n *nodeImpl) Length() int {
	switch n.nodeType {
	case DocumentTypeNode, AttributeNode:
		return 0
	case TextNode, CommentNode, ProcessingInstructionNode, CDataSectionNode:
		// For CharacterData nodes, return the length of the data
		return len([]rune(n.nodeValue)) // Use runes for proper Unicode support
	default:
		// For other nodes (Document, Element, DocumentFragment), return number of children
		return len(n.childNodes)
	}
}

// IsEmpty returns true if the node's length is 0
func (n *nodeImpl) IsEmpty() bool {
	return n.Length() == 0
}

// validateHierarchy implements the DOM hierarchy validation rules
func (n *nodeImpl) validateHierarchy(child Node) *DOMException {
	return n.validateHierarchyWithExclusion(child, nil)
}

// validateHierarchyWithExclusion validates hierarchy but excludes a specific node from uniqueness checks
// This is used during replacement operations where the old node is being replaced
func (n *nodeImpl) validateHierarchyWithExclusion(child Node, excludeNode Node) *DOMException {
	// Check for circular references (child cannot be an ancestor of n.self)
	current := n.self
	for current != nil {
		if current == child {
			return NewHierarchyRequestError("Cannot insert a node as a child of itself or its ancestors")
		}
		current = current.ParentNode()
	}

	// Check node type compatibility based on WHATWG DOM specification
	parentType := n.nodeType
	childType := child.NodeType()

	switch parentType {
	case DocumentNode:
		// Document can only have one Element, one DocumentType, and PIs/Comments
		switch childType {
		case ElementNode:
			// Check if Document already has an element child (excluding the node being replaced)
			for _, existing := range n.childNodes {
				if existing.NodeType() == ElementNode && existing != excludeNode {
					return NewHierarchyRequestError("Document can only have one Element child")
				}
			}
		case DocumentTypeNode:
			// Check if Document already has a doctype child (excluding the node being replaced)
			for _, existing := range n.childNodes {
				if existing.NodeType() == DocumentTypeNode && existing != excludeNode {
					return NewHierarchyRequestError("Document can only have one DocumentType child")
				}
			}
		case ProcessingInstructionNode, CommentNode:
			// These are allowed
		case DocumentFragmentNode:
			// Fragment contents will be validated individually
		default:
			return NewHierarchyRequestError("Document cannot contain this node type")
		}

	case DocumentTypeNode, TextNode, CDataSectionNode, ProcessingInstructionNode, CommentNode:
		// These node types cannot have children
		return NewHierarchyRequestError("This node type cannot have children")

	case ElementNode, DocumentFragmentNode:
		// Elements and DocumentFragments can contain Elements, Text, CData, PI, Comments
		switch childType {
		case ElementNode, TextNode, CDataSectionNode, ProcessingInstructionNode, CommentNode:
			// These are allowed
		case DocumentFragmentNode:
			// Fragment contents will be validated individually
		case DocumentNode, DocumentTypeNode:
			return NewHierarchyRequestError("Elements cannot contain Document or DocumentType nodes")
		default:
			return NewHierarchyRequestError("Elements cannot contain this node type")
		}

	case AttributeNode:
		// Attributes cannot have children
		return NewHierarchyRequestError("Attribute nodes cannot have children")

	default:
		return NewHierarchyRequestError("Unknown node type cannot have children")
	}

	return nil
}

// registerMutationObserver registers a mutation observer for this node
func (n *nodeImpl) registerMutationObserver(observer *MutationObserver, options MutationObserverInit) {
	// Get the document's observer registry
	if n.ownerDocument != nil {
		n.ownerDocument.getObserverRegistry().RegisterObserver(n.self, observer, options)
	}
}

// unregisterMutationObserver removes a mutation observer from this node
func (n *nodeImpl) unregisterMutationObserver(observer *MutationObserver) {
	// Get the document's observer registry
	if n.ownerDocument != nil {
		n.ownerDocument.getObserverRegistry().UnregisterObserver(n.self, observer)
	}
}

// toJS is a placeholder for JavaScript binding.
func (n *nodeImpl) toJS(vm *goja.Runtime) goja.Value {
	// This will be implemented later when integrating with Goja
	return vm.ToValue(nil)
}

// Traverse performs a depth-first traversal of the DOM tree starting from the given node.
// The visit function is called for each node. If visit returns false, traversal stops.
func Traverse(n Node, visit func(Node) bool) {
	if n == nil {
		return
	}

	if !visit(n) {
		return
	}

	for _, child := range n.ChildNodes() {
		Traverse(child, visit)
	}
}

// ChildNode represents the ChildNode mixin as defined in the WHATWG DOM specification.
// This interface provides convenience methods for manipulating child nodes.
type ChildNode interface {
	Before(nodes ...interface{}) error
	After(nodes ...interface{}) error
	ReplaceWith(nodes ...interface{}) error
	Remove() error
}

// convertNodesToNode implements the "converting nodes into a node" algorithm from the WHATWG DOM spec.
// This function takes a slice of Node or string values and converts them into a single Node.
func convertNodesToNode(nodes []interface{}, doc *Document) (Node, error) {
	if len(nodes) == 0 {
		return nil, nil
	}

	var convertedNodes []Node

	// Convert strings to Text nodes and validate Node types
	for _, node := range nodes {
		switch v := node.(type) {
		case Node:
			convertedNodes = append(convertedNodes, v)
		case string:
			textNode := NewText(v, doc)
			convertedNodes = append(convertedNodes, textNode)
		default:
			return nil, NewInvalidCharacterError("invalid node type: expected Node or string")
		}
	}

	// If only one node, return it directly
	if len(convertedNodes) == 1 {
		return convertedNodes[0], nil
	}

	// If multiple nodes, create a DocumentFragment to contain them
	fragment := NewDocumentFragment(doc)
	for _, node := range convertedNodes {
		fragment.AppendChild(node)
	}

	return fragment, nil
}

// findViablePreviousSibling finds the first preceding sibling that is not in the excludeNodes list.
// This implements the viable sibling logic from the WHATWG DOM specification.
func findViablePreviousSibling(node Node, excludeNodes []Node) Node {
	if node.ParentNode() == nil {
		return nil
	}

	// Create a map for faster lookup of excluded nodes
	excluded := make(map[Node]bool)
	for _, n := range excludeNodes {
		excluded[n] = true
	}

	// Find the node's position and search backwards for a viable sibling
	siblings := node.ParentNode().ChildNodes()
	for i, sibling := range siblings {
		if sibling == node {
			// Search backwards from this position
			for j := i - 1; j >= 0; j-- {
				if !excluded[siblings[j]] {
					return siblings[j]
				}
			}
			break
		}
	}

	return nil
}

// findViableNextSibling finds the first following sibling that is not in the excludeNodes list.
// This implements the viable sibling logic from the WHATWG DOM specification.
func findViableNextSibling(node Node, excludeNodes []Node) Node {
	if node.ParentNode() == nil {
		return nil
	}

	// Create a map for faster lookup of excluded nodes
	excluded := make(map[Node]bool)
	for _, n := range excludeNodes {
		excluded[n] = true
	}

	// Find the node's position and search forwards for a viable sibling
	siblings := node.ParentNode().ChildNodes()
	for i, sibling := range siblings {
		if sibling == node {
			// Search forwards from this position
			for j := i + 1; j < len(siblings); j++ {
				if !excluded[siblings[j]] {
					return siblings[j]
				}
			}
			break
		}
	}

	return nil
}

// preInsert is a helper function that implements the pre-insert algorithm.
// This is used by the ChildNode methods to safely insert nodes.
func preInsert(node, parent, child Node) error {
	if parent == nil {
		return NewNotFoundError("parent is null")
	}

	// Validate that the node can be inserted
	if !canInsertNode(node, parent) {
		return NewHierarchyRequestError("node cannot be inserted at the specified point in the hierarchy")
	}

	// Insert the node
	if child == nil {
		parent.AppendChild(node)
	} else {
		parent.InsertBefore(node, child)
	}

	return nil
}

// canInsertNode checks if a node can be legally inserted as a child of the given parent.
// This implements basic hierarchy validation rules.
func canInsertNode(node, parent Node) bool {
	if node == nil || parent == nil {
		return false
	}

	// Check for circular references (node cannot be an ancestor of parent)
	current := parent
	for current != nil {
		if current == node {
			return false
		}
		current = current.ParentNode()
	}

	// Additional type-based validation can be added here
	return true
}
