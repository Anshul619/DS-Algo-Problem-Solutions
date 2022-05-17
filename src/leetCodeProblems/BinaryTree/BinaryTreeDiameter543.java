package leetCodeProblems.BinaryTree;

/**
 * LeetCode - https://leetcode.com/problems/diameter-of-binary-tree/
 */

public class BinaryTreeDiameter543 {

    public int depth(TreeNode node) {

        if (node == null) {
            return 0;
        }

        //System.out.println("Current node value -> " + node.val);
        int leftChildHeight = depth(node.left);
        int rightChildHeight = depth(node.right);

        //System.out.println("LeftChildHeight => " + leftChildHeight);
        //System.out.println("RightChildHeight => " + rightChildHeight);

        int currentNodeHeight = 1 + Math.max(leftChildHeight, rightChildHeight);

        //System.out.println("Current node value -> " + node.val);
        //System.out.println("currentNodeHeight => " + currentNodeHeight);

        //System.out.println("-----");
        return currentNodeHeight;
    }

    public int diameter(TreeNode node) {

        if (node == null) {
            return 0;
        }

        //int height = depth(node) - 1;
        int leftChildHeight = depth(node.left) - 1;
        int rightChildHeight = depth(node.right) -1;

        int leftChildDiameter = diameter(node.left);
        int rightChildDiameter = diameter(node.right);

        return Math.max(leftChildHeight + rightChildHeight + 2, Math.max(leftChildDiameter, rightChildDiameter));
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

        BinaryTreeDiameter543 ob = new BinaryTreeDiameter543();

        int diameter = ob.diameter(root);

        System.out.println(diameter);

    }
}
