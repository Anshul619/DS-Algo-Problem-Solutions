package main

/*
- Leetcode - https://leetcode.com/problems/search-a-2d-matrix/description/
- Time - O(log(m*n))
- Space - O(1)
*/
func searchMatrix(matrix [][]int, target int) bool {
	row, col := len(matrix), len(matrix[0])

	// consider matrix to be 1 array for binary search
	start, end := 0, row*col-1

	for start <= end {
		mid := start + (end-start)/2

		tc := mid % col
		tr := mid / col

		if matrix[tr][tc] == target {
			return true
		}

		if matrix[tr][tc] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return false
}
