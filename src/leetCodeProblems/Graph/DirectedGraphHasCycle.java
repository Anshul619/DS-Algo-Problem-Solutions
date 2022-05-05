package leetCodeProblems.Graph;

import java.util.*;

/**
 * LeetCode Problem Link - <TBD>
 *
 * InterviewBit - https://www.interviewbit.com/problems/cycle-in-directed-graph/

 * */

public class DirectedGraphHasCycle {
		
	  //Consider Zero-based indexing carefully, while building Graph in ArrayList.
	  ArrayList<ArrayList<Integer>> buildGraph(int A, int[][] B) {

	       ArrayList<ArrayList<Integer>> graph = new ArrayList<ArrayList<Integer>>();

	       for(int i = 0; i < A; i++) {
	           graph.add(new ArrayList<Integer>());
	       }

	       for(int j=0; j < B.length; j++) {
	           ArrayList<Integer> temp = graph.get(B[j][0] - 1);
	           temp.add(B[j][1] - 1);
	       }

	       return graph;
	   }

	   // A DFS based function to check if there is a cycle in the directed graph.
	   boolean dfsCycle(ArrayList<ArrayList<Integer>> graph, int node,
	           boolean[] visitedOnPath) {

	       visitedOnPath[node] = true;

	       ArrayList<Integer> neighbors = graph.get(node);

	       for(int j=0; j < neighbors.size(); j++) {
	           if (visitedOnPath[neighbors.get(j)] || dfsCycle(graph, neighbors.get(j), visitedOnPath)) {
	               return true;
	           }
	       }

	       visitedOnPath[node] = false;

	       return false;

	   }

	   public int solve(int A, int[][] B) {

	       ArrayList<ArrayList<Integer>> graph = buildGraph(A, B);

	       boolean[] visitedOnPath = new boolean[A];

	       for(int i=0; i<A; i++) {
	           if (dfsCycle(graph, i, visitedOnPath)) {
	               return 1;
	           }
	       }

	       return 0;
	   }


}
