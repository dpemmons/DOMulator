package dom

import (
	"testing"
)

// TestNamedNodeMapSpecCompliance tests the NamedNodeMap implementation against WHATWG DOM Section 4.9.1
func TestNamedNodeMapSpecCompliance(t *testing.T) {
	// Create test document and element
	doc := NewDocument()
	doc.documentType = "html" // Set as HTML document
	elem := NewElement("div", doc)
	elem.namespaceURI = htmlNamespace // Set as HTML namespace element

	// Test initial state
	if elem.attributes.Length() != 0 {
		t.Errorf("Expected initial attribute count to be 0, got %d", elem.attributes.Length())
	}

	if elem.attributes.HasAttributes() {
		t.Error("Expected HasAttributes() to return false for empty collection")
	}

	// Test GetNamedItem returns null for non-existent attribute
	if attr := elem.attributes.GetNamedItem("nonexistent"); attr != nil {
		t.Error("Expected GetNamedItem to return nil for non-existent attribute")
	}

	// Test Item with invalid index
	if attr := elem.attributes.Item(-1); attr != nil {
		t.Error("Expected Item(-1) to return nil")
	}

	if attr := elem.attributes.Item(0); attr != nil {
		t.Error("Expected Item(0) to return nil for empty collection")
	}

	// Test adding an attribute
	attr1 := NewAttr("id", "test-id", doc)
	oldAttr := elem.attributes.SetNamedItem(attr1)

	if oldAttr != nil {
		t.Error("Expected SetNamedItem to return nil when adding new attribute")
	}

	if elem.attributes.Length() != 1 {
		t.Errorf("Expected attribute count to be 1, got %d", elem.attributes.Length())
	}

	if !elem.attributes.HasAttributes() {
		t.Error("Expected HasAttributes() to return true after adding attribute")
	}

	// Test retrieving the attribute
	retrieved := elem.attributes.GetNamedItem("id")
	if retrieved == nil {
		t.Error("Expected to retrieve added attribute")
	}

	if retrieved != attr1 {
		t.Error("Expected retrieved attribute to be the same instance")
	}

	if retrieved.Value() != "test-id" {
		t.Errorf("Expected attribute value to be 'test-id', got '%s'", retrieved.Value())
	}

	// Test Item method
	itemAttr := elem.attributes.Item(0)
	if itemAttr == nil {
		t.Error("Expected Item(0) to return the first attribute")
	}

	if itemAttr != attr1 {
		t.Error("Expected Item(0) to return the same attribute")
	}

	// Test case insensitivity for HTML elements in HTML documents
	retrieved = elem.attributes.GetNamedItem("ID") // Uppercase
	if retrieved == nil {
		t.Error("Expected case-insensitive retrieval for HTML elements in HTML documents")
	}

	if retrieved != attr1 {
		t.Error("Expected case-insensitive retrieval to return same attribute")
	}
}

// TestNamedNodeMapHTMLCaseInsensitivity tests case sensitivity behavior per spec
func TestNamedNodeMapHTMLCaseInsensitivity(t *testing.T) {
	// HTML element in HTML document should be case-insensitive
	htmlDoc := NewDocument()
	htmlDoc.documentType = "html"
	htmlElem := NewElement("div", htmlDoc)
	htmlElem.namespaceURI = htmlNamespace

	attr := NewAttr("Class", "test-class", htmlDoc)
	htmlElem.attributes.SetNamedItem(attr)

	// Should be retrievable with lowercase
	retrieved := htmlElem.attributes.GetNamedItem("class")
	if retrieved == nil {
		t.Error("Expected case-insensitive retrieval for HTML elements")
	}

	// Should be retrievable with original case
	retrieved = htmlElem.attributes.GetNamedItem("Class")
	if retrieved == nil {
		t.Error("Expected to retrieve with original case")
	}

	// XML element should be case-sensitive
	xmlDoc := NewDocument()
	xmlDoc.documentType = "xml"
	xmlElem := NewElement("div", xmlDoc)
	xmlElem.namespaceURI = "" // Not HTML namespace

	attr2 := NewAttr("Class", "test-class", xmlDoc)
	xmlElem.attributes.SetNamedItem(attr2)

	// Should NOT be retrievable with lowercase for XML
	retrieved = xmlElem.attributes.GetNamedItem("class")
	if retrieved != nil {
		t.Error("Expected case-sensitive retrieval for XML elements")
	}

	// Should be retrievable with exact case
	retrieved = xmlElem.attributes.GetNamedItem("Class")
	if retrieved == nil {
		t.Error("Expected to retrieve with exact case for XML")
	}
}

// TestNamedNodeMapSetNamedItemReplacement tests attribute replacement behavior
func TestNamedNodeMapSetNamedItemReplacement(t *testing.T) {
	doc := NewDocument()
	elem := NewElement("div", doc)

	// Add first attribute
	attr1 := NewAttr("id", "first-value", doc)
	oldAttr := elem.attributes.SetNamedItem(attr1)

	if oldAttr != nil {
		t.Error("Expected no old attribute when adding new")
	}

	// Replace with same attribute - should return same instance
	sameAttr := elem.attributes.SetNamedItem(attr1)
	if sameAttr != attr1 {
		t.Error("Expected SetNamedItem to return same attribute when setting same instance")
	}

	// Replace with different attribute with same name
	attr2 := NewAttr("id", "second-value", doc)
	oldAttr = elem.attributes.SetNamedItem(attr2)

	if oldAttr != attr1 {
		t.Error("Expected SetNamedItem to return old attribute when replacing")
	}

	if elem.attributes.Length() != 1 {
		t.Error("Expected collection to still have only one attribute after replacement")
	}

	retrieved := elem.attributes.GetNamedItem("id")
	if retrieved != attr2 {
		t.Error("Expected retrieved attribute to be the new one")
	}

	if retrieved.Value() != "second-value" {
		t.Errorf("Expected new value 'second-value', got '%s'", retrieved.Value())
	}

	// Check that old attribute no longer has owner element
	if attr1.OwnerElement() != nil {
		t.Error("Expected old attribute to have nil owner element")
	}
}

// TestNamedNodeMapInUseAttributeError tests error when setting in-use attribute
func TestNamedNodeMapInUseAttributeError(t *testing.T) {
	doc := NewDocument()
	elem1 := NewElement("div", doc)
	elem2 := NewElement("span", doc)

	// Add attribute to first element
	attr := NewAttr("id", "test", doc)
	elem1.attributes.SetNamedItem(attr)

	// Try to add same attribute to second element - should panic
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic when setting in-use attribute to different element")
		}
	}()

	elem2.attributes.SetNamedItem(attr)
}

// TestNamedNodeMapRemoveNamedItem tests attribute removal
func TestNamedNodeMapRemoveNamedItem(t *testing.T) {
	doc := NewDocument()
	elem := NewElement("div", doc)

	// Add attribute
	attr := NewAttr("id", "test", doc)
	elem.attributes.SetNamedItem(attr)

	if elem.attributes.Length() != 1 {
		t.Error("Expected one attribute before removal")
	}

	// Remove attribute
	removed := elem.attributes.RemoveNamedItem("id")
	if removed != attr {
		t.Error("Expected RemoveNamedItem to return the removed attribute")
	}

	if elem.attributes.Length() != 0 {
		t.Error("Expected zero attributes after removal")
	}

	if removed.OwnerElement() != nil {
		t.Error("Expected removed attribute to have nil owner element")
	}

	// Try to remove non-existent attribute - should panic
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic when removing non-existent attribute")
		}
	}()

	elem.attributes.RemoveNamedItem("nonexistent")
}

// TestNamedNodeMapNamespaceSupport tests namespace-aware methods
func TestNamedNodeMapNamespaceSupport(t *testing.T) {
	doc := NewDocument()
	elem := NewElement("div", doc)

	// Add namespaced attribute
	nsAttr, err := NewAttrNS("http://example.com/ns", "example:attr", "value", doc)
	if err != nil {
		t.Fatalf("Failed to create namespaced attribute: %v", err)
	}

	oldAttr := elem.attributes.SetNamedItemNS(nsAttr)
	if oldAttr != nil {
		t.Error("Expected no old attribute when adding new namespaced attribute")
	}

	// Retrieve by namespace and local name
	retrieved := elem.attributes.GetNamedItemNS("http://example.com/ns", "attr")
	if retrieved == nil {
		t.Error("Expected to retrieve namespaced attribute")
	}

	if retrieved != nsAttr {
		t.Error("Expected retrieved attribute to be the same instance")
	}

	// Test empty string namespace handling
	retrieved = elem.attributes.GetNamedItemNS("", "attr")
	if retrieved != nil {
		t.Error("Expected no match for empty namespace")
	}

	// Remove by namespace
	removed := elem.attributes.RemoveNamedItemNS("http://example.com/ns", "attr")
	if removed != nsAttr {
		t.Error("Expected RemoveNamedItemNS to return the removed attribute")
	}

	if elem.attributes.Length() != 0 {
		t.Error("Expected zero attributes after namespace removal")
	}

	// Try to remove non-existent namespaced attribute - should panic
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic when removing non-existent namespaced attribute")
		}
	}()

	elem.attributes.RemoveNamedItemNS("http://example.com/ns", "nonexistent")
}

// TestNamedNodeMapSupportedPropertyNames tests the supported property names algorithm
func TestNamedNodeMapSupportedPropertyNames(t *testing.T) {
	// Test HTML element in HTML document
	htmlDoc := NewDocument()
	htmlDoc.documentType = "html"
	htmlElem := NewElement("div", htmlDoc)
	htmlElem.namespaceURI = htmlNamespace

	// Add mixed case attributes
	attr1 := NewAttr("id", "test", htmlDoc)
	attr2 := NewAttr("Class", "test-class", htmlDoc)
	attr3 := NewAttr("DATA-VALUE", "test-data", htmlDoc)

	htmlElem.attributes.SetNamedItem(attr1)
	htmlElem.attributes.SetNamedItem(attr2)
	htmlElem.attributes.SetNamedItem(attr3)

	names := htmlElem.attributes.SupportedPropertyNames()

	// For HTML elements in HTML documents, only lowercase names should be in supported property names
	expectedNames := []string{"id", "class", "data-value"} // All should be normalized to lowercase
	if len(names) != len(expectedNames) {
		t.Errorf("Expected %d supported property names, got %d: %v", len(expectedNames), len(names), names)
	}

	for _, expected := range expectedNames {
		found := false
		for _, name := range names {
			if name == expected {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected '%s' in supported property names", expected)
		}
	}

	// Test XML element - all names should be included
	xmlDoc := NewDocument()
	xmlDoc.documentType = "xml"
	xmlElem := NewElement("div", xmlDoc)

	xmlElem.attributes.SetNamedItem(NewAttr("id", "test", xmlDoc))
	xmlElem.attributes.SetNamedItem(NewAttr("Class", "test", xmlDoc))
	xmlElem.attributes.SetNamedItem(NewAttr("DATA-VALUE", "test", xmlDoc))

	xmlNames := xmlElem.attributes.SupportedPropertyNames()
	expectedXMLNames := []string{"id", "Class", "DATA-VALUE"}

	if len(xmlNames) != len(expectedXMLNames) {
		t.Errorf("Expected %d XML supported property names, got %d: %v", len(expectedXMLNames), len(xmlNames), xmlNames)
	}
}

// TestNamedNodeMapGetAttributeNames tests the convenience method
func TestNamedNodeMapGetAttributeNames(t *testing.T) {
	doc := NewDocument()
	elem := NewElement("div", doc)

	// Add attributes
	attrs := []string{"id", "class", "data-value"}
	for _, name := range attrs {
		attr := NewAttr(name, "test", doc)
		elem.attributes.SetNamedItem(attr)
	}

	names := elem.attributes.GetAttributeNames()
	if len(names) != len(attrs) {
		t.Errorf("Expected %d attribute names, got %d", len(attrs), len(names))
	}

	for i, expected := range attrs {
		if names[i] != expected {
			t.Errorf("Expected attribute name '%s' at index %d, got '%s'", expected, i, names[i])
		}
	}
}

// TestNamedNodeMapContainsMethods tests the Contains and ContainsNS methods
func TestNamedNodeMapContainsMethods(t *testing.T) {
	doc := NewDocument()
	elem := NewElement("div", doc)

	// Test empty collection
	if elem.attributes.Contains("id") {
		t.Error("Expected Contains to return false for empty collection")
	}

	if elem.attributes.ContainsNS("http://example.com", "attr") {
		t.Error("Expected ContainsNS to return false for empty collection")
	}

	// Add attributes
	attr1 := NewAttr("id", "test", doc)
	elem.attributes.SetNamedItem(attr1)

	nsAttr, _ := NewAttrNS("http://example.com", "example:attr", "value", doc)
	elem.attributes.SetNamedItem(nsAttr)

	// Test Contains
	if !elem.attributes.Contains("id") {
		t.Error("Expected Contains to return true for existing attribute")
	}

	if elem.attributes.Contains("nonexistent") {
		t.Error("Expected Contains to return false for non-existent attribute")
	}

	// Test ContainsNS
	if !elem.attributes.ContainsNS("http://example.com", "attr") {
		t.Error("Expected ContainsNS to return true for existing namespaced attribute")
	}

	if elem.attributes.ContainsNS("http://other.com", "attr") {
		t.Error("Expected ContainsNS to return false for wrong namespace")
	}
}

// TestNamedNodeMapClear tests the Clear method
func TestNamedNodeMapClear(t *testing.T) {
	doc := NewDocument()
	elem := NewElement("div", doc)

	// Add multiple attributes
	attrs := []*Attr{
		NewAttr("id", "test1", doc),
		NewAttr("class", "test2", doc),
		NewAttr("data-value", "test3", doc),
	}

	for _, attr := range attrs {
		elem.attributes.SetNamedItem(attr)
	}

	if elem.attributes.Length() != 3 {
		t.Error("Expected 3 attributes before clear")
	}

	// Clear all attributes
	elem.attributes.Clear()

	if elem.attributes.Length() != 0 {
		t.Error("Expected 0 attributes after clear")
	}

	if elem.attributes.HasAttributes() {
		t.Error("Expected HasAttributes to return false after clear")
	}

	// Check that all attributes have nil owner element
	for _, attr := range attrs {
		if attr.OwnerElement() != nil {
			t.Error("Expected cleared attribute to have nil owner element")
		}
	}
}

// TestNamedNodeMapToSlice tests the ToSlice method
func TestNamedNodeMapToSlice(t *testing.T) {
	doc := NewDocument()
	elem := NewElement("div", doc)

	// Add attributes
	originalAttrs := []*Attr{
		NewAttr("id", "test1", doc),
		NewAttr("class", "test2", doc),
	}

	for _, attr := range originalAttrs {
		elem.attributes.SetNamedItem(attr)
	}

	slice := elem.attributes.ToSlice()

	if len(slice) != len(originalAttrs) {
		t.Errorf("Expected slice length %d, got %d", len(originalAttrs), len(slice))
	}

	// Verify it's a copy (modifying slice shouldn't affect original)
	slice[0] = nil
	if elem.attributes.Item(0) == nil {
		t.Error("ToSlice should return a copy, not reference to internal slice")
	}
}

// TestNamedNodeMapForEach tests the ForEach method
func TestNamedNodeMapForEach(t *testing.T) {
	doc := NewDocument()
	elem := NewElement("div", doc)

	// Add attributes
	attrs := []string{"id", "class", "data-value"}
	for _, name := range attrs {
		attr := NewAttr(name, "test", doc)
		elem.attributes.SetNamedItem(attr)
	}

	// Test complete iteration
	var visited []string
	elem.attributes.ForEach(func(attr *Attr) bool {
		visited = append(visited, attr.Name())
		return true // Continue iteration
	})

	if len(visited) != len(attrs) {
		t.Errorf("Expected to visit %d attributes, visited %d", len(attrs), len(visited))
	}

	// Test early termination
	var earlyVisited []string
	elem.attributes.ForEach(func(attr *Attr) bool {
		earlyVisited = append(earlyVisited, attr.Name())
		return len(earlyVisited) < 2 // Stop after 2 items
	})

	if len(earlyVisited) != 2 {
		t.Errorf("Expected early termination after 2 items, got %d", len(earlyVisited))
	}
}

// TestNamedNodeMapSortByName tests the SortByName method
func TestNamedNodeMapSortByName(t *testing.T) {
	doc := NewDocument()
	elem := NewElement("div", doc)

	// Add attributes in random order
	names := []string{"zebra", "alpha", "beta"}
	for _, name := range names {
		attr := NewAttr(name, "test", doc)
		elem.attributes.SetNamedItem(attr)
	}

	// Sort by name
	elem.attributes.SortByName()

	// Verify sorted order
	expectedOrder := []string{"alpha", "beta", "zebra"}
	for i, expected := range expectedOrder {
		attr := elem.attributes.Item(i)
		if attr == nil || attr.Name() != expected {
			t.Errorf("Expected attribute '%s' at index %d, got '%s'", expected, i, attr.Name())
		}
	}
}
