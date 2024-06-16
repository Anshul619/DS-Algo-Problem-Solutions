package main

import (
	"reflect"
	"testing"
)

func TestAllPathsSourceTarget(t *testing.T) {
	tests := []struct {
		graph [][]int
		paths [][]int
	}{
		{[][]int{{1, 2}, {3}, {3}, {}}, [][]int{{0, 1, 3}, {0, 2, 3}}},
		{[][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}}, [][]int{{0, 4}, {0, 3, 4}, {0, 1, 3, 4}, {0, 1, 2, 3, 4}, {0, 1, 4}}},
	}

	for i, v := range tests {
		if !reflect.DeepEqual(allPathsSourceTarget(v.graph), v.paths) {
			t.Errorf("failed %v", i)
		}
	}
}
