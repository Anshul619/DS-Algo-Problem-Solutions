package leetCodeProblems.Graph;

/**
 * InterviewBit - https://www.interviewbit.com/problems/possibility-of-finishing-all-courses-given-prerequisites/
 */

import java.util.*;

public class CanFinishAllCourses {
	ArrayList<ArrayList<Integer>> buildGraph(int A, int[] B, int[] C) {

        ArrayList<ArrayList<Integer>> graph = new ArrayList<ArrayList<Integer>>();

        for(int i = 0; i < A; i++) {
            //System.out.println(i);
            //System.out.println("Graph add");
            graph.add(new ArrayList<Integer>());
        }

        for(int j=0; j < B.length; j++) {
            //System.out.println(j);
            //System.out.println("Before get");
            ArrayList<Integer> temp = graph.get(B[j]-1);
            //System.out.println("After");
            temp.add(C[j]-1);
        }

        return graph;
    }

    /**
     * A DFS based function to check if there is a cycle in the directed graph.
     */
    boolean dfsCycle(ArrayList<ArrayList<Integer>> graph, int node,
            boolean[] visitedOnPath, boolean[] allVisited) {

        if (allVisited[node]) {
            return false;
        }

        allVisited[node] = true;
        visitedOnPath[node] = true;

        ArrayList<Integer> neighbors = graph.get(node);

        for(int j=0; j < neighbors.size(); j++) {
            if (visitedOnPath[neighbors.get(j)] || dfsCycle(graph, neighbors.get(j), visitedOnPath, allVisited)) {
                return true;
            }
        }

        visitedOnPath[node] = false;

        return false;

    }

    boolean canFinish(int A, int[] B, int[] C) {

        ArrayList<ArrayList<Integer>> graph = buildGraph(A, B, C);

        boolean[] visitedOnPath = new boolean[A];
        boolean[] allVisited = new boolean[A]; // This is needed to remove redundant looping/cycles.

        for(int j=0; j < A; j++) {
            if (!allVisited[j] && dfsCycle(graph, j, visitedOnPath, allVisited)) {
                return false;
            }
        }

        return true;

    }

    public int solve(int A, int[] B, int[] C) {

        boolean result = canFinish(A, B, C);

        if (result) {
            return 1;
        }
        else {
            return 0;
        }
    }
}
