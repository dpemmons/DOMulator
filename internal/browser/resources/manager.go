package resources

import (
	"fmt"
	"sync"

	"github.com/dpemmons/DOMulator/internal/dom"
)

// ResourceManager orchestrates loading of different resource types
type ResourceManager struct {
	mu          sync.RWMutex
	loaders     map[ResourceType]ResourceLoader
	activeTasks map[string]*ResourceTask // URL -> Task
	cache       ResourceCache
	document    *dom.Document
	fetcher     FetchAPI
}

// NewResourceManager creates a new resource manager
func NewResourceManager(document *dom.Document, fetcher FetchAPI) *ResourceManager {
	return &ResourceManager{
		loaders:     make(map[ResourceType]ResourceLoader),
		activeTasks: make(map[string]*ResourceTask),
		cache:       NewMemoryCache(),
		document:    document,
		fetcher:     fetcher,
	}
}

// RegisterLoader registers a loader for a specific resource type
func (rm *ResourceManager) RegisterLoader(resourceType ResourceType, loader ResourceLoader) {
	rm.mu.Lock()
	defer rm.mu.Unlock()
	rm.loaders[resourceType] = loader
}

// UnregisterLoader removes a loader for a specific resource type
func (rm *ResourceManager) UnregisterLoader(resourceType ResourceType) {
	rm.mu.Lock()
	defer rm.mu.Unlock()
	delete(rm.loaders, resourceType)
}

// LoadResource attempts to load a resource for the given element
func (rm *ResourceManager) LoadResource(element *dom.Element) error {
	rm.mu.RLock()
	defer rm.mu.RUnlock()

	// Find a loader that can handle this element
	for _, loader := range rm.loaders {
		if loader.CanLoad(element) {
			task, err := loader.Load(element)
			if err != nil {
				return fmt.Errorf("failed to load resource with %T: %w", loader, err)
			}

			// Track the active task
			if task != nil {
				rm.activeTasks[task.URL] = task
			}
			return nil
		}
	}

	// No loader found - this might be okay (e.g., unsupported resource type)
	return nil
}

// GetActiveTask returns the active task for a given URL
func (rm *ResourceManager) GetActiveTask(url string) (*ResourceTask, bool) {
	rm.mu.RLock()
	defer rm.mu.RUnlock()
	task, found := rm.activeTasks[url]
	return task, found
}

// GetActiveTasks returns all currently active tasks
func (rm *ResourceManager) GetActiveTasks() []*ResourceTask {
	rm.mu.RLock()
	defer rm.mu.RUnlock()

	tasks := make([]*ResourceTask, 0, len(rm.activeTasks))
	for _, task := range rm.activeTasks {
		tasks = append(tasks, task)
	}
	return tasks
}

// CompleteTask marks a task as completed and removes it from active tasks
func (rm *ResourceManager) CompleteTask(taskID string) {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	// Find and remove the task
	for url, task := range rm.activeTasks {
		if task.ID == taskID {
			delete(rm.activeTasks, url)
			break
		}
	}
}

// GetCache returns the resource cache
func (rm *ResourceManager) GetCache() ResourceCache {
	return rm.cache
}

// SetCache sets a custom resource cache
func (rm *ResourceManager) SetCache(cache ResourceCache) {
	rm.mu.Lock()
	defer rm.mu.Unlock()
	rm.cache = cache
}

// GetRegisteredLoaders returns all registered resource loaders
func (rm *ResourceManager) GetRegisteredLoaders() map[ResourceType]ResourceLoader {
	rm.mu.RLock()
	defer rm.mu.RUnlock()

	// Return a copy to prevent external modification
	loaders := make(map[ResourceType]ResourceLoader)
	for resourceType, loader := range rm.loaders {
		loaders[resourceType] = loader
	}
	return loaders
}

// HasLoader checks if a loader is registered for the given resource type
func (rm *ResourceManager) HasLoader(resourceType ResourceType) bool {
	rm.mu.RLock()
	defer rm.mu.RUnlock()
	_, found := rm.loaders[resourceType]
	return found
}

// GetTasksByType returns all active tasks of a specific type
func (rm *ResourceManager) GetTasksByType(resourceType ResourceType) []*ResourceTask {
	rm.mu.RLock()
	defer rm.mu.RUnlock()

	var tasks []*ResourceTask
	for _, task := range rm.activeTasks {
		if task.Type == resourceType {
			tasks = append(tasks, task)
		}
	}
	return tasks
}

// GetTasksByState returns all active tasks in a specific state
func (rm *ResourceManager) GetTasksByState(state LoadState) []*ResourceTask {
	rm.mu.RLock()
	defer rm.mu.RUnlock()

	var tasks []*ResourceTask
	for _, task := range rm.activeTasks {
		if task.State == state {
			tasks = append(tasks, task)
		}
	}
	return tasks
}

// IsLoading checks if any resources are currently loading
func (rm *ResourceManager) IsLoading() bool {
	rm.mu.RLock()
	defer rm.mu.RUnlock()

	for _, task := range rm.activeTasks {
		if task.IsLoading() {
			return true
		}
	}
	return false
}

// GetLoadingCount returns the number of resources currently loading
func (rm *ResourceManager) GetLoadingCount() int {
	rm.mu.RLock()
	defer rm.mu.RUnlock()

	count := 0
	for _, task := range rm.activeTasks {
		if task.IsLoading() {
			count++
		}
	}
	return count
}

// ClearCompletedTasks removes all completed tasks from the active task list
func (rm *ResourceManager) ClearCompletedTasks() {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	for url, task := range rm.activeTasks {
		if task.IsLoaded() || task.HasError() {
			delete(rm.activeTasks, url)
		}
	}
}
