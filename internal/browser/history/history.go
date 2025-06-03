package history

import (
	"encoding/json"
	"fmt"
	"net/url"
	"sync"
)

// History represents the browser history API
type History struct {
	stack   []HistoryEntry
	current int
	maxLen  int
	mutex   sync.RWMutex

	// Event callback for popstate events
	onPopState func(state interface{})
}

// HistoryEntry represents a single history entry
type HistoryEntry struct {
	URL   string      `json:"url"`
	State interface{} `json:"state"`
	Title string      `json:"title"`
}

// NewHistory creates a new History instance
func NewHistory() *History {
	return &History{
		stack:   make([]HistoryEntry, 0, 50),
		current: -1,
		maxLen:  50, // Default browser history limit
	}
}

// PushState adds a new entry to the history stack
func (h *History) PushState(state interface{}, title, urlStr string) error {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	// Validate URL
	if urlStr != "" {
		if _, err := url.Parse(urlStr); err != nil {
			return fmt.Errorf("invalid URL: %w", err)
		}
	}

	// Create new entry
	entry := HistoryEntry{
		URL:   urlStr,
		State: state,
		Title: title,
	}

	// Remove any entries after current position
	if h.current >= 0 && h.current < len(h.stack)-1 {
		h.stack = h.stack[:h.current+1]
	}

	// Add new entry
	h.stack = append(h.stack, entry)
	h.current = len(h.stack) - 1

	// Enforce max length
	if len(h.stack) > h.maxLen {
		// Remove oldest entries
		excess := len(h.stack) - h.maxLen
		h.stack = h.stack[excess:]
		h.current -= excess
	}

	return nil
}

// ReplaceState replaces the current history entry
func (h *History) ReplaceState(state interface{}, title, urlStr string) error {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	// Validate URL
	if urlStr != "" {
		if _, err := url.Parse(urlStr); err != nil {
			return fmt.Errorf("invalid URL: %w", err)
		}
	}

	// If no entries exist, create initial entry
	if len(h.stack) == 0 {
		h.stack = append(h.stack, HistoryEntry{})
		h.current = 0
	}

	// Replace current entry
	if h.current >= 0 && h.current < len(h.stack) {
		h.stack[h.current] = HistoryEntry{
			URL:   urlStr,
			State: state,
			Title: title,
		}
	}

	return nil
}

// Back navigates back in history
func (h *History) Back() bool {
	return h.Go(-1)
}

// Forward navigates forward in history
func (h *History) Forward() bool {
	return h.Go(1)
}

// Go navigates by the specified delta
func (h *History) Go(delta int) bool {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	newPos := h.current + delta

	// Check bounds
	if newPos < 0 || newPos >= len(h.stack) {
		return false
	}

	// Get states directly without additional locking to avoid deadlock
	var oldState, newState interface{}
	if h.current >= 0 && h.current < len(h.stack) {
		oldState = h.stack[h.current].State
	}

	h.current = newPos

	if h.current >= 0 && h.current < len(h.stack) {
		newState = h.stack[h.current].State
	}

	// Trigger popstate event if callback is set and state changed
	if h.onPopState != nil && !statesEqual(oldState, newState) {
		// Call in goroutine to avoid deadlocks
		go h.onPopState(newState)
	}

	return true
}

// Length returns the number of entries in history
func (h *History) Length() int {
	h.mutex.RLock()
	defer h.mutex.RUnlock()
	return len(h.stack)
}

// GetState returns the current history state
func (h *History) GetState() interface{} {
	h.mutex.RLock()
	defer h.mutex.RUnlock()

	if h.current >= 0 && h.current < len(h.stack) {
		return h.stack[h.current].State
	}
	return nil
}

// GetURL returns the current history URL
func (h *History) GetURL() string {
	h.mutex.RLock()
	defer h.mutex.RUnlock()

	if h.current >= 0 && h.current < len(h.stack) {
		return h.stack[h.current].URL
	}
	return ""
}

// GetTitle returns the current history title
func (h *History) GetTitle() string {
	h.mutex.RLock()
	defer h.mutex.RUnlock()

	if h.current >= 0 && h.current < len(h.stack) {
		return h.stack[h.current].Title
	}
	return ""
}

// SetPopStateCallback sets the callback for popstate events
func (h *History) SetPopStateCallback(callback func(state interface{})) {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	h.onPopState = callback
}

// GetEntries returns a copy of all history entries (for testing)
func (h *History) GetEntries() []HistoryEntry {
	h.mutex.RLock()
	defer h.mutex.RUnlock()

	entries := make([]HistoryEntry, len(h.stack))
	copy(entries, h.stack)
	return entries
}

// GetCurrentIndex returns the current position in history (for testing)
func (h *History) GetCurrentIndex() int {
	h.mutex.RLock()
	defer h.mutex.RUnlock()
	return h.current
}

// statesEqual compares two states for equality
func statesEqual(a, b interface{}) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}

	// For complex objects, compare JSON representation
	aJSON, aErr := json.Marshal(a)
	bJSON, bErr := json.Marshal(b)

	if aErr != nil || bErr != nil {
		// Fallback to direct comparison
		return a == b
	}

	return string(aJSON) == string(bJSON)
}
