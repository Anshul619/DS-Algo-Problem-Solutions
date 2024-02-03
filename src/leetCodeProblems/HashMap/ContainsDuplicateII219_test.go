package main

/*
- LeetCode - https://leetcode.com/problems/contains-duplicate-ii
*/
import "testing"

func TestContainsNearbyDuplicate(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected bool
	}{
		{[]int{1, 2, 3, 1}, 3, true},
		{[]int{1, 0, 1, 1}, 1, true},
		{[]int{1, 2, 3, 1, 2, 3}, 2, false},
	}

	for _, v := range tests {
		if containsNearbyDuplicate(v.nums, v.k) != v.expected {
			t.Fail()
		}
	}
}
