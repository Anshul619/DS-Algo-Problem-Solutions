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
    
    private void bfsTraversal(TreeNode root) {
    	Queue<TreeVNode> queue = new LinkedList<TreeVNode>();
        
        TreeVNode nodeV = new TreeVNode(root, 0);
        queue.add(nodeV);
        
        TreeVNode currentVNode;
        TreeNode currentNode; 
        int currentNodeVOrder;
        
        TreeVNode currentLeftVNode;
        TreeVNode currentRightVNode;
        
        while (!queue.isEmpty()) {
            
            currentVNode = queue.remove();
            currentNode = currentVNode.node;
            currentNodeVOrder = currentVNode.vOrder;
            
            if ( currentNode != null) {
                
                int currentNodeVOrderArrayListIndex = currentNodeVOrder-min;
                result.get(currentNodeVOrderArrayListIndex).add(currentNode.val);
                
                if (currentNode.left != null) {
                	currentLeftVNode = new TreeVNode(currentNode.left, currentNodeVOrder-1);
                    queue.add(currentLeftVNode);
                }
                
                if (currentNode.right != null) {
                	currentRightVNode = new TreeVNode(currentNode.right, currentNodeVOrder+1);
                    queue.add(currentRightVNode);
                }
            }
        }
        
        return;
    }
    
    // Driver method
    public List<List<Integer>> verticalOrder(TreeNode root) {
        
        if (root == null) {
            return new ArrayList<>();    
        }
        
        findMinMax(root, 0);
        
        int numberOfVTraversalRows = max - min;
        
        for(int i = 0; i <= numberOfVTraversalRows; i++) {
            result.add(new ArrayList<Integer>());
        }
        
        bfsTraversal(root);
        
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
    
    static class TreeVNode {
        TreeNode node;
        int vOrder;
        
        TreeVNode(TreeNode node, int vOrder) {
            this.node = node;
            this.vOrder = vOrder;
        }
    }
    
    public static void main(String[] args) {
        //System.out.println("Hello World!");
        
        /*TreeNode root = new TreeNode(1);
        root.left = new TreeNode(2);
        root.right = new TreeNode(3);
        root.left.left = new TreeNode(4);
        root.left.right = new TreeNode(5);
        root.right.left = new TreeNode(6);
        root.right.right = new TreeNode(7);*/
        
        TreeNode root = new TreeNode(3);
        root.left = new TreeNode(9);
        root.right = new TreeNode(8);
        root.left.left = new TreeNode(4);
        root.left.right = new TreeNode(0);
        root.right.left = new TreeNode(1);
        root.right.right = new TreeNode(7);
        
        root.left.right.left = new TreeNode(5);
        root.left.right.right = new TreeNode(2);
        
        
        BinaryTreeVerticalOrderTraversal314 obj = new BinaryTreeVerticalOrderTraversal314();
        
        List<List<Integer>> result = obj.verticalOrder(root);
        
        for(int i=0; i < result.size(); i++) {
            System.out.println(result.get(i));
        }
        
    }
}
