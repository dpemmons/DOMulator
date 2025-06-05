package dom

import (
	"regexp"
	"strings"
)

// Namespace constants per WHATWG specifications
const (
	HTMLNamespace = "http://www.w3.org/1999/xhtml"
	SVGNamespace  = "http://www.w3.org/2000/svg"
)

// XML Name production validation regex (per XML 1.0 specification)
// NameStartChar ::= ":" | [A-Z] | "_" | [a-z] | [#xC0-#xD6] | [#xD8-#xF6] | [#xF8-#x2FF] | [#x370-#x37D] | [#x37F-#x1FFF] | [#x200C-#x200D] | [#x2070-#x218F] | [#x2C00-#x2FEF] | [#x3001-#xD7FF] | [#xF900-#xFDCF] | [#xFDF0-#xFFFD] | [#x10000-#xEFFFF]
// NameChar ::= NameStartChar | "-" | "." | [0-9] | #xB7 | [#x0300-#x036F] | [#x203F-#x2040]
// Name ::= NameStartChar (NameChar)*
var nameStartCharRegex = regexp.MustCompile(`^[a-zA-Z_:]`)
var nameCharRegex = regexp.MustCompile(`^[a-zA-Z_:0-9.\-]*$`)

// validateQualifiedName validates a qualified name per XML specification
func validateQualifiedName(qualifiedName string) error {
	// If qualifiedName does not match the Name production, throw an "InvalidCharacterError" DOMException
	if qualifiedName == "" {
		return NewInvalidCharacterError("qualifiedName cannot be empty")
	}

	// If qualifiedName does not match the QName production, throw a "NamespaceError" DOMException
	// QName production allows qualified names with namespaces (prefix:localName)
	parts := strings.Split(qualifiedName, ":")
	if len(parts) > 2 {
		return NewNamespaceError("qualifiedName contains more than one colon")
	}

	// Handle different cases
	if len(parts) == 1 {
		// Simple name without colon
		part := parts[0]
		if !nameStartCharRegex.MatchString(part) {
			return NewInvalidCharacterError("qualifiedName does not match Name production")
		}
		if !nameCharRegex.MatchString(part) {
			return NewInvalidCharacterError("qualifiedName does not match Name production")
		}
	} else if len(parts) == 2 {
		// Qualified name with colon
		prefix := parts[0]
		localName := parts[1]

		// Check for invalid cases
		if prefix == "" && localName == "" {
			return NewNamespaceError("qualifiedName contains empty prefix and local name")
		}
		if localName == "" {
			return NewNamespaceError("qualifiedName contains empty local name")
		}

		// Validate prefix if it exists
		if prefix != "" {
			if !nameStartCharRegex.MatchString(prefix) {
				return NewInvalidCharacterError("qualifiedName prefix does not match Name production")
			}
			if !nameCharRegex.MatchString(prefix) {
				return NewInvalidCharacterError("qualifiedName prefix does not match Name production")
			}
		}

		// Validate local name
		if !nameStartCharRegex.MatchString(localName) {
			return NewInvalidCharacterError("qualifiedName local name does not match Name production")
		}
		if !nameCharRegex.MatchString(localName) {
			return NewInvalidCharacterError("qualifiedName local name does not match Name production")
		}
	}

	return nil
}

// DOMImplementation provides a number of methods for performing operations
// that are independent of any particular instance of the document object model.
// Per specification, each DOMImplementation is associated with a document.
type DOMImplementation struct {
	document *Document // Associated document per specification
}

// NewDOMImplementation creates a new DOMImplementation instance associated with the given document
func NewDOMImplementation(document *Document) *DOMImplementation {
	return &DOMImplementation{
		document: document,
	}
}

// CreateDocumentType creates a DocumentType node
// Per WHATWG DOM Section 4.5.1: createDocumentType(qualifiedName, publicId, systemId)
func (impl *DOMImplementation) CreateDocumentType(qualifiedName, publicID, systemID string) (*DocumentType, error) {
	// Step 1: Validate qualifiedName
	if err := validateQualifiedName(qualifiedName); err != nil {
		return nil, err
	}

	// Step 2: Return a new doctype, with qualifiedName as its name, publicId as its public ID,
	// and systemId as its system ID, and with its node document set to the associated document of this
	doctype := NewDocumentType(qualifiedName, publicID, systemID, impl.document)

	return doctype, nil
}

// CreateDocument creates an XMLDocument node
// Per WHATWG DOM Section 4.5.1: createDocument(namespace, qualifiedName, doctype)
func (impl *DOMImplementation) CreateDocument(namespace, qualifiedName string, doctype *DocumentType) (*Document, error) {
	// Step 1: Let document be a new XMLDocument
	document := NewDocument()
	document.documentType = "xml" // Ensure this is an XML document

	// Step 2: Let element be null
	var element *Element

	// Step 3: If qualifiedName is not the empty string, then set element to the result of
	// running the internal createElementNS steps, given document, namespace, qualifiedName, and an empty dictionary
	if qualifiedName != "" {
		elem, err := document.CreateElementNS(namespace, qualifiedName)
		if err != nil {
			// This method throws the same exceptions as the createElementNS() method
			return nil, err
		}
		element = elem
	}

	// Step 4: If doctype is non-null, append doctype to document
	if doctype != nil {
		document.AppendChild(doctype)
	}

	// Step 5: If element is non-null, append element to document
	if element != nil {
		document.AppendChild(element)
	}

	// Step 6: document's origin is this's associated document's origin
	// TODO: Implement origin when security model is added

	// Step 7: document's content type is determined by namespace
	switch namespace {
	case HTMLNamespace:
		document.contentType = "application/xhtml+xml"
	case SVGNamespace:
		document.contentType = "image/svg+xml"
	default:
		document.contentType = "application/xml"
	}

	// Step 8: Return document
	return document, nil
}

// CreateHTMLDocument creates an HTML document
// Per WHATWG DOM Section 4.5.1: createHTMLDocument(optional DOMString title)
func (impl *DOMImplementation) CreateHTMLDocument(title ...string) *Document {
	// Step 1: Let doc be a new document that is an HTML document
	doc := NewDocument()

	// Step 2: Set doc's content type to "text/html"
	doc.contentType = "text/html"
	doc.documentType = "html"
	doc.mode = "no-quirks"

	// Step 3: Append a new doctype, with "html" as its name and with its node document set to doc, to doc
	doctype := NewDocumentType("html", "", "", doc)
	doc.AppendChild(doctype)

	// Step 4: Append the result of creating an element given doc, "html", and the HTML namespace, to doc
	html, _ := doc.CreateElementNS(HTMLNamespace, "html")
	doc.AppendChild(html)

	// Step 5: Append the result of creating an element given doc, "head", and the HTML namespace, to the html element created earlier
	head, _ := doc.CreateElementNS(HTMLNamespace, "head")
	html.AppendChild(head)

	// Step 6: If title is given
	if len(title) > 0 {
		// Append the result of creating an element given doc, "title", and the HTML namespace, to the head element created earlier
		titleElem, _ := doc.CreateElementNS(HTMLNamespace, "title")
		head.AppendChild(titleElem)

		// Append a new Text node, with its data set to title (which could be the empty string) and its node document set to doc, to the title element created earlier
		titleText := doc.CreateTextNode(title[0])
		titleElem.AppendChild(titleText)
	}

	// Step 7: Append the result of creating an element given doc, "body", and the HTML namespace, to the html element created earlier
	body, _ := doc.CreateElementNS(HTMLNamespace, "body")
	html.AppendChild(body)

	// Step 8: doc's origin is this's associated document's origin
	// TODO: Implement origin when security model is added

	// Step 9: Return doc
	return doc
}

// HasFeature always returns true (legacy method)
func (impl *DOMImplementation) HasFeature(feature, version string) bool {
	// Per specification, this method is legacy and should always return true
	return true
}
