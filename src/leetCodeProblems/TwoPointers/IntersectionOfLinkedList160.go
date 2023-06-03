package main

/*
- LeetCode - https://leetcode.com/problems/intersection-of-two-linked-lists/description/
- Time - O(m+n)
- Space - O(1)
*/

type ListNode struct {
	Next *ListNode
	Val  int
}

func countNodes(head *ListNode) int {

	next, count := head, 0

	for next != nil {
		count++
		next = next.Next
	}

	return count
}

func getTargetNode(head *ListNode, target int) *ListNode {

	next := head

	for target > 0 {
		next = next.Next
		target--
	}

	return next
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {

	countA := countNodes(headA)
	countB := countNodes(headB)

	nextA, nextB := headA, headB

	if countA < countB {
		nextB = getTargetNode(headB, countB-countA)
	} else {
		nextA = getTargetNode(headA, countA-countB)
	}

	for nextA != nil && nextB != nil {
		if nextA == nextB {
			break
		}

		nextA = nextA.Next
		nextB = nextB.Next
	}

	return nextA
}

// func main() {

// 	// intersection := new(ListNode)
// 	// intersection.Val = 8
// 	// intersection.Next = new(ListNode)
// 	// intersection.Next.Val = 4
// 	// intersection.Next.Next = new(ListNode)
// 	// intersection.Next.Next.Val = 5

// 	// head := new(ListNode)
// 	// head.Val = 4
// 	// head.Next = new(ListNode)
// 	// head.Next.Val = 1
// 	// head.Next.Next = intersection

// 	// head1 := new(ListNode)
// 	// head1.Val = 5
// 	// head1.Next = new(ListNode)
// 	// head1.Next.Val = 6
// 	// head1.Next.Next = new(ListNode)
// 	// head1.Next.Next.Val = 1
// 	// head1.Next.Next.Next = intersection

// 	// head := new(ListNode)
// 	// head.Val = 1

// 	// head1 := head

// 	head := new(ListNode)
// 	head.Val = 1

// 	head1 := new(ListNode)
// 	head1.Val = 2

// 	out := getIntersectionNode(head, head1)

// 	if out != nil {
// 		log.Println("Intersection", out.Val)
// 	} else {
// 		log.Println("No Intersection")
// 	}

// }
