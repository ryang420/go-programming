package queues

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestArrayQueue(t *testing.T) {
	queue := NewQueue()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	assert.Equal(t, 3, queue.Size(), "Queue size should be 3")
	assert.False(t, queue.IsEmpty(), "Queue should not be empty")

	front, _ := queue.Front()
	assert.Equal(t, 1, front, "Front element should be 1")
	rear, _ := queue.Rear()
	assert.Equal(t, 3, rear, "Rear element should be 3")

	queue.Dequeue()
	assert.Equal(t, 2, queue.Size(), "Queue size should be 2 after dequeue")
	
}