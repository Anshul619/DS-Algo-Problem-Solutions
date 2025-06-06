package main

/*
- Leetcode - https://leetcode.com/problems/insert-interval/description/
*/
import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	tests := []struct {
		intervals   [][]int
		newInterval []int
		expected    [][]int
	}{
		{[][]int{{1, 3}, {6, 9}}, []int{2, 5}, [][]int{{1, 5}, {6, 9}}},
		{[][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}, [][]int{{1, 2}, {3, 10}, {12, 16}}},
		{[][]int{}, []int{5, 7}, [][]int{{5, 7}}},
		{[][]int{{2, 5}, {6, 7}, {8, 9}}, []int{0, 1}, [][]int{{0, 1}, {2, 5}, {6, 7}, {8, 9}}},
		{[][]int{{0, 4}, {7, 12}}, []int{0, 5}, [][]int{{0, 5}, {7, 12}}},
		{[][]int{{2, 3}, {5, 5}, {6, 6}, {7, 7}, {8, 11}}, []int{6, 13}, [][]int{{2, 3}, {5, 5}, {6, 13}}},
		{[][]int{{1, 5}}, []int{2, 3}, [][]int{{1, 5}}},
	}

	for i, v := range tests {
		if !reflect.DeepEqual(insert(v.intervals, v.newInterval), v.expected) {
			t.Errorf("test failed - %v", i)
		}
	}
}
