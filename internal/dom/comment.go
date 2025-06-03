package dom

import (
	"github.com/dop251/goja"
)

// Comment represents a comment node in the DOM tree.
type Comment struct {
	nodeImpl
}

// NewComment creates a new Comment node.
func NewComment(data string, doc *Document) *Comment {
	comment := &Comment{
		nodeImpl: nodeImpl{
			nodeType:      CommentNode,
			nodeName:      "#comment",
			nodeValue:     data,
			ownerDocument: doc,
		},
	}
	return comment
}

// toJS is a placeholder for JavaScript binding.
func (c *Comment) toJS(vm *goja.Runtime) goja.Value {
	// This will be implemented later when integrating with Goja
	return vm.ToValue(nil)
}
