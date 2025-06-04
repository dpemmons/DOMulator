package dom

// NodeList represents a collection of nodes as defined in the WHATWG DOM specification.
//
// The NodeList interface provides access to an ordered collection of nodes.
// NodeList objects are collections of nodes, usually returned by properties such as
// Node.childNodes and methods such as document.querySelectorAll().
type NodeList struct {
	nodes []Node
	// For live NodeLists, we store a reference to the parent and a function to get current nodes
	parent   Node
	getNodes func() []Node
	isLive   bool
}

// NewNodeList creates a new NodeList from a slice of nodes
func NewNodeList(nodes []Node) *NodeList {
	// Create a copy to ensure the NodeList owns its data
	nodesCopy := make([]Node, len(nodes))
	copy(nodesCopy, nodes)
	return &NodeList{nodes: nodesCopy, isLive: false}
}

// NewLiveNodeList creates a new NodeList that dynamically reflects changes to the parent's children.
// This is used for live collections where the NodeList should always reflect the current state.
func NewLiveNodeList(parent Node, getNodes func() []Node) *NodeList {
	return &NodeList{
		parent:   parent,
		getNodes: getNodes,
		isLive:   true,
	}
}

// getCurrentNodes returns the current nodes, updating live collections if necessary
func (nl *NodeList) getCurrentNodes() []Node {
	if nl == nil {
		return nil
	}
	if nl.isLive && nl.getNodes != nil {
		return nl.getNodes()
	}
	return nl.nodes
}

// Length returns the number of nodes in the collection.
//
// From the spec: "The length attribute must return the number of nodes represented by the collection."
func (nl *NodeList) Length() int {
	nodes := nl.getCurrentNodes()
	return len(nodes)
}

// Item returns the node with index index from the collection.
// The nodes are sorted in tree order.
//
// From the spec: "The item(index) method must return the indexth node in the collection.
// If there is no indexth node in the collection, then the method must return null."
func (nl *NodeList) Item(index int) Node {
	nodes := nl.getCurrentNodes()
	if index < 0 || index >= len(nodes) {
		return nil
	}
	return nodes[index]
}

// ToSlice returns a copy of the underlying slice of nodes.
// This is a convenience method not in the spec but useful for Go code.
func (nl *NodeList) ToSlice() []Node {
	nodes := nl.getCurrentNodes()
	if nodes == nil {
		return nil
	}
	result := make([]Node, len(nodes))
	copy(result, nodes)
	return result
}

// forEach provides iteration over the NodeList.
// This is a convenience method for Go code.
func (nl *NodeList) ForEach(fn func(node Node, index int)) {
	nodes := nl.getCurrentNodes()
	for i, node := range nodes {
		fn(node, i)
	}
}

// Contains checks if a node is present in the NodeList.
// This is a convenience method for Go code.
func (nl *NodeList) Contains(node Node) bool {
	nodes := nl.getCurrentNodes()
	for _, n := range nodes {
		if n == node {
			return true
		}
	}
	return false
}

// IndexOf returns the index of the node in the NodeList, or -1 if not found.
// This is a convenience method for Go code.
func (nl *NodeList) IndexOf(node Node) int {
	nodes := nl.getCurrentNodes()
	for i, n := range nodes {
		if n == node {
			return i
		}
	}
	return -1
}

// IsEmpty returns true if the NodeList contains no nodes.
// This is a convenience method for Go code.
func (nl *NodeList) IsEmpty() bool {
	return nl.Length() == 0
}
