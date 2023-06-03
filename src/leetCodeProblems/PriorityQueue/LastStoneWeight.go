package main

/*
- LeetCode - https://leetcode.com/problems/last-stone-weight/
- Time - O(nlogn)
- Space - O(n)
*/
import (
	"container/heap"
)

type PriorityQueue []int

func (pq PriorityQueue) Len() int {
	return len(pq)
}

// Max-heap
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i] > pq[j]
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(num interface{}) {
	*pq = append(*pq, num.(int))
}

func (pq *PriorityQueue) Pop() interface{} {
	len := len(*pq)
	num := (*pq)[len-1]
	*pq = (*pq)[:len-1]
	return num
}

func lastStoneWeight(stones []int) int {
	pq := new(PriorityQueue)

	for _, v := range stones {
		heap.Push(pq, v)
	}

	for pq.Len() > 1 {
		elem1 := heap.Pop(pq).(int)
		elem2 := heap.Pop(pq).(int)

		switch {
		case elem2 > elem1:
			heap.Push(pq, elem2-elem1)
		case elem1 > elem2:
			heap.Push(pq, elem1-elem2)
		}
	}

	if pq.Len() == 1 {
		return heap.Pop(pq).(int)
	} else {
		return 0
	}
}

// func main() {
// 	log.Println(lastStoneWeight([]int{2, 7, 4, 1, 8, 1}))
// 	log.Println(lastStoneWeight([]int{1}))
// 	log.Println(lastStoneWeight([]int{2, 2}))
// }
