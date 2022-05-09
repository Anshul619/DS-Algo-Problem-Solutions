package leetCodeProblems.StacksAndQueues;

import java.util.*;

public class MeetingRoomsII253 {
	
	// Helper class extending Comparator interface
    static class CompareStartIntervalArray implements Comparator<int[]> {
        public int compare(int[] a, int[] b)
        {
            // if positive, then it would be in the same order
            return a[0] - b[0];
        }
    }
    
    // Helper class extending Comparator interface
    static class PriorityQueueCompare implements Comparator<Integer> {
		@Override
		public int compare(Integer o1, Integer o2) {
			// TODO Auto-generated method stub
			return o1 - o2;
		}
    }
    
    public int minMeetingRooms(int[][] intervals) {
        
        // Check for the base case. If there are no intervals, return 0
        if (intervals.length == 0) {
          return 0;
        }

        PriorityQueue<Integer> ongoingMeetingRoomsQueue =
            new PriorityQueue<Integer>(intervals.length, new PriorityQueueCompare());

        // Sort the intervals by start time
        Arrays.sort(intervals, new CompareStartIntervalArray());

        // Add the first meeting's end time in Priority Queue
        ongoingMeetingRoomsQueue.add(intervals[0][1]);

        // Iterate over remaining intervals
        for (int i = 1; i < intervals.length; i++) {
        
          // For every meeting, check if the minimum element of the heap i.e. the room at the top of the heap is free or not.
          if (intervals[i][0] >= ongoingMeetingRoomsQueue.peek()) {
            ongoingMeetingRoomsQueue.remove();
          }
          
          // If the room is free, then we extract the topmost element and add it back with the ending time of the current meeting we are processing.
          ongoingMeetingRoomsQueue.add(intervals[i][1]);
        }

        // The size of the heap tells us the minimum rooms required for all the meetings.
        return ongoingMeetingRoomsQueue.size();
        
    }
 
    public static void main(String[] args) {
        //System.out.println("Hello World!");
        
        //int[][] intervals = {{0,30},{5,10},{15,20}};
        
        //int[][] intervals = {{7,10},{2,4}};
        
        int[][] intervals = {{4,9},{4,17},{9,10}};
        
        MeetingRoomsII253 obj = new MeetingRoomsII253();
    	
        int minMeetingRoomsNeeded = obj.minMeetingRooms(intervals);
        
    	//int minMeetingRoomsNeeded = obj.minMeetingRooms(intervals);
    	
    	System.out.println(minMeetingRoomsNeeded);
    }

}
