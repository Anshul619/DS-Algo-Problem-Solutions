package main

/*
LeetCode - https://leetcode.com/problems/sum-of-nodes-with-even-valued-grandparent/
*/

func util(node *TreeNode, sum *int) {
	if node == nil {
		return
	}

	if node.Val%2 == 0 {

		if node.Left != nil {
			if node.Left.Left != nil {
				*sum += node.Left.Left.Val
			}

			if node.Left.Right != nil {
				*sum += node.Left.Right.Val
			}
		}

		if node.Right != nil {
			if node.Right.Left != nil {
				*sum += node.Right.Left.Val
			}

			if node.Right.Right != nil {
				*sum += node.Right.Right.Val
			}
		}
	}

	util(node.Left, sum)
	util(node.Right, sum)

}
func sumEvenGrandparent(root *TreeNode) int {

	out := 0
	util(root, &out)
	return out
}

// func main() {
// 	root := new(TreeNode) //Returns pointer to TreeNode object
// 	root.Val = 6
// 	root.Left = new(TreeNode)
// 	root.Left.Val = 7
// 	root.Left.Left = new(TreeNode)
// 	root.Left.Left.Val = 2
// 	root.Left.Right = new(TreeNode)
// 	root.Left.Right.Val = 7
// 	root.Left.Right.Right = new(TreeNode)
// 	root.Left.Right.Right.Val = 4
// 	root.Left.Left.Left = new(TreeNode)
// 	root.Left.Left.Left.Val = 9
// 	root.Left.Right.Left = new(TreeNode)
// 	root.Left.Right.Left.Val = 1

// 	root.Right = new(TreeNode)
// 	root.Right.Val = 8
// 	root.Right.Left = new(TreeNode)
// 	root.Right.Left.Val = 1
// 	root.Right.Right = new(TreeNode)
// 	root.Right.Right.Val = 3
// 	root.Right.Right.Right = new(TreeNode)
// 	root.Right.Right.Right.Val = 5

// 	log.Println(sumEvenGrandparent(root))
// }
