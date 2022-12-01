package main

/*
- LeetCode - https://leetcode.com/problems/number-of-closed-islands/
*/

type cell struct {
	row int
	col int
}

type queue []cell

func (q *queue) push(c cell) {
	*q = append(*q, c)
}

func (q *queue) pop() cell {
	c := (*q)[0]
	*q = (*q)[1:]
	return c
}

func (q queue) isEmpty() bool {
	return len(q) == 0
}

func bfs(row int, col int, grid [][]int, visited [][]bool) bool {

	q := new(queue)

	q.push(cell{row, col})

	dx := []int{1, 0, -1, 0}
	dy := []int{0, 1, 0, -1}

	isClosed := true

	for !q.isEmpty() {

		c := q.pop()

		row = c.row
		col = c.col

		if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) {
			isClosed = false
			continue // this is important and tricky
		}

		if grid[row][col] == 1 || visited[row][col] {
			continue
		}

		visited[row][col] = true

		for i := 0; i < 4; i++ {
			q.push(cell{row + dx[i], col + dy[i]})
		}

		//log.Println(visited)
	}

	return isClosed
}

func closedIsland(grid [][]int) int {

	visited := make([][]bool, len(grid))

	for i := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}

	closedIslandCount := 0

	for i := 0; i < len(grid); i++ {

		for j := 0; j < len(grid[0]); j++ {

			if grid[i][j] == 0 && !visited[i][j] {

				if bfs(i, j, grid, visited) {
					closedIslandCount++
				}
			}
		}
	}

	return closedIslandCount
}

// func main() {
// 	//grid := [][]int{{1, 1, 1, 1, 1, 1, 1, 0}, {1, 0, 0, 0, 0, 1, 1, 0}, {1, 0, 1, 0, 1, 1, 1, 0}, {1, 0, 0, 0, 0, 1, 0, 1}, {1, 1, 1, 1, 1, 1, 1, 0}}

// 	//grid := [][]int{{0, 0, 1, 0, 0}, {0, 1, 0, 1, 0}, {0, 1, 1, 1, 0}}

// 	//grid := [][]int{{1, 1, 1, 1, 1, 1, 1}, {1, 0, 0, 0, 0, 0, 1}, {1, 0, 1, 1, 1, 0, 1}, {1, 0, 1, 0, 1, 0, 1}, {1, 0, 1, 1, 1, 0, 1}, {1, 0, 0, 0, 0, 0, 1}, {1, 1, 1, 1, 1, 1, 1}}

// 	grid := [][]int{{0, 0, 1, 1, 0, 1, 0, 0, 1, 0}, {1, 1, 0, 1, 1, 0, 1, 1, 1, 0}, {1, 0, 1, 1, 1, 0, 0, 1, 1, 0}, {0, 1, 1, 0, 0, 0, 0, 1, 0, 1}, {0, 0, 0, 0, 0, 0, 1, 1, 1, 0}, {0, 1, 0, 1, 0, 1, 0, 1, 1, 1}, {1, 0, 1, 0, 1, 1, 0, 0, 0, 1}, {1, 1, 1, 1, 1, 1, 0, 0, 0, 0}, {1, 1, 1, 0, 0, 1, 0, 1, 0, 1}, {1, 1, 1, 0, 1, 1, 0, 1, 1, 0}}

// 	log.Println(closedIsland(grid))
// }
