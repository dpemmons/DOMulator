package dom

import (
	"sync"
	"testing"
	"time"
)

// TestMutationObserver_SpecCompliance_ObserveMethod tests all spec requirements for the observe method
func TestMutationObserver_SpecCompliance_ObserveMethod(t *testing.T) {
	doc := NewDocument()
	target := doc.CreateElement("div")
	doc.AppendChild(target)

	callback := func(records []*MutationRecord, observer *MutationObserver) {}
	observer := NewMutationObserver(callback)

	t.Run("Step1: Auto-set attributes when attributeOldValue exists", func(t *testing.T) {
		// Should not panic and should auto-set attributes=true
		observer.Observe(target, MutationObserverInit{
			ChildList:         true,
			AttributeOldValue: true,
			// attributes not set, should be auto-set to true
		})
		observer.Disconnect()
	})

	t.Run("Step1: Auto-set attributes when attributeFilter exists", func(t *testing.T) {
		// Should not panic and should auto-set attributes=true
		observer.Observe(target, MutationObserverInit{
			ChildList:       true,
			AttributeFilter: []string{"class"},
			// attributes not set, should be auto-set to true
		})
		observer.Disconnect()
	})

	t.Run("Step2: Auto-set characterData when characterDataOldValue exists", func(t *testing.T) {
		// Should not panic and should auto-set characterData=true
		observer.Observe(target, MutationObserverInit{
			ChildList:             true,
			CharacterDataOldValue: true,
			// characterData not set, should be auto-set to true
		})
		observer.Disconnect()
	})

	t.Run("Step3: Throw if no observation types enabled", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected panic when no observation types are enabled")
			}
		}()
		observer.Observe(target, MutationObserverInit{
			ChildList:     false,
			Attributes:    false,
			CharacterData: false,
		})
	})

	t.Run("Step4: Throw if attributeOldValue=true but attributes=false", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected panic when attributeOldValue=true but attributes=false")
			}
		}()
		observer.Observe(target, MutationObserverInit{
			ChildList:         true,
			Attributes:        false,
			AttributesSet:     true, // Explicitly set to false
			AttributeOldValue: true,
		})
	})

	t.Run("Step5: Throw if attributeFilter present but attributes=false", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected panic when attributeFilter present but attributes=false")
			}
		}()
		observer.Observe(target, MutationObserverInit{
			ChildList:       true,
			Attributes:      false,
			AttributesSet:   true, // Explicitly set to false
			AttributeFilter: []string{"class"},
		})
	})

	t.Run("Step6: Throw if characterDataOldValue=true but characterData=false", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected panic when characterDataOldValue=true but characterData=false")
			}
		}()
		observer.Observe(target, MutationObserverInit{
			ChildList:             true,
			CharacterData:         false,
			CharacterDataSet:      true, // Explicitly set to false
			CharacterDataOldValue: true,
		})
	})
}

// TestMutationObserver_SpecCompliance_MutationRecordStructure tests MutationRecord compliance
func TestMutationObserver_SpecCompliance_MutationRecordStructure(t *testing.T) {
	doc := NewDocument()
	target := doc.CreateElement("div")
	doc.AppendChild(target)

	var receivedRecords []*MutationRecord
	var mu sync.Mutex

	callback := func(records []*MutationRecord, observer *MutationObserver) {
		mu.Lock()
		defer mu.Unlock()
		receivedRecords = append(receivedRecords, records...)
	}

	observer := NewMutationObserver(callback)

	t.Run("childList record structure", func(t *testing.T) {
		observer.Observe(target, MutationObserverInit{
			ChildList: true,
		})

		addedNode := doc.CreateElement("span")
		record := &MutationRecord{
			Type:            "childList",
			Target:          target,
			AddedNodes:      []Node{addedNode},
			RemovedNodes:    []Node{},
			PreviousSibling: nil,
			NextSibling:     nil,
			AttributeName:   "",
			OldValue:        "",
		}

		doc.getObserverRegistry().NotifyMutation(record)
		doc.getObserverRegistry().ProcessMutationObservers()

		mu.Lock()
		defer mu.Unlock()

		if len(receivedRecords) != 1 {
			t.Fatalf("Expected 1 record, got %d", len(receivedRecords))
		}

		r := receivedRecords[0]
		if r.Type != "childList" {
			t.Errorf("Expected type 'childList', got '%s'", r.Type)
		}
		if r.Target != target {
			t.Errorf("Expected target to match")
		}
		if len(r.AddedNodes) != 1 || r.AddedNodes[0] != addedNode {
			t.Errorf("Expected correct addedNodes")
		}
		if len(r.RemovedNodes) != 0 {
			t.Errorf("Expected empty removedNodes")
		}
		if r.AttributeName != "" {
			t.Errorf("Expected empty attributeName for childList")
		}
		if r.OldValue != "" {
			t.Errorf("Expected empty oldValue for childList")
		}

		receivedRecords = receivedRecords[:0] // Clear
		observer.Disconnect()
	})

	t.Run("attributes record structure", func(t *testing.T) {
		observer.Observe(target, MutationObserverInit{
			Attributes:        true,
			AttributeOldValue: true,
		})

		record := &MutationRecord{
			Type:               "attributes",
			Target:             target,
			AddedNodes:         []Node{},
			RemovedNodes:       []Node{},
			PreviousSibling:    nil,
			NextSibling:        nil,
			AttributeName:      "class",
			AttributeNamespace: "",
			OldValue:           "old-value",
		}

		doc.getObserverRegistry().NotifyMutation(record)
		doc.getObserverRegistry().ProcessMutationObservers()

		mu.Lock()
		defer mu.Unlock()

		if len(receivedRecords) != 1 {
			t.Fatalf("Expected 1 record, got %d", len(receivedRecords))
		}

		r := receivedRecords[0]
		if r.Type != "attributes" {
			t.Errorf("Expected type 'attributes', got '%s'", r.Type)
		}
		if r.AttributeName != "class" {
			t.Errorf("Expected attributeName 'class', got '%s'", r.AttributeName)
		}
		if r.OldValue != "old-value" {
			t.Errorf("Expected oldValue 'old-value', got '%s'", r.OldValue)
		}
		if len(r.AddedNodes) != 0 {
			t.Errorf("Expected empty addedNodes for attributes")
		}

		receivedRecords = receivedRecords[:0] // Clear
		observer.Disconnect()
	})

	t.Run("characterData record structure", func(t *testing.T) {
		textNode := doc.CreateTextNode("original")
		target.AppendChild(textNode)

		observer.Observe(textNode, MutationObserverInit{
			CharacterData:         true,
			CharacterDataOldValue: true,
		})

		record := &MutationRecord{
			Type:            "characterData",
			Target:          textNode,
			AddedNodes:      []Node{},
			RemovedNodes:    []Node{},
			PreviousSibling: nil,
			NextSibling:     nil,
			AttributeName:   "",
			OldValue:        "original",
		}

		doc.getObserverRegistry().NotifyMutation(record)
		doc.getObserverRegistry().ProcessMutationObservers()

		mu.Lock()
		defer mu.Unlock()

		if len(receivedRecords) != 1 {
			t.Fatalf("Expected 1 record, got %d", len(receivedRecords))
		}

		r := receivedRecords[0]
		if r.Type != "characterData" {
			t.Errorf("Expected type 'characterData', got '%s'", r.Type)
		}
		if r.Target != textNode {
			t.Errorf("Expected target to be text node")
		}
		if r.OldValue != "original" {
			t.Errorf("Expected oldValue 'original', got '%s'", r.OldValue)
		}
		if r.AttributeName != "" {
			t.Errorf("Expected empty attributeName for characterData")
		}

		receivedRecords = receivedRecords[:0] // Clear
		observer.Disconnect()
	})
}

// TestMutationObserver_SpecCompliance_QueueingAlgorithm tests the spec's mutation record queuing
func TestMutationObserver_SpecCompliance_QueueingAlgorithm(t *testing.T) {
	doc := NewDocument()
	target := doc.CreateElement("div")
	child := doc.CreateElement("span")
	grandchild := doc.CreateElement("p")
	target.AppendChild(child)
	child.AppendChild(grandchild)
	doc.AppendChild(target)

	var receivedRecords []*MutationRecord
	var mu sync.Mutex

	callback := func(records []*MutationRecord, observer *MutationObserver) {
		mu.Lock()
		defer mu.Unlock()
		receivedRecords = append(receivedRecords, records...)
	}

	observer := NewMutationObserver(callback)

	t.Run("Interested observers - subtree observation", func(t *testing.T) {
		observer.Observe(target, MutationObserverInit{
			ChildList: true,
			Subtree:   true, // Should observe mutations on descendants
		})

		// Mutation on grandchild should be observed
		record := &MutationRecord{
			Type:       "childList",
			Target:     grandchild,
			AddedNodes: []Node{doc.CreateElement("b")},
		}

		doc.getObserverRegistry().NotifyMutation(record)
		doc.getObserverRegistry().ProcessMutationObservers()

		mu.Lock()
		defer mu.Unlock()

		if len(receivedRecords) != 1 {
			t.Errorf("Expected 1 record for subtree observation, got %d", len(receivedRecords))
		}

		receivedRecords = receivedRecords[:0] // Clear
		observer.Disconnect()
	})

	t.Run("Interested observers - no subtree observation", func(t *testing.T) {
		observer.Observe(target, MutationObserverInit{
			ChildList: true,
			Subtree:   false, // Should NOT observe mutations on descendants
		})

		// Mutation on grandchild should NOT be observed
		record := &MutationRecord{
			Type:       "childList",
			Target:     grandchild,
			AddedNodes: []Node{doc.CreateElement("b")},
		}

		doc.getObserverRegistry().NotifyMutation(record)
		doc.getObserverRegistry().ProcessMutationObservers()

		mu.Lock()
		defer mu.Unlock()

		if len(receivedRecords) != 0 {
			t.Errorf("Expected 0 records for non-subtree observation, got %d", len(receivedRecords))
		}

		observer.Disconnect()
	})

	t.Run("Attribute filter enforcement", func(t *testing.T) {
		observer.Observe(target, MutationObserverInit{
			Attributes:      true,
			AttributeFilter: []string{"class", "id"}, // Only observe class and id
		})

		// class attribute - should be observed
		record1 := &MutationRecord{
			Type:          "attributes",
			Target:        target,
			AttributeName: "class",
		}
		doc.getObserverRegistry().NotifyMutation(record1)

		// style attribute - should NOT be observed
		record2 := &MutationRecord{
			Type:          "attributes",
			Target:        target,
			AttributeName: "style",
		}
		doc.getObserverRegistry().NotifyMutation(record2)

		doc.getObserverRegistry().ProcessMutationObservers()

		mu.Lock()
		defer mu.Unlock()

		if len(receivedRecords) != 1 {
			t.Errorf("Expected 1 record for filtered attributes, got %d", len(receivedRecords))
		}

		if len(receivedRecords) > 0 && receivedRecords[0].AttributeName != "class" {
			t.Errorf("Expected filtered attribute 'class', got '%s'", receivedRecords[0].AttributeName)
		}

		receivedRecords = receivedRecords[:0] // Clear
		observer.Disconnect()
	})
}

// TestMutationObserver_SpecCompliance_TakeRecords tests takeRecords method spec compliance
func TestMutationObserver_SpecCompliance_TakeRecords(t *testing.T) {
	doc := NewDocument()
	target := doc.CreateElement("div")
	doc.AppendChild(target)

	var callbackRecords []*MutationRecord
	var mu sync.Mutex

	callback := func(records []*MutationRecord, observer *MutationObserver) {
		mu.Lock()
		defer mu.Unlock()
		callbackRecords = append(callbackRecords, records...)
	}

	observer := NewMutationObserver(callback)
	observer.Observe(target, MutationObserverInit{
		ChildList: true,
	})

	// Queue some records without processing
	record1 := &MutationRecord{
		Type:       "childList",
		Target:     target,
		AddedNodes: []Node{doc.CreateElement("span")},
	}
	record2 := &MutationRecord{
		Type:       "childList",
		Target:     target,
		AddedNodes: []Node{doc.CreateElement("div")},
	}

	observer.queueMutationRecord(record1)
	observer.queueMutationRecord(record2)

	// takeRecords should return and clear the queue
	takenRecords := observer.TakeRecords()

	if len(takenRecords) != 2 {
		t.Errorf("Expected 2 records from takeRecords, got %d", len(takenRecords))
	}

	// Queue should now be empty
	secondTake := observer.TakeRecords()
	if len(secondTake) != 0 {
		t.Errorf("Expected 0 records from second takeRecords, got %d", len(secondTake))
	}

	// Process observers should not call callback since queue is empty
	doc.getObserverRegistry().ProcessMutationObservers()

	mu.Lock()
	defer mu.Unlock()

	if len(callbackRecords) != 0 {
		t.Errorf("Expected callback not to be called after takeRecords, but got %d records", len(callbackRecords))
	}
}

// TestMutationObserver_SpecCompliance_Disconnect tests disconnect method spec compliance
func TestMutationObserver_SpecCompliance_Disconnect(t *testing.T) {
	doc := NewDocument()
	root := doc.CreateElement("div")
	target1 := doc.CreateElement("div")
	target2 := doc.CreateElement("span")
	doc.AppendChild(root)
	root.AppendChild(target1)
	root.AppendChild(target2)

	var callbackRecords []*MutationRecord
	var mu sync.Mutex

	callback := func(records []*MutationRecord, observer *MutationObserver) {
		mu.Lock()
		defer mu.Unlock()
		callbackRecords = append(callbackRecords, records...)
	}

	observer := NewMutationObserver(callback)

	// Observe multiple targets
	observer.Observe(target1, MutationObserverInit{
		ChildList: true,
	})
	observer.Observe(target2, MutationObserverInit{
		Attributes: true,
	})

	// Queue some records
	record1 := &MutationRecord{
		Type:       "childList",
		Target:     target1,
		AddedNodes: []Node{doc.CreateElement("p")},
	}
	observer.queueMutationRecord(record1)

	// Disconnect should remove observer from all nodes and clear record queue
	observer.Disconnect()

	// Try to notify mutations - should not be received
	record2 := &MutationRecord{
		Type:       "childList",
		Target:     target1,
		AddedNodes: []Node{doc.CreateElement("b")},
	}
	record3 := &MutationRecord{
		Type:          "attributes",
		Target:        target2,
		AttributeName: "class",
	}

	doc.getObserverRegistry().NotifyMutation(record2)
	doc.getObserverRegistry().NotifyMutation(record3)
	doc.getObserverRegistry().ProcessMutationObservers()

	// takeRecords should return empty after disconnect
	takenRecords := observer.TakeRecords()
	if len(takenRecords) != 0 {
		t.Errorf("Expected 0 records after disconnect, got %d", len(takenRecords))
	}

	mu.Lock()
	defer mu.Unlock()

	if len(callbackRecords) != 0 {
		t.Errorf("Expected no callback records after disconnect, got %d", len(callbackRecords))
	}
}

// TestMutationObserver_SpecCompliance_MicrotaskQueuing tests microtask queuing behavior
func TestMutationObserver_SpecCompliance_MicrotaskQueuing(t *testing.T) {
	doc := NewDocument()
	target := doc.CreateElement("div")
	doc.AppendChild(target)

	var callbackCount int
	var mu sync.Mutex

	callback := func(records []*MutationRecord, observer *MutationObserver) {
		mu.Lock()
		defer mu.Unlock()
		callbackCount++
	}

	observer := NewMutationObserver(callback)
	observer.Observe(target, MutationObserverInit{
		ChildList: true,
	})

	// Queue multiple mutations before processing
	record1 := &MutationRecord{
		Type:       "childList",
		Target:     target,
		AddedNodes: []Node{doc.CreateElement("span")},
	}
	record2 := &MutationRecord{
		Type:       "childList",
		Target:     target,
		AddedNodes: []Node{doc.CreateElement("div")},
	}

	doc.getObserverRegistry().NotifyMutation(record1)
	doc.getObserverRegistry().NotifyMutation(record2)

	// Process should batch mutations and call callback once
	doc.getObserverRegistry().ProcessMutationObservers()

	mu.Lock()
	defer mu.Unlock()

	if callbackCount != 1 {
		t.Errorf("Expected callback to be called once for batched mutations, got %d calls", callbackCount)
	}
}

// TestMutationObserver_SpecCompliance_CallbackInvocation tests proper callback invocation
func TestMutationObserver_SpecCompliance_CallbackInvocation(t *testing.T) {
	doc := NewDocument()
	target := doc.CreateElement("div")
	doc.AppendChild(target)

	var receivedObserver *MutationObserver
	var receivedRecords []*MutationRecord
	var mu sync.Mutex

	callback := func(records []*MutationRecord, observer *MutationObserver) {
		mu.Lock()
		defer mu.Unlock()
		receivedRecords = records
		receivedObserver = observer
	}

	observer := NewMutationObserver(callback)
	observer.Observe(target, MutationObserverInit{
		ChildList: true,
	})

	record := &MutationRecord{
		Type:       "childList",
		Target:     target,
		AddedNodes: []Node{doc.CreateElement("span")},
	}

	doc.getObserverRegistry().NotifyMutation(record)
	doc.getObserverRegistry().ProcessMutationObservers()

	mu.Lock()
	defer mu.Unlock()

	// Callback should be invoked with records and observer
	if receivedObserver != observer {
		t.Errorf("Expected callback to receive the observer instance")
	}

	if len(receivedRecords) != 1 {
		t.Errorf("Expected callback to receive 1 record, got %d", len(receivedRecords))
	}

	if len(receivedRecords) > 0 && receivedRecords[0] != record {
		t.Errorf("Expected callback to receive the correct record")
	}
}

// TestMutationObserver_SpecCompliance_ConcurrentSafety tests thread safety
func TestMutationObserver_SpecCompliance_ConcurrentSafety(t *testing.T) {
	doc := NewDocument()
	target := doc.CreateElement("div")
	doc.AppendChild(target)

	var receivedCount int
	var mu sync.Mutex

	callback := func(records []*MutationRecord, observer *MutationObserver) {
		mu.Lock()
		defer mu.Unlock()
		receivedCount += len(records)
	}

	observer := NewMutationObserver(callback)
	observer.Observe(target, MutationObserverInit{
		ChildList: true,
	})

	const numGoroutines = 100
	var wg sync.WaitGroup

	// Concurrently queue mutations
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func(index int) {
			defer wg.Done()
			record := &MutationRecord{
				Type:       "childList",
				Target:     target,
				AddedNodes: []Node{doc.CreateElement("span")},
			}
			doc.getObserverRegistry().NotifyMutation(record)
		}(i)
	}

	wg.Wait()

	// Process all mutations
	doc.getObserverRegistry().ProcessMutationObservers()

	// Allow some time for goroutines to complete
	time.Sleep(100 * time.Millisecond)

	mu.Lock()
	defer mu.Unlock()

	if receivedCount != numGoroutines {
		t.Errorf("Expected %d mutation records from concurrent access, got %d", numGoroutines, receivedCount)
	}
}
