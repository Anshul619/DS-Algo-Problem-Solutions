package main

/*
- LeetCode - https://leetcode.com/problems/roman-to-integer/
- Time - O(n)
- Space - O(1)
*/
func romanToInt(s string) int {
	out := 0

	m := map[string]int{
		"M":  1000,
		"CM": 900,
		"D":  500,
		"CD": 400,
		"C":  100,
		"XC": 90,
		"L":  50,
		"XL": 40,
		"X":  10,
		"IX": 9,
		"V":  5,
		"IV": 4,
		"I":  1,
	}

	for i := 0; i < len(s); i++ {
		// compare pair
		if i+1 < len(s) {
			p := s[i : i+2]

			if v, ok := m[p]; ok {
				out += v
				i++
				continue
			}
		}

		out += m[string(s[i])]
	}
	return out
}
