package leetCodeProblems.BruteForce;

/**
 * LeetCode - https://leetcode.com/problems/length-of-last-word/
 */

public class 58 {

    public int (String s) {

        String[] input = s.split(" ");
        return input[input.length-1].length();
    }

    public static void main(String[] args) {
        58 obj = new 58();
        System.out.println(obj.("   fly me   to   the moon  "));
    }
}
