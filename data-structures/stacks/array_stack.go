package stacks

import (
	"errors"
	"fmt"
	"strings"
)

// Stack represents a LIFO (Last In, First Out) data structure
type Stack struct {
	items []int
}

// NewStack creates a new empty stack
func NewStack() *Stack {
	return &Stack{
		items: make([]int, 0),
	}
}

// Push adds an element to the top of the stack
func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top element from the stack
func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}

	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]
	return item, nil
}

// Peek returns the top element without removing it
func (s *Stack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}

	return s.items[len(s.items)-1], nil
}

// IsEmpty checks if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of elements in the stack
func (s *Stack) Size() int {
	return len(s.items)
}

// Clear removes all elements from the stack
func (s *Stack) Clear() {
	s.items = s.items[:0]
}

// ToSlice returns a copy of the stack elements as a slice (bottom to top)
func (s *Stack) ToSlice() []int {
	result := make([]int, len(s.items))
	copy(result, s.items)
	return result
}

// String returns a string representation of the stack
func (s *Stack) String() string {
	if s.IsEmpty() {
		return "Stack: []"
	}

	var result strings.Builder
	result.WriteString("Stack: [")

	for i, item := range s.items {
		result.WriteString(fmt.Sprintf("%d", item))
		if i < len(s.items)-1 {
			result.WriteString(", ")
		}
	}

	result.WriteString("] <- top")
	return result.String()
}

// Contains checks if the stack contains the given value
func (s *Stack) Contains(item int) bool {
	for _, value := range s.items {
		if value == item {
			return true
		}
	}
	return false
}
