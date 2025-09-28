package arrays

import (
	"errors"
	"fmt"
)

// DynamicArray represents a resizable array
type DynamicArray[T any] struct {
	data     []T
	size     int
	capacity int
}

// NewDynamicArray creates a new dynamic array with initial capacity
func NewDynamicArray[T any](initialCapacity int) *DynamicArray[T] {
	if initialCapacity <= 0 {
		initialCapacity = 10
	}
	return &DynamicArray[T]{
		data:     make([]T, 0, initialCapacity),
		size:     0,
		capacity: initialCapacity,
	}
}

// Append adds an element to the end of the array
func (da *DynamicArray[T]) Append(value T) {
	if da.size >= da.capacity {
		da.increaseCapacity()
	}
	da.data = append(da.data, value)
	da.size++
}

// Get returns the element at the given index
func (da *DynamicArray[T]) Get(index int) (T, error) {
	if index < 0 || index >= da.size {
		var zero T
		return zero, errors.New("index out of bounds")
	}
	return da.data[index], nil
}

// Set updates the element at the given index
func (da *DynamicArray[T]) Set(index int, value T) error {
	if index < 0 || index >= da.size {
		return errors.New("index out of bounds")
	}
	// Set operation doesn't change size, no need to resize
	da.data[index] = value
	return nil
}

// Insert inserts an element at the given index
func (da *DynamicArray[T]) Insert(index int, value T) error {
	if index < 0 || index > da.size {
		return errors.New("index out of bounds")
	}

	if da.size >= da.capacity {
		da.increaseCapacity()
	}

	da.data = append(da.data[:index], value)
	copy(da.data[index+1:], da.data[index:])
	da.data[index] = value
	da.size++
	return nil
}

// Delete removes the element at the given index
func (da *DynamicArray[T]) Delete(index int) error {
	if index < 0 || index >= da.size {
		return errors.New("index out of bounds")
	}

	copy(da.data[index:], da.data[index+1:]) // 左移覆盖index位置
	da.data = da.data[:da.size-1]
	da.size--

	// Shrink if usage is below 25% and capacity > minimum
	if da.size > 0 && da.size < da.capacity/4 && da.capacity > 16 {
		da.decreaseCapacity()
	}
	return nil
}

// Size returns the current size of the array
func (da *DynamicArray[T]) Size() int {
	return da.size
}

// IsEmpty checks if the array is empty
func (da *DynamicArray[T]) IsEmpty() bool {
	return da.size == 0
}

// ToSlice returns a copy of the current elements as a slice
func (da *DynamicArray[T]) ToSlice() []T {
	result := make([]T, da.size)
	copy(result, da.data[:da.size])
	return result
}

// resize doubles the capacity of the array
func (da *DynamicArray[T]) increaseCapacity() {
	da.capacity *= 2
	newData := make([]T, da.capacity)
	copy(newData, da.data)
	da.data = newData
}

func (da *DynamicArray[T]) decreaseCapacity() {
	da.capacity /= 2
	// Ensure minimum capacity
	if da.capacity < 8 {
		da.capacity = 8
	}
	newData := make([]T, da.capacity)
	copy(newData, da.data[:da.size])
	da.data = newData
}

// String returns a string representation of the array
func (da *DynamicArray[T]) String() string {
	return fmt.Sprintf("DynamicArray{size: %d, capacity: %d, data: %v}",
		da.size, da.capacity, da.data[:da.size])
}
