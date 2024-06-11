package main

import "testing"

func TestFindKthLargest(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected int
	}{
		{[]int{3, 2, 1, 5, 6, 4}, 2, 5},
		{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4},
	}
	for i, v := range tests {
		if findKthLargest(v.nums, v.k) != v.expected {
			t.Errorf("failed %v", i)
		}
	}
}
