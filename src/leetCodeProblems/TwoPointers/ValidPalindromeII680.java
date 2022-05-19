package leetCodeProblems.TwoPointers;

/**
 * LeetCode - https://leetcode.com/problems/valid-palindrome-ii/
 */
public class ValidPalindromeII680 {

    public boolean isPalindrome(String s, int leftPointer, int rightPointer) {

        while (leftPointer < rightPointer) {

            if (s.charAt(leftPointer) == s.charAt(rightPointer)) {

                leftPointer++;
                rightPointer--;

            } else {
                return false;
            }
        }

        return true;
    }

    public boolean validPalindrome(String s) {

        int leftPointer = 0;
        int rightPointer = s.length()-1;

        int numberOfDeletions = 0;

        while (leftPointer < rightPointer) {

            if (s.charAt(leftPointer) != s.charAt(rightPointer)) {

                if (isPalindrome(s, leftPointer+1, rightPointer) ||
                    isPalindrome(s, leftPointer, rightPointer-1)) {
                   return true;
                }
                else {
                    return false;
                }
            }

            leftPointer++;
            rightPointer--;
        }

        return true;
    }

    // Driver Code
    public static void main(String[] args)
    {
        ValidPalindromeII680 obj = new ValidPalindromeII680();

        //String str = "A man, a plan, a canal: Panama";

        //String str = "aba";
        //String str =  "abca";
        //String str = "abc";
        String str = "ebcbbececabbacecbbcbe"; // true
        //String str = "deee"; // true
        System.out.println(obj.validPalindrome(str));
    }
}
