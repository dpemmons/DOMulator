package dom

import (
	"testing"
)

func TestNodeList_BasicSpecCompliance(t *testing.T) {
	doc := NewDocument()
	parent := NewElement("div", doc)

	// Create test nodes
	text1 := NewText("text1", doc)
	elem1 := NewElement("span", doc)
	text2 := NewText("text2", doc)
	comment1 := NewComment("comment", doc)
	elem2 := NewElement("p", doc)

	// Add nodes to parent
	parent.AppendChild(text1)
	parent.AppendChild(elem1)
	parent.AppendChild(text2)
	parent.AppendChild(comment1)
	parent.AppendChild(elem2)

	nodeList := parent.ChildNodes()

	t.Run("Length attribute compliance", func(t *testing.T) {
		// The length attribute must return the number of nodes represented by the collection
		if nodeList.Length() != 5 {
			t.Errorf("Expected length 5, got %d", nodeList.Length())
		}
	})

	t.Run("Item method compliance", func(t *testing.T) {
		// The item(index) method must return the indexth node in the collection
		if nodeList.Item(0) != text1 {
			t.Errorf("Expected text1 at index 0")
		}
		if nodeList.Item(1) != elem1 {
			t.Errorf("Expected elem1 at index 1")
		}
		if nodeList.Item(2) != text2 {
			t.Errorf("Expected text2 at index 2")
		}
		if nodeList.Item(3) != comment1 {
			t.Errorf("Expected comment1 at index 3")
		}
		if nodeList.Item(4) != elem2 {
			t.Errorf("Expected elem2 at index 4")
		}

		// If there is no indexth node in the collection, then the method must return null
		if nodeList.Item(5) != nil {
			t.Errorf("Expected nil for out-of-bounds index 5")
		}
		if nodeList.Item(10) != nil {
			t.Errorf("Expected nil for out-of-bounds index 10")
		}
	})

	t.Run("Tree order compliance", func(t *testing.T) {
		// The nodes are sorted in tree order
		// This is already tested above by checking that items are returned in insertion order
		// which matches tree order for a simple linear structure
	})

	t.Run("Supported property indices", func(t *testing.T) {
		// The object's supported property indices are the numbers in the range zero to one less than the number of nodes
		// For a collection with 5 nodes, valid indices are 0, 1, 2, 3, 4
		for i := 0; i < 5; i++ {
			if nodeList.Item(i) == nil {
				t.Errorf("Expected non-nil node at valid index %d", i)
			}
		}

		// Invalid indices should return nil
		invalidIndices := []int{5, 6, 100, -1}
		for _, idx := range invalidIndices {
			if nodeList.Item(idx) != nil {
				t.Errorf("Expected nil for invalid index %d", idx)
			}
		}
	})
}

func TestNodeList_EmptyCollection(t *testing.T) {
	doc := NewDocument()
	parent := NewElement("div", doc)
	nodeList := parent.ChildNodes()

	t.Run("Empty collection length", func(t *testing.T) {
		if nodeList.Length() != 0 {
			t.Errorf("Expected length 0 for empty collection, got %d", nodeList.Length())
		}
	})

	t.Run("Empty collection item access", func(t *testing.T) {
		// If there are no elements, then there are no supported property indices
		if nodeList.Item(0) != nil {
			t.Errorf("Expected nil for index 0 in empty collection")
		}
	})
}

func TestNodeList_LiveCollection(t *testing.T) {
	doc := NewDocument()
	parent := NewElement("div", doc)
	elem1 := NewElement("span", doc)
	elem2 := NewElement("p", doc)

	parent.AppendChild(elem1)
	nodeList := parent.ChildNodes()

	t.Run("Initial state", func(t *testing.T) {
		if nodeList.Length() != 1 {
			t.Errorf("Expected length 1, got %d", nodeList.Length())
		}
		if nodeList.Item(0) != elem1 {
			t.Errorf("Expected elem1 at index 0")
		}
	})

	t.Run("Collection updates when DOM changes", func(t *testing.T) {
		// Add another child
		parent.AppendChild(elem2)

		if nodeList.Length() != 2 {
			t.Errorf("Expected length 2 after adding child, got %d", nodeList.Length())
		}
		if nodeList.Item(1) != elem2 {
			t.Errorf("Expected elem2 at index 1")
		}
	})

	t.Run("Collection updates when child removed", func(t *testing.T) {
		// Remove first child
		parent.RemoveChild(elem1)

		if nodeList.Length() != 1 {
			t.Errorf("Expected length 1 after removing child, got %d", nodeList.Length())
		}
		if nodeList.Item(0) != elem2 {
			t.Errorf("Expected elem2 at index 0 after removal")
		}
		if nodeList.Item(1) != nil {
			t.Errorf("Expected nil at index 1 after removal")
		}
	})
}

func TestNodeList_AllNodeTypes(t *testing.T) {
	doc := NewDocument()
	parent := NewElement("div", doc)

	// Test with all types of nodes that can be children
	text := NewText("text", doc)
	element := NewElement("span", doc)
	comment := NewComment("comment", doc)
	cdata := NewCDATASection("cdata", doc)
	pi := NewProcessingInstruction("target", "data", doc)

	parent.AppendChild(text)
	parent.AppendChild(element)
	parent.AppendChild(comment)
	parent.AppendChild(cdata)
	parent.AppendChild(pi)

	nodeList := parent.ChildNodes()

	t.Run("All node types included", func(t *testing.T) {
		if nodeList.Length() != 5 {
			t.Errorf("Expected length 5, got %d", nodeList.Length())
		}

		// Verify each node type
		if nodeList.Item(0).NodeType() != TextNode {
			t.Errorf("Expected TextNode at index 0")
		}
		if nodeList.Item(1).NodeType() != ElementNode {
			t.Errorf("Expected ElementNode at index 1")
		}
		if nodeList.Item(2).NodeType() != CommentNode {
			t.Errorf("Expected CommentNode at index 2")
		}
		if nodeList.Item(3).NodeType() != CDataSectionNode {
			t.Errorf("Expected CDataSectionNode at index 3")
		}
		if nodeList.Item(4).NodeType() != ProcessingInstructionNode {
			t.Errorf("Expected ProcessingInstructionNode at index 4")
		}
	})
}

func TestNodeList_TreeOrderCompliance(t *testing.T) {
	doc := NewDocument()
	parent := NewElement("div", doc)

	// Create nested structure to test tree order
	child1 := NewElement("child1", doc)
	child2 := NewElement("child2", doc)
	grandchild1 := NewElement("grandchild1", doc)
	grandchild2 := NewElement("grandchild2", doc)

	child1.AppendChild(grandchild1)
	child1.AppendChild(grandchild2)
	parent.AppendChild(child1)
	parent.AppendChild(child2)

	nodeList := parent.ChildNodes()

	t.Run("Direct children in tree order", func(t *testing.T) {
		if nodeList.Length() != 2 {
			t.Errorf("Expected length 2, got %d", nodeList.Length())
		}
		if nodeList.Item(0) != child1 {
			t.Errorf("Expected child1 at index 0")
		}
		if nodeList.Item(1) != child2 {
			t.Errorf("Expected child2 at index 1")
		}
	})

	t.Run("Nested children have their own NodeList", func(t *testing.T) {
		child1NodeList := child1.ChildNodes()
		if child1NodeList.Length() != 2 {
			t.Errorf("Expected child1 to have 2 children, got %d", child1NodeList.Length())
		}
		if child1NodeList.Item(0) != grandchild1 {
			t.Errorf("Expected grandchild1 at index 0 of child1")
		}
		if child1NodeList.Item(1) != grandchild2 {
			t.Errorf("Expected grandchild2 at index 1 of child1")
		}
	})
}

func TestNodeList_SpecificationMethods(t *testing.T) {
	doc := NewDocument()
	parent := NewElement("div", doc)
	elem1 := NewElement("span", doc)
	elem2 := NewElement("p", doc)

	parent.AppendChild(elem1)
	parent.AppendChild(elem2)
	nodeList := parent.ChildNodes()

	t.Run("Length method returns correct count", func(t *testing.T) {
		// Returns the number of nodes in the collection
		if nodeList.Length() != 2 {
			t.Errorf("Expected length 2, got %d", nodeList.Length())
		}
	})

	t.Run("Item method returns correct node", func(t *testing.T) {
		// Returns the node with index from the collection
		if nodeList.Item(0) != elem1 {
			t.Errorf("Expected elem1 at index 0")
		}
		if nodeList.Item(1) != elem2 {
			t.Errorf("Expected elem2 at index 1")
		}
	})

	t.Run("Item method with invalid index", func(t *testing.T) {
		// Returns null if there is no indexth node
		if nodeList.Item(2) != nil {
			t.Errorf("Expected nil for index 2")
		}
		if nodeList.Item(-1) != nil {
			t.Errorf("Expected nil for negative index")
		}
	})
}

func TestNodeList_DocumentContext(t *testing.T) {
	doc := NewDocument()
	html := NewElement("html", doc)
	head := NewElement("head", doc)
	body := NewElement("body", doc)

	html.AppendChild(head)
	html.AppendChild(body)
	doc.AppendChild(html)

	t.Run("Document children NodeList", func(t *testing.T) {
		docChildren := doc.ChildNodes()
		if docChildren.Length() != 1 {
			t.Errorf("Expected document to have 1 child, got %d", docChildren.Length())
		}
		if docChildren.Item(0) != html {
			t.Errorf("Expected html element as document child")
		}
	})

	t.Run("HTML element children NodeList", func(t *testing.T) {
		htmlChildren := html.ChildNodes()
		if htmlChildren.Length() != 2 {
			t.Errorf("Expected html to have 2 children, got %d", htmlChildren.Length())
		}
		if htmlChildren.Item(0) != head {
			t.Errorf("Expected head element at index 0")
		}
		if htmlChildren.Item(1) != body {
			t.Errorf("Expected body element at index 1")
		}
	})
}

func TestNodeList_ModificationDuringIteration(t *testing.T) {
	doc := NewDocument()
	parent := NewElement("div", doc)

	// Add initial children
	for i := 0; i < 3; i++ {
		elem := NewElement("span", doc)
		parent.AppendChild(elem)
	}

	nodeList := parent.ChildNodes()

	t.Run("NodeList reflects modifications", func(t *testing.T) {
		initialLength := nodeList.Length()
		if initialLength != 3 {
			t.Errorf("Expected initial length 3, got %d", initialLength)
		}

		// Add a child during "iteration"
		newElem := NewElement("p", doc)
		parent.AppendChild(newElem)

		// NodeList should reflect the change
		if nodeList.Length() != 4 {
			t.Errorf("Expected length 4 after addition, got %d", nodeList.Length())
		}
		if nodeList.Item(3) != newElem {
			t.Errorf("Expected new element at index 3")
		}
	})
}
