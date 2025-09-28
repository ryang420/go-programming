package easy

import (
	"reflect"
	"testing"
)

// 辅助函数: 将 []int 转为链表
func sliceToList(nums []int) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for _, v := range nums {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return dummy.Next
}

// 辅助函数: 将链表转为 []int
func listToSlice(head *ListNode) []int {
	if head == nil {
		return []int{} // 保证不是 nil
	}
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		l1       []int
		l2       []int
		expected []int
	}{
		{[]int{1, 2, 4}, []int{1, 3, 4}, []int{1, 1, 2, 3, 4, 4}},
		{[]int{}, []int{}, []int{}},
		{[]int{}, []int{0}, []int{0}},
		{[]int{2, 5, 7}, []int{1, 3, 6}, []int{1, 2, 3, 5, 6, 7}},
		{[]int{1, 1, 2}, []int{1, 1, 2}, []int{1, 1, 1, 1, 2, 2}},
		{[]int{1, 2, 3}, []int{}, []int{1, 2, 3}},
		{[]int{}, []int{4, 5, 6}, []int{4, 5, 6}},
	}

	for i, tt := range tests {
		l1 := sliceToList(tt.l1)
		l2 := sliceToList(tt.l2)
		got := mergeTwoLists(l1, l2)
		gotSlice := listToSlice(got)

		if !reflect.DeepEqual(gotSlice, tt.expected) {
			t.Errorf("test %d failed: expected %v, got %v", i+1, tt.expected, gotSlice)
		}
	}
}
