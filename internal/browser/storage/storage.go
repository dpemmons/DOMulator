package storage

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/dop251/goja"
)

// Storage represents the Web Storage API (localStorage/sessionStorage)
type Storage struct {
	mu   sync.RWMutex
	data map[string]string
}

// NewStorage creates a new Storage instance
func NewStorage() *Storage {
	return &Storage{
		data: make(map[string]string),
	}
}

// Length returns the number of items in storage
func (s *Storage) Length() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.data)
}

// Key returns the name of the nth key in storage
func (s *Storage) Key(index int) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if index < 0 || index >= len(s.data) {
		return "", false
	}

	// Convert map to ordered slice for consistent indexing
	keys := make([]string, 0, len(s.data))
	for k := range s.data {
		keys = append(keys, k)
	}

	return keys[index], true
}

// GetItem returns the value for the given key, or empty string if not found
func (s *Storage) GetItem(key string) string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.data[key]
}

// SetItem sets the value for the given key
func (s *Storage) SetItem(key, value string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Check storage quota (simplified - 5MB limit)
	totalSize := 0
	for k, v := range s.data {
		if k != key { // Don't count existing value for this key
			totalSize += len(k) + len(v)
		}
	}
	totalSize += len(key) + len(value)

	if totalSize > 5*1024*1024 { // 5MB limit
		return fmt.Errorf("quota exceeded")
	}

	s.data[key] = value
	return nil
}

// RemoveItem removes the item with the given key
func (s *Storage) RemoveItem(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.data, key)
}

// Clear removes all items from storage
func (s *Storage) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data = make(map[string]string)
}

// Keys returns all keys in storage
func (s *Storage) Keys() []string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	keys := make([]string, 0, len(s.data))
	for k := range s.data {
		keys = append(keys, k)
	}
	return keys
}

// Values returns all values in storage
func (s *Storage) Values() []string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	values := make([]string, 0, len(s.data))
	for _, v := range s.data {
		values = append(values, v)
	}
	return values
}

// Entries returns all key-value pairs in storage
func (s *Storage) Entries() [][2]string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	entries := make([][2]string, 0, len(s.data))
	for k, v := range s.data {
		entries = append(entries, [2]string{k, v})
	}
	return entries
}

// ToJSON exports storage contents as JSON
func (s *Storage) ToJSON() (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	bytes, err := json.Marshal(s.data)
	if err != nil {
		return "", fmt.Errorf("failed to marshal storage: %w", err)
	}
	return string(bytes), nil
}

// FromJSON imports storage contents from JSON
func (s *Storage) FromJSON(jsonStr string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	var data map[string]string
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		return fmt.Errorf("failed to unmarshal storage: %w", err)
	}

	s.data = data
	return nil
}

// CreateStorageObject creates a JavaScript Storage object
func CreateStorageObject(vm *goja.Runtime, storage *Storage) *goja.Object {
	obj := vm.NewObject()

	// Length property
	obj.Set("length", storage.Length())

	// getItem method
	obj.Set("getItem", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			return goja.Null()
		}

		key := call.Arguments[0].String()
		value := storage.GetItem(key)

		if value == "" && !storage.hasKey(key) {
			return goja.Null()
		}

		return vm.ToValue(value)
	})

	// setItem method
	obj.Set("setItem", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 2 {
			panic(vm.NewTypeError("Storage.setItem requires 2 arguments"))
		}

		key := call.Arguments[0].String()
		value := call.Arguments[1].String()

		err := storage.SetItem(key, value)
		if err != nil {
			panic(vm.NewTypeError(err.Error()))
		}

		// Update length property
		obj.Set("length", storage.Length())

		return goja.Undefined()
	})

	// removeItem method
	obj.Set("removeItem", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			return goja.Undefined()
		}

		key := call.Arguments[0].String()
		storage.RemoveItem(key)

		// Update length property
		obj.Set("length", storage.Length())

		return goja.Undefined()
	})

	// clear method
	obj.Set("clear", func(call goja.FunctionCall) goja.Value {
		storage.Clear()

		// Update length property
		obj.Set("length", storage.Length())

		return goja.Undefined()
	})

	// key method
	obj.Set("key", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			return goja.Null()
		}

		indexValue := call.Arguments[0]
		index := int(indexValue.ToInteger())

		if key, exists := storage.Key(index); exists {
			return vm.ToValue(key)
		}

		return goja.Null()
	})

	// Add internal reference to Go Storage for other APIs to use
	obj.Set("__storage__", vm.ToValue(storage))

	return obj
}

// hasKey checks if a key exists in storage (helper method)
func (s *Storage) hasKey(key string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	_, exists := s.data[key]
	return exists
}

// StorageManager manages localStorage and sessionStorage instances
type StorageManager struct {
	localStorage   *Storage
	sessionStorage *Storage
}

// NewStorageManager creates a new storage manager
func NewStorageManager() *StorageManager {
	return &StorageManager{
		localStorage:   NewStorage(),
		sessionStorage: NewStorage(),
	}
}

// GetLocalStorage returns the localStorage instance
func (sm *StorageManager) GetLocalStorage() *Storage {
	return sm.localStorage
}

// GetSessionStorage returns the sessionStorage instance
func (sm *StorageManager) GetSessionStorage() *Storage {
	return sm.sessionStorage
}

// SetupStorageAPIs sets up localStorage and sessionStorage in the JavaScript runtime
func SetupStorageAPIs(vm *goja.Runtime, manager *StorageManager) {
	// Create localStorage object
	localStorage := CreateStorageObject(vm, manager.GetLocalStorage())
	vm.Set("localStorage", localStorage)

	// Create sessionStorage object
	sessionStorage := CreateStorageObject(vm, manager.GetSessionStorage())
	vm.Set("sessionStorage", sessionStorage)
}
