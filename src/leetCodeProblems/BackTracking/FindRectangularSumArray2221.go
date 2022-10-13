package main

/*
- LeetCode - https://leetcode.com/problems/find-triangular-sum-of-an-array/submissions/
*/
import "log"

func triangularSum(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}

	out := make([]int, len(nums)-1)

	for i := 0; i < len(nums)-1; i++ {
		out[i] = (nums[i] + nums[i+1]) % 10
	}

	return triangularSum(out)
}

func main() {

	nums := []int{1, 2, 3, 4, 5}

	log.Println(triangularSum(nums))
}
