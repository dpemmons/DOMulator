package dom

import (
	"fmt"
)

// CloneOptions represents options for cloning operations
type CloneOptions struct {
	Document         *Document
	Subtree          bool
	Parent           Node
	FallbackRegistry interface{} // For custom element registry (not fully implemented yet)
}

// cloneNode implements the "clone a node" algorithm from the DOM specification
// https://dom.spec.whatwg.org/#concept-node-clone
func cloneNode(node Node, options CloneOptions) Node {
	// Assert: node is not a document or node is document
	if options.Document == nil {
		options.Document = node.OwnerDocument()
	}

	// Let copy be the result of cloning a single node given node, document, and fallbackRegistry
	copy := cloneSingleNode(node, options.Document, options.FallbackRegistry)

	// Run any cloning steps defined for node in other applicable specifications
	// TODO: Implement cloning steps hooks for extensibility

	// If parent is non-null, then append copy to parent
	if options.Parent != nil {
		options.Parent.AppendChild(copy)
	}

	// If subtree is true, then for each child of node's children, in tree order:
	// clone a node given child with document set to document, subtree set to subtree,
	// parent set to copy, and fallbackRegistry set to fallbackRegistry
	if options.Subtree {
		children := node.ChildNodes()
		for i := 0; i < children.Length(); i++ {
			child := children.Item(i)
			childOptions := CloneOptions{
				Document:         options.Document,
				Subtree:          true,
				Parent:           copy,
				FallbackRegistry: options.FallbackRegistry,
			}
			cloneNode(child, childOptions)
		}
	}

	// TODO: Handle shadow DOM cloning when shadow DOM is implemented
	// If node is an element, node is a shadow host, and node's shadow root's clonable is true:
	// - Assert: copy is not a shadow host
	// - Attach shadow root and clone shadow children

	return copy
}

// cloneSingleNode implements the "clone a single node" algorithm from the DOM specification
// https://dom.spec.whatwg.org/#concept-node-clone-single
func cloneSingleNode(node Node, document *Document, fallbackRegistry interface{}) Node {
	if node == nil {
		return nil
	}

	var copy Node

	switch n := node.(type) {
	case *Element:
		copy = cloneElement(n, document, fallbackRegistry)

	case *Document:
		copy = cloneDocument(n)

	case *DocumentType:
		copy = cloneDocumentType(n, document)

	case *Attr:
		copy = cloneAttr(n, document)

	case *Text:
		copy = cloneText(n, document)

	case *Comment:
		copy = cloneComment(n, document)

	case *ProcessingInstruction:
		copy = cloneProcessingInstruction(n, document)

	case *DocumentFragment:
		copy = cloneDocumentFragment(n, document)

	case *CDATASection:
		copy = cloneCDATASection(n, document)

	default:
		// For unknown node types, create a basic copy
		copy = &nodeImpl{
			nodeType:      node.NodeType(),
			nodeName:      node.NodeName(),
			nodeValue:     node.NodeValue(),
			ownerDocument: document,
		}
	}

	// Assert: copy is a node
	if copy == nil {
		panic("cloning failed to produce a valid node")
	}

	// If node is a document, then set document to copy
	if _, isDocument := node.(*Document); isDocument {
		document = copy.(*Document)
	}

	// Set copy's node document to document
	copy.setOwnerDocument(document)

	return copy
}

// cloneElement implements element-specific cloning logic
func cloneElement(elem *Element, document *Document, fallbackRegistry interface{}) Node {
	// TODO: Handle custom element registry when implemented
	// For now, use the basic element creation

	var copy *Element
	var err error

	// Create element with proper namespace support
	if elem.namespaceURI != "" {
		copy, err = NewElementNS(elem.namespaceURI, elem.tagName, document)
		if err != nil {
			// Fallback to basic element creation
			copy = NewElement(elem.tagName, document)
		}
	} else {
		copy = NewElement(elem.tagName, document)
	}

	// Ensure the copy has the same nodeName as the original
	copy.nodeName = elem.nodeName

	// Copy namespace information
	copy.namespaceURI = elem.namespaceURI
	copy.prefix = elem.prefix
	copy.localName = elem.localName

	// For each attribute of node's attribute list:
	// Let copyAttribute be the result of cloning a single node given attribute, document, and null
	// Append copyAttribute to copy
	for name, value := range elem.attributes {
		copy.SetAttribute(name, value)
	}

	return copy
}

// cloneDocument implements document-specific cloning logic
func cloneDocument(doc *Document) Node {
	copy := NewDocument()

	// Set copy's encoding, content type, URL, origin, type, mode, and custom element registry to those of node
	// TODO: Implement these properties when Document is enhanced
	// For now, we just copy basic properties

	return copy
}

// cloneDocumentType implements DocumentType-specific cloning logic
func cloneDocumentType(dt *DocumentType, document *Document) Node {
	// Set copy's name, public ID, and system ID to those of node
	copy := NewDocumentType(dt.name, dt.publicID, dt.systemID, document)
	copy.nodeImpl.self = copy
	return copy
}

// cloneAttr implements Attr-specific cloning logic
func cloneAttr(attr *Attr, document *Document) Node {
	// Set copy's namespace, namespace prefix, local name, and value to those of node
	copy := NewAttr(attr.name, attr.value, document)

	// TODO: Copy namespace information when Attr supports namespaces
	// copy.namespaceURI = attr.namespaceURI
	// copy.prefix = attr.prefix
	// copy.localName = attr.localName

	return copy
}

// cloneText implements Text-specific cloning logic
func cloneText(text *Text, document *Document) Node {
	// Set copy's data to that of node
	copy := NewText(text.Data(), document)
	return copy
}

// cloneComment implements Comment-specific cloning logic
func cloneComment(comment *Comment, document *Document) Node {
	// Set copy's data to that of node
	copy := NewComment(comment.Data(), document)
	return copy
}

// cloneProcessingInstruction implements ProcessingInstruction-specific cloning logic
func cloneProcessingInstruction(pi *ProcessingInstruction, document *Document) Node {
	// Set copy's target and data to those of node
	copy := NewProcessingInstruction(pi.Target(), pi.Data(), document)
	return copy
}

// cloneDocumentFragment implements DocumentFragment-specific cloning logic
func cloneDocumentFragment(fragment *DocumentFragment, document *Document) Node {
	// DocumentFragment has no special properties to copy
	copy := NewDocumentFragment(document)
	return copy
}

// cloneCDATASection implements CDATASection-specific cloning logic
func cloneCDATASection(cdata *CDATASection, document *Document) Node {
	// Set copy's data to that of node
	copy := NewCDATASection(cdata.Data(), document)
	return copy
}

// CloningStepsHook represents a function that can be called during node cloning
// to perform additional cloning operations defined by other specifications
type CloningStepsHook func(original Node, copy Node, subtree bool)

// Global registry for cloning steps hooks
var cloningStepsHooks = make(map[NodeType][]CloningStepsHook)

// RegisterCloningSteps registers a cloning steps hook for a specific node type
// This allows other packages to extend the cloning behavior
func RegisterCloningSteps(nodeType NodeType, hook CloningStepsHook) {
	cloningStepsHooks[nodeType] = append(cloningStepsHooks[nodeType], hook)
}

// runCloningSteps executes all registered cloning steps for a node
func runCloningSteps(original Node, copy Node, subtree bool) {
	if hooks, exists := cloningStepsHooks[original.NodeType()]; exists {
		for _, hook := range hooks {
			hook(original, copy, subtree)
		}
	}
}

// CloneNodeSpec implements the spec-compliant cloneNode method
// This is used internally by the Node implementations
func CloneNodeSpec(node Node, subtree bool) Node {
	// If this is a shadow root, throw NotSupportedError
	// TODO: Implement when shadow DOM is added

	// Return the result of cloning a node given this with subtree set to subtree
	options := CloneOptions{
		Document:         node.OwnerDocument(),
		Subtree:          subtree,
		Parent:           nil,
		FallbackRegistry: nil,
	}
	return cloneNode(node, options)
}

// CloneNodeWithDocument clones a node into a different document
func CloneNodeWithDocument(node Node, targetDocument *Document, subtree bool) Node {
	options := CloneOptions{
		Document:         targetDocument,
		Subtree:          subtree,
		Parent:           nil,
		FallbackRegistry: nil,
	}
	return cloneNode(node, options)
}

// validateClonedNode performs post-cloning validation to ensure spec compliance
func validateClonedNode(original Node, cloned Node, subtree bool) error {
	// Basic validation
	if cloned == nil {
		return fmt.Errorf("cloned node cannot be nil")
	}

	// Node type must match
	if original.NodeType() != cloned.NodeType() {
		return fmt.Errorf("cloned node type %d does not match original type %d",
			cloned.NodeType(), original.NodeType())
	}

	// Node name must match for most node types
	if original.NodeName() != cloned.NodeName() {
		return fmt.Errorf("cloned node name %q does not match original name %q",
			cloned.NodeName(), original.NodeName())
	}

	// If subtree cloning, child count should match
	if subtree {
		originalChildren := original.ChildNodes()
		clonedChildren := cloned.ChildNodes()
		if originalChildren.Length() != clonedChildren.Length() {
			return fmt.Errorf("cloned node has %d children, original has %d",
				clonedChildren.Length(), originalChildren.Length())
		}
	} else {
		// Shallow clone should have no children
		clonedChildren := cloned.ChildNodes()
		if clonedChildren.Length() > 0 {
			return fmt.Errorf("shallow cloned node should have no children, but has %d",
				clonedChildren.Length())
		}
	}

	return nil
}
