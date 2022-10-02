package leetCodeProblems.HashSearch;

import java.util.*;

public class MaxPointsOnSameLine149 {
	
	static HashMap<Double, Integer> slopeCountHashMap = new HashMap<Double, Integer>();
    
    /**
    * Calculate slope of start and end point
    */
    public double calculateSlope(int[] startPoint, int[] endPoint) {
        
        int yDif = endPoint[1] - startPoint[1];
        int xDif = endPoint[0] - startPoint[0];
        
        if (xDif == 0) {
            return Integer.MAX_VALUE;
        }
        else {
            return 0.0 + ((double)(yDif)/(double)(xDif)); // Impoortant
            //return yDif/xDif;
        }     
    }
    
    /**
    * Increment Slope Count in the slopeCountHashMap
    */
    public void incrementSlopeCount(double slope) {
        if (slopeCountHashMap.containsKey(slope)) {
                    
            int angleCount = slopeCountHashMap.get(slope);
            angleCount++;

            slopeCountHashMap.put(slope, angleCount);
        }
        
        else {
            slopeCountHashMap.put(slope, 2);
        }
    }
    
    public int maxPoints(int[][] points) {
            
        int maxCountAnswer = 0;
        
        if (points.length > 0) {
            maxCountAnswer = 1;
        }
         
        for(int i=0; i<points.length;i++) {
            
            slopeCountHashMap.clear();
            
            for (int j=i+1; j<points.length; j++) {
                    
                double slope = calculateSlope(points[i], points[j]);
                
                //System.out.println("Slope - "+ slope);
                incrementSlopeCount(slope);
            }
            
            for(int count: slopeCountHashMap.values()) {
                
                if (count > maxCountAnswer) {
                    maxCountAnswer = count;
                }
            }
            
        }
        
        return maxCountAnswer;
        
    }

}
