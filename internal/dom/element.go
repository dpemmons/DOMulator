package dom

import (
	"errors"
	"strings"

	"github.com/dop251/goja"
)

// Element represents an element node in the DOM tree.
type Element struct {
	nodeImpl

	// Namespace support
	namespaceURI string
	prefix       string
	localName    string
	tagName      string // Computed from prefix and localName

	attributes   *NamedNodeMap     // NamedNodeMap to store attributes per spec
	attributeMap map[string]string // Legacy compatibility map - will be phased out
	classList    *DOMTokenList     // DOMTokenList for class attribute
	// dataset        *Dataset      // To be implemented
	// style          *CSSStyleDeclaration // To be implemented

	// CustomElement support per WHATWG DOM specification
	isValue string // The "is" attribute for customized built-in elements

	// Lazy-loaded properties
	innerHTML   *string
	outerHTML   *string
	textContent *string

	// Event handling
	// eventListeners map[EventPhase]map[string][]*EventListener // To be implemented

	// Shadow DOM support
	shadowRoot *ShadowRoot

	// Slottable interface support (Element and Text nodes are slottables)
	slotName     string // For slot attribute
	assignedSlot *Slot  // Current assigned slot
	manualSlot   *Slot  // Manual slot assignment

	// Form elements
	value    string
	checked  bool
	selected bool

	// Mutation tracking
	// mutationRecord *MutationRecord // To be implemented
}

// NewElement creates a new Element node without namespace.
func NewElement(tagName string, doc *Document) *Element {
	// Parse qualified name for legacy compatibility
	info := ParseQualifiedName(tagName)

	// NewElement does not automatically assign namespaces - use NewElementNS for that
	namespaceURI := info.NamespaceURI

	// According to DOM spec, Element NodeName (and TagName) should be "Its HTML-uppercased qualified name"
	// for HTML elements in an HTML document.
	nodeNameValue := tagName         // tagName is lowercase from parser
	localNameValue := info.LocalName // Should be the lowercase name

	if namespaceURI == htmlNamespace || (namespaceURI == "" && isHTMLElement(localNameValue)) {
		nodeNameValue = strings.ToUpper(tagName)
	}

	elem := &Element{
		nodeImpl: nodeImpl{
			nodeType:      ElementNode,
			nodeName:      nodeNameValue, // Uppercase for HTML elements
			ownerDocument: doc,
		},
		namespaceURI: namespaceURI,
		prefix:       info.Prefix,
		localName:    localNameValue,          // Store the original (lowercase) local name
		tagName:      nodeNameValue,           // tagName field should mirror nodeName for elements
		attributeMap: make(map[string]string), // Legacy compatibility map
	}
	elem.attributes = NewNamedNodeMap(elem) // Initialize the NamedNodeMap
	elem.nodeImpl.self = elem               // Set the self reference
	return elem
}

// isHTMLElement checks if a tag name is a known HTML element
func isHTMLElement(tagName string) bool {
	htmlElements := map[string]bool{
		"a": true, "abbr": true, "address": true, "area": true, "article": true, "aside": true, "audio": true,
		"b": true, "base": true, "bdi": true, "bdo": true, "blockquote": true, "body": true, "br": true, "button": true,
		"canvas": true, "caption": true, "cite": true, "code": true, "col": true, "colgroup": true,
		"data": true, "datalist": true, "dd": true, "del": true, "details": true, "dfn": true, "dialog": true, "div": true, "dl": true, "dt": true,
		"em": true, "embed": true,
		"fieldset": true, "figcaption": true, "figure": true, "footer": true, "form": true,
		"h1": true, "h2": true, "h3": true, "h4": true, "h5": true, "h6": true, "head": true, "header": true, "hgroup": true, "hr": true, "html": true,
		"i": true, "iframe": true, "img": true, "input": true, "ins": true,
		"kbd":   true,
		"label": true, "legend": true, "li": true, "link": true,
		"main": true, "map": true, "mark": true, "menu": true, "meta": true, "meter": true,
		"nav": true, "noscript": true,
		"object": true, "ol": true, "optgroup": true, "option": true, "output": true,
		"p": true, "picture": true, "pre": true, "progress": true,
		"q":  true,
		"rp": true, "rt": true, "ruby": true,
		"s": true, "samp": true, "script": true, "section": true, "select": true, "slot": true, "small": true, "source": true, "span": true, "strong": true, "style": true, "sub": true, "summary": true, "sup": true,
		"table": true, "tbody": true, "td": true, "template": true, "textarea": true, "tfoot": true, "th": true, "thead": true, "time": true, "title": true, "tr": true, "track": true,
		"u": true, "ul": true,
		"var": true, "video": true,
		"wbr": true,
	}
	return htmlElements[strings.ToLower(tagName)]
}

// NewElementNS creates a new Element node with namespace support.
func NewElementNS(namespaceURI, qualifiedName string, doc *Document) (*Element, error) {
	// Validate and extract namespace information
	ns, prefix, localName, err := ValidateAndExtract(namespaceURI, qualifiedName)
	if err != nil {
		return nil, err
	}

	elem := &Element{
		nodeImpl: nodeImpl{
			nodeType:      ElementNode,
			nodeName:      qualifiedName, // NodeName for Element is its qualified name
			ownerDocument: doc,
		},
		namespaceURI: ns,
		prefix:       prefix,
		localName:    localName,
		tagName:      qualifiedName,
		attributeMap: make(map[string]string), // Legacy compatibility map
	}
	elem.attributes = NewNamedNodeMap(elem) // Initialize the NamedNodeMap
	elem.nodeImpl.self = elem               // Set the self reference
	return elem, nil
}

// AppendChild overrides the nodeImpl version to handle cache invalidation
func (e *Element) AppendChild(child Node) Node {
	result := e.nodeImpl.AppendChild(child)
	if result != nil {
		e.invalidateCaches()
	}
	return result
}

// RemoveChild overrides the nodeImpl version to handle cache invalidation
func (e *Element) RemoveChild(child Node) Node {
	result := e.nodeImpl.RemoveChild(child)
	if result != nil {
		e.invalidateCaches()
	}
	return result
}

// InsertBefore overrides the nodeImpl version to handle cache invalidation
func (e *Element) InsertBefore(newChild, refChild Node) Node {
	result := e.nodeImpl.InsertBefore(newChild, refChild)
	if result != nil {
		e.invalidateCaches()
	}
	return result
}

// ReplaceChild overrides the nodeImpl version to handle cache invalidation
func (e *Element) ReplaceChild(newChild, oldChild Node) Node {
	result := e.nodeImpl.ReplaceChild(newChild, oldChild)
	if result != nil {
		e.invalidateCaches()
	}
	return result
}

// invalidateCaches clears cached content when the element's children change
func (e *Element) invalidateCaches() {
	e.innerHTML = nil
	e.outerHTML = nil
	e.textContent = nil
}

// TagName returns the tag name of the element.
// For HTML elements in an HTML document, this is an uppercase name.
func (e *Element) TagName() string {
	return e.tagName // This field is now set to the (potentially uppercased) nodeName
}

// NamespaceURI returns the namespace URI of the element.
func (e *Element) NamespaceURI() string {
	return e.namespaceURI
}

// Prefix returns the namespace prefix of the element.
func (e *Element) Prefix() string {
	return e.prefix
}

// LocalName returns the local name of the element.
func (e *Element) LocalName() string {
	return e.localName
}

// IsValue returns the "is" attribute value for customized built-in elements.
func (e *Element) IsValue() string {
	return e.isValue
}

// SetIsValue sets the "is" attribute value for customized built-in elements.
func (e *Element) SetIsValue(is string) {
	e.isValue = is
}

// SetAttribute sets the value of an attribute on the specified element.
// If the attribute already exists, its value is updated; otherwise, a new attribute is added.
func (e *Element) SetAttribute(name, value string) {
	attr := NewAttr(name, value, e.ownerDocument)
	e.attributes.SetNamedItem(attr)

	// Notify DOMTokenList if class attribute changed
	if name == "class" && e.classList != nil {
		e.classList.invalidate()
	}
}

// GetAttribute returns the value of the named attribute on the specified element.
// If the named attribute does not exist, the value will be an empty string.
func (e *Element) GetAttribute(name string) string {
	attr := e.attributes.GetNamedItem(name)
	if attr != nil {
		return attr.Value()
	}
	return ""
}

// HasAttribute returns true if the element has the specified attribute
func (e *Element) HasAttribute(name string) bool {
	return e.attributes.Contains(name)
}

// RemoveAttribute removes the named attribute from the element
func (e *Element) RemoveAttribute(name string) {
	// According to the DOM spec, removeAttribute should be silent if the attribute doesn't exist
	defer func() {
		if r := recover(); r != nil {
			// Silently ignore NotFoundError as per spec
		}
	}()
	e.attributes.RemoveNamedItem(name)

	// Notify DOMTokenList if class attribute changed
	if name == "class" && e.classList != nil {
		e.classList.invalidate()
	}
}

// GetAttributeNS returns the value of the attribute with the specified namespace and local name.
func (e *Element) GetAttributeNS(namespaceURI, localName string) string {
	attr := e.attributes.GetNamedItemNS(namespaceURI, localName)
	if attr != nil {
		return attr.Value()
	}
	return ""
}

// SetAttributeNS sets the value of the attribute with the specified namespace and qualified name.
func (e *Element) SetAttributeNS(namespaceURI, qualifiedName, value string) error {
	// Create namespace-aware attribute
	attr, err := NewAttrNS(namespaceURI, qualifiedName, value, e.ownerDocument)
	if err != nil {
		return err
	}

	e.attributes.SetNamedItemNS(attr)
	return nil
}

// HasAttributeNS returns true if the element has an attribute with the specified namespace and local name.
func (e *Element) HasAttributeNS(namespaceURI, localName string) bool {
	return e.attributes.ContainsNS(namespaceURI, localName)
}

// RemoveAttributeNS removes the attribute with the specified namespace and local name.
func (e *Element) RemoveAttributeNS(namespaceURI, localName string) {
	// According to the DOM spec, removeAttributeNS should be silent if the attribute doesn't exist
	defer func() {
		if r := recover(); r != nil {
			// Silently ignore NotFoundError as per spec
		}
	}()
	e.attributes.RemoveNamedItemNS(namespaceURI, localName)
}

// ClassList returns the DOMTokenList for the class attribute
func (e *Element) ClassList() *DOMTokenList {
	if e.classList == nil {
		e.classList = NewDOMTokenList(e, "class")
	}
	return e.classList
}

// HasClass returns true if the element has the specified class
func (e *Element) HasClass(className string) bool {
	return e.ClassList().Contains(className)
}

// GetElementsByTagName returns a live HTMLCollection of all descendant elements with the specified tag name
func (e *Element) GetElementsByTagName(tagName string) *HTMLCollection {
	return NewElementsByTagNameCollection(e, tagName)
}

// GetElementsByClassName returns a live HTMLCollection of all descendant elements with the specified class name
func (e *Element) GetElementsByClassName(className string) *HTMLCollection {
	return NewElementsByClassNameCollection(e, className)
}

// GetElementsByTagNameNS returns a live HTMLCollection of all descendant elements with the specified namespace URI and local name
func (e *Element) GetElementsByTagNameNS(namespaceURI, localName string) *HTMLCollection {
	return NewElementsByTagNameNSCollection(e, namespaceURI, localName)
}

// Children returns a live HTMLCollection of all child elements
func (e *Element) Children() *HTMLCollection {
	return NewChildElementsCollection(e)
}

// FirstElementChild returns the first child that is an element; otherwise null
func (e *Element) FirstElementChild() *Element {
	for _, child := range e.childNodes {
		if child.NodeType() == ElementNode {
			return child.(*Element)
		}
	}
	return nil
}

// LastElementChild returns the last child that is an element; otherwise null
func (e *Element) LastElementChild() *Element {
	for i := len(e.childNodes) - 1; i >= 0; i-- {
		if e.childNodes[i].NodeType() == ElementNode {
			return e.childNodes[i].(*Element)
		}
	}
	return nil
}

// ChildElementCount returns the number of children of this that are elements
func (e *Element) ChildElementCount() int {
	count := 0
	for _, child := range e.childNodes {
		if child.NodeType() == ElementNode {
			count++
		}
	}
	return count
}

// Prepend inserts nodes before the first child of this element
func (e *Element) Prepend(nodes ...interface{}) error {
	// Convert nodes into a single node or document fragment
	convertedNode, err := convertNodesToNode(nodes, e.ownerDocument)
	if err != nil {
		return err
	}

	if convertedNode == nil {
		return nil // Nothing to insert
	}

	// Pre-insert node before the first child
	if e.FirstChild() != nil {
		e.InsertBefore(convertedNode, e.FirstChild())
	} else {
		e.AppendChild(convertedNode)
	}

	return nil
}

// Append inserts nodes after the last child of this element
func (e *Element) Append(nodes ...interface{}) error {
	// Convert nodes into a single node or document fragment
	convertedNode, err := convertNodesToNode(nodes, e.ownerDocument)
	if err != nil {
		return err
	}

	if convertedNode == nil {
		return nil // Nothing to insert
	}

	// Append node to this element
	e.AppendChild(convertedNode)
	return nil
}

// ReplaceChildren replaces all children of this element with the given nodes
func (e *Element) ReplaceChildren(nodes ...interface{}) error {
	// Convert nodes into a single node or document fragment
	convertedNode, err := convertNodesToNode(nodes, e.ownerDocument)
	if err != nil {
		return err
	}

	// Ensure pre-insert validity (even though we're not inserting before a specific child)
	if convertedNode != nil {
		if err := e.validateHierarchy(convertedNode); err != nil {
			return err
		}
	}

	// Remove all existing children
	for len(e.childNodes) > 0 {
		e.RemoveChild(e.childNodes[0])
	}

	// Add the new node if any
	if convertedNode != nil {
		e.AppendChild(convertedNode)
	}

	return nil
}

// MoveBefore moves node into this element before child, preserving state
func (e *Element) MoveBefore(node Node, child Node) error {
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
	if err := e.validateHierarchy(node); err != nil {
		return err
	}

	// Move the node (this preserves state by not removing first)
	// Note: InsertBefore handles moving nodes within the same parent correctly
	if referenceChild != nil {
		e.InsertBefore(node, referenceChild)
	} else {
		e.AppendChild(node)
	}

	return nil
}

// QuerySelector returns the first element that is a descendant of this element that matches selectors
func (e *Element) QuerySelector(selectors string) *Element {
	// Simple implementation for now - supports basic selectors
	// TODO: Integrate with full CSS selector engine to avoid circular dependency
	return e.querySelectorRecursive(selectors, false)
}

// querySelectorRecursive performs depth-first search for elements matching the selector
func (e *Element) querySelectorRecursive(selectors string, includeRoot bool) *Element {
	// Check if this element matches (only if we should include the root)
	if includeRoot && matchSimpleSelector(e, selectors) {
		return e
	}

	// Search through all direct children in order
	children := e.ChildNodes()
	for i := 0; i < children.Length(); i++ {
		child := children.Item(i)
		if elem, ok := child.(*Element); ok {
			// Check if this child element matches
			if matchSimpleSelector(elem, selectors) {
				return elem
			}
			// Recursively search in this element's descendants
			if found := elem.querySelectorRecursive(selectors, true); found != nil {
				return found
			}
		}
	}
	return nil
}

// QuerySelectorAll returns all element descendants of this element that match selectors
func (e *Element) QuerySelectorAll(selectors string) *NodeList {
	// Simple implementation for now - supports basic selectors
	// TODO: Integrate with full CSS selector engine to avoid circular dependency
	var matchingNodes []Node
	Traverse(e, func(n Node) bool {
		if n == e {
			return true // Skip the root element itself
		}
		if elem, ok := n.(*Element); ok {
			if matchSimpleSelector(elem, selectors) {
				matchingNodes = append(matchingNodes, elem)
			}
		}
		return true // Continue traversal
	})
	return NewNodeList(matchingNodes)
}

// findElementByIdRecursive performs a depth-first search for an element with the given ID.
// It properly stops traversal when the first match is found.
func (e *Element) findElementByIdRecursive(id string) *Element {
	// Search through all direct children
	children := e.ChildNodes()
	for i := 0; i < children.Length(); i++ {
		child := children.Item(i)
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

// InnerHTML returns the HTML content of the element.
func (e *Element) InnerHTML() string {
	if e.innerHTML == nil {
		// TODO: Implement actual HTML serialization
		// For now, a placeholder
		content := ""
		children := e.ChildNodes()
		for i := 0; i < children.Length(); i++ {
			child := children.Item(i)
			// This is a very basic placeholder.
			// A proper implementation would recursively serialize children.
			if child.NodeType() == TextNode {
				content += child.NodeValue()
			} else if child.NodeType() == ElementNode {
				content += child.(*Element).OuterHTML()
			}
		}
		e.innerHTML = &content
	}
	return *e.innerHTML
}

// SetInnerHTML sets the HTML content of the element.
func (e *Element) SetInnerHTML(html string) {
	// TODO: Implement HTML parsing and DOM manipulation
	// For now, clear existing children and set a text node
	e.childNodes = nil
	textNode := NewText(html, e.ownerDocument)
	e.AppendChild(textNode)
	e.innerHTML = &html // Invalidate cached innerHTML
	e.outerHTML = nil   // Invalidate cached outerHTML
	e.textContent = nil // Invalidate cached textContent
}

// OuterHTML returns the HTML content of the element including itself.
func (e *Element) OuterHTML() string {
	if e.outerHTML == nil {
		// TODO: Implement actual HTML serialization
		// For now, a placeholder
		content := "<" + e.tagName + ">" + e.InnerHTML() + "</" + e.tagName + ">"
		e.outerHTML = &content
	}
	return *e.outerHTML
}

// TextContent returns the text content of the element.
func (e *Element) TextContent() string {
	if e.textContent == nil {
		// TODO: Implement actual text content extraction
		// For now, a placeholder
		content := ""
		for _, child := range e.childNodes {
			if child.NodeType() == TextNode {
				content += child.NodeValue()
			} else if child.NodeType() == ElementNode {
				content += child.(*Element).TextContent()
			}
		}
		e.textContent = &content
	}
	return *e.textContent
}

// SetTextContent sets the text content of the element using the DOM specification algorithm.
func (e *Element) SetTextContent(text string) {
	// Use the spec-compliant implementation from nodeImpl
	e.nodeImpl.SetTextContent(text)

	// Invalidate cached properties
	e.innerHTML = nil
	e.outerHTML = nil
	e.textContent = &text
}

// splitBySpace splits a string by whitespace characters
func splitBySpace(s string) []string {
	var result []string
	current := ""
	for _, r := range s {
		if r == ' ' || r == '\t' || r == '\n' || r == '\r' {
			if current != "" {
				result = append(result, current)
				current = ""
			}
		} else {
			current += string(r)
		}
	}
	if current != "" {
		result = append(result, current)
	}
	return result
}

// CloneNode creates a copy of the element using the spec-compliant cloning implementation.
func (e *Element) CloneNode(deep bool) Node {
	// Use the spec-compliant cloning implementation
	return CloneNodeSpec(e, deep)
}

// getEventListeners returns the event listeners map for use in event dispatching
func (e *Element) getEventListeners() map[string][]func(Event) {
	return e.nodeImpl.getEventListeners()
}

// InsertAdjacentHTML parses the specified text as HTML and inserts the resulting nodes
// into the DOM tree at a position relative to the element on which it was invoked.
func (e *Element) InsertAdjacentHTML(position, html string) error {
	// Validate position
	position = strings.ToLower(position)
	if position != "beforebegin" && position != "afterbegin" && position != "beforeend" && position != "afterend" {
		return errors.New("invalid position: must be 'beforebegin', 'afterbegin', 'beforeend', or 'afterend'")
	}

	// Parse HTML fragment - for now, use a simple text node approach
	// TODO: Implement proper HTML parsing when parser circular dependency is resolved
	parsedNodes := e.parseHTMLFragment(html)

	switch position {
	case "beforebegin":
		// Insert before the element (as siblings)
		parent := e.ParentNode()
		if parent == nil {
			return errors.New("cannot insert beforebegin: element has no parent")
		}
		for i, node := range parsedNodes {
			if i == 0 {
				parent.InsertBefore(node, e)
			} else {
				// Insert subsequent nodes after the previously inserted node
				prevNode := parsedNodes[i-1]
				var nextSibling Node
				parentChildren := parent.ChildNodes()

				// Find the next sibling after the previously inserted node
				for j := 0; j < parentChildren.Length(); j++ {
					child := parentChildren.Item(j)
					if child == prevNode && j+1 < parentChildren.Length() {
						nextSibling = parentChildren.Item(j + 1)
						break
					}
				}

				if nextSibling != nil {
					parent.InsertBefore(node, nextSibling)
				} else {
					parent.AppendChild(node)
				}
			}
		}

	case "afterbegin":
		// Insert as first children of the element
		for i := len(parsedNodes) - 1; i >= 0; i-- {
			if len(e.childNodes) > 0 {
				e.InsertBefore(parsedNodes[i], e.childNodes[0])
			} else {
				e.AppendChild(parsedNodes[i])
			}
		}

	case "beforeend":
		// Insert as last children of the element
		for _, node := range parsedNodes {
			e.AppendChild(node)
		}

	case "afterend":
		// Insert after the element (as siblings)
		parent := e.ParentNode()
		if parent == nil {
			return errors.New("cannot insert afterend: element has no parent")
		}

		// Find the next sibling
		var nextSibling Node
		found := false
		parentChildren := parent.ChildNodes()
		for i := 0; i < parentChildren.Length(); i++ {
			child := parentChildren.Item(i)
			if found {
				nextSibling = child
				break
			}
			if child == e {
				found = true
			}
		}

		for _, node := range parsedNodes {
			if nextSibling != nil {
				parent.InsertBefore(node, nextSibling)
			} else {
				parent.AppendChild(node)
			}
		}
	}

	// Invalidate caches
	e.invalidateCaches()

	return nil
}

// parseHTMLFragment is a simple HTML parser for basic fragments
// TODO: Replace with proper HTML parsing when circular dependency is resolved
func (e *Element) parseHTMLFragment(html string) []Node {
	// For now, implement a very basic parser for simple cases
	// This handles basic elements and text content
	html = strings.TrimSpace(html)
	if html == "" {
		return []Node{}
	}

	// Check if it's a simple text node (no < or > characters)
	if !strings.Contains(html, "<") {
		return []Node{NewText(html, e.ownerDocument)}
	}

	// Basic element parsing - handle simple cases like <div>content</div>
	var nodes []Node

	// Very simple parsing for demonstration - handles basic elements
	if strings.HasPrefix(html, "<") && strings.HasSuffix(html, ">") {
		// Check for self-closing tag
		if strings.HasSuffix(html, "/>") {
			tagContent := html[1 : len(html)-2]
			tagParts := strings.Fields(tagContent)
			if len(tagParts) > 0 {
				element := NewElement(tagParts[0], e.ownerDocument)
				// Parse basic attributes
				for i := 1; i < len(tagParts); i++ {
					if strings.Contains(tagParts[i], "=") {
						attrParts := strings.SplitN(tagParts[i], "=", 2)
						if len(attrParts) == 2 {
							key := attrParts[0]
							value := strings.Trim(attrParts[1], `"'`)
							element.SetAttribute(key, value)
						}
					}
				}
				nodes = append(nodes, element)
			}
		} else {
			// Try to parse opening and closing tags
			openEnd := strings.Index(html, ">")
			if openEnd > 0 {
				openTag := html[1:openEnd]
				tagParts := strings.Fields(openTag)
				if len(tagParts) > 0 {
					tagName := tagParts[0]
					closeTag := "</" + tagName + ">"
					closeStart := strings.LastIndex(html, closeTag)

					if closeStart > openEnd {
						element := NewElement(tagName, e.ownerDocument)

						// Parse basic attributes
						for i := 1; i < len(tagParts); i++ {
							if strings.Contains(tagParts[i], "=") {
								attrParts := strings.SplitN(tagParts[i], "=", 2)
								if len(attrParts) == 2 {
									key := attrParts[0]
									value := strings.Trim(attrParts[1], `"'`)
									element.SetAttribute(key, value)
								}
							}
						}

						// Get content between tags
						content := html[openEnd+1 : closeStart]
						if strings.TrimSpace(content) != "" {
							// Recursively parse content
							childNodes := e.parseHTMLFragment(content)
							for _, child := range childNodes {
								element.AppendChild(child)
							}
						}

						nodes = append(nodes, element)
					}
				}
			}
		}
	} else {
		// Fallback: treat as text
		nodes = append(nodes, NewText(html, e.ownerDocument))
	}

	if len(nodes) == 0 {
		// Fallback: create text node
		nodes = append(nodes, NewText(html, e.ownerDocument))
	}

	return nodes
}

// toJS is a placeholder for JavaScript binding.
func (e *Element) toJS(vm *goja.Runtime) goja.Value {
	// This will be implemented later when integrating with Goja
	return vm.ToValue(nil)
}

// ChildNode mixin methods

// Before inserts nodes just before this element in its parent's children list.
func (e *Element) Before(nodes ...interface{}) error {
	parent := e.ParentNode()
	if parent == nil {
		return nil // No parent, nothing to do
	}

	// Convert nodes to a single node or document fragment
	convertedNode, err := convertNodesToNode(nodes, e.ownerDocument)
	if err != nil {
		return err
	}

	if convertedNode == nil {
		return nil // Nothing to insert
	}

	// Insert before this element
	parent.InsertBefore(convertedNode, e)
	return nil
}

// After inserts nodes just after this element in its parent's children list.
func (e *Element) After(nodes ...interface{}) error {
	parent := e.ParentNode()
	if parent == nil {
		return nil // No parent, nothing to do
	}

	// Convert nodes to a single node or document fragment
	convertedNode, err := convertNodesToNode(nodes, e.ownerDocument)
	if err != nil {
		return err
	}

	if convertedNode == nil {
		return nil // Nothing to insert
	}

	// Find the next viable sibling (not in the nodes being inserted)
	var excludeNodes []Node
	if fragment, ok := convertedNode.(*DocumentFragment); ok {
		fragmentChildren := fragment.ChildNodes()
		for i := 0; i < fragmentChildren.Length(); i++ {
			excludeNodes = append(excludeNodes, fragmentChildren.Item(i))
		}
	} else {
		excludeNodes = []Node{convertedNode}
	}

	nextSibling := findViableNextSibling(e, excludeNodes)
	if nextSibling != nil {
		parent.InsertBefore(convertedNode, nextSibling)
	} else {
		parent.AppendChild(convertedNode)
	}
	return nil
}

// ReplaceWith replaces this element with the given nodes.
func (e *Element) ReplaceWith(nodes ...interface{}) error {
	parent := e.ParentNode()
	if parent == nil {
		return nil // No parent, nothing to do
	}

	// Convert nodes to a single node or document fragment
	convertedNode, err := convertNodesToNode(nodes, e.ownerDocument)
	if err != nil {
		return err
	}

	if convertedNode == nil {
		// Just remove this element if no replacement nodes
		parent.RemoveChild(e)
		return nil
	}

	// Replace this element with the converted node
	parent.ReplaceChild(convertedNode, e)
	return nil
}

// Remove removes this element from its parent's children list.
func (e *Element) Remove() error {
	parent := e.ParentNode()
	if parent == nil {
		return nil // No parent, nothing to do
	}

	parent.RemoveChild(e)
	return nil
}

// NonDocumentTypeChildNode mixin methods

// PreviousElementSibling returns the first preceding sibling that is an element; otherwise null.
func (e *Element) PreviousElementSibling() *Element {
	parent := e.ParentNode()
	if parent == nil {
		return nil
	}

	siblings := parent.ChildNodes()
	for i := siblings.Length() - 1; i >= 0; i-- {
		if siblings.Item(i) == e {
			// Found this element, now look backwards for an element sibling
			for j := i - 1; j >= 0; j-- {
				sibling := siblings.Item(j)
				if sibling.NodeType() == ElementNode {
					return sibling.(*Element)
				}
			}
			break
		}
	}
	return nil
}

// NextElementSibling returns the first following sibling that is an element; otherwise null.
func (e *Element) NextElementSibling() *Element {
	parent := e.ParentNode()
	if parent == nil {
		return nil
	}

	siblings := parent.ChildNodes()
	for i := 0; i < siblings.Length(); i++ {
		sibling := siblings.Item(i)
		if sibling == e {
			// Found this element, now look forwards for an element sibling
			for j := i + 1; j < siblings.Length(); j++ {
				nextSibling := siblings.Item(j)
				if nextSibling.NodeType() == ElementNode {
					return nextSibling.(*Element)
				}
			}
			break
		}
	}
	return nil
}

// Shadow DOM Methods

// ShadowRoot returns the shadow root attached to this element, or nil if none exists
func (e *Element) ShadowRoot() *ShadowRoot {
	return e.shadowRoot
}

// AttachShadow creates and attaches a shadow root to this element
func (e *Element) AttachShadow(init ShadowRootInit) (*ShadowRoot, error) {
	// Check if element already has a shadow root
	if e.shadowRoot != nil {
		return nil, NewInvalidStateError("Element already has a shadow root")
	}

	// Validate that this element can have a shadow root
	if !e.canAttachShadowRoot() {
		return nil, NewNotSupportedError("This element type cannot have a shadow root")
	}

	// Create the shadow root
	shadowRoot := NewShadowRoot(e, init.Mode)
	if init.SlotAssignment != "" {
		shadowRoot.SetSlotAssignment(init.SlotAssignment)
	}

	// Attach it to this element
	e.shadowRoot = shadowRoot

	return shadowRoot, nil
}

// canAttachShadowRoot checks if this element type can have a shadow root
func (e *Element) canAttachShadowRoot() bool {
	// Per HTML spec, only certain elements can have shadow roots
	switch strings.ToLower(e.localName) {
	case "article", "aside", "blockquote", "body", "div", "footer", "h1", "h2", "h3", "h4", "h5", "h6",
		"header", "main", "nav", "p", "section", "span":
		return true
	default:
		return false
	}
}

// Slottable interface implementation

// GetSlotName returns the value of the slot attribute
func (e *Element) GetSlotName() string {
	slotAttr := e.GetAttribute("slot")
	if slotAttr != "" {
		return slotAttr
	}
	return e.slotName
}

// SetSlotName sets the slot attribute value
func (e *Element) SetSlotName(name string) {
	e.slotName = name
	if name == "" {
		e.RemoveAttribute("slot")
	} else {
		e.SetAttribute("slot", name)
	}

	// Trigger slot assignment update
	if shadowRoot := e.getShadowRootFromParent(); shadowRoot != nil {
		shadowRoot.AssignSlot(e)
	}
}

// GetAssignedSlot returns the slot this element is assigned to
func (e *Element) GetAssignedSlot() *Slot {
	return e.assignedSlot
}

// AssignedSlot implements the Slottable mixin assignedSlot getter per WHATWG DOM Section 4.2.9
// The assignedSlot getter steps are to return the result of find a slot given this and true.
func (e *Element) AssignedSlot() *Slot {
	return findSlotForSlottable(e, true)
}

// findSlotForSlottable implements the "find a slot" algorithm for a slottable node
// This is used by both Element and Text nodes for their AssignedSlot() method
func findSlotForSlottable(slottable Slottable, open bool) *Slot {
	if slottable == nil {
		return nil
	}

	parent := slottable.ParentNode()
	if parent == nil {
		return nil
	}

	// Get the parent's shadow root
	var shadow *ShadowRoot
	if elem, ok := parent.(*Element); ok {
		shadow = elem.ShadowRoot()
	}

	if shadow == nil {
		return nil
	}

	// Use the ShadowRoot's FindSlot method which implements the full algorithm
	return shadow.FindSlot(slottable, open)
}

// SetAssignedSlot sets the slot this element is assigned to
func (e *Element) SetAssignedSlot(slot *Slot) {
	e.assignedSlot = slot
}

// GetManualSlot returns the manual slot assignment
func (e *Element) GetManualSlot() *Slot {
	return e.manualSlot
}

// SetManualSlot sets the manual slot assignment
func (e *Element) SetManualSlot(slot *Slot) {
	e.manualSlot = slot
}

// getShadowRootFromParent finds the shadow root that contains this element
func (e *Element) getShadowRootFromParent() *ShadowRoot {
	parent := e.ParentNode()
	if parent == nil {
		return nil
	}

	if shadowRoot, ok := parent.(*ShadowRoot); ok {
		return shadowRoot
	}

	if elem, ok := parent.(*Element); ok {
		return elem.ShadowRoot()
	}

	return nil
}

// GetRootNode returns the root node, optionally crossing shadow boundaries
func (e *Element) GetRootNode(options *GetRootNodeOptions) Node {
	if options != nil && options.Composed {
		// For shadow-including root, traverse up through shadow boundaries
		current := Node(e)
		for {
			parent := current.ParentNode()
			if parent != nil {
				current = parent
				continue
			}

			// Check if current is a shadow root
			if shadowRoot, ok := current.(*ShadowRoot); ok {
				if shadowRoot.Host() != nil {
					current = shadowRoot.Host()
					continue
				}
			}

			break
		}
		return current
	}

	// Normal root (don't cross shadow boundaries)
	return e.nodeImpl.GetRootNode(options)
}
