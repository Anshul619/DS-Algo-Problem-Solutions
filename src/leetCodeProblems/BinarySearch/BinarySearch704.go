package main

/*
- LeetCode - https://leetcode.com/problems/binary-search/
- Time - O(logn)
- Space - O(1)
*/

func searchUtil(start, end int, nums []int, target int) int {
	if start < 0 || end >= len(nums) || start > end {
		return -1
	}

	mid := start + (end-start)/2

	if nums[mid] == target {
		return mid
	}

	if nums[mid] > target {
		return searchUtil(start, mid-1, nums, target)
	}

	return searchUtil(mid+1, end, nums, target)
}
func search(nums []int, target int) int {
	return searchUtil(0, len(nums)-1, nums, target)
}

// func main() {
// 	log.Println(search([]int{-1, 0, 3, 5, 9, 12}, 9))
// 	log.Println(search([]int{-1, 0, 3, 5, 9, 12}, 2))
// }

// func main() {

// 	head := new(TreeNode)
// 	head.Val = 1
// 	head.Right = new(TreeNode)
// 	head.Right.Val = 2
// 	head.Right.Right = new(TreeNode)
// 	head.Right.Right.Val = 3
// 	head.Right.Right.Right = new(TreeNode)
// 	head.Right.Right.Right.Val = 4

// 	printTree(balanceBST(head))

// }
