package leetCodeProblems.BackTracking;

/**
 * LeetCode - https://leetcode.com/problems/all-nodes-distance-k-in-binary-tree/
 * InterviewBit - https://www.interviewbit.com/problems/nodes-at-distance-k/
 * 
 */
import java.util.*;

public class BinaryTreeNodesAtKDistance863 {
	
	public ArrayList<Integer> output = new ArrayList<Integer>();
    
    private void downwardsRecursion(TreeNode node, int targetDistance) {
        
        if (node == null) {
            return;
        }
        
        System.out.println("downwardsRecursion - node ->" + node.val + ", targetDistance ->" + targetDistance);
        
        if (targetDistance == 0) {
            System.out.println("downwardsRecursion - targetDistance is MATCHED at node ->" + node.val);
            output.add(node.val);
            return;
        }
        
        downwardsRecursion(node.left, targetDistance-1);
        downwardsRecursion(node.right, targetDistance-1);
        
    }
    
    public int distanceKRecursively(TreeNode node, TreeNode targetNode, int distance, int targetDistance) {
        
        if (node == null) {
            return -1;
        }
        
        System.out.println("distanceKRecursively - Current node -> " + node.val);
        
        // Target node is FOUND
        if (node == targetNode) {
            System.out.println("distanceKRecursively - Target node is FOUND, targetDistance -> " + targetDistance);
            downwardsRecursion(node, targetDistance);
            return 0;
        }
        
        int distanceOfCurrentNodeWithTargetNode = 0;
        
        // find distance of currentNode with the targetNode.
        int distanceOfLeftChildWithTargetNode = distanceKRecursively(node.left, targetNode, distance, targetDistance);
        
        System.out.println("distanceKRecursively - distanceOfLeftChildWithTargetNode -> " + distanceOfLeftChildWithTargetNode);
        
        // targetNode is in LEFT subtree
        if (distanceOfLeftChildWithTargetNode != -1) {
                    
            distanceOfCurrentNodeWithTargetNode = distanceOfLeftChildWithTargetNode+1;
            
            if (distanceOfCurrentNodeWithTargetNode == targetDistance) {
                output.add(node.val);
            }
            downwardsRecursion(node.right, targetDistance-distanceOfCurrentNodeWithTargetNode-1);
            
            return distanceOfCurrentNodeWithTargetNode;
        }
        
        int distanceOfRightChildWithTargetNode = distanceKRecursively(node.right, targetNode, distance, targetDistance);
    
        // Distance b/w current node's right child & the target node is greater than 0. Hence target node is surely in left child. 
        // Now go downwards on the right child, with corresponding distance.
        if (distanceOfRightChildWithTargetNode != -1) {
            
            distanceOfCurrentNodeWithTargetNode = distanceOfRightChildWithTargetNode+1;
            
            if (distanceOfCurrentNodeWithTargetNode == targetDistance) {
                output.add(node.val);
            }
            
            downwardsRecursion(node.left, targetDistance-distanceOfCurrentNodeWithTargetNode-1);
            
            return distanceOfCurrentNodeWithTargetNode;
        }
        
        return -1;
        
    }
    
    public List<Integer> distanceK(TreeNode root, TreeNode target, int k) {
        
        distanceKRecursively(root, target, k, k);
        
        return output;
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
        
        /*TreeNode root = new TreeNode(20);
        root.left = new TreeNode(8);
        root.right = new TreeNode(22);
        root.left.left = new TreeNode(4);
        root.left.right = new TreeNode(12);
        root.left.right.left = new TreeNode(10);
        root.left.right.right = new TreeNode(14);
        
        TreeNode target = root.left;
        
        Main obj = new Main();
        obj.distanceKRecursively(root, target, 2, 2); // O/P = 10, 14, 22*/
        
        TreeNode root = new TreeNode(0);
        root.right = new TreeNode(1);
        root.right.right = new TreeNode(2);
        root.right.right.right = new TreeNode(3);
        
        TreeNode target = root.right;
    
        BinaryTreeNodesAtKDistance863 obj = new BinaryTreeNodesAtKDistance863();
        obj.distanceKRecursively(root, target, 2, 2); // O/P = 3
        
        /*TreeNode root = new TreeNode(0);
        root.left = new TreeNode(1);
        root.left.left = new TreeNode(2);
        root.left.left.left = new TreeNode(3);
        
        TreeNode target = root.left;
        
        Main obj = new Main();
        obj.distanceKRecursively(root, target, 2, 2); // O/P = 3*/
        
        /*TreeNode root = new TreeNode(0);
        root.left = new TreeNode(1);
        root.left.left = new TreeNode(3);
        root.left.right = new TreeNode(2);
        
        TreeNode target = root.left.right;
    
        Main obj = new Main();
        obj.distanceKRecursively(root, target, 1, 1);*/
        
        
        System.out.println(obj.output);
    }

}
