package main

/*
- LeetCode - https://leetcode.com/problems/sort-list/description/
- Time - O(n*log n)
- Space - O(log n) // auxiliary (due to recursion stack), not counting the O(1) pointer rearrangements
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

// classic fast/slow pointer middle finding, splitting in-place.
func findMidAndSplit(head *ListNode) *ListNode {
	slow, fast := head, head

	var prev *ListNode

	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	if prev != nil {
		prev.Next = nil // split the list
	}
	return slow
}

// Standard merge of two sorted lists.
func merge(list1, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy

	for list1 != nil || list2 != nil {
		if list1 == nil {
			cur.Next = list2
			break
		}

		if list2 == nil {
			cur.Next = list1
			break
		}

		if list1.Val < list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}

		cur = cur.Next
	}
	return dummy.Next
}

// Recursive merge sort
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	mid := findMidAndSplit(head) // mid is start of second list
	left := sortList(head)       // head is already cut by findMid
	right := sortList(mid)
	return merge(left, right)
}
