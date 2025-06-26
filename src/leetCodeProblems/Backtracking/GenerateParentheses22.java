package leetCodeProblems.BackTracking;

/**
 * LeetCode - https://leetcode.com/problems/generate-parentheses/
 */

import java.util.ArrayList;
import java.util.List;

public class GenerateParentheses22 {

    public void generateParenthesisUtil(int n, int open, int close, String parenthesesString, List<String> output) {

        if (open == n && close == n) {
            output.add(parenthesesString);
            return;
        }

        if (open < n) {
            generateParenthesisUtil(n, open+1, close, parenthesesString + "(", output);
        }

        if (close < open) {
            generateParenthesisUtil(n, open, close+1, parenthesesString + ")", output);
        }

        return;
    }

    public List<String> generateParenthesis(int n) {

        List<String> output = new ArrayList<>();

        generateParenthesisUtil(n, 0, 0, "", output );

        return output;
    }

    public static void main(String[] args) {

        GenerateParentheses22 obj = new GenerateParentheses22();

        System.out.println(obj.generateParenthesis(4));

    }
}
