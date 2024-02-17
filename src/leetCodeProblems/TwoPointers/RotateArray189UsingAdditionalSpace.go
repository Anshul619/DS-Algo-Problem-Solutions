package main

/*
- Leetcode - https://leetcode.com/problems/rotate-array/description/
- Space - O(k)
- Time - O(n)
*/
func rotateUsingSpace(nums []int, k int) {
	if k > len(nums) {
		k = k % len(nums)
	}

	out := make([]int, k)
	outIndex := 0

	for i := len(nums) - k; i < len(nums); i++ {
		out[outIndex] = nums[i]
		outIndex++
	}

	for i := len(nums) - 1; i >= k; i-- {
		nums[i] = nums[i-k]
	}

	for i := 0; i < k; i++ {
		nums[i] = out[i]
	}
}
