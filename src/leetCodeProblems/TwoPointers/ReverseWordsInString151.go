package main

/*
- Leetcode - https://leetcode.com/problems/reverse-words-in-a-string/description/
- Time - O(n)
- Space - O(n)
*/

func reverseWords(s string) string {

	out := []rune{}

	lastIndex := -1

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			if lastIndex == -1 {
				lastIndex = i
			}
		} else {
			if lastIndex != -1 {

				if len(out) > 0 {
					out = append(out, ' ')
				}
				out = append(out, []rune(s[i+1:lastIndex+1])...)
				lastIndex = -1
			}
		}
	}

	if lastIndex != -1 {
		if len(out) > 0 {
			out = append(out, ' ')
		}
		out = append(out, []rune(s[0:lastIndex+1])...)
	}

	return string(out)
}

// func main() {
// 	log.Println(reverseWords("the sky is blue"))
// 	log.Println(reverseWords("  hello world  "))
// 	log.Println(reverseWords("a good   example"))
// }
