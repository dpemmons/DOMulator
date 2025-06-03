package loop

import (
	"sync"
	"sync/atomic"
	"time"

	"github.com/dop251/goja"
)

// AnimationFrameCallback represents a callback scheduled for the next animation frame
type AnimationFrameCallback struct {
	ID        int64
	Callback  goja.Callable
	Timestamp time.Time
}

// AnimationQueue manages animation frame callbacks with proper scheduling
type AnimationQueue struct {
	callbacks []AnimationFrameCallback
	mu        sync.RWMutex
	nextID    atomic.Int64
}

// NewAnimationQueue creates a new animation queue
func NewAnimationQueue() *AnimationQueue {
	return &AnimationQueue{
		callbacks: make([]AnimationFrameCallback, 0),
	}
}

// Schedule adds an animation frame callback to the queue
func (aq *AnimationQueue) Schedule(callback goja.Callable) int64 {
	id := aq.nextID.Add(1)

	aq.mu.Lock()
	defer aq.mu.Unlock()

	aq.callbacks = append(aq.callbacks, AnimationFrameCallback{
		ID:        id,
		Callback:  callback,
		Timestamp: time.Now(),
	})

	return id
}

// Cancel removes an animation frame callback by ID
func (aq *AnimationQueue) Cancel(id int64) bool {
	aq.mu.Lock()
	defer aq.mu.Unlock()

	for i, callback := range aq.callbacks {
		if callback.ID == id {
			// Remove the callback
			aq.callbacks = append(aq.callbacks[:i], aq.callbacks[i+1:]...)
			return true
		}
	}
	return false
}

// GetCallbacks returns all scheduled callbacks and optionally clears the queue
func (aq *AnimationQueue) GetCallbacks() []AnimationFrameCallback {
	aq.mu.RLock()
	defer aq.mu.RUnlock()

	callbacks := make([]AnimationFrameCallback, len(aq.callbacks))
	copy(callbacks, aq.callbacks)
	return callbacks
}

// Clear removes all callbacks from the queue
func (aq *AnimationQueue) Clear() {
	aq.mu.Lock()
	defer aq.mu.Unlock()
	aq.callbacks = aq.callbacks[:0]
}

// HasCallbacks returns true if there are callbacks waiting
func (aq *AnimationQueue) HasCallbacks() bool {
	aq.mu.RLock()
	defer aq.mu.RUnlock()
	return len(aq.callbacks) > 0
}

// Count returns the number of callbacks in the queue
func (aq *AnimationQueue) Count() int {
	aq.mu.RLock()
	defer aq.mu.RUnlock()
	return len(aq.callbacks)
}

// GetOldestCallback returns the oldest callback without removing it
func (aq *AnimationQueue) GetOldestCallback() *AnimationFrameCallback {
	aq.mu.RLock()
	defer aq.mu.RUnlock()

	if len(aq.callbacks) == 0 {
		return nil
	}

	// Find the oldest callback by timestamp
	oldest := &aq.callbacks[0]
	for i := 1; i < len(aq.callbacks); i++ {
		if aq.callbacks[i].Timestamp.Before(oldest.Timestamp) {
			oldest = &aq.callbacks[i]
		}
	}

	return oldest
}

// DrainCallbacks returns all callbacks and clears the queue
func (aq *AnimationQueue) DrainCallbacks() []AnimationFrameCallback {
	aq.mu.Lock()
	defer aq.mu.Unlock()

	callbacks := make([]AnimationFrameCallback, len(aq.callbacks))
	copy(callbacks, aq.callbacks)
	aq.callbacks = aq.callbacks[:0]

	return callbacks
}
