package leetCodeProblems.BinaryTree;

import java.util.*;

import leetCodeProblems.BinaryTree.BinaryTreeConstructFromInOrderPreOrder105.TreeNode;

public class BinaryTreeConstructFromInOrderPostOrder106 {
	static HashMap<Integer, Integer> inHashMap = new HashMap<Integer, Integer>();
    
    public TreeNode buildTreeRecursion(int[] inorder, int[] postorder, int inStart, 
                                        int inEnd, int postStart, int postEnd) {
        
        if (inStart > inEnd) {
            return null;    
        }
        
        if (postStart > postEnd) {
            return null;
        }
        
        //System.out.println("PostIndex - " + postIndex);
        
        int postOrderValue = postorder[postEnd];
        
        TreeNode node = new TreeNode(postOrderValue);
        
        // There is no left or right childs of this node.
        if (inStart == inEnd) {
            return node;
        }
        
        int inorderIndex = inHashMap.get(postOrderValue);
        
        int inLeftStart = inStart;
        int inLeftEnd = inorderIndex-1;
        int lengthOfInorderLeft = inLeftEnd-inLeftStart +1; //Since inLeftStart, inLeftEnd are ZERO based, hence incrementing it by 1 is needed.
        
        int postLeftStart = postStart;
        int postLeftEnd = postStart + lengthOfInorderLeft - 1; //Since lengthOfInorderLeft is NON-Zero based, hence decrement by 1 is needed.
        
        node.left = buildTreeRecursion(inorder, postorder, 
                                       inLeftStart, inLeftEnd, 
                                       postLeftStart, postLeftEnd);
        
        int inRightStart = inorderIndex+1;
        int inRightEnd = inEnd;
        int lengthOfInorderRight = inRightEnd-inRightStart+1; // Since inRightStart, inRightEnd are ZERO based, hence incrementing it by 1 is needed.
        
        //System.out.println("lengthOfInorderRight - " + lengthOfInorderRight);
        
        int postRightStart = postEnd - lengthOfInorderRight;
        int postRightEnd = postEnd-1;
        
        node.right = buildTreeRecursion(inorder, postorder, 
                                        inRightStart, inRightEnd,
                                        postRightStart, postRightEnd);
        
        return node;
    }
    
    public TreeNode buildTree(int[] inorder, int[] postorder) {
        
        for(int i=0; i < inorder.length; i++) {
            inHashMap.put(inorder[i], i);
        }
        
        return buildTreeRecursion(inorder, postorder, 0, inorder.length-1, 0, postorder.length-1);
    }
    
    /* This function is here just to test buildTree() */
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
    
    public static void main(String[] args) {
        
        int in[] = new int[] {4, 2, 5, 1, 6, 3, 7};
        int post[] = new int[] { 4, 5, 2, 6, 7, 3, 1};
        
        
        BinaryTreeConstructFromInOrderPostOrder106 ob = new BinaryTreeConstructFromInOrderPostOrder106();

        TreeNode output = ob.buildTree(in, post);
        
        ob.printInorder(output);

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
