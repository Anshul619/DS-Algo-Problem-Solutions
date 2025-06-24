package main

/*
- LeetCode - https://leetcode.com/problems/3sum/
- Time - O(n^2)
- Space - O(n)
*/
import (
	"sort"
	"strconv"
)

func threeSumHashMap(nums []int) [][]int {

	targetSum := 0
	uniqueOutMap := make(map[string]bool)

	var out [][]int

	sort.Ints(nums)

	for i, v := range nums {
		left := i + 1
		right := len(nums) - 1

		for left < right {

			sum := v + nums[left] + nums[right]

			if sum == targetSum {
				pairKey := strconv.Itoa(nums[left]) + "_" + strconv.Itoa(nums[right]) + "_" + strconv.Itoa(v)
				if _, uOk := uniqueOutMap[pairKey]; !uOk {
					out = append(out, []int{v, nums[left], nums[right]})
					uniqueOutMap[pairKey] = true
				}
			}

			if sum > targetSum {
				right--
			} else {
				left++
			}
		}
	}
	return out
}
