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
		{[]int{10, 4, 8, 3}, []int{15, 1, 11, 22}},
		{[]int{1}, []int{0}},
	}

	for i, v := range tests {
		if !reflect.DeepEqual(leftRightDifference(v.nums), v.out) {
			t.Errorf("failed %v", i)
		}
	}
}
