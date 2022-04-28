package leetCodeProblems.BinaryTree;

import java.util.*;

/**
 * LeetCode Problem Link - https://leetcode.com/problems/path-sum-ii/
 *
 * InterviewBit - https://www.interviewbit.com/problems/root-to-leaf-paths-with-sum/

 * */

public class BinaryTreeAllRootToLeafPathsWithSum113 {

	List<List<Integer>> result = new ArrayList<>();

    public void dfs(TreeNode node, int targetSum, ArrayList<Integer> path) {

        if (node == null) {
            return;
        }

        int subSum =  targetSum - node.val;

        path.add(node.val);

        if (subSum == 0 && node.left == null && node.right == null) {
            result.add(path);
        }

        dfs(node.left, subSum, new ArrayList<>(path));
        dfs(node.right, subSum, new ArrayList<>(path));
    }

    public List<List<Integer>> pathSum(TreeNode root, int targetSum) {

        if (root == null) {
            return result;
        }

        dfs(root, targetSum, new ArrayList<Integer>());

        return result;
    }

	public static void main(String[] args) {

    	TreeNode root = new TreeNode(1);
        root.left = new TreeNode(2);
        root.right = new TreeNode(3);
        root.left.left = new TreeNode(4);
        root.left.right = new TreeNode(5);
        root.right.left = new TreeNode(6);
        root.right.right = new TreeNode(7);

        BinaryTreeAllRootToLeafPathsWithSum113 ob = new BinaryTreeAllRootToLeafPathsWithSum113();

        List<List<Integer>> output = ob.pathSum(root, 8);

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
