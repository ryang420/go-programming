package stacks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedListStack(t *testing.T) {
	stack := NewLinkedListStack()

	// 测试入栈操作
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	// 使用assert替代if判断
	assert.Equal(t, 3, stack.Size(), "Stack size should be 3")
	assert.False(t, stack.IsEmpty(), "Stack should not be empty")

	// 测试栈顶元素
	top := stack.Peek()
	assert.Equal(t, 3, top, "Top element should be 3")

	// 测试出栈操作
	popped := stack.Pop()
	assert.Equal(t, 3, popped, "Popped element should be 3")
	assert.Equal(t, 2, stack.Size(), "Stack size should be 2 after pop")

	// 继续出栈
	popped = stack.Pop()
	assert.Equal(t, 2, popped, "Popped element should be 2")

	popped = stack.Pop()
	assert.Equal(t, 1, popped, "Popped element should be 1")

	// 测试空栈
	assert.True(t, stack.IsEmpty(), "Stack should be empty")
	assert.Equal(t, 0, stack.Size(), "Empty stack size should be 0")

	// 测试空栈操作
	emptyPop := stack.Pop()
	assert.Nil(t, emptyPop, "Pop from empty stack should return nil")

	emptyPeek := stack.Peek()
	assert.Nil(t, emptyPeek, "Peek from empty stack should return nil")
}
