package main

/*
- Leetcode - https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/description/
*/
import (
	"testing"
)

func TestValidateBinaryTree(t *testing.T) {

	tests := []struct {
		n          int
		leftChild  []int
		rightChild []int
		expected   bool
	}{
		{4, []int{1, -1, 3, -1}, []int{2, -1, -1, -1}, true},
		{4, []int{1, -1, 3, -1}, []int{2, 3, -1, -1}, false},
		{2, []int{1, 0}, []int{-1, -1}, false},
		{4, []int{1, 0, 3, -1}, []int{-1, -1, -1, -1}, false},
	}

	for i, v := range tests {
		if validateBinaryTreeNodes(v.n, v.leftChild, v.rightChild) != v.expected {
			t.Errorf("failed test %v", i)
		}
	}
}
