package performance

import (
	"sync"
	"time"
)

// Performance represents the browser Performance API
type Performance struct {
	navigationStart time.Time
	entries         []PerformanceEntry
	marks           map[string]PerformanceEntry
	mutex           sync.RWMutex
}

// PerformanceEntry represents a performance measurement
type PerformanceEntry struct {
	Name      string  `json:"name"`
	StartTime float64 `json:"startTime"`
	Duration  float64 `json:"duration"`
	EntryType string  `json:"entryType"`
}

// NewPerformance creates a new Performance instance
func NewPerformance() *Performance {
	return &Performance{
		navigationStart: time.Now(),
		entries:         make([]PerformanceEntry, 0),
		marks:           make(map[string]PerformanceEntry),
	}
}

// Now returns a high-resolution timestamp in milliseconds
func (p *Performance) Now() float64 {
	return float64(time.Since(p.navigationStart).Nanoseconds()) / 1000000.0
}

// Mark creates a named performance mark
func (p *Performance) Mark(name string) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	now := p.Now()
	entry := PerformanceEntry{
		Name:      name,
		StartTime: now,
		Duration:  0,
		EntryType: "mark",
	}

	p.marks[name] = entry
	p.entries = append(p.entries, entry)
}

// Measure creates a performance measure between two marks or timestamps
func (p *Performance) Measure(name, startMark, endMark string) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	var startTime, endTime float64
	now := p.Now()

	// Determine start time
	if startMark == "" {
		startTime = 0 // Start of navigation
	} else if mark, exists := p.marks[startMark]; exists {
		startTime = mark.StartTime
	} else {
		startTime = now // Use current time if mark doesn't exist
	}

	// Determine end time
	if endMark == "" {
		endTime = now // Current time
	} else if mark, exists := p.marks[endMark]; exists {
		endTime = mark.StartTime
	} else {
		endTime = now // Use current time if mark doesn't exist
	}

	// Create measure entry
	entry := PerformanceEntry{
		Name:      name,
		StartTime: startTime,
		Duration:  endTime - startTime,
		EntryType: "measure",
	}

	p.entries = append(p.entries, entry)
	return nil
}

// GetEntries returns all performance entries
func (p *Performance) GetEntries() []PerformanceEntry {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	// Return a copy to prevent external modification
	entries := make([]PerformanceEntry, len(p.entries))
	copy(entries, p.entries)
	return entries
}

// GetEntriesByType returns entries filtered by type
func (p *Performance) GetEntriesByType(entryType string) []PerformanceEntry {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	var filtered []PerformanceEntry
	for _, entry := range p.entries {
		if entry.EntryType == entryType {
			filtered = append(filtered, entry)
		}
	}
	return filtered
}

// GetEntriesByName returns entries filtered by name
func (p *Performance) GetEntriesByName(name string) []PerformanceEntry {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	var filtered []PerformanceEntry
	for _, entry := range p.entries {
		if entry.Name == name {
			filtered = append(filtered, entry)
		}
	}
	return filtered
}

// ClearMarks removes marks from the performance buffer
func (p *Performance) ClearMarks(markName string) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if markName == "" {
		// Clear all marks
		p.marks = make(map[string]PerformanceEntry)
		// Remove all mark entries
		filtered := make([]PerformanceEntry, 0)
		for _, entry := range p.entries {
			if entry.EntryType != "mark" {
				filtered = append(filtered, entry)
			}
		}
		p.entries = filtered
	} else {
		// Clear specific mark
		delete(p.marks, markName)
		// Remove specific mark entries
		filtered := make([]PerformanceEntry, 0)
		for _, entry := range p.entries {
			if !(entry.EntryType == "mark" && entry.Name == markName) {
				filtered = append(filtered, entry)
			}
		}
		p.entries = filtered
	}
}

// ClearMeasures removes measures from the performance buffer
func (p *Performance) ClearMeasures(measureName string) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if measureName == "" {
		// Clear all measures
		filtered := make([]PerformanceEntry, 0)
		for _, entry := range p.entries {
			if entry.EntryType != "measure" {
				filtered = append(filtered, entry)
			}
		}
		p.entries = filtered
	} else {
		// Clear specific measure
		filtered := make([]PerformanceEntry, 0)
		for _, entry := range p.entries {
			if !(entry.EntryType == "measure" && entry.Name == measureName) {
				filtered = append(filtered, entry)
			}
		}
		p.entries = filtered
	}
}

// ClearResourceTimings removes resource timing entries
func (p *Performance) ClearResourceTimings() {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	filtered := make([]PerformanceEntry, 0)
	for _, entry := range p.entries {
		if entry.EntryType != "resource" {
			filtered = append(filtered, entry)
		}
	}
	p.entries = filtered
}

// SetResourceBufferSize sets the maximum number of resource timing entries
func (p *Performance) SetResourceBufferSize(maxSize int) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	// This is a simplified implementation - in a real browser this would
	// control how many resource timing entries are kept
	if maxSize < len(p.entries) {
		// Keep the most recent entries
		p.entries = p.entries[len(p.entries)-maxSize:]
	}
}

// GetNavigationStart returns the navigation start time
func (p *Performance) GetNavigationStart() time.Time {
	return p.navigationStart
}

// AddResourceTiming adds a resource timing entry (for internal use)
func (p *Performance) AddResourceTiming(name string, startTime, duration float64) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	entry := PerformanceEntry{
		Name:      name,
		StartTime: startTime,
		Duration:  duration,
		EntryType: "resource",
	}

	p.entries = append(p.entries, entry)
}

// TimeOrigin returns the time origin for the performance timeline
func (p *Performance) TimeOrigin() float64 {
	// Return navigation start as milliseconds since Unix epoch
	return float64(p.navigationStart.UnixNano()) / 1000000.0
}
