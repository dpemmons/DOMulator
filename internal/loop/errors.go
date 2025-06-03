package loop

import "errors"

// Common event loop errors
var (
	ErrAlreadyRunning = errors.New("event loop is already running")
	ErrNotRunning     = errors.New("event loop is not running")
	ErrInvalidTask    = errors.New("invalid task")
	ErrQueueFull      = errors.New("queue is full")
)
