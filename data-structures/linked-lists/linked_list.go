package linkedlists

import (
	"fmt"
	"strings"
)

// Node represents a node in the linked list
type Node struct {
	Data int
	Next *Node
}

// LinkedList represents a singly linked list
type LinkedList struct {
	Head *Node
	size int
}

// NewLinkedList creates a new empty linked list
func NewLinkedList() *LinkedList {
	return &LinkedList{
		Head: nil,
		size: 0,
	}
}

// Prepend adds a node at the beginning of the list
func (ll *LinkedList) Prepend(data int) {
	newNode := &Node{Data: data, Next: ll.Head}
	ll.Head = newNode
	ll.size++
}

// Append adds a node at the end of the list
func (ll *LinkedList) Append(data int) {
	newNode := &Node{Data: data, Next: nil}

	if ll.Head == nil {
		ll.Head = newNode
		ll.size++
		return
	}

	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
	ll.size++
}

// InsertAt inserts a node at the specified index
func (ll *LinkedList) InsertAt(index int, data int) error {
	if index < 0 || index > ll.size {
		return fmt.Errorf("index out of bounds")
	}

	if index == 0 {
		ll.Prepend(data)
		return nil
	}

	newNode := &Node{Data: data, Next: nil}
	current := ll.Head
	for i := 0; i < index-1; i++ {
		current = current.Next
	}

	newNode.Next = current.Next
	current.Next = newNode
	ll.size++
	return nil
}

// DeleteAt removes a node at the specified index
func (ll *LinkedList) DeleteAt(index int) error {
	if index < 0 || index >= ll.size {
		return fmt.Errorf("index out of bounds")
	}

	if index == 0 {
		ll.Head = ll.Head.Next
		ll.size--
		return nil
	}

	current := ll.Head
	for i := 0; i < index-1; i++ {
		current = current.Next
	}

	current.Next = current.Next.Next
	ll.size--
	return nil
}

// DeleteValue removes the first occurrence of the specified value
func (ll *LinkedList) DeleteValue(data int) bool {
	if ll.Head == nil {
		return false
	}

	if ll.Head.Data == data {
		ll.Head = ll.Head.Next
		ll.size--
		return true
	}

	current := ll.Head
	for current.Next != nil {
		if current.Next.Data == data {
			current.Next = current.Next.Next
			ll.size--
			return true
		}
		current = current.Next
	}

	return false
}

// Find returns the index of the first occurrence of the value
func (ll *LinkedList) Find(data int) int {
	current := ll.Head
	index := 0

	for current != nil {
		if current.Data == data {
			return index
		}
		current = current.Next
		index++
	}

	return -1
}

// Get returns the value at the specified index
func (ll *LinkedList) Get(index int) (int, error) {
	if index < 0 || index >= ll.size {
		return 0, fmt.Errorf("index out of bounds")
	}

	current := ll.Head
	for i := 0; i < index; i++ {
		current = current.Next
	}

	return current.Data, nil
}

// Size returns the number of elements in the list
func (ll *LinkedList) Size() int {
	return ll.size
}

// IsEmpty checks if the list is empty
func (ll *LinkedList) IsEmpty() bool {
	return ll.size == 0
}

// ToSlice converts the linked list to a slice
func (ll *LinkedList) ToSlice() []int {
	result := make([]int, ll.size)
	current := ll.Head
	index := 0

	for current != nil {
		result[index] = current.Data
		current = current.Next
		index++
	}

	return result
}

// Reverse reverses the linked list in place
func (ll *LinkedList) Reverse() {
	var prev *Node
	current := ll.Head

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	ll.Head = prev
}

// String returns a string representation of the linked list
func (ll *LinkedList) String() string {
	if ll.Head == nil {
		return "[]"
	}

	var result strings.Builder
	result.WriteString("[")

	current := ll.Head
	for current != nil {
		result.WriteString(fmt.Sprintf("%d", current.Data))
		if current.Next != nil {
			result.WriteString(" -> ")
		}
		current = current.Next
	}

	result.WriteString("]")
	return result.String()
}
