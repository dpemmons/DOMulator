package resources

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/dpemmons/DOMulator/internal/dom"
)

// JSRuntime interface to abstract the JavaScript runtime
type JSRuntime interface {
	RunScript(name, code string) (interface{}, error)
	VM() interface{} // Returns the underlying JS VM
}

// ScriptLoader handles loading and execution of JavaScript scripts
type ScriptLoader struct {
	*BaseResourceLoader
	runtime     JSRuntime
	mu          sync.Mutex
	scriptQueue []*ScriptExecution
	executing   bool
}

// ScriptExecution represents a script waiting to be executed
type ScriptExecution struct {
	Task        *ResourceTask
	Content     []byte
	ExecuteTime time.Time
	Priority    int // Lower number = higher priority
}

// NewScriptLoader creates a new script loader
func NewScriptLoader(baseLoader *BaseResourceLoader, runtime JSRuntime) *ScriptLoader {
	return &ScriptLoader{
		BaseResourceLoader: baseLoader,
		runtime:            runtime,
		scriptQueue:        make([]*ScriptExecution, 0),
	}
}

// CanLoad determines if this loader can handle the given element
func (sl *ScriptLoader) CanLoad(element *dom.Element) bool {
	return strings.ToLower(element.TagName()) == "script" && element.GetAttribute("src") != ""
}

// Load starts loading a script resource
func (sl *ScriptLoader) Load(element *dom.Element) (*ResourceTask, error) {
	// Create the base task
	task, err := sl.CreateTask(element, ResourceTypeScript)
	if err != nil {
		return nil, err
	}

	// Set loading state BEFORE starting goroutine to avoid race condition
	task.State = LoadStateLoading
	task.StartTime = time.Now()

	// Start loading asynchronously
	go sl.loadAndQueue(task)

	return task, nil
}

// GetPriority returns the loading priority for scripts
func (sl *ScriptLoader) GetPriority() int {
	return 100 // High priority for scripts
}

// GetType returns the resource type this loader handles
func (sl *ScriptLoader) GetType() ResourceType {
	return ResourceTypeScript
}

// loadAndQueue loads the script content and queues it for execution
func (sl *ScriptLoader) loadAndQueue(task *ResourceTask) {
	// Fetch the script content
	content, err := sl.FetchContent(task.URL)
	if err != nil {
		task.State = LoadStateError
		task.Error = err
		task.EndTime = time.Now()
		sl.EmitErrorEvent(task.Element, err)
		return
	}

	task.Content = content
	task.State = LoadStateLoaded
	task.EndTime = time.Now()

	// Determine execution priority and timing
	execution := &ScriptExecution{
		Task:     task,
		Content:  content,
		Priority: sl.calculatePriority(task),
	}

	// Handle async/defer attributes
	isAsync := task.Attributes["async"] != ""
	isDefer := task.Attributes["defer"] != ""

	if isAsync {
		// Async scripts execute immediately when loaded
		sl.executeScript(execution)
	} else if isDefer {
		// Deferred scripts wait until parsing is complete
		execution.ExecuteTime = time.Now().Add(50 * time.Millisecond) // Longer delay to ensure DOM is ready
		sl.queueForExecution(execution)
	} else {
		// Regular scripts execute in document order
		sl.queueForExecution(execution)
	}
}

// calculatePriority determines execution priority based on script attributes
func (sl *ScriptLoader) calculatePriority(task *ResourceTask) int {
	// Lower number = higher priority
	if task.Attributes["async"] != "" {
		return 1 // Highest priority - execute immediately
	}
	if task.Attributes["defer"] != "" {
		return 3 // Lower priority - execute after parsing
	}
	return 2 // Normal priority - document order
}

// queueForExecution adds a script to the execution queue
func (sl *ScriptLoader) queueForExecution(execution *ScriptExecution) {
	sl.mu.Lock()
	defer sl.mu.Unlock()

	// Insert in priority order
	inserted := false
	for i, existing := range sl.scriptQueue {
		if execution.Priority < existing.Priority {
			// Insert before this element
			sl.scriptQueue = append(sl.scriptQueue[:i], append([]*ScriptExecution{execution}, sl.scriptQueue[i:]...)...)
			inserted = true
			break
		}
	}

	if !inserted {
		sl.scriptQueue = append(sl.scriptQueue, execution)
	}

	// Process the queue if we're not already executing
	if !sl.executing {
		go sl.processQueue()
	}
}

// processQueue processes scripts in the execution queue
func (sl *ScriptLoader) processQueue() {
	sl.mu.Lock()
	if sl.executing {
		sl.mu.Unlock()
		return
	}
	sl.executing = true
	sl.mu.Unlock()

	defer func() {
		sl.mu.Lock()
		sl.executing = false
		sl.mu.Unlock()
	}()

	for {
		sl.mu.Lock()
		if len(sl.scriptQueue) == 0 {
			sl.mu.Unlock()
			break
		}

		// Get the next script to execute
		execution := sl.scriptQueue[0]
		sl.scriptQueue = sl.scriptQueue[1:]
		sl.mu.Unlock()

		// Wait if this is a deferred script and we need to delay
		if !execution.ExecuteTime.IsZero() && time.Now().Before(execution.ExecuteTime) {
			time.Sleep(execution.ExecuteTime.Sub(time.Now()))
		}

		sl.executeScript(execution)
	}
}

// executeScript executes a script in the JavaScript runtime
func (sl *ScriptLoader) executeScript(execution *ScriptExecution) {
	task := execution.Task

	// Create a script name for debugging
	scriptName := task.URL
	if scriptName == "" {
		scriptName = fmt.Sprintf("inline-script-%s", task.ID)
	}

	// Execute the script
	_, err := sl.runtime.RunScript(scriptName, string(execution.Content))
	if err != nil {
		task.State = LoadStateError
		task.Error = fmt.Errorf("script execution error: %w", err)
		sl.EmitErrorEvent(task.Element, err)
		return
	}

	// Emit load event
	sl.EmitLoadEvent(task.Element)
}

// ExecuteInlineScript executes an inline script immediately
func (sl *ScriptLoader) ExecuteInlineScript(element *dom.Element, content string) error {
	// Create a task for tracking
	task := &ResourceTask{
		ID:      sl.GenerateID(),
		Element: element,
		URL:     "", // Inline scripts don't have URLs
		Type:    ResourceTypeScript,
		State:   LoadStateLoaded,
		Content: []byte(content),
	}

	execution := &ScriptExecution{
		Task:     task,
		Content:  []byte(content),
		Priority: 1, // High priority for inline scripts
	}

	sl.executeScript(execution)
	return nil
}

// GetQueueLength returns the number of scripts waiting to execute
func (sl *ScriptLoader) GetQueueLength() int {
	sl.mu.Lock()
	defer sl.mu.Unlock()
	return len(sl.scriptQueue)
}

// IsExecuting returns true if scripts are currently being executed
func (sl *ScriptLoader) IsExecuting() bool {
	sl.mu.Lock()
	defer sl.mu.Unlock()
	return sl.executing
}

// WaitForCompletion waits for all queued scripts to finish executing
func (sl *ScriptLoader) WaitForCompletion(timeout time.Duration) error {
	deadline := time.Now().Add(timeout)

	for time.Now().Before(deadline) {
		sl.mu.Lock()
		queueEmpty := len(sl.scriptQueue) == 0
		notExecuting := !sl.executing
		sl.mu.Unlock()

		if queueEmpty && notExecuting {
			return nil
		}

		time.Sleep(10 * time.Millisecond)
	}

	return fmt.Errorf("timeout waiting for script execution to complete")
}

// ClearQueue removes all pending scripts from the execution queue
func (sl *ScriptLoader) ClearQueue() {
	sl.mu.Lock()
	defer sl.mu.Unlock()
	sl.scriptQueue = sl.scriptQueue[:0]
}
