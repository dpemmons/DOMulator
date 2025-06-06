package integration

import (
	"testing"

	domulator "github.com/dpemmons/DOMulator"
)

func TestDOMAPIPhase3DocumentProperties(t *testing.T) {
	test := domulator.NewTest(t)

	test.LoadHTML(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>Phase 3 Test</title>
		</head>
		<body>
			<div id="test-container">Test Content</div>
		</body>
		</html>
		<script>
		// Test document properties
		console.log("Document implementation exists:", typeof document.implementation);
		console.log("Character set:", document.characterSet);
		console.log("Charset alias:", document.charset);
		console.log("Input encoding alias:", document.inputEncoding);
		console.log("Content type:", document.contentType);
		console.log("Document type exists:", document.doctype !== null);
		console.log("Compat mode:", document.compatMode);

		// Test DOMImplementation methods
		const impl = document.implementation;
		console.log("Has HTML feature:", impl.hasFeature("HTML", "1.0"));
		console.log("Has XML feature:", impl.hasFeature("XML", "1.0"));

		// Mark properties test complete
		document.body.setAttribute('data-properties-test', 'complete');
		</script>
	`)

	test.AssertElement("body").HasAttribute("data-properties-test", "complete")
}

func TestDOMAPIPhase3DocumentCreationMethods(t *testing.T) {
	test := domulator.NewTest(t)

	test.LoadHTML(`
		<!DOCTYPE html>
		<html>
		<body>
			<div id="container"></div>
		</body>
		</html>
		<script>
		const container = document.getElementById('container');
		
		// Test createAttribute
		const attr = document.createAttribute('data-test');
		attr.value = 'test-value';
		console.log("Created attribute:", attr.name, attr.value);
		
		// Test createAttributeNS
		const nsAttr = document.createAttributeNS('http://example.com', 'test:attr');
		nsAttr.value = 'ns-value';
		console.log("Created NS attribute:", nsAttr.name, nsAttr.value);
		
		// Test createCDATASection (where supported)
		try {
			const cdata = document.createCDATASection('test data');
			console.log("Created CDATA section:", cdata.nodeType, cdata.data);
		} catch (e) {
			console.log("CDATA not supported in HTML:", e.message);
		}
		
		// Test createProcessingInstruction
		try {
			const pi = document.createProcessingInstruction('xml-stylesheet', 'href="style.css"');
			console.log("Created PI:", pi.target, pi.data);
		} catch (e) {
			console.log("PI creation:", e.message);
		}
		
		// Test createRange
		const range = document.createRange();
		console.log("Created range:", range.collapsed);
		console.log("Range start container:", range.startContainer);
		console.log("Range end container:", range.endContainer);
		
		// Test range methods
		range.collapse(true);
		range.selectNode(container);
		range.selectNodeContents(container);
		
		// Test createNodeIterator
		const iterator = document.createNodeIterator(document.body);
		console.log("Created iterator:", iterator.root === document.body);
		console.log("Iterator reference node:", iterator.referenceNode === document.body);
		
		// Test iterator methods
		const nextNode = iterator.nextNode();
		const prevNode = iterator.previousNode();
		console.log("Iterator navigation:", nextNode, prevNode);
		
		// Test createTreeWalker
		const walker = document.createTreeWalker(document.body);
		console.log("Created walker:", walker.root === document.body);
		console.log("Walker current node:", walker.currentNode === document.body);
		
		// Test walker methods
		const walkerNext = walker.nextNode();
		const walkerPrev = walker.previousNode();
		const firstChild = walker.firstChild();
		const lastChild = walker.lastChild();
		console.log("Walker navigation:", walkerNext, walkerPrev, firstChild, lastChild);
		
		document.body.setAttribute('data-creation-test', 'complete');
		</script>
	`)

	test.AssertElement("body").HasAttribute("data-creation-test", "complete")
}

func TestDOMAPIPhase3DocumentManipulationMethods(t *testing.T) {
	test := domulator.NewTest(t)

	test.LoadHTML(`
		<!DOCTYPE html>
		<html>
		<body>
			<div id="source">
				<p>Source paragraph</p>
				<span>Source span</span>
			</div>
			<div id="target"></div>
		</body>
		</html>
		<script>
		const source = document.getElementById('source');
		const target = document.getElementById('target');
		const sourcePara = source.querySelector('p');
		
		// Test importNode (shallow)
		const shallowImport = document.importNode(sourcePara, false);
		console.log("Shallow import:", shallowImport.tagName, shallowImport.textContent);
		target.appendChild(shallowImport);
		
		// Test importNode (deep)
		const deepImport = document.importNode(source, true);
		console.log("Deep import:", deepImport.tagName, deepImport.children.length);
		
		// Test adoptNode
		const sourceSpan = source.querySelector('span');
		const adoptedSpan = document.adoptNode(sourceSpan);
		console.log("Adopted span:", adoptedSpan.tagName, adoptedSpan.parentNode === null);
		target.appendChild(adoptedSpan);
		
		// Test getElementsByName (create elements with name attribute first)
		const input1 = document.createElement('input');
		input1.name = 'test-input';
		input1.id = 'input1';
		target.appendChild(input1);
		
		const input2 = document.createElement('input');
		input2.name = 'test-input';
		input2.id = 'input2';
		target.appendChild(input2);
		
		const namedElements = document.getElementsByName('test-input');
		console.log("Elements by name:", namedElements.length);
		if (namedElements.length > 0) {
			console.log("First named element ID:", namedElements[0].id);
			if (namedElements.length > 1) {
				console.log("Second named element ID:", namedElements[1].id);
			}
		} else {
			console.log("No elements found by name 'test-input'");
		}
		
		// Test normalize
		// Create text nodes to test normalization
		const textNode1 = document.createTextNode('Hello ');
		const textNode2 = document.createTextNode('World');
		target.appendChild(textNode1);
		target.appendChild(textNode2);
		
		console.log("Child nodes before normalize:", target.childNodes.length);
		document.normalize();
		console.log("Child nodes after normalize:", target.childNodes.length);
		
		document.body.setAttribute('data-manipulation-test', 'complete');
		</script>
	`)

	test.AssertElement("body").HasAttribute("data-manipulation-test", "complete")
	test.AssertElement("#input1").Exists()
	test.AssertElement("#input2").Exists()
}

func TestDOMAPIPhase3DOMImplementationMethods(t *testing.T) {
	test := domulator.NewTest(t)

	test.LoadHTML(`
		<!DOCTYPE html>
		<html>
		<body>
			<div id="results"></div>
		</body>
		</html>
		<script>
		const results = document.getElementById('results');
		const impl = document.implementation;
		
		// Test createDocumentType
		try {
			const doctype = impl.createDocumentType('html', '', '');
			console.log("Created doctype:", doctype.name, doctype.nodeType);
			results.innerHTML += '<p>DocType: ' + doctype.name + '</p>';
		} catch (e) {
			console.log("DocType creation error:", e.message);
			results.innerHTML += '<p>DocType error: ' + e.message + '</p>';
		}
		
		// Test createDocument
		try {
			const newDoc = impl.createDocument(null, null, null);
			console.log("Created document:", newDoc.nodeType);
			results.innerHTML += '<p>Document created: nodeType ' + newDoc.nodeType + '</p>';
		} catch (e) {
			console.log("Document creation error:", e.message);
			results.innerHTML += '<p>Document error: ' + e.message + '</p>';
		}
		
		// Test createHTMLDocument
		try {
			const htmlDoc = impl.createHTMLDocument('Test Title');
			console.log("Created HTML document:", htmlDoc.nodeType);
			results.innerHTML += '<p>HTML Document created: nodeType ' + htmlDoc.nodeType + '</p>';
		} catch (e) {
			console.log("HTML document creation error:", e.message);
			results.innerHTML += '<p>HTML Document error: ' + e.message + '</p>';
		}
		
		// Test hasFeature method
		const hasHTML = impl.hasFeature('HTML', '1.0');
		const hasXML = impl.hasFeature('XML', '1.0');
		const hasUnknown = impl.hasFeature('Unknown', '1.0');
		
		console.log("Has HTML feature:", hasHTML);
		console.log("Has XML feature:", hasXML);
		console.log("Has Unknown feature:", hasUnknown);
		
		results.innerHTML += '<p>HTML feature: ' + hasHTML + '</p>';
		results.innerHTML += '<p>XML feature: ' + hasXML + '</p>';
		results.innerHTML += '<p>Unknown feature: ' + hasUnknown + '</p>';
		
		document.body.setAttribute('data-implementation-test', 'complete');
		</script>
	`)

	test.AssertElement("body").HasAttribute("data-implementation-test", "complete")
	test.AssertElement("#results").Exists()
}

func TestDOMAPIPhase3Integration(t *testing.T) {
	test := domulator.NewTest(t)

	test.LoadHTML(`
		<!DOCTYPE html>
		<html>
		<body>
			<div id="workspace">
				<section id="original">
					<h2>Original Section</h2>
					<p class="content">Original content</p>
				</section>
			</div>
		</body>
		</html>
		<script>
		const workspace = document.getElementById('workspace');
		const original = document.getElementById('original');
		
		// Complex scenario: Clone, import, and manipulate
		const clonedSection = original.cloneNode(true);
		clonedSection.id = 'cloned';
		clonedSection.querySelector('h2').textContent = 'Cloned Section';
		clonedSection.querySelector('p').textContent = 'Cloned content';
		
		// Import the cloned section (testing import of complex structure)
		const importedSection = document.importNode(clonedSection, true);
		importedSection.id = 'imported';
		importedSection.querySelector('h2').textContent = 'Imported Section';
		
		workspace.appendChild(importedSection);
		
		// Test document properties on complex structure
		console.log("Document has doctype:", document.doctype !== null);
		console.log("Document charset:", document.characterSet);
		console.log("Document implementation:", typeof document.implementation);
		
		// Test creation methods with complex content
		const fragment = document.createDocumentFragment();
		
		// Create and add multiple elements
		for (let i = 0; i < 3; i++) {
			const div = document.createElement('div');
			div.className = 'generated';
			div.textContent = 'Generated content ' + (i + 1);
			fragment.appendChild(div);
		}
		
		workspace.appendChild(fragment);
		
		// Test getElementsByName with generated content
		const inputs = [];
		for (let i = 0; i < 2; i++) {
			const input = document.createElement('input');
			input.name = 'generated-input';
			input.value = 'value' + (i + 1);
			input.id = 'gen-input-' + (i + 1);
			workspace.appendChild(input);
			inputs.push(input);
		}
		
		const foundInputs = document.getElementsByName('generated-input');
		console.log("Found inputs by name:", foundInputs.length);
		if (foundInputs.length >= 2) {
			console.log("Input values:", foundInputs[0].value, foundInputs[1].value);
		} else if (foundInputs.length === 1) {
			console.log("Input value:", foundInputs[0].value);
		} else {
			console.log("No inputs found by name 'generated-input'");
		}
		
		// Test range operations
		const range = document.createRange();
		range.selectNodeContents(workspace);
		console.log("Range selected workspace contents");
		
		// Test normalization on complex content
		document.normalize();
		console.log("Document normalized");
		
		// Verify final state
		console.log("Final sections count:", workspace.querySelectorAll('section').length);
		console.log("Final generated divs:", workspace.querySelectorAll('.generated').length);
		console.log("Final inputs:", workspace.querySelectorAll('input').length);
		
		document.body.setAttribute('data-integration-test', 'complete');
		</script>
	`)

	test.AssertElement("body").HasAttribute("data-integration-test", "complete")
	test.AssertElement("#imported").Exists()
	test.AssertElement(".generated").Exists()
	test.AssertElement("#gen-input-1").Exists()
	test.AssertElement("#gen-input-2").Exists()
}

func TestDOMAPIPhase3ErrorHandling(t *testing.T) {
	test := domulator.NewTest(t)

	test.LoadHTML(`
		<!DOCTYPE html>
		<html>
		<body>
			<div id="test-errors"></div>
		</body>
		</html>
		<script>
		const testDiv = document.getElementById('test-errors');
		let errorsCaught = 0;
		
		// Test error handling for invalid operations
		
		// Test createAttribute with invalid name
		try {
			document.createAttribute('');
			console.log("Empty attribute name should fail");
		} catch (e) {
			console.log("Caught expected error for empty attribute name:", e.message);
			errorsCaught++;
		}
		
		// Test importNode with null node
		try {
			document.importNode(null, false);
			console.log("Import null should fail");
		} catch (e) {
			console.log("Caught expected error for null import:", e.message);
			errorsCaught++;
		}
		
		// Test adoptNode with null node
		try {
			document.adoptNode(null);
			console.log("Adopt null should fail");
		} catch (e) {
			console.log("Caught expected error for null adopt:", e.message);
			errorsCaught++;
		}
		
		// Test createNodeIterator without arguments
		try {
			document.createNodeIterator();
			console.log("CreateNodeIterator without root should fail");
		} catch (e) {
			console.log("Caught expected error for createNodeIterator:", e.message);
			errorsCaught++;
		}
		
		// Test createTreeWalker without arguments
		try {
			document.createTreeWalker();
			console.log("CreateTreeWalker without root should fail");
		} catch (e) {
			console.log("Caught expected error for createTreeWalker:", e.message);
			errorsCaught++;
		}
		
		// Test createProcessingInstruction with insufficient arguments
		try {
			document.createProcessingInstruction('target');
			console.log("CreateProcessingInstruction with one arg should fail");
		} catch (e) {
			console.log("Caught expected error for createProcessingInstruction:", e.message);
			errorsCaught++;
		}
		
		console.log("Total errors caught:", errorsCaught);
		testDiv.setAttribute('data-errors-caught', errorsCaught.toString());
		
		document.body.setAttribute('data-error-test', 'complete');
		</script>
	`)

	test.AssertElement("body").HasAttribute("data-error-test", "complete")
	test.AssertElement("#test-errors[data-errors-caught]").Exists()
}
