package main

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{[]int{0}, []int{0}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{0, 0, 0}, []int{0, 0, 0}},
		{[]int{4, 0, 5, 0, 6}, []int{4, 5, 6, 0, 0}},
		{[]int{}, []int{}},
		{[]int{0, 0, 1}, []int{1, 0, 0}},
		{[]int{1, 0, 0}, []int{1, 0, 0}},
		{[]int{0, 1, 0}, []int{1, 0, 0}},
	}

	for _, tt := range tests {
		nums := make([]int, len(tt.input))
		copy(nums, tt.input)
		moveZeroes(nums)
		if !reflect.DeepEqual(nums, tt.expected) {
			t.Errorf("moveZeroes(%v) = %v; want %v", tt.input, nums, tt.expected)
		}
	}
}
