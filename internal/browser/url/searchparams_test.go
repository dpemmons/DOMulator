package url

import (
	"reflect"
	"testing"
)

func TestNewURLSearchParams_Empty(t *testing.T) {
	params := NewURLSearchParams()
	if params.Size() != 0 {
		t.Errorf("Expected empty params, got size %d", params.Size())
	}
	if params.ToString() != "" {
		t.Errorf("Expected empty string, got %s", params.ToString())
	}
}

func TestNewURLSearchParams_FromString(t *testing.T) {
	tests := []struct {
		input    string
		expected map[string][]string
	}{
		{"foo=bar", map[string][]string{"foo": {"bar"}}},
		{"foo=bar&baz=qux", map[string][]string{"foo": {"bar"}, "baz": {"qux"}}},
		{"foo=bar&foo=baz", map[string][]string{"foo": {"bar", "baz"}}},
		{"?foo=bar&baz=qux", map[string][]string{"foo": {"bar"}, "baz": {"qux"}}},
		{"", map[string][]string{}},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			params := NewURLSearchParams(test.input)

			for name, expectedValues := range test.expected {
				actualValues := params.GetAll(name)
				if !reflect.DeepEqual(actualValues, expectedValues) {
					t.Errorf("Expected %s to have values %v, got %v", name, expectedValues, actualValues)
				}
			}
		})
	}
}

func TestURLSearchParams_Append(t *testing.T) {
	params := NewURLSearchParams()

	params.Append("foo", "bar")
	if params.Get("foo") != "bar" {
		t.Errorf("Expected foo to be bar, got %s", params.Get("foo"))
	}

	params.Append("foo", "baz")
	values := params.GetAll("foo")
	expected := []string{"bar", "baz"}
	if !reflect.DeepEqual(values, expected) {
		t.Errorf("Expected foo to have values %v, got %v", expected, values)
	}
}

func TestURLSearchParams_Delete(t *testing.T) {
	params := NewURLSearchParams("foo=bar&foo=baz&qux=quux")

	params.Delete("foo")
	if params.Has("foo") {
		t.Error("Expected foo to be deleted")
	}
	if !params.Has("qux") {
		t.Error("Expected qux to still exist")
	}
	if params.Get("qux") != "quux" {
		t.Errorf("Expected qux to be quux, got %s", params.Get("qux"))
	}
}

func TestURLSearchParams_Get(t *testing.T) {
	params := NewURLSearchParams("foo=bar&foo=baz&qux=quux")

	// Get should return first value
	if params.Get("foo") != "bar" {
		t.Errorf("Expected foo to be bar, got %s", params.Get("foo"))
	}

	if params.Get("qux") != "quux" {
		t.Errorf("Expected qux to be quux, got %s", params.Get("qux"))
	}

	// Non-existent key should return empty string
	if params.Get("nonexistent") != "" {
		t.Errorf("Expected nonexistent key to return empty string, got %s", params.Get("nonexistent"))
	}
}

func TestURLSearchParams_GetAll(t *testing.T) {
	params := NewURLSearchParams("foo=bar&foo=baz&qux=quux")

	fooValues := params.GetAll("foo")
	expected := []string{"bar", "baz"}
	if !reflect.DeepEqual(fooValues, expected) {
		t.Errorf("Expected foo values to be %v, got %v", expected, fooValues)
	}

	quxValues := params.GetAll("qux")
	expected = []string{"quux"}
	if !reflect.DeepEqual(quxValues, expected) {
		t.Errorf("Expected qux values to be %v, got %v", expected, quxValues)
	}

	// Non-existent key should return empty slice
	nonexistentValues := params.GetAll("nonexistent")
	if len(nonexistentValues) != 0 {
		t.Errorf("Expected empty slice for nonexistent key, got %v", nonexistentValues)
	}
}

func TestURLSearchParams_Has(t *testing.T) {
	params := NewURLSearchParams("foo=bar&baz=qux")

	if !params.Has("foo") {
		t.Error("Expected params to have foo")
	}

	if !params.Has("baz") {
		t.Error("Expected params to have baz")
	}

	if params.Has("nonexistent") {
		t.Error("Expected params not to have nonexistent")
	}
}

func TestURLSearchParams_Set(t *testing.T) {
	params := NewURLSearchParams("foo=bar&foo=baz&qux=quux")

	// Set should replace all existing values
	params.Set("foo", "newvalue")
	if params.Get("foo") != "newvalue" {
		t.Errorf("Expected foo to be newvalue, got %s", params.Get("foo"))
	}

	fooValues := params.GetAll("foo")
	expected := []string{"newvalue"}
	if !reflect.DeepEqual(fooValues, expected) {
		t.Errorf("Expected foo to have only one value %v, got %v", expected, fooValues)
	}

	// Set non-existent key should add it
	params.Set("newkey", "newval")
	if params.Get("newkey") != "newval" {
		t.Errorf("Expected newkey to be newval, got %s", params.Get("newkey"))
	}
}

func TestURLSearchParams_Sort(t *testing.T) {
	params := NewURLSearchParams("z=1&a=2&m=3")
	params.Sort()

	// Check that keys are in alphabetical order
	keys := params.Keys()
	expected := []string{"a", "m", "z"}
	if !reflect.DeepEqual(keys, expected) {
		t.Errorf("Expected keys to be sorted %v, got %v", expected, keys)
	}
}

func TestURLSearchParams_ToString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"foo=bar&baz=qux", "baz=qux&foo=bar"}, // URL encoding may change order
		{"", ""},
		{"key=value%20with%20spaces", "key=value+with+spaces"}, // Space encoding
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			params := NewURLSearchParams(test.input)
			result := params.ToString()

			// Parse both to compare content since order might differ
			if test.expected == "" && result == "" {
				return // Both empty, test passes
			}

			expectedParams := NewURLSearchParams(test.expected)
			resultParams := NewURLSearchParams(result)

			// Compare all keys and values
			expectedKeys := expectedParams.Keys()
			resultKeys := resultParams.Keys()

			if len(expectedKeys) != len(resultKeys) {
				t.Errorf("Different number of keys: expected %v, got %v", expectedKeys, resultKeys)
				return
			}

			for _, key := range expectedKeys {
				if !resultParams.Has(key) {
					t.Errorf("Missing key %s in result", key)
				}
				expectedVals := expectedParams.GetAll(key)
				resultVals := resultParams.GetAll(key)
				if !reflect.DeepEqual(expectedVals, resultVals) {
					t.Errorf("Different values for key %s: expected %v, got %v", key, expectedVals, resultVals)
				}
			}
		})
	}
}

func TestURLSearchParams_Keys(t *testing.T) {
	params := NewURLSearchParams("foo=bar&baz=qux&foo=baz")
	keys := params.Keys()

	// Should return unique keys in order of first appearance
	expected := []string{"foo", "baz"}
	if !reflect.DeepEqual(keys, expected) {
		t.Errorf("Expected keys %v, got %v", expected, keys)
	}
}

func TestURLSearchParams_Values(t *testing.T) {
	params := NewURLSearchParams("foo=bar&baz=qux&foo=baz")
	values := params.Values()

	// Should return all values in order
	expected := []string{"bar", "qux", "baz"}
	if !reflect.DeepEqual(values, expected) {
		t.Errorf("Expected values %v, got %v", expected, values)
	}
}

func TestURLSearchParams_Entries(t *testing.T) {
	params := NewURLSearchParams("foo=bar&baz=qux&foo=baz")
	entries := params.Entries()

	expected := [][2]string{
		{"foo", "bar"},
		{"baz", "qux"},
		{"foo", "baz"},
	}
	if !reflect.DeepEqual(entries, expected) {
		t.Errorf("Expected entries %v, got %v", expected, entries)
	}
}

func TestURLSearchParams_ForEach(t *testing.T) {
	params := NewURLSearchParams("foo=bar&baz=qux")

	var results []string
	params.ForEach(func(value, name string) {
		results = append(results, name+"="+value)
	})

	expected := []string{"foo=bar", "baz=qux"}
	if !reflect.DeepEqual(results, expected) {
		t.Errorf("Expected forEach results %v, got %v", expected, results)
	}
}

func TestURLSearchParams_Size(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"", 0},
		{"foo=bar", 1},
		{"foo=bar&baz=qux", 2},
		{"foo=bar&foo=baz&qux=quux", 3},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			params := NewURLSearchParams(test.input)
			if params.Size() != test.expected {
				t.Errorf("Expected size %d, got %d", test.expected, params.Size())
			}
		})
	}
}

func TestURLSearchParams_EncodingDecoding(t *testing.T) {
	params := NewURLSearchParams()

	// Test special characters
	params.Append("key with spaces", "value with spaces")
	params.Append("special", "!@#$%^&*()")
	params.Append("unicode", "café")

	// Convert to string and back
	queryString := params.ToString()
	newParams := NewURLSearchParams(queryString)

	if newParams.Get("key with spaces") != "value with spaces" {
		t.Error("Failed to preserve spaces in key and value")
	}

	if newParams.Get("special") != "!@#$%^&*()" {
		t.Error("Failed to preserve special characters")
	}

	if newParams.Get("unicode") != "café" {
		t.Error("Failed to preserve unicode characters")
	}
}

func TestURLSearchParams_EdgeCases(t *testing.T) {
	// Test empty values
	params := NewURLSearchParams("foo=&bar=baz&=empty")

	if params.Get("foo") != "" {
		t.Errorf("Expected empty value for foo, got %s", params.Get("foo"))
	}

	if params.Get("bar") != "baz" {
		t.Errorf("Expected bar to be baz, got %s", params.Get("bar"))
	}

	// Test duplicate equals signs
	params = NewURLSearchParams("foo=bar=baz")
	if params.Get("foo") != "bar=baz" {
		t.Errorf("Expected foo to be bar=baz, got %s", params.Get("foo"))
	}
}
