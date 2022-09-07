package main

/*
- LeetCode - https://leetcode.com/problems/maximum-width-of-binary-tree
*/
// import "log"

// type TreeNode struct {
// 	Left  *TreeNode
// 	Right *TreeNode
// 	Val   int
// }

type TreeNodeIndexed struct {
	Node  *TreeNode
	Index int
}

type queue []TreeNodeIndexed

func (q *queue) offer(node TreeNodeIndexed) {
	*q = append(*q, node)
}

func (q *queue) poll() (bool, TreeNodeIndexed) {

	if !q.isEmpty() {
		node := (*q)[0]
		*q = (*q)[1:]
		return true, node
	}

	return false, *new(TreeNodeIndexed)
}

func (q *queue) size() int {
	return len(*q)
}

func (q *queue) isEmpty() bool {
	return q.size() <= 0
}

func widthOfBinaryTree(root *TreeNode) int {

	if root == nil {
		return 0
	}

	maxWidth := 1
	nodesQ := new(queue)

	if root.Left != nil {
		nodesQ.offer(TreeNodeIndexed{root.Left, 2})
	}

	if root.Right != nil {
		nodesQ.offer(TreeNodeIndexed{root.Right, 3})
	}

	//log.Println("Outer Size", nodesQ.size())
	//log.Println(maxWidth)
	for !nodesQ.isEmpty() {

		size := nodesQ.size()
		leftMostIndex, rightMostIndex := 0, 0

		for size > 0 {

			_, nodeI := nodesQ.poll()

			if leftMostIndex == 0 {
				leftMostIndex = nodeI.Index
			}

			rightMostIndex = nodeI.Index

			if nodeI.Node.Left != nil {
				nodesQ.offer(TreeNodeIndexed{nodeI.Node.Left, nodeI.Index * 2})
			}

			if nodeI.Node.Right != nil {
				nodesQ.offer(TreeNodeIndexed{nodeI.Node.Right, nodeI.Index*2 + 1})
			}

			size--

			//log.Println("Decremented Size", size)
		}

		width := rightMostIndex - leftMostIndex + 1

		//log.Println(width)
		if maxWidth < width {
			maxWidth = width
		}
	}

	return maxWidth
}

// func main() {

// 	root := new(TreeNode) //Returns pointer to TreeNode object
// 	root.Val = 1
// 	root.Left = new(TreeNode)
// 	root.Left.Val = 3
// 	root.Right = new(TreeNode)
// 	root.Right.Val = 2
// 	root.Left.Left = new(TreeNode)
// 	root.Left.Left.Val = 5
// 	root.Left.Right = new(TreeNode)
// 	root.Left.Right.Val = 3

// 	// 	root.Left.Left.Left = new(TreeNode)
// 	// 	root.Left.Left.Left.Val = 7

// 	root.Right.Right = new(TreeNode)
// 	root.Right.Right.Val = 9
// 	// 	root.Right.Right.Right = new(TreeNode)
// 	// 	root.Right.Right.Right.Val = 8

// 	log.Println(widthOfBinaryTree(root))
// }
