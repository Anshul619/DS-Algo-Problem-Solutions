package main

import "testing"

func TestRob(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 2, 3, 1}, 4},
		{[]int{2, 7, 9, 3, 1}, 12},
		{[]int{2, 1, 1, 2}, 4},
	}

	for i, v := range tests {
		if rob(v.nums) != v.expected {
			t.Errorf("failed %v", i)
		}
	}
}
