package dom

// EventHandler represents an event handler function type
type EventHandler func(event *Event)

// ShadowRootMode represents the mode of a shadow root
type ShadowRootMode string

const (
	ShadowRootModeOpen   ShadowRootMode = "open"
	ShadowRootModeClosed ShadowRootMode = "closed"
)

// SlotAssignmentMode represents the slot assignment mode
type SlotAssignmentMode string

const (
	SlotAssignmentNamed  SlotAssignmentMode = "named"
	SlotAssignmentManual SlotAssignmentMode = "manual"
)

// ShadowRootInit represents the init dictionary for creating shadow roots
type ShadowRootInit struct {
	Mode           ShadowRootMode
	SlotAssignment SlotAssignmentMode
	DelegatesFocus bool
	Clonable       bool
	Serializable   bool
}

// ShadowRoot represents a shadow root as defined in WHATWG DOM Section 4.8
type ShadowRoot struct {
	*DocumentFragment
	host           *Element
	mode           ShadowRootMode
	slotAssignment SlotAssignmentMode

	// WHATWG DOM Section 4.8 properties
	delegatesFocus                bool // Initially false
	clonable                      bool // Initially false
	serializable                  bool // Initially false
	availableToElementInternals   bool // Initially false
	declarative                   bool // Initially false
	keepCustomElementRegistryNull bool // Initially false

	// For manual slot assignment - maps slot to manually assigned slottables
	manualSlotMap map[*Slot][]*Element

	// DocumentOrShadowRoot mixin per WHATWG DOM Section 4.2.5
	customElementRegistry *CustomElementRegistry // Custom element registry

	// Event handlers
	onslotchange EventHandler // onslotchange event handler
}

// NewShadowRoot creates a new shadow root
func NewShadowRoot(host *Element, mode ShadowRootMode) *ShadowRoot {
	if host == nil {
		panic(NewInvalidStateError("Host element cannot be null"))
	}

	doc := host.OwnerDocument()
	if doc == nil {
		panic(NewInvalidStateError("Host element must have an owner document"))
	}

	sr := &ShadowRoot{
		DocumentFragment: NewDocumentFragment(doc),
		host:             host,
		mode:             mode,
		slotAssignment:   SlotAssignmentNamed, // Default to named assignment
		// All boolean properties initially false per WHATWG DOM Section 4.8
		delegatesFocus:                false,
		clonable:                      false,
		serializable:                  false,
		availableToElementInternals:   false,
		declarative:                   false,
		keepCustomElementRegistryNull: false,
		manualSlotMap:                 make(map[*Slot][]*Element),
	}

	// Set up the DocumentFragment's self reference correctly
	sr.DocumentFragment.nodeImpl.self = sr

	return sr
}

// NewShadowRootWithInit creates a new shadow root with initialization options
func NewShadowRootWithInit(host *Element, init *ShadowRootInit) *ShadowRoot {
	if host == nil {
		panic(NewInvalidStateError("Host element cannot be null"))
	}

	doc := host.OwnerDocument()
	if doc == nil {
		panic(NewInvalidStateError("Host element must have an owner document"))
	}

	mode := ShadowRootModeOpen
	slotAssignment := SlotAssignmentNamed
	delegatesFocus := false
	clonable := false
	serializable := false

	if init != nil {
		mode = init.Mode
		slotAssignment = init.SlotAssignment
		delegatesFocus = init.DelegatesFocus
		clonable = init.Clonable
		serializable = init.Serializable
	}

	sr := &ShadowRoot{
		DocumentFragment:              NewDocumentFragment(doc),
		host:                          host,
		mode:                          mode,
		slotAssignment:                slotAssignment,
		delegatesFocus:                delegatesFocus,
		clonable:                      clonable,
		serializable:                  serializable,
		availableToElementInternals:   false,
		declarative:                   false,
		keepCustomElementRegistryNull: false,
		manualSlotMap:                 make(map[*Slot][]*Element),
	}

	// Set up the DocumentFragment's self reference correctly
	sr.DocumentFragment.nodeImpl.self = sr

	return sr
}

// Host returns the shadow root's host element
func (sr *ShadowRoot) Host() *Element {
	return sr.host
}

// Mode returns the shadow root's mode
func (sr *ShadowRoot) Mode() ShadowRootMode {
	return sr.mode
}

// SlotAssignment returns the shadow root's slot assignment mode
func (sr *ShadowRoot) SlotAssignment() SlotAssignmentMode {
	return sr.slotAssignment
}

// SetSlotAssignment sets the shadow root's slot assignment mode
func (sr *ShadowRoot) SetSlotAssignment(mode SlotAssignmentMode) {
	sr.slotAssignment = mode
}

// DelegatesFocus returns the shadow root's delegates focus
func (sr *ShadowRoot) DelegatesFocus() bool {
	return sr.delegatesFocus
}

// Clonable returns the shadow root's clonable
func (sr *ShadowRoot) Clonable() bool {
	return sr.clonable
}

// Serializable returns the shadow root's serializable
func (sr *ShadowRoot) Serializable() bool {
	return sr.serializable
}

// OnSlotChange returns the onslotchange event handler
func (sr *ShadowRoot) OnSlotChange() EventHandler {
	return sr.onslotchange
}

// SetOnSlotChange sets the onslotchange event handler
func (sr *ShadowRoot) SetOnSlotChange(handler EventHandler) {
	sr.onslotchange = handler
}

// GetManualSlotAssignments returns a copy of the manual slot assignments
func (sr *ShadowRoot) GetManualSlotAssignments() map[*Slot][]*Element {
	result := make(map[*Slot][]*Element)
	for slot, slottables := range sr.manualSlotMap {
		// Make a copy of the slice
		result[slot] = append([]*Element(nil), slottables...)
	}
	return result
}

// SetManualSlotAssignment sets manual slot assignment for a slot
func (sr *ShadowRoot) SetManualSlotAssignment(slot *Slot, slottables []*Element) {
	if slot == nil {
		return
	}

	if slottables == nil {
		delete(sr.manualSlotMap, slot)
	} else {
		// Make a copy of the slice to avoid external modification
		sr.manualSlotMap[slot] = append([]*Element(nil), slottables...)
	}
}

// FindSlot implements the "find a slot" algorithm from WHATWG DOM Section 4.2.2.3
func (sr *ShadowRoot) FindSlot(slottable Slottable, open bool) *Slot {
	if slottable == nil {
		return nil
	}

	parent := slottable.ParentNode()
	if parent == nil {
		return nil
	}

	// Get the parent's shadow root
	var shadow *ShadowRoot
	if elem, ok := parent.(*Element); ok {
		shadow = elem.ShadowRoot()
	}

	if shadow == nil {
		return nil
	}

	// If open is true and shadow's mode is not "open", return null
	if open && shadow.mode != ShadowRootModeOpen {
		return nil
	}

	// If shadow's slot assignment is "manual"
	if shadow.slotAssignment == SlotAssignmentManual {
		// Return the slot in shadow's descendants whose manually assigned nodes contains slottable
		for slot, nodes := range shadow.manualSlotMap {
			for _, node := range nodes {
				if node == slottable {
					return slot
				}
			}
		}
		return nil
	}

	// Return the first slot in tree order in shadow's descendants whose name is slottable's name
	slottableName := slottable.GetSlotName()
	return sr.findSlotByName(shadow, slottableName)
}

// findSlotByName finds the first slot with the given name in tree order
func (sr *ShadowRoot) findSlotByName(shadow *ShadowRoot, name string) *Slot {
	var result *Slot

	// Traverse in tree order to find the first slot with matching name
	Traverse(shadow, func(node Node) bool {
		if slot, ok := node.(*Slot); ok {
			if slot.GetSlotName() == name {
				result = slot
				return false // Stop traversal
			}
		}
		return true // Continue traversal
	})

	return result
}

// FindSlottables implements the "find slottables" algorithm from WHATWG DOM Section 4.2.2.3
func (sr *ShadowRoot) FindSlottables(slot *Slot) []Slottable {
	if slot == nil {
		return nil
	}

	var result []Slottable

	root := slot.GetRootNode(nil)
	shadowRoot, ok := root.(*ShadowRoot)
	if !ok {
		return result
	}

	host := shadowRoot.Host()
	if host == nil {
		return result
	}

	// If root's slot assignment is "manual"
	if shadowRoot.slotAssignment == SlotAssignmentManual {
		// For each slottable of slot's manually assigned nodes
		if manualNodes, exists := shadowRoot.manualSlotMap[slot]; exists {
			for _, slottable := range manualNodes {
				// If slottable's parent is host, append slottable to result
				if slottable.ParentNode() == host {
					result = append(result, slottable)
				}
			}
		}
	} else {
		// For each slottable child of host, in tree order
		children := host.ChildNodes()
		for i := 0; i < children.Length(); i++ {
			child := children.Item(i)
			if slottable, ok := child.(Slottable); ok {
				// Find a slot for this slottable
				foundSlot := sr.FindSlot(slottable, false)
				if foundSlot == slot {
					result = append(result, slottable)
				}
			}
		}
	}

	return result
}

// FindFlattenedSlottables implements the "find flattened slottables" algorithm from WHATWG DOM Section 4.2.2.3
func (sr *ShadowRoot) FindFlattenedSlottables(slot *Slot) []Slottable {
	if slot == nil {
		return nil
	}

	var result []Slottable

	root := slot.GetRootNode(nil)
	if _, ok := root.(*ShadowRoot); !ok {
		return result
	}

	slottables := sr.FindSlottables(slot)

	// If slottables is empty, append each slottable child of slot, in tree order
	if len(slottables) == 0 {
		children := slot.ChildNodes()
		for i := 0; i < children.Length(); i++ {
			child := children.Item(i)
			if slottable, ok := child.(Slottable); ok {
				slottables = append(slottables, slottable)
			}
		}
	}

	// For each node of slottables
	for _, node := range slottables {
		if slot, ok := node.(*Slot); ok {
			// If node is a slot whose root is a shadow root
			if slotRoot := slot.GetRootNode(nil); slotRoot != nil {
				if _, ok := slotRoot.(*ShadowRoot); ok {
					// Recursively get flattened slottables
					temporaryResult := sr.FindFlattenedSlottables(slot)
					result = append(result, temporaryResult...)
				}
			}
		} else {
			// Otherwise, append node to result
			result = append(result, node)
		}
	}

	return result
}

// assignSlottablesForSlot implements the "assign slottables for a slot" algorithm from WHATWG DOM Section 4.2.2.4
func (sr *ShadowRoot) assignSlottablesForSlot(slot *Slot) {
	if slot == nil {
		return
	}

	slottables := sr.FindSlottables(slot)

	// Check if slottables and slot's assigned nodes are not identical
	if !sr.slottablesIdentical(slottables, slot.GetAssignedNodes()) {
		sr.signalSlotChange(slot)
	}

	// Set slot's assigned nodes to slottables
	slot.SetAssignedNodes(slottables)

	// For each slottable of slottables, set slottable's assigned slot to slot
	for _, slottable := range slottables {
		slottable.SetAssignedSlot(slot)
	}
}

// slottablesIdentical checks if two slottable slices are identical
func (sr *ShadowRoot) slottablesIdentical(a, b []Slottable) bool {
	if len(a) != len(b) {
		return false
	}
	for i, slottable := range a {
		if slottable != b[i] {
			return false
		}
	}
	return true
}

// AssignSlottablesForTree implements the "assign slottables for a tree" algorithm from WHATWG DOM Section 4.2.2.4
func (sr *ShadowRoot) AssignSlottablesForTree(root Node) {
	if root == nil {
		return
	}

	// Run assign slottables for each slot of root's inclusive descendants, in tree order
	Traverse(root, func(node Node) bool {
		if slot, ok := node.(*Slot); ok {
			sr.assignSlottablesForSlot(slot)
		}
		return true // Continue traversal
	})
}

// AssignSlot implements the "assign a slot" algorithm from WHATWG DOM Section 4.2.2.4
func (sr *ShadowRoot) AssignSlot(slottable Slottable) {
	if slottable == nil {
		return
	}

	slot := sr.FindSlot(slottable, false)
	if slot != nil {
		sr.assignSlottablesForSlot(slot)
	}
}

// signalSlotChange implements the "signal a slot change" algorithm from WHATWG DOM Section 4.2.2.5
func (sr *ShadowRoot) signalSlotChange(slot *Slot) {
	if slot == nil {
		return
	}

	// Get the document to access its signal slots
	doc := slot.OwnerDocument()
	if doc != nil {
		// Append slot to slot's relevant agent's signal slots
		doc.addSignalSlot(slot)

		// Queue a mutation observer microtask
		doc.queueMutationObserverMicrotask()
	}
}

// GetEventParent implements the shadow root's "get the parent" algorithm per WHATWG DOM Section 4.8
// Returns null if event's composed flag is unset and shadow root is the root of event's path's first struct's invocation target;
// otherwise shadow root's host.
func (sr *ShadowRoot) GetEventParent(event Event) Node {
	if event == nil {
		return sr.host
	}

	// TODO: Implement full composed flag support in Event interface
	// For now, always return the host as a simplified implementation
	// Full implementation would check:
	// - If event's composed flag is unset and this shadow root is the root of event's path's first struct's invocation target, return null
	// - Otherwise return the host
	return sr.host
}

// Override GetRootNode to handle shadow-including root
func (sr *ShadowRoot) GetRootNode(options *GetRootNodeOptions) Node {
	if options != nil && options.Composed {
		// For shadow-including root, traverse up through shadow boundaries
		current := Node(sr)
		for {
			parent := current.ParentNode()
			if parent != nil {
				current = parent
				continue
			}

			// Check if current is a shadow root
			if shadowRoot, ok := current.(*ShadowRoot); ok {
				if shadowRoot.host != nil {
					current = shadowRoot.host
					continue
				}
			}

			break
		}
		return current
	}

	// Normal root (don't cross shadow boundaries)
	return sr.DocumentFragment.GetRootNode(options)
}

// Override NodeName for ShadowRoot
func (sr *ShadowRoot) NodeName() string {
	return "#shadow-root"
}

// CustomElementRegistry implements the DocumentOrShadowRoot mixin getter per WHATWG DOM Section 4.2.5
// Returns this shadow root's CustomElementRegistry object, if any; otherwise null.
func (sr *ShadowRoot) CustomElementRegistry() *CustomElementRegistry {
	return sr.customElementRegistry
}

// ShadowIncludingRoot returns the shadow-including root of an object per WHATWG DOM Section 4.8
// The shadow-including root of an object is its root's host's shadow-including root,
// if the object's root is a shadow root; otherwise its root.
func ShadowIncludingRoot(node Node) Node {
	if node == nil {
		return nil
	}

	root := node.GetRootNode(nil)

	// If root is a shadow root
	if shadowRoot, ok := root.(*ShadowRoot); ok {
		// Return the shadow-including root of the host
		if shadowRoot.host != nil {
			return ShadowIncludingRoot(shadowRoot.host)
		}
	}

	// Otherwise return the root
	return root
}

// IsShadowIncludingDescendant returns true if A is a shadow-including descendant of B per WHATWG DOM Section 4.8
// An object A is a shadow-including descendant of an object B, if A is a descendant of B,
// or A's root is a shadow root and A's root's host is a shadow-including inclusive descendant of B.
func IsShadowIncludingDescendant(a, b Node) bool {
	if a == nil || b == nil {
		return false
	}

	// If A is a descendant of B (B contains A)
	if b.Contains(a) {
		return true
	}

	// If A's root is a shadow root and A's root's host is a shadow-including inclusive descendant of B
	aRoot := a.GetRootNode(nil)
	if shadowRoot, ok := aRoot.(*ShadowRoot); ok {
		if shadowRoot.host != nil {
			return IsShadowIncludingInclusiveDescendant(shadowRoot.host, b)
		}
	}

	return false
}

// IsShadowIncludingInclusiveDescendant returns true if A is a shadow-including inclusive descendant of B
// A shadow-including inclusive descendant is an object or one of its shadow-including descendants.
func IsShadowIncludingInclusiveDescendant(a, b Node) bool {
	if a == nil || b == nil {
		return false
	}

	// If A is B
	if a == b {
		return true
	}

	// If A is a shadow-including descendant of B
	return IsShadowIncludingDescendant(a, b)
}

// IsShadowIncludingAncestor returns true if A is a shadow-including ancestor of B
// An object A is a shadow-including ancestor of an object B, if and only if B is a shadow-including descendant of A.
func IsShadowIncludingAncestor(a, b Node) bool {
	return IsShadowIncludingDescendant(b, a)
}

// IsShadowIncludingInclusiveAncestor returns true if A is a shadow-including inclusive ancestor of B
// A shadow-including inclusive ancestor is an object or one of its shadow-including ancestors.
func IsShadowIncludingInclusiveAncestor(a, b Node) bool {
	if a == nil || b == nil {
		return false
	}

	// If A is B
	if a == b {
		return true
	}

	// If A is a shadow-including ancestor of B
	return IsShadowIncludingAncestor(a, b)
}

// IsClosedShadowHidden returns true if A is closed-shadow-hidden from B per WHATWG DOM Section 4.8
// A node A is closed-shadow-hidden from a node B if all of the following conditions are true:
// 1. A's root is a shadow root.
// 2. A's root is not a shadow-including inclusive ancestor of B.
// 3. A's root is a shadow root whose mode is "closed" or A's root's host is closed-shadow-hidden from B.
func IsClosedShadowHidden(a, b Node) bool {
	if a == nil || b == nil {
		return false
	}

	// A's root is a shadow root
	aRoot := a.GetRootNode(nil)
	shadowRoot, ok := aRoot.(*ShadowRoot)
	if !ok {
		return false
	}

	// A's root is not a shadow-including inclusive ancestor of B
	if IsShadowIncludingInclusiveAncestor(aRoot, b) {
		return false
	}

	// A's root is a shadow root whose mode is "closed" or A's root's host is closed-shadow-hidden from B
	if shadowRoot.mode == ShadowRootModeClosed {
		return true
	}

	if shadowRoot.host != nil {
		return IsClosedShadowHidden(shadowRoot.host, b)
	}

	return false
}

// Retarget implements the retargeting algorithm per WHATWG DOM Section 4.8
// To retarget an object A against an object B, repeat these steps until they return an object:
// 1. If one of the following is true:
//   - A is not a node
//   - A's root is not a shadow root
//   - B is a node and A's root is a shadow-including inclusive ancestor of B
//     then return A.
//
// 2. Set A to A's root's host.
func Retarget(a, b Node) Node {
	if a == nil {
		return nil
	}

	current := a
	for {
		// If A is not a node (already checked above, A is a Node)
		// If A's root is not a shadow root
		currentRoot := current.GetRootNode(nil)
		shadowRoot, ok := currentRoot.(*ShadowRoot)
		if !ok {
			return current
		}

		// If B is a node and A's root is a shadow-including inclusive ancestor of B
		if b != nil && IsShadowIncludingInclusiveAncestor(currentRoot, b) {
			return current
		}

		// Set A to A's root's host
		if shadowRoot.host == nil {
			return current
		}
		current = shadowRoot.host
	}
}

// TraverseShadowIncludingTreeOrder traverses nodes in shadow-including tree order per WHATWG DOM Section 4.8
// Shadow-including tree order is shadow-including preorder, depth-first traversal of a node tree.
// Shadow-including preorder, depth-first traversal of a node tree is preorder, depth-first traversal of tree,
// with for each shadow host encountered in tree, shadow-including preorder, depth-first traversal
// of that element's shadow root's node tree just after it is encountered.
func TraverseShadowIncludingTreeOrder(root Node, visitor func(Node) bool) {
	if root == nil || visitor == nil {
		return
	}

	// Visit the current node
	if !visitor(root) {
		return // Stop traversal if visitor returns false
	}

	// If this is a shadow host, traverse its shadow root
	if element, ok := root.(*Element); ok {
		if shadowRoot := element.ShadowRoot(); shadowRoot != nil {
			TraverseShadowIncludingTreeOrder(shadowRoot, visitor)
		}
	}

	// Traverse children in normal tree order
	children := root.ChildNodes()
	if children != nil {
		for i := 0; i < children.Length(); i++ {
			child := children.Item(i)
			TraverseShadowIncludingTreeOrder(child, visitor)
		}
	}
}

// String representation for debugging
func (sr *ShadowRoot) String() string {
	var mode string
	if sr.mode == ShadowRootModeOpen {
		mode = "open"
	} else {
		mode = "closed"
	}

	hostTag := ""
	if sr.host != nil {
		hostTag = sr.host.TagName()
	}

	return "#shadow-root (" + mode + ") host: " + hostTag
}
