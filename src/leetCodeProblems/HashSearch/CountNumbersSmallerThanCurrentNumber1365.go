package main

/*
	LeetCode - https://leetcode.com/problems/how-many-numbers-are-smaller-than-the-current-number/
*/

import (
	"log"
	"sort"
)

func smallerNumbersThanCurrent(nums []int) []int {

	sorted := make([]int, len(nums))
	copy(sorted, nums)
	sort.Ints(sorted)

	m := make(map[int]int)

	for i, val := range sorted {
		if _, ok := m[val]; !ok {
			m[val] = i
		}
	}

	for i, val := range nums {

		if mIndex, ok := m[val]; ok {
			nums[i] = mIndex
		}
	}

	return nums
}

func main() {

	input := []int{8, 1, 2, 2, 3}

	log.Println(smallerNumbersThanCurrent(input))
}
