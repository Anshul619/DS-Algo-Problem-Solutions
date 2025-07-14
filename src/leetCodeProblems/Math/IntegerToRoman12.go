package main

/*
- Leetcode - https://leetcode.com/problems/integer-to-roman/description/
- Time - O(n)
- Space - O(1)
*/

func intToRoman(num int) string {

	out := ""

	m := map[int]string{
		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		5:    "V",
		4:    "IV",
		1:    "I",
	}

	fractions := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	for _, v := range fractions {
		c := num / v
		for j := 0; j < c; j++ {
			out += m[v]
		}
		num = num % v
	}
	return out
}
