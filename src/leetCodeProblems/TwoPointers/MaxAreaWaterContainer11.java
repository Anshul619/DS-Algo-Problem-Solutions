package leetCodeProblems.TwoPointers;

public class MaxAreaWaterContainer11 {
	
	public int maxArea(int[] A) {
        int leftPointer = 0;
        int rightPointer = A.length - 1;
        int areaMax = 0;
        
        while(leftPointer < rightPointer) {
        
            areaMax = Math.max(areaMax, (rightPointer - leftPointer) * Math.min(A[leftPointer], A[rightPointer]));
            
            if (A[leftPointer] < A[rightPointer]) {
                leftPointer += 1;
            }
            else {
                rightPointer -= 1;
            }
            
        }
        
        return areaMax;
    }

}
