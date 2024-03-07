package main

/*
- LeetCode - https://leetcode.com/problems/spiral-matrix/
*/
import (
	"reflect"
	"testing"
)

func TestSpiralMatrix(t *testing.T) {
	tests := []struct {
		matrix [][]int
		out    []int
	}{
		//{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
		{[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}, []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}},
	}

	for i, v := range tests {
		if !reflect.DeepEqual(spiralOrder(v.matrix), v.out) {
			t.Errorf("failed %v", i)
		}
	}
}
