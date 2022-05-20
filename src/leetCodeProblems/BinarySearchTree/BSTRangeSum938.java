package leetCodeProblems.BinarySearchTree;

/**
 * https://leetcode.com/problems/range-sum-of-bst/
 */
public class BSTRangeSum938 {

    static class TreeNode {

        int val;
        TreeNode left;
        TreeNode right;

        TreeNode(int val) {
            this.val = val;
        }
    }

    public int rangeSumBST(TreeNode root, int low, int high) {

        if (root == null) {
            return 0;
        }

        int sum = rangeSumBST(root.left, low, high);

        if (low <= root.val && root.val <= high) {
            sum += root.val;
        }

        sum += rangeSumBST(root.right, low, high);

        return sum;
    }

    public static void main(String[] args) {

        BSTRangeSum938 ob = new BSTRangeSum938();

        /*TreeNode root = new TreeNode(10);
        root.left = new TreeNode(5);
        root.right = new TreeNode(15);
        root.left.left = new TreeNode(3);
        root.left.right = new TreeNode(7);
        root.right.right = new TreeNode(18);*/

        //int sum = ob.rangeSumBST(root, 7, 15);

        TreeNode root = new TreeNode(10);
        root.left = new TreeNode(5);
        root.right = new TreeNode(15);
        root.left.left = new TreeNode(3);
        root.left.left.left = new TreeNode(1);
        root.left.right = new TreeNode(7);
        root.left.right.left = new TreeNode(6);
        root.right.right = new TreeNode(18);

        System.out.println(ob.rangeSumBST(root, 6, 10));
    }
}
