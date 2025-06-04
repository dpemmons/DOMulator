package css

import (
	"strings"

	"github.com/dpemmons/DOMulator/internal/dom"
)

// SimpleSelector represents a single CSS selector part.
type SimpleSelector struct {
	tag     string
	id      string
	classes []string
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

// parseSimpleSelector parses a single selector part like "#id", ".class", "tag", "#id.class", etc.
func parseSimpleSelector(selectorPart string) (*SimpleSelector, error) {
	s := &SimpleSelector{}
	remaining := selectorPart

	// 1. Extract ID (e.g., #myid)
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

	// 2. Extract Classes (e.g., .myclass)
	classParts := strings.Split(remaining, ".")
	if len(classParts) > 1 {
		for _, class := range classParts[1:] {
			if class != "" { // Avoid empty strings from leading/trailing dots
				s.classes = append(s.classes, class)
			}
		}
		remaining = classParts[0] // The part before the first dot
	}

	// 3. Extract Tag (the remaining part)
	s.tag = strings.TrimSpace(remaining)

	return s, nil
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

	return true
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
