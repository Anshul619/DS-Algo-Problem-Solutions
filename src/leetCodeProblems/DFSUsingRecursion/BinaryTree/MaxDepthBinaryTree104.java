package leetCodeProblems.BinaryTree;

/**
 * LeetCode - https://leetcode.com/problems/maximum--of-binary-tree/
 */
public class BinaryTreeMax104 {

    public int maxUtil(TreeNode node) {
        if (node == null) {
            return 0;
        }

        System.out.println("Current node value -> " + node.val);
        int leftChildHeight = maxUtil(node.left);
        int rightChildHeight = maxUtil(node.right);

        System.out.println("LeftChildHeight => " + leftChildHeight);
        System.out.println("RightChildHeight => " + rightChildHeight);

        int currentNodeHeight = 1 + Math.max(leftChildHeight, rightChildHeight);

        System.out.println("Current node value -> " + node.val);
        System.out.println("currentNodeHeight => " + currentNodeHeight);

        System.out.println("-----");
        return currentNodeHeight;
    }

    public int max(TreeNode root) {
        return maxUtil(root);
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

        BinaryTreeMax104 ob = new BinaryTreeMax104();

        int max = ob.max(root);

        System.out.println(max);

    }
}
