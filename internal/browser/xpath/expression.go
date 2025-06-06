package xpath

import (
	"fmt"

	"github.com/dpemmons/DOMulator/internal/dom"
)

// XPathExpression represents a compiled XPath expression
type XPathExpression struct {
	expression *ParsedExpression
	resolver   NSResolver
}

// Evaluate executes the XPath expression against a context node
func (xe *XPathExpression) Evaluate(contextNode dom.Node, resultType int, result *XPathResult) (*XPathResult, error) {
	if contextNode == nil {
		return nil, fmt.Errorf("context node cannot be nil")
	}

	// Execute the parsed expression
	nodes, err := xe.expression.Execute(contextNode)
	if err != nil {
		return nil, err
	}

	// Convert results based on requested type
	if result == nil {
		result = &XPathResult{}
	}

	result.resultType = resultType

	switch resultType {
	case UNORDERED_NODE_ITERATOR_TYPE, ORDERED_NODE_ITERATOR_TYPE:
		result.nodes = nodes
		result.iteratorIndex = 0

	case UNORDERED_NODE_SNAPSHOT_TYPE, ORDERED_NODE_SNAPSHOT_TYPE:
		result.nodes = nodes
		result.snapshotLength = len(nodes)

	case FIRST_ORDERED_NODE_TYPE:
		if len(nodes) > 0 {
			result.singleNodeValue = nodes[0]
		}

	case ANY_UNORDERED_NODE_TYPE:
		if len(nodes) > 0 {
			result.singleNodeValue = nodes[0]
		}

	case NUMBER_TYPE:
		result.numberValue = float64(len(nodes))

	case STRING_TYPE:
		if len(nodes) > 0 {
			if element, ok := nodes[0].(*dom.Element); ok {
				result.stringValue = element.TextContent()
			} else {
				result.stringValue = nodes[0].NodeValue()
			}
		}

	case BOOLEAN_TYPE:
		result.booleanValue = len(nodes) > 0

	default:
		return nil, fmt.Errorf("unsupported result type: %d", resultType)
	}

	return result, nil
}
