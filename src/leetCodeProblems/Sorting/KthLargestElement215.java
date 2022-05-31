package leetCodeProblems.Sorting;

/**
 * LeetCode - https://leetcode.com/problems/kth-largest-element-in-an-array/solution/
 * Time Complexity - O(nlogn)
 * Space Complexity - O(1)
 *
 * More optimization - Use PriorityQueue
 */

import java.util.Arrays;

public class KthLargestElement215 {
    public int findKthLargest(int[] nums, int k) {
        Arrays.sort(nums);

        if (nums.length - k >= 0) {
            return nums[nums.length - k];
        }
        return -1;
    }

    public static void main(String[] args) {

        KthLargestElement215 obj = new KthLargestElement215();

        int[] input = {3,2,1,5,6,4};
        int k = 2;

        System.out.println(obj.findKthLargest(input, k)); //
    }
}
