package main

/*
- Leetcode - https://leetcode.com/problems/symmetric-tree/
- Time - O(n)
- Space - O(1)
*/

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	s := stackTreeNode{}

	s.push(root)
	s.push(root)

	for !s.isEmpty() {

		s.print()
		node1 := s.pop()
		node2 := s.pop()

		if node1 == nil && node2 == nil {
			continue
		}

		if node1 == nil || node2 == nil {
			return false
		}

		if node1.Val != node2.Val {
			return false
		}

		s.push(node1.Left)
		s.push(node2.Right)

		if node1 != node2 {
			s.push(node1.Right)
			s.push(node2.Left)
		}
	}

	return true
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

// 	// root := new(TreeNode) //Returns pointer to TreeNode object
// 	// root.Val = 1
// 	// root.Left = new(TreeNode)
// 	// root.Left.Val = 2
// 	// root.Left.Right = new(TreeNode)
// 	// root.Left.Right.Val = 3

// 	// root.Right = new(TreeNode)
// 	// root.Right.Val = 2
// 	// root.Right.Right = new(TreeNode)
// 	// root.Right.Right.Val = 3

// 	// log.Println(isSymmetric(root))

// 	// root := new(TreeNode) //Returns pointer to TreeNode object
// 	// root.Val = 2
// 	// root.Left = new(TreeNode)
// 	// root.Left.Val = 3
// 	// root.Left.Left = new(TreeNode)
// 	// root.Left.Left.Val = 4
// 	// root.Left.Right = new(TreeNode)
// 	// root.Left.Right.Val = 5

// 	// root.Right = new(TreeNode)
// 	// root.Right.Val = 3
// 	// root.Right.Left = new(TreeNode)
// 	// root.Right.Left.Val = 5

// 	// log.Println(isSymmetric(root))

// 	root := new(TreeNode) //Returns pointer to TreeNode object
// 	root.Val = 2
// 	root.Left = new(TreeNode)
// 	root.Left.Val = 3
// 	root.Left.Left = new(TreeNode)
// 	root.Left.Left.Val = 4
// 	root.Left.Right = new(TreeNode)
// 	root.Left.Right.Val = 5

// 	root.Right = new(TreeNode)
// 	root.Right.Val = 3
// 	root.Right.Right = new(TreeNode)
// 	root.Right.Right.Val = 4

// 	log.Println(isSymmetric(root))

// }
