package storage

import (
	"testing"

	"github.com/dop251/goja"
)

func TestNewStorage(t *testing.T) {
	storage := NewStorage()

	if storage == nil {
		t.Fatal("NewStorage should not return nil")
	}

	if storage.data == nil {
		t.Error("Storage data map should be initialized")
	}

	if storage.Length() != 0 {
		t.Errorf("New storage should be empty, got length %d", storage.Length())
	}
}

func TestStorageBasicOperations(t *testing.T) {
	storage := NewStorage()

	// Test SetItem and GetItem
	err := storage.SetItem("key1", "value1")
	if err != nil {
		t.Fatalf("SetItem failed: %v", err)
	}

	value := storage.GetItem("key1")
	if value != "value1" {
		t.Errorf("Expected 'value1', got '%s'", value)
	}

	// Test Length
	if storage.Length() != 1 {
		t.Errorf("Expected length 1, got %d", storage.Length())
	}

	// Test non-existent key
	value = storage.GetItem("nonexistent")
	if value != "" {
		t.Errorf("Expected empty string for non-existent key, got '%s'", value)
	}
}

func TestStorageRemoveItem(t *testing.T) {
	storage := NewStorage()

	// Add items
	storage.SetItem("key1", "value1")
	storage.SetItem("key2", "value2")

	if storage.Length() != 2 {
		t.Errorf("Expected length 2, got %d", storage.Length())
	}

	// Remove one item
	storage.RemoveItem("key1")

	if storage.Length() != 1 {
		t.Errorf("Expected length 1 after removal, got %d", storage.Length())
	}

	value := storage.GetItem("key1")
	if value != "" {
		t.Errorf("Expected empty string for removed key, got '%s'", value)
	}

	// Verify remaining item still exists
	value = storage.GetItem("key2")
	if value != "value2" {
		t.Errorf("Expected 'value2', got '%s'", value)
	}
}

func TestStorageClear(t *testing.T) {
	storage := NewStorage()

	// Add multiple items
	storage.SetItem("key1", "value1")
	storage.SetItem("key2", "value2")
	storage.SetItem("key3", "value3")

	if storage.Length() != 3 {
		t.Errorf("Expected length 3, got %d", storage.Length())
	}

	// Clear all items
	storage.Clear()

	if storage.Length() != 0 {
		t.Errorf("Expected length 0 after clear, got %d", storage.Length())
	}

	// Verify all items are gone
	if storage.GetItem("key1") != "" {
		t.Error("key1 should be empty after clear")
	}
	if storage.GetItem("key2") != "" {
		t.Error("key2 should be empty after clear")
	}
	if storage.GetItem("key3") != "" {
		t.Error("key3 should be empty after clear")
	}
}

func TestStorageKey(t *testing.T) {
	storage := NewStorage()

	// Add items
	storage.SetItem("alpha", "value1")
	storage.SetItem("beta", "value2")
	storage.SetItem("gamma", "value3")

	// Test valid indices
	for i := 0; i < storage.Length(); i++ {
		key, exists := storage.Key(i)
		if !exists {
			t.Errorf("Key at index %d should exist", i)
		}
		if key == "" {
			t.Errorf("Key at index %d should not be empty", i)
		}
	}

	// Test invalid indices
	_, exists := storage.Key(-1)
	if exists {
		t.Error("Key at index -1 should not exist")
	}

	_, exists = storage.Key(storage.Length())
	if exists {
		t.Error("Key at index beyond length should not exist")
	}
}

func TestStorageIteratorMethods(t *testing.T) {
	storage := NewStorage()

	// Add test data
	storage.SetItem("key1", "value1")
	storage.SetItem("key2", "value2")
	storage.SetItem("key3", "value3")

	// Test Keys
	keys := storage.Keys()
	if len(keys) != 3 {
		t.Errorf("Expected 3 keys, got %d", len(keys))
	}

	// Test Values
	values := storage.Values()
	if len(values) != 3 {
		t.Errorf("Expected 3 values, got %d", len(values))
	}

	// Test Entries
	entries := storage.Entries()
	if len(entries) != 3 {
		t.Errorf("Expected 3 entries, got %d", len(entries))
	}

	// Verify entries structure
	for _, entry := range entries {
		if len(entry) != 2 {
			t.Errorf("Entry should have 2 elements, got %d", len(entry))
		}
		if entry[0] == "" || entry[1] == "" {
			t.Error("Entry key and value should not be empty")
		}
	}
}

func TestStorageQuotaLimit(t *testing.T) {
	storage := NewStorage()

	// Create a large value that exceeds quota
	largeValue := make([]byte, 6*1024*1024) // 6MB
	for i := range largeValue {
		largeValue[i] = 'a'
	}

	err := storage.SetItem("large", string(largeValue))
	if err == nil {
		t.Error("Expected quota exceeded error for large value")
	}

	if err.Error() != "quota exceeded" {
		t.Errorf("Expected 'quota exceeded' error, got '%s'", err.Error())
	}
}

func TestStorageJSONSerialization(t *testing.T) {
	storage := NewStorage()

	// Add test data
	storage.SetItem("key1", "value1")
	storage.SetItem("key2", "value2")

	// Export to JSON
	json, err := storage.ToJSON()
	if err != nil {
		t.Fatalf("ToJSON failed: %v", err)
	}

	if json == "" {
		t.Error("JSON output should not be empty")
	}

	// Create new storage and import
	newStorage := NewStorage()
	err = newStorage.FromJSON(json)
	if err != nil {
		t.Fatalf("FromJSON failed: %v", err)
	}

	// Verify data was imported correctly
	if newStorage.Length() != 2 {
		t.Errorf("Expected length 2 after import, got %d", newStorage.Length())
	}

	if newStorage.GetItem("key1") != "value1" {
		t.Errorf("Expected 'value1', got '%s'", newStorage.GetItem("key1"))
	}

	if newStorage.GetItem("key2") != "value2" {
		t.Errorf("Expected 'value2', got '%s'", newStorage.GetItem("key2"))
	}
}

func TestStorageManager(t *testing.T) {
	manager := NewStorageManager()

	if manager == nil {
		t.Fatal("NewStorageManager should not return nil")
	}

	localStorage := manager.GetLocalStorage()
	sessionStorage := manager.GetSessionStorage()

	if localStorage == nil {
		t.Error("localStorage should not be nil")
	}

	if sessionStorage == nil {
		t.Error("sessionStorage should not be nil")
	}

	// Verify they are separate instances
	localStorage.SetItem("test", "local")
	sessionStorage.SetItem("test", "session")

	if localStorage.GetItem("test") != "local" {
		t.Error("localStorage should contain 'local'")
	}

	if sessionStorage.GetItem("test") != "session" {
		t.Error("sessionStorage should contain 'session'")
	}
}

func TestCreateStorageObject(t *testing.T) {
	vm := goja.New()
	storage := NewStorage()

	obj := CreateStorageObject(vm, storage)

	if obj == nil {
		t.Fatal("CreateStorageObject should not return nil")
	}

	// Test length property
	length := obj.Get("length")
	if length == nil {
		t.Error("Storage object should have length property")
	}

	// Test methods exist
	methods := []string{"getItem", "setItem", "removeItem", "clear", "key"}
	for _, method := range methods {
		if obj.Get(method) == nil {
			t.Errorf("Storage object should have %s method", method)
		}
	}
}

func TestStorageJavaScriptIntegration(t *testing.T) {
	vm := goja.New()
	storage := NewStorage()

	obj := CreateStorageObject(vm, storage)
	vm.Set("storage", obj)

	// Test setItem and getItem
	_, err := vm.RunString(`
		storage.setItem('key1', 'value1');
		var result = storage.getItem('key1');
	`)
	if err != nil {
		t.Fatalf("JavaScript execution failed: %v", err)
	}

	result, err := vm.RunString(`storage.getItem('key1')`)
	if err != nil {
		t.Fatalf("Failed to get item: %v", err)
	}

	if result.String() != "value1" {
		t.Errorf("Expected 'value1', got '%s'", result.String())
	}

	// Test length property
	_, err = vm.RunString(`storage.setItem('key2', 'value2')`)
	if err != nil {
		t.Fatalf("Failed to set second item: %v", err)
	}

	lengthResult, err := vm.RunString(`storage.length`)
	if err != nil {
		t.Fatalf("Failed to get length: %v", err)
	}

	if lengthResult.ToInteger() != 2 {
		t.Errorf("Expected length 2, got %d", lengthResult.ToInteger())
	}
}

func TestStorageJavaScriptRemoveAndClear(t *testing.T) {
	vm := goja.New()
	storage := NewStorage()

	obj := CreateStorageObject(vm, storage)
	vm.Set("storage", obj)

	// Add items
	_, err := vm.RunString(`
		storage.setItem('key1', 'value1');
		storage.setItem('key2', 'value2');
		storage.setItem('key3', 'value3');
	`)
	if err != nil {
		t.Fatalf("Failed to set items: %v", err)
	}

	// Test removeItem
	_, err = vm.RunString(`storage.removeItem('key2')`)
	if err != nil {
		t.Fatalf("Failed to remove item: %v", err)
	}

	result, err := vm.RunString(`storage.getItem('key2')`)
	if err != nil {
		t.Fatalf("Failed to get removed item: %v", err)
	}

	if !goja.IsNull(result) {
		t.Error("Removed item should return null")
	}

	// Test clear
	_, err = vm.RunString(`storage.clear()`)
	if err != nil {
		t.Fatalf("Failed to clear storage: %v", err)
	}

	lengthResult, err := vm.RunString(`storage.length`)
	if err != nil {
		t.Fatalf("Failed to get length after clear: %v", err)
	}

	if lengthResult.ToInteger() != 0 {
		t.Errorf("Expected length 0 after clear, got %d", lengthResult.ToInteger())
	}
}

func TestStorageJavaScriptKeyMethod(t *testing.T) {
	vm := goja.New()
	storage := NewStorage()

	obj := CreateStorageObject(vm, storage)
	vm.Set("storage", obj)

	// Add items
	_, err := vm.RunString(`
		storage.setItem('alpha', 'value1');
		storage.setItem('beta', 'value2');
	`)
	if err != nil {
		t.Fatalf("Failed to set items: %v", err)
	}

	// Test key method with valid index
	result, err := vm.RunString(`storage.key(0)`)
	if err != nil {
		t.Fatalf("Failed to get key at index 0: %v", err)
	}

	key := result.String()
	if key != "alpha" && key != "beta" {
		t.Errorf("Expected 'alpha' or 'beta', got '%s'", key)
	}

	// Test key method with invalid index
	result, err = vm.RunString(`storage.key(10)`)
	if err != nil {
		t.Fatalf("Failed to get key at invalid index: %v", err)
	}

	if !goja.IsNull(result) {
		t.Error("Invalid index should return null")
	}
}

func TestStorageJavaScriptErrorHandling(t *testing.T) {
	vm := goja.New()
	storage := NewStorage()

	obj := CreateStorageObject(vm, storage)
	vm.Set("storage", obj)

	// Test setItem with insufficient arguments
	_, err := vm.RunString(`
		try {
			storage.setItem('key');
		} catch(e) {
			var errorCaught = true;
		}
	`)
	if err != nil {
		t.Fatalf("Error handling test failed: %v", err)
	}

	result, err := vm.RunString(`typeof errorCaught !== 'undefined'`)
	if err != nil {
		t.Fatalf("Failed to check error: %v", err)
	}

	if !result.ToBoolean() {
		t.Error("Expected error to be caught for insufficient arguments")
	}
}

func TestSetupStorageAPIs(t *testing.T) {
	vm := goja.New()
	manager := NewStorageManager()

	SetupStorageAPIs(vm, manager)

	// Test localStorage exists
	result, err := vm.RunString(`typeof localStorage`)
	if err != nil {
		t.Fatalf("Failed to check localStorage: %v", err)
	}

	if result.String() != "object" {
		t.Error("localStorage should be an object")
	}

	// Test sessionStorage exists
	result, err = vm.RunString(`typeof sessionStorage`)
	if err != nil {
		t.Fatalf("Failed to check sessionStorage: %v", err)
	}

	if result.String() != "object" {
		t.Error("sessionStorage should be an object")
	}

	// Test that they are separate
	_, err = vm.RunString(`
		localStorage.setItem('test', 'local');
		sessionStorage.setItem('test', 'session');
	`)
	if err != nil {
		t.Fatalf("Failed to set items: %v", err)
	}

	localResult, err := vm.RunString(`localStorage.getItem('test')`)
	if err != nil {
		t.Fatalf("Failed to get localStorage item: %v", err)
	}

	sessionResult, err := vm.RunString(`sessionStorage.getItem('test')`)
	if err != nil {
		t.Fatalf("Failed to get sessionStorage item: %v", err)
	}

	if localResult.String() != "local" {
		t.Errorf("Expected 'local', got '%s'", localResult.String())
	}

	if sessionResult.String() != "session" {
		t.Errorf("Expected 'session', got '%s'", sessionResult.String())
	}
}

func TestStorageConcurrency(t *testing.T) {
	storage := NewStorage()

	// Test concurrent operations
	done := make(chan bool, 2)

	go func() {
		for i := 0; i < 100; i++ {
			storage.SetItem("key1", "value1")
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 100; i++ {
			storage.GetItem("key1")
		}
		done <- true
	}()

	// Wait for both goroutines
	<-done
	<-done

	// Verify final state
	if storage.GetItem("key1") != "value1" {
		t.Error("Concurrent operations should not corrupt data")
	}
}
