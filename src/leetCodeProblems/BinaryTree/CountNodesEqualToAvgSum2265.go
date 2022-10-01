package main

/*
- LeetCode - https://leetcode.com/problems/count-nodes-equal-to-average-of-subtree/submissions/
*/
import "log"

func averageOfSubtreeUtil(node *TreeNode, out *int) (int, int) {

	if node == nil {
		return 0, 0
	}

	if node.Left == nil && node.Right == nil {
		*out++
		return node.Val, 1
	}

	leftSum, leftCount := averageOfSubtreeUtil(node.Left, out)
	rightSum, rightCount := averageOfSubtreeUtil(node.Right, out)

	// log.Println("node.Val ->", node.Val)

	// log.Println("leftSum -", leftSum)
	// log.Println("leftCount -", leftCount)

	// log.Println("rightSum -", rightSum)
	// log.Println("rightCount -", rightCount)

	sum := leftSum + rightSum + node.Val
	count := leftCount + rightCount + 1

	avg := sum / count

	// log.Println("Avg", avg)

	if node.Val == avg {
		// log.Println("INCREMENT")
		*out++
	}

	return sum, count
}

func averageOfSubtree(root *TreeNode) int {

	count := 0
	averageOfSubtreeUtil(root, &count)

	return count
}

func main() {
	// root := new(TreeNode) //Returns pointer to TreeNode object
	// root.Val = 4
	// root.Left = new(TreeNode)
	// root.Left.Val = 8
	// root.Left.Left = new(TreeNode)
	// root.Left.Left.Val = 0
	// root.Left.Right = new(TreeNode)
	// root.Left.Right.Val = 1

	// root.Right = new(TreeNode)
	// root.Right.Val = 5
	// root.Right.Right = new(TreeNode)
	// root.Right.Right.Val = 6

	root := new(TreeNode) //Returns pointer to TreeNode object
	root.Val = 1

	log.Println(averageOfSubtree(root))
}
