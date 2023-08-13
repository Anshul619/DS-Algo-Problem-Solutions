package main

/*
- LeetCode - https://leetcode.com/problems/number-of-provinces/submissions/
*/

func visitNeighbours(isConnected [][]int, visited []bool, vertex int) {

	for i, v := range isConnected[vertex] {
		if v == 1 && !visited[i] {
			visited[i] = true
			visitNeighbours(isConnected, visited, i)
		}
	}
}

func findCircleNum(isConnected [][]int) int {

	vertices, out := len(isConnected), 0
	visited := make([]bool, vertices)

	for i := 0; i < vertices; i++ {
		if !visited[i] {
			visited[i] = true
			out++
			visitNeighbours(isConnected, visited, i)
		}
	}

	return out
}

// func main() {
// 	//input := [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}
// 	input := [][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
// 	log.Println(findCircleNum(input))
// }
