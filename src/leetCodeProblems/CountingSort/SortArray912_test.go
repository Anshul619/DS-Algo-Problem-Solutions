package main

import (
	"reflect"
	"testing"
)

func TestSortArray(t *testing.T) {
	tests := []struct {
		nums     []int
		expected []int
	}{
		{[]int{5, 2, 3, 1}, []int{1, 2, 3, 5}},
	}

	for i, v := range tests {
		if !reflect.DeepEqual(sortArray(v.nums), v.expected) {
			t.Errorf("fail %v", i)
		}
	}
}
