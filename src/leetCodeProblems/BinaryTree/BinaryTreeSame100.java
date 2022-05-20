package leetCodeProblems.BinaryTree;

public class BinaryTreeSame100 {

    public boolean isSameTree(TreeNode p, TreeNode q) {

        if (p == null && q == null) {
            return true;
        }

        if (p == null || q == null) {
            return false;
        }

        boolean isLeftChildSame = isSameTree(p.left, q.left);
        boolean isRightChildSame = isSameTree(p.right, q.right);

        if (p.val == q.val && isLeftChildSame && isRightChildSame) {
            return true;
        }

        return false;
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

        /*TreeNode root = new TreeNode(1);
        root.left = new TreeNode(2);
        root.right = new TreeNode(3);

        TreeNode root1 = new TreeNode(1);
        root1.left = new TreeNode(2);
        root1.right = new TreeNode(3);*/

        TreeNode root = new TreeNode(1);
        root.left = new TreeNode(2);

        TreeNode root1 = new TreeNode(1);
        root1.right = new TreeNode(2);

        BinaryTreeSame100 obj = new BinaryTreeSame100();

        System.out.println(obj.isSameTree(root, root1));
    }
}
