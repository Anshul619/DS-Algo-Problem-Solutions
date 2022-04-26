package leetCodeProblems;

import java.util.ArrayList;
import java.util.List;

/**
 * LeetCode Problem Link - https://leetcode.com/problems/binary-tree-postorder-traversal/
 * */

public class BinaryTreePostOrderTraversalRecursion145 {
	
	void recursion(TreeNode node, List<Integer> output) {
        
        if (node == null) {
            return;
        }
        
        if (node.left != null) {
            recursion(node.left, output);
        }
        
        if (node.right != null) {
            recursion(node.right, output);
        }    
        
        output.add(node.val);
    }
    
    public List<Integer> postorderTraversal(TreeNode root) {
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
        
        BinaryTreePostOrderTraversalRecursion145 ob = new BinaryTreePostOrderTraversalRecursion145();
        
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
