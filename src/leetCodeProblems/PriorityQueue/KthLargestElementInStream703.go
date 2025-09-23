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

func (pq *elementsQueue) Push(num any) {
	*pq = append(*pq, num.(int))
}

// This is how PQ works - same way as stack
func (pq *elementsQueue) Pop() any {
	len := len(*pq)
	num := (*pq)[len-1]
	*pq = (*pq)[:len-1]
	return num
}

// This is again special case of PQ peek - peek from start
func (pq elementsQueue) Peek() any {
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
