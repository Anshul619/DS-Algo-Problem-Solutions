package leetCodeProblems.PrefixSum;

/**
 * LeetCode - https://leetcode.com/problems/random-pick-with-weight/submissions/
 */

import java.util.Random;

public class RandomPickWithWeight528 {

    int[] prefixSums;
    int totalSum;

    public RandomPickWithWeight528(int[] w) {
        prefixSums = new int[w.length];

        int currentSum = 0;

        for(int i=0; i < w.length; i++) {
            currentSum += w[i];
            prefixSums[i] = currentSum;
        }

        totalSum = currentSum;
    }

    public int pickIndex() {

        int random = new Random().nextInt(totalSum);

        int low = 0;
        int high = prefixSums.length;

        while (low < high) {

            int mid = low + (high-low)/2;

            if (random == prefixSums[mid]) {
                return prefixSums[mid];
            }
            else if (random > prefixSums[mid]) {
                low = mid+1;
            }
            else {
                high = mid-1;
            }
        }

        return low;
    }
}
