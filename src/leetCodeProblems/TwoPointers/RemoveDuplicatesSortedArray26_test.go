package main

/*
- LeetCode - https://leetcode.com/problems/remove-duplicates-from-sorted-array/
*/

import "testing"

func TestRemoveDuplicates(t *testing.T) {

	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 1, 2}, []int{1, 2}},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, []int{0, 1, 2, 3, 4}},
	}

	for _, v := range tests {
		out := removeDuplicates(v.input)

		for j := 0; j < out; j++ {
			if v.input[j] != v.expected[j] {
				t.Fail()
			}
		}
	}
}
