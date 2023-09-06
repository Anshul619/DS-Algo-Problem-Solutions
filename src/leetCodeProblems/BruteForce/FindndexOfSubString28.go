package main

/*
- LeetCode - https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/
- Time - O(n*m)
- Space - O(1)
*/

func strStr(haystack string, needle string) int {
	for i := range haystack {

		j := i
		for _, v := range needle {
			if rune(haystack[j]) != v {
				break
			}
			j++
		}
		if j-i == len(needle) {
			return i
		}
	}

	return -1
}

// func main() {
// 	//s := "abc"
// 	//t := "ahbgdc"

// 	// s := "axc"
// 	// t := "ahbgdc"

// 	s := "sadbutsad"
// 	t := "sad"

// 	log.Println(strStr(s, t))

// 	s = "leetcode"
// 	t = "leeto"

// 	log.Println(strStr(s, t))
// }
