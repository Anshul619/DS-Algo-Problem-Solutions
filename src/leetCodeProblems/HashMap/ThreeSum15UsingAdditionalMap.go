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
	sort.Ints(nums)

	out := [][]int{}
	m := make(map[string]struct{})

	for i := 0; i < len(nums)-2; i++ {
		l := i + 1
		r := len(nums) - 1

		for l < r {
			s := nums[i] + nums[l] + nums[r]

			if s == 0 {

				p := strconv.Itoa(nums[i]) + strconv.Itoa(nums[l]) + strconv.Itoa(nums[r])

				if _, ok := m[p]; !ok {
					out = append(out, []int{nums[i], nums[l], nums[r]})
					m[p] = struct{}{}
				}

				l++
				r--
				continue
			}

			if s < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return out
}
