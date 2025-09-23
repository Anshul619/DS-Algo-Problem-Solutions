package main

/*
- LeetCode - https://leetcode.com/problems/maximum-number-of-events-that-can-be-attended/description/
- Time - O(nlogn)
- Space - O(n)
*/
import (
	"container/heap"
	"sort"
)

type ongoingEventsQueue []int

func (pq ongoingEventsQueue) isEmpty() bool {
	return pq.Len() == 0
}

func (pq ongoingEventsQueue) Len() int {
	return len(pq)
}

// Min-heap
func (pq ongoingEventsQueue) Less(i, j int) bool {
	return pq[i] < pq[j]
}

func (pq ongoingEventsQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *ongoingEventsQueue) Push(num any) {
	*pq = append(*pq, num.(int))
}

func (pq *ongoingEventsQueue) Pop() any {
	len := len(*pq)
	num := (*pq)[len-1]
	*pq = (*pq)[:len-1]
	return num
}

func (pq ongoingEventsQueue) Peek() any {
	return pq[0]
}

func maxEvents(events [][]int) int {
	sort.SliceStable(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})

	pq := new(ongoingEventsQueue)
	eventsIndex, out := 0, 0

	for currentDay := 1; currentDay <= 100000; currentDay++ {

		for !pq.isEmpty() && pq.Peek().(int) < currentDay {
			heap.Pop(pq)
		}

		for eventsIndex < len(events) && events[eventsIndex][0] == currentDay {
			heap.Push(pq, events[eventsIndex][1])
			eventsIndex++
		}

		if !pq.isEmpty() {
			heap.Pop(pq)
			out++
		}
	}

	return out
}

// func main() {
// 	// log.Println(maxEvents([][]int{{1, 2}, {2, 3}, {3, 4}}))
// 	// log.Println(maxEvents([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 2}}))
// 	// log.Println(maxEvents([][]int{{1, 1}, {2, 2}, {1, 2}}))
// 	// log.Println(maxEvents([][]int{{1, 1}, {1, 2}, {1, 3}, {1, 4}, {1, 5}, {1, 6}, {1, 7}}))
// 	// log.Println(maxEvents([][]int{{1, 2}, {1, 2}, {3, 3}, {1, 5}, {1, 5}}))
// 	// log.Println(maxEvents([][]int{{1, 5}, {1, 5}, {1, 5}, {2, 3}, {2, 3}}))
// 	log.Println(maxEvents([][]int{{1, 2}, {1, 2}, {1, 6}, {1, 2}, {1, 2}}))
// }
