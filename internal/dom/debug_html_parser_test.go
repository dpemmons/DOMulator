package dom

import (
	"fmt"
	"testing"
)

func TestHTMLParserDebug(t *testing.T) {
	doc := NewDocument()
	element := doc.CreateElement("div")

	// Test simple HTML parsing
	html := `<em>emphasis</em>`
	nodes := element.parseHTMLFragment(html)

	fmt.Printf("Input HTML: %s\n", html)
	fmt.Printf("Number of parsed nodes: %d\n", len(nodes))
	for i, node := range nodes {
		fmt.Printf("Node %d: Type=%d, Name=%s, Value=%s\n", i, node.NodeType(), node.NodeName(), node.NodeValue())
		if elem, ok := node.(*Element); ok {
			fmt.Printf("  Element tag: %s, children: %d\n", elem.TagName(), elem.ChildNodes().Length())
		}
	}

	// Test complex HTML
	html2 := `<div class="test" id="myid">content</div>`
	nodes2 := element.parseHTMLFragment(html2)

	fmt.Printf("\nInput HTML: %s\n", html2)
	fmt.Printf("Number of parsed nodes: %d\n", len(nodes2))
	for i, node := range nodes2 {
		fmt.Printf("Node %d: Type=%d, Name=%s, Value=%s\n", i, node.NodeType(), node.NodeName(), node.NodeValue())
		if elem, ok := node.(*Element); ok {
			fmt.Printf("  Element tag: %s, children: %d\n", elem.TagName(), elem.ChildNodes().Length())
		}
	}
}
