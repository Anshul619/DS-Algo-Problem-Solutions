package main

import (
	"reflect"
	"testing"
)

func TestKSmallestPairs(t *testing.T) {
	tests := []struct {
		nums1    []int
		nums2    []int
		k        int
		expected [][]int
	}{
		{[]int{1, 7, 11}, []int{2, 4, 6}, 3, [][]int{{1, 2}, {1, 4}, {1, 6}}},
		{[]int{1, 1, 2}, []int{1, 2, 3}, 2, [][]int{{1, 1}, {1, 1}}},
		{[]int{1, 2, 4, 5, 6}, []int{3, 5, 7, 9}, 3, [][]int{{1, 3}, {2, 3}, {1, 5}}},
	}

	for i, v := range tests {
		if !reflect.DeepEqual(kSmallestPairs(v.nums1, v.nums2, v.k), v.expected) {
			t.Errorf("test failed - %v", i)
		}
	}
}
