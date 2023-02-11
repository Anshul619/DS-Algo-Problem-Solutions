package main

/*
- LeetCode - https://leetcode.com/problems/minimize-maximum-pair-sum-in-array/description/
- Time - O(nlogn)
- Space - O(1)
*/
import (
	"math"
	"sort"
)

func minPairSum(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	maxSum := math.MinInt
	first, last := 0, len(nums)-1

	sort.Ints(nums)

	for first < last {

		if nums[first]+nums[last] > maxSum {
			maxSum = nums[first] + nums[last]
		}
		first++
		last--
	}

	return maxSum
}

// func main() {
// 	//nums := []int{3, 5, 2, 3}
// 	//nums := []int{3, 5, 4, 2, 4, 6}
// 	nums := []int{3, 2, 4, 1, 1, 5, 1, 3, 5, 1}
// 	log.Println(minPairSum(nums))
// }
