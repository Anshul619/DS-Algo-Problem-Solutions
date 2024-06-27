package main

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	tests := []struct {
		digits []int
		out    []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 4}},
		{[]int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
		{[]int{9}, []int{1, 0}},
	}

	for i, v := range tests {
		if !reflect.DeepEqual(plusOne(v.digits), v.out) {
			t.Errorf("failed %v", i)
		}
	}
}
