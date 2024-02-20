package main

import "testing"

func TestCanJump(t *testing.T) {
	tests := []struct {
		nums     []int
		expected bool
	}{
		{[]int{2, 3, 1, 1, 4}, true},
		{[]int{3, 2, 1, 0, 4}, false},
		{[]int{0}, true},
		{[]int{1, 2}, true},
		{[]int{1, 2, 3}, true},
	}

	for i, v := range tests {
		if canJump1(v.nums) != v.expected {
			t.Errorf("test failed - %v", i)
		}
	}
}
