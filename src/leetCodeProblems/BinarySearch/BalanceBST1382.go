package main

/*
- LeetCode - https://leetcode.com/problems/balance-a-binary-search-tree/
- Time - O(n)
- Space - O(n)
*/
import "log"

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func bstToSortedArray(node *TreeNode, sortedArray *[]int) {

	if node == nil {
		return
	}

	bstToSortedArray(node.Left, sortedArray)

	*sortedArray = append(*sortedArray, node.Val)

	bstToSortedArray(node.Right, sortedArray)
}

func sortedArrayToBST1(sortedArray []int, low int, high int) *TreeNode {

	if low < 0 || low >= len(sortedArray) || high < 0 || high >= len(sortedArray) || low > high {
		return nil
	}

	mid := low + (high-low)/2

	node := new(TreeNode)
	node.Val = sortedArray[mid]

	node.Left = sortedArrayToBST1(sortedArray, low, mid-1)
	node.Right = sortedArrayToBST1(sortedArray, mid+1, high)

	return node
}

func balanceBST(root *TreeNode) *TreeNode {

	sortedArray := make([]int, 0)

	bstToSortedArray(root, &sortedArray)

	return sortedArrayToBST1(sortedArray, 0, len(sortedArray)-1)
}

func printTree(node *TreeNode) {

	if node == nil {
		return
	}

	printTree(node.Left)
	log.Println(node.Val)
	printTree(node.Right)
}

func main() {

	head := new(TreeNode)
	head.Val = 1
	head.Right = new(TreeNode)
	head.Right.Val = 2
	head.Right.Right = new(TreeNode)
	head.Right.Right.Val = 3
	head.Right.Right.Right = new(TreeNode)
	head.Right.Right.Right.Val = 4

	printTree(balanceBST(head))

}
