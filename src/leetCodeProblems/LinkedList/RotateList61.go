package main

/*
- LeetCode - https://leetcode.com/problems/rotate-list/
*/
import "log"

type ListNode struct {
	Next *ListNode
	Val  int
}

func rotateRight(head *ListNode, k int) *ListNode {

	fast := head
	slow := head
	total := 0

	if head == nil || head.Next == nil {
		return head
	}

	for fast.Next != nil {
		fast = fast.Next
		total++
	}

	total++

	if k > total {
		k = k % total
	}

	fast = head
	for i := 0; i < k; i++ {
		if fast.Next != nil {
			fast = fast.Next
		} else {
			return head
		}
	}

	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	fast.Next = head
	head = slow.Next
	slow.Next = nil

	return head
}

func print(node *ListNode) {
	for node != nil {
		log.Println(node.Val)
		node = node.Next
	}
}

// func main() {

// 	//head := ListNode{
// 	//	Val: 10,
// 	//	Next: &ListNode{
// 	//		Val:  10,
// 	//		Next: new(ListNode),
// 	//	},
// 	//}

// 	// head := new(ListNode)
// 	// head.Val = 0
// 	// head.Next = new(ListNode)
// 	// head.Next.Val = 1
// 	// head.Next.Next = new(ListNode)
// 	// head.Next.Next.Val = 2

// 	// print(rotateRight(head, 4))

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

// 	head := new(ListNode)
// 	head.Val = 1
// 	head.Next = new(ListNode)
// 	head.Next.Val = 2

// 	print(rotateRight(head, 5))

// 	// //var head *ListNode

// 	// print(rotateRight(head, 2))
// }
