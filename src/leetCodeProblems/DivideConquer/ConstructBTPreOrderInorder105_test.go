package main

import (
	"testing"
)

func TestBuildTree(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		node := buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})

		printInOrder(node)
	})
}
