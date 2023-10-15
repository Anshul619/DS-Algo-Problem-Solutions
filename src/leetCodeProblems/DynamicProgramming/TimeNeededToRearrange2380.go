package main

import "log"

/*
- LeetCode - https://leetcode.com/problems/time-needed-to-rearrange-a-binary-string/
- Time - O(n)
- Space - O(1)
*/

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func secondsToRemoveOccurrences(s string) int {

	out := 0
	zeros := 0

	for _, v := range s {
		if string(v) == "0" {
			zeros++
		} else if zeros > 0 {
			out = maxInt(out+1, zeros)
		}
	}

	return out
}

func main() {
	log.Println(secondsToRemoveOccurrences("0110101"))
	log.Println(secondsToRemoveOccurrences("11100"))
	log.Println(secondsToRemoveOccurrences("001011"))
}
