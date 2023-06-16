package main

/*
- LeetCode - https://leetcode.com/problems/plus-one/description/
- Time - O(n)
- Space - O(1)
*/
import "log"

func plusOne(digits []int) []int {

	for i := len(digits) - 1; i >= 0; i-- {

		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}

	return append([]int{1}, digits...)
}

func main() {
	log.Println(plusOne([]int{1, 2, 3}))
	log.Println(plusOne([]int{4, 3, 2, 1}))
	log.Println(plusOne([]int{9}))
}
