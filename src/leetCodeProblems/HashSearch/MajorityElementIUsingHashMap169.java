package leetCodeProblems.HashSearch;

/**
 * LeetCode - https://leetcode.com/problems/majority-element/
 * 
 * InterviewBit - https://www.interviewbit.com/problems/majority-element/
 */
import java.util.*;

public class MajorityElementIUsingHashMap169 {
	
	public int majorityElement(int[] nums) {
        
		HashMap<Integer, Integer> countsMap = new HashMap<Integer, Integer>();
        
        int localCurrentCount = 0;
            
        for(int i=0; i<nums.length; i++) {
            
            localCurrentCount = 0;
            
            if (countsMap.containsKey(nums[i])) {
                localCurrentCount = countsMap.get(nums[i]);
                localCurrentCount++;
                
                countsMap.put(nums[i], localCurrentCount);
            }
            else {
                localCurrentCount = 1;
                countsMap.put(nums[i], 1);
            }  
            
            if (localCurrentCount > (nums.length/2)) {
                return nums[i];
            }
        }
        
        return -1;
        
    }
    
    public static void main(String[] args) {
        
        //int[] inputArray = {3, 2, 3};
        
        int[] inputArray = {2,2,1,1,1,2,2};
        
        MajorityElementIUsingHashMap169 obj = new MajorityElementIUsingHashMap169();
        
        int elementWithMaxCount = obj.majorityElement(inputArray);

        System.out.println(elementWithMaxCount);
    }

}
