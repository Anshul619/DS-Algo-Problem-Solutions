package main

/*
- Leetcode - https://leetcode.com/problems/convert-binary-number-in-a-linked-list-to-integer/description/
- Time - O(n)
- Space - O(1)
*/
import (
	"math"
)

func getDecimalValueRecursively(head *ListNode, out *int) int {

	if head.Next == nil {
		*out += head.Val
		return 0
	}

	bit := getDecimalValueRecursively(head.Next, out) + 1
	*out += head.Val * int(math.Pow(float64(2), float64(bit)))
	return bit
}

func getDecimalValue(head *ListNode) int {
	out := 0
	getDecimalValueRecursively(head, &out)
	return out
}

// func main() {
// 	head := new(ListNode)
// 	head.Val = 1
// 	head.Next = new(ListNode)
// 	head.Next.Val = 0
// 	head.Next.Next = new(ListNode)
// 	head.Next.Next.Val = 1

// 	log.Println(getDecimalValue(head))

// 	head = new(ListNode)
// 	head.Val = 0

// 	log.Println(getDecimalValue(head))
// }
