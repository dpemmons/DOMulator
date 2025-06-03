package parser

import (
	"io"

	"github.com/dpemmons/DOMulator/internal/dom"
	"golang.org/x/net/html"
)

// Parser is responsible for parsing HTML and building a DOM tree.
type Parser struct {
	doc *dom.Document
}

// NewParser creates a new Parser instance.
func NewParser() *Parser {
	return &Parser{}
}

// Parse parses the given HTML content from an io.Reader and returns a Document.
func (p *Parser) Parse(r io.Reader) (*dom.Document, error) {
	// Use the full html.Parse() function to get a complete parsed tree
	htmlNode, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	p.doc = dom.NewDocument()

	// Convert the html.Node tree to our DOM tree
	p.convertNode(htmlNode, p.doc)

	return p.doc, nil
}

// convertNode recursively converts an html.Node tree to our DOM tree
func (p *Parser) convertNode(htmlNode *html.Node, parent dom.Node) {
	switch htmlNode.Type {
	case html.DocumentNode:
		// Document node - just process children
		for child := htmlNode.FirstChild; child != nil; child = child.NextSibling {
			p.convertNode(child, parent)
		}

	case html.ElementNode:
		// Create DOM element
		elem := dom.NewElement(htmlNode.Data, p.doc)

		// Copy attributes
		for _, attr := range htmlNode.Attr {
			elem.SetAttribute(attr.Key, attr.Val)
		}

		// Append to parent
		parent.AppendChild(elem)

		// Process children
		for child := htmlNode.FirstChild; child != nil; child = child.NextSibling {
			p.convertNode(child, elem)
		}

	case html.TextNode:
		// Create text node
		if len(htmlNode.Data) > 0 {
			textNode := dom.NewText(htmlNode.Data, p.doc)
			parent.AppendChild(textNode)
		}

	case html.CommentNode:
		// Create comment node
		commentNode := dom.NewComment(htmlNode.Data, p.doc)
		parent.AppendChild(commentNode)

	case html.DoctypeNode:
		// Create doctype node
		doctypeNode := dom.NewDocumentType(htmlNode.Data, "", "", p.doc)
		parent.AppendChild(doctypeNode)

		// Note: html.RawNode and html.ErrorNode are not handled as they're not commonly needed
	}
}
