package leetCodeProblems.Arrays;

import java.util.*;

public class ArrayMergeOverlappingIntervals56 {
	
	static class Interval {
		int start;
		int end;
	}
	
	// Helper class extending Comparator interface
    static class CompareStartInterval implements Comparator<Interval> {
        public int compare(Interval a, Interval b)
        {
            // if positive, then it would be in the same order
            return a.start - b.start;
        }
    }
    
 // Helper class extending Comparator interface
    static class CompareStartIntervalArray implements Comparator<int[]> {
        public int compare(int[] a, int[] b)
        {
            // if positive, then it would be in the same order
            return a[0] - b[0];
        }
    }

    public ArrayList<Interval> merge(ArrayList<Interval> intervals) {

        if (intervals.size() == 0) {
            return new ArrayList<Interval>();
        }
        
        ArrayList<Interval> outputList = new ArrayList<Interval>();
        
        Collections.sort(intervals, new CompareStartInterval());
        
        outputList.add(intervals.get(0));
        
        for(int i=1; i<intervals.size(); i++) {
            
            // Get latest element from output arrayList
            int latestIndex = outputList.size() -1;
            Interval latestPushedElement = outputList.get(latestIndex); 
            
            if (latestPushedElement.start <= intervals.get(i).start && 
                intervals.get(i).start <= latestPushedElement.end) {
                
                //System.out.println("Merge - " + i);
                if (latestPushedElement.end < intervals.get(i).end) {
                    latestPushedElement.end = intervals.get(i).end;
                }
                
                outputList.set(latestIndex, latestPushedElement);
            }
            else {
                outputList.add(intervals.get(i));
            }
        }
        
        
        return outputList;

    }
    
    
    public int[][] merge(int[][] intervals) {
        
        if (intervals.length == 0) {
            return new int[0][0];
        }
        
        ArrayList<int[]> outputList = new ArrayList<int[]>();
        
        Arrays.sort(intervals, new CompareStartIntervalArray());
        
        outputList.add(intervals[0]);
        
        for(int i=1; i<intervals.length; i++) {
            
            // Get latest element from output arrayList
            int latestIndex = outputList.size() -1;
            int[] latestPushedElement = outputList.get(latestIndex); 
            
            if (latestPushedElement[0] <= intervals[i][0] && 
                intervals[i][0] <= latestPushedElement[1]) {
                
                //System.out.println("Merge - " + i);
                if (latestPushedElement[1] < intervals[i][1]) {
                    latestPushedElement[1] = intervals[i][1];
                }
                
                outputList.set(latestIndex, latestPushedElement);
            }
            else {
                outputList.add(intervals[i]);
            }
        }
        
        int[][] output = new int[outputList.size()][2];
        
        for(int i=0; i<outputList.size(); i++) {
            output[i] = outputList.get(i);
        }
        
        return output;
        
    }

}
