package main

/*
- LeetCode - https://leetcode.com/problems/merge-two-sorted-lists/description/
- Time - O(m+n)
- Space - O(1)
*/

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	list1Next := list1
	list2Next := list2

	var list1Prev *ListNode

	for list1Next != nil && list2Next != nil {

		if list1Next.Val > list2Next.Val {
			list2Temp := list2Next
			list2Next = list2Next.Next

			if list1Prev == nil {
				list1 = list2Temp
			} else {
				list1Prev.Next = list2Temp
			}

			list2Temp.Next = list1Next
			list1Prev = list2Temp
		} else {
			if list1Prev == nil {
				list1Prev = list1
			} else {
				list1Prev = list1Prev.Next
			}

			list1Next = list1Next.Next
		}
	}

	if list2Next != nil {

		if list1Prev == nil {
			list1 = list2Next
		} else {
			list1Prev.Next = list2Next
		}
	}
	return list1
}

func main() {
	head := new(ListNode)
	head.Val = 1
	head.Next = new(ListNode)
	head.Next.Val = 2
	head.Next.Next = new(ListNode)
	head.Next.Next.Val = 4

	head1 := new(ListNode)
	head1.Val = 1
	head1.Next = new(ListNode)
	head1.Next.Val = 3
	head1.Next.Next = new(ListNode)
	head1.Next.Next.Val = 5

	printList(mergeTwoLists(nil, head1))
	printList(mergeTwoLists(head, head1))
	printList(mergeTwoLists(nil, nil))
}
