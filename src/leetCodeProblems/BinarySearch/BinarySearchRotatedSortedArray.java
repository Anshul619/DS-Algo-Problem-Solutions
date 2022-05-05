package leetCodeProblems.BinarySearch;

/**
 * Leet Code - https://leetcode.com/problems/binary-search/submissions/
 *
 * @author anshul.agrawal
 *
 */

public class BinarySearchRotatedSortedArray {

	public static int binarySearch(int start, int end, int[] nums1, int target) {

        int middle = (start+end)/2;

        if (start < 0 || end < 0 || start > end) {
            return -1;
        }

        if (nums1[middle] == target) {
            return middle;
        }
        else if (start == end) {
            return -1;
        }
        else if (nums1[start] <= nums1[middle]) { // Left side is sorted array
        	 
            if (nums1[start] <= target && target <= nums1[middle]) { // Apply Binary Search i.e. target element value is b/w start and middle
            	// start = start; // unchanged
                end = middle - 1;
            }
            else {
                start = middle + 1;
                // end = end; // unchanged
            }

        }
        else {  // right side is SORTED

        	if (nums1[middle] <= target && target <= nums1[end]) {
        		start = middle + 1;
        	}
        	else {
        		end = middle - 1;
        	}
        }

        return binarySearch(start, end, nums1, target);
    }

    public int search(int[] nums, int target) {

        int searchElement = binarySearch(0, nums.length-1, nums, target);

        return searchElement;
    }

}
