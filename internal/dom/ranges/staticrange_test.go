package ranges

import (
	"strings"
	"testing"
)

// TestStaticRangeConstructor tests StaticRange constructor per WHATWG DOM Section 5.4
func TestStaticRangeConstructor(t *testing.T) {
	doc := &mockDocument{}
	elem := &mockElement{nodeType: 1, nodeName: "div", ownerDocument: doc}
	text := &mockText{nodeType: 3, nodeName: "#text", data: "Hello", ownerDocument: doc}
	doctype := &mockDocumentType{nodeType: 10, nodeName: "html", ownerDocument: doc}
	attr := &mockAttr{nodeType: 2, nodeName: "id", ownerDocument: doc}

	tests := []struct {
		name          string
		init          StaticRangeInit
		shouldError   bool
		expectedError string
	}{
		{
			name: "valid range",
			init: StaticRangeInit{
				StartContainer: elem,
				StartOffset:    0,
				EndContainer:   text,
				EndOffset:      5,
			},
			shouldError: false,
		},
		{
			name: "DocumentType startContainer",
			init: StaticRangeInit{
				StartContainer: doctype,
				StartOffset:    0,
				EndContainer:   elem,
				EndOffset:      0,
			},
			shouldError:   true,
			expectedError: "InvalidNodeTypeError",
		},
		{
			name: "DocumentType endContainer",
			init: StaticRangeInit{
				StartContainer: elem,
				StartOffset:    0,
				EndContainer:   doctype,
				EndOffset:      0,
			},
			shouldError:   true,
			expectedError: "InvalidNodeTypeError",
		},
		{
			name: "Attr startContainer",
			init: StaticRangeInit{
				StartContainer: attr,
				StartOffset:    0,
				EndContainer:   elem,
				EndOffset:      0,
			},
			shouldError:   true,
			expectedError: "InvalidNodeTypeError",
		},
		{
			name: "Attr endContainer",
			init: StaticRangeInit{
				StartContainer: elem,
				StartOffset:    0,
				EndContainer:   attr,
				EndOffset:      0,
			},
			shouldError:   true,
			expectedError: "InvalidNodeTypeError",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			staticRange, err := NewStaticRange(tt.init)

			if tt.shouldError {
				if err == nil {
					t.Error("Expected NewStaticRange to fail")
				} else if !strings.Contains(err.Error(), tt.expectedError) {
					t.Errorf("Expected error to contain %s, got %v", tt.expectedError, err)
				}
			} else {
				if err != nil {
					t.Errorf("Expected NewStaticRange to succeed, got error: %v", err)
				}
				if staticRange == nil {
					t.Error("Expected NewStaticRange to return a StaticRange")
				} else {
					// Test that properties are set correctly
					if staticRange.StartContainer() != tt.init.StartContainer {
						t.Error("Expected startContainer to be set correctly")
					}
					if staticRange.StartOffset() != tt.init.StartOffset {
						t.Errorf("Expected startOffset to be %d, got %d", tt.init.StartOffset, staticRange.StartOffset())
					}
					if staticRange.EndContainer() != tt.init.EndContainer {
						t.Error("Expected endContainer to be set correctly")
					}
					if staticRange.EndOffset() != tt.init.EndOffset {
						t.Errorf("Expected endOffset to be %d, got %d", tt.init.EndOffset, staticRange.EndOffset())
					}
				}
			}
		})
	}
}

// TestStaticRangeValidation tests StaticRange validation per WHATWG DOM specification
func TestStaticRangeValidation(t *testing.T) {
	doc := &mockDocument{}
	root := &mockElement{nodeType: 1, nodeName: "root", ownerDocument: doc}
	child1 := &mockElement{nodeType: 1, nodeName: "child1", ownerDocument: doc, parent: root}
	text1 := &mockText{nodeType: 3, nodeName: "#text", data: "Hello", ownerDocument: doc, parent: child1}

	// Set up relationships
	root.children = []Node{child1}
	child1.children = []Node{text1}

	// Create another document tree for testing cross-document ranges
	doc2 := &mockDocument{}
	elem2 := &mockElement{nodeType: 1, nodeName: "div", ownerDocument: doc2}

	tests := []struct {
		name     string
		init     StaticRangeInit
		expected bool
	}{
		{
			name: "valid same node range",
			init: StaticRangeInit{
				StartContainer: text1,
				StartOffset:    0,
				EndContainer:   text1,
				EndOffset:      5,
			},
			expected: true,
		},
		{
			name: "valid collapsed range",
			init: StaticRangeInit{
				StartContainer: text1,
				StartOffset:    2,
				EndContainer:   text1,
				EndOffset:      2,
			},
			expected: true,
		},
		{
			name: "start offset out of range",
			init: StaticRangeInit{
				StartContainer: text1,
				StartOffset:    10, // text1 has length 5
				EndContainer:   text1,
				EndOffset:      5,
			},
			expected: false,
		},
		{
			name: "end offset out of range",
			init: StaticRangeInit{
				StartContainer: text1,
				StartOffset:    0,
				EndContainer:   text1,
				EndOffset:      10, // text1 has length 5
			},
			expected: false,
		},
		{
			name: "start after end",
			init: StaticRangeInit{
				StartContainer: text1,
				StartOffset:    3,
				EndContainer:   text1,
				EndOffset:      1,
			},
			expected: false,
		},
		{
			name: "different document trees",
			init: StaticRangeInit{
				StartContainer: text1,
				StartOffset:    0,
				EndContainer:   elem2,
				EndOffset:      0,
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			staticRange, err := NewStaticRange(tt.init)
			if err != nil {
				t.Fatalf("Failed to create StaticRange: %v", err)
			}

			valid := staticRange.IsValid()
			if valid != tt.expected {
				t.Errorf("Expected IsValid() to return %v, got %v", tt.expected, valid)
			}
		})
	}
}

// TestStaticRangeCollapsed tests StaticRange collapsed property
func TestStaticRangeCollapsed(t *testing.T) {
	doc := &mockDocument{}
	text := &mockText{nodeType: 3, nodeName: "#text", data: "Hello", ownerDocument: doc}

	tests := []struct {
		name        string
		startOffset int
		endOffset   int
		expected    bool
	}{
		{
			name:        "collapsed at start",
			startOffset: 0,
			endOffset:   0,
			expected:    true,
		},
		{
			name:        "collapsed in middle",
			startOffset: 2,
			endOffset:   2,
			expected:    true,
		},
		{
			name:        "collapsed at end",
			startOffset: 5,
			endOffset:   5,
			expected:    true,
		},
		{
			name:        "not collapsed",
			startOffset: 0,
			endOffset:   3,
			expected:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			init := StaticRangeInit{
				StartContainer: text,
				StartOffset:    tt.startOffset,
				EndContainer:   text,
				EndOffset:      tt.endOffset,
			}

			staticRange, err := NewStaticRange(init)
			if err != nil {
				t.Fatalf("Failed to create StaticRange: %v", err)
			}

			collapsed := staticRange.Collapsed()
			if collapsed != tt.expected {
				t.Errorf("Expected Collapsed() to return %v, got %v", tt.expected, collapsed)
			}
		})
	}
}

// Additional mock types for testing

type mockAttr struct {
	nodeType      int
	nodeName      string
	ownerDocument Document
}

func (m *mockAttr) NodeType() int           { return m.nodeType }
func (m *mockAttr) ParentNode() Node        { return nil }
func (m *mockAttr) ChildNodes() NodeList    { return &mockNodeList{nodes: []Node{}} }
func (m *mockAttr) OwnerDocument() Document { return m.ownerDocument }
func (m *mockAttr) NextSibling() Node       { return nil }
func (m *mockAttr) PreviousSibling() Node   { return nil }
func (m *mockAttr) Length() int             { return 0 }
