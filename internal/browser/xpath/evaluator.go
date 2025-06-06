package xpath

import (
	"github.com/dpemmons/DOMulator/internal/dom"
)

// XPathEvaluator implements the Web API XPathEvaluator interface
type XPathEvaluator struct {
	// Document for context when needed
	document *dom.Document
}

// NewXPathEvaluator creates a new XPathEvaluator instance
func NewXPathEvaluator() *XPathEvaluator {
	return &XPathEvaluator{}
}

// CreateExpression compiles an XPath expression with optional namespace resolver
func (xe *XPathEvaluator) CreateExpression(expression string, resolver NSResolver) (*XPathExpression, error) {
	expr, err := ParseXPathExpression(expression)
	if err != nil {
		return nil, err
	}

	return &XPathExpression{
		expression: expr,
		resolver:   resolver,
	}, nil
}

// CreateNSResolver creates a namespace resolver (deprecated but still used)
func (xe *XPathEvaluator) CreateNSResolver(node dom.Node) NSResolver {
	// Per spec, this method is deprecated and should return the input as-is
	// For now, return a simple resolver
	return &SimpleNSResolver{node: node}
}

// Evaluate evaluates an XPath expression and returns results of the specified type
func (xe *XPathEvaluator) Evaluate(expression string, contextNode dom.Node, resolver NSResolver, resultType int, result *XPathResult) (*XPathResult, error) {
	expr, err := ParseXPathExpression(expression)
	if err != nil {
		return nil, err
	}

	xpathExpr := &XPathExpression{
		expression: expr,
		resolver:   resolver,
	}

	return xpathExpr.Evaluate(contextNode, resultType, result)
}

// NSResolver interface for namespace resolution
type NSResolver interface {
	LookupNamespaceURI(prefix string) string
}

// SimpleNSResolver provides a basic namespace resolver
type SimpleNSResolver struct {
	node dom.Node
}

// LookupNamespaceURI looks up namespace URI for a prefix
func (r *SimpleNSResolver) LookupNamespaceURI(prefix string) string {
	// For now, return empty string (no namespace)
	// In a full implementation, this would use the node's namespace context
	return ""
}
