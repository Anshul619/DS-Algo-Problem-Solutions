package main

/*
- LeetCode - https://leetcode.com/problems/valid-anagram/description/
- Time complexity: O(n)
- Space complexity: O(26)
*/
import (
	"log"
)

func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	a := make([]int, 26)

	for _, v := range s {
		a[v-'a']++
	}

	for _, v := range t {
		a[v-'a']--
	}

	for _, v := range a {
		if v != 0 {
			return false
		}
	}

	return true
}

func main() {
	s := "anagram"
	t := "nagaram"

	// s := "rat"
	// t := "car"

	log.Println(isAnagram(s, t))
}
