package CodingInterviewQuestions;

/**
 * Asked in Twillo Coding Round, 28thMay2022.
 *
 * To-Be-Done.
 *
 * Similar LeetCode Question - https://leetcode.com/problems/path-sum-iii/submissions/
 */
import java.util.ArrayList;
import java.util.List;

public class HackTree {

    //public int countVerticalPaths(List<Integer> cost, int edgeNodes, List<Integer> edgeFrom, List<Integer> edgeTo, int k) {

        // I know the logic to solve this problem
        // Constuct the first first
        // And then do the recursive traversal ( DFS ) from root to bottom. And keep calculating the sum and increase the count.
        // My limitation - Im not able to understand the inputs properly to construct the tree. Otherwise, i can do it quickly.

    //}

    public static void main(String[] args) {

        HackTree obj = new HackTree();

        List<Integer> cost = new ArrayList<>();
        cost.add(1);
        cost.add(2);
        cost.add(2);
        cost.add(1);
        cost.add(2);

        int edgeNodes = 4;

        List<Integer> edgeFrom = new ArrayList<>();
        edgeFrom.add(1);
        edgeFrom.add(2);
        edgeFrom.add(2);
        edgeFrom.add(2);

        List<Integer> edgeTo = new ArrayList<>();
        edgeFrom.add(4);
        edgeFrom.add(1);
        edgeFrom.add(5);
        edgeFrom.add(3);

        int k = 2;

        //obj.countVerticalPaths(cost, edgeNodes, edgeFrom, edgeTo, k);
    }
}
