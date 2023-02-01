package main

/**
- LeetCode - https://leetcode.com/problems/maximum-twin-sum-of-a-linked-list/submissions/
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSumUsingHashMap(head *ListNode) int {

	i := 0
	maxSum := 0

	nodesMap := make(map[int]int)

	for head != nil {
		nodesMap[i] = head.Val
		head = head.Next
		i++
	}

	for j := 0; j < i/2; j++ {

		twinNode := i - 1 - j

		if v, ok := nodesMap[twinNode]; ok {

			if maxSum < (v + nodesMap[j]) {
				maxSum = v + nodesMap[j]
			}
		}
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

// 	log.Println(pairSumUsingHashMap(head))

// }
