package leetCodeProblems.GreedyAlgo;

/**
 * LeetCode - https://leetcode.com/problems/maximum-product-subarray/
 * 
 * @author anshul.agrawal
 *
 */

public class MaxProductSubArray152 {

	// DO NOT MODIFY THE ARGUMENTS WITH "final" PREFIX. IT IS READ ONLY
    public int maxProduct(int[] nums) {

        int maxProductSoFar = Integer.MIN_VALUE;
        int maxProductEndingHere = 1;

        for (int element : nums) {

            maxProductEndingHere = maxProductEndingHere*element;

            /*if (maxProductEndingHere < 0) {
                maxProductEndingHere = -maxProductEndingHere;
            }*/

            if ( maxProductSoFar < maxProductEndingHere) {
                maxProductSoFar = maxProductEndingHere;
            }

            if (maxProductEndingHere == 0) {
                maxProductEndingHere = 1;
            }
        }

        return maxProductSoFar;
    }

    public static void main(String[] args) {

        MaxProductSubArray152 obj = new MaxProductSubArray152();

        //int[] input = {-2,1,-3,4,-1,2,1,-5,4}; // Expected output = 6

        //int[] input = {2,3,-2,4}; // Expected output = 6

        //int[] input = {-2,0,-1}; // Expected output = 6

        //int[] input = {6, -3, -10, 0, 2}; // Expected output = 180

        //int[] input = {-1, -3, -10, 0, 60}; // Expected output = 60

        //int[] input = {-2, -40, 0, -2, -3}; // Expected output = 80

        int[] input = {-2,1,-3,4,-1,2,1,-5,4}; // Expected output = 6

        System.out.println(obj.maxProduct(input));
    }
}
