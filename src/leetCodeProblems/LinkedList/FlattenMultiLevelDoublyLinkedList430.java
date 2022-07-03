package leetCodeProblems.LinkedList;

/**
 * LeetCode - https://leetcode.com/problems/flatten-a-multilevel-doubly-linked-list/submissions/
 *
 * TimeComplexity - O(N)
 * SpaceComplexity - O(N)
 */
public class FlattenMultiLevelDoublyLinkedList430 {

    static class Node {
        public int val;
        public Node prev;
        public Node next;
        public Node child;

    }

    Node outputHead;
    Node outputLast;

    private void addNodeInOutputList(int val) {

        Node temp = new Node();
        temp.val = val;

        if (outputLast != null) {
            temp.prev = outputLast;
            outputLast.next = temp;
            outputLast = temp;
        }
        else {
            outputHead = temp;
            outputLast = temp;
        }
    }

    public Node flatten(Node head) {

        if (head == null) {
            return null;
        }

        while (head != null) {

            addNodeInOutputList(head.val);

            if (head.child != null) {
                flatten(head.child);
            }

            head = head.next;

        }

        return outputHead;
    }
}
