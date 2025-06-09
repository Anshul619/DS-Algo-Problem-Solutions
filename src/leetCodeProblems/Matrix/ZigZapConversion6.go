package main

/*
- Leetcode - https://leetcode.com/problems/zigzag-conversion/description/
- Time - O(n * numRows)
- Space - O(n * numRows)
*/
func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}

	// temporary matrix
	matrix := [][]byte{}

	matrixIndex, sIndex := 0, 0

	for sIndex < len(s) {
		t := make([]byte, numRows)

		if matrixIndex%(numRows-1) == 0 {
			for k := 0; k < numRows && sIndex < len(s); k++ {
				t[k] = s[sIndex]
				sIndex++
			}

		} else {
			t[numRows-matrixIndex%(numRows-1)-1] = s[sIndex]
			sIndex++
		}

		matrix = append(matrix, t)
		matrixIndex++
	}

	var out string

	for i := 0; i < numRows; i++ {
		for j := 0; j < len(matrix); j++ {
			if matrix[j][i] != 0 {
				out += string(matrix[j][i])
			}
		}
	}

	return out
}
