package leetCodeProblems.GreedyAlgo;

/**
 * LeetCode - https://leetcode.com/problems/maximum-subarray/
 * 
 * @author anshul.agrawal
 *
 */

public class MaxSumSubArray53 {

	// DO NOT MODIFY THE ARGUMENTS WITH "final" PREFIX. IT IS READ ONLY
    public int maxSubArray(final int[] A) {

        int maxSumSoFar = Integer.MIN_VALUE;
        int maxSumEndingHere = 0;

        for (int element : A) {

            maxSumEndingHere = maxSumEndingHere + element;

            if (maxSumSoFar < maxSumEndingHere) {
                maxSumSoFar = maxSumEndingHere;
            }

            if (maxSumEndingHere < 0) {
                maxSumEndingHere = 0;
            }
        }

        return maxSumSoFar;
    }

    public static void main(String[] args) {

        MaxSumSubArray53 obj = new MaxSumSubArray53();

        //int[] input = {-2,1,-3,4,-1,2,1,-5,4}; // Expected output = 6

        int[] input = {5,4,-1,7,8}; // Expected output = 23

        System.out.println(obj.maxSubArray(input));
    }
}
