package leetCodeProblems.LinkedList;

/**
 * LeetCode - https://leetcode.com/problems/palindrome-linked-list/submissions/
 *
 * TimeComplexity - O(n)
 * SpaceComplexity - O(n)
 */

import java.util.ArrayList;

public class PalindromeLinkedList234 {
    static class ListNode {
        int val;
        ListNode next;
        ListNode() {}
        ListNode(int val) { this.val = val; }
        ListNode(int val, ListNode next) { this.val = val; this.next = next; }
    }

    public boolean isPalindrome(ListNode head) {

        if (head == null) {
            return false;
        }

        ListNode temp = head;
        ArrayList nodes = new ArrayList();

        while(temp != null) {
            nodes.add(temp.val);
            temp = temp.next;
        }

        //System.out.println(nodes);

        int left = 0;
        int right = nodes.size()-1;

        while (left < right) {

            if (nodes.get(left) != nodes.get(right)) {
                return false;
            }

            left++;
            right--;
        }

        return true;
    }

    public static void main(String[] args) {

        ListNode l1 = new ListNode(1);
        l1.next = new ListNode(2);
        //l1.next.next = new ListNode(2);
        //l1.next.next.next = new ListNode(1);
        //l1.next.next.next.next = new ListNode(5);
        //l1.next.next.next.next.next = new ListNode(6);

        PalindromeLinkedList234 obj = new PalindromeLinkedList234();

        System.out.println(obj.isPalindrome(l1));
    }
}
