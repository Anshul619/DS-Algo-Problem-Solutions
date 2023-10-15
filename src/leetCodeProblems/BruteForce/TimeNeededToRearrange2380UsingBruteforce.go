package main

/*
- LeetCode - https://leetcode.com/problems/time-needed-to-rearrange-a-binary-string/
- Time - O(n)
- Space - O(1)
*/

func secondsToRemoveOccurrence1(s string) int {

	out := 0
	found := true

	for found {

		i := 0
		temp := ""
		found = false

		for i < len(s) {
			if string(s[i]) == "0" && i+1 < len(s) &&
				string(s[i+1]) == "1" {
				temp += "10"
				i = i + 2
				found = true
			} else {
				temp += string(s[i])
				i++
			}
		}
		s = temp
		if found {
			out++
		}
	}
	return out
}

// func main() {
// 	log.Println(secondsToRemoveOccurrences("0110101"))
// 	log.Println(secondsToRemoveOccurrences("11100"))
// 	log.Println(secondsToRemoveOccurrences("001011"))
// }
