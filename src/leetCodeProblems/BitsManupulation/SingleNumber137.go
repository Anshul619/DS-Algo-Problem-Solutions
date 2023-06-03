package main

/*
- LeetCode - https://leetcode.com/problems/single-number-ii
- Time - O(n)
- Space - O(1)
*/
import "log"

func singleNumber(nums []int) int {
	low, high := 0, 0

	for _, v := range nums {
		low = (v ^ low) & ^high
		high = (v ^ high) & ^low
	}
	return low
}

func main() {
	log.Println(singleNumber([]int{2, 2, 3, 2}))
	log.Println(singleNumber([]int{0, 1, 0, 1, 0, 1, 99}))
}
