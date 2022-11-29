package main

/*
- LeetCode - https://leetcode.com/problems/merge-k-sorted-lists/submissions/
- Time - O(n*k*logk)
- Space - O(k)
*/
import (
	"container/heap"
	"log"
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

	log.Println(lists)
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

func main() {

	input := []*ListNode{}

	// list1 := new(ListNode)
	// list1.Val = 1
	// list1.Next = new(ListNode)
	// list1.Next.Val = 4
	// list1.Next.Next = new(ListNode)
	// list1.Next.Next.Val = 5

	// input = append(input, list1)

	// list2 := new(ListNode)
	// list2.Val = 1
	// list2.Next = new(ListNode)
	// list2.Next.Val = 3
	// list2.Next.Next = new(ListNode)
	// list2.Next.Next.Val = 4

	// input = append(input, list2)

	// list3 := new(ListNode)
	// list3.Val = 2
	// list3.Next = new(ListNode)
	// list3.Next.Val = 6

	// input = append(input, list3)

	// list4 := new(ListNode)
	// input = append(input, list4)

	out := mergeKLists(input)

	for out != nil {
		log.Println(out.Val)
		out = out.Next
	}
}
