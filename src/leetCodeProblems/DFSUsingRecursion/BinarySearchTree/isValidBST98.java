package leetCodeProblems.BinarySearchTree;

/**
 * https://leetcode.com/problems/validate-binary-search-tree/
 * 
 * @author anshul.agrawal
 *
 */
public class BSTValidate98 {
	
	public boolean isValidBSTUtil(TreeNode root, Integer lowerRange, Integer upperRange) {
	    
        if (root == null) {
            return true;
        }
        
        if (lowerRange != null && root.val <= lowerRange) {
            return false;
        }
        
        if (upperRange != null && root.val >= upperRange) {
            return false;
        }
        
        boolean isLeftBST = isValidBSTUtil(root.left, lowerRange, root.val);
        boolean isRightBST = isValidBSTUtil(root.right, root.val, upperRange);
            
        if (!isLeftBST || !isRightBST) {
            return false;
        }
        
        return true;
        
    }
    
    public boolean isValidBST(TreeNode root) {
        return isValidBSTUtil(root, null, null);
    }
    
    public static void main(String[] args) {
        
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

        BSTValidate98 ob = new BSTValidate98();

        System.out.println(ob.isValidBST(root));
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
