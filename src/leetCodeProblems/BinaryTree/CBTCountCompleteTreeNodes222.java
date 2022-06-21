package leetCodeProblems.BinaryTree;

import com.sun.source.tree.Tree;

/**
 * LeetCode - https://leetcode.com/problems/count-complete-tree-nodes/
 *
 * Approach
 * -
 */
public class CBTCountCompleteTreeNodes222 {

    static class TreeNode {
        int val;
        TreeNode left;
        TreeNode right;

        TreeNode(int val) {
            this.val = val;
        }
    }

    private int calculateDepth(TreeNode root) {

        TreeNode node = root;
        int depth = 0;

        while(node != null) {
            depth++;
            node = node.left;
        }

        return depth;
    }

    public boolean exists(int indexToCheck, int depth, TreeNode node) {

        int left = 1;
        int right = (int) Math.pow(2, depth-1);

        int mid;

        for(int i=0; i < depth; i++) {

            mid = (left+right)/2;

            System.out.println("Mid ->" + mid);
            System.out.println("indexToCheck ->" + indexToCheck);

            if (indexToCheck <= mid) {
                node = node.left;
                right = mid;
            }
            else {
                node = node.right;
                left = mid+1;
            }
        }

        if (node != null) {
            System.out.println("Index to check ->"+ indexToCheck);
            return true;
        }
        return false;
    }

    public int countNodes(TreeNode root) {

        int depth = calculateDepth(root);

        int left = 1;
        int right = (int) Math.pow(2, depth-1);

        //System.out.println(right);

        while(left < right) {

            int mid = (left + right)/2;

            if (exists(mid, depth, root)) {
                left = mid;
            }
            else {
                right = mid;
            }
        }

        //System.out.println("Depth ->" + depth);
        int nodes = 0;

        for(int i=0; i < depth-1; i++) {
            //System.out.println(Math.pow(2, i));
            nodes += Math.pow(2, i);
        }

        //System.out.println(nodes);
        nodes += left;

        return nodes;

    }

    public static void main(String[] args) {

        TreeNode root = new TreeNode(1);
        root.left = new TreeNode(2);
        root.left.left = new TreeNode(4);
        root.left.left.left = new TreeNode(1);
        root.left.left.right = new TreeNode(7);
        root.left.right = new TreeNode(5);
        root.left.right.left = new TreeNode(9);
        root.left.right.right = new TreeNode(3);
        root.right = new TreeNode(3);
        root.right.left = new TreeNode(6);
        root.right.left.left = new TreeNode(4);
        root.right.left.right = new TreeNode(7);
        root.right.right = new TreeNode(7);

        CBTCountCompleteTreeNodes222 obj = new CBTCountCompleteTreeNodes222();

        System.out.println(obj.countNodes(root));

    }
}
