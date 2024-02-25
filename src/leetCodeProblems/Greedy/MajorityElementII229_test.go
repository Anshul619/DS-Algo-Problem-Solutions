package main

import (
	"reflect"
	"testing"
)

func TestMajorityElementII(t *testing.T) {
	tests := []struct {
		nums     []int
		expected []int
	}{
		{[]int{3, 2, 3}, []int{3}},
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{1, 2}},
		{[]int{2, 2}, []int{2}},
		{[]int{0, 3, 4, 0}, []int{0}},
	}

	for i, v := range tests {
		if !reflect.DeepEqual(majorityElement(v.nums), v.expected) {
			t.Errorf("fail %v", i)
		}
	}
}
