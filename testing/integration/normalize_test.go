package integration

import (
	"testing"

	domulator "github.com/dpemmons/DOMulator"
	testpkg "github.com/dpemmons/DOMulator/testing"
)

func TestNormalizeBasic(t *testing.T) {
	test := domulator.NewTest(t)

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

	// Assert DOM state
	test.AssertElement("body").HasAttribute("data-before-count", "3")
	test.AssertElement("body").HasAttribute("data-after-count", "1")
	test.AssertElement("body").HasAttribute("data-text-content", "Hello World!")
	test.AssertElement("body").HasAttribute("data-normalized", "true")
}

func TestNormalizeBasicWithConsoleCapture(t *testing.T) {
	harness := testpkg.NewTestHarness()
	defer harness.Close()

	consoleCapture := harness.CaptureConsole(t)

	harness.LoadHTML(`
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

	// Assert console output
	consoleCapture.AssertLogContains("Before normalize:")
	consoleCapture.AssertLogContains("Child count:")
	consoleCapture.AssertLogContains("Text content:")
	consoleCapture.AssertLogContains("After normalize:")
	consoleCapture.AssertLogContains("3") // Before count
	consoleCapture.AssertLogContains("1") // After count
	consoleCapture.AssertLogContains("Hello World!")

	// Assert DOM state using harness
	body := harness.Document().QuerySelector("body")
	if body == nil {
		t.Fatal("Could not find body element")
	}

	if beforeCount := body.GetAttribute("data-before-count"); beforeCount != "3" {
		t.Errorf("Expected before count 3, got %s", beforeCount)
	}
	if afterCount := body.GetAttribute("data-after-count"); afterCount != "1" {
		t.Errorf("Expected after count 1, got %s", afterCount)
	}
	if textContent := body.GetAttribute("data-text-content"); textContent != "Hello World!" {
		t.Errorf("Expected text content 'Hello World!', got %s", textContent)
	}
	if normalized := body.GetAttribute("data-normalized"); normalized != "true" {
		t.Errorf("Expected normalized true, got %s", normalized)
	}
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

func TestNormalizeWithExistingContentWithConsoleCapture(t *testing.T) {
	harness := testpkg.NewTestHarness()
	defer harness.Close()

	consoleCapture := harness.CaptureConsole(t)

	harness.LoadHTML(`
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

	// Assert console output
	consoleCapture.AssertLogContains("Initial state:")
	consoleCapture.AssertLogContains("After adding nodes:")
	consoleCapture.AssertLogContains("After normalize:")
	consoleCapture.AssertLogContains("4") // Before count (1 + 3 added)
	consoleCapture.AssertLogContains("1") // After count
	consoleCapture.AssertLogContains("Initial text")
	consoleCapture.AssertLogContains("Extra")
	consoleCapture.AssertLogContains("Node")
	consoleCapture.AssertLogContains(":") // From node logging

	// Assert DOM state
	body := harness.Document().QuerySelector("body")
	if body == nil {
		t.Fatal("Could not find body element")
	}

	if beforeCount := body.GetAttribute("data-before"); beforeCount != "4" {
		t.Errorf("Expected before count 4, got %s", beforeCount)
	}
	if afterCount := body.GetAttribute("data-after"); afterCount != "1" {
		t.Errorf("Expected after count 1, got %s", afterCount)
	}
	if reduced := body.GetAttribute("data-reduced"); reduced != "true" {
		t.Errorf("Expected reduced true, got %s", reduced)
	}
	if finalText := body.GetAttribute("data-final-text"); finalText != "Initial text Extra text content" {
		t.Errorf("Expected text 'Initial text Extra text content', got %s", finalText)
	}
}
