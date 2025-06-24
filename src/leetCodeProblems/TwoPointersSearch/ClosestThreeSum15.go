package main

import (
	"log"
	"math"
	"sort"
)

/*
- LeetCode - https://leetcode.com/problems/3sum-closest/description/
- Time - O(n^2)
- Space - O(1)
*/

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	closest := math.MaxInt

	log.Println(nums)
	for i, v := range nums {

		left := i + 1
		right := len(nums) - 1

		for left < right {
			sum := v + nums[left] + nums[right]
			if abs(sum-target) < abs(closest-target) {
				closest = sum
			}
			if sum < target {
				left++
			} else {
				right--
			}
		}
	}
	return closest
}
