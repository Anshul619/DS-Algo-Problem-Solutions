package main

import "log"

/*
- LeetCode - https://leetcode.com/problems/optimal-partition-of-string/description/
- Time - O(n)
- Space - O(1)
*/
func partitionString(s string) int {

	characters := make([]int, 26)
	out := 0

	for i, v := range s {
		if characters[v-rune('a')] != 0 {
			out++
			characters = make([]int, 26)
			characters[v-rune('a')] = 1
		} else {
			characters[v-rune('a')] = 1
		}

		if i == len(s)-1 {
			out++
		}
	}

	return out
}

func main() {

	//s := "abacaba"

	s := "ssssss"
	log.Println(partitionString(s))
}
