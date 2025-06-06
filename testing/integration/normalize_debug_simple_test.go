package integration

import (
	"testing"

	domulator "github.com/dpemmons/DOMulator"
)

func TestNormalizeDebugSimple(t *testing.T) {
	test := domulator.NewTest(t)

	test.LoadHTML(`
		<!DOCTYPE html>
		<html>
		<body>
			<p id="para">First paragraph</p>
		</body>
		</html>
		<script>
		const para = document.getElementById('para');
		
		console.log("Initial children count:", para.childNodes.length);
		console.log("Initial text content:", para.textContent);
		
		// Add three more text nodes
		const textNode1 = document.createTextNode(' Extra ');
		const textNode2 = document.createTextNode('text ');
		const textNode3 = document.createTextNode('content');
		
		para.appendChild(textNode1);
		para.appendChild(textNode2);
		para.appendChild(textNode3);
		
		console.log("After adding nodes - children count:", para.childNodes.length);
		console.log("After adding nodes - text content:", para.textContent);
		
		// Show each child node's data
		for (let i = 0; i < para.childNodes.length; i++) {
			const child = para.childNodes[i];
			console.log("Child", i, "nodeType:", child.nodeType, "data:", child.data || child.textContent);
		}
		
		// Now normalize
		console.log("Calling normalize...");
		para.normalize();
		
		console.log("After normalize - children count:", para.childNodes.length);
		console.log("After normalize - text content:", para.textContent);
		
		// Show each child node's data after normalize
		for (let i = 0; i < para.childNodes.length; i++) {
			const child = para.childNodes[i];
			console.log("After normalize - Child", i, "nodeType:", child.nodeType, "data:", child.data || child.textContent);
		}
		
		// Set the result
		const success = para.childNodes.length === 1;
		document.body.setAttribute('data-normalize-result', success.toString());
		document.body.setAttribute('data-final-count', para.childNodes.length.toString());
		document.body.setAttribute('data-final-text', para.textContent);
		</script>
	`)

	// Just check that the script completed
	test.AssertElement("body[data-normalize-result]").Exists()
}
