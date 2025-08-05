package main

/*
- Leetcode - https://leetcode.com/problems/zigzag-conversion/description/
- Time - O(n)
- Space - O(n)
*/
func convert(s string, numRows int) string {
	// If only 1 row, just return original string
	if numRows < 2 {
		return s
	}

	// Create a slice of strings to hold characters for each row
	m := make([]string, numRows)
	goDown := false

	row := 0

	// Iterate over the characters of the input string
	for _, v := range s {
		if row == 0 || row == numRows-1 {
			goDown = !goDown // Change direction at top or bottom row
		}

		m[row] += string(v) // Append character to the current row

		// Update the row index depending on the direction
		if goDown {
			row++
		} else {
			row--
		}
	}

	var out string

	// Combine all rows into one string
	for _, v := range m {
		out += v
	}
	return out
}
