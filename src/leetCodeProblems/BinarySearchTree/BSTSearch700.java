package leetCodeProblems.BinarySearchTree;

/**
 * LeetCode - https://leetcode.com/problems/search-in-a-binary-search-tree/
 */
public class BSTSearch66 {
    static class TreeNode {

        int val;
        TreeNode left;
        TreeNode right;

        TreeNode(int val) {
            this.val = val;
        }
    }

    public TreeNode searchBST(TreeNode root, int val) {

        if (root == null) {
            return null;
        }

        if (root.val == val) {
            return root;
        }

        if (root.val > val) {
            return searchBST(root.left, val);
        }
        else {
            return searchBST(root.right, val);
        }
    }

    public static void main(String[] args) {

        TreeNode root = new TreeNode(5);
        root.left = new TreeNode(1);
        root.right = new TreeNode(7);
        root.right.left = new TreeNode(6);
        root.right.right = new TreeNode(8); // Return false

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

        BSTSearch66 ob = new BSTSearch66();

        TreeNode outputNode = ob.searchBST(root, 8);

        if (outputNode != null) {
            System.out.println(outputNode.val);
        }

    }
}
