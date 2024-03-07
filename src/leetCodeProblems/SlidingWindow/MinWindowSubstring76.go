package main

/*
- LeetCode - https://leetcode.com/problems/minimum-window-substring/description
- Time - O(n)
- Space - O(n)
*/
import (
	"math"
)

func minWindow(s string, t string) string {
	mT := make(map[rune]int)
	mS := make(map[rune]int)

	for _, v := range t {
		mT[v]++
	}

	found := 0
	out := math.MaxInt
	start, startIndex := 0, -1

	for i, v := range s {
		mS[v]++

		if mS[v] <= mT[v] {
			found++
		}

		// all found
		if found == len(t) {
			for mS[rune(s[start])] > mT[rune(s[start])] || mT[rune(s[start])] == 0 {

				if mS[rune(s[start])] > mT[rune(s[start])] {
					mS[rune(s[start])]--
				}
				start++
			}
			if i-start+1 < out {
				out = i - start + 1
				startIndex = start
			}
		}
	}

	if startIndex == -1 {
		return ""
	}
	return s[startIndex : out+startIndex]
}
