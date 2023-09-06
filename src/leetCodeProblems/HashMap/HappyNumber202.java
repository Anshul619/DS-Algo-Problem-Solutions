package leetCodeProblems.MathCalculations;

/**
 * LeetCode - https://leetcode.com/problems/happy-number/submissions/
 *
 * TimeComplexity - O(n)
 * SpaceComplexity - O(n)
 */

import java.util.HashSet;

public class 202 {

    private int getDigitsSum(int n) {

        int leftNumber = n;
        int digit = 0;
        int sum = 0;

        while(leftNumber > 0) {
            digit = leftNumber%10;
            sum += digit*digit;
            leftNumber = leftNumber/10;
        }

        return sum;
    }

    public boolean isHappy(int n) {

        HashSet sumAlreadySeen = new HashSet();

        while (n!=1 && !sumAlreadySeen.contains(n)) {
            sumAlreadySeen.add(n);
            n = getDigitsSum(n);
        }

        if (n == 1) {
            return true;
        }

        return false;
    }

    public static void main(String[] args) {

        202 obj = new HappyNumber202();

        System.out.println(obj.isHappy(2));
    }
}
