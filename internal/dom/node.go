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
		return nil
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
		return nil
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
	return nil // Child not found
}

func (n *nodeImpl) InsertBefore(newChild, refChild Node) Node {
	if newChild == nil {
		return nil
	}

	if refChild == nil {
		return n.AppendChild(newChild)
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
	if newChild == nil || oldChild == nil {
		return nil
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
