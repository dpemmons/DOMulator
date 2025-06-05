package parser

import (
	"io"
	"strings"

	"github.com/dpemmons/DOMulator/internal/dom"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
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
		tagName := strings.ToLower(htmlNode.Data)
		elem := dom.NewElement(tagName, p.doc)

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

// ParseFragment parses an HTML fragment and returns the parsed nodes
// This is used for insertAdjacentHTML and similar operations
func (p *Parser) ParseFragment(htmlFragment string, doc *dom.Document) ([]dom.Node, error) {
	// Create a temporary container to parse the fragment
	containerHTML := "<div>" + htmlFragment + "</div>"

	// Parse the container
	htmlNodes, err := html.ParseFragment(strings.NewReader(containerHTML), &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Div,
		Data:     "div",
	})
	if err != nil {
		return nil, err
	}

	var nodes []dom.Node

	// Extract the children of the container (skip the div wrapper)
	for _, htmlNode := range htmlNodes {
		if htmlNode.Type == html.ElementNode && htmlNode.Data == "div" {
			// Process the children of the div
			for child := htmlNode.FirstChild; child != nil; child = child.NextSibling {
				if node := p.convertNodeToDOM(child, doc); node != nil {
					nodes = append(nodes, node)
				}
			}
		} else {
			// Process the node directly
			if node := p.convertNodeToDOM(htmlNode, doc); node != nil {
				nodes = append(nodes, node)
			}
		}
	}

	return nodes, nil
}

// convertNodeToDOM converts a single html.Node to a DOM node
func (p *Parser) convertNodeToDOM(htmlNode *html.Node, doc *dom.Document) dom.Node {
	switch htmlNode.Type {
	case html.ElementNode:
		// Create DOM element
		tagName := strings.ToLower(htmlNode.Data)
		elem := dom.NewElement(tagName, doc)

		// Copy attributes
		for _, attr := range htmlNode.Attr {
			elem.SetAttribute(attr.Key, attr.Val)
		}

		// Process children
		for child := htmlNode.FirstChild; child != nil; child = child.NextSibling {
			if childNode := p.convertNodeToDOM(child, doc); childNode != nil {
				elem.AppendChild(childNode)
			}
		}

		return elem

	case html.TextNode:
		// Create text node
		if len(htmlNode.Data) > 0 {
			return dom.NewText(htmlNode.Data, doc)
		}

	case html.CommentNode:
		// Create comment node
		return dom.NewComment(htmlNode.Data, doc)

	case html.DoctypeNode:
		// Create doctype node
		return dom.NewDocumentType(htmlNode.Data, "", "", doc)
	}

	return nil
}
