package main

/*
- Leetcode - https://leetcode.com/problems/middle-of-the-linked-list/
- Time - O(n)
- Space - O(1)
*/
func middleNode(head *ListNode) *ListNode {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

// func main() {
// 	head := new(ListNode)
// 	head.Val = 1
// 	head.Next = new(ListNode)
// 	head.Next.Val = 2
// 	head.Next.Next = new(ListNode)
// 	head.Next.Next.Val = 3
// 	head.Next.Next.Next = new(ListNode)
// 	head.Next.Next.Next.Val = 4
// 	head.Next.Next.Next.Next = new(ListNode)
// 	head.Next.Next.Next.Next.Val = 5
// 	head.Next.Next.Next.Next.Next = new(ListNode)
// 	head.Next.Next.Next.Next.Next.Val = 6
// 	log.Println(middleNode(head))
// }
