package leetCodeProblems.Sorting;

/**
 * LeetCode - https://leetcode.com/problems/majority-element/
 * 
 * InterviewBit - https://www.interviewbit.com/problems/majority-element/
 */
import java.util.*;

public class MajorityElementIUsingSorting169 {
	
	public int majorityElement(int[] nums) {
        
		Arrays.sort(nums);        
        
        int localCurrentCount = 0;
        int localCurrentElement = 0;
            
        for(int i=0; i<nums.length; i++) {
            
            if (nums[i] != localCurrentElement) {
                localCurrentElement = nums[i];
                localCurrentCount = 0;
            }
            
            localCurrentCount++;
            
            if (localCurrentCount > (nums.length/2)) {
                return nums[i];
            }
        }
        
        return -1;
        
    }
    
    public static void main(String[] args) {
        
        //int[] inputArray = {3, 2, 3};
        
        int[] inputArray = {2,2,1,1,1,2,2};
        
        MajorityElementIUsingSorting169 obj = new MajorityElementIUsingSorting169();
        
        int elementWithMaxCount = obj.majorityElement(inputArray);

        System.out.println(elementWithMaxCount);
    }

}
