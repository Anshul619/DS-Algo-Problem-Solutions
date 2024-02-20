package main

/*
- Leetcode - https://leetcode.com/problems/partition-list
*/
import (
	"reflect"
	"testing"
)

func TestPartition(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		head := new(ListNode)
		head.Val = 1
		head.Next = new(ListNode)
		head.Next.Val = 4
		head.Next.Next = new(ListNode)
		head.Next.Next.Val = 3
		head.Next.Next.Next = new(ListNode)
		head.Next.Next.Next.Val = 2
		head.Next.Next.Next.Next = new(ListNode)
		head.Next.Next.Next.Next.Val = 5
		head.Next.Next.Next.Next.Next = new(ListNode)
		head.Next.Next.Next.Next.Next.Val = 2

		if !reflect.DeepEqual(getSlice(partition(head, 3)), []int{1, 2, 2, 4, 3, 5}) {
			t.Fail()
		}
	})

	t.Run("test2", func(t *testing.T) {
		head := new(ListNode)
		head.Val = 2
		head.Next = new(ListNode)
		head.Next.Val = 1

		if !reflect.DeepEqual(getSlice(partition(head, 2)), []int{1, 2}) {
			t.Fail()
		}
	})

	t.Run("test3", func(t *testing.T) {
		head := new(ListNode)
		head.Val = 1
		head.Next = new(ListNode)
		head.Next.Val = 1

		if !reflect.DeepEqual(getSlice(partition(head, 2)), []int{1, 1}) {
			t.Fail()
		}
	})

	t.Run("test4", func(t *testing.T) {
		head := new(ListNode)
		head.Val = 5
		head.Next = new(ListNode)
		head.Next.Val = 4
		head.Next.Next = new(ListNode)
		head.Next.Next.Val = 3
		head.Next.Next.Next = new(ListNode)
		head.Next.Next.Next.Val = 2
		head.Next.Next.Next.Next = new(ListNode)
		head.Next.Next.Next.Next.Val = 5
		head.Next.Next.Next.Next.Next = new(ListNode)
		head.Next.Next.Next.Next.Next.Val = 2

		if !reflect.DeepEqual(getSlice(partition(head, 3)), []int{2, 2, 5, 4, 3, 5}) {
			t.Fail()
		}
	})
}
