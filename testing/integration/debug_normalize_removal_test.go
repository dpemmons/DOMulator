package integration

import (
	"testing"

	domulator "github.com/dpemmons/DOMulator"
)

func TestDebugNormalizeRemoval(t *testing.T) {
	test := domulator.NewTest(t)

	test.LoadHTML(`
		<!DOCTYPE html>
		<html>
		<body>
			<p id="test">First</p>
		</body>
		</html>
		<script>
		const p = document.getElementById('test');
		
		console.log("=== REMOVAL DEBUG TEST ===");
		console.log("Initial children count:", p.childNodes.length);
		console.log("Initial text content:", p.textContent);
		
		// Add exactly one more text node
		const extraNode = document.createTextNode(' Second');
		p.appendChild(extraNode);
		
		console.log("After adding one node - children count:", p.childNodes.length);
		console.log("Child 0 data:", p.childNodes[0].data);
		console.log("Child 1 data:", p.childNodes[1].data);
		
		// Now normalize
		console.log("Calling normalize...");
		p.normalize();
		
		console.log("After normalize - children count:", p.childNodes.length);
		if (p.childNodes.length > 0) {
			console.log("Child 0 data:", p.childNodes[0].data);
		}
		if (p.childNodes.length > 1) {
			console.log("Child 1 still exists with data:", p.childNodes[1].data);
		}
		
		console.log("Final text content:", p.textContent);
		console.log("=== END REMOVAL DEBUG TEST ===");
		</script>
	`)
}
