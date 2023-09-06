package leetCodeProblems.Sorting;

import java.util.*;

public class MergeOverlappingIntervals56 {
	
	// Helper class extending Comparator interface
    static class CompareStartIntervalArray implements Comparator<int[]> {
        public int compare(int[] a, int[] b)
        {
            // if positive, then it would be in the same order
            return a[0] - b[0];
        }
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
    
    public static void main(String[] args) {
    	
    	//int[][] intervals = {{7,10},{2,4}};
    	
    	//int[][] intervals = {{0,30},{5,10},{15,20}};
    	
    	int[][] intervals = {{13,15},{1,13}};
    	
    	MergeOverlappingIntervals56 obj = new MergeOverlappingIntervals56();
    	
    	int[][] output = obj.merge(intervals);
    	
    	//System.out.println(Arrays.toString(obj.merge(intervals)));
    	
    	for (int i=0; i < output.length; i++) {
    		System.out.println(Arrays.toString(output[i]));
    	}
    	
    }

}
