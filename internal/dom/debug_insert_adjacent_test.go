package dom

import (
	"fmt"
	"testing"
)

func TestInsertAdjacentHTMLDebug(t *testing.T) {
	doc := NewDocument()

	// Create test structure like in the failing test
	container := doc.CreateElement("div")
	target := doc.CreateElement("p")
	target.SetTextContent("target")
	container.AppendChild(target)

	fmt.Printf("Initial setup:\n")
	fmt.Printf("Target children: %d\n", target.ChildNodes().Length())
	for i := 0; i < target.ChildNodes().Length(); i++ {
		child := target.ChildNodes().Item(i)
		fmt.Printf("  Child %d: Type=%d, Name=%s, Value='%s'\n", i, child.NodeType(), child.NodeName(), child.NodeValue())
	}

	// Test afterbegin position
	fmt.Printf("\nInserting '<em>emphasis</em>' with afterbegin...\n")
	err := target.InsertAdjacentHTML("afterbegin", "<em>emphasis</em>")
	if err != nil {
		t.Fatalf("InsertAdjacentHTML failed: %v", err)
	}

	fmt.Printf("After afterbegin:\n")
	fmt.Printf("Target children: %d\n", target.ChildNodes().Length())
	for i := 0; i < target.ChildNodes().Length(); i++ {
		child := target.ChildNodes().Item(i)
		fmt.Printf("  Child %d: Type=%d, Name=%s, Value='%s'\n", i, child.NodeType(), child.NodeName(), child.NodeValue())
		if elem, ok := child.(*Element); ok {
			fmt.Printf("    Element tag: %s, children: %d\n", elem.TagName(), elem.ChildNodes().Length())
		}
	}

	// Test beforeend position
	fmt.Printf("\nInserting '<strong>strong</strong>' with beforeend...\n")
	err = target.InsertAdjacentHTML("beforeend", "<strong>strong</strong>")
	if err != nil {
		t.Fatalf("InsertAdjacentHTML failed: %v", err)
	}

	fmt.Printf("After beforeend:\n")
	fmt.Printf("Target children: %d\n", target.ChildNodes().Length())
	for i := 0; i < target.ChildNodes().Length(); i++ {
		child := target.ChildNodes().Item(i)
		fmt.Printf("  Child %d: Type=%d, Name=%s, Value='%s'\n", i, child.NodeType(), child.NodeName(), child.NodeValue())
		if elem, ok := child.(*Element); ok {
			fmt.Printf("    Element tag: %s, children: %d\n", elem.TagName(), elem.ChildNodes().Length())
		}
	}
}
