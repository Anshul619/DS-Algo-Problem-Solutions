package leetCodeProblems;

/**
 * LeetCode Problem Link - https://leetcode.com/problems/remove-duplicates-from-sorted-array/submissions/
 * */

public class SortedArrayRemoveDuplicates26 {
	
	public int removeDuplicates(int[] nums) {
        
        if (nums.length == 0 || nums.length == 1) {
            return nums.length;
        }
        
        int answerIndex = 0;
        
        for(int i=1; i< nums.length;i++) {
            
            if (nums[answerIndex] != nums[i]) {
                answerIndex += 1;
                nums[answerIndex] = nums[i];
            }
            
        }
         
        return answerIndex + 1;
    }
	
	public static void main (String[] args)
    {
        int arr[] = {1, 2, 2, 3, 4, 4, 4, 5, 5};
         
        SortedArrayRemoveDuplicates26 obj = new SortedArrayRemoveDuplicates26();
        
        int n = obj.removeDuplicates(arr);
  
        // Print updated array
        for (int i=0; i<n; i++) {
        	System.out.println(arr[i]+" ");
        }
           
    }


}
