package main

/*
- Leetcode - https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/description/
- Time - O(n)
- Space - O(1)
*/

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root == p {
		return root
	}

	if root == q {
		return root
	}

	leftLCA := lowestCommonAncestor(root.Left, p, q)
	rightLCA := lowestCommonAncestor(root.Right, p, q)

	if leftLCA != nil && rightLCA != nil {
		return root
	}

	if leftLCA != nil {
		return leftLCA
	} else {
		return rightLCA
	}
}

// func main() {
// 	root := new(TreeNode) //Returns pointer to TreeNode object
// 	root.Val = 3
// 	root.Left = new(TreeNode)
// 	root.Left.Val = 5
// 	root.Left.Left = new(TreeNode)
// 	root.Left.Left.Val = 6
// 	root.Left.Right = new(TreeNode)
// 	root.Left.Right.Val = 2
// 	root.Left.Right.Left = new(TreeNode)
// 	root.Left.Right.Left.Val = 7
// 	root.Left.Right.Right = new(TreeNode)
// 	root.Left.Right.Right.Val = 4

// 	root.Right = new(TreeNode)
// 	root.Right.Val = 1
// 	root.Right.Left = new(TreeNode)
// 	root.Right.Left.Val = 0
// 	root.Right.Right = new(TreeNode)
// 	root.Right.Right.Val = 8

// 	//log.Println(lowestCommonAncestor(root, root.Left, root.Right))
// 	log.Println(lowestCommonAncestor(root, root.Left, root.Left.Right.Right))

// }
