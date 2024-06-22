package main

import (
	"reflect"
	"testing"
)

func TestLeftRightDiff(t *testing.T) {
	tests := []struct {
		nums []int
		out  []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{-3, -1, 1, 3, 5}},
		{[]int{3, 2, 3, 4, 2}, []int{-2, -1, 0, 2, 3}},
	}

	for i, v := range tests {
		if !reflect.DeepEqual(distinctDifferenceArray(v.nums), v.out) {
			t.Errorf("failed %v", i)
		}
	}
}
