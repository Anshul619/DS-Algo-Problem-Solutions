package main

/*
- LeetCode - https://leetcode.com/problems/delete-the-middle-node-of-a-linked-list/description/
- Time - O(n)
- Space - O(1)
*/
import "log"

func deleteMiddle(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return nil
	}

	var previous *ListNode
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		previous = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	previous.Next = slow.Next
	return head
}

func printList(head *ListNode) {

	for head != nil {
		log.Println(head.Val)
		head = head.Next
	}
}

func main() {

	// head := new(ListNode)
	// head.Val = 1
	// head.Next = new(ListNode)
	// head.Next.Val = 3
	// head.Next.Next = new(ListNode)
	// head.Next.Next.Val = 4
	// head.Next.Next.Next = new(ListNode)
	// head.Next.Next.Next.Val = 7
	// head.Next.Next.Next.Next = new(ListNode)
	// head.Next.Next.Next.Next.Val = 1
	// head.Next.Next.Next.Next.Next = new(ListNode)
	// head.Next.Next.Next.Next.Next.Val = 2
	// head.Next.Next.Next.Next.Next.Next = new(ListNode)
	// head.Next.Next.Next.Next.Next.Next.Val = 6

	// printList(deleteMiddle(head))

	head := new(ListNode)
	head.Val = 1
	head.Next = new(ListNode)
	head.Next.Val = 2
	head.Next.Next = new(ListNode)
	head.Next.Next.Val = 3
	head.Next.Next.Next = new(ListNode)
	head.Next.Next.Next.Val = 4

	printList(deleteMiddle(head))

	// head := new(ListNode)
	// head.Val = 2
	// head.Next = new(ListNode)
	// head.Next.Val = 1

	// //log.Println(deleteMiddle(head))

	// printList(deleteMiddle(head))
}
