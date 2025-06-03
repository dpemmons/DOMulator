package dom

import (
	"github.com/dop251/goja"
)

// DocumentType represents a DocumentType node in the DOM tree.
type DocumentType struct {
	nodeImpl
	name     string
	publicID string
	systemID string
}

// NewDocumentType creates a new DocumentType node.
func NewDocumentType(name, publicID, systemID string, doc *Document) *DocumentType {
	dt := &DocumentType{
		nodeImpl: nodeImpl{
			nodeType:      DocumentTypeNode,
			nodeName:      name,
			ownerDocument: doc,
		},
		name:     name,
		publicID: publicID,
		systemID: systemID,
	}
	return dt
}

// Name returns the name of the document type.
func (dt *DocumentType) Name() string {
	return dt.name
}

// PublicID returns the public ID of the document type.
func (dt *DocumentType) PublicID() string {
	return dt.publicID
}

// SystemID returns the system ID of the document type.
func (dt *DocumentType) SystemID() string {
	return dt.systemID
}

// toJS is a placeholder for JavaScript binding.
func (dt *DocumentType) toJS(vm *goja.Runtime) goja.Value {
	// This will be implemented later when integrating with Goja
	return vm.ToValue(nil)
}
