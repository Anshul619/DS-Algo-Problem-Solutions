package main

/*
- LeetCode - https://leetcode.com/problems/merge-two-sorted-lists/description/
- Time - O(m+n)
- Space - O(m+n)
*/
func insertNewNode(outHead *ListNode, outNext *ListNode, newNode *ListNode) (*ListNode, *ListNode) {
	if outHead == nil {
		outHead = newNode
		outNext = newNode
	} else {
		outNext.Next = newNode
		outNext = outNext.Next
	}

	return outHead, outNext
}

func mergeTwoLists1(list1 *ListNode, list2 *ListNode) *ListNode {

	list1Next := list1
	list2Next := list2

	var outHead, outNext *ListNode

	for list1Next != nil && list2Next != nil {

		newNode := new(ListNode)

		if list1Next.Val < list2Next.Val {
			newNode.Val = list1Next.Val
			list1Next = list1Next.Next
		} else {
			newNode.Val = list2Next.Val
			list2Next = list2Next.Next
		}

		outHead, outNext = insertNewNode(outHead, outNext, newNode)
	}

	for list1Next != nil {
		newNode := new(ListNode)
		newNode.Val = list1Next.Val
		list1Next = list1Next.Next
		outHead, outNext = insertNewNode(outHead, outNext, newNode)
	}

	for list2Next != nil {
		newNode := new(ListNode)
		newNode.Val = list2Next.Val
		list2Next = list2Next.Next
		outHead, outNext = insertNewNode(outHead, outNext, newNode)
	}

	return outHead
}

// func main() {
// 	head := new(ListNode)
// 	head.Val = 1
// 	head.Next = new(ListNode)
// 	head.Next.Val = 2
// 	head.Next.Next = new(ListNode)
// 	head.Next.Next.Val = 4

// 	head1 := new(ListNode)
// 	head1.Val = 1
// 	head1.Next = new(ListNode)
// 	head1.Next.Val = 3
// 	head1.Next.Next = new(ListNode)
// 	head1.Next.Next.Val = 5

// 	out := mergeTwoLists1(head, head1)

// 	printList(out)
// }
