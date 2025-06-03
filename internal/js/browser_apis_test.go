package js

import (
	"testing"

	"github.com/dop251/goja"
	"github.com/dpemmons/DOMulator/internal/dom"
)

func TestHistoryAPIIntegration(t *testing.T) {
	// Create JavaScript runtime with browser APIs
	vm := goja.New()
	document := dom.NewDocument()
	bindings := NewDOMBindings(vm, document)

	// Setup browser APIs including history
	bindings.SetupBrowserAPIs()
	bindings.SetupGlobalAPIs()

	// Test History API access
	script := `
		// Verify history object exists
		if (!window.history) {
			throw new Error("window.history not available");
		}

		// Test initial state
		var initialLength = window.history.length;
		
		// Test pushState
		window.history.pushState({page: 1}, "Page 1", "/page1");
		
		// Test that length increased
		if (window.history.length !== initialLength + 1) {
			throw new Error("History length did not increase after pushState");
		}
		
		// Test state property
		var currentState = window.history.state;
		if (!currentState || currentState.page !== 1) {
			throw new Error("History state not set correctly");
		}
		
		// Test replaceState
		window.history.replaceState({page: 2}, "Page 2", "/page2");
		
		// Length should remain the same after replace
		if (window.history.length !== initialLength + 1) {
			throw new Error("History length changed after replaceState");
		}
		
		// State should be updated
		currentState = window.history.state;
		if (!currentState || currentState.page !== 2) {
			throw new Error("History state not replaced correctly");
		}
		
		// Test navigation methods
		window.history.pushState({page: 3}, "Page 3", "/page3");
		
		// Test back navigation
		var backResult = window.history.back();
		if (!backResult) {
			throw new Error("History.back() returned false");
		}
		
		// Test forward navigation
		var forwardResult = window.history.forward();
		if (!forwardResult) {
			throw new Error("History.forward() returned false");
		}
		
		// Test go method
		window.history.pushState({page: 4}, "Page 4", "/page4");
		var goResult = window.history.go(-1);
		if (!goResult) {
			throw new Error("History.go(-1) returned false");
		}
		
		"History API test passed";
	`

	result, err := vm.RunString(script)
	if err != nil {
		t.Fatalf("History API test failed: %v", err)
	}

	if result.String() != "History API test passed" {
		t.Errorf("Expected success message, got: %s", result.String())
	}
}

func TestPerformanceAPIIntegration(t *testing.T) {
	// Create JavaScript runtime with browser APIs
	vm := goja.New()
	document := dom.NewDocument()
	bindings := NewDOMBindings(vm, document)

	// Setup browser APIs including performance
	bindings.SetupBrowserAPIs()
	bindings.SetupGlobalAPIs()

	// Test Performance API access
	script := `
		// Verify performance object exists
		if (!window.performance) {
			throw new Error("window.performance not available");
		}

		// Test performance.now()
		var start = window.performance.now();
		if (typeof start !== 'number' || start < 0) {
			throw new Error("performance.now() did not return a valid timestamp");
		}
		
		// Test performance.mark()
		window.performance.mark("test-mark-1");
		window.performance.mark("test-mark-2");
		
		// Test performance.measure()
		window.performance.measure("test-measure", "test-mark-1", "test-mark-2");
		
		// Test performance.getEntries()
		var entries = window.performance.getEntries();
		if (!Array.isArray(entries) || entries.length < 3) {
			throw new Error("getEntries() did not return expected entries");
		}
		
		// Verify entries have correct structure
		var markFound = false;
		var measureFound = false;
		for (var i = 0; i < entries.length; i++) {
			var entry = entries[i];
			if (!entry.name || typeof entry.startTime !== 'number' || typeof entry.duration !== 'number' || !entry.entryType) {
				throw new Error("Entry missing required properties: " + JSON.stringify(entry));
			}
			if (entry.entryType === 'mark' && entry.name === 'test-mark-1') {
				markFound = true;
			}
			if (entry.entryType === 'measure' && entry.name === 'test-measure') {
				measureFound = true;
			}
		}
		
		if (!markFound) {
			throw new Error("Expected mark not found in entries");
		}
		if (!measureFound) {
			throw new Error("Expected measure not found in entries");
		}
		
		// Test performance.getEntriesByType()
		var marks = window.performance.getEntriesByType("mark");
		if (!Array.isArray(marks) || marks.length < 2) {
			throw new Error("getEntriesByType('mark') did not return expected entries");
		}
		
		var measures = window.performance.getEntriesByType("measure");
		if (!Array.isArray(measures) || measures.length < 1) {
			throw new Error("getEntriesByType('measure') did not return expected entries");
		}
		
		// Test performance.getEntriesByName()
		var namedEntries = window.performance.getEntriesByName("test-mark-1");
		if (!Array.isArray(namedEntries) || namedEntries.length !== 1) {
			throw new Error("getEntriesByName() did not return expected entries");
		}
		
		// Test clearMarks()
		window.performance.clearMarks("test-mark-1");
		var remainingMarks = window.performance.getEntriesByName("test-mark-1");
		if (remainingMarks.length !== 0) {
			throw new Error("clearMarks() did not remove the specified mark");
		}
		
		// Test clearMeasures()
		window.performance.clearMeasures("test-measure");
		var remainingMeasures = window.performance.getEntriesByName("test-measure");
		if (remainingMeasures.length !== 0) {
			throw new Error("clearMeasures() did not remove the specified measure");
		}
		
		// Test clearMarks() with no argument (clear all)
		window.performance.clearMarks();
		var allMarks = window.performance.getEntriesByType("mark");
		if (allMarks.length !== 0) {
			throw new Error("clearMarks() without arguments did not clear all marks");
		}
		
		// Test timeOrigin property
		var timeOrigin = window.performance.timeOrigin;
		if (typeof timeOrigin !== 'number' || timeOrigin <= 0) {
			throw new Error("timeOrigin is not a valid timestamp");
		}
		
		"Performance API test passed";
	`

	result, err := vm.RunString(script)
	if err != nil {
		t.Fatalf("Performance API test failed: %v", err)
	}

	if result.String() != "Performance API test passed" {
		t.Errorf("Expected success message, got: %s", result.String())
	}
}

func TestBrowserAPIsCoexistence(t *testing.T) {
	// Create JavaScript runtime with all browser APIs
	vm := goja.New()
	document := dom.NewDocument()
	bindings := NewDOMBindings(vm, document)

	// Setup all browser APIs
	bindings.SetupBrowserAPIs()
	bindings.SetupGlobalAPIs()

	// Test that all APIs coexist without conflicts
	script := `
		// Test all APIs are available
		var apis = {
			history: typeof window.history,
			performance: typeof window.performance,
			URL: typeof URL,
			URLSearchParams: typeof URLSearchParams,
			CustomEvent: typeof CustomEvent
		};
		
		// Verify all APIs are objects/functions
		for (var api in apis) {
			if (apis[api] !== 'object' && apis[api] !== 'function') {
				throw new Error(api + " API not properly exposed: " + apis[api]);
			}
		}
		
		// Test concurrent usage
		var perfStart = window.performance.now();
		window.history.pushState({test: true}, "Test", "/test");
		window.performance.mark("after-history");
		
		var url = new URL("https://example.com/path?query=value");
		window.performance.mark("after-url");
		
		var params = new URLSearchParams("key=value");
		window.performance.mark("after-params");
		
		var perfEnd = window.performance.now();
		window.performance.measure("total-test", "after-history");
		
		// Verify everything still works
		if (window.history.state.test !== true) {
			throw new Error("History state was corrupted");
		}
		
		if (url.hostname !== "example.com") {
			throw new Error("URL parsing was corrupted");
		}
		
		if (params.get("key") !== "value") {
			throw new Error("URLSearchParams was corrupted");
		}
		
		if (perfEnd <= perfStart) {
			throw new Error("Performance timing was corrupted");
		}
		
		var entries = window.performance.getEntries();
		if (entries.length === 0) {
			throw new Error("Performance entries were lost");
		}
		
		"All APIs coexist successfully";
	`

	result, err := vm.RunString(script)
	if err != nil {
		t.Fatalf("APIs coexistence test failed: %v", err)
	}

	if result.String() != "All APIs coexist successfully" {
		t.Errorf("Expected success message, got: %s", result.String())
	}
}

func TestPerformanceTimingAccuracy(t *testing.T) {
	// Create JavaScript runtime with performance API
	vm := goja.New()
	document := dom.NewDocument()
	bindings := NewDOMBindings(vm, document)

	bindings.SetupBrowserAPIs()
	bindings.SetupGlobalAPIs()

	// Test timing accuracy and behavior
	script := `
		// Test that consecutive now() calls return increasing values
		var times = [];
		for (var i = 0; i < 5; i++) {
			times.push(window.performance.now());
		}
		
		// Verify times are increasing
		for (var i = 1; i < times.length; i++) {
			if (times[i] < times[i-1]) {
				throw new Error("Performance.now() returned decreasing values: " + times[i-1] + " -> " + times[i]);
			}
		}
		
		// Test mark timing
		var beforeMark = window.performance.now();
		window.performance.mark("timing-test");
		var afterMark = window.performance.now();
		
		var entries = window.performance.getEntriesByName("timing-test");
		if (entries.length !== 1) {
			throw new Error("Expected exactly one mark entry");
		}
		
		var markEntry = entries[0];
		if (markEntry.startTime < beforeMark || markEntry.startTime > afterMark) {
			throw new Error("Mark timestamp out of expected range: " + markEntry.startTime + " not between " + beforeMark + " and " + afterMark);
		}
		
		// Test measure duration
		window.performance.mark("start-mark");
		window.performance.mark("end-mark");
		window.performance.measure("duration-test", "start-mark", "end-mark");
		
		var measureEntries = window.performance.getEntriesByName("duration-test");
		if (measureEntries.length !== 1) {
			throw new Error("Expected exactly one measure entry");
		}
		
		var measureEntry = measureEntries[0];
		if (measureEntry.duration < 0) {
			throw new Error("Measure duration should not be negative: " + measureEntry.duration);
		}
		
		"Performance timing accuracy test passed";
	`

	result, err := vm.RunString(script)
	if err != nil {
		t.Fatalf("Performance timing accuracy test failed: %v", err)
	}

	if result.String() != "Performance timing accuracy test passed" {
		t.Errorf("Expected success message, got: %s", result.String())
	}
}
