package main

/*
- LeetCode - https://leetcode.com/problems/delete-node-in-a-linked-list
- Time - O(1)
- Space - O(1)
*/
import "log"

func deleteNode(node *ListNode) {
	nextVal := node.Next.Val
	nextToNextNode := node.Next.Next

	node.Next = nextToNextNode
	node.Val = nextVal
}

func printLinkedList(head *ListNode) {

	for head != nil {
		log.Println(head.Val)
		head = head.Next
	}
}

// func main() {

// 	head := new(ListNode)
// 	head.Val = 4
// 	head.Next = new(ListNode)
// 	head.Next.Val = 5
// 	head.Next.Next = new(ListNode)
// 	head.Next.Next.Val = 1
// 	head.Next.Next.Next = new(ListNode)
// 	head.Next.Next.Next.Val = 9

// 	deleteNode(head.Next)

// 	printLinkedList(head)

// }
