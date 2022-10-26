package main

/*
- LeetCode - https://leetcode.com/problems/number-of-islands/
- TimeComplexity - O(m*n)
- SpaceComplexity - O(1)
*/
import (
	"bytes"
	"log"
)

func dfs(grid [][]byte, x int, y int) {

	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || !bytes.Equal([]byte{grid[x][y]}, []byte{'1'}) {
		return
	}

	grid[x][y] = 0

	dfs(grid, x-1, y)
	dfs(grid, x+1, y)
	dfs(grid, x, y-1)
	dfs(grid, x, y+1)
}

func numIslands(grid [][]byte) int {

	numberOfIslands := 0
	for i := range grid {
		for j := range grid[i] {
			if bytes.Equal([]byte{grid[i][j]}, []byte{'1'}) {
				numberOfIslands++
				dfs(grid, i, j)
			}
		}
	}

	return numberOfIslands
}

func main() {

	// grid := [][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'},
	// 	{'1', '1', '0', '0', '0'},
	// 	{'0', '0', '0', '0', '0'}}

	// log.Println(numIslands(grid))

	grid := [][]byte{{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'}}

	log.Println(numIslands(grid))
}
