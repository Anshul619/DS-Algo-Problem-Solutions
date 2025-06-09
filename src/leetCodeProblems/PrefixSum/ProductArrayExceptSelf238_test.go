package main

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		nums     []int
		expected []int
	}{
		{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{[]int{2, 2, 2, 2}, []int{8, 8, 8, 8}},
		{[]int{1, 2, 0, 4}, []int{0, 0, 8, 0}},
		{[]int{0, 2, 0, 4}, []int{0, 0, 0, 0}},
		{[]int{3, 5}, []int{5, 3}},
		{[]int{42}, []int{1}},
	}

	for _, test := range tests {
		result := productExceptSelf(test.nums)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For input %v, expected %v, got %v", test.nums, test.expected, result)
		}
	}
}
