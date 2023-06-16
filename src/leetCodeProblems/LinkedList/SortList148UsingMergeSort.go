package main

/*
- LeetCode - https://leetcode.com/problems/sort-list/description/
- Time - O(n*log n)
- Space - O(n)
*/

func addNode(head *ListNode, next *ListNode, val int) (*ListNode, *ListNode) {
	temp := new(ListNode)
	temp.Val = val

	if next == nil {
		return temp, temp
	} else {
		next.Next = temp
		return head, temp
	}
}

func merge(list1 *ListNode, list2 *ListNode) *ListNode {

	var outHead, next *ListNode

	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			outHead, next = addNode(outHead, next, list2.Val)
			list2 = list2.Next
		} else {
			outHead, next = addNode(outHead, next, list1.Val)
			list1 = list1.Next
		}
	}

	for list1 != nil {
		outHead, next = addNode(outHead, next, list1.Val)
		list1 = list1.Next
	}

	for list2 != nil {
		outHead, next = addNode(outHead, next, list2.Val)
		list2 = list2.Next
	}

	return outHead
}

func getMiddle(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	slow, fast := head, head.Next

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func sortList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	mid := getMiddle(head)
	list2 := mid.Next
	mid.Next = nil
	head = sortList(head)
	list2 = sortList(list2)

	return merge(head, list2)
}

func main() {
	// head := new(ListNode)
	// head.Val = 4
	// head.Next = new(ListNode)
	// head.Next.Val = 2
	// head.Next.Next = new(ListNode)
	// head.Next.Next.Val = 1
	// head.Next.Next.Next = new(ListNode)
	// head.Next.Next.Next.Val = 3

	// print(sortList(head))

	// head := new(ListNode)
	// head.Val = -1
	// head.Next = new(ListNode)
	// head.Next.Val = 5
	// head.Next.Next = new(ListNode)
	// head.Next.Next.Val = 3
	// head.Next.Next.Next = new(ListNode)
	// head.Next.Next.Next.Val = 4
	// head.Next.Next.Next.Next = new(ListNode)
	// head.Next.Next.Next.Next.Val = 0

	// print(sortList(head))

	// head := new(ListNode)
	// head.Val = -1
	// head.Next = new(ListNode)
	// head.Next.Val = 5

	// print(sortList(head))

	head := new(ListNode)
	head.Val = 1
	head.Next = new(ListNode)
	head.Next.Val = 3
	head.Next.Next = new(ListNode)
	head.Next.Next.Val = 3
	head.Next.Next.Next = new(ListNode)
	head.Next.Next.Next.Val = 1
	head.Next.Next.Next.Next = new(ListNode)
	head.Next.Next.Next.Next.Val = 3

	print(sortList(head))

	//print(sortList(nil))
}
