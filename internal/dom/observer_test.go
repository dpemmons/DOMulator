package dom

import (
	"sync"
	"testing"
)

func TestMutationObserver_BasicFunctionality(t *testing.T) {
	doc := NewDocument()
	root := doc.CreateElement("div")
	doc.AppendChild(root)

	var receivedRecords []*MutationRecord
	var receivedObserver *MutationObserver
	var mu sync.Mutex

	callback := func(records []*MutationRecord, observer *MutationObserver) {
		mu.Lock()
		defer mu.Unlock()
		receivedRecords = append(receivedRecords, records...)
		receivedObserver = observer
	}

	observer := NewMutationObserver(callback)

	// Test basic observe
	observer.Observe(root, MutationObserverInit{
		ChildList: true,
	})

	// Make a mutation
	child := doc.CreateElement("span")
	root.AppendChild(child)

	// Manually trigger mutation notification
	record := &MutationRecord{
		Type:         "childList",
		Target:       root,
		AddedNodes:   []Node{child},
		RemovedNodes: []Node{},
	}
	doc.getObserverRegistry().NotifyMutation(record)

	// Process mutation observers
	doc.getObserverRegistry().ProcessMutationObservers()

	// Check that callback was called
	mu.Lock()
	defer mu.Unlock()

	if len(receivedRecords) != 1 {
		t.Errorf("Expected 1 mutation record, got %d", len(receivedRecords))
	}

	if receivedRecords[0].Type != "childList" {
		t.Errorf("Expected mutation type 'childList', got '%s'", receivedRecords[0].Type)
	}

	if receivedRecords[0].Target != root {
		t.Errorf("Expected target to be root element")
	}

	if len(receivedRecords[0].AddedNodes) != 1 || receivedRecords[0].AddedNodes[0] != child {
		t.Errorf("Expected one added node")
	}

	if receivedObserver != observer {
		t.Errorf("Expected observer to match")
	}
}

func TestMutationObserver_TakeRecords(t *testing.T) {
	doc := NewDocument()
	root := doc.CreateElement("div")
	doc.AppendChild(root)

	var callbackCalled bool
	callback := func(records []*MutationRecord, observer *MutationObserver) {
		callbackCalled = true
	}

	observer := NewMutationObserver(callback)
	observer.Observe(root, MutationObserverInit{
		ChildList: true,
	})

	// Queue a mutation without processing
	record := &MutationRecord{
		Type:       "childList",
		Target:     root,
		AddedNodes: []Node{doc.CreateElement("span")},
	}
	observer.queueMutationRecord(record)

	// Take records should return pending records and clear queue
	records := observer.TakeRecords()

	if len(records) != 1 {
		t.Errorf("Expected 1 record from TakeRecords, got %d", len(records))
	}

	if records[0].Type != "childList" {
		t.Errorf("Expected record type 'childList', got '%s'", records[0].Type)
	}

	// Queue should now be empty
	records2 := observer.TakeRecords()
	if len(records2) != 0 {
		t.Errorf("Expected empty records after TakeRecords, got %d", len(records2))
	}

	// Process should not call callback since queue is empty
	doc.getObserverRegistry().ProcessMutationObservers()

	if callbackCalled {
		t.Errorf("Expected callback not to be called after TakeRecords")
	}
}

func TestMutationObserver_Disconnect(t *testing.T) {
	doc := NewDocument()
	root := doc.CreateElement("div")
	doc.AppendChild(root)

	var callbackCalled bool
	callback := func(records []*MutationRecord, observer *MutationObserver) {
		callbackCalled = true
	}

	observer := NewMutationObserver(callback)
	observer.Observe(root, MutationObserverInit{
		ChildList: true,
	})

	// Disconnect observer
	observer.Disconnect()

	// Make a mutation
	record := &MutationRecord{
		Type:       "childList",
		Target:     root,
		AddedNodes: []Node{doc.CreateElement("span")},
	}
	doc.getObserverRegistry().NotifyMutation(record)
	doc.getObserverRegistry().ProcessMutationObservers()

	// Callback should not have been called
	if callbackCalled {
		t.Errorf("Expected callback not to be called after disconnect")
	}
}

func TestMutationObserver_AttributeChanges(t *testing.T) {
	doc := NewDocument()
	root := doc.CreateElement("div")
	doc.AppendChild(root)

	var receivedRecords []*MutationRecord
	var mu sync.Mutex

	callback := func(records []*MutationRecord, observer *MutationObserver) {
		mu.Lock()
		defer mu.Unlock()
		receivedRecords = append(receivedRecords, records...)
	}

	observer := NewMutationObserver(callback)
	observer.Observe(root, MutationObserverInit{
		Attributes:        true,
		AttributeOldValue: true,
	})

	// Make an attribute mutation
	record := &MutationRecord{
		Type:          "attributes",
		Target:        root,
		AttributeName: "class",
		OldValue:      "old-class",
	}
	doc.getObserverRegistry().NotifyMutation(record)
	doc.getObserverRegistry().ProcessMutationObservers()

	mu.Lock()
	defer mu.Unlock()

	if len(receivedRecords) != 1 {
		t.Errorf("Expected 1 mutation record, got %d", len(receivedRecords))
	}

	if receivedRecords[0].Type != "attributes" {
		t.Errorf("Expected mutation type 'attributes', got '%s'", receivedRecords[0].Type)
	}

	if receivedRecords[0].AttributeName != "class" {
		t.Errorf("Expected attribute name 'class', got '%s'", receivedRecords[0].AttributeName)
	}

	if receivedRecords[0].OldValue != "old-class" {
		t.Errorf("Expected old value 'old-class', got '%s'", receivedRecords[0].OldValue)
	}
}

func TestMutationObserver_AttributeFilter(t *testing.T) {
	doc := NewDocument()
	root := doc.CreateElement("div")
	doc.AppendChild(root)

	var receivedRecords []*MutationRecord
	var mu sync.Mutex

	callback := func(records []*MutationRecord, observer *MutationObserver) {
		mu.Lock()
		defer mu.Unlock()
		receivedRecords = append(receivedRecords, records...)
	}

	observer := NewMutationObserver(callback)
	observer.Observe(root, MutationObserverInit{
		Attributes:      true,
		AttributeFilter: []string{"class", "id"},
	})

	// Test filtered attribute - should be observed
	record1 := &MutationRecord{
		Type:          "attributes",
		Target:        root,
		AttributeName: "class",
	}
	doc.getObserverRegistry().NotifyMutation(record1)

	// Test non-filtered attribute - should be ignored
	record2 := &MutationRecord{
		Type:          "attributes",
		Target:        root,
		AttributeName: "style",
	}
	doc.getObserverRegistry().NotifyMutation(record2)

	doc.getObserverRegistry().ProcessMutationObservers()

	mu.Lock()
	defer mu.Unlock()

	if len(receivedRecords) != 1 {
		t.Errorf("Expected 1 mutation record (filtered), got %d", len(receivedRecords))
	}

	if receivedRecords[0].AttributeName != "class" {
		t.Errorf("Expected filtered attribute 'class', got '%s'", receivedRecords[0].AttributeName)
	}
}

func TestMutationObserver_CharacterData(t *testing.T) {
	doc := NewDocument()
	text := doc.CreateTextNode("original text")
	root := doc.CreateElement("div")
	root.AppendChild(text)
	doc.AppendChild(root)

	var receivedRecords []*MutationRecord
	var mu sync.Mutex

	callback := func(records []*MutationRecord, observer *MutationObserver) {
		mu.Lock()
		defer mu.Unlock()
		receivedRecords = append(receivedRecords, records...)
	}

	observer := NewMutationObserver(callback)
	observer.Observe(text, MutationObserverInit{
		CharacterData:         true,
		CharacterDataOldValue: true,
	})

	// Make a character data mutation
	record := &MutationRecord{
		Type:     "characterData",
		Target:   text,
		OldValue: "original text",
	}
	doc.getObserverRegistry().NotifyMutation(record)
	doc.getObserverRegistry().ProcessMutationObservers()

	mu.Lock()
	defer mu.Unlock()

	if len(receivedRecords) != 1 {
		t.Errorf("Expected 1 mutation record, got %d", len(receivedRecords))
	}

	if receivedRecords[0].Type != "characterData" {
		t.Errorf("Expected mutation type 'characterData', got '%s'", receivedRecords[0].Type)
	}

	if receivedRecords[0].OldValue != "original text" {
		t.Errorf("Expected old value 'original text', got '%s'", receivedRecords[0].OldValue)
	}
}

func TestMutationObserver_SubtreeObservation(t *testing.T) {
	doc := NewDocument()
	root := doc.CreateElement("div")
	child := doc.CreateElement("span")
	grandchild := doc.CreateElement("p")

	root.AppendChild(child)
	child.AppendChild(grandchild)
	doc.AppendChild(root)

	var receivedRecords []*MutationRecord
	var mu sync.Mutex

	callback := func(records []*MutationRecord, observer *MutationObserver) {
		mu.Lock()
		defer mu.Unlock()
		receivedRecords = append(receivedRecords, records...)
	}

	observer := NewMutationObserver(callback)
	observer.Observe(root, MutationObserverInit{
		ChildList: true,
		Subtree:   true,
	})

	// Make a mutation on the grandchild (should be observed due to subtree=true)
	newElement := doc.CreateElement("b")
	record := &MutationRecord{
		Type:       "childList",
		Target:     grandchild,
		AddedNodes: []Node{newElement},
	}
	doc.getObserverRegistry().NotifyMutation(record)
	doc.getObserverRegistry().ProcessMutationObservers()

	mu.Lock()
	defer mu.Unlock()

	if len(receivedRecords) != 1 {
		t.Errorf("Expected 1 mutation record (subtree), got %d", len(receivedRecords))
	}

	if receivedRecords[0].Target != grandchild {
		t.Errorf("Expected target to be grandchild element")
	}
}

func TestMutationObserver_SubtreeObservationWithoutSubtree(t *testing.T) {
	doc := NewDocument()
	root := doc.CreateElement("div")
	child := doc.CreateElement("span")
	grandchild := doc.CreateElement("p")

	root.AppendChild(child)
	child.AppendChild(grandchild)
	doc.AppendChild(root)

	var receivedRecords []*MutationRecord
	var mu sync.Mutex

	callback := func(records []*MutationRecord, observer *MutationObserver) {
		mu.Lock()
		defer mu.Unlock()
		receivedRecords = append(receivedRecords, records...)
	}

	observer := NewMutationObserver(callback)
	observer.Observe(root, MutationObserverInit{
		ChildList: true,
		Subtree:   false, // Only observe direct children
	})

	// Make a mutation on the grandchild (should NOT be observed due to subtree=false)
	newElement := doc.CreateElement("b")
	record := &MutationRecord{
		Type:       "childList",
		Target:     grandchild,
		AddedNodes: []Node{newElement},
	}
	doc.getObserverRegistry().NotifyMutation(record)
	doc.getObserverRegistry().ProcessMutationObservers()

	mu.Lock()
	defer mu.Unlock()

	if len(receivedRecords) != 0 {
		t.Errorf("Expected 0 mutation records (no subtree), got %d", len(receivedRecords))
	}
}

func TestMutationObserver_MultipleObservers(t *testing.T) {
	doc := NewDocument()
	root := doc.CreateElement("div")
	doc.AppendChild(root)

	var receivedRecords1 []*MutationRecord
	var receivedRecords2 []*MutationRecord
	var mu1, mu2 sync.Mutex

	callback1 := func(records []*MutationRecord, observer *MutationObserver) {
		mu1.Lock()
		defer mu1.Unlock()
		receivedRecords1 = append(receivedRecords1, records...)
	}

	callback2 := func(records []*MutationRecord, observer *MutationObserver) {
		mu2.Lock()
		defer mu2.Unlock()
		receivedRecords2 = append(receivedRecords2, records...)
	}

	observer1 := NewMutationObserver(callback1)
	observer2 := NewMutationObserver(callback2)

	observer1.Observe(root, MutationObserverInit{
		ChildList: true,
	})

	observer2.Observe(root, MutationObserverInit{
		Attributes: true,
	})

	// Make a childList mutation (should notify observer1 only)
	record1 := &MutationRecord{
		Type:       "childList",
		Target:     root,
		AddedNodes: []Node{doc.CreateElement("span")},
	}
	doc.getObserverRegistry().NotifyMutation(record1)

	// Make an attribute mutation (should notify observer2 only)
	record2 := &MutationRecord{
		Type:          "attributes",
		Target:        root,
		AttributeName: "class",
	}
	doc.getObserverRegistry().NotifyMutation(record2)

	doc.getObserverRegistry().ProcessMutationObservers()

	mu1.Lock()
	mu2.Lock()
	defer mu1.Unlock()
	defer mu2.Unlock()

	if len(receivedRecords1) != 1 {
		t.Errorf("Expected observer1 to receive 1 record, got %d", len(receivedRecords1))
	}

	if len(receivedRecords2) != 1 {
		t.Errorf("Expected observer2 to receive 1 record, got %d", len(receivedRecords2))
	}

	if receivedRecords1[0].Type != "childList" {
		t.Errorf("Expected observer1 to receive childList mutation")
	}

	if receivedRecords2[0].Type != "attributes" {
		t.Errorf("Expected observer2 to receive attributes mutation")
	}
}

func TestMutationObserver_ValidationErrors(t *testing.T) {
	doc := NewDocument()
	root := doc.CreateElement("div")
	doc.AppendChild(root)

	callback := func(records []*MutationRecord, observer *MutationObserver) {}
	observer := NewMutationObserver(callback)

	// Test invalid options - no observation types specified
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for invalid options (no observation types)")
		}
	}()

	observer.Observe(root, MutationObserverInit{
		ChildList:     false,
		Attributes:    false,
		CharacterData: false,
	})
}

func TestMutationObserver_ValidationErrors_AttributeOldValue(t *testing.T) {
	doc := NewDocument()
	root := doc.CreateElement("div")
	doc.AppendChild(root)

	callback := func(records []*MutationRecord, observer *MutationObserver) {}
	observer := NewMutationObserver(callback)

	// Test invalid options - attributeOldValue when attributes is explicitly false
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for attributeOldValue with explicitly false attributes")
		}
	}()

	observer.Observe(root, MutationObserverInit{
		ChildList:         true,
		Attributes:        false,
		AttributesSet:     true, // Explicitly set to false
		AttributeOldValue: true,
	})
}

func TestMutationObserver_ValidationErrors_CharacterDataOldValue(t *testing.T) {
	doc := NewDocument()
	root := doc.CreateElement("div")
	doc.AppendChild(root)

	callback := func(records []*MutationRecord, observer *MutationObserver) {}
	observer := NewMutationObserver(callback)

	// Test invalid options - characterDataOldValue when characterData is explicitly false
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for characterDataOldValue with explicitly false characterData")
		}
	}()

	observer.Observe(root, MutationObserverInit{
		ChildList:             true,
		CharacterData:         false,
		CharacterDataSet:      true, // Explicitly set to false
		CharacterDataOldValue: true,
	})
}

func TestMutationObserver_ConcurrentAccess(t *testing.T) {
	doc := NewDocument()
	root := doc.CreateElement("div")
	doc.AppendChild(root)

	var receivedRecords []*MutationRecord
	var mu sync.Mutex

	callback := func(records []*MutationRecord, observer *MutationObserver) {
		mu.Lock()
		defer mu.Unlock()
		receivedRecords = append(receivedRecords, records...)
	}

	observer := NewMutationObserver(callback)
	observer.Observe(root, MutationObserverInit{
		ChildList: true,
	})

	// Test concurrent mutations
	var wg sync.WaitGroup
	numGoroutines := 10

	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func(index int) {
			defer wg.Done()

			record := &MutationRecord{
				Type:       "childList",
				Target:     root,
				AddedNodes: []Node{doc.CreateElement("span")},
			}
			doc.getObserverRegistry().NotifyMutation(record)
		}(i)
	}

	wg.Wait()
	doc.getObserverRegistry().ProcessMutationObservers()

	mu.Lock()
	defer mu.Unlock()

	if len(receivedRecords) != numGoroutines {
		t.Errorf("Expected %d mutation records from concurrent access, got %d", numGoroutines, len(receivedRecords))
	}
}
