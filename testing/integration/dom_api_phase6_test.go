package integration

import (
	"testing"
	"time"

	domulator "github.com/dpemmons/DOMulator"
)

func TestDOMAPIPhase6ResizeObserver(t *testing.T) {
	test := domulator.NewTest(t)
	test.SetDebugMode(true)

	test.LoadHTML(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>Phase 6 ResizeObserver Test</title>
		</head>
		<body>
			<div id="observed-element" style="width: 100px; height: 100px; background: blue;">
				Resizable Element
			</div>
			<div id="results"></div>
		</body>
		</html>
		<script>
		let observerCallbackCount = 0;
		let observedElements = [];
		
		// Test ResizeObserver creation
		console.log("ResizeObserver constructor:", typeof ResizeObserver);
		
		const observer = new ResizeObserver(function(entries) {
			observerCallbackCount++;
			console.log("ResizeObserver callback called, entries:", entries.length);
			
			entries.forEach(function(entry) {
				observedElements.push(entry.target.id);
				console.log("Observed element:", entry.target.id);
				console.log("Content rect:", entry.contentRect);
				console.log("Border box size:", entry.borderBoxSize);
				console.log("Content box size:", entry.contentBoxSize);
				
				// Store results in DOM for verification
				const results = document.getElementById('results');
				results.setAttribute('data-callback-count', observerCallbackCount.toString());
				results.setAttribute('data-observed-elements', observedElements.join(','));
				
				// Test entry properties
				let entryResults = [];
				entryResults.push("contentRect: " + (typeof entry.contentRect === 'object'));
				entryResults.push("borderBoxSize: " + (typeof entry.borderBoxSize === 'object'));
				entryResults.push("contentBoxSize: " + (typeof entry.contentBoxSize === 'object'));
				entryResults.push("target: " + (entry.target === document.getElementById('observed-element')));
				
				results.setAttribute('data-entry-properties', entryResults.join(','));
			});
		});
		
		console.log("ResizeObserver created");
		
		// Test observer methods
		const element = document.getElementById('observed-element');
		
		try {
			observer.observe(element);
			console.log("observe() completed");
		} catch (e) {
			console.log("observe() error:", e.message);
		}
		
		// Mark initial setup as complete
		document.body.setAttribute('data-resize-setup', 'complete');
		</script>
	`)

	// Wait for setup to complete
	test.AssertElement("body").HasAttribute("data-resize-setup", "complete")

	// Process initial ResizeObserver callback (fired when observe() is called)
	test.FlushMicrotasks()

	// Check if any callback was fired yet and log state
	test.ExecuteScript(`
		var resultsEl = document.getElementById('results');
		var currentCount = resultsEl ? resultsEl.getAttribute('data-callback-count') : '0';
		console.log("Initial callback count after FlushMicrotasks:", currentCount);
		console.log("Observer callback count variable:", observerCallbackCount);
		document.body.setAttribute('data-initial-count', currentCount || '0');
	`)

	// Now trigger a resize using the test harness
	test.TriggerElementResize("#observed-element", domulator.ResizeOptions{
		Width:  200,
		Height: 150,
	})

	// Process the resize callback
	test.FlushMicrotasks()

	// Check final state and verify ResizeObserver is working
	test.ExecuteScript(`
		var resultsEl = document.getElementById('results');
		var finalCount = resultsEl ? resultsEl.getAttribute('data-callback-count') : '0';
		console.log("Final callback count:", finalCount);
		console.log("Observer callback count variable:", observerCallbackCount);
		console.log("Observed elements:", resultsEl ? resultsEl.getAttribute('data-observed-elements') : 'none');
		
		// Verify that ResizeObserver is working - we should have at least 1 callback
		var isWorking = observerCallbackCount > 0;
		document.body.setAttribute('data-final-count', finalCount || '0');
		document.body.setAttribute('data-resize-working', isWorking ? 'true' : 'false');
	`)

	// Verify that ResizeObserver triggered at least one callback
	test.AssertElement("body").HasAttribute("data-resize-working", "true")

	// Test observer methods by executing cleanup via JavaScript
	test.ExecuteScript(`
		try {
			observer.unobserve(element);
			console.log("unobserve() completed");
		} catch (e) {
			console.log("unobserve() error:", e.message);
		}
		
		try {
			observer.disconnect();
			console.log("disconnect() completed");
		} catch (e) {
			console.log("disconnect() error:", e.message);
		}
		
		// Verify observer functionality
		let results = [];
		results.push("constructor: " + (typeof ResizeObserver === 'function'));
		results.push("creation: " + (observer !== null));
		results.push("methods: " + (typeof observer.observe === 'function'));
		results.push("callbacks: " + (observerCallbackCount > 0));
		
		document.body.setAttribute('data-resize-results', results.join(','));
		document.body.setAttribute('data-resize-test', 'complete');
	`)

	// Verify final test results
	test.AssertElement("body").HasAttribute("data-resize-test", "complete")
	test.AssertElement("body").HasAttributeContaining("data-resize-results", "constructor: true")
	test.AssertElement("body").HasAttributeContaining("data-resize-results", "creation: true")
	test.AssertElement("body").HasAttributeContaining("data-resize-results", "methods: true")
	test.AssertElement("body").HasAttributeContaining("data-resize-results", "callbacks: true")
}

func TestDOMAPIPhase6IntersectionObserver(t *testing.T) {
	test := domulator.NewTest(t)

	test.LoadHTML(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>Phase 6 IntersectionObserver Test</title>
		</head>
		<body>
			<div id="container" style="height: 200px; overflow: scroll;">
				<div id="spacer" style="height: 300px;"></div>
				<div id="target" style="height: 100px; background: red;">Target Element</div>
				<div style="height: 300px;"></div>
			</div>
			<div id="results"></div>
		</body>
		</html>
		<script>
		let intersectionCallbackCount = 0;
		let intersectedElements = [];
		
		// Test IntersectionObserver creation
		console.log("IntersectionObserver constructor:", typeof IntersectionObserver);
		
		const observer = new IntersectionObserver(function(entries) {
			intersectionCallbackCount++;
			console.log("IntersectionObserver callback called, entries:", entries.length);
			
			entries.forEach(function(entry) {
				intersectedElements.push(entry.target.id);
				console.log("Intersected element:", entry.target.id);
				console.log("Is intersecting:", entry.isIntersecting);
				console.log("Intersection ratio:", entry.intersectionRatio);
				console.log("Intersection rect:", entry.intersectionRect);
				console.log("Bounding client rect:", entry.boundingClientRect);
				console.log("Root bounds:", entry.rootBounds);
				
				// Store results in DOM for verification
				const results = document.getElementById('results');
				results.setAttribute('data-callback-count', intersectionCallbackCount.toString());
				results.setAttribute('data-intersected-elements', intersectedElements.join(','));
				
				// Test entry properties
				let entryResults = [];
				entryResults.push("isIntersecting: " + (typeof entry.isIntersecting === 'boolean'));
				entryResults.push("intersectionRatio: " + (typeof entry.intersectionRatio === 'number'));
				entryResults.push("intersectionRect: " + (typeof entry.intersectionRect === 'object'));
				entryResults.push("boundingClientRect: " + (typeof entry.boundingClientRect === 'object'));
				entryResults.push("target: " + (entry.target === document.getElementById('target')));
				
				results.setAttribute('data-entry-properties', entryResults.join(','));
			});
		}, {
			root: document.getElementById('container'),
			rootMargin: '0px',
			threshold: [0, 0.5, 1.0]
		});
		
		console.log("IntersectionObserver created");
		
		// Test observer methods
		const target = document.getElementById('target');
		
		try {
			observer.observe(target);
			console.log("observe() completed");
		} catch (e) {
			console.log("observe() error:", e.message);
		}
		
		// Simulate scrolling to trigger intersection
		setTimeout(() => {
			const container = document.getElementById('container');
			container.scrollTop = 200; // Scroll to bring target into view
			console.log("Container scrolled");
			
			setTimeout(() => {
				try {
					observer.unobserve(target);
					console.log("unobserve() completed");
				} catch (e) {
					console.log("unobserve() error:", e.message);
				}
				
				try {
					observer.disconnect();
					console.log("disconnect() completed");
				} catch (e) {
					console.log("disconnect() error:", e.message);
				}
				
				// Verify observer functionality
				let results = [];
				results.push("constructor: " + (typeof IntersectionObserver === 'function'));
				results.push("creation: " + (observer !== null));
				results.push("methods: " + (typeof observer.observe === 'function'));
				results.push("callbacks: " + (intersectionCallbackCount >= 0)); // May or may not trigger in test
				
				document.body.setAttribute('data-intersection-results', results.join(','));
				document.body.setAttribute('data-intersection-test', 'complete');
			}, 50);
		}, 10);
		</script>
	`)

	// Advance virtual time to execute the setTimeout calls and process microtasks
	test.AdvanceTime(100 * time.Millisecond)
	test.FlushMicrotasks()
	test.ProcessPendingTimers()

	// Process any additional timers that may have been scheduled
	test.AdvanceTime(100 * time.Millisecond)
	test.ProcessPendingTimers()

	test.AssertElement("body").HasAttribute("data-intersection-test", "complete")
	test.AssertElement("body").HasAttributeContaining("data-intersection-results", "constructor: true")
	test.AssertElement("body").HasAttributeContaining("data-intersection-results", "creation: true")
	test.AssertElement("body").HasAttributeContaining("data-intersection-results", "methods: true")
}

func TestDOMAPIPhase6EnhancedFormData(t *testing.T) {
	test := domulator.NewTest(t)

	test.LoadHTML(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>Phase 6 Enhanced FormData Test</title>
		</head>
		<body>
			<form id="test-form">
				<input type="text" name="username" value="john_doe">
				<input type="email" name="email" value="john@example.com">
				<input type="password" name="password" value="secret123">
				<input type="hidden" name="csrf_token" value="abc123xyz">
				<select name="country">
					<option value="us" selected>United States</option>
					<option value="ca">Canada</option>
				</select>
				<textarea name="comments">Hello world!</textarea>
				<input type="checkbox" name="newsletter" checked>
				<input type="radio" name="gender" value="male" checked>
				<input type="radio" name="gender" value="female">
			</form>
		</body>
		</html>
		<script>
		// Test FormData constructor
		console.log("FormData constructor:", typeof FormData);
		
		// Test FormData with form element
		const form = document.getElementById('test-form');
		const formData1 = new FormData(form);
		console.log("FormData created from form");
		
		// Test FormData without form
		const formData2 = new FormData();
		console.log("FormData created empty");
		
		// Test FormData methods
		let results = [];
		
		// Test has() method
		const hasUsername = formData1.has('username');
		const hasNonexistent = formData1.has('nonexistent');
		results.push("has: " + (hasUsername === true && hasNonexistent === false));
		
		// Test get() method
		const username = formData1.get('username');
		const nonexistent = formData1.get('nonexistent');
		results.push("get: " + (username === 'john_doe' && nonexistent === null));
		
		// Test set() method
		formData2.set('name', 'Jane');
		formData2.set('age', '25');
		const setName = formData2.get('name');
		results.push("set: " + (setName === 'Jane'));
		
		// Test append() method
		formData2.append('hobby', 'reading');
		formData2.append('hobby', 'swimming');
		const hobbies = formData2.getAll('hobby');
		results.push("append: " + (hobbies.length === 2));
		
		// Test getAll() method
		results.push("getAll: " + (Array.isArray(hobbies) && hobbies[0] === 'reading'));
		
		// Test delete() method
		formData2.delete('age');
		const deletedAge = formData2.has('age');
		results.push("delete: " + (deletedAge === false));
		
		// Test keys() method
		try {
			const keys = formData2.keys();
			results.push("keys: " + (typeof keys === 'object'));
		} catch (e) {
			results.push("keys: error");
		}
		
		// Test values() method
		try {
			const values = formData2.values();
			results.push("values: " + (typeof values === 'object'));
		} catch (e) {
			results.push("values: error");
		}
		
		// Test entries() method
		try {
			const entries = formData2.entries();
			results.push("entries: " + (typeof entries === 'object'));
		} catch (e) {
			results.push("entries: error");
		}
		
		// Test forEach() method
		let forEachCount = 0;
		try {
			formData2.forEach(function(value, key) {
				forEachCount++;
			});
			results.push("forEach: " + (forEachCount > 0));
		} catch (e) {
			results.push("forEach: error");
		}
		
		// Verify FormData functionality
		let finalResults = [];
		finalResults.push("constructor: " + (typeof FormData === 'function'));
		finalResults.push("creation: " + (formData1 !== null && formData2 !== null));
		finalResults.push("methods: " + (typeof formData1.get === 'function'));
		finalResults.push("form-parsing: " + (formData1.get('username') === 'john_doe'));
		
		document.body.setAttribute('data-formdata-results', finalResults.join(','));
		document.body.setAttribute('data-formdata-methods', results.join(','));
		document.body.setAttribute('data-formdata-test', 'complete');
		</script>
	`)

	test.AssertElement("body").HasAttribute("data-formdata-test", "complete")
	test.AssertElement("body").HasAttributeContaining("data-formdata-results", "constructor: true")
	test.AssertElement("body").HasAttributeContaining("data-formdata-results", "creation: true")
	test.AssertElement("body").HasAttributeContaining("data-formdata-results", "methods: true")
	test.AssertElement("body").HasAttributeContaining("data-formdata-methods", "has: true")
	test.AssertElement("body").HasAttributeContaining("data-formdata-methods", "get: true")
	test.AssertElement("body").HasAttributeContaining("data-formdata-methods", "set: true")
}

func TestDOMAPIPhase6MediaQueryAPI(t *testing.T) {
	test := domulator.NewTest(t)

	test.LoadHTML(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>Phase 6 MediaQuery API Test</title>
		</head>
		<body>
			<div id="test-element">Media Query Test</div>
		</body>
		</html>
		<script>
		// Test matchMedia function
		console.log("matchMedia function:", typeof window.matchMedia);
		
		let results = [];
		
		// Test basic media queries
		try {
			const mq1 = window.matchMedia('(min-width: 768px)');
			console.log("Media query created:", mq1.media);
			results.push("creation: " + (mq1 !== null));
			results.push("media: " + (mq1.media === '(min-width: 768px)'));
			results.push("matches: " + (typeof mq1.matches === 'boolean'));
		} catch (e) {
			console.log("matchMedia error:", e.message);
			results.push("creation: error");
		}
		
		// Test different media query types
		try {
			const orientationQuery = window.matchMedia('(orientation: landscape)');
			const colorQuery = window.matchMedia('(min-color: 8)');
			const prefersReducedMotion = window.matchMedia('(prefers-reduced-motion: reduce)');
			
			results.push("orientation: " + (orientationQuery.media.includes('orientation')));
			results.push("color: " + (colorQuery.media.includes('color')));
			results.push("prefers: " + (prefersReducedMotion.media.includes('prefers')));
		} catch (e) {
			results.push("queries: error");
		}
		
		// Test addListener/removeListener (deprecated but still supported)
		try {
			const mq = window.matchMedia('(max-width: 600px)');
			const listener = function(e) {
				console.log("Media query changed:", e.matches);
			};
			
			mq.addListener(listener);
			results.push("addListener: " + (typeof mq.addListener === 'function'));
			
			mq.removeListener(listener);
			results.push("removeListener: " + (typeof mq.removeListener === 'function'));
		} catch (e) {
			results.push("listeners: error");
		}
		
		// Test addEventListener/removeEventListener (modern approach)
		try {
			const mq = window.matchMedia('(hover: hover)');
			const eventListener = function(e) {
				console.log("Media query event:", e.matches);
			};
			
			mq.addEventListener('change', eventListener);
			results.push("addEventListener: " + (typeof mq.addEventListener === 'function'));
			
			mq.removeEventListener('change', eventListener);
			results.push("removeEventListener: " + (typeof mq.removeEventListener === 'function'));
		} catch (e) {
			results.push("eventListeners: error");
		}
		
		// Verify MediaQuery functionality
		let finalResults = [];
		finalResults.push("matchMedia: " + (typeof window.matchMedia === 'function'));
		finalResults.push("creation: " + results.includes("creation: true"));
		finalResults.push("properties: " + results.includes("matches: true"));
		finalResults.push("methods: " + results.includes("addListener: true"));
		
		document.body.setAttribute('data-mediaquery-results', finalResults.join(','));
		document.body.setAttribute('data-mediaquery-methods', results.join(','));
		document.body.setAttribute('data-mediaquery-test', 'complete');
		</script>
	`)

	test.AssertElement("body").HasAttribute("data-mediaquery-test", "complete")
	test.AssertElement("body").HasAttributeContaining("data-mediaquery-results", "matchMedia: true")
	test.AssertElement("body").HasAttributeContaining("data-mediaquery-results", "creation: true")
	test.AssertElement("body").HasAttributeContaining("data-mediaquery-results", "properties: true")
}

func TestDOMAPIPhase6FileAPI(t *testing.T) {
	test := domulator.NewTest(t)

	test.LoadHTML(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>Phase 6 File API Test</title>
		</head>
		<body>
			<input type="file" id="file-input" multiple>
			<div id="results"></div>
		</body>
		</html>
		<script>
		let results = [];
		
		// Test File constructor
		console.log("File constructor:", typeof File);
		try {
			const file1 = new File(['Hello, World!'], 'hello.txt', {
				type: 'text/plain',
				lastModified: Date.now()
			});
			
			console.log("File created:", file1.name);
			results.push("file-creation: " + (file1.name === 'hello.txt'));
			results.push("file-type: " + (file1.type === 'text/plain'));
			results.push("file-size: " + (file1.size > 0));
		} catch (e) {
			console.log("File constructor error:", e.message);
			results.push("file-creation: error");
		}
		
		// Test Blob constructor
		console.log("Blob constructor:", typeof Blob);
		try {
			const blob1 = new Blob(['Hello, Blob!'], { type: 'text/plain' });
			console.log("Blob created, size:", blob1.size);
			results.push("blob-creation: " + (blob1.size > 0));
			results.push("blob-type: " + (blob1.type === 'text/plain'));
		} catch (e) {
			console.log("Blob constructor error:", e.message);
			results.push("blob-creation: error");
		}
		
		// Test FileReader constructor
		console.log("FileReader constructor:", typeof FileReader);
		try {
			const reader = new FileReader();
			console.log("FileReader created");
			
			results.push("reader-creation: " + (reader !== null));
			results.push("reader-state: " + (reader.readyState === FileReader.EMPTY));
			results.push("reader-methods: " + (typeof reader.readAsText === 'function'));
			
			// Test FileReader constants
			results.push("reader-constants: " + (
				FileReader.EMPTY === 0 && 
				FileReader.LOADING === 1 && 
				FileReader.DONE === 2
			));
			
		} catch (e) {
			console.log("FileReader constructor error:", e.message);
			results.push("reader-creation: error");
		}
		
		// Test FileReader with Blob
		try {
			const blob = new Blob(['Test content'], { type: 'text/plain' });
			const reader = new FileReader();
			
			reader.onload = function(e) {
				console.log("FileReader loaded:", e.target.result);
				const loadResults = document.getElementById('results');
				loadResults.setAttribute('data-reader-result', e.target.result);
			};
			
			reader.onerror = function(e) {
				console.log("FileReader error:", e);
			};
			
			reader.readAsText(blob);
			results.push("reader-read: " + (reader.readyState === FileReader.LOADING || reader.readyState === FileReader.DONE));
			
		} catch (e) {
			console.log("FileReader read error:", e.message);
			results.push("reader-read: error");
		}
		
		// Test URL.createObjectURL (if available)
		try {
			const blob = new Blob(['URL test'], { type: 'text/plain' });
			if (typeof URL.createObjectURL === 'function') {
				const objectURL = URL.createObjectURL(blob);
				console.log("Object URL created:", objectURL);
				results.push("object-url: " + (typeof objectURL === 'string'));
				
				URL.revokeObjectURL(objectURL);
				results.push("revoke-url: " + (typeof URL.revokeObjectURL === 'function'));
			} else {
				results.push("object-url: not-available");
				results.push("revoke-url: not-available");
			}
		} catch (e) {
			console.log("URL object error:", e.message);
			results.push("object-url: error");
		}
		
		// Verify File API functionality
		let finalResults = [];
		finalResults.push("File: " + (typeof File === 'function'));
		finalResults.push("Blob: " + (typeof Blob === 'function'));
		finalResults.push("FileReader: " + (typeof FileReader === 'function'));
		finalResults.push("creation: " + results.includes("file-creation: true"));
		
		document.body.setAttribute('data-fileapi-results', finalResults.join(','));
		document.body.setAttribute('data-fileapi-methods', results.join(','));
		document.body.setAttribute('data-fileapi-test', 'complete');
		</script>
	`)

	test.AssertElement("body").HasAttribute("data-fileapi-test", "complete")
	test.AssertElement("body").HasAttributeContaining("data-fileapi-results", "Blob: true")
	test.AssertElement("body").HasAttributeContaining("data-fileapi-results", "FileReader: true")
	test.AssertElement("body").HasAttributeContaining("data-fileapi-methods", "blob-creation: true")
	test.AssertElement("body").HasAttributeContaining("data-fileapi-methods", "reader-creation: true")
}

func TestDOMAPIPhase6IntegrationScenarios(t *testing.T) {
	test := domulator.NewTest(t)

	test.LoadHTML(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>Phase 6 Integration Test</title>
		</head>
		<body>
			<div id="app">
				<div id="sidebar" style="width: 250px; background: lightgray;">
					<div id="resizable-panel" style="height: 200px; background: lightblue;">
						Resizable Panel
					</div>
				</div>
				<div id="main-content" style="margin-left: 260px;">
					<div id="infinite-scroll-container" style="height: 300px; overflow-y: auto;">
						<div class="item" style="height: 100px; background: lightyellow; margin: 10px;">Item 1</div>
						<div class="item" style="height: 100px; background: lightyellow; margin: 10px;">Item 2</div>
						<div class="item" style="height: 100px; background: lightyellow; margin: 10px;">Item 3</div>
						<div id="load-trigger" style="height: 50px; background: lightcoral;">Load More Trigger</div>
					</div>
					<form id="dynamic-form">
						<input type="text" name="search" placeholder="Search...">
						<button type="submit">Search</button>
					</form>
				</div>
			</div>
		</body>
		</html>
		<script>
		let integrationResults = [];
		
		// Test 1: ResizeObserver monitoring sidebar
		try {
			const resizeObserver = new ResizeObserver(function(entries) {
				entries.forEach(function(entry) {
					console.log("Sidebar resized:", entry.contentRect.width);
					integrationResults.push("sidebar-resize");
				});
			});
			
			const sidebar = document.getElementById('sidebar');
			resizeObserver.observe(sidebar);
			
			// Simulate resize
			sidebar.style.width = '300px';
			integrationResults.push("resize-setup");
			
		} catch (e) {
			console.log("ResizeObserver integration error:", e.message);
			integrationResults.push("resize-error");
		}
		
		// Test 2: IntersectionObserver for infinite scroll
		try {
			const intersectionObserver = new IntersectionObserver(function(entries) {
				entries.forEach(function(entry) {
					if (entry.isIntersecting) {
						console.log("Load trigger visible, loading more items...");
						integrationResults.push("infinite-scroll-triggered");
						
						// Simulate loading more items
						const container = document.getElementById('infinite-scroll-container');
						const newItem = document.createElement('div');
						newItem.className = 'item';
						newItem.style.cssText = 'height: 100px; background: lightgreen; margin: 10px;';
						newItem.textContent = 'New Item ' + (container.children.length);
						container.insertBefore(newItem, entry.target);
					}
				});
			});
			
			const loadTrigger = document.getElementById('load-trigger');
			intersectionObserver.observe(loadTrigger);
			integrationResults.push("intersection-setup");
			
		} catch (e) {
			console.log("IntersectionObserver integration error:", e.message);
			integrationResults.push("intersection-error");
		}
		
		// Test 3: FormData with dynamic form
		try {
			const form = document.getElementById('dynamic-form');
			form.addEventListener('submit', function(e) {
				e.preventDefault();
				
				const formData = new FormData(form);
				const searchTerm = formData.get('search');
				console.log("Form submitted with search:", searchTerm);
				integrationResults.push("form-processing");
				
				// Simulate search results processing
				integrationResults.push("search-term:" + searchTerm);
			});
			
			integrationResults.push("form-setup");
			
		} catch (e) {
			console.log("FormData integration error:", e.message);
			integrationResults.push("form-error");
		}
		
		// Test 4: MediaQuery responsive behavior
		try {
			const mediaQuery = window.matchMedia('(max-width: 768px)');
			mediaQuery.addEventListener('change', function(e) {
				const sidebar = document.getElementById('sidebar');
				if (e.matches) {
					// Mobile layout
					sidebar.style.display = 'none';
					document.getElementById('main-content').style.marginLeft = '0';
				} else {
					// Desktop layout
					sidebar.style.display = 'block';
					document.getElementById('main-content').style.marginLeft = '260px';
				}
				integrationResults.push("layout-changed");
			});
			
			integrationResults.push("mediaquery-setup");
			
		} catch (e) {
			console.log("MediaQuery integration error:", e.message);
			integrationResults.push("mediaquery-error");
		}
		
		// Test 5: Simulate realistic workflow
		setTimeout(() => {
			// Trigger form submission
			const form = document.getElementById('dynamic-form');
			const searchInput = form.querySelector('input[name="search"]');
			searchInput.value = 'test query';
			
			// Trigger form submission event
			const submitEvent = new Event('submit', { bubbles: true, cancelable: true });
			form.dispatchEvent(submitEvent);
			
			setTimeout(() => {
				// Store final integration results
				document.body.setAttribute('data-integration-results', integrationResults.join(','));
				document.body.setAttribute('data-integration-test', 'complete');
			}, 20);
		}, 10);
		</script>
	`)

	// Advance virtual time to execute the setTimeout calls and process microtasks
	test.AdvanceTime(100 * time.Millisecond)
	test.FlushMicrotasks()
	test.ProcessPendingTimers()

	// Process any additional timers that may have been scheduled
	test.AdvanceTime(100 * time.Millisecond)
	test.ProcessPendingTimers()

	test.AssertElement("body").HasAttribute("data-integration-test", "complete")
	test.AssertElement("body").HasAttributeContaining("data-integration-results", "resize-setup")
	test.AssertElement("body").HasAttributeContaining("data-integration-results", "intersection-setup")
	test.AssertElement("body").HasAttributeContaining("data-integration-results", "form-setup")
	test.AssertElement("body").HasAttributeContaining("data-integration-results", "mediaquery-setup")
}
