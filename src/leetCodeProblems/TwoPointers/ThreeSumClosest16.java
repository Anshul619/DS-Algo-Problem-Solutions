package leetCodeProblems.TwoPointers;

import java.util.Arrays;

public class ThreeSumClosest16 {
	public int threeSumClosest(int[] nums, int target) {
        
        int targetSum = target;
        int closestSum;

        if (nums.length > 2) {
            closestSum = nums[0] + nums[1] + nums[nums.length-1];
        }
        else {
            closestSum = -1;
        }

        Arrays.sort(nums);
    
        for(int fixedPointer=0; fixedPointer < nums.length; fixedPointer++) {
            
            int leftPointer = fixedPointer+1;
            int rightPointer = nums.length-1;
            
            while(leftPointer < rightPointer) {
                
                int sum = nums[fixedPointer] + nums[leftPointer] + nums[rightPointer];
                
                if (sum == targetSum) {
                    return targetSum;
                }
                else if(sum < targetSum) {
                    leftPointer++;
                }
                else {
                    rightPointer--;
                }

                //System.out.println("Sum diff -> " + Math.abs(sumDiff));
                //System.out.println("Closest diff -> " + Math.abs(closestDiff));
                
                int sumDiff = sum - targetSum;
                int closestDiff = closestSum - targetSum;
                
                if (Math.abs(sumDiff) < Math.abs(closestDiff)) {
                    closestSum = sum;
                }
                
            }
        }
        
        return closestSum;
        
    }
}
