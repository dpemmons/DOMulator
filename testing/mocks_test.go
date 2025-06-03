package testing

import (
	"net/http"
	"strings"
	"testing"
)

func TestNewNetworkMocks(t *testing.T) {
	mocks := NewNetworkMocks()

	if mocks == nil {
		t.Fatal("NewNetworkMocks() returned nil")
	}

	if mocks.routes == nil {
		t.Error("NetworkMocks.routes should not be nil")
	}

	if len(mocks.routes) != 0 {
		t.Error("New NetworkMocks should have empty routes")
	}

	if mocks.Count() != 0 {
		t.Error("New NetworkMocks should have count of 0")
	}
}

func TestNetworkMocks_MockGET(t *testing.T) {
	mocks := NewNetworkMocks()

	response := MockResponse{
		Status: http.StatusOK,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: `{"message": "success"}`,
	}

	mocks.MockGET("/api/test", response)

	if mocks.Count() != 1 {
		t.Errorf("Expected 1 mock route, got %d", mocks.Count())
	}

	if !mocks.HasMock("GET", "/api/test") {
		t.Error("Should have GET mock for /api/test")
	}

	// Test matching
	match := mocks.Match("GET", "/api/test")
	if match == nil {
		t.Fatal("Should match GET /api/test")
	}

	if match.Status != http.StatusOK {
		t.Errorf("Expected status 200, got %d", match.Status)
	}

	if match.Headers["Content-Type"] != "application/json" {
		t.Errorf("Expected Content-Type 'application/json', got '%s'", match.Headers["Content-Type"])
	}

	if match.Body != `{"message": "success"}` {
		t.Errorf("Expected JSON body, got '%s'", match.Body)
	}
}

func TestNetworkMocks_MockPOST(t *testing.T) {
	mocks := NewNetworkMocks()

	response := MockResponse{
		Status: http.StatusCreated,
		Body:   "Created successfully",
	}

	mocks.MockPOST("/api/create", response)

	if !mocks.HasMock("POST", "/api/create") {
		t.Error("Should have POST mock for /api/create")
	}

	match := mocks.Match("POST", "/api/create")
	if match == nil {
		t.Fatal("Should match POST /api/create")
	}

	if match.Status != http.StatusCreated {
		t.Errorf("Expected status 201, got %d", match.Status)
	}
}

func TestNetworkMocks_MockPUT(t *testing.T) {
	mocks := NewNetworkMocks()

	response := MockResponse{
		Status: http.StatusOK,
		Body:   "Updated",
	}

	mocks.MockPUT("/api/update", response)

	if !mocks.HasMock("PUT", "/api/update") {
		t.Error("Should have PUT mock for /api/update")
	}

	match := mocks.Match("PUT", "/api/update")
	if match == nil {
		t.Fatal("Should match PUT /api/update")
	}

	if match.Status != http.StatusOK {
		t.Errorf("Expected status 200, got %d", match.Status)
	}
}

func TestNetworkMocks_MockDELETE(t *testing.T) {
	mocks := NewNetworkMocks()

	response := MockResponse{
		Status: http.StatusNoContent,
	}

	mocks.MockDELETE("/api/delete", response)

	if !mocks.HasMock("DELETE", "/api/delete") {
		t.Error("Should have DELETE mock for /api/delete")
	}

	match := mocks.Match("DELETE", "/api/delete")
	if match == nil {
		t.Fatal("Should match DELETE /api/delete")
	}

	if match.Status != http.StatusNoContent {
		t.Errorf("Expected status 204, got %d", match.Status)
	}
}

func TestNetworkMocks_Mock_GenericMethod(t *testing.T) {
	mocks := NewNetworkMocks()

	response := MockResponse{
		Status: http.StatusOK,
		Body:   "PATCH response",
	}

	mocks.Mock("PATCH", "/api/patch", response)

	if !mocks.HasMock("PATCH", "/api/patch") {
		t.Error("Should have PATCH mock for /api/patch")
	}

	match := mocks.Match("PATCH", "/api/patch")
	if match == nil {
		t.Fatal("Should match PATCH /api/patch")
	}

	if match.Status != http.StatusOK {
		t.Errorf("Expected status 200, got %d", match.Status)
	}

	if match.Body != "PATCH response" {
		t.Errorf("Expected 'PATCH response', got '%s'", match.Body)
	}
}

func TestNetworkMocks_Mock_CaseInsensitiveMethod(t *testing.T) {
	mocks := NewNetworkMocks()

	response := MockResponse{
		Status: http.StatusOK,
		Body:   "lowercase method",
	}

	// Add with lowercase method
	mocks.Mock("get", "/test", response)

	// Should match uppercase
	match := mocks.Match("GET", "/test")
	if match == nil {
		t.Fatal("Should match case-insensitive method")
	}

	if match.Body != "lowercase method" {
		t.Errorf("Expected 'lowercase method', got '%s'", match.Body)
	}
}

func TestNetworkMocks_Match_ExactPath(t *testing.T) {
	mocks := NewNetworkMocks()

	mocks.MockGET("/exact/path", MockResponse{
		Status: http.StatusOK,
		Body:   "exact match",
	})

	// Should match exact path
	match := mocks.Match("GET", "/exact/path")
	if match == nil {
		t.Fatal("Should match exact path")
	}

	// Should not match different path
	match = mocks.Match("GET", "/exact/different")
	if match != nil {
		t.Error("Should not match different path")
	}

	// Should not match partial path
	match = mocks.Match("GET", "/exact")
	if match != nil {
		t.Error("Should not match partial path")
	}
}

func TestNetworkMocks_Match_WildcardPattern(t *testing.T) {
	mocks := NewNetworkMocks()

	mocks.MockGET("/api/*", MockResponse{
		Status: http.StatusOK,
		Body:   "wildcard match",
	})

	// Should match paths with wildcard
	match := mocks.Match("GET", "/api/users")
	if match == nil {
		t.Fatal("Should match wildcard pattern")
	}

	match = mocks.Match("GET", "/api/users/123")
	if match == nil {
		t.Fatal("Should match nested wildcard pattern")
	}

	match = mocks.Match("GET", "/api/")
	if match == nil {
		t.Fatal("Should match wildcard with trailing slash")
	}

	// Should not match different prefix
	match = mocks.Match("GET", "/different/users")
	if match != nil {
		t.Error("Should not match different prefix")
	}
}

func TestNetworkMocks_Match_RegexPattern(t *testing.T) {
	mocks := NewNetworkMocks()

	// Add regex pattern for user IDs
	mocks.MockGET(`/api/users/\d+`, MockResponse{
		Status: http.StatusOK,
		Body:   "user details",
	})

	// Should match numeric user IDs
	match := mocks.Match("GET", "/api/users/123")
	if match == nil {
		t.Fatal("Should match regex pattern for /api/users/123")
	}

	match = mocks.Match("GET", "/api/users/456")
	if match == nil {
		t.Fatal("Should match regex pattern for /api/users/456")
	}

	// Should not match non-numeric IDs
	match = mocks.Match("GET", "/api/users/abc")
	if match != nil {
		t.Error("Should not match non-numeric user ID")
	}

	// Should not match without ID
	match = mocks.Match("GET", "/api/users/")
	if match != nil {
		t.Error("Should not match user path without ID")
	}
}

func TestNetworkMocks_Match_ComplexRegex(t *testing.T) {
	mocks := NewNetworkMocks()

	// Add complex regex for query parameters
	mocks.MockGET(`/search\?q=\w+`, MockResponse{
		Status: http.StatusOK,
		Body:   "search results",
	})

	// Should match search with query
	match := mocks.Match("GET", "/search?q=test")
	if match == nil {
		t.Fatal("Should match search query regex")
	}

	match = mocks.Match("GET", "/search?q=javascript")
	if match == nil {
		t.Fatal("Should match different search query")
	}

	// Should not match without query
	match = mocks.Match("GET", "/search")
	if match != nil {
		t.Error("Should not match search without query")
	}
}

func TestNetworkMocks_Match_MethodSpecific(t *testing.T) {
	mocks := NewNetworkMocks()

	// Add different responses for same path, different methods
	mocks.MockGET("/resource", MockResponse{
		Status: http.StatusOK,
		Body:   "GET resource",
	})

	mocks.MockPOST("/resource", MockResponse{
		Status: http.StatusCreated,
		Body:   "POST resource",
	})

	// Should match specific methods
	getMatch := mocks.Match("GET", "/resource")
	if getMatch == nil {
		t.Fatal("Should match GET /resource")
	}
	if getMatch.Body != "GET resource" {
		t.Errorf("Expected 'GET resource', got '%s'", getMatch.Body)
	}

	postMatch := mocks.Match("POST", "/resource")
	if postMatch == nil {
		t.Fatal("Should match POST /resource")
	}
	if postMatch.Body != "POST resource" {
		t.Errorf("Expected 'POST resource', got '%s'", postMatch.Body)
	}

	// Should not match different method
	putMatch := mocks.Match("PUT", "/resource")
	if putMatch != nil {
		t.Error("Should not match PUT /resource")
	}
}

func TestNetworkMocks_Match_FirstMatchWins(t *testing.T) {
	mocks := NewNetworkMocks()

	// Add multiple patterns that could match
	mocks.MockGET("/api/*", MockResponse{
		Status: http.StatusOK,
		Body:   "wildcard first",
	})

	mocks.MockGET("/api/users", MockResponse{
		Status: http.StatusOK,
		Body:   "exact second",
	})

	// First pattern should win
	match := mocks.Match("GET", "/api/users")
	if match == nil {
		t.Fatal("Should match some pattern")
	}

	if match.Body != "wildcard first" {
		t.Errorf("Expected first pattern to win, got '%s'", match.Body)
	}
}

func TestNetworkMocks_Match_NoMatch(t *testing.T) {
	mocks := NewNetworkMocks()

	mocks.MockGET("/api/test", MockResponse{Status: http.StatusOK})

	// Should not match different method
	match := mocks.Match("POST", "/api/test")
	if match != nil {
		t.Error("Should not match different method")
	}

	// Should not match different path
	match = mocks.Match("GET", "/api/different")
	if match != nil {
		t.Error("Should not match different path")
	}

	// Should not match when no mocks
	emptyMocks := NewNetworkMocks()
	match = emptyMocks.Match("GET", "/anything")
	if match != nil {
		t.Error("Should not match when no mocks configured")
	}
}

func TestNetworkMocks_Clear(t *testing.T) {
	mocks := NewNetworkMocks()

	// Add some mocks
	mocks.MockGET("/test1", MockResponse{Status: http.StatusOK})
	mocks.MockPOST("/test2", MockResponse{Status: http.StatusCreated})

	if mocks.Count() != 2 {
		t.Errorf("Expected 2 mocks, got %d", mocks.Count())
	}

	// Clear mocks
	mocks.Clear()

	if mocks.Count() != 0 {
		t.Errorf("Expected 0 mocks after clear, got %d", mocks.Count())
	}

	// Should not match after clear
	match := mocks.Match("GET", "/test1")
	if match != nil {
		t.Error("Should not match after clear")
	}
}

func TestNetworkMocks_Reset(t *testing.T) {
	mocks := NewNetworkMocks()

	// Add some mocks
	mocks.MockGET("/test1", MockResponse{Status: http.StatusOK})
	mocks.MockPOST("/test2", MockResponse{Status: http.StatusCreated})

	if mocks.Count() != 2 {
		t.Errorf("Expected 2 mocks, got %d", mocks.Count())
	}

	// Reset should work the same as Clear
	mocks.Reset()

	if mocks.Count() != 0 {
		t.Errorf("Expected 0 mocks after reset, got %d", mocks.Count())
	}

	// Should not match after reset
	match := mocks.Match("GET", "/test1")
	if match != nil {
		t.Error("Should not match after reset")
	}
}

func TestNetworkMocks_HasMock(t *testing.T) {
	mocks := NewNetworkMocks()

	// Should not have mocks initially
	if mocks.HasMock("GET", "/test") {
		t.Error("Should not have mock initially")
	}

	// Add mock
	mocks.MockGET("/test", MockResponse{Status: http.StatusOK})

	// Should have mock now
	if !mocks.HasMock("GET", "/test") {
		t.Error("Should have mock after adding")
	}

	// Should not have different method
	if mocks.HasMock("POST", "/test") {
		t.Error("Should not have POST mock")
	}

	// Should not have different path
	if mocks.HasMock("GET", "/different") {
		t.Error("Should not have different path mock")
	}
}

func TestNetworkMocks_Count(t *testing.T) {
	mocks := NewNetworkMocks()

	// Should start with 0
	if mocks.Count() != 0 {
		t.Errorf("Expected count 0, got %d", mocks.Count())
	}

	// Add mocks one by one
	mocks.MockGET("/test1", MockResponse{Status: http.StatusOK})
	if mocks.Count() != 1 {
		t.Errorf("Expected count 1, got %d", mocks.Count())
	}

	mocks.MockPOST("/test2", MockResponse{Status: http.StatusCreated})
	if mocks.Count() != 2 {
		t.Errorf("Expected count 2, got %d", mocks.Count())
	}

	mocks.MockPUT("/test3", MockResponse{Status: http.StatusOK})
	if mocks.Count() != 3 {
		t.Errorf("Expected count 3, got %d", mocks.Count())
	}

	// Clear and check count
	mocks.Clear()
	if mocks.Count() != 0 {
		t.Errorf("Expected count 0 after clear, got %d", mocks.Count())
	}
}

func TestNetworkMocks_InvalidRegex(t *testing.T) {
	mocks := NewNetworkMocks()

	// Add pattern with invalid regex
	mocks.MockGET("[invalid", MockResponse{Status: http.StatusOK})

	// Should still work with exact match since regex compilation failed
	match := mocks.Match("GET", "[invalid")
	if match == nil {
		t.Fatal("Should match as exact string when regex fails")
	}

	// Should not match as regex
	match = mocks.Match("GET", "invalid")
	if match != nil {
		t.Error("Should not match as regex when compilation fails")
	}
}

func TestNetworkMocks_EmptyResponse(t *testing.T) {
	mocks := NewNetworkMocks()

	// Test empty response
	mocks.MockGET("/empty", MockResponse{})

	match := mocks.Match("GET", "/empty")
	if match == nil {
		t.Fatal("Should match empty response")
	}

	if match.Status != 0 {
		t.Errorf("Expected status 0 for empty response, got %d", match.Status)
	}

	if match.Body != "" {
		t.Errorf("Expected empty body, got '%s'", match.Body)
	}

	if match.Headers != nil {
		t.Error("Expected nil headers for empty response")
	}
}

func TestNetworkMocks_MultipleHeaders(t *testing.T) {
	mocks := NewNetworkMocks()

	response := MockResponse{
		Status: http.StatusOK,
		Headers: map[string]string{
			"Content-Type":  "application/json",
			"Cache-Control": "no-cache",
			"X-Custom":      "custom-value",
			"Authorization": "Bearer token123",
		},
		Body: `{"data": "test"}`,
	}

	mocks.MockGET("/headers", response)

	match := mocks.Match("GET", "/headers")
	if match == nil {
		t.Fatal("Should match headers test")
	}

	if match.Headers["Content-Type"] != "application/json" {
		t.Error("Should preserve Content-Type header")
	}

	if match.Headers["Cache-Control"] != "no-cache" {
		t.Error("Should preserve Cache-Control header")
	}

	if match.Headers["X-Custom"] != "custom-value" {
		t.Error("Should preserve custom header")
	}

	if match.Headers["Authorization"] != "Bearer token123" {
		t.Error("Should preserve Authorization header")
	}
}

// Integration test with HTTP client
func TestNetworkMocks_Integration(t *testing.T) {
	mocks := NewNetworkMocks()
	client := NewHTTPClient(mocks)

	// Setup multiple mocks
	mocks.MockGET("/api/users", MockResponse{
		Status: http.StatusOK,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: `[{"id": 1, "name": "Alice"}, {"id": 2, "name": "Bob"}]`,
	})

	mocks.MockPOST("/api/users", MockResponse{
		Status:  http.StatusCreated,
		Headers: map[string]string{"Location": "/api/users/3"},
		Body:    `{"id": 3, "name": "Charlie"}`,
	})

	mocks.MockGET(`/api/users/\d+`, MockResponse{
		Status: http.StatusOK,
		Body:   `{"id": 1, "name": "Alice"}`,
	})

	// Test GET users
	response := client.GET("/api/users")
	if response.Error != nil {
		t.Errorf("GET users failed: %v", response.Error)
	}
	if response.Status != http.StatusOK {
		t.Errorf("Expected status 200, got %d", response.Status)
	}
	if !strings.Contains(response.Body, "Alice") {
		t.Error("Response should contain Alice")
	}

	// Test POST user
	response = client.POST("/api/users", `{"name": "Charlie"}`)
	if response.Error != nil {
		t.Errorf("POST user failed: %v", response.Error)
	}
	if response.Status != http.StatusCreated {
		t.Errorf("Expected status 201, got %d", response.Status)
	}
	if response.Headers["Location"] != "/api/users/3" {
		t.Error("Should have Location header")
	}

	// Test GET specific user (regex pattern)
	response = client.GET("/api/users/123")
	if response.Error != nil {
		t.Errorf("GET specific user failed: %v", response.Error)
	}
	if response.Status != http.StatusOK {
		t.Errorf("Expected status 200, got %d", response.Status)
	}
}
