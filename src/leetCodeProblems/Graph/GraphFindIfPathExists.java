package leetCodeProblems.Graph;

import java.util.*;

public class GraphFindIfPathExists {

	ArrayList<ArrayList<Integer>> buildGraph(int A, int[][] B) {

		ArrayList<ArrayList<Integer>> graph = new ArrayList<ArrayList<Integer>>();

		for (int i = 0; i < A; i++) {
			graph.add(new ArrayList<Integer>());
		}

		for (int j = 0; j < B.length; j++) {
			ArrayList<Integer> temp = graph.get(B[j][0]);
			temp.add(B[j][1]);

			temp = graph.get(B[j][1]);
			temp.add(B[j][0]);
		}

		return graph;
	}

	boolean bfsPath(ArrayList<ArrayList<Integer>> graph, int source, int destination, boolean[] visitedOnPath) {
		Queue<Integer> queue = new LinkedList<Integer>();
		visitedOnPath[source] = true;

		// https://www.geeksforgeeks.org/queue-offer-method-in-java/
		queue.add(source);

		while (!queue.isEmpty()) {

			int node = queue.remove();

			if (node == destination) {
				return true;
			}

			ArrayList<Integer> neighbors = graph.get(node);

			for (int j = 0; j < neighbors.size(); j++) {

				int next = neighbors.get(j);

				if (!visitedOnPath[next]) {
					visitedOnPath[next] = true;
					queue.add(next);
				}
			}
		}

		return false;

	}

	public boolean validPath(int n, int[][] edges, int source, int destination) {

		ArrayList<ArrayList<Integer>> graph = buildGraph(n, edges);

		// This is important to ignore cycles in graph.
		boolean[] visitedOnPath = new boolean[n];

		return bfsPath(graph, source, destination, visitedOnPath);
	}
}
