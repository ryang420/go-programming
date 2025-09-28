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

// CircularQueue represents a fixed-size circular queue
type CircularQueue struct {
	items    []int
	front    int
	rear     int
	size     int
	capacity int
}

// NewCircularQueue creates a new circular queue with the specified capacity
func NewCircularQueue(capacity int) *CircularQueue {
	return &CircularQueue{
		items:    make([]int, capacity),
		front:    0,
		rear:     -1,
		size:     0,
		capacity: capacity,
	}
}

// Enqueue adds an element to the circular queue
func (cq *CircularQueue) Enqueue(item int) error {
	if cq.IsFull() {
		return errors.New("queue is full")
	}

	cq.rear = (cq.rear + 1) % cq.capacity
	cq.items[cq.rear] = item
	cq.size++
	return nil
}

// Dequeue removes and returns the front element from the circular queue
func (cq *CircularQueue) Dequeue() (int, error) {
	if cq.IsEmpty() {
		return 0, errors.New("queue is empty")
	}

	item := cq.items[cq.front]
	cq.front = (cq.front + 1) % cq.capacity
	cq.size--
	return item, nil
}

// IsEmpty checks if the circular queue is empty
func (cq *CircularQueue) IsEmpty() bool {
	return cq.size == 0
}

// IsFull checks if the circular queue is full
func (cq *CircularQueue) IsFull() bool {
	return cq.size == cq.capacity
}

// Size returns the current size of the circular queue
func (cq *CircularQueue) Size() int {
	return cq.size
}

// Capacity returns the maximum capacity of the circular queue
func (cq *CircularQueue) Capacity() int {
	return cq.capacity
}
