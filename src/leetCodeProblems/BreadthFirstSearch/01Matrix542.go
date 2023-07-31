package main

import (
	"math"
)

/**
- LeetCode - https://leetcode.com/problems/01-matrix/
*/

type cell struct {
	row    int
	column int
}

type mQueue []cell

func (s *mQueue) pop() cell {
	c := (*s)[0]
	*s = (*s)[1:]
	return c
}

func (s *mQueue) push(c cell) {
	*s = append(*s, c)
}

func (s mQueue) isEmpty() bool {
	return len(s) == 0
}

func updateMatrix(mat [][]int) [][]int {

	out := make([][]int, len(mat))
	queue := new(mQueue)

	if len(mat) == 0 {
		return out
	}

	for i, _ := range out {
		out[i] = make([]int, len(mat[0]))

		for j, _ := range out[i] {
			out[i][j] = math.MaxInt

			if mat[i][j] == 0 {
				out[i][j] = 0
				queue.push(cell{i, j})
			}
		}
	}

	dx := []int{1, 0, -1, 0}
	dy := []int{0, 1, 0, -1}

	for !queue.isEmpty() {

		c := queue.pop()

		for k := 0; k < 4; k++ {

			x := c.row + dx[k]
			y := c.column + dy[k]

			if x >= 0 &&
				x < len(mat) &&
				y >= 0 &&
				y < len(mat[0]) {

				if out[x][y] > out[c.row][c.column]+1 {
					out[x][y] = out[c.row][c.column] + 1
					queue.push(cell{x, y})
				}
			}
		}
	}

	return out
}

// func main() {
// 	//mat := [][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}

// 	//mat := [][]int{{0, 1, 0, 1, 1}, {1, 1, 0, 0, 1}, {0, 0, 0, 1, 0}, {1, 0, 1, 1, 1}, {1, 0, 0, 0, 1}}
// 	//mat := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}

// 	//mat := [][]int{{1, 1, 0, 0, 1, 0, 0, 1, 1, 0}, {1, 0, 0, 1, 0, 1, 1, 1, 1, 1}, {1, 1, 1, 0, 0, 1, 1, 1, 1, 0}, {0, 1, 1, 1, 0, 1, 1, 1, 1, 1}, {0, 0, 1, 1, 1, 1, 1, 1, 1, 0}, {1, 1, 1, 1, 1, 1, 0, 1, 1, 1}, {0, 1, 1, 1, 1, 1, 1, 0, 0, 1}, {1, 1, 1, 1, 1, 0, 0, 1, 1, 1}, {0, 1, 0, 1, 1, 0, 1, 1, 1, 1}, {1, 1, 1, 0, 1, 0, 1, 1, 1, 1}}

// 	mat := [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {0, 0, 0}}
// 	log.Println(updateMatrix(mat))
// }
