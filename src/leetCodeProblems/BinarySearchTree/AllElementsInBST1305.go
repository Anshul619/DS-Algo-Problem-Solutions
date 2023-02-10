package main

/**
- LeetCode - https://leetcode.com/problems/all-elements-in-two-binary-search-trees/description/
- Time - O(m) + O(n) = O(n)
- Space - O(m) + O(n) = O(n)
*/

func inorder(node *TreeNode, out *[]int) {

	if node == nil {
		return
	}

	inorder(node.Left, out)
	*out = append(*out, node.Val)
	inorder(node.Right, out)
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {

	inorder1 := []int{}
	inorder(root1, &inorder1)

	inorder2 := []int{}
	inorder(root2, &inorder2)

	out := []int{}
	i, j := 0, 0

	for i < len(inorder1) && j < len(inorder2) {

		if inorder1[i] <= inorder2[j] {
			out = append(out, inorder1[i])
			i++
		} else {
			out = append(out, inorder2[j])
			j++
		}
	}

	if i != len(inorder1) {
		out = append(out, inorder1[i:]...)
	} else {
		out = append(out, inorder2[j:]...)
	}

	return out
}

// func main() {

// 	root1 := new(TreeNode)
// 	root1.Val = 2
// 	root1.Left = new(TreeNode)
// 	root1.Left.Val = 1
// 	root1.Right = new(TreeNode)
// 	root1.Right.Val = 4

// 	root2 := new(TreeNode)
// 	root2.Val = 1
// 	root2.Left = new(TreeNode)
// 	root2.Left.Val = 0
// 	root2.Right = new(TreeNode)
// 	root2.Right.Val = 3

// 	log.Println(getAllElements(root1, root2))
// }
