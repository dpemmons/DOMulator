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

	// Event methods (inherited from Node)
	db.addEventMethods(doc, db.document)

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

	// Add node methods
	db.addNodeMethods(elem, element)

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
	obj.Set("nodeValue", node.NodeValue())
	obj.Set("textContent", db.getTextContent(node))

	switch n := node.(type) {
	case *dom.Text:
		obj.Set("data", n.NodeValue())
	case *dom.Comment:
		obj.Set("data", n.NodeValue())
	}

	// Add common node methods
	db.addNodeMethods(obj, node)

	// Add event methods
	db.addEventMethods(obj, node)

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

	// Navigation properties - avoid infinite recursion by setting to null initially
	obj.Set("parentNode", goja.Null())
	obj.Set("nextSibling", goja.Null())
	obj.Set("previousSibling", goja.Null())

	// Set up navigation properties properly
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
