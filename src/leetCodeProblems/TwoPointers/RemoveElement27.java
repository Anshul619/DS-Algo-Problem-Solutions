package leetCodeProblems.TwoPointers;

import java.util.Arrays;

/**
 * LeetCode - https://leetcode.com/problems/remove-element/
 *
 * TimeComplexity - O(n)
 * SpaceComplexity - O(1)
 */
public class RemoveElement27 {

    public int removeElement(int[] nums, int val) {

        if (nums.length == 0) {
            return 0;
        }

        //System.out.println("Before ->" + Arrays.toString(nums));

        int firstPointer = 0;
        int lastPointer = nums.length;

        while(firstPointer < lastPointer) {

            if (nums[lastPointer-1] == val) {
                lastPointer--;
            }
            else if (nums[firstPointer] == val) {
                lastPointer--;
                int temp = nums[lastPointer];
                nums[lastPointer] = nums[firstPointer];
                nums[firstPointer] = temp;
                firstPointer++;
            }
            else {
                firstPointer++;
            }

            //System.out.println("After ->" + Arrays.toString(nums));
        }

        System.out.println("FirstPointer ->" + firstPointer);
        System.out.println("LastPointer ->" + lastPointer);
        System.out.println("After ->" + Arrays.toString(nums));

        return lastPointer;
    }

    public static void main(String[] args) {

        //int[] nums = {3,2,2,3};

        //int[] nums = {0,1,2,2,3,0,4,2};

        int[] nums = {2};

        //int[] nums = {1};

        RemoveElement27 obj = new RemoveElement27();
        System.out.println(obj.removeElement(nums, 3));

    }
}
