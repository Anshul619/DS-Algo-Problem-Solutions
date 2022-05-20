package leetCodeProblems.BinaryTree;

import com.sun.source.tree.Tree;

/**
 * LeetCode - https://leetcode.com/problems/invert-binary-tree/
 *
 */
public class BinaryTreeInvert226 {

    public TreeNode invertTree(TreeNode root) {
        if ( root == null ) {
            return null;
        }

        TreeNode leftChild = invertTree(root.left);
        TreeNode rightChild = invertTree(root.right);

        root.right = leftChild;
        root.left = rightChild;

        return root;
    }

    void printInorder(TreeNode node)
    {
        if (node == null)
            return;

        /* first recur on left child */
        printInorder(node.left);

        /* then print the data of node */
        System.out.print(node.val + " ");

        /* now recur on right child */
        printInorder(node.right);
    }

    static class TreeNode {
        int val;
        TreeNode left;
        TreeNode right;

        TreeNode(int val) {
            this.val = val;
        }

        TreeNode(int val, TreeNode left, TreeNode right) {
            this.val = val;
            this.left = left;
            this.right = right;
        }
    }

    public static void main(String[] args) {

        /*TreeNode root = new TreeNode(4);
        root.left = new TreeNode(2);
        root.left.left = new TreeNode(1);
        root.left.right = new TreeNode(3);
        root.right = new TreeNode(7);
        root.right.left = new TreeNode(6);
        root.right.right = new TreeNode(9);*/


        TreeNode root = new TreeNode(1);
        root.left = new TreeNode(2);
        root.left.left = new TreeNode(5);
        root.left.right = new TreeNode(3);

        BinaryTreeInvert226 obj = new BinaryTreeInvert226();

        TreeNode invertedNode = obj.invertTree(root);
        obj.printInorder(invertedNode);

    }
}
