package main

/*
- LeetCode - https://leetcode.com/problems/next-greater-element-i/description/
*/
import "testing"

func TestNextGreaterElement(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		out := nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2})

		if out[1] != 3 {
			t.Fail()
		}
	})

	t.Run("test2", func(t *testing.T) {
		out := nextGreaterElement([]int{2, 4}, []int{1, 2, 3, 4})

		if out[0] != 3 {
			t.Fail()
		}
	})
}
