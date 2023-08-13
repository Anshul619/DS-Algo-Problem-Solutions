package main

/*
- Leetcode - https://leetcode.com/problems/find-if-path-exists-in-graph/description/
- Time - O(n+m)
- Space - O(n+m)
*/
import "log"

type vQueue []int

func (q *vQueue) pop() int {
	temp := (*q)[0]
	*q = (*q)[1:]
	return temp
}

func (q *vQueue) push(num int) {
	*q = append(*q, num)
}

func (q vQueue) isEmpty() bool {
	return len(q) == 0
}

func (q vQueue) Len() int {
	return len(q)
}

func buildGraph(edges [][]int, n int) map[int][]int {
	g := make(map[int][]int)

	for i := 0; i < n; i++ {
		g[i] = []int{}
	}

	for _, v := range edges {
		g[v[0]] = append(g[v[0]], v[1])
		g[v[1]] = append(g[v[1]], v[0])
	}
	return g
}

func validPath(n int, edges [][]int, source int, destination int) bool {
	g := buildGraph(edges, n)
	visited := make([]bool, n)

	q := new(vQueue)
	q.push(source)

	for !q.isEmpty() {
		v := q.pop()

		if v == destination {
			return true
		}

		if !visited[v] {
			for _, v1 := range g[v] {
				q.push(v1)
			}
		}

		visited[v] = true
	}
	return false
}

func main() {
	log.Println(validPath(3, [][]int{{0, 1}, {1, 2}, {2, 0}}, 0, 2))
	log.Println(validPath(6, [][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}}, 0, 5))
	//log.Println(canVisitAllRooms([][]int{{1, 3}, {3, 0, 1}, {2}, {0}}))
}
