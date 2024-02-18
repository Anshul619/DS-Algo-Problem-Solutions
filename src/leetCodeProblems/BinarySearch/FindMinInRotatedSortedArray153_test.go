package main

/*
- Leetcode - https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/description/
*/
import "testing"

func TestFindMin(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{3, 4, 5, 1, 2}, 1},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0},
		{[]int{11, 13, 15, 17}, 11},
		{[]int{3, 1, 2}, 1},
	}

	for i, v := range tests {
		if findMin(v.nums) != v.expected {
			t.Errorf("test failed - %v", i)
		}
	}
}
