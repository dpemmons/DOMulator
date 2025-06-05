package ranges

import (
	"testing"
)

// Mock types for testing
type mockDocument struct {
	nodeType      int
	nodeName      string
	children      []Node
	parent        Node
	ownerDocument Document
}

func (m *mockDocument) NodeType() int           { return 9 } // Document node
func (m *mockDocument) ParentNode() Node        { return nil }
func (m *mockDocument) ChildNodes() NodeList    { return &mockNodeList{nodes: m.children} }
func (m *mockDocument) OwnerDocument() Document { return m }
func (m *mockDocument) NextSibling() Node       { return nil }
func (m *mockDocument) PreviousSibling() Node   { return nil }
func (m *mockDocument) Length() int             { return len(m.children) }

type mockElement struct {
	nodeType      int
	nodeName      string
	children      []Node
	parent        Node
	ownerDocument Document
}

func (m *mockElement) NodeType() int           { return m.nodeType }
func (m *mockElement) ParentNode() Node        { return m.parent }
func (m *mockElement) ChildNodes() NodeList    { return &mockNodeList{nodes: m.children} }
func (m *mockElement) OwnerDocument() Document { return m.ownerDocument }
func (m *mockElement) NextSibling() Node       { return nil } // Simplified for testing
func (m *mockElement) PreviousSibling() Node   { return nil } // Simplified for testing
func (m *mockElement) Length() int             { return len(m.children) }

type mockText struct {
	nodeType      int
	nodeName      string
	data          string
	parent        Node
	ownerDocument Document
}

func (m *mockText) NodeType() int           { return m.nodeType }
func (m *mockText) ParentNode() Node        { return m.parent }
func (m *mockText) ChildNodes() NodeList    { return &mockNodeList{nodes: []Node{}} }
func (m *mockText) OwnerDocument() Document { return m.ownerDocument }
func (m *mockText) NextSibling() Node       { return nil }
func (m *mockText) PreviousSibling() Node   { return nil }
func (m *mockText) Length() int             { return len(m.data) }

// Implement CharacterData interface
func (m *mockText) Data() string        { return m.data }
func (m *mockText) SetData(data string) { m.data = data }

type mockNodeList struct {
	nodes []Node
}

func (m *mockNodeList) Length() int { return len(m.nodes) }
func (m *mockNodeList) Item(index int) Node {
	if index >= 0 && index < len(m.nodes) {
		return m.nodes[index]
	}
	return nil
}

// boundaryPointPosition returns the position of boundary point A relative to boundary point B
func boundaryPointPosition(nodeA Node, offsetA int, nodeB Node, offsetB int) string {
	pointA := &BoundaryPoint{Node: nodeA, Offset: offsetA}
	pointB := &BoundaryPoint{Node: nodeB, Offset: offsetB}

	result := CompareBoundaryPoints(pointA, pointB)
	switch result {
	case -1:
		return "before"
	case 0:
		return "equal"
	case 1:
		return "after"
	default:
		return "before"
	}
}

// TestBoundaryPoint tests boundary point operations per WHATWG DOM specification
func TestBoundaryPoint(t *testing.T) {
	// Create a mock document for testing
	doc := &mockDocument{}
	root := &mockElement{nodeType: 1, nodeName: "root", ownerDocument: doc}
	child1 := &mockElement{nodeType: 1, nodeName: "child1", ownerDocument: doc, parent: root}
	child2 := &mockElement{nodeType: 1, nodeName: "child2", ownerDocument: doc, parent: root}
	text1 := &mockText{nodeType: 3, nodeName: "#text", data: "Hello", ownerDocument: doc, parent: child1}

	// Set up parent-child relationships
	root.children = []Node{child1, child2}
	child1.children = []Node{text1}

	tests := []struct {
		name     string
		nodeA    Node
		offsetA  int
		nodeB    Node
		offsetB  int
		expected string
	}{
		{
			name:     "same node, same offset",
			nodeA:    root,
			offsetA:  0,
			nodeB:    root,
			offsetB:  0,
			expected: "equal",
		},
		{
			name:     "same node, offsetA < offsetB",
			nodeA:    root,
			offsetA:  0,
			nodeB:    root,
			offsetB:  1,
			expected: "before",
		},
		{
			name:     "same node, offsetA > offsetB",
			nodeA:    root,
			offsetA:  1,
			nodeB:    root,
			offsetB:  0,
			expected: "after",
		},
		{
			name:     "ancestor relationship",
			nodeA:    root,
			offsetA:  0,
			nodeB:    child1,
			offsetB:  0,
			expected: "before",
		},
		{
			name:     "ancestor relationship 2",
			nodeA:    root,
			offsetA:  1,
			nodeB:    child1,
			offsetB:  0,
			expected: "after",
		},
		{
			name:     "descendant relationship",
			nodeA:    child1,
			offsetA:  0,
			nodeB:    root,
			offsetB:  1,
			expected: "before",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := boundaryPointPosition(tt.nodeA, tt.offsetA, tt.nodeB, tt.offsetB)
			if result != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, result)
			}
		})
	}
}

// TestBoundaryPointValidation tests boundary point validation
func TestBoundaryPointValidation(t *testing.T) {
	doc := &mockDocument{}
	elem := &mockElement{nodeType: 1, nodeName: "div", ownerDocument: doc}
	text := &mockText{nodeType: 3, nodeName: "#text", data: "Hello", ownerDocument: doc}

	tests := []struct {
		name   string
		node   Node
		offset int
		valid  bool
	}{
		{
			name:   "valid element offset 0",
			node:   elem,
			offset: 0,
			valid:  true,
		},
		{
			name:   "valid text offset within bounds",
			node:   text,
			offset: 3,
			valid:  true,
		},
		{
			name:   "valid text offset at end",
			node:   text,
			offset: 5,
			valid:  true,
		},
		{
			name:   "invalid negative offset",
			node:   elem,
			offset: -1,
			valid:  false,
		},
		{
			name:   "invalid text offset beyond length",
			node:   text,
			offset: 6,
			valid:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid := isValidBoundaryPoint(tt.node, tt.offset)
			if valid != tt.valid {
				t.Errorf("Expected valid=%v, got valid=%v", tt.valid, valid)
			}
		})
	}
}

// isValidBoundaryPoint validates a boundary point
func isValidBoundaryPoint(node Node, offset int) bool {
	if offset < 0 {
		return false
	}
	return offset <= node.Length()
}
