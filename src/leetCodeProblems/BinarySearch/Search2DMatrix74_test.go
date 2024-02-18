package main

import "testing"

func TestSearchMatrix(t *testing.T) {
	tests := []struct {
		matrix   [][]int
		target   int
		expected bool
	}{
		{[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3, true},
		{[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13, false},
	}

	for i, v := range tests {
		if searchMatrix(v.matrix, v.target) != v.expected {
			t.Errorf("test failed - %v", i)
		}
	}
}
