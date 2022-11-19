package leetCodeProblems.GreedyAlgo;

/**
 * LeetCode - https://leetcode.com/problems/minimum-add-to-make-parentheses-valid/
 */

public class MinAddToMakeValidParentheses921 {

    public int minAddToMakeValid(String s) {
        int c_need = 0;
        int ans = 0;

        for (int i = 0; i < s.length(); i++) {

            if (s.charAt(i) == '(') {
                c_need += 1;
            }
            else {
                c_need -= 1;

                // This is TRICKY and IMPORTANT CODE. To reset the variables. This is the only closing bracket.
                if (c_need == -1) {
                    c_need = 0;
                    ans++; // an open bracket is needed here. Hence increment "ans".
                }
            }
        }

        return c_need + ans;
    }

    public static void main(String[] args) {

        MinAddToMakeValidParentheses921 obj = new MinAddToMakeValidParentheses921();

        System.out.println(obj.minAddToMakeValid("())"));
        System.out.println(obj.minAddToMakeValid("((("));
    }
}
