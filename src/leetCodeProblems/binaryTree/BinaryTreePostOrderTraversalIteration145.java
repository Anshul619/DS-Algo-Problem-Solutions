package leetCodeProblems.binaryTree;

/**
 * LeetCode Problem Link - https://leetcode.com/problems/binary-tree-postorder-traversal/
 * */

import java.util.*;

class BinaryTreePostOrderTraversalIteration145 {

    public List<Integer> postorderTraversal(TreeNode root) {

    	List<Integer> output = new ArrayList<Integer>();

        Stack<TreeNode> stack1 = new Stack<TreeNode>();
        Stack<TreeNode> stack2 = new Stack<TreeNode>();

        if (root != null) {
            stack1.push(root);
        }

        while(!stack1.isEmpty()) {

        	TreeNode temp = stack1.pop();
            stack2.push(temp);

            if (temp.left != null) {
                stack1.push(temp.left);
            }

            if (temp.right != null) {
                stack1.push(temp.right);
            }
        }

        while(!stack2.isEmpty()) {

        	TreeNode temp = stack2.pop();
        	output.add(temp.val);

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

        BinaryTreePostOrderTraversalIteration145 ob = new BinaryTreePostOrderTraversalIteration145();

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

