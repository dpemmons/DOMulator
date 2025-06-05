package dom

import (
	"github.com/dop251/goja"
)

// Attr represents an attribute of an Element.
type Attr struct {
	nodeImpl
	name         string
	value        string
	ownerElement *Element

	// Namespace support
	namespaceURI string
	prefix       string
	localName    string
}

// NewAttr creates a new Attr node.
func NewAttr(name, value string, doc *Document) *Attr {
	// Parse qualified name for namespace information
	info := ParseQualifiedName(name)

	attr := &Attr{
		nodeImpl: nodeImpl{
			nodeType:      AttributeNode,
			nodeName:      name,
			nodeValue:     value,
			ownerDocument: doc,
		},
		name:         name,
		value:        value,
		namespaceURI: info.NamespaceURI,
		prefix:       info.Prefix,
		localName:    info.LocalName,
	}
	attr.nodeImpl.self = attr // Set the self reference
	return attr
}

// NewAttrNS creates a new Attr node with namespace support.
func NewAttrNS(namespaceURI, qualifiedName, value string, doc *Document) (*Attr, error) {
	// Validate and extract namespace information
	ns, prefix, localName, err := ValidateAndExtract(namespaceURI, qualifiedName)
	if err != nil {
		return nil, err
	}

	attr := &Attr{
		nodeImpl: nodeImpl{
			nodeType:      AttributeNode,
			nodeName:      qualifiedName,
			nodeValue:     value,
			ownerDocument: doc,
		},
		name:         qualifiedName,
		value:        value,
		namespaceURI: ns,
		prefix:       prefix,
		localName:    localName,
	}
	attr.nodeImpl.self = attr // Set the self reference
	return attr, nil
}

// Name returns the name of the attribute.
func (a *Attr) Name() string {
	return a.name
}

// Value returns the value of the attribute.
func (a *Attr) Value() string {
	return a.value
}

// SetValue sets the value of the attribute.
func (a *Attr) SetValue(value string) {
	a.value = value
	a.nodeImpl.nodeValue = value
}

// OwnerElement returns the Element that owns this attribute.
func (a *Attr) OwnerElement() *Element {
	return a.ownerElement
}

// NamespaceURI returns the namespace URI of the attribute.
func (a *Attr) NamespaceURI() string {
	return a.namespaceURI
}

// Prefix returns the namespace prefix of the attribute.
func (a *Attr) Prefix() string {
	return a.prefix
}

// LocalName returns the local name of the attribute.
func (a *Attr) LocalName() string {
	return a.localName
}

// CloneNode creates a copy of the attribute using the spec-compliant cloning implementation.
func (a *Attr) CloneNode(deep bool) Node {
	// Use the spec-compliant cloning implementation
	return CloneNodeSpec(a, deep)
}

// toJS is a placeholder for JavaScript binding.
func (a *Attr) toJS(vm *goja.Runtime) goja.Value {
	// This will be implemented later when integrating with Goja
	return vm.ToValue(nil)
}

// AttributeMap is a map of attribute names to Attr nodes.
type AttributeMap map[string]*Attr

// Get returns an attribute by name.
func (am AttributeMap) Get(name string) *Attr {
	return am[name]
}

// Set sets an attribute.
func (am AttributeMap) Set(attr *Attr) {
	am[attr.Name()] = attr
	attr.ownerElement = am.ownerElement() // Set owner element
}

// Remove removes an attribute by name.
func (am AttributeMap) Remove(name string) {
	delete(am, name)
}

// ownerElement is a helper to get the owner element from the map.
// This is a bit of a hack, ideally AttributeMap would be part of Element.
func (am AttributeMap) ownerElement() *Element {
	for _, attr := range am {
		return attr.ownerElement // All attributes in the map should have the same owner
	}
	return nil
}
