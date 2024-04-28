package main

import (
	"reflect"
	"testing"
)

func getSlice(head *ListNode) []int {
	out := []int{}

	for head != nil {
		out = append(out, head.Val)
		head = head.Next
	}
	return out
}

func TestReverseList(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
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

		if !reflect.DeepEqual(getSlice(reverseListUsingRecursion(head)), []int{5, 4, 3, 2, 1}) {
			t.Fail()
		}

	})

	t.Run("test2", func(t *testing.T) {
		head := new(ListNode)
		head.Val = 1
		head.Next = new(ListNode)
		head.Next.Val = 2

		if !reflect.DeepEqual(getSlice(reverseListUsingRecursion(head)), []int{2, 1}) {
			t.Fail()
		}
	})

	t.Run("test3", func(t *testing.T) {
		if !reflect.DeepEqual(getSlice(reverseListUsingRecursion(nil)), []int{}) {
			t.Fail()
		}
	})
}
