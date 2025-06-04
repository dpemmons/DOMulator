package dom

// DOMImplementation provides a number of methods for performing operations
// that are independent of any particular instance of the document object model.
type DOMImplementation struct {
	// No state needed for basic implementation
}

// NewDOMImplementation creates a new DOMImplementation instance
func NewDOMImplementation() *DOMImplementation {
	return &DOMImplementation{}
}

// CreateDocumentType creates a DocumentType node
func (impl *DOMImplementation) CreateDocumentType(qualifiedName, publicID, systemID string) (*DocumentType, error) {
	// TODO: Add Name production validation
	if qualifiedName == "" {
		return nil, NewInvalidCharacterError("qualifiedName cannot be empty")
	}

	return NewDocumentType(qualifiedName, publicID, systemID, nil), nil
}

// CreateDocument creates a Document node
func (impl *DOMImplementation) CreateDocument(namespace, qualifiedName string, doctype *DocumentType) (*Document, error) {
	doc := NewDocument()

	// Set document properties
	if namespace != "" || qualifiedName != "" {
		doc.contentType = "application/xml"
		doc.documentType = "xml"
	}

	// Add doctype if provided
	if doctype != nil {
		doc.AppendChild(doctype)
	}

	// Create document element if qualified name provided
	if qualifiedName != "" {
		elem, err := doc.CreateElementNS(namespace, qualifiedName)
		if err != nil {
			return nil, err
		}
		doc.AppendChild(elem)
	}

	return doc, nil
}

// CreateHTMLDocument creates an HTML document
func (impl *DOMImplementation) CreateHTMLDocument(title string) *Document {
	doc := NewDocument()

	// Set HTML document properties
	doc.contentType = "text/html"
	doc.documentType = "html"
	doc.mode = "no-quirks"

	// Create DOCTYPE html
	doctype := NewDocumentType("html", "", "", doc)
	doc.AppendChild(doctype)

	// Create HTML structure
	html := doc.CreateElement("html")
	head := doc.CreateElement("head")
	body := doc.CreateElement("body")

	doc.AppendChild(html)
	html.AppendChild(head)
	html.AppendChild(body)

	// Add title if provided
	if title != "" {
		titleElem := doc.CreateElement("title")
		titleText := doc.CreateTextNode(title)
		titleElem.AppendChild(titleText)
		head.AppendChild(titleElem)
	}

	return doc
}

// HasFeature always returns true (legacy method)
func (impl *DOMImplementation) HasFeature(feature, version string) bool {
	// Per specification, this method is legacy and should always return true
	return true
}
