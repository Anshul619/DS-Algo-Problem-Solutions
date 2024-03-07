package main

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

func TestMergeKLists(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		list1 := new(ListNode)
		list1.Val = 1
		list1.Next = new(ListNode)
		list1.Next.Val = 4
		list1.Next.Next = new(ListNode)
		list1.Next.Next.Val = 5

		list2 := new(ListNode)
		list2.Val = 1
		list2.Next = new(ListNode)
		list2.Next.Val = 3
		list2.Next.Next = new(ListNode)
		list2.Next.Next.Val = 4

		list3 := new(ListNode)
		list3.Val = 2
		list3.Next = new(ListNode)
		list3.Next.Val = 6

		lists := []*ListNode{list1, list2, list3}

		if !reflect.DeepEqual(getSlice(mergeKLists(lists)), []int{1, 1, 2, 3, 4, 4, 5, 6}) {
			t.Fail()
		}

	})

	t.Run("test2", func(t *testing.T) {
		list1 := new(ListNode)
		list1.Val = 1

		list2 := new(ListNode)
		list2.Val = 0

		lists := []*ListNode{list1, list2}

		if !reflect.DeepEqual(getSlice(mergeKLists(lists)), []int{0, 1}) {
			t.Fail()
		}

	})
}
