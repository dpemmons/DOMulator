package dom

import (
	"sync"
	"sync/atomic"
)

// MutationRecord represents a single mutation to the DOM tree per WHATWG DOM spec
type MutationRecord struct {
	Type               string // "childList", "attributes", "characterData"
	Target             Node   // The node that was mutated
	AddedNodes         []Node // Nodes that were added
	RemovedNodes       []Node // Nodes that were removed
	PreviousSibling    Node   // Previous sibling of added/removed nodes
	NextSibling        Node   // Next sibling of added/removed nodes
	AttributeName      string // Name of changed attribute (for "attributes" type)
	AttributeNamespace string // Namespace of changed attribute (renamed from AttributeNS for spec compliance)
	OldValue           string // Old value (for attributes/characterData if configured)
}

// MutationObserverInit represents the configuration for a MutationObserver
type MutationObserverInit struct {
	ChildList             bool     // Observe direct children changes
	Attributes            bool     // Observe attribute changes
	CharacterData         bool     // Observe text content changes
	Subtree               bool     // Observe changes to entire subtree
	AttributeOldValue     bool     // Include old attribute values
	CharacterDataOldValue bool     // Include old character data
	AttributeFilter       []string // Only observe these attributes
	// Internal fields to track if options were explicitly set
	AttributesSet    bool // True if Attributes was explicitly set
	CharacterDataSet bool // True if CharacterData was explicitly set
}

// MutationCallback is the function signature for mutation observer callbacks
type MutationCallback func([]*MutationRecord, *MutationObserver)

// RegisteredObserver represents a registered observer per WHATWG DOM spec
type RegisteredObserver struct {
	Observer *MutationObserver
	Options  MutationObserverInit
	// For transient observers only:
	IsTransient bool
	Source      *RegisteredObserver // The source registered observer
}

// MutationObserver observes mutations to DOM nodes
type MutationObserver struct {
	callback    MutationCallback
	records     []*MutationRecord
	nodes       map[Node]MutationObserverInit
	mu          sync.RWMutex
	isScheduled atomic.Bool
}

// NewMutationObserver creates a new MutationObserver with the given callback
func NewMutationObserver(callback MutationCallback) *MutationObserver {
	return &MutationObserver{
		callback: callback,
		records:  make([]*MutationRecord, 0),
		nodes:    make(map[Node]MutationObserverInit),
	}
}

// Observe starts observing mutations on the target node per WHATWG DOM spec
func (mo *MutationObserver) Observe(target Node, options MutationObserverInit) {
	// Step 1: Auto-set attributes if attributeOldValue or attributeFilter exists and attributes was not explicitly set
	if (options.AttributeOldValue || len(options.AttributeFilter) > 0) && !options.AttributesSet {
		options.Attributes = true
	}

	// Step 2: Auto-set characterData if characterDataOldValue exists and characterData was not explicitly set
	if options.CharacterDataOldValue && !options.CharacterDataSet {
		options.CharacterData = true
	}

	// Step 3: Validate that at least one observation type is enabled
	if !options.ChildList && !options.Attributes && !options.CharacterData {
		panic("Must observe at least one of: childList, attributes, or characterData")
	}

	// Step 4: Validate attributeOldValue constraint (after potential auto-setting)
	if options.AttributeOldValue && !options.Attributes {
		panic("attributeOldValue requires attributes to be true")
	}

	// Step 5: Validate attributeFilter constraint (after potential auto-setting)
	if len(options.AttributeFilter) > 0 && !options.Attributes {
		panic("attributeFilter requires attributes to be true")
	}

	// Step 6: Validate characterDataOldValue constraint (after potential auto-setting)
	if options.CharacterDataOldValue && !options.CharacterData {
		panic("characterDataOldValue requires characterData to be true")
	}

	mo.mu.Lock()
	defer mo.mu.Unlock()

	// Steps 7-8: Check if already observing this target
	// For now, we'll use the simplified approach of updating the options
	// TODO: Implement full registered observer list per spec

	// Register observer with the target node
	mo.nodes[target] = options
	target.registerMutationObserver(mo, options)
}

// Disconnect stops observing all nodes
func (mo *MutationObserver) Disconnect() {
	mo.mu.Lock()
	defer mo.mu.Unlock()

	// Unregister from all observed nodes
	for node := range mo.nodes {
		node.unregisterMutationObserver(mo)
	}

	// Clear observer state
	mo.nodes = make(map[Node]MutationObserverInit)
	mo.records = make([]*MutationRecord, 0)
	mo.isScheduled.Store(false)
}

// TakeRecords returns and clears all pending mutation records
func (mo *MutationObserver) TakeRecords() []*MutationRecord {
	mo.mu.Lock()
	defer mo.mu.Unlock()

	records := mo.records
	mo.records = make([]*MutationRecord, 0)
	return records
}

// queueMutationRecord adds a mutation record to the observer's queue
func (mo *MutationObserver) queueMutationRecord(record *MutationRecord) {
	mo.mu.Lock()
	defer mo.mu.Unlock()

	mo.records = append(mo.records, record)

	// Schedule delivery if not already scheduled
	if !mo.isScheduled.Load() {
		mo.isScheduled.Store(true)
		// In a real implementation, this would be scheduled as a microtask
		// For now, we'll deliver immediately in tests or when explicitly triggered
	}
}

// deliverMutationRecords delivers all pending records to the callback
func (mo *MutationObserver) deliverMutationRecords() {
	mo.mu.Lock()
	if len(mo.records) == 0 {
		mo.mu.Unlock()
		return
	}

	records := mo.records
	mo.records = make([]*MutationRecord, 0)
	mo.isScheduled.Store(false)
	mo.mu.Unlock()

	// Call the callback with the mutation records
	mo.callback(records, mo)
}

// ObserverRegistry manages all mutation observers for a document
type ObserverRegistry struct {
	observers map[Node][]*MutationObserver
	mu        sync.RWMutex
}

// NewObserverRegistry creates a new observer registry
func NewObserverRegistry() *ObserverRegistry {
	return &ObserverRegistry{
		observers: make(map[Node][]*MutationObserver),
	}
}

// RegisterObserver registers a mutation observer for a node
func (or *ObserverRegistry) RegisterObserver(node Node, observer *MutationObserver, options MutationObserverInit) {
	or.mu.Lock()
	defer or.mu.Unlock()

	if or.observers[node] == nil {
		or.observers[node] = make([]*MutationObserver, 0)
	}
	or.observers[node] = append(or.observers[node], observer)
}

// UnregisterObserver removes a mutation observer from a node
func (or *ObserverRegistry) UnregisterObserver(node Node, observer *MutationObserver) {
	or.mu.Lock()
	defer or.mu.Unlock()

	observers := or.observers[node]
	for i, obs := range observers {
		if obs == observer {
			// Remove observer from slice
			or.observers[node] = append(observers[:i], observers[i+1:]...)
			break
		}
	}

	// Clean up empty entries
	if len(or.observers[node]) == 0 {
		delete(or.observers, node)
	}
}

// NotifyMutation notifies all relevant observers about a mutation
func (or *ObserverRegistry) NotifyMutation(record *MutationRecord) {
	or.mu.RLock()
	defer or.mu.RUnlock()

	// Find all observers that should be notified
	target := record.Target

	// Check observers on the target node and its ancestors if subtree observation is enabled
	var nodesToCheck []Node

	// Always check the target node
	nodesToCheck = append(nodesToCheck, target)

	// Check ancestors for subtree observers
	current := target.ParentNode()
	for current != nil {
		nodesToCheck = append(nodesToCheck, current)
		current = current.ParentNode()
	}

	// Notify all relevant observers
	for _, node := range nodesToCheck {
		if observers := or.observers[node]; observers != nil {
			for _, observer := range observers {
				if or.shouldNotifyObserver(observer, node, record) {
					observer.queueMutationRecord(record)
				}
			}
		}
	}
}

// shouldNotifyObserver determines if an observer should be notified about a mutation
func (or *ObserverRegistry) shouldNotifyObserver(observer *MutationObserver, observedNode Node, record *MutationRecord) bool {
	observer.mu.RLock()
	options, exists := observer.nodes[observedNode]
	observer.mu.RUnlock()

	if !exists {
		return false
	}

	// Check if this mutation type is being observed
	switch record.Type {
	case "childList":
		if !options.ChildList {
			return false
		}
	case "attributes":
		if !options.Attributes {
			return false
		}
		// Check attribute filter
		if len(options.AttributeFilter) > 0 {
			found := false
			for _, attr := range options.AttributeFilter {
				if attr == record.AttributeName {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}
	case "characterData":
		if !options.CharacterData {
			return false
		}
	}

	// Check subtree requirement
	if observedNode != record.Target && !options.Subtree {
		return false
	}

	return true
}

// removeTransientObservers removes all transient registered observers for this observer
func (mo *MutationObserver) removeTransientObservers() {
	mo.mu.RLock()
	defer mo.mu.RUnlock()

	// For each node in the observer's node list
	for node := range mo.nodes {
		// TODO: Implement transient observer removal
		// This would remove transient registered observers whose observer is mo
		// from the node's registered observer list
		_ = node // Placeholder for now
	}
}

// ProcessMutationObservers processes all scheduled mutation observers
func (or *ObserverRegistry) ProcessMutationObservers() {
	or.mu.RLock()
	var observersToProcess []*MutationObserver

	// Collect all observers that have scheduled records
	for _, observers := range or.observers {
		for _, observer := range observers {
			if observer.isScheduled.Load() {
				observersToProcess = append(observersToProcess, observer)
			}
		}
	}
	or.mu.RUnlock()

	// Process each observer
	for _, observer := range observersToProcess {
		observer.deliverMutationRecords()
	}
}
