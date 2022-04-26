package leetCodeProblems;/*
    Problem -
    Solution -
    Time Complexity -
    Space Complexity -
 */


/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode() {}
 *     TreeNode(int val) { this.val = val; }
 *     TreeNode(int val, TreeNode left, TreeNode right) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */

import java.util.*;

class BinaryTreePostOrderTraversalRecursion145 {
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
}
