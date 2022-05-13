package leetCodeProblems.BinaryTree;

import java.util.*;

public class BinaryTreeConstructFromInOrderPreOrder105 {
	
	static int preIndex = 0;
	static int preOrderLength = 0;
	
    static HashMap<Integer, Integer> inHashMap = new HashMap<Integer, Integer>();
    
    public TreeNode buildTreeRecursion(int[] preorder, int inStart, int inEnd) {
        
        if (inStart > inEnd) {
            return null;    
        }
        
        if (preIndex >= preOrderLength) {
            return null;
        }
        
        //System.out.println("Preindex - " + preIndex);
        
        int preOrderValue = preorder[preIndex];
        
        TreeNode node = new TreeNode(preOrderValue);
        preIndex += 1;
        
        // There is no left or right childs of this node.
        if (inStart == inEnd) {
            return node;
        }
        
        int inorderIndex = inHashMap.get(preOrderValue);
        
        node.left = buildTreeRecursion(preorder, inStart, inorderIndex-1);
        node.right = buildTreeRecursion(preorder, inorderIndex+1, inEnd);
        
        return node;
    }
    
    public TreeNode buildTree(int[] preorder, int[] inorder) {
        
        for(int i=0; i < inorder.length; i++) {
            inHashMap.put(inorder[i], i);
        }
        
        preIndex = 0;
        preOrderLength = preorder.length;
        
        return buildTreeRecursion(preorder, 0, inorder.length-1);
    }
    
    /**
     * This function is here just to test buildTree() 
     **/
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
        int pre[] = new int[] { 1, 2, 4, 5, 3, 6, 7};
        
        
        BinaryTreeConstructFromInOrderPreOrder105 ob = new BinaryTreeConstructFromInOrderPreOrder105();

        TreeNode output = ob.buildTree(pre, in);
        
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
