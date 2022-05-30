package leetCodeProblems.PrefixSum;

/**
 * LeetCode - https://leetcode.com/problems/continuous-subarray-sum/
 * Time Complexity - O(n)
 * Space Complexity - O(n)
 */

import java.util.HashMap;

public class ContinuousSubarraySum523 {
    public boolean checkSubarraySum(int[] nums, int k) {
        HashMap<Integer, Integer> remainderMap = new HashMap<>();
        //remainderMap.put(0, 1);

        int currentSum = 0;
        int remainder = 0;

        for (int i=0; i < nums.length; i++) {

            if (nums[i] == 0) {
                remainder = 0;
            }
            else {
                currentSum += nums[i];
                remainder = currentSum%k;
            }

            System.out.println(currentSum);
            System.out.println(remainderMap);

            if (currentSum%k == 0) {
                return true;
            }

            if (remainderMap.containsKey(remainder)) {

                int index = remainderMap.get(remainder);

                // Continuous subarray of size at least two
                if (i-index >= 1) {
                    return true;
                }
            }

            remainderMap.put(remainder, i);
        }

        System.out.println("map ->" + remainderMap);
        return false;
    }

    public static void main(String[] args) {

        ContinuousSubarraySum523 obj = new ContinuousSubarraySum523();

        //int[] inputArray = {23,2,4,6,7};
        //int targetSumBase = 6;

        //int[] inputArray = {23,2,6,4,7};
        //int targetSumBase = 6;

        int[] inputArray = {23,2,4,6,6};
        int targetSumBase = 7;

        //int[] inputArray = {23,2,6,4,7};
        //int targetSumBase = 13;

        //int[] inputArray = {5,0,0,0};
        //int targetSumBase = 3;

        //int[] inputArray = {0,0};
        //int targetSumBase = 1;

        //int[] inputArray = {0};
        //int targetSumBase = 1;

        System.out.println(obj.checkSubarraySum(inputArray, targetSumBase));
    }

}
