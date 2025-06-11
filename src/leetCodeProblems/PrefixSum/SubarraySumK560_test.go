package main

import "testing"

func TestSubArraySum(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected int
	}{
		{[]int{1, 1, 1}, 2, 2},
		{[]int{1, 2, 3}, 3, 2},
		{[]int{1}, 0, 0},
	}

	for i, v := range tests {
		if subarraySum(v.nums, v.k) != v.expected {
			t.Errorf("failed %v", i)
		}
	}
}
