package integration

import (
	"testing"

	domulator "github.com/dpemmons/DOMulator"
)

func TestDOMAPIPhase2NavigationProperties(t *testing.T) {
	test := domulator.NewTest(t)

	test.LoadHTML(`
		<!DOCTYPE html>
		<html>
		<body>
			<div id="parent">
				<span id="first">First</span>
				Text node between
				<p id="middle">Middle</p>
				More text
				<strong id="last">Last</strong>
			</div>
		</body>
		</html>
		<script>
		// Test element navigation properties
		const parent = document.getElementById('parent');
		const first = document.getElementById('first');
		const middle = document.getElementById('middle');
		const last = document.getElementById('last');
		
		// Test children property - should only include element children
		console.log("Parent children length:", parent.children.length);
		console.log("First child tag:", parent.children[0].tagName);
		console.log("Last child tag:", parent.children[2].tagName);
		
		// Test firstElementChild and lastElementChild
		console.log("First element child:", parent.firstElementChild.id);
		console.log("Last element child:", parent.lastElementChild.id);
		
		// Test childElementCount
		console.log("Child element count:", parent.childElementCount);
		
		// Test sibling navigation
		console.log("Middle previous element:", middle.previousElementSibling.id);
		console.log("Middle next element:", middle.nextElementSibling.id);
		
		// Test edge cases - first/last elements
		console.log("First previous element:", first.previousElementSibling);
		console.log("Last next element:", last.nextElementSibling);
		
		// Mark navigation test complete
		document.body.setAttribute('data-navigation-test', 'complete');
		</script>
	`)

	test.AssertElement("body").HasAttribute("data-navigation-test", "complete")
}

func TestDOMAPIPhase2ModernManipulation(t *testing.T) {
	test := domulator.NewTest(t)

	test.LoadHTML(`
		<!DOCTYPE html>
		<html>
		<body>
			<div id="container">
				<p id="existing">Existing content</p>
			</div>
		</body>
		</html>
		<script>
		const container = document.getElementById('container');
		const existing = document.getElementById('existing');
		
		// Test append() with strings and elements
		const newSpan = document.createElement('span');
		newSpan.textContent = 'New span';
		newSpan.id = 'new-span';
		
		container.append('Text string', newSpan);
		
		// Test prepend() with strings and elements  
		const newDiv = document.createElement('div');
		newDiv.textContent = 'Prepended div';
		newDiv.id = 'prepended';
		
		container.prepend(newDiv, 'Prepended text');
		
		// Test replaceChildren()
		const replacement = document.createElement('h1');
		replacement.textContent = 'Replaced all';
		replacement.id = 'replacement';
		
		// Test current state first
		console.log("Children before replace:", container.children.length);
		
		// Replace all children
		container.replaceChildren(replacement, 'Replacement text');
		
		console.log("Children after replace:", container.children.length);
		console.log("First child after replace:", container.firstElementChild.id);
		
		document.body.setAttribute('data-modern-test', 'complete');
		</script>
	`)

	test.AssertElement("body").HasAttribute("data-modern-test", "complete")
	test.AssertElement("#replacement").Exists()
}

func TestDOMAPIPhase2ChildNodeMixin(t *testing.T) {
	test := domulator.NewTest(t)

	test.LoadHTML(`
		<!DOCTYPE html>
		<html>
		<body>
			<div id="container">
				<p id="target">Target element</p>
			</div>
		</body>
		</html>
		<script>
		const container = document.getElementById('container');
		const target = document.getElementById('target');
		
		// Test before() method
		const beforeDiv = document.createElement('div');
		beforeDiv.textContent = 'Before target';
		beforeDiv.id = 'before';
		
		target.before(beforeDiv, 'Before text');
		
		// Test after() method
		const afterDiv = document.createElement('div');
		afterDiv.textContent = 'After target';
		afterDiv.id = 'after';
		
		target.after('After text', afterDiv);
		
		// Check current state
		console.log("Container children:", container.children.length);
		console.log("Children order:", 
			container.children[0].id, 
			container.children[1].id,
			container.children[2].id);
		
		// Test replaceWith() method
		const replacement = document.createElement('section');
		replacement.textContent = 'Replacement section';
		replacement.id = 'replacement';
		
		target.replaceWith(replacement, 'Replacement text');
		
		console.log("After replace - children:", container.children.length);
		console.log("Middle child after replace:", container.children[1].id);
		
		// Test remove() method
		const toRemove = document.getElementById('before');
		toRemove.remove();
		
		console.log("After remove - children:", container.children.length);
		console.log("First child after remove:", container.children[0].id);
		
		document.body.setAttribute('data-childnode-test', 'complete');
		</script>
	`)

	test.AssertElement("body").HasAttribute("data-childnode-test", "complete")
	test.AssertElement("#replacement").Exists()
	test.AssertElement("#after").Exists()
}

func TestDOMAPIPhase2ElementsIntegration(t *testing.T) {
	test := domulator.NewTest(t)

	test.LoadHTML(`
		<!DOCTYPE html>
		<html>
		<body>
			<div id="root">
				<div class="section">
					<h2>Section 1</h2>
					<p>Paragraph 1</p>
				</div>
				<div class="section">
					<h2>Section 2</h2>
					<p>Paragraph 2</p>
				</div>
			</div>
		</body>
		</html>
		<script>
		const root = document.getElementById('root');
		
		// Test navigation with real DOM structure
		console.log("Root child element count:", root.childElementCount);
		console.log("First section children:", root.firstElementChild.children.length);
		console.log("Second section first child:", root.lastElementChild.firstElementChild.tagName);
		
		// Test modern manipulation with complex structure
		const newSection = document.createElement('div');
		newSection.className = 'section';
		const newH2 = document.createElement('h2');
		newH2.textContent = 'Section 3';
		const newP = document.createElement('p');
		newP.textContent = 'Paragraph 3';
		newSection.append(newH2, newP);
		
		root.append(newSection);
		
		console.log("After adding section:", root.childElementCount);
		console.log("Last section children:", root.lastElementChild.children.length);
		
		// Test ChildNode methods with nested structure
		const firstSection = root.firstElementChild;
		const insertedDiv = document.createElement('div');
		insertedDiv.className = 'inserted';
		insertedDiv.textContent = 'Inserted between sections';
		
		firstSection.after(insertedDiv);
		
		console.log("After inserting between:", root.childElementCount);
		console.log("Middle element class:", root.children[1].className);
		
		// Test element removal
		const toRemove = root.children[1]; // The inserted div
		toRemove.remove();
		
		console.log("After removal:", root.childElementCount);
		console.log("Second child is section:", root.children[1].className);
		
		document.body.setAttribute('data-integration-test', 'complete');
		</script>
	`)

	test.AssertElement("body").HasAttribute("data-integration-test", "complete")
	test.AssertElement(".section").Exists()
}

func TestDOMAPIPhase2NavigationUpdates(t *testing.T) {
	test := domulator.NewTest(t)

	test.LoadHTML(`
		<!DOCTYPE html>
		<html>
		<body>
			<div id="dynamic">
				<span>Initial</span>
			</div>
		</body>
		</html>
		<script>
		const dynamic = document.getElementById('dynamic');
		
		// Initial state
		console.log("Initial children:", dynamic.childElementCount);
		console.log("Initial first child:", dynamic.firstElementChild.tagName);
		console.log("Initial last child:", dynamic.lastElementChild.tagName);
		
		// Add elements and verify navigation updates
		const first = document.createElement('div');
		first.textContent = 'First';
		first.id = 'first';
		
		const last = document.createElement('section');
		last.textContent = 'Last';
		last.id = 'last';
		
		dynamic.prepend(first);
		dynamic.append(last);
		
		console.log("After adding - children:", dynamic.childElementCount);
		console.log("New first child:", dynamic.firstElementChild.id);
		console.log("New last child:", dynamic.lastElementChild.id);
		console.log("Middle child:", dynamic.children[1].tagName);
		
		// Test sibling navigation after manipulation
		const middle = dynamic.children[1]; // The original span
		console.log("Middle previous:", middle.previousElementSibling.id);
		console.log("Middle next:", middle.nextElementSibling.id);
		
		// Test replaceChildren and verify updates
		const replacement1 = document.createElement('p');
		replacement1.textContent = 'Replacement 1';
		replacement1.id = 'repl1';
		
		const replacement2 = document.createElement('p');
		replacement2.textContent = 'Replacement 2';  
		replacement2.id = 'repl2';
		
		dynamic.replaceChildren(replacement1, replacement2);
		
		console.log("After replaceChildren:", dynamic.childElementCount);
		console.log("New first after replace:", dynamic.firstElementChild.id);
		console.log("New last after replace:", dynamic.lastElementChild.id);
		
		document.body.setAttribute('data-updates-test', 'complete');
		</script>
	`)

	test.AssertElement("body").HasAttribute("data-updates-test", "complete")
	test.AssertElement("#repl1").Exists()
	test.AssertElement("#repl2").Exists()
}
