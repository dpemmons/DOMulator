package resources

import (
	"time"

	"github.com/dpemmons/DOMulator/internal/browser/fetch"
	"github.com/dpemmons/DOMulator/internal/dom"
)

// ResourceType identifies different types of resources that can be loaded
type ResourceType string

const (
	ResourceTypeScript     ResourceType = "script"
	ResourceTypeStylesheet ResourceType = "stylesheet"
	ResourceTypeImage      ResourceType = "image"
	ResourceTypeModule     ResourceType = "module"
	ResourceTypeImportMap  ResourceType = "importmap"
)

// LoadState represents the loading state of a resource
type LoadState int

const (
	LoadStateUnloaded LoadState = iota
	LoadStateLoading
	LoadStateLoaded
	LoadStateError
)

// String returns a string representation of the load state
func (s LoadState) String() string {
	switch s {
	case LoadStateUnloaded:
		return "unloaded"
	case LoadStateLoading:
		return "loading"
	case LoadStateLoaded:
		return "loaded"
	case LoadStateError:
		return "error"
	default:
		return "unknown"
	}
}

// ResourceTask represents a resource loading task
type ResourceTask struct {
	ID         string
	Element    *dom.Element
	URL        string
	Type       ResourceType
	Priority   int
	State      LoadState
	Content    []byte
	Error      error
	StartTime  time.Time
	EndTime    time.Time
	Attributes map[string]string
}

// Duration returns the time taken to load the resource
func (rt *ResourceTask) Duration() time.Duration {
	if rt.EndTime.IsZero() {
		return 0
	}
	return rt.EndTime.Sub(rt.StartTime)
}

// IsLoading returns true if the resource is currently loading
func (rt *ResourceTask) IsLoading() bool {
	return rt.State == LoadStateLoading
}

// IsLoaded returns true if the resource has loaded successfully
func (rt *ResourceTask) IsLoaded() bool {
	return rt.State == LoadStateLoaded
}

// HasError returns true if the resource failed to load
func (rt *ResourceTask) HasError() bool {
	return rt.State == LoadStateError
}

// ResourceLoader interface defines the contract for loading specific resource types
type ResourceLoader interface {
	// CanLoad determines if this loader can handle the given element
	CanLoad(element *dom.Element) bool

	// Load starts loading the resource for the given element
	Load(element *dom.Element) (*ResourceTask, error)

	// GetPriority returns the loading priority for this resource type
	GetPriority() int

	// GetType returns the resource type this loader handles
	GetType() ResourceType
}

// ResourceEventEmitter interface for emitting resource loading events
type ResourceEventEmitter interface {
	EmitLoadEvent(element *dom.Element)
	EmitErrorEvent(element *dom.Element, err error)
	EmitProgressEvent(element *dom.Element, loaded, total int64)
}

// URLResolver interface for resolving relative URLs
type URLResolver interface {
	ResolveURL(base, relative string) (string, error)
}

// ResourceCache interface for caching loaded resources
type ResourceCache interface {
	Get(url string) ([]byte, bool)
	Set(url string, data []byte)
	Clear()
	Size() int
}

// FetchAPI interface to abstract the fetch implementation
type FetchAPI interface {
	Fetch(url string, options map[string]interface{}) (*fetch.Response, error)
}
