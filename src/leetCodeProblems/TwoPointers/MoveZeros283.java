package leetCodeProblems.TwoPointers;

import java.util.Arrays;

/**
 * LeetCode - https://leetcode.com/problems/move-zeroes/submissions/
 *
 * TimeComplexity - O(n)
 * SpaceComplexity - O(1)
 */
public class MoveZeros283 {

    public void moveZeroes(int[] nums) {

        int leftPointer = 0;

        for (int i=0; i < nums.length; i++) {
            if (nums[i] != 0) {
                nums[leftPointer] = nums[i];
                leftPointer++;
            }
        }

        while(leftPointer < nums.length) {
            nums[leftPointer] = 0;
            leftPointer++;
        }
    }

    public static void main(String[] args) {
        int[] nums = {0,1,0,3,12};

        MoveZeros283 obj = new MoveZeros283();
        obj.moveZeroes(nums);

        System.out.println(Arrays.toString(nums));
    }
}
