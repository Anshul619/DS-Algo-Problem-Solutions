package leetCodeProblems.BinaryTree;

/**
 * LeetCode - https://leetcode.com/problems/maximum-depth-of-binary-tree/
 */
public class BinaryTreeMaxDepth104 {

    public int maxDepthUtil(TreeNode node) {
        if (node == null) {
            return 0;
        }

        System.out.println("Current node value -> " + node.val);
        int leftChildHeight = maxDepthUtil(node.left);
        int rightChildHeight = maxDepthUtil(node.right);

        System.out.println("LeftChildHeight => " + leftChildHeight);
        System.out.println("RightChildHeight => " + rightChildHeight);

        int currentNodeHeight = 1 + Math.max(leftChildHeight, rightChildHeight);

        System.out.println("Current node value -> " + node.val);
        System.out.println("currentNodeHeight => " + currentNodeHeight);

        System.out.println("-----");
        return currentNodeHeight;
    }

    public int maxDepth(TreeNode root) {
        return maxDepthUtil(root);
    }

    static class TreeNode {

        int val;
        TreeNode left;
        TreeNode right;

        TreeNode(int val) {
            this.val = val;
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

        BinaryTreeMaxDepth104 ob = new BinaryTreeMaxDepth104();

        int maxDepth = ob.maxDepth(root);

        System.out.println(maxDepth);

    }
}
