package leetCodeProblems.Sorting;

/**
 * LeetCode - https://leetcode.com/problems/maximum-product-of-three-numbers/submissions/
 * InterviewBit - https://www.interviewbit.com/problems/highest-product/
 */

import java.util.Arrays;

public class MaximumProductOfThreeNumbers628 {

    public int maximumProduct(int[] nums) {
        Arrays.sort(nums);

        if (nums.length < 3) {
            return -1;
        }

        int length = nums.length;

        int maxProduct = Math.max(nums[length-1]*nums[length-2]*nums[length-3], nums[length-1]*nums[0]*nums[1]);

        return maxProduct;
    }

    public static void main(String[] args) {

        MaximumProductOfThreeNumbers628 obj = new MaximumProductOfThreeNumbers628();

        //int[] nums = {1,2,3,4};

        int[] nums = {-1,-2,-3};

        System.out.println(obj.maximumProduct(nums));

    }
}
