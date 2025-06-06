package xpath

import (
	"fmt"

	"github.com/dpemmons/DOMulator/internal/dom"
)

// XPathResult type constants
const (
	ANY_TYPE                     = 0
	NUMBER_TYPE                  = 1
	STRING_TYPE                  = 2
	BOOLEAN_TYPE                 = 3
	UNORDERED_NODE_ITERATOR_TYPE = 4
	ORDERED_NODE_ITERATOR_TYPE   = 5
	UNORDERED_NODE_SNAPSHOT_TYPE = 6
	ORDERED_NODE_SNAPSHOT_TYPE   = 7
	ANY_UNORDERED_NODE_TYPE      = 8
	FIRST_ORDERED_NODE_TYPE      = 9
)

// XPathResult represents the result of an XPath evaluation
type XPathResult struct {
	// Result type
	resultType int

	// Value results
	numberValue     float64
	stringValue     string
	booleanValue    bool
	singleNodeValue dom.Node

	// Node collection results
	nodes          []dom.Node
	snapshotLength int

	// Iterator state
	iteratorIndex        int
	invalidIteratorState bool
}

// ResultType returns the type of this result
func (xr *XPathResult) ResultType() int {
	return xr.resultType
}

// NumberValue returns the numeric value (for NUMBER_TYPE)
func (xr *XPathResult) NumberValue() float64 {
	if xr.resultType != NUMBER_TYPE {
		panic("XPathResult is not a number type")
	}
	return xr.numberValue
}

// StringValue returns the string value (for STRING_TYPE)
func (xr *XPathResult) StringValue() string {
	if xr.resultType != STRING_TYPE {
		panic("XPathResult is not a string type")
	}
	return xr.stringValue
}

// BooleanValue returns the boolean value (for BOOLEAN_TYPE)
func (xr *XPathResult) BooleanValue() bool {
	if xr.resultType != BOOLEAN_TYPE {
		panic("XPathResult is not a boolean type")
	}
	return xr.booleanValue
}

// SingleNodeValue returns the single node value (for node types)
func (xr *XPathResult) SingleNodeValue() dom.Node {
	if xr.resultType != ANY_UNORDERED_NODE_TYPE && xr.resultType != FIRST_ORDERED_NODE_TYPE {
		panic("XPathResult is not a single node type")
	}
	return xr.singleNodeValue
}

// SnapshotLength returns the number of nodes in the snapshot
func (xr *XPathResult) SnapshotLength() int {
	if xr.resultType != UNORDERED_NODE_SNAPSHOT_TYPE && xr.resultType != ORDERED_NODE_SNAPSHOT_TYPE {
		panic("XPathResult is not a snapshot type")
	}
	return xr.snapshotLength
}

// SnapshotItem returns the node at the specified index in the snapshot
func (xr *XPathResult) SnapshotItem(index int) dom.Node {
	if xr.resultType != UNORDERED_NODE_SNAPSHOT_TYPE && xr.resultType != ORDERED_NODE_SNAPSHOT_TYPE {
		panic("XPathResult is not a snapshot type")
	}

	if index < 0 || index >= len(xr.nodes) {
		return nil
	}

	return xr.nodes[index]
}

// IterateNext returns the next node in the iterator
func (xr *XPathResult) IterateNext() dom.Node {
	if xr.resultType != UNORDERED_NODE_ITERATOR_TYPE && xr.resultType != ORDERED_NODE_ITERATOR_TYPE {
		panic("XPathResult is not an iterator type")
	}

	if xr.invalidIteratorState {
		panic("The document has been mutated since the result was returned")
	}

	if xr.iteratorIndex >= len(xr.nodes) {
		return nil
	}

	node := xr.nodes[xr.iteratorIndex]
	xr.iteratorIndex++
	return node
}

// InvalidateIterator marks the iterator as invalid (called when DOM is mutated)
func (xr *XPathResult) InvalidateIterator() {
	if xr.resultType == UNORDERED_NODE_ITERATOR_TYPE || xr.resultType == ORDERED_NODE_ITERATOR_TYPE {
		xr.invalidIteratorState = true
	}
}

// String returns a string representation of the result
func (xr *XPathResult) String() string {
	switch xr.resultType {
	case NUMBER_TYPE:
		return fmt.Sprintf("XPathResult{NUMBER: %g}", xr.numberValue)
	case STRING_TYPE:
		return fmt.Sprintf("XPathResult{STRING: %q}", xr.stringValue)
	case BOOLEAN_TYPE:
		return fmt.Sprintf("XPathResult{BOOLEAN: %t}", xr.booleanValue)
	case ANY_UNORDERED_NODE_TYPE, FIRST_ORDERED_NODE_TYPE:
		if xr.singleNodeValue != nil {
			return fmt.Sprintf("XPathResult{NODE: %s}", xr.singleNodeValue.NodeName())
		}
		return "XPathResult{NODE: null}"
	case UNORDERED_NODE_SNAPSHOT_TYPE, ORDERED_NODE_SNAPSHOT_TYPE:
		return fmt.Sprintf("XPathResult{SNAPSHOT: %d nodes}", xr.snapshotLength)
	case UNORDERED_NODE_ITERATOR_TYPE, ORDERED_NODE_ITERATOR_TYPE:
		return fmt.Sprintf("XPathResult{ITERATOR: %d nodes}", len(xr.nodes))
	default:
		return fmt.Sprintf("XPathResult{UNKNOWN: %d}", xr.resultType)
	}
}
