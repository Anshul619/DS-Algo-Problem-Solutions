package leetCodeProblems.BinaryTree;

/**
 * LeetCode - https://leetcode.com/problems/find-leaves-of-binary-tree/
 */

import java.util.ArrayList;
import java.util.List;

public class BinaryTreeFindLeaves366 {

    static class TreeNode {
        int val;
        TreeNode left;
        TreeNode right;
        TreeNode(int val) {
            this.val = val;
        }
    }

    List<List<Integer>> output = new ArrayList<List<Integer>>();

    public void addInArrayList(int nodeVal, int level) {

        if (output.size() <= level) {

            ArrayList<Integer> temp = new ArrayList<Integer>();
            temp.add(nodeVal);

            output.add(temp);
        }
        else {
            output.get(level).add(nodeVal);
        }

        return;
    }

    public int findLeavesUtils(TreeNode node) {

        if (node == null) {
            return -1;
        }
        // This is leaf node
        if (node.left == null && node.right == null) {
            return 0;
        }

        int leftChildLevel = findLeavesUtils(node.left);

        if (leftChildLevel != -1) {
            addInArrayList(node.left.val, leftChildLevel);
        }

        int rightChildLevel = findLeavesUtils(node.right);

        if (rightChildLevel != -1) {
            addInArrayList(node.right.val, rightChildLevel);
        }

        return Math.max(leftChildLevel, rightChildLevel) + 1;

    }

    public List<List<Integer>> findLeaves(TreeNode root) {
        output = new ArrayList<List<Integer>>();

        if (root == null) {
            return output;
        }

        int rootLevel = findLeavesUtils(root);

        addInArrayList(root.val, rootLevel);

        return output;

    }

    public static void main(String[] args) {

        /*TreeNode root = new TreeNode(1);
        root.left = new TreeNode(2);
        root.right = new TreeNode(3);
        root.left.left = new TreeNode(4);
        root.left.right = new TreeNode(5);
        root.right.left = new TreeNode(6);
        root.right.right = new TreeNode(7);*/

        /*TreeNode root = new TreeNode(1);
        root.left = new TreeNode(2);
        root.right = new TreeNode(3);
        root.left.left = new TreeNode(4);
        root.left.right = new TreeNode(5);*/

        TreeNode root = new TreeNode(1);
        root.left = new TreeNode(2);
        root.left.left = new TreeNode(3);

        BinaryTreeFindLeaves366 ob = new BinaryTreeFindLeaves366();

        List<List<Integer>> output = ob.findLeaves(root);

        System.out.println(output);

    }
}
