package main

import "testing"

func TestCanFinish(t *testing.T) {
	tests := []struct {
		numCourses    int
		prerequisites [][]int
		expected      bool
	}{
		{2, [][]int{{1, 0}}, true},
		{2, [][]int{{1, 0}, {0, 1}}, false},
		{20, [][]int{{0, 10}, {3, 18}, {5, 5}, {6, 11}, {11, 14}, {13, 1}, {15, 1}, {17, 4}}, false},
		{2, [][]int{}, true},
		{2, [][]int{{0, 1}}, true},
		{5, [][]int{{1, 4}, {2, 4}, {3, 1}, {3, 2}}, true},
	}

	for i, v := range tests {
		if canFinish(v.numCourses, v.prerequisites) != v.expected {
			t.Errorf("failed %v", i)
		}
	}
}
