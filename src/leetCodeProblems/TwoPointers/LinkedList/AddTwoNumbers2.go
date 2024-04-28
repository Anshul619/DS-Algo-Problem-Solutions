package main

/*
- Leetcode - https://leetcode.com/problems/add-two-numbers/
- Time - O(n)
- Space - O(1)
*/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0

	var outHead, outNext *ListNode

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

		temp := &ListNode{nil, sum % 10}
		carry = sum / 10

		if outHead == nil {
			outHead = temp
			outNext = temp
		} else {
			outNext.Next = temp
			outNext = temp
		}
	}

	return outHead
}
