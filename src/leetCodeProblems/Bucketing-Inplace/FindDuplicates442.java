package leetCodeProblems.Bucketing;

/**
 * LeetCode - https://leetcode.com/problems/find-all-duplicates-in-an-array/
 * TimeComplexity - O(n)
 * SpaceComplexity - O(1)
 */

import java.util.*;

public class FindDuplicates442 {

	public List<Integer> findDuplicates(int[] nums) {

        List<Integer> output = new ArrayList<Integer>();

        for(int i=0; i<nums.length; i++) {

            // This is important, to get original value at nums[i].
            // Because it might be changed due to array operations
            int hashMapIndexInArray = nums[i] % nums.length;

            if (hashMapIndexInArray >= 0) {
                //System.out.println (hashMapIndexInArray);
                nums[hashMapIndexInArray] = nums[hashMapIndexInArray] + nums.length;
            }
        }

        for(int i=0; i<nums.length; i++) {

            //System.out.println(nums[i]);
            if (nums[i] > (nums.length*2)) {

                if (i == 0) {
                    output.add(nums.length);
                }
                else {
                    output.add(i);
                }

            }
        }

        return output;
    }

}
