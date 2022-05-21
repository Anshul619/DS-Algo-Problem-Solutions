package leetCodeProblems.BinaryTree;

/**
 * LeetCode - https://leetcode.com/problems/flatten-binary-tree-to-linked-list/
 */

public class BinaryTreeToLinkedList114 {

    static class TreeNode {

        int val;
        TreeNode left;
        TreeNode right;

        TreeNode(int val) {
            this.val = val;
        }
    }

    public TreeNode flattenUtils(TreeNode node) {

        if (node == null) {
            return null;
        }

        TreeNode leftChildStartPointer = node.left;
        TreeNode leftChildLastPointer = flattenUtils(node.left);

        TreeNode rightChildStartPointer = node.right;
        TreeNode rightChildLastPointer  = flattenUtils(node.right);

        if (leftChildStartPointer != null) {
            node.right = leftChildStartPointer;
        }

        node.left = null;

        if (leftChildLastPointer != null && rightChildStartPointer != null) {
            leftChildLastPointer.right = rightChildStartPointer;
        }

        if (rightChildLastPointer != null) {
            return rightChildLastPointer;
        }

        if (leftChildLastPointer != null) {
            return leftChildLastPointer;
        }

        return node;
    }

    public TreeNode flatten(TreeNode root) {

        if (root == null) {
            return null;
        }

        flattenUtils(root);

        return root;

    }

    void printList(TreeNode node)
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
        BinaryTreeToLinkedList114 obj = new BinaryTreeToLinkedList114();

        // Let us create the tree shown in above diagram
        /*TreeNode root = new TreeNode(1);
        root.left = new TreeNode(2);
        root.right = new TreeNode(5);
        root.left.left = new TreeNode(3);
        root.left.right = new TreeNode(4);
        root.right.right = new TreeNode(6);
        root.right.right.right = new TreeNode(7);*/

        /*TreeNode root = new TreeNode(1);
        root.left = new TreeNode(2);
        root.left.left = new TreeNode(3);*/

        TreeNode root = new TreeNode(1);
        root.left = new TreeNode(2);
        root.left.left = new TreeNode(3);
        root.left.right = new TreeNode(4);
        root.left.left.left = new TreeNode(5);

        // Convert to DLL
        TreeNode head = obj.flatten(root);

        // Print the converted list
        obj.printList(head);
    }
}
