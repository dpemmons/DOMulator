package history

import (
	"sync"
	"testing"
	"time"
)

func TestNewHistory(t *testing.T) {
	h := NewHistory()

	if h == nil {
		t.Fatal("NewHistory returned nil")
	}

	if h.Length() != 0 {
		t.Errorf("Expected length 0, got %d", h.Length())
	}

	if h.GetCurrentIndex() != -1 {
		t.Errorf("Expected current index -1, got %d", h.GetCurrentIndex())
	}

	if h.GetState() != nil {
		t.Errorf("Expected nil state, got %v", h.GetState())
	}
}

func TestPushState(t *testing.T) {
	h := NewHistory()

	// Test basic push
	err := h.PushState("test-state", "Test Title", "/test")
	if err != nil {
		t.Fatalf("PushState failed: %v", err)
	}

	if h.Length() != 1 {
		t.Errorf("Expected length 1, got %d", h.Length())
	}

	if h.GetCurrentIndex() != 0 {
		t.Errorf("Expected current index 0, got %d", h.GetCurrentIndex())
	}

	if h.GetState() != "test-state" {
		t.Errorf("Expected state 'test-state', got %v", h.GetState())
	}

	if h.GetURL() != "/test" {
		t.Errorf("Expected URL '/test', got %s", h.GetURL())
	}

	if h.GetTitle() != "Test Title" {
		t.Errorf("Expected title 'Test Title', got %s", h.GetTitle())
	}
}

func TestPushStateMultiple(t *testing.T) {
	h := NewHistory()

	// Push multiple entries
	states := []string{"state1", "state2", "state3"}
	urls := []string{"/page1", "/page2", "/page3"}
	titles := []string{"Page 1", "Page 2", "Page 3"}

	for i, state := range states {
		err := h.PushState(state, titles[i], urls[i])
		if err != nil {
			t.Fatalf("PushState %d failed: %v", i, err)
		}
	}

	if h.Length() != 3 {
		t.Errorf("Expected length 3, got %d", h.Length())
	}

	if h.GetCurrentIndex() != 2 {
		t.Errorf("Expected current index 2, got %d", h.GetCurrentIndex())
	}

	if h.GetState() != "state3" {
		t.Errorf("Expected state 'state3', got %v", h.GetState())
	}
}

func TestPushStateInvalidURL(t *testing.T) {
	h := NewHistory()

	// Test invalid URL
	err := h.PushState("state", "title", "://invalid-url")
	if err == nil {
		t.Error("Expected error for invalid URL, got nil")
	}
}

func TestPushStateEmptyURL(t *testing.T) {
	h := NewHistory()

	// Test empty URL (should be allowed)
	err := h.PushState("state", "title", "")
	if err != nil {
		t.Errorf("Expected no error for empty URL, got %v", err)
	}

	if h.GetURL() != "" {
		t.Errorf("Expected empty URL, got %s", h.GetURL())
	}
}

func TestReplaceState(t *testing.T) {
	h := NewHistory()

	// Push initial state
	h.PushState("initial", "Initial", "/initial")

	// Replace it
	err := h.ReplaceState("replaced", "Replaced", "/replaced")
	if err != nil {
		t.Fatalf("ReplaceState failed: %v", err)
	}

	if h.Length() != 1 {
		t.Errorf("Expected length 1, got %d", h.Length())
	}

	if h.GetState() != "replaced" {
		t.Errorf("Expected state 'replaced', got %v", h.GetState())
	}

	if h.GetURL() != "/replaced" {
		t.Errorf("Expected URL '/replaced', got %s", h.GetURL())
	}
}

func TestReplaceStateEmpty(t *testing.T) {
	h := NewHistory()

	// Replace state on empty history
	err := h.ReplaceState("new-state", "New", "/new")
	if err != nil {
		t.Fatalf("ReplaceState on empty history failed: %v", err)
	}

	if h.Length() != 1 {
		t.Errorf("Expected length 1, got %d", h.Length())
	}

	if h.GetState() != "new-state" {
		t.Errorf("Expected state 'new-state', got %v", h.GetState())
	}
}

func TestNavigation(t *testing.T) {
	h := NewHistory()

	// Create history stack
	states := []string{"state1", "state2", "state3", "state4"}
	for i, state := range states {
		h.PushState(state, "", "")
		if i == len(states)-1 {
			continue // Don't check intermediate states
		}
	}

	// Should be at last entry
	if h.GetCurrentIndex() != 3 {
		t.Errorf("Expected current index 3, got %d", h.GetCurrentIndex())
	}

	// Go back
	if !h.Back() {
		t.Error("Back() returned false")
	}

	if h.GetCurrentIndex() != 2 {
		t.Errorf("Expected current index 2 after back, got %d", h.GetCurrentIndex())
	}

	if h.GetState() != "state3" {
		t.Errorf("Expected state 'state3' after back, got %v", h.GetState())
	}

	// Go back by 2
	if !h.Go(-2) {
		t.Error("Go(-2) returned false")
	}

	if h.GetCurrentIndex() != 0 {
		t.Errorf("Expected current index 0 after Go(-2), got %d", h.GetCurrentIndex())
	}

	if h.GetState() != "state1" {
		t.Errorf("Expected state 'state1' after Go(-2), got %v", h.GetState())
	}

	// Go forward
	if !h.Forward() {
		t.Error("Forward() returned false")
	}

	if h.GetCurrentIndex() != 1 {
		t.Errorf("Expected current index 1 after forward, got %d", h.GetCurrentIndex())
	}

	if h.GetState() != "state2" {
		t.Errorf("Expected state 'state2' after forward, got %v", h.GetState())
	}
}

func TestNavigationBounds(t *testing.T) {
	h := NewHistory()

	// Test navigation on empty history
	if h.Back() {
		t.Error("Back() on empty history should return false")
	}

	if h.Forward() {
		t.Error("Forward() on empty history should return false")
	}

	if h.Go(-1) {
		t.Error("Go(-1) on empty history should return false")
	}

	// Add one entry
	h.PushState("state1", "", "")

	// Test bounds
	if h.Back() {
		t.Error("Back() at first entry should return false")
	}

	if h.Forward() {
		t.Error("Forward() at last entry should return false")
	}

	if h.Go(-1) {
		t.Error("Go(-1) beyond bounds should return false")
	}

	if h.Go(1) {
		t.Error("Go(1) beyond bounds should return false")
	}
}

func TestPushAfterNavigation(t *testing.T) {
	h := NewHistory()

	// Create stack: [state1, state2, state3]
	h.PushState("state1", "", "")
	h.PushState("state2", "", "")
	h.PushState("state3", "", "")

	// Go back to state1
	h.Go(-2)

	if h.GetCurrentIndex() != 0 {
		t.Errorf("Expected current index 0, got %d", h.GetCurrentIndex())
	}

	// Push new state - should truncate forward history
	h.PushState("state4", "", "")

	if h.Length() != 2 {
		t.Errorf("Expected length 2 after push, got %d", h.Length())
	}

	if h.GetCurrentIndex() != 1 {
		t.Errorf("Expected current index 1 after push, got %d", h.GetCurrentIndex())
	}

	if h.GetState() != "state4" {
		t.Errorf("Expected state 'state4', got %v", h.GetState())
	}

	// Verify truncation
	entries := h.GetEntries()
	if len(entries) != 2 {
		t.Errorf("Expected 2 entries, got %d", len(entries))
	}

	if entries[0].State != "state1" || entries[1].State != "state4" {
		t.Errorf("History not properly truncated: %v", entries)
	}
}

func TestMaxLength(t *testing.T) {
	h := NewHistory()
	h.maxLen = 3 // Set small limit for testing

	// Push more entries than limit
	for i := 0; i < 5; i++ {
		h.PushState(i, "", "")
	}

	if h.Length() != 3 {
		t.Errorf("Expected length 3 (max), got %d", h.Length())
	}

	// Should have entries 2, 3, 4 (oldest removed)
	entries := h.GetEntries()
	expectedStates := []int{2, 3, 4}

	for i, expected := range expectedStates {
		if entries[i].State != expected {
			t.Errorf("Entry %d: expected state %d, got %v", i, expected, entries[i].State)
		}
	}

	if h.GetCurrentIndex() != 2 {
		t.Errorf("Expected current index 2, got %d", h.GetCurrentIndex())
	}

	if h.GetState() != 4 {
		t.Errorf("Expected current state 4, got %v", h.GetState())
	}
}

func TestPopStateCallback(t *testing.T) {
	h := NewHistory()

	var callbackCalled bool
	var callbackState interface{}
	var wg sync.WaitGroup

	// Set callback
	h.SetPopStateCallback(func(state interface{}) {
		defer wg.Done()
		callbackCalled = true
		callbackState = state
	})

	// Create history
	h.PushState("state1", "", "")
	h.PushState("state2", "", "")

	// Navigate back - should trigger callback
	wg.Add(1)
	if !h.Back() {
		t.Error("Back() failed")
	}

	// Wait for callback (with timeout)
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		// Callback completed
	case <-time.After(100 * time.Millisecond):
		t.Error("Callback not called within timeout")
	}

	if !callbackCalled {
		t.Error("Callback was not called")
	}

	if callbackState != "state1" {
		t.Errorf("Expected callback state 'state1', got %v", callbackState)
	}
}

func TestPopStateCallbackNoChange(t *testing.T) {
	h := NewHistory()

	var callbackCalled bool

	// Set callback
	h.SetPopStateCallback(func(state interface{}) {
		callbackCalled = true
	})

	// Navigate with no state change
	h.PushState("same-state", "", "")
	h.ReplaceState("same-state", "", "") // Replace with same state

	// Navigate - should not trigger callback since state didn't change
	h.Go(0) // Navigate to same position

	// Give callback time to be called (if it would be)
	time.Sleep(10 * time.Millisecond)

	if callbackCalled {
		t.Error("Callback should not be called when state doesn't change")
	}
}

func TestComplexState(t *testing.T) {
	h := NewHistory()

	// Test complex state object
	complexState := map[string]interface{}{
		"user":  "john",
		"count": 42,
		"data":  []string{"a", "b", "c"},
	}

	err := h.PushState(complexState, "Complex", "/complex")
	if err != nil {
		t.Fatalf("PushState with complex state failed: %v", err)
	}

	retrievedState := h.GetState()
	if retrievedState == nil {
		t.Fatal("Retrieved state is nil")
	}

	// Should be able to navigate and preserve complex state
	h.PushState("simple", "", "")
	h.Back()

	if !statesEqual(h.GetState(), complexState) {
		t.Errorf("Complex state not preserved: expected %v, got %v", complexState, h.GetState())
	}
}

func TestConcurrency(t *testing.T) {
	h := NewHistory()

	var wg sync.WaitGroup
	numGoroutines := 10

	// Concurrent pushes
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			h.PushState(id, "", "")
		}(i)
	}

	wg.Wait()

	if h.Length() != numGoroutines {
		t.Errorf("Expected length %d, got %d", numGoroutines, h.Length())
	}

	// Concurrent reads
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			h.GetState()
			h.GetURL()
			h.GetTitle()
			h.Length()
		}()
	}

	wg.Wait()
}

func TestStatesEqual(t *testing.T) {
	// Test nil states
	if !statesEqual(nil, nil) {
		t.Error("nil states should be equal")
	}

	if statesEqual(nil, "test") {
		t.Error("nil and non-nil should not be equal")
	}

	// Test primitive states
	if !statesEqual("test", "test") {
		t.Error("Same strings should be equal")
	}

	if statesEqual("test1", "test2") {
		t.Error("Different strings should not be equal")
	}

	// Test complex states
	state1 := map[string]interface{}{"a": 1, "b": 2}
	state2 := map[string]interface{}{"a": 1, "b": 2}
	state3 := map[string]interface{}{"a": 1, "b": 3}

	if !statesEqual(state1, state2) {
		t.Error("Equivalent complex states should be equal")
	}

	if statesEqual(state1, state3) {
		t.Error("Different complex states should not be equal")
	}
}
