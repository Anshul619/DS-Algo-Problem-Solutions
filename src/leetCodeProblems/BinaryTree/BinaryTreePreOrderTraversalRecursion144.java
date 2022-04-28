package leetCodeProblems.BinaryTree;

import java.util.ArrayList;
import java.util.List;

/**
 * LeetCode Problem Link - https://leetcode.com/problems/binary-tree-preorder-traversal/submissions/
 *
 * InterviewBit - <TBD>
 * */

public class BinaryTreePreOrderTraversalRecursion144 {

	void recursion(TreeNode node, List<Integer> output) {

        if (node == null) {
            return;
        }

        output.add(node.val);

        if (node.left != null) {
            recursion(node.left, output);
        }

        if (node.right != null) {
            recursion(node.right, output);
        }
    }

    public List<Integer> preorderTraversal(TreeNode root) {
        List<Integer> output = new ArrayList<Integer>();

        recursion(root, output);

        return output;
    }

    public static void main(String[] args) {

    	TreeNode root = new TreeNode(1);
        root.left = new TreeNode(2);
        root.right = new TreeNode(3);
        root.left.left = new TreeNode(4);
        root.left.right = new TreeNode(5);
        root.right.left = new TreeNode(6);
        root.right.right = new TreeNode(7);

        BinaryTreePreOrderTraversalRecursion144 ob = new BinaryTreePreOrderTraversalRecursion144();

        List<Integer> output = ob.postorderTraversal(root);

        System.out.println(output);

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

}
