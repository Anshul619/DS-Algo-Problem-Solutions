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
	}

	for i, v := range tests {
		if !reflect.DeepEqual(threeSum(v.nums), v.out) {
			t.Errorf("failed %v", i)
		}
	}
}
