package integration

import (
	"testing"

	domulator "github.com/dpemmons/DOMulator"
)

func TestDOMAPIPhase4NodeComparison(t *testing.T) {
	test := domulator.NewTest(t)

	test.LoadHTML(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>Phase 4 Test</title>
		</head>
		<body>
			<div id="parent">
				<p id="child1">First child</p>
				<p id="child2">Second child</p>
			</div>
			<div id="other">Other element</div>
		</body>
		</html>
		<script>
		const parent = document.getElementById('parent');
		const child1 = document.getElementById('child1');
		const child2 = document.getElementById('child2');
		const other = document.getElementById('other');
		
		// Test compareDocumentPosition
		const position1 = parent.compareDocumentPosition(child1);
		const position2 = child1.compareDocumentPosition(child2);
		const position3 = child1.compareDocumentPosition(parent);
		
		console.log("Parent vs Child1 position:", position1);
		console.log("Child1 vs Child2 position:", position2);
		console.log("Child1 vs Parent position:", position3);
		
		// Test contains
		const parentContainsChild1 = parent.contains(child1);
		const parentContainsChild2 = parent.contains(child2);
		const child1ContainsParent = child1.contains(parent);
		const parentContainsOther = parent.contains(other);
		
		console.log("Parent contains child1:", parentContainsChild1);
		console.log("Parent contains child2:", parentContainsChild2);
		console.log("Child1 contains parent:", child1ContainsParent);
		console.log("Parent contains other:", parentContainsOther);
		
		// Test isEqualNode
		const clonedChild1 = child1.cloneNode(true);
		const isEqual1 = child1.isEqualNode(clonedChild1);
		const isEqual2 = child1.isEqualNode(child2);
		
		console.log("Child1 equals cloned child1:", isEqual1);
		console.log("Child1 equals child2:", isEqual2);
		
		// Test isSameNode
		const isSame1 = child1.isSameNode(child1);
		const isSame2 = child1.isSameNode(clonedChild1);
		const isSame3 = child1.isSameNode(child2);
		
		console.log("Child1 same as itself:", isSame1);
		console.log("Child1 same as cloned:", isSame2);
		console.log("Child1 same as child2:", isSame3);
		
		// Verify results and mark test complete
		let results = [];
		results.push("compareDocumentPosition: " + (typeof position1 === 'number'));
		results.push("contains: " + (parentContainsChild1 === true && parentContainsOther === false));
		results.push("isEqualNode: " + (isEqual1 === true && isEqual2 === false));
		results.push("isSameNode: " + (isSame1 === true && isSame2 === false));
		
		document.body.setAttribute('data-comparison-results', results.join(','));
		document.body.setAttribute('data-comparison-test', 'complete');
		</script>
	`)

	test.AssertElement("body").HasAttribute("data-comparison-test", "complete")
	test.AssertElement("body").HasAttributeContaining("data-comparison-results", "compareDocumentPosition: true")
	test.AssertElement("body").HasAttributeContaining("data-comparison-results", "contains: true")
	test.AssertElement("body").HasAttributeContaining("data-comparison-results", "isEqualNode: true")
	test.AssertElement("body").HasAttributeContaining("data-comparison-results", "isSameNode: true")
}

func TestDOMAPIPhase4NamespaceSupport(t *testing.T) {
	test := domulator.NewTest(t)

	test.LoadHTML(`
		<!DOCTYPE html>
		<html xmlns="http://www.w3.org/1999/xhtml">
		<head>
			<title>Phase 4 Namespace Test</title>
		</head>
		<body>
			<div id="test-element">Test</div>
		</body>
		</html>
		<script>
		const element = document.getElementById('test-element');
		const htmlElement = document.documentElement;
		
		// Test namespace properties
		const namespaceURI = element.namespaceURI;
		const prefix = element.prefix;
		const localName = element.localName;
		
		console.log("Element namespaceURI:", namespaceURI);
		console.log("Element prefix:", prefix);
		console.log("Element localName:", localName);
		
		// Test HTML element namespace
		const htmlNamespaceURI = htmlElement.namespaceURI;
		const htmlLocalName = htmlElement.localName;
		
		console.log("HTML namespaceURI:", htmlNamespaceURI);
		console.log("HTML localName:", htmlLocalName);
		
		// Test namespace lookup methods
		const lookupPrefix1 = element.lookupPrefix("http://www.w3.org/1999/xhtml");
		const lookupPrefix2 = element.lookupPrefix("http://example.com/unknown");
		
		console.log("Lookup prefix for XHTML:", lookupPrefix1);
		console.log("Lookup prefix for unknown:", lookupPrefix2);
		
		const lookupNS1 = element.lookupNamespaceURI(null);
		const lookupNS2 = element.lookupNamespaceURI("unknown");
		
		console.log("Lookup namespace for null prefix:", lookupNS1);
		console.log("Lookup namespace for unknown prefix:", lookupNS2);
		
		// Test default namespace
		const isDefaultNS1 = element.isDefaultNamespace("http://www.w3.org/1999/xhtml");
		const isDefaultNS2 = element.isDefaultNamespace("http://example.com/unknown");
		
		console.log("Is XHTML default namespace:", isDefaultNS1);
		console.log("Is unknown default namespace:", isDefaultNS2);
		
		// Verify namespace properties are read-only (attempting to set should not work)
		try {
			element.prefix = "test";
			console.log("Prefix after attempt to set:", element.prefix);
		} catch (e) {
			console.log("Prefix setter error:", e.message);
		}
		
		// Mark test complete with results
		let results = [];
		results.push("namespaceURI: " + (typeof namespaceURI));
		results.push("localName: " + (localName === 'div'));
		results.push("lookupMethods: " + (typeof lookupPrefix1 !== 'undefined' && typeof lookupNS1 !== 'undefined'));
		results.push("defaultNamespace: " + (typeof isDefaultNS1 === 'boolean'));
		
		document.body.setAttribute('data-namespace-results', results.join(','));
		document.body.setAttribute('data-namespace-test', 'complete');
		</script>
	`)

	test.AssertElement("body").HasAttribute("data-namespace-test", "complete")
	test.AssertElement("body").HasAttributeContaining("data-namespace-results", "localName: true")
	test.AssertElement("body").HasAttributeContaining("data-namespace-results", "lookupMethods: true")
	test.AssertElement("body").HasAttributeContaining("data-namespace-results", "defaultNamespace: true")
}

func TestDOMAPIPhase4AdvancedMethods(t *testing.T) {
	test := domulator.NewTest(t)

	test.LoadHTML(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>Phase 4 Advanced Methods Test</title>
		</head>
		<body>
			<div id="parent">
				<text>Some text</text>
				<p id="child">Child element</p>
				<text>More text</text>
			</div>
			<div id="empty"></div>
		</body>
		</html>
		<script>
		const parent = document.getElementById('parent');
		const child = document.getElementById('child');
		const empty = document.getElementById('empty');
		
		// Test hasChildNodes
		const parentHasChildren = parent.hasChildNodes();
		const childHasChildren = child.hasChildNodes();
		const emptyHasChildren = empty.hasChildNodes();
		
		console.log("Parent has child nodes:", parentHasChildren);
		console.log("Child has child nodes:", childHasChildren);
		console.log("Empty has child nodes:", emptyHasChildren);
		
		// Test getRootNode
		const parentRoot = parent.getRootNode();
		const childRoot = child.getRootNode();
		
		console.log("Parent root node:", parentRoot.nodeType);
		console.log("Child root node:", childRoot.nodeType);
		console.log("Root nodes are same:", parentRoot === childRoot);
		
		// Test getRootNode with options
		const rootWithOptions = parent.getRootNode({ composed: false });
		console.log("Root with options:", rootWithOptions.nodeType);
		
		// Test normalize method
		// Create text nodes for normalization test
		const textNode1 = document.createTextNode('Hello ');
		const textNode2 = document.createTextNode('World');
		empty.appendChild(textNode1);
		empty.appendChild(textNode2);
		
		console.log("Child nodes before normalize:", empty.childNodes.length);
		empty.normalize();
		console.log("Child nodes after normalize:", empty.childNodes.length);
		console.log("Text content after normalize:", empty.textContent);
		
		// Document normalize
		document.normalize();
		console.log("Document normalized successfully");
		
		// Verify all methods work correctly
		let results = [];
		results.push("hasChildNodes: " + (parentHasChildren === true && emptyHasChildren === false));
		results.push("getRootNode: " + (parentRoot.nodeType === 9)); // Document node
		results.push("normalize: " + (empty.textContent === 'Hello World'));
		
		document.body.setAttribute('data-advanced-results', results.join(','));
		document.body.setAttribute('data-advanced-test', 'complete');
		</script>
	`)

	test.AssertElement("body").HasAttribute("data-advanced-test", "complete")
	test.AssertElement("body").HasAttributeContaining("data-advanced-results", "hasChildNodes: true")
	test.AssertElement("body").HasAttributeContaining("data-advanced-results", "getRootNode: true")
	test.AssertElement("body").HasAttributeContaining("data-advanced-results", "normalize: true")
}

func TestDOMAPIPhase4IntegrationScenarios(t *testing.T) {
	test := domulator.NewTest(t)

	test.LoadHTML(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>Phase 4 Integration Test</title>
		</head>
		<body>
			<div id="container">
				<section id="section1" class="test-section">
					<h2>Section 1</h2>
					<p id="para1">First paragraph</p>
				</section>
				<section id="section2" class="test-section">
					<h2>Section 2</h2>
					<p id="para2">Second paragraph</p>
				</section>
			</div>
		</body>
		</html>
		<script>
		const container = document.getElementById('container');
		const section1 = document.getElementById('section1');
		const section2 = document.getElementById('section2');
		const para1 = document.getElementById('para1');
		const para2 = document.getElementById('para2');
		
		// Complex scenario: Clone, compare, and manipulate
		const clonedSection = section1.cloneNode(true);
		
		// Test node relationships after cloning (BEFORE modification)
		const originalIsEqual = section1.isEqualNode(clonedSection);
		const originalIsSame = section1.isSameNode(clonedSection);
		
		console.log("Original equals cloned:", originalIsEqual);
		console.log("Original same as cloned:", originalIsSame);
		
		// Now modify the clone
		clonedSection.id = 'cloned-section';
		
		// Test containment relationships
		const containerContainsSection1 = container.contains(section1);
		const section1ContainsPara1 = section1.contains(para1);
		const section1ContainsPara2 = section1.contains(para2);
		
		console.log("Container contains section1:", containerContainsSection1);
		console.log("Section1 contains para1:", section1ContainsPara1);
		console.log("Section1 contains para2:", section1ContainsPara2);
		
		// Test document position between different elements
		const pos1 = section1.compareDocumentPosition(section2);
		const pos2 = para1.compareDocumentPosition(para2);
		const pos3 = container.compareDocumentPosition(para1);
		
		console.log("Section1 vs Section2 position:", pos1);
		console.log("Para1 vs Para2 position:", pos2);
		console.log("Container vs Para1 position:", pos3);
		
		// Test namespace operations on complex structure
		const section1NS = section1.namespaceURI;
		const section1LocalName = section1.localName;
		const para1NS = para1.lookupNamespaceURI(null);
		
		console.log("Section1 namespace:", section1NS);
		console.log("Section1 local name:", section1LocalName);
		console.log("Para1 default namespace:", para1NS);
		
		// Test advanced methods on complex structure
		const hasChildren = section1.hasChildNodes();
		const rootNode = para1.getRootNode();
		
		console.log("Section1 has children:", hasChildren);
		console.log("Para1 root is document:", rootNode.nodeType === 9);
		
		// Append cloned section and verify relationships change
		container.appendChild(clonedSection);
		const newContainment = container.contains(clonedSection);
		
		console.log("Container now contains cloned section:", newContainment);
		
		// Test normalization on complex content
		// Add multiple text nodes to test normalization
		const textNode1 = document.createTextNode(' Extra ');
		const textNode2 = document.createTextNode('text ');
		const textNode3 = document.createTextNode('content');
		
		para1.appendChild(textNode1);
		para1.appendChild(textNode2);
		para1.appendChild(textNode3);
		
		const childCountBefore = para1.childNodes.length;
		para1.normalize();
		const childCountAfter = para1.childNodes.length;
		
		console.log("Para1 child count before normalize:", childCountBefore);
		console.log("Para1 child count after normalize:", childCountAfter);
		console.log("Para1 final text:", para1.textContent);
		
		// Verify all complex operations worked
		let results = [];
		results.push("cloneComparison: " + (originalIsEqual === true && originalIsSame === false));
		results.push("containment: " + (containerContainsSection1 === true && section1ContainsPara2 === false));
		results.push("positioning: " + (typeof pos1 === 'number' && typeof pos2 === 'number'));
		results.push("namespace: " + (section1LocalName === 'section'));
		results.push("advanced: " + (hasChildren === true && rootNode.nodeType === 9));
		results.push("newRelation: " + (newContainment === true));
		results.push("normalization: " + (childCountAfter < childCountBefore));
		
		document.body.setAttribute('data-integration-results', results.join(','));
		document.body.setAttribute('data-integration-test', 'complete');
		</script>
	`)

	test.AssertElement("body").HasAttribute("data-integration-test", "complete")
	test.AssertElement("body").HasAttributeContaining("data-integration-results", "cloneComparison: true")
	test.AssertElement("body").HasAttributeContaining("data-integration-results", "containment: true")
	test.AssertElement("body").HasAttributeContaining("data-integration-results", "positioning: true")
	test.AssertElement("body").HasAttributeContaining("data-integration-results", "namespace: true")
	test.AssertElement("body").HasAttributeContaining("data-integration-results", "advanced: true")
	test.AssertElement("body").HasAttributeContaining("data-integration-results", "newRelation: true")
	test.AssertElement("body").HasAttributeContaining("data-integration-results", "normalization: true")
}

func TestDOMAPIPhase4ErrorHandling(t *testing.T) {
	test := domulator.NewTest(t)

	test.LoadHTML(`
		<!DOCTYPE html>
		<html>
		<body>
			<div id="test-element">Test</div>
		</body>
		</html>
		<script>
		const element = document.getElementById('test-element');
		let errorsCaught = 0;
		
		// Test error handling for node comparison methods
		
		// Test compareDocumentPosition with invalid node
		try {
			element.compareDocumentPosition(null);
			console.log("compareDocumentPosition with null should fail");
		} catch (e) {
			console.log("Caught expected error for compareDocumentPosition:", e.message);
			errorsCaught++;
		}
		
		// Test contains with invalid node
		try {
			element.contains(null);
			console.log("contains with null returns false");
		} catch (e) {
			console.log("contains with null threw error:", e.message);
			// contains with null should return false, not throw
		}
		
		// Test isEqualNode with invalid node
		try {
			element.isEqualNode(null);
			console.log("isEqualNode with null returns false");
		} catch (e) {
			console.log("isEqualNode with null threw error:", e.message);
			// isEqualNode with null should return false, not throw
		}
		
		// Test isSameNode with invalid node  
		try {
			element.isSameNode(null);
			console.log("isSameNode with null returns false");
		} catch (e) {
			console.log("isSameNode with null threw error:", e.message);
			// isSameNode with null should return false, not throw
		}
		
		// Test lookupPrefix with invalid namespace
		try {
			const result = element.lookupPrefix(null);
			console.log("lookupPrefix with null returned:", result);
		} catch (e) {
			console.log("lookupPrefix error:", e.message);
		}
		
		// Test lookupNamespaceURI with invalid prefix
		try {
			const result = element.lookupNamespaceURI(null);
			console.log("lookupNamespaceURI with null returned:", result);
		} catch (e) {
			console.log("lookupNamespaceURI error:", e.message);
		}
		
		// Test isDefaultNamespace with invalid namespace
		try {
			const result = element.isDefaultNamespace(null);
			console.log("isDefaultNamespace with null returned:", result);
		} catch (e) {
			console.log("isDefaultNamespace error:", e.message);
		}
		
		console.log("Total errors caught:", errorsCaught);
		element.setAttribute('data-errors-caught', errorsCaught.toString());
		
		document.body.setAttribute('data-error-test', 'complete');
		</script>
	`)

	test.AssertElement("body").HasAttribute("data-error-test", "complete")
	test.AssertElement("#test-element[data-errors-caught]").Exists()
}
