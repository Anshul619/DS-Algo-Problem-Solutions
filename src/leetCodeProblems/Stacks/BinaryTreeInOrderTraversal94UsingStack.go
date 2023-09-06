package main

/*
- Leetcode - https://leetcode.com/problems/binary-tree-inorder-traversal/
- Time - O(n)
- Space - O(n)
*/

type stackTreeNode []*TreeNode

func (s *stackTreeNode) pop() *TreeNode {
	temp := (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
	return temp
}

func (s *stackTreeNode) push(node *TreeNode) {
	*s = append(*s, node)
}
func (s stackTreeNode) isEmpty() bool {
	return len(s) == 0
}

func (s stackTreeNode) peek() *TreeNode {
	return s[len(s)-1]
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	out := []int{}
	toProcess := root

	s := stackTreeNode{}

	for !s.isEmpty() || toProcess != nil {

		for toProcess != nil {
			s.push(toProcess)
			toProcess = toProcess.Left
		}

		temp := s.pop()
		out = append(out, temp.Val)
		toProcess = temp.Right
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

// 	log.Println(inorderTraversal(root))
// }
