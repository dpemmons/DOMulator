package dom

// NodeFilter constants per WHATWG DOM Section 6.3 - Interface NodeFilter
//
// Note: When exposed to JavaScript, these constants should use spec-compliant names:
// - NodeFilterAccept -> FILTER_ACCEPT
// - NodeFilterReject -> FILTER_REJECT
// - NodeFilterSkip -> FILTER_SKIP
// - NodeFilterShowAll -> SHOW_ALL
// - NodeFilterShowElement -> SHOW_ELEMENT
// etc.
const (
	// NodeFilter show constants for whatToShow bitmask
	// SHOW_ALL (4294967295, FFFFFFFF in hexadecimal)
	NodeFilterShowAll = 0xFFFFFFFF

	// SHOW_ELEMENT (1)
	NodeFilterShowElement = 0x1

	// SHOW_ATTRIBUTE (2) - LEGACY: This constant is historical and should not be used
	NodeFilterShowAttribute = 0x2

	// SHOW_TEXT (4)
	NodeFilterShowText = 0x4

	// SHOW_CDATA_SECTION (8)
	NodeFilterShowCDataSection = 0x8

	// SHOW_ENTITY_REFERENCE (16, 10 in hexadecimal) - LEGACY: This constant is historical and should not be used
	NodeFilterShowEntityReference = 0x10

	// SHOW_ENTITY (32, 20 in hexadecimal) - LEGACY: This constant is historical and should not be used
	NodeFilterShowEntity = 0x20

	// SHOW_PROCESSING_INSTRUCTION (64, 40 in hexadecimal)
	NodeFilterShowProcessingInstruction = 0x40

	// SHOW_COMMENT (128, 80 in hexadecimal)
	NodeFilterShowComment = 0x80

	// SHOW_DOCUMENT (256, 100 in hexadecimal)
	NodeFilterShowDocument = 0x100

	// SHOW_DOCUMENT_TYPE (512, 200 in hexadecimal)
	NodeFilterShowDocumentType = 0x200

	// SHOW_DOCUMENT_FRAGMENT (1024, 400 in hexadecimal)
	NodeFilterShowDocumentFragment = 0x400

	// SHOW_NOTATION (2048, 800 in hexadecimal) - LEGACY: This constant is historical and should not be used
	NodeFilterShowNotation = 0x800

	// NodeFilter accept constants for acceptNode() return values
	// FILTER_ACCEPT (1) - Accept the node
	NodeFilterAccept = 1

	// FILTER_REJECT (2) - Reject the node and its children
	NodeFilterReject = 2

	// FILTER_SKIP (3) - Skip the node but consider its children
	NodeFilterSkip = 3
)

// NodeFilter represents a function that tests whether a node should be accepted by a NodeIterator or TreeWalker
type NodeFilter interface {
	AcceptNode(node Node) int
}

// NodeFilterFunc is a function type that implements NodeFilter
type NodeFilterFunc func(node Node) int

// AcceptNode implements the NodeFilter interface for function types
func (f NodeFilterFunc) AcceptNode(node Node) int {
	return f(node)
}

// DefaultNodeFilter accepts all nodes
var DefaultNodeFilter = NodeFilterFunc(func(node Node) int {
	return NodeFilterAccept
})

// whatToShowFilter creates a NodeFilter based on whatToShow bitmask
func whatToShowFilter(whatToShow uint32) NodeFilter {
	return NodeFilterFunc(func(node Node) int {
		nodeType := node.NodeType()

		// Map node types to whatToShow bits
		var bit uint32
		switch nodeType {
		case ElementNode:
			bit = NodeFilterShowElement
		case AttributeNode:
			bit = NodeFilterShowAttribute
		case TextNode:
			bit = NodeFilterShowText
		case CDataSectionNode:
			bit = NodeFilterShowCDataSection
		case EntityReferenceNode:
			bit = NodeFilterShowEntityReference
		case EntityNode:
			bit = NodeFilterShowEntity
		case ProcessingInstructionNode:
			bit = NodeFilterShowProcessingInstruction
		case CommentNode:
			bit = NodeFilterShowComment
		case DocumentNode:
			bit = NodeFilterShowDocument
		case DocumentTypeNode:
			bit = NodeFilterShowDocumentType
		case DocumentFragmentNode:
			bit = NodeFilterShowDocumentFragment
		case NotationNode:
			bit = NodeFilterShowNotation
		default:
			return NodeFilterReject
		}

		// Check if this node type should be shown
		if whatToShow&bit != 0 {
			return NodeFilterAccept
		}
		return NodeFilterSkip
	})
}

// combinedFilter combines whatToShow filtering with custom filter
func combinedFilter(whatToShow uint32, filter NodeFilter) NodeFilter {
	whatToShowFilt := whatToShowFilter(whatToShow)

	if filter == nil {
		return whatToShowFilt
	}

	return NodeFilterFunc(func(node Node) int {
		// First check whatToShow
		result := whatToShowFilt.AcceptNode(node)
		if result != NodeFilterAccept {
			return result
		}

		// Then apply custom filter
		return filter.AcceptNode(node)
	})
}
