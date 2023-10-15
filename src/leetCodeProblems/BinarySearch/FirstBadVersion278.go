package main

import "log"

/*
- LeetCode - https://leetcode.com/problems/first-bad-version/description/
- Time - O(logn)
- Space - O(1)
*/

func isBadVersion(version int) bool {
	return version == 4
}

func firstBadVersion(n int) int {
	start := 0
	end := n

	for start < end {
		mid := start + (end-start)/2

		if isBadVersion(mid) {
			end = mid
		} else {
			start = mid + 1
		}
	}

	return start
}

func main() {
	log.Println(firstBadVersion(5))
}
