package resources

import (
	"github.com/dpemmons/DOMulator/internal/browser/fetch"
	"github.com/dpemmons/DOMulator/internal/dom"
)

// New creates a fully configured ResourceManager with default loaders
func New(document *dom.Document, fetchAPI *fetch.FetchAPI, runtime JSRuntime) *ResourceManager {
	// Create fetch adapter
	fetchAdapter := NewFetchAdapter(fetchAPI)

	// Create resource manager
	manager := NewResourceManager(document, fetchAdapter)

	// Create base loader with cache
	baseLoader := NewBaseResourceLoader(document, fetchAdapter, manager.GetCache())

	// Create and register script loader
	scriptLoader := NewScriptLoader(baseLoader, runtime)
	manager.RegisterLoader(ResourceTypeScript, scriptLoader)

	return manager
}

// NewWithCustomCache creates a ResourceManager with a custom cache
func NewWithCustomCache(document *dom.Document, fetchAPI *fetch.FetchAPI, runtime JSRuntime, cache ResourceCache) *ResourceManager {
	// Create fetch adapter
	fetchAdapter := NewFetchAdapter(fetchAPI)

	// Create resource manager with custom cache
	manager := NewResourceManager(document, fetchAdapter)
	manager.SetCache(cache)

	// Create base loader with custom cache
	baseLoader := NewBaseResourceLoader(document, fetchAdapter, cache)

	// Create and register script loader
	scriptLoader := NewScriptLoader(baseLoader, runtime)
	manager.RegisterLoader(ResourceTypeScript, scriptLoader)

	return manager
}

// CreateDefaultScriptLoader creates a script loader with default configuration
func CreateDefaultScriptLoader(document *dom.Document, fetchAPI *fetch.FetchAPI, runtime JSRuntime) *ScriptLoader {
	fetchAdapter := NewFetchAdapter(fetchAPI)
	cache := NewMemoryCache()
	baseLoader := NewBaseResourceLoader(document, fetchAdapter, cache)
	return NewScriptLoader(baseLoader, runtime)
}

// JSRuntimeAdapter adapts different JavaScript runtime types
type JSRuntimeAdapter struct {
	runScript func(name, code string) (interface{}, error)
	vm        interface{}
}

// NewJSRuntimeAdapter creates an adapter for any JS runtime
func NewJSRuntimeAdapter(runScript func(name, code string) (interface{}, error), vm interface{}) *JSRuntimeAdapter {
	return &JSRuntimeAdapter{
		runScript: runScript,
		vm:        vm,
	}
}

// RunScript executes JavaScript code
func (jra *JSRuntimeAdapter) RunScript(name, code string) (interface{}, error) {
	return jra.runScript(name, code)
}

// VM returns the underlying JavaScript VM
func (jra *JSRuntimeAdapter) VM() interface{} {
	return jra.vm
}
