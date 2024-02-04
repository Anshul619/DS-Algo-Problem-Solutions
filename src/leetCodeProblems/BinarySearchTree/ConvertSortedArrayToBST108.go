package main

/*
- LeetCode - https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/submissions/
- Time - O(nlogn)
- Space - O(1)
*/

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	if len(nums) == 1 {
		return &TreeNode{nums[0], nil, nil}
	}

	mid := len(nums) / 2
	node := new(TreeNode)
	node.Val = nums[mid]
	node.Left = sortedArrayToBST(nums[:mid])
	node.Right = sortedArrayToBST(nums[mid+1:])
	return node
}
