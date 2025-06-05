package resources

import (
	"errors"
	"testing"

	"github.com/dpemmons/DOMulator/internal/browser/fetch"
	"github.com/dpemmons/DOMulator/internal/dom"
)

// MockJSRuntime implements JSRuntime for testing
type MockJSRuntime struct {
	executedScripts []TestScriptExecution
	shouldError     bool
}

type TestScriptExecution struct {
	Name string
	Code string
}

func NewMockJSRuntime() *MockJSRuntime {
	return &MockJSRuntime{
		executedScripts: make([]TestScriptExecution, 0),
		shouldError:     false,
	}
}

func (mjr *MockJSRuntime) RunScript(name, code string) (interface{}, error) {
	mjr.executedScripts = append(mjr.executedScripts, TestScriptExecution{
		Name: name,
		Code: code,
	})

	if mjr.shouldError {
		return nil, errors.New("Mock script execution error")
	}

	return nil, nil
}

func (mjr *MockJSRuntime) VM() interface{} {
	return "mock-vm"
}

func (mjr *MockJSRuntime) SetShouldError(shouldError bool) {
	mjr.shouldError = shouldError
}

func (mjr *MockJSRuntime) GetExecutedScripts() []TestScriptExecution {
	return mjr.executedScripts
}

func (mjr *MockJSRuntime) Reset() {
	mjr.executedScripts = mjr.executedScripts[:0]
	mjr.shouldError = false
}

func TestResourceManager_Creation(t *testing.T) {
	// Create test dependencies
	document := dom.NewDocument()
	fetchAPI := fetch.New(nil) // No network mocks for now
	runtime := NewMockJSRuntime()

	// Create resource manager
	manager := New(document, fetchAPI, runtime)

	// Verify manager was created with script loader
	if !manager.HasLoader(ResourceTypeScript) {
		t.Error("Expected script loader to be registered")
	}

	loaders := manager.GetRegisteredLoaders()
	if len(loaders) != 1 {
		t.Errorf("Expected 1 loader, got %d", len(loaders))
	}

	if _, exists := loaders[ResourceTypeScript]; !exists {
		t.Error("Expected script loader to be registered")
	}
}

func TestScriptLoader_CanLoad(t *testing.T) {
	// Create test dependencies
	document := dom.NewDocument()
	fetchAPI := fetch.New(nil)
	runtime := NewMockJSRuntime()

	manager := New(document, fetchAPI, runtime)
	loaders := manager.GetRegisteredLoaders()
	scriptLoader := loaders[ResourceTypeScript]

	tests := []struct {
		name       string
		tagName    string
		src        string
		shouldLoad bool
	}{
		{"External script", "script", "app.js", true},
		{"External script uppercase", "SCRIPT", "app.js", true},
		{"Inline script", "script", "", false},
		{"Non-script element", "div", "app.js", false},
		{"Link element", "link", "style.css", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			element := document.CreateElement(tt.tagName)
			if tt.src != "" {
				element.SetAttribute("src", tt.src)
			}

			canLoad := scriptLoader.CanLoad(element)
			if canLoad != tt.shouldLoad {
				t.Errorf("Expected CanLoad() = %v, got %v", tt.shouldLoad, canLoad)
			}
		})
	}
}

func TestResourceCache(t *testing.T) {
	cache := NewMemoryCache()

	// Test Set and Get
	url := "test.js"
	content := []byte("console.log('test');")

	cache.Set(url, content)

	retrieved, found := cache.Get(url)
	if !found {
		t.Error("Expected to find cached content")
	}

	if string(retrieved) != string(content) {
		t.Errorf("Expected content '%s', got '%s'", string(content), string(retrieved))
	}

	// Test Size
	if cache.Size() != 1 {
		t.Errorf("Expected cache size 1, got %d", cache.Size())
	}

	// Test Has
	if !cache.Has(url) {
		t.Error("Expected cache to have URL")
	}

	// Test Clear
	cache.Clear()
	if cache.Size() != 0 {
		t.Errorf("Expected cache size 0 after clear, got %d", cache.Size())
	}

	if cache.Has(url) {
		t.Error("Expected cache not to have URL after clear")
	}
}

func TestResourceManager_MultipleLoaders(t *testing.T) {
	// Create test dependencies
	document := dom.NewDocument()
	fetchAPI := fetch.New(nil)

	manager := NewResourceManager(document, NewFetchAdapter(fetchAPI))

	// Mock loader for testing
	mockLoader := &MockResourceLoader{
		canLoadFunc: func(element *dom.Element) bool {
			return element.TagName() == "mock"
		},
		resourceType: ResourceType("mock"),
	}

	// Register multiple loaders
	manager.RegisterLoader(ResourceTypeScript, mockLoader)
	manager.RegisterLoader(ResourceType("mock"), mockLoader)

	// Verify both are registered
	if !manager.HasLoader(ResourceTypeScript) {
		t.Error("Expected script loader to be registered")
	}

	if !manager.HasLoader(ResourceType("mock")) {
		t.Error("Expected mock loader to be registered")
	}

	loaders := manager.GetRegisteredLoaders()
	if len(loaders) != 2 {
		t.Errorf("Expected 2 loaders, got %d", len(loaders))
	}

	// Unregister one
	manager.UnregisterLoader(ResourceType("mock"))
	if manager.HasLoader(ResourceType("mock")) {
		t.Error("Expected mock loader to be unregistered")
	}
}

// MockResourceLoader for testing
type MockResourceLoader struct {
	canLoadFunc  func(*dom.Element) bool
	resourceType ResourceType
}

func (mrl *MockResourceLoader) CanLoad(element *dom.Element) bool {
	if mrl.canLoadFunc != nil {
		return mrl.canLoadFunc(element)
	}
	return false
}

func (mrl *MockResourceLoader) Load(element *dom.Element) (*ResourceTask, error) {
	return &ResourceTask{
		ID:      "mock-task",
		Element: element,
		Type:    mrl.resourceType,
		State:   LoadStateLoaded,
	}, nil
}

func (mrl *MockResourceLoader) GetPriority() int {
	return 100
}

func (mrl *MockResourceLoader) GetType() ResourceType {
	return mrl.resourceType
}
