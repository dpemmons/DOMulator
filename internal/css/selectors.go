package css

import (
	"strings"

	"github.com/dpemmons/DOMulator/internal/dom"
)

// SimpleSelector represents a single CSS selector part.
type SimpleSelector struct {
	tag        string
	id         string
	classes    []string
	attributes []AttributeSelector
}

// AttributeSelector represents an attribute selector like [name=value]
type AttributeSelector struct {
	name     string
	value    string
	operator string // "=", "~=", "|=", "^=", "$=", "*=", or "" for just existence
}

// Selector represents a compiled CSS selector that may include descendant relationships.
type Selector struct {
	parts []SimpleSelector // For descendant selectors like "ul li"
}

// CompileSelector parses a CSS selector string into a compilable form.
// Supports descendant selectors like "#todo-list li", "ul li", etc.
func CompileSelector(selector string) (*Selector, error) {
	s := &Selector{}

	// Split by whitespace to handle descendant selectors
	parts := strings.Fields(strings.TrimSpace(selector))

	for _, part := range parts {
		simpleSelector, err := parseSimpleSelector(part)
		if err != nil {
			return nil, err
		}
		s.parts = append(s.parts, *simpleSelector)
	}

	return s, nil
}

// parseSimpleSelector parses a single selector part like "#id", ".class", "tag", "#id.class", "[attr=value]", etc.
func parseSimpleSelector(selectorPart string) (*SimpleSelector, error) {
	s := &SimpleSelector{}
	remaining := selectorPart

	// 1. Extract Attribute selectors (e.g., [name=value])
	for {
		startIdx := strings.Index(remaining, "[")
		if startIdx == -1 {
			break
		}
		endIdx := strings.Index(remaining[startIdx:], "]")
		if endIdx == -1 {
			break
		}
		endIdx += startIdx

		attrStr := remaining[startIdx+1 : endIdx]
		attr, err := parseAttributeSelector(attrStr)
		if err != nil {
			return nil, err
		}
		s.attributes = append(s.attributes, *attr)

		// Remove the attribute selector from remaining
		remaining = remaining[:startIdx] + remaining[endIdx+1:]
	}

	// 2. Extract ID (e.g., #myid)
	if idx := strings.Index(remaining, "#"); idx != -1 {
		idPart := remaining[idx+1:]
		if spaceIdx := strings.IndexAny(idPart, ". "); spaceIdx != -1 {
			s.id = idPart[:spaceIdx]
			remaining = remaining[:idx] + idPart[spaceIdx:]
		} else {
			s.id = idPart
			remaining = remaining[:idx]
		}
	}

	// 3. Extract Classes (e.g., .myclass)
	classParts := strings.Split(remaining, ".")
	if len(classParts) > 1 {
		for _, class := range classParts[1:] {
			if class != "" { // Avoid empty strings from leading/trailing dots
				s.classes = append(s.classes, class)
			}
		}
		remaining = classParts[0] // The part before the first dot
	}

	// 4. Extract Tag (the remaining part)
	s.tag = strings.TrimSpace(remaining)

	return s, nil
}

// parseAttributeSelector parses an attribute selector string like "name=value" or "class~=active"
func parseAttributeSelector(attrStr string) (*AttributeSelector, error) {
	attr := &AttributeSelector{}

	// Check for different operators in order of length (longest first)
	operators := []string{"~=", "|=", "^=", "$=", "*=", "="}

	for _, op := range operators {
		if idx := strings.Index(attrStr, op); idx != -1 {
			attr.name = strings.TrimSpace(attrStr[:idx])
			attr.operator = op
			attr.value = strings.Trim(strings.TrimSpace(attrStr[idx+len(op):]), `"'`)
			return attr, nil
		}
	}

	// No operator found, it's just an attribute existence check
	attr.name = strings.TrimSpace(attrStr)
	attr.operator = ""
	attr.value = ""
	return attr, nil
}

// Matches checks if a given dom.Node matches this simple selector.
func (s *SimpleSelector) Matches(n dom.Node) bool {
	// Only ElementNodes can match CSS selectors
	if n.NodeType() != dom.ElementNode {
		return false
	}

	// Check tag
	if s.tag != "" && !strings.EqualFold(s.tag, n.NodeName()) {
		return false
	}

	// Check ID
	if s.id != "" {
		if elem, ok := n.(*dom.Element); ok {
			if elem.GetAttribute("id") != s.id {
				return false
			}
		} else {
			return false // Only elements have IDs
		}
	}

	// Check classes
	if len(s.classes) > 0 {
		if elem, ok := n.(*dom.Element); ok {
			nodeClasses := strings.Fields(elem.GetAttribute("class"))
			for _, requiredClass := range s.classes {
				found := false
				for _, nodeClass := range nodeClasses {
					if requiredClass == nodeClass {
						found = true
						break
					}
				}
				if !found {
					return false
				}
			}
		} else {
			return false // Only elements have classes
		}
	}

	// Check attributes
	if len(s.attributes) > 0 {
		if elem, ok := n.(*dom.Element); ok {
			for _, attr := range s.attributes {
				if !s.matchesAttribute(elem, attr) {
					return false
				}
			}
		} else {
			return false // Only elements have attributes
		}
	}

	return true
}

// matchesAttribute checks if an element matches an attribute selector
func (s *SimpleSelector) matchesAttribute(elem *dom.Element, attr AttributeSelector) bool {
	attrValue := elem.GetAttribute(attr.name)

	switch attr.operator {
	case "":
		// Just check for existence
		return elem.HasAttribute(attr.name)
	case "=":
		// Exact match
		return attrValue == attr.value
	case "~=":
		// Word match (space-separated list)
		words := strings.Fields(attrValue)
		for _, word := range words {
			if word == attr.value {
				return true
			}
		}
		return false
	case "|=":
		// Exact match or starts with value followed by hyphen
		return attrValue == attr.value || strings.HasPrefix(attrValue, attr.value+"-")
	case "^=":
		// Starts with
		return strings.HasPrefix(attrValue, attr.value)
	case "$=":
		// Ends with
		return strings.HasSuffix(attrValue, attr.value)
	case "*=":
		// Contains
		return strings.Contains(attrValue, attr.value)
	default:
		return false
	}
}

// MatchesDescendant checks if a node matches a descendant selector.
// For example, for selector "ul li", it checks if the node is an "li"
// and has an ancestor "ul".
func (s *Selector) MatchesDescendant(n dom.Node) bool {
	if len(s.parts) == 0 {
		return false
	}

	// If it's a simple selector (no descendants), just check the last part
	if len(s.parts) == 1 {
		return s.parts[0].Matches(n)
	}

	// For descendant selectors, the node must match the last part
	lastPart := s.parts[len(s.parts)-1]
	if !lastPart.Matches(n) {
		return false
	}

	// Then check if it has the required ancestors in order
	return s.hasRequiredAncestors(n, len(s.parts)-2)
}

// hasRequiredAncestors checks if the node has the required ancestors
// in the correct order (from closest to furthest).
func (s *Selector) hasRequiredAncestors(n dom.Node, partIndex int) bool {
	if partIndex < 0 {
		return true // All required ancestors found
	}

	requiredAncestor := s.parts[partIndex]

	// Look for the required ancestor
	current := n.ParentNode()
	for current != nil {
		if requiredAncestor.Matches(current) {
			// Found this ancestor, check for the next one
			return s.hasRequiredAncestors(current, partIndex-1)
		}
		current = current.ParentNode()
	}

	return false // Required ancestor not found
}

// QuerySelector returns the first element within the given root that matches the specified selector.
func QuerySelector(root dom.Node, selector string) dom.Node {
	sel, err := CompileSelector(selector)
	if err != nil {
		return nil // Or return error
	}

	var foundNode dom.Node
	dom.Traverse(root, func(n dom.Node) bool {
		if sel.MatchesDescendant(n) {
			foundNode = n
			return false // Stop traversal after finding the first match
		}
		return true // Continue traversal
	})
	return foundNode
}

// QuerySelectorAll returns a NodeList of all elements within the given root that match the specified selector.
func QuerySelectorAll(root dom.Node, selector string) *dom.NodeList {
	sel, err := CompileSelector(selector)
	if err != nil {
		return dom.NewNodeList(nil) // Return empty NodeList on error
	}

	var matchingNodes []dom.Node
	dom.Traverse(root, func(n dom.Node) bool {
		if sel.MatchesDescendant(n) {
			matchingNodes = append(matchingNodes, n)
		}
		return true // Continue traversal
	})
	return dom.NewNodeList(matchingNodes)
}
