package main

/*
- LeetCode - https://leetcode.com/problems/difference-between-ones-and-zeros-in-row-and-column
- Time - O(mn)
- Space - O(mn)
*/

func onesMinusZeros(grid [][]int) [][]int {

	rowLen, colLen := len(grid), len(grid[0])

	onesRow := make([]int, rowLen)
	onesCol := make([]int, colLen)

	out := make([][]int, rowLen)

	for i := range out {
		out[i] = make([]int, colLen)
	}

	for i, v := range grid {
		for j, v1 := range v {
			if v1 > 0 {
				onesRow[i]++
				onesCol[j]++
			}
		}
	}

	for i, v := range out {
		for j := range v {
			out[i][j] = onesRow[i] + onesCol[j] - (rowLen - onesRow[i]) - (colLen - onesCol[j])
		}
	}
	return out
}

// func main() {
// 	log.Println(onesMinusZeros([][]int{{0, 1, 1}, {1, 0, 1}, {0, 0, 1}}))
// }
