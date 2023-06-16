package main

/*
- LeetCode - https://leetcode.com/problems/sort-list/description/
- Time - O(nlogn)
- Space - O(n)
*/
import "sort"

func sortListUsingExtraSpace(head *ListNode) *ListNode {
	slice := []int{}

	for head != nil {
		slice = append(slice, head.Val)
		head = head.Next
	}

	sort.Ints(slice)

	var (
		newHead *ListNode
		next    *ListNode
	)

	for _, v := range slice {
		temp := new(ListNode)
		temp.Val = v

		if newHead == nil {
			newHead = temp
			next = temp
		} else {
			next.Next = temp
			next = temp
		}
	}
	return newHead
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
// }
