package main

/*
- LeetCode - https://leetcode.com/problems/roman-to-integer/
- Time - O(n)
- Space - O(1)
*/
func romanToInt(s string) int {
	out := 0

	for i := 0; i < len(s); i++ {
		if i+1 < len(s) {
			t := string(s[i]) + string(s[i+1])
			n := 0

			switch t {
			case "IV":
				n = 4
			case "IX":
				n = 9
			case "XL":
				n = 40
			case "XC":
				n = 90
			case "CD":
				n = 400
			case "CM":
				n = 900
			}

			if n != 0 {
				out += n
				i++
				continue
			}
		}

		switch string(s[i]) {
		case "I":
			out += 1
		case "V":
			out += 5
		case "X":
			out += 10
		case "L":
			out += 50
		case "C":
			out += 100
		case "D":
			out += 500
		case "M":
			out += 1000
		}
	}
	return out
}
