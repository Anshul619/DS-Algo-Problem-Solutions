package leetCodeProblems.BinaryTree;

import java.util.*;

/**
 * LeetCode Problem Link - https://leetcode.com/problems/binary-tree-inorder-traversal/
 *
 * InterviewBit - https://www.interviewbit.com/problems/inorder-traversal/
 * */


public class BinaryTreePreOrderTraversalIteration144 {

	public List<Integer> preorderTraversal(TreeNode root) {

        List<Integer> output = new ArrayList<Integer>();

        Stack<TreeNode> stackIn = new Stack<TreeNode>();

        TreeNode current = root;

        while(!stackIn.isEmpty() || current != null) {

            if (current != null) {
                output.add(current.val);
                stackIn.push(current);

                if (current.left != null) {
                    current = current.left;
                }
                else {
                    current = null;
                }

            }
            else {
                TreeNode temp = stackIn.pop();

                if (temp.right != null) {
                    current = temp.right;
                }
                else {
                    current = null;
                }

            }
        }

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

        BinaryTreePreOrderTraversalIteration144 ob = new BinaryTreePreOrderTraversalIteration144();

        List<Integer> output = ob.preorderTraversal(root);

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
