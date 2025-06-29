package main

import (
	"reflect"
	"testing"
)

func TestCanFinish1(t *testing.T) {
	tests := []struct {
		numCourses    int
		prerequisites [][]int
		expected      []int
	}{
		{2, [][]int{{1, 0}}, []int{0, 1}},
		{4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}, []int{0, 2, 1, 3}},
		{1, [][]int{}, []int{0}},
		{3, [][]int{{1, 0}, {2, 0}, {0, 2}}, []int{}},
	}

	for i, v := range tests {
		if !reflect.DeepEqual(findOrder(v.numCourses, v.prerequisites), v.expected) {
			t.Errorf("failed %v", i)
		}
	}
}
