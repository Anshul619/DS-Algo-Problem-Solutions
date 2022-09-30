package main

/*
- LeetCode - https://leetcode.com/problems/is-subsequence/submissions/
*/
import "log"

func isSubsequence(s string, t string) bool {

	sIndex := 0

	if len(s) == 0 {
		return true
	}

	for i, _ := range t {

		if s[sIndex] == t[i] {
			sIndex++
		}

		if sIndex == len(s) {
			return true
		}
	}

	return false
}

func main() {
	//s := "abc"
	//t := "ahbgdc"

	// s := "axc"
	// t := "ahbgdc"

	s := "ahb"
	t := "ahbgdc"

	log.Println(isSubsequence(s, t))
}
