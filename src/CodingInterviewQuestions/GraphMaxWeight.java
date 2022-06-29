package CodingInterviewQuestions;

/**
 * Asked in Amex, 29thJune2022
 *
 * To-Be-Completed
 */

import java.util.ArrayList;
import java.util.Comparator;
import java.util.PriorityQueue;

class GraphMaxWeight {
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

        for(int j=0; j < C.length; j++) {
            //System.out.println(j);
            //System.out.println("Before get");
            ArrayList<Integer> temp = graph.get(C[j]-1);
            //System.out.println("After");
            temp.add(B[j]-1);
        }

        return graph;
    }
    public int solution(int N, int[] A, int[] B) {
        // write your code in Java SE 8

        ArrayList<ArrayList<Integer>> graph = buildGraph(A.length, A, B);

        int[] weighedArray = new int[A.length];

        PriorityQueue<Integer> pq = new PriorityQueue(Comparator.reverseOrder());

        for(int i=0; i < graph.size(); i++) {
            pq.add(graph.get(i).size());
        }

        int currentCounter = N;

        for(int i=0; i < pq.size(); i++) {
            weighedArray[pq.remove()] = currentCounter;
            currentCounter--;
        }

        int sum = 0;

        for(int i=0; i < graph.size(); i++) {

            for(int j=0; i < graph.get(i).size(); i++) {
                // Code to be written, due to time limitation

                // Calculate the sum here, since we have assigned weight to every node in the graph.

                sum += weighedArray[graph.get(i).get(j)] + weighedArray[graph.get(i).get(i)];

            }

        }

        return sum;
    }
}


