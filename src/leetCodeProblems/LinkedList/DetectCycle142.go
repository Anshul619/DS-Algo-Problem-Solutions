/*
- LeetCode - https://leetcode.com/problems/linked-list-cycle-ii/
*/

package main

import "log"

func detectCycle(head *ListNode) *ListNode {

	nodesMap := make(map[*ListNode]bool)

	if head == nil {
		return nil
	}

	nodesMap[head] = true

	for head.Next != nil {

		head = head.Next

		if _, ok := nodesMap[head]; ok {
			return head
		}

		nodesMap[head] = true
	}

	return nil
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

	// head := new(ListNode)
	// head.Val = 3
	// head.Next = new(ListNode)
	// head.Next.Val = 2
	// head.Next.Next = new(ListNode)
	// head.Next.Next.Val = 0
	// head.Next.Next.Next = new(ListNode)
	// head.Next.Next.Next.Val = -4
	// head.Next.Next.Next.Next = head.Next

	// head := new(ListNode)
	// head.Val = 1
	// head.Next = new(ListNode)
	// head.Next.Val = 2
	// head.Next.Next = head

	// head := new(ListNode)
	// head.Val = -1
	// head.Next = new(ListNode)
	// head.Next.Val = -7
	// head.Next.Next = new(ListNode)
	// head.Next.Next.Val = 7
	// head.Next.Next.Next = new(ListNode)
	// head.Next.Next.Next.Val = -4
	// head.Next.Next.Next.Next = new(ListNode)
	// head.Next.Next.Next.Next.Val = 19
	// head.Next.Next.Next.Next.Next = new(ListNode)
	// head.Next.Next.Next.Next.Next.Val = 6
	// head.Next.Next.Next.Next.Next.Next = new(ListNode)
	// head.Next.Next.Next.Next.Next.Next.Val = -9
	// head.Next.Next.Next.Next.Next.Next.Next = new(ListNode)
	// head.Next.Next.Next.Next.Next.Next.Next.Val = -5
	// head.Next.Next.Next.Next.Next.Next.Next.Next = new(ListNode)
	// head.Next.Next.Next.Next.Next.Next.Next.Next.Val = -2
	// head.Next.Next.Next.Next.Next.Next.Next.Next.Next = new(ListNode)
	// head.Next.Next.Next.Next.Next.Next.Next.Next.Next.Val = -5
	// head.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next = head.Next.Next.Next.Next.Next.Next

	head := new(ListNode)
	head.Val = 1
	head.Next = nil
	// head.Next.Val = 2
	// head.Next.Next = head

	//log.Println
	log.Println(detectCycle(head))
}
