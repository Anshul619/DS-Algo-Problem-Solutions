package main

/*
- LeetCode - https://leetcode.com/problems/odd-even-linked-list/submissions/
- TimeComplexity - O(n)
- SpaceComplexity - O(1)
*/
//import "log"

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func oddEvenList(head *ListNode) *ListNode {

	cur, fast, counter := head, head, 1

	for fast != nil {

		for i := 1; i <= counter; i++ {
			fast = fast.Next
		}

		if fast == nil || fast.Next == nil {
			break
		}

		swappedNode := new(ListNode)
		swappedNode.Val = fast.Next.Val

		if fast.Next.Next != nil {
			fast.Next = fast.Next.Next
		} else {
			fast.Next = nil
		}

		temp := cur.Next
		swappedNode.Next = temp
		cur.Next = swappedNode

		cur = cur.Next
		fast = cur
		counter++
	}

	return head
}

// func print(node *ListNode) {
// 	for node != nil {
// 		log.Println(node.Val)
// 		node = node.Next
// 	}
// }
// func main() {

// 	//head := ListNode{
// 	//	Val: 10,
// 	//	Next: &ListNode{
// 	//		Val:  10,
// 	//		Next: new(ListNode),
// 	//	},
// 	//}

// 	//head := new(ListNode)
// 	//head.Val = 1
// 	//head.Next = new(ListNode)
// 	//head.Next.Val = 2
// 	//head.Next.Next = new(ListNode)
// 	//head.Next.Next.Val = 3

// 	//head := new(ListNode)
// 	//head.Val = 1
// 	//head.Next = new(ListNode)
// 	//head.Next.Val = 2
// 	//head.Next.Next = new(ListNode)
// 	//head.Next.Next.Val = 3
// 	//head.Next.Next.Next = new(ListNode)
// 	//head.Next.Next.Next.Val = 4
// 	//head.Next.Next.Next.Next = new(ListNode)
// 	//head.Next.Next.Next.Next.Val = 5

// 	head := new(ListNode)
// 	head.Val = 1

// 	print(oddEvenList(head))
// }
