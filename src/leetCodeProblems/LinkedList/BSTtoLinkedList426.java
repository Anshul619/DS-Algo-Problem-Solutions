package leetCodeProblems.LinkedList;

/**
 * LeetCode - https://leetcode.com/problems/convert-binary-search-tree-to-sorted-doubly-linked-list/solution/
 */

// Note - You can think of the left and right pointers as synonymous to the predecessor and successor pointers in a doubly-linked list.

public class BSTtoLinkedList426 {

    static class Node {

        int val;
        Node left;
        Node right;

        Node(int val) {
            this.val = val;
        }
    }

    Node firstPointer = null; // first element of linked list
    Node lastPointer = null; // last element of linked list

    public void treeToDoublyListUtils(Node node) {

        if (node == null) {
            return;
        }

        treeToDoublyListUtils(node.left);

        if (lastPointer != null) {
            //node.right = lastPointer;
            lastPointer.right = node;
            node.left = lastPointer;
        }
        else {
            firstPointer = node;
        }

        lastPointer = node;

        treeToDoublyListUtils(node.right);

        return;
    }

    public Node treeToDoublyList(Node root) {

        if (root == null) {
            return null;
        }

        firstPointer = null;
        lastPointer = null;
        treeToDoublyListUtils(root);

        firstPointer.left = lastPointer;
        lastPointer.right = firstPointer;

        return firstPointer;

    }

    void printList(Node node)
    {
        while (node != null)
        {
            System.out.print(node.val + " ");
            node = node.right;
        }
    }

    /* Driver program to test above functions*/
    public static void main(String[] args)
    {
        BSTtoLinkedList426 obj = new BSTtoLinkedList426();

        // Let us create the tree shown in above diagram
        Node root = new Node(4);
        root.left = new Node(2);
        root.right = new Node(5);
        root.left.left = new Node(1);
        root.left.right = new Node(3);

        // Convert to DLL
        Node head = obj.treeToDoublyList(root);

        // Print the converted list
        obj.printList(head);
    }
}
