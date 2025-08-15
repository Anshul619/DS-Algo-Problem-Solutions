package main

/*
- Leetcode - https://leetcode.com/problems/search-a-2d-matrix/description/
- Time - O(log(m*n))
- Space - O(1)
*/
func searchMatrix(matrix [][]int, target int) bool {
	numRows := len(matrix)
	numCols := len(matrix[0])

	start := 0
	end := numRows*numCols - 1

	for start <= end {
		mid := start + (end-start)/2

		// flattened index trick for mid
		midRow := mid / numCols // integer division — tells you how many full rows fit before this index
		midCol := mid % numCols // remainder — tells you the position within that row

		if target == matrix[midRow][midCol] {
			return true
		} else if target < matrix[midRow][midCol] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return false
}
