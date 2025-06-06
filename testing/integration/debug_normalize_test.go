package integration

import (
	"testing"

	domulator "github.com/dpemmons/DOMulator"
)

func TestDebugNormalize(t *testing.T) {
	test := domulator.NewTest(t).SetDebugMode(true)

	test.LoadHTML(`
		<!DOCTYPE html>
		<html>
		<body>
			<p id="para">First paragraph</p>
		</body>
		</html>
		<script>
		const para = document.getElementById('para');
		
		console.log("=== NORMALIZE DEBUG TEST ===");
		console.log("Initial children count:", para.childNodes.length);
		console.log("Initial text content:", JSON.stringify(para.textContent));
		
		// Show initial children
		for (let i = 0; i < para.childNodes.length; i++) {
			const child = para.childNodes[i];
			console.log("Initial child", i, "nodeType:", child.nodeType, "data:", JSON.stringify(child.data || child.textContent));
		}
		
		// Add three more text nodes
		const textNode1 = document.createTextNode(' Extra ');
		const textNode2 = document.createTextNode('text ');
		const textNode3 = document.createTextNode('content');
		
		para.appendChild(textNode1);
		para.appendChild(textNode2);
		para.appendChild(textNode3);
		
		console.log("After adding nodes - children count:", para.childNodes.length);
		
		// Show each child node's data after adding
		for (let i = 0; i < para.childNodes.length; i++) {
			const child = para.childNodes[i];
			console.log("Before normalize - Child", i, "nodeType:", child.nodeType, "data:", JSON.stringify(child.data || child.textContent));
		}
		
		const childCountBefore = para.childNodes.length;
		
		// Now normalize
		console.log("Calling normalize...");
		para.normalize();
		
		const childCountAfter = para.childNodes.length;
		console.log("After normalize - children count:", childCountAfter);
		console.log("childCountBefore:", childCountBefore, "childCountAfter:", childCountAfter);
		console.log("childCountAfter < childCountBefore:", childCountAfter < childCountBefore);
		
		// Show each child node's data after normalize
		for (let i = 0; i < para.childNodes.length; i++) {
			const child = para.childNodes[i];
			console.log("After normalize - Child", i, "nodeType:", child.nodeType, "data:", JSON.stringify(child.data || child.textContent));
		}
		
		// Set the result
		const success = childCountAfter < childCountBefore;
		document.body.setAttribute('data-normalize-result', success.toString());
		document.body.setAttribute('data-child-count-before', childCountBefore.toString());
		document.body.setAttribute('data-child-count-after', childCountAfter.toString());
		document.body.setAttribute('data-final-text', para.textContent);
		console.log("=== END NORMALIZE DEBUG TEST ===");
		</script>
	`)

	// Check that the script completed
	test.AssertElement("body[data-normalize-result]").Exists()
}
