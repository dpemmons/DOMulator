package dom

import (
	"strings"

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

// Document position constants for compareDocumentPosition method
const (
	DOCUMENT_POSITION_DISCONNECTED            = 0x01
	DOCUMENT_POSITION_PRECEDING               = 0x02
	DOCUMENT_POSITION_FOLLOWING               = 0x04
	DOCUMENT_POSITION_CONTAINS                = 0x08
	DOCUMENT_POSITION_CONTAINED_BY            = 0x10
	DOCUMENT_POSITION_IMPLEMENTATION_SPECIFIC = 0x20
)

// GetRootNodeOptions represents the options dictionary for getRootNode method
type GetRootNodeOptions struct {
	Composed bool
}

// Node represents a single node in the DOM tree.
type Node interface {
	NodeType() NodeType
	NodeName() string
	NodeValue() string
	SetNodeValue(value string)
	OwnerDocument() *Document

	// Phase 1: Core Properties & Simple Methods
	IsConnected() bool
	ParentElement() *Element
	BaseURI() string
	HasChildNodes() bool

	// Phase 2: Text Content & Normalization
	TextContent() string
	SetTextContent(value string)
	Normalize()

	// Phase 3: Comparison & Traversal Methods
	GetRootNode(options *GetRootNodeOptions) Node
	IsEqualNode(otherNode Node) bool
	IsSameNode(otherNode Node) bool
	CompareDocumentPosition(other Node) int
	Contains(other Node) bool

	ParentNode() Node
	ChildNodes() *NodeList
	FirstChild() Node
	LastChild() Node
	PreviousSibling() Node
	NextSibling() Node

	AppendChild(child Node) Node
	RemoveChild(child Node) Node
	InsertBefore(newChild, refChild Node) Node
	ReplaceChild(newChild, oldChild Node) Node

	CloneNode(deep bool) Node

	// Namespace lookup methods
	LookupPrefix(namespace string) string
	LookupNamespaceURI(prefix string) string
	IsDefaultNamespace(namespace string) bool

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
	getRegisteredObservers() []*RegisteredObserver

	// Internal methods for DOM manipulation and JS binding
	setParent(parent Node)
	setOwnerDocument(doc *Document)
	toJS(vm *goja.Runtime) goja.Value
}

// nodeImpl provides a common base for all Node types.
type nodeImpl struct {
	nodeType            NodeType
	nodeName            string
	nodeValue           string
	parentNode          Node
	childNodes          []Node // Internal storage as slice
	ownerDocument       *Document
	self                Node // Reference to the containing Node
	eventListeners      map[string][]func(Event)
	registeredObservers []*RegisteredObserver // Per-spec registered observer list
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

func (n *nodeImpl) ChildNodes() *NodeList {
	return NewLiveNodeList(n.self, func() []Node {
		return n.childNodes
	})
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
	siblings := n.parentNode.ChildNodes()
	for i := 0; i < siblings.Length(); i++ {
		child := siblings.Item(i)
		if child == n.self {
			if i > 0 {
				return siblings.Item(i - 1)
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
	siblings := n.parentNode.ChildNodes()
	for i := 0; i < siblings.Length(); i++ {
		child := siblings.Item(i)
		if child == n.self {
			if i < siblings.Length()-1 {
				return siblings.Item(i + 1)
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
		// Collect all fragment children first (since removing them modifies the live NodeList)
		fragmentChildren := fragment.ChildNodes()
		var childrenToMove []Node
		for i := 0; i < fragmentChildren.Length(); i++ {
			childrenToMove = append(childrenToMove, fragmentChildren.Item(i))
		}

		// Move all children from the fragment to this node
		for _, fragmentChild := range childrenToMove {
			// Remove from fragment first
			fragment.RemoveChild(fragmentChild)

			// Add directly to this node without calling AppendChild (to avoid recursion)
			n.childNodes = append(n.childNodes, fragmentChild)
			fragmentChild.setParent(n.self)
			fragmentChild.setOwnerDocument(n.ownerDocument)
		}

		// Increment modification time
		if n.ownerDocument != nil {
			n.ownerDocument.incrementModificationTime()
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

		// Collect all fragment children first (since removing them modifies the live NodeList)
		fragmentChildren := fragment.ChildNodes()
		var childrenToInsert []Node
		for i := 0; i < fragmentChildren.Length(); i++ {
			childrenToInsert = append(childrenToInsert, fragmentChildren.Item(i))
		}

		// Insert all children from the fragment before refChild
		for _, fragmentChild := range childrenToInsert {
			fragment.RemoveChild(fragmentChild)
			n.childNodes = append(n.childNodes[:refIndex], append([]Node{fragmentChild}, n.childNodes[refIndex:]...)...)
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

		// If already at the correct position, do nothing
		if currentIndex == targetIndex || currentIndex == targetIndex-1 {
			return newChild
		}

		// Remove from current position first
		n.childNodes = append(n.childNodes[:currentIndex], n.childNodes[currentIndex+1:]...)

		// Adjust target index if necessary (when moving from earlier position to later position)
		adjustedTargetIndex := targetIndex
		if currentIndex < targetIndex {
			adjustedTargetIndex = targetIndex - 1
		}

		// Insert at the adjusted target index
		n.childNodes = append(n.childNodes[:adjustedTargetIndex], append([]Node{newChild}, n.childNodes[adjustedTargetIndex:]...)...)
		return newChild
	}

	// Remove from old parent if it has one
	if newChild.ParentNode() != nil {
		newChild.ParentNode().RemoveChild(newChild)
	}

	for i, c := range n.childNodes {
		if c == refChild {
			n.childNodes = append(n.childNodes[:i], append([]Node{newChild}, n.childNodes[i:]...)...)
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
		fragmentChildren := fragment.ChildNodes()
		for fragmentChildren.Length() > 0 {
			fragmentChild := fragmentChildren.Item(0)
			fragment.RemoveChild(fragmentChild)
			n.childNodes = append(n.childNodes[:insertIndex], append([]Node{fragmentChild}, n.childNodes[insertIndex:]...)...)
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
	// Use the spec-compliant cloning implementation
	return CloneNodeSpec(n.self, deep)
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

// Phase 1: Core Properties & Simple Methods

// IsConnected returns true if the node is connected to a document
// Per WHATWG DOM Section 4.2.1: A node is in a document tree if its shadow-including root is a document
func (n *nodeImpl) IsConnected() bool {
	// A node is connected if its shadow-including root is a document
	// Use GetRootNode with composed: true to cross shadow boundaries
	shadowIncludingRoot := n.self.GetRootNode(&GetRootNodeOptions{Composed: true})
	_, isDocument := shadowIncludingRoot.(*Document)
	return isDocument
}

// ParentElement returns the parent if it's an Element, null otherwise
func (n *nodeImpl) ParentElement() *Element {
	if n.parentNode == nil {
		return nil
	}
	if elem, ok := n.parentNode.(*Element); ok {
		return elem
	}
	return nil
}

// BaseURI returns the node's document base URL, serialized
func (n *nodeImpl) BaseURI() string {
	// TODO: Implement proper document URL handling
	// For now, return a default base URI
	if n.ownerDocument != nil {
		return "about:blank"
	}
	return ""
}

// HasChildNodes returns true if the node has children
func (n *nodeImpl) HasChildNodes() bool {
	return len(n.childNodes) > 0
}

// Phase 2: Text Content & Normalization

// SetNodeValue implements the nodeValue setter from the DOM specification
func (n *nodeImpl) SetNodeValue(value string) {
	// Handle null as empty string per spec
	if value == "" {
		value = ""
	}

	// Switch based on node type per specification
	switch n.nodeType {
	case AttributeNode:
		// For Attr nodes: delegate to SetValue if it's an Attr
		if attr, ok := n.self.(*Attr); ok {
			attr.SetValue(value)
		}
	case TextNode, CommentNode, ProcessingInstructionNode, CDataSectionNode:
		// For CharacterData nodes: delegate to SetData if it implements CharacterData
		if charData, ok := n.self.(CharacterData); ok {
			charData.SetData(value)
		} else {
			// Fallback: set nodeValue directly
			n.nodeValue = value
		}
	default:
		// For other nodes: do nothing per specification
	}
}

// Phase 3: Comparison & Traversal Methods

// GetRootNode implements the getRootNode(options) method from the DOM specification
func (n *nodeImpl) GetRootNode(options *GetRootNodeOptions) Node {
	// If options is nil, create default options
	if options == nil {
		options = &GetRootNodeOptions{Composed: false}
	}

	// For now, we don't support shadow DOM, so composed doesn't matter
	// Just find the root node by traversing up the parent chain
	root := n.self
	for root.ParentNode() != nil {
		root = root.ParentNode()
	}
	return root
}

// IsEqualNode implements the isEqualNode(otherNode) method from the DOM specification
func (n *nodeImpl) IsEqualNode(otherNode Node) bool {
	if otherNode == nil {
		return false
	}

	// A and B implement the same interfaces (check node type)
	if n.nodeType != otherNode.NodeType() {
		return false
	}

	// Check type-specific properties per specification
	switch n.nodeType {
	case DocumentTypeNode:
		// Compare name, public ID, and system ID
		if selfDT, ok := n.self.(*DocumentType); ok {
			if otherDT, ok := otherNode.(*DocumentType); ok {
				return selfDT.Name() == otherDT.Name() &&
					selfDT.PublicID() == otherDT.PublicID() &&
					selfDT.SystemID() == otherDT.SystemID()
			}
		}
		return false

	case ElementNode:
		// Compare namespace, namespace prefix, local name, and attribute list size
		if selfElem, ok := n.self.(*Element); ok {
			if otherElem, ok := otherNode.(*Element); ok {
				// Check basic properties
				if selfElem.TagName() != otherElem.TagName() ||
					selfElem.NamespaceURI() != otherElem.NamespaceURI() ||
					selfElem.Prefix() != otherElem.Prefix() {
					return false
				}

				// Check attribute list size and content
				if selfElem.attributes.Length() != otherElem.attributes.Length() {
					return false
				}

				// Each attribute in A has an equal attribute in B
				selfAttrs := selfElem.attributes.ToSlice()
				for _, attr := range selfAttrs {
					otherAttr := otherElem.attributes.GetNamedItem(attr.Name())
					if otherAttr == nil || otherAttr.Value() != attr.Value() {
						return false
					}
				}
			} else {
				return false
			}
		} else {
			return false
		}

	case AttributeNode:
		// Compare namespace, local name, and value
		if selfAttr, ok := n.self.(*Attr); ok {
			if otherAttr, ok := otherNode.(*Attr); ok {
				// For now, use Name() for both namespace and local name comparison
				// TODO: Add proper namespace support to Attr struct
				return selfAttr.Name() == otherAttr.Name() &&
					selfAttr.Value() == otherAttr.Value()
			}
		}
		return false

	case ProcessingInstructionNode:
		// Compare target and data
		if selfPI, ok := n.self.(*ProcessingInstruction); ok {
			if otherPI, ok := otherNode.(*ProcessingInstruction); ok {
				return selfPI.Target() == otherPI.Target() &&
					selfPI.Data() == otherPI.Data()
			}
		}
		return false

	case TextNode, CommentNode:
		// Compare data
		if selfCharData, ok := n.self.(CharacterData); ok {
			if otherCharData, ok := otherNode.(CharacterData); ok {
				return selfCharData.Data() == otherCharData.Data()
			}
		}
		return false

	default:
		// For other node types, no additional properties to compare
	}

	// Check that A and B have the same number of children
	selfChildren := n.self.ChildNodes()
	otherChildren := otherNode.ChildNodes()
	if selfChildren.Length() != otherChildren.Length() {
		return false
	}

	// Each child of A equals the child of B at the identical index
	for i := 0; i < selfChildren.Length(); i++ {
		if !selfChildren.Item(i).IsEqualNode(otherChildren.Item(i)) {
			return false
		}
	}

	return true
}

// IsSameNode implements the isSameNode(otherNode) method from the DOM specification
func (n *nodeImpl) IsSameNode(otherNode Node) bool {
	// isSameNode is a legacy alias for === comparison
	return n.self == otherNode
}

// CompareDocumentPosition implements the compareDocumentPosition(other) method from the DOM specification
func (n *nodeImpl) CompareDocumentPosition(other Node) int {
	if other == nil {
		return DOCUMENT_POSITION_DISCONNECTED | DOCUMENT_POSITION_IMPLEMENTATION_SPECIFIC | DOCUMENT_POSITION_PRECEDING
	}

	// If this is other, return zero
	if n.self == other {
		return 0
	}

	node1 := other
	node2 := n.self

	// Handle attributes (not fully implemented as we don't have attribute ownership tracking)
	// For now, we'll skip the attribute handling part of the algorithm

	// Check if nodes are in different trees
	node1Root := node1
	for node1Root.ParentNode() != nil {
		node1Root = node1Root.ParentNode()
	}

	node2Root := node2
	for node2Root.ParentNode() != nil {
		node2Root = node2Root.ParentNode()
	}

	if node1Root != node2Root {
		// Nodes are disconnected - return consistent implementation-specific result
		// Use a simple comparison to ensure consistency
		if node1Root.NodeName() < node2Root.NodeName() {
			return DOCUMENT_POSITION_DISCONNECTED | DOCUMENT_POSITION_IMPLEMENTATION_SPECIFIC | DOCUMENT_POSITION_FOLLOWING
		}
		return DOCUMENT_POSITION_DISCONNECTED | DOCUMENT_POSITION_IMPLEMENTATION_SPECIFIC | DOCUMENT_POSITION_PRECEDING
	}

	// Check if node1 is an ancestor of node2
	current := node2.ParentNode()
	for current != nil {
		if current == node1 {
			return DOCUMENT_POSITION_CONTAINS | DOCUMENT_POSITION_PRECEDING
		}
		current = current.ParentNode()
	}

	// Check if node1 is a descendant of node2
	current = node1.ParentNode()
	for current != nil {
		if current == node2 {
			return DOCUMENT_POSITION_CONTAINED_BY | DOCUMENT_POSITION_FOLLOWING
		}
		current = current.ParentNode()
	}

	// Nodes are siblings or in different branches - find which comes first in tree order
	if n.isNodePreceding(node1, node2) {
		return DOCUMENT_POSITION_PRECEDING
	}

	return DOCUMENT_POSITION_FOLLOWING
}

// isNodePreceding determines if node1 comes before node2 in tree order
func (n *nodeImpl) isNodePreceding(node1, node2 Node) bool {
	// Find common ancestor
	ancestors1 := n.getAncestors(node1)
	ancestors2 := n.getAncestors(node2)

	// Find the deepest common ancestor
	var commonAncestor Node
	var child1, child2 Node

	// Compare ancestors from root to leaf to find the deepest common ancestor
	minLen := len(ancestors1)
	if len(ancestors2) < minLen {
		minLen = len(ancestors2)
	}

	// Find common ancestor by comparing from the root down
	for i := 1; i <= minLen; i++ {
		idx1 := len(ancestors1) - i
		idx2 := len(ancestors2) - i

		if ancestors1[idx1] == ancestors2[idx2] {
			commonAncestor = ancestors1[idx1]
			// The children are the next level down (closer to the original nodes)
			if idx1 > 0 {
				child1 = ancestors1[idx1-1]
			} else {
				child1 = node1
			}
			if idx2 > 0 {
				child2 = ancestors2[idx2-1]
			} else {
				child2 = node2
			}
		} else {
			break // Found the deepest common ancestor
		}
	}

	if commonAncestor == nil {
		return false // Shouldn't happen if nodes are in same tree
	}

	// Find which child comes first among the common ancestor's children
	children := commonAncestor.ChildNodes()
	for i := 0; i < children.Length(); i++ {
		child := children.Item(i)
		if child == child1 {
			return true
		}
		if child == child2 {
			return false
		}
	}

	return false
}

// getAncestors returns a slice of ancestors from the node up to the root
func (n *nodeImpl) getAncestors(node Node) []Node {
	var ancestors []Node
	current := node
	for current != nil {
		ancestors = append(ancestors, current)
		current = current.ParentNode()
	}
	return ancestors
}

// Contains implements the contains(other) method from the DOM specification
func (n *nodeImpl) Contains(other Node) bool {
	if other == nil {
		return false
	}

	// A node contains itself
	if n.self == other {
		return true
	}

	// Check if other is a descendant of this node
	current := other.ParentNode()
	for current != nil {
		if current == n.self {
			return true
		}
		current = current.ParentNode()
	}

	return false
}

// getTextContent implements the "get text content" algorithm from the DOM specification
func getTextContent(node Node) string {
	if node == nil {
		return ""
	}

	switch node.NodeType() {
	case DocumentFragmentNode, ElementNode:
		// Return descendant text content
		var result strings.Builder
		children := node.ChildNodes()
		for i := 0; i < children.Length(); i++ {
			child := children.Item(i)
			childText := getTextContent(child)
			result.WriteString(childText)
		}
		return result.String()
	case AttributeNode:
		// Return the attribute's value
		if attr, ok := node.(*Attr); ok {
			return attr.Value()
		}
		return ""
	case TextNode, CommentNode, ProcessingInstructionNode, CDataSectionNode:
		// Return the node's data
		if charData, ok := node.(CharacterData); ok {
			return charData.Data()
		}
		return node.NodeValue()
	default:
		// Return null (empty string) for other node types
		return ""
	}
}

// TextContent implements the textContent getter from the DOM specification
func (n *nodeImpl) TextContent() string {
	return getTextContent(n.self)
}

// stringReplaceAll implements the "string replace all" algorithm from the DOM specification
func stringReplaceAll(str string, parent Node) {
	// Per spec: "To string replace all with a string string within a node parent, run these steps:"

	// Remove all existing children first
	children := parent.ChildNodes()
	for children.Length() > 0 {
		parent.RemoveChild(children.Item(0))
	}

	// 1. Let node be null.
	// 2. If string is not the empty string, then set node to a new Text node whose data is string and node document is parent's node document.
	// 3. Replace all with node within parent.

	// CRITICAL: Only create and append a text node if the string is NOT empty
	if str != "" && parent.OwnerDocument() != nil {
		node := NewText(str, parent.OwnerDocument())
		parent.AppendChild(node)
	}
	// If str is empty, we leave parent with no children (this is correct per spec)
}

// setTextContent implements the "set text content" algorithm from the DOM specification
func setTextContent(node Node, value string) {
	if node == nil {
		return
	}

	// Handle null as empty string per spec
	if value == "" {
		value = ""
	}

	switch node.NodeType() {
	case DocumentFragmentNode, ElementNode:
		// String replace all with value within node
		stringReplaceAll(value, node)
	case AttributeNode:
		// Set existing attribute value
		if attr, ok := node.(*Attr); ok {
			attr.SetValue(value)
		}
	case TextNode, CommentNode, ProcessingInstructionNode, CDataSectionNode:
		// Replace data with node, offset 0, count node's length, and data value
		if charData, ok := node.(CharacterData); ok {
			charData.SetData(value)
		}
	default:
		// Do nothing for other node types
	}
}

// SetTextContent implements the textContent setter from the DOM specification
func (n *nodeImpl) SetTextContent(value string) {
	setTextContent(n.self, value)
}

// isExclusiveTextNode checks if a node is an exclusive Text node (not CDATA)
func isExclusiveTextNode(node Node) bool {
	return node.NodeType() == TextNode
}

// Normalize implements the normalize() method from the DOM specification
func (n *nodeImpl) Normalize() {
	// Run these steps for each descendant exclusive Text node of this
	var textNodesToProcess []Node

	// Collect all descendant exclusive Text nodes
	Traverse(n.self, func(node Node) bool {
		if isExclusiveTextNode(node) {
			textNodesToProcess = append(textNodesToProcess, node)
		}
		return true // Continue traversal
	})

	// Process each Text node
	for _, textNode := range textNodesToProcess {
		if !isExclusiveTextNode(textNode) {
			continue // Skip if it's no longer a Text node
		}

		charData, ok := textNode.(CharacterData)
		if !ok {
			continue
		}

		length := charData.Length()

		// If length is zero, remove node and continue
		if length == 0 {
			if textNode.ParentNode() != nil {
				textNode.ParentNode().RemoveChild(textNode)
			}
			continue
		}

		// Find contiguous exclusive Text nodes
		var contiguousNodes []Node
		currentNode := textNode.NextSibling()

		for currentNode != nil && isExclusiveTextNode(currentNode) {
			contiguousNodes = append(contiguousNodes, currentNode)
			currentNode = currentNode.NextSibling()
		}

		if len(contiguousNodes) == 0 {
			continue // No contiguous text nodes to concatenate
		}

		// Get the data of contiguous exclusive Text nodes in tree order
		var dataBuilder strings.Builder
		for _, node := range contiguousNodes {
			if cd, ok := node.(CharacterData); ok {
				dataBuilder.WriteString(cd.Data())
			}
		}
		concatenatedData := dataBuilder.String()

		// Replace data with node textNode, offset length, count 0, and data concatenatedData
		if concatenatedData != "" {
			charData.ReplaceData(length, 0, concatenatedData)
		}

		// TODO: Handle live ranges here when implemented
		// The spec requires updating range start/end offsets for each contiguous node

		// Remove contiguous exclusive Text nodes in tree order
		for _, node := range contiguousNodes {
			if node.ParentNode() != nil {
				node.ParentNode().RemoveChild(node)
			}
		}
	}
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

// getRegisteredObservers returns the list of registered observers for this node
func (n *nodeImpl) getRegisteredObservers() []*RegisteredObserver {
	if n.registeredObservers == nil {
		n.registeredObservers = make([]*RegisteredObserver, 0)
	}
	return n.registeredObservers
}

// addRegisteredObserver adds a registered observer to this node
func (n *nodeImpl) addRegisteredObserver(ro *RegisteredObserver) {
	if n.registeredObservers == nil {
		n.registeredObservers = make([]*RegisteredObserver, 0)
	}
	n.registeredObservers = append(n.registeredObservers, ro)
}

// removeRegisteredObserver removes a registered observer from this node
func (n *nodeImpl) removeRegisteredObserver(observer *MutationObserver) {
	if n.registeredObservers == nil {
		return
	}

	for i := len(n.registeredObservers) - 1; i >= 0; i-- {
		if n.registeredObservers[i].Observer == observer {
			n.registeredObservers = append(n.registeredObservers[:i], n.registeredObservers[i+1:]...)
		}
	}
}

// locateNamespacePrefix implements the "locate a namespace prefix" algorithm from the DOM specification
func (n *nodeImpl) locateNamespacePrefix(element *Element, namespace string) string {
	// If element's namespace is namespace and its namespace prefix is non-null, return its namespace prefix
	if element.NamespaceURI() == namespace && element.Prefix() != "" {
		return element.Prefix()
	}

	// If element has an attribute whose namespace prefix is "xmlns" and value is namespace,
	// then return element's first such attribute's local name
	attrs := element.attributes.ToSlice()
	for _, attr := range attrs {
		attrName := attr.Name()
		if strings.HasPrefix(attrName, "xmlns:") && attr.Value() == namespace {
			return strings.TrimPrefix(attrName, "xmlns:")
		}
	}

	// If element's parent element is not null, return the result of running
	// locate a namespace prefix on that element using namespace
	if parentElement := element.ParentElement(); parentElement != nil {
		return n.locateNamespacePrefix(parentElement, namespace)
	}

	// Return null
	return ""
}

// locateNamespace implements the "locate a namespace" algorithm from the DOM specification
func (n *nodeImpl) locateNamespace(node Node, prefix string) string {
	if node == nil {
		return ""
	}

	switch node.NodeType() {
	case ElementNode:
		element := node.(*Element)

		// If prefix is "xml", return the XML namespace
		if prefix == "xml" {
			return "http://www.w3.org/XML/1998/namespace"
		}

		// If prefix is "xmlns", return the XMLNS namespace
		if prefix == "xmlns" {
			return "http://www.w3.org/2000/xmlns/"
		}

		// If its namespace is non-null and its namespace prefix is prefix, return namespace
		if element.NamespaceURI() != "" && element.Prefix() == prefix {
			return element.NamespaceURI()
		}

		// Check for xmlns attributes
		if prefix != "" {
			// Look for xmlns:prefix attribute
			attr := element.attributes.GetNamedItem("xmlns:" + prefix)
			if attr != nil {
				value := attr.Value()
				if value == "" {
					return ""
				}
				return value
			}
		} else {
			// Look for xmlns attribute (default namespace)
			attr := element.attributes.GetNamedItem("xmlns")
			if attr != nil {
				value := attr.Value()
				if value == "" {
					return ""
				}
				return value
			}
		}

		// If its parent element is null, return null
		if parentElement := element.ParentElement(); parentElement == nil {
			return ""
		} else {
			// Return the result of running locate a namespace on its parent element using prefix
			return n.locateNamespace(parentElement, prefix)
		}

	case DocumentNode:
		document := node.(*Document)

		// If its document element is null, return null
		if document.DocumentElement() == nil {
			return ""
		}

		// Return the result of running locate a namespace on its document element using prefix
		return n.locateNamespace(document.DocumentElement(), prefix)

	case DocumentTypeNode, DocumentFragmentNode:
		// Return null
		return ""

	case AttributeNode:
		attr := node.(*Attr)

		// If its element is null, return null
		if attr.OwnerElement() == nil {
			return ""
		}

		// Return the result of running locate a namespace on its element using prefix
		return n.locateNamespace(attr.OwnerElement(), prefix)

	default:
		// For other node types (Text, Comment, ProcessingInstruction, etc.)
		// If its parent element is null, return null
		if parentElement := node.ParentElement(); parentElement == nil {
			return ""
		} else {
			// Return the result of running locate a namespace on its parent element using prefix
			return n.locateNamespace(parentElement, prefix)
		}
	}
}

// LookupPrefix implements the lookupPrefix method from the DOM specification
func (n *nodeImpl) LookupPrefix(namespace string) string {
	// If namespace is null or the empty string, return null
	if namespace == "" {
		return ""
	}

	// Switch on the interface this implements
	switch n.nodeType {
	case ElementNode:
		element := n.self.(*Element)
		// Return the result of locating a namespace prefix for this using namespace
		return n.locateNamespacePrefix(element, namespace)

	case DocumentNode:
		document := n.self.(*Document)
		// If this's document element is null, return null
		if document.DocumentElement() == nil {
			return ""
		}
		// Return the result of locating a namespace prefix for this's document element using namespace
		return n.locateNamespacePrefix(document.DocumentElement(), namespace)

	case DocumentTypeNode, DocumentFragmentNode:
		// Return null
		return ""

	case AttributeNode:
		attr := n.self.(*Attr)
		// If this's element is null, return null
		if attr.OwnerElement() == nil {
			return ""
		}
		// Return the result of locating a namespace prefix for this's element using namespace
		return n.locateNamespacePrefix(attr.OwnerElement(), namespace)

	default:
		// For other node types (Text, Comment, ProcessingInstruction, etc.)
		// If this's parent element is null, return null
		if parentElement := n.self.ParentElement(); parentElement == nil {
			return ""
		} else {
			// Return the result of locating a namespace prefix for this's parent element using namespace
			return n.locateNamespacePrefix(parentElement, namespace)
		}
	}
}

// LookupNamespaceURI implements the lookupNamespaceURI method from the DOM specification
func (n *nodeImpl) LookupNamespaceURI(prefix string) string {
	// If prefix is the empty string, set it to null
	if prefix == "" {
		prefix = ""
	}

	// Return the result of running locate a namespace for this using prefix
	return n.locateNamespace(n.self, prefix)
}

// IsDefaultNamespace implements the isDefaultNamespace method from the DOM specification
func (n *nodeImpl) IsDefaultNamespace(namespace string) bool {
	// If namespace is the empty string, set it to null
	if namespace == "" {
		namespace = ""
	}

	// Let defaultNamespace be the result of running locate a namespace for this using null
	defaultNamespace := n.locateNamespace(n.self, "")

	// Return true if defaultNamespace is the same as namespace; otherwise false
	return defaultNamespace == namespace
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

	children := n.ChildNodes()
	for i := 0; i < children.Length(); i++ {
		child := children.Item(i)
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
	for i := 0; i < siblings.Length(); i++ {
		sibling := siblings.Item(i)
		if sibling == node {
			// Search backwards from this position
			for j := i - 1; j >= 0; j-- {
				prevSibling := siblings.Item(j)
				if !excluded[prevSibling] {
					return prevSibling
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
	for i := 0; i < siblings.Length(); i++ {
		sibling := siblings.Item(i)
		if sibling == node {
			// Search forwards from this position
			for j := i + 1; j < siblings.Length(); j++ {
				nextSibling := siblings.Item(j)
				if !excluded[nextSibling] {
					return nextSibling
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

// matchSimpleSelector checks if an element matches a basic CSS selector
// This is a shared helper function for simple selector matching across DOM nodes
func matchSimpleSelector(elem *Element, selector string) bool {
	selector = strings.TrimSpace(selector)
	if selector == "" {
		return false
	}

	// Handle wildcard selector
	if selector == "*" {
		return true
	}

	// Parse simple selectors: tag, #id, .class, tag#id, tag.class, etc.
	remaining := selector
	var tagName, id string
	var classes []string

	// Extract ID (e.g., #myid)
	if idx := strings.Index(remaining, "#"); idx != -1 {
		idPart := remaining[idx+1:]
		if spaceIdx := strings.IndexAny(idPart, ". "); spaceIdx != -1 {
			id = idPart[:spaceIdx]
			remaining = remaining[:idx] + idPart[spaceIdx:]
		} else {
			id = idPart
			remaining = remaining[:idx]
		}
	}

	// Extract Classes (e.g., .myclass)
	classParts := strings.Split(remaining, ".")
	if len(classParts) > 1 {
		for _, class := range classParts[1:] {
			if class != "" {
				classes = append(classes, class)
			}
		}
		remaining = classParts[0]
	}

	// Extract Tag (the remaining part)
	tagName = strings.TrimSpace(remaining)

	// Check tag match
	if tagName != "" && !strings.EqualFold(tagName, elem.TagName()) {
		return false
	}

	// Check ID match
	if id != "" && elem.GetAttribute("id") != id {
		return false
	}

	// Check class matches
	if len(classes) > 0 {
		elemClasses := strings.Fields(elem.GetAttribute("class"))
		for _, requiredClass := range classes {
			found := false
			for _, elemClass := range elemClasses {
				if requiredClass == elemClass {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}
	}

	return true
}
