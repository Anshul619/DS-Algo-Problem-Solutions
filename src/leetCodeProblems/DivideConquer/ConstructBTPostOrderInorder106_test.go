package main

import (
	"testing"
)

func TestConstructBTPostOrderInorder(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		node := buildTree1([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3})

		printInOrder(node)
	})
}
