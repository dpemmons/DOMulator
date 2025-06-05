package ranges

// StaticRangeInit represents the dictionary used to initialize a StaticRange
// per WHATWG DOM Section 5.4
type StaticRangeInit struct {
	StartContainer Node
	StartOffset    int
	EndContainer   Node
	EndOffset      int
}

// StaticRange represents a static range that does not update when the node tree mutates
// per WHATWG DOM Section 5.4
type StaticRange struct {
	AbstractRangeImpl
}

// NewStaticRange creates a new StaticRange with the given initialization parameters
// The new StaticRange(init) constructor steps are:
//  1. If init["startContainer"] or init["endContainer"] is a DocumentType or Attr node,
//     then throw an "InvalidNodeTypeError" DOMException.
//  2. Set this's start to (init["startContainer"], init["startOffset"]) and
//     end to (init["endContainer"], init["endOffset"]).
func NewStaticRange(init StaticRangeInit) (*StaticRange, error) {
	// Step 1: Validate node types
	if init.StartContainer.NodeType() == DocumentTypeNode {
		return nil, &DOMException{
			Code:    InvalidNodeTypeError,
			Message: "InvalidNodeTypeError",
		}
	}

	if init.EndContainer.NodeType() == DocumentTypeNode {
		return nil, &DOMException{
			Code:    InvalidNodeTypeError,
			Message: "InvalidNodeTypeError",
		}
	}

	// Check for Attr nodes (node type 2)
	if init.StartContainer.NodeType() == AttributeNode {
		return nil, &DOMException{
			Code:    InvalidNodeTypeError,
			Message: "InvalidNodeTypeError",
		}
	}

	if init.EndContainer.NodeType() == AttributeNode {
		return nil, &DOMException{
			Code:    InvalidNodeTypeError,
			Message: "InvalidNodeTypeError",
		}
	}

	// TODO: Add Attr node check when Attr nodes are properly implemented
	// if init.StartContainer.NodeType() == AttributeNode {
	//     return nil, &DOMException{
	//         Code:    InvalidNodeTypeError,
	//         Message: "startContainer cannot be an Attr node",
	//     }
	// }

	// Step 2: Set start and end boundary points
	staticRange := &StaticRange{
		AbstractRangeImpl: AbstractRangeImpl{
			start: BoundaryPoint{
				Node:   init.StartContainer,
				Offset: init.StartOffset,
			},
			end: BoundaryPoint{
				Node:   init.EndContainer,
				Offset: init.EndOffset,
			},
		},
	}

	return staticRange, nil
}

// IsValid checks if a StaticRange is valid per WHATWG DOM Section 5.4
// A StaticRange is valid if all of the following are true:
// - Its start and end are in the same node tree.
// - Its start offset is between 0 and its start node's length, inclusive.
// - Its end offset is between 0 and its end node's length, inclusive.
// - Its start is before or equal to its end.
func (sr *StaticRange) IsValid() bool {
	// Check if start and end are in the same node tree
	if !inSameTree(sr.start.Node, sr.end.Node) {
		return false
	}

	// Check if start offset is valid
	startLength := getNodeLength(sr.start.Node)
	if sr.start.Offset < 0 || sr.start.Offset > startLength {
		return false
	}

	// Check if end offset is valid
	endLength := getNodeLength(sr.end.Node)
	if sr.end.Offset < 0 || sr.end.Offset > endLength {
		return false
	}

	// Check if start is before or equal to end
	comparison := CompareBoundaryPoints(&sr.start, &sr.end)
	if comparison == 1 { // start is after end
		return false
	}

	return true
}

// inSameTree checks if two nodes are in the same node tree
func inSameTree(nodeA, nodeB Node) bool {
	// Get the root of both nodes and compare
	rootA := getRoot(nodeA)
	rootB := getRoot(nodeB)
	return rootA == rootB
}
