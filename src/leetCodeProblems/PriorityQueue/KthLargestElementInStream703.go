package main

import (
	"container/heap"
)

/*
- LeetCode - https://leetcode.com/problems/kth-largest-element-in-a-stream/description/
- Time - O(nlogn)
- Space - O(k)
*/

type elementsQueue []int

func (pq elementsQueue) isEmpty() bool {
	return pq.Len() == 0
}

func (pq elementsQueue) Len() int {
	return len(pq)
}

// Min-heap
func (pq elementsQueue) Less(i, j int) bool {
	return pq[i] < pq[j]
}

func (pq elementsQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *elementsQueue) Push(num interface{}) {
	*pq = append(*pq, num.(int))
}

func (pq *elementsQueue) Pop() interface{} {
	len := len(*pq)
	num := (*pq)[len-1]
	*pq = (*pq)[:len-1]
	return num
}

func (pq elementsQueue) Peek() interface{} {
	return pq[0]
}

type KthLargest struct {
	pq *elementsQueue
	k  int
}

func Constructor(k int, nums []int) KthLargest {
	pq := new(elementsQueue)
	for _, v := range nums {
		heap.Push(pq, v)
	}
	for pq.Len() > k {
		heap.Pop(pq)
	}
	return KthLargest{
		pq: pq,
		k:  k,
	}
}

func (this *KthLargest) Add(val int) int {
	if this.pq.Len() == this.k {
		if val <= this.pq.Peek().(int) {
			return this.pq.Peek().(int)
		}

		heap.Pop(this.pq)
	}

	heap.Push(this.pq, val)
	return this.pq.Peek().(int)
}

// func main() {
// 	// obj := Constructor(3, []int{4, 5, 8, 2})
// 	// log.Println(obj.Add(3))
// 	// log.Println(obj.Add(5))
// 	// log.Println(obj.Add(10))
// 	// log.Println(obj.Add(9))
// 	// log.Println(obj.Add(4))

// 	obj := Constructor(2, []int{0})
// 	log.Println(obj.Add(-1))
// 	log.Println(obj.Add(1))
// 	log.Println(obj.Add(-2))
// 	log.Println(obj.Add(-4))
// 	log.Println(obj.Add(3))
// }
