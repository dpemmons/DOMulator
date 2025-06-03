package js

import (
	"strings"
	"testing"

	"github.com/dpemmons/DOMulator/internal/dom"
	"github.com/dpemmons/DOMulator/internal/parser"
)

func TestJavaScriptRuntimeBasics(t *testing.T) {
	// Create a document
	doc := dom.NewDocument()

	// Create runtime
	runtime := New(doc)
	defer runtime.Shutdown()

	// Test basic JavaScript execution
	result, err := runtime.RunString("1 + 1")
	if err != nil {
		t.Fatalf("Failed to execute JavaScript: %v", err)
	}

	if result.ToInteger() != 2 {
		t.Errorf("Expected 2, got %d", result.ToInteger())
	}
}

func TestConsoleAPI(t *testing.T) {
	doc := dom.NewDocument()
	runtime := New(doc)
	defer runtime.Shutdown()

	// Test console.log
	_, err := runtime.RunString(`
		console.log("Hello, World!");
		console.error("Test error");
		console.warn("Test warning");
		console.info("Test info");
	`)
	if err != nil {
		t.Fatalf("Failed to execute console API: %v", err)
	}
}

func TestTimerAPI(t *testing.T) {
	doc := dom.NewDocument()
	runtime := New(doc)
	defer runtime.Shutdown()

	// Test setTimeout
	result, err := runtime.RunString(`
		var executed = false;
		var timerId = setTimeout(function() {
			executed = true;
		}, 10);
		timerId; // Return the timer ID
	`)
	if err != nil {
		t.Fatalf("Failed to execute setTimeout: %v", err)
	}

	// Verify timer ID was returned
	if result.ToInteger() <= 0 {
		t.Errorf("Expected positive timer ID, got %d", result.ToInteger())
	}
}

func TestDOMDocument(t *testing.T) {
	doc := dom.NewDocument()
	runtime := New(doc)
	defer runtime.Shutdown()

	// Test document global
	_, err := runtime.RunString(`
		if (typeof document === 'undefined') {
			throw new Error('document is not defined');
		}
		
		if (document.nodeType !== 9) { // DOCUMENT_NODE
			throw new Error('document.nodeType should be 9');
		}
		
		if (document.nodeName !== '#document') {
			throw new Error('document.nodeName should be #document');
		}
	`)
	if err != nil {
		t.Fatalf("Document API test failed: %v", err)
	}
}

func TestElementCreation(t *testing.T) {
	doc := dom.NewDocument()
	runtime := New(doc)
	defer runtime.Shutdown()

	// Test createElement
	_, err := runtime.RunString(`
		var div = document.createElement('div');
		
		if (!div) {
			throw new Error('createElement returned null');
		}
		
		if (div.nodeType !== 1) { // ELEMENT_NODE
			throw new Error('element.nodeType should be 1');
		}
		
		if (div.tagName !== 'div') {
			throw new Error('element.tagName should be div');
		}
		
		if (div.nodeName !== 'div') {
			throw new Error('element.nodeName should be div');
		}
	`)
	if err != nil {
		t.Fatalf("createElement test failed: %v", err)
	}
}

func TestTextNodeCreation(t *testing.T) {
	doc := dom.NewDocument()
	runtime := New(doc)
	defer runtime.Shutdown()

	// Test createTextNode
	_, err := runtime.RunString(`
		var textNode = document.createTextNode('Hello, World!');
		
		if (!textNode) {
			throw new Error('createTextNode returned null');
		}
		
		if (textNode.nodeType !== 3) { // TEXT_NODE
			throw new Error('textNode.nodeType should be 3');
		}
		
		if (textNode.nodeName !== '#text') {
			throw new Error('textNode.nodeName should be #text');
		}
		
		if (textNode.nodeValue !== 'Hello, World!') {
			throw new Error('textNode.nodeValue should be "Hello, World!"');
		}
	`)
	if err != nil {
		t.Fatalf("createTextNode test failed: %v", err)
	}
}

func TestDOMManipulation(t *testing.T) {
	doc := dom.NewDocument()
	runtime := New(doc)
	defer runtime.Shutdown()

	// Test DOM manipulation
	_, err := runtime.RunString(`
		var div = document.createElement('div');
		var text = document.createTextNode('Hello');
		
		div.appendChild(text);
		
		if (div.childNodes.length !== 1) {
			throw new Error('div should have 1 child');
		}
		
		if (div.firstChild !== text) {
			throw new Error('div.firstChild should be the text node');
		}
		
		if (text.parentNode !== div) {
			throw new Error('text.parentNode should be the div');
		}
	`)
	if err != nil {
		t.Fatalf("DOM manipulation test failed: %v", err)
	}
}

func TestAttributeManipulation(t *testing.T) {
	doc := dom.NewDocument()
	runtime := New(doc)
	defer runtime.Shutdown()

	// Test attribute manipulation
	_, err := runtime.RunString(`
		var div = document.createElement('div');
		
		// Test setAttribute/getAttribute
		div.setAttribute('id', 'test-id');
		if (div.getAttribute('id') !== 'test-id') {
			throw new Error('getAttribute should return set value');
		}
		
		// Test hasAttribute
		if (!div.hasAttribute('id')) {
			throw new Error('hasAttribute should return true for set attribute');
		}
		
		if (div.hasAttribute('class')) {
			throw new Error('hasAttribute should return false for unset attribute');
		}
		
		// Test removeAttribute
		div.removeAttribute('id');
		if (div.hasAttribute('id')) {
			throw new Error('hasAttribute should return false after removeAttribute');
		}
	`)
	if err != nil {
		t.Fatalf("Attribute manipulation test failed: %v", err)
	}
}

func TestDOMWithParser(t *testing.T) {
	// Parse HTML and test JavaScript interaction
	html := `<html><head><title>Test</title></head><body><div id="main">Hello</div></body></html>`

	p := parser.NewParser()
	doc, err := p.Parse(strings.NewReader(html))
	if err != nil {
		t.Fatalf("Failed to parse HTML: %v", err)
	}

	runtime := New(doc)
	defer runtime.Shutdown()

	// Test getElementById
	_, err = runtime.RunString(`
		var mainDiv = document.getElementById('main');
		if (!mainDiv) {
			throw new Error('getElementById should find the main div');
		}
		
		if (mainDiv.tagName !== 'div') {
			throw new Error('Found element should be a div');
		}
		
		if (mainDiv.getAttribute('id') !== 'main') {
			throw new Error('Found element should have id="main"');
		}
	`)
	if err != nil {
		t.Fatalf("getElementById test failed: %v", err)
	}
}

func TestNodeConstants(t *testing.T) {
	doc := dom.NewDocument()
	runtime := New(doc)
	defer runtime.Shutdown()

	// Test Node constants
	_, err := runtime.RunString(`
		if (Node.ELEMENT_NODE !== 1) {
			throw new Error('Node.ELEMENT_NODE should be 1');
		}
		
		if (Node.TEXT_NODE !== 3) {
			throw new Error('Node.TEXT_NODE should be 3');
		}
		
		if (Node.COMMENT_NODE !== 8) {
			throw new Error('Node.COMMENT_NODE should be 8');
		}
		
		if (Node.DOCUMENT_NODE !== 9) {
			throw new Error('Node.DOCUMENT_NODE should be 9');
		}
	`)
	if err != nil {
		t.Fatalf("Node constants test failed: %v", err)
	}
}

func TestStorageAPIs(t *testing.T) {
	doc := dom.NewDocument()
	runtime := New(doc)
	defer runtime.Shutdown()

	// Test localStorage exists
	_, err := runtime.RunString(`
		if (typeof localStorage === 'undefined') {
			throw new Error('localStorage is not defined');
		}
		
		if (typeof sessionStorage === 'undefined') {
			throw new Error('sessionStorage is not defined');
		}
	`)
	if err != nil {
		t.Fatalf("Storage APIs existence test failed: %v", err)
	}
}

func TestLocalStorageOperations(t *testing.T) {
	doc := dom.NewDocument()
	runtime := New(doc)
	defer runtime.Shutdown()

	// Test localStorage operations
	_, err := runtime.RunString(`
		// Test setItem and getItem
		localStorage.setItem('key1', 'value1');
		if (localStorage.getItem('key1') !== 'value1') {
			throw new Error('localStorage.getItem should return set value');
		}
		
		// Test length property
		if (localStorage.length !== 1) {
			throw new Error('localStorage.length should be 1 after setting one item');
		}
		
		// Test key method
		if (localStorage.key(0) !== 'key1') {
			throw new Error('localStorage.key(0) should return "key1"');
		}
		
		// Test removeItem
		localStorage.removeItem('key1');
		if (localStorage.getItem('key1') !== null) {
			throw new Error('localStorage.getItem should return null after removeItem');
		}
		
		if (localStorage.length !== 0) {
			throw new Error('localStorage.length should be 0 after removing item');
		}
	`)
	if err != nil {
		t.Fatalf("localStorage operations test failed: %v", err)
	}
}

func TestSessionStorageOperations(t *testing.T) {
	doc := dom.NewDocument()
	runtime := New(doc)
	defer runtime.Shutdown()

	// Test sessionStorage operations
	_, err := runtime.RunString(`
		// Test setItem and getItem
		sessionStorage.setItem('session-key', 'session-value');
		if (sessionStorage.getItem('session-key') !== 'session-value') {
			throw new Error('sessionStorage.getItem should return set value');
		}
		
		// Test that sessionStorage and localStorage are separate
		if (localStorage.getItem('session-key') !== null) {
			throw new Error('localStorage should not contain sessionStorage items');
		}
		
		// Test clear
		sessionStorage.clear();
		if (sessionStorage.length !== 0) {
			throw new Error('sessionStorage.length should be 0 after clear');
		}
		
		if (sessionStorage.getItem('session-key') !== null) {
			throw new Error('sessionStorage should be empty after clear');
		}
	`)
	if err != nil {
		t.Fatalf("sessionStorage operations test failed: %v", err)
	}
}

func TestStorageErrorHandling(t *testing.T) {
	doc := dom.NewDocument()
	runtime := New(doc)
	defer runtime.Shutdown()

	// Test error handling for insufficient arguments
	_, err := runtime.RunString(`
		var errorCaught = false;
		try {
			localStorage.setItem('key');
		} catch(e) {
			errorCaught = true;
		}
		
		if (!errorCaught) {
			throw new Error('Expected error for insufficient arguments to setItem');
		}
	`)
	if err != nil {
		t.Fatalf("Storage error handling test failed: %v", err)
	}
}

func TestStorageWindowReference(t *testing.T) {
	doc := dom.NewDocument()
	runtime := New(doc)
	defer runtime.Shutdown()

	// Test that storage is available on window object
	_, err := runtime.RunString(`
		if (window.localStorage !== localStorage) {
			throw new Error('window.localStorage should reference same object as localStorage');
		}
		
		if (window.sessionStorage !== sessionStorage) {
			throw new Error('window.sessionStorage should reference same object as sessionStorage');
		}
		
		// Test operations through window references
		window.localStorage.setItem('window-test', 'window-value');
		if (localStorage.getItem('window-test') !== 'window-value') {
			throw new Error('localStorage and window.localStorage should be the same object');
		}
	`)
	if err != nil {
		t.Fatalf("Storage window reference test failed: %v", err)
	}
}
