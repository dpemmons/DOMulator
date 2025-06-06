package js

import (
	"fmt"
	"testing"

	"github.com/dop251/goja"
	"github.com/dpemmons/DOMulator/internal/dom"
)

func TestInsertAdjacentHTMLJSDebug(t *testing.T) {
	doc := dom.NewDocument()
	vm := goja.New()
	bindings := NewDOMBindings(vm, doc)

	// Create test structure
	target := doc.CreateElement("p")
	target.SetTextContent("target")

	// Wrap element for JavaScript
	targetJS := bindings.WrapElement(target)
	vm.Set("target", targetJS)

	fmt.Printf("Initial setup:\n")
	fmt.Printf("Go DOM - Target children: %d\n", target.ChildNodes().Length())
	for i := 0; i < target.ChildNodes().Length(); i++ {
		child := target.ChildNodes().Item(i)
		fmt.Printf("  Child %d: Type=%d, Name=%s, Value='%s'\n", i, child.NodeType(), child.NodeName(), child.NodeValue())
	}

	// Check JavaScript side before operation
	result, err := vm.RunString(`target.childNodes.length`)
	if err != nil {
		t.Fatalf("Failed to access JS childNodes: %v", err)
	}
	fmt.Printf("JS - Target children: %d\n", result.ToInteger())

	// Test afterbegin position from JavaScript
	fmt.Printf("\nCalling insertAdjacentHTML from JavaScript...\n")
	_, err = vm.RunString(`target.insertAdjacentHTML('afterbegin', '<em>emphasis</em>')`)
	if err != nil {
		t.Fatalf("insertAdjacentHTML afterbegin failed: %v", err)
	}

	fmt.Printf("After insertAdjacentHTML:\n")
	fmt.Printf("Go DOM - Target children: %d\n", target.ChildNodes().Length())
	for i := 0; i < target.ChildNodes().Length(); i++ {
		child := target.ChildNodes().Item(i)
		fmt.Printf("  Child %d: Type=%d, Name=%s, Value='%s'\n", i, child.NodeType(), child.NodeName(), child.NodeValue())
		if elem, ok := child.(*dom.Element); ok {
			fmt.Printf("    Element tag: %s, children: %d\n", elem.TagName(), elem.ChildNodes().Length())
		}
	}

	// Check JavaScript side after operation
	result, err = vm.RunString(`target.childNodes.length`)
	if err != nil {
		t.Fatalf("Failed to access JS childNodes: %v", err)
	}
	fmt.Printf("JS - Target children: %d\n", result.ToInteger())

	// Check if first child is element from JavaScript
	result, err = vm.RunString(`target.childNodes[0] ? target.childNodes[0].nodeType : null`)
	if err != nil {
		t.Fatalf("Failed to access first child nodeType: %v", err)
	}
	fmt.Printf("JS - First child nodeType: %v\n", result.Export())

	// Check textContent from JavaScript
	result, err = vm.RunString(`target.textContent`)
	if err != nil {
		t.Fatalf("Failed to access textContent: %v", err)
	}
	fmt.Printf("JS - textContent: '%s'\n", result.String())
}
