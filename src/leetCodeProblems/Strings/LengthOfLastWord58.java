package leetCodeProblems.Strings;

/**
 * LeetCode - https://leetcode.com/problems/length-of-last-word/
 */

public class LengthOfLastWord58 {

    public int lengthOfLastWord(String s) {

        String[] input = s.split(" ");
        return input[input.length-1].length();
    }

    public static void main(String[] args) {
        LengthOfLastWord58 obj = new LengthOfLastWord58();
        System.out.println(obj.lengthOfLastWord("   fly me   to   the moon  "));
    }
}
