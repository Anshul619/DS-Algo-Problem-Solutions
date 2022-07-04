package CodingInterviewQuestions;

import java.util.LinkedList;
import java.util.Queue;

/*

Flatten a multilevel Linked List, level by level, where each node has a 'next' and 'down' pointer.
- Reference - https://www.geeksforgeeks.org/flatten-a-linked-list-with-next-and-child-pointers/amp/
- Asked in Alation Coding Interview Round, 4thJuly2022.

Input:
1 - 2  -    3 - 4
|   |       |   |
5   6 - 8   9   12 - 13 - 14
    |   |
    7   10
        |
        11

Output:
1 - 2 - 3 - 4 - 5 - 6 - 8 - 9 - 12 - 13 - 14 - 7 - 10 - 11
*/
public class LevelOrderFlattenMultiLevelLinkedList {
    static class ListNode {
        ListNode down;
        ListNode next;

        int val;

        ListNode(int val) {
            this.val = val;
        }
    }

    ListNode outputHead;
    ListNode outputTail;

    Queue<ListNode> queue = new LinkedList();

    private void addNodeInOutput(ListNode temp) {

        if (outputHead == null) {
            outputHead = temp;
        }
        else {
            outputTail.next = temp;
        }
        outputTail = temp;
    }

    public ListNode flatten(ListNode head) {

        while(head != null) {

            addNodeInOutput(head);

            if (head.down != null) {
                queue.add(head.down);
            }

            head = head.next;
        }

        while(!queue.isEmpty()) {
            flatten(queue.remove());
        }

        return outputHead;
    }

    public static void main(String[] args) {

        LevelOrderFlattenMultiLevelLinkedList obj = new LevelOrderFlattenMultiLevelLinkedList();

        // Input =
        // 1 - 2  -    3 - 4
        // |   |
        // 5   6 - 8
        //     |
        //     7

        // Output = 1->2->3->4->5->6->8->7

        ListNode node = new ListNode(1);
        node.down = new ListNode(5);
        node.next = new ListNode(2);
        node.next.down = new ListNode(6);
        node.next.down.next = new ListNode(8);
        //node.next.down.next = new ListNode(8);
        node.next.down.down = new ListNode(7);
        node.next.next = new ListNode(3);
        node.next.next.next = new ListNode(4);

        ListNode output = obj.flatten(node);

        while(output != null) {
            System.out.println(output.val);
            output = output.next;
        }
    }
}