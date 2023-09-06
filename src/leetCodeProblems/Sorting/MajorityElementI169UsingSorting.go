package main

import (
	"sort"
)

/*
- LeetCode - https://leetcode.com/problems/majority-element/
- Time - O(nlogn)
- Space - O(1)
*/

func majorityElement1(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

// func main() {
// 	log.Println(majorityElement1([]int{3, 2, 3}))
// 	log.Println(majorityElement1([]int{2, 2, 1, 1, 1, 2, 2}))
// }
