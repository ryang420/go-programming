package stacks

import (
	"testing"
)

func TestStack(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	if stack.Size() != 3 {
		t.Errorf("Expected size 3, got %d", stack.Size())
	}
	
	val, err := stack.Peek()
	if err != nil {
		t.Errorf("Unexpected error on Peek: %v", err)
	} else if val != 3 {
		t.Errorf("Expected 3, got %d", val)
	}
	val, err = stack.Pop()
	if err != nil {
		t.Errorf("Unexpected error on Pop: %v", err)
	} else if val != 3 {
		t.Errorf("Expected 3, got %d", val)
	}
	if stack.Size() != 2 {
		t.Errorf("Expected size 2, got %d", stack.Size())
	}
	if !stack.Contains(1) {
		t.Errorf("Expected stack to contain 1")
	}

	stackString := stack.String()
	if stackString != "Stack: [1, 2] <- top" {
		t.Errorf("Expected Stack: [1, 2] <- top, got %s", stackString)
	}

	stack.Clear()
	if stack.Size() != 0 {
		t.Errorf("Expected size 0, got %d", stack.Size())
	}
	if !stack.IsEmpty() {
		t.Errorf("Expected stack to be empty")
	}
}