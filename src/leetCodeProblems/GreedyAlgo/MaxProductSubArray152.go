package main

/*
- LeetCode - https://leetcode.com/problems/maximum-product-subarray
*/
import (
	"log"
)

func min(n1 int, n2 int) int {
	if n1 > n2 {
		return n2
	}
	return n1
}

func max(n1 int, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

func maxProduct(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	max_so_far := nums[0]
	max_ending := nums[0]
	min_ending := nums[0]

	for i := 1; i < len(nums); i++ {

		temp_min := min(min(nums[i], nums[i]*min_ending), nums[i]*max_ending)  // get min of all possible three
		max_ending = max(max(nums[i], nums[i]*max_ending), nums[i]*min_ending) // get max of all possible three
		min_ending = temp_min
		max_so_far = max(max_so_far, max_ending)
	}

	return max_so_far
}

func main() {
	nums := []int{2, -5, -2, -4, 3}

	//nums := []int{3, -1, 4}
	//nums := []int{-2, 0, -1}
	log.Println(maxProduct(nums))
}
