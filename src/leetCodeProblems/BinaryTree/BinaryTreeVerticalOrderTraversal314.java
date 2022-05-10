package leetCodeProblems.BinaryTree;

/**
 * LeetCode - https://leetcode.com/problems/binary-tree-vertical-order-traversal/
 * 
 */
import java.util.*;

public class BinaryTreeVerticalOrderTraversal314 {

	int min = 0;
    int max = 0;
    
    public List<List<Integer>> result = new ArrayList<List<Integer>>();
    
    private void findMinMax(TreeNode node, int currentVerticalOrder) {
        
        if (node == null) {
            return;
        }
        
        min = Math.min(min, currentVerticalOrder);
        max = Math.max(max, currentVerticalOrder);
        
        findMinMax(node.left, currentVerticalOrder-1);
        findMinMax(node.right, currentVerticalOrder+1);
    }
    
    private void verticalOrderTraversalRecursion(TreeNode node, int currentVerticalOrder, int targetVerticalOrder) {
        if (node == null) {
            return;
        }
        
        if (currentVerticalOrder == targetVerticalOrder) {
            int targetVerticalOrderIndex = targetVerticalOrder-min;
            
            System.out.println("targetVerticalOrder ->"+ targetVerticalOrderIndex);
            List<Integer> temp = result.get(targetVerticalOrderIndex);
            temp.add(node.val);
        }
            
        verticalOrderTraversalRecursion(node.left, currentVerticalOrder-1, targetVerticalOrder);
        verticalOrderTraversalRecursion(node.right, currentVerticalOrder+1, targetVerticalOrder);
        
    }
    
    public List<List<Integer>> verticalOrder(TreeNode root) {
        
        findMinMax(root, 0);
        
        int numberOfVTraversalRows = max - min;
            
        System.out.println("numberOfVTraversalRows -> " + numberOfVTraversalRows);
        for(int i = 0; i <= numberOfVTraversalRows; i++) {
            result.add(new ArrayList<Integer>());
        }
        //obj.print();
        
        for (int i=min; i <= max; i++) {
            verticalOrderTraversalRecursion(root, 0, i);
        }
        
        return result;
        
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
        //System.out.println("Hello World!");
        
        TreeNode root = new TreeNode(1);
        root.left = new TreeNode(2);
        root.right = new TreeNode(3);
        root.left.left = new TreeNode(4);
        root.left.right = new TreeNode(5);
        root.right.left = new TreeNode(6);
        root.right.right = new TreeNode(7);
        
        BinaryTreeVerticalOrderTraversal314 obj = new BinaryTreeVerticalOrderTraversal314();
        
        List<List<Integer>> result = obj.verticalOrder(root);
        
        for(int i=0; i < result.size(); i++) {
            System.out.println(result.get(i));
        }
        
    }
}
