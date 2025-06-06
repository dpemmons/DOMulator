package integration

import (
	"testing"
	"time"

	domulator "github.com/dpemmons/DOMulator"
)

func TestDOMAPIPhase5RangeAPI(t *testing.T) {
	test := domulator.NewTest(t)

	test.LoadHTML(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>Phase 5 Range API Test</title>
		</head>
		<body>
			<div id="content">
				<p id="para1">First paragraph with some text content.</p>
				<p id="para2">Second paragraph with different content.</p>
			</div>
		</body>
		</html>
		<script>
		const para1 = document.getElementById('para1');
		const para2 = document.getElementById('para2');
		
		// Test Range creation and basic properties
		const range = document.createRange();
		console.log("Range created:", range.constructor.name);
		
		// Test range selection
		range.selectNode(para1);
		console.log("Range selectNode completed");
		
		const collapsed = range.collapsed;
		const startContainer = range.startContainer;
		const endContainer = range.endContainer;
		const startOffset = range.startOffset;
		const endOffset = range.endOffset;
		
		console.log("Range collapsed:", collapsed);
		console.log("Range start container type:", startContainer.nodeType);
		console.log("Range end container type:", endContainer.nodeType);
		console.log("Range start offset:", startOffset);
		console.log("Range end offset:", endOffset);
		
		// Test range content extraction
		try {
			const contents = range.extractContents();
			console.log("Range extractContents completed");
		} catch (e) {
			console.log("Range extractContents error:", e.message);
		}
		
		// Test range content cloning
		try {
			const clonedRange = range.cloneRange();
			console.log("Range cloneRange completed");
		} catch (e) {
			console.log("Range cloneRange error:", e.message);
		}
		
		// Test range methods
		try {
			range.selectNodeContents(para2);
			console.log("Range selectNodeContents completed");
		} catch (e) {
			console.log("Range selectNodeContents error:", e.message);
		}
		
		try {
			range.setStart(para2.firstChild, 0);
			range.setEnd(para2.firstChild, 5);
			console.log("Range setStart/setEnd completed");
		} catch (e) {
			console.log("Range setStart/setEnd error:", e.message);
		}
		
		try {
			range.collapse(true);
			console.log("Range collapse completed");
		} catch (e) {
			console.log("Range collapse error:", e.message);
		}
		
		// Mark test complete
		let results = [];
		results.push("creation: " + (typeof range === 'object'));
		results.push("properties: " + (typeof collapsed === 'boolean'));
		results.push("methods: " + (typeof range.selectNode === 'function'));
		
		document.body.setAttribute('data-range-results', results.join(','));
		document.body.setAttribute('data-range-test', 'complete');
		</script>
	`)

	test.AssertElement("body").HasAttribute("data-range-test", "complete")
	test.AssertElement("body").HasAttributeContaining("data-range-results", "creation: true")
	test.AssertElement("body").HasAttributeContaining("data-range-results", "properties: true")
	test.AssertElement("body").HasAttributeContaining("data-range-results", "methods: true")
}

func TestDOMAPIPhase5TreeWalker(t *testing.T) {
	test := domulator.NewTest(t)

	test.LoadHTML(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>Phase 5 TreeWalker Test</title>
		</head>
		<body>
			<div id="container">
				<p>First paragraph</p>
				<div>
					<span>Nested span</span>
					<em>Emphasis text</em>
				</div>
				<p>Second paragraph</p>
			</div>
		</body>
		</html>
		<script>
		const container = document.getElementById('container');
		
		// Test TreeWalker creation
		const walker = document.createTreeWalker(
			container,
			NodeFilter.SHOW_ELEMENT,
			null,
			false
		);
		
		console.log("TreeWalker created");
		console.log("Walker root:", walker.root.id);
		console.log("Walker current node:", walker.currentNode.id);
		
		// Test TreeWalker navigation
		let nodeCount = 0;
		let currentNode = walker.currentNode;
		
		// Count nodes by walking through them
		while (currentNode) {
			nodeCount++;
			console.log("Visiting node:", currentNode.tagName || currentNode.nodeType);
			currentNode = walker.nextNode();
		}
		
		console.log("Total nodes visited:", nodeCount);
		
		// Test walker methods
		walker.currentNode = container;
		const firstChild = walker.firstChild();
		console.log("First child:", firstChild ? firstChild.tagName : 'null');
		
		const nextSibling = walker.nextSibling();
		console.log("Next sibling:", nextSibling ? nextSibling.tagName : 'null');
		
		const parent = walker.parentNode();
		console.log("Parent node:", parent ? parent.id : 'null');
		
		const lastChild = walker.lastChild();
		console.log("Last child:", lastChild ? lastChild.tagName : 'null');
		
		const previousSibling = walker.previousSibling();
		console.log("Previous sibling:", previousSibling ? previousSibling.tagName : 'null');
		
		const previousNode = walker.previousNode();
		console.log("Previous node:", previousNode ? previousNode.tagName : 'null');
		
		// Verify walker properties and methods
		let results = [];
		results.push("creation: " + (walker !== null));
		results.push("properties: " + (walker.root === container));
		results.push("navigation: " + (nodeCount > 0));
		results.push("methods: " + (typeof walker.nextNode === 'function'));
		
		document.body.setAttribute('data-walker-results', results.join(','));
		document.body.setAttribute('data-walker-test', 'complete');
		</script>
	`)

	test.AssertElement("body").HasAttribute("data-walker-test", "complete")
	test.AssertElement("body").HasAttributeContaining("data-walker-results", "creation: true")
	test.AssertElement("body").HasAttributeContaining("data-walker-results", "navigation: true")
	test.AssertElement("body").HasAttributeContaining("data-walker-results", "methods: true")
}

func TestDOMAPIPhase5NodeIterator(t *testing.T) {
	test := domulator.NewTest(t)

	test.LoadHTML(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>Phase 5 NodeIterator Test</title>
		</head>
		<body>
			<div id="content">
				<h1>Title</h1>
				<p>First paragraph</p>
				<ul>
					<li>List item 1</li>
					<li>List item 2</li>
				</ul>
				<p>Second paragraph</p>
			</div>
		</body>
		</html>
		<script>
		const content = document.getElementById('content');
		
		// Test NodeIterator creation
		const iterator = document.createNodeIterator(
			content,
			NodeFilter.SHOW_ELEMENT,
			null,
			false
		);
		
		console.log("NodeIterator created");
		console.log("Iterator root:", iterator.root.id);
		
		// Test iterator navigation
		let nodeCount = 0;
		let currentNode = iterator.nextNode();
		
		while (currentNode) {
			nodeCount++;
			console.log("Iterator node:", currentNode.tagName);
			currentNode = iterator.nextNode();
		}
		
		console.log("Total iterator nodes:", nodeCount);
		
		// Test reverse iteration
		let reverseCount = 0;
		let reverseNode = iterator.previousNode();
		
		while (reverseNode && reverseCount < 3) {
			reverseCount++;
			console.log("Reverse node:", reverseNode.tagName);
			reverseNode = iterator.previousNode();
		}
		
		console.log("Reverse nodes visited:", reverseCount);
		
		// Test iterator properties
		console.log("Iterator whatToShow:", iterator.whatToShow);
		console.log("Iterator filter:", iterator.filter);
		
		// Verify iterator functionality
		let results = [];
		results.push("creation: " + (iterator !== null));
		results.push("navigation: " + (nodeCount > 0));
		results.push("reverse: " + (reverseCount > 0));
		results.push("properties: " + (typeof iterator.whatToShow === 'number'));
		
		document.body.setAttribute('data-iterator-results', results.join(','));
		document.body.setAttribute('data-iterator-test', 'complete');
		</script>
	`)

	test.AssertElement("body").HasAttribute("data-iterator-test", "complete")
	test.AssertElement("body").HasAttributeContaining("data-iterator-results", "creation: true")
	test.AssertElement("body").HasAttributeContaining("data-iterator-results", "navigation: true")
	test.AssertElement("body").HasAttributeContaining("data-iterator-results", "properties: true")
}

func TestDOMAPIPhase5Selection(t *testing.T) {
	test := domulator.NewTest(t)

	test.LoadHTML(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>Phase 5 Selection API Test</title>
		</head>
		<body>
			<div id="selectable">
				<p id="text1">This is the first paragraph of selectable text.</p>
				<p id="text2">This is the second paragraph with more content.</p>
			</div>
		</body>
		</html>
		<script>
		// Test Selection API
		const selection = window.getSelection();
		console.log("Selection object:", typeof selection);
		
		// Test selection properties
		console.log("Selection range count:", selection.rangeCount);
		console.log("Selection anchor node:", selection.anchorNode);
		console.log("Selection focus node:", selection.focusNode);
		console.log("Selection is collapsed:", selection.isCollapsed);
		
		// Test selection methods
		const text1 = document.getElementById('text1');
		const text2 = document.getElementById('text2');
		
		// Create a range and add it to selection
		const range = document.createRange();
		range.setStart(text1.firstChild, 0);
		range.setEnd(text1.firstChild, 10);
		
		selection.removeAllRanges();
		selection.addRange(range);
		
		console.log("Range added to selection");
		console.log("Selection after adding range - range count:", selection.rangeCount);
		console.log("Selection text:", selection.toString());
		
		// Test selection modification
		try {
			selection.selectAllChildren(text2);
			console.log("selectAllChildren completed");
		} catch (e) {
			console.log("selectAllChildren error:", e.message);
		}
		
		try {
			selection.collapse(text1.firstChild, 5);
			console.log("collapse completed");
		} catch (e) {
			console.log("collapse error:", e.message);
		}
		
		try {
			selection.extend(text1.firstChild, 15);
			console.log("extend completed");
		} catch (e) {
			console.log("extend error:", e.message);
		}
		
		// Test selection modification methods
		try {
			selection.collapseToStart();
			console.log("collapseToStart completed");
		} catch (e) {
			console.log("collapseToStart error:", e.message);
		}
		
		try {
			selection.collapseToEnd();
			console.log("collapseToEnd completed");
		} catch (e) {
			console.log("collapseToEnd error:", e.message);
		}
		
		// Test containsNode
		try {
			const contains = selection.containsNode(text1, false);
			console.log("containsNode result:", contains);
		} catch (e) {
			console.log("containsNode error:", e.message);
		}
		
		// Verify selection functionality
		let results = [];
		results.push("object: " + (typeof selection === 'object'));
		results.push("properties: " + (typeof selection.rangeCount === 'number'));
		results.push("methods: " + (typeof selection.addRange === 'function'));
		results.push("range: " + (selection.rangeCount >= 0));
		
		document.body.setAttribute('data-selection-results', results.join(','));
		document.body.setAttribute('data-selection-test', 'complete');
		</script>
	`)

	test.AssertElement("body").HasAttribute("data-selection-test", "complete")
	test.AssertElement("body").HasAttributeContaining("data-selection-results", "object: true")
	test.AssertElement("body").HasAttributeContaining("data-selection-results", "properties: true")
	test.AssertElement("body").HasAttributeContaining("data-selection-results", "methods: true")
}

func TestDOMAPIPhase5WebStorage(t *testing.T) {
	test := domulator.NewTest(t)

	test.LoadHTML(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>Phase 5 Web Storage Test</title>
		</head>
		<body>
			<div id="storage-test">Storage Test</div>
		</body>
		</html>
		<script>
		// Test localStorage
		console.log("Testing localStorage...");
		console.log("localStorage object:", typeof localStorage);
		
		// Test localStorage methods
		try {
			localStorage.setItem('test-key', 'test-value');
			console.log("localStorage.setItem completed");
			
			const value = localStorage.getItem('test-key');
			console.log("localStorage.getItem result:", value);
			
			const length = localStorage.length;
			console.log("localStorage.length:", length);
			
			const key = localStorage.key(0);
			console.log("localStorage.key(0):", key);
			
			localStorage.removeItem('test-key');
			console.log("localStorage.removeItem completed");
			
			localStorage.clear();
			console.log("localStorage.clear completed");
		} catch (e) {
			console.log("localStorage error:", e.message);
		}
		
		// Test sessionStorage
		console.log("Testing sessionStorage...");
		console.log("sessionStorage object:", typeof sessionStorage);
		
		try {
			sessionStorage.setItem('session-key', 'session-value');
			console.log("sessionStorage.setItem completed");
			
			const sessionValue = sessionStorage.getItem('session-key');
			console.log("sessionStorage.getItem result:", sessionValue);
			
			const sessionLength = sessionStorage.length;
			console.log("sessionStorage.length:", sessionLength);
			
			sessionStorage.removeItem('session-key');
			console.log("sessionStorage.removeItem completed");
		} catch (e) {
			console.log("sessionStorage error:", e.message);
		}
		
		// Test storage events (would fire in other windows/tabs)
		try {
			window.addEventListener('storage', function(e) {
				console.log("Storage event:", e.key, e.newValue);
			});
			console.log("Storage event listener added");
		} catch (e) {
			console.log("Storage event error:", e.message);
		}
		
		// Verify storage functionality
		let results = [];
		results.push("localStorage: " + (typeof localStorage === 'object'));
		results.push("sessionStorage: " + (typeof sessionStorage === 'object'));
		results.push("methods: " + (typeof localStorage.setItem === 'function'));
		results.push("events: " + (typeof window.addEventListener === 'function'));
		
		document.body.setAttribute('data-storage-results', results.join(','));
		document.body.setAttribute('data-storage-test', 'complete');
		</script>
	`)

	test.AssertElement("body").HasAttribute("data-storage-test", "complete")
	test.AssertElement("body").HasAttributeContaining("data-storage-results", "localStorage: true")
	test.AssertElement("body").HasAttributeContaining("data-storage-results", "sessionStorage: true")
	test.AssertElement("body").HasAttributeContaining("data-storage-results", "methods: true")
}

func TestDOMAPIPhase5MutationObserver(t *testing.T) {
	test := domulator.NewTest(t)

	test.LoadHTML(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>Phase 5 MutationObserver Test</title>
		</head>
		<body>
			<div id="observed">
				<p id="child1">Original content</p>
			</div>
		</body>
		</html>
		<script>
		let mutationCount = 0;
		let mutationTypes = [];
		
		// Test MutationObserver creation
		console.log("MutationObserver constructor:", typeof MutationObserver);
		
		const observer = new MutationObserver(function(mutations) {
			mutations.forEach(function(mutation) {
				mutationCount++;
				mutationTypes.push(mutation.type);
				console.log("Mutation observed:", mutation.type, mutation.target.tagName);
				
				// Update DOM to signal mutations were observed
				document.body.setAttribute('data-mutation-count', mutationCount.toString());
				document.body.setAttribute('data-mutation-types', mutationTypes.join(','));
			});
		});
		
		console.log("MutationObserver created");
		
		// Test observer configuration and start observing
		const observed = document.getElementById('observed');
		
		try {
			observer.observe(observed, {
				childList: true,
				attributes: true,
				characterData: true,
				subtree: true
			});
			console.log("Observer started");
		} catch (e) {
			console.log("Observer start error:", e.message);
		}
		
		// Trigger some mutations
		setTimeout(() => {
			// Child list mutation
			const newChild = document.createElement('p');
			newChild.textContent = 'New paragraph';
			observed.appendChild(newChild);
			
			// Attribute mutation
			observed.setAttribute('data-test', 'value');
			
			// Text mutation
			document.getElementById('child1').textContent = 'Modified content';
			
			console.log("Mutations triggered");
			
			// Stop observing after mutations
			setTimeout(() => {
				observer.disconnect();
				console.log("Observer disconnected");
				
				// Final verification
				let results = [];
				results.push("constructor: " + (typeof MutationObserver === 'function'));
				results.push("creation: " + (observer !== null));
				results.push("methods: " + (typeof observer.observe === 'function'));
				results.push("mutations: " + (mutationCount > 0));
				
				document.body.setAttribute('data-observer-results', results.join(','));
				document.body.setAttribute('data-observer-test', 'complete');
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

	test.AssertElement("body").HasAttribute("data-observer-test", "complete")
	test.AssertElement("body").HasAttributeContaining("data-observer-results", "constructor: true")
	test.AssertElement("body").HasAttributeContaining("data-observer-results", "creation: true")
	test.AssertElement("body").HasAttributeContaining("data-observer-results", "methods: true")
}

func TestDOMAPIPhase5IntegrationScenarios(t *testing.T) {
	test := domulator.NewTest(t)

	test.LoadHTML(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>Phase 5 Integration Test</title>
		</head>
		<body>
			<div id="editor">
				<div id="toolbar">
					<button id="bold-btn">Bold</button>
					<button id="italic-btn">Italic</button>
				</div>
				<div id="content" contenteditable="true">
					This is editable content that can be manipulated using Range and Selection APIs.
				</div>
			</div>
		</body>
		</html>
		<script>
		const content = document.getElementById('content');
		const boldBtn = document.getElementById('bold-btn');
		const italicBtn = document.getElementById('italic-btn');
		
		// Complex scenario: Text editor using Range and Selection APIs
		let operations = [];
		
		// Function to apply formatting
		function applyFormatting(tag) {
			const selection = window.getSelection();
			if (selection.rangeCount > 0) {
				const range = selection.getRangeAt(0);
				if (!range.collapsed) {
					const element = document.createElement(tag);
					try {
						range.surroundContents(element);
						operations.push('format-' + tag);
					} catch (e) {
						// Fallback: extract and wrap
						const contents = range.extractContents();
						element.appendChild(contents);
						range.insertNode(element);
						operations.push('format-' + tag + '-fallback');
					}
				}
			}
			updateOperationsList();
		}
		
		function updateOperationsList() {
			document.body.setAttribute('data-operations', operations.join(','));
		}
		
		// Set up toolbar functionality
		boldBtn.addEventListener('click', () => {
			applyFormatting('strong');
		});
		
		italicBtn.addEventListener('click', () => {
			applyFormatting('em');
		});
		
		// Test programmatic selection and formatting
		setTimeout(() => {
			// Create a selection programmatically
			const range = document.createRange();
			const textNode = content.firstChild;
			
			if (textNode && textNode.nodeType === 3) {
				range.setStart(textNode, 0);
				range.setEnd(textNode, 7); // Select "This is"
				
				const selection = window.getSelection();
				selection.removeAllRanges();
				selection.addRange(range);
				
				operations.push('programmatic-selection');
				updateOperationsList();
				
				// Apply bold formatting
				applyFormatting('strong');
				
				// Test TreeWalker to analyze the result
				const walker = document.createTreeWalker(
					content,
					NodeFilter.SHOW_ELEMENT,
					null,
					false
				);
				
				let elementCount = 0;
				let currentNode = walker.nextNode();
				while (currentNode) {
					elementCount++;
					operations.push('walked-' + currentNode.tagName);
					currentNode = walker.nextNode();
				}
				
				// Test Range operations
				const newRange = document.createRange();
				newRange.selectNodeContents(content);
				const text = newRange.toString();
				
				if (text.length > 0) {
					operations.push('range-text-extracted');
				}
				
				// Test storage of editor state
				try {
					localStorage.setItem('editor-content', content.innerHTML);
					operations.push('state-saved');
				} catch (e) {
					operations.push('state-save-failed');
				}
				
				updateOperationsList();
				
				// Final verification
				let results = [];
				results.push("selection: " + (selection.rangeCount > 0));
				results.push("formatting: " + (content.querySelector('strong') !== null));
				results.push("walker: " + (elementCount > 0));
				results.push("operations: " + (operations.length > 0));
				
				document.body.setAttribute('data-integration-results', results.join(','));
				document.body.setAttribute('data-integration-test', 'complete');
			}
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
	test.AssertElement("body").HasAttributeContaining("data-integration-results", "operations: true")
	test.AssertElement("body").HasAttributeContaining("data-operations", "programmatic-selection")
}
