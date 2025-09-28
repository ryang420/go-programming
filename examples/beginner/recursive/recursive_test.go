package recursive

import (
	"testing"
)

func TestFactorial(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{0, 1},
		{1, 1},
		{3, 6},
		{5, 120},
		{10, 3_628_800},
	}

	for _, tt := range tests {
		res := factorial(tt.n)
		if res != tt.expected {
			t.Errorf("factorial(%d) = %d; want %d", tt.n, res, tt.expected)
		}
	}
}

func TestMergeSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{38, 27, 43, 3, 9, 82, 10}, []int{3, 9, 10, 27, 38, 43, 82}},
		{[]int{5, 2, 9, 1, 5, 6}, []int{1, 2, 5, 5, 6, 9}},
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{2, 1}, []int{1, 2}},
	}

	for _, tt := range tests {
		res := mergeSort(tt.input)
		if !equal(res, tt.expected) {
			t.Errorf("mergeSort(%v) = %v; want %v", tt.input, res, tt.expected)
		}
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
