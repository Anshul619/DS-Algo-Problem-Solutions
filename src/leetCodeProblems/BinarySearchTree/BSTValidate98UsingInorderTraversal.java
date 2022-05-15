package leetCodeProblems.BinarySearchTree;

/**
 * LeetCode - https://leetcode.com/problems/recover-binary-search-tree/solution/
 * 
 * @author anshul.agrawal
 *
 */
public class BSTValidate98UsingInorderTraversal {
	
	TreeNode predessor = null;
	
	/**
	 * Check if the BST is valid or not
	 * 
	 * @param root
	 */
	public boolean isValidBSTUtil(TreeNode root) {
		
		if (root == null) {
			return true;
		}
		
		boolean isLeftBST = isValidBSTUtil(root.left);
    
		if (!isLeftBST) {
			return false;
		}
		
		if (predessor != null && root.val <= predessor.val) {
			return false;
		}
	
		predessor = root;
		
		System.out.println("predessor -> " + predessor.val);
		
		boolean isRightBST = isValidBSTUtil(root.right);
		
		if (!isRightBST) {
			return false;
		}
		
		return true;
	}
    
	public boolean isValidBST(TreeNode root) {
		return isValidBSTUtil(root);
	}
	
	public static void main(String[] args) {
        
		/*TreeNode root = new TreeNode(1);
        root.left = new TreeNode(3);
        root.left.right = new TreeNode(2);*/
        
        TreeNode root = new TreeNode(5);
        root.left = new TreeNode(1);
        root.right = new TreeNode(4);
        root.right.left = new TreeNode(3);
        root.right.right = new TreeNode(6); // Return false
        
        /*TreeNode root = new TreeNode(5);
        root.left = new TreeNode(4);
        root.right = new TreeNode(6);
        root.right.left = new TreeNode(3);
        root.right.right = new TreeNode(7); // Return false*/
        
    
        /*TreeNode root = new TreeNode(2);
        root.left = new TreeNode(1);
        root.right = new TreeNode(3); // Return true*/
        
        /*TreeNode root = new TreeNode(22);
        root.left = new TreeNode(-4);
        root.left.left = new TreeNode(-60);
        root.left.left.left = new TreeNode(-90);*/
        
        /*TreeNode root = new TreeNode(2);
        root.left = new TreeNode(2);
        root.right = new TreeNode(2); // Return false*/
        
        BSTValidate98UsingInorderTraversal obj = new BSTValidate98UsingInorderTraversal();
        
        System.out.println(obj.isValidBST(root));
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
