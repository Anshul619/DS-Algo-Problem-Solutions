package leetCodeProblems.LinkedList;


/**
 * LeetCode - https://leetcode.com/problems/middle-of-the-linked-list/
 *
 * Single linked list find middle element without using streams ,set, collection try to do your own logic.
 *
 * Time Complexity - O(n)
 * Space Complexity - O(1)
 */
public class MiddleLinkedList876 {

    ListNode head; // head of list

    static class ListNode {

        int data;
        ListNode next;

        // Constructor
        ListNode(int d)
        {
            data = d;
            next = null;
        }
    }

    public ListNode middleNode(ListNode head) {

        ListNode slowPointer = head;
        ListNode fastPointer = head;

        while (fastPointer.next != null) {

            if (fastPointer.next.next != null) {
                slowPointer = slowPointer.next;
                fastPointer = fastPointer.next.next;
            }
            else {
                slowPointer = slowPointer.next;
                fastPointer = fastPointer.next;
            }

        }

        return slowPointer;
    }

    public static void main(String[] args) {

        ListNode l1 = new ListNode(1);
        l1.next = new ListNode(2);
        l1.next.next = new ListNode(3);
        l1.next.next.next = new ListNode(4);
        l1.next.next.next.next = new ListNode(5);
        l1.next.next.next.next.next = new ListNode(6);

        MiddleLinkedList876 obj = new MiddleLinkedList876();
        ListNode output = obj.middleNode(l1);

        System.out.println("Output value ->" + output.data);
    }
}
