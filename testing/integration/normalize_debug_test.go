package integration

import (
	"testing"

	domulator "github.com/dpemmons/DOMulator"
)

func TestNormalizeDebug(t *testing.T) {
	test := domulator.NewTest(t)
	test.SetDebugMode(true)

	test.LoadHTML(`
		<!DOCTYPE html>
		<html>
		<body>
			<div id="test"></div>
		</body>
		</html>
		<script>
		const div = document.getElementById('test');
		
		// Add multiple text nodes
		const text1 = document.createTextNode('Hello ');
		const text2 = document.createTextNode('World');
		const text3 = document.createTextNode('!');
		
		div.appendChild(text1);
		div.appendChild(text2);
		div.appendChild(text3);
		
		console.log('Before normalize:');
		console.log('Child count:', div.childNodes.length);
		console.log('Children types:');
		for (let i = 0; i < div.childNodes.length; i++) {
			const child = div.childNodes[i];
			console.log('  Child', i, '- nodeType:', child.nodeType, 'nodeValue:', JSON.stringify(child.nodeValue));
		}
		
		// Check siblings
		console.log('Sibling relationships:');
		console.log('  text1.nextSibling:', text1.nextSibling ? text1.nextSibling.nodeValue : 'null');
		console.log('  text2.previousSibling:', text2.previousSibling ? text2.previousSibling.nodeValue : 'null');
		console.log('  text2.nextSibling:', text2.nextSibling ? text2.nextSibling.nodeValue : 'null');
		console.log('  text3.previousSibling:', text3.previousSibling ? text3.previousSibling.nodeValue : 'null');
		
		// Call normalize
		console.log('Calling normalize...');
		div.normalize();
		
		console.log('After normalize:');
		console.log('Child count:', div.childNodes.length);
		console.log('Children after normalize:');
		for (let i = 0; i < div.childNodes.length; i++) {
			const child = div.childNodes[i];
			console.log('  Child', i, '- nodeType:', child.nodeType, 'nodeValue:', JSON.stringify(child.nodeValue));
		}
		
		// Test if CharacterData methods are available
		console.log('Testing CharacterData methods on first child:');
		const firstChild = div.childNodes[0];
		console.log('  typeof appendData:', typeof firstChild.appendData);
		console.log('  typeof replaceData:', typeof firstChild.replaceData);
		console.log('  typeof data property:', typeof firstChild.data);
		console.log('  data value:', firstChild.data);
		
		document.body.setAttribute('data-child-count', div.childNodes.length.toString());
		document.body.setAttribute('data-text-content', div.textContent);
		</script>
	`)

	// We're not asserting anything specific here, just want to see the debug output
	t.Log("Test completed - check debug output above")
}
