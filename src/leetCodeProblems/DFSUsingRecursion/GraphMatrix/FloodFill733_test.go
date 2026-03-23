package main

import (
	"reflect"
	"testing"
)

func TestFloodFill(t *testing.T) {

	tests := []struct {
		image [][]int
		sr    int
		sc    int
		color int
		out   [][]int
	}{
		{[][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}, 1, 1, 2, [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}}},
		{[][]int{{0, 0, 0}, {0, 0, 0}}, 0, 0, 0, [][]int{{0, 0, 0}, {0, 0, 0}}},
	}

	for i, v := range tests {
		if !reflect.DeepEqual(floodFill(v.image, v.sr, v.sc, v.color), v.out) {
			t.Errorf("failed %v", i)
		}
	}
}
