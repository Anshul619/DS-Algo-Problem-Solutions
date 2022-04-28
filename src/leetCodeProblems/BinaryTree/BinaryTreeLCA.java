package leetCodeProblems.BinaryTree;

public class BinaryTreeLCA {
	static boolean v1 = false;
    static boolean v2 = false;

    TreeNode findLCAUtil(TreeNode node, int n1, int n2) {

        if (node == null) {
            return null;
        }

        TreeNode temp = null;

        if (node.val == n1) {
            temp = node;
            v1 = true;
        }

        if (node.val == n2) {
            temp = node;
            v2 = true;
        }

        TreeNode left_lca = findLCAUtil(node.left, n1, n2);
        TreeNode right_lca = findLCAUtil(node.right, n1, n2);

        if (temp != null) {
            return temp;
        }

        // This is IMPORTANT code
        if (left_lca != null && right_lca != null) {
            return node;
        }

        if (left_lca != null) {
            return left_lca;
        }
        else {
            return right_lca;
        }
    }

    public int lca(TreeNode A, int B, int C) {
        v1 = v2 = false;

        TreeNode LCA = findLCAUtil(A, B, C);

        if (v1 && v2) {
            return LCA.val;
        }
        else {
            return -1;
        }
    }

}
