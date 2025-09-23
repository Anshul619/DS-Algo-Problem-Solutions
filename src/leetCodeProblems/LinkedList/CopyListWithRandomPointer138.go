package main

/*
- LeetCode - https://leetcode.com/problems/copy-list-with-random-pointer/description/
- Time - O(n)
- Space - O(n)
*/

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	// Map from original node -> copied node
	m := make(map[*Node]*Node)

	dummyHead := &Node{}
	prev := dummyHead

	oldNext := head

	// First pass: copy nodes (Val only), set up Next links later
	for oldNext != nil {
		node := &Node{}
		node.Val = oldNext.Val

		m[oldNext] = node

		prev.Next = node
		prev = prev.Next

		oldNext = oldNext.Next
	}

	oldNext = head

	for oldNext != nil {
		m[oldNext].Random = m[oldNext.Random]
		m[oldNext].Next = m[oldNext.Next]

		oldNext = oldNext.Next
	}

	return dummyHead.Next
}

// func main() {
// 	head := new(Node)
// 	head.Val = 7
// 	head.Next = new(Node)

// 	head.Next.Val = 13
// 	head.Next.Next = new(Node)
// 	head.Next.Next.Val = 11
// 	head.Next.Next.Next = new(Node)
// 	head.Next.Next.Next.Val = 10
// 	head.Next.Next.Next.Next = new(Node)
// 	head.Next.Next.Next.Next.Val = 1

// 	head.Random = nil
// 	head.Next.Random = head
// 	head.Next.Next.Random = head.Next.Next.Next.Next

// 	newHead := copyRandomList(head)

// 	next := newHead

// 	for next != nil {
// 		log.Println(next.Val)
// 		log.Println(next.Random)
// 		next = next.Next
// 	}
// }
