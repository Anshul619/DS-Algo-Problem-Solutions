package main

/*
- LeetCode - https://leetcode.com/problems/search-in-rotated-sorted-array/description/
*/

import "testing"

func TestSearchSortedArray(t *testing.T) {
	tests := []struct {
		nums     []int
		target   int
		expected int
	}{
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{[]int{1}, 0, -1},
		{[]int{1, 3}, 0, -1},
		{[]int{3, 1}, 1, 1},
	}

	for i, v := range tests {
		if searchSortedArray(v.nums, v.target) != v.expected {
			t.Fatalf("test failed %v", i)
		}
	}
}
