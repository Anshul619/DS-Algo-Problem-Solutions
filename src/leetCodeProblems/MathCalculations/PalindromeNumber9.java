package leetCodeProblems.MathCalculations;

import java.util.ArrayList;

/**
 * LeetCode - https://leetcode.com/problems/palindrome-number/
 *
 */
public class PalindromeNumber9 {

    String reversedNumber = "";

    public void calculateReversedNumber(int number) {

        int digit = number%10;
        int lowerNumber = number/10;

        //System.out.println(digit);
        reversedNumber = reversedNumber + String.valueOf(digit);

        if (lowerNumber == 0) {
            return;
        }

        calculateReversedNumber(lowerNumber);

        return;

    }

    public boolean isPalindrome(int x) {

        // Negative numbers can NOT palindrome
        if (x < 0) {
           return false;
        }
        reversedNumber = "";

        calculateReversedNumber(x);

        //System.out.println(reversedNumber);
        if (String.valueOf(x).equals(reversedNumber)) {
            return true;
        }

        return false;
    }

    public static void main(String[] args) {
        PalindromeNumber9 obj = new PalindromeNumber9();
        System.out.println(obj.isPalindrome(123)); // No
        System.out.println(obj.isPalindrome(121)); // Yes
        System.out.println(obj.isPalindrome(-1)); // No
    }
}
