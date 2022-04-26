package leetCodeProblems;

import java.util.*;

/**
 * LeetCode Problem Link - https://leetcode.com/problems/binary-tree-inorder-traversal/
 * 
 * InterviewBit - https://www.interviewbit.com/problems/inorder-traversal/
 * */


public class BinaryTreeInOrderTraversalIteration94 {
	
	public List<Integer> inorderTraversal(TreeNode A) {
		 
       ArrayList<Integer> outputList = new ArrayList<Integer>();
 
       Stack<TreeNode> inputStack = new Stack<TreeNode>();
 
       TreeNode current = A;
 
       while(!inputStack.isEmpty() || current != null) {
          
           if (current != null) {
               inputStack.push(current);
               current = current.left;
           }
           else {
               TreeNode temp = inputStack.pop();
               outputList.add(temp.val);
               current = temp.right;
           }
       }
      
       return outputList;
 
    }
	
	public static void main(String[] args) {
    	
    	TreeNode root = new TreeNode(1);
        root.left = new TreeNode(2);
        root.right = new TreeNode(3);
        root.left.left = new TreeNode(4);
        root.left.right = new TreeNode(5);
        root.right.left = new TreeNode(6);
        root.right.right = new TreeNode(7);
        
        BinaryTreeInOrderTraversalIteration94 ob = new BinaryTreeInOrderTraversalIteration94();
        
        List<Integer> output = ob.inorderTraversal(root);
        
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
