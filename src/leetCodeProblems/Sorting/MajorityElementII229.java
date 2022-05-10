package leetCodeProblems.Sorting;

/**
 * LeetCode - https://leetcode.com/problems/majority-element-ii/
 */
import java.util.*;

public class MajorityElementII229 {
	
	public List<Integer> majorityElement(int[] nums) {
        
		Arrays.sort(nums);        
        
        List<Integer> outputList = new ArrayList<Integer>();
        
        int localCurrentCount = 0;
        int localCurrentElement = 0;
            
        for(int i=0; i<nums.length; i++) {
            
            if (nums[i] != localCurrentElement) {
                localCurrentElement = nums[i];
                localCurrentCount = 0;
            }
            
            localCurrentCount++;
            
            if (localCurrentCount > (nums.length/3) && !outputList.contains(nums[i])) {
                outputList.add(nums[i]);
            }
        }
        
        return outputList;
        
    }
    
    public static void main(String[] args) {
        
    	//int[] inputArray = {1};
    	
        //int[] inputArray = {3, 2, 3};
        
        int[] inputArray = {2,2,1,1,1,2,2};
        
        MajorityElementII229 obj = new MajorityElementII229();
        
        System.out.println(obj.majorityElement(inputArray));
    }

}
