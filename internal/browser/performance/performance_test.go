package performance

import (
	"testing"
	"time"
)

func TestNewPerformance(t *testing.T) {
	p := NewPerformance()

	if p == nil {
		t.Fatal("NewPerformance returned nil")
	}

	if len(p.GetEntries()) != 0 {
		t.Errorf("Expected 0 entries, got %d", len(p.GetEntries()))
	}

	// Navigation start should be recent
	if time.Since(p.GetNavigationStart()) > time.Second {
		t.Error("Navigation start should be recent")
	}
}

func TestNow(t *testing.T) {
	p := NewPerformance()

	start := p.Now()
	time.Sleep(1 * time.Millisecond)
	end := p.Now()

	if end <= start {
		t.Errorf("Expected end (%f) > start (%f)", end, start)
	}

	// Should be roughly 1ms difference
	diff := end - start
	if diff < 0.5 || diff > 5.0 {
		t.Errorf("Expected ~1ms difference, got %f ms", diff)
	}
}

func TestMark(t *testing.T) {
	p := NewPerformance()

	// Create a mark
	p.Mark("test-mark")

	entries := p.GetEntries()
	if len(entries) != 1 {
		t.Fatalf("Expected 1 entry, got %d", len(entries))
	}

	entry := entries[0]
	if entry.Name != "test-mark" {
		t.Errorf("Expected name 'test-mark', got %s", entry.Name)
	}

	if entry.EntryType != "mark" {
		t.Errorf("Expected type 'mark', got %s", entry.EntryType)
	}

	if entry.Duration != 0 {
		t.Errorf("Expected duration 0, got %f", entry.Duration)
	}

	if entry.StartTime < 0 {
		t.Errorf("Expected positive start time, got %f", entry.StartTime)
	}
}

func TestMultipleMarks(t *testing.T) {
	p := NewPerformance()

	p.Mark("mark1")
	time.Sleep(1 * time.Millisecond)
	p.Mark("mark2")
	time.Sleep(1 * time.Millisecond)
	p.Mark("mark3")

	entries := p.GetEntries()
	if len(entries) != 3 {
		t.Fatalf("Expected 3 entries, got %d", len(entries))
	}

	// Check marks are in chronological order
	for i := 1; i < len(entries); i++ {
		if entries[i].StartTime <= entries[i-1].StartTime {
			t.Errorf("Marks not in chronological order: %f <= %f",
				entries[i].StartTime, entries[i-1].StartTime)
		}
	}
}

func TestMeasure(t *testing.T) {
	p := NewPerformance()

	// Create marks
	p.Mark("start")
	time.Sleep(2 * time.Millisecond)
	p.Mark("end")

	// Create measure
	err := p.Measure("test-measure", "start", "end")
	if err != nil {
		t.Fatalf("Measure failed: %v", err)
	}

	entries := p.GetEntries()
	if len(entries) != 3 { // 2 marks + 1 measure
		t.Fatalf("Expected 3 entries, got %d", len(entries))
	}

	// Find the measure entry
	var measure *PerformanceEntry
	for _, entry := range entries {
		if entry.EntryType == "measure" {
			measure = &entry
			break
		}
	}

	if measure == nil {
		t.Fatal("Measure entry not found")
	}

	if measure.Name != "test-measure" {
		t.Errorf("Expected name 'test-measure', got %s", measure.Name)
	}

	if measure.Duration <= 0 {
		t.Errorf("Expected positive duration, got %f", measure.Duration)
	}

	// Duration should be roughly 2ms
	if measure.Duration < 1.0 || measure.Duration > 10.0 {
		t.Errorf("Expected ~2ms duration, got %f ms", measure.Duration)
	}
}

func TestMeasureNoMarks(t *testing.T) {
	p := NewPerformance()

	// Measure without creating marks (should use current time)
	err := p.Measure("no-marks", "", "")
	if err != nil {
		t.Fatalf("Measure failed: %v", err)
	}

	entries := p.GetEntries()
	if len(entries) != 1 {
		t.Fatalf("Expected 1 entry, got %d", len(entries))
	}

	measure := entries[0]
	if measure.EntryType != "measure" {
		t.Errorf("Expected type 'measure', got %s", measure.EntryType)
	}

	// Duration should be small (near zero)
	if measure.Duration < 0 || measure.Duration > 1.0 {
		t.Errorf("Expected small duration, got %f ms", measure.Duration)
	}
}

func TestMeasureNonExistentMarks(t *testing.T) {
	p := NewPerformance()

	// Measure with non-existent marks (should use current time)
	err := p.Measure("missing-marks", "missing-start", "missing-end")
	if err != nil {
		t.Fatalf("Measure failed: %v", err)
	}

	entries := p.GetEntries()
	if len(entries) != 1 {
		t.Fatalf("Expected 1 entry, got %d", len(entries))
	}

	measure := entries[0]
	if measure.EntryType != "measure" {
		t.Errorf("Expected type 'measure', got %s", measure.EntryType)
	}
}

func TestGetEntriesByType(t *testing.T) {
	p := NewPerformance()

	// Create mixed entries
	p.Mark("mark1")
	p.Mark("mark2")
	p.Measure("measure1", "", "")
	p.AddResourceTiming("resource1", 0, 10)

	// Get marks only
	marks := p.GetEntriesByType("mark")
	if len(marks) != 2 {
		t.Errorf("Expected 2 marks, got %d", len(marks))
	}

	for _, mark := range marks {
		if mark.EntryType != "mark" {
			t.Errorf("Expected mark type, got %s", mark.EntryType)
		}
	}

	// Get measures only
	measures := p.GetEntriesByType("measure")
	if len(measures) != 1 {
		t.Errorf("Expected 1 measure, got %d", len(measures))
	}

	// Get resources only
	resources := p.GetEntriesByType("resource")
	if len(resources) != 1 {
		t.Errorf("Expected 1 resource, got %d", len(resources))
	}

	// Get non-existent type
	none := p.GetEntriesByType("nonexistent")
	if len(none) != 0 {
		t.Errorf("Expected 0 entries for nonexistent type, got %d", len(none))
	}
}

func TestGetEntriesByName(t *testing.T) {
	p := NewPerformance()

	// Create entries with same name
	p.Mark("test")
	p.Measure("test", "", "")
	p.Mark("other")

	entries := p.GetEntriesByName("test")
	if len(entries) != 2 {
		t.Errorf("Expected 2 entries named 'test', got %d", len(entries))
	}

	for _, entry := range entries {
		if entry.Name != "test" {
			t.Errorf("Expected name 'test', got %s", entry.Name)
		}
	}

	// Get non-existent name
	none := p.GetEntriesByName("nonexistent")
	if len(none) != 0 {
		t.Errorf("Expected 0 entries for nonexistent name, got %d", len(none))
	}
}

func TestClearMarks(t *testing.T) {
	p := NewPerformance()

	p.Mark("mark1")
	p.Mark("mark2")
	p.Measure("measure1", "", "")

	// Clear specific mark
	p.ClearMarks("mark1")

	entries := p.GetEntries()
	if len(entries) != 2 { // mark2 + measure1
		t.Errorf("Expected 2 entries after clearing mark1, got %d", len(entries))
	}

	// Verify mark1 is gone
	for _, entry := range entries {
		if entry.Name == "mark1" {
			t.Error("mark1 should have been cleared")
		}
	}

	// Clear all marks
	p.ClearMarks("")

	entries = p.GetEntries()
	if len(entries) != 1 { // only measure1
		t.Errorf("Expected 1 entry after clearing all marks, got %d", len(entries))
	}

	if entries[0].EntryType != "measure" {
		t.Errorf("Expected remaining entry to be measure, got %s", entries[0].EntryType)
	}
}

func TestClearMeasures(t *testing.T) {
	p := NewPerformance()

	p.Mark("mark1")
	p.Measure("measure1", "", "")
	p.Measure("measure2", "", "")

	// Clear specific measure
	p.ClearMeasures("measure1")

	measures := p.GetEntriesByType("measure")
	if len(measures) != 1 {
		t.Errorf("Expected 1 measure after clearing measure1, got %d", len(measures))
	}

	if measures[0].Name != "measure2" {
		t.Errorf("Expected remaining measure to be measure2, got %s", measures[0].Name)
	}

	// Clear all measures
	p.ClearMeasures("")

	measures = p.GetEntriesByType("measure")
	if len(measures) != 0 {
		t.Errorf("Expected 0 measures after clearing all, got %d", len(measures))
	}

	// Mark should still exist
	marks := p.GetEntriesByType("mark")
	if len(marks) != 1 {
		t.Errorf("Expected 1 mark to remain, got %d", len(marks))
	}
}

func TestClearResourceTimings(t *testing.T) {
	p := NewPerformance()

	p.Mark("mark1")
	p.AddResourceTiming("resource1", 0, 10)
	p.AddResourceTiming("resource2", 5, 15)

	p.ClearResourceTimings()

	resources := p.GetEntriesByType("resource")
	if len(resources) != 0 {
		t.Errorf("Expected 0 resources after clearing, got %d", len(resources))
	}

	// Mark should still exist
	marks := p.GetEntriesByType("mark")
	if len(marks) != 1 {
		t.Errorf("Expected 1 mark to remain, got %d", len(marks))
	}
}

func TestAddResourceTiming(t *testing.T) {
	p := NewPerformance()

	p.AddResourceTiming("test-resource", 100.5, 50.25)

	entries := p.GetEntries()
	if len(entries) != 1 {
		t.Fatalf("Expected 1 entry, got %d", len(entries))
	}

	resource := entries[0]
	if resource.Name != "test-resource" {
		t.Errorf("Expected name 'test-resource', got %s", resource.Name)
	}

	if resource.EntryType != "resource" {
		t.Errorf("Expected type 'resource', got %s", resource.EntryType)
	}

	if resource.StartTime != 100.5 {
		t.Errorf("Expected start time 100.5, got %f", resource.StartTime)
	}

	if resource.Duration != 50.25 {
		t.Errorf("Expected duration 50.25, got %f", resource.Duration)
	}
}

func TestSetResourceBufferSize(t *testing.T) {
	p := NewPerformance()

	// Add many resource entries
	for i := 0; i < 10; i++ {
		p.AddResourceTiming("resource", float64(i), 1.0)
	}

	if len(p.GetEntries()) != 10 {
		t.Fatalf("Expected 10 entries, got %d", len(p.GetEntries()))
	}

	// Limit to 5 entries
	p.SetResourceBufferSize(5)

	entries := p.GetEntries()
	if len(entries) != 5 {
		t.Errorf("Expected 5 entries after buffer size limit, got %d", len(entries))
	}

	// Should keep the most recent entries (5-9)
	for i, entry := range entries {
		expectedStartTime := float64(i + 5)
		if entry.StartTime != expectedStartTime {
			t.Errorf("Entry %d: expected start time %f, got %f",
				i, expectedStartTime, entry.StartTime)
		}
	}
}

func TestTimeOrigin(t *testing.T) {
	p := NewPerformance()

	origin := p.TimeOrigin()

	// Should be a reasonable timestamp (milliseconds since Unix epoch)
	now := float64(time.Now().UnixNano()) / 1000000.0

	if origin <= 0 {
		t.Error("Time origin should be positive")
	}

	// Should be within a few seconds of current time
	diff := now - origin
	if diff < 0 || diff > 5000 { // 5 seconds tolerance
		t.Errorf("Time origin seems incorrect: diff=%f ms", diff)
	}
}

func TestConcurrency(t *testing.T) {
	p := NewPerformance()

	done := make(chan struct{})
	numGoroutines := 10

	// Concurrent operations
	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			defer func() { done <- struct{}{} }()

			p.Mark("mark")
			p.Measure("measure", "", "")
			p.GetEntries()
			p.GetEntriesByType("mark")
			p.GetEntriesByName("mark")
		}(i)
	}

	// Wait for all goroutines
	for i := 0; i < numGoroutines; i++ {
		<-done
	}

	// Verify we have entries (exact count may vary due to timing)
	entries := p.GetEntries()
	if len(entries) == 0 {
		t.Error("Expected some entries after concurrent operations")
	}
}

func TestNowPrecision(t *testing.T) {
	p := NewPerformance()

	// Test precision by taking multiple measurements
	times := make([]float64, 100)
	for i := 0; i < 100; i++ {
		times[i] = p.Now()
	}

	// Verify times are increasing
	for i := 1; i < len(times); i++ {
		if times[i] < times[i-1] {
			t.Errorf("Time went backwards: %f -> %f", times[i-1], times[i])
		}
	}

	// Verify reasonable precision (should be sub-millisecond)
	totalTime := times[99] - times[0]
	if totalTime <= 0 {
		t.Error("Expected positive total time")
	}

	// Should complete in reasonable time (less than 10ms)
	if totalTime > 10 {
		t.Errorf("Expected completion in <10ms, took %f ms", totalTime)
	}
}
