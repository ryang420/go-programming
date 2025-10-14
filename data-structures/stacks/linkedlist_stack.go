package stacks

import "container/list"

/* 基于链表实现的栈 */
type LinkedListStack struct {
    // 使用内置包 list 来实现栈
    data *list.List
}

/* 初始化栈 */
func NewLinkedListStack() *LinkedListStack {
    return &LinkedListStack{
        data: list.New(),
    }
}

/* 入栈 */
func (s *LinkedListStack) Push(value int) {
    s.data.PushBack(value)
}

/* 出栈 */
func (s *LinkedListStack) Pop() any {
    if s.IsEmpty() {
        return nil
    }
    e := s.data.Back()
    s.data.Remove(e)
    return e.Value
}

/* 访问栈顶元素 */
func (s *LinkedListStack) Peek() any {
    if s.IsEmpty() {
        return nil
    }
    e := s.data.Back()
    return e.Value
}

/* 获取栈的长度 */
func (s *LinkedListStack) Size() int {
    return s.data.Len()
}

/* 判断栈是否为空 */
func (s *LinkedListStack) IsEmpty() bool {
    return s.data.Len() == 0
}

/* 获取 List 用于打印 */
func (s *LinkedListStack) ToList() *list.List {
    return s.data
}