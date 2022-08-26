package main

/*
- LeetCode - https://leetcode.com/problems/rotting-oranges/
*/
import "log"

func orangesRotting(grid [][]int) int {

	if len(grid) == 0 {
		return 0
	}

	dx, dy, out, changed := []int{0, 1, 0, -1}, []int{1, 0, -1, 0}, 2, false

	for true {
		for i := range grid {
			for j := range grid[i] {

				if grid[i][j] == out {

					for k := 0; k < 4; k++ {

						x := i + dx[k]
						y := j + dy[k]

						if x > -1 && x < len(grid) && y > -1 && y < len(grid[0]) && grid[x][y] == 1 {
							grid[x][y] = grid[i][j] + 1
							//changed = true
							changed = true
						}
					}
				}
			}
		}

		if !changed {
			break
		}

		out++
		changed = false
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}

	return out - 2
}

func main() {

	input := [][]int{{2, 0, 1, 1, 1, 1, 1, 1, 1, 1}, {1, 0, 1, 0, 0, 0, 0, 0, 0, 1}, {1, 0, 1, 0, 1, 1, 1, 1, 0, 1}, {1, 0, 1, 0, 1, 0, 0, 1, 0, 1}, {1, 0, 1, 0, 1, 0, 0, 1, 0, 1}, {1, 0, 1, 0, 1, 1, 0, 1, 0, 1}, {1, 0, 1, 0, 0, 0, 0, 1, 0, 1}, {1, 0, 1, 1, 1, 1, 1, 1, 0, 1}, {1, 0, 0, 0, 0, 0, 0, 0, 0, 1}, {1, 1, 1, 1, 1, 1, 1, 1, 1, 1}}

	//input := [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}} // expected o/p = 4

	//input := [][]int{{2, 1, 1}, {1, 1, 1}, {0, 1, 2}} // expected o/p = 2

	//input := [][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}

	//input := [][]int{{2}, {1}}
	log.Println(orangesRotting(input))
}
