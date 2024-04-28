package main

/*
- Leetcode - https://leetcode.com/problems/add-two-numbers-ii/
- Time - O(n)
- Space - O(1)
*/

func getSize(l *ListNode) int {
	out := 0

	for l != nil {
		out++
		l = l.Next
	}
	return out
}

func addTwoNumbersUtil(l1, l2 *ListNode, diff int) (*ListNode, int) {

	if l1 == nil {
		return l2, 0
	}

	if l2 == nil {
		return l1, 0
	}

	var (
		child *ListNode
		carry int
		sum   int
	)

	if diff > 0 {
		child, carry = addTwoNumbersUtil(l1.Next, l2, diff-1)
		sum = l1.Val
	} else {
		child, carry = addTwoNumbersUtil(l1.Next, l2.Next, 0)
		sum = l1.Val + l2.Val
	}

	if carry > 0 {
		sum += carry
	}

	return &ListNode{child, sum % 10}, sum / 10

}

func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	size1 := getSize(l1)
	size2 := getSize(l2)

	diff := size1 - size2

	if size1 != size2 {
		if size1 < size2 {
			temp := l1
			l1 = l2
			l2 = temp

			diff = -diff
		}
	}

	out, carry := addTwoNumbersUtil(l1, l2, diff)

	if carry > 0 {
		out = &ListNode{out, carry}
	}

	return out
}

// func main() {
// 	head := new(ListNode)
// 	head.Val = 7
// 	head.Next = new(ListNode)
// 	head.Next.Val = 2
// 	head.Next.Next = new(ListNode)
// 	head.Next.Next.Val = 4
// 	head.Next.Next.Next = new(ListNode)
// 	head.Next.Next.Next.Val = 3

// 	head1 := new(ListNode)
// 	head1.Val = 5
// 	head1.Next = new(ListNode)
// 	head1.Next.Val = 6
// 	head1.Next.Next = new(ListNode)
// 	head1.Next.Next.Val = 4

// 	out := addTwoNumbers(head, head1)

// 	printLinkedList(out)

// 	head = new(ListNode)
// 	head.Val = 5
// 	head.Next = new(ListNode)
// 	head.Next.Val = 6
// 	head.Next.Next = new(ListNode)
// 	head.Next.Next.Val = 3

// 	head1 = new(ListNode)
// 	head1.Val = 8
// 	head1.Next = new(ListNode)
// 	head1.Next.Val = 4
// 	head1.Next.Next = new(ListNode)
// 	head1.Next.Next.Val = 2

// 	out = addTwoNumbers(head, head1)

// 	printLinkedList(out)
// }
