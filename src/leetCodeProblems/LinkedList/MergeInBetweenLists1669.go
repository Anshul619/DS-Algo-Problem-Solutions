package main

/*
- LeetCode - https://leetcode.com/problems/merge-in-between-linked-lists/description/
- Time - O(n)
- Space - O(1)
*/
import "log"

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {

	list1Iterator, list1IteratorPrev, list1Counter := list1, list1, 0
	list2Iterator, list2End := list2, list2

	for list2Iterator != nil {
		list2End = list2Iterator
		list2Iterator = list2Iterator.Next
	}

	for list1Iterator != nil {

		if list1Counter == a {
			list1IteratorPrev.Next = list2
		}

		if list1Counter == b {
			list2End.Next = list1Iterator.Next
			break
		}

		list1Counter++

		list1IteratorPrev = list1Iterator
		list1Iterator = list1Iterator.Next
	}

	return list1
}

func main() {

	head := new(ListNode)
	head.Val = 0
	head.Next = new(ListNode)
	head.Next.Val = 1
	head.Next.Next = new(ListNode)
	head.Next.Next.Val = 2
	head.Next.Next.Next = new(ListNode)
	head.Next.Next.Next.Val = 3
	head.Next.Next.Next.Next = new(ListNode)
	head.Next.Next.Next.Next.Val = 4
	head.Next.Next.Next.Next.Next = new(ListNode)
	head.Next.Next.Next.Next.Next.Val = 5

	head1 := new(ListNode)
	head1.Val = 1000000
	head1.Next = new(ListNode)
	head1.Next.Val = 1000001
	head1.Next.Next = new(ListNode)
	head1.Next.Next.Val = 1000002

	new := mergeInBetween(head, 3, 4, head1)

	for new != nil {
		log.Println(new.Val)
		new = new.Next
	}
}
