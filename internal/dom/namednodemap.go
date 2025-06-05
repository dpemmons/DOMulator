package dom

import (
	"sort"
	"strings"
)

// NamedNodeMap represents a collection of Attr nodes.
// It is defined in WHATWG DOM Section 4.9.1 as part of the Element interface.
type NamedNodeMap struct {
	element *Element
	attrs   []*Attr // Ordered list of attributes
}

// NewNamedNodeMap creates a new NamedNodeMap for the given element.
func NewNamedNodeMap(element *Element) *NamedNodeMap {
	return &NamedNodeMap{
		element: element,
		attrs:   make([]*Attr, 0),
	}
}

// Length returns the number of attributes in the collection.
// Per WHATWG DOM spec: "The length getter steps are to return the attribute list's size."
func (nnm *NamedNodeMap) Length() int {
	return len(nnm.attrs)
}

// Item returns the attribute at the specified index.
// Per WHATWG DOM spec: "The item(index) method steps are:
// 1. If index is equal to or greater than this's attribute list's size, then return null.
// 2. Otherwise, return this's attribute list[index]."
func (nnm *NamedNodeMap) Item(index int) *Attr {
	if index < 0 || index >= len(nnm.attrs) {
		return nil
	}
	return nnm.attrs[index]
}

// GetNamedItem returns the attribute with the specified qualified name.
// Per WHATWG DOM spec: "The getNamedItem(qualifiedName) method steps are to return
// the result of getting an attribute given qualifiedName and element."
func (nnm *NamedNodeMap) GetNamedItem(qualifiedName string) *Attr {
	// Implementation of "get an attribute by name" algorithm from WHATWG DOM spec:
	// If element is in the HTML namespace and its node document is an HTML document,
	// then set qualifiedName to qualifiedName in ASCII lowercase.
	isHTMLInHTML := nnm.element != nil &&
		nnm.element.NamespaceURI() == htmlNamespace &&
		nnm.element.OwnerDocument() != nil &&
		nnm.element.OwnerDocument().documentType == "html"

	if isHTMLInHTML {
		qualifiedName = strings.ToLower(qualifiedName)
	}

	// Return the first attribute in element's attribute list whose qualified name is qualifiedName
	for _, attr := range nnm.attrs {
		attrName := attr.Name()
		if isHTMLInHTML {
			attrName = strings.ToLower(attrName)
		}
		if attrName == qualifiedName {
			return attr
		}
	}
	return nil
}

// GetNamedItemNS returns the attribute with the specified namespace URI and local name.
// Per WHATWG DOM spec: "The getNamedItemNS(namespace, localName) method steps are to return
// the result of getting an attribute given namespace, localName, and element."
func (nnm *NamedNodeMap) GetNamedItemNS(namespace, localName string) *Attr {
	// Implementation of "get an attribute by namespace and local name" algorithm:
	// If namespace is the empty string, then set it to null.
	if namespace == "" {
		namespace = ""
	}

	// Return the attribute in element's attribute list whose namespace is namespace
	// and local name is localName, if any; otherwise null.
	for _, attr := range nnm.attrs {
		if attr.LocalName() == localName && attr.NamespaceURI() == namespace {
			return attr
		}
	}
	return nil
}

// SetNamedItem adds an attribute to the collection.
// Per WHATWG DOM spec: "The setNamedItem(attr) and setNamedItemNS(attr) method steps are to return
// the result of setting an attribute given attr and element."
func (nnm *NamedNodeMap) SetNamedItem(attr *Attr) *Attr {
	if attr == nil {
		return nil
	}

	// Implementation of "set an attribute" algorithm:
	// If attr's element is neither null nor element, throw an "InUseAttributeError" DOMException.
	if attr.OwnerElement() != nil && attr.OwnerElement() != nnm.element {
		panic(NewInvalidStateError("Attribute is already in use by another element"))
	}

	// Let oldAttr be the result of getting an attribute given attr's namespace, attr's local name, and element.
	var oldAttr *Attr
	for i, existingAttr := range nnm.attrs {
		if existingAttr.NamespaceURI() == attr.NamespaceURI() && existingAttr.LocalName() == attr.LocalName() {
			oldAttr = existingAttr
			// If oldAttr is attr, return attr.
			if oldAttr == attr {
				return attr
			}
			// If oldAttr is non-null, then replace oldAttr with attr.
			nnm.attrs[i] = attr
			break
		}
	}

	// Otherwise, append attr to element.
	if oldAttr == nil {
		nnm.attrs = append(nnm.attrs, attr)
	}

	// Set attr's element and node document
	attr.ownerElement = nnm.element
	if nnm.element != nil {
		attr.ownerDocument = nnm.element.OwnerDocument()
	}

	// Clear old attr's element if it exists
	if oldAttr != nil {
		oldAttr.ownerElement = nil
	}

	return oldAttr
}

// SetNamedItemNS adds an attribute with namespace support to the collection.
// Per WHATWG DOM spec: "The setNamedItem(attr) and setNamedItemNS(attr) method steps are to return
// the result of setting an attribute given attr and element."
func (nnm *NamedNodeMap) SetNamedItemNS(attr *Attr) *Attr {
	// Both setNamedItem and setNamedItemNS use the same "set an attribute" algorithm
	return nnm.SetNamedItem(attr)
}

// RemoveNamedItem removes the attribute with the specified qualified name.
// Per WHATWG DOM spec: "The removeNamedItem(qualifiedName) method steps are:
// 1. Let attr be the result of removing an attribute given qualifiedName and element.
// 2. If attr is null, then throw a "NotFoundError" DOMException.
// 3. Return attr."
func (nnm *NamedNodeMap) RemoveNamedItem(qualifiedName string) *Attr {
	// Implementation of "remove an attribute by name" algorithm:
	// If element is in the HTML namespace and its node document is an HTML document,
	// then set qualifiedName to qualifiedName in ASCII lowercase.
	if nnm.element != nil &&
		nnm.element.NamespaceURI() == htmlNamespace &&
		nnm.element.OwnerDocument() != nil &&
		nnm.element.OwnerDocument().documentType == "html" {
		qualifiedName = strings.ToLower(qualifiedName)
	}

	// Let attr be the result of getting an attribute given qualifiedName and element.
	for i, attr := range nnm.attrs {
		if attr.Name() == qualifiedName {
			// Remove the attribute
			nnm.attrs = append(nnm.attrs[:i], nnm.attrs[i+1:]...)
			attr.ownerElement = nil
			return attr
		}
	}

	// If attr is non-null, then remove attr. Return attr.
	// If attr is null, throw a "NotFoundError" DOMException.
	panic(NewNotFoundError("No attribute named '" + qualifiedName + "' found"))
}

// RemoveNamedItemNS removes the attribute with the specified namespace URI and local name.
// Per WHATWG DOM spec: "The removeNamedItemNS(namespace, localName) method steps are:
// 1. Let attr be the result of removing an attribute given namespace, localName, and element.
// 2. If attr is null, then throw a "NotFoundError" DOMException.
// 3. Return attr."
func (nnm *NamedNodeMap) RemoveNamedItemNS(namespace, localName string) *Attr {
	// Implementation of "remove an attribute by namespace and local name" algorithm:
	// If namespace is the empty string, then set it to null.
	if namespace == "" {
		namespace = ""
	}

	// Let attr be the result of getting an attribute given namespace, localName, and element.
	for i, attr := range nnm.attrs {
		if attr.LocalName() == localName && attr.NamespaceURI() == namespace {
			// Remove the attribute
			nnm.attrs = append(nnm.attrs[:i], nnm.attrs[i+1:]...)
			attr.ownerElement = nil
			return attr
		}
	}

	// If attr is non-null, then remove attr. Return attr.
	// If attr is null, throw a "NotFoundError" DOMException.
	panic(NewNotFoundError("No attribute with namespace '" + namespace + "' and local name '" + localName + "' found"))
}

// SupportedPropertyNames returns the supported property names for this NamedNodeMap.
// Per WHATWG DOM spec: "A NamedNodeMap object's supported property names are the return value of running these steps:
//  1. Let names be the qualified names of the attributes in this NamedNodeMap object's attribute list, with duplicates omitted, in order.
//  2. If this NamedNodeMap object's element is in the HTML namespace and its node document is an HTML document, then for each name of names:
//     a. Let lowercaseName be name, in ASCII lowercase.
//     b. If lowercaseName is not equal to name, remove name from names.
//  3. Return names."
func (nnm *NamedNodeMap) SupportedPropertyNames() []string {
	// Step 1: Let names be the qualified names of the attributes in this NamedNodeMap object's attribute list,
	// with duplicates omitted, in order.
	var names []string
	seen := make(map[string]bool)

	isHTMLInHTML := nnm.element != nil &&
		nnm.element.NamespaceURI() == htmlNamespace &&
		nnm.element.OwnerDocument() != nil &&
		nnm.element.OwnerDocument().documentType == "html"

	for _, attr := range nnm.attrs {
		name := attr.Name()

		// For HTML elements in HTML documents, normalize to lowercase
		if isHTMLInHTML {
			name = strings.ToLower(name)
		}

		if !seen[name] {
			names = append(names, name)
			seen[name] = true
		}
	}

	// Step 2: Already handled above for HTML elements in HTML documents
	// For non-HTML elements, no filtering is needed

	// Step 3: Return names.
	return names
}

// GetAttributeNames returns the qualified names of all attributes.
// This is a convenience method that returns all attribute names (not filtered by supported property names).
func (nnm *NamedNodeMap) GetAttributeNames() []string {
	names := make([]string, len(nnm.attrs))
	for i, attr := range nnm.attrs {
		names[i] = attr.Name()
	}
	return names
}

// HasAttributes returns true if the collection contains any attributes.
func (nnm *NamedNodeMap) HasAttributes() bool {
	return len(nnm.attrs) > 0
}

// ToSlice returns a slice of all attributes for iteration.
func (nnm *NamedNodeMap) ToSlice() []*Attr {
	// Return a copy to prevent external modification
	result := make([]*Attr, len(nnm.attrs))
	copy(result, nnm.attrs)
	return result
}

// ForEach iterates over all attributes and calls the provided function for each.
func (nnm *NamedNodeMap) ForEach(fn func(*Attr) bool) {
	for _, attr := range nnm.attrs {
		if !fn(attr) {
			break
		}
	}
}

// SortByName sorts the attributes alphabetically by name.
// This is not required by the specification but can be useful for deterministic output.
func (nnm *NamedNodeMap) SortByName() {
	sort.Slice(nnm.attrs, func(i, j int) bool {
		return nnm.attrs[i].Name() < nnm.attrs[j].Name()
	})
}

// Clear removes all attributes from the collection.
func (nnm *NamedNodeMap) Clear() {
	for _, attr := range nnm.attrs {
		attr.ownerElement = nil
	}
	nnm.attrs = nnm.attrs[:0]
}

// Contains checks if an attribute with the given name exists.
func (nnm *NamedNodeMap) Contains(name string) bool {
	return nnm.GetNamedItem(name) != nil
}

// ContainsNS checks if an attribute with the given namespace and local name exists.
func (nnm *NamedNodeMap) ContainsNS(namespaceURI, localName string) bool {
	return nnm.GetNamedItemNS(namespaceURI, localName) != nil
}
