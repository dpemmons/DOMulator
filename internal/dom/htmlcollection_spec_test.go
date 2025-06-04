package dom

import (
	"testing"
)

func TestHTMLCollection_SpecCompliance_Length(t *testing.T) {
	doc := NewDocument()
	parent := NewElement("div", doc)

	// Create test elements
	elem1 := NewElement("span", doc)
	elem2 := NewElement("p", doc)
	elem3 := NewElement("div", doc)
	text1 := NewText("text", doc) // Should be ignored by HTMLCollection

	parent.AppendChild(elem1)
	parent.AppendChild(text1) // Text nodes should not be in HTMLCollection
	parent.AppendChild(elem2)
	parent.AppendChild(elem3)

	collection := parent.Children()

	t.Run("Length returns number of elements", func(t *testing.T) {
		// Returns the number of elements in the collection
		if collection.Length() != 3 {
			t.Errorf("Expected length 3 (elements only), got %d", collection.Length())
		}
	})

	t.Run("Only elements are included", func(t *testing.T) {
		// HTMLCollection should only contain Element nodes, not Text nodes
		for i := 0; i < collection.Length(); i++ {
			item := collection.Item(i)
			if item == nil {
				t.Errorf("Expected non-nil item at index %d", i)
				continue
			}
			if item.NodeType() != ElementNode {
				t.Errorf("Expected ElementNode at index %d, got %d", i, item.NodeType())
			}
		}
	})
}

func TestHTMLCollection_SpecCompliance_Item(t *testing.T) {
	doc := NewDocument()
	parent := NewElement("div", doc)

	// Create test elements in specific order
	span := NewElement("span", doc)
	p := NewElement("p", doc)
	div := NewElement("div", doc)

	parent.AppendChild(span)
	parent.AppendChild(p)
	parent.AppendChild(div)

	collection := parent.Children()

	t.Run("Item method returns elements in tree order", func(t *testing.T) {
		// The elements are sorted in tree order
		if collection.Item(0) != span {
			t.Errorf("Expected span at index 0")
		}
		if collection.Item(1) != p {
			t.Errorf("Expected p at index 1")
		}
		if collection.Item(2) != div {
			t.Errorf("Expected div at index 2")
		}
	})

	t.Run("Item method returns nil for invalid indices", func(t *testing.T) {
		// If there is no indexth element in the collection, then the method must return null
		if collection.Item(3) != nil {
			t.Errorf("Expected nil for out-of-bounds index 3")
		}
		if collection.Item(10) != nil {
			t.Errorf("Expected nil for out-of-bounds index 10")
		}
		if collection.Item(-1) != nil {
			t.Errorf("Expected nil for negative index -1")
		}
	})

	t.Run("Supported property indices", func(t *testing.T) {
		// The object's supported property indices are the numbers in the range zero to one less than the number of elements
		for i := 0; i < 3; i++ {
			if collection.Item(i) == nil {
				t.Errorf("Expected non-nil element at valid index %d", i)
			}
		}
	})
}

func TestHTMLCollection_SpecCompliance_NamedItem(t *testing.T) {
	doc := NewDocument()
	parent := NewElement("div", doc)

	// Create elements with IDs and names
	elemWithId := NewElement("span", doc)
	elemWithId.SetAttribute("id", "myid")

	elemWithName := NewElement("input", doc)
	elemWithName.SetAttribute("name", "myname")

	elemWithBoth := NewElement("form", doc)
	elemWithBoth.SetAttribute("id", "formid")
	elemWithBoth.SetAttribute("name", "formname")

	elemPlain := NewElement("p", doc)

	parent.AppendChild(elemWithId)
	parent.AppendChild(elemWithName)
	parent.AppendChild(elemWithBoth)
	parent.AppendChild(elemPlain)

	collection := parent.Children()

	t.Run("NamedItem returns element by ID", func(t *testing.T) {
		// Returns the first element with ID which is key
		result := collection.NamedItem("myid")
		if result != elemWithId {
			t.Errorf("Expected elemWithId for ID 'myid'")
		}
	})

	t.Run("NamedItem returns element by name attribute", func(t *testing.T) {
		// Returns the first element with name attribute whose value is key
		result := collection.NamedItem("myname")
		if result != elemWithName {
			t.Errorf("Expected elemWithName for name 'myname'")
		}
	})

	t.Run("NamedItem prefers ID over name", func(t *testing.T) {
		// When both ID and name exist, ID should take precedence
		result := collection.NamedItem("formid")
		if result != elemWithBoth {
			t.Errorf("Expected elemWithBoth for ID 'formid'")
		}

		// But name should also work for the same element
		result = collection.NamedItem("formname")
		if result != elemWithBoth {
			t.Errorf("Expected elemWithBoth for name 'formname'")
		}
	})

	t.Run("NamedItem returns nil for empty string", func(t *testing.T) {
		// If key is the empty string, return null
		result := collection.NamedItem("")
		if result != nil {
			t.Errorf("Expected nil for empty string key")
		}
	})

	t.Run("NamedItem returns nil for non-existent key", func(t *testing.T) {
		// Return null if there is no such element
		result := collection.NamedItem("nonexistent")
		if result != nil {
			t.Errorf("Expected nil for non-existent key 'nonexistent'")
		}
	})
}

func TestHTMLCollection_SpecCompliance_TreeOrder(t *testing.T) {
	doc := NewDocument()
	parent := NewElement("div", doc)

	// Create nested structure to test tree order
	child1 := NewElement("child1", doc)
	child2 := NewElement("child2", doc)
	child3 := NewElement("child3", doc)

	// Add some text nodes between elements (should be ignored)
	parent.AppendChild(NewText("text1", doc))
	parent.AppendChild(child1)
	parent.AppendChild(NewText("text2", doc))
	parent.AppendChild(child2)
	parent.AppendChild(NewComment("comment", doc))
	parent.AppendChild(child3)
	parent.AppendChild(NewText("text3", doc))

	collection := parent.Children()

	t.Run("Elements returned in tree order ignoring non-elements", func(t *testing.T) {
		if collection.Length() != 3 {
			t.Errorf("Expected 3 elements, got %d", collection.Length())
		}

		// Should return elements in tree order, skipping non-element nodes
		if collection.Item(0) != child1 {
			t.Errorf("Expected child1 at index 0")
		}
		if collection.Item(1) != child2 {
			t.Errorf("Expected child2 at index 1")
		}
		if collection.Item(2) != child3 {
			t.Errorf("Expected child3 at index 2")
		}
	})
}

func TestHTMLCollection_SpecCompliance_LiveCollection(t *testing.T) {
	doc := NewDocument()
	parent := NewElement("div", doc)

	elem1 := NewElement("span", doc)
	elem2 := NewElement("p", doc)

	parent.AppendChild(elem1)
	collection := parent.Children()

	t.Run("Collection is live - reflects DOM changes", func(t *testing.T) {
		// Initial state
		if collection.Length() != 1 {
			t.Errorf("Expected initial length 1, got %d", collection.Length())
		}

		// Add another element
		parent.AppendChild(elem2)
		if collection.Length() != 2 {
			t.Errorf("Expected length 2 after adding element, got %d", collection.Length())
		}
		if collection.Item(1) != elem2 {
			t.Errorf("Expected elem2 at index 1")
		}

		// Remove an element
		parent.RemoveChild(elem1)
		if collection.Length() != 1 {
			t.Errorf("Expected length 1 after removing element, got %d", collection.Length())
		}
		if collection.Item(0) != elem2 {
			t.Errorf("Expected elem2 at index 0 after removal")
		}
	})
}

func TestHTMLCollection_SpecCompliance_NamespaceHandling(t *testing.T) {
	doc := NewDocument()
	parent := NewElement("div", doc)

	// Create HTML elements (in HTML namespace)
	htmlInput := NewElement("input", doc)
	htmlInput.SetAttribute("name", "htmlname")

	// Create SVG element (different namespace)
	svgRect, err := doc.CreateElementNS("http://www.w3.org/2000/svg", "rect")
	if err != nil {
		t.Fatalf("Failed to create SVG element: %v", err)
	}
	svgRect.SetAttribute("name", "svgname")

	parent.AppendChild(htmlInput)
	parent.AppendChild(svgRect)

	collection := parent.Children()

	t.Run("Name attribute only works for HTML namespace elements", func(t *testing.T) {
		// HTML element with name attribute should be found
		result := collection.NamedItem("htmlname")
		if result != htmlInput {
			t.Errorf("Expected to find HTML input by name attribute")
		}

		// SVG element with name attribute should NOT be found by name
		// (only by ID according to spec)
		result = collection.NamedItem("svgname")
		if result != nil {
			t.Errorf("Expected nil for SVG element name attribute (not in HTML namespace)")
		}
	})

	t.Run("ID attribute works for all namespaces", func(t *testing.T) {
		// Set ID on SVG element
		svgRect.SetAttribute("id", "svgid")

		// Should be found by ID regardless of namespace
		result := collection.NamedItem("svgid")
		if result != svgRect {
			t.Errorf("Expected to find SVG element by ID attribute")
		}
	})
}

func TestHTMLCollection_SpecCompliance_DuplicateNames(t *testing.T) {
	doc := NewDocument()
	parent := NewElement("div", doc)

	// Create multiple elements with same name/id
	elem1 := NewElement("input", doc)
	elem1.SetAttribute("name", "duplicate")

	elem2 := NewElement("input", doc)
	elem2.SetAttribute("name", "duplicate")

	elem3 := NewElement("span", doc)
	elem3.SetAttribute("id", "duplicate")

	parent.AppendChild(elem1)
	parent.AppendChild(elem2)
	parent.AppendChild(elem3)

	collection := parent.Children()

	t.Run("NamedItem returns first element with matching name/ID", func(t *testing.T) {
		// Should return the first element in tree order
		result := collection.NamedItem("duplicate")
		if result != elem1 {
			t.Errorf("Expected first element with name 'duplicate', got different element")
		}
	})
}

func TestHTMLCollection_SpecCompliance_EmptyCollection(t *testing.T) {
	doc := NewDocument()
	parent := NewElement("div", doc)

	// Add only non-element nodes
	parent.AppendChild(NewText("text", doc))
	parent.AppendChild(NewComment("comment", doc))

	collection := parent.Children()

	t.Run("Empty collection behavior", func(t *testing.T) {
		// If there are no elements, then there are no supported property indices
		if collection.Length() != 0 {
			t.Errorf("Expected length 0 for collection with no elements, got %d", collection.Length())
		}

		if collection.Item(0) != nil {
			t.Errorf("Expected nil for index 0 in empty collection")
		}

		if collection.NamedItem("anything") != nil {
			t.Errorf("Expected nil for any name in empty collection")
		}
	})
}

func TestHTMLCollection_SpecCompliance_SupportedPropertyNames(t *testing.T) {
	doc := NewDocument()
	parent := NewElement("div", doc)

	// Create elements to test the supported property names algorithm
	elem1 := NewElement("span", doc)
	elem1.SetAttribute("id", "uniqueid")

	elem2 := NewElement("input", doc)
	elem2.SetAttribute("name", "uniquename")

	elem3 := NewElement("form", doc)
	elem3.SetAttribute("id", "duplicateid")
	elem3.SetAttribute("name", "duplicatename")

	elem4 := NewElement("div", doc)
	elem4.SetAttribute("id", "duplicateid") // Duplicate ID should not be added twice

	elem5 := NewElement("input", doc)
	elem5.SetAttribute("name", "duplicatename") // Duplicate name should not be added twice

	elem6 := NewElement("input", doc)
	elem6.SetAttribute("name", "") // Empty name should be ignored

	parent.AppendChild(elem1)
	parent.AppendChild(elem2)
	parent.AppendChild(elem3)
	parent.AppendChild(elem4)
	parent.AppendChild(elem5)
	parent.AppendChild(elem6)

	collection := parent.Children()

	t.Run("All unique IDs are supported property names", func(t *testing.T) {
		// Test that we can access elements by their IDs
		if collection.NamedItem("uniqueid") != elem1 {
			t.Errorf("Expected elem1 for 'uniqueid'")
		}
		if collection.NamedItem("duplicateid") != elem3 { // First in tree order
			t.Errorf("Expected elem3 (first) for 'duplicateid'")
		}
	})

	t.Run("All unique names are supported property names", func(t *testing.T) {
		// Test that we can access elements by their names
		if collection.NamedItem("uniquename") != elem2 {
			t.Errorf("Expected elem2 for 'uniquename'")
		}
		if collection.NamedItem("duplicatename") != elem3 { // First in tree order
			t.Errorf("Expected elem3 (first) for 'duplicatename'")
		}
	})

	t.Run("Empty names are not supported property names", func(t *testing.T) {
		// Empty string should return nil as per spec
		if collection.NamedItem("") != nil {
			t.Errorf("Expected nil for empty string name")
		}
	})
}

func TestHTMLCollection_SpecCompliance_GetElementsByTagName(t *testing.T) {
	doc := NewDocument()
	parent := NewElement("div", doc)

	// Create elements with different tag names
	span1 := NewElement("span", doc)
	span1.SetAttribute("id", "span1")

	span2 := NewElement("span", doc)
	span2.SetAttribute("id", "span2")

	p1 := NewElement("p", doc)
	p1.SetAttribute("id", "p1")

	div1 := NewElement("div", doc)
	div1.SetAttribute("id", "div1")

	parent.AppendChild(span1)
	parent.AppendChild(p1)
	parent.AppendChild(span2)
	parent.AppendChild(div1)

	t.Run("GetElementsByTagName returns HTMLCollection", func(t *testing.T) {
		spanCollection := parent.GetElementsByTagName("span")

		// Should contain both span elements in tree order
		if spanCollection.Length() != 2 {
			t.Errorf("Expected 2 span elements, got %d", spanCollection.Length())
		}
		if spanCollection.Item(0) != span1 {
			t.Errorf("Expected span1 at index 0")
		}
		if spanCollection.Item(1) != span2 {
			t.Errorf("Expected span2 at index 1")
		}

		// Should support namedItem by ID
		if spanCollection.NamedItem("span1") != span1 {
			t.Errorf("Expected span1 for namedItem('span1')")
		}
		if spanCollection.NamedItem("span2") != span2 {
			t.Errorf("Expected span2 for namedItem('span2')")
		}

		// Should not find elements not in the collection
		if spanCollection.NamedItem("p1") != nil {
			t.Errorf("Expected nil for p1 in span collection")
		}
	})

	t.Run("Wildcard selector returns all elements", func(t *testing.T) {
		allCollection := parent.GetElementsByTagName("*")

		if allCollection.Length() != 4 {
			t.Errorf("Expected 4 elements with wildcard, got %d", allCollection.Length())
		}
	})
}
