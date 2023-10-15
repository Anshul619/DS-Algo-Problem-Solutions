package main

/*
- LeetCode - https://leetcode.com/problems/4sum/description/
- Time - O(n^3)
- Space - O(n)
*/
import (
	"sort"
	"strconv"
)

func fourSum(nums []int, target int) [][]int {

	uniqueOutMap := make(map[string]bool)

	var out [][]int

	sort.Ints(nums)

	for i, v := range nums {

		for j := i + 1; j < len(nums); j++ {

			left := j + 1
			right := len(nums) - 1

			for left < right {
				sum := v + nums[j] + nums[left] + nums[right]

				if sum == target {
					pairKey := strconv.Itoa(v) + "_" + strconv.Itoa(nums[j]) + "_" + strconv.Itoa(nums[left]) + "_" + strconv.Itoa(nums[right])

					if _, uOk := uniqueOutMap[pairKey]; !uOk {
						out = append(out, append([]int{}, v, nums[j], nums[left], nums[right]))
						uniqueOutMap[pairKey] = true
					}
				}

				if sum < target {
					left++
				} else {
					right--
				}
			}

		}
	}

	return out
}

// func main() {
// 	log.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
// }
