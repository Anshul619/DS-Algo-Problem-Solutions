package main

/*
- LeetCode - https://leetcode.com/problems/sort-list/description/
- Time - O(n)
- Space - O(1)
*/
func insertNode(nodeToInsert *ListNode, sortedListHead *ListNode) *ListNode {

	newNode := new(ListNode)
	newNode.Val = nodeToInsert.Val

	if sortedListHead == nil || sortedListHead.Val >= nodeToInsert.Val {
		newNode.Next = sortedListHead
		sortedListHead = newNode
	} else {

		sortedListNext := sortedListHead

		for sortedListNext.Next != nil && sortedListNext.Next.Val < nodeToInsert.Val {
			sortedListNext = sortedListNext.Next
		}

		newNode.Next = sortedListNext.Next
		sortedListNext.Next = newNode
	}

	return sortedListHead
}

func sortListUsingInsertionSort(head *ListNode) *ListNode {

	var sortedListHead *ListNode

	for head != nil {
		sortedListHead = insertNode(head, sortedListHead)
		head = head.Next
	}

	return sortedListHead
}

// func main() {
// 	// head := new(ListNode)
// 	// head.Val = 4
// 	// head.Next = new(ListNode)
// 	// head.Next.Val = 2
// 	// head.Next.Next = new(ListNode)
// 	// head.Next.Next.Val = 1
// 	// head.Next.Next.Next = new(ListNode)
// 	// head.Next.Next.Next.Val = 3

// 	// print(sortList(head))

// 	head := new(ListNode)
// 	head.Val = -1
// 	head.Next = new(ListNode)
// 	head.Next.Val = 5
// 	head.Next.Next = new(ListNode)
// 	head.Next.Next.Val = 3
// 	head.Next.Next.Next = new(ListNode)
// 	head.Next.Next.Next.Val = 4
// 	head.Next.Next.Next.Next = new(ListNode)
// 	head.Next.Next.Next.Next.Val = 0

// 	print(sortList(head))

// 	// head := new(ListNode)
// 	// head.Val = 1
// 	// head.Next = new(ListNode)
// 	// head.Next.Val = 3
// 	// head.Next.Next = new(ListNode)
// 	// head.Next.Next.Val = 3
// 	// head.Next.Next.Next = new(ListNode)
// 	// head.Next.Next.Next.Val = 1
// 	// head.Next.Next.Next.Next = new(ListNode)
// 	// head.Next.Next.Next.Next.Val = 3

// 	// print(sortList(head))
// }
