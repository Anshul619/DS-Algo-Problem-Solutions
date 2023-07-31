package main

/*
- Leetcode - https://leetcode.com/problems/keys-and-rooms/description/
- Time - O(n)
- Space - O(n)
*/
import "log"

type keysQueue []int

func (q *keysQueue) pop() int {
	temp := (*q)[0]
	*q = (*q)[1:]
	return temp
}

func (q *keysQueue) push(num int) {
	*q = append(*q, num)
}

func (q keysQueue) isEmpty() bool {
	return len(q) == 0
}

func (q keysQueue) Len() int {
	return len(q)
}

func canVisitAllRooms(rooms [][]int) bool {
	visited := make(map[int]bool)
	q := new(keysQueue)

	q.push(0)

	for !q.isEmpty() {
		key := q.pop()

		visited[key] = true

		for i := 0; i < len(rooms[key]); i++ {
			if _, ok := visited[rooms[key][i]]; !ok {
				q.push(rooms[key][i])
			}
		}
	}

	for i := 0; i < len(rooms); i++ {
		if _, ok := visited[i]; !ok {
			return false
		}
	}

	return true
}

func main() {
	log.Println(canVisitAllRooms([][]int{{1}, {2}, {3}, {}}))
	log.Println(canVisitAllRooms([][]int{{1, 3}, {3, 0, 1}, {2}, {0}}))
}
