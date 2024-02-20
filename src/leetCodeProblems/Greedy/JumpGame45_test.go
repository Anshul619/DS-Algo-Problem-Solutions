package main

import "testing"

func TestJump(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{2, 3, 1, 1, 4}, 2},
		{[]int{2, 3, 0, 1, 4}, 2},
	}

	for i, v := range tests {
		if jump(v.nums) != v.expected {
			t.Errorf("test failed - %v", i)
		}
	}
}
