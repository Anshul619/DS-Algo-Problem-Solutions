package main

/*
- LeetCode - https://leetcode.com/problems/merge-k-sorted-lists/submissions/
- Time - O(N * LogK)
- Space - O(K)
*/
import (
	"container/heap"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type minHeap2 []*ListNode

func (h minHeap2) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h minHeap2) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap2) Pop() interface{} {
	t := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return t
}

func (h *minHeap2) Push(i interface{}) {
	*h = append(*h, i.(*ListNode))
}

func (h minHeap2) Len() int {
	return len(h)
}

func mergeKLists(lists []*ListNode) *ListNode {
	h := new(minHeap2)
	heap.Init(h)

	dummy := &ListNode{}
	cur := dummy

	for _, v := range lists {
		if v != nil {
			heap.Push(h, v)
		}
	}

	for h.Len() > 0 {
		e := heap.Pop(h).(*ListNode)
		cur.Next = e
		cur = cur.Next
		if e.Next != nil {
			heap.Push(h, e.Next)
		}
	}
	return dummy.Next
}
