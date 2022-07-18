package leetCodeProblems.Sorting;

/**
 * LeetCode - https://leetcode.com/problems/valid-anagram/solution/
 *
 * TimeComplexity - O(nlogn)
 * SpaceComplexity - O(n)
 */

import java.util.Arrays;

public class ValidAnagram242 {

    public boolean isAnagram(String s, String t) {

        if (s.length() != t.length()) {
            return false;
        }

        char[] sArr = s.toCharArray();
        Arrays.sort(sArr);
        s = String.valueOf(sArr);

        char[] tArr = t.toCharArray();
        Arrays.sort(tArr);
        t = String.valueOf(tArr);


        if (s.equals(t)) {
            return true;
        }

        return false;
    }

    public static void main(String[] args) {

        String s = "anagram";
        String t = "nagaram";

        //String s = "rat";
        //String t = "car";

        ValidAnagram242 obj = new ValidAnagram242();
        System.out.println(obj.isAnagram(s, t));
    }
}
