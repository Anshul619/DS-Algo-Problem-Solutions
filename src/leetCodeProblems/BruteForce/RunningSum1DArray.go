package main

/*
- LeetCode - https://leetcode.com/problems/running-sum-of-1d-array/
*/
import "log"

func runningSum(nums []int) []int {

	for index := range nums {
		if index > 0 {
			nums[index] += nums[index-1]
		}
	}

	return nums
}

func main() {

	input := []int{1, 2, 3, 4}

	log.Println(runningSum(input))
}
