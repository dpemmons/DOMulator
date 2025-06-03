package forms

import (
	"strings"
	"testing"

	"github.com/dop251/goja"
)

func TestFormData_BasicOperations(t *testing.T) {
	fd := NewFormData()

	// Test Append
	fd.Append("name", "John")
	fd.Append("email", "john@example.com")
	fd.Append("name", "Jane") // Multiple values for same key

	// Test Get
	value, exists := fd.Get("name")
	if !exists || value != "John" {
		t.Errorf("Expected 'John', got '%s', exists: %v", value, exists)
	}

	// Test GetAll
	names := fd.GetAll("name")
	if len(names) != 2 || names[0] != "John" || names[1] != "Jane" {
		t.Errorf("Expected ['John', 'Jane'], got %v", names)
	}

	// Test Has
	if !fd.Has("email") {
		t.Error("Expected FormData to have 'email' key")
	}
	if fd.Has("nonexistent") {
		t.Error("Expected FormData not to have 'nonexistent' key")
	}

	// Test Set (should replace all existing values)
	fd.Set("name", "Bob")
	names = fd.GetAll("name")
	if len(names) != 1 || names[0] != "Bob" {
		t.Errorf("Expected ['Bob'], got %v", names)
	}

	// Test Delete
	fd.Delete("email")
	if fd.Has("email") {
		t.Error("Expected 'email' to be deleted")
	}
}

func TestFormData_FileOperations(t *testing.T) {
	fd := NewFormData()

	// Test AppendFile
	fd.AppendFile("avatar", "image data", "avatar.jpg")
	fd.Append("description", "Profile picture")

	// Test file entry
	entries := fd.Entries()
	if len(entries) != 2 {
		t.Errorf("Expected 2 entries, got %d", len(entries))
	}

	// Check file entry details
	found := false
	for _, entry := range fd.entries {
		if entry.Name == "avatar" && entry.Filename == "avatar.jpg" {
			found = true
			break
		}
	}
	if !found {
		t.Error("Expected to find file entry with filename")
	}
}

func TestFormData_IteratorMethods(t *testing.T) {
	fd := NewFormData()
	fd.Append("a", "1")
	fd.Append("b", "2")
	fd.Append("a", "3") // Duplicate key

	// Test Keys
	keys := fd.Keys()
	if len(keys) != 2 || keys[0] != "a" || keys[1] != "b" {
		t.Errorf("Expected ['a', 'b'], got %v", keys)
	}

	// Test Values
	values := fd.Values()
	if len(values) != 3 || values[0] != "1" || values[1] != "2" || values[2] != "3" {
		t.Errorf("Expected ['1', '2', '3'], got %v", values)
	}

	// Test Entries
	entries := fd.Entries()
	if len(entries) != 3 {
		t.Errorf("Expected 3 entries, got %d", len(entries))
	}
	if entries[0][0] != "a" || entries[0][1] != "1" {
		t.Errorf("Expected ['a', '1'], got %v", entries[0])
	}
}

func TestFormData_MultipartConversion(t *testing.T) {
	fd := NewFormData()
	fd.Append("name", "John")
	fd.AppendFile("file", "content", "test.txt")

	reader, contentType, err := fd.ToMultipartReader()
	if err != nil {
		t.Fatalf("Failed to convert to multipart: %v", err)
	}

	if !strings.Contains(contentType, "multipart/form-data") {
		t.Errorf("Expected multipart/form-data content type, got %s", contentType)
	}

	// Read the content to ensure it's properly formatted
	buf := make([]byte, 1024)
	n, _ := reader.Read(buf)
	content := string(buf[:n])

	if !strings.Contains(content, "name=\"name\"") {
		t.Error("Expected to find name field in multipart content")
	}
	if !strings.Contains(content, "John") {
		t.Error("Expected to find John value in multipart content")
	}
}

func TestFormData_URLEncodedConversion(t *testing.T) {
	fd := NewFormData()
	fd.Append("name", "John Doe")
	fd.Append("email", "john@example.com")
	fd.AppendFile("file", "content", "test.txt") // Should be skipped

	encoded := fd.ToURLEncoded()
	expected := "name=John+Doe&email=john@example.com"

	if encoded != expected {
		t.Errorf("Expected '%s', got '%s'", expected, encoded)
	}
}

func TestParseURLEncoded(t *testing.T) {
	content := "name=John+Doe&email=john%40example.com&city=New+York"
	fd := ParseURLEncoded(content)

	// Test parsed values
	name, exists := fd.Get("name")
	if !exists || name != "John Doe" {
		t.Errorf("Expected 'John Doe', got '%s'", name)
	}

	email, exists := fd.Get("email")
	if !exists || email != "john@example.com" {
		t.Errorf("Expected 'john@example.com', got '%s'", email)
	}

	city, exists := fd.Get("city")
	if !exists || city != "New York" {
		t.Errorf("Expected 'New York', got '%s'", city)
	}
}

func TestFormData_JavaScriptBinding(t *testing.T) {
	vm := goja.New()
	constructor := CreateFormDataConstructor(vm)

	// Test constructor call
	args := []goja.Value{}
	call := goja.ConstructorCall{Arguments: args}
	obj := constructor(call)

	if obj == nil {
		t.Fatal("FormData constructor returned nil")
	}

	// Test append method
	appendFn, exists := goja.AssertFunction(obj.Get("append"))
	if !exists {
		t.Fatal("FormData object missing append method")
	}

	// Call append
	_, err := appendFn(goja.Undefined(), vm.ToValue("name"), vm.ToValue("John"))
	if err != nil {
		t.Errorf("Failed to call append: %v", err)
	}

	// Test get method
	getFn, exists := goja.AssertFunction(obj.Get("get"))
	if !exists {
		t.Fatal("FormData object missing get method")
	}

	// Call get
	result, err := getFn(goja.Undefined(), vm.ToValue("name"))
	if err != nil {
		t.Errorf("Failed to call get: %v", err)
	}

	if result.String() != "John" {
		t.Errorf("Expected 'John', got '%s'", result.String())
	}
}

func TestFormData_JavaScriptIntegration(t *testing.T) {
	vm := goja.New()

	// Register FormData constructor
	vm.Set("FormData", CreateFormDataConstructor(vm))

	// Test JavaScript code
	script := `
		var fd = new FormData();
		fd.append('name', 'John');
		fd.append('email', 'john@example.com');
		fd.append('name', 'Jane');
		
		var results = {
			hasName: fd.has('name'),
			hasEmail: fd.has('email'),
			hasAge: fd.has('age'),
			getName: fd.get('name'),
			getAllNames: fd.getAll('name'),
			keys: fd.keys(),
			values: fd.values(),
			entries: fd.entries()
		};
		
		// Test set method
		fd.set('name', 'Bob');
		results.getNameAfterSet = fd.get('name');
		results.getAllNamesAfterSet = fd.getAll('name');
		
		// Test delete method
		fd.delete('email');
		results.hasEmailAfterDelete = fd.has('email');
		
		results;
	`

	result, err := vm.RunString(script)
	if err != nil {
		t.Fatalf("JavaScript execution failed: %v", err)
	}

	obj := result.ToObject(vm)

	// Test has method results
	if !obj.Get("hasName").ToBoolean() {
		t.Error("Expected FormData to have 'name'")
	}
	if !obj.Get("hasEmail").ToBoolean() {
		t.Error("Expected FormData to have 'email'")
	}
	if obj.Get("hasAge").ToBoolean() {
		t.Error("Expected FormData not to have 'age'")
	}

	// Test get method result
	if obj.Get("getName").String() != "John" {
		t.Errorf("Expected 'John', got '%s'", obj.Get("getName").String())
	}

	// Test set method effect
	if obj.Get("getNameAfterSet").String() != "Bob" {
		t.Errorf("Expected 'Bob' after set, got '%s'", obj.Get("getNameAfterSet").String())
	}

	// Test that set replaced all values
	setNamesValue := obj.Get("getAllNamesAfterSet")
	setNames := setNamesValue.Export().([]string)
	if len(setNames) != 1 || setNames[0] != "Bob" {
		t.Errorf("Expected ['Bob'] after set, got %v", setNames)
	}

	// Test delete method effect
	if obj.Get("hasEmailAfterDelete").ToBoolean() {
		t.Error("Expected 'email' to be deleted")
	}
}

func TestFormData_EdgeCases(t *testing.T) {
	// Test empty FormData
	fd := NewFormData()
	if fd.Has("nonexistent") {
		t.Error("Empty FormData should not have any keys")
	}

	keys := fd.Keys()
	if len(keys) != 0 {
		t.Errorf("Empty FormData should have no keys, got %v", keys)
	}

	// Test empty values
	fd.Append("empty", "")
	value, exists := fd.Get("empty")
	if !exists || value != "" {
		t.Errorf("Expected empty string, got '%s', exists: %v", value, exists)
	}

	// Test special characters in names and values
	fd.Append("special!@#", "value&=+")
	value, exists = fd.Get("special!@#")
	if !exists || value != "value&=+" {
		t.Errorf("Expected 'value&=+', got '%s'", value)
	}
}

func TestFormData_URLEncoding(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello world", "hello+world"},
		{"test&data", "test%26data"},
		{"key=value", "key%3Dvalue"},
		{"100%", "100%25"},
		{"", ""},
	}

	for _, test := range tests {
		result := urlEncode(test.input)
		if result != test.expected {
			t.Errorf("urlEncode(%q) = %q, expected %q", test.input, result, test.expected)
		}

		// Test round-trip
		decoded := urlDecode(result)
		if decoded != test.input {
			t.Errorf("Round-trip failed: %q -> %q -> %q", test.input, result, decoded)
		}
	}
}

func TestParseURLEncoded_EdgeCases(t *testing.T) {
	// Test empty content
	fd := ParseURLEncoded("")
	if len(fd.entries) != 0 {
		t.Error("Empty content should result in empty FormData")
	}

	// Test malformed pairs
	fd = ParseURLEncoded("name=John&=invalid&email")

	// Should have parsed the valid parts
	if !fd.Has("name") {
		t.Error("Should have parsed 'name' field")
	}

	// Should handle key without value
	if !fd.Has("email") {
		t.Error("Should have created entry for 'email' with empty value")
	}

	value, _ := fd.Get("email")
	if value != "" {
		t.Errorf("Expected empty value for 'email', got '%s'", value)
	}
}
