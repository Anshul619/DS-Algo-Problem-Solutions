package main

/*
- LeetCode - https://leetcode.com/problems/cousins-in-binary-tree/description/
- Time - O(n)
- Space - O(n)
*/

func findDepth(node *TreeNode, num int, depth int) (*TreeNode, int, bool) {
	if node == nil {
		return nil, -1, false
	}

	if (node.Left != nil && node.Left.Val == num) || (node.Right != nil && node.Right.Val == num) {
		return node, depth + 1, true
	}

	lParent, lDepth, lFound := findDepth(node.Left, num, depth+1)

	if lFound {
		return lParent, lDepth, true
	}

	rParent, rDepth, rFound := findDepth(node.Right, num, depth+1)

	if rFound {
		return rParent, rDepth, true
	}

	return nil, -1, false
}

func isCousins(root *TreeNode, x int, y int) bool {
	xParent, xDepth, _ := findDepth(root, x, 0)
	yParent, yDepth, _ := findDepth(root, y, 0)

	if xParent != yParent && xDepth == yDepth {
		return true
	}

	return false
}

// func main() {
// 	// root := new(TreeNode) //Returns pointer to TreeNode object
// 	// root.Val = 1
// 	// root.Left = new(TreeNode)
// 	// root.Left.Val = 2
// 	// root.Right = new(TreeNode)
// 	// root.Right.Val = 3
// 	// root.Left.Right = new(TreeNode)
// 	// root.Left.Right.Val = 4

// 	// log.Println(isCousins(root, 2, 3))

// 	root := new(TreeNode) //Returns pointer to TreeNode object
// 	root.Val = 1
// 	root.Left = new(TreeNode)
// 	root.Left.Val = 2
// 	root.Right = new(TreeNode)
// 	root.Right.Val = 3
// 	root.Left.Right = new(TreeNode)
// 	root.Left.Right.Val = 4
// 	root.Right.Right = new(TreeNode)
// 	root.Right.Right.Val = 5

// 	log.Println(isCousins(root, 5, 4))

// }
