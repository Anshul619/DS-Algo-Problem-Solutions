package main

import "log"

/*
- LeetCode - https://leetcode.com/problems/reverse-linked-list/
- TimeComplexity - O(n)
- SpaceComplexity - O(n)
*/

func reverseListUsingRecursion(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	new := reverseListUsingRecursion(head.Next)
	log.Println(head, "new", new.Val, "head.Next.Next", head.Next.Next)
	head.Next.Next = head
	head.Next = nil
	return new
}

// func main() {

// 	// head := new(ListNode)
// 	// head.Val = 1
// 	//head.Next = n

// 	head := new(ListNode)
// 	head.Val = 1
// 	head.Next = new(ListNode)
// 	head.Next.Val = 2
// 	head.Next.Next = new(ListNode)
// 	head.Next.Next.Val = 3
// 	head.Next.Next.Next = new(ListNode)
// 	head.Next.Next.Next.Val = 4
// 	head.Next.Next.Next.Next = new(ListNode)
// 	head.Next.Next.Next.Next.Val = 5

// 	printLinkedList(reverseListUsingRecursion(head))

// 	head = new(ListNode)
// 	head.Val = 1
// 	head.Next = new(ListNode)
// 	head.Next.Val = 2

// 	printLinkedList(reverseListUsingRecursion(head))

// 	head = new(ListNode)
// 	head.Val = 1
// 	head.Next = new(ListNode)
// 	head.Next.Val = 2
// 	head.Next.Next = new(ListNode)
// 	head.Next.Next.Val = 3
// 	printLinkedList(reverseListUsingRecursion(head))

// 	// head := new(ListNode)
// 	// head.Val = 1
// 	// head.Next = new(ListNode)
// 	// head.Next.Val = 2
// 	// head.Next.Next = new(ListNode)
// 	// head.Next.Next.Val = 3
// 	// head.Next.Next.Next = new(ListNode)
// 	// head.Next.Next.Next.Val = 4
// 	// head.Next.Next.Next.Next = new(ListNode)
// 	// head.Next.Next.Next.Next.Val = 5
// 	// head.Next.Next.Next.Next.Next = new(ListNode)
// 	// head.Next.Next.Next.Next.Next.Val = 6
// 	// head.Next.Next.Next.Next.Next.Next = new(ListNode)
// 	// head.Next.Next.Next.Next.Next.Next.Val = 7
// 	// head.Next.Next.Next.Next.Next.Next.Next = new(ListNode)
// 	// head.Next.Next.Next.Next.Next.Next.Next.Val = 8

// 	// reorderList(head)

// 	// /*head := new(ListNode)
// 	// head.Val = 1

// 	// output := removeNthFromEnd(head, 1)*/

// 	// output := removeNthFromEnd(head, 2)*/

// 	// temp := head

// 	// for temp != nil {
// 	// 	log.Println(temp.Val)
// 	// 	temp = temp.Next
// 	// }

// }
