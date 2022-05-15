package leetCodeProblems.BinarySearchTree;

/**
 * LeetCode - https://leetcode.com/problems/recover-binary-search-tree/solution/
 * 
 * @author anshul.agrawal
 *
 */
public class BSTRecover99UsingInorderTraversal {
	
	TreeNode firstIncorrectNode = null;
	TreeNode lastIncorrectNode = null;
	TreeNode predessor = null;
	
	/**
	 * Swap the two nodes
	 * 
	 * @param a
	 * @param b
	 */
	public void swap(TreeNode a, TreeNode b) {
	    int tmp = a.val;
	    a.val = b.val;
	    b.val = tmp;
	}
	
	/**
	 * Find the nodes which we need to swap
	 * 
	 * @param root
	 */
	public void findTwoSwapped(TreeNode root) {
		
		if (root == null) {
			return;
		}
		
		findTwoSwapped(root.left);
    
		if (predessor != null && root.val <= predessor.val) {
			
			lastIncorrectNode = root;
			
			if (firstIncorrectNode == null) {
				firstIncorrectNode = predessor;
			}	
		}
	
		predessor = root;
		
		System.out.println("predessor -> " + predessor.val);
		
		findTwoSwapped(root.right);
	}
    
	public void recoverTree(TreeNode root) {
		findTwoSwapped(root);
        
		if (firstIncorrectNode != null & lastIncorrectNode != null) {
        	
        	System.out.println("firstIncorrectNode val ->" + firstIncorrectNode.val);
            System.out.println("lastIncorrectNode val ->" + lastIncorrectNode.val);
        	swap(firstIncorrectNode, lastIncorrectNode);
        }
	}
	
	public static void main(String[] args) {
        
		TreeNode root = new TreeNode(1);
        root.left = new TreeNode(3);
        root.left.right = new TreeNode(2);
        
        BSTRecover99UsingInorderTraversal ob = new BSTRecover99UsingInorderTraversal();
        
        ob.recoverTree(root);
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
