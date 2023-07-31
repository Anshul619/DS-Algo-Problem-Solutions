package leetCodeProblems.BinaryTree;

/**
 * LeetCode Problem Link - https://leetcode.com/problems/path-sum/submissions/
 *
 * InterviewBit - https://www.interviewbit.com/problems/path-sum/
 * */

public class BinaryTreeHasPathSum112 {

	public boolean hasPathSum(TreeNode root, int targetSum) {

	   if (root == null) {
           return false;
       }

       boolean ans = false;

       int subSum = targetSum - root.val;

       // This is leaf node
       if (subSum == 0 && root.left == null && root.right == null) {
           return true;
       }

       if (root.left != null) {
           ans = ans || hasPathSum(root.left, subSum);
       }

       if (root.right != null) {
           ans = ans || hasPathSum(root.right, subSum);
       }

       return ans;
   }

   public static void main(String[] args) {

    	TreeNode root = new TreeNode(1);
        root.left = new TreeNode(2);
        root.right = new TreeNode(3);
        root.left.left = new TreeNode(4);
        root.left.right = new TreeNode(5);
        root.right.left = new TreeNode(6);
        root.right.right = new TreeNode(7);

        BinaryTreeHasPathSum112 ob = new BinaryTreeHasPathSum112();

        boolean output = ob.hasPathSum(root, 8);

        System.out.println(output);

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
