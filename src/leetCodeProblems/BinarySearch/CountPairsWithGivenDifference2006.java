package leetCodeProblems.BinarySearch;

/**
 * Leet Code - https://leetcode.com/problems/count-number-of-pairs-with-absolute-difference-k/
 */

import java.util.*;

public class CountPairsWithGivenDifference2006 {

	private int binarySearch(int[] arr, int X, int low)  {
        int high = arr.length - 1;
         
        int ans = arr.length;
         
        while(low <= high) {
            int mid = low + (high - low) / 2;
            if(arr[mid] >= X) {
                ans = mid;
                high = mid - 1;
            }
            else low = mid + 1;
        }
        return ans;
    }
    
    
    public int countKDifference(int[] nums, int k) {
        
        Arrays.sort(nums);
        
        int count = 0;
        
        for(int i=0; i<nums.length; i++) {

            int targetNumUsingSum = k+nums[i];
            
            //System.out.println("nums[i] ->" + nums[i]);
            
            int searchIndex = binarySearch(nums, targetNumUsingSum, i+1);
            
            if (searchIndex != nums.length) {
                int incrementalSearchIndex = binarySearch(nums, targetNumUsingSum+1, i+1);
                count += incrementalSearchIndex - searchIndex;
            }
            
            //System.out.println("searchIndex ->" + searchIndex);
            
        }
        
        return count;
        
    }
    
    public static void main(String[] args) {
				
		int[] inputArray = {1, 2, 2, 1};
		int targetDifference = 1;
		
		CountPairsWithGivenDifference2006 obj = new CountPairsWithGivenDifference2006();
		
		System.out.println(obj.countKDifference(inputArray, targetDifference));
	}
}
