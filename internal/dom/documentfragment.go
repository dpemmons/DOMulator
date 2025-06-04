package dom

import (
	"github.com/dop251/goja"
)

// DocumentFragment represents a lightweight Document object whose contents are not part of the main document.
type DocumentFragment struct {
	nodeImpl
}

// NewDocumentFragment creates a new DocumentFragment node.
func NewDocumentFragment(doc *Document) *DocumentFragment {
	df := &DocumentFragment{
		nodeImpl: nodeImpl{
			nodeType:      DocumentFragmentNode,
			nodeName:      "#document-fragment",
			ownerDocument: doc,
		},
	}
	df.nodeImpl.self = df // Set the self reference
	return df
}

// GetElementById returns the first element within the fragment's descendants whose ID is elementId.
// This implements the NonElementParentNode mixin as specified in WHATWG DOM Standard 4.2.4.
// The method steps are to return the first element, in tree order, within this's descendants,
// whose ID is elementId; otherwise, if there is no such element, null.
func (df *DocumentFragment) GetElementById(id string) *Element {
	if id == "" {
		return nil // Don't search for empty IDs
	}
	return df.findElementByIdRecursive(id)
}

// findElementByIdRecursive performs a depth-first search for an element with the given ID.
// It properly stops traversal when the first match is found.
func (df *DocumentFragment) findElementByIdRecursive(id string) *Element {
	// Search through all direct children
	for _, child := range df.ChildNodes() {
		if elem, ok := child.(*Element); ok {
			// Check if this element has the target ID
			if elem.GetAttribute("id") == id {
				return elem
			}
			// Recursively search in this element's descendants
			if found := elem.findElementByIdRecursive(id); found != nil {
				return found
			}
		}
	}
	return nil
}

// Children returns a live HTMLCollection of all child elements
func (df *DocumentFragment) Children() *HTMLCollection {
	return NewChildElementsCollection(df)
}

// FirstElementChild returns the first child that is an element; otherwise null
func (df *DocumentFragment) FirstElementChild() *Element {
	for _, child := range df.childNodes {
		if child.NodeType() == ElementNode {
			return child.(*Element)
		}
	}
	return nil
}

// LastElementChild returns the last child that is an element; otherwise null
func (df *DocumentFragment) LastElementChild() *Element {
	for i := len(df.childNodes) - 1; i >= 0; i-- {
		if df.childNodes[i].NodeType() == ElementNode {
			return df.childNodes[i].(*Element)
		}
	}
	return nil
}

// ChildElementCount returns the number of children of this that are elements
func (df *DocumentFragment) ChildElementCount() int {
	count := 0
	for _, child := range df.childNodes {
		if child.NodeType() == ElementNode {
			count++
		}
	}
	return count
}

// Prepend inserts nodes before the first child of this document fragment
func (df *DocumentFragment) Prepend(nodes ...interface{}) error {
	// Convert nodes into a single node or document fragment
	convertedNode, err := convertNodesToNode(nodes, df.ownerDocument)
	if err != nil {
		return err
	}

	if convertedNode == nil {
		return nil // Nothing to insert
	}

	// Pre-insert node before the first child
	if df.FirstChild() != nil {
		df.InsertBefore(convertedNode, df.FirstChild())
	} else {
		df.AppendChild(convertedNode)
	}

	return nil
}

// Append inserts nodes after the last child of this document fragment
func (df *DocumentFragment) Append(nodes ...interface{}) error {
	// Convert nodes into a single node or document fragment
	convertedNode, err := convertNodesToNode(nodes, df.ownerDocument)
	if err != nil {
		return err
	}

	if convertedNode == nil {
		return nil // Nothing to insert
	}

	// Append node to this document fragment
	df.AppendChild(convertedNode)
	return nil
}

// ReplaceChildren replaces all children of this document fragment with the given nodes
func (df *DocumentFragment) ReplaceChildren(nodes ...interface{}) error {
	// Convert nodes into a single node or document fragment
	convertedNode, err := convertNodesToNode(nodes, df.ownerDocument)
	if err != nil {
		return err
	}

	// Ensure pre-insert validity (even though we're not inserting before a specific child)
	if convertedNode != nil {
		if err := df.validateHierarchy(convertedNode); err != nil {
			return err
		}
	}

	// Remove all existing children
	for len(df.childNodes) > 0 {
		df.RemoveChild(df.childNodes[0])
	}

	// Add the new node if any
	if convertedNode != nil {
		df.AppendChild(convertedNode)
	}

	return nil
}

// MoveBefore moves node into this document fragment before child, preserving state
func (df *DocumentFragment) MoveBefore(node Node, child Node) error {
	if node == nil {
		return NewNotFoundError("node cannot be null")
	}

	// Determine reference child
	referenceChild := child
	if referenceChild == node {
		// If reference is the node being moved, use its next sibling
		referenceChild = node.NextSibling()
	}

	// Validate hierarchy constraints
	if err := df.validateHierarchy(node); err != nil {
		return err
	}

	// Move the node (this preserves state by not removing first)
	if referenceChild != nil {
		df.InsertBefore(node, referenceChild)
	} else {
		df.AppendChild(node)
	}

	return nil
}

// QuerySelector returns the first element that is a descendant of this document fragment that matches selectors
func (df *DocumentFragment) QuerySelector(selectors string) *Element {
	// Simple implementation for now - supports basic selectors
	// TODO: Integrate with full CSS selector engine to avoid circular dependency
	var foundElement *Element
	Traverse(df, func(n Node) bool {
		if n == df {
			return true // Skip the root document fragment itself
		}
		if elem, ok := n.(*Element); ok {
			if df.matchesSimpleSelector(elem, selectors) {
				foundElement = elem
				return false // Stop traversal
			}
		}
		return true // Continue traversal
	})
	return foundElement
}

// QuerySelectorAll returns all element descendants of this document fragment that match selectors
func (df *DocumentFragment) QuerySelectorAll(selectors string) NodeList {
	// Simple implementation for now - supports basic selectors
	// TODO: Integrate with full CSS selector engine to avoid circular dependency
	var matchingNodes NodeList
	Traverse(df, func(n Node) bool {
		if n == df {
			return true // Skip the root document fragment itself
		}
		if elem, ok := n.(*Element); ok {
			if df.matchesSimpleSelector(elem, selectors) {
				matchingNodes = append(matchingNodes, elem)
			}
		}
		return true // Continue traversal
	})
	return matchingNodes
}

// matchesSimpleSelector checks if an element matches a basic CSS selector
func (df *DocumentFragment) matchesSimpleSelector(elem *Element, selector string) bool {
	return matchSimpleSelector(elem, selector)
}

// toJS is a placeholder for JavaScript binding.
func (df *DocumentFragment) toJS(vm *goja.Runtime) goja.Value {
	// This will be implemented later when integrating with Goja
	return vm.ToValue(nil)
}
