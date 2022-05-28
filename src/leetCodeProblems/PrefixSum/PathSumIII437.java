package leetCodeProblems.PrefixSum;

/**
 * LeetCode - https://leetcode.com/problems/path-sum-iii/submissions/
 */

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;

public class PathSumIII437 {

    static class TreeNode {
        int val;
        TreeNode left;
        TreeNode right;
        TreeNode(int val) {
            this.val = val;
        }
    }

    int ansCount = 0;
    int targetSum = 0;

    HashMap<Integer, Integer> pathSumMap;

    public void pathSumUtil(TreeNode node, int pathSum) {

        if (node == null) {
            return;
        }

        pathSum += node.val;

        int prefixSum = pathSum-targetSum;

        if (pathSumMap.containsKey(prefixSum)) {
            ansCount += pathSumMap.get(prefixSum);
        }

        if (pathSumMap.containsKey(pathSum)) {
            pathSumMap.put(pathSum, pathSumMap.get(pathSum) +1);
        }
        else {
            pathSumMap.put(pathSum, 1);
        }

        System.out.println("--Prefix Sum ->" + prefixSum);
        System.out.println(pathSumMap);
        System.out.println("Ans --->" + ansCount);

        pathSumUtil(node.left, pathSum);
        pathSumUtil(node.right, pathSum);

        pathSumMap.put(pathSum, pathSumMap.get(pathSum)-1);
    }

    public int pathSum(TreeNode root, int k) {

        targetSum = k;
        ansCount = 0;

        pathSumMap = new HashMap<>();
        pathSumMap.put(0, 1);

        pathSumUtil(root, 0);

        return ansCount;
    }

    public static void main(String[] args) {

        TreeNode root = new TreeNode(1);
        root.left = new TreeNode(1);
        root.right = new TreeNode(3);
        root.right.left = new TreeNode(2);
        root.right.right = new TreeNode(1);

        PathSumIII437 ob = new PathSumIII437();

        int output = ob.pathSum(root, 2); // o/p =

        System.out.println(output);

    }
}
