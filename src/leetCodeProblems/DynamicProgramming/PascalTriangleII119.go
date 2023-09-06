package main

import "log"

/*
- LeetCode - https://leetcode.com/problems/pascals-triangle-ii/
- Time - O(n)
- Space - O(n)
*/

func getRow(rowIndex int) []int {
	rowIndex++
	out := make([]int, rowIndex)

	c := 1

	for i := 1; i <= rowIndex; i++ {
		out[i-1] = c
		c = c * (rowIndex - i) / i
	}
	return out
}

func main() {
	log.Println(getRow(3))
	log.Println(getRow(0))
	log.Println(getRow(5))
	log.Println(getRow(1))
}
