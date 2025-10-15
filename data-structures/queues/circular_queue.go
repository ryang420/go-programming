package queues

import "errors"

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
