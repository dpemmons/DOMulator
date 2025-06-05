package resources

import (
	"fmt"
	"net/url"
	"strings"
	"sync/atomic"

	"github.com/dpemmons/DOMulator/internal/dom"
)

// BaseResourceLoader provides common functionality for all resource loaders
type BaseResourceLoader struct {
	document  *dom.Document
	fetcher   FetchAPI
	cache     ResourceCache
	idCounter int64
}

// NewBaseResourceLoader creates a new base resource loader
func NewBaseResourceLoader(document *dom.Document, fetcher FetchAPI, cache ResourceCache) *BaseResourceLoader {
	return &BaseResourceLoader{
		document: document,
		fetcher:  fetcher,
		cache:    cache,
	}
}

// GenerateID generates a unique ID for a resource task
func (bl *BaseResourceLoader) GenerateID() string {
	id := atomic.AddInt64(&bl.idCounter, 1)
	return fmt.Sprintf("resource_%d", id)
}

// ResolveURL resolves a relative URL against the document's base URL
func (bl *BaseResourceLoader) ResolveURL(relative string) (string, error) {
	if relative == "" {
		return "", fmt.Errorf("empty URL")
	}

	// If already absolute, return as-is
	if strings.HasPrefix(relative, "http://") || strings.HasPrefix(relative, "https://") {
		return relative, nil
	}

	// Get document base URL (defaults to "about:blank")
	baseURL := bl.document.URL()
	if baseURL == "" || baseURL == "about:blank" {
		// For testing/relative URLs, assume current working directory
		baseURL = "file:///"
	}

	// Parse base URL
	base, err := url.Parse(baseURL)
	if err != nil {
		return "", fmt.Errorf("invalid base URL %q: %w", baseURL, err)
	}

	// Parse relative URL
	rel, err := url.Parse(relative)
	if err != nil {
		return "", fmt.Errorf("invalid relative URL %q: %w", relative, err)
	}

	// Resolve relative against base
	resolved := base.ResolveReference(rel)
	return resolved.String(), nil
}

// FetchContent fetches content from a URL, using cache if available
func (bl *BaseResourceLoader) FetchContent(resourceURL string) ([]byte, error) {
	// Check cache first
	if bl.cache != nil {
		if content, found := bl.cache.Get(resourceURL); found {
			return content, nil
		}
	}

	// Fetch from network
	response, err := bl.fetcher.Fetch(resourceURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch %q: %w", resourceURL, err)
	}

	// Check for HTTP errors
	if response.Status >= 400 {
		return nil, fmt.Errorf("HTTP %d fetching %q", response.Status, resourceURL)
	}

	content := []byte(response.Body)

	// Cache the content
	if bl.cache != nil {
		bl.cache.Set(resourceURL, content)
	}

	return content, nil
}

// CreateTask creates a new resource task for the given element
func (bl *BaseResourceLoader) CreateTask(element *dom.Element, resourceType ResourceType) (*ResourceTask, error) {
	// Get the source URL from element
	src := bl.getSourceURL(element)
	if src == "" {
		return nil, fmt.Errorf("no source URL found for element")
	}

	// Resolve URL
	resolvedURL, err := bl.ResolveURL(src)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve URL %q: %w", src, err)
	}

	// Extract attributes
	attributes := make(map[string]string)
	if element.HasAttribute("async") {
		attributes["async"] = element.GetAttribute("async")
	}
	if element.HasAttribute("defer") {
		attributes["defer"] = element.GetAttribute("defer")
	}
	if element.HasAttribute("type") {
		attributes["type"] = element.GetAttribute("type")
	}

	return &ResourceTask{
		ID:         bl.GenerateID(),
		Element:    element,
		URL:        resolvedURL,
		Type:       resourceType,
		State:      LoadStateUnloaded,
		Attributes: attributes,
	}, nil
}

// getSourceURL extracts the source URL from an element based on its type
func (bl *BaseResourceLoader) getSourceURL(element *dom.Element) string {
	switch strings.ToLower(element.TagName()) {
	case "script":
		return element.GetAttribute("src")
	case "link":
		// For stylesheets: <link rel="stylesheet" href="...">
		if element.GetAttribute("rel") == "stylesheet" {
			return element.GetAttribute("href")
		}
	case "img":
		return element.GetAttribute("src")
	default:
		// Check for src attribute as fallback
		return element.GetAttribute("src")
	}
	return ""
}

// EmitLoadEvent emits a load event on the element
func (bl *BaseResourceLoader) EmitLoadEvent(element *dom.Element) {
	event := dom.NewEvent("load", true, false) // bubbles: true, cancelable: false
	element.DispatchEvent(event)
}

// EmitErrorEvent emits an error event on the element
func (bl *BaseResourceLoader) EmitErrorEvent(element *dom.Element, err error) {
	event := dom.NewEvent("error", true, false) // bubbles: true, cancelable: false
	// TODO: Add error details to event when ErrorEvent is implemented
	element.DispatchEvent(event)
}

// EmitProgressEvent emits a progress event on the element
func (bl *BaseResourceLoader) EmitProgressEvent(element *dom.Element, loaded, total int64) {
	event := dom.NewEvent("progress", true, false) // bubbles: true, cancelable: false
	// TODO: Add progress details to event when ProgressEvent is implemented
	element.DispatchEvent(event)
}
