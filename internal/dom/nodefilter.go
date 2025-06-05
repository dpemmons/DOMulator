package dom

// NodeFilter constants per WHATWG DOM specification
const (
	// NodeFilter show constants
	NodeFilterShowAll                   = 0xFFFFFFFF
	NodeFilterShowElement               = 0x1
	NodeFilterShowAttribute             = 0x2 // Deprecated
	NodeFilterShowText                  = 0x4
	NodeFilterShowCDataSection          = 0x8
	NodeFilterShowEntityReference       = 0x10 // Deprecated
	NodeFilterShowEntity                = 0x20 // Deprecated
	NodeFilterShowProcessingInstruction = 0x40
	NodeFilterShowComment               = 0x80
	NodeFilterShowDocument              = 0x100
	NodeFilterShowDocumentType          = 0x200
	NodeFilterShowDocumentFragment      = 0x400
	NodeFilterShowNotation              = 0x800 // Deprecated

	// NodeFilter accept constants
	NodeFilterAccept = 1
	NodeFilterReject = 2
	NodeFilterSkip   = 3
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
