package main

import (
	"sort"
)

/*
- LeetCode - https://leetcode.com/problems/majority-element-ii/
- Time - O(nlogn)
- Space - O(1)
*/

func majorityElement(nums []int) []int {
	out := []int{}

	sort.Ints(nums)

	current, count := 0, 0

	for _, v := range nums {
		if v != current {

			if count > len(nums)/3 {
				out = append(out, current)
			}

			current, count = v, 0
		}

		count++
	}

	if count > len(nums)/3 {
		out = append(out, current)
	}
	return out
}

// func main() {
// 	log.Println(majorityElement([]int{3, 2, 3}))
// 	log.Println(majorityElement([]int{1}))
// 	log.Println(majorityElement([]int{1, 2}))
// }
