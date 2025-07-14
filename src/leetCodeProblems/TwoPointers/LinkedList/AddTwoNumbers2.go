package main

/*
- Leetcode - https://leetcode.com/problems/add-two-numbers/
- Time - O(n)
- Space - O(1)
*/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	var tail *ListNode

	carry := 0

	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10

		newNode := &ListNode{}
		newNode.Val = sum % 10

		if head != nil {
			tail.Next = newNode
		} else {
			head = newNode
		}
		tail = newNode
	}

	return head
}
