package main

/**
- LeetCode - https://leetcode.com/problems/minimum-depth-of-binary-tree/
- Time - O(n)
- Space - O(1)
*/

func maxDepthUtil(node *TreeNode, height int, maxHeight *int) {
	if node == nil {
		return
	}
	if height > *maxHeight {
		*maxHeight = height
	}

	maxDepthUtil(node.Left, height+1, maxHeight)
	maxDepthUtil(node.Right, height+1, maxHeight)
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	out := 0
	maxDepthUtil(root, 0, &out)
	return out + 1
}

// func main() {
// 	root := new(TreeNode) //Returns pointer to TreeNode object
// 	root.Val = 4
// 	root.Left = new(TreeNode)
// 	root.Left.Val = 2
// 	root.Left.Left = new(TreeNode)
// 	root.Left.Left.Val = 1
// 	root.Left.Right = new(TreeNode)
// 	root.Left.Right.Val = 3

// 	root.Right = new(TreeNode)
// 	root.Right.Val = 7
// 	root.Right.Left = new(TreeNode)
// 	root.Right.Left.Val = 6
// 	root.Right.Right = new(TreeNode)
// 	root.Right.Right.Val = 9

// 	log.Println(maxDepth(root))
// }
