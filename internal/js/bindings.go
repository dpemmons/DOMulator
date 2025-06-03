package js

import (
	"strconv"

	"github.com/dop251/goja"
	"github.com/dpemmons/DOMulator/internal/browser/url"
	"github.com/dpemmons/DOMulator/internal/css"
	"github.com/dpemmons/DOMulator/internal/dom"
)

// DOMBindings manages the binding between Go DOM objects and JavaScript
type DOMBindings struct {
	vm        *goja.Runtime
	document  *dom.Document
	nodeCache map[dom.Node]*goja.Object // Cache to maintain object identity
}

// NewDOMBindings creates a new DOM bindings instance
func NewDOMBindings(vm *goja.Runtime, document *dom.Document) *DOMBindings {
	return &DOMBindings{
		vm:        vm,
		document:  document,
		nodeCache: make(map[dom.Node]*goja.Object),
	}
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
		elements := db.document.GetElementsByTagName(tagName)
		return db.WrapNodeList(elements)
	})

	doc.Set("getElementsByClassName", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			return db.vm.NewArray()
		}

		className := call.Arguments[0].String()
		elements := db.document.GetElementsByClassName(className)
		return db.WrapNodeList(elements)
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
		for _, node := range nodeList {
			if element, ok := node.(*dom.Element); ok {
				elements = append(elements, element)
			}
		}
		return db.WrapNodeList(elements)
	})

	// Document properties - avoid infinite recursion by not wrapping elements here
	doc.Set("documentElement", goja.Null())
	doc.Set("body", goja.Null())
	doc.Set("head", goja.Null())

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
	elem.Set("innerHTML", element.InnerHTML())
	elem.Set("outerHTML", element.OuterHTML())

	// Dynamic textContent property - set current value
	elem.Set("textContent", element.TextContent())

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

	// Class methods
	elem.Set("className", element.GetAttribute("class"))

	// ID property
	elem.Set("id", element.GetAttribute("id"))

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
		for _, node := range nodeList {
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
		elements := element.GetElementsByTagName(tagName)
		return db.WrapNodeList(elements)
	})

	elem.Set("getElementsByClassName", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			return db.vm.NewArray()
		}

		className := call.Arguments[0].String()
		elements := element.GetElementsByClassName(className)
		return db.WrapNodeList(elements)
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

		// Update cached properties after DOM modification
		elem.Set("innerHTML", element.InnerHTML())
		elem.Set("outerHTML", element.OuterHTML())
		elem.Set("textContent", element.TextContent())

		// Update childNodes if this was afterbegin or beforeend
		if position == "afterbegin" || position == "beforeend" {
			children := db.vm.NewArray()
			childNodes := element.ChildNodes()
			for i, childNode := range childNodes {
				children.Set(strconv.Itoa(i), db.WrapNode(childNode))
			}
			children.Set("length", len(childNodes))
			elem.Set("childNodes", children)

			// Update navigation properties
			if element.FirstChild() != nil {
				elem.Set("firstChild", db.WrapNode(element.FirstChild()))
			} else {
				elem.Set("firstChild", goja.Null())
			}

			if element.LastChild() != nil {
				elem.Set("lastChild", db.WrapNode(element.LastChild()))
			} else {
				elem.Set("lastChild", goja.Null())
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

// WrapNodeList wraps a slice of elements as a JavaScript array-like object
func (db *DOMBindings) WrapNodeList(elements []*dom.Element) *goja.Object {
	arr := db.vm.NewArray()

	for i, element := range elements {
		arr.Set(strconv.Itoa(i), db.WrapElement(element))
	}

	arr.Set("length", len(elements))

	return arr
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

		node.AppendChild(child)

		// Update the childNodes array
		children := db.vm.NewArray()
		childNodes := node.ChildNodes()
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

		// Update textContent after DOM modification for elements
		if element, ok := node.(*dom.Element); ok {
			obj.Set("textContent", element.TextContent())
		}

		// Update the child's parentNode property
		childJS := call.Arguments[0].ToObject(db.vm)
		if childJS != nil {
			childJS.Set("parentNode", obj)
		}

		return call.Arguments[0] // Return the appended child
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
	obj.Set("firstChild", goja.Null())
	obj.Set("lastChild", goja.Null())
	obj.Set("nextSibling", goja.Null())
	obj.Set("previousSibling", goja.Null())

	// Children array - create initial empty array
	children := db.vm.NewArray()
	children.Set("length", len(node.ChildNodes()))
	obj.Set("childNodes", children)
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

		// Extract Go event from JavaScript wrapper
		event := db.extractEvent(call.Arguments[0])
		if event == nil {
			panic(db.vm.NewTypeError("Invalid event"))
		}

		return db.vm.ToValue(node.DispatchEvent(event))
	})
}

// WrapEvent wraps a DOM event for JavaScript access
func (db *DOMBindings) WrapEvent(event dom.Event) *goja.Object {
	evt := db.vm.NewObject()

	evt.Set("type", event.Type())
	evt.Set("target", db.WrapNode(event.Target()))
	evt.Set("currentTarget", db.WrapNode(event.CurrentTarget()))
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

// SetupBrowserAPIs sets up browser APIs like CustomEvent, URL, and URLSearchParams
func (db *DOMBindings) SetupBrowserAPIs() {
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

	// CustomEvent constructor stub
	db.vm.Set("CustomEvent", func(call goja.ConstructorCall) *goja.Object {
		obj := db.vm.NewObject()
		obj.Set("type", "")
		obj.Set("detail", goja.Null())
		return obj
	})
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
