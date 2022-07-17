package leetCodeProblems.BruteForce;

/**
 * LeetCode - https://leetcode.com/problems/find-n-unique-integers-sum-up-to-zero/submissions/
 *
 * TimeComplexity - O(n)
 * SpaceComplexity - O(n)
 */

import java.util.Arrays;

public class FindNUniqueIntegersSumUpToZero1304 {

    public int[] sumZero(int n) {

        int baseElement = 1;
        int currentIndex = 0;

        int remaining = n;

        int[] output = new int[n];

        while (currentIndex < n && remaining > 1) {

            output[currentIndex] = baseElement;
            currentIndex++;

            output[currentIndex] = -baseElement;
            currentIndex++;

            baseElement++;
            remaining -= 2;
        }

        return output;
    }

    public static void main(String[] args) {

        FindNUniqueIntegersSumUpToZero1304 obj = new FindNUniqueIntegersSumUpToZero1304();

        int[] output = obj.sumZero(5);
        System.out.println(Arrays.toString(output));

        int[] output1 = obj.sumZero(4);

        System.out.println(Arrays.toString(output1));
    }
}
