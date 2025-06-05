package dom

import (
	"testing"
)

func TestImportNode_LegacyBooleanParameter(t *testing.T) {
	// Test backward compatibility with boolean parameter
	sourceDoc := NewDocument()
	targetDoc := NewDocument()

	// Create a source element with children
	sourceElement := sourceDoc.CreateElement("div")
	sourceElement.SetAttribute("id", "test")
	sourceText := sourceDoc.CreateTextNode("Hello")
	sourceChild := sourceDoc.CreateElement("span")
	sourceChild.AppendChild(sourceDoc.CreateTextNode("Child"))
	sourceElement.AppendChild(sourceText)
	sourceElement.AppendChild(sourceChild)

	// Test deep import (legacy: true)
	deepImported, err := targetDoc.ImportNode(sourceElement, true)
	if err != nil {
		t.Fatalf("ImportNode with deep=true failed: %v", err)
	}

	deepElement := deepImported.(*Element)
	if deepElement.GetAttribute("id") != "test" {
		t.Errorf("Expected id='test', got %q", deepElement.GetAttribute("id"))
	}
	if deepElement.ChildNodes().Length() != 2 {
		t.Errorf("Expected 2 children, got %d", deepElement.ChildNodes().Length())
	}
	if deepElement.OwnerDocument() != targetDoc {
		t.Errorf("Expected owner document to be targetDoc")
	}

	// Test shallow import (legacy: false)
	shallowImported, err := targetDoc.ImportNode(sourceElement, false)
	if err != nil {
		t.Fatalf("ImportNode with deep=false failed: %v", err)
	}

	shallowElement := shallowImported.(*Element)
	if shallowElement.GetAttribute("id") != "test" {
		t.Errorf("Expected id='test', got %q", shallowElement.GetAttribute("id"))
	}
	if shallowElement.ChildNodes().Length() != 0 {
		t.Errorf("Expected 0 children for shallow import, got %d", shallowElement.ChildNodes().Length())
	}
	if shallowElement.OwnerDocument() != targetDoc {
		t.Errorf("Expected owner document to be targetDoc")
	}
}

func TestImportNode_ImportNodeOptions(t *testing.T) {
	sourceDoc := NewDocument()
	targetDoc := NewDocument()

	// Create a source element with children
	sourceElement := sourceDoc.CreateElement("div")
	sourceElement.SetAttribute("class", "test")
	sourceText := sourceDoc.CreateTextNode("Content")
	sourceChild := sourceDoc.CreateElement("p")
	sourceChild.AppendChild(sourceDoc.CreateTextNode("Paragraph"))
	sourceElement.AppendChild(sourceText)
	sourceElement.AppendChild(sourceChild)

	// Test with ImportNodeOptions (selfOnly: false - deep clone)
	deepOptions := ImportNodeOptions{SelfOnly: false}
	deepImported, err := targetDoc.ImportNode(sourceElement, deepOptions)
	if err != nil {
		t.Fatalf("ImportNode with ImportNodeOptions (selfOnly: false) failed: %v", err)
	}

	deepElement := deepImported.(*Element)
	if deepElement.GetAttribute("class") != "test" {
		t.Errorf("Expected class='test', got %q", deepElement.GetAttribute("class"))
	}
	if deepElement.ChildNodes().Length() != 2 {
		t.Errorf("Expected 2 children, got %d", deepElement.ChildNodes().Length())
	}

	// Test with ImportNodeOptions (selfOnly: true - shallow clone)
	shallowOptions := ImportNodeOptions{SelfOnly: true}
	shallowImported, err := targetDoc.ImportNode(sourceElement, shallowOptions)
	if err != nil {
		t.Fatalf("ImportNode with ImportNodeOptions (selfOnly: true) failed: %v", err)
	}

	shallowElement := shallowImported.(*Element)
	if shallowElement.GetAttribute("class") != "test" {
		t.Errorf("Expected class='test', got %q", shallowElement.GetAttribute("class"))
	}
	if shallowElement.ChildNodes().Length() != 0 {
		t.Errorf("Expected 0 children for selfOnly import, got %d", shallowElement.ChildNodes().Length())
	}
}

func TestImportNode_MapInput(t *testing.T) {
	sourceDoc := NewDocument()
	targetDoc := NewDocument()

	sourceElement := sourceDoc.CreateElement("section")
	sourceElement.AppendChild(sourceDoc.CreateTextNode("Test content"))

	// Test with map input (JavaScript object style)
	mapOptions := map[string]interface{}{
		"selfOnly": true,
	}
	imported, err := targetDoc.ImportNode(sourceElement, mapOptions)
	if err != nil {
		t.Fatalf("ImportNode with map options failed: %v", err)
	}

	importedElement := imported.(*Element)
	if importedElement.ChildNodes().Length() != 0 {
		t.Errorf("Expected 0 children for selfOnly=true, got %d", importedElement.ChildNodes().Length())
	}

	// Test with map input including custom element registry
	registryOptions := map[string]interface{}{
		"selfOnly":              false,
		"customElementRegistry": "mock-registry",
	}
	imported2, err := targetDoc.ImportNode(sourceElement, registryOptions)
	if err != nil {
		t.Fatalf("ImportNode with registry options failed: %v", err)
	}

	importedElement2 := imported2.(*Element)
	if importedElement2.ChildNodes().Length() != 1 {
		t.Errorf("Expected 1 child for selfOnly=false, got %d", importedElement2.ChildNodes().Length())
	}
}

func TestImportNode_JSONStringInput(t *testing.T) {
	sourceDoc := NewDocument()
	targetDoc := NewDocument()

	sourceElement := sourceDoc.CreateElement("article")
	sourceElement.AppendChild(sourceDoc.CreateTextNode("Article content"))

	// Test with JSON string input
	jsonOptions := `{"selfOnly": true}`
	imported, err := targetDoc.ImportNode(sourceElement, jsonOptions)
	if err != nil {
		t.Fatalf("ImportNode with JSON options failed: %v", err)
	}

	importedElement := imported.(*Element)
	if importedElement.ChildNodes().Length() != 0 {
		t.Errorf("Expected 0 children for JSON selfOnly=true, got %d", importedElement.ChildNodes().Length())
	}

	// Test with invalid JSON
	invalidJSON := `{"selfOnly": true` // Missing closing brace
	_, err = targetDoc.ImportNode(sourceElement, invalidJSON)
	if err == nil {
		t.Errorf("Expected error for invalid JSON, got nil")
	}
}

func TestImportNode_ErrorCases(t *testing.T) {
	sourceDoc := NewDocument()
	targetDoc := NewDocument()

	// Test importing a document (should fail)
	_, err := targetDoc.ImportNode(sourceDoc)
	if err == nil {
		t.Errorf("Expected error when importing a document")
	}
	if domErr, ok := err.(*DOMException); !ok || domErr.Name() != "NotSupportedError" {
		t.Errorf("Expected NotSupportedError, got %T", err)
	}

	// Test with unsupported input type
	sourceElement := sourceDoc.CreateElement("div")
	_, err = targetDoc.ImportNode(sourceElement, 123) // Invalid type
	if err == nil {
		t.Errorf("Expected error for unsupported input type")
	}
}

func TestImportNode_VariousNodeTypes(t *testing.T) {
	sourceDoc := NewDocument()
	targetDoc := NewDocument()

	tests := []struct {
		name string
		node func() Node
	}{
		{
			name: "Text",
			node: func() Node { return sourceDoc.CreateTextNode("Test text") },
		},
		{
			name: "Comment",
			node: func() Node { return sourceDoc.CreateComment("Test comment") },
		},
		{
			name: "ProcessingInstruction",
			node: func() Node {
				pi, _ := sourceDoc.CreateProcessingInstruction("xml", "version='1.0'")
				return pi
			},
		},
		{
			name: "DocumentFragment",
			node: func() Node { return sourceDoc.CreateDocumentFragment() },
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			node := test.node()
			imported, err := targetDoc.ImportNode(node, true)
			if err != nil {
				t.Fatalf("ImportNode failed for %s: %v", test.name, err)
			}
			if imported.NodeType() != node.NodeType() {
				t.Errorf("Node type mismatch: expected %d, got %d", node.NodeType(), imported.NodeType())
			}
			if imported.OwnerDocument() != targetDoc {
				t.Errorf("Owner document not set correctly for %s", test.name)
			}
		})
	}
}

func TestAdoptNode_BasicFunctionality(t *testing.T) {
	sourceDoc := NewDocument()
	targetDoc := NewDocument()

	// Create element in source document
	element := sourceDoc.CreateElement("div")
	element.SetAttribute("data-test", "value")
	text := sourceDoc.CreateTextNode("Adopted text")
	element.AppendChild(text)

	// Verify initial state
	if element.OwnerDocument() != sourceDoc {
		t.Errorf("Initial owner document should be sourceDoc")
	}

	// Adopt the element
	adopted, err := targetDoc.AdoptNode(element)
	if err != nil {
		t.Fatalf("AdoptNode failed: %v", err)
	}

	// Verify adoption worked
	if adopted != element {
		t.Errorf("AdoptNode should return the same node")
	}
	if adopted.OwnerDocument() != targetDoc {
		t.Errorf("Adopted node should have targetDoc as owner")
	}
	if text.OwnerDocument() != targetDoc {
		t.Errorf("Child nodes should also be adopted")
	}

	// Verify attributes are preserved
	adoptedElement := adopted.(*Element)
	if adoptedElement.GetAttribute("data-test") != "value" {
		t.Errorf("Attributes should be preserved during adoption")
	}
}

func TestAdoptNode_RemovalFromParent(t *testing.T) {
	sourceDoc := NewDocument()
	targetDoc := NewDocument()

	// Create parent and child in source document
	parent := sourceDoc.CreateElement("parent")
	child := sourceDoc.CreateElement("child")
	parent.AppendChild(child)

	// Verify initial parent-child relationship
	if child.ParentNode() != parent {
		t.Errorf("Initial parent should be set")
	}
	if parent.ChildNodes().Length() != 1 {
		t.Errorf("Parent should have 1 child initially")
	}

	// Adopt the child
	adopted, err := targetDoc.AdoptNode(child)
	if err != nil {
		t.Fatalf("AdoptNode failed: %v", err)
	}

	// Verify child was removed from original parent
	if adopted.ParentNode() != nil {
		t.Errorf("Adopted node should not have a parent")
	}
	if parent.ChildNodes().Length() != 0 {
		t.Errorf("Original parent should have no children after adoption")
	}
}

func TestAdoptNode_ErrorCases(t *testing.T) {
	targetDoc := NewDocument()

	// Test adopting a document (should fail)
	sourceDoc := NewDocument()
	_, err := targetDoc.AdoptNode(sourceDoc)
	if err == nil {
		t.Errorf("Expected error when adopting a document")
	}
	if domErr, ok := err.(*DOMException); !ok || domErr.Name() != "NotSupportedError" {
		t.Errorf("Expected NotSupportedError, got %T", err)
	}

	// TODO: Test adopting a shadow root when shadow DOM is implemented
}

func TestAdoptNode_CustomElementCallback(t *testing.T) {
	sourceDoc := NewDocument()
	targetDoc := NewDocument()

	// Create a custom element (with hyphen in name)
	customElement := sourceDoc.CreateElement("my-custom")
	customElement.SetAttribute("data-custom", "true")

	// Adopt the custom element
	adopted, err := targetDoc.AdoptNode(customElement)
	if err != nil {
		t.Fatalf("AdoptNode failed for custom element: %v", err)
	}

	// Verify the element was adopted
	if adopted.OwnerDocument() != targetDoc {
		t.Errorf("Custom element should be adopted to target document")
	}

	// Test with customized built-in element (has "is" attribute)
	builtinElement := sourceDoc.CreateElement("button")
	builtinElement.SetAttribute("is", "my-button")

	adopted2, err := targetDoc.AdoptNode(builtinElement)
	if err != nil {
		t.Fatalf("AdoptNode failed for customized built-in: %v", err)
	}

	if adopted2.OwnerDocument() != targetDoc {
		t.Errorf("Customized built-in should be adopted to target document")
	}
}

func TestParseImportNodeOptions_AllInputTypes(t *testing.T) {
	tests := []struct {
		name     string
		input    ImportNodeOptionsInput
		expected ImportNodeOptions
		wantErr  bool
	}{
		{
			name:     "nil input",
			input:    nil,
			expected: ImportNodeOptions{SelfOnly: false},
			wantErr:  false,
		},
		{
			name:     "boolean true (deep)",
			input:    true,
			expected: ImportNodeOptions{SelfOnly: false}, // true becomes selfOnly=false
			wantErr:  false,
		},
		{
			name:     "boolean false (shallow)",
			input:    false,
			expected: ImportNodeOptions{SelfOnly: true}, // false becomes selfOnly=true
			wantErr:  false,
		},
		{
			name:     "ImportNodeOptions struct",
			input:    ImportNodeOptions{SelfOnly: true, CustomElementRegistry: "test"},
			expected: ImportNodeOptions{SelfOnly: true, CustomElementRegistry: "test"},
			wantErr:  false,
		},
		{
			name: "map input",
			input: map[string]interface{}{
				"selfOnly":              true,
				"customElementRegistry": "registry",
			},
			expected: ImportNodeOptions{SelfOnly: true, CustomElementRegistry: "registry"},
			wantErr:  false,
		},
		{
			name:     "JSON string",
			input:    `{"selfOnly": true}`,
			expected: ImportNodeOptions{SelfOnly: true},
			wantErr:  false,
		},
		{
			name:    "invalid JSON",
			input:   `{"selfOnly": true`,
			wantErr: true,
		},
		{
			name:    "unsupported type",
			input:   123,
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := parseImportNodeOptions(test.input)
			if test.wantErr {
				if err == nil {
					t.Errorf("Expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if result.SelfOnly != test.expected.SelfOnly {
				t.Errorf("SelfOnly mismatch: expected %v, got %v", test.expected.SelfOnly, result.SelfOnly)
			}
			if result.CustomElementRegistry != test.expected.CustomElementRegistry {
				t.Errorf("CustomElementRegistry mismatch: expected %v, got %v",
					test.expected.CustomElementRegistry, result.CustomElementRegistry)
			}
		})
	}
}

func TestImportNode_DefaultBehavior(t *testing.T) {
	// Test that calling ImportNode without options defaults to deep clone
	sourceDoc := NewDocument()
	targetDoc := NewDocument()

	sourceElement := sourceDoc.CreateElement("div")
	sourceElement.AppendChild(sourceDoc.CreateTextNode("Content"))

	// Call ImportNode without options
	imported, err := targetDoc.ImportNode(sourceElement)
	if err != nil {
		t.Fatalf("ImportNode without options failed: %v", err)
	}

	importedElement := imported.(*Element)
	if importedElement.ChildNodes().Length() != 1 {
		t.Errorf("Expected 1 child (default deep clone), got %d", importedElement.ChildNodes().Length())
	}
}

func TestCloneOptions_SelfOnlyBehavior(t *testing.T) {
	// Test that SelfOnly option in CloneOptions works correctly
	doc := NewDocument()

	parent := doc.CreateElement("parent")
	child1 := doc.CreateElement("child1")
	child2 := doc.CreateElement("child2")
	parent.AppendChild(child1)
	parent.AppendChild(child2)

	// Test selfOnly=true (should not clone children)
	selfOnlyOpts := CloneOptions{
		Document: doc,
		Subtree:  true,
		SelfOnly: true,
		Parent:   nil,
	}
	selfOnlyClone := cloneNode(parent, selfOnlyOpts)
	if selfOnlyClone.ChildNodes().Length() != 0 {
		t.Errorf("SelfOnly clone should have 0 children, got %d", selfOnlyClone.ChildNodes().Length())
	}

	// Test selfOnly=false (should clone children)
	deepOpts := CloneOptions{
		Document: doc,
		Subtree:  true,
		SelfOnly: false,
		Parent:   nil,
	}
	deepClone := cloneNode(parent, deepOpts)
	if deepClone.ChildNodes().Length() != 2 {
		t.Errorf("Deep clone should have 2 children, got %d", deepClone.ChildNodes().Length())
	}
}
