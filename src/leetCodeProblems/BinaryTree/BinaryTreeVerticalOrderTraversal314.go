package main

/*
- LeetCode - https://leetcode.com/problems/binary-tree-vertical-order-traversal/
*/
import (
	"log"
	"sort"
)

type TreeVNode struct {
	node   *TreeNode
	vOrder int
}

type tQueue []TreeVNode

func (q *tQueue) pop() TreeVNode {
	out := (*q)[0]
	*q = (*q)[1:]
	return out
}

func (q *tQueue) push(t TreeVNode) {
	*q = append(*q, t)
}

func (q *tQueue) isEmpty() bool {
	return len(*q) == 0
}

func findMinMaxVOrder(node *TreeNode, current int, min *int, max *int) {

	if node == nil {
		return
	}

	leftVOrder := current
	rightVOrder := current

	if node.Left != nil {
		leftVOrder--
	}

	if node.Right != nil {
		rightVOrder++
	}

	if *min > leftVOrder {
		*min = leftVOrder
	}

	if *max < rightVOrder {
		*max = rightVOrder
	}

	findMinMaxVOrder(node.Left, leftVOrder, min, max)
	findMinMaxVOrder(node.Right, rightVOrder, min, max)
}

func verticalOrder(root *TreeNode) [][]int {
	min, max := 0, 0

	findMinMaxVOrder(root, 0, &min, &max)

	total := max - min + 1

	log.Println(min, max, total)

	queue := new(tQueue)
	queue.push(TreeVNode{root, 0})

	out := make([][]int, total)

	for i := range out {
		out[i] = []int{}
	}

	for !queue.isEmpty() {
		cTreeVNode := queue.pop()
		outIndex := cTreeVNode.vOrder - min
		log.Println(outIndex)
		out[outIndex] = append(out[outIndex], cTreeVNode.node.Val)

		if cTreeVNode.node.Left != nil {
			queue.push(TreeVNode{cTreeVNode.node.Left, cTreeVNode.vOrder - 1})
		}

		if cTreeVNode.node.Right != nil {
			queue.push(TreeVNode{cTreeVNode.node.Right, cTreeVNode.vOrder + 1})
		}
	}

	for i := range out {
		sort.Ints(out[i])
	}
	return out

}

func main() {
	// root := new(TreeNode) //Returns pointer to TreeNode object
	// root.Val = 3
	// root.Left = new(TreeNode)
	// root.Left.Val = 9
	// root.Right = new(TreeNode)
	// root.Right.Val = 20
	// root.Right.Left = new(TreeNode)
	// root.Right.Left.Val = 15
	// root.Right.Right = new(TreeNode)
	// root.Right.Right.Val = 7

	// root := new(TreeNode) //Returns pointer to TreeNode object
	// root.Val = 1
	// root.Left = new(TreeNode)
	// root.Left.Val = 2
	// root.Left.Left = new(TreeNode)
	// root.Left.Left.Val = 4
	// root.Left.Right = new(TreeNode)
	// root.Left.Right.Val = 5

	// root.Right = new(TreeNode)
	// root.Right.Val = 3
	// root.Right.Left = new(TreeNode)
	// root.Right.Left.Val = 6
	// root.Right.Right = new(TreeNode)
	// root.Right.Right.Val = 7

	root := new(TreeNode) //Returns pointer to TreeNode object
	root.Val = 0
	root.Left = new(TreeNode)
	root.Left.Val = 1
	root.Left.Right = new(TreeNode)
	root.Left.Right.Val = 2
	root.Left.Right.Left = new(TreeNode)
	root.Left.Right.Left.Val = 6
	root.Left.Right.Right = new(TreeNode)
	root.Left.Right.Right.Val = 3
	root.Left.Right.Right.Right = new(TreeNode)
	root.Left.Right.Right.Right.Val = 4
	root.Left.Right.Right.Right.Right = new(TreeNode)
	root.Left.Right.Right.Right.Right.Val = 5

	log.Println(verticalTraversal(root))
}
