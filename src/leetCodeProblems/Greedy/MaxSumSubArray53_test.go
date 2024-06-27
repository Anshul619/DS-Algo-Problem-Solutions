package main

import (
	"reflect"
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		nums []int
		out  int
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{[]int{1}, 1},
		{[]int{5, 4, -1, 7, 8}, 23},
	}

	for i, v := range tests {
		if !reflect.DeepEqual(maxSubArray(v.nums), v.out) {
			t.Errorf("failed %v", i)
		}
	}
}
