package main

/*
- LeetCode - https://leetcode.com/problems/swapping-nodes-in-a-linked-list/description/
*/
import (
	"reflect"
	"testing"
)

func TestSwapNodes(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		list := new(ListNode)
		list.Val = 1
		list.Next = new(ListNode)
		list.Next.Val = 2
		list.Next.Next = new(ListNode)
		list.Next.Next.Val = 3
		list.Next.Next.Next = new(ListNode)
		list.Next.Next.Next.Val = 4
		list.Next.Next.Next.Next = new(ListNode)
		list.Next.Next.Next.Next.Val = 5

		if !reflect.DeepEqual(getSlice(swapNodes(list, 2)), []int{1, 4, 3, 2, 5}) {
			t.Fail()
		}
	})
}
