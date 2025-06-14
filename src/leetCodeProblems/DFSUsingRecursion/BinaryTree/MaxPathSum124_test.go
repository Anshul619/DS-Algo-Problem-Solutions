package main

import (
	"fmt"
	"testing"
)

func TestMaxPathSum(t *testing.T) {
	tests := []struct {
		tree     *TreeNode
		expected int
	}{
		// Example 1
		{
			&TreeNode{1,
				&TreeNode{2, nil, nil},
				&TreeNode{3, nil, nil}},
			6,
		},
		// Example 2
		{
			&TreeNode{-10,
				&TreeNode{9, nil, nil},
				&TreeNode{20,
					&TreeNode{15, nil, nil},
					&TreeNode{7, nil, nil}}},
			42,
		},
		// Single node
		{
			&TreeNode{1, nil, nil},
			1,
		},
		// All negative nodes
		{
			&TreeNode{-3,
				&TreeNode{-2, nil, nil},
				&TreeNode{-1, nil, nil}},
			-1,
		},
		// Skewed right
		{
			&TreeNode{1,
				nil,
				&TreeNode{2,
					nil,
					&TreeNode{3, nil, nil}}},
			6,
		},
		// Complex mixed
		{
			&TreeNode{10,
				&TreeNode{2,
					&TreeNode{20, nil, nil},
					&TreeNode{1, nil, nil}},
				&TreeNode{10,
					nil,
					&TreeNode{-25,
						&TreeNode{3, nil, nil},
						&TreeNode{4, nil, nil}}}},
			42,
		},
		// Big negative root + positive subtree
		{
			&TreeNode{-10,
				&TreeNode{-20, nil, nil},
				&TreeNode{5,
					nil,
					&TreeNode{25, nil, nil}}},
			30,
		},
	}

	for i, test := range tests {
		result := maxPathSum(test.tree)
		fmt.Printf("Test %d: Expected %d, Got %d\n", i+1, test.expected, result)
	}
}
