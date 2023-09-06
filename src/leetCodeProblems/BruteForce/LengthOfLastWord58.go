package main

/*
- Leetcode - https://leetcode.com/problems/length-of-last-word
- Time - O(n)
- Space - O(1)
*/

func lengthOfLastWord(s string) int {
	out := 0
	wordLength := 0

	for _, v := range s {
		if string(v) != " " {
			wordLength++
		} else {
			if wordLength != 0 {
				out = wordLength
			}

			wordLength = 0
		}
	}

	if wordLength != 0 {
		out = wordLength
	}
	return out
}

// func main() {
// 	log.Println(lengthOfLastWord("Hello World"))
// 	log.Println(lengthOfLastWord("   fly me   to   the moon  "))
// 	log.Println(lengthOfLastWord("luffy is still joyboy"))
// }
