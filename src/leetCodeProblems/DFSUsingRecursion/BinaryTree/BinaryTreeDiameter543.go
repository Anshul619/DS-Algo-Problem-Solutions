package main

/*
- Leetcode - https://leetcode.com/problems/diameter-of-binary-tree/
- Time - O(n^2)
- Space - O(1)
*/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func heightUtil(node *TreeNode, prevHeight int, maxHeight *int) {
	if node == nil {
		return
	}

	if prevHeight > *maxHeight {
		*maxHeight = prevHeight
	}

	prevHeight++

	heightUtil(node.Left, prevHeight, maxHeight)
	heightUtil(node.Right, prevHeight, maxHeight)
}

func height(node *TreeNode) int {
	out := 0
	heightUtil(node, 0, &out)
	return out
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight, rightHeight := 0, 0

	if root.Left != nil {
		leftHeight = height(root.Left) + 1
	}

	if root.Right != nil {
		rightHeight = height(root.Right) + 1
	}

	return max(max(diameterOfBinaryTree(root.Left), diameterOfBinaryTree(root.Right)),
		leftHeight+rightHeight)
}

// func main() {
// 	// root := new(TreeNode) //Returns pointer to TreeNode object
// 	// root.Val = 4
// 	// root.Left = new(TreeNode)
// 	// root.Left.Val = 2
// 	// root.Left.Left = new(TreeNode)
// 	// root.Left.Left.Val = 1
// 	// root.Left.Right = new(TreeNode)
// 	// root.Left.Right.Val = 3

// 	// root.Right = new(TreeNode)
// 	// root.Right.Val = 7
// 	// root.Right.Left = new(TreeNode)
// 	// root.Right.Left.Val = 6
// 	// root.Right.Right = new(TreeNode)
// 	// root.Right.Right.Val = 9

// 	// // maxHeight := 0
// 	// // height(root, 0, &maxHeight)
// 	// // log.Println(maxHeight)

// 	// log.Println(diameterOfBinaryTree(root))

// 	root := new(TreeNode) //Returns pointer to TreeNode object
// 	root.Val = 1
// 	root.Left = new(TreeNode)
// 	root.Left.Val = 2
// 	root.Left.Left = new(TreeNode)
// 	root.Left.Left.Val = 4
// 	root.Left.Right = new(TreeNode)
// 	root.Left.Right.Val = 5

// 	root.Right = new(TreeNode)
// 	root.Right.Val = 3

// 	// maxHeight := 0
// 	// height(root, 0, &maxHeight)
// 	// log.Println(maxHeight)

// 	log.Println(diameterOfBinaryTree(root))
// }
