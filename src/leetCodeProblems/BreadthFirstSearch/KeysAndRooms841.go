package main

/*
- Leetcode - https://leetcode.com/problems/keys-and-rooms/description/
- Time - O(n+e) where n = number of rooms & e = number of keys
- Space - O(n)
*/

type queue []int

func (q *queue) push(n int) {
	*q = append(*q, n)
}

func (q *queue) pop() int {
	t := (*q)[0]
	*q = (*q)[1:]
	return t
}

func (q queue) isEmpty() bool {
	return len(q) == 0
}

func canVisitAllRooms(rooms [][]int) bool {
	visited := make([]bool, len(rooms))
	q := new(queue)

	q.push(0)

	for !q.isEmpty() {
		k := q.pop()

		for _, v := range rooms[k] {
			if !visited[v] {
				q.push(v)
			}
		}

		visited[k] = true
	}

	for _, v := range visited {
		if !v {
			return false
		}
	}
	return true
}
