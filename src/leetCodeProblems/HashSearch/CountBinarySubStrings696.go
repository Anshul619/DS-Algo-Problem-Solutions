package main

/*
- LeetCode - https://leetcode.com/problems/count-binary-substrings/
*/

func countBinarySubstrings(s string) int {

	out := 0
	zeroCount := 0
	oneCount := 0

	i := 0

	for i < len(s) {

		oneCount, zeroCount = 0, 0
		if s[i] == '0' {

			for i < len(s) && s[i] == '0' {
				zeroCount++
				i++
			}

			j := i

			for j < len(s) && s[j] == '1' {
				oneCount++
				j++
			}

		} else {

			for i < len(s) && s[i] == '1' {
				oneCount++
				i++
			}

			j := i

			for j < len(s) && s[j] == '0' {
				zeroCount++
				j++
			}
		}

		if oneCount < zeroCount {
			out += oneCount
		} else {
			out += zeroCount
		}
	}

	return out
}

// func main() {

// 	//s := "00110011"
// 	s := "00100"
// 	log.Println(countBinarySubstrings(s))
// }
