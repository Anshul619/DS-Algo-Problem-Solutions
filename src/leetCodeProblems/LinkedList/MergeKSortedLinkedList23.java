package leetCodeProblems.LinkedList;
import java.util.*;

/**
 * LeetCode - https://leetcode.com/problems/merge-k-sorted-lists/
 *
 * TimeComplexity - O(n)
 * SpaceComplexity - O(n)
 *
 * Where n is the length of the max linked list in the input list.
 */
public class MergeKSortedLinkedList23
{
    static class ListNode {
        int val;
        ListNode next;

        ListNode(int val) {
            this.val = val;
        }
    }

    public ListNode merge (ArrayList<ListNode> input) {

        PriorityQueue<Integer> sortedQueue = new PriorityQueue();

        ListNode output = null;
        ListNode lastPointer = null;

        for (int i =0; i < input.size(); i++) {

            ListNode node = input.get(i);

            while(node != null) {
                sortedQueue.add(node.val);
                node = node.next;
            }
        }

        System.out.println("Queue ->" + sortedQueue);

        while (!sortedQueue.isEmpty()) {

            int value = sortedQueue.remove();

            ListNode temp = new ListNode(value);

            if (output != null) {
                lastPointer.next = temp;
                lastPointer = temp;
            }
            else {
                output = temp;
                lastPointer = temp;
            }
        }

        return output;
    }

    public static void main(String[] args) {

        //Input: lists = [[1,4,5],[1,3,4],[2,6]]

        ArrayList<ListNode> input = new ArrayList();

        ListNode list1 = new ListNode(1);
        list1.next = new ListNode(4);
        list1.next.next = new ListNode(5);
        input.add(list1);

        ListNode list2 = new ListNode(1);
        list2.next = new ListNode(3);
        list2.next.next = new ListNode(4);
        input.add(list2);

        ListNode list3 = new ListNode(2);
        list3.next = new ListNode(6);
        input.add(list3);

        MergeKSortedLinkedList23 obj = new MergeKSortedLinkedList23();
        ListNode out = obj.merge(input);

        while(out != null) {
            System.out.println("Number ->" + out.val);
            out = out.next;
        }

        //System.out.println("Hello World");
    }
}
