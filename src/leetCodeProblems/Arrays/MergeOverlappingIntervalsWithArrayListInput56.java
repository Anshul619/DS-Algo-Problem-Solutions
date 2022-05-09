package leetCodeProblems.Arrays;

import java.util.ArrayList;
import java.util.Collections;
import java.util.Comparator;

public class MergeOverlappingIntervalsWithArrayListInput56 {
	
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
}
