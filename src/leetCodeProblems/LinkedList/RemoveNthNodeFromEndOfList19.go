package main

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	previous := head
	current := head
	fast := head

	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	for fast != nil {
		previous = current
		current = current.Next
		fast = fast.Next
	}

	if previous == current {
		head = current.Next
	} else if previous.Next != nil {
		previous.Next = previous.Next.Next
	} else {
		head = nil
	}

	return head
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

// 	output := removeNthFromEnd(head, 2)

// 	/*head := new(ListNode)
// 	head.Val = 1

// 	output := removeNthFromEnd(head, 1)*/

// 	/*head := new(ListNode)
// 	head.Val = 1
// 	head.Next = new(ListNode)
// 	head.Next.Val = 2

// 	output := removeNthFromEnd(head, 2)*/

// 	for output != nil {
// 		log.Println(output.Val)
// 		output = output.Next
// 	}

// }
