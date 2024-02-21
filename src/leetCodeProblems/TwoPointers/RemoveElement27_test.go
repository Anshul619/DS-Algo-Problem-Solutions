package main

import (
	"testing"
)

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		nums     []int
		val      int
		expected int
	}{
		{[]int{3, 2, 2, 3}, 3, 2},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5},
		{[]int{}, 0, 0},
		{[]int{1}, 1, 0},
	}

	for i, v := range tests {
		if removeElement(v.nums, v.val) != v.expected {
			t.Errorf("failed test - %v", i)
		}
	}
}
