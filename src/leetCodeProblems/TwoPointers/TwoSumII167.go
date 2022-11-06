package main

/*
- LeetCode - https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
- TimeComplexity - O(n)
- Space - O(1)
*/
import "log"

func twoSum(numbers []int, target int) []int {

	out := []int{}

	left := 0
	right := len(numbers) - 1

	for left < right {

		currentSum := numbers[left] + numbers[right]

		if currentSum == target {
			out = append(out, left+1)
			out = append(out, right+1)
			break
		} else if currentSum < target {
			left++
		} else {
			right--
		}
	}

	return out
}

func main() {
	// numbers := []int{2, 7, 11, 15}
	// target := 9

	// numbers := []int{2, 3, 4}
	// target := 6

	numbers := []int{-1, 0}
	target := -1

	log.Println(twoSum(numbers, target))

}
