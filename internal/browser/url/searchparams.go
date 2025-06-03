package url

import (
	"net/url"
	"sort"
	"strings"
)

// URLSearchParams represents a URLSearchParams object following the Web API specification
type URLSearchParams struct {
	params []param
}

// param represents a single name-value pair
type param struct {
	name  string
	value string
}

// NewURLSearchParams creates a new URLSearchParams object from a query string
func NewURLSearchParams(init ...string) *URLSearchParams {
	params := &URLSearchParams{
		params: make([]param, 0),
	}

	if len(init) > 0 && init[0] != "" {
		query := init[0]
		// Remove leading ? if present
		query = strings.TrimPrefix(query, "?")

		// Parse query string manually to preserve order
		pairs := strings.Split(query, "&")
		for _, pair := range pairs {
			if pair == "" {
				continue
			}
			parts := strings.SplitN(pair, "=", 2)
			name := parts[0]
			value := ""
			if len(parts) > 1 {
				value = parts[1]
			}
			// URL decode
			decodedName, _ := url.QueryUnescape(name)
			decodedValue, _ := url.QueryUnescape(value)
			params.params = append(params.params, param{name: decodedName, value: decodedValue})
		}
	}

	return params
}

// Append adds a new name-value pair to the end of the list
func (p *URLSearchParams) Append(name, value string) {
	p.params = append(p.params, param{name: name, value: value})
}

// Delete removes all name-value pairs whose name is name
func (p *URLSearchParams) Delete(name string) {
	newParams := make([]param, 0, len(p.params))
	for _, param := range p.params {
		if param.name != name {
			newParams = append(newParams, param)
		}
	}
	p.params = newParams
}

// Get returns the first value associated to the given search parameter
func (p *URLSearchParams) Get(name string) string {
	for _, param := range p.params {
		if param.name == name {
			return param.value
		}
	}
	return ""
}

// GetAll returns all values associated to the given search parameter
func (p *URLSearchParams) GetAll(name string) []string {
	var values []string
	for _, param := range p.params {
		if param.name == name {
			values = append(values, param.value)
		}
	}
	return values
}

// Has returns true if there is at least one name-value pair whose name is name
func (p *URLSearchParams) Has(name string) bool {
	for _, param := range p.params {
		if param.name == name {
			return true
		}
	}
	return false
}

// Set sets the value associated to a given search parameter
// If there were several matching values, this method deletes the others
// If the search parameter doesn't exist, this method creates it
func (p *URLSearchParams) Set(name, value string) {
	// First remove all existing entries with this name
	p.Delete(name)
	// Then add the new entry
	p.Append(name, value)
}

// Sort sorts all name-value pairs by their names
func (p *URLSearchParams) Sort() {
	sort.Slice(p.params, func(i, j int) bool {
		return p.params[i].name < p.params[j].name
	})
}

// ToString returns a query string suitable for use in a URL
func (p *URLSearchParams) ToString() string {
	if len(p.params) == 0 {
		return ""
	}

	values := url.Values{}
	for _, param := range p.params {
		values.Add(param.name, param.value)
	}
	return values.Encode()
}

// String implements the Stringer interface
func (p *URLSearchParams) String() string {
	return p.ToString()
}

// Keys returns an iterator over all names in the search parameters
func (p *URLSearchParams) Keys() []string {
	var keys []string
	seen := make(map[string]bool)

	for _, param := range p.params {
		if !seen[param.name] {
			keys = append(keys, param.name)
			seen[param.name] = true
		}
	}
	return keys
}

// Values returns an iterator over all values in the search parameters
func (p *URLSearchParams) Values() []string {
	var values []string
	for _, param := range p.params {
		values = append(values, param.value)
	}
	return values
}

// Entries returns an iterator over all name-value pairs in the search parameters
func (p *URLSearchParams) Entries() [][2]string {
	var entries [][2]string
	for _, param := range p.params {
		entries = append(entries, [2]string{param.name, param.value})
	}
	return entries
}

// ForEach executes a provided function once for each name-value pair
func (p *URLSearchParams) ForEach(fn func(value, name string)) {
	for _, param := range p.params {
		fn(param.value, param.name)
	}
}

// Size returns the number of search parameter name-value pairs
func (p *URLSearchParams) Size() int {
	return len(p.params)
}
