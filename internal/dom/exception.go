package dom

import (
	"github.com/dop251/goja"
)

// DOMException represents a DOM exception per WHATWG DOM specification.
// https://dom.spec.whatwg.org/#interface-domexception
type DOMException struct {
	name    string
	message string
}

// Legacy code constants as defined in WHATWG DOM specification
const (
	INDEX_SIZE_ERR              = 1
	DOMSTRING_SIZE_ERR          = 2
	HIERARCHY_REQUEST_ERR       = 3
	WRONG_DOCUMENT_ERR          = 4
	INVALID_CHARACTER_ERR       = 5
	NO_DATA_ALLOWED_ERR         = 6
	NO_MODIFICATION_ALLOWED_ERR = 7
	NOT_FOUND_ERR               = 8
	NOT_SUPPORTED_ERR           = 9
	INUSE_ATTRIBUTE_ERR         = 10
	INVALID_STATE_ERR           = 11
	SYNTAX_ERR                  = 12
	INVALID_MODIFICATION_ERR    = 13
	NAMESPACE_ERR               = 14
	INVALID_ACCESS_ERR          = 15
	VALIDATION_ERR              = 16
	TYPE_MISMATCH_ERR           = 17
	SECURITY_ERR                = 18
	NETWORK_ERR                 = 19
	ABORT_ERR                   = 20
	URL_MISMATCH_ERR            = 21
	QUOTA_EXCEEDED_ERR          = 22
	TIMEOUT_ERR                 = 23
	INVALID_NODE_TYPE_ERR       = 24
	DATA_CLONE_ERR              = 25
)

// DOMException names table mapping names to legacy codes
var domExceptionCodes = map[string]int{
	"IndexSizeError":             INDEX_SIZE_ERR,
	"HierarchyRequestError":      HIERARCHY_REQUEST_ERR,
	"WrongDocumentError":         WRONG_DOCUMENT_ERR,
	"InvalidCharacterError":      INVALID_CHARACTER_ERR,
	"NoModificationAllowedError": NO_MODIFICATION_ALLOWED_ERR,
	"NotFoundError":              NOT_FOUND_ERR,
	"NotSupportedError":          NOT_SUPPORTED_ERR,
	"InUseAttributeError":        INUSE_ATTRIBUTE_ERR,
	"InvalidStateError":          INVALID_STATE_ERR,
	"SyntaxError":                SYNTAX_ERR,
	"InvalidModificationError":   INVALID_MODIFICATION_ERR,
	"NamespaceError":             NAMESPACE_ERR,
	"InvalidAccessError":         INVALID_ACCESS_ERR,
	"SecurityError":              SECURITY_ERR,
	"NetworkError":               NETWORK_ERR,
	"AbortError":                 ABORT_ERR,
	"URLMismatchError":           URL_MISMATCH_ERR,
	"QuotaExceededError":         QUOTA_EXCEEDED_ERR,
	"TimeoutError":               TIMEOUT_ERR,
	"InvalidNodeTypeError":       INVALID_NODE_TYPE_ERR,
	"DataCloneError":             DATA_CLONE_ERR,
}

// NewDOMException creates a new DOMException with the specified message and name.
// The constructor steps are:
// 1. Set this's name to name.
// 2. Set this's message to message.
func NewDOMException(message, name string) *DOMException {
	if name == "" {
		name = "Error"
	}
	return &DOMException{
		name:    name,
		message: message,
	}
}

// Name getter steps are to return this's name.
func (e *DOMException) Name() string {
	return e.name
}

// Message getter steps are to return this's message.
func (e *DOMException) Message() string {
	return e.message
}

// Code getter steps are to return the legacy code indicated in the DOMException names table
// for this's name, or 0 if no such entry exists in the table.
func (e *DOMException) Code() int {
	if code, exists := domExceptionCodes[e.name]; exists {
		return code
	}
	return 0
}

// Error implements the Go error interface.
func (e *DOMException) Error() string {
	if e.message != "" {
		return e.name + ": " + e.message
	}
	return e.name
}

// Convenience constructors for common DOM exceptions

// NewHierarchyRequestError creates a HierarchyRequestError.
// Used when the insertion point is of a type that does not allow children of the type of node being inserted.
func NewHierarchyRequestError(message string) *DOMException {
	return NewDOMException(message, "HierarchyRequestError")
}

// NewNotFoundError creates a NotFoundError.
// Used when a reference node is not found.
func NewNotFoundError(message string) *DOMException {
	return NewDOMException(message, "NotFoundError")
}

// NewIndexSizeError creates an IndexSizeError.
// Used when an index or size is negative or greater than the allowed amount.
func NewIndexSizeError(message string) *DOMException {
	return NewDOMException(message, "IndexSizeError")
}

// NewNamespaceError creates a NamespaceError.
// Used when an attempt is made to create or change an object in a way that is incorrect with regard to namespaces.
func NewNamespaceError(message string) *DOMException {
	return NewDOMException(message, "NamespaceError")
}

// NewSyntaxError creates a SyntaxError.
// Used when an invalid or illegal string is specified.
func NewSyntaxError(message string) *DOMException {
	return NewDOMException(message, "SyntaxError")
}

// NewInvalidCharacterError creates an InvalidCharacterError.
// Used when an invalid character is specified (e.g., in an XML name).
func NewInvalidCharacterError(message string) *DOMException {
	return NewDOMException(message, "InvalidCharacterError")
}

// NewWrongDocumentError creates a WrongDocumentError.
// Used when a node is used in a document other than the one that created it.
func NewWrongDocumentError(message string) *DOMException {
	return NewDOMException(message, "WrongDocumentError")
}

// NewNoModificationAllowedError creates a NoModificationAllowedError.
// Used when an attempt is made to modify an object where modifications are not allowed.
func NewNoModificationAllowedError(message string) *DOMException {
	return NewDOMException(message, "NoModificationAllowedError")
}

// NewNotSupportedError creates a NotSupportedError.
// Used when the requested type of object or operation is not supported.
func NewNotSupportedError(message string) *DOMException {
	return NewDOMException(message, "NotSupportedError")
}

// NewInvalidStateError creates an InvalidStateError.
// Used when an attempt is made to use an object that is not, or is no longer, usable.
func NewInvalidStateError(message string) *DOMException {
	return NewDOMException(message, "InvalidStateError")
}

// NewAbortError creates an AbortError.
// Used when an operation was aborted.
func NewAbortError(message string) *DOMException {
	return NewDOMException(message, "AbortError")
}

// NewNetworkError creates a NetworkError.
// Used when a network error occurs.
func NewNetworkError(message string) *DOMException {
	return NewDOMException(message, "NetworkError")
}

// NewSecurityError creates a SecurityError.
// Used when the operation is insecure.
func NewSecurityError(message string) *DOMException {
	return NewDOMException(message, "SecurityError")
}

// toJS converts the DOMException to a JavaScript object.
func (e *DOMException) toJS(vm *goja.Runtime) goja.Value {
	obj := vm.NewObject()
	obj.Set("name", e.Name())
	obj.Set("message", e.Message())
	obj.Set("code", e.Code())
	obj.Set("toString", func() string {
		return e.Error()
	})
	return obj
}
