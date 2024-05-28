package main

/*
- LeetCode - https://leetcode.com/problems/set-matrix-zeroes/description/
- Time - O(m*n)
- Space - O(1)
*/
import (
	"reflect"
	"testing"
)

func TestSetMatrixZeros(t *testing.T) {
	tests := []struct {
		matrix [][]int
		output [][]int
	}{
		{[][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}, [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}},
		{[][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}, [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}}},
	}

	for i, v := range tests {
		setZeroes(v.matrix)
		if !reflect.DeepEqual(v.matrix, v.output) {
			t.Errorf("%v", i)
		}
	}
}
