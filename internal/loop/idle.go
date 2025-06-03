package loop

import (
	"sync"
	"sync/atomic"
	"time"

	"github.com/dop251/goja"
)

// IdleCallback represents a callback scheduled for idle time execution
type IdleCallback struct {
	ID       int64
	Callback goja.Callable
	Timeout  time.Duration
	Created  time.Time
}

// IdleQueue manages idle callbacks with proper scheduling and timeout handling
type IdleQueue struct {
	callbacks []IdleCallback
	mu        sync.RWMutex
	nextID    atomic.Int64
}

// NewIdleQueue creates a new idle queue
func NewIdleQueue() *IdleQueue {
	return &IdleQueue{
		callbacks: make([]IdleCallback, 0),
	}
}

// Schedule adds an idle callback to the queue with optional timeout
func (iq *IdleQueue) Schedule(callback goja.Callable, timeout time.Duration) int64 {
	id := iq.nextID.Add(1)

	iq.mu.Lock()
	defer iq.mu.Unlock()

	iq.callbacks = append(iq.callbacks, IdleCallback{
		ID:       id,
		Callback: callback,
		Timeout:  timeout,
		Created:  time.Now(),
	})

	return id
}

// Cancel removes an idle callback by ID
func (iq *IdleQueue) Cancel(id int64) bool {
	iq.mu.Lock()
	defer iq.mu.Unlock()

	for i, callback := range iq.callbacks {
		if callback.ID == id {
			// Remove the callback
			iq.callbacks = append(iq.callbacks[:i], iq.callbacks[i+1:]...)
			return true
		}
	}
	return false
}

// GetNextCallback returns the next callback ready for execution
func (iq *IdleQueue) GetNextCallback() *IdleCallback {
	iq.mu.Lock()
	defer iq.mu.Unlock()

	if len(iq.callbacks) == 0 {
		return nil
	}

	// For now, just return the first callback (FIFO)
	// In a more sophisticated implementation, we could prioritize by timeout
	callback := iq.callbacks[0]
	iq.callbacks = iq.callbacks[1:]
	return &callback
}

// GetTimeoutCallbacks returns callbacks that have exceeded their timeout
func (iq *IdleQueue) GetTimeoutCallbacks(now time.Time) []IdleCallback {
	iq.mu.RLock()
	defer iq.mu.RUnlock()

	var timeoutCallbacks []IdleCallback
	for _, callback := range iq.callbacks {
		if callback.Timeout > 0 && now.Sub(callback.Created) >= callback.Timeout {
			timeoutCallbacks = append(timeoutCallbacks, callback)
		}
	}
	return timeoutCallbacks
}

// RemoveTimeoutCallbacks removes callbacks that have exceeded their timeout
func (iq *IdleQueue) RemoveTimeoutCallbacks(now time.Time) []IdleCallback {
	iq.mu.Lock()
	defer iq.mu.Unlock()

	var timeoutCallbacks []IdleCallback
	var remainingCallbacks []IdleCallback

	for _, callback := range iq.callbacks {
		if callback.Timeout > 0 && now.Sub(callback.Created) >= callback.Timeout {
			timeoutCallbacks = append(timeoutCallbacks, callback)
		} else {
			remainingCallbacks = append(remainingCallbacks, callback)
		}
	}

	iq.callbacks = remainingCallbacks
	return timeoutCallbacks
}

// Count returns the number of callbacks in the queue
func (iq *IdleQueue) Count() int {
	iq.mu.RLock()
	defer iq.mu.RUnlock()
	return len(iq.callbacks)
}

// IsEmpty returns true if the queue is empty
func (iq *IdleQueue) IsEmpty() bool {
	iq.mu.RLock()
	defer iq.mu.RUnlock()
	return len(iq.callbacks) == 0
}

// Clear removes all callbacks from the queue
func (iq *IdleQueue) Clear() {
	iq.mu.Lock()
	defer iq.mu.Unlock()
	iq.callbacks = iq.callbacks[:0]
}

// GetCallbacks returns a copy of all callbacks (for debugging/testing)
func (iq *IdleQueue) GetCallbacks() []IdleCallback {
	iq.mu.RLock()
	defer iq.mu.RUnlock()

	callbacks := make([]IdleCallback, len(iq.callbacks))
	copy(callbacks, iq.callbacks)
	return callbacks
}

// HasCallbacks returns true if there are callbacks waiting
func (iq *IdleQueue) HasCallbacks() bool {
	iq.mu.RLock()
	defer iq.mu.RUnlock()
	return len(iq.callbacks) > 0
}
