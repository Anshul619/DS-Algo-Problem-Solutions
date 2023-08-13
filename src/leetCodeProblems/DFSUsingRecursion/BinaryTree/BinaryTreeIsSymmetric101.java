package leetCodeProblems.BinaryTree;

/**
 * LeetCode - https://leetcode.com/problems/symmetric-tree/
 */
public class BinaryTreeIsSymmetric {

    static class TreeNode {
        int val;

        TreeNode left;

        TreeNode right;

        TreeNode(int val) {
            this.val = val;
        }
    }

    private boolean isSymmetricUtil(TreeNode node, TreeNode symmetricNode) {

        if (node == null && symmetricNode == null) {
            return true;
        }

        if (node == null || symmetricNode == null || node.val != symmetricNode.val) {
            return false;
        }

        boolean isLeftChildSymmetric = isSymmetricUtil(node.left, symmetricNode.right);
        boolean isRightChildSymmetric = isSymmetricUtil(node.right, symmetricNode.left);

        if (!isLeftChildSymmetric || !isRightChildSymmetric) {
            return false;
        }

        return true;
    }

    public boolean isSymmetric(TreeNode root) {

        if ( root == null ) {
            return true;
        }

        return isSymmetricUtil(root.left, root.right);
    }

    public static void main(String[] args) {

        /*TreeNode root = new TreeNode(1);
        root.left = new TreeNode(2);
        root.right = new TreeNode(2);
        root.left.left = new TreeNode(3);
        root.left.right = new TreeNode(4);
        root.right.left = new TreeNode(4);
        root.right.right = new TreeNode(3); // Expected O/P = true*/

        /*TreeNode root = new TreeNode(1);
        root.left = new TreeNode(2);
        root.left.right = new TreeNode(3);
        root.right = new TreeNode(2);
        root.right.right = new TreeNode(3); // Expected O/P = false*/

        TreeNode root = new TreeNode(1);
        root.right = new TreeNode(2); // Expected O/P = false

        BinaryTreeIsSymmetric obj = new BinaryTreeIsSymmetric();
        System.out.println(obj.isSymmetric(root));

    }
}
