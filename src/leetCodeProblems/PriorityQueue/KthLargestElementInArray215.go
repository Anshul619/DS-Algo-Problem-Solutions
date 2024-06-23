package main

import (
	"container/heap"
)

/*
- LeetCode - https://leetcode.com/problems/kth-largest-element-in-an-array/
- Time - O(N*LogK)
- Space - O(k)
*/
func findKthLargest(nums []int, k int) int {
	pq := new(elementsQueue)

	for _, v := range nums {
		if pq.Len() == k {
			if v <= pq.Peek().(int) {
				continue
			}

			heap.Pop(pq)
		}
		heap.Push(pq, v)
	}
	return pq.Peek().(int)
}
