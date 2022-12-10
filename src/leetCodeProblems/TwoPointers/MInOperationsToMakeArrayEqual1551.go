package main

/*
- LeetCode - https://leetcode.com/problems/minimum-operations-to-make-array-equal/description/
- Time - O(1)
- Space - O(1)
*/
import "log"

func minOperations(n int) int {
	return (n * n) / 4
}

func main() {
	//log.Println(minOperations(3))
	log.Println(minOperations(6))
	//log.Println(minOperations(31))
}
