package main

import (
	"container/heap"
	"log"
)

/*
- LeetCode - https://leetcode.com/problems/kth-largest-element-in-an-array/
- Time - O(n + klogn)
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

func main() {
	log.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
	log.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
}
