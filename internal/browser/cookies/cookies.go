package cookies

import (
	"fmt"
	"net/url"
	"sort"
	"strings"
	"sync"
	"time"
)

// SameSite represents the SameSite attribute of a cookie
type SameSite int

const (
	SameSiteNone SameSite = iota
	SameSiteLax
	SameSiteStrict
)

// String returns the string representation of SameSite
func (s SameSite) String() string {
	switch s {
	case SameSiteNone:
		return "None"
	case SameSiteLax:
		return "Lax"
	case SameSiteStrict:
		return "Strict"
	default:
		return "Lax" // Default
	}
}

// Cookie represents an HTTP cookie
type Cookie struct {
	Name     string
	Value    string
	Path     string
	Domain   string
	Expires  time.Time
	MaxAge   int
	Secure   bool
	HttpOnly bool
	SameSite SameSite
}

// IsExpired checks if the cookie has expired
func (c *Cookie) IsExpired() bool {
	if c.MaxAge < 0 {
		return true
	}
	if c.MaxAge > 0 {
		return false // MaxAge takes precedence over Expires
	}
	if !c.Expires.IsZero() {
		return time.Now().After(c.Expires)
	}
	return false
}

// String returns the string representation of the cookie for Set-Cookie header
func (c *Cookie) String() string {
	var parts []string
	parts = append(parts, fmt.Sprintf("%s=%s", c.Name, c.Value))

	if c.Path != "" {
		parts = append(parts, fmt.Sprintf("Path=%s", c.Path))
	}
	if c.Domain != "" {
		parts = append(parts, fmt.Sprintf("Domain=%s", c.Domain))
	}
	if !c.Expires.IsZero() {
		parts = append(parts, fmt.Sprintf("Expires=%s", c.Expires.UTC().Format(time.RFC1123)))
	}
	if c.MaxAge > 0 {
		parts = append(parts, fmt.Sprintf("Max-Age=%d", c.MaxAge))
	}
	if c.Secure {
		parts = append(parts, "Secure")
	}
	if c.HttpOnly {
		parts = append(parts, "HttpOnly")
	}
	if c.SameSite != SameSiteLax { // Only include if not default
		parts = append(parts, fmt.Sprintf("SameSite=%s", c.SameSite.String()))
	}

	return strings.Join(parts, "; ")
}

// CookieManager manages cookies for a domain
type CookieManager struct {
	cookies map[string]*Cookie
	domain  string
	path    string
	mutex   sync.RWMutex
}

// NewCookieManager creates a new cookie manager
func NewCookieManager(domain, path string) *CookieManager {
	if path == "" {
		path = "/"
	}
	return &CookieManager{
		cookies: make(map[string]*Cookie),
		domain:  domain,
		path:    path,
	}
}

// SetCookie adds or updates a cookie
func (cm *CookieManager) SetCookie(cookie *Cookie) error {
	if cookie.Name == "" {
		return fmt.Errorf("cookie name cannot be empty")
	}

	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	// Set defaults if not specified
	if cookie.Path == "" {
		cookie.Path = cm.path
	}
	if cookie.Domain == "" {
		cookie.Domain = cm.domain
	}

	// If MaxAge is negative, delete the cookie
	if cookie.MaxAge < 0 {
		delete(cm.cookies, cookie.Name)
		return nil
	}

	// Store the cookie
	cm.cookies[cookie.Name] = cookie
	return nil
}

// GetCookie retrieves a cookie by name
func (cm *CookieManager) GetCookie(name string) *Cookie {
	cm.mutex.RLock()
	defer cm.mutex.RUnlock()

	cookie, exists := cm.cookies[name]
	if !exists || cookie.IsExpired() {
		return nil
	}
	return cookie
}

// DeleteCookie removes a cookie by name
func (cm *CookieManager) DeleteCookie(name string) {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()
	delete(cm.cookies, name)
}

// GetAllCookies returns all non-expired cookies
func (cm *CookieManager) GetAllCookies() []*Cookie {
	cm.mutex.RLock()
	defer cm.mutex.RUnlock()

	var result []*Cookie
	for _, cookie := range cm.cookies {
		if !cookie.IsExpired() {
			result = append(result, cookie)
		}
	}

	// Sort by path length (longest first) and name for consistent ordering
	sort.Slice(result, func(i, j int) bool {
		if len(result[i].Path) != len(result[j].Path) {
			return len(result[i].Path) > len(result[j].Path)
		}
		return result[i].Name < result[j].Name
	})

	return result
}

// GetCookiesForDomain returns cookies that match the given domain and path
func (cm *CookieManager) GetCookiesForDomain(domain, path string) []*Cookie {
	cm.mutex.RLock()
	defer cm.mutex.RUnlock()

	var result []*Cookie
	for _, cookie := range cm.cookies {
		if !cookie.IsExpired() && cm.domainMatches(cookie.Domain, domain) && cm.pathMatches(cookie.Path, path) {
			result = append(result, cookie)
		}
	}

	// Sort by path length (longest first) and name for consistent ordering
	sort.Slice(result, func(i, j int) bool {
		if len(result[i].Path) != len(result[j].Path) {
			return len(result[i].Path) > len(result[j].Path)
		}
		return result[i].Name < result[j].Name
	})

	return result
}

// FormatCookieString formats all cookies as a string for the Cookie header
func (cm *CookieManager) FormatCookieString() string {
	cookies := cm.GetAllCookies()
	if len(cookies) == 0 {
		return ""
	}

	var parts []string
	for _, cookie := range cookies {
		parts = append(parts, fmt.Sprintf("%s=%s", cookie.Name, cookie.Value))
	}

	return strings.Join(parts, "; ")
}

// ParseCookieString parses a cookie string from document.cookie and updates the manager
func (cm *CookieManager) ParseCookieString(cookieStr string) error {
	if cookieStr == "" {
		return nil
	}

	// Parse the cookie string
	parts := strings.Split(cookieStr, ";")
	if len(parts) == 0 {
		return nil
	}

	// First part is name=value
	nameValue := strings.TrimSpace(parts[0])
	idx := strings.Index(nameValue, "=")
	if idx == -1 {
		return fmt.Errorf("invalid cookie format: missing '='")
	}

	name := strings.TrimSpace(nameValue[:idx])
	value := strings.TrimSpace(nameValue[idx+1:])

	if name == "" {
		return fmt.Errorf("cookie name cannot be empty")
	}

	// Create cookie with defaults
	cookie := &Cookie{
		Name:     name,
		Value:    value,
		Path:     cm.path,
		Domain:   cm.domain,
		SameSite: SameSiteLax,
	}

	// Parse attributes
	for i := 1; i < len(parts); i++ {
		attr := strings.TrimSpace(parts[i])
		if attr == "" {
			continue
		}

		if strings.EqualFold(attr, "Secure") {
			cookie.Secure = true
		} else if strings.EqualFold(attr, "HttpOnly") {
			cookie.HttpOnly = true
		} else if idx := strings.Index(attr, "="); idx != -1 {
			attrName := strings.TrimSpace(attr[:idx])
			attrValue := strings.TrimSpace(attr[idx+1:])

			switch strings.ToLower(attrName) {
			case "path":
				cookie.Path = attrValue
			case "domain":
				cookie.Domain = attrValue
			case "expires":
				if expires, err := time.Parse(time.RFC1123, attrValue); err == nil {
					cookie.Expires = expires
				}
			case "max-age":
				if maxAge, err := parseMaxAge(attrValue); err == nil {
					cookie.MaxAge = maxAge
				}
			case "samesite":
				switch strings.ToLower(attrValue) {
				case "none":
					cookie.SameSite = SameSiteNone
				case "strict":
					cookie.SameSite = SameSiteStrict
				default:
					cookie.SameSite = SameSiteLax
				}
			}
		}
	}

	return cm.SetCookie(cookie)
}

// ParseSetCookieHeader parses a Set-Cookie header value
func (cm *CookieManager) ParseSetCookieHeader(setCookieValue string) error {
	return cm.ParseCookieString(setCookieValue)
}

// CleanupExpiredCookies removes all expired cookies
func (cm *CookieManager) CleanupExpiredCookies() {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	for name, cookie := range cm.cookies {
		if cookie.IsExpired() {
			delete(cm.cookies, name)
		}
	}
}

// domainMatches checks if the cookie domain matches the request domain
func (cm *CookieManager) domainMatches(cookieDomain, requestDomain string) bool {
	if cookieDomain == "" {
		return true // No domain restriction
	}

	// Check if cookie domain starts with a dot (subdomain matching allowed)
	allowSubdomains := strings.HasPrefix(cookieDomain, ".")
	if allowSubdomains {
		cookieDomain = cookieDomain[1:] // Remove the leading dot
	}

	// Exact match
	if strings.EqualFold(cookieDomain, requestDomain) {
		return true
	}

	// Subdomain match only if cookie domain had leading dot
	if allowSubdomains && strings.HasSuffix(strings.ToLower(requestDomain), "."+strings.ToLower(cookieDomain)) {
		return true
	}

	return false
}

// pathMatches checks if the cookie path matches the request path
func (cm *CookieManager) pathMatches(cookiePath, requestPath string) bool {
	if cookiePath == "" {
		cookiePath = "/"
	}
	if requestPath == "" {
		requestPath = "/"
	}

	// Exact match
	if cookiePath == requestPath {
		return true
	}

	// Path prefix match
	if strings.HasPrefix(requestPath, cookiePath) {
		// Cookie path must end with '/' or be followed by '/'
		if strings.HasSuffix(cookiePath, "/") {
			return true
		}
		if len(requestPath) > len(cookiePath) && requestPath[len(cookiePath)] == '/' {
			return true
		}
	}

	return false
}

// parseMaxAge parses the Max-Age attribute value
func parseMaxAge(value string) (int, error) {
	// Simple integer parsing for Max-Age
	var maxAge int
	n, err := fmt.Sscanf(value, "%d", &maxAge)
	if err != nil || n != 1 {
		return 0, fmt.Errorf("invalid Max-Age value: %s", value)
	}
	return maxAge, nil
}

// GetCookieHeader returns the Cookie header value for HTTP requests
func (cm *CookieManager) GetCookieHeader(requestURL *url.URL) string {
	domain := requestURL.Hostname()
	path := requestURL.Path
	if path == "" {
		path = "/"
	}

	cookies := cm.GetCookiesForDomain(domain, path)
	if len(cookies) == 0 {
		return ""
	}

	var parts []string
	for _, cookie := range cookies {
		parts = append(parts, fmt.Sprintf("%s=%s", cookie.Name, cookie.Value))
	}

	return strings.Join(parts, "; ")
}
