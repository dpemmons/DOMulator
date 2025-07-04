package css

import (
	"testing"

	"github.com/dpemmons/DOMulator/internal/dom"
)

// Helper function to create a simple DOM structure for testing
func createTestDOM() *dom.Document {
	doc := dom.NewDocument()

	// <html>
	htmlElem := dom.NewElement("html", doc)
	doc.AppendChild(htmlElem)

	// <head>
	headElem := dom.NewElement("head", doc)
	htmlElem.AppendChild(headElem)
	titleElem := dom.NewElement("title", doc)
	headElem.AppendChild(titleElem)
	titleElem.AppendChild(dom.NewText("Test Title", doc))

	// <body>
	bodyElem := dom.NewElement("body", doc)
	htmlElem.AppendChild(bodyElem)

	// <div id="container" class="wrapper main">
	containerDiv := dom.NewElement("div", doc)
	containerDiv.SetAttribute("id", "container")
	containerDiv.SetAttribute("class", "wrapper main")
	bodyElem.AppendChild(containerDiv)

	// <p class="text-content">Paragraph 1</p>
	p1 := dom.NewElement("p", doc)
	p1.SetAttribute("class", "text-content")
	p1.AppendChild(dom.NewText("Paragraph 1", doc))
	containerDiv.AppendChild(p1)

	// <div class="item">
	itemDiv1 := dom.NewElement("div", doc)
	itemDiv1.SetAttribute("class", "item")
	containerDiv.AppendChild(itemDiv1)

	// <span id="span1" class="highlight">Span 1</span>
	span1 := dom.NewElement("span", doc)
	span1.SetAttribute("id", "span1")
	span1.SetAttribute("class", "highlight")
	span1.AppendChild(dom.NewText("Span 1", doc))
	itemDiv1.AppendChild(span1)

	// <p class="text-content">Paragraph 2</p>
	p2 := dom.NewElement("p", doc)
	p2.SetAttribute("class", "text-content")
	p2.AppendChild(dom.NewText("Paragraph 2", doc))
	containerDiv.AppendChild(p2)

	// <div class="item active">
	itemDiv2 := dom.NewElement("div", doc)
	itemDiv2.SetAttribute("class", "item active")
	containerDiv.AppendChild(itemDiv2)

	// <a href="#" class="link">Link</a>
	// The link element is created and appended as part of setting up the test DOM.
	// Its variable is not directly used in the test assertions, so we assign it to a blank identifier.
	_ = func() *dom.Element {
		linkElem := dom.NewElement("a", doc)
		linkElem.SetAttribute("href", "#")
		linkElem.SetAttribute("class", "link")
		linkElem.AppendChild(dom.NewText("Link", doc))
		itemDiv2.AppendChild(linkElem)
		return linkElem
	}()

	return doc
}

func TestCompileSelector(t *testing.T) {
	tests := []struct {
		name     string
		selector string
		expected SimpleSelector // Just test the first part for simple selectors
		hasError bool
	}{
		{"Tag Selector", "div", SimpleSelector{tag: "div"}, false},
		{"ID Selector", "#container", SimpleSelector{id: "container"}, false},
		{"Class Selector", ".item", SimpleSelector{classes: []string{"item"}}, false},
		{"Tag and Class", "p.text-content", SimpleSelector{tag: "p", classes: []string{"text-content"}}, false},
		{"Tag and ID", "div#container", SimpleSelector{tag: "div", id: "container"}, false},
		{"Multiple Classes", ".wrapper.main", SimpleSelector{classes: []string{"wrapper", "main"}}, false},
		{"Complex Selector", "div#container.wrapper.main", SimpleSelector{tag: "div", id: "container", classes: []string{"wrapper", "main"}}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := CompileSelector(tt.selector)
			if (err != nil) != tt.hasError {
				t.Errorf("CompileSelector() error = %v, hasError %v", err, tt.hasError)
				return
			}
			if !tt.hasError {
				// For simple selectors, there should be exactly one part
				if len(s.parts) != 1 {
					t.Errorf("Expected 1 selector part, got %d", len(s.parts))
					return
				}
				part := s.parts[0]
				if part.tag != tt.expected.tag {
					t.Errorf("Expected tag %q, got %q", tt.expected.tag, part.tag)
				}
				if part.id != tt.expected.id {
					t.Errorf("Expected id %q, got %q", tt.expected.id, part.id)
				}
				if len(part.classes) != len(tt.expected.classes) {
					t.Errorf("Expected %d classes, got %d", len(tt.expected.classes), len(part.classes))
				} else {
					for i, class := range part.classes {
						if class != tt.expected.classes[i] {
							t.Errorf("Expected class %q, got %q at index %d", tt.expected.classes[i], class, i)
						}
					}
				}
			}
		})
	}
}

func TestDescendantSelector(t *testing.T) {
	tests := []struct {
		name      string
		selector  string
		partCount int
	}{
		{"Simple Selector", "div", 1},
		{"Two-part Descendant", "div p", 2},
		{"Three-part Descendant", "body div p", 3},
		{"Complex Descendant", "#container .item", 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := CompileSelector(tt.selector)
			if err != nil {
				t.Errorf("CompileSelector() error = %v", err)
				return
			}
			if len(s.parts) != tt.partCount {
				t.Errorf("Expected %d parts, got %d", tt.partCount, len(s.parts))
			}
		})
	}
}

func TestMatches(t *testing.T) {
	doc := createTestDOM()
	containerDiv := doc.ChildNodes().Item(0).(*dom.Element).ChildNodes().Item(1).(*dom.Element).ChildNodes().Item(0).(*dom.Element) // body -> div#container
	p1 := containerDiv.ChildNodes().Item(0)                                                                                         // p.text-content
	span1 := containerDiv.ChildNodes().Item(1).(*dom.Element).ChildNodes().Item(0)                                                  // div.item -> span#span1.highlight
	itemDiv2 := containerDiv.ChildNodes().Item(3)                                                                                   // div.item.active

	tests := []struct {
		name     string
		node     dom.Node
		selector string
		expected bool
	}{
		{"Matches Tag", containerDiv, "div", true},
		{"Does Not Match Tag", containerDiv, "span", false},
		{"Matches ID", containerDiv, "#container", true},
		{"Does Not Match ID", containerDiv, "#nonexistent", false},
		{"Matches Class", p1, ".text-content", true},
		{"Does Not Match Class", p1, ".nonexistent", false},
		{"Matches Tag and Class", p1, "p.text-content", true},
		{"Does Not Match Tag and Class", p1, "div.text-content", false},
		{"Matches Tag and ID", containerDiv, "div#container", true},
		{"Matches Multiple Classes", itemDiv2, ".item.active", true},
		{"Matches One of Multiple Classes", itemDiv2, ".item", true},
		{"Does Not Match All Multiple Classes", itemDiv2, ".item.nonexistent", false},
		{"Matches ID and Class", span1, "#span1.highlight", true},
		{"Matches Tag, ID, and Class", span1, "span#span1.highlight", true},
		{"Non-Element Node", doc, "html", false}, // Document node should not match element selector
		// Add descendant selector tests
		{"Descendant: div p matches", p1, "div p", true},
		{"Descendant: body p matches", p1, "body p", true},
		{"Descendant: body div p matches", p1, "body div p", true},
		{"Descendant: span p does not match", p1, "span p", false},
		{"Descendant: #container .item matches", containerDiv.ChildNodes().Item(1), "#container .item", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sel, err := CompileSelector(tt.selector)
			if err != nil {
				t.Fatalf("CompileSelector failed: %v", err)
			}
			if sel.MatchesDescendant(tt.node) != tt.expected {
				t.Errorf("Selector %q on node %q: Expected %v, got %v", tt.selector, tt.node.NodeName(), tt.expected, sel.MatchesDescendant(tt.node))
			}
		})
	}
}

func TestQuerySelector(t *testing.T) {
	doc := createTestDOM()

	tests := []struct {
		name     string
		root     dom.Node
		selector string
		expected string // NodeName of the expected node
	}{
		{"Select HTML", doc, "html", "HTML"},
		{"Select Body", doc, "body", "BODY"},
		{"Select Container by ID", doc, "#container", "DIV"},
		{"Select First P", doc, "p", "P"},
		{"Select Highlight Span", doc, "span.highlight", "SPAN"},
		{"Select Active Item Div", doc, "div.active", "DIV"},
		{"Select Link", doc, "a.link", "A"},
		{"Non-existent selector", doc, "h1", ""},
		{"Select from Element Root", doc.ChildNodes().Item(0).(*dom.Element).ChildNodes().Item(1).(*dom.Element).ChildNodes().Item(0).(*dom.Element), "p.text-content", "P"}, // From #container, select p
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			foundNode := QuerySelector(tt.root, tt.selector)
			if tt.expected == "" {
				if foundNode != nil {
					t.Errorf("Expected no node for selector %q, got %q", tt.selector, foundNode.NodeName())
				}
			} else {
				if foundNode == nil {
					t.Errorf("Expected node %q for selector %q, got nil", tt.expected, tt.selector)
				} else if foundNode.NodeName() != tt.expected {
					t.Errorf("Expected node %q for selector %q, got %q", tt.expected, tt.selector, foundNode.NodeName())
				}
			}
		})
	}
}

func TestQuerySelectorAll(t *testing.T) {
	doc := createTestDOM()

	tests := []struct {
		name     string
		root     dom.Node
		selector string
		expected []string // NodeNames of the expected nodes
	}{
		{"Select All Divs", doc, "div", []string{"DIV", "DIV", "DIV"}},
		{"Select All P", doc, "p", []string{"P", "P"}},
		{"Select All Items", doc, ".item", []string{"DIV", "DIV"}},
		{"Select All Text Content Paragraphs", doc, "p.text-content", []string{"P", "P"}},
		{"Select All Spans", doc, "span", []string{"SPAN"}},
		{"Non-existent selector", doc, "h1", []string{}},
		{"Select All from Element Root", doc.ChildNodes().Item(0).(*dom.Element).ChildNodes().Item(1).(*dom.Element).ChildNodes().Item(0).(*dom.Element), "p", []string{"P", "P"}}, // From #container, select all p
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			foundNodes := QuerySelectorAll(tt.root, tt.selector)
			foundSlice := foundNodes.ToSlice()
			if foundNodes.Length() != len(tt.expected) {
				t.Errorf("Selector %q: Expected %d nodes, got %d", tt.selector, len(tt.expected), foundNodes.Length())
				return
			}
			for i, node := range foundSlice {
				if node.NodeName() != tt.expected[i] {
					t.Errorf("Selector %q: Expected node %q at index %d, got %q", tt.selector, tt.expected[i], i, node.NodeName())
				}
			}
		})
	}
}
