package xpath

import (
	"fmt"
	"strings"

	"github.com/dpemmons/DOMulator/internal/dom"
)

// ParsedExpression represents a parsed XPath expression
type ParsedExpression struct {
	pattern    string
	axis       string
	nodeTest   string
	predicates []string
}

// ParseXPathExpression parses an XPath expression string
func ParseXPathExpression(expression string) (*ParsedExpression, error) {
	expression = strings.TrimSpace(expression)

	if expression == "" {
		return nil, fmt.Errorf("empty XPath expression")
	}

	// Handle simple patterns that HTMX might use
	parsed := &ParsedExpression{
		pattern: expression,
	}

	// Parse basic XPath patterns
	if strings.HasPrefix(expression, "//") {
		// Descendant-or-self axis (e.g., "//div", "//div[@class='test']")
		parsed.axis = "descendant-or-self"
		rest := expression[2:]

		// Extract node test and predicates
		if idx := strings.Index(rest, "["); idx != -1 {
			parsed.nodeTest = rest[:idx]
			predicateStr := rest[idx:]
			parsed.predicates = []string{predicateStr}
		} else {
			parsed.nodeTest = rest
		}
	} else if strings.HasPrefix(expression, "/") {
		// Child axis (e.g., "/div", "/html/body")
		parsed.axis = "child"
		parsed.nodeTest = expression[1:]
	} else {
		// Simple node test (e.g., "div", "span")
		parsed.axis = "child"
		parsed.nodeTest = expression
	}

	return parsed, nil
}

// Execute runs the parsed XPath expression against a context node
func (pe *ParsedExpression) Execute(contextNode dom.Node) ([]dom.Node, error) {
	var results []dom.Node

	switch pe.axis {
	case "descendant-or-self":
		results = pe.findDescendantNodes(contextNode)
	case "child":
		results = pe.findChildNodes(contextNode)
	default:
		return nil, fmt.Errorf("unsupported axis: %s", pe.axis)
	}

	// Apply predicates if any
	if len(pe.predicates) > 0 {
		filtered := make([]dom.Node, 0, len(results))
		for _, node := range results {
			if pe.matchesPredicates(node) {
				filtered = append(filtered, node)
			}
		}
		results = filtered
	}

	return results, nil
}

// findDescendantNodes finds all descendant nodes matching the node test
func (pe *ParsedExpression) findDescendantNodes(contextNode dom.Node) []dom.Node {
	var results []dom.Node

	// Include context node if it matches
	if pe.matchesNodeTest(contextNode) {
		results = append(results, contextNode)
	}

	// Recursively search children
	pe.searchNode(contextNode, &results)

	return results
}

// findChildNodes finds direct child nodes matching the node test
func (pe *ParsedExpression) findChildNodes(contextNode dom.Node) []dom.Node {
	var results []dom.Node

	children := contextNode.ChildNodes()
	for i := 0; i < children.Length(); i++ {
		child := children.Item(i)
		if pe.matchesNodeTest(child) {
			results = append(results, child)
		}
	}

	return results
}

// searchNode recursively searches a node and its descendants
func (pe *ParsedExpression) searchNode(node dom.Node, results *[]dom.Node) {
	children := node.ChildNodes()
	for i := 0; i < children.Length(); i++ {
		child := children.Item(i)

		if pe.matchesNodeTest(child) {
			*results = append(*results, child)
		}

		// Recursively search child nodes
		pe.searchNode(child, results)
	}
}

// matchesNodeTest checks if a node matches the node test
func (pe *ParsedExpression) matchesNodeTest(node dom.Node) bool {
	if pe.nodeTest == "*" {
		// Match any element
		return node.NodeType() == dom.ElementNode
	}

	if pe.nodeTest == "." {
		// Match current node
		return true
	}

	if pe.nodeTest == ".." {
		// Match parent node (not implemented for now)
		return false
	}

	// Match by tag name
	if element, ok := node.(*dom.Element); ok {
		return strings.EqualFold(element.TagName(), pe.nodeTest)
	}

	return false
}

// matchesPredicates checks if a node matches all predicates
func (pe *ParsedExpression) matchesPredicates(node dom.Node) bool {
	element, ok := node.(*dom.Element)
	if !ok {
		return false
	}

	for _, predicate := range pe.predicates {
		if !pe.matchesPredicate(element, predicate) {
			return false
		}
	}

	return true
}

// matchesPredicate checks if a node matches a single predicate
func (pe *ParsedExpression) matchesPredicate(element *dom.Element, predicate string) bool {
	// Remove brackets
	predicate = strings.Trim(predicate, "[]")

	// Handle simple attribute tests
	if strings.Contains(predicate, "=") {
		parts := strings.SplitN(predicate, "=", 2)
		if len(parts) != 2 {
			return false
		}

		attrName := strings.TrimSpace(parts[0])
		attrValue := strings.Trim(strings.TrimSpace(parts[1]), "'\"")

		// Handle @attribute syntax
		if strings.HasPrefix(attrName, "@") {
			attrName = attrName[1:]
		}

		return element.GetAttribute(attrName) == attrValue
	}

	// Handle attribute existence tests
	if strings.HasPrefix(predicate, "@") {
		attrName := predicate[1:]
		return element.HasAttribute(attrName)
	}

	// Handle positional predicates (like [1], [2])
	// For now, just return true for numeric predicates
	if _, err := fmt.Sscanf(predicate, "%d"); err == nil {
		return true
	}

	return false
}
