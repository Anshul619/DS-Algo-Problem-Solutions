package main

/*
- LeetCode - https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
*/

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		numbers  []int
		target   int
		expected []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{1, 2}},
		{[]int{2, 3, 4}, 6, []int{1, 3}},
		{[]int{-1, 0}, -1, []int{1, 2}},
	}

	for i, v := range tests {
		if !reflect.DeepEqual(twoSum1(v.numbers, v.target), v.expected) {
			t.Errorf("failed %v", i)
		}
	}
}
