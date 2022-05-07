package leetCodeProblems.BinarySearch;

/**
 * Leet Code - https://leetcode.com/problems/count-number-of-pairs-with-absolute-difference-k/
 */

import java.util.*;

public class CountPairsWithGivenDifference2006 {

	/**
	 * This is little different binary search.
	 * - It is handling duplicates.
	 * 
	 * For incremented target value, 
	 * - Instead of returning -1 ( if not found ), we would return next index ( with different value ) of the target value's index.
	 * 
	 * @param arr
	 * @param target
	 * @param low
	 * @return
	 */
	private int binarySearchWithDuplicatesHandling(int[] arr, int target, int low)  {
		
        int high = arr.length - 1;
         
        int ans = arr.length;//-1;
         
        while(low <= high) {
        	
            int mid = (low + high) / 2;
            
            if (arr[mid] == target) {
            	return mid;
            }
            else if (arr[mid] > target) {
            	
            	// This is special thing in this Binary Search. 
            	// This is needed to handle duplicates, when incremented target value is searched.
                ans = mid; 
                
                // low = low; // unchanged
                high = mid - 1;
            }
            else {
            	
            	low = mid + 1;
            	// high = high; // unchanged
            	
            }
        }
        
        return ans;
        
    }
    
    
    public int countKDifference(int[] nums, int k) {
        
        Arrays.sort(nums);
        
        int count = 0;
        
        for(int i=0; i<nums.length; i++) {

            int targetNum = k+nums[i];
            
            int targetNumIndex = binarySearchWithDuplicatesHandling(nums, targetNum, i+1);
            
            if (targetNumIndex != nums.length) {
            	
            	int incrementedTargetSum = targetNum+1;
            	
                int incrementedTargetSumIndex = binarySearchWithDuplicatesHandling(nums, incrementedTargetSum, i+1);
                
                System.out.println("nums[i] -> " + nums[i]);
                
                System.out.println("targetNum -> " + targetNum + ", targetSumIndex -> " + targetNumIndex);

            	System.out.println("incrementedTargetSum -> " + incrementedTargetSum + ", incrementedTargetSumIndex -> " + incrementedTargetSumIndex);
            	
            	System.out.println("---");
            	
                count += incrementedTargetSumIndex - targetNumIndex;
            }
            
            //
            
        }
        
        return count;
        
    }
    
    public static void main(String[] args) {
				
		//int[] inputArray = {1, 2, 2, 1};
		//int targetDifference = 1; // Output = 4
		
		int[] inputArray = {3,2,1,5,4};
		int targetDifference = 2; // Output = 3
		
		//int[] inputArray = {1, 3, 6};
		//int targetDifference = 2; // Output = 1
		
		CountPairsWithGivenDifference2006 obj = new CountPairsWithGivenDifference2006();
		
		System.out.println(obj.countKDifference(inputArray, targetDifference));
	}
}
