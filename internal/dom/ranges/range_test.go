package ranges

import (
	"strings"
	"testing"
)

// Range comparison constants per WHATWG DOM specification
const (
	StartToStart = 0
	StartToEnd   = 1
	EndToEnd     = 2
	EndToStart   = 3
)

// TestRangeConstructor tests Range constructor per WHATWG DOM Section 5.5
func TestRangeConstructor(t *testing.T) {
	doc := &mockDocument{}
	docAdapter := &mockDocumentAdapter{document: doc}

	range_ := NewRange(docAdapter)
	if range_ == nil {
		t.Error("Expected NewRange to return a Range")
	}

	// Test initial state per specification - should be (document, 0) to (document, 0)
	if range_.StartContainer() != doc {
		t.Error("Expected startContainer to be the document")
	}
	if range_.StartOffset() != 0 {
		t.Errorf("Expected startOffset to be 0, got %d", range_.StartOffset())
	}
	if range_.EndContainer() != doc {
		t.Error("Expected endContainer to be the document")
	}
	if range_.EndOffset() != 0 {
		t.Errorf("Expected endOffset to be 0, got %d", range_.EndOffset())
	}
	if !range_.Collapsed() {
		t.Error("Expected new range to be collapsed")
	}
}

// TestRangeSetStart tests setStart method per WHATWG DOM specification
func TestRangeSetStart(t *testing.T) {
	doc := &mockDocument{}
	docAdapter := &mockDocumentAdapter{document: doc}
	elem := &mockElement{nodeType: 1, nodeName: "div", ownerDocument: doc}
	text := &mockText{nodeType: 3, nodeName: "#text", data: "Hello", ownerDocument: doc}
	doctype := &mockDocumentType{nodeType: 10, nodeName: "html", ownerDocument: doc}

	tests := []struct {
		name          string
		node          Node
		offset        int
		shouldError   bool
		expectedError string
	}{
		{
			name:        "valid element node",
			node:        elem,
			offset:      0,
			shouldError: false,
		},
		{
			name:        "valid text node",
			node:        text,
			offset:      3,
			shouldError: false,
		},
		{
			name:          "doctype node",
			node:          doctype,
			offset:        0,
			shouldError:   true,
			expectedError: "InvalidNodeTypeError",
		},
		{
			name:          "offset beyond node length",
			node:          text,
			offset:        10, // text has length 5
			shouldError:   true,
			expectedError: "IndexSizeError",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			range_ := NewRange(docAdapter)
			err := range_.SetStart(tt.node, tt.offset)

			if tt.shouldError {
				if err == nil {
					t.Error("Expected SetStart to fail")
				} else if !strings.Contains(err.Error(), tt.expectedError) {
					t.Errorf("Expected error to contain %s, got %v", tt.expectedError, err)
				}
			} else {
				if err != nil {
					t.Errorf("Expected SetStart to succeed, got error: %v", err)
				}
				if range_.StartContainer() != tt.node {
					t.Error("Expected startContainer to be set correctly")
				}
				if range_.StartOffset() != tt.offset {
					t.Errorf("Expected startOffset to be %d, got %d", tt.offset, range_.StartOffset())
				}
			}
		})
	}
}

// TestRangeSetEnd tests setEnd method per WHATWG DOM specification
func TestRangeSetEnd(t *testing.T) {
	doc := &mockDocument{}
	docAdapter := &mockDocumentAdapter{document: doc}
	elem := &mockElement{nodeType: 1, nodeName: "div", ownerDocument: doc}
	text := &mockText{nodeType: 3, nodeName: "#text", data: "Hello", ownerDocument: doc}

	range_ := NewRange(docAdapter)

	// Set up a valid start position first
	err := range_.SetStart(elem, 0)
	if err != nil {
		t.Fatalf("Failed to set start: %v", err)
	}

	// Test setting end
	err = range_.SetEnd(text, 3)
	if err != nil {
		t.Errorf("Expected SetEnd to succeed, got error: %v", err)
	}

	if range_.EndContainer() != text {
		t.Error("Expected endContainer to be set correctly")
	}
	if range_.EndOffset() != 3 {
		t.Errorf("Expected endOffset to be 3, got %d", range_.EndOffset())
	}
}

// TestRangeSetStartBefore tests setStartBefore method
func TestRangeSetStartBefore(t *testing.T) {
	doc := &mockDocument{}
	docAdapter := &mockDocumentAdapter{document: doc}
	parent := &mockElement{nodeType: 1, nodeName: "parent", ownerDocument: doc}
	child := &mockElement{nodeType: 1, nodeName: "child", ownerDocument: doc, parent: parent}

	// Set up parent-child relationship
	parent.children = []Node{child}

	tests := []struct {
		name          string
		node          Node
		shouldError   bool
		expectedError string
	}{
		{
			name:        "valid node with parent",
			node:        child,
			shouldError: false,
		},
		{
			name:          "node without parent",
			node:          parent, // has no parent in our setup
			shouldError:   true,
			expectedError: "InvalidNodeTypeError",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			range_ := NewRange(docAdapter)
			err := range_.SetStartBefore(tt.node)

			if tt.shouldError {
				if err == nil {
					t.Error("Expected SetStartBefore to fail")
				} else if !strings.Contains(err.Error(), tt.expectedError) {
					t.Errorf("Expected error to contain %s, got %v", tt.expectedError, err)
				}
			} else {
				if err != nil {
					t.Errorf("Expected SetStartBefore to succeed, got error: %v", err)
				}
				// Should set start to (parent, child's index)
				if range_.StartContainer() != parent {
					t.Error("Expected startContainer to be the parent")
				}
				if range_.StartOffset() != 0 { // child is at index 0
					t.Errorf("Expected startOffset to be 0, got %d", range_.StartOffset())
				}
			}
		})
	}
}

// TestRangeCollapse tests collapse method
func TestRangeCollapse(t *testing.T) {
	doc := &mockDocument{}
	docAdapter := &mockDocumentAdapter{document: doc}
	elem := &mockElement{nodeType: 1, nodeName: "div", ownerDocument: doc}
	text := &mockText{nodeType: 3, nodeName: "#text", data: "Hello", ownerDocument: doc}

	range_ := NewRange(docAdapter)

	// Set up a non-collapsed range
	range_.SetStart(elem, 0)
	range_.SetEnd(text, 3)

	// Test collapse to start (default)
	range_.Collapse(true)
	if !range_.Collapsed() {
		t.Error("Expected range to be collapsed after Collapse(true)")
	}
	if range_.EndContainer() != range_.StartContainer() {
		t.Error("Expected end container to match start container after collapse")
	}
	if range_.EndOffset() != range_.StartOffset() {
		t.Error("Expected end offset to match start offset after collapse")
	}

	// Set up non-collapsed range again
	range_.SetStart(elem, 0)
	range_.SetEnd(text, 3)

	// Test collapse to end
	range_.Collapse(false)
	if !range_.Collapsed() {
		t.Error("Expected range to be collapsed after Collapse(false)")
	}
	if range_.StartContainer() != text {
		t.Error("Expected start container to match end container after collapse to end")
	}
	if range_.StartOffset() != 3 {
		t.Errorf("Expected start offset to be 3 after collapse to end, got %d", range_.StartOffset())
	}
}

// TestRangeSelectNode tests selectNode method
func TestRangeSelectNode(t *testing.T) {
	doc := &mockDocument{}
	docAdapter := &mockDocumentAdapter{document: doc}
	parent := &mockElement{nodeType: 1, nodeName: "parent", ownerDocument: doc}
	child := &mockElement{nodeType: 1, nodeName: "child", ownerDocument: doc, parent: parent}

	// Set up parent-child relationship
	parent.children = []Node{child}

	tests := []struct {
		name          string
		node          Node
		shouldError   bool
		expectedError string
	}{
		{
			name:        "valid node with parent",
			node:        child,
			shouldError: false,
		},
		{
			name:          "node without parent",
			node:          parent,
			shouldError:   true,
			expectedError: "InvalidNodeTypeError",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			range_ := NewRange(docAdapter)
			err := range_.SelectNode(tt.node)

			if tt.shouldError {
				if err == nil {
					t.Error("Expected SelectNode to fail")
				} else if !strings.Contains(err.Error(), tt.expectedError) {
					t.Errorf("Expected error to contain %s, got %v", tt.expectedError, err)
				}
			} else {
				if err != nil {
					t.Errorf("Expected SelectNode to succeed, got error: %v", err)
				}
				// Should set range to (parent, index) to (parent, index+1)
				if range_.StartContainer() != parent {
					t.Error("Expected startContainer to be the parent")
				}
				if range_.EndContainer() != parent {
					t.Error("Expected endContainer to be the parent")
				}
				if range_.StartOffset() != 0 {
					t.Errorf("Expected startOffset to be 0, got %d", range_.StartOffset())
				}
				if range_.EndOffset() != 1 {
					t.Errorf("Expected endOffset to be 1, got %d", range_.EndOffset())
				}
			}
		})
	}
}

// TestRangeSelectNodeContents tests selectNodeContents method
func TestRangeSelectNodeContents(t *testing.T) {
	doc := &mockDocument{}
	docAdapter := &mockDocumentAdapter{document: doc}
	elem := &mockElement{nodeType: 1, nodeName: "div", ownerDocument: doc}
	text := &mockText{nodeType: 3, nodeName: "#text", data: "Hello", ownerDocument: doc}
	doctype := &mockDocumentType{nodeType: 10, nodeName: "html", ownerDocument: doc}

	tests := []struct {
		name          string
		node          Node
		shouldError   bool
		expectedError string
		expectedStart int
		expectedEnd   int
	}{
		{
			name:          "valid element node",
			node:          elem,
			shouldError:   false,
			expectedStart: 0,
			expectedEnd:   0, // elem has no children
		},
		{
			name:          "valid text node",
			node:          text,
			shouldError:   false,
			expectedStart: 0,
			expectedEnd:   5, // text length
		},
		{
			name:          "doctype node",
			node:          doctype,
			shouldError:   true,
			expectedError: "InvalidNodeTypeError",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			range_ := NewRange(docAdapter)
			err := range_.SelectNodeContents(tt.node)

			if tt.shouldError {
				if err == nil {
					t.Error("Expected SelectNodeContents to fail")
				} else if !strings.Contains(err.Error(), tt.expectedError) {
					t.Errorf("Expected error to contain %s, got %v", tt.expectedError, err)
				}
			} else {
				if err != nil {
					t.Errorf("Expected SelectNodeContents to succeed, got error: %v", err)
				}
				if range_.StartContainer() != tt.node {
					t.Error("Expected startContainer to be the node")
				}
				if range_.EndContainer() != tt.node {
					t.Error("Expected endContainer to be the node")
				}
				if range_.StartOffset() != tt.expectedStart {
					t.Errorf("Expected startOffset to be %d, got %d", tt.expectedStart, range_.StartOffset())
				}
				if range_.EndOffset() != tt.expectedEnd {
					t.Errorf("Expected endOffset to be %d, got %d", tt.expectedEnd, range_.EndOffset())
				}
			}
		})
	}
}

// TestRangeCompareBoundaryPoints tests compareBoundaryPoints method
func TestRangeCompareBoundaryPoints(t *testing.T) {
	doc := &mockDocument{}
	// Initialize children slice to avoid nil slice issues
	if doc.children == nil {
		doc.children = make([]Node, 0)
	}

	// Set up mock nodes with proper parent relationships
	elem := &mockElement{nodeType: 1, nodeName: "div", ownerDocument: doc, parent: doc}
	text := &mockText{nodeType: 3, nodeName: "#text", data: "Hello", ownerDocument: doc, parent: doc}

	// Set the children of 'doc' in the order required by the test logic.
	// For range1.start (elem,0) to be before range2.start (text,1),
	// 'elem' must appear before 'text' in the document's child list.
	doc.children = []Node{elem, text}

	docAdapter := &mockDocumentAdapter{document: doc}

	range1 := NewRange(docAdapter)
	range1.SetStart(elem, 0)
	range1.SetEnd(text, 3)

	range2 := NewRange(docAdapter)
	range2.SetStart(text, 1)
	range2.SetEnd(text, 5)

	tests := []struct {
		name     string
		how      int
		expected int
	}{
		{
			name:     "START_TO_START",
			how:      StartToStart,
			expected: -1, // range1.start is before range2.start
		},
		{
			name:     "START_TO_END",
			how:      StartToEnd,
			expected: -1, // range1.end is before range2.start
		},
		{
			name:     "END_TO_END",
			how:      EndToEnd,
			expected: -1, // range1.end is before range2.end
		},
		{
			name:     "END_TO_START",
			how:      EndToStart,
			expected: 1, // range1.start is after range2.end
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := range1.CompareBoundaryPoints(tt.how, range2)
			if err != nil {
				t.Errorf("Expected CompareBoundaryPoints to succeed, got error: %v", err)
			}
			if result != tt.expected {
				t.Errorf("Expected result to be %d, got %d", tt.expected, result)
			}
		})
	}

	// Test invalid how parameter
	_, err := range1.CompareBoundaryPoints(999, range2)
	if err == nil {
		t.Error("Expected CompareBoundaryPoints to fail with invalid how parameter")
	} else if !strings.Contains(err.Error(), "NotSupportedError") {
		t.Errorf("Expected NotSupportedError, got %v", err)
	}
}

// TestRangeIsPointInRange tests isPointInRange method
func TestRangeIsPointInRange(t *testing.T) {
	doc := &mockDocument{}
	docAdapter := &mockDocumentAdapter{document: doc}
	text := &mockText{nodeType: 3, nodeName: "#text", data: "Hello", ownerDocument: doc}
	doctype := &mockDocumentType{nodeType: 10, nodeName: "html", ownerDocument: doc}

	range_ := NewRange(docAdapter)
	range_.SetStart(text, 1)
	range_.SetEnd(text, 4)

	tests := []struct {
		name          string
		node          Node
		offset        int
		expected      bool
		shouldError   bool
		expectedError string
	}{
		{
			name:     "point before range",
			node:     text,
			offset:   0,
			expected: false,
		},
		{
			name:     "point at start",
			node:     text,
			offset:   1,
			expected: true,
		},
		{
			name:     "point in middle",
			node:     text,
			offset:   2,
			expected: true,
		},
		{
			name:     "point at end",
			node:     text,
			offset:   4,
			expected: false, // end is exclusive
		},
		{
			name:     "point after range",
			node:     text,
			offset:   5,
			expected: false,
		},
		{
			name:          "doctype node",
			node:          doctype,
			offset:        0,
			shouldError:   true,
			expectedError: "InvalidNodeTypeError",
		},
		{
			name:          "offset beyond node length",
			node:          text,
			offset:        10,
			shouldError:   true,
			expectedError: "IndexSizeError",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := range_.IsPointInRange(tt.node, tt.offset)

			if tt.shouldError {
				if err == nil {
					t.Error("Expected IsPointInRange to fail")
				} else if !strings.Contains(err.Error(), tt.expectedError) {
					t.Errorf("Expected error to contain %s, got %v", tt.expectedError, err)
				}
			} else {
				if err != nil {
					t.Errorf("Expected IsPointInRange to succeed, got error: %v", err)
				}
				if result != tt.expected {
					t.Errorf("Expected result to be %v, got %v", tt.expected, result)
				}
			}
		})
	}
}

// TestRangeCloneRange tests cloneRange method
func TestRangeCloneRange(t *testing.T) {
	doc := &mockDocument{}
	docAdapter := &mockDocumentAdapter{document: doc}
	text := &mockText{nodeType: 3, nodeName: "#text", data: "Hello", ownerDocument: doc}

	original := NewRange(docAdapter)
	original.SetStart(text, 1)
	original.SetEnd(text, 4)

	clone := original.CloneRange()
	if clone == nil {
		t.Error("Expected CloneRange to return a range")
	}

	// Test that clone has same boundary points
	if clone.StartContainer() != original.StartContainer() {
		t.Error("Expected clone startContainer to match original")
	}
	if clone.StartOffset() != original.StartOffset() {
		t.Error("Expected clone startOffset to match original")
	}
	if clone.EndContainer() != original.EndContainer() {
		t.Error("Expected clone endContainer to match original")
	}
	if clone.EndOffset() != original.EndOffset() {
		t.Error("Expected clone endOffset to match original")
	}

	// Test that they are different objects
	if clone == original {
		t.Error("Expected clone to be a different object than original")
	}

	// Test that modifying clone doesn't affect original
	clone.SetStart(text, 0)
	if original.StartOffset() == 0 {
		t.Error("Expected original to be unaffected by clone modification")
	}
}

// Mock document adapter for testing
type mockDocumentAdapter struct {
	document *mockDocument
}

func (m *mockDocumentAdapter) GetDocument() Document {
	return m.document
}

// Implement Document interface by delegating to the document field
func (m *mockDocumentAdapter) NodeType() int           { return m.document.NodeType() }
func (m *mockDocumentAdapter) ParentNode() Node        { return m.document.ParentNode() }
func (m *mockDocumentAdapter) ChildNodes() NodeList    { return m.document.ChildNodes() }
func (m *mockDocumentAdapter) OwnerDocument() Document { return m.document.OwnerDocument() }
func (m *mockDocumentAdapter) NextSibling() Node       { return m.document.NextSibling() }
func (m *mockDocumentAdapter) PreviousSibling() Node   { return m.document.PreviousSibling() }
func (m *mockDocumentAdapter) Length() int             { return m.document.Length() }

// Mock document type for testing
type mockDocumentType struct {
	nodeType      int
	nodeName      string
	ownerDocument Document
	parent        Node
}

func (m *mockDocumentType) NodeType() int           { return m.nodeType }
func (m *mockDocumentType) ParentNode() Node        { return m.parent }
func (m *mockDocumentType) ChildNodes() NodeList    { return &mockNodeList{nodes: []Node{}} }
func (m *mockDocumentType) OwnerDocument() Document { return m.ownerDocument }
func (m *mockDocumentType) NextSibling() Node       { return nil }
func (m *mockDocumentType) PreviousSibling() Node   { return nil }
func (m *mockDocumentType) Length() int             { return 0 }
