package main

import "testing"

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{100, 4, 200, 1, 3, 2}, 4},
	}

	for i, v := range tests {
		if longestConsecutive(v.nums) != v.expected {
			t.Errorf("failed %v", i)
		}
	}
}
