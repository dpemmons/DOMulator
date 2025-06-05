package parser

import (
	"strings"
	"testing"

	"github.com/dpemmons/DOMulator/internal/dom"
)

func TestParseBasicHTML(t *testing.T) {
	htmlContent := `<!DOCTYPE html><html><head><title>Test</title></head><body><h1>Hello</h1><p>World</p></body></html>`
	parser := NewParser()
	doc, err := parser.Parse(strings.NewReader(htmlContent))

	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}
	if doc == nil {
		t.Fatal("Document is nil")
	}

	// Verify DocumentType
	doctype := doc.ChildNodes().Item(0)
	if doctype.NodeType() != dom.DocumentTypeNode || doctype.NodeName() != "html" { // Doctype name is not uppercased by spec
		t.Errorf("Expected doctype 'html', got %v %v", doctype.NodeType(), doctype.NodeName())
	}

	// Verify html element
	htmlElem := doc.ChildNodes().Item(1)
	if htmlElem.NodeType() != dom.ElementNode || htmlElem.NodeName() != "HTML" {
		t.Errorf("Expected html element, got %v %v", htmlElem.NodeType(), htmlElem.NodeName())
	}

	// Verify head and body
	headElem := htmlElem.ChildNodes().Item(0)
	if headElem.NodeType() != dom.ElementNode || headElem.NodeName() != "HEAD" {
		t.Errorf("Expected head element, got %v %v", headElem.NodeType(), headElem.NodeName())
	}
	bodyElem := htmlElem.ChildNodes().Item(1)
	if bodyElem.NodeType() != dom.ElementNode || bodyElem.NodeName() != "BODY" {
		t.Errorf("Expected body element, got %v %v", bodyElem.NodeType(), bodyElem.NodeName())
	}

	// Verify title inside head
	titleElem := headElem.ChildNodes().Item(0)
	if titleElem.NodeType() != dom.ElementNode || titleElem.NodeName() != "TITLE" {
		t.Errorf("Expected title element, got %v %v", titleElem.NodeType(), titleElem.NodeName())
	}
	titleText := titleElem.ChildNodes().Item(0)
	if titleText.NodeType() != dom.TextNode || titleText.NodeValue() != "Test" {
		t.Errorf("Expected title text 'Test', got %v %q", titleText.NodeType(), titleText.NodeValue())
	}

	// Verify h1 and p inside body
	h1Elem := bodyElem.ChildNodes().Item(0)
	if h1Elem.NodeType() != dom.ElementNode || h1Elem.NodeName() != "H1" {
		t.Errorf("Expected h1 element, got %v %v", h1Elem.NodeType(), h1Elem.NodeName())
	}
	h1Text := h1Elem.ChildNodes().Item(0)
	if h1Text.NodeType() != dom.TextNode || h1Text.NodeValue() != "Hello" {
		t.Errorf("Expected h1 text 'Hello', got %v %q", h1Text.NodeType(), h1Text.NodeValue())
	}

	pElem := bodyElem.ChildNodes().Item(1)
	if pElem.NodeType() != dom.ElementNode || pElem.NodeName() != "P" {
		t.Errorf("Expected p element, got %v %v", pElem.NodeType(), pElem.NodeName())
	}
	pText := pElem.ChildNodes().Item(0)
	if pText.NodeType() != dom.TextNode || pText.NodeValue() != "World" {
		t.Errorf("Expected p text 'World', got %v %q", pText.NodeType(), pText.NodeValue())
	}
}

func TestParseWithAttributes(t *testing.T) {
	htmlContent := `<div id="main" class="container"><a href="/test">Link</a></div>`
	parser := NewParser()
	doc, err := parser.Parse(strings.NewReader(htmlContent))
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	// The parser now ensures html and body elements exist.
	// doc -> html -> body -> div
	htmlElem := doc.ChildNodes().Item(0)
	bodyElem := htmlElem.ChildNodes().Item(1)
	div := bodyElem.ChildNodes().Item(0).(*dom.Element)
	if div.NodeName() != "DIV" {
		t.Fatalf("Expected div, got %s", div.NodeName())
	}
	if div.GetAttribute("id") != "main" {
		t.Errorf("Expected id 'main', got %s", div.GetAttribute("id"))
	}
	if div.GetAttribute("class") != "container" {
		t.Errorf("Expected class 'container', got %s", div.GetAttribute("class"))
	}

	a := div.ChildNodes().Item(0).(*dom.Element)
	if a.NodeName() != "A" {
		t.Fatalf("Expected a, got %s", a.NodeName())
	}
	if a.GetAttribute("href") != "/test" {
		t.Errorf("Expected href '/test', got %s", a.GetAttribute("href"))
	}
	if a.ChildNodes().Item(0).NodeValue() != "Link" {
		t.Errorf("Expected text 'Link', got %s", a.ChildNodes().Item(0).NodeValue())
	}
}

func TestParseSelfClosingTags(t *testing.T) {
	htmlContent := `<div><img src="test.jpg"><br></div>`
	parser := NewParser()
	doc, err := parser.Parse(strings.NewReader(htmlContent))
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	htmlElem := doc.ChildNodes().Item(0)
	bodyElem := htmlElem.ChildNodes().Item(1)
	div := bodyElem.ChildNodes().Item(0)
	if div.NodeName() != "DIV" {
		t.Fatalf("Expected div, got %s", div.NodeName())
	}

	img := div.ChildNodes().Item(0).(*dom.Element)
	if img.NodeName() != "IMG" {
		t.Errorf("Expected img, got %s", img.NodeName())
	}
	if img.GetAttribute("src") != "test.jpg" {
		t.Errorf("Expected src 'test.jpg', got %s", img.GetAttribute("src"))
	}

	br := div.ChildNodes().Item(1)
	if br.NodeName() != "BR" {
		t.Errorf("Expected br, got %s", br.NodeName())
	}
}

func TestParseNestedElements(t *testing.T) {
	htmlContent := `<div><span><p>Text</p></span></div>`
	parser := NewParser()
	doc, err := parser.Parse(strings.NewReader(htmlContent))
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	htmlElem := doc.ChildNodes().Item(0)
	bodyElem := htmlElem.ChildNodes().Item(1)
	div := bodyElem.ChildNodes().Item(0)
	span := div.ChildNodes().Item(0)
	p := span.ChildNodes().Item(0)
	text := p.ChildNodes().Item(0)

	if div.NodeName() != "DIV" || span.NodeName() != "SPAN" || p.NodeName() != "P" || text.NodeValue() != "Text" {
		t.Errorf("Nested elements parsing failed. Got: %s > %s > %s > %s", div.NodeName(), span.NodeName(), p.NodeName(), text.NodeValue())
	}
}

func TestParseMismatchedTags(t *testing.T) {
	htmlContent := `<div><p>Text</div></p>` // Mismatched closing tags
	parser := NewParser()
	doc, err := parser.Parse(strings.NewReader(htmlContent))
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	htmlElem := doc.ChildNodes().Item(0)
	bodyElem := htmlElem.ChildNodes().Item(1)
	div := bodyElem.ChildNodes().Item(0)
	p := div.ChildNodes().Item(0)
	text := p.ChildNodes().Item(0)

	if div.NodeName() != "DIV" || p.NodeName() != "P" || text.NodeValue() != "Text" {
		t.Errorf("Mismatched tags parsing failed. Got: %s > %s > %s", div.NodeName(), p.NodeName(), text.NodeValue())
	}
	// The parser should gracefully handle the </p> after </div> by popping div and then p.
	// The current implementation will pop div, then p, then the stack will be empty.
	// The extra </p> will be ignored.
	if div.ChildNodes().Length() != 1 {
		t.Errorf("Expected div to have 1 child, got %d", div.ChildNodes().Length())
	}
}

func TestParseComments(t *testing.T) {
	htmlContent := `<div><!-- This is a comment --><span>Text</span></div>`
	parser := NewParser()
	doc, err := parser.Parse(strings.NewReader(htmlContent))
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	htmlElem := doc.ChildNodes().Item(0)
	bodyElem := htmlElem.ChildNodes().Item(1)
	div := bodyElem.ChildNodes().Item(0)
	comment := div.ChildNodes().Item(0)
	span := div.ChildNodes().Item(1)

	if comment.NodeType() != dom.CommentNode || comment.NodeValue() != " This is a comment " {
		t.Errorf("Expected comment node with value ' This is a comment ', got %v %q", comment.NodeType(), comment.NodeValue())
	}
	if span.NodeName() != "SPAN" {
		t.Errorf("Expected span, got %s", span.NodeName())
	}
}

func TestParseEmptyHTML(t *testing.T) {
	htmlContent := ``
	parser := NewParser()
	doc, err := parser.Parse(strings.NewReader(htmlContent))
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}
	// HTML5 parser should create proper document structure even for empty input
	if doc.ChildNodes().Length() != 1 {
		t.Errorf("Expected 1 child (html element), got %d children", doc.ChildNodes().Length())
	}

	htmlElem := doc.ChildNodes().Item(0)
	if htmlElem.NodeName() != "HTML" {
		t.Errorf("Expected html element, got %s", htmlElem.NodeName())
	}

	// Should have head and body
	if htmlElem.ChildNodes().Length() != 2 {
		t.Errorf("Expected html to have 2 children (head and body), got %d", htmlElem.ChildNodes().Length())
	}

	if htmlElem.ChildNodes().Item(0).NodeName() != "HEAD" {
		t.Errorf("Expected first child to be head, got %s", htmlElem.ChildNodes().Item(0).NodeName())
	}

	if htmlElem.ChildNodes().Item(1).NodeName() != "BODY" {
		t.Errorf("Expected second child to be body, got %s", htmlElem.ChildNodes().Item(1).NodeName())
	}
}

func TestParseFragment(t *testing.T) {
	htmlContent := `<span>Hello</span><p>World</p>`
	parser := NewParser()
	doc, err := parser.Parse(strings.NewReader(htmlContent))
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	// The tokenizer implicitly adds html and body.
	// doc -> html -> body -> span, p
	htmlElem := doc.ChildNodes().Item(0)
	bodyElem := htmlElem.ChildNodes().Item(1)
	if bodyElem.ChildNodes().Length() != 2 {
		t.Fatalf("Expected 2 children in body, got %d", bodyElem.ChildNodes().Length())
	}
	if bodyElem.ChildNodes().Item(0).NodeName() != "SPAN" {
		t.Errorf("Expected first child to be span, got %s", bodyElem.ChildNodes().Item(0).NodeName())
	}
	if bodyElem.ChildNodes().Item(1).NodeName() != "P" {
		t.Errorf("Expected second child to be p, got %s", bodyElem.ChildNodes().Item(1).NodeName())
	}
}
