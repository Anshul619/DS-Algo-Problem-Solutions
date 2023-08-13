package main

/*
- Leetcode - https://leetcode.com/problems/add-two-numbers/
- Time - O(n)
- Space - O(1)
*/

func appendNodeToList(outHead, outNext *ListNode, val int) (*ListNode, *ListNode) {
	temp := new(ListNode)
	temp.Val = val

	if outHead == nil {
		outHead, outNext = temp, temp
	} else {
		outNext.Next = temp
		outNext = temp
	}

	return outHead, outNext
}

func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0

	var (
		outHead, outNext *ListNode
	)

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

		outHead, outNext = appendNodeToList(outHead, outNext, sum%10)
	}

	return outHead
}

// func main() {
// 	head := new(ListNode)
// 	head.Val = 2
// 	head.Next = new(ListNode)
// 	head.Next.Val = 4
// 	head.Next.Next = new(ListNode)
// 	head.Next.Next.Val = 3

// 	head1 := new(ListNode)
// 	head1.Val = 5
// 	head1.Next = new(ListNode)
// 	head1.Next.Val = 6
// 	head1.Next.Next = new(ListNode)
// 	head1.Next.Next.Val = 4

// 	out := addTwoNumbers1(head, head1)

// 	printLinkedList(out)
// }
