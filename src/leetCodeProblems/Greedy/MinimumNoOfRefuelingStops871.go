package main

/*
- LeetCode - https://leetcode.com/problems/minimum-number-of-refueling-stops/
- Time - O(N*logN)
- Space - O(N)
*/
import (
	"container/heap"
	"log"
	"sort"
)

type stopsQueue []int

func (h stopsQueue) Len() int {
	return len(h)
}

func (h stopsQueue) Less(i int, j int) bool {
	return h[i] > h[j]
}

func (h stopsQueue) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *stopsQueue) Push(a interface{}) {
	*h = append(*h, a.(int))
}

func (h *stopsQueue) Pop() interface{} {
	l := len(*h)
	res := (*h)[l-1]
	*h = (*h)[:l-1]
	return res
}

func minRefuelStops(target int, startFuel int, stations [][]int) int {
	sort.SliceStable(stations, func(i, j int) bool {
		return stations[i][0] < stations[j][0]
	})

	pq := new(stopsQueue)
	out := 0

	stationsLength := len(stations)

	for i := 0; i <= stationsLength; i++ {

		log.Println(pq)
		dist := 0

		if i == stationsLength {
			dist = target
		} else {
			dist = stations[i][0]
		}

		log.Println(startFuel, dist)
		for startFuel < dist {

			if pq.Len() == 0 {
				return -1
			}

			startFuel += heap.Pop(pq).(int)
			out++
		}

		if i < stationsLength {
			heap.Push(pq, stations[i][1])
		}
	}

	return out
}

// func main() {

// 	// stations := [][]int{{10, 60}, {20, 30}, {30, 30}, {60, 40}}
// 	// log.Println(minRefuelStops(100, 10, stations))

// 	stations := [][]int{{25, 30}}
// 	log.Println(minRefuelStops(100, 50, stations))

// }
