package dom

import (
	"strings"
	"sync"

	"github.com/dop251/goja"
)

// Document represents the entire HTML document.
type Document struct {
	nodeImpl
	observerRegistry  *ObserverRegistry
	modificationTime  int64 // Counter incremented on each DOM modification
	modificationMutex sync.RWMutex

	// MutationObserver agent-level state per WHATWG DOM spec
	mutationObserverMicrotaskQueued bool
	pendingMutationObservers        map[*MutationObserver]bool // Set implementation
	signalSlots                     []*Element                 // For slotchange events
	mutationStateMutex              sync.Mutex                 // Protects mutation observer state

	// Add document-specific properties here
	// For example, a reference to the window object, if we implement one
	// defaultView *Window
}

// NewDocument creates a new Document node.
func NewDocument() *Document {
	doc := &Document{
		nodeImpl: nodeImpl{
			nodeType: DocumentNode,
			nodeName: "#document",
		},
		observerRegistry:         NewObserverRegistry(),
		pendingMutationObservers: make(map[*MutationObserver]bool),
		signalSlots:              make([]*Element, 0),
	}
	doc.ownerDocument = doc // A document is its own owner document
	doc.nodeImpl.self = doc // Set the self reference
	return doc
}

// OwnerDocument returns nil for documents per spec
func (d *Document) OwnerDocument() *Document {
	return nil
}

// CreateElement creates a new element with the given tag name
func (d *Document) CreateElement(tagName string) *Element {
	return NewElement(tagName, d)
}

// CreateElementNS creates a new element with the given namespace URI and qualified name
func (d *Document) CreateElementNS(namespaceURI, qualifiedName string) (*Element, error) {
	return NewElementNS(namespaceURI, qualifiedName, d)
}

// CreateTextNode creates a new text node with the given data
func (d *Document) CreateTextNode(data string) *Text {
	return NewText(data, d)
}

// CreateComment creates a new comment node with the given data
func (d *Document) CreateComment(data string) *Comment {
	return NewComment(data, d)
}

// CreateDocumentFragment creates a new document fragment
func (d *Document) CreateDocumentFragment() *DocumentFragment {
	return NewDocumentFragment(d)
}

// GetElementById returns the element with the specified ID
func (d *Document) GetElementById(id string) *Element {
	if id == "" {
		return nil // Don't search for empty IDs
	}
	var result *Element
	Traverse(d, func(node Node) bool {
		if elem, ok := node.(*Element); ok {
			if elem.GetAttribute("id") == id {
				result = elem
				return false // Stop traversal
			}
		}
		return true // Continue traversal
	})
	return result
}

// GetElementsByTagName returns a live HTMLCollection of all elements with the specified tag name
func (d *Document) GetElementsByTagName(tagName string) *HTMLCollection {
	return NewElementsByTagNameCollection(d, tagName)
}

// GetElementsByTagNameNS returns a live HTMLCollection of all elements with the specified namespace URI and local name
func (d *Document) GetElementsByTagNameNS(namespaceURI, localName string) *HTMLCollection {
	return NewElementsByTagNameNSCollection(d, namespaceURI, localName)
}

// GetElementsByClassName returns a live HTMLCollection of all elements with the specified class name
func (d *Document) GetElementsByClassName(className string) *HTMLCollection {
	return NewElementsByClassNameCollection(d, className)
}

// GetElementsByName returns a live HTMLCollection of all elements with the specified name attribute
func (d *Document) GetElementsByName(name string) *HTMLCollection {
	return NewElementsByNameCollection(d, name)
}

// DocumentElement returns the document element (usually <html>)
func (d *Document) DocumentElement() *Element {
	for _, child := range d.childNodes {
		if elem, ok := child.(*Element); ok && elem.TagName() == "html" {
			return elem
		}
	}
	return nil
}

// Body returns the body element
func (d *Document) Body() *Element {
	if docElem := d.DocumentElement(); docElem != nil {
		children := docElem.ChildNodes()
		for i := 0; i < children.Length(); i++ {
			child := children.Item(i)
			if elem, ok := child.(*Element); ok && elem.TagName() == "body" {
				return elem
			}
		}
	}
	return nil
}

// Head returns the head element
func (d *Document) Head() *Element {
	if docElem := d.DocumentElement(); docElem != nil {
		children := docElem.ChildNodes()
		for i := 0; i < children.Length(); i++ {
			child := children.Item(i)
			if elem, ok := child.(*Element); ok && elem.TagName() == "head" {
				return elem
			}
		}
	}
	return nil
}

// getObserverRegistry returns the document's observer registry
func (d *Document) getObserverRegistry() *ObserverRegistry {
	return d.observerRegistry
}

// getModificationTime returns the current modification time
func (d *Document) getModificationTime() int64 {
	d.modificationMutex.RLock()
	defer d.modificationMutex.RUnlock()
	return d.modificationTime
}

// incrementModificationTime increments the modification counter
func (d *Document) incrementModificationTime() {
	d.modificationMutex.Lock()
	defer d.modificationMutex.Unlock()
	d.modificationTime++
}

// Children returns a live HTMLCollection of all child elements
func (d *Document) Children() *HTMLCollection {
	return NewChildElementsCollection(d)
}

// FirstElementChild returns the first child that is an element; otherwise null
func (d *Document) FirstElementChild() *Element {
	for _, child := range d.childNodes {
		if child.NodeType() == ElementNode {
			return child.(*Element)
		}
	}
	return nil
}

// LastElementChild returns the last child that is an element; otherwise null
func (d *Document) LastElementChild() *Element {
	for i := len(d.childNodes) - 1; i >= 0; i-- {
		if d.childNodes[i].NodeType() == ElementNode {
			return d.childNodes[i].(*Element)
		}
	}
	return nil
}

// ChildElementCount returns the number of children of this that are elements
func (d *Document) ChildElementCount() int {
	count := 0
	for _, child := range d.childNodes {
		if child.NodeType() == ElementNode {
			count++
		}
	}
	return count
}

// Prepend inserts nodes before the first child of this document
func (d *Document) Prepend(nodes ...interface{}) error {
	// Convert nodes into a single node or document fragment
	convertedNode, err := convertNodesToNode(nodes, d)
	if err != nil {
		return err
	}

	if convertedNode == nil {
		return nil // Nothing to insert
	}

	// Pre-insert node before the first child
	if d.FirstChild() != nil {
		d.InsertBefore(convertedNode, d.FirstChild())
	} else {
		d.AppendChild(convertedNode)
	}

	return nil
}

// Append inserts nodes after the last child of this document
func (d *Document) Append(nodes ...interface{}) error {
	// Convert nodes into a single node or document fragment
	convertedNode, err := convertNodesToNode(nodes, d)
	if err != nil {
		return err
	}

	if convertedNode == nil {
		return nil // Nothing to insert
	}

	// Append node to this document
	d.AppendChild(convertedNode)
	return nil
}

// ReplaceChildren replaces all children of this document with the given nodes
func (d *Document) ReplaceChildren(nodes ...interface{}) error {
	// Convert nodes into a single node or document fragment
	convertedNode, err := convertNodesToNode(nodes, d)
	if err != nil {
		return err
	}

	// Ensure pre-insert validity (even though we're not inserting before a specific child)
	if convertedNode != nil {
		if err := d.validateHierarchy(convertedNode); err != nil {
			return err
		}
	}

	// Remove all existing children
	for len(d.childNodes) > 0 {
		d.RemoveChild(d.childNodes[0])
	}

	// Add the new node if any
	if convertedNode != nil {
		d.AppendChild(convertedNode)
	}

	return nil
}

// MoveBefore moves node into this document before child, preserving state
func (d *Document) MoveBefore(node Node, child Node) error {
	if node == nil {
		return NewNotFoundError("node cannot be null")
	}

	// Determine reference child
	referenceChild := child
	if referenceChild == node {
		// If reference is the node being moved, use its next sibling
		referenceChild = node.NextSibling()
	}

	// Validate hierarchy constraints
	if err := d.validateHierarchy(node); err != nil {
		return err
	}

	// Move the node (this preserves state by not removing first)
	if referenceChild != nil {
		d.InsertBefore(node, referenceChild)
	} else {
		d.AppendChild(node)
	}

	return nil
}

// QuerySelector returns the first element that is a descendant of this document that matches selectors
func (d *Document) QuerySelector(selectors string) *Element {
	// Simple implementation for now - supports basic selectors
	// TODO: Integrate with full CSS selector engine to avoid circular dependency
	var foundElement *Element
	Traverse(d, func(n Node) bool {
		if n == d {
			return true // Skip the root document itself
		}
		if elem, ok := n.(*Element); ok {
			if d.matchesSimpleSelector(elem, selectors) {
				foundElement = elem
				return false // Stop traversal
			}
		}
		return true // Continue traversal
	})
	return foundElement
}

// QuerySelectorAll returns all element descendants of this document that match selectors
func (d *Document) QuerySelectorAll(selectors string) *NodeList {
	// Simple implementation for now - supports basic selectors
	// TODO: Integrate with full CSS selector engine to avoid circular dependency
	var matchingNodes []Node
	Traverse(d, func(n Node) bool {
		if n == d {
			return true // Skip the root document itself
		}
		if elem, ok := n.(*Element); ok {
			if d.matchesSimpleSelector(elem, selectors) {
				matchingNodes = append(matchingNodes, elem)
			}
		}
		return true // Continue traversal
	})
	return NewNodeList(matchingNodes)
}

// matchesSimpleSelector checks if an element matches a basic CSS selector
func (d *Document) matchesSimpleSelector(elem *Element, selector string) bool {
	selector = strings.TrimSpace(selector)
	if selector == "" {
		return false
	}

	// Handle wildcard selector
	if selector == "*" {
		return true
	}

	// Parse simple selectors: tag, #id, .class, tag#id, tag.class, etc.
	remaining := selector
	var tagName, id string
	var classes []string

	// Extract ID (e.g., #myid)
	if idx := strings.Index(remaining, "#"); idx != -1 {
		idPart := remaining[idx+1:]
		if spaceIdx := strings.IndexAny(idPart, ". "); spaceIdx != -1 {
			id = idPart[:spaceIdx]
			remaining = remaining[:idx] + idPart[spaceIdx:]
		} else {
			id = idPart
			remaining = remaining[:idx]
		}
	}

	// Extract Classes (e.g., .myclass)
	classParts := strings.Split(remaining, ".")
	if len(classParts) > 1 {
		for _, class := range classParts[1:] {
			if class != "" {
				classes = append(classes, class)
			}
		}
		remaining = classParts[0]
	}

	// Extract Tag (the remaining part)
	tagName = strings.TrimSpace(remaining)

	// Check tag match
	if tagName != "" && !strings.EqualFold(tagName, elem.TagName()) {
		return false
	}

	// Check ID match
	if id != "" && elem.GetAttribute("id") != id {
		return false
	}

	// Check class matches
	if len(classes) > 0 {
		elemClasses := strings.Fields(elem.GetAttribute("class"))
		for _, requiredClass := range classes {
			found := false
			for _, elemClass := range elemClasses {
				if requiredClass == elemClass {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}
	}

	return true
}

// MutationObserver agent-level methods per WHATWG DOM specification

// queueMutationObserverMicrotask implements the "queue a mutation observer microtask" algorithm
func (d *Document) queueMutationObserverMicrotask() {
	d.mutationStateMutex.Lock()
	defer d.mutationStateMutex.Unlock()

	// Step 1: If already queued, return
	if d.mutationObserverMicrotaskQueued {
		return
	}

	// Step 2: Set flag to true
	d.mutationObserverMicrotaskQueued = true

	// Step 3: Queue microtask to notify observers
	// TODO: Integrate with event loop microtask queue when available
	// For now, we'll use a goroutine to simulate microtask behavior
	go func() {
		d.notifyMutationObservers()
	}()
}

// notifyMutationObservers implements the "notify mutation observers" algorithm
func (d *Document) notifyMutationObservers() {
	d.mutationStateMutex.Lock()

	// Step 1: Set flag to false
	d.mutationObserverMicrotaskQueued = false

	// Step 2: Clone pending mutation observers
	notifySet := make([]*MutationObserver, 0, len(d.pendingMutationObservers))
	for mo := range d.pendingMutationObservers {
		notifySet = append(notifySet, mo)
	}

	// Step 3: Empty pending mutation observers
	d.pendingMutationObservers = make(map[*MutationObserver]bool)

	// Step 4: Clone signal slots
	signalSet := make([]*Element, len(d.signalSlots))
	copy(signalSet, d.signalSlots)

	// Step 5: Empty signal slots
	d.signalSlots = d.signalSlots[:0]

	d.mutationStateMutex.Unlock()

	// Step 6: For each observer in notifySet
	for _, mo := range notifySet {
		// Step 6.1: Clone records and empty queue
		records := mo.TakeRecords()

		// Step 6.2: Remove transient registered observers
		mo.removeTransientObservers()

		// Step 6.3: If records not empty, invoke callback
		if len(records) > 0 {
			mo.callback(records, mo)
		}
	}

	// Step 7: Fire slotchange events
	for _, slot := range signalSet {
		// TODO: Implement slotchange event when event system supports it
		// For now, this is a placeholder for future slot implementation
		_ = slot
	}
}

// addPendingMutationObserver adds an observer to the pending set
func (d *Document) addPendingMutationObserver(observer *MutationObserver) {
	d.mutationStateMutex.Lock()
	defer d.mutationStateMutex.Unlock()
	d.pendingMutationObservers[observer] = true
}

// queueMutationRecord implements the "queue a mutation record" algorithm per WHATWG DOM spec
func (d *Document) queueMutationRecord(recordType string, target Node, name, namespace, oldValue string,
	addedNodes, removedNodes []Node, previousSibling, nextSibling Node) {

	// Step 1: Create interested observers map
	interestedObservers := make(map[*MutationObserver]string) // observer -> oldValue

	// Step 2: Get inclusive ancestors
	nodes := d.getInclusiveAncestors(target)

	// Step 3: For each node and registered observer
	for _, node := range nodes {
		registeredObservers := node.getRegisteredObservers()
		for _, registered := range registeredObservers {
			options := registered.Options

			// Check if observer is interested (exact spec conditions)
			if d.shouldNotifyForMutation(node, target, recordType, options, name, namespace) {
				mo := registered.Observer

				// Determine if old value needed
				if _, exists := interestedObservers[mo]; !exists {
					if (recordType == "attributes" && options.AttributeOldValue) ||
						(recordType == "characterData" && options.CharacterDataOldValue) {
						interestedObservers[mo] = oldValue
					} else {
						interestedObservers[mo] = "" // Interested but no old value
					}
				}
			}
		}
	}

	// Step 4: For each interested observer
	for observer, mappedOldValue := range interestedObservers {
		// Create mutation record
		record := &MutationRecord{
			Type:               recordType,
			Target:             target,
			AddedNodes:         addedNodes,
			RemovedNodes:       removedNodes,
			PreviousSibling:    previousSibling,
			NextSibling:        nextSibling,
			AttributeName:      name,
			AttributeNamespace: namespace,
			OldValue:           mappedOldValue,
		}

		// Queue to observer
		observer.queueMutationRecord(record)

		// Add to pending observers
		d.addPendingMutationObserver(observer)
	}

	// Step 5: Queue mutation observer microtask
	d.queueMutationObserverMicrotask()
}

// queueTreeMutationRecord implements "queue a tree mutation record"
func (d *Document) queueTreeMutationRecord(target Node, addedNodes, removedNodes []Node, previousSibling, nextSibling Node) {
	// Assert: either addedNodes or removedNodes is not empty
	if len(addedNodes) == 0 && len(removedNodes) == 0 {
		panic("queueTreeMutationRecord: either addedNodes or removedNodes must not be empty")
	}

	// Queue a mutation record of "childList"
	d.queueMutationRecord("childList", target, "", "", "", addedNodes, removedNodes, previousSibling, nextSibling)
}

// getInclusiveAncestors returns the target node and all its ancestors
func (d *Document) getInclusiveAncestors(target Node) []Node {
	var nodes []Node
	current := target
	for current != nil {
		nodes = append(nodes, current)
		current = current.ParentNode()
	}
	return nodes
}

// shouldNotifyForMutation determines if an observer should be notified based on spec conditions
func (d *Document) shouldNotifyForMutation(node, target Node, recordType string, options MutationObserverInit, name, namespace string) bool {
	// Check spec conditions - return false if any of these are true:

	// 1. node is not target and options["subtree"] is false
	if node != target && !options.Subtree {
		return false
	}

	// 2. type is "attributes" and options["attributes"] either does not exist or is false
	if recordType == "attributes" && !options.Attributes {
		return false
	}

	// 3. type is "attributes", options["attributeFilter"] exists, and options["attributeFilter"] does not contain name or namespace is non-null
	if recordType == "attributes" && len(options.AttributeFilter) > 0 {
		found := false
		for _, attr := range options.AttributeFilter {
			if attr == name {
				found = true
				break
			}
		}
		if !found || namespace != "" {
			return false
		}
	}

	// 4. type is "characterData" and options["characterData"] either does not exist or is false
	if recordType == "characterData" && !options.CharacterData {
		return false
	}

	// 5. type is "childList" and options["childList"] is false
	if recordType == "childList" && !options.ChildList {
		return false
	}

	return true
}

// CloneNode creates a copy of the document using the spec-compliant cloning implementation.
func (d *Document) CloneNode(deep bool) Node {
	// Use the spec-compliant cloning implementation
	return CloneNodeSpec(d, deep)
}

// toJS is a placeholder for JavaScript binding.
func (d *Document) toJS(vm *goja.Runtime) goja.Value {
	// This will be implemented later when integrating with Goja
	return vm.ToValue(nil)
}
