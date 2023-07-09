package leetCodeProblems.TwoPointers;

/**
 * LeetCode - https://leetcode.com/problems/merge-sorted-array/
 * InterviewBit - https://www.interviewbit.com/problems/merge-two-sorted-lists-ii/
 */

import java.util.*;

public class MergeTwoSortedListsInPlace88 {
	
	public void setElementToList(ArrayList<Integer> nums1, int index, int value) {

        if (index < nums1.size()) {
            nums1.set(index, value);
        }
        else {
            nums1.add(value);
        }

    }
    public void shiftArray(ArrayList<Integer> nums1, int arrayLength, int startPointer, int startValueToBeAdded) {
        
        int endPointer = arrayLength;
        
        //System.out.println("Start Pointer ->"+ startPointer);
        //System.out.println("End Pointer ->"+ endPointer);
        System.out.println("End Pointer ->"+ endPointer);
        
        while ( startPointer < endPointer ) {

            setElementToList(nums1, endPointer, nums1.get(endPointer-1));
            endPointer--;
        }
        
        setElementToList(nums1, startPointer, startValueToBeAdded);
        return;
    }

	public void merge(ArrayList<Integer> a, ArrayList<Integer> b) {

        int firstArrayPointer = 0;
        int secondArrayPointer = 0;
        
        int firstArrayLength = a.size();
        int n = b.size();

        while (secondArrayPointer < n && firstArrayPointer < firstArrayLength) {
                
              /*System.out.println("----");
              System.out.println("first array pointer ->" + firstArrayPointer);
              System.out.println("second array pointer ->" + secondArrayPointer);
              System.out.println("first array ->" + Arrays.toString(nums1));
              System.out.println("second array ->" + Arrays.toString(nums2));
              System.out.println("----");*/
            
              if (a.get(firstArrayPointer) == b.get(secondArrayPointer)) {
                  
                  shiftArray(a, firstArrayLength, firstArrayPointer, b.get(secondArrayPointer));
                  firstArrayLength++;
                  
                  firstArrayPointer++;
                  secondArrayPointer++;
              }
              else if (a.get(firstArrayPointer) < b.get(secondArrayPointer)) {
                  firstArrayPointer++;
              }      
              else {
                  shiftArray(a, firstArrayLength, firstArrayPointer, b.get(secondArrayPointer));
                  firstArrayLength++;
                  secondArrayPointer++;
              }
        }
        
        while(secondArrayPointer < n) {
            
            setElementToList(a, firstArrayPointer, b.get(secondArrayPointer));
            secondArrayPointer++;
            firstArrayPointer++;
            //System.out.println("first array ->" + Arrays.toString(nums1));
        }
        
        System.out.println("FINAL merged array ->" + Arrays.toString(a.toArray()));
	}
	
	public static void main(String[] args) {
        
        //int[] firstArray = {1,2,3,0,0,0};
        //int[] secondArray = {2,5,6};
        //int m = 3;
        //int n = 3;
        
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
        
        ArrayList<Integer> firstArrayList = new ArrayList<Integer>();
        firstArrayList.add(-4);
        firstArrayList.add(3);
        
        ArrayList<Integer> secondArrayList = new ArrayList<Integer>();
        secondArrayList.add(-2);
        secondArrayList.add(-2);

        System.out.println("Merged array1 ->" + Arrays.toString(firstArrayList.toArray()));
        System.out.println("Merged array2 ->" + Arrays.toString(secondArrayList.toArray()));
        
        MergeTwoSortedListsInPlace88 obj = new MergeTwoSortedListsInPlace88();
        
        obj.merge(firstArrayList, secondArrayList);
        
        System.out.println("Output");
    }

}
