package main

/*
- Leetcode - https://leetcode.com/problems/remove-duplicates-from-sorted-list/description/
*/
import (
	"log"
	"reflect"
	"testing"
)

func getSlice(list *ListNode) []int {
	out := []int{}

	for list != nil {
		out = append(out, list.Val)
		list = list.Next
	}
	log.Println(out)
	return out
}

func TestDeleteDuplicates(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		list := new(ListNode)
		list.Val = 1
		list.Next = new(ListNode)
		list.Next.Val = 1
		list.Next.Next = new(ListNode)
		list.Next.Next.Val = 2

		if !reflect.DeepEqual(getSlice(deleteDuplicates(list)), []int{1, 2}) {
			t.Fail()
		}
	})

	t.Run("test2", func(t *testing.T) {
		list := new(ListNode)
		list.Val = 1
		list.Next = new(ListNode)
		list.Next.Val = 1
		list.Next.Next = new(ListNode)
		list.Next.Next.Val = 2
		list.Next.Next.Next = new(ListNode)
		list.Next.Next.Next.Val = 3
		list.Next.Next.Next.Next = new(ListNode)
		list.Next.Next.Next.Next.Val = 3

		if !reflect.DeepEqual(getSlice(deleteDuplicates(list)), []int{1, 2, 3}) {
			t.Fail()
		}
	})

	t.Run("test3", func(t *testing.T) {
		if !reflect.DeepEqual(getSlice(deleteDuplicates(nil)), []int{}) {
			t.Fail()
		}
	})
}
