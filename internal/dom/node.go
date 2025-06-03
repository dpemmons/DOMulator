package dom

import (
	"github.com/dop251/goja"
)

// NodeType represents the type of a Node.
type NodeType int

const (
	ElementNode               NodeType = 1
	AttributeNode             NodeType = 2
	TextNode                  NodeType = 3
	CDataSectionNode          NodeType = 4
	EntityReferenceNode       NodeType = 5
	EntityNode                NodeType = 6
	ProcessingInstructionNode NodeType = 7
	CommentNode               NodeType = 8
	DocumentNode              NodeType = 9
	DocumentTypeNode          NodeType = 10
	DocumentFragmentNode      NodeType = 11
	NotationNode              NodeType = 12
)

// Node represents a single node in the DOM tree.
type Node interface {
	NodeType() NodeType
	NodeName() string
	NodeValue() string

	ParentNode() Node
	ChildNodes() NodeList
	FirstChild() Node
	LastChild() Node
	PreviousSibling() Node
	NextSibling() Node

	AppendChild(child Node) Node
	RemoveChild(child Node) Node
	InsertBefore(newChild, refChild Node) Node
	ReplaceChild(newChild, oldChild Node) Node

	CloneNode(deep bool) Node

	// Internal methods for DOM manipulation and JS binding
	setParent(parent Node)
	setOwnerDocument(doc *Document)
	toJS(vm *goja.Runtime) goja.Value
}

// NodeList represents a collection of nodes.
type NodeList []Node

// nodeImpl provides a common base for all Node types.
type nodeImpl struct {
	nodeType      NodeType
	nodeName      string
	nodeValue     string
	parentNode    Node
	childNodes    NodeList
	ownerDocument *Document
}

func (n *nodeImpl) NodeType() NodeType {
	return n.nodeType
}

func (n *nodeImpl) NodeName() string {
	return n.nodeName
}

func (n *nodeImpl) NodeValue() string {
	return n.nodeValue
}

func (n *nodeImpl) ParentNode() Node {
	return n.parentNode
}

func (n *nodeImpl) ChildNodes() NodeList {
	return n.childNodes
}

func (n *nodeImpl) FirstChild() Node {
	if len(n.childNodes) > 0 {
		return n.childNodes[0]
	}
	return nil
}

func (n *nodeImpl) LastChild() Node {
	if len(n.childNodes) > 0 {
		return n.childNodes[len(n.childNodes)-1]
	}
	return nil
}

func (n *nodeImpl) PreviousSibling() Node {
	if n.parentNode == nil {
		return nil
	}
	for i, child := range n.parentNode.ChildNodes() {
		if child == n {
			if i > 0 {
				return n.parentNode.ChildNodes()[i-1]
			}
			return nil
		}
	}
	return nil
}

func (n *nodeImpl) NextSibling() Node {
	if n.parentNode == nil {
		return nil
	}
	for i, child := range n.parentNode.ChildNodes() {
		if child == n {
			if i < len(n.parentNode.ChildNodes())-1 {
				return n.parentNode.ChildNodes()[i+1]
			}
			return nil
		}
	}
	return nil
}

func (n *nodeImpl) AppendChild(child Node) Node {
	if child == nil {
		return nil
	}
	if child.ParentNode() != nil {
		child.ParentNode().RemoveChild(child)
	}
	n.childNodes = append(n.childNodes, child)
	child.setParent(n)
	child.setOwnerDocument(n.ownerDocument)
	return child
}

func (n *nodeImpl) RemoveChild(child Node) Node {
	if child == nil {
		return nil
	}
	for i, c := range n.childNodes {
		if c == child {
			n.childNodes = append(n.childNodes[:i], n.childNodes[i+1:]...)
			child.setParent(nil)
			return child
		}
	}
	return nil // Child not found
}

func (n *nodeImpl) InsertBefore(newChild, refChild Node) Node {
	if newChild == nil {
		return nil
	}
	if newChild.ParentNode() != nil {
		newChild.ParentNode().RemoveChild(newChild)
	}

	if refChild == nil {
		return n.AppendChild(newChild)
	}

	for i, c := range n.childNodes {
		if c == refChild {
			n.childNodes = append(n.childNodes[:i], append(NodeList{newChild}, n.childNodes[i:]...)...)
			newChild.setParent(n)
			newChild.setOwnerDocument(n.ownerDocument)
			return newChild
		}
	}
	return nil // refChild not found
}

func (n *nodeImpl) ReplaceChild(newChild, oldChild Node) Node {
	if newChild == nil || oldChild == nil {
		return nil
	}
	if newChild.ParentNode() != nil {
		newChild.ParentNode().RemoveChild(newChild)
	}

	for i, c := range n.childNodes {
		if c == oldChild {
			n.childNodes[i] = newChild
			newChild.setParent(n)
			newChild.setOwnerDocument(n.ownerDocument)
			oldChild.setParent(nil)
			return oldChild
		}
	}
	return nil // oldChild not found
}

func (n *nodeImpl) CloneNode(deep bool) Node {
	clone := &nodeImpl{
		nodeType:      n.nodeType,
		nodeName:      n.nodeName,
		nodeValue:     n.nodeValue,
		ownerDocument: n.ownerDocument,
	}
	if deep {
		for _, child := range n.childNodes {
			clone.AppendChild(child.CloneNode(true))
		}
	}
	return clone
}

func (n *nodeImpl) setParent(parent Node) {
	n.parentNode = parent
}

func (n *nodeImpl) setOwnerDocument(doc *Document) {
	n.ownerDocument = doc
	for _, child := range n.childNodes {
		child.setOwnerDocument(doc)
	}
}

// toJS is a placeholder for JavaScript binding.
func (n *nodeImpl) toJS(vm *goja.Runtime) goja.Value {
	// This will be implemented later when integrating with Goja
	return vm.ToValue(nil)
}
