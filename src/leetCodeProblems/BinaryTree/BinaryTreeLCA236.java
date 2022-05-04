package leetCodeProblems.BinaryTree;

public class BinaryTreeLCA236 {
	static boolean isN1Found = false;
    static boolean isN2Found = false;
    
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
    
    // node is nothing BUT root of that recursive TREE.
    TreeNode findLCAUtil(TreeNode node, int n1, int n2) {

        if (node == null) {
            return null;
        }

        TreeNode temp = null;
        
        // Root matches with n1
        if (node.val == n1) {
            temp = node;
            isN1Found = true;
        }
        
        // Root matches with n2
        if (node.val == n2) {
            temp = node;
            isN2Found = true;
        }

        if (temp != null) {
            return temp;
        }
        
        // Either n1 or n2 exits as CHILDs in the left child of node
        TreeNode leftLCA = findLCAUtil(node.left, n1, n2); 
        TreeNode rightLCA = findLCAUtil(node.right, n1, n2);

        // LCA is found - This is IMPORTANT code
        if (leftLCA != null && rightLCA != null) {
            return node; // Note we are returning parent of leftLCA & rightLCA. NOT leftLCA or rightLCA.
        }

        if (leftLCA != null) {
            return leftLCA; // Return leftLCA, not PARENT.
        }
        else {
            return rightLCA;
        }
    }

    public int lca(TreeNode A, int B, int C) {
    	isN1Found = isN2Found = false;

        TreeNode LCA = findLCAUtil(A, B, C);

        if (isN1Found && isN2Found) {
            return LCA.val;
        }
        else {
            return -1;
        }
    }
    
    public static void main(String[] args) {
        
    	TreeNode root = new TreeNode(1);
        root.left = new TreeNode(2);
        root.right = new TreeNode(3);
        root.left.left = new TreeNode(4);
        root.left.right = new TreeNode(5);
        root.right.left = new TreeNode(6);
        root.right.right = new TreeNode(7);
        
        BinaryTreeLCA236 ob = new BinaryTreeLCA236();

        int output = ob.lca(root, 4, 7);
        
        System.out.println(output);

    }

}
