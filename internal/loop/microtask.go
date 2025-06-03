package loop

import (
	"sync"
	"sync/atomic"

	"github.com/dop251/goja"
)

// MicrotaskSource identifies the source of a microtask for debugging and metrics
type MicrotaskSource string

const (
	PromiseTaskSource      MicrotaskSource = "promise"
	QueueMicrotaskSource   MicrotaskSource = "queueMicrotask"
	MutationObserverSource MicrotaskSource = "MutationObserver"
	AnimationFrameSource   MicrotaskSource = "animationFrame"
)

// Microtask represents a unit of work to be executed at the next microtask checkpoint
type Microtask struct {
	ID       int64
	Source   MicrotaskSource
	Callback goja.Callable
	Args     []goja.Value
}

// MicrotaskQueue manages the queue of microtasks with FIFO ordering
type MicrotaskQueue struct {
	microtasks []Microtask
	mu         sync.Mutex
	nextID     atomic.Int64
}

// NewMicrotaskQueue creates a new microtask queue
func NewMicrotaskQueue() *MicrotaskQueue {
	return &MicrotaskQueue{
		microtasks: make([]Microtask, 0),
	}
}

// Queue adds a microtask to the queue
func (mq *MicrotaskQueue) Queue(source MicrotaskSource, callback goja.Callable, args []goja.Value) int64 {
	microtask := Microtask{
		ID:       mq.nextID.Add(1),
		Source:   source,
		Callback: callback,
		Args:     args,
	}

	mq.mu.Lock()
	defer mq.mu.Unlock()

	mq.microtasks = append(mq.microtasks, microtask)
	return microtask.ID
}

// Dequeue removes and returns the next microtask, or nil if the queue is empty
func (mq *MicrotaskQueue) Dequeue() *Microtask {
	mq.mu.Lock()
	defer mq.mu.Unlock()

	if len(mq.microtasks) == 0 {
		return nil
	}

	// FIFO ordering - take from the front
	microtask := mq.microtasks[0]
	mq.microtasks = mq.microtasks[1:]
	return &microtask
}

// Count returns the number of microtasks in the queue
func (mq *MicrotaskQueue) Count() int {
	mq.mu.Lock()
	defer mq.mu.Unlock()
	return len(mq.microtasks)
}

// IsEmpty returns true if the queue is empty
func (mq *MicrotaskQueue) IsEmpty() bool {
	mq.mu.Lock()
	defer mq.mu.Unlock()
	return len(mq.microtasks) == 0
}

// Clear removes all microtasks from the queue
func (mq *MicrotaskQueue) Clear() {
	mq.mu.Lock()
	defer mq.mu.Unlock()
	mq.microtasks = mq.microtasks[:0]
}

// GetMicrotasks returns a copy of all microtasks (for debugging/testing)
func (mq *MicrotaskQueue) GetMicrotasks() []Microtask {
	mq.mu.Lock()
	defer mq.mu.Unlock()

	microtasks := make([]Microtask, len(mq.microtasks))
	copy(microtasks, mq.microtasks)
	return microtasks
}

// GetMicrotasksBySource returns all microtasks from a specific source
func (mq *MicrotaskQueue) GetMicrotasksBySource(source MicrotaskSource) []Microtask {
	mq.mu.Lock()
	defer mq.mu.Unlock()

	var result []Microtask
	for _, microtask := range mq.microtasks {
		if microtask.Source == source {
			result = append(result, microtask)
		}
	}
	return result
}

// DrainBySource removes and returns all microtasks from a specific source
func (mq *MicrotaskQueue) DrainBySource(source MicrotaskSource) []Microtask {
	mq.mu.Lock()
	defer mq.mu.Unlock()

	var result []Microtask
	var remaining []Microtask

	for _, microtask := range mq.microtasks {
		if microtask.Source == source {
			result = append(result, microtask)
		} else {
			remaining = append(remaining, microtask)
		}
	}

	mq.microtasks = remaining
	return result
}
