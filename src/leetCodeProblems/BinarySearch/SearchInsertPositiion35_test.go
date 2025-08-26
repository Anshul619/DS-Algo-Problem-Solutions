package main

/*
- LeetCode - https://leetcode.com/problems/search-in-rotated-sorted-array/description/
*/

import "testing"

func TestSearchInsertPosition(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		out    int
	}{
		{[]int{1, 3, 5, 6}, 5, 2},
		{[]int{1, 3, 5, 6}, 2, 1},
		{[]int{1, 3, 5, 6}, 7, 4},
		{[]int{1, 3, 6, 10, 15}, 11, 4},
		{[]int{1, 3, 6, 10, 15}, 15, 4},
	}

	for i, v := range tests {
		if searchInsert(v.nums, v.target) != v.out {
			t.Errorf("failed test case %v", i)
		}
	}
}
