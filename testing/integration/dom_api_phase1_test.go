package integration

import (
	"testing"

	domulator "github.com/dpemmons/DOMulator"
)

func TestDOMAPIPhase1CoreConstructorsAndProperties(t *testing.T) {
	test := domulator.NewTest(t)

	// Load HTML with basic structure
	test.LoadHTML(`
		<!DOCTYPE html>
		<html>
		<head><title>Phase 1 Test</title></head>
		<body>
			<div id="container">
				<p id="test-p">Original text</p>
			</div>
		</body>
		</html>
	`)

	// Test Phase 1: Core Node Constructors & Properties - inline script
	test.LoadHTML(`
		<!DOCTYPE html>
		<html>
		<head><title>Phase 1 Test</title></head>
		<body>
			<div id="container">
				<p id="test-p">Original text</p>
			</div>
		</body>
		</html>
		<script>
		// Test Text constructor
		const textNode = new Text("Hello World");
		console.log("Text constructor works:", textNode.nodeType === 3);
		console.log("Text data:", textNode.nodeValue);
		
		// Test Comment constructor  
		const commentNode = new Comment("This is a comment");
		console.log("Comment constructor works:", commentNode.nodeType === 8);
		console.log("Comment data:", commentNode.nodeValue);
		
		// Test DocumentType constructor (should throw)
		try {
			new DocumentType();
			console.log("DocumentType constructor should throw");
		} catch (e) {
			console.log("DocumentType correctly throws:", e.message);
		}
		
		// Test Attr constructor (should throw)
		try {
			new Attr();
			console.log("Attr constructor should throw"); 
		} catch (e) {
			console.log("Attr correctly throws:", e.message);
		}
		
		// Test basic properties on existing element
		const testP = document.getElementById('test-p');
		
		// Test isConnected property
		console.log("Element isConnected:", testP.isConnected);
		
		// Test baseURI property
		console.log("Element baseURI:", testP.baseURI);
		
		// Test nodeValue getter/setter
		console.log("Element nodeValue (should be null):", testP.nodeValue);
		testP.nodeValue = "new value";
		console.log("Element nodeValue after setter:", testP.nodeValue);
		
		// Test properties on text node
		const textChild = testP.firstChild;
		console.log("Text node isConnected:", textChild.isConnected);
		console.log("Text node nodeValue:", textChild.nodeValue);
		textChild.nodeValue = "Modified text";
		console.log("Text node after modification:", textChild.nodeValue);
		
		// Test textContent setter
		testP.textContent = "Brand new content";
		console.log("After textContent setter:", testP.textContent);
		
		// Append created nodes to test connectivity
		const container = document.getElementById('container');
		container.appendChild(textNode);
		container.appendChild(commentNode);
		
		console.log("Text node connected after append:", textNode.isConnected);
		console.log("Comment node connected after append:", commentNode.isConnected);
		
		// Mark test completion
		document.body.setAttribute('data-phase1-test', 'complete');
		document.getElementById('test-p').setAttribute('data-final-text', document.getElementById('test-p').textContent);
		</script>
	`)

	// Verify test completion
	test.AssertElement("body").HasAttribute("data-phase1-test", "complete")

	// Verify DOM modifications were successful
	test.AssertElement("#test-p").HasAttribute("data-final-text", "Brand new content")

	// Verify elements still exist
	test.AssertElement("#container").Exists()
	test.AssertElement("#test-p").Exists()
}

func TestDOMAPIPhase1IsConnectedBehavior(t *testing.T) {
	test := domulator.NewTest(t)

	test.LoadHTML(`
		<!DOCTYPE html>
		<html>
		<body>
			<div id="parent"></div>
		</body>
		</html>
		<script>
		// Test isConnected behavior for disconnected vs connected nodes
		const disconnectedText = new Text("disconnected");
		const disconnectedComment = new Comment("disconnected");
		
		console.log("Disconnected text isConnected:", disconnectedText.isConnected);
		console.log("Disconnected comment isConnected:", disconnectedComment.isConnected);
		
		const parent = document.getElementById('parent');
		console.log("Parent isConnected:", parent.isConnected);
		
		// Connect the nodes
		parent.appendChild(disconnectedText);
		parent.appendChild(disconnectedComment);
		
		console.log("Text after connection:", disconnectedText.isConnected);
		console.log("Comment after connection:", disconnectedComment.isConnected);
		
		// Remove and test disconnection
		parent.removeChild(disconnectedText);
		console.log("Text after removal:", disconnectedText.isConnected);
		
		document.body.setAttribute('data-connectivity-test', 'complete');
		</script>
	`)

	test.AssertElement("body").HasAttribute("data-connectivity-test", "complete")
}

func TestDOMAPIPhase1BaseURIProperty(t *testing.T) {
	test := domulator.NewTest(t)

	test.LoadHTML(`
		<!DOCTYPE html>
		<html>
		<body>
			<div id="test"></div>
		</body>
		</html>
		<script>
		// Test baseURI property on different node types
		const div = document.getElementById('test');
		const text = new Text("test");
		const comment = new Comment("test");
		
		console.log("Div baseURI:", div.baseURI);
		console.log("Text baseURI:", text.baseURI);  
		console.log("Comment baseURI:", comment.baseURI);
		
		// baseURI should be read-only
		try {
			div.baseURI = "http://example.com";
			console.log("baseURI should be read-only");
		} catch (e) {
			console.log("baseURI correctly read-only");
		}
		
		document.body.setAttribute('data-baseuri-test', 'complete');
		</script>
	`)

	test.AssertElement("body").HasAttribute("data-baseuri-test", "complete")
}
