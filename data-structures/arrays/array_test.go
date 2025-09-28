package arrays

import (
	"testing"
)

func TestDynamicArray(t *testing.T) {
	// Test creation
	da := NewDynamicArray[int](5)
	if da.Size() != 0 {
		t.Errorf("Expected size 0, got %d", da.Size())
	}
	if !da.IsEmpty() {
		t.Error("Expected array to be empty")
	}

	// Test append
	da.Append(1)
	da.Append(2)
	da.Append(3)
	if da.Size() != 3 {
		t.Errorf("Expected size 3, got %d", da.Size())
	}

	// Test get
	val, err := da.Get(1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if val != 2 {
		t.Errorf("Expected 2, got %d", val)
	}

	// Test set
	err = da.Set(1, 10)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	val, _ = da.Get(1)
	if val != 10 {
		t.Errorf("Expected 10, got %d", val)
	}

	// Test insert
	err = da.Insert(1, 5)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if da.Size() != 4 {
		t.Errorf("Expected size 4, got %d", da.Size())
	}
	val, _ = da.Get(1)
	if val != 5 {
		t.Errorf("Expected 5, got %d", val)
	}

	// Test delete
	err = da.Delete(1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if da.Size() != 3 {
		t.Errorf("Expected size 3, got %d", da.Size())
	}

	// Test bounds checking
	_, err = da.Get(10)
	if err == nil {
		t.Error("Expected bounds error")
	}

	// Test resize functionality
	for i := range 10 {
		da.Append(i)
	}
	if da.Size() != 13 {
		t.Errorf("Expected size 13, got %d", da.Size())
	}

	// Test shrink functionality
	for range 10 {
		da.Delete(0)
	}
	if da.Size() != 3 {
		t.Errorf("Expected size 3, got %d", da.Size())
	}
}

func BenchmarkDynamicArrayAppend(b *testing.B) {
	da := NewDynamicArray[int](10)
	for i := 0; i < b.N; i++ {
		da.Append(i)
	}
}

func BenchmarkDynamicArrayGet(b *testing.B) {
	da := NewDynamicArray[int](1000)
	for i := 0; i < 1000; i++ {
		da.Append(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		da.Get(i % 1000)
	}
}
