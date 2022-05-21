package leetCodeProblems.BinaryTree;

/**
 * LeetCode - https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree-iii/
 */
public class BinaryTreeLCAIII1650 {

    public boolean recusivelyFindDownwards (Node node, int target){

        if (node == null) {
            return false;
        }

        if (node.val == target) {
            return true;
        }

        if (recusivelyFindDownwards(node.left, target) || recusivelyFindDownwards(node.right, target)) {
            return true;
        }

        return false;
    }

    public Node findOutSibling(Node node) {

        Node parentNode = node.parent;

        if (parentNode != null) {

            if (parentNode.left == node) {
                return parentNode.right;
            }
            else {
                return parentNode.left;
            }
        }

        return null;
    }

    public Node lowestCommonAncestor(Node p, Node q) {

        if (p.val == q.val) {
            return p;
        }

        if (recusivelyFindDownwards(p, q.val)) {
            return p;
        }

        while (p.parent != null) {

            if (p.parent.val == q.val) {
                return p.parent;
            }

            Node siblingNode = findOutSibling(p);

            if (siblingNode != null && recusivelyFindDownwards(siblingNode, q.val)) {
                 return p.parent;
            }

            p = p.parent;
        }

        return null;
    }

    public static void main(String[] args) {

        /*Node root = new Node(3);

        root.left = new Node(5);
        root.left.parent = root;

        root.right = new Node(1);
        root.right.parent = root;

        root.left.left = new Node(6);
        root.left.left.parent = root.left;

        root.left.right = new Node(2);
        root.left.right.parent = root.left;

        root.left.right.left = new Node(7);
        root.left.right.left.parent = root.left.right;

        root.left.right.right = new Node(4);
        root.left.right.right.parent = root.left.right;

        root.right.left = new Node(0);
        root.right.left.parent = root.right;

        root.right.right = new Node(8);
        root.right.right.parent = root.right;*/

        Node root = new Node(1);

        root.left = new Node(2);
        root.left.parent = root;

        root.right = new Node(3);
        root.right.parent = root;

        root.left.right = new Node(4);
        root.left.right.parent = root.left;


        BinaryTreeLCAIII1650 ob = new BinaryTreeLCAIII1650();

        Node output = ob.lowestCommonAncestor(root.left.right, root);

        System.out.println(output.val);

    }

    static class Node {

        Node left;
        Node right;
        int val;
        Node parent;

        Node(int val) {
            this.val = val;
        }
    }
}
