package css

import (
	"strings"

	"github.com/dpemmons/DOMulator/internal/dom"
)

// Selector represents a compiled CSS selector.
type Selector struct {
	tag     string
	id      string
	classes []string
	// Add more fields for other selector types (e.g., attributes, pseudo-classes)
}

// CompileSelector parses a CSS selector string into a compilable form.
// This is a very basic implementation and will only handle simple selectors like tag, #id, .class.
func CompileSelector(selector string) (*Selector, error) {
	s := &Selector{}
	remaining := selector

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
	s.tag = remaining

	return s, nil
}

// Matches checks if a given dom.Node matches this selector.
// This is a very basic implementation.
func (s *Selector) Matches(n dom.Node) bool {
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

// QuerySelector returns the first element within the given root that matches the specified selector.
func QuerySelector(root dom.Node, selector string) dom.Node {
	sel, err := CompileSelector(selector)
	if err != nil {
		return nil // Or return error
	}

	var foundNode dom.Node
	dom.Traverse(root, func(n dom.Node) bool {
		if sel.Matches(n) {
			foundNode = n
			return false // Stop traversal after finding the first match
		}
		return true // Continue traversal
	})
	return foundNode
}

// QuerySelectorAll returns a NodeList of all elements within the given root that match the specified selector.
func QuerySelectorAll(root dom.Node, selector string) dom.NodeList {
	sel, err := CompileSelector(selector)
	if err != nil {
		return nil // Or return error
	}

	var matchingNodes dom.NodeList
	dom.Traverse(root, func(n dom.Node) bool {
		if sel.Matches(n) {
			matchingNodes = append(matchingNodes, n)
		}
		return true // Continue traversal
	})
	return matchingNodes
}
