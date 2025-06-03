package parser

import (
	"io"

	"github.com/dpemmons/DOMulator/internal/dom"
	"golang.org/x/net/html"
)

// Parser is responsible for parsing HTML and building a DOM tree.
type Parser struct {
	tokenizer *html.Tokenizer
	doc       *dom.Document
	errors    []error
}

// NewParser creates a new Parser instance.
func NewParser() *Parser {
	return &Parser{}
}

// Parse parses the given HTML content from an io.Reader and returns a Document.
func (p *Parser) Parse(r io.Reader) (*dom.Document, error) {
	p.tokenizer = html.NewTokenizer(r)
	p.doc = dom.NewDocument()

	// Build DOM tree
	if err := p.parseDocument(); err != nil {
		return nil, err
	}

	// TODO: Post-process (e.g., assign IDs, compute styles)
	// p.assignIDs()
	// p.computeStyles()

	return p.doc, nil
}

// parseDocument recursively parses tokens and builds the DOM tree.
func (p *Parser) parseDocument() error {
	// For now, a very basic implementation that just adds a body element
	// and any text content. This will be expanded significantly.
	body := dom.NewElement("body", p.doc)
	p.doc.AppendChild(body)

	for {
		tt := p.tokenizer.Next()
		switch tt {
		case html.ErrorToken:
			err := p.tokenizer.Err()
			if err == io.EOF {
				return nil // End of document
			}
			p.errors = append(p.errors, err)
			return err // Or continue parsing, depending on error handling strategy
		case html.TextToken:
			text := string(p.tokenizer.Text())
			if len(text) > 0 {
				body.AppendChild(dom.NewText(text, p.doc))
			}
		case html.StartTagToken:
			// TODO: Handle elements and their attributes
		case html.EndTagToken:
			// TODO: Handle closing tags
		case html.SelfClosingTagToken:
			// TODO: Handle self-closing tags
		case html.CommentToken:
			// TODO: Handle comments
		case html.DoctypeToken:
			// TODO: Handle doctype
		}
	}
}
