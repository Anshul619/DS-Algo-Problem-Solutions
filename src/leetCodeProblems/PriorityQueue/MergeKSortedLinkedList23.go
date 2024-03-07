package main

/*
- LeetCode - https://leetcode.com/problems/merge-k-sorted-lists/submissions/
- Time - O(N * K * logK)
- Space - O(K)
*/
import (
	"container/heap"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type nodesPQ []*ListNode

func (h nodesPQ) Len() int {
	return len(h)
}

func (h nodesPQ) Less(i int, j int) bool {
	return h[i].Val < h[j].Val
}

func (h nodesPQ) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *nodesPQ) Push(a interface{}) {
	*h = append(*h, a.(*ListNode))
}

func (h *nodesPQ) Pop() interface{} {
	l := len(*h)
	res := (*h)[l-1]
	*h = (*h)[:l-1]
	return res
}

func mergeKLists(lists []*ListNode) *ListNode {
	pq := new(nodesPQ)

	//log.Println(lists)
	var head *ListNode
	var cur *ListNode

	for i := 0; i < len(lists); i++ {
		heap.Push(pq, lists[i])
	}

	for pq.Len() != 0 {
		node := heap.Pop(pq).(*ListNode)

		if head == nil {
			head = node
			cur = node
		} else {
			cur.Next = node
			cur = cur.Next
		}

		if node.Next != nil {
			heap.Push(pq, node.Next)
		}
	}

	return head
}
