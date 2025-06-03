package forms

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"strings"

	"github.com/dop251/goja"
)

// FormData represents the FormData Web API
type FormData struct {
	entries []FormEntry
}

// FormEntry represents a single form field entry
type FormEntry struct {
	Name     string
	Value    string
	Filename string // For file uploads, empty for regular fields
}

// NewFormData creates a new FormData instance
func NewFormData() *FormData {
	return &FormData{
		entries: make([]FormEntry, 0),
	}
}

// Append adds a new value to an existing key, or creates the key if it doesn't exist
func (fd *FormData) Append(name, value string) {
	fd.entries = append(fd.entries, FormEntry{
		Name:  name,
		Value: value,
	})
}

// AppendFile adds a file field (with filename)
func (fd *FormData) AppendFile(name, value, filename string) {
	fd.entries = append(fd.entries, FormEntry{
		Name:     name,
		Value:    value,
		Filename: filename,
	})
}

// Set sets a new value for an existing key, or creates the key if it doesn't exist
// If the key already exists, it replaces all existing values
func (fd *FormData) Set(name, value string) {
	// Remove all existing entries with this name
	fd.Delete(name)
	// Add the new entry
	fd.Append(name, value)
}

// SetFile sets a file field, replacing any existing entries with the same name
func (fd *FormData) SetFile(name, value, filename string) {
	// Remove all existing entries with this name
	fd.Delete(name)
	// Add the new file entry
	fd.AppendFile(name, value, filename)
}

// Get returns the first value associated with the given name
func (fd *FormData) Get(name string) (string, bool) {
	for _, entry := range fd.entries {
		if entry.Name == name {
			return entry.Value, true
		}
	}
	return "", false
}

// GetAll returns all values associated with the given name
func (fd *FormData) GetAll(name string) []string {
	var values []string
	for _, entry := range fd.entries {
		if entry.Name == name {
			values = append(values, entry.Value)
		}
	}
	return values
}

// Has returns true if the FormData contains the given name
func (fd *FormData) Has(name string) bool {
	for _, entry := range fd.entries {
		if entry.Name == name {
			return true
		}
	}
	return false
}

// Delete removes all entries with the given name
func (fd *FormData) Delete(name string) {
	var filtered []FormEntry
	for _, entry := range fd.entries {
		if entry.Name != name {
			filtered = append(filtered, entry)
		}
	}
	fd.entries = filtered
}

// Keys returns an iterator over all keys
func (fd *FormData) Keys() []string {
	seen := make(map[string]bool)
	var keys []string
	for _, entry := range fd.entries {
		if !seen[entry.Name] {
			keys = append(keys, entry.Name)
			seen[entry.Name] = true
		}
	}
	return keys
}

// Values returns an iterator over all values
func (fd *FormData) Values() []string {
	var values []string
	for _, entry := range fd.entries {
		values = append(values, entry.Value)
	}
	return values
}

// Entries returns all form entries as name-value pairs
func (fd *FormData) Entries() [][2]string {
	var entries [][2]string
	for _, entry := range fd.entries {
		entries = append(entries, [2]string{entry.Name, entry.Value})
	}
	return entries
}

// ToMultipartReader converts FormData to multipart/form-data format
func (fd *FormData) ToMultipartReader() (io.Reader, string, error) {
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	for _, entry := range fd.entries {
		if entry.Filename != "" {
			// File field
			part, err := writer.CreateFormFile(entry.Name, entry.Filename)
			if err != nil {
				return nil, "", fmt.Errorf("failed to create form file: %w", err)
			}
			_, err = part.Write([]byte(entry.Value))
			if err != nil {
				return nil, "", fmt.Errorf("failed to write file data: %w", err)
			}
		} else {
			// Regular field
			err := writer.WriteField(entry.Name, entry.Value)
			if err != nil {
				return nil, "", fmt.Errorf("failed to write field: %w", err)
			}
		}
	}

	err := writer.Close()
	if err != nil {
		return nil, "", fmt.Errorf("failed to close multipart writer: %w", err)
	}

	return &buf, writer.FormDataContentType(), nil
}

// ToURLEncoded converts FormData to application/x-www-form-urlencoded format
func (fd *FormData) ToURLEncoded() string {
	var parts []string
	for _, entry := range fd.entries {
		// Skip file entries in URL encoding
		if entry.Filename == "" {
			parts = append(parts, fmt.Sprintf("%s=%s",
				urlEncode(entry.Name),
				urlEncode(entry.Value)))
		}
	}
	return strings.Join(parts, "&")
}

// urlEncode provides basic URL encoding for form data
func urlEncode(s string) string {
	// Basic implementation - replace special characters
	// NOTE: % must be encoded first to avoid double-encoding
	s = strings.ReplaceAll(s, "%", "%25")
	s = strings.ReplaceAll(s, " ", "+")
	s = strings.ReplaceAll(s, "&", "%26")
	s = strings.ReplaceAll(s, "=", "%3D")
	return s
}

// ParseMultipart parses multipart/form-data content into FormData
func ParseMultipart(content string, boundary string) (*FormData, error) {
	fd := NewFormData()

	reader := multipart.NewReader(strings.NewReader(content), boundary)

	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to read multipart: %w", err)
		}

		// Read part data
		data, err := io.ReadAll(part)
		if err != nil {
			return nil, fmt.Errorf("failed to read part data: %w", err)
		}

		name := part.FormName()
		filename := part.FileName()

		if filename != "" {
			fd.AppendFile(name, string(data), filename)
		} else {
			fd.Append(name, string(data))
		}
	}

	return fd, nil
}

// ParseURLEncoded parses application/x-www-form-urlencoded content into FormData
func ParseURLEncoded(content string) *FormData {
	fd := NewFormData()

	if content == "" {
		return fd
	}

	pairs := strings.Split(content, "&")
	for _, pair := range pairs {
		if pair == "" {
			continue
		}

		parts := strings.SplitN(pair, "=", 2)
		name := urlDecode(parts[0])
		value := ""
		if len(parts) > 1 {
			value = urlDecode(parts[1])
		}

		fd.Append(name, value)
	}

	return fd
}

// urlDecode provides basic URL decoding for form data
func urlDecode(s string) string {
	// Basic implementation - reverse the encoding
	s = strings.ReplaceAll(s, "+", " ")
	s = strings.ReplaceAll(s, "%40", "@")
	s = strings.ReplaceAll(s, "%26", "&")
	s = strings.ReplaceAll(s, "%3D", "=")
	s = strings.ReplaceAll(s, "%25", "%")
	return s
}

// CreateFormDataConstructor creates the FormData constructor for JavaScript binding
func CreateFormDataConstructor(vm *goja.Runtime) func(goja.ConstructorCall) *goja.Object {
	return func(call goja.ConstructorCall) *goja.Object {
		fd := NewFormData()

		// If form element is passed, parse it
		if len(call.Arguments) > 0 && !goja.IsUndefined(call.Arguments[0]) {
			// TODO: Parse HTML form element when DOM integration is needed
			// For now, just create empty FormData
		}

		obj := vm.NewObject()

		// Set up the FormData prototype methods
		obj.Set("append", func(call goja.FunctionCall) goja.Value {
			if len(call.Arguments) < 2 {
				panic(vm.NewTypeError("FormData.append requires at least 2 arguments"))
			}

			name := call.Arguments[0].String()
			value := call.Arguments[1].String()

			// Check if third argument is provided (for files)
			if len(call.Arguments) > 2 && !goja.IsUndefined(call.Arguments[2]) {
				filename := call.Arguments[2].String()
				fd.AppendFile(name, value, filename)
			} else {
				fd.Append(name, value)
			}

			return goja.Undefined()
		})

		obj.Set("set", func(call goja.FunctionCall) goja.Value {
			if len(call.Arguments) < 2 {
				panic(vm.NewTypeError("FormData.set requires at least 2 arguments"))
			}

			name := call.Arguments[0].String()
			value := call.Arguments[1].String()

			// Check if third argument is provided (for files)
			if len(call.Arguments) > 2 && !goja.IsUndefined(call.Arguments[2]) {
				filename := call.Arguments[2].String()
				fd.SetFile(name, value, filename)
			} else {
				fd.Set(name, value)
			}

			return goja.Undefined()
		})

		obj.Set("get", func(call goja.FunctionCall) goja.Value {
			if len(call.Arguments) < 1 {
				panic(vm.NewTypeError("FormData.get requires 1 argument"))
			}

			name := call.Arguments[0].String()
			if value, exists := fd.Get(name); exists {
				return vm.ToValue(value)
			}
			return goja.Null()
		})

		obj.Set("getAll", func(call goja.FunctionCall) goja.Value {
			if len(call.Arguments) < 1 {
				panic(vm.NewTypeError("FormData.getAll requires 1 argument"))
			}

			name := call.Arguments[0].String()
			values := fd.GetAll(name)
			return vm.ToValue(values)
		})

		obj.Set("has", func(call goja.FunctionCall) goja.Value {
			if len(call.Arguments) < 1 {
				panic(vm.NewTypeError("FormData.has requires 1 argument"))
			}

			name := call.Arguments[0].String()
			return vm.ToValue(fd.Has(name))
		})

		obj.Set("delete", func(call goja.FunctionCall) goja.Value {
			if len(call.Arguments) < 1 {
				panic(vm.NewTypeError("FormData.delete requires 1 argument"))
			}

			name := call.Arguments[0].String()
			fd.Delete(name)
			return goja.Undefined()
		})

		obj.Set("keys", func(call goja.FunctionCall) goja.Value {
			keys := fd.Keys()
			return vm.ToValue(keys)
		})

		obj.Set("values", func(call goja.FunctionCall) goja.Value {
			values := fd.Values()
			return vm.ToValue(values)
		})

		obj.Set("entries", func(call goja.FunctionCall) goja.Value {
			entries := fd.Entries()
			return vm.ToValue(entries)
		})

		// Add internal reference to Go FormData for other APIs to use
		obj.Set("__formdata__", vm.ToValue(fd))

		return obj
	}
}
