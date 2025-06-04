package dom

import (
	"strings"
	"sync"
)

// HTMLCollection represents a live collection of DOM elements
type HTMLCollection struct {
	// Filter function to determine which nodes to include
	filter func(Node) bool

	// Root node to search from
	root Node

	// Cached results and validity
	cache     []*Element
	cacheTime int64 // DOM modification counter when cache was built
	mutex     sync.RWMutex

	// Document reference for change detection
	document *Document
}

// NewHTMLCollection creates a new HTMLCollection with the given filter
func NewHTMLCollection(root Node, filter func(Node) bool) *HTMLCollection {
	var doc *Document
	if root != nil {
		doc = root.OwnerDocument()
		if doc == nil {
			// If root is a document, use it directly
			if d, ok := root.(*Document); ok {
				doc = d
			}
		}
	}

	return &HTMLCollection{
		filter:   filter,
		root:     root,
		document: doc,
	}
}

// Length returns the number of elements in the collection
func (hc *HTMLCollection) Length() int {
	hc.ensureCache()
	hc.mutex.RLock()
	defer hc.mutex.RUnlock()
	return len(hc.cache)
}

// Item returns the element at the specified index, or nil if out of bounds
func (hc *HTMLCollection) Item(index int) *Element {
	hc.ensureCache()
	hc.mutex.RLock()
	defer hc.mutex.RUnlock()

	if index < 0 || index >= len(hc.cache) {
		return nil
	}
	return hc.cache[index]
}

// NamedItem returns the first element with the specified id or name, or nil if not found
func (hc *HTMLCollection) NamedItem(name string) *Element {
	if name == "" {
		return nil
	}

	hc.ensureCache()
	hc.mutex.RLock()
	defer hc.mutex.RUnlock()

	// Search through elements in tree order and check both ID and name for each
	// According to spec: "Return the first element in the collection for which at least one of the following is true"
	for _, elem := range hc.cache {
		// Check if it has an ID which is name
		if elem.GetAttribute("id") == name {
			return elem
		}

		// Check if it is in the HTML namespace OR is a known HTML element (for legacy compatibility)
		// and has a name attribute whose value is name
		isHTMLElement := elem.NamespaceURI() == htmlNamespace ||
			(elem.NamespaceURI() == "" && isHTMLElementName(elem.LocalName()))

		if isHTMLElement && elem.GetAttribute("name") == name {
			return elem
		}
	}

	return nil
}

// ToSlice returns a slice copy of all elements in the collection
func (hc *HTMLCollection) ToSlice() []*Element {
	hc.ensureCache()
	hc.mutex.RLock()
	defer hc.mutex.RUnlock()

	result := make([]*Element, len(hc.cache))
	copy(result, hc.cache)
	return result
}

// ensureCache builds or updates the cache if necessary
func (hc *HTMLCollection) ensureCache() {
	if hc.document == nil {
		return
	}

	// Get current DOM modification time
	currentTime := hc.document.getModificationTime()

	hc.mutex.RLock()
	cacheValid := hc.cacheTime == currentTime
	hc.mutex.RUnlock()

	if cacheValid {
		return // Cache is still valid
	}

	// Rebuild cache
	hc.mutex.Lock()
	defer hc.mutex.Unlock()

	// Double-check after acquiring write lock
	if hc.cacheTime == currentTime {
		return
	}

	hc.cache = hc.cache[:0] // Clear cache but keep capacity
	hc.buildCache(hc.root)
	hc.cacheTime = currentTime
}

// buildCache recursively builds the cache of matching elements
func (hc *HTMLCollection) buildCache(node Node) {
	hc.buildCacheRecursive(node, true)
}

// buildCacheRecursive recursively builds the cache, with option to skip root
func (hc *HTMLCollection) buildCacheRecursive(node Node, isRoot bool) {
	if node == nil {
		return
	}

	// Check if this node matches the filter (skip root element for descendant searches)
	if elem, ok := node.(*Element); ok && hc.filter(node) && !isRoot {
		hc.cache = append(hc.cache, elem)
	}

	// Recursively check children
	children := node.ChildNodes()
	for i := 0; i < children.Length(); i++ {
		child := children.Item(i)
		hc.buildCacheRecursive(child, false)
	}
}

// NewElementsByTagNameCollection creates an HTMLCollection that contains all elements with the specified tag name
func NewElementsByTagNameCollection(root Node, tagName string) *HTMLCollection {
	filter := func(node Node) bool {
		if elem, ok := node.(*Element); ok {
			return elem.TagName() == tagName || tagName == "*"
		}
		return false
	}
	return NewHTMLCollection(root, filter)
}

// NewElementsByTagNameNSCollection creates an HTMLCollection that contains all elements with the specified namespace and local name
func NewElementsByTagNameNSCollection(root Node, namespaceURI, localName string) *HTMLCollection {
	filter := func(node Node) bool {
		if elem, ok := node.(*Element); ok {
			nsMatch := namespaceURI == "*" || elem.NamespaceURI() == namespaceURI
			nameMatch := localName == "*" || elem.LocalName() == localName
			return nsMatch && nameMatch
		}
		return false
	}
	return NewHTMLCollection(root, filter)
}

// NewElementsByClassNameCollection creates an HTMLCollection that contains all elements with the specified class names
func NewElementsByClassNameCollection(root Node, classNames string) *HTMLCollection {
	filter := func(node Node) bool {
		if elem, ok := node.(*Element); ok {
			// Split class names by whitespace and check that element has ALL classes
			return hasAllClasses(elem, classNames)
		}
		return false
	}
	return NewHTMLCollection(root, filter)
}

// hasAllClasses checks if an element has all the classes in the space-separated list
func hasAllClasses(elem *Element, classNames string) bool {
	// Split by whitespace characters
	classes := strings.Fields(classNames)
	if len(classes) == 0 {
		return false
	}

	// Check that the element has ALL of the specified classes
	for _, className := range classes {
		if !elem.HasClass(className) {
			return false
		}
	}
	return true
}

// NewElementsByNameCollection creates an HTMLCollection that contains all elements with the specified name attribute
func NewElementsByNameCollection(root Node, name string) *HTMLCollection {
	filter := func(node Node) bool {
		if elem, ok := node.(*Element); ok {
			return elem.GetAttribute("name") == name
		}
		return false
	}
	return NewHTMLCollection(root, filter)
}

// NewChildElementsCollection creates an HTMLCollection that contains all child elements
func NewChildElementsCollection(parent Node) *HTMLCollection {
	filter := func(node Node) bool {
		// Only include direct children that are elements
		if elem, ok := node.(*Element); ok {
			return elem.ParentNode() == parent
		}
		return false
	}
	return NewHTMLCollection(parent, filter)
}

// isHTMLElementName checks if a tag name is a known HTML element
func isHTMLElementName(tagName string) bool {
	// Known HTML elements - for backwards compatibility, allow name attribute on all HTML elements
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
	return htmlElements[tagName]
}
