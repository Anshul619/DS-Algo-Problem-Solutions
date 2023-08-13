package main

/*
- Leetcode - https://leetcode.com/problems/symmetric-tree/
- Time - O(n)
- Space - O(1)
*/

func isSymmetricUtil(node1 *TreeNode, node2 *TreeNode) bool {

	if node1 == nil && node2 == nil {
		return true
	}

	if node1 == nil || node2 == nil {
		return false
	}

	if node1.Val == node2.Val &&
		isSymmetricUtil(node1.Left, node2.Right) &&
		isSymmetricUtil(node1.Right, node2.Left) {
		return true
	}

	return false
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isSymmetricUtil(root.Left, root.Right)
}

// func main() {
// 	// root := new(TreeNode) //Returns pointer to TreeNode object
// 	// root.Val = 1
// 	// root.Left = new(TreeNode)
// 	// root.Left.Val = 2
// 	// root.Left.Left = new(TreeNode)
// 	// root.Left.Left.Val = 3
// 	// root.Left.Right = new(TreeNode)
// 	// root.Left.Right.Val = 4

// 	// root.Right = new(TreeNode)
// 	// root.Right.Val = 2
// 	// root.Right.Left = new(TreeNode)
// 	// root.Right.Left.Val = 4
// 	// root.Right.Right = new(TreeNode)
// 	// root.Right.Right.Val = 3

// 	// log.Println(isSymmetric(root))

// 	root := new(TreeNode) //Returns pointer to TreeNode object
// 	root.Val = 1
// 	root.Left = new(TreeNode)
// 	root.Left.Val = 2
// 	root.Left.Right = new(TreeNode)
// 	root.Left.Right.Val = 3

// 	root.Right = new(TreeNode)
// 	root.Right.Val = 2
// 	root.Right.Right = new(TreeNode)
// 	root.Right.Right.Val = 3

// 	log.Println(isSymmetric(root))

// }
