package main

/*
- Leetcode - https://leetcode.com/problems/subarrays-with-k-different-integers/
- Time - O(n)
- Space - O(n)
*/

// subarraysWithKDistinct returns the number of subarrays with exactly k distinct integers.
// It leverages the fact that:
//
//	Number of subarrays with exactly K distinct elements =
//	    Number of subarrays with at most K distinct elements -
//	    Number of subarrays with at most (K - 1) distinct elements.
func subarraysWithKDistinct(nums []int, k int) int {
	return atMostK(nums, k) - atMostK(nums, k-1)
}

// atMostK returns the number of subarrays with at most k distinct integers.
// It uses the sliding window technique along with a frequency map to count valid subarrays.
func atMostK(nums []int, k int) int {
	count := 0                // Total number of valid subarrays
	left := 0                 // Left boundary of the sliding window
	freq := make(map[int]int) // Frequency map to track elements in the current window

	for right := 0; right < len(nums); right++ {
		// If nums[right] is new to the window, decrease k
		if freq[nums[right]] == 0 {
			k--
		}
		freq[nums[right]]++

		// Shrink the window until there are at most k distinct numbers
		for k < 0 {
			freq[nums[left]]--
			if freq[nums[left]] == 0 {
				k++
			}
			left++
		}

		// Add the number of subarrays ending at 'right'
		count += right - left + 1
	}

	return count
}
