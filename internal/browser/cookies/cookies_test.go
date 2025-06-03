package cookies

import (
	"net/url"
	"testing"
	"time"
)

func TestCookieCreation(t *testing.T) {
	cookie := &Cookie{
		Name:     "test",
		Value:    "value",
		Path:     "/",
		Domain:   "example.com",
		Secure:   true,
		HttpOnly: true,
		SameSite: SameSiteStrict,
	}

	if cookie.Name != "test" {
		t.Errorf("Expected name 'test', got '%s'", cookie.Name)
	}
	if cookie.Value != "value" {
		t.Errorf("Expected value 'value', got '%s'", cookie.Value)
	}
	if cookie.SameSite.String() != "Strict" {
		t.Errorf("Expected SameSite 'Strict', got '%s'", cookie.SameSite.String())
	}
}

func TestCookieExpiration(t *testing.T) {
	// Test MaxAge < 0 (expired)
	expiredCookie := &Cookie{
		Name:   "expired",
		Value:  "value",
		MaxAge: -1,
	}
	if !expiredCookie.IsExpired() {
		t.Error("Cookie with MaxAge < 0 should be expired")
	}

	// Test MaxAge > 0 (not expired)
	validCookie := &Cookie{
		Name:   "valid",
		Value:  "value",
		MaxAge: 3600,
	}
	if validCookie.IsExpired() {
		t.Error("Cookie with MaxAge > 0 should not be expired")
	}

	// Test past Expires date
	pastCookie := &Cookie{
		Name:    "past",
		Value:   "value",
		Expires: time.Now().Add(-time.Hour),
	}
	if !pastCookie.IsExpired() {
		t.Error("Cookie with past Expires date should be expired")
	}

	// Test future Expires date
	futureCookie := &Cookie{
		Name:    "future",
		Value:   "value",
		Expires: time.Now().Add(time.Hour),
	}
	if futureCookie.IsExpired() {
		t.Error("Cookie with future Expires date should not be expired")
	}

	// Test MaxAge takes precedence over Expires
	precedenceCookie := &Cookie{
		Name:    "precedence",
		Value:   "value",
		MaxAge:  3600,
		Expires: time.Now().Add(-time.Hour),
	}
	if precedenceCookie.IsExpired() {
		t.Error("Cookie with positive MaxAge should not be expired even with past Expires")
	}
}

func TestCookieString(t *testing.T) {
	cookie := &Cookie{
		Name:     "sessionid",
		Value:    "abc123",
		Path:     "/",
		Domain:   "example.com",
		MaxAge:   3600,
		Secure:   true,
		HttpOnly: true,
		SameSite: SameSiteStrict,
	}

	cookieStr := cookie.String()
	expected := "sessionid=abc123; Path=/; Domain=example.com; Max-Age=3600; Secure; HttpOnly; SameSite=Strict"
	if cookieStr != expected {
		t.Errorf("Expected cookie string '%s', got '%s'", expected, cookieStr)
	}
}

func TestSameSiteString(t *testing.T) {
	tests := []struct {
		sameSite SameSite
		expected string
	}{
		{SameSiteNone, "None"},
		{SameSiteLax, "Lax"},
		{SameSiteStrict, "Strict"},
	}

	for _, test := range tests {
		if test.sameSite.String() != test.expected {
			t.Errorf("Expected SameSite %d to string '%s', got '%s'", test.sameSite, test.expected, test.sameSite.String())
		}
	}
}

func TestCookieManagerCreation(t *testing.T) {
	cm := NewCookieManager("example.com", "/app")
	if cm.domain != "example.com" {
		t.Errorf("Expected domain 'example.com', got '%s'", cm.domain)
	}
	if cm.path != "/app" {
		t.Errorf("Expected path '/app', got '%s'", cm.path)
	}

	// Test default path
	cm2 := NewCookieManager("example.com", "")
	if cm2.path != "/" {
		t.Errorf("Expected default path '/', got '%s'", cm2.path)
	}
}

func TestSetAndGetCookie(t *testing.T) {
	cm := NewCookieManager("example.com", "/")

	cookie := &Cookie{
		Name:  "test",
		Value: "value",
	}

	err := cm.SetCookie(cookie)
	if err != nil {
		t.Errorf("Unexpected error setting cookie: %v", err)
	}

	retrieved := cm.GetCookie("test")
	if retrieved == nil {
		t.Error("Expected to retrieve cookie, got nil")
	}
	if retrieved.Name != "test" {
		t.Errorf("Expected name 'test', got '%s'", retrieved.Name)
	}
	if retrieved.Value != "value" {
		t.Errorf("Expected value 'value', got '%s'", retrieved.Value)
	}
}

func TestSetCookieDefaults(t *testing.T) {
	cm := NewCookieManager("example.com", "/app")

	cookie := &Cookie{
		Name:  "test",
		Value: "value",
		// No path or domain specified
	}

	err := cm.SetCookie(cookie)
	if err != nil {
		t.Errorf("Unexpected error setting cookie: %v", err)
	}

	retrieved := cm.GetCookie("test")
	if retrieved.Path != "/app" {
		t.Errorf("Expected default path '/app', got '%s'", retrieved.Path)
	}
	if retrieved.Domain != "example.com" {
		t.Errorf("Expected default domain 'example.com', got '%s'", retrieved.Domain)
	}
}

func TestSetCookieEmptyName(t *testing.T) {
	cm := NewCookieManager("example.com", "/")

	cookie := &Cookie{
		Name:  "",
		Value: "value",
	}

	err := cm.SetCookie(cookie)
	if err == nil {
		t.Error("Expected error for empty cookie name")
	}
}

func TestDeleteCookieViaMaxAge(t *testing.T) {
	cm := NewCookieManager("example.com", "/")

	// Set a cookie
	cookie := &Cookie{
		Name:  "test",
		Value: "value",
	}
	cm.SetCookie(cookie)

	// Delete it with MaxAge < 0
	deleteCookie := &Cookie{
		Name:   "test",
		Value:  "value",
		MaxAge: -1,
	}
	err := cm.SetCookie(deleteCookie)
	if err != nil {
		t.Errorf("Unexpected error deleting cookie: %v", err)
	}

	// Should be gone
	retrieved := cm.GetCookie("test")
	if retrieved != nil {
		t.Error("Expected cookie to be deleted")
	}
}

func TestDeleteCookie(t *testing.T) {
	cm := NewCookieManager("example.com", "/")

	cookie := &Cookie{
		Name:  "test",
		Value: "value",
	}
	cm.SetCookie(cookie)

	cm.DeleteCookie("test")

	retrieved := cm.GetCookie("test")
	if retrieved != nil {
		t.Error("Expected cookie to be deleted")
	}
}

func TestGetAllCookies(t *testing.T) {
	cm := NewCookieManager("example.com", "/")

	cookies := []*Cookie{
		{Name: "cookie1", Value: "value1", Path: "/"},
		{Name: "cookie2", Value: "value2", Path: "/app"},
		{Name: "cookie3", Value: "value3", Path: "/app/sub"},
	}

	for _, cookie := range cookies {
		cm.SetCookie(cookie)
	}

	allCookies := cm.GetAllCookies()
	if len(allCookies) != 3 {
		t.Errorf("Expected 3 cookies, got %d", len(allCookies))
	}

	// Check sorting (longest path first)
	if allCookies[0].Path != "/app/sub" {
		t.Errorf("Expected first cookie to have path '/app/sub', got '%s'", allCookies[0].Path)
	}
}

func TestFormatCookieString(t *testing.T) {
	cm := NewCookieManager("example.com", "/")

	cm.SetCookie(&Cookie{Name: "cookie1", Value: "value1"})
	cm.SetCookie(&Cookie{Name: "cookie2", Value: "value2"})

	cookieStr := cm.FormatCookieString()
	// Should contain both cookies
	if cookieStr == "" {
		t.Error("Expected non-empty cookie string")
	}

	// Basic format check (order may vary due to sorting)
	if !contains(cookieStr, "cookie1=value1") {
		t.Error("Expected cookie string to contain 'cookie1=value1'")
	}
	if !contains(cookieStr, "cookie2=value2") {
		t.Error("Expected cookie string to contain 'cookie2=value2'")
	}
}

func TestParseCookieString(t *testing.T) {
	cm := NewCookieManager("example.com", "/")

	// Simple cookie
	err := cm.ParseCookieString("sessionid=abc123")
	if err != nil {
		t.Errorf("Unexpected error parsing simple cookie: %v", err)
	}

	cookie := cm.GetCookie("sessionid")
	if cookie == nil {
		t.Error("Expected to find parsed cookie")
	}
	if cookie.Value != "abc123" {
		t.Errorf("Expected value 'abc123', got '%s'", cookie.Value)
	}

	// Cookie with attributes
	err = cm.ParseCookieString("secure_cookie=value; Path=/admin; Secure; HttpOnly; SameSite=Strict")
	if err != nil {
		t.Errorf("Unexpected error parsing cookie with attributes: %v", err)
	}

	secureC := cm.GetCookie("secure_cookie")
	if secureC == nil {
		t.Error("Expected to find secure cookie")
	}
	if secureC.Path != "/admin" {
		t.Errorf("Expected path '/admin', got '%s'", secureC.Path)
	}
	if !secureC.Secure {
		t.Error("Expected secure flag to be true")
	}
	if !secureC.HttpOnly {
		t.Error("Expected HttpOnly flag to be true")
	}
	if secureC.SameSite != SameSiteStrict {
		t.Errorf("Expected SameSite Strict, got %v", secureC.SameSite)
	}
}

func TestParseCookieStringWithMaxAge(t *testing.T) {
	cm := NewCookieManager("example.com", "/")

	err := cm.ParseCookieString("maxage_cookie=value; Max-Age=3600")
	if err != nil {
		t.Errorf("Unexpected error parsing cookie with Max-Age: %v", err)
	}

	cookie := cm.GetCookie("maxage_cookie")
	if cookie == nil {
		t.Error("Expected to find cookie with Max-Age")
	}
	if cookie.MaxAge != 3600 {
		t.Errorf("Expected MaxAge 3600, got %d", cookie.MaxAge)
	}
}

func TestParseCookieStringErrors(t *testing.T) {
	cm := NewCookieManager("example.com", "/")

	// Missing equals sign
	err := cm.ParseCookieString("invalid_cookie")
	if err == nil {
		t.Error("Expected error for cookie without equals sign")
	}

	// Empty name
	err = cm.ParseCookieString("=value")
	if err == nil {
		t.Error("Expected error for cookie with empty name")
	}
}

func TestDomainMatching(t *testing.T) {
	cm := NewCookieManager("example.com", "/")

	tests := []struct {
		cookieDomain  string
		requestDomain string
		shouldMatch   bool
	}{
		{"", "example.com", true},                 // No domain restriction
		{"example.com", "example.com", true},      // Exact match
		{".example.com", "example.com", true},     // Domain with leading dot
		{"example.com", "sub.example.com", false}, // Subdomain mismatch
		{".example.com", "sub.example.com", true}, // Subdomain match with dot
		{"other.com", "example.com", false},       // Different domain
		{"EXAMPLE.COM", "example.com", true},      // Case insensitive
	}

	for _, test := range tests {
		result := cm.domainMatches(test.cookieDomain, test.requestDomain)
		if result != test.shouldMatch {
			t.Errorf("Domain match test failed: cookie='%s', request='%s', expected=%v, got=%v",
				test.cookieDomain, test.requestDomain, test.shouldMatch, result)
		}
	}
}

func TestPathMatching(t *testing.T) {
	cm := NewCookieManager("example.com", "/")

	tests := []struct {
		cookiePath  string
		requestPath string
		shouldMatch bool
	}{
		{"", "/any/path", true},         // No path restriction
		{"/", "/any/path", true},        // Root path matches all
		{"/app", "/app", true},          // Exact match
		{"/app", "/app/page", true},     // Prefix match with slash
		{"/app/", "/app/page", true},    // Prefix match ending with slash
		{"/app", "/application", false}, // No slash separator
		{"/admin", "/app", false},       // Different path
		{"/app/admin", "/app", false},   // Longer path doesn't match shorter
	}

	for _, test := range tests {
		result := cm.pathMatches(test.cookiePath, test.requestPath)
		if result != test.shouldMatch {
			t.Errorf("Path match test failed: cookie='%s', request='%s', expected=%v, got=%v",
				test.cookiePath, test.requestPath, test.shouldMatch, result)
		}
	}
}

func TestGetCookiesForDomain(t *testing.T) {
	cm := NewCookieManager("example.com", "/")

	// Add cookies with different domains and paths
	cookies := []*Cookie{
		{Name: "cookie1", Value: "value1", Domain: "example.com", Path: "/"},
		{Name: "cookie2", Value: "value2", Domain: ".example.com", Path: "/app"},
		{Name: "cookie3", Value: "value3", Domain: "other.com", Path: "/"},
		{Name: "cookie4", Value: "value4", Domain: "example.com", Path: "/admin"},
	}

	for _, cookie := range cookies {
		cm.SetCookie(cookie)
	}

	// Get cookies for example.com, /app
	matchingCookies := cm.GetCookiesForDomain("example.com", "/app")

	expectedCount := 2 // cookie1 (/) and cookie2 (/app)
	if len(matchingCookies) != expectedCount {
		t.Errorf("Expected %d matching cookies, got %d", expectedCount, len(matchingCookies))
	}

	// Verify sorting (longest path first)
	if len(matchingCookies) >= 2 && len(matchingCookies[0].Path) < len(matchingCookies[1].Path) {
		t.Error("Expected cookies to be sorted by path length (longest first)")
	}
}

func TestGetCookieHeader(t *testing.T) {
	cm := NewCookieManager("example.com", "/")

	cm.SetCookie(&Cookie{Name: "cookie1", Value: "value1"})
	cm.SetCookie(&Cookie{Name: "cookie2", Value: "value2"})

	testURL, _ := url.Parse("https://example.com/")
	header := cm.GetCookieHeader(testURL)

	if header == "" {
		t.Error("Expected non-empty cookie header")
	}

	// Should contain both cookies
	if !contains(header, "cookie1=value1") {
		t.Error("Expected header to contain 'cookie1=value1'")
	}
	if !contains(header, "cookie2=value2") {
		t.Error("Expected header to contain 'cookie2=value2'")
	}
}

func TestCleanupExpiredCookies(t *testing.T) {
	cm := NewCookieManager("example.com", "/")

	// Add valid and expired cookies
	validCookie := &Cookie{Name: "valid", Value: "value", MaxAge: 3600}
	expiredCookie := &Cookie{Name: "expired", Value: "value", MaxAge: -1}

	cm.SetCookie(validCookie)
	cm.SetCookie(expiredCookie)

	// Should have both initially
	if len(cm.cookies) != 1 { // expired cookie should be deleted immediately
		t.Errorf("Expected 1 cookie after setting (expired deleted immediately), got %d", len(cm.cookies))
	}

	// Add a cookie that will expire
	futureExpired := &Cookie{
		Name:    "future_expired",
		Value:   "value",
		Expires: time.Now().Add(-time.Second), // Already expired
	}
	cm.SetCookie(futureExpired)

	cm.CleanupExpiredCookies()

	// Should only have the valid cookie
	remaining := cm.GetAllCookies()
	if len(remaining) != 1 {
		t.Errorf("Expected 1 cookie after cleanup, got %d", len(remaining))
	}
	if remaining[0].Name != "valid" {
		t.Errorf("Expected remaining cookie to be 'valid', got '%s'", remaining[0].Name)
	}
}

// Helper function to check if a string contains a substring
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > len(substr) &&
		(s[:len(substr)] == substr || s[len(s)-len(substr):] == substr ||
			indexOfSubstring(s, substr) >= 0))
}

func indexOfSubstring(s, substr string) int {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}
