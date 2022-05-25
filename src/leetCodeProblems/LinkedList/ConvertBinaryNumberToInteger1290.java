package leetCodeProblems.LinkedList;

/**
 * LeetCode - https://leetcode.com/problems/convert-binary-number-in-a-linked-list-to-integer/
 */

public class ConvertBinaryNumberToInteger1290 {

    static class ListNode {
        int val;
        ListNode next;
        ListNode() {}
        ListNode(int val) { this.val = val; }
        ListNode(int val, ListNode next) { this.val = val; this.next = next; }
    }

    static int output = 0;

    private static int getNumberUtil(ListNode node) {

        if (node.next == null) {
            output += node.val*1; // 2^0 = 1
            return 0;
        }

        int numberPower = getNumberUtil(node.next);

        numberPower++;

        output += node.val*Math.pow(2, numberPower); // 2^0 = 1

        return numberPower;

    }


    public int getDecimalValue(ListNode head) {
        output = 0;

        getNumberUtil(head);

        return output;
    }

    public static void main(String[] args) {

        /**ListNode root = new ListNode(1);
        root.next = new ListNode(0);
        root.next.next = new ListNode(1);**/

        ListNode root = new ListNode(0);

        ConvertBinaryNumberToInteger1290 obj = new ConvertBinaryNumberToInteger1290();
        System.out.println(obj.getDecimalValue(root));
    }
}
