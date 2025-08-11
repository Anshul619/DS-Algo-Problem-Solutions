package main

import (
	"testing"
)

func TestFindPeakElement(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 2, 3, 1}, 2},
		{[]int{1, 2, 1, 3, 5, 6, 4}, 5},
		{[]int{1}, 0},
		{[]int{6, 5, 4, 3, 2, 3, 2}, 0},
		{[]int{2, 5, 7, 3, 2, 7, 0}, 2}, // 2 peaks
	}

	for i, v := range tests {
		if findPeakElement(v.nums) != v.expected {
			t.Errorf("failed %v", i)
		}
	}
}
