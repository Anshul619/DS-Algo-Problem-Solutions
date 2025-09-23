package main

import (
	"container/heap"
)

/*
- LeetCode - https://leetcode.com/problems/kth-largest-element-in-an-array/
- Time - O(N*LogK)
- Space - O(k)
*/
type minHeap1 []int

func (h minHeap1) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h minHeap1) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h minHeap1) Len() int {
	return len(h)
}

func (h *minHeap1) Pop() any {
	t := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return t
}

func (h *minHeap1) Push(t any) {
	*h = append(*h, t.(int))
}

func (h minHeap1) Peek() int {
	return h[0]
}

func findKthLargest(nums []int, k int) int {
	pq := new(minHeap1)
	heap.Init(pq)

	for _, v := range nums {
		if pq.Len() < k {
			heap.Push(pq, v)
		} else if pq.Peek() < v {
			heap.Pop(pq)
			heap.Push(pq, v)
		}
	}

	return heap.Pop(pq).(int)
}
