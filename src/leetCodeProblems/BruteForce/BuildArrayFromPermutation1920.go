package main

/**
- LeetCode Link - https://leetcode.com/problems/build-array-from-permutation/
- TimeComplexity - O(n)
- SpaceComplexity - O(n)
*/

//import "log"

func buildArray(nums []int) []int {

	output := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		output[i] = nums[nums[i]]
	}

	return output
}

// func main() {

// 	input := []int{0, 2, 1, 5, 3, 4}

// 	log.Println(buildArray(input)) //expected o/p = [0,1,2,4,5,3]
// }
