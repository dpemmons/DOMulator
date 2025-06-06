package dom

import (
	"fmt"
)

// CloneOptions represents options for cloning operations
type CloneOptions struct {
	Document         *Document
	Subtree          bool
	SelfOnly         bool // When true, only clone the node itself, not children
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
	runCloningSteps(node, copy, options.Subtree)
	// TODO: Implement cloning steps hooks for extensibility (now called via runCloningSteps)

	// If parent is non-null, then append copy to parent
	if options.Parent != nil {
		options.Parent.AppendChild(copy)
	}

	// If subtree is true and selfOnly is false, then for each child of node's children, in tree order:
	// clone a node given child with document set to document, subtree set to subtree,
	// parent set to copy, and fallbackRegistry set to fallbackRegistry
	if options.Subtree && !options.SelfOnly {
		children := node.ChildNodes()
		for i := 0; i < children.Length(); i++ {
			child := children.Item(i)
			childOptions := CloneOptions{
				Document:         options.Document,
				Subtree:          true,
				SelfOnly:         false, // Children are always deep cloned
				Parent:           copy,
				FallbackRegistry: options.FallbackRegistry,
			}
			cloneNode(child, childOptions)
		}
	}

	// Handle shadow DOM cloning per spec
	// If node is an element, node is a shadow host, and node's shadow root's clonable is true:
	if originalElement, ok := node.(*Element); ok {
		if originalShadowRoot := originalElement.ShadowRoot(); originalShadowRoot != nil && originalShadowRoot.Clonable() {
			copiedElement, ok := copy.(*Element)
			if !ok {
				// This should not happen if cloneSingleNode works correctly for Elements
				panic(fmt.Sprintf("Cloned node is not an Element when original was, during shadow DOM cloning. Original: %T, Cloned: %T", node, copy))
			}
			if copiedElement.ShadowRoot() != nil {
				panic(NewInvalidStateError("Copy is already a shadow host during shadow root cloning."))
			}

			initOptions := ShadowRootInit{
				Mode:           originalShadowRoot.Mode(),
				SlotAssignment: originalShadowRoot.SlotAssignment(),
				DelegatesFocus: originalShadowRoot.DelegatesFocus(),
				Clonable:       true, // Per spec: "Attach a shadow root with ... true (for clonable)..."
				Serializable:   originalShadowRoot.Serializable(),
			}
			newShadowRoot, err := copiedElement.AttachShadow(initOptions)
			if err != nil {
				// Handle error, e.g., panic or return an error if the function signature allows
				panic(fmt.Sprintf("Failed to attach shadow root during clone: %v", err))
			}

			// Set custom element registry and declarative state per spec
			newShadowRoot.customElementRegistry = originalShadowRoot.CustomElementRegistry()
			// Note: declarative is a private field, so we can access it directly within the same package
			newShadowRoot.declarative = originalShadowRoot.declarative

			// Clone children of the original shadow root into the new shadow root
			shadowChildren := originalShadowRoot.ChildNodes()
			for i := 0; i < shadowChildren.Length(); i++ {
				child := shadowChildren.Item(i)
				childCloneOptions := CloneOptions{
					Document:         options.Document, // Document for the clone operation
					Subtree:          true,             // Children of shadow root are always deep cloned as part of this step
					SelfOnly:         false,
					Parent:           newShadowRoot, // Parent is the new shadow root
					FallbackRegistry: nil,           // Per spec: "This intentionally does not pass the fallbackRegistry argument."
				}
				cloneNode(child, childCloneOptions)
			}
		}
	}

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
	// TODO: Spec requires using a resolved registry (elem's, or fallbackRegistry if elem's is null)
	// for element creation. Current NewElement/NewElementNS use document's registry.
	// FallbackRegistry (if a *CustomElementRegistry) is not currently plumbed through.

	var copy *Element
	var err error

	// Determine the name to use for creation.
	// For HTML elements (or elements in no namespace that are treated as HTML),
	// NewElement expects the lowercase tag name, and it will handle uppercasing for nodeName/tagName.
	// For other namespaced elements, NewElementNS expects the qualified name.
	if elem.NamespaceURI() != "" && elem.NamespaceURI() != htmlNamespace {
		// Non-HTML namespaced element: use its tagName (qualified name)
		copy, err = NewElementNS(elem.NamespaceURI(), elem.TagName(), document)
		if err != nil {
			// This is a more serious error if a namespaced element cannot be cloned with its namespace.
			// Depending on desired strictness, this could panic or attempt a less accurate fallback.
			// For now, we'll let it proceed, but NewElementNS failing is problematic.
			// A simple fallback might be: copy = NewElement(elem.LocalName(), document)
			// but this loses namespace. The original code had a fallback to NewElement(elem.tagName)
			// which is also not ideal for namespaced elements if elem.tagName was just localName.
			// Sticking to NewElementNS and letting error propagate or be handled is better.
			// If NewElementNS fails, the resulting 'copy' might be nil or an imperfect clone.
			// The panic below is a placeholder for more robust error handling.
			panic(fmt.Sprintf("failed to clone namespaced element %s with namespace %s: %v", elem.TagName(), elem.NamespaceURI(), err))
		}
	} else {
		// HTML element (or no namespace, treated as HTML by NewElement):
		// Pass the localName (which is lowercase) to NewElement.
		// NewElement is responsible for setting nodeName and tagName to uppercase for HTML elements.
		copy = NewElement(elem.LocalName(), document)
	}

	// Copy attributes using NamedNodeMap API
	attrs := elem.attributes.ToSlice()
	for _, attr := range attrs {
		copy.SetAttribute(attr.Name(), attr.Value())
	}

	// Note: The fields like copy.nodeName, copy.tagName, copy.localName, copy.namespaceURI, copy.prefix
	// are set by NewElement or NewElementNS. We don't need to copy them manually here after creation
	// if those constructors do their job correctly based on the input name and namespace.
	// The original code had manual copying of these, which might be redundant or conflict
	// if NewElement/NewElementNS already set them based on their input.

	// Set "is" value per spec
	if copy != nil && elem.IsValue() != "" {
		copy.SetIsValue(elem.IsValue())
	}

	return copy
}

// cloneDocument implements document-specific cloning logic
func cloneDocument(doc *Document) Node {
	copy := NewDocument()

	// Set copy's encoding, content type, URL, origin, type, mode, and custom element registry to those of node
	// Per spec: "Set copy's encoding, content type, URL, origin, type, mode, and custom element registry to those of node."
	copy.url = doc.URL()
	copy.documentURI = doc.DocumentURI() // Alias for URL
	copy.contentType = doc.ContentType()
	copy.characterSet = doc.CharacterSet()
	copy.charset = doc.CharacterSet()       // Legacy alias
	copy.inputEncoding = doc.CharacterSet() // Legacy alias
	copy.mode = doc.mode                    // Internal field for CompatMode()
	copy.documentType = doc.documentType    // Internal field for "type" (xml/html)

	// Create a new CustomElementRegistry associated with the cloned document
	// instead of copying the reference to ensure proper association
	if doc.CustomElementRegistry() != nil {
		copy.customElementRegistry = NewCustomElementRegistry(copy)
	}
	// TODO: Copy 'origin' when it's available on the Document struct.

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
	// Create the new attribute using its qualified name and value.
	// NewAttr will parse the qualified name to set initial namespaceURI, prefix, and localName.
	clonedAttr := NewAttr(attr.Name(), attr.Value(), document)

	// Per spec: "Set copy’s namespace, namespace prefix, local name, and value to those of node."
	// Explicitly copy these properties to ensure exact replication, especially if the original
	// Attr was created with specific namespace details not fully reconstructible from its name alone.
	clonedAttr.namespaceURI = attr.NamespaceURI()
	clonedAttr.prefix = attr.Prefix()
	clonedAttr.localName = attr.LocalName()
	// Value is already set by NewAttr, but ensure nodeValue is also consistent.
	clonedAttr.nodeImpl.nodeValue = attr.Value()

	return clonedAttr
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
	if _, ok := node.(*ShadowRoot); ok {
		panic(NewNotSupportedError("ShadowRoot nodes cannot be cloned."))
	}

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
