package stacks

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestArrayStack(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	assert.Equal(t, 3, stack.Size(), "Stack size should be 3")
	
	val, _ := stack.Peek()
	assert.Equal(t, 3, val, "Top element should be 3")
	val, _ = stack.Pop()
	assert.Equal(t, 3, val, "Popped element should be 3")
	assert.Equal(t, 2, stack.Size(), "Stack size should be 2 after pop")
	
	assert.Equal(t, 2, stack.Size(), "Stack size should be 2 after pop")
	assert.True(t, stack.Contains(1), "Stack should contain 1")
	
	stackString := stack.String()
	assert.Equal(t, "Stack: [1, 2] <- top", stackString, "Stack string should be [1, 2] <- top")

	stack.Clear()
	assert.Equal(t, 0, stack.Size(), "Stack size should be 0 after clear")
	assert.True(t, stack.IsEmpty(), "Stack should be empty after clear")
	
}