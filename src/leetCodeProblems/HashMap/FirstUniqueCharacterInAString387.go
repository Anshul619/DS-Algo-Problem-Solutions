package main

/*
- LeetCode - https://leetcode.com/problems/first-unique-character-in-a-string/
- Time - O(n)
- Space - O(n)
*/

func firstUniqChar(s string) int {

	m := make(map[rune]int)

	for _, v := range s {
		m[v]++
	}

	for i, v := range s {
		if m[v] == 1 {
			return i
		}
	}
	return -1
}

// func main() {
// 	fmt.Println(firstUniqChar("leetcode"))
// 	fmt.Println(firstUniqChar("loveleetcode"))
// 	fmt.Println(firstUniqChar("aabb"))
// }
