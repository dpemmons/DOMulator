package dom

import (
	"strings"
	"sync"

	"github.com/dop251/goja"
)

// Document represents the entire HTML document.
type Document struct {
	nodeImpl
	observerRegistry  *ObserverRegistry
	modificationTime  int64 // Counter incremented on each DOM modification
	modificationMutex sync.RWMutex
	// Add document-specific properties here
	// For example, a reference to the window object, if we implement one
	// defaultView *Window
}

// NewDocument creates a new Document node.
func NewDocument() *Document {
	doc := &Document{
		nodeImpl: nodeImpl{
			nodeType: DocumentNode,
			nodeName: "#document",
		},
		observerRegistry: NewObserverRegistry(),
	}
	doc.ownerDocument = doc // A document is its own owner document
	doc.nodeImpl.self = doc // Set the self reference
	return doc
}

// CreateElement creates a new element with the given tag name
func (d *Document) CreateElement(tagName string) *Element {
	return NewElement(tagName, d)
}

// CreateElementNS creates a new element with the given namespace URI and qualified name
func (d *Document) CreateElementNS(namespaceURI, qualifiedName string) (*Element, error) {
	return NewElementNS(namespaceURI, qualifiedName, d)
}

// CreateTextNode creates a new text node with the given data
func (d *Document) CreateTextNode(data string) *Text {
	return NewText(data, d)
}

// CreateComment creates a new comment node with the given data
func (d *Document) CreateComment(data string) *Comment {
	return NewComment(data, d)
}

// CreateDocumentFragment creates a new document fragment
func (d *Document) CreateDocumentFragment() *DocumentFragment {
	return NewDocumentFragment(d)
}

// GetElementById returns the element with the specified ID
func (d *Document) GetElementById(id string) *Element {
	if id == "" {
		return nil // Don't search for empty IDs
	}
	var result *Element
	Traverse(d, func(node Node) bool {
		if elem, ok := node.(*Element); ok {
			if elem.GetAttribute("id") == id {
				result = elem
				return false // Stop traversal
			}
		}
		return true // Continue traversal
	})
	return result
}

// GetElementsByTagName returns a live HTMLCollection of all elements with the specified tag name
func (d *Document) GetElementsByTagName(tagName string) *HTMLCollection {
	return NewElementsByTagNameCollection(d, tagName)
}

// GetElementsByTagNameNS returns a live HTMLCollection of all elements with the specified namespace URI and local name
func (d *Document) GetElementsByTagNameNS(namespaceURI, localName string) *HTMLCollection {
	return NewElementsByTagNameNSCollection(d, namespaceURI, localName)
}

// GetElementsByClassName returns a live HTMLCollection of all elements with the specified class name
func (d *Document) GetElementsByClassName(className string) *HTMLCollection {
	return NewElementsByClassNameCollection(d, className)
}

// GetElementsByName returns a live HTMLCollection of all elements with the specified name attribute
func (d *Document) GetElementsByName(name string) *HTMLCollection {
	return NewElementsByNameCollection(d, name)
}

// DocumentElement returns the document element (usually <html>)
func (d *Document) DocumentElement() *Element {
	for _, child := range d.childNodes {
		if elem, ok := child.(*Element); ok && elem.TagName() == "html" {
			return elem
		}
	}
	return nil
}

// Body returns the body element
func (d *Document) Body() *Element {
	if docElem := d.DocumentElement(); docElem != nil {
		children := docElem.ChildNodes()
		for i := 0; i < children.Length(); i++ {
			child := children.Item(i)
			if elem, ok := child.(*Element); ok && elem.TagName() == "body" {
				return elem
			}
		}
	}
	return nil
}

// Head returns the head element
func (d *Document) Head() *Element {
	if docElem := d.DocumentElement(); docElem != nil {
		children := docElem.ChildNodes()
		for i := 0; i < children.Length(); i++ {
			child := children.Item(i)
			if elem, ok := child.(*Element); ok && elem.TagName() == "head" {
				return elem
			}
		}
	}
	return nil
}

// getObserverRegistry returns the document's observer registry
func (d *Document) getObserverRegistry() *ObserverRegistry {
	return d.observerRegistry
}

// getModificationTime returns the current modification time
func (d *Document) getModificationTime() int64 {
	d.modificationMutex.RLock()
	defer d.modificationMutex.RUnlock()
	return d.modificationTime
}

// incrementModificationTime increments the modification counter
func (d *Document) incrementModificationTime() {
	d.modificationMutex.Lock()
	defer d.modificationMutex.Unlock()
	d.modificationTime++
}

// Children returns a live HTMLCollection of all child elements
func (d *Document) Children() *HTMLCollection {
	return NewChildElementsCollection(d)
}

// FirstElementChild returns the first child that is an element; otherwise null
func (d *Document) FirstElementChild() *Element {
	for _, child := range d.childNodes {
		if child.NodeType() == ElementNode {
			return child.(*Element)
		}
	}
	return nil
}

// LastElementChild returns the last child that is an element; otherwise null
func (d *Document) LastElementChild() *Element {
	for i := len(d.childNodes) - 1; i >= 0; i-- {
		if d.childNodes[i].NodeType() == ElementNode {
			return d.childNodes[i].(*Element)
		}
	}
	return nil
}

// ChildElementCount returns the number of children of this that are elements
func (d *Document) ChildElementCount() int {
	count := 0
	for _, child := range d.childNodes {
		if child.NodeType() == ElementNode {
			count++
		}
	}
	return count
}

// Prepend inserts nodes before the first child of this document
func (d *Document) Prepend(nodes ...interface{}) error {
	// Convert nodes into a single node or document fragment
	convertedNode, err := convertNodesToNode(nodes, d)
	if err != nil {
		return err
	}

	if convertedNode == nil {
		return nil // Nothing to insert
	}

	// Pre-insert node before the first child
	if d.FirstChild() != nil {
		d.InsertBefore(convertedNode, d.FirstChild())
	} else {
		d.AppendChild(convertedNode)
	}

	return nil
}

// Append inserts nodes after the last child of this document
func (d *Document) Append(nodes ...interface{}) error {
	// Convert nodes into a single node or document fragment
	convertedNode, err := convertNodesToNode(nodes, d)
	if err != nil {
		return err
	}

	if convertedNode == nil {
		return nil // Nothing to insert
	}

	// Append node to this document
	d.AppendChild(convertedNode)
	return nil
}

// ReplaceChildren replaces all children of this document with the given nodes
func (d *Document) ReplaceChildren(nodes ...interface{}) error {
	// Convert nodes into a single node or document fragment
	convertedNode, err := convertNodesToNode(nodes, d)
	if err != nil {
		return err
	}

	// Ensure pre-insert validity (even though we're not inserting before a specific child)
	if convertedNode != nil {
		if err := d.validateHierarchy(convertedNode); err != nil {
			return err
		}
	}

	// Remove all existing children
	for len(d.childNodes) > 0 {
		d.RemoveChild(d.childNodes[0])
	}

	// Add the new node if any
	if convertedNode != nil {
		d.AppendChild(convertedNode)
	}

	return nil
}

// MoveBefore moves node into this document before child, preserving state
func (d *Document) MoveBefore(node Node, child Node) error {
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
	if err := d.validateHierarchy(node); err != nil {
		return err
	}

	// Move the node (this preserves state by not removing first)
	if referenceChild != nil {
		d.InsertBefore(node, referenceChild)
	} else {
		d.AppendChild(node)
	}

	return nil
}

// QuerySelector returns the first element that is a descendant of this document that matches selectors
func (d *Document) QuerySelector(selectors string) *Element {
	// Simple implementation for now - supports basic selectors
	// TODO: Integrate with full CSS selector engine to avoid circular dependency
	var foundElement *Element
	Traverse(d, func(n Node) bool {
		if n == d {
			return true // Skip the root document itself
		}
		if elem, ok := n.(*Element); ok {
			if d.matchesSimpleSelector(elem, selectors) {
				foundElement = elem
				return false // Stop traversal
			}
		}
		return true // Continue traversal
	})
	return foundElement
}

// QuerySelectorAll returns all element descendants of this document that match selectors
func (d *Document) QuerySelectorAll(selectors string) *NodeList {
	// Simple implementation for now - supports basic selectors
	// TODO: Integrate with full CSS selector engine to avoid circular dependency
	var matchingNodes []Node
	Traverse(d, func(n Node) bool {
		if n == d {
			return true // Skip the root document itself
		}
		if elem, ok := n.(*Element); ok {
			if d.matchesSimpleSelector(elem, selectors) {
				matchingNodes = append(matchingNodes, elem)
			}
		}
		return true // Continue traversal
	})
	return NewNodeList(matchingNodes)
}

// matchesSimpleSelector checks if an element matches a basic CSS selector
func (d *Document) matchesSimpleSelector(elem *Element, selector string) bool {
	selector = strings.TrimSpace(selector)
	if selector == "" {
		return false
	}

	// Handle wildcard selector
	if selector == "*" {
		return true
	}

	// Parse simple selectors: tag, #id, .class, tag#id, tag.class, etc.
	remaining := selector
	var tagName, id string
	var classes []string

	// Extract ID (e.g., #myid)
	if idx := strings.Index(remaining, "#"); idx != -1 {
		idPart := remaining[idx+1:]
		if spaceIdx := strings.IndexAny(idPart, ". "); spaceIdx != -1 {
			id = idPart[:spaceIdx]
			remaining = remaining[:idx] + idPart[spaceIdx:]
		} else {
			id = idPart
			remaining = remaining[:idx]
		}
	}

	// Extract Classes (e.g., .myclass)
	classParts := strings.Split(remaining, ".")
	if len(classParts) > 1 {
		for _, class := range classParts[1:] {
			if class != "" {
				classes = append(classes, class)
			}
		}
		remaining = classParts[0]
	}

	// Extract Tag (the remaining part)
	tagName = strings.TrimSpace(remaining)

	// Check tag match
	if tagName != "" && !strings.EqualFold(tagName, elem.TagName()) {
		return false
	}

	// Check ID match
	if id != "" && elem.GetAttribute("id") != id {
		return false
	}

	// Check class matches
	if len(classes) > 0 {
		elemClasses := strings.Fields(elem.GetAttribute("class"))
		for _, requiredClass := range classes {
			found := false
			for _, elemClass := range elemClasses {
				if requiredClass == elemClass {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}
	}

	return true
}

// toJS is a placeholder for JavaScript binding.
func (d *Document) toJS(vm *goja.Runtime) goja.Value {
	// This will be implemented later when integrating with Goja
	return vm.ToValue(nil)
}
