package main

import (
	"reflect"
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		nums []int
		out  bool
	}{
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 2, 3, 4}, false},
		{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
	}

	for i, v := range tests {
		if !reflect.DeepEqual(containsDuplicate(v.nums), v.out) {
			t.Errorf("failed %v", i)
		}
	}
}
