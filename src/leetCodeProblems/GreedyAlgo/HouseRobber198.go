package main

/*
- LeetCode - https://leetcode.com/problems/house-robber/
*/
//import "log"

func max(num1 int, num2 int) int {

	if num1 > num2 {
		return num1
	}
	return num2
}

func rob(nums []int) int {

	length := len(nums)

	if length == 0 {
		return 0
	}

	value1 := nums[0]

	if length == 1 {
		return value1
	}

	value2 := max(value1, nums[1])

	if length == 2 {
		return value2
	}

	out := 0

	for i := 2; i < length; i++ {
		out = max(nums[i]+value1, value2)
		value1 = value2
		value2 = out
	}

	return out
}

// func main() {

// 	//input := []int{1, 2, 3, 1}
// 	//input := []int{2, 7, 9, 3, 1}

// 	input := []int{2, 1, 1, 2}

// 	log.Println(rob(input))
// }
