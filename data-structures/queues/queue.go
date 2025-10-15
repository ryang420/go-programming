package queues

import (
	"errors"
	"fmt"
	"strings"
)

// Queue represents a FIFO (First In, First Out) data structure
type Queue struct {
	items []int
}

// NewQueue creates a new empty queue
func NewQueue() *Queue {
	return &Queue{
		items: make([]int, 0),
	}
}

// Enqueue adds an element to the rear of the queue
func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

// Dequeue removes and returns the front element from the queue
func (q *Queue) Dequeue() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("queue is empty")
	}

	item := q.items[0]
	q.items = q.items[1:]
	return item, nil
}

// Front returns the front element without removing it
func (q *Queue) Front() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("queue is empty")
	}

	return q.items[0], nil
}

// Rear returns the rear element without removing it
func (q *Queue) Rear() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("queue is empty")
	}

	return q.items[len(q.items)-1], nil
}

// IsEmpty checks if the queue is empty
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

// Size returns the number of elements in the queue
func (q *Queue) Size() int {
	return len(q.items)
}

// Clear removes all elements from the queue
func (q *Queue) Clear() {
	q.items = q.items[:0]
}

// ToSlice returns a copy of the queue elements as a slice (front to rear)
func (q *Queue) ToSlice() []int {
	result := make([]int, len(q.items))
	copy(result, q.items)
	return result
}

// String returns a string representation of the queue
func (q *Queue) String() string {
	if q.IsEmpty() {
		return "Queue: []"
	}

	var result strings.Builder
	result.WriteString("Queue: front -> [")

	for i, item := range q.items {
		result.WriteString(fmt.Sprintf("%d", item))
		if i < len(q.items)-1 {
			result.WriteString(", ")
		}
	}

	result.WriteString("] <- rear")
	return result.String()
}

// Contains checks if the queue contains the given value
func (q *Queue) Contains(item int) bool {
	for _, value := range q.items {
		if value == item {
			return true
		}
	}
	return false
}

