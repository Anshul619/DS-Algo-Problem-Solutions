package main

/*
- LeetCode - https://leetcode.com/problems/two-sum/description/
- Time - O(n)
- Space - O(n)
*/

import (
	"testing"
)

func TestTwoSum(t *testing.T) {

	t.Run("TestTwoSum1", func(t *testing.T) {
		out := twoSum([]int{2, 7, 11, 15}, 9)

		if out[0] != 0 && out[1] != 1 {
			t.Fail()
		}
	})

	t.Run("TestTwoSum2", func(t *testing.T) {
		out := twoSum([]int{3, 2, 4}, 6)

		if out[0] != 1 && out[1] != 2 {
			t.Fail()
		}
	})

	t.Run("TestTwoSum3", func(t *testing.T) {
		out := twoSum([]int{3, 3}, 6)

		if out[0] != 0 && out[1] != 1 {
			t.Fail()
		}
	})
}
