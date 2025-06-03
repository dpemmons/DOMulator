package url

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// URL represents a URL object following the Web API specification
type URL struct {
	parsedURL *url.URL
}

// NewURL creates a new URL object from a string and optional base URL
// Follows the Web API URL constructor specification
func NewURL(urlString string, base ...string) (*URL, error) {
	var parsedURL *url.URL
	var err error

	if len(base) > 0 && base[0] != "" {
		// Parse with base URL
		baseURL, err := url.Parse(base[0])
		if err != nil {
			return nil, fmt.Errorf("invalid base URL: %w", err)
		}
		parsedURL, err = baseURL.Parse(urlString)
		if err != nil {
			return nil, fmt.Errorf("invalid URL: %w", err)
		}
	} else {
		// Parse absolute URL
		parsedURL, err = url.Parse(urlString)
		if err != nil {
			return nil, fmt.Errorf("invalid URL: %w", err)
		}

		// Web API requires absolute URLs when no base is provided
		if !parsedURL.IsAbs() {
			return nil, fmt.Errorf("invalid URL: relative URL without base")
		}
	}

	return &URL{parsedURL: parsedURL}, nil
}

// Href returns the complete URL as a string
func (u *URL) Href() string {
	return u.parsedURL.String()
}

// SetHref sets the complete URL from a string
func (u *URL) SetHref(href string) error {
	parsedURL, err := url.Parse(href)
	if err != nil {
		return fmt.Errorf("invalid URL: %w", err)
	}
	u.parsedURL = parsedURL
	return nil
}

// Origin returns the origin (protocol + host) of the URL
func (u *URL) Origin() string {
	if u.parsedURL.Scheme == "" {
		return "null"
	}

	origin := u.parsedURL.Scheme + "://"
	if u.parsedURL.Host != "" {
		origin += u.parsedURL.Host
	} else {
		// For file:// URLs and others without host
		return "null"
	}

	return origin
}

// Protocol returns the protocol scheme of the URL (including the trailing colon)
func (u *URL) Protocol() string {
	if u.parsedURL.Scheme == "" {
		return ":"
	}
	return u.parsedURL.Scheme + ":"
}

// SetProtocol sets the protocol scheme of the URL
func (u *URL) SetProtocol(protocol string) {
	// Remove trailing colon if present
	protocol = strings.TrimSuffix(protocol, ":")
	u.parsedURL.Scheme = protocol
}

// Host returns the host (hostname + port) of the URL
func (u *URL) Host() string {
	return u.parsedURL.Host
}

// SetHost sets the host (hostname + port) of the URL
func (u *URL) SetHost(host string) {
	u.parsedURL.Host = host
}

// Hostname returns the hostname of the URL (without port)
func (u *URL) Hostname() string {
	return u.parsedURL.Hostname()
}

// SetHostname sets the hostname of the URL (preserving port)
func (u *URL) SetHostname(hostname string) {
	port := u.parsedURL.Port()
	if port != "" {
		u.parsedURL.Host = hostname + ":" + port
	} else {
		u.parsedURL.Host = hostname
	}
}

// Port returns the port number of the URL
func (u *URL) Port() string {
	return u.parsedURL.Port()
}

// SetPort sets the port number of the URL
func (u *URL) SetPort(port string) {
	hostname := u.parsedURL.Hostname()
	if port == "" {
		u.parsedURL.Host = hostname
	} else {
		// Validate port number
		if portNum, err := strconv.Atoi(port); err != nil || portNum < 0 || portNum > 65535 {
			return // Invalid port, ignore
		}
		u.parsedURL.Host = hostname + ":" + port
	}
}

// Pathname returns the path of the URL
func (u *URL) Pathname() string {
	path := u.parsedURL.Path
	if path == "" {
		return "/"
	}
	return path
}

// SetPathname sets the path of the URL
func (u *URL) SetPathname(pathname string) {
	u.parsedURL.Path = pathname
}

// Search returns the search/query string of the URL (including the leading ?)
func (u *URL) Search() string {
	if u.parsedURL.RawQuery == "" {
		return ""
	}
	return "?" + u.parsedURL.RawQuery
}

// SetSearch sets the search/query string of the URL
func (u *URL) SetSearch(search string) {
	// Remove leading ? if present
	search = strings.TrimPrefix(search, "?")
	u.parsedURL.RawQuery = search
}

// Hash returns the fragment identifier of the URL (including the leading #)
func (u *URL) Hash() string {
	if u.parsedURL.Fragment == "" {
		return ""
	}
	return "#" + u.parsedURL.Fragment
}

// SetHash sets the fragment identifier of the URL
func (u *URL) SetHash(hash string) {
	// Remove leading # if present
	hash = strings.TrimPrefix(hash, "#")
	u.parsedURL.Fragment = hash
}

// SearchParams returns a URLSearchParams object for the URL's query string
func (u *URL) SearchParams() *URLSearchParams {
	return NewURLSearchParams(u.parsedURL.RawQuery)
}

// ToString returns the string representation of the URL (same as Href)
func (u *URL) ToString() string {
	return u.Href()
}

// String implements the Stringer interface
func (u *URL) String() string {
	return u.Href()
}

// ToJSON returns the JSON representation of the URL
func (u *URL) ToJSON() string {
	return u.Href()
}
