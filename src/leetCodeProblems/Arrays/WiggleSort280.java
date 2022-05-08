package leetCodeProblems.Arrays;

/**
 * LeetCode - https://leetcode.com/problems/wiggle-sort/
 * 
 * Interview Bit - https://www.interviewbit.com/problems/wave-array/
 */

import java.util.*;

public class WiggleSort280 {
	public void wiggleSort(int[] nums) {
        
        Arrays.sort(nums);

        for(int i = 1; i < (nums.length-1); i+=2) {

            int temp = nums[i];
            nums[i] = nums[i+1];
            nums[i+1] = temp;
        }

        return;
        
    }
	
	public static void main(String[] args) {
		
		int[] inputArray = {3,5,2,1,6,4};
		
		WiggleSort280 obj = new WiggleSort280();
		
		obj.wiggleSort(inputArray);
		
		System.out.println(Arrays.toString(inputArray));
		
	}

}
