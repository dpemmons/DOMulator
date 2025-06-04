package dom

import (
	"github.com/dop251/goja"
)

// Text represents a text node in the DOM tree.
type Text struct {
	nodeImpl
}

// NewText creates a new Text node.
func NewText(data string, doc *Document) *Text {
	text := &Text{
		nodeImpl: nodeImpl{
			nodeType:      TextNode,
			nodeName:      "#text",
			nodeValue:     data,
			ownerDocument: doc,
		},
	}
	text.nodeImpl.self = text // Set the self reference
	return text
}

// Data returns the text content of the text node.
func (t *Text) Data() string {
	return t.nodeValue
}

// SetData sets the text content of the text node.
func (t *Text) SetData(data string) {
	t.nodeValue = data
}

// Length returns the length of the text data.
func (t *Text) Length() int {
	return len([]rune(t.nodeValue)) // Use runes for proper Unicode support
}

// SubstringData extracts a range of data from the text node.
func (t *Text) SubstringData(offset, count int) string {
	runes := []rune(t.nodeValue)

	// Handle negative offset
	if offset < 0 {
		offset = 0
	}

	// If offset is beyond the text, return empty string
	if offset >= len(runes) {
		return ""
	}

	// Calculate the end position
	end := offset + count
	if end > len(runes) {
		end = len(runes)
	}

	return string(runes[offset:end])
}

// AppendData appends data to the end of the text node.
func (t *Text) AppendData(data string) {
	t.nodeValue += data
}

// InsertData inserts data at the specified offset.
func (t *Text) InsertData(offset int, data string) {
	runes := []rune(t.nodeValue)

	// Handle negative offset
	if offset < 0 {
		offset = 0
	}

	// If offset is beyond the text, append
	if offset >= len(runes) {
		t.nodeValue += data
		return
	}

	// Insert the data
	newRunes := make([]rune, 0, len(runes)+len([]rune(data)))
	newRunes = append(newRunes, runes[:offset]...)
	newRunes = append(newRunes, []rune(data)...)
	newRunes = append(newRunes, runes[offset:]...)

	t.nodeValue = string(newRunes)
}

// DeleteData deletes data from the text node.
func (t *Text) DeleteData(offset, count int) {
	runes := []rune(t.nodeValue)

	// Handle negative offset
	if offset < 0 {
		offset = 0
	}

	// If offset is beyond the text, do nothing
	if offset >= len(runes) {
		return
	}

	// Calculate the end position
	end := offset + count
	if end > len(runes) {
		end = len(runes)
	}

	// Delete the data
	newRunes := make([]rune, 0, len(runes)-(end-offset))
	newRunes = append(newRunes, runes[:offset]...)
	newRunes = append(newRunes, runes[end:]...)

	t.nodeValue = string(newRunes)
}

// ReplaceData replaces data in the text node.
func (t *Text) ReplaceData(offset, count int, data string) {
	runes := []rune(t.nodeValue)

	// Handle negative offset
	if offset < 0 {
		offset = 0
	}

	// If offset is beyond the text, append
	if offset >= len(runes) {
		t.nodeValue += data
		return
	}

	// Calculate the end position
	end := offset + count
	if end > len(runes) {
		end = len(runes)
	}

	// Replace the data
	newRunes := make([]rune, 0, len(runes)-(end-offset)+len([]rune(data)))
	newRunes = append(newRunes, runes[:offset]...)
	newRunes = append(newRunes, []rune(data)...)
	newRunes = append(newRunes, runes[end:]...)

	t.nodeValue = string(newRunes)
}

// SplitText splits this text node into two text nodes at the specified offset.
func (t *Text) SplitText(offset int) *Text {
	runes := []rune(t.nodeValue)

	// Handle negative offset
	if offset < 0 {
		offset = 0
	}

	// If offset is beyond the text, create empty new text
	if offset >= len(runes) {
		newText := NewText("", t.ownerDocument)
		return newText
	}

	// Split the text
	originalData := string(runes[:offset])
	newData := string(runes[offset:])

	// Update original text
	t.nodeValue = originalData

	// Create new text node
	newText := NewText(newData, t.ownerDocument)

	return newText
}

// CloneNode creates a copy of the text node.
func (t *Text) CloneNode(deep bool) Node {
	clone := NewText(t.nodeValue, t.ownerDocument)
	return clone
}

// toJS is a placeholder for JavaScript binding.
func (t *Text) toJS(vm *goja.Runtime) goja.Value {
	// This will be implemented later when integrating with Goja
	return vm.ToValue(nil)
}

// ChildNode mixin methods

// Before inserts nodes just before this text node in its parent's children list.
func (t *Text) Before(nodes ...interface{}) error {
	parent := t.ParentNode()
	if parent == nil {
		return nil // No parent, nothing to do
	}

	// Convert nodes to a single node or document fragment
	convertedNode, err := convertNodesToNode(nodes, t.ownerDocument)
	if err != nil {
		return err
	}

	if convertedNode == nil {
		return nil // Nothing to insert
	}

	// Insert before this text node
	parent.InsertBefore(convertedNode, t)
	return nil
}

// After inserts nodes just after this text node in its parent's children list.
func (t *Text) After(nodes ...interface{}) error {
	parent := t.ParentNode()
	if parent == nil {
		return nil // No parent, nothing to do
	}

	// Convert nodes to a single node or document fragment
	convertedNode, err := convertNodesToNode(nodes, t.ownerDocument)
	if err != nil {
		return err
	}

	if convertedNode == nil {
		return nil // Nothing to insert
	}

	// Find the next viable sibling (not in the nodes being inserted)
	var excludeNodes []Node
	if fragment, ok := convertedNode.(*DocumentFragment); ok {
		excludeNodes = fragment.ChildNodes()
	} else {
		excludeNodes = []Node{convertedNode}
	}

	nextSibling := findViableNextSibling(t, excludeNodes)
	if nextSibling != nil {
		parent.InsertBefore(convertedNode, nextSibling)
	} else {
		parent.AppendChild(convertedNode)
	}
	return nil
}

// ReplaceWith replaces this text node with the given nodes.
func (t *Text) ReplaceWith(nodes ...interface{}) error {
	parent := t.ParentNode()
	if parent == nil {
		return nil // No parent, nothing to do
	}

	// Convert nodes to a single node or document fragment
	convertedNode, err := convertNodesToNode(nodes, t.ownerDocument)
	if err != nil {
		return err
	}

	if convertedNode == nil {
		// Just remove this text node
		parent.RemoveChild(t)
		return nil
	}

	// Replace this text node with the converted node
	parent.ReplaceChild(convertedNode, t)
	return nil
}

// Remove removes this text node from its parent.
func (t *Text) Remove() error {
	parent := t.ParentNode()
	if parent == nil {
		return nil // No parent, nothing to do
	}

	parent.RemoveChild(t)
	return nil
}

// NonDocumentTypeChildNode mixin methods

// PreviousElementSibling returns the first preceding sibling that is an element; otherwise null.
func (t *Text) PreviousElementSibling() *Element {
	parent := t.ParentNode()
	if parent == nil {
		return nil
	}

	siblings := parent.ChildNodes()
	for i := len(siblings) - 1; i >= 0; i-- {
		if siblings[i] == t {
			// Found this text node, now look backwards for an element sibling
			for j := i - 1; j >= 0; j-- {
				if siblings[j].NodeType() == ElementNode {
					return siblings[j].(*Element)
				}
			}
			break
		}
	}
	return nil
}

// NextElementSibling returns the first following sibling that is an element; otherwise null.
func (t *Text) NextElementSibling() *Element {
	parent := t.ParentNode()
	if parent == nil {
		return nil
	}

	siblings := parent.ChildNodes()
	for i, sibling := range siblings {
		if sibling == t {
			// Found this text node, now look forwards for an element sibling
			for j := i + 1; j < len(siblings); j++ {
				if siblings[j].NodeType() == ElementNode {
					return siblings[j].(*Element)
				}
			}
			break
		}
	}
	return nil
}
