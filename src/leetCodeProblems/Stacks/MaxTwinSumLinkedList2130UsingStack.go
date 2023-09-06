package main

/**
- LeetCode - https://leetcode.com/problems/maximum-twin-sum-of-a-linked-list/submissions/
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

type NodesStack []int

func (s *NodesStack) pop() int {
	out := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return out
}

func (s *NodesStack) push(elem int) {
	*s = append(*s, elem)
}

func pairSum(head *ListNode) int {

	maxSum := 0

	nodesStack := new(NodesStack)

	slow := head
	fast := head

	// Traverse the first half & push elements to stack
	for slow != nil {

		nodesStack.push(slow.Val)
		slow = slow.Next

		if fast.Next.Next != nil {
			fast = fast.Next.Next
		} else {
			break
		}
	}

	// Traverse second half, pop elements from stack & compare max sum
	for slow != nil {

		sum := nodesStack.pop() + slow.Val

		if maxSum < sum {
			maxSum = sum
		}

		slow = slow.Next
	}

	return maxSum
}

// func main() {

// 	head := new(ListNode)
// 	head.Val = 5
// 	head.Next = new(ListNode)
// 	head.Next.Val = 4
// 	head.Next.Next = new(ListNode)
// 	head.Next.Next.Val = 2
// 	head.Next.Next.Next = new(ListNode)
// 	head.Next.Next.Next.Val = 1

// 	log.Println(pairSum(head))

// }
