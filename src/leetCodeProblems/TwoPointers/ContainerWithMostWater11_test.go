package main

import "testing"

func TestMaxArea(t *testing.T) {
	tests := []struct {
		height []int
		out    int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1}, 1},
	}

	for i, v := range tests {
		if maxArea(v.height) != v.out {
			t.Errorf("failed %v", i)
		}
	}
}
