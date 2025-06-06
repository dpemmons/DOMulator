package integration

import (
	"testing"

	domulator "github.com/dpemmons/DOMulator"
)

func TestNormalizeBasic(t *testing.T) {
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
		div.appendChild(document.createTextNode('Hello '));
		div.appendChild(document.createTextNode('World'));
		div.appendChild(document.createTextNode('!'));
		
		console.log('Before normalize:');
		console.log('Child count:', div.childNodes.length);
		console.log('Text content:', div.textContent);
		
		// Store before state
		const beforeCount = div.childNodes.length;
		
		// Normalize
		div.normalize();
		
		console.log('After normalize:');
		console.log('Child count:', div.childNodes.length);
		console.log('Text content:', div.textContent);
		
		// Store results
		const afterCount = div.childNodes.length;
		
		document.body.setAttribute('data-before-count', beforeCount.toString());
		document.body.setAttribute('data-after-count', afterCount.toString());
		document.body.setAttribute('data-text-content', div.textContent);
		document.body.setAttribute('data-normalized', (afterCount < beforeCount).toString());
		</script>
	`)

	test.AssertElement("body").HasAttribute("data-before-count", "3")
	test.AssertElement("body").HasAttribute("data-after-count", "1")
	test.AssertElement("body").HasAttribute("data-text-content", "Hello World!")
	test.AssertElement("body").HasAttribute("data-normalized", "true")
}

func TestNormalizeWithExistingContent(t *testing.T) {
	test := domulator.NewTest(t)

	test.LoadHTML(`
		<!DOCTYPE html>
		<html>
		<body>
			<p id="para">Initial text</p>
		</body>
		</html>
		<script>
		const para = document.getElementById('para');
		
		console.log('Initial state:');
		console.log('Child count:', para.childNodes.length);
		console.log('Text content:', para.textContent);
		
		// Add more text nodes
		para.appendChild(document.createTextNode(' Extra '));
		para.appendChild(document.createTextNode('text '));
		para.appendChild(document.createTextNode('content'));
		
		console.log('After adding nodes:');
		console.log('Child count:', para.childNodes.length);
		for (let i = 0; i < para.childNodes.length; i++) {
			console.log('Node', i, ':', para.childNodes[i].nodeValue);
		}
		
		const beforeCount = para.childNodes.length;
		para.normalize();
		const afterCount = para.childNodes.length;
		
		console.log('After normalize:');
		console.log('Child count:', para.childNodes.length);
		console.log('Text content:', para.textContent);
		
		document.body.setAttribute('data-before', beforeCount.toString());
		document.body.setAttribute('data-after', afterCount.toString());
		document.body.setAttribute('data-reduced', (afterCount < beforeCount).toString());
		document.body.setAttribute('data-final-text', para.textContent);
		</script>
	`)

	test.AssertElement("body").HasAttribute("data-before", "4")
	test.AssertElement("body").HasAttribute("data-after", "1")
	test.AssertElement("body").HasAttribute("data-reduced", "true")
	test.AssertElement("body").HasAttribute("data-final-text", "Initial text Extra text content")
}
