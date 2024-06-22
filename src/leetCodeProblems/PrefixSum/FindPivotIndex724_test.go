package main

import "testing"

func TestPivot(t *testing.T) {
	tests := []struct {
		nums []int
		out  int
	}{
		{[]int{1, 7, 3, 6, 5, 6}, 3},
		{[]int{1, 2, 3}, -1},
		{[]int{2, 1, -1}, 0},
	}

	for i, v := range tests {
		if pivotIndex(v.nums) != v.out {
			t.Errorf("failed %v", i)
		}
	}
}
