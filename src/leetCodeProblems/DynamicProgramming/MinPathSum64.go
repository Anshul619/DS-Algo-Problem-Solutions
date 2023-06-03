package main

/*
- LeetCode - https://leetcode.com/problems/minimum-path-sum/submissions/960464393/
- Time - O(m*n)
- Space - O(1)
*/
import "log"

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func minPathSum(grid [][]int) int {

	for i, v := range grid {
		for j := range v {
			if i > 0 && j > 0 {
				grid[i][j] += min(grid[i-1][j], grid[i][j-1])
			} else if i > 0 {
				grid[i][0] += grid[i-1][0]
			} else if j > 0 {
				grid[0][j] += grid[0][j-1]
			}
		}
	}

	return grid[len(grid)-1][len(grid[0])-1]
}

func main() {
	log.Println(minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
	log.Println(minPathSum([][]int{{1, 2, 3}, {4, 5, 6}}))
}
