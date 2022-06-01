package leetCodeProblems.HashSearch;

/**
 * LeetCode - https://leetcode.com/problems/longest-consecutive-sequence/
 * TimeComplexity - O(n2)
 * SpaceComplexity - O(n)
 */

import java.util.*;

public class LongestConsecutiveSequence128 {

	public int longestConsecutive(int[] nums) {
        HashSet<Integer> hashSet = new HashSet<Integer>();
        int ans = 0;

        for(int i=0; i < nums.length; i++) {
            hashSet.add(nums[i]);
        }

        for(int i=0; i < nums.length; i++) {

            // This is important, to remove UNREQUIRED loops.
            if (!hashSet.contains(nums[i]-1)) {

                int j = 0;

                while(hashSet.contains(nums[i]+j)) {
                    j++;
                }

                if (j > ans) {
                    ans = j;
                }
            }
        }

        return ans;
    }

    public static void main(String[] args) {

        LongestConsecutiveSequence128 obj = new LongestConsecutiveSequence128();

        int[] nums = {100,4,200,1,3,2};

        System.out.println(obj.longestConsecutive(nums)); // Expected output 4
    }
}
