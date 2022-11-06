package main

/*
- LeetCode - https://leetcode.com/problems/reorder-list/
- TimeComplexity - O(N)
- SpaceComplexity - O(1)
*/

func findMid(head *ListNode) *ListNode {
	mid := head
	temp := head

	for temp.Next != nil && temp.Next.Next != nil {
		mid = mid.Next
		temp = temp.Next.Next
	}

	return mid
}

func reverseList(mid *ListNode) {
	cur := mid.Next

	for cur != nil && cur.Next != nil {
		temp := mid.Next
		temp1 := cur.Next.Next

		mid.Next = cur.Next
		mid.Next.Next = temp

		cur.Next = temp1
	}
}

func reorderList(head *ListNode) {

	mid := findMid(head)

	reverseList(mid)

	cur := head

	for cur.Next != nil && cur.Next.Next != nil {
		temp := cur.Next
		temp1 := mid.Next.Next

		cur.Next = mid.Next
		cur.Next.Next = temp

		mid.Next = temp1

		cur = cur.Next.Next
	}

}

// func main() {

// 	head := new(ListNode)
// 	head.Val = 1
// 	//head.Next = n

// 	// head := new(ListNode)
// 	// head.Val = 1
// 	// head.Next = new(ListNode)
// 	// head.Next.Val = 2
// 	// head.Next.Next = new(ListNode)
// 	// head.Next.Next.Val = 3
// 	// head.Next.Next.Next = new(ListNode)
// 	// head.Next.Next.Next.Val = 4
// 	// head.Next.Next.Next.Next = new(ListNode)
// 	// head.Next.Next.Next.Next.Val = 5

// 	// head := new(ListNode)
// 	// head.Val = 1
// 	// head.Next = new(ListNode)
// 	// head.Next.Val = 2
// 	// head.Next.Next = new(ListNode)
// 	// head.Next.Next.Val = 3
// 	// head.Next.Next.Next = new(ListNode)
// 	// head.Next.Next.Next.Val = 4
// 	// head.Next.Next.Next.Next = new(ListNode)
// 	// head.Next.Next.Next.Next.Val = 5
// 	// head.Next.Next.Next.Next.Next = new(ListNode)
// 	// head.Next.Next.Next.Next.Next.Val = 6
// 	// head.Next.Next.Next.Next.Next.Next = new(ListNode)
// 	// head.Next.Next.Next.Next.Next.Next.Val = 7
// 	// head.Next.Next.Next.Next.Next.Next.Next = new(ListNode)
// 	// head.Next.Next.Next.Next.Next.Next.Next.Val = 8

// 	reorderList(head)

// 	/*head := new(ListNode)
// 	head.Val = 1

// 	output := removeNthFromEnd(head, 1)*/

// 	/*head := new(ListNode)
// 	head.Val = 1
// 	head.Next = new(ListNode)
// 	head.Next.Val = 2

// 	output := removeNthFromEnd(head, 2)*/

// 	temp := head

// 	for temp != nil {
// 		log.Println(temp.Val)
// 		temp = temp.Next
// 	}

// }
