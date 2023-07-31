package main

/*
- Leetcode - https://leetcode.com/problems/binary-tree-preorder-traversal/description/
- Time - O(n)
- Space - O(n)
*/
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	out := []int{}

	s := stackTreeNode{}

	s.push(root)

	for !s.isEmpty() {
		node := s.pop()

		out = append(out, node.Val)

		if node.Right != nil {
			s.push(node.Right)
		}

		if node.Left != nil {
			s.push(node.Left)
		}
	}

	return out
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

// 	log.Println(preorderTraversal(root))
// }
