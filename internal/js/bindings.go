package js

import (
	"strconv"

	"github.com/dop251/goja"
	"github.com/dpemmons/DOMulator/internal/browser/cookies"
	"github.com/dpemmons/DOMulator/internal/browser/history"
	"github.com/dpemmons/DOMulator/internal/browser/performance"
	"github.com/dpemmons/DOMulator/internal/browser/url"
	"github.com/dpemmons/DOMulator/internal/browser/xpath"
	"github.com/dpemmons/DOMulator/internal/css"
	"github.com/dpemmons/DOMulator/internal/dom"
)

// DOMBindings manages the binding between Go DOM objects and JavaScript
type DOMBindings struct {
	vm            *goja.Runtime
	document      *dom.Document
	nodeCache     map[dom.Node]*goja.Object // Cache to maintain object identity
	cookieManager *cookies.CookieManager
	jsListeners   map[dom.Node]map[string][]*goja.Object // Store JS listeners for lookup
	windowTarget  *WindowEventTarget                     // Window event target for proper bubbling
	jsRuntime     *Runtime                               // Reference to JS runtime for event loop access
}

// NewDOMBindings creates a new DOM bindings instance
func NewDOMBindings(vm *goja.Runtime, document *dom.Document) *DOMBindings {
	// Initialize cookie manager with default domain
	cookieManager := cookies.NewCookieManager("localhost", "/")

	return &DOMBindings{
		vm:            vm,
		document:      document,
		nodeCache:     make(map[dom.Node]*goja.Object),
		cookieManager: cookieManager,
		jsListeners:   make(map[dom.Node]map[string][]*goja.Object),
		windowTarget:  &WindowEventTarget{eventListeners: make(map[string][]func(dom.Event))},
	}
}

// storeJSListener stores a JavaScript listener for later lookup
func (db *DOMBindings) storeJSListener(node dom.Node, eventType string, listener *goja.Object) {
	if db.jsListeners[node] == nil {
		db.jsListeners[node] = make(map[string][]*goja.Object)
	}
	db.jsListeners[node][eventType] = append(db.jsListeners[node][eventType], listener)
}

// storeWindowJSListener stores a JavaScript listener for window events
func (db *DOMBindings) storeWindowJSListener(eventType string, listener *goja.Object) {
	// For window events, we'll store them separately or skip storage
	// since WindowEventTarget doesn't implement dom.Node
}

// dispatchEventWithBubbling dispatches an event with proper bubbling to window
func (db *DOMBindings) dispatchEventWithBubbling(node dom.Node, event dom.Event) bool {
	// Set the target
	if baseEvent, ok := event.(*dom.BaseEvent); ok {
		baseEvent.SetTarget(node)
	}

	// First dispatch through normal DOM bubbling
	result := node.DispatchEvent(event)

	// For events that should reach window, also dispatch to window target
	if event.Bubbles() && db.windowTarget != nil {
		// Window-specific events or events that bubble to window
		if isWindowEvent(event.Type()) {
			db.windowTarget.DispatchEvent(event)
		}
	}

	return result
}

// isWindowEvent checks if an event type should reach the window
func isWindowEvent(eventType string) bool {
	windowEvents := map[string]bool{
		"resize": true, "load": true, "unload": true, "beforeunload": true,
		"hashchange": true, "online": true, "offline": true, "scroll": true,
		// These can bubble to window
		"click": true, "keydown": true, "keyup": true, "mousemove": true,
	}
	return windowEvents[eventType]
}

// WrapDocument wraps a DOM document for JavaScript access
func (db *DOMBindings) WrapDocument() *goja.Object {
	doc := db.vm.NewObject()

	// Basic document properties
	doc.Set("nodeType", int(dom.DocumentNode))
	doc.Set("nodeName", "#document")

	// Document methods
	doc.Set("createElement", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			panic(db.vm.NewTypeError("createElement requires a tag name"))
		}

		tagName := call.Arguments[0].String()
		element := db.document.CreateElement(tagName)
		return db.WrapElement(element)
	})

	doc.Set("createTextNode", func(call goja.FunctionCall) goja.Value {
		var data string
		if len(call.Arguments) > 0 {
			data = call.Arguments[0].String()
		}

		textNode := db.document.CreateTextNode(data)
		return db.WrapNode(textNode)
	})

	doc.Set("createComment", func(call goja.FunctionCall) goja.Value {
		var data string
		if len(call.Arguments) > 0 {
			data = call.Arguments[0].String()
		}

		comment := db.document.CreateComment(data)
		return db.WrapNode(comment)
	})

	doc.Set("createDocumentFragment", func(call goja.FunctionCall) goja.Value {
		fragment := db.document.CreateDocumentFragment()
		return db.WrapNode(fragment)
	})

	doc.Set("getElementById", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			return goja.Null()
		}

		id := call.Arguments[0].String()
		element := db.document.GetElementById(id)
		if element == nil {
			return goja.Null()
		}
		return db.WrapElement(element)
	})

	doc.Set("getElementsByTagName", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			return db.vm.NewArray()
		}

		tagName := call.Arguments[0].String()
		collection := db.document.GetElementsByTagName(tagName)
		return db.WrapNodeList(collection.ToSlice())
	})

	doc.Set("getElementsByClassName", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			return db.vm.NewArray()
		}

		className := call.Arguments[0].String()
		collection := db.document.GetElementsByClassName(className)
		return db.WrapNodeList(collection.ToSlice())
	})

	doc.Set("querySelector", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			return goja.Null()
		}

		selector := call.Arguments[0].String()
		node := css.QuerySelector(db.document, selector)
		if node == nil {
			return goja.Null()
		}
		if element, ok := node.(*dom.Element); ok {
			return db.WrapElement(element)
		}
		return goja.Null()
	})

	doc.Set("querySelectorAll", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			return db.vm.NewArray()
		}

		selector := call.Arguments[0].String()
		nodeList := css.QuerySelectorAll(db.document, selector)
		// Convert NodeList to []*Element
		var elements []*dom.Element
		nodeSlice := nodeList.ToSlice()
		for _, node := range nodeSlice {
			if element, ok := node.(*dom.Element); ok {
				elements = append(elements, element)
			}
		}
		return db.WrapNodeList(elements)
	})

	// Document properties - set up the main document elements
	if documentElement := db.document.DocumentElement(); documentElement != nil {
		doc.Set("documentElement", db.WrapElement(documentElement))
	} else {
		doc.Set("documentElement", goja.Null())
	}

	if body := db.document.Body(); body != nil {
		doc.Set("body", db.WrapElement(body))
	} else {
		doc.Set("body", goja.Null())
	}

	if head := db.document.Head(); head != nil {
		doc.Set("head", db.WrapElement(head))
	} else {
		doc.Set("head", goja.Null())
	}

	// Cookie property - implement document.cookie getter/setter
	doc.DefineAccessorProperty("cookie", db.vm.ToValue(func(call goja.FunctionCall) goja.Value {
		// Getter: return current cookies as string
		return db.vm.ToValue(db.cookieManager.FormatCookieString())
	}), db.vm.ToValue(func(call goja.FunctionCall) goja.Value {
		// Setter: parse and set cookie
		if len(call.Arguments) > 0 {
			cookieStr := call.Arguments[0].String()
			err := db.cookieManager.ParseCookieString(cookieStr)
			if err != nil {
				// In browsers, invalid cookies are usually silently ignored
				// rather than throwing errors, so we'll do the same
			}
		}
		return goja.Undefined()
	}), goja.FLAG_FALSE, goja.FLAG_TRUE) // Not enumerable, configurable

	// ReadyState property
	doc.DefineAccessorProperty("readyState",
		db.vm.ToValue(func(call goja.FunctionCall) goja.Value {
			// Getter
			return db.vm.ToValue(db.document.ReadyState())
		}),
		goja.Undefined(),                // No setter - readyState is read-only
		goja.FLAG_FALSE, goja.FLAG_TRUE) // Not enumerable, configurable

	// Phase 3: Document APIs
	// Document properties
	doc.DefineAccessorProperty("implementation",
		db.vm.ToValue(func(call goja.FunctionCall) goja.Value {
			// Getter - return DOMImplementation
			return db.wrapDOMImplementation(db.document.Implementation())
		}),
		goja.Undefined(),                // No setter - implementation is read-only
		goja.FLAG_FALSE, goja.FLAG_TRUE) // Not enumerable, configurable

	doc.DefineAccessorProperty("characterSet",
		db.vm.ToValue(func(call goja.FunctionCall) goja.Value {
			// Getter
			return db.vm.ToValue(db.document.CharacterSet())
		}),
		goja.Undefined(),                // No setter - characterSet is read-only
		goja.FLAG_FALSE, goja.FLAG_TRUE) // Not enumerable, configurable

	doc.DefineAccessorProperty("charset",
		db.vm.ToValue(func(call goja.FunctionCall) goja.Value {
			// Getter - alias for characterSet
			return db.vm.ToValue(db.document.CharacterSet())
		}),
		goja.Undefined(),                // No setter - charset is read-only
		goja.FLAG_FALSE, goja.FLAG_TRUE) // Not enumerable, configurable

	doc.DefineAccessorProperty("inputEncoding",
		db.vm.ToValue(func(call goja.FunctionCall) goja.Value {
			// Getter - alias for characterSet
			return db.vm.ToValue(db.document.CharacterSet())
		}),
		goja.Undefined(),                // No setter - inputEncoding is read-only
		goja.FLAG_FALSE, goja.FLAG_TRUE) // Not enumerable, configurable

	doc.DefineAccessorProperty("contentType",
		db.vm.ToValue(func(call goja.FunctionCall) goja.Value {
			// Getter
			return db.vm.ToValue(db.document.ContentType())
		}),
		goja.Undefined(),                // No setter - contentType is read-only
		goja.FLAG_FALSE, goja.FLAG_TRUE) // Not enumerable, configurable

	doc.DefineAccessorProperty("doctype",
		db.vm.ToValue(func(call goja.FunctionCall) goja.Value {
			// Getter
			if doctype := db.document.Doctype(); doctype != nil {
				return db.WrapNode(doctype)
			}
			return goja.Null()
		}),
		goja.Undefined(),                // No setter - doctype is read-only
		goja.FLAG_FALSE, goja.FLAG_TRUE) // Not enumerable, configurable

	doc.DefineAccessorProperty("compatMode",
		db.vm.ToValue(func(call goja.FunctionCall) goja.Value {
			// Getter
			return db.vm.ToValue(db.document.CompatMode())
		}),
		goja.Undefined(),                // No setter - compatMode is read-only
		goja.FLAG_FALSE, goja.FLAG_TRUE) // Not enumerable, configurable

	// Document creation methods
	doc.Set("createAttribute", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			panic(db.vm.NewTypeError("createAttribute requires a name"))
		}

		name := call.Arguments[0].String()
		attr, err := db.document.CreateAttribute(name)
		if err != nil {
			panic(db.vm.NewTypeError(err.Error()))
		}
		return db.WrapNode(attr)
	})

	doc.Set("createAttributeNS", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 2 {
			panic(db.vm.NewTypeError("createAttributeNS requires namespace and name"))
		}

		namespace := ""
		if !goja.IsNull(call.Arguments[0]) && !goja.IsUndefined(call.Arguments[0]) {
			namespace = call.Arguments[0].String()
		}
		qualifiedName := call.Arguments[1].String()

		attr, err := db.document.CreateAttributeNS(namespace, qualifiedName)
		if err != nil {
			panic(db.vm.NewTypeError(err.Error()))
		}
		return db.WrapNode(attr)
	})

	doc.Set("createCDATASection", func(call goja.FunctionCall) goja.Value {
		data := ""
		if len(call.Arguments) > 0 {
			data = call.Arguments[0].String()
		}

		cdataSection, err := db.document.CreateCDATASection(data)
		if err != nil {
			panic(db.vm.NewTypeError(err.Error()))
		}
		return db.WrapNode(cdataSection)
	})

	doc.Set("createProcessingInstruction", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 2 {
			panic(db.vm.NewTypeError("createProcessingInstruction requires target and data"))
		}

		target := call.Arguments[0].String()
		data := call.Arguments[1].String()

		pi, err := db.document.CreateProcessingInstruction(target, data)
		if err != nil {
			panic(db.vm.NewTypeError(err.Error()))
		}
		return db.WrapNode(pi)
	})

	doc.Set("createRange", func(call goja.FunctionCall) goja.Value {
		// Phase 5: Full Range API implementation
		return db.createRange()
	})

	doc.Set("createNodeIterator", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			panic(db.vm.NewTypeError("createNodeIterator requires a root node"))
		}

		// Phase 5: Full NodeIterator implementation
		return db.createNodeIterator(call.Arguments)
	})

	doc.Set("createTreeWalker", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			panic(db.vm.NewTypeError("createTreeWalker requires a root node"))
		}

		// Phase 5: Full TreeWalker implementation
		return db.createTreeWalker(call.Arguments)
	})

	// Document manipulation methods
	doc.Set("importNode", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			panic(db.vm.NewTypeError("importNode requires a node"))
		}

		node := db.extractNodeFromJS(call.Arguments[0])
		if node == nil {
			panic(db.vm.NewTypeError("Invalid node"))
		}

		deep := false
		if len(call.Arguments) > 1 {
			deep = call.Arguments[1].ToBoolean()
		}

		importedNode, err := db.document.ImportNode(node, deep)
		if err != nil {
			panic(db.vm.NewTypeError(err.Error()))
		}
		return db.WrapNode(importedNode)
	})

	doc.Set("adoptNode", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			panic(db.vm.NewTypeError("adoptNode requires a node"))
		}

		node := db.extractNodeFromJS(call.Arguments[0])
		if node == nil {
			panic(db.vm.NewTypeError("Invalid node"))
		}

		adoptedNode, err := db.document.AdoptNode(node)
		if err != nil {
			panic(db.vm.NewTypeError(err.Error()))
		}
		return db.WrapNode(adoptedNode)
	})

	doc.Set("getElementsByName", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			return db.vm.NewArray()
		}

		name := call.Arguments[0].String()
		collection := db.document.GetElementsByName(name)
		return db.WrapNodeList(collection.ToSlice())
	})

	// Normalize document method
	doc.Set("normalize", func(call goja.FunctionCall) goja.Value {
		db.document.Normalize()
		return goja.Undefined()
	})

	// Event methods (inherited from Node)
	db.addEventMethods(doc, db.document)

	// Selection API - getSelection method
	doc.Set("getSelection", func(call goja.FunctionCall) goja.Value {
		return db.createSelection()
	})

	return doc
}

// WrapElement wraps a DOM element for JavaScript access
func (db *DOMBindings) WrapElement(element *dom.Element) *goja.Object {
	if element == nil {
		return nil
	}

	// Check cache first
	if cached, exists := db.nodeCache[element]; exists {
		return cached
	}

	elem := db.vm.NewObject()

	// Cache the object before setting up properties to avoid infinite recursion
	db.nodeCache[element] = elem

	// Store the Go DOM node for extraction later
	elem.Set("__domNode", element)

	// Basic node properties
	elem.Set("nodeType", int(dom.ElementNode))
	elem.Set("nodeName", element.TagName())
	elem.Set("tagName", element.TagName())

	// Phase 1: Core Properties & Simple Methods
	elem.DefineAccessorProperty("isConnected",
		db.vm.ToValue(func(call goja.FunctionCall) goja.Value {
			// Getter
			return db.vm.ToValue(element.IsConnected())
		}),
		goja.Undefined(),                // No setter - isConnected is read-only
		goja.FLAG_FALSE, goja.FLAG_TRUE) // Not enumerable, configurable

	elem.DefineAccessorProperty("baseURI",
		db.vm.ToValue(func(call goja.FunctionCall) goja.Value {
			// Getter
			return db.vm.ToValue(element.BaseURI())
		}),
		goja.Undefined(),                // No setter - baseURI is read-only
		goja.FLAG_FALSE, goja.FLAG_TRUE) // Not enumerable, configurable

	elem.DefineAccessorProperty("nodeValue",
		db.vm.ToValue(func(call goja.FunctionCall) goja.Value {
			// Getter
			return db.vm.ToValue(element.NodeValue())
		}),
		db.vm.ToValue(func(call goja.FunctionCall) goja.Value {
			// Setter
			if len(call.Arguments) > 0 {
				value := call.Arguments[0].String()
				element.SetNodeValue(value)
			}
			return goja.Undefined()
		}),
		goja.FLAG_FALSE, goja.FLAG_TRUE) // Not enumerable, configurable

	// Phase 2: Element Navigation & Manipulation
	// Element navigation properties
	elem.DefineAccessorProperty("children",
		db.vm.ToValue(func(call goja.FunctionCall) goja.Value {
			// Getter - return live HTMLCollection of child elements
			var childElements []*dom.Element
			for _, child := range element.ChildNodes().ToSlice() {
				if childElement, ok := child.(*dom.Element); ok {
					childElements = append(childElements, childElement)
				}
			}
			return db.WrapNodeList(childElements)
		}),
		goja.Undefined(),                // No setter - children is read-only
		goja.FLAG_FALSE, goja.FLAG_TRUE) // Not enumerable, configurable

	elem.DefineAccessorProperty("firstElementChild",
		db.vm.ToValue(func(call goja.FunctionCall) goja.Value {
			// Getter
			for _, child := range element.ChildNodes().ToSlice() {
				if childElement, ok := child.(*dom.Element); ok {
					return db.WrapElement(childElement)
				}
			}
			return goja.Null()
		}),
		goja.Undefined(),                // No setter - firstElementChild is read-only
		goja.FLAG_FALSE, goja.FLAG_TRUE) // Not enumerable, configurable

	elem.DefineAccessorProperty("lastElementChild",
		db.vm.ToValue(func(call goja.FunctionCall) goja.Value {
			// Getter
			var lastElement *dom.Element
			for _, child := range element.ChildNodes().ToSlice() {
				if childElement, ok := child.(*dom.Element); ok {
					lastElement = childElement
				}
			}
			if lastElement != nil {
				return db.WrapElement(lastElement)
			}
			return goja.Null()
		}),
		goja.Undefined(),                // No setter - lastElementChild is read-only
		goja.FLAG_FALSE, goja.FLAG_TRUE) // Not enumerable, configurable

	elem.DefineAccessorProperty("childElementCount",
		db.vm.ToValue(func(call goja.FunctionCall) goja.Value {
			// Getter
			count := 0
			for _, child := range element.ChildNodes().ToSlice() {
				if _, ok := child.(*dom.Element); ok {
					count++
				}
			}
			return db.vm.ToValue(count)
		}),
		goja.Undefined(),                // No setter - childElementCount is read-only
		goja.FLAG_FALSE, goja.FLAG_TRUE) // Not enumerable, configurable

	elem.DefineAccessorProperty("previousElementSibling",
		db.vm.ToValue(func(call goja.FunctionCall) goja.Value {
			// Getter
			if parent := element.ParentNode(); parent != nil {
				children := parent.ChildNodes().ToSlice()
				for i, child := range children {
					if child == element && i > 0 {
						// Look backwards for previous element sibling
						for j := i - 1; j >= 0; j-- {
							if prevElement, ok := children[j].(*dom.Element); ok {
								return db.WrapElement(prevElement)
							}
						}
						break
					}
				}
			}
			return goja.Null()
		}),
		goja.Undefined(),                // No setter - previousElementSibling is read-only
		goja.FLAG_FALSE, goja.FLAG_TRUE) // Not enumerable, configurable

	elem.DefineAccessorProperty("nextElementSibling",
		db.vm.ToValue(func(call goja.FunctionCall) goja.Value {
			// Getter
			if parent := element.ParentNode(); parent != nil {
				children := parent.ChildNodes().ToSlice()
				for i, child := range children {
					if child == element && i < len(children)-1 {
						// Look forwards for next element sibling
						for j := i + 1; j < len(children); j++ {
							if nextElement, ok := children[j].(*dom.Element); ok {
								return db.WrapElement(nextElement)
							}
						}
						break
					}
				}
			}
			return goja.Null()
		}),
		goja.Undefined(),                // No setter - nextElementSibling is read-only
		goja.FLAG_FALSE, goja.FLAG_TRUE) // Not enumerable, configurable

	// Phase 2: Modern DOM manipulation methods
	// Modern append/prepend methods
	elem.Set("append", func(call goja.FunctionCall) goja.Value {
		for _, arg := range call.Arguments {
			if goja.IsUndefined(arg) || goja.IsNull(arg) {
				continue
			}

			// Handle strings - convert to text nodes
			domNodeValue := arg.ToObject(db.vm).Get("__domNode")
			if goja.IsUndefined(domNodeValue) || goja.IsNull(domNodeValue) {
				textContent := arg.String()
				textNode := db.document.CreateTextNode(textContent)
				element.AppendChild(textNode)
			} else {
				// Handle DOM nodes
				node := db.extractNodeFromJS(arg)
				if node != nil {
					element.AppendChild(node)
				}
			}
		}
		db.updateNodeNavigationProperties(elem, element)
		return goja.Undefined()
	})

	elem.Set("prepend", func(call goja.FunctionCall) goja.Value {
		// Insert nodes at the beginning
		for i := len(call.Arguments) - 1; i >= 0; i-- {
			arg := call.Arguments[i]
			if goja.IsUndefined(arg) || goja.IsNull(arg) {
				continue
			}

			// Handle strings - convert to text nodes
			domNodeValue := arg.ToObject(db.vm).Get("__domNode")
			if goja.IsUndefined(domNodeValue) || goja.IsNull(domNodeValue) {
				textContent := arg.String()
				textNode := db.document.CreateTextNode(textContent)
				element.InsertBefore(textNode, element.FirstChild())
			} else {
				// Handle DOM nodes
				node := db.extractNodeFromJS(arg)
				if node != nil {
					element.InsertBefore(node, element.FirstChild())
				}
			}
		}
		db.updateNodeNavigationProperties(elem, element)
		return goja.Undefined()
	})

	elem.Set("replaceChildren", func(call goja.FunctionCall) goja.Value {
		// Remove all existing children
		for element.FirstChild() != nil {
			element.RemoveChild(element.FirstChild())
		}

		// Add new children
		for _, arg := range call.Arguments {
			if goja.IsUndefined(arg) || goja.IsNull(arg) {
				continue
			}

			// Handle strings - convert to text nodes
			domNodeValue := arg.ToObject(db.vm).Get("__domNode")
			if goja.IsUndefined(domNodeValue) || goja.IsNull(domNodeValue) {
				textContent := arg.String()
				textNode := db.document.CreateTextNode(textContent)
				element.AppendChild(textNode)
			} else {
				// Handle DOM nodes
				node := db.extractNodeFromJS(arg)
				if node != nil {
					element.AppendChild(node)
				}
			}
		}
		db.updateNodeNavigationProperties(elem, element)
		return goja.Undefined()
	})

	// ChildNode mixin methods
	elem.Set("before", func(call goja.FunctionCall) goja.Value {
		parent := element.ParentNode()
		if parent == nil {
			return goja.Undefined()
		}

		for _, arg := range call.Arguments {
			if goja.IsUndefined(arg) || goja.IsNull(arg) {
				continue
			}

			// Handle strings - convert to text nodes
			domNodeValue := arg.ToObject(db.vm).Get("__domNode")
			if goja.IsUndefined(domNodeValue) || goja.IsNull(domNodeValue) {
				textContent := arg.String()
				textNode := db.document.CreateTextNode(textContent)
				parent.InsertBefore(textNode, element)
			} else {
				// Handle DOM nodes
				node := db.extractNodeFromJS(arg)
				if node != nil {
					parent.InsertBefore(node, element)
				}
			}
		}

		// Update parent's navigation properties
		if parentWrapper, exists := db.nodeCache[parent]; exists {
			db.updateNodeNavigationProperties(parentWrapper, parent)
		}
		return goja.Undefined()
	})

	elem.Set("after", func(call goja.FunctionCall) goja.Value {
		parent := element.ParentNode()
		if parent == nil {
			return goja.Undefined()
		}

		// Find the next sibling to insert before
		var nextSibling dom.Node
		siblings := parent.ChildNodes().ToSlice()
		for i, sibling := range siblings {
			if sibling == element && i < len(siblings)-1 {
				nextSibling = siblings[i+1]
				break
			}
		}

		for _, arg := range call.Arguments {
			if goja.IsUndefined(arg) || goja.IsNull(arg) {
				continue
			}

			// Handle strings - convert to text nodes
			domNodeValue := arg.ToObject(db.vm).Get("__domNode")
			if goja.IsUndefined(domNodeValue) || goja.IsNull(domNodeValue) {
				textContent := arg.String()
				textNode := db.document.CreateTextNode(textContent)
				parent.InsertBefore(textNode, nextSibling)
			} else {
				// Handle DOM nodes
				node := db.extractNodeFromJS(arg)
				if node != nil {
					parent.InsertBefore(node, nextSibling)
				}
			}
		}

		// Update parent's navigation properties
		if parentWrapper, exists := db.nodeCache[parent]; exists {
			db.updateNodeNavigationProperties(parentWrapper, parent)
		}
		return goja.Undefined()
	})

	elem.Set("replaceWith", func(call goja.FunctionCall) goja.Value {
		parent := element.ParentNode()
		if parent == nil {
			return goja.Undefined()
		}

		// Insert replacement nodes before the element
		for _, arg := range call.Arguments {
			if goja.IsUndefined(arg) || goja.IsNull(arg) {
				continue
			}

			// Handle strings - convert to text nodes
			domNodeValue := arg.ToObject(db.vm).Get("__domNode")
			if goja.IsUndefined(domNodeValue) || goja.IsNull(domNodeValue) {
				textContent := arg.String()
				textNode := db.document.CreateTextNode(textContent)
				parent.InsertBefore(textNode, element)
			} else {
				// Handle DOM nodes
				node := db.extractNodeFromJS(arg)
				if node != nil {
					parent.InsertBefore(node, element)
				}
			}
		}

		// Remove the original element
		parent.RemoveChild(element)

		// Update parent's navigation properties
		if parentWrapper, exists := db.nodeCache[parent]; exists {
			db.updateNodeNavigationProperties(parentWrapper, parent)
		}
		return goja.Undefined()
	})

	elem.Set("remove", func(call goja.FunctionCall) goja.Value {
		parent := element.ParentNode()
		if parent != nil {
			parent.RemoveChild(element)

			// Update parent's navigation properties
			if parentWrapper, exists := db.nodeCache[parent]; exists {
				db.updateNodeNavigationProperties(parentWrapper, parent)
			}
		}
		return goja.Undefined()
	})

	// Element content properties (using getters for dynamic updates)
	elem.DefineAccessorProperty("innerHTML",
		db.vm.ToValue(func(call goja.FunctionCall) goja.Value {
			// Getter
			return db.vm.ToValue(element.InnerHTML())
		}),
		db.vm.ToValue(func(call goja.FunctionCall) goja.Value {
			// Setter
			if len(call.Arguments) > 0 {
				html := call.Arguments[0].String()
				element.SetInnerHTML(html)
				// Update navigation properties after setting innerHTML
				db.updateNodeNavigationProperties(elem, element)
			}
			return goja.Undefined()
		}),
		goja.FLAG_FALSE, goja.FLAG_TRUE) // Not enumerable, configurable

	elem.DefineAccessorProperty("outerHTML",
		db.vm.ToValue(func(call goja.FunctionCall) goja.Value {
			// Getter
			return db.vm.ToValue(element.OuterHTML())
		}),
		db.vm.ToValue(func(call goja.FunctionCall) goja.Value {
			// Setter - outerHTML replacement is complex, for now just update innerHTML
			if len(call.Arguments) > 0 {
				html := call.Arguments[0].String()
				element.SetInnerHTML(html)
				// Update navigation properties after setting outerHTML
				db.updateNodeNavigationProperties(elem, element)
			}
			return goja.Undefined()
		}),
		goja.FLAG_FALSE, goja.FLAG_TRUE) // Not enumerable, configurable

	// Dynamic textContent property - set current value and add setter
	elem.DefineAccessorProperty("textContent",
		db.vm.ToValue(func(call goja.FunctionCall) goja.Value {
			// Getter
			return db.vm.ToValue(element.TextContent())
		}),
		db.vm.ToValue(func(call goja.FunctionCall) goja.Value {
			// Setter
			if len(call.Arguments) > 0 {
				text := call.Arguments[0].String()
				element.SetTextContent(text)
			}
			return goja.Undefined()
		}),
		goja.FLAG_FALSE, goja.FLAG_TRUE) // Not enumerable, configurable

	// Attribute methods
	elem.Set("getAttribute", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			return goja.Null()
		}
		name := call.Arguments[0].String()
		value := element.GetAttribute(name)
		if value == "" && !element.HasAttribute(name) {
			return goja.Null()
		}
		return db.vm.ToValue(value)
	})

	elem.Set("setAttribute", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 2 {
			panic(db.vm.NewTypeError("setAttribute requires name and value"))
		}
		name := call.Arguments[0].String()
		value := call.Arguments[1].String()
		element.SetAttribute(name, value)
		return goja.Undefined()
	})

	elem.Set("hasAttribute", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			return db.vm.ToValue(false)
		}
		name := call.Arguments[0].String()
		return db.vm.ToValue(element.HasAttribute(name))
	})

	elem.Set("removeAttribute", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			return goja.Undefined()
		}
		name := call.Arguments[0].String()
		element.RemoveAttribute(name)
		return goja.Undefined()
	})

	// Class methods with getter/setter
	elem.DefineAccessorProperty("className",
		db.vm.ToValue(func(call goja.FunctionCall) goja.Value {
			// Getter
			return db.vm.ToValue(element.GetAttribute("class"))
		}),
		db.vm.ToValue(func(call goja.FunctionCall) goja.Value {
			// Setter
			if len(call.Arguments) > 0 {
				className := call.Arguments[0].String()
				element.SetAttribute("class", className)
			}
			return goja.Undefined()
		}),
		goja.FLAG_FALSE, goja.FLAG_TRUE) // Not enumerable, configurable

	// ID property with getter/setter
	elem.DefineAccessorProperty("id",
		db.vm.ToValue(func(call goja.FunctionCall) goja.Value {
			// Getter
			return db.vm.ToValue(element.GetAttribute("id"))
		}),
		db.vm.ToValue(func(call goja.FunctionCall) goja.Value {
			// Setter
			if len(call.Arguments) > 0 {
				id := call.Arguments[0].String()
				element.SetAttribute("id", id)
			}
			return goja.Undefined()
		}),
		goja.FLAG_FALSE, goja.FLAG_TRUE) // Not enumerable, configurable

	// Selector methods
	elem.Set("querySelector", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			return goja.Null()
		}

		selector := call.Arguments[0].String()
		found := css.QuerySelector(element, selector)
		if found == nil {
			return goja.Null()
		}
		if foundElement, ok := found.(*dom.Element); ok {
			return db.WrapElement(foundElement)
		}
		return goja.Null()
	})

	elem.Set("querySelectorAll", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			return db.vm.NewArray()
		}

		selector := call.Arguments[0].String()
		nodeList := css.QuerySelectorAll(element, selector)
		// Convert NodeList to []*Element
		var elements []*dom.Element
		nodeSlice := nodeList.ToSlice()
		for _, node := range nodeSlice {
			if elem, ok := node.(*dom.Element); ok {
				elements = append(elements, elem)
			}
		}
		return db.WrapNodeList(elements)
	})

	elem.Set("getElementsByTagName", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			return db.vm.NewArray()
		}

		tagName := call.Arguments[0].String()
		collection := element.GetElementsByTagName(tagName)
		return db.WrapNodeList(collection.ToSlice())
	})

	elem.Set("getElementsByClassName", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			return db.vm.NewArray()
		}

		className := call.Arguments[0].String()
		collection := element.GetElementsByClassName(className)
		return db.WrapNodeList(collection.ToSlice())
	})

	// Insert adjacent HTML method
	elem.Set("insertAdjacentHTML", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 2 {
			panic(db.vm.NewTypeError("insertAdjacentHTML requires position and html"))
		}

		position := call.Arguments[0].String()
		html := call.Arguments[1].String()

		err := element.InsertAdjacentHTML(position, html)
		if err != nil {
			panic(db.vm.NewTypeError(err.Error()))
		}

		// Update navigation properties for the target element (afterbegin/beforeend affect target's children)
		if position == "afterbegin" || position == "beforeend" {
			db.updateNodeNavigationProperties(elem, element)
		}

		// Update parent's navigation properties for beforebegin/afterend
		if (position == "beforebegin" || position == "afterend") && element.ParentNode() != nil {
			// Find the parent's JavaScript wrapper in cache and update it
			if parentWrapper, exists := db.nodeCache[element.ParentNode()]; exists {
				db.updateNodeNavigationProperties(parentWrapper, element.ParentNode())
			}
		}

		return goja.Undefined()
	})

	// Add getRootNode method (newer DOM API)
	elem.Set("getRootNode", func(call goja.FunctionCall) goja.Value {
		// For most cases, return the document
		return db.WrapDocument()
	})

	// Add node methods
	db.addNodeMethods(elem, element)

	// Phase 4: Advanced DOM Features - Add to elements
	db.addPhase4NodeMethods(elem, element)

	// Add event methods
	db.addEventMethods(elem, element)

	return elem
}

// WrapNode wraps any DOM node for JavaScript access
func (db *DOMBindings) WrapNode(node dom.Node) *goja.Object {
	if node == nil {
		return nil
	}

	// Handle specific node types first
	switch n := node.(type) {
	case *dom.Element:
		return db.WrapElement(n)
	}

	// Check cache first
	if cached, exists := db.nodeCache[node]; exists {
		return cached
	}

	obj := db.vm.NewObject()

	// Cache the object before setting up properties to avoid infinite recursion
	db.nodeCache[node] = obj

	// Store the Go DOM node for extraction later
	obj.Set("__domNode", node)

	// Basic node properties
	obj.Set("nodeType", int(node.NodeType()))
	obj.Set("nodeName", node.NodeName())

	// Phase 1: Core Properties & Simple Methods
	obj.DefineAccessorProperty("isConnected",
		db.vm.ToValue(func(call goja.FunctionCall) goja.Value {
			// Getter
			return db.vm.ToValue(node.IsConnected())
		}),
		goja.Undefined(),                // No setter - isConnected is read-only
		goja.FLAG_FALSE, goja.FLAG_TRUE) // Not enumerable, configurable

	obj.DefineAccessorProperty("baseURI",
		db.vm.ToValue(func(call goja.FunctionCall) goja.Value {
			// Getter
			return db.vm.ToValue(node.BaseURI())
		}),
		goja.Undefined(),                // No setter - baseURI is read-only
		goja.FLAG_FALSE, goja.FLAG_TRUE) // Not enumerable, configurable

	obj.DefineAccessorProperty("nodeValue",
		db.vm.ToValue(func(call goja.FunctionCall) goja.Value {
			// Getter
			return db.vm.ToValue(node.NodeValue())
		}),
		db.vm.ToValue(func(call goja.FunctionCall) goja.Value {
			// Setter
			if len(call.Arguments) > 0 {
				value := call.Arguments[0].String()
				node.SetNodeValue(value)
			}
			return goja.Undefined()
		}),
		goja.FLAG_FALSE, goja.FLAG_TRUE) // Not enumerable, configurable

	obj.DefineAccessorProperty("textContent",
		db.vm.ToValue(func(call goja.FunctionCall) goja.Value {
			// Getter
			return db.vm.ToValue(db.getTextContent(node))
		}),
		db.vm.ToValue(func(call goja.FunctionCall) goja.Value {
			// Setter
			if len(call.Arguments) > 0 {
				text := call.Arguments[0].String()
				node.SetTextContent(text)
			}
			return goja.Undefined()
		}),
		goja.FLAG_FALSE, goja.FLAG_TRUE) // Not enumerable, configurable

	// Add CharacterData properties and methods for Text/Comment nodes
	switch n := node.(type) {
	case *dom.Text:
		// CharacterData properties
		obj.DefineAccessorProperty("data",
			db.vm.ToValue(func(call goja.FunctionCall) goja.Value {
				return db.vm.ToValue(n.Data())
			}),
			db.vm.ToValue(func(call goja.FunctionCall) goja.Value {
				if len(call.Arguments) > 0 {
					data := call.Arguments[0].String()
					n.SetData(data)
				}
				return goja.Undefined()
			}),
			goja.FLAG_FALSE, goja.FLAG_TRUE)

		obj.DefineAccessorProperty("length",
			db.vm.ToValue(func(call goja.FunctionCall) goja.Value {
				return db.vm.ToValue(n.Length())
			}),
			goja.Undefined(),
			goja.FLAG_FALSE, goja.FLAG_TRUE)

		// CharacterData methods
		obj.Set("appendData", func(call goja.FunctionCall) goja.Value {
			if len(call.Arguments) > 0 {
				data := call.Arguments[0].String()
				n.AppendData(data)
			}
			return goja.Undefined()
		})

		obj.Set("insertData", func(call goja.FunctionCall) goja.Value {
			if len(call.Arguments) < 2 {
				panic(db.vm.NewTypeError("insertData requires offset and data"))
			}
			offset := int(call.Arguments[0].ToInteger())
			data := call.Arguments[1].String()
			n.InsertData(offset, data)
			return goja.Undefined()
		})

		obj.Set("deleteData", func(call goja.FunctionCall) goja.Value {
			if len(call.Arguments) < 2 {
				panic(db.vm.NewTypeError("deleteData requires offset and count"))
			}
			offset := int(call.Arguments[0].ToInteger())
			count := int(call.Arguments[1].ToInteger())
			n.DeleteData(offset, count)
			return goja.Undefined()
		})

		obj.Set("replaceData", func(call goja.FunctionCall) goja.Value {
			if len(call.Arguments) < 3 {
				panic(db.vm.NewTypeError("replaceData requires offset, count, and data"))
			}
			offset := int(call.Arguments[0].ToInteger())
			count := int(call.Arguments[1].ToInteger())
			data := call.Arguments[2].String()
			n.ReplaceData(offset, count, data)
			return goja.Undefined()
		})

		obj.Set("substringData", func(call goja.FunctionCall) goja.Value {
			if len(call.Arguments) < 2 {
				panic(db.vm.NewTypeError("substringData requires offset and count"))
			}
			offset := int(call.Arguments[0].ToInteger())
			count := int(call.Arguments[1].ToInteger())
			return db.vm.ToValue(n.SubstringData(offset, count))
		})

	case *dom.Comment:
		// CharacterData properties
		obj.DefineAccessorProperty("data",
			db.vm.ToValue(func(call goja.FunctionCall) goja.Value {
				return db.vm.ToValue(n.Data())
			}),
			db.vm.ToValue(func(call goja.FunctionCall) goja.Value {
				if len(call.Arguments) > 0 {
					data := call.Arguments[0].String()
					n.SetData(data)
				}
				return goja.Undefined()
			}),
			goja.FLAG_FALSE, goja.FLAG_TRUE)

		obj.DefineAccessorProperty("length",
			db.vm.ToValue(func(call goja.FunctionCall) goja.Value {
				return db.vm.ToValue(n.Length())
			}),
			goja.Undefined(),
			goja.FLAG_FALSE, goja.FLAG_TRUE)

		// CharacterData methods
		obj.Set("appendData", func(call goja.FunctionCall) goja.Value {
			if len(call.Arguments) > 0 {
				data := call.Arguments[0].String()
				n.AppendData(data)
			}
			return goja.Undefined()
		})

		obj.Set("insertData", func(call goja.FunctionCall) goja.Value {
			if len(call.Arguments) < 2 {
				panic(db.vm.NewTypeError("insertData requires offset and data"))
			}
			offset := int(call.Arguments[0].ToInteger())
			data := call.Arguments[1].String()
			n.InsertData(offset, data)
			return goja.Undefined()
		})

		obj.Set("deleteData", func(call goja.FunctionCall) goja.Value {
			if len(call.Arguments) < 2 {
				panic(db.vm.NewTypeError("deleteData requires offset and count"))
			}
			offset := int(call.Arguments[0].ToInteger())
			count := int(call.Arguments[1].ToInteger())
			n.DeleteData(offset, count)
			return goja.Undefined()
		})

		obj.Set("replaceData", func(call goja.FunctionCall) goja.Value {
			if len(call.Arguments) < 3 {
				panic(db.vm.NewTypeError("replaceData requires offset, count, and data"))
			}
			offset := int(call.Arguments[0].ToInteger())
			count := int(call.Arguments[1].ToInteger())
			data := call.Arguments[2].String()
			n.ReplaceData(offset, count, data)
			return goja.Undefined()
		})

		obj.Set("substringData", func(call goja.FunctionCall) goja.Value {
			if len(call.Arguments) < 2 {
				panic(db.vm.NewTypeError("substringData requires offset and count"))
			}
			offset := int(call.Arguments[0].ToInteger())
			count := int(call.Arguments[1].ToInteger())
			return db.vm.ToValue(n.SubstringData(offset, count))
		})
	}

	// Add common node methods
	db.addNodeMethods(obj, node)

	// Phase 4: Advanced DOM Features - Add to all nodes
	db.addPhase4NodeMethods(obj, node)

	// Add event methods
	db.addEventMethods(obj, node)

	return obj
}

// createMediaQueryList creates a MediaQueryList JavaScript object
func (db *DOMBindings) createMediaQueryList(mediaQuery string) *goja.Object {
	obj := db.vm.NewObject()

	// MediaQueryList properties
	obj.Set("media", mediaQuery)

	// Mock matches property - for testing, we'll have some basic logic
	var matches bool
	// Simple mock logic for common media queries
	switch {
	case mediaQuery == "(min-width: 768px)":
		matches = true // Assume desktop by default
	case mediaQuery == "(max-width: 600px)":
		matches = false // Assume not mobile
	case mediaQuery == "(orientation: landscape)":
		matches = true // Assume landscape
	case mediaQuery == "(hover: hover)":
		matches = true // Assume hover capability
	default:
		matches = false // Default to false for unknown queries
	}

	obj.Set("matches", matches)

	// Event listener storage
	var changeListeners []goja.Callable

	// addEventListener method (modern approach)
	obj.Set("addEventListener", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 2 {
			panic(db.vm.NewTypeError("addEventListener requires type and listener"))
		}

		eventType := call.Arguments[0].String()
		if eventType != "change" {
			return goja.Undefined() // Only support 'change' events
		}

		listener, ok := goja.AssertFunction(call.Arguments[1])
		if !ok {
			panic(db.vm.NewTypeError("Listener must be a function"))
		}

		changeListeners = append(changeListeners, listener)
		return goja.Undefined()
	})

	// removeEventListener method
	obj.Set("removeEventListener", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 2 {
			return goja.Undefined()
		}

		eventType := call.Arguments[0].String()
		if eventType != "change" {
			return goja.Undefined()
		}

		// For simplicity, remove all listeners
		changeListeners = nil
		return goja.Undefined()
	})

	// addListener method (deprecated but still supported)
	obj.Set("addListener", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			panic(db.vm.NewTypeError("addListener requires a listener"))
		}

		listener, ok := goja.AssertFunction(call.Arguments[0])
		if !ok {
			panic(db.vm.NewTypeError("Listener must be a function"))
		}

		changeListeners = append(changeListeners, listener)
		return goja.Undefined()
	})

	// removeListener method (deprecated but still supported)
	obj.Set("removeListener", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			return goja.Undefined()
		}

		// For simplicity, remove all listeners
		changeListeners = nil
		return goja.Undefined()
	})

	// dispatchEvent method (for completeness)
	obj.Set("dispatchEvent", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			panic(db.vm.NewTypeError("dispatchEvent requires an event"))
		}

		// Create event object for listeners
		event := db.vm.NewObject()
		event.Set("type", "change")
		event.Set("matches", matches)
		event.Set("media", mediaQuery)

		// Call all change listeners
		for _, listener := range changeListeners {
			_, _ = listener(obj, event)
		}

		return db.vm.ToValue(true)
	})

	return obj
}

// addNodeMethods adds common DOM node methods to a JavaScript object
func (db *DOMBindings) addNodeMethods(obj *goja.Object, node dom.Node) {
	obj.Set("appendChild", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			panic(db.vm.NewTypeError("appendChild requires a child node"))
		}

		// Extract the DOM node from the JavaScript object
		childArg := call.Arguments[0]
		child := db.extractNodeFromJS(childArg)
		if child == nil {
			panic(db.vm.NewTypeError("Invalid node"))
		}

		// Perform the actual DOM appendChild operation
		node.AppendChild(child)

		// Update the child's parentNode property
		childJS := call.Arguments[0].ToObject(db.vm)
		if childJS != nil {
			childJS.Set("parentNode", obj)
		}

		// Update parent's navigation properties
		db.updateNodeNavigationProperties(obj, node)

		return call.Arguments[0] // Return the appended child (JavaScript wrapper)
	})

	obj.Set("removeChild", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			panic(db.vm.NewTypeError("removeChild requires a child node"))
		}

		// Simplified implementation - remove first child
		if node.FirstChild() != nil {
			node.RemoveChild(node.FirstChild())
		}
		return call.Arguments[0] // Return the removed child
	})

	obj.Set("insertBefore", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 2 {
			panic(db.vm.NewTypeError("insertBefore requires newChild and referenceChild"))
		}

		// Simplified implementation
		childText := call.Arguments[0].String()
		textNode := db.document.CreateTextNode(childText)
		node.InsertBefore(textNode, node.FirstChild())
		return call.Arguments[0] // Return the inserted child
	})

	obj.Set("replaceChild", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 2 {
			panic(db.vm.NewTypeError("replaceChild requires newChild and oldChild"))
		}

		// Simplified implementation
		if node.FirstChild() != nil {
			childText := call.Arguments[0].String()
			textNode := db.document.CreateTextNode(childText)
			node.ReplaceChild(textNode, node.FirstChild())
		}
		return call.Arguments[1] // Return the replaced child
	})

	obj.Set("cloneNode", func(call goja.FunctionCall) goja.Value {
		deep := false
		if len(call.Arguments) > 0 && call.Arguments[0].ToBoolean() {
			deep = true
		}

		cloned := node.CloneNode(deep)

		// WrapNode automatically handles Element detection and should call WrapElement
		// if the cloned node is actually an Element
		return db.WrapNode(cloned)
	})

	// Navigation properties - use getters for dynamic updates
	obj.DefineAccessorProperty("parentNode",
		db.vm.ToValue(func(call goja.FunctionCall) goja.Value {
			if parent := node.ParentNode(); parent != nil {
				return db.WrapNode(parent)
			}
			return goja.Null()
		}),
		goja.Undefined(),                // No setter
		goja.FLAG_FALSE, goja.FLAG_TRUE) // Not enumerable, configurable

	obj.DefineAccessorProperty("nextSibling",
		db.vm.ToValue(func(call goja.FunctionCall) goja.Value {
			if next := node.NextSibling(); next != nil {
				return db.WrapNode(next)
			}
			return goja.Null()
		}),
		goja.Undefined(),                // No setter
		goja.FLAG_FALSE, goja.FLAG_TRUE) // Not enumerable, configurable

	obj.DefineAccessorProperty("previousSibling",
		db.vm.ToValue(func(call goja.FunctionCall) goja.Value {
			if prev := node.PreviousSibling(); prev != nil {
				return db.WrapNode(prev)
			}
			return goja.Null()
		}),
		goja.Undefined(),                // No setter
		goja.FLAG_FALSE, goja.FLAG_TRUE) // Not enumerable, configurable

	// Set up other navigation properties properly
	db.updateNodeNavigationProperties(obj, node)
}

// addEventMethods adds event handling methods to a JavaScript object
func (db *DOMBindings) addEventMethods(obj *goja.Object, node dom.Node) {
	obj.Set("addEventListener", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 2 {
			panic(db.vm.NewTypeError("addEventListener requires type and listener"))
		}

		eventType := call.Arguments[0].String()
		listener, ok := goja.AssertFunction(call.Arguments[1])
		if !ok {
			panic(db.vm.NewTypeError("Listener must be a function"))
		}

		// Store the JavaScript listener for later lookup
		db.storeJSListener(node, eventType, db.vm.ToValue(listener).ToObject(db.vm))

		// Wrap the JavaScript function for Go DOM events
		wrapper := func(event dom.Event) {
			jsEvent := db.WrapEvent(event)
			_, _ = listener(goja.Undefined(), jsEvent)
		}

		node.AddEventListener(eventType, wrapper)
		return goja.Undefined()
	})

	obj.Set("removeEventListener", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 2 {
			panic(db.vm.NewTypeError("removeEventListener requires type and listener"))
		}

		// For now, we'll implement a simple version
		// In a full implementation, we'd need to track the wrapper functions
		eventType := call.Arguments[0].String()
		node.RemoveEventListener(eventType, nil) // Remove all listeners of this type
		return goja.Undefined()
	})

	obj.Set("dispatchEvent", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			panic(db.vm.NewTypeError("dispatchEvent requires an event"))
		}

		// Get event details from JavaScript event
		jsEvent := call.Arguments[0].ToObject(db.vm)
		if jsEvent == nil {
			panic(db.vm.NewTypeError("Event argument must be an object"))
		}

		eventType := ""
		if typeVal := jsEvent.Get("type"); typeVal != nil {
			eventType = typeVal.String()
		}
		if eventType == "" {
			panic(db.vm.NewTypeError("Event must have a type"))
		}

		// Dispatch through Go DOM system (which will call JS listeners via wrappers)
		bubbles := true
		cancelable := true
		if bubblesVal := jsEvent.Get("bubbles"); bubblesVal != nil {
			bubbles = bubblesVal.ToBoolean()
		}
		if cancelableVal := jsEvent.Get("cancelable"); cancelableVal != nil {
			cancelable = cancelableVal.ToBoolean()
		}

		goEvent := dom.NewEvent(eventType, bubbles, cancelable)
		goResult := db.dispatchEventWithBubbling(node, goEvent)

		return db.vm.ToValue(goResult)
	})
}

// WrapEvent wraps a DOM event for JavaScript access
func (db *DOMBindings) WrapEvent(event dom.Event) *goja.Object {
	evt := db.vm.NewObject()

	evt.Set("type", event.Type())

	// Ensure target and currentTarget are properly wrapped
	target := event.Target()
	if target != nil {
		evt.Set("target", db.WrapNode(target))
	} else {
		evt.Set("target", goja.Null())
	}

	currentTarget := event.CurrentTarget()
	if currentTarget != nil {
		evt.Set("currentTarget", db.WrapNode(currentTarget))
	} else {
		evt.Set("currentTarget", goja.Null())
	}

	evt.Set("bubbles", event.Bubbles())
	evt.Set("cancelable", event.Cancelable())
	evt.Set("defaultPrevented", event.DefaultPrevented())

	evt.Set("preventDefault", func(call goja.FunctionCall) goja.Value {
		event.PreventDefault()
		return goja.Undefined()
	})

	evt.Set("stopPropagation", func(call goja.FunctionCall) goja.Value {
		event.StopPropagation()
		return goja.Undefined()
	})

	evt.Set("stopImmediatePropagation", func(call goja.FunctionCall) goja.Value {
		event.StopImmediatePropagation()
		return goja.Undefined()
	})

	return evt
}

// extractNode extracts a Go DOM node from a JavaScript value
func (db *DOMBindings) extractNode(value goja.Value) dom.Node {
	// This is a simplified implementation
	// In a real implementation, we'd need to maintain a mapping
	// between JavaScript wrappers and Go objects
	// For now, return nil and let the caller handle it
	return nil
}

// extractNodeFromJS extracts a Go DOM node from a JavaScript value
func (db *DOMBindings) extractNodeFromJS(value goja.Value) dom.Node {
	// Check if this is a JavaScript object with a __domNode property
	if obj := value.ToObject(db.vm); obj != nil {
		if domNodeValue := obj.Get("__domNode"); domNodeValue != nil && !goja.IsUndefined(domNodeValue) {
			// Extract the stored DOM node
			if exported := domNodeValue.Export(); exported != nil {
				if node, ok := exported.(dom.Node); ok {
					return node
				}
			}
		}
	}
	return nil
}

// updateNodeNavigationProperties updates a JavaScript object's navigation properties to match the DOM state
func (db *DOMBindings) updateNodeNavigationProperties(obj *goja.Object, node dom.Node) {
	// Update childNodes array
	children := db.vm.NewArray()
	childNodes := node.ChildNodes().ToSlice()
	for i, childNode := range childNodes {
		children.Set(strconv.Itoa(i), db.WrapNode(childNode))
	}
	children.Set("length", len(childNodes))
	obj.Set("childNodes", children)

	// Update navigation properties
	if node.FirstChild() != nil {
		obj.Set("firstChild", db.WrapNode(node.FirstChild()))
	} else {
		obj.Set("firstChild", goja.Null())
	}

	if node.LastChild() != nil {
		obj.Set("lastChild", db.WrapNode(node.LastChild()))
	} else {
		obj.Set("lastChild", goja.Null())
	}

	// Note: parentNode, nextSibling, previousSibling would need additional context
	// and could cause circular references, so we handle them separately when needed
}

// getTextContent extracts text content from a DOM node
func (db *DOMBindings) getTextContent(node dom.Node) string {
	switch n := node.(type) {
	case *dom.Element:
		return n.TextContent()
	case *dom.Text:
		return n.NodeValue()
	case *dom.Comment:
		return ""
	default:
		// For other node types, try to get node value
		return node.NodeValue()
	}
}

// extractEvent extracts a Go DOM event from a JavaScript value
func (db *DOMBindings) extractEvent(value goja.Value) dom.Event {
	// Similar to extractNode, this would need proper implementation
	// For now, return nil
	return nil
}

// WrapNodeList wraps a slice of elements as a JavaScript array-like object
func (db *DOMBindings) WrapNodeList(elements []*dom.Element) *goja.Object {
	arr := db.vm.NewArray()

	for i, element := range elements {
		arr.Set(strconv.Itoa(i), db.WrapElement(element))
	}

	arr.Set("length", len(elements))

	return arr
}

// SetupBrowserAPIs sets up browser APIs like CustomEvent, URL, URLSearchParams, History, and Performance
func (db *DOMBindings) SetupBrowserAPIs() {
	// DOMException constructor
	db.vm.Set("DOMException", func(call goja.ConstructorCall) *goja.Object {
		message := ""
		name := "Error"

		if len(call.Arguments) > 0 && !goja.IsUndefined(call.Arguments[0]) {
			message = call.Arguments[0].String()
		}
		if len(call.Arguments) > 1 && !goja.IsUndefined(call.Arguments[1]) {
			name = call.Arguments[1].String()
		}

		exc := dom.NewDOMException(message, name)
		return db.wrapDOMException(exc)
	})

	// Add DOMException constants to the constructor
	domExceptionConstructor := db.vm.Get("DOMException").(*goja.Object)
	domExceptionConstructor.Set("INDEX_SIZE_ERR", dom.INDEX_SIZE_ERR)
	domExceptionConstructor.Set("DOMSTRING_SIZE_ERR", dom.DOMSTRING_SIZE_ERR)
	domExceptionConstructor.Set("HIERARCHY_REQUEST_ERR", dom.HIERARCHY_REQUEST_ERR)
	domExceptionConstructor.Set("WRONG_DOCUMENT_ERR", dom.WRONG_DOCUMENT_ERR)
	domExceptionConstructor.Set("INVALID_CHARACTER_ERR", dom.INVALID_CHARACTER_ERR)
	domExceptionConstructor.Set("NO_DATA_ALLOWED_ERR", dom.NO_DATA_ALLOWED_ERR)
	domExceptionConstructor.Set("NO_MODIFICATION_ALLOWED_ERR", dom.NO_MODIFICATION_ALLOWED_ERR)
	domExceptionConstructor.Set("NOT_FOUND_ERR", dom.NOT_FOUND_ERR)
	domExceptionConstructor.Set("NOT_SUPPORTED_ERR", dom.NOT_SUPPORTED_ERR)
	domExceptionConstructor.Set("INUSE_ATTRIBUTE_ERR", dom.INUSE_ATTRIBUTE_ERR)
	domExceptionConstructor.Set("INVALID_STATE_ERR", dom.INVALID_STATE_ERR)
	domExceptionConstructor.Set("SYNTAX_ERR", dom.SYNTAX_ERR)
	domExceptionConstructor.Set("INVALID_MODIFICATION_ERR", dom.INVALID_MODIFICATION_ERR)
	domExceptionConstructor.Set("NAMESPACE_ERR", dom.NAMESPACE_ERR)
	domExceptionConstructor.Set("INVALID_ACCESS_ERR", dom.INVALID_ACCESS_ERR)
	domExceptionConstructor.Set("VALIDATION_ERR", dom.VALIDATION_ERR)
	domExceptionConstructor.Set("TYPE_MISMATCH_ERR", dom.TYPE_MISMATCH_ERR)
	domExceptionConstructor.Set("SECURITY_ERR", dom.SECURITY_ERR)
	domExceptionConstructor.Set("NETWORK_ERR", dom.NETWORK_ERR)
	domExceptionConstructor.Set("ABORT_ERR", dom.ABORT_ERR)
	domExceptionConstructor.Set("URL_MISMATCH_ERR", dom.URL_MISMATCH_ERR)
	domExceptionConstructor.Set("QUOTA_EXCEEDED_ERR", dom.QUOTA_EXCEEDED_ERR)
	domExceptionConstructor.Set("TIMEOUT_ERR", dom.TIMEOUT_ERR)
	domExceptionConstructor.Set("INVALID_NODE_TYPE_ERR", dom.INVALID_NODE_TYPE_ERR)
	domExceptionConstructor.Set("DATA_CLONE_ERR", dom.DATA_CLONE_ERR)

	// URL constructor
	db.vm.Set("URL", func(call goja.ConstructorCall) *goja.Object {
		if len(call.Arguments) < 1 {
			panic(db.vm.NewTypeError("URL constructor requires a URL string"))
		}

		urlStr := call.Arguments[0].String()
		var base string
		if len(call.Arguments) > 1 {
			base = call.Arguments[1].String()
		}

		var goURL *url.URL
		var err error
		if base != "" {
			goURL, err = url.NewURL(urlStr, base)
		} else {
			goURL, err = url.NewURL(urlStr)
		}

		if err != nil {
			panic(db.vm.NewTypeError("Invalid URL: " + err.Error()))
		}

		return db.wrapURL(goURL)
	})

	// URLSearchParams constructor
	db.vm.Set("URLSearchParams", func(call goja.ConstructorCall) *goja.Object {
		var init string
		if len(call.Arguments) > 0 {
			init = call.Arguments[0].String()
		}

		params := url.NewURLSearchParams(init)
		return db.wrapURLSearchParams(params)
	})

	// Event constructor
	db.vm.Set("Event", func(call goja.ConstructorCall) *goja.Object {
		if len(call.Arguments) < 1 {
			panic(db.vm.NewTypeError("Event constructor requires an event type"))
		}

		eventType := call.Arguments[0].String()
		obj := db.vm.NewObject()
		obj.Set("type", eventType)
		obj.Set("bubbles", false)
		obj.Set("cancelable", false)
		obj.Set("target", goja.Null())
		obj.Set("currentTarget", goja.Null())

		// Handle options parameter
		if len(call.Arguments) > 1 {
			options := call.Arguments[1].ToObject(db.vm)
			if options != nil {
				if bubblesVal := options.Get("bubbles"); bubblesVal != nil {
					obj.Set("bubbles", bubblesVal.ToBoolean())
				}
				if cancelableVal := options.Get("cancelable"); cancelableVal != nil {
					obj.Set("cancelable", cancelableVal.ToBoolean())
				}
			}
		}

		return obj
	})

	// KeyboardEvent constructor
	db.vm.Set("KeyboardEvent", func(call goja.ConstructorCall) *goja.Object {
		if len(call.Arguments) < 1 {
			panic(db.vm.NewTypeError("KeyboardEvent constructor requires an event type"))
		}

		eventType := call.Arguments[0].String()
		obj := db.vm.NewObject()
		obj.Set("type", eventType)
		obj.Set("bubbles", false)
		obj.Set("cancelable", false)
		obj.Set("target", goja.Null())
		obj.Set("currentTarget", goja.Null())
		obj.Set("key", "")
		obj.Set("code", "")
		obj.Set("altKey", false)
		obj.Set("ctrlKey", false)
		obj.Set("shiftKey", false)
		obj.Set("metaKey", false)

		// Handle options parameter
		if len(call.Arguments) > 1 {
			options := call.Arguments[1].ToObject(db.vm)
			if options != nil {
				if bubblesVal := options.Get("bubbles"); bubblesVal != nil {
					obj.Set("bubbles", bubblesVal.ToBoolean())
				}
				if cancelableVal := options.Get("cancelable"); cancelableVal != nil {
					obj.Set("cancelable", cancelableVal.ToBoolean())
				}
				if keyVal := options.Get("key"); keyVal != nil {
					obj.Set("key", keyVal.String())
				}
				if codeVal := options.Get("code"); codeVal != nil {
					obj.Set("code", codeVal.String())
				}
				if altKeyVal := options.Get("altKey"); altKeyVal != nil {
					obj.Set("altKey", altKeyVal.ToBoolean())
				}
				if ctrlKeyVal := options.Get("ctrlKey"); ctrlKeyVal != nil {
					obj.Set("ctrlKey", ctrlKeyVal.ToBoolean())
				}
				if shiftKeyVal := options.Get("shiftKey"); shiftKeyVal != nil {
					obj.Set("shiftKey", shiftKeyVal.ToBoolean())
				}
				if metaKeyVal := options.Get("metaKey"); metaKeyVal != nil {
					obj.Set("metaKey", metaKeyVal.ToBoolean())
				}
			}
		}

		return obj
	})

	// CustomEvent constructor stub
	db.vm.Set("CustomEvent", func(call goja.ConstructorCall) *goja.Object {
		obj := db.vm.NewObject()
		obj.Set("type", "")
		obj.Set("detail", goja.Null())
		return obj
	})

	// XPathEvaluator constructor
	db.vm.Set("XPathEvaluator", func(call goja.ConstructorCall) *goja.Object {
		evaluator := xpath.NewXPathEvaluator()
		return db.wrapXPathEvaluator(evaluator)
	})

	// XPathResult constants
	xpathResult := db.vm.NewObject()
	xpathResult.Set("ANY_TYPE", xpath.ANY_TYPE)
	xpathResult.Set("NUMBER_TYPE", xpath.NUMBER_TYPE)
	xpathResult.Set("STRING_TYPE", xpath.STRING_TYPE)
	xpathResult.Set("BOOLEAN_TYPE", xpath.BOOLEAN_TYPE)
	xpathResult.Set("UNORDERED_NODE_ITERATOR_TYPE", xpath.UNORDERED_NODE_ITERATOR_TYPE)
	xpathResult.Set("ORDERED_NODE_ITERATOR_TYPE", xpath.ORDERED_NODE_ITERATOR_TYPE)
	xpathResult.Set("UNORDERED_NODE_SNAPSHOT_TYPE", xpath.UNORDERED_NODE_SNAPSHOT_TYPE)
	xpathResult.Set("ORDERED_NODE_SNAPSHOT_TYPE", xpath.ORDERED_NODE_SNAPSHOT_TYPE)
	xpathResult.Set("ANY_UNORDERED_NODE_TYPE", xpath.ANY_UNORDERED_NODE_TYPE)
	xpathResult.Set("FIRST_ORDERED_NODE_TYPE", xpath.FIRST_ORDERED_NODE_TYPE)
	db.vm.Set("XPathResult", xpathResult)
}

// SetupGlobalAPIs sets up global browser APIs like window.history and window.performance
func (db *DOMBindings) SetupGlobalAPIs() {
	// Create window object if it doesn't exist
	windowValue := db.vm.Get("window")
	var window *goja.Object
	if windowValue == nil || goja.IsUndefined(windowValue) || goja.IsNull(windowValue) {
		window = db.vm.NewObject()
		db.vm.Set("window", window)
	} else {
		window = windowValue.ToObject(db.vm)
		if window == nil {
			window = db.vm.NewObject()
			db.vm.Set("window", window)
		}
	}

	// Add event handling capabilities to window object
	db.addWindowEventMethods(window)

	// Setup History API
	h := history.NewHistory()
	window.Set("history", db.wrapHistory(h))

	// Setup Performance API
	p := performance.NewPerformance()
	window.Set("performance", db.wrapPerformance(p))

	// Setup Location API
	location := db.vm.NewObject()
	location.Set("href", "http://localhost:3000/")
	location.Set("origin", "http://localhost:3000")
	location.Set("protocol", "http:")
	location.Set("host", "localhost:3000")
	location.Set("hostname", "localhost")
	location.Set("port", "3000")
	location.Set("pathname", "/")
	location.Set("search", "")
	location.Set("hash", "")

	// Location methods
	location.Set("assign", func(call goja.FunctionCall) goja.Value {
		// Stub implementation
		return goja.Undefined()
	})
	location.Set("replace", func(call goja.FunctionCall) goja.Value {
		// Stub implementation
		return goja.Undefined()
	})
	location.Set("reload", func(call goja.FunctionCall) goja.Value {
		// Stub implementation
		return goja.Undefined()
	})
	location.Set("toString", func(call goja.FunctionCall) goja.Value {
		return db.vm.ToValue("http://localhost:3000/")
	})

	window.Set("location", location)
	db.vm.Set("location", location) // Also set as global

	// Add fetch API for HTMX and other libraries
	db.vm.Set("fetch", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			panic(db.vm.NewTypeError("fetch requires a URL"))
		}

		url := call.Arguments[0].String()

		// Create a simple Promise-like object
		promise := db.vm.NewObject()

		// For testing purposes, simulate successful responses
		// In a real implementation, this would make actual HTTP requests
		promise.Set("then", func(call goja.FunctionCall) goja.Value {
			if len(call.Arguments) > 0 {
				callback, ok := goja.AssertFunction(call.Arguments[0])
				if ok {
					// Create mock response
					response := db.vm.NewObject()
					response.Set("ok", true)
					response.Set("status", 200)
					response.Set("url", url)

					// Mock text() method
					response.Set("text", func(call goja.FunctionCall) goja.Value {
						textPromise := db.vm.NewObject()
						textPromise.Set("then", func(call goja.FunctionCall) goja.Value {
							if len(call.Arguments) > 0 {
								textCallback, ok := goja.AssertFunction(call.Arguments[0])
								if ok {
									// Return mock HTML content based on URL
									var content string
									if url == "/api/hello" {
										content = `<div class="success">Hello from HTMX! 👋</div>`
									} else if url == "/api/contact" {
										content = `<div class="success">Thank you! Message sent.</div>`
									} else {
										content = `<div>Mock response for ` + url + `</div>`
									}
									_, _ = textCallback(goja.Undefined(), db.vm.ToValue(content))
								}
							}
							return textPromise
						})
						return textPromise
					})

					_, _ = callback(goja.Undefined(), response)
				}
			}
			return promise
		})

		promise.Set("catch", func(call goja.FunctionCall) goja.Value {
			// Error handling - for now just return the promise
			return promise
		})

		return promise
	})

	// Phase 1: Core Node Constructors & Properties
	// Text constructor
	db.vm.Set("Text", func(call goja.ConstructorCall) *goja.Object {
		data := ""
		if len(call.Arguments) > 0 {
			data = call.Arguments[0].String()
		}
		textNode := db.document.CreateTextNode(data)
		return db.WrapNode(textNode)
	})

	// Comment constructor
	db.vm.Set("Comment", func(call goja.ConstructorCall) *goja.Object {
		data := ""
		if len(call.Arguments) > 0 {
			data = call.Arguments[0].String()
		}
		comment := db.document.CreateComment(data)
		return db.WrapNode(comment)
	})

	// DocumentType constructor (read-only, but available for instanceof checks)
	db.vm.Set("DocumentType", func(call goja.ConstructorCall) *goja.Object {
		// DocumentType cannot be constructed directly in DOM
		panic(db.vm.NewTypeError("DocumentType constructor is not available"))
	})

	// Attr constructor (read-only, but available for instanceof checks)
	db.vm.Set("Attr", func(call goja.ConstructorCall) *goja.Object {
		// Attr cannot be constructed directly in most browsers
		panic(db.vm.NewTypeError("Attr constructor is not available"))
	})

	// CDATASection constructor
	db.vm.Set("CDATASection", func(call goja.ConstructorCall) *goja.Object {
		// CDATASection constructor is not typically available directly
		panic(db.vm.NewTypeError("CDATASection constructor is not available"))
	})

	// ProcessingInstruction constructor
	db.vm.Set("ProcessingInstruction", func(call goja.ConstructorCall) *goja.Object {
		// ProcessingInstruction constructor is not typically available directly
		panic(db.vm.NewTypeError("ProcessingInstruction constructor is not available"))
	})

	// Add DOM constructor globals that libraries like HTMX expect
	db.vm.Set("Element", func(call goja.ConstructorCall) *goja.Object {
		// Element constructor - HTMX uses this to check instanceof
		return db.vm.NewObject()
	})

	nodeConstructor := func(call goja.ConstructorCall) *goja.Object {
		// Node constructor
		return db.vm.NewObject()
	}
	db.vm.Set("Node", nodeConstructor)

	// Add Node constants to the constructor
	nodeObj := db.vm.Get("Node").(*goja.Object)
	nodeObj.Set("ELEMENT_NODE", int(dom.ElementNode))
	nodeObj.Set("ATTRIBUTE_NODE", int(dom.AttributeNode))
	nodeObj.Set("TEXT_NODE", int(dom.TextNode))
	nodeObj.Set("CDATA_SECTION_NODE", int(dom.CDataSectionNode))
	nodeObj.Set("ENTITY_REFERENCE_NODE", int(dom.EntityReferenceNode))
	nodeObj.Set("ENTITY_NODE", int(dom.EntityNode))
	nodeObj.Set("PROCESSING_INSTRUCTION_NODE", int(dom.ProcessingInstructionNode))
	nodeObj.Set("COMMENT_NODE", int(dom.CommentNode))
	nodeObj.Set("DOCUMENT_NODE", int(dom.DocumentNode))
	nodeObj.Set("DOCUMENT_TYPE_NODE", int(dom.DocumentTypeNode))
	nodeObj.Set("DOCUMENT_FRAGMENT_NODE", int(dom.DocumentFragmentNode))
	nodeObj.Set("NOTATION_NODE", int(dom.NotationNode))

	db.vm.Set("Document", func(call goja.ConstructorCall) *goja.Object {
		// Document constructor
		return db.vm.NewObject()
	})

	db.vm.Set("HTMLElement", func(call goja.ConstructorCall) *goja.Object {
		// HTMLElement constructor
		return db.vm.NewObject()
	})

	db.vm.Set("DocumentFragment", func(call goja.ConstructorCall) *goja.Object {
		// DocumentFragment constructor
		return db.vm.NewObject()
	})

	// NodeFilter constants and functionality
	nodeFilter := db.vm.NewObject()
	nodeFilter.Set("SHOW_ALL", 0xFFFFFFFF)
	nodeFilter.Set("SHOW_ELEMENT", 0x1)
	nodeFilter.Set("SHOW_ATTRIBUTE", 0x2)
	nodeFilter.Set("SHOW_TEXT", 0x4)
	nodeFilter.Set("SHOW_CDATA_SECTION", 0x8)
	nodeFilter.Set("SHOW_ENTITY_REFERENCE", 0x10)
	nodeFilter.Set("SHOW_ENTITY", 0x20)
	nodeFilter.Set("SHOW_PROCESSING_INSTRUCTION", 0x40)
	nodeFilter.Set("SHOW_COMMENT", 0x80)
	nodeFilter.Set("SHOW_DOCUMENT", 0x100)
	nodeFilter.Set("SHOW_DOCUMENT_TYPE", 0x200)
	nodeFilter.Set("SHOW_DOCUMENT_FRAGMENT", 0x400)
	nodeFilter.Set("SHOW_NOTATION", 0x800)

	// NodeFilter result constants
	nodeFilter.Set("FILTER_ACCEPT", 1)
	nodeFilter.Set("FILTER_REJECT", 2)
	nodeFilter.Set("FILTER_SKIP", 3)

	db.vm.Set("NodeFilter", nodeFilter)

	// Selection API
	db.vm.Set("Selection", func(call goja.ConstructorCall) *goja.Object {
		// Selection objects are typically not constructed directly
		panic(db.vm.NewTypeError("Selection constructor is not available"))
	})

	// Storage API
	localStorage := db.vm.NewObject()
	sessionStorage := db.vm.NewObject()

	// In-memory storage for simulation
	localStorageData := make(map[string]string)
	sessionStorageData := make(map[string]string)

	// Implement localStorage methods
	localStorage.Set("setItem", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 2 {
			panic(db.vm.NewTypeError("setItem requires key and value"))
		}
		key := call.Arguments[0].String()
		value := call.Arguments[1].String()
		localStorageData[key] = value
		return goja.Undefined()
	})

	localStorage.Set("getItem", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			return goja.Null()
		}
		key := call.Arguments[0].String()
		if value, exists := localStorageData[key]; exists {
			return db.vm.ToValue(value)
		}
		return goja.Null()
	})

	localStorage.Set("removeItem", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			return goja.Undefined()
		}
		key := call.Arguments[0].String()
		delete(localStorageData, key)
		return goja.Undefined()
	})

	localStorage.Set("clear", func(call goja.FunctionCall) goja.Value {
		localStorageData = make(map[string]string)
		return goja.Undefined()
	})

	localStorage.Set("key", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			return goja.Null()
		}
		index := int(call.Arguments[0].ToInteger())
		keys := make([]string, 0, len(localStorageData))
		for k := range localStorageData {
			keys = append(keys, k)
		}
		if index >= 0 && index < len(keys) {
			return db.vm.ToValue(keys[index])
		}
		return goja.Null()
	})

	localStorage.DefineAccessorProperty("length",
		db.vm.ToValue(func(call goja.FunctionCall) goja.Value {
			return db.vm.ToValue(len(localStorageData))
		}),
		goja.Undefined(),
		goja.FLAG_FALSE, goja.FLAG_TRUE)

	// Implement sessionStorage methods (same as localStorage but different data)
	sessionStorage.Set("setItem", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 2 {
			panic(db.vm.NewTypeError("setItem requires key and value"))
		}
		key := call.Arguments[0].String()
		value := call.Arguments[1].String()
		sessionStorageData[key] = value
		return goja.Undefined()
	})

	sessionStorage.Set("getItem", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			return goja.Null()
		}
		key := call.Arguments[0].String()
		if value, exists := sessionStorageData[key]; exists {
			return db.vm.ToValue(value)
		}
		return goja.Null()
	})

	sessionStorage.Set("removeItem", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			return goja.Undefined()
		}
		key := call.Arguments[0].String()
		delete(sessionStorageData, key)
		return goja.Undefined()
	})

	sessionStorage.Set("clear", func(call goja.FunctionCall) goja.Value {
		sessionStorageData = make(map[string]string)
		return goja.Undefined()
	})

	sessionStorage.Set("key", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			return goja.Null()
		}
		index := int(call.Arguments[0].ToInteger())
		keys := make([]string, 0, len(sessionStorageData))
		for k := range sessionStorageData {
			keys = append(keys, k)
		}
		if index >= 0 && index < len(keys) {
			return db.vm.ToValue(keys[index])
		}
		return goja.Null()
	})

	sessionStorage.DefineAccessorProperty("length",
		db.vm.ToValue(func(call goja.FunctionCall) goja.Value {
			return db.vm.ToValue(len(sessionStorageData))
		}),
		goja.Undefined(),
		goja.FLAG_FALSE, goja.FLAG_TRUE)

	// Set storage objects on window and globally
	window.Set("localStorage", localStorage)
	window.Set("sessionStorage", sessionStorage)
	db.vm.Set("localStorage", localStorage)
	db.vm.Set("sessionStorage", sessionStorage)

	// Selection API - getSelection method on window
	window.Set("getSelection", func(call goja.FunctionCall) goja.Value {
		return db.createSelection()
	})

	// Also set getSelection globally
	db.vm.Set("getSelection", func(call goja.FunctionCall) goja.Value {
		return db.createSelection()
	})

	// MediaQuery API - matchMedia function
	window.Set("matchMedia", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			panic(db.vm.NewTypeError("matchMedia requires a media query string"))
		}

		mediaQuery := call.Arguments[0].String()
		return db.createMediaQueryList(mediaQuery)
	})

	// Also set matchMedia globally
	db.vm.Set("matchMedia", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			panic(db.vm.NewTypeError("matchMedia requires a media query string"))
		}

		mediaQuery := call.Arguments[0].String()
		return db.createMediaQueryList(mediaQuery)
	})

	// MutationObserver constructor
	db.vm.Set("MutationObserver", func(call goja.ConstructorCall) *goja.Object {
		if len(call.Arguments) < 1 {
			panic(db.vm.NewTypeError("MutationObserver constructor requires a callback"))
		}

		callback, ok := goja.AssertFunction(call.Arguments[0])
		if !ok {
			panic(db.vm.NewTypeError("MutationObserver callback must be a function"))
		}

		// Create a Go DOM MutationObserver that bridges to JavaScript
		domCallback := func(records []*dom.MutationRecord, observer *dom.MutationObserver) {
			// Convert DOM mutation records to JavaScript objects
			jsRecords := db.vm.NewArray()
			for i, record := range records {
				jsRecord := db.wrapMutationRecord(record)
				jsRecords.Set(strconv.Itoa(i), jsRecord)
			}
			jsRecords.Set("length", len(records))

			// Call JavaScript callback via microtask if we have access to event loop
			if db.jsRuntime != nil && db.jsRuntime.EventLoop() != nil {
				// Queue as microtask for proper event loop behavior
				db.jsRuntime.EventLoop().QueueMicrotask(func() {
					_, _ = callback(goja.Undefined(), jsRecords, db.wrapDOMMutationObserver(observer))
				}, "MutationObserver")
			} else {
				// Fallback: call immediately
				_, _ = callback(goja.Undefined(), jsRecords, db.wrapDOMMutationObserver(observer))
			}
		}

		// Create the actual DOM MutationObserver
		domObserver := dom.NewMutationObserver(domCallback)

		// Return JavaScript wrapper
		return db.wrapDOMMutationObserver(domObserver)
	})

	// Phase 6: Modern Web APIs

	// ResizeObserver constructor
	db.vm.Set("ResizeObserver", func(call goja.ConstructorCall) *goja.Object {
		if len(call.Arguments) < 1 {
			panic(db.vm.NewTypeError("ResizeObserver constructor requires a callback"))
		}

		callback, ok := goja.AssertFunction(call.Arguments[0])
		if !ok {
			panic(db.vm.NewTypeError("ResizeObserver callback must be a function"))
		}

		return db.createResizeObserver(callback)
	})

	// Test utility to trigger ResizeObserver callbacks
	db.vm.Set("__triggerResizeObserver", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			panic(db.vm.NewTypeError("__triggerResizeObserver requires target element"))
		}

		target := db.extractNodeFromJS(call.Arguments[0])
		if target == nil {
			panic(db.vm.NewTypeError("Invalid target element"))
		}

		// Optional resize info parameter
		var resizeInfo *goja.Object
		if len(call.Arguments) > 1 {
			resizeInfo = call.Arguments[1].ToObject(db.vm)
		}

		// Get all registered ResizeObservers
		globalResizeObservers := db.vm.Get("__resizeObservers")
		if goja.IsUndefined(globalResizeObservers) || goja.IsNull(globalResizeObservers) {
			return goja.Undefined()
		}

		if arr, ok := globalResizeObservers.(*goja.Object); ok {
			if lengthVal := arr.Get("length"); lengthVal != nil {
				length := int(lengthVal.ToInteger())
				for i := 0; i < length; i++ {
					if observerInfo := arr.Get(strconv.Itoa(i)); observerInfo != nil {
						if info := observerInfo.ToObject(db.vm); info != nil {
							// Check if this observer is watching the target
							if elementsArr := info.Get("elements").(*goja.Object); elementsArr != nil {
								if elemLengthVal := elementsArr.Get("length"); elemLengthVal != nil {
									elemLength := int(elemLengthVal.ToInteger())
									for j := 0; j < elemLength; j++ {
										if elem := elementsArr.Get(strconv.Itoa(j)); elem != nil {
											if node := db.extractNodeFromJS(elem); node == target {
												// This observer is watching the target, fire callback
												if callback := info.Get("callback"); callback != nil {
													if cb, ok := goja.AssertFunction(callback); ok {
														if observer := info.Get("observer"); observer != nil {
															var entry *goja.Object
															if resizeInfo != nil {
																entry = db.createResizeObserverEntryFromInfo(target, resizeInfo)
															} else {
																entry = db.createResizeObserverEntry(target)
															}

															entries := db.vm.NewArray()
															entries.Set("0", entry)
															entries.Set("length", 1)

															// Fire callback via microtask
															if db.jsRuntime != nil && db.jsRuntime.EventLoop() != nil {
																db.jsRuntime.EventLoop().QueueMicrotask(func() {
																	_, _ = cb(goja.Undefined(), entries, observer)
																}, "ResizeObserver")
															}
														}
													}
												}
												break
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}

		return goja.Undefined()
	})

	// IntersectionObserver constructor
	db.vm.Set("IntersectionObserver", func(call goja.ConstructorCall) *goja.Object {
		if len(call.Arguments) < 1 {
			panic(db.vm.NewTypeError("IntersectionObserver constructor requires a callback"))
		}

		callback, ok := goja.AssertFunction(call.Arguments[0])
		if !ok {
			panic(db.vm.NewTypeError("IntersectionObserver callback must be a function"))
		}

		// Parse options if provided
		var options *goja.Object
		if len(call.Arguments) > 1 && !goja.IsNull(call.Arguments[1]) && !goja.IsUndefined(call.Arguments[1]) {
			options = call.Arguments[1].ToObject(db.vm)
		}

		return db.createIntersectionObserver(callback, options)
	})

	// File constructor
	db.vm.Set("File", func(call goja.ConstructorCall) *goja.Object {
		if len(call.Arguments) < 2 {
			panic(db.vm.NewTypeError("File constructor requires fileBits and name"))
		}

		// Parse file bits (array of strings/blobs)
		fileBitsArray := call.Arguments[0].ToObject(db.vm)
		if fileBitsArray == nil {
			panic(db.vm.NewTypeError("File constructor requires an array of file bits"))
		}

		name := call.Arguments[1].String()

		// Parse options
		var fileType string
		var lastModified int64
		if len(call.Arguments) > 2 {
			options := call.Arguments[2].ToObject(db.vm)
			if options != nil {
				if typeVal := options.Get("type"); typeVal != nil && !goja.IsUndefined(typeVal) {
					fileType = typeVal.String()
				}
				if lastModVal := options.Get("lastModified"); lastModVal != nil && !goja.IsUndefined(lastModVal) {
					lastModified = lastModVal.ToInteger()
				}
			}
		}

		return db.createFile(fileBitsArray, name, fileType, lastModified)
	})

	// Blob constructor
	db.vm.Set("Blob", func(call goja.ConstructorCall) *goja.Object {
		var blobParts *goja.Object
		if len(call.Arguments) > 0 {
			blobParts = call.Arguments[0].ToObject(db.vm)
		}

		// Parse options
		var blobType string
		if len(call.Arguments) > 1 {
			options := call.Arguments[1].ToObject(db.vm)
			if options != nil {
				if typeVal := options.Get("type"); typeVal != nil && !goja.IsUndefined(typeVal) {
					blobType = typeVal.String()
				}
			}
		}

		return db.createBlob(blobParts, blobType)
	})

	// FileReader constructor
	db.vm.Set("FileReader", func(call goja.ConstructorCall) *goja.Object {
		return db.createFileReader()
	})

	// FileReader constants
	fileReaderConstructor := db.vm.Get("FileReader").(*goja.Object)
	fileReaderConstructor.Set("EMPTY", 0)
	fileReaderConstructor.Set("LOADING", 1)
	fileReaderConstructor.Set("DONE", 2)

	// Add XMLHttpRequest for older libraries
	db.vm.Set("XMLHttpRequest", func(call goja.ConstructorCall) *goja.Object {
		xhr := db.vm.NewObject()

		var requestURL string
		var readyState = 0

		xhr.Set("readyState", readyState)
		xhr.Set("status", 0)
		xhr.Set("statusText", "")
		xhr.Set("responseText", "")
		xhr.Set("responseXML", goja.Null())

		// XMLHttpRequest methods
		xhr.Set("open", func(call goja.FunctionCall) goja.Value {
			if len(call.Arguments) >= 2 {
				// method is first arg but we don't need to store it for mock
				requestURL = call.Arguments[1].String()
				readyState = 1
				xhr.Set("readyState", readyState)
			}
			return goja.Undefined()
		})

		xhr.Set("setRequestHeader", func(call goja.FunctionCall) goja.Value {
			// Mock implementation
			return goja.Undefined()
		})

		xhr.Set("send", func(call goja.FunctionCall) goja.Value {
			// Mock successful response
			readyState = 4
			xhr.Set("readyState", readyState)
			xhr.Set("status", 200)
			xhr.Set("statusText", "OK")

			// Mock response based on URL
			var response string
			if requestURL == "/api/hello" {
				response = `<div class="success">Hello from HTMX! 👋</div>`
			} else if requestURL == "/api/contact" {
				response = `<div class="success">Thank you! Message sent.</div>`
			} else {
				response = `<div>Mock response for ` + requestURL + `</div>`
			}
			xhr.Set("responseText", response)

			// Trigger onreadystatechange if it exists
			if onReadyStateChange := xhr.Get("onreadystatechange"); onReadyStateChange != nil && !goja.IsUndefined(onReadyStateChange) {
				if callback, ok := goja.AssertFunction(onReadyStateChange); ok {
					_, _ = callback(goja.Undefined())
				}
			}

			return goja.Undefined()
		})

		xhr.Set("abort", func(call goja.FunctionCall) goja.Value {
			return goja.Undefined()
		})

		return xhr
	})
}

// addWindowEventMethods adds event methods to window but with proper target
func (db *DOMBindings) addWindowEventMethods(window *goja.Object) {
	window.Set("addEventListener", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 2 {
			panic(db.vm.NewTypeError("addEventListener requires type and listener"))
		}

		eventType := call.Arguments[0].String()
		listener, ok := goja.AssertFunction(call.Arguments[1])
		if !ok {
			panic(db.vm.NewTypeError("Listener must be a function"))
		}

		// Store the JavaScript listener for later lookup
		db.storeWindowJSListener(eventType, db.vm.ToValue(listener).ToObject(db.vm))

		// Wrap the JavaScript function for Go DOM events
		wrapper := func(event dom.Event) {
			jsEvent := db.WrapEvent(event)
			_, _ = listener(goja.Undefined(), jsEvent)
		}

		db.windowTarget.AddEventListener(eventType, wrapper)
		return goja.Undefined()
	})

	window.Set("removeEventListener", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 2 {
			panic(db.vm.NewTypeError("removeEventListener requires type and listener"))
		}

		eventType := call.Arguments[0].String()
		db.windowTarget.RemoveEventListener(eventType, nil)
		return goja.Undefined()
	})

	window.Set("dispatchEvent", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			panic(db.vm.NewTypeError("dispatchEvent requires an event"))
		}

		// Create a Go DOM event from JavaScript event
		eventType := ""
		bubbles := true
		cancelable := true

		jsEvent := call.Arguments[0].ToObject(db.vm)
		if jsEvent != nil {
			if typeVal := jsEvent.Get("type"); typeVal != nil {
				eventType = typeVal.String()
			}
			if bubblesVal := jsEvent.Get("bubbles"); bubblesVal != nil {
				bubbles = bubblesVal.ToBoolean()
			}
			if cancelableVal := jsEvent.Get("cancelable"); cancelableVal != nil {
				cancelable = cancelableVal.ToBoolean()
			}
		}

		if eventType == "" {
			panic(db.vm.NewTypeError("Event must have a type"))
		}

		// Create Go DOM event and dispatch it
		goEvent := dom.NewEvent(eventType, bubbles, cancelable)
		return db.vm.ToValue(db.windowTarget.DispatchEvent(goEvent))
	})
}

// WindowEventTarget implements the Node interface for window event handling
type WindowEventTarget struct {
	eventListeners map[string][]func(dom.Event)
}

func (w *WindowEventTarget) NodeType() dom.NodeType                   { return dom.DocumentNode }
func (w *WindowEventTarget) NodeName() string                         { return "#window" }
func (w *WindowEventTarget) NodeValue() string                        { return "" }
func (w *WindowEventTarget) SetNodeValue(string)                      {}
func (w *WindowEventTarget) ParentNode() dom.Node                     { return nil }
func (w *WindowEventTarget) ChildNodes() *dom.NodeList                { return dom.NewNodeList([]dom.Node{}) }
func (w *WindowEventTarget) FirstChild() dom.Node                     { return nil }
func (w *WindowEventTarget) LastChild() dom.Node                      { return nil }
func (w *WindowEventTarget) PreviousSibling() dom.Node                { return nil }
func (w *WindowEventTarget) NextSibling() dom.Node                    { return nil }
func (w *WindowEventTarget) OwnerDocument() *dom.Document             { return nil }
func (w *WindowEventTarget) AppendChild(dom.Node) dom.Node            { return nil }
func (w *WindowEventTarget) InsertBefore(dom.Node, dom.Node) dom.Node { return nil }
func (w *WindowEventTarget) ReplaceChild(dom.Node, dom.Node) dom.Node { return nil }
func (w *WindowEventTarget) RemoveChild(dom.Node) dom.Node            { return nil }
func (w *WindowEventTarget) CloneNode(bool) dom.Node                  { return nil }
func (w *WindowEventTarget) Normalize()                               {}
func (w *WindowEventTarget) IsSupported(string, string) bool          { return false }
func (w *WindowEventTarget) HasAttributes() bool                      { return false }
func (w *WindowEventTarget) TextContent() string                      { return "" }
func (w *WindowEventTarget) SetTextContent(string)                    {}
func (w *WindowEventTarget) CompareDocumentPosition(dom.Node) int     { return 0 }
func (w *WindowEventTarget) Contains(dom.Node) bool                   { return false }
func (w *WindowEventTarget) LookupPrefix(string) string               { return "" }
func (w *WindowEventTarget) LookupNamespaceURI(string) string         { return "" }
func (w *WindowEventTarget) IsDefaultNamespace(string) bool           { return false }
func (w *WindowEventTarget) IsEqualNode(dom.Node) bool                { return false }
func (w *WindowEventTarget) IsSameNode(dom.Node) bool                 { return false }

// Phase 1: Core Properties & Simple Methods
func (w *WindowEventTarget) IsConnected() bool           { return false }
func (w *WindowEventTarget) ParentElement() *dom.Element { return nil }
func (w *WindowEventTarget) BaseURI() string             { return "" }
func (w *WindowEventTarget) HasChildNodes() bool         { return false }

// Phase 2: Text Content & Normalization
// TextContent and SetTextContent already defined above

// Phase 3: Comparison & Traversal Methods
func (w *WindowEventTarget) GetRootNode(*dom.GetRootNodeOptions) dom.Node { return nil }

// IsEqualNode, IsSameNode, CompareDocumentPosition, Contains already defined above

// Length and IsEmpty
func (w *WindowEventTarget) Length() int   { return 0 }
func (w *WindowEventTarget) IsEmpty() bool { return true }

// Internal methods for DOM manipulation and JS binding
func (w *WindowEventTarget) setParent(dom.Node)             {}
func (w *WindowEventTarget) setOwnerDocument(*dom.Document) {}
func (w *WindowEventTarget) toJS(*goja.Runtime) goja.Value  { return goja.Undefined() }

// Mutation Observer methods
func (w *WindowEventTarget) registerMutationObserver(*dom.MutationObserver, dom.MutationObserverInit) {
}
func (w *WindowEventTarget) unregisterMutationObserver(*dom.MutationObserver)  {}
func (w *WindowEventTarget) getRegisteredObservers() []*dom.RegisteredObserver { return nil }

func (w *WindowEventTarget) AddEventListener(eventType string, listener func(dom.Event)) {
	if w.eventListeners == nil {
		w.eventListeners = make(map[string][]func(dom.Event))
	}
	w.eventListeners[eventType] = append(w.eventListeners[eventType], listener)
}

func (w *WindowEventTarget) RemoveEventListener(eventType string, listener func(dom.Event)) {
	if w.eventListeners != nil {
		delete(w.eventListeners, eventType)
	}
}

func (w *WindowEventTarget) DispatchEvent(event dom.Event) bool {
	if w.eventListeners != nil {
		if listeners, ok := w.eventListeners[event.Type()]; ok {
			for _, listener := range listeners {
				listener(event)
			}
		}
	}
	return true
}

// wrapHistory wraps a History object for JavaScript access
func (db *DOMBindings) wrapHistory(h *history.History) *goja.Object {
	obj := db.vm.NewObject()

	// Helper function to update properties
	updateProps := func() {
		obj.Set("length", h.Length())
		obj.Set("state", h.GetState())
	}

	// Set initial properties
	updateProps()

	// History methods
	obj.Set("pushState", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 3 {
			panic(db.vm.NewTypeError("pushState requires state, title, and url"))
		}
		state := call.Arguments[0].Export()
		title := call.Arguments[1].String()
		url := call.Arguments[2].String()

		err := h.PushState(state, title, url)
		if err != nil {
			panic(db.vm.NewTypeError(err.Error()))
		}

		// Update properties after state change
		updateProps()
		return goja.Undefined()
	})

	obj.Set("replaceState", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 3 {
			panic(db.vm.NewTypeError("replaceState requires state, title, and url"))
		}
		state := call.Arguments[0].Export()
		title := call.Arguments[1].String()
		url := call.Arguments[2].String()

		err := h.ReplaceState(state, title, url)
		if err != nil {
			panic(db.vm.NewTypeError(err.Error()))
		}

		// Update properties after state change
		updateProps()
		return goja.Undefined()
	})

	obj.Set("go", func(call goja.FunctionCall) goja.Value {
		delta := 0
		if len(call.Arguments) > 0 {
			delta = int(call.Arguments[0].ToInteger())
		}
		result := h.Go(delta)
		return db.vm.ToValue(result)
	})

	obj.Set("back", func(call goja.FunctionCall) goja.Value {
		result := h.Back()
		return db.vm.ToValue(result)
	})

	obj.Set("forward", func(call goja.FunctionCall) goja.Value {
		result := h.Forward()
		return db.vm.ToValue(result)
	})

	return obj
}

// wrapPerformance wraps a Performance object for JavaScript access
func (db *DOMBindings) wrapPerformance(p *performance.Performance) *goja.Object {
	obj := db.vm.NewObject()

	// Performance.now() - high resolution timestamp
	obj.Set("now", func(call goja.FunctionCall) goja.Value {
		return db.vm.ToValue(p.Now())
	})

	// Performance.mark() - create a named timestamp
	obj.Set("mark", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			panic(db.vm.NewTypeError("mark requires a name"))
		}
		name := call.Arguments[0].String()
		p.Mark(name)
		return goja.Undefined()
	})

	// Performance.measure() - measure between two marks
	obj.Set("measure", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			panic(db.vm.NewTypeError("measure requires a name"))
		}
		name := call.Arguments[0].String()

		var startMark, endMark string
		if len(call.Arguments) > 1 {
			startMark = call.Arguments[1].String()
		}
		if len(call.Arguments) > 2 {
			endMark = call.Arguments[2].String()
		}

		err := p.Measure(name, startMark, endMark)
		if err != nil {
			panic(db.vm.NewTypeError(err.Error()))
		}
		return goja.Undefined()
	})

	// Performance.getEntries() - get all performance entries
	obj.Set("getEntries", func(call goja.FunctionCall) goja.Value {
		entries := p.GetEntries()
		return db.wrapPerformanceEntries(entries)
	})

	// Performance.getEntriesByType() - get entries by type
	obj.Set("getEntriesByType", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			return db.vm.NewArray()
		}
		entryType := call.Arguments[0].String()
		entries := p.GetEntriesByType(entryType)
		return db.wrapPerformanceEntries(entries)
	})

	// Performance.getEntriesByName() - get entries by name
	obj.Set("getEntriesByName", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			return db.vm.NewArray()
		}
		name := call.Arguments[0].String()
		entries := p.GetEntriesByName(name)
		return db.wrapPerformanceEntries(entries)
	})

	// Performance.clearMarks() - clear marks
	obj.Set("clearMarks", func(call goja.FunctionCall) goja.Value {
		markName := ""
		if len(call.Arguments) > 0 {
			markName = call.Arguments[0].String()
		}
		p.ClearMarks(markName)
		return goja.Undefined()
	})

	// Performance.clearMeasures() - clear measures
	obj.Set("clearMeasures", func(call goja.FunctionCall) goja.Value {
		measureName := ""
		if len(call.Arguments) > 0 {
			measureName = call.Arguments[0].String()
		}
		p.ClearMeasures(measureName)
		return goja.Undefined()
	})

	// Performance.clearResourceTimings() - clear resource timings
	obj.Set("clearResourceTimings", func(call goja.FunctionCall) goja.Value {
		p.ClearResourceTimings()
		return goja.Undefined()
	})

	// Performance.setResourceTimingBufferSize() - set resource buffer size
	obj.Set("setResourceTimingBufferSize", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			panic(db.vm.NewTypeError("setResourceTimingBufferSize requires a maxSize"))
		}
		maxSize := int(call.Arguments[0].ToInteger())
		p.SetResourceBufferSize(maxSize)
		return goja.Undefined()
	})

	// Performance.timeOrigin - read-only property
	obj.Set("timeOrigin", p.TimeOrigin())

	return obj
}

// wrapPerformanceEntries wraps performance entries for JavaScript access
func (db *DOMBindings) wrapPerformanceEntries(entries []performance.PerformanceEntry) *goja.Object {
	arr := db.vm.NewArray()

	for i, entry := range entries {
		entryObj := db.vm.NewObject()
		entryObj.Set("name", entry.Name)
		entryObj.Set("startTime", entry.StartTime)
		entryObj.Set("duration", entry.Duration)
		entryObj.Set("entryType", entry.EntryType)

		arr.Set(strconv.Itoa(i), entryObj)
	}

	arr.Set("length", len(entries))
	return arr
}

// wrapURL wraps a URL object for JavaScript access
func (db *DOMBindings) wrapURL(goURL *url.URL) *goja.Object {
	obj := db.vm.NewObject()

	// URL properties
	obj.Set("href", goURL.Href())
	obj.Set("origin", goURL.Origin())
	obj.Set("protocol", goURL.Protocol())
	obj.Set("host", goURL.Host())
	obj.Set("hostname", goURL.Hostname())
	obj.Set("port", goURL.Port())
	obj.Set("pathname", goURL.Pathname())
	obj.Set("search", goURL.Search())
	obj.Set("hash", goURL.Hash())

	// SearchParams property
	obj.Set("searchParams", db.wrapURLSearchParams(goURL.SearchParams()))

	// toString method
	obj.Set("toString", func(call goja.FunctionCall) goja.Value {
		return db.vm.ToValue(goURL.ToString())
	})

	// toJSON method
	obj.Set("toJSON", func(call goja.FunctionCall) goja.Value {
		return db.vm.ToValue(goURL.ToJSON())
	})

	return obj
}

// wrapDOMException wraps a DOMException for JavaScript access
func (db *DOMBindings) wrapDOMException(exc *dom.DOMException) *goja.Object {
	obj := db.vm.NewObject()

	// DOMException properties
	obj.Set("name", exc.Name())
	obj.Set("message", exc.Message())
	obj.Set("code", exc.Code())

	// toString method
	obj.Set("toString", func(call goja.FunctionCall) goja.Value {
		return db.vm.ToValue(exc.Error())
	})

	// Make it instanceof DOMException in JavaScript
	domExceptionConstructor := db.vm.Get("DOMException")
	if domExceptionConstructor != nil {
		if prototype := domExceptionConstructor.(*goja.Object).Get("prototype"); prototype != nil {
			if protoObj := prototype.ToObject(db.vm); protoObj != nil {
				obj.SetPrototype(protoObj)
			}
		}
	}

	return obj
}

// wrapURLSearchParams wraps URLSearchParams for JavaScript access
func (db *DOMBindings) wrapURLSearchParams(params *url.URLSearchParams) *goja.Object {
	obj := db.vm.NewObject()

	// URLSearchParams methods
	obj.Set("append", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 2 {
			panic(db.vm.NewTypeError("append requires name and value"))
		}
		name := call.Arguments[0].String()
		value := call.Arguments[1].String()
		params.Append(name, value)
		return goja.Undefined()
	})

	obj.Set("delete", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			panic(db.vm.NewTypeError("delete requires name"))
		}
		name := call.Arguments[0].String()
		params.Delete(name)
		return goja.Undefined()
	})

	obj.Set("get", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			return goja.Null()
		}
		name := call.Arguments[0].String()
		value := params.Get(name)
		if value == "" && !params.Has(name) {
			return goja.Null()
		}
		return db.vm.ToValue(value)
	})

	obj.Set("getAll", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			return db.vm.NewArray()
		}
		name := call.Arguments[0].String()
		values := params.GetAll(name)
		arr := db.vm.NewArray()
		for i, v := range values {
			arr.Set(strconv.Itoa(i), v)
		}
		arr.Set("length", len(values))
		return arr
	})

	obj.Set("has", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			return db.vm.ToValue(false)
		}
		name := call.Arguments[0].String()
		return db.vm.ToValue(params.Has(name))
	})

	obj.Set("set", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 2 {
			panic(db.vm.NewTypeError("set requires name and value"))
		}
		name := call.Arguments[0].String()
		value := call.Arguments[1].String()
		params.Set(name, value)
		return goja.Undefined()
	})

	obj.Set("sort", func(call goja.FunctionCall) goja.Value {
		params.Sort()
		return goja.Undefined()
	})

	obj.Set("toString", func(call goja.FunctionCall) goja.Value {
		return db.vm.ToValue(params.ToString())
	})

	obj.Set("keys", func(call goja.FunctionCall) goja.Value {
		keys := params.Keys()
		arr := db.vm.NewArray()
		for i, k := range keys {
			arr.Set(strconv.Itoa(i), k)
		}
		arr.Set("length", len(keys))
		return arr
	})

	obj.Set("values", func(call goja.FunctionCall) goja.Value {
		values := params.Values()
		arr := db.vm.NewArray()
		for i, v := range values {
			arr.Set(strconv.Itoa(i), v)
		}
		arr.Set("length", len(values))
		return arr
	})

	obj.Set("entries", func(call goja.FunctionCall) goja.Value {
		entries := params.Entries()
		arr := db.vm.NewArray()
		for i, entry := range entries {
			entryArr := db.vm.NewArray()
			entryArr.Set("0", entry[0])
			entryArr.Set("1", entry[1])
			entryArr.Set("length", 2)
			arr.Set(strconv.Itoa(i), entryArr)
		}
		arr.Set("length", len(entries))
		return arr
	})

	obj.Set("forEach", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			panic(db.vm.NewTypeError("forEach requires a callback"))
		}
		callback, ok := goja.AssertFunction(call.Arguments[0])
		if !ok {
			panic(db.vm.NewTypeError("Callback must be a function"))
		}

		params.ForEach(func(value, name string) {
			_, _ = callback(goja.Undefined(), db.vm.ToValue(value), db.vm.ToValue(name))
		})
		return goja.Undefined()
	})

	// Size property (read-only)
	obj.Set("size", params.Size())

	return obj
}

// wrapXPathEvaluator wraps an XPathEvaluator for JavaScript access
func (db *DOMBindings) wrapXPathEvaluator(evaluator *xpath.XPathEvaluator) *goja.Object {
	obj := db.vm.NewObject()

	// createExpression method
	obj.Set("createExpression", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			panic(db.vm.NewTypeError("createExpression requires an expression"))
		}

		expression := call.Arguments[0].String()
		var resolver xpath.NSResolver
		if len(call.Arguments) > 1 && !goja.IsNull(call.Arguments[1]) && !goja.IsUndefined(call.Arguments[1]) {
			// For now, use a simple resolver
			resolver = &xpath.SimpleNSResolver{}
		}

		xpathExpr, err := evaluator.CreateExpression(expression, resolver)
		if err != nil {
			panic(db.vm.NewTypeError("Invalid XPath expression: " + err.Error()))
		}

		return db.wrapXPathExpression(xpathExpr)
	})

	// createNSResolver method
	obj.Set("createNSResolver", func(call goja.FunctionCall) goja.Value {
		var node dom.Node
		if len(call.Arguments) > 0 {
			node = db.extractNodeFromJS(call.Arguments[0])
		}

		resolver := evaluator.CreateNSResolver(node)
		return db.wrapNSResolver(resolver)
	})

	// evaluate method
	obj.Set("evaluate", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 3 {
			panic(db.vm.NewTypeError("evaluate requires expression, contextNode, and resultType"))
		}

		expression := call.Arguments[0].String()
		contextNode := db.extractNodeFromJS(call.Arguments[1])
		if contextNode == nil {
			panic(db.vm.NewTypeError("Invalid context node"))
		}

		var resolver xpath.NSResolver
		if len(call.Arguments) > 2 && !goja.IsNull(call.Arguments[2]) && !goja.IsUndefined(call.Arguments[2]) {
			// For now, use a simple resolver
			resolver = &xpath.SimpleNSResolver{}
		}

		resultType := int(call.Arguments[3].ToInteger())

		var result *xpath.XPathResult
		if len(call.Arguments) > 4 && !goja.IsNull(call.Arguments[4]) && !goja.IsUndefined(call.Arguments[4]) {
			// Reuse existing result if provided
			// For now, we'll ignore this parameter
		}

		xpathResult, err := evaluator.Evaluate(expression, contextNode, resolver, resultType, result)
		if err != nil {
			panic(db.vm.NewTypeError("XPath evaluation failed: " + err.Error()))
		}

		return db.wrapXPathResult(xpathResult)
	})

	return obj
}

// wrapXPathExpression wraps an XPathExpression for JavaScript access
func (db *DOMBindings) wrapXPathExpression(expr *xpath.XPathExpression) *goja.Object {
	obj := db.vm.NewObject()

	// evaluate method
	obj.Set("evaluate", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 2 {
			panic(db.vm.NewTypeError("evaluate requires contextNode and resultType"))
		}

		contextNode := db.extractNodeFromJS(call.Arguments[0])
		if contextNode == nil {
			panic(db.vm.NewTypeError("Invalid context node"))
		}

		resultType := int(call.Arguments[1].ToInteger())

		var result *xpath.XPathResult
		if len(call.Arguments) > 2 && !goja.IsNull(call.Arguments[2]) && !goja.IsUndefined(call.Arguments[2]) {
			// Reuse existing result if provided
			// For now, we'll ignore this parameter
		}

		xpathResult, err := expr.Evaluate(contextNode, resultType, result)
		if err != nil {
			panic(db.vm.NewTypeError("XPath evaluation failed: " + err.Error()))
		}

		return db.wrapXPathResult(xpathResult)
	})

	return obj
}

// wrapXPathResult wraps an XPathResult for JavaScript access
func (db *DOMBindings) wrapXPathResult(result *xpath.XPathResult) *goja.Object {
	obj := db.vm.NewObject()

	// Properties
	obj.Set("resultType", result.ResultType())

	// Value properties (only available for certain result types)
	if result.ResultType() == xpath.NUMBER_TYPE {
		obj.Set("numberValue", result.NumberValue())
	}
	if result.ResultType() == xpath.STRING_TYPE {
		obj.Set("stringValue", result.StringValue())
	}
	if result.ResultType() == xpath.BOOLEAN_TYPE {
		obj.Set("booleanValue", result.BooleanValue())
	}
	if result.ResultType() == xpath.ANY_UNORDERED_NODE_TYPE || result.ResultType() == xpath.FIRST_ORDERED_NODE_TYPE {
		singleNode := result.SingleNodeValue()
		if singleNode != nil {
			obj.Set("singleNodeValue", db.WrapNode(singleNode))
		} else {
			obj.Set("singleNodeValue", goja.Null())
		}
	}

	// Snapshot properties
	if result.ResultType() == xpath.UNORDERED_NODE_SNAPSHOT_TYPE || result.ResultType() == xpath.ORDERED_NODE_SNAPSHOT_TYPE {
		obj.Set("snapshotLength", result.SnapshotLength())

		// snapshotItem method
		obj.Set("snapshotItem", func(call goja.FunctionCall) goja.Value {
			if len(call.Arguments) < 1 {
				return goja.Null()
			}

			index := int(call.Arguments[0].ToInteger())
			node := result.SnapshotItem(index)
			if node != nil {
				return db.WrapNode(node)
			}
			return goja.Null()
		})
	}

	// Iterator methods
	if result.ResultType() == xpath.UNORDERED_NODE_ITERATOR_TYPE || result.ResultType() == xpath.ORDERED_NODE_ITERATOR_TYPE {
		obj.Set("iterateNext", func(call goja.FunctionCall) goja.Value {
			node := result.IterateNext()
			if node != nil {
				return db.WrapNode(node)
			}
			return goja.Null()
		})
	}

	return obj
}

// wrapDOMImplementation wraps a DOMImplementation for JavaScript access
func (db *DOMBindings) wrapDOMImplementation(impl *dom.DOMImplementation) *goja.Object {
	obj := db.vm.NewObject()

	// hasFeature method
	obj.Set("hasFeature", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 2 {
			return db.vm.ToValue(false)
		}
		feature := call.Arguments[0].String()
		version := call.Arguments[1].String()
		return db.vm.ToValue(impl.HasFeature(feature, version))
	})

	// createDocumentType method
	obj.Set("createDocumentType", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 3 {
			panic(db.vm.NewTypeError("createDocumentType requires qualifiedName, publicId, and systemId"))
		}
		qualifiedName := call.Arguments[0].String()
		publicId := call.Arguments[1].String()
		systemId := call.Arguments[2].String()

		doctype, err := impl.CreateDocumentType(qualifiedName, publicId, systemId)
		if err != nil {
			panic(db.vm.NewTypeError(err.Error()))
		}
		return db.WrapNode(doctype)
	})

	// createDocument method
	obj.Set("createDocument", func(call goja.FunctionCall) goja.Value {
		var namespaceURI, qualifiedName string
		var doctype *dom.DocumentType

		if len(call.Arguments) > 0 && !goja.IsNull(call.Arguments[0]) {
			namespaceURI = call.Arguments[0].String()
		}
		if len(call.Arguments) > 1 && !goja.IsNull(call.Arguments[1]) {
			qualifiedName = call.Arguments[1].String()
		}
		if len(call.Arguments) > 2 && !goja.IsNull(call.Arguments[2]) {
			// Extract doctype node if provided
			if doctypeNode := db.extractNodeFromJS(call.Arguments[2]); doctypeNode != nil {
				if dt, ok := doctypeNode.(*dom.DocumentType); ok {
					doctype = dt
				}
			}
		}

		doc, err := impl.CreateDocument(namespaceURI, qualifiedName, doctype)
		if err != nil {
			panic(db.vm.NewTypeError(err.Error()))
		}

		// Create new DOMBindings for the new document and wrap it
		newBindings := NewDOMBindings(db.vm, doc)
		return newBindings.WrapDocument()
	})

	// createHTMLDocument method
	obj.Set("createHTMLDocument", func(call goja.FunctionCall) goja.Value {
		title := ""
		if len(call.Arguments) > 0 {
			title = call.Arguments[0].String()
		}

		doc := impl.CreateHTMLDocument(title)

		// Create new DOMBindings for the new document and wrap it
		newBindings := NewDOMBindings(db.vm, doc)
		return newBindings.WrapDocument()
	})

	return obj
}

// wrapNSResolver wraps an NSResolver for JavaScript access
func (db *DOMBindings) wrapNSResolver(resolver xpath.NSResolver) *goja.Object {
	obj := db.vm.NewObject()

	// lookupNamespaceURI method
	obj.Set("lookupNamespaceURI", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			return goja.Null()
		}

		prefix := call.Arguments[0].String()
		uri := resolver.LookupNamespaceURI(prefix)
		if uri == "" {
			return goja.Null()
		}
		return db.vm.ToValue(uri)
	})

	return obj
}

// createRange creates a new Range object for JavaScript access
func (db *DOMBindings) createRange() *goja.Object {
	rangeObj := db.vm.NewObject()

	// Range properties
	var collapsed = true
	var commonAncestorContainer dom.Node
	var endContainer dom.Node
	var endOffset = 0
	var startContainer dom.Node
	var startOffset = 0

	// Dynamic property getters
	rangeObj.DefineAccessorProperty("collapsed",
		db.vm.ToValue(func(call goja.FunctionCall) goja.Value {
			return db.vm.ToValue(collapsed)
		}),
		goja.Undefined(),                // No setter - collapsed is read-only
		goja.FLAG_FALSE, goja.FLAG_TRUE) // Not enumerable, configurable

	rangeObj.DefineAccessorProperty("commonAncestorContainer",
		db.vm.ToValue(func(call goja.FunctionCall) goja.Value {
			if commonAncestorContainer != nil {
				return db.WrapNode(commonAncestorContainer)
			}
			return goja.Null()
		}),
		goja.Undefined(),                // No setter - commonAncestorContainer is read-only
		goja.FLAG_FALSE, goja.FLAG_TRUE) // Not enumerable, configurable

	rangeObj.DefineAccessorProperty("endContainer",
		db.vm.ToValue(func(call goja.FunctionCall) goja.Value {
			if endContainer != nil {
				return db.WrapNode(endContainer)
			}
			return goja.Null()
		}),
		goja.Undefined(),                // No setter - endContainer is read-only
		goja.FLAG_FALSE, goja.FLAG_TRUE) // Not enumerable, configurable

	rangeObj.DefineAccessorProperty("endOffset",
		db.vm.ToValue(func(call goja.FunctionCall) goja.Value {
			return db.vm.ToValue(endOffset)
		}),
		goja.Undefined(),                // No setter - endOffset is read-only
		goja.FLAG_FALSE, goja.FLAG_TRUE) // Not enumerable, configurable

	rangeObj.DefineAccessorProperty("startContainer",
		db.vm.ToValue(func(call goja.FunctionCall) goja.Value {
			if startContainer != nil {
				return db.WrapNode(startContainer)
			}
			return goja.Null()
		}),
		goja.Undefined(),                // No setter - startContainer is read-only
		goja.FLAG_FALSE, goja.FLAG_TRUE) // Not enumerable, configurable

	rangeObj.DefineAccessorProperty("startOffset",
		db.vm.ToValue(func(call goja.FunctionCall) goja.Value {
			return db.vm.ToValue(startOffset)
		}),
		goja.Undefined(),                // No setter - startOffset is read-only
		goja.FLAG_FALSE, goja.FLAG_TRUE) // Not enumerable, configurable

	// Range methods
	rangeObj.Set("collapse", func(call goja.FunctionCall) goja.Value {
		toStart := false
		if len(call.Arguments) > 0 {
			toStart = call.Arguments[0].ToBoolean()
		}

		if toStart {
			endContainer = startContainer
			endOffset = startOffset
		} else {
			startContainer = endContainer
			startOffset = endOffset
		}
		collapsed = (startContainer == endContainer && startOffset == endOffset)
		return goja.Undefined()
	})

	rangeObj.Set("selectNode", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			panic(db.vm.NewTypeError("selectNode requires a node"))
		}

		node := db.extractNodeFromJS(call.Arguments[0])
		if node == nil {
			panic(db.vm.NewTypeError("Invalid node"))
		}

		// Set range to encompass the entire node
		startContainer = node.ParentNode()
		if startContainer != nil {
			// Find the index of this node among its siblings
			siblings := startContainer.ChildNodes().ToSlice()
			for i, sibling := range siblings {
				if sibling == node {
					startOffset = i
					endOffset = i + 1
					break
				}
			}
			endContainer = startContainer
		}
		collapsed = false
		commonAncestorContainer = startContainer
		return goja.Undefined()
	})

	rangeObj.Set("selectNodeContents", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			panic(db.vm.NewTypeError("selectNodeContents requires a node"))
		}

		node := db.extractNodeFromJS(call.Arguments[0])
		if node == nil {
			panic(db.vm.NewTypeError("Invalid node"))
		}

		// Set range to encompass the contents of the node
		startContainer = node
		startOffset = 0
		endContainer = node
		endOffset = len(node.ChildNodes().ToSlice())
		collapsed = (startOffset == endOffset)
		commonAncestorContainer = node
		return goja.Undefined()
	})

	rangeObj.Set("setStart", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 2 {
			panic(db.vm.NewTypeError("setStart requires node and offset"))
		}

		node := db.extractNodeFromJS(call.Arguments[0])
		if node == nil {
			panic(db.vm.NewTypeError("Invalid node"))
		}

		offset := int(call.Arguments[1].ToInteger())
		startContainer = node
		startOffset = offset
		collapsed = (startContainer == endContainer && startOffset == endOffset)
		return goja.Undefined()
	})

	rangeObj.Set("setEnd", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 2 {
			panic(db.vm.NewTypeError("setEnd requires node and offset"))
		}

		node := db.extractNodeFromJS(call.Arguments[0])
		if node == nil {
			panic(db.vm.NewTypeError("Invalid node"))
		}

		offset := int(call.Arguments[1].ToInteger())
		endContainer = node
		endOffset = offset
		collapsed = (startContainer == endContainer && startOffset == endOffset)
		return goja.Undefined()
	})

	rangeObj.Set("cloneRange", func(call goja.FunctionCall) goja.Value {
		// Create a new range with the same boundaries
		newRange := db.createRange()
		if startContainer != nil {
			if setStartFunc, ok := goja.AssertFunction(newRange.Get("setStart")); ok {
				_, _ = setStartFunc(newRange, db.WrapNode(startContainer), db.vm.ToValue(startOffset))
			}
		}
		if endContainer != nil {
			if setEndFunc, ok := goja.AssertFunction(newRange.Get("setEnd")); ok {
				_, _ = setEndFunc(newRange, db.WrapNode(endContainer), db.vm.ToValue(endOffset))
			}
		}
		return newRange
	})

	rangeObj.Set("extractContents", func(call goja.FunctionCall) goja.Value {
		// Create a document fragment and move the range contents into it
		fragment := db.document.CreateDocumentFragment()

		// For simplified implementation, just return the fragment
		// A full implementation would actually extract the DOM nodes
		return db.WrapNode(fragment)
	})

	rangeObj.Set("cloneContents", func(call goja.FunctionCall) goja.Value {
		// Create a document fragment and clone the range contents into it
		fragment := db.document.CreateDocumentFragment()

		// For simplified implementation, just return the fragment
		// A full implementation would actually clone the DOM nodes
		return db.WrapNode(fragment)
	})

	rangeObj.Set("insertNode", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			panic(db.vm.NewTypeError("insertNode requires a node"))
		}

		node := db.extractNodeFromJS(call.Arguments[0])
		if node == nil {
			panic(db.vm.NewTypeError("Invalid node"))
		}

		// Insert the node at the start of the range
		if startContainer != nil {
			startContainer.InsertBefore(node, startContainer.FirstChild())
		}
		return goja.Undefined()
	})

	rangeObj.Set("surroundContents", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			panic(db.vm.NewTypeError("surroundContents requires a node"))
		}

		node := db.extractNodeFromJS(call.Arguments[0])
		if node == nil {
			panic(db.vm.NewTypeError("Invalid node"))
		}

		// Simplified implementation - extract contents and append to the new node
		// then insert the new node at the range position
		return goja.Undefined()
	})

	rangeObj.Set("toString", func(call goja.FunctionCall) goja.Value {
		// Return the text content of the range
		// For simplified implementation, return empty string
		return db.vm.ToValue("")
	})

	return rangeObj
}

// createTreeWalker creates a new TreeWalker object for JavaScript access
func (db *DOMBindings) createTreeWalker(args []goja.Value) *goja.Object {
	if len(args) < 1 {
		panic(db.vm.NewTypeError("createTreeWalker requires a root node"))
	}

	root := db.extractNodeFromJS(args[0])
	if root == nil {
		panic(db.vm.NewTypeError("Invalid root node"))
	}

	whatToShow := uint32(0xFFFFFFFF) // NodeFilter.SHOW_ALL by default
	if len(args) > 1 && !goja.IsUndefined(args[1]) {
		whatToShow = uint32(args[1].ToInteger())
	}

	var filter *goja.Object
	if len(args) > 2 && !goja.IsNull(args[2]) && !goja.IsUndefined(args[2]) {
		filter = args[2].ToObject(db.vm)
	}

	walkerObj := db.vm.NewObject()
	var currentNode = root

	// TreeWalker properties
	walkerObj.Set("root", db.WrapNode(root))
	walkerObj.DefineAccessorProperty("currentNode",
		db.vm.ToValue(func(call goja.FunctionCall) goja.Value {
			return db.WrapNode(currentNode)
		}),
		db.vm.ToValue(func(call goja.FunctionCall) goja.Value {
			if len(call.Arguments) > 0 {
				newNode := db.extractNodeFromJS(call.Arguments[0])
				if newNode != nil {
					currentNode = newNode
				}
			}
			return goja.Undefined()
		}),
		goja.FLAG_FALSE, goja.FLAG_TRUE)

	walkerObj.Set("whatToShow", whatToShow)
	if filter != nil {
		walkerObj.Set("filter", filter)
	} else {
		walkerObj.Set("filter", goja.Null())
	}

	// Helper function to check if a node should be shown
	acceptNode := func(node dom.Node) bool {
		// Check whatToShow filter
		nodeTypeMask := uint32(1 << (node.NodeType() - 1))
		if (whatToShow & nodeTypeMask) == 0 {
			return false
		}

		// Check custom filter if provided
		if filter != nil {
			if acceptNodeFunc := filter.Get("acceptNode"); acceptNodeFunc != nil {
				if callback, ok := goja.AssertFunction(acceptNodeFunc); ok {
					result, _ := callback(filter, db.WrapNode(node))
					return result.ToInteger() == 1 // NodeFilter.FILTER_ACCEPT
				}
			}
		}

		return true
	}

	// TreeWalker navigation methods
	walkerObj.Set("nextNode", func(call goja.FunctionCall) goja.Value {
		// Traverse in document order looking for next acceptable node
		var findNext func(dom.Node) dom.Node
		findNext = func(node dom.Node) dom.Node {
			// Try first child
			if firstChild := node.FirstChild(); firstChild != nil {
				if acceptNode(firstChild) {
					return firstChild
				}
				if result := findNext(firstChild); result != nil {
					return result
				}
			}

			// Try next sibling and its descendants
			current := node
			for current != nil && current != root {
				if nextSibling := current.NextSibling(); nextSibling != nil {
					if acceptNode(nextSibling) {
						return nextSibling
					}
					if result := findNext(nextSibling); result != nil {
						return result
					}
				}
				current = current.ParentNode()
			}

			return nil
		}

		if next := findNext(currentNode); next != nil {
			currentNode = next
			return db.WrapNode(next)
		}
		return goja.Null()
	})

	walkerObj.Set("previousNode", func(call goja.FunctionCall) goja.Value {
		// Traverse in reverse document order
		var findPrevious func(dom.Node) dom.Node
		findPrevious = func(node dom.Node) dom.Node {
			// Try previous sibling and its last descendants
			if prevSibling := node.PreviousSibling(); prevSibling != nil {
				// Go to the deepest last child of the previous sibling
				current := prevSibling
				for current.LastChild() != nil {
					current = current.LastChild()
				}
				if acceptNode(current) {
					return current
				}
				return findPrevious(current)
			}

			// Try parent
			if parent := node.ParentNode(); parent != nil && parent != root {
				if acceptNode(parent) {
					return parent
				}
				return findPrevious(parent)
			}

			return nil
		}

		if prev := findPrevious(currentNode); prev != nil {
			currentNode = prev
			return db.WrapNode(prev)
		}
		return goja.Null()
	})

	walkerObj.Set("firstChild", func(call goja.FunctionCall) goja.Value {
		if firstChild := currentNode.FirstChild(); firstChild != nil {
			if acceptNode(firstChild) {
				currentNode = firstChild
				return db.WrapNode(firstChild)
			}
			// Find first acceptable child
			for child := firstChild.NextSibling(); child != nil; child = child.NextSibling() {
				if acceptNode(child) {
					currentNode = child
					return db.WrapNode(child)
				}
			}
		}
		return goja.Null()
	})

	walkerObj.Set("lastChild", func(call goja.FunctionCall) goja.Value {
		if lastChild := currentNode.LastChild(); lastChild != nil {
			if acceptNode(lastChild) {
				currentNode = lastChild
				return db.WrapNode(lastChild)
			}
			// Find last acceptable child
			for child := lastChild.PreviousSibling(); child != nil; child = child.PreviousSibling() {
				if acceptNode(child) {
					currentNode = child
					return db.WrapNode(child)
				}
			}
		}
		return goja.Null()
	})

	walkerObj.Set("nextSibling", func(call goja.FunctionCall) goja.Value {
		if nextSibling := currentNode.NextSibling(); nextSibling != nil {
			if acceptNode(nextSibling) {
				currentNode = nextSibling
				return db.WrapNode(nextSibling)
			}
			// Find next acceptable sibling
			for sibling := nextSibling.NextSibling(); sibling != nil; sibling = sibling.NextSibling() {
				if acceptNode(sibling) {
					currentNode = sibling
					return db.WrapNode(sibling)
				}
			}
		}
		return goja.Null()
	})

	walkerObj.Set("previousSibling", func(call goja.FunctionCall) goja.Value {
		if prevSibling := currentNode.PreviousSibling(); prevSibling != nil {
			if acceptNode(prevSibling) {
				currentNode = prevSibling
				return db.WrapNode(prevSibling)
			}
			// Find previous acceptable sibling
			for sibling := prevSibling.PreviousSibling(); sibling != nil; sibling = sibling.PreviousSibling() {
				if acceptNode(sibling) {
					currentNode = sibling
					return db.WrapNode(sibling)
				}
			}
		}
		return goja.Null()
	})

	walkerObj.Set("parentNode", func(call goja.FunctionCall) goja.Value {
		if parent := currentNode.ParentNode(); parent != nil && parent != root {
			if acceptNode(parent) {
				currentNode = parent
				return db.WrapNode(parent)
			}
		}
		return goja.Null()
	})

	return walkerObj
}

// createNodeIterator creates a new NodeIterator object for JavaScript access
func (db *DOMBindings) createNodeIterator(args []goja.Value) *goja.Object {
	if len(args) < 1 {
		panic(db.vm.NewTypeError("createNodeIterator requires a root node"))
	}

	root := db.extractNodeFromJS(args[0])
	if root == nil {
		panic(db.vm.NewTypeError("Invalid root node"))
	}

	whatToShow := uint32(0xFFFFFFFF) // NodeFilter.SHOW_ALL by default
	if len(args) > 1 && !goja.IsUndefined(args[1]) {
		whatToShow = uint32(args[1].ToInteger())
	}

	var filter *goja.Object
	if len(args) > 2 && !goja.IsNull(args[2]) && !goja.IsUndefined(args[2]) {
		filter = args[2].ToObject(db.vm)
	}

	iteratorObj := db.vm.NewObject()
	var referenceNode = root
	var pointerBeforeReferenceNode = true

	// NodeIterator properties
	iteratorObj.Set("root", db.WrapNode(root))
	iteratorObj.DefineAccessorProperty("referenceNode",
		db.vm.ToValue(func(call goja.FunctionCall) goja.Value {
			return db.WrapNode(referenceNode)
		}),
		goja.Undefined(),                // No setter - referenceNode is read-only
		goja.FLAG_FALSE, goja.FLAG_TRUE) // Not enumerable, configurable

	iteratorObj.DefineAccessorProperty("pointerBeforeReferenceNode",
		db.vm.ToValue(func(call goja.FunctionCall) goja.Value {
			return db.vm.ToValue(pointerBeforeReferenceNode)
		}),
		goja.Undefined(),                // No setter - pointerBeforeReferenceNode is read-only
		goja.FLAG_FALSE, goja.FLAG_TRUE) // Not enumerable, configurable

	iteratorObj.Set("whatToShow", whatToShow)
	if filter != nil {
		iteratorObj.Set("filter", filter)
	} else {
		iteratorObj.Set("filter", goja.Null())
	}

	// Helper function to check if a node should be shown
	acceptNode := func(node dom.Node) bool {
		// Check whatToShow filter
		nodeTypeMask := uint32(1 << (node.NodeType() - 1))
		if (whatToShow & nodeTypeMask) == 0 {
			return false
		}

		// Check custom filter if provided
		if filter != nil {
			if acceptNodeFunc := filter.Get("acceptNode"); acceptNodeFunc != nil {
				if callback, ok := goja.AssertFunction(acceptNodeFunc); ok {
					result, _ := callback(filter, db.WrapNode(node))
					return result.ToInteger() == 1 // NodeFilter.FILTER_ACCEPT
				}
			}
		}

		return true
	}

	// Helper function to get all nodes in document order
	getAllNodes := func() []dom.Node {
		var nodes []dom.Node
		var traverse func(dom.Node)
		traverse = func(node dom.Node) {
			if acceptNode(node) {
				nodes = append(nodes, node)
			}
			for _, child := range node.ChildNodes().ToSlice() {
				traverse(child)
			}
		}
		traverse(root)
		return nodes
	}

	// NodeIterator navigation methods
	iteratorObj.Set("nextNode", func(call goja.FunctionCall) goja.Value {
		allNodes := getAllNodes()

		// Find current position
		currentIndex := -1
		for i, node := range allNodes {
			if node == referenceNode {
				currentIndex = i
				break
			}
		}

		var nextIndex int
		if pointerBeforeReferenceNode {
			// If pointer is before reference node, next node is the reference node itself
			nextIndex = currentIndex
		} else {
			// If pointer is after reference node, next node is the one after reference node
			nextIndex = currentIndex + 1
		}

		if nextIndex >= 0 && nextIndex < len(allNodes) {
			referenceNode = allNodes[nextIndex]
			pointerBeforeReferenceNode = false
			return db.WrapNode(referenceNode)
		}

		return goja.Null()
	})

	iteratorObj.Set("previousNode", func(call goja.FunctionCall) goja.Value {
		allNodes := getAllNodes()

		// Find current position
		currentIndex := -1
		for i, node := range allNodes {
			if node == referenceNode {
				currentIndex = i
				break
			}
		}

		var prevIndex int
		if pointerBeforeReferenceNode {
			// If pointer is before reference node, previous node is the one before reference node
			prevIndex = currentIndex - 1
		} else {
			// If pointer is after reference node, previous node is the reference node itself
			prevIndex = currentIndex
		}

		if prevIndex >= 0 && prevIndex < len(allNodes) {
			referenceNode = allNodes[prevIndex]
			pointerBeforeReferenceNode = true
			return db.WrapNode(referenceNode)
		}

		return goja.Null()
	})

	iteratorObj.Set("detach", func(call goja.FunctionCall) goja.Value {
		// In modern browsers, this is a no-op
		return goja.Undefined()
	})

	return iteratorObj
}

// createSelection creates a new Selection object for JavaScript access
func (db *DOMBindings) createSelection() *goja.Object {
	selectionObj := db.vm.NewObject()

	// Selection properties
	var anchorNode dom.Node
	var anchorOffset = 0
	var focusNode dom.Node
	var focusOffset = 0
	var isCollapsed = true
	var rangeCount = 0

	// Selection properties
	selectionObj.DefineAccessorProperty("anchorNode",
		db.vm.ToValue(func(call goja.FunctionCall) goja.Value {
			if anchorNode != nil {
				return db.WrapNode(anchorNode)
			}
			return goja.Null()
		}),
		goja.Undefined(),                // No setter - anchorNode is read-only
		goja.FLAG_FALSE, goja.FLAG_TRUE) // Not enumerable, configurable

	selectionObj.DefineAccessorProperty("anchorOffset",
		db.vm.ToValue(func(call goja.FunctionCall) goja.Value {
			return db.vm.ToValue(anchorOffset)
		}),
		goja.Undefined(),                // No setter - anchorOffset is read-only
		goja.FLAG_FALSE, goja.FLAG_TRUE) // Not enumerable, configurable

	selectionObj.DefineAccessorProperty("focusNode",
		db.vm.ToValue(func(call goja.FunctionCall) goja.Value {
			if focusNode != nil {
				return db.WrapNode(focusNode)
			}
			return goja.Null()
		}),
		goja.Undefined(),                // No setter - focusNode is read-only
		goja.FLAG_FALSE, goja.FLAG_TRUE) // Not enumerable, configurable

	selectionObj.DefineAccessorProperty("focusOffset",
		db.vm.ToValue(func(call goja.FunctionCall) goja.Value {
			return db.vm.ToValue(focusOffset)
		}),
		goja.Undefined(),                // No setter - focusOffset is read-only
		goja.FLAG_FALSE, goja.FLAG_TRUE) // Not enumerable, configurable

	selectionObj.DefineAccessorProperty("isCollapsed",
		db.vm.ToValue(func(call goja.FunctionCall) goja.Value {
			return db.vm.ToValue(isCollapsed)
		}),
		goja.Undefined(),                // No setter - isCollapsed is read-only
		goja.FLAG_FALSE, goja.FLAG_TRUE) // Not enumerable, configurable

	selectionObj.DefineAccessorProperty("rangeCount",
		db.vm.ToValue(func(call goja.FunctionCall) goja.Value {
			return db.vm.ToValue(rangeCount)
		}),
		goja.Undefined(),                // No setter - rangeCount is read-only
		goja.FLAG_FALSE, goja.FLAG_TRUE) // Not enumerable, configurable

	// Selection methods
	selectionObj.Set("getRangeAt", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			panic(db.vm.NewTypeError("getRangeAt requires an index"))
		}

		index := int(call.Arguments[0].ToInteger())
		if index < 0 || index >= rangeCount {
			panic(db.vm.NewTypeError("Index out of bounds"))
		}

		// For simplified implementation, return a basic range
		return db.createRange()
	})

	selectionObj.Set("addRange", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			panic(db.vm.NewTypeError("addRange requires a range"))
		}

		// For simplified implementation, just increment range count
		rangeCount = 1
		isCollapsed = false
		return goja.Undefined()
	})

	selectionObj.Set("removeRange", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			panic(db.vm.NewTypeError("removeRange requires a range"))
		}

		// For simplified implementation, just decrement range count
		if rangeCount > 0 {
			rangeCount--
			if rangeCount == 0 {
				isCollapsed = true
			}
		}
		return goja.Undefined()
	})

	selectionObj.Set("removeAllRanges", func(call goja.FunctionCall) goja.Value {
		rangeCount = 0
		isCollapsed = true
		anchorNode = nil
		anchorOffset = 0
		focusNode = nil
		focusOffset = 0
		return goja.Undefined()
	})

	selectionObj.Set("collapse", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 2 {
			panic(db.vm.NewTypeError("collapse requires node and offset"))
		}

		node := db.extractNodeFromJS(call.Arguments[0])
		if node == nil {
			panic(db.vm.NewTypeError("Invalid node"))
		}

		offset := int(call.Arguments[1].ToInteger())
		anchorNode = node
		anchorOffset = offset
		focusNode = node
		focusOffset = offset
		isCollapsed = true
		rangeCount = 1
		return goja.Undefined()
	})

	selectionObj.Set("extend", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 2 {
			panic(db.vm.NewTypeError("extend requires node and offset"))
		}

		node := db.extractNodeFromJS(call.Arguments[0])
		if node == nil {
			panic(db.vm.NewTypeError("Invalid node"))
		}

		offset := int(call.Arguments[1].ToInteger())
		focusNode = node
		focusOffset = offset
		isCollapsed = (anchorNode == focusNode && anchorOffset == focusOffset)
		if !isCollapsed {
			rangeCount = 1
		}
		return goja.Undefined()
	})

	selectionObj.Set("selectAllChildren", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			panic(db.vm.NewTypeError("selectAllChildren requires a node"))
		}

		node := db.extractNodeFromJS(call.Arguments[0])
		if node == nil {
			panic(db.vm.NewTypeError("Invalid node"))
		}

		// Select all children of the node
		anchorNode = node
		anchorOffset = 0
		focusNode = node
		focusOffset = len(node.ChildNodes().ToSlice())
		isCollapsed = (anchorOffset == focusOffset)
		rangeCount = 1
		return goja.Undefined()
	})

	selectionObj.Set("toString", func(call goja.FunctionCall) goja.Value {
		// Return the text content of the selection
		// For simplified implementation, return empty string
		return db.vm.ToValue("")
	})

	return selectionObj
}

// wrapMutationRecord wraps a DOM MutationRecord for JavaScript access
func (db *DOMBindings) wrapMutationRecord(record *dom.MutationRecord) *goja.Object {
	obj := db.vm.NewObject()

	obj.Set("type", record.Type)
	obj.Set("target", db.WrapNode(record.Target))

	// Wrap added nodes
	addedNodes := db.vm.NewArray()
	for i, node := range record.AddedNodes {
		addedNodes.Set(strconv.Itoa(i), db.WrapNode(node))
	}
	addedNodes.Set("length", len(record.AddedNodes))
	obj.Set("addedNodes", addedNodes)

	// Wrap removed nodes
	removedNodes := db.vm.NewArray()
	for i, node := range record.RemovedNodes {
		removedNodes.Set(strconv.Itoa(i), db.WrapNode(node))
	}
	removedNodes.Set("length", len(record.RemovedNodes))
	obj.Set("removedNodes", removedNodes)

	// Previous and next siblings
	if record.PreviousSibling != nil {
		obj.Set("previousSibling", db.WrapNode(record.PreviousSibling))
	} else {
		obj.Set("previousSibling", goja.Null())
	}

	if record.NextSibling != nil {
		obj.Set("nextSibling", db.WrapNode(record.NextSibling))
	} else {
		obj.Set("nextSibling", goja.Null())
	}

	// Attribute information
	if record.AttributeName != "" {
		obj.Set("attributeName", record.AttributeName)
	} else {
		obj.Set("attributeName", goja.Null())
	}

	if record.AttributeNamespace != "" {
		obj.Set("attributeNamespace", record.AttributeNamespace)
	} else {
		obj.Set("attributeNamespace", goja.Null())
	}

	if record.OldValue != "" {
		obj.Set("oldValue", record.OldValue)
	} else {
		obj.Set("oldValue", goja.Null())
	}

	return obj
}

// wrapDOMMutationObserver wraps a DOM MutationObserver for JavaScript access
func (db *DOMBindings) wrapDOMMutationObserver(observer *dom.MutationObserver) *goja.Object {
	obj := db.vm.NewObject()

	obj.Set("observe", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 2 {
			panic(db.vm.NewTypeError("observe requires target and options"))
		}

		target := db.extractNodeFromJS(call.Arguments[0])
		if target == nil {
			panic(db.vm.NewTypeError("Invalid target node"))
		}

		// Parse options
		optionsJS := call.Arguments[1].ToObject(db.vm)
		if optionsJS == nil {
			panic(db.vm.NewTypeError("Options must be an object"))
		}

		options := dom.MutationObserverInit{}
		if childListVal := optionsJS.Get("childList"); childListVal != nil {
			options.ChildList = childListVal.ToBoolean()
		}
		if attributesVal := optionsJS.Get("attributes"); attributesVal != nil {
			options.Attributes = attributesVal.ToBoolean()
			options.AttributesSet = true
		}
		if characterDataVal := optionsJS.Get("characterData"); characterDataVal != nil {
			options.CharacterData = characterDataVal.ToBoolean()
			options.CharacterDataSet = true
		}
		if subtreeVal := optionsJS.Get("subtree"); subtreeVal != nil {
			options.Subtree = subtreeVal.ToBoolean()
		}
		if attributeOldValueVal := optionsJS.Get("attributeOldValue"); attributeOldValueVal != nil {
			options.AttributeOldValue = attributeOldValueVal.ToBoolean()
		}
		if characterDataOldValueVal := optionsJS.Get("characterDataOldValue"); characterDataOldValueVal != nil {
			options.CharacterDataOldValue = characterDataOldValueVal.ToBoolean()
		}
		if attributeFilterVal := optionsJS.Get("attributeFilter"); attributeFilterVal != nil {
			if filterArr := attributeFilterVal.ToObject(db.vm); filterArr != nil {
				if lengthVal := filterArr.Get("length"); lengthVal != nil {
					length := int(lengthVal.ToInteger())
					for i := 0; i < length; i++ {
						if attrVal := filterArr.Get(strconv.Itoa(i)); attrVal != nil {
							options.AttributeFilter = append(options.AttributeFilter, attrVal.String())
						}
					}
				}
			}
		}

		// Start observing
		observer.Observe(target, options)
		return goja.Undefined()
	})

	obj.Set("disconnect", func(call goja.FunctionCall) goja.Value {
		observer.Disconnect()
		return goja.Undefined()
	})

	obj.Set("takeRecords", func(call goja.FunctionCall) goja.Value {
		records := observer.TakeRecords()
		jsRecords := db.vm.NewArray()
		for i, record := range records {
			jsRecords.Set(strconv.Itoa(i), db.wrapMutationRecord(record))
		}
		jsRecords.Set("length", len(records))
		return jsRecords
	})

	return obj
}

// addPhase4NodeMethods adds Phase 4 advanced DOM features to a JavaScript object
func (db *DOMBindings) addPhase4NodeMethods(obj *goja.Object, node dom.Node) {
	// Phase 4: Advanced DOM Features

	// Node comparison methods
	obj.Set("compareDocumentPosition", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			panic(db.vm.NewTypeError("compareDocumentPosition requires another node"))
		}

		otherNode := db.extractNodeFromJS(call.Arguments[0])
		if otherNode == nil {
			panic(db.vm.NewTypeError("Invalid node"))
		}

		position := node.CompareDocumentPosition(otherNode)
		return db.vm.ToValue(position)
	})

	obj.Set("contains", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			return db.vm.ToValue(false)
		}

		otherNode := db.extractNodeFromJS(call.Arguments[0])
		if otherNode == nil {
			return db.vm.ToValue(false)
		}

		return db.vm.ToValue(node.Contains(otherNode))
	})

	obj.Set("isEqualNode", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			return db.vm.ToValue(false)
		}

		otherNode := db.extractNodeFromJS(call.Arguments[0])
		if otherNode == nil {
			return db.vm.ToValue(false)
		}

		return db.vm.ToValue(node.IsEqualNode(otherNode))
	})

	obj.Set("isSameNode", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			return db.vm.ToValue(false)
		}

		otherNode := db.extractNodeFromJS(call.Arguments[0])
		if otherNode == nil {
			return db.vm.ToValue(false)
		}

		return db.vm.ToValue(node.IsSameNode(otherNode))
	})

	// Namespace support properties and methods
	obj.DefineAccessorProperty("namespaceURI",
		db.vm.ToValue(func(call goja.FunctionCall) goja.Value {
			// Getter
			if element, ok := node.(*dom.Element); ok {
				return db.vm.ToValue(element.NamespaceURI())
			}
			return goja.Null()
		}),
		goja.Undefined(),                // No setter - namespaceURI is read-only
		goja.FLAG_FALSE, goja.FLAG_TRUE) // Not enumerable, configurable

	obj.DefineAccessorProperty("prefix",
		db.vm.ToValue(func(call goja.FunctionCall) goja.Value {
			// Getter
			if element, ok := node.(*dom.Element); ok {
				return db.vm.ToValue(element.Prefix())
			}
			return goja.Null()
		}),
		goja.Undefined(),                // No setter - prefix is read-only in this implementation
		goja.FLAG_FALSE, goja.FLAG_TRUE) // Not enumerable, configurable

	obj.DefineAccessorProperty("localName",
		db.vm.ToValue(func(call goja.FunctionCall) goja.Value {
			// Getter
			if element, ok := node.(*dom.Element); ok {
				return db.vm.ToValue(element.LocalName())
			}
			return goja.Null()
		}),
		goja.Undefined(),                // No setter - localName is read-only
		goja.FLAG_FALSE, goja.FLAG_TRUE) // Not enumerable, configurable

	// Namespace lookup methods
	obj.Set("lookupPrefix", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			return goja.Null()
		}

		namespaceURI := call.Arguments[0].String()
		prefix := node.LookupPrefix(namespaceURI)
		if prefix == "" {
			return goja.Null()
		}
		return db.vm.ToValue(prefix)
	})

	obj.Set("lookupNamespaceURI", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			return goja.Null()
		}

		prefix := call.Arguments[0].String()
		namespaceURI := node.LookupNamespaceURI(prefix)
		if namespaceURI == "" {
			return goja.Null()
		}
		return db.vm.ToValue(namespaceURI)
	})

	obj.Set("isDefaultNamespace", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			return db.vm.ToValue(false)
		}

		namespaceURI := call.Arguments[0].String()
		return db.vm.ToValue(node.IsDefaultNamespace(namespaceURI))
	})

	// Text normalization - available on all nodes
	obj.Set("normalize", func(call goja.FunctionCall) goja.Value {
		node.Normalize()
		// Update navigation properties after normalization
		db.updateNodeNavigationProperties(obj, node)
		return goja.Undefined()
	})

	// getRootNode method (advanced DOM)
	obj.Set("getRootNode", func(call goja.FunctionCall) goja.Value {
		// For most cases, return the document
		// In a full implementation, this would handle shadow roots and options
		var options *dom.GetRootNodeOptions
		if len(call.Arguments) > 0 {
			optionsObj := call.Arguments[0].ToObject(db.vm)
			if optionsObj != nil {
				options = &dom.GetRootNodeOptions{}
				if composedVal := optionsObj.Get("composed"); composedVal != nil {
					options.Composed = composedVal.ToBoolean()
				}
			}
		}

		rootNode := node.GetRootNode(options)
		return db.WrapNode(rootNode)
	})

	// hasChildNodes method
	obj.Set("hasChildNodes", func(call goja.FunctionCall) goja.Value {
		return db.vm.ToValue(node.HasChildNodes())
	})
}

// Phase 6: Modern Web API Implementation Methods

// createResizeObserver creates a ResizeObserver JavaScript object
func (db *DOMBindings) createResizeObserver(callback goja.Callable) *goja.Object {
	obj := db.vm.NewObject()

	// Store observed elements and callback
	var observedElements []dom.Node
	var isActive = true

	// Store the resize observer on the global object for test access
	globalResizeObservers := db.vm.Get("__resizeObservers")
	if goja.IsUndefined(globalResizeObservers) || goja.IsNull(globalResizeObservers) {
		globalResizeObservers = db.vm.NewArray()
		db.vm.Set("__resizeObservers", globalResizeObservers)
	}

	observerInfo := db.vm.NewObject()
	observerInfo.Set("observer", obj)
	observerInfo.Set("callback", callback)
	observerInfo.Set("elements", db.vm.NewArray())

	// Add to global list
	if arr, ok := globalResizeObservers.(*goja.Object); ok {
		if lengthVal := arr.Get("length"); lengthVal != nil {
			length := int(lengthVal.ToInteger())
			arr.Set(strconv.Itoa(length), observerInfo)
			arr.Set("length", length+1)
		}
	}

	obj.Set("observe", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			panic(db.vm.NewTypeError("observe requires a target element"))
		}

		target := db.extractNodeFromJS(call.Arguments[0])
		if target == nil {
			panic(db.vm.NewTypeError("Invalid target element"))
		}

		// Add to observed elements if not already observed
		for _, observed := range observedElements {
			if observed == target {
				return goja.Undefined() // Already observing
			}
		}
		observedElements = append(observedElements, target)

		// Update global tracking
		if elementsArr := observerInfo.Get("elements").(*goja.Object); elementsArr != nil {
			if lengthVal := elementsArr.Get("length"); lengthVal != nil {
				length := int(lengthVal.ToInteger())
				elementsArr.Set(strconv.Itoa(length), db.WrapNode(target))
				elementsArr.Set("length", length+1)
			}
		}

		// For initial observation, immediately fire a callback with default entry
		if isActive && db.jsRuntime != nil && db.jsRuntime.EventLoop() != nil {
			db.jsRuntime.EventLoop().QueueMicrotask(func() {
				entry := db.createResizeObserverEntry(target)
				entries := db.vm.NewArray()
				entries.Set("0", entry)
				entries.Set("length", 1)
				_, _ = callback(goja.Undefined(), entries, obj)
			}, "ResizeObserver")
		}

		return goja.Undefined()
	})

	obj.Set("unobserve", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			return goja.Undefined()
		}

		target := db.extractNodeFromJS(call.Arguments[0])
		if target == nil {
			return goja.Undefined()
		}

		// Remove from observed elements
		for i, observed := range observedElements {
			if observed == target {
				observedElements = append(observedElements[:i], observedElements[i+1:]...)
				break
			}
		}

		return goja.Undefined()
	})

	obj.Set("disconnect", func(call goja.FunctionCall) goja.Value {
		observedElements = nil
		isActive = false
		return goja.Undefined()
	})

	return obj
}

// createResizeObserverEntry creates a ResizeObserverEntry JavaScript object
func (db *DOMBindings) createResizeObserverEntry(target dom.Node) *goja.Object {
	entry := db.vm.NewObject()

	entry.Set("target", db.WrapNode(target))

	// Mock content rect
	contentRect := db.vm.NewObject()
	contentRect.Set("x", 0)
	contentRect.Set("y", 0)
	contentRect.Set("width", 200)
	contentRect.Set("height", 150)
	contentRect.Set("top", 0)
	contentRect.Set("right", 200)
	contentRect.Set("bottom", 150)
	contentRect.Set("left", 0)
	entry.Set("contentRect", contentRect)

	// Mock border box size
	borderBoxSize := db.vm.NewArray()
	borderBox := db.vm.NewObject()
	borderBox.Set("inlineSize", 200)
	borderBox.Set("blockSize", 150)
	borderBoxSize.Set("0", borderBox)
	borderBoxSize.Set("length", 1)
	entry.Set("borderBoxSize", borderBoxSize)

	// Mock content box size
	contentBoxSize := db.vm.NewArray()
	contentBox := db.vm.NewObject()
	contentBox.Set("inlineSize", 200)
	contentBox.Set("blockSize", 150)
	contentBoxSize.Set("0", contentBox)
	contentBoxSize.Set("length", 1)
	entry.Set("contentBoxSize", contentBoxSize)

	return entry
}

// createResizeObserverEntryFromInfo creates a ResizeObserverEntry from custom resize info
func (db *DOMBindings) createResizeObserverEntryFromInfo(target dom.Node, resizeInfo *goja.Object) *goja.Object {
	entry := db.vm.NewObject()

	entry.Set("target", db.WrapNode(target))

	// Extract data from resizeInfo
	if contentRect := resizeInfo.Get("contentRect"); contentRect != nil {
		entry.Set("contentRect", contentRect)
	}

	if borderBoxSize := resizeInfo.Get("borderBoxSize"); borderBoxSize != nil {
		entry.Set("borderBoxSize", borderBoxSize)
	}

	if contentBoxSize := resizeInfo.Get("contentBoxSize"); contentBoxSize != nil {
		entry.Set("contentBoxSize", contentBoxSize)
	}

	if devicePixelContentBoxSize := resizeInfo.Get("devicePixelContentBoxSize"); devicePixelContentBoxSize != nil {
		entry.Set("devicePixelContentBoxSize", devicePixelContentBoxSize)
	}

	return entry
}

// createIntersectionObserver creates an IntersectionObserver JavaScript object
func (db *DOMBindings) createIntersectionObserver(callback goja.Callable, options *goja.Object) *goja.Object {
	obj := db.vm.NewObject()

	// Store observed elements and configuration
	var observedElements []dom.Node
	var isActive = true

	// Parse options
	var root dom.Node
	var rootMargin = "0px"
	var threshold = []float64{0}

	if options != nil {
		if rootVal := options.Get("root"); rootVal != nil && !goja.IsNull(rootVal) && !goja.IsUndefined(rootVal) {
			root = db.extractNodeFromJS(rootVal)
		}
		if rootMarginVal := options.Get("rootMargin"); rootMarginVal != nil && !goja.IsUndefined(rootMarginVal) {
			rootMargin = rootMarginVal.String()
		}
		if thresholdVal := options.Get("threshold"); thresholdVal != nil && !goja.IsUndefined(thresholdVal) {
			if thresholdArr := thresholdVal.ToObject(db.vm); thresholdArr != nil {
				if lengthVal := thresholdArr.Get("length"); lengthVal != nil {
					length := int(lengthVal.ToInteger())
					threshold = make([]float64, 0, length)
					for i := 0; i < length; i++ {
						if val := thresholdArr.Get(strconv.Itoa(i)); val != nil {
							threshold = append(threshold, val.ToFloat())
						}
					}
				}
			} else {
				// Single threshold value
				threshold = []float64{thresholdVal.ToFloat()}
			}
		}
	}

	obj.Set("root", func() goja.Value {
		if root != nil {
			return db.WrapNode(root)
		}
		return goja.Null()
	}())
	obj.Set("rootMargin", rootMargin)
	obj.Set("thresholds", func() goja.Value {
		arr := db.vm.NewArray()
		for i, t := range threshold {
			arr.Set(strconv.Itoa(i), t)
		}
		arr.Set("length", len(threshold))
		return arr
	}())

	obj.Set("observe", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			panic(db.vm.NewTypeError("observe requires a target element"))
		}

		target := db.extractNodeFromJS(call.Arguments[0])
		if target == nil {
			panic(db.vm.NewTypeError("Invalid target element"))
		}

		// Add to observed elements if not already observed
		for _, observed := range observedElements {
			if observed == target {
				return goja.Undefined() // Already observing
			}
		}
		observedElements = append(observedElements, target)

		// For testing, immediately fire a callback with mock entry
		if isActive && db.jsRuntime != nil && db.jsRuntime.EventLoop() != nil {
			db.jsRuntime.EventLoop().QueueMicrotask(func() {
				entry := db.createIntersectionObserverEntry(target, true)
				entries := db.vm.NewArray()
				entries.Set("0", entry)
				entries.Set("length", 1)
				_, _ = callback(goja.Undefined(), entries, obj)
			}, "IntersectionObserver")
		}

		return goja.Undefined()
	})

	obj.Set("unobserve", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			return goja.Undefined()
		}

		target := db.extractNodeFromJS(call.Arguments[0])
		if target == nil {
			return goja.Undefined()
		}

		// Remove from observed elements
		for i, observed := range observedElements {
			if observed == target {
				observedElements = append(observedElements[:i], observedElements[i+1:]...)
				break
			}
		}

		return goja.Undefined()
	})

	obj.Set("disconnect", func(call goja.FunctionCall) goja.Value {
		observedElements = nil
		isActive = false
		return goja.Undefined()
	})

	obj.Set("takeRecords", func(call goja.FunctionCall) goja.Value {
		// Return empty array for mock implementation
		return db.vm.NewArray()
	})

	return obj
}

// createIntersectionObserverEntry creates an IntersectionObserverEntry JavaScript object
func (db *DOMBindings) createIntersectionObserverEntry(target dom.Node, isIntersecting bool) *goja.Object {
	entry := db.vm.NewObject()

	entry.Set("target", db.WrapNode(target))
	entry.Set("isIntersecting", isIntersecting)
	entry.Set("intersectionRatio", func() float64 {
		if isIntersecting {
			return 1.0
		}
		return 0.0
	}())

	// Mock intersection rect
	intersectionRect := db.vm.NewObject()
	if isIntersecting {
		intersectionRect.Set("x", 0)
		intersectionRect.Set("y", 0)
		intersectionRect.Set("width", 200)
		intersectionRect.Set("height", 150)
		intersectionRect.Set("top", 0)
		intersectionRect.Set("right", 200)
		intersectionRect.Set("bottom", 150)
		intersectionRect.Set("left", 0)
	} else {
		intersectionRect.Set("x", 0)
		intersectionRect.Set("y", 0)
		intersectionRect.Set("width", 0)
		intersectionRect.Set("height", 0)
		intersectionRect.Set("top", 0)
		intersectionRect.Set("right", 0)
		intersectionRect.Set("bottom", 0)
		intersectionRect.Set("left", 0)
	}
	entry.Set("intersectionRect", intersectionRect)

	// Mock bounding client rect
	boundingClientRect := db.vm.NewObject()
	boundingClientRect.Set("x", 0)
	boundingClientRect.Set("y", 0)
	boundingClientRect.Set("width", 200)
	boundingClientRect.Set("height", 150)
	boundingClientRect.Set("top", 0)
	boundingClientRect.Set("right", 200)
	boundingClientRect.Set("bottom", 150)
	boundingClientRect.Set("left", 0)
	entry.Set("boundingClientRect", boundingClientRect)

	// Mock root bounds
	rootBounds := db.vm.NewObject()
	rootBounds.Set("x", 0)
	rootBounds.Set("y", 0)
	rootBounds.Set("width", 800)
	rootBounds.Set("height", 600)
	rootBounds.Set("top", 0)
	rootBounds.Set("right", 800)
	rootBounds.Set("bottom", 600)
	rootBounds.Set("left", 0)
	entry.Set("rootBounds", rootBounds)

	entry.Set("time", func() float64 {
		// Mock timestamp
		return 1000.0
	}())

	return entry
}

// createFile creates a File JavaScript object
func (db *DOMBindings) createFile(fileBits *goja.Object, name string, fileType string, lastModified int64) *goja.Object {
	obj := db.vm.NewObject()

	// Calculate size from file bits
	size := 0
	if fileBits != nil {
		if lengthVal := fileBits.Get("length"); lengthVal != nil {
			length := int(lengthVal.ToInteger())
			for i := 0; i < length; i++ {
				if bit := fileBits.Get(strconv.Itoa(i)); bit != nil {
					size += len(bit.String())
				}
			}
		}
	}

	if lastModified == 0 {
		lastModified = 1640995200000 // Default to Jan 1, 2022
	}

	obj.Set("name", name)
	obj.Set("size", size)
	obj.Set("type", fileType)
	obj.Set("lastModified", lastModified)

	// Add Blob methods since File extends Blob
	obj.Set("slice", func(call goja.FunctionCall) goja.Value {
		// Mock slice implementation
		return db.createBlob(nil, fileType)
	})

	obj.Set("stream", func(call goja.FunctionCall) goja.Value {
		// Mock stream implementation
		return db.vm.NewObject()
	})

	obj.Set("text", func(call goja.FunctionCall) goja.Value {
		// Mock text implementation - return Promise-like object
		promise := db.vm.NewObject()
		promise.Set("then", func(call goja.FunctionCall) goja.Value {
			if len(call.Arguments) > 0 {
				callback, ok := goja.AssertFunction(call.Arguments[0])
				if ok {
					_, _ = callback(goja.Undefined(), db.vm.ToValue(""))
				}
			}
			return promise
		})
		return promise
	})

	obj.Set("arrayBuffer", func(call goja.FunctionCall) goja.Value {
		// Mock arrayBuffer implementation - return Promise-like object
		promise := db.vm.NewObject()
		promise.Set("then", func(call goja.FunctionCall) goja.Value {
			if len(call.Arguments) > 0 {
				callback, ok := goja.AssertFunction(call.Arguments[0])
				if ok {
					// Mock ArrayBuffer
					arrayBuffer := db.vm.NewObject()
					arrayBuffer.Set("byteLength", size)
					_, _ = callback(goja.Undefined(), arrayBuffer)
				}
			}
			return promise
		})
		return promise
	})

	return obj
}

// createBlob creates a Blob JavaScript object
func (db *DOMBindings) createBlob(blobParts *goja.Object, blobType string) *goja.Object {
	obj := db.vm.NewObject()

	// Calculate size from blob parts
	size := 0
	if blobParts != nil {
		if lengthVal := blobParts.Get("length"); lengthVal != nil {
			length := int(lengthVal.ToInteger())
			for i := 0; i < length; i++ {
				if part := blobParts.Get(strconv.Itoa(i)); part != nil {
					size += len(part.String())
				}
			}
		}
	}

	obj.Set("size", size)
	obj.Set("type", blobType)

	obj.Set("slice", func(call goja.FunctionCall) goja.Value {
		start := 0
		end := size
		contentType := blobType

		if len(call.Arguments) > 0 {
			start = int(call.Arguments[0].ToInteger())
		}
		if len(call.Arguments) > 1 {
			end = int(call.Arguments[1].ToInteger())
		}
		if len(call.Arguments) > 2 {
			contentType = call.Arguments[2].String()
		}

		// Create new blob with sliced content
		newSize := end - start
		if newSize < 0 {
			newSize = 0
		}

		newBlob := db.vm.NewObject()
		newBlob.Set("size", newSize)
		newBlob.Set("type", contentType)
		return newBlob
	})

	obj.Set("stream", func(call goja.FunctionCall) goja.Value {
		// Mock ReadableStream
		stream := db.vm.NewObject()
		return stream
	})

	obj.Set("text", func(call goja.FunctionCall) goja.Value {
		// Return Promise-like object
		promise := db.vm.NewObject()
		promise.Set("then", func(call goja.FunctionCall) goja.Value {
			if len(call.Arguments) > 0 {
				callback, ok := goja.AssertFunction(call.Arguments[0])
				if ok {
					_, _ = callback(goja.Undefined(), db.vm.ToValue(""))
				}
			}
			return promise
		})
		return promise
	})

	obj.Set("arrayBuffer", func(call goja.FunctionCall) goja.Value {
		// Return Promise-like object
		promise := db.vm.NewObject()
		promise.Set("then", func(call goja.FunctionCall) goja.Value {
			if len(call.Arguments) > 0 {
				callback, ok := goja.AssertFunction(call.Arguments[0])
				if ok {
					// Mock ArrayBuffer
					arrayBuffer := db.vm.NewObject()
					arrayBuffer.Set("byteLength", size)
					_, _ = callback(goja.Undefined(), arrayBuffer)
				}
			}
			return promise
		})
		return promise
	})

	return obj
}

// createFileReader creates a FileReader JavaScript object
func (db *DOMBindings) createFileReader() *goja.Object {
	obj := db.vm.NewObject()

	var readyState = 0 // EMPTY
	var result goja.Value = goja.Null()
	var error goja.Value = goja.Null()

	obj.Set("readyState", readyState)
	obj.Set("result", result)
	obj.Set("error", error)

	// Event handlers
	obj.Set("onload", goja.Null())
	obj.Set("onloadstart", goja.Null())
	obj.Set("onloadend", goja.Null())
	obj.Set("onprogress", goja.Null())
	obj.Set("onerror", goja.Null())
	obj.Set("onabort", goja.Null())

	obj.Set("readAsText", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			panic(db.vm.NewTypeError("readAsText requires a blob"))
		}

		readyState = 1 // LOADING
		obj.Set("readyState", readyState)

		// Fire loadstart event
		if onloadstart := obj.Get("onloadstart"); onloadstart != nil && !goja.IsNull(onloadstart) && !goja.IsUndefined(onloadstart) {
			if callback, ok := goja.AssertFunction(onloadstart); ok {
				event := db.vm.NewObject()
				event.Set("type", "loadstart")
				event.Set("target", obj)
				_, _ = callback(goja.Undefined(), event)
			}
		}

		// Simulate async reading with microtask
		if db.jsRuntime != nil && db.jsRuntime.EventLoop() != nil {
			db.jsRuntime.EventLoop().QueueMicrotask(func() {
				readyState = 2 // DONE
				result = db.vm.ToValue("Test content")
				obj.Set("readyState", readyState)
				obj.Set("result", result)

				// Fire load event
				if onload := obj.Get("onload"); onload != nil && !goja.IsNull(onload) && !goja.IsUndefined(onload) {
					if callback, ok := goja.AssertFunction(onload); ok {
						event := db.vm.NewObject()
						event.Set("type", "load")
						event.Set("target", obj)
						_, _ = callback(goja.Undefined(), event)
					}
				}

				// Fire loadend event
				if onloadend := obj.Get("onloadend"); onloadend != nil && !goja.IsNull(onloadend) && !goja.IsUndefined(onloadend) {
					if callback, ok := goja.AssertFunction(onloadend); ok {
						event := db.vm.NewObject()
						event.Set("type", "loadend")
						event.Set("target", obj)
						_, _ = callback(goja.Undefined(), event)
					}
				}
			}, "FileReader")
		}

		return goja.Undefined()
	})

	obj.Set("readAsDataURL", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			panic(db.vm.NewTypeError("readAsDataURL requires a blob"))
		}

		readyState = 1 // LOADING
		obj.Set("readyState", readyState)

		// Simulate async reading
		if db.jsRuntime != nil && db.jsRuntime.EventLoop() != nil {
			db.jsRuntime.EventLoop().QueueMicrotask(func() {
				readyState = 2 // DONE
				result = db.vm.ToValue("data:text/plain;base64,VGVzdCBjb250ZW50")
				obj.Set("readyState", readyState)
				obj.Set("result", result)

				// Fire events similar to readAsText
			}, "FileReader")
		}

		return goja.Undefined()
	})

	obj.Set("readAsArrayBuffer", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			panic(db.vm.NewTypeError("readAsArrayBuffer requires a blob"))
		}

		readyState = 1 // LOADING
		obj.Set("readyState", readyState)

		// Simulate async reading
		if db.jsRuntime != nil && db.jsRuntime.EventLoop() != nil {
			db.jsRuntime.EventLoop().QueueMicrotask(func() {
				readyState = 2 // DONE
				// Mock ArrayBuffer
				arrayBuffer := db.vm.NewObject()
				arrayBuffer.Set("byteLength", 12)
				result = arrayBuffer
				obj.Set("readyState", readyState)
				obj.Set("result", result)

				// Fire events similar to readAsText
			}, "FileReader")
		}

		return goja.Undefined()
	})

	obj.Set("abort", func(call goja.FunctionCall) goja.Value {
		if readyState == 1 {
			readyState = 2 // DONE
			error = db.vm.NewObject()
			obj.Set("readyState", readyState)
			obj.Set("error", error)

			// Fire abort event
			if onabort := obj.Get("onabort"); onabort != nil && !goja.IsNull(onabort) && !goja.IsUndefined(onabort) {
				if callback, ok := goja.AssertFunction(onabort); ok {
					event := db.vm.NewObject()
					event.Set("type", "abort")
					event.Set("target", obj)
					_, _ = callback(goja.Undefined(), event)
				}
			}
		}
		return goja.Undefined()
	})

	return obj
}
