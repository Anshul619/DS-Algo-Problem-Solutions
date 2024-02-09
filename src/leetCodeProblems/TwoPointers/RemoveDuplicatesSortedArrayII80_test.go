package main

/*
- LeetCode - https://leetcode.com/problems/remove-duplicates-from-sorted-array/
*/

import (
	"testing"
)

func TestRemoveDuplicatesII(t *testing.T) {

	tests := []struct {
		input     []int
		expected  []int
		expectedK int
	}{
		{[]int{1, 1, 1, 2, 2, 3}, []int{1, 1, 2, 2, 3}, 5},
		{[]int{0, 0, 1, 1, 1, 1, 2, 3, 3}, []int{0, 0, 1, 1, 2, 3, 3}, 7},
	}

	for _, v := range tests {
		out := removeDuplicatesII(v.input)

		if out != v.expectedK {
			t.Fail()
		}
		for j := 0; j < out; j++ {
			if v.input[j] != v.expected[j] {
				t.Fail()
			}
		}
	}
}
