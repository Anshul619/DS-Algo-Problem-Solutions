package main

import (
	"reflect"
	"testing"
)

func TestSearchRange(t *testing.T) {
	tests := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
		{[]int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
		{[]int{}, 0, []int{-1, -1}},
		{[]int{2, 2}, 3, []int{-1, -1}},
	}

	for i, v := range tests {
		if !reflect.DeepEqual(searchRange(v.nums, v.target), v.expected) {
			t.Errorf("test failed - %v", i)
		}
	}
}
