package linkedlists

import (
	"testing"
)

func TestLinkedList(t *testing.T) {
	// Test creation
	ll := NewLinkedList()
	if !ll.IsEmpty() {
		t.Error("Expected list to be empty")
	}
	if ll.Size() != 0 {
		t.Errorf("Expected size 0, got %d", ll.Size())
	}

	// Test append
	ll.Append(1)
	ll.Append(2)
	ll.Append(3)
	if ll.Size() != 3 {
		t.Errorf("Expected size 3, got %d", ll.Size())
	}

	// Test prepend
	ll.Prepend(0)
	if ll.Size() != 4 {
		t.Errorf("Expected size 4, got %d", ll.Size())
	}
	val, _ := ll.Get(0)
	if val != 0 {
		t.Errorf("Expected 0 at index 0, got %d", val)
	}

	// Test get
	val, err := ll.Get(2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if val != 2 {
		t.Errorf("Expected 2, got %d", val)
	}

	// Test find
	index := ll.Find(2)
	if index != 2 {
		t.Errorf("Expected index 2, got %d", index)
	}

	index = ll.Find(99)
	if index != -1 {
		t.Errorf("Expected -1 for non-existent value, got %d", index)
	}

	// Test insert at
	err = ll.InsertAt(2, 15)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if ll.Size() != 5 {
		t.Errorf("Expected size 5, got %d", ll.Size())
	}
	val, _ = ll.Get(2)
	if val != 15 {
		t.Errorf("Expected 15 at index 2, got %d", val)
	}

	// Test delete at
	err = ll.DeleteAt(2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if ll.Size() != 4 {
		t.Errorf("Expected size 4, got %d", ll.Size())
	}

	// Test delete value
	found := ll.DeleteValue(1)
	if !found {
		t.Error("Expected to find and delete value 1")
	}
	if ll.Size() != 3 {
		t.Errorf("Expected size 3, got %d", ll.Size())
	}

	// Test to slice
	slice := ll.ToSlice()
	expected := []int{0, 2, 3}
	for i, val := range expected {
		if slice[i] != val {
			t.Errorf("Expected %d at index %d, got %d", val, i, slice[i])
		}
	}

	// Test reverse
	ll.Reverse()
	slice = ll.ToSlice()
	expected = []int{3, 2, 0}
	for i, val := range expected {
		if slice[i] != val {
			t.Errorf("After reverse, expected %d at index %d, got %d", val, i, slice[i])
		}
	}

	// Test bounds checking
	_, err = ll.Get(10)
	if err == nil {
		t.Error("Expected bounds error")
	}
}

func BenchmarkLinkedListAppend(b *testing.B) {
	ll := NewLinkedList()
	for i := 0; i < b.N; i++ {
		ll.Append(i)
	}
}

func BenchmarkLinkedListPrepend(b *testing.B) {
	ll := NewLinkedList()
	for i := 0; i < b.N; i++ {
		ll.Prepend(i)
	}
}
