package main

/*
- LeetCode - https://leetcode.com/problems/linked-list-cycle-ii/
*/

import (
	"log"
)

func detectCycle(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	slow := head
	fast := head
	isCycle := false

	for slow != nil && fast != nil && fast.Next != nil {

		slow = slow.Next
		fast = fast.Next.Next

		log.Println("Slow - ", slow.Val)
		log.Println("Fast - ", fast.Val)
		if slow == fast {
			isCycle = true
			break
		}
	}

	if !isCycle {
		return nil
	}

	for head != slow {
		head = head.Next
		slow = slow.Next
	}

	return head
}

// func main() {
// 	// head := new(ListNode)
// 	// head.Val = 0
// 	// head.Next = new(ListNode)
// 	// head.Next.Val = 3
// 	// head.Next.Next = new(ListNode)
// 	// head.Next.Next.Val = 1
// 	// head.Next.Next.Next = new(ListNode)
// 	// head.Next.Next.Next.Val = 0
// 	// head.Next.Next.Next.Next = new(ListNode)
// 	// head.Next.Next.Next.Next.Val = 4
// 	// head.Next.Next.Next.Next.Next = new(ListNode)
// 	// head.Next.Next.Next.Next.Next.Val = 5
// 	// head.Next.Next.Next.Next.Next.Next = new(ListNode)
// 	// head.Next.Next.Next.Next.Next.Next.Val = 2
// 	// head.Next.Next.Next.Next.Next.Next.Next = new(ListNode)
// 	// head.Next.Next.Next.Next.Next.Next.Next.Val = 0

// 	// head := new(ListNode)
// 	// head.Val = 3
// 	// head.Next = new(ListNode)
// 	// head.Next.Val = 2
// 	// head.Next.Next = new(ListNode)
// 	// head.Next.Next.Val = 0
// 	// head.Next.Next.Next = new(ListNode)
// 	// head.Next.Next.Next.Val = -4
// 	// head.Next.Next.Next.Next = head.Next

// 	// head := new(ListNode)
// 	// head.Val = 1
// 	// head.Next = new(ListNode)
// 	// head.Next.Val = 2
// 	// head.Next.Next = head

// 	head := new(ListNode)
// 	head.Val = -1
// 	head.Next = new(ListNode)
// 	head.Next.Val = -7
// 	head.Next.Next = new(ListNode)
// 	head.Next.Next.Val = 7
// 	head.Next.Next.Next = new(ListNode)
// 	head.Next.Next.Next.Val = -4
// 	head.Next.Next.Next.Next = new(ListNode)
// 	head.Next.Next.Next.Next.Val = 19
// 	head.Next.Next.Next.Next.Next = new(ListNode)
// 	head.Next.Next.Next.Next.Next.Val = 6
// 	head.Next.Next.Next.Next.Next.Next = new(ListNode)
// 	head.Next.Next.Next.Next.Next.Next.Val = -9
// 	head.Next.Next.Next.Next.Next.Next.Next = new(ListNode)
// 	head.Next.Next.Next.Next.Next.Next.Next.Val = -5
// 	head.Next.Next.Next.Next.Next.Next.Next.Next = new(ListNode)
// 	head.Next.Next.Next.Next.Next.Next.Next.Next.Val = -2
// 	head.Next.Next.Next.Next.Next.Next.Next.Next.Next = new(ListNode)
// 	head.Next.Next.Next.Next.Next.Next.Next.Next.Next.Val = -5
// 	head.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next = head.Next.Next.Next.Next.Next.Next

// 	// head := new(ListNode)
// 	// head.Val = 1
// 	// head.Next = new(ListNode)
// 	// head.Next.Val = 2
// 	// head.Next.Next = head

// 	log.Println(detectCycle(head))
// }
