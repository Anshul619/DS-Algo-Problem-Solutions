package main

import (
	"reflect"
	"testing"
)

func TestIntersect(t *testing.T) {
	tests := []struct {
		nums1 []int
		nums2 []int
		out   []int
	}{
		{[]int{1, 2, 2, 1}, []int{2, 2}, []int{2, 2}},
		{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}, []int{9, 4}},
		{[]int{3, 1, 2}, []int{1, 1}, []int{1}},
	}

	for i, v := range tests {
		if !reflect.DeepEqual(intersect2(v.nums1, v.nums2), v.out) {
			t.Errorf("failed %v", i)
		}
	}
}
