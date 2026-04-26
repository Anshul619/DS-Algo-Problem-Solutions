package main

/*
- LeetCode - https://leetcode.com/problems/balanced-binary-tree
- Time - O(n)
- Space - O(h)
*/

func dfs2(node *TreeNode) (int, bool) {
	if node == nil {
		return 0, true
	}

	leftHeight, leftBalanced := dfs2(node.Left)
	rightHeight, rightBalanced := dfs2(node.Right)

	diff := leftHeight - rightHeight
	if diff < 0 {
		diff = -diff
	}

	isBalanced := diff <= 1 && leftBalanced && rightBalanced

	return max(rightHeight, leftHeight) + 1, isBalanced
}
func isBalanced(root *TreeNode) bool {
	_, isBalanced := dfs2(root)
	return isBalanced
}

// func main() {
// 	root := new(TreeNode) //Returns pointer to TreeNode object
// 	root.Val = 2
// 	root.Left = new(TreeNode)
// 	root.Left.Val = 1
// 	root.Right = new(TreeNode)
// 	root.Right.Val = 3
// 	root.Right.Left = new(TreeNode)
// 	root.Right.Left.Val = 0
// 	root.Right.Right = new(TreeNode)
// 	root.Right.Right.Val = 1
// 	root.Right.Right.Left = new(TreeNode)
// 	root.Right.Right.Left.Val = 4
// 	root.Right.Right.Left.Left = new(TreeNode)
// 	root.Right.Right.Left.Left.Val = 5

// 	//root := new(TreeNode) //Returns pointer to TreeNode object
// 	//root.Val = 1
// 	//root.Left = new(TreeNode)
// 	//root.Left.Val = 2
// 	//root.Left.Left = new(TreeNode)
// 	//root.Left.Left.Val = 3
// 	//root.Left.Left.Left = new(TreeNode)
// 	//root.Left.Left.Left.Val = 4
// 	//root.Right = new(TreeNode)
// 	//root.Right.Val = 2
// 	//root.Right.Right = new(TreeNode)
// 	//root.Right.Right.Val = 3
// 	//root.Right.Right.Right = new(TreeNode)
// 	//root.Right.Right.Right.Val = 4

// 	log.Println(isBalanced(root))

// }
