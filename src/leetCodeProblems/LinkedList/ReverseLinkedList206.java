package leetCodeProblems.LinkedList;

/**
 * LeetCode - https://leetcode.com/problems/reverse-linked-list/solution/
 *
 * TimeComplexity - O(n)
 * SpaceComplexity - O(1)
 */
public class ReverseLinkedList {

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

        ListNode outputHead = null;

        while (head != null) {

            ListNode temp = new ListNode(head.val);

            if (outputHead != null) {
                temp.next = outputHead;
            }

            outputHead = temp;
            head = head.next;
        }

        return outputHead;
    }

    public static void main(String[] args) {

        ReverseLinkedList obj = new ReverseLinkedList();

        ListNode l1 = new ListNode(1);
        l1.next = new ListNode(2);
        l1.next.next = new ListNode(3);
        l1.next.next.next = new ListNode(4);
        l1.next.next.next.next = new ListNode(5);
        l1.next.next.next.next.next = new ListNode(6);

        ListNode output = obj.reverseList(l1);

        while(output != null) {
            System.out.println(output.val);
            output = output.next;
        }
    }
}
