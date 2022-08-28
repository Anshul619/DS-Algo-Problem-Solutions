package main

/*
- LeetCode - https://leetcode.com/problems/merge-nodes-in-between-zeros/
- TimeComplexity - O(n)
*/
import "log"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeNodes(head *ListNode) *ListNode {

	var outHead *ListNode
	var outCurrent *ListNode

	next, sum := head.Next, 0

	for next != nil {

		if next.Val == 0 {

			//log.Println(sum)
			temp := new(ListNode)
			temp.Val = sum

			if outHead == nil {
				outHead = temp
				outCurrent = temp
			} else {
				outCurrent.Next = temp
				outCurrent = temp
			}

			sum = 0

		} else {
			sum += next.Val
		}

		next = next.Next
	}

	return outHead
}

func print(node *ListNode) {
	for node != nil {
		log.Println(node.Val)
		node = node.Next
	}
}

func main() {
	// head := new(ListNode)
	// head.Val = 0
	// head.Next = new(ListNode)
	// head.Next.Val = 3
	// head.Next.Next = new(ListNode)
	// head.Next.Next.Val = 1
	// head.Next.Next.Next = new(ListNode)
	// head.Next.Next.Next.Val = 0
	// head.Next.Next.Next.Next = new(ListNode)
	// head.Next.Next.Next.Next.Val = 4
	// head.Next.Next.Next.Next.Next = new(ListNode)
	// head.Next.Next.Next.Next.Next.Val = 5
	// head.Next.Next.Next.Next.Next.Next = new(ListNode)
	// head.Next.Next.Next.Next.Next.Next.Val = 2
	// head.Next.Next.Next.Next.Next.Next.Next = new(ListNode)
	// head.Next.Next.Next.Next.Next.Next.Next.Val = 0

	head := new(ListNode)
	head.Val = 0
	head.Next = new(ListNode)
	head.Next.Val = 1
	head.Next.Next = new(ListNode)
	head.Next.Next.Val = 0
	head.Next.Next.Next = new(ListNode)
	head.Next.Next.Next.Val = 3
	head.Next.Next.Next.Next = new(ListNode)
	head.Next.Next.Next.Next.Val = 0
	head.Next.Next.Next.Next.Next = new(ListNode)
	head.Next.Next.Next.Next.Next.Val = 2
	head.Next.Next.Next.Next.Next.Next = new(ListNode)
	head.Next.Next.Next.Next.Next.Next.Val = 2
	head.Next.Next.Next.Next.Next.Next.Next = new(ListNode)
	head.Next.Next.Next.Next.Next.Next.Next.Val = 0

	log.Println("A")
	print(mergeNodes(head))
}
