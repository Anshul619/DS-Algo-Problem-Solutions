package main

import "log"

/*
- Leetcode - https://leetcode.com/problems/palindrome-linked-list/
- Time - O(n)
- Space - O(n) if function call stack is considered, otherwise O(1)
*/

var left *ListNode

func isPalindromeUtil(right *ListNode) bool {
	if right == nil {
		return true
	}

	if !isPalindromeUtil(right.Next) {
		return false
	}

	isPalindrome := left.Val == right.Val
	left = left.Next

	return isPalindrome
}

func isPalindrome(head *ListNode) bool {
	left = head
	return isPalindromeUtil(head)
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
}
