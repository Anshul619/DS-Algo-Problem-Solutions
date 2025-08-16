package main

import "log"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func printInOrder(node *TreeNode) {

	if node == nil {
		log.Println("null")
		return
	}

	log.Println(node.Val)
	printInOrder(node.Left)

	printInOrder(node.Right)
}
