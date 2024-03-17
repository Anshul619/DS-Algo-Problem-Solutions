package main

import "testing"

func TestMinArrows(t *testing.T) {
	tests := []struct {
		points   [][]int
		expected int
	}{
		{[][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}, 2},
		{[][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}, 4},
		{[][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}, 2},
	}

	for i, v := range tests {
		if findMinArrowShots(v.points) != v.expected {
			t.Errorf("failed %v", i)
		}
	}
}
