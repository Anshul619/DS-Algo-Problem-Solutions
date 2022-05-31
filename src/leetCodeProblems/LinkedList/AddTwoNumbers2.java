package leetCodeProblems.LinkedList;

/**
 * LeetCode - https://leetcode.com/problems/add-two-numbers/
 * Time Complexity - O(n)
 * Space Complexity - O(1)
 * */

public class AddTwoNumbers2 {

    static class ListNode {
        int val;
        ListNode next;
        ListNode() {}
        ListNode(int val) { this.val = val; }
        ListNode(int val, ListNode next) { this.val = val; this.next = next; }
    }

    private ListNode startPointer = null;

    private void addTwoNumbersUtils(ListNode l1, ListNode l2, ListNode out, int carry) {

        if (l1 == null && l2 == null) {

            if (carry != 0) {

                ListNode newNode = new ListNode(carry);

                if (out == null) {
                    startPointer = newNode;
                }
                else {
                    out.next = newNode;
                }
            }
            return;
        }

        int firstVal = 0;
        int secondVal = 0;
        int finalVal = 0;

        ListNode firstNextNode = null;
        ListNode secondNextNode = null;

        if (l1 != null) {

            firstVal = l1.val;
            firstNextNode = l1.next;
            //System.out.println("firstVal ->" + firstVal);
        }

        if (l2 != null) {

            secondVal = l2.val;
            secondNextNode = l2.next;
            //System.out.println("secondVal ->" + secondVal);
        }

        finalVal = firstVal + secondVal + carry;

        //System.out.println("finalVal ->" + finalVal);
        ListNode newNode = new ListNode();
        newNode.val = finalVal%10;

        if (finalVal != 0) {
            carry = finalVal / 10;
        }

        if (out == null) {
            //out = newNode;
            startPointer = newNode;
            System.out.println(newNode.val);
        }
        else {
            out.next = newNode;
        }

        addTwoNumbersUtils(firstNextNode, secondNextNode, newNode, carry);
    }

    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {

        addTwoNumbersUtils(l1, l2, null, 0);

        return startPointer;
    }

    public static void printList(ListNode currNode)
    {
        System.out.print("LinkedList: ");

        // Traverse through the LinkedList
        while (currNode != null) {
            // Print the data at current node
            System.out.print(currNode.val + " ");

            // Go to next node
            currNode = currNode.next;
        }
    }

    public static void main(String[] args) {
        ListNode l1 = new ListNode(9);
        l1.next = new ListNode(9);
        l1.next.next = new ListNode(9);
        l1.next.next.next = new ListNode(9);
        l1.next.next.next.next = new ListNode(9);
        l1.next.next.next.next.next = new ListNode(9);
        l1.next.next.next.next.next.next = new ListNode(9);

        ListNode l2 = new ListNode(9);
        l2.next = new ListNode(9);
        l2.next.next = new ListNode(9);
        l2.next.next.next = new ListNode(9);

        AddTwoNumbers2 obj = new AddTwoNumbers2();
        ListNode output = obj.addTwoNumbers(l1, l2);

        obj.printList(output);

    }
}
