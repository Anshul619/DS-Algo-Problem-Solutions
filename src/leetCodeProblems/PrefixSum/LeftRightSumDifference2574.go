package main

/*
- LeetCode - https://leetcode.com/problems/left-and-right-sum-differences/description/
- Time - O(n)
- Extra Space - O(1)
*/

func diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func leftRightDifference(nums []int) []int {
	totalSum := 0

	for _, v := range nums {
		totalSum += v
	}

	out := make([]int, len(nums))
	leftSum := 0

	for i, v := range nums {
		rightSum := totalSum - v - leftSum

		out[i] = diff(leftSum, rightSum)
		leftSum += v
	}

	return out
}

// func main() {
// 	log.Println(leftRightDifference([]int{10, 4, 8, 3}))
// 	log.Println(leftRightDifference([]int{1}))
// }
