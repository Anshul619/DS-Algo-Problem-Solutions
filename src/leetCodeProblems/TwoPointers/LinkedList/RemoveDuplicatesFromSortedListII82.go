package main

/*
- Leetcode - https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/description
- Time - O(n)
- Space - O(1)
*/
func deleteNode(p, prev, head *ListNode) *ListNode {
	if prev != nil {
		prev.Next = p.Next
	} else {
		head = p.Next
	}
	return head
}

func deleteDuplicatesII(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var prev *ListNode
	removedVal := 101

	p := head

	for p != nil {

		if p.Next != nil && p.Val == p.Next.Val {
			head = deleteNode(p.Next, prev, head)
			removedVal = p.Val
			p = p.Next.Next

			continue
		}

		if p.Val == removedVal {
			head = deleteNode(p, prev, head)
			p = p.Next

			continue
		}

		prev = p
		p = p.Next
	}
	return head
}
