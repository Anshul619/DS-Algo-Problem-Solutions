package main

/*
- LeetCode - https://leetcode.com/problems/excel-sheet-column-number/description/
- Time - O(n)
- Space - O(1)
*/
import (
	"math"
)

func titleToNumber(columnTitle string) int {

	out := 0

	for i, v := range columnTitle {
		if len(columnTitle)-i != 1 {
			out += int(math.Pow(26.0, float64(len(columnTitle)-i-1))) * int(v-64)
		} else {
			out += int(v - 64)
		}
	}

	return out
}

// func main() {

// 	log.Println(titleToNumber("A"))
// 	log.Println(titleToNumber("AB"))
// 	log.Println(titleToNumber("ZY"))
// 	log.Println(titleToNumber("CDA"))
// }
