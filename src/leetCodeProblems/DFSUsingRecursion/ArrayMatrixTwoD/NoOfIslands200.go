package main

/*
- LeetCode - https://leetcode.com/problems/number-of-islands/
- TimeComplexity - O(m*n)
- SpaceComplexity - O(1)
*/
import (
	"bytes"
)

func dfs(grid [][]byte, x, y int) {
	if x < 0 || y < 0 ||
		x >= len(grid) || y >= len(grid[0]) ||
		!bytes.Equal([]byte{grid[x][y]}, []byte{'1'}) {
		return
	}

	grid[x][y] = 0
	dfs(grid, x+1, y)
	dfs(grid, x-1, y)
	dfs(grid, x, y+1)
	dfs(grid, x, y-1)
}
func numIslands(grid [][]byte) int {
	out := 0

	for i, v := range grid {
		for j, v1 := range v {
			if bytes.Equal([]byte{v1}, []byte{'1'}) {
				out++
				dfs(grid, i, j)
			}
		}
	}
	return out
}
