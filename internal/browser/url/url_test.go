package url

import (
	"testing"
)

func TestNewURL_AbsoluteURL(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"https://example.com", "https://example.com"},
		{"https://example.com/path", "https://example.com/path"},
		{"https://example.com:8080/path?query=value#fragment", "https://example.com:8080/path?query=value#fragment"},
		{"http://user:pass@example.com/path", "http://user:pass@example.com/path"},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			url, err := NewURL(test.input)
			if err != nil {
				t.Fatalf("Expected no error, got %v", err)
			}
			if url.Href() != test.expected {
				t.Errorf("Expected %s, got %s", test.expected, url.Href())
			}
		})
	}
}

func TestNewURL_WithBase(t *testing.T) {
	tests := []struct {
		url      string
		base     string
		expected string
	}{
		{"path", "https://example.com", "https://example.com/path"},
		{"/absolute", "https://example.com/some/path", "https://example.com/absolute"},
		{"../relative", "https://example.com/some/path/", "https://example.com/some/relative"},
		{"?query=value", "https://example.com/path", "https://example.com/path?query=value"},
		{"#fragment", "https://example.com/path", "https://example.com/path#fragment"},
	}

	for _, test := range tests {
		t.Run(test.url+"_with_"+test.base, func(t *testing.T) {
			url, err := NewURL(test.url, test.base)
			if err != nil {
				t.Fatalf("Expected no error, got %v", err)
			}
			if url.Href() != test.expected {
				t.Errorf("Expected %s, got %s", test.expected, url.Href())
			}
		})
	}
}

func TestNewURL_InvalidURL(t *testing.T) {
	tests := []string{
		"relative/path",       // relative without base
		"://invalid",          // invalid scheme
		"ht tp://example.com", // space in scheme
	}

	for _, test := range tests {
		t.Run(test, func(t *testing.T) {
			_, err := NewURL(test)
			if err == nil {
				t.Errorf("Expected error for invalid URL: %s", test)
			}
		})
	}
}

func TestURL_Properties(t *testing.T) {
	url, err := NewURL("https://user:pass@example.com:8080/path/to/resource?query=value&foo=bar#section")
	if err != nil {
		t.Fatalf("Failed to create URL: %v", err)
	}

	tests := []struct {
		property string
		expected string
	}{
		{"href", "https://user:pass@example.com:8080/path/to/resource?query=value&foo=bar#section"},
		{"origin", "https://example.com:8080"},
		{"protocol", "https:"},
		{"host", "example.com:8080"},
		{"hostname", "example.com"},
		{"port", "8080"},
		{"pathname", "/path/to/resource"},
		{"search", "?query=value&foo=bar"},
		{"hash", "#section"},
	}

	for _, test := range tests {
		t.Run(test.property, func(t *testing.T) {
			var actual string
			switch test.property {
			case "href":
				actual = url.Href()
			case "origin":
				actual = url.Origin()
			case "protocol":
				actual = url.Protocol()
			case "host":
				actual = url.Host()
			case "hostname":
				actual = url.Hostname()
			case "port":
				actual = url.Port()
			case "pathname":
				actual = url.Pathname()
			case "search":
				actual = url.Search()
			case "hash":
				actual = url.Hash()
			}

			if actual != test.expected {
				t.Errorf("Expected %s to be %s, got %s", test.property, test.expected, actual)
			}
		})
	}
}

func TestURL_SetProperties(t *testing.T) {
	url, err := NewURL("https://example.com/path?query=value#fragment")
	if err != nil {
		t.Fatalf("Failed to create URL: %v", err)
	}

	// Test SetProtocol
	url.SetProtocol("http:")
	if url.Protocol() != "http:" {
		t.Errorf("Expected protocol to be http:, got %s", url.Protocol())
	}

	// Test SetHost
	url.SetHost("newhost.com:9000")
	if url.Host() != "newhost.com:9000" {
		t.Errorf("Expected host to be newhost.com:9000, got %s", url.Host())
	}

	// Test SetHostname (preserving port)
	url.SetHostname("another.com")
	if url.Hostname() != "another.com" {
		t.Errorf("Expected hostname to be another.com, got %s", url.Hostname())
	}
	if url.Port() != "9000" {
		t.Errorf("Expected port to be preserved as 9000, got %s", url.Port())
	}

	// Test SetPort
	url.SetPort("8080")
	if url.Port() != "8080" {
		t.Errorf("Expected port to be 8080, got %s", url.Port())
	}

	// Test SetPathname
	url.SetPathname("/new/path")
	if url.Pathname() != "/new/path" {
		t.Errorf("Expected pathname to be /new/path, got %s", url.Pathname())
	}

	// Test SetSearch
	url.SetSearch("?new=query")
	if url.Search() != "?new=query" {
		t.Errorf("Expected search to be ?new=query, got %s", url.Search())
	}

	// Test SetHash
	url.SetHash("#newhash")
	if url.Hash() != "#newhash" {
		t.Errorf("Expected hash to be #newhash, got %s", url.Hash())
	}
}

func TestURL_SetPortValidation(t *testing.T) {
	url, err := NewURL("https://example.com:8080")
	if err != nil {
		t.Fatalf("Failed to create URL: %v", err)
	}

	// Test invalid port (should be ignored)
	url.SetPort("99999")
	if url.Port() != "8080" {
		t.Errorf("Expected invalid port to be ignored, but port changed to %s", url.Port())
	}

	// Test negative port (should be ignored)
	url.SetPort("-1")
	if url.Port() != "8080" {
		t.Errorf("Expected negative port to be ignored, but port changed to %s", url.Port())
	}

	// Test empty port (should remove port)
	url.SetPort("")
	if url.Port() != "" {
		t.Errorf("Expected empty port to remove port, but got %s", url.Port())
	}
}

func TestURL_Origin(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"https://example.com", "https://example.com"},
		{"https://example.com:8080", "https://example.com:8080"},
		{"http://localhost:3000", "http://localhost:3000"},
		{"file:///path/to/file", "null"},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			url, err := NewURL(test.input)
			if err != nil {
				t.Fatalf("Failed to create URL: %v", err)
			}
			if url.Origin() != test.expected {
				t.Errorf("Expected origin to be %s, got %s", test.expected, url.Origin())
			}
		})
	}
}

func TestURL_SearchParams(t *testing.T) {
	url, err := NewURL("https://example.com?foo=bar&baz=qux")
	if err != nil {
		t.Fatalf("Failed to create URL: %v", err)
	}

	params := url.SearchParams()
	if params.Get("foo") != "bar" {
		t.Errorf("Expected foo to be bar, got %s", params.Get("foo"))
	}
	if params.Get("baz") != "qux" {
		t.Errorf("Expected baz to be qux, got %s", params.Get("baz"))
	}
}

func TestURL_ToStringAndJSON(t *testing.T) {
	urlStr := "https://example.com/path?query=value#fragment"
	url, err := NewURL(urlStr)
	if err != nil {
		t.Fatalf("Failed to create URL: %v", err)
	}

	if url.ToString() != urlStr {
		t.Errorf("Expected toString to return %s, got %s", urlStr, url.ToString())
	}

	if url.ToJSON() != urlStr {
		t.Errorf("Expected toJSON to return %s, got %s", urlStr, url.ToJSON())
	}

	if url.String() != urlStr {
		t.Errorf("Expected String() to return %s, got %s", urlStr, url.String())
	}
}

func TestURL_SetHref(t *testing.T) {
	url, err := NewURL("https://example.com")
	if err != nil {
		t.Fatalf("Failed to create URL: %v", err)
	}

	newHref := "https://newsite.com/path?query=value"
	err = url.SetHref(newHref)
	if err != nil {
		t.Fatalf("Failed to set href: %v", err)
	}

	if url.Href() != newHref {
		t.Errorf("Expected href to be %s, got %s", newHref, url.Href())
	}

	// Test invalid href
	err = url.SetHref("://invalid")
	if err == nil {
		t.Error("Expected error when setting invalid href")
	}
}

func TestURL_EmptyComponents(t *testing.T) {
	url, err := NewURL("https://example.com")
	if err != nil {
		t.Fatalf("Failed to create URL: %v", err)
	}

	// Test empty search
	if url.Search() != "" {
		t.Errorf("Expected empty search, got %s", url.Search())
	}

	// Test empty hash
	if url.Hash() != "" {
		t.Errorf("Expected empty hash, got %s", url.Hash())
	}

	// Test default pathname
	if url.Pathname() != "/" {
		t.Errorf("Expected default pathname to be /, got %s", url.Pathname())
	}
}
