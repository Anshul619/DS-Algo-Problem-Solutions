package leetCodeProblems.PriorityQueue;

import java.util.PriorityQueue;

public class SamplePriorityQueue {
	public static void main(String[] args) {
		
		PriorityQueue<Integer> priorityQueue =new PriorityQueue<Integer>();
     
		priorityQueue.add(5);
		priorityQueue.add(1);
		priorityQueue.add(4);
		
		priorityQueue.add(10);
		priorityQueue.add(2);
		
		priorityQueue.add(8);
		
		priorityQueue.add(20);
		
		priorityQueue.add(3);
		
		System.out.println(priorityQueue);
		
		while(!priorityQueue.isEmpty()) {
			System.out.println(priorityQueue.remove());
		}
		
	}
}
