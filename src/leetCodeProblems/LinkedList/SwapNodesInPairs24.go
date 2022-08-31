package main

/*
- LeetCode - https://leetcode.com/problems/swap-nodes-in-pairs/submissions/
- TimeCompelxity - O(n)
*/

import "log"

type ListNode struct {
	Next *ListNode
	Val  int
}

func swapPairs(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	var prev *ListNode
	cur := head

	for cur.Next != nil {

		temp := cur.Next

		cur.Next = temp.Next
		temp.Next = cur

		if prev != nil {
			prev.Next = temp
		} else {
			head = temp
		}

		if cur.Next != nil {
			prev = cur
			cur = cur.Next
		} else {
			break
		}
	}

	return head
}

func print(node *ListNode) {
	for node != nil {
		log.Println(node.Val)
		node = node.Next
	}
}
func main() {

	//head := ListNode{
	//	Val: 10,
	//	Next: &ListNode{
	//		Val:  10,
	//		Next: new(ListNode),
	//	},
	//}

	//head := new(ListNode)
	//head.Val = 1
	//head.Next = new(ListNode)
	//head.Next.Val = 2
	//head.Next.Next = new(ListNode)
	//head.Next.Next.Val = 3

	head := new(ListNode)
	head.Val = 1
	head.Next = new(ListNode)
	head.Next.Val = 2
	head.Next.Next = new(ListNode)
	head.Next.Next.Val = 3
	head.Next.Next.Next = new(ListNode)
	head.Next.Next.Next.Val = 4
	head.Next.Next.Next.Next = new(ListNode)
	head.Next.Next.Next.Next.Val = 5

	// head := new(ListNode)
	// head.Val = 1

	//var head *ListNode

	print(swapPairs(head))
}
