package main

/*
- LeetCode - https://leetcode.com/problems/two-sum/description/
- Time - O(n)
- Space - O(n)
*/

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {

	tests := []struct {
		nums []int
		out  [][]int
	}{
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{[]int{0, 1, 1}, [][]int{}},
		{[]int{0, 0, 0}, [][]int{{0, 0, 0}}},
		{[]int{2, -3, 0, -2, -5, -5, -4, 1, 2, -2, 2, 0, 2, -4, 5, 5, -10}, [][]int{{-10, 5, 5}, {-5, 0, 5}, {-4, 2, 2}, {-3, -2, 5}, {-3, 1, 2}, {-2, 0, 2}}},
		{[]int{-2, 0, 3, -1, 4, 0, 3, 4, 1, 1, 1, -3, -5, 4, 0}, [][]int{{-5, 1, 4}, {-3, -1, 4}, {-3, 0, 3}, {-2, -1, 3}, {-2, 1, 1}, {-1, 0, 1}, {0, 0, 0}}},
	}

	for i, v := range tests {
		if !reflect.DeepEqual(threeSum(v.nums), v.out) {
			t.Errorf("failed %v", i)
		}
	}
}
