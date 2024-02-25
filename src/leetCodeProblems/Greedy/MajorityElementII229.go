package main

/*
- LeetCode - https://leetcode.com/problems/majority-element-ii/description/
- Time - O(n)
- Space - O(1)
*/

func findCandidates(nums []int) []int {
	fIndex, fCount := -1, 0
	sIndex, sCount := -1, 0

	for i := 0; i < len(nums); i++ {
		if fIndex == -1 || nums[i] == nums[fIndex] {
			fIndex = i
			fCount++
		} else if sIndex == -1 || nums[i] == nums[sIndex] {
			sIndex = i
			sCount++
		} else if sCount == 0 {
			sIndex, sCount = i, 1
		} else if fCount == 0 {
			fIndex, fCount = i, 1
		} else {
			fCount--
			sCount--
		}
	}

	out := []int{}

	if fIndex != -1 {
		out = append(out, nums[fIndex])
	}

	if sIndex != -1 {
		out = append(out, nums[sIndex])
	}
	return out
}

func isMajority(nums []int, n int) bool {
	count := 0

	for _, v := range nums {
		if v == n {
			count++
		}
	}

	return count > len(nums)/3
}

func majorityElement(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	if len(nums) == 1 {
		return nums
	}

	out := []int{}

	c := findCandidates(nums)

	for _, v := range c {
		if isMajority(nums, v) {
			out = append(out, v)
		}
	}

	return out
}
