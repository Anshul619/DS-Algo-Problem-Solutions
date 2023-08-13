package main

/*
- Leetcode - https://leetcode.com/problems/roman-to-integer/
- Time - O(n)
- Space - O(1)
*/
func romanToInt(s string) int {
	out := 0

	for i := 0; i < len(s); i++ {
		switch string(s[i]) {
		case "M":
			out += 1000
		case "D":
			out += 500
		case "C":
			if i+1 < len(s) {
				switch string(s[i+1]) {
				case "D":
					out += 400
					i++
					continue
				case "M":
					out += 900
					i++
					continue
				}
			}
			out += 100
		case "L":
			out += 50
		case "X":
			if i+1 < len(s) {
				switch string(s[i+1]) {
				case "L":
					out += 40
					i++
					continue
				case "C":
					out += 90
					i++
					continue
				}
			}
			out += 10
		case "V":
			out += 5
		case "I":
			if i+1 < len(s) {
				switch string(s[i+1]) {
				case "V":
					out += 4
					i++
					continue
				case "X":
					out += 9
					i++
					continue
				}
			}
			out += 1
		}
	}
	return out
}

// func main() {
// 	log.Println(romanToInt("III"))
// 	log.Println(romanToInt("LVIII"))
// 	log.Println(romanToInt("MCMXCIV"))
// }
