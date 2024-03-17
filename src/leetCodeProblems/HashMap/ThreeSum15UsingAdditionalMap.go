package main

/*
- LeetCode - https://leetcode.com/problems/3sum/
- Time - O(n^2)
- Space - O(n)
*/
import (
	//"log"

	"sort"
	"strconv"
)

func threeSum(nums []int) [][]int {

	hashMap := make(map[int]int)
	targetSum := 0
	uniqueOutMap := make(map[string]bool)

	var out [][]int

	sort.Ints(nums)

	for i, v := range nums {
		hashMap[v] = i
	}

	for i, v := range nums {

		for j := i + 1; j < len(nums); j++ {

			sum := targetSum - v - nums[j]

			if mapV, ok := hashMap[sum]; ok && mapV > i && mapV > j {

				pairKey := strconv.Itoa(v) + "_" + strconv.Itoa(nums[j]) + "_" + strconv.Itoa(sum)
				temp := append([]int{}, v, nums[j], sum)

				if _, uOk := uniqueOutMap[pairKey]; !uOk {
					out = append(out, temp)
					uniqueOutMap[pairKey] = true
				}
			}
		}
	}

	return out
}
