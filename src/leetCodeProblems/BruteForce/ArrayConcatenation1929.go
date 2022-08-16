package main

/**
- LeetCode - https://leetcode.com/problems/concatenation-of-array/submissions/
- TimeComplexity - O(n)
- SpaceComplexity - O(n)
*/
import "log"

func getConcatenation(nums []int) []int {

	output := make([]int, len(nums)*2) //Initialize slice instead of array, since length is non-constant

	for i := 0; i < len(nums); i++ {
		output[i] = nums[i]
		output[i+len(nums)] = nums[i]
	}

	return output
}

func main() {

	//input := []int{1, 2, 1}
	//log.Println(getConcatenation(input))

	input := []int{1, 3, 2, 1}
	log.Println(getConcatenation(input))
}
