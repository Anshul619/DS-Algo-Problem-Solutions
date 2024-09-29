package main

import "testing"

/*
- LeetCode - https://leetcode.com/problems/majority-element/
*/
func TestMajorityElemenet(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{3, 2, 3}, 3},
		{[]int{2, 2, 1, 1, 1, 2, 2}, 2},
	}

	for i, v := range tests {
		if majorityElement1(v.nums) != v.expected {
			t.Errorf("failed case %v", i)
		}
	}
}
