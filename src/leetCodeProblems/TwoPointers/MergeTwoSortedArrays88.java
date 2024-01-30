package leetCodeProblems.TwoPointers;

/**
 * LeetCode - https://leetcode.com/problems/merge-sorted-array/
 * InterviewBit - https://www.interviewbit.com/problems/merge-two-sorted-lists-ii/
 */
import java.util.*;

public class MergeTwoSortedArraysInPlace88 {
	
	public void shiftArray(int[] nums1, int arrayLength, int startPointer, int startValueToBeAdded) {
        
        int endPointer = arrayLength;
        
        System.out.println("Array ->"+ Arrays.toString(nums1));
        System.out.println("Array Length ->"+ arrayLength);
        System.out.println("startValueToBeAdded ->"+ startValueToBeAdded);
        System.out.println("Start Pointer ->"+ startPointer);
        System.out.println("End Pointer ->"+ endPointer);
        
        while ( startPointer < endPointer ) {
            nums1[endPointer] = nums1[endPointer-1];
            endPointer--;
        }
        
        nums1[startPointer] = startValueToBeAdded;
        
        return;
    }
    
    public void merge(int[] nums1, int m, int[] nums2, int n) {
        
        int firstArrayPointer = 0;
        int secondArrayPointer = 0;
        
        int firstArrayLength = m;
        
        while (secondArrayPointer < n && firstArrayPointer < firstArrayLength) {
                
              /*System.out.println("----");
              System.out.println("first array pointer ->" + firstArrayPointer);
              System.out.println("second array pointer ->" + secondArrayPointer);
              System.out.println("first array ->" + Arrays.toString(nums1));
              System.out.println("second array ->" + Arrays.toString(nums2));
              System.out.println("----");*/
            
              if (nums1[firstArrayPointer] == nums2[secondArrayPointer]) {
                  
                  shiftArray(nums1, firstArrayLength, firstArrayPointer, nums2[secondArrayPointer]);
                  firstArrayLength++;
                  
                  firstArrayPointer++;
                  secondArrayPointer++;
              }
              else if (nums1[firstArrayPointer] < nums2[secondArrayPointer]) {
                  firstArrayPointer++;
              }      
              else {
                  shiftArray(nums1, firstArrayLength, firstArrayPointer, nums2[secondArrayPointer]);
                  firstArrayLength++;
                  secondArrayPointer++;
              }
        }
        
        while(secondArrayPointer < n) {
            
            nums1[firstArrayPointer] = nums2[secondArrayPointer];
            secondArrayPointer++;
            firstArrayPointer++;
            //System.out.println("first array ->" + Arrays.toString(nums1));
        }
        
        System.out.println("FINAL merged array ->" + Arrays.toString(nums1));
    }
    
    public static void main(String[] args) {
        
        int[] firstArray = {1,2,3,0,0,0};
        int[] secondArray = {2,5,6};
        int m = 3;
        int n = 3;
        
        //int[] firstArray = {1,2,4,5,6,0};
        //int m = 5;
        //int[] secondArray = {3};
        //int n = 1;
        
        //int[] firstArray = {1};
        //int[] secondArray = {};
        //int m = 1;
        //int n = 0;
        
        //int[] firstArray = {0};
        //int[] secondArray = {1};
        //int m = 0;
        //int n = 1;
        
        //int[] firstArray = {2,0};
        //int[] secondArray = {1};
        //int m = 1;
        //int n = 1;

        MergeTwoSortedArraysInPlace88 obj = new MergeTwoSortedArraysInPlace88();
        
        obj.merge(firstArray, m, secondArray, n);
    }

}
