package testing

import (
	"regexp"
	"strings"
)

// NetworkMocks handles HTTP request mocking for testing
type NetworkMocks struct {
	routes []MockRoute
}

// MockRoute represents a single mock route configuration
type MockRoute struct {
	Method   string
	Pattern  string
	Response MockResponse
	regex    *regexp.Regexp
}

// MockResponse represents a mock HTTP response
type MockResponse struct {
	Status  int
	Headers map[string]string
	Body    string
}

// NewNetworkMocks creates a new network mocking system
func NewNetworkMocks() *NetworkMocks {
	return &NetworkMocks{
		routes: make([]MockRoute, 0),
	}
}

// MockGET adds a mock for GET requests to the specified pattern
func (m *NetworkMocks) MockGET(pattern string, response MockResponse) {
	m.addRoute("GET", pattern, response)
}

// MockPOST adds a mock for POST requests to the specified pattern
func (m *NetworkMocks) MockPOST(pattern string, response MockResponse) {
	m.addRoute("POST", pattern, response)
}

// MockPUT adds a mock for PUT requests to the specified pattern
func (m *NetworkMocks) MockPUT(pattern string, response MockResponse) {
	m.addRoute("PUT", pattern, response)
}

// MockDELETE adds a mock for DELETE requests to the specified pattern
func (m *NetworkMocks) MockDELETE(pattern string, response MockResponse) {
	m.addRoute("DELETE", pattern, response)
}

// Mock adds a mock for any HTTP method to the specified pattern
func (m *NetworkMocks) Mock(method, pattern string, response MockResponse) {
	m.addRoute(strings.ToUpper(method), pattern, response)
}

// addRoute adds a new mock route with optional regex compilation
func (m *NetworkMocks) addRoute(method, pattern string, response MockResponse) {
	route := MockRoute{
		Method:   method,
		Pattern:  pattern,
		Response: response,
	}

	// If pattern contains regex characters, compile it
	if strings.ContainsAny(pattern, "[]{}()*+?.\\^$|") {
		if regex, err := regexp.Compile(pattern); err == nil {
			route.regex = regex
		}
	}

	m.routes = append(m.routes, route)
}

// Match finds a matching mock response for the given method and path
func (m *NetworkMocks) Match(method, path string) *MockResponse {
	method = strings.ToUpper(method)

	for _, route := range m.routes {
		if route.Method != method {
			continue
		}

		// Try regex match first if available
		if route.regex != nil {
			if route.regex.MatchString(path) {
				return &route.Response
			}
			continue
		}

		// Fallback to exact string match
		if route.Pattern == path {
			return &route.Response
		}

		// Also try simple wildcard matching for convenience
		if strings.HasSuffix(route.Pattern, "*") {
			prefix := strings.TrimSuffix(route.Pattern, "*")
			if strings.HasPrefix(path, prefix) {
				return &route.Response
			}
		}
	}

	return nil
}

// Clear removes all mock routes
func (m *NetworkMocks) Clear() {
	m.routes = m.routes[:0]
}

// Reset is an alias for Clear for convenience
func (m *NetworkMocks) Reset() {
	m.Clear()
}

// Count returns the number of configured mock routes
func (m *NetworkMocks) Count() int {
	return len(m.routes)
}

// HasMock checks if a mock exists for the given method and pattern
func (m *NetworkMocks) HasMock(method, pattern string) bool {
	method = strings.ToUpper(method)
	for _, route := range m.routes {
		if route.Method == method && route.Pattern == pattern {
			return true
		}
	}
	return false
}
