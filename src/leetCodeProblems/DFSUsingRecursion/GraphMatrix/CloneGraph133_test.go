package main

import (
	"log"
	"testing"
)

func TestCloneGraph(t *testing.T) {

	t.Run("test1", func(t *testing.T) {
		node1 := new(Node)
		node1.Val = 1

		node2 := new(Node)
		node2.Val = 2

		node3 := new(Node)
		node3.Val = 3

		node4 := new(Node)
		node4.Val = 4

		node1.Neighbors = []*Node{node2, node4}
		node2.Neighbors = []*Node{node1, node3}
		node3.Neighbors = []*Node{node2, node4}
		node4.Neighbors = []*Node{node1, node3}

		newNode1 := cloneGraph(node1)

		if newNode1.Val != 1 {
			t.Errorf("failed")
		}

		for _, v := range newNode1.Neighbors {
			log.Println(v.Val)
		}
	})
}
