package main

/*
- LeetCode - https://leetcode.com/problems/search-insert-position/description/
- Time - O(logn)
- Space - O(1)
*/
import "log"

func binarySearchUtil(nums []int, start int, end int, target int) int {

	if start > end || end >= len(nums) {
		return end + 1
	}

	if start < 0 {
		return 0
	}

	mid := start + (end-start)/2

	if nums[mid] == target {
		return mid
	}

	if nums[mid] < target {
		return binarySearchUtil(nums, mid+1, end, target)
	} else {
		return binarySearchUtil(nums, start, mid-1, target)
	}
}

func searchInsert(nums []int, target int) int {
	return binarySearchUtil(nums, 0, len(nums)-1, target)
}

func main() {
	log.Println(searchInsert([]int{1, 3, 5, 6}, 5))
	log.Println(searchInsert([]int{1, 3, 5, 6}, 2))
	log.Println(searchInsert([]int{1, 3, 5, 6}, 7))
}
