package main

/*
- LeetCode - https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/submissions/
*/
import "log"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func printInOrder(root *TreeNode) {

	if root == nil {
		return
	}

	printInOrder(root.Left)
	log.Println(root.Val)
	printInOrder(root.Right)

}
func sortedArrayToBSTUtil(nums []int, start int, end int) *TreeNode {

	mid := start + (end-start)/2

	if start > end || start < 0 || end > len(nums) || mid >= len(nums) {
		return nil
	}

	node := new(TreeNode)
	node.Val = nums[mid]
	node.Left = sortedArrayToBSTUtil(nums, start, mid-1)
	node.Right = sortedArrayToBSTUtil(nums, mid+1, end)

	return node
}

func sortedArrayToBST(nums []int) *TreeNode {
	return sortedArrayToBSTUtil(nums, 0, len(nums))
}

// func main() {

// 	input := []int{-10, -3, 0, 5, 9}
// 	printInOrder(sortedArrayToBST(input))
// }
