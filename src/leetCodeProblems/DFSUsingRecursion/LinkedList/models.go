package main

import "log"

type ListNode struct {
	Val  int
	Next *ListNode
}

func printDoublyLinkedList(head *Node) {

	for head != nil {
		log.Println(head.Val)
		if head.Prev != nil {
			log.Println("Prev", head.Prev.Val)
		}
		if head.Child != nil {
			log.Println("Child", head.Child.Val)
		}
		head = head.Next
	}
}

func printLinkedList(head *ListNode) {

	for head != nil {
		log.Println(head.Val)
		head = head.Next
	}
}
