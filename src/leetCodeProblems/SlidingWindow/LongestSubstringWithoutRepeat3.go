package main

/*
- Leet Code - https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
- Space - O(n)
- Time - O(n)
*/
import "log"

func lengthOfLongestSubstring(s string) int {
	m := make(map[rune]int)
	out := 0
	startIndex := 0

	for i, v := range s {
		if k, ok := m[v]; ok && k >= startIndex {
			if (i - startIndex) > out {
				out = i - startIndex
			}

			startIndex = k + 1
		}
		m[v] = i
	}

	if (len(s) - startIndex) > out {
		out = len(s) - startIndex
	}

	return out
}

func main() {
	//log.Println(lengthOfLongestSubstring("abcabcbb"))
	// log.Println(lengthOfLongestSubstring("bbbbb"))
	//log.Println(lengthOfLongestSubstring("abba"))

	log.Println(lengthOfLongestSubstring(" "))
}
