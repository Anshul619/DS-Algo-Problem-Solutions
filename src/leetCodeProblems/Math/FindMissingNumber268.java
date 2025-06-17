package leetCodeProblems.MathCalculations;

/**
 * LeetCode Link - https://leetcode.com/problems/missing-number/
 * 
 * @author anshul.agrawal
 *
 */
public class FindMissingNumber268 {
	
	public int missingNumber(int[] nums) {
        
        int arraySize = nums.length;
        int currentSum = 0;
        
        int expectedSum = (arraySize*(arraySize+1))/2;
        
        for(int i=0; i<arraySize;i++) {
            currentSum += nums[i];
        }
        
        return expectedSum - currentSum;
        
        //System.out.println("MI")
    }

	public static void main(String args[]) {

		//int[] inputArray = {3,0,1};

		int[] inputArray = {0,2};
		
		FindMissingNumber268 obj = new FindMissingNumber268();

		System.out.println(obj.missingNumber(inputArray));
		// Your code goes here
	}

}
