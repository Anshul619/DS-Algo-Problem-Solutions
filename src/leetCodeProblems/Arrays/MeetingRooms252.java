package leetCodeProblems.Arrays;

import java.util.*;

public class MeetingRooms252 {
	
	// Helper class extending Comparator interface
    static class CompareStartIntervalArray implements Comparator<int[]> {
        public int compare(int[] a, int[] b)
        {
            // if positive, then it would be in the same order
            return a[0] - b[0];
        }
    }
    
    public boolean canAttendMeetings(int[][] intervals) {
        
        Arrays.sort(intervals, new CompareStartIntervalArray());
        
        for (int i=0; i<intervals.length-1; i++) {
            
            int nextSlot = i+1;
            
            if (intervals[i][0] == intervals[nextSlot][0] && 
            	intervals[i][1] == intervals[nextSlot][1]) {
                return false;
            }
            
            if (intervals[i][1] > intervals[nextSlot][0]) {
                return false;
            }
        }
        
        return true;
    }
    
    public static void main(String[] args) {
    	
    	//int[][] intervals = {{7,10},{2,4}}; // Sorted - {{2,4}, {7,10}}
    	
    	//int[][] intervals = {{0,30},{5,10},{15,20}};
    	
    	int[][] intervals = {{4,9},{4,17},{9,10}};
    	
    	//int[][] intervals = {{13,15},{1,13}};
    	
    	//int[][] intervals = {{8, 11}, {17,20}, {17,20}};
    	
    	MeetingRooms252 obj = new MeetingRooms252();
    	
    	boolean result = obj.canAttendMeetings(intervals);
    	
    	System.out.println(result);
    	
    	
    }

}
