package leetCodeProblems.LinkedList;

/**
 * LeetCode - https://leetcode.com/problems/reverse-linked-list/solution/
 */
public class ReverseLinkedListUsingRecursion {

    static class ListNode {

        int val;
        ListNode next;

        // Constructor
        ListNode(int d)
        {
            val = d;
            next = null;
        }
    }

    public ListNode reverseList(ListNode head) {

        if (head == null || head.next == null) {
            return head;
        }

        ListNode p = reverseList(head.next);
        head.next.next = head;
        head.next = null;
        return p;

    }
}
