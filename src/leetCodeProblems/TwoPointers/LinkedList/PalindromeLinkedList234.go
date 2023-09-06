package main

import "log"

/*
- Leetcode - https://leetcode.com/problems/palindrome-linked-list/
- Time - O(n)
- Space - O(1)
*/

func compareLists(node1, node2 *ListNode) bool {
	for node1 != nil && node2 != nil {
		if node1.Val != node2.Val {
			return false
		}
		node1 = node1.Next
		node2 = node2.Next
	}
	return true
}
func isPalindrome(head *ListNode) bool {

	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return compareLists(reverseList(slow), head)
}

func main() {
	head := new(ListNode)
	head.Val = 1
	head.Next = new(ListNode)
	head.Next.Val = 2
	head.Next.Next = new(ListNode)
	head.Next.Next.Val = 2
	head.Next.Next.Next = new(ListNode)
	head.Next.Next.Next.Val = 1

	log.Println(isPalindrome(head))

	head = new(ListNode)
	head.Val = 1
	head.Next = new(ListNode)
	head.Next.Val = 2
	log.Println(isPalindrome(head))

	head = new(ListNode)
	head.Val = 1
	head.Next = new(ListNode)
	head.Next.Val = 2
	head.Next.Next = new(ListNode)
	head.Next.Next.Val = 3
	head.Next.Next.Next = new(ListNode)
	head.Next.Next.Next.Val = 2
	head.Next.Next.Next.Next = new(ListNode)
	head.Next.Next.Next.Next.Val = 1

	log.Println(isPalindrome(head))
}
