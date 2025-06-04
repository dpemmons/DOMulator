package dom

import (
	"testing"
)

func TestElement_ClassList_Integration(t *testing.T) {
	// Create a document and an element
	doc := NewDocument()
	element := NewElement("div", doc)

	// Test initial classList
	classList := element.ClassList()
	if classList == nil {
		t.Error("ClassList() should not return nil")
	}

	// Test adding classes via classList
	err := classList.Add("class1", "class2")
	if err != nil {
		t.Errorf("Add() error: %v", err)
	}

	// Verify the class attribute was updated
	classAttr := element.GetAttribute("class")
	if classAttr != "class1 class2" {
		t.Errorf("class attribute = %q, want %q", classAttr, "class1 class2")
	}

	// Test HasClass method
	if !element.HasClass("class1") {
		t.Error("HasClass(class1) should return true")
	}
	if !element.HasClass("class2") {
		t.Error("HasClass(class2) should return true")
	}
	if element.HasClass("class3") {
		t.Error("HasClass(class3) should return false")
	}

	// Test that classList reflects changes made via setAttribute
	element.SetAttribute("class", "newclass1 newclass2 newclass3")

	// Test live updates
	if classList.Length() != 3 {
		t.Errorf("classList.Length() = %d, want 3", classList.Length())
	}

	if !classList.Contains("newclass1") {
		t.Error("classList should contain newclass1")
	}

	if classList.Contains("class1") {
		t.Error("classList should not contain class1 after setAttribute")
	}

	// Test toggle
	result := classList.Toggle("newclass4", nil)
	if !result {
		t.Error("Toggle(newclass4) should return true when adding")
	}

	if element.GetAttribute("class") != "newclass1 newclass2 newclass3 newclass4" {
		t.Errorf("class after toggle = %q, want %q", element.GetAttribute("class"), "newclass1 newclass2 newclass3 newclass4")
	}

	// Test toggle remove
	result = classList.Toggle("newclass2", nil)
	if result {
		t.Error("Toggle(newclass2) should return false when removing")
	}

	if element.GetAttribute("class") != "newclass1 newclass3 newclass4" {
		t.Errorf("class after toggle remove = %q, want %q", element.GetAttribute("class"), "newclass1 newclass3 newclass4")
	}

	// Test remove
	classList.Remove("newclass3", "nonexistent")
	if element.GetAttribute("class") != "newclass1 newclass4" {
		t.Errorf("class after remove = %q, want %q", element.GetAttribute("class"), "newclass1 newclass4")
	}

	// Test replace
	success := classList.Replace("newclass1", "replacedclass")
	if !success {
		t.Error("Replace should return true when successful")
	}

	if element.GetAttribute("class") != "replacedclass newclass4" {
		t.Errorf("class after replace = %q, want %q", element.GetAttribute("class"), "replacedclass newclass4")
	}

	// Test classList persistence - getting classList again should return the same instance
	classList2 := element.ClassList()
	if classList != classList2 {
		t.Error("ClassList() should return the same instance on subsequent calls")
	}
}

func TestElement_ClassList_EmptyClass(t *testing.T) {
	doc := NewDocument()
	element := NewElement("div", doc)

	classList := element.ClassList()

	// Test with empty class attribute
	if classList.Length() != 0 {
		t.Errorf("Length() for empty class = %d, want 0", classList.Length())
	}

	if classList.String() != "" {
		t.Errorf("String() for empty class = %q, want empty string", classList.String())
	}

	// Test adding to empty
	err := classList.Add("firstclass")
	if err != nil {
		t.Errorf("Add() error: %v", err)
	}

	if element.GetAttribute("class") != "firstclass" {
		t.Errorf("class attribute = %q, want %q", element.GetAttribute("class"), "firstclass")
	}
}

func TestElement_ClassList_InvalidTokens(t *testing.T) {
	doc := NewDocument()
	element := NewElement("div", doc)

	classList := element.ClassList()

	// Test adding invalid tokens
	err := classList.Add("valid", "invalid token", "valid2")
	if err == nil {
		t.Error("Add() should return error for invalid tokens")
	}

	// Verify no changes were made
	if element.GetAttribute("class") != "" {
		t.Errorf("class attribute should be empty after failed Add(), got %q", element.GetAttribute("class"))
	}

	// Test that valid operations still work
	err = classList.Add("valid1", "valid2")
	if err != nil {
		t.Errorf("Add() error for valid tokens: %v", err)
	}

	if element.GetAttribute("class") != "valid1 valid2" {
		t.Errorf("class attribute = %q, want %q", element.GetAttribute("class"), "valid1 valid2")
	}
}

func TestElement_GetElementsByClassName_WithDOMTokenList(t *testing.T) {
	doc := NewDocument()
	parent := NewElement("div", doc)

	// Create child elements
	child1 := NewElement("span", doc)
	child1.ClassList().Add("class1", "class2")

	child2 := NewElement("p", doc)
	child2.ClassList().Add("class2", "class3")

	child3 := NewElement("div", doc)
	child3.ClassList().Add("class1")

	// Build DOM structure
	parent.AppendChild(child1)
	parent.AppendChild(child2)
	parent.AppendChild(child3)

	// Test getElementsByClassName
	class1Elements := parent.GetElementsByClassName("class1")
	if class1Elements.Length() != 2 {
		t.Errorf("getElementsByClassName(class1) length = %d, want 2", class1Elements.Length())
	}

	class2Elements := parent.GetElementsByClassName("class2")
	if class2Elements.Length() != 2 {
		t.Errorf("getElementsByClassName(class2) length = %d, want 2", class2Elements.Length())
	}

	class3Elements := parent.GetElementsByClassName("class3")
	if class3Elements.Length() != 1 {
		t.Errorf("getElementsByClassName(class3) length = %d, want 1", class3Elements.Length())
	}

	nonexistentElements := parent.GetElementsByClassName("nonexistent")
	if nonexistentElements.Length() != 0 {
		t.Errorf("getElementsByClassName(nonexistent) length = %d, want 0", nonexistentElements.Length())
	}
}
