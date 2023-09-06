package main

import "log"

/*
- LeetCode - https://leetcode.com/problems/find-n-unique-integers-sum-up-to-zero/
- Time - O(n)
- Space - O(1)
*/

func sumZero(n int) []int {
	out := []int{}

	if n%2 != 0 {
		out = append(out, 0)
	}

	for i := 1; i <= n/2; i++ {
		out = append(out, i)
		out = append(out, -i)
	}
	return out
}

func main() {
	log.Println(sumZero(5))
	log.Println(sumZero(3))
	log.Println(sumZero(1))
	log.Println(sumZero(4))
}
