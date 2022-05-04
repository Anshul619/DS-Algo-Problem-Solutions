package leetCodeProblems.TwoPointers;

import java.util.*;

public class ThreeSum15 {
	
	public List<List<Integer>> threeSum2PointerApproach(int[] nums) {
        
        int targetSum = 0;
        
        Arrays.sort(nums);
        
        List<List<Integer>> output = new ArrayList<List<Integer>>();
        
        for(int fixedPointer=0; fixedPointer < nums.length; fixedPointer++) {
            
            int leftPointer = fixedPointer+1;
            int rightPointer = nums.length-1;
            
            while(leftPointer < rightPointer) {
                
                int sum = nums[fixedPointer] + nums[leftPointer] + nums[rightPointer];
                
                if (sum == targetSum) {
                    ArrayList<Integer> temp = new ArrayList<Integer>();
                    temp.add(nums[fixedPointer]);
                    temp.add(nums[leftPointer]);
                    temp.add(nums[rightPointer]);
                    
                    /*if (!s.contains(temp)) {
                        output.add(temp);
                    }*/
                    
                    break;
                }
                else if(sum < targetSum) {
                    leftPointer++;
                }
                else {
                    rightPointer--;
                }
            }
        }
        
        return output;
        
    }

}
