package main

import "log"

type ListNode struct {
	Next *ListNode
	Val  int
}

func getSlice(list *ListNode) []int {
	out := []int{}

	for list != nil {
		out = append(out, list.Val)
		list = list.Next
	}
	log.Println(out)
	return out
}
