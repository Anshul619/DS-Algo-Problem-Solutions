package main

/*
- Leetcode - https://leetcode.com/problems/palindrome-linked-list/
- Time - O(n)
- Space - O(n)
*/
func isPalindrome1(head *ListNode) bool {
	list := []int{}

	for head != nil {
		list = append(list, head.Val)
		head = head.Next
	}

	start, end := 0, len(list)-1

	for start < end {
		if list[start] != list[end] {
			return false
		}
		start++
		end--
	}

	return true
}

// func main() {
// 	head := new(ListNode)
// 	head.Val = 1
// 	head.Next = new(ListNode)
// 	head.Next.Val = 2
// 	head.Next.Next = new(ListNode)
// 	head.Next.Next.Val = 2
// 	head.Next.Next.Next = new(ListNode)
// 	head.Next.Next.Next.Val = 1

// 	log.Println(isPalindrome1(head))

// 	head = new(ListNode)
// 	head.Val = 1
// 	head.Next = new(ListNode)
// 	head.Next.Val = 2
// 	log.Println(isPalindrome1(head))
// }
