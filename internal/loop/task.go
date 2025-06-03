package loop

import (
	"sync"
	"sync/atomic"
	"time"

	"github.com/dop251/goja"
)

// TaskType defines the type of task for prioritization and scheduling
type TaskType int

const (
	TimerTask TaskType = iota
	DOMEventTask
	NetworkTask
	UserInteractionTask
	HistoryTraversalTask
)

// TaskSource identifies the source of a task for debugging and metrics
type TaskSource string

const (
	TimerTaskSource           TaskSource = "timer"
	DOMManipulationTaskSource TaskSource = "DOM manipulation"
	UserInteractionTaskSource TaskSource = "user interaction"
	NetworkingTaskSource      TaskSource = "networking"
	HistoryTraversalSource    TaskSource = "history traversal"
)

// Task represents a unit of work to be executed by the event loop
type Task struct {
	ID          int64
	Type        TaskType
	Source      TaskSource
	Callback    goja.Callable
	Args        []goja.Value
	Delay       time.Duration
	ScheduledAt time.Time
	ExecuteAt   time.Time
}

// TaskQueue manages the queue of tasks with proper scheduling and prioritization
type TaskQueue struct {
	tasks  []*Task
	mu     sync.Mutex
	nextID atomic.Int64
}

// NewTaskQueue creates a new task queue
func NewTaskQueue() *TaskQueue {
	return &TaskQueue{
		tasks: make([]*Task, 0),
	}
}

// ScheduleTask adds a task to the queue with the specified delay
func (tq *TaskQueue) ScheduleTask(taskType TaskType, source TaskSource, callback goja.Callable, args []goja.Value, delay time.Duration, currentTime time.Time) int64 {
	task := &Task{
		ID:          tq.nextID.Add(1),
		Type:        taskType,
		Source:      source,
		Callback:    callback,
		Args:        args,
		Delay:       delay,
		ScheduledAt: currentTime,
		ExecuteAt:   currentTime.Add(delay),
	}

	tq.mu.Lock()
	defer tq.mu.Unlock()

	// Insert task in sorted order by execution time
	tq.insertTaskSorted(task)

	return task.ID
}

// insertTaskSorted inserts a task in the correct position based on execution time
func (tq *TaskQueue) insertTaskSorted(task *Task) {
	// Find the correct position to maintain sorted order
	insertPos := len(tq.tasks)
	for i, existing := range tq.tasks {
		if task.ExecuteAt.Before(existing.ExecuteAt) {
			insertPos = i
			break
		}
	}

	// Insert at the correct position
	tq.tasks = append(tq.tasks, nil)
	copy(tq.tasks[insertPos+1:], tq.tasks[insertPos:])
	tq.tasks[insertPos] = task
}

// SelectTask selects the next task ready for execution
func (tq *TaskQueue) SelectTask(currentTime time.Time) *Task {
	tq.mu.Lock()
	defer tq.mu.Unlock()

	if len(tq.tasks) == 0 {
		return nil
	}

	// Check if the first task is ready to execute
	if tq.tasks[0].ExecuteAt.After(currentTime) {
		return nil // No tasks ready yet
	}

	// Remove and return the first task
	task := tq.tasks[0]
	tq.tasks = tq.tasks[1:]
	return task
}

// SelectTaskByTime selects the next task that should execute by the specified time
func (tq *TaskQueue) SelectTaskByTime(targetTime time.Time) *Task {
	tq.mu.Lock()
	defer tq.mu.Unlock()

	if len(tq.tasks) == 0 {
		return nil
	}

	// Check if the first task should execute by the target time
	if tq.tasks[0].ExecuteAt.After(targetTime) {
		return nil // No tasks ready yet
	}

	// Remove and return the first task
	task := tq.tasks[0]
	tq.tasks = tq.tasks[1:]
	return task
}

// CancelTask removes a task from the queue by ID
func (tq *TaskQueue) CancelTask(id int64) bool {
	tq.mu.Lock()
	defer tq.mu.Unlock()

	for i, task := range tq.tasks {
		if task.ID == id {
			// Remove the task
			tq.tasks = append(tq.tasks[:i], tq.tasks[i+1:]...)
			return true
		}
	}
	return false
}

// Count returns the number of tasks in the queue
func (tq *TaskQueue) Count() int {
	tq.mu.Lock()
	defer tq.mu.Unlock()
	return len(tq.tasks)
}

// IsEmpty returns true if the queue is empty
func (tq *TaskQueue) IsEmpty() bool {
	tq.mu.Lock()
	defer tq.mu.Unlock()
	return len(tq.tasks) == 0
}

// GetNextExecutionTime returns the execution time of the next task, or zero time if empty
func (tq *TaskQueue) GetNextExecutionTime() time.Time {
	tq.mu.Lock()
	defer tq.mu.Unlock()

	if len(tq.tasks) == 0 {
		return time.Time{}
	}

	return tq.tasks[0].ExecuteAt
}

// Clear removes all tasks from the queue
func (tq *TaskQueue) Clear() {
	tq.mu.Lock()
	defer tq.mu.Unlock()
	tq.tasks = tq.tasks[:0]
}

// GetTasks returns a copy of all tasks (for debugging/testing)
func (tq *TaskQueue) GetTasks() []*Task {
	tq.mu.Lock()
	defer tq.mu.Unlock()

	tasks := make([]*Task, len(tq.tasks))
	copy(tasks, tq.tasks)
	return tasks
}

// GetTasksByType returns all tasks of a specific type
func (tq *TaskQueue) GetTasksByType(taskType TaskType) []*Task {
	tq.mu.Lock()
	defer tq.mu.Unlock()

	var result []*Task
	for _, task := range tq.tasks {
		if task.Type == taskType {
			result = append(result, task)
		}
	}
	return result
}

// GetTasksBySource returns all tasks from a specific source
func (tq *TaskQueue) GetTasksBySource(source TaskSource) []*Task {
	tq.mu.Lock()
	defer tq.mu.Unlock()

	var result []*Task
	for _, task := range tq.tasks {
		if task.Source == source {
			result = append(result, task)
		}
	}
	return result
}
